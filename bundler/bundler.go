package bundler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/vintlang/vintlang/config"
	"github.com/vintlang/vintlang/utils"
)

// logError logs errors to a file for debugging
func logError(err error) {
	f, ferr := os.OpenFile("bundler.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", ferr)
		return
	}
	defer f.Close()
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Fprintf(f, "[%s] %v\n", timestamp, err)
}

// printVerbose prints only if verbose mode is enabled
func printVerbose(verbose bool, a ...any) {
	if verbose {
		fmt.Print(a...)
	}
}
func printlnVerbose(verbose bool, a ...any) {
	if verbose {
		fmt.Println(a...)

	}
}

// Bundles the vintlang code to a go binary
func Bundle(args []string) error {
	if len(args) == 0 {
		err := fmt.Errorf("no input file provided to bundler")
		logError(err)
		return err
	}
	vintFile := args[0]
	fmt.Println(len(vintFile))

	// Verbose/Quiet mode
	verbose := true
	if len(args) >= 7 && args[6] == "quiet" {
		verbose = false
	}

	keepTemp := len(args) >= 8 && args[7] == "keep"

	printlnVerbose(verbose, ">> Starting Enhanced Bundle for '", filepath.Base(vintFile), "'")

	// Analyze dependencies
	printVerbose(verbose, "=> Analyzing dependencies... ")
	analyzer := NewDependencyAnalyzer()
	bundle, err := analyzer.AnalyzeDependencies(vintFile)
	if err != nil {
		err = fmt.Errorf("failed to analyze dependencies: %w", err)
		logError(err)
		return err
	}
	printlnVerbose(verbose, fmt.Sprintf("=> Found %d files", len(bundle.Files)))

	printVerbose(verbose, "=> Creating temp Bundle directory... ")
	tempDir, err := os.MkdirTemp("", "vint-bundle-*")
	if err != nil {
		err = fmt.Errorf("failed to create temp dir: %w", err)
		logError(err)
		return err
	}
	if !keepTemp {
		defer os.RemoveAll(tempDir)
	}
	printlnVerbose(verbose, "OK")

	bundlerVersion := config.VINT_VERSION
	buildTime := time.Now().Format(time.RFC3339)

	// Generate bundled code using the new evaluator
	printVerbose(verbose, "=> Generating Go code with bundled files... ")
	bundledEvaluator := NewBundledEvaluator(bundle)
	goCode, err := bundledEvaluator.GenerateBundledCode(bundlerVersion, buildTime)
	if err != nil {
		err = fmt.Errorf("failed to generate bundled code: %w", err)
		logError(err)
		return err
	}

	mainPath := filepath.Join(tempDir, "main.go")
	if err := os.WriteFile(mainPath, []byte(goCode), 0644); err != nil {
		err = fmt.Errorf("failed to create main.go in temp dir '%s': %w", tempDir, err)
		logError(err)
		return err
	}
	printlnVerbose(verbose, "OK")

	printVerbose(verbose, "=> Initializing modules... ")

	// Use the current vintlang version for the dependency
	vintVersion := config.VINT_VERSION
	if !strings.HasPrefix(vintVersion, "v") {
		vintVersion = "v" + vintVersion
	}

	goMod := fmt.Sprintf(`module vint-bundled

go 1.25

require github.com/vintlang/vintlang %s
`, vintVersion)

	if err := os.WriteFile(filepath.Join(tempDir, "go.mod"), []byte(goMod), 0644); err != nil {
		err = fmt.Errorf("failed to create go.mod in temp dir '%s': %w", tempDir, err)
		logError(err)
		return err
	}
	printlnVerbose(verbose, "OK")

	// Bundle binary
	binaryName := strings.TrimSuffix(filepath.Base(vintFile), ".vint")

	printlnVerbose(verbose, args)
	if len(args) >= 3 && args[2] != "" {
		binaryName = args[2]
	}
	printlnVerbose(verbose, "=> Bundling binary '", binaryName, "'...")

	spinner := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	done := make(chan bool)
	if verbose {
		go func() {
			for i := 0; ; i++ {
				select {
				case <-done:
					return
				default:
					fmt.Printf("\r%s Bundling...", spinner[i%len(spinner)])
					time.Sleep(100 * time.Millisecond)
				}
			}
		}()
	}

	// Cross-compilation: GOOS and GOARCH
	goos := ""
	goarch := ""
	if len(args) >= 5 && args[4] != "" {
		goos = args[4]
	}
	if len(args) >= 6 && args[5] != "" {
		goarch = args[5]
	}

	// Build command with cross-compilation support
	// Skip 'go mod tidy' — go build with -mod=mod resolves modules directly,
	// avoiding a separate slow network + verification pass.
	buildEnv := "CGO_ENABLED=0 GONOSUMDB=* GOFLAGS=-mod=mod "
	if goos != "" {
		buildEnv += fmt.Sprintf("GOOS=%s ", goos)
	}
	if goarch != "" {
		buildEnv += fmt.Sprintf("GOARCH=%s ", goarch)
	}
	BundleCmd := fmt.Sprintf("cd %s && %sgo build -trimpath -o %s", tempDir, buildEnv, binaryName)
	if err := utils.RunShell(BundleCmd); err != nil {
		err = fmt.Errorf("bundle command failed: %w", err)
		logError(err)
		done <- true
		return fmt.Errorf("\n!! Bundle failed: %w", err)
	}

	done <- true
	printVerbose(verbose, "\r=> Bundle successful! Moving binary... ")

	// Determine output directory
	outputDir := "."
	if len(args) >= 4 && args[3] != "" {
		outputDir = args[3]
		if _, err := os.Stat(outputDir); os.IsNotExist(err) {
			err = fmt.Errorf("output directory '%s' does not exist", outputDir)
			logError(err)
			return err
		}
	}
	finalBinary := filepath.Join(tempDir, binaryName)
	outputPath := filepath.Join(outputDir, binaryName)
	if err := os.Rename(finalBinary, outputPath); err != nil {
		err = fmt.Errorf("failed to move binary from '%s' to '%s': %w", finalBinary, outputPath, err)
		logError(err)
		return fmt.Errorf("\n!! Failed to move binary: %w", err)
	}

	printlnVerbose(verbose, "OK")
	fmt.Printf("\n=> Successfully created binary with %d bundled files: %s\n", len(bundle.Files), outputPath)

	if keepTemp {
		printlnVerbose(verbose, "Temp directory kept for debugging:", tempDir)
	}
	return nil
}

// target represents a single GOOS/GOARCH build target with an output name
type target struct {
	goos, goarch, name string
}

// BundleMulti bundles a vint file into multiple binaries for different platforms
// in a single invocation: deps + codegen run once, then go build runs per target.
//
// args: <file> <placeholder> <name_prefix> <output_dir> <targets> [quiet]
//
//	targets = "darwin/arm64:name,linux/amd64:name,..." or "darwin/arm64,linux/amd64,..."
//	If :name is omitted, the binary name is <name_prefix>-<goos>-<goarch>[.exe].
func BundleMulti(args []string) error {
	if len(args) < 5 {
		return fmt.Errorf("usage: vint bundle-multi <file> \"\" <name_prefix> <output_dir> <targets> [quiet]")
	}

	vintFile := args[0]
	namePrefix := args[2]
	outputDir := args[3]
	targetsStr := args[4]

	verbose := true
	if len(args) >= 6 && args[5] == "quiet" {
		verbose = false
	}

	// Parse targets
	targets, err := parseTargets(targetsStr, namePrefix)
	if err != nil {
		return err
	}

	printlnVerbose(verbose, ">> Multi-target bundle for '"+filepath.Base(vintFile)+"' →", len(targets), "targets")

	// ── Phase 1: Analyze dependencies (once) ─────────────────────────────
	printVerbose(verbose, "=> Analyzing dependencies... ")
	analyzer := NewDependencyAnalyzer()
	bundle, err := analyzer.AnalyzeDependencies(vintFile)
	if err != nil {
		return fmt.Errorf("failed to analyze dependencies: %w", err)
	}
	printlnVerbose(verbose, fmt.Sprintf("Found %d files", len(bundle.Files)))

	// ── Phase 2: Generate Go code (once) ─────────────────────────────────
	printVerbose(verbose, "=> Generating Go code... ")
	bundlerVersion := config.VINT_VERSION
	buildTime := time.Now().Format(time.RFC3339)
	bundledEvaluator := NewBundledEvaluator(bundle)
	goCode, err := bundledEvaluator.GenerateBundledCode(bundlerVersion, buildTime)
	if err != nil {
		return fmt.Errorf("failed to generate bundled code: %w", err)
	}
	printlnVerbose(verbose, "OK")

	// ── Phase 3: Create temp dir + go.mod (once) ─────────────────────────
	tempDir, err := os.MkdirTemp("", "vint-bundle-multi-*")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	if err := os.WriteFile(filepath.Join(tempDir, "main.go"), []byte(goCode), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %w", err)
	}

	vintVersion := config.VINT_VERSION
	if !strings.HasPrefix(vintVersion, "v") {
		vintVersion = "v" + vintVersion
	}
	goMod := fmt.Sprintf("module vint-bundled\n\ngo 1.25\n\nrequire github.com/vintlang/vintlang %s\n", vintVersion)
	if err := os.WriteFile(filepath.Join(tempDir, "go.mod"), []byte(goMod), 0644); err != nil {
		return fmt.Errorf("failed to write go.mod: %w", err)
	}

	// Ensure output dir exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		return fmt.Errorf("output directory '%s' does not exist", outputDir)
	}

	// Resolve modules once before parallel builds to avoid go.mod races.
	// go mod tidy resolves deps for all platforms (not just the host), populating go.sum fully.
	printVerbose(verbose, "=> Resolving modules... ")
	resolveCmd := fmt.Sprintf("cd %s && GONOSUMDB=* go mod tidy", tempDir)
	if err := utils.RunShell(resolveCmd); err != nil {
		return fmt.Errorf("module resolution failed: %w", err)
	}
	printlnVerbose(verbose, "OK")

	// ── Phase 4: Build all targets concurrently ──────────────────────────
	printlnVerbose(verbose, "=> Building", len(targets), "targets...")

	type buildResult struct {
		t   target
		err error
	}
	results := make([]buildResult, len(targets))
	var wg sync.WaitGroup

	for i, t := range targets {
		wg.Add(1)
		go func(idx int, tgt target) {
			defer wg.Done()
			buildEnv := fmt.Sprintf("CGO_ENABLED=0 GONOSUMDB=* GOFLAGS=-mod=readonly GOOS=%s GOARCH=%s ", tgt.goos, tgt.goarch)
			cmd := fmt.Sprintf("cd %s && %sgo build -trimpath -o %s", tempDir, buildEnv, tgt.name)
			bErr := utils.RunShell(cmd)
			if bErr == nil {
				// Move binary to output dir
				src := filepath.Join(tempDir, tgt.name)
				dst := filepath.Join(outputDir, tgt.name)
				bErr = os.Rename(src, dst)
			}
			results[idx] = buildResult{t: tgt, err: bErr}
		}(i, t)
	}
	wg.Wait()

	// Report results
	failed := 0
	for _, r := range results {
		if r.err != nil {
			fmt.Printf("  ✗ %s/%s (%s): %v\n", r.t.goos, r.t.goarch, r.t.name, r.err)
			failed++
		} else {
			printlnVerbose(verbose, fmt.Sprintf("  ✓ %s/%s → %s/%s", r.t.goos, r.t.goarch, outputDir, r.t.name))
		}
	}

	if failed > 0 {
		return fmt.Errorf("%d of %d targets failed", failed, len(targets))
	}
	fmt.Printf("\n=> Successfully built %d binaries (%d files each) in %s/\n", len(targets), len(bundle.Files), outputDir)
	return nil
}

// parseTargets parses "darwin/arm64:name,linux/amd64,..." into target structs.
// If :name is omitted, generates <prefix>-<goos>-<goarch>[.exe].
func parseTargets(s, prefix string) ([]target, error) {
	parts := strings.Split(s, ",")
	var targets []target
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		var name, osArch string
		if idx := strings.Index(p, ":"); idx != -1 {
			osArch = p[:idx]
			name = p[idx+1:]
		} else {
			osArch = p
		}
		slash := strings.SplitN(osArch, "/", 2)
		if len(slash) != 2 || slash[0] == "" || slash[1] == "" {
			return nil, fmt.Errorf("invalid target '%s': expected os/arch", p)
		}
		goos, goarch := slash[0], slash[1]
		if name == "" {
			name = prefix + "-" + goos + "-" + goarch
			if goos == "windows" {
				name += ".exe"
			}
		}
		targets = append(targets, target{goos: goos, goarch: goarch, name: name})
	}
	if len(targets) == 0 {
		return nil, fmt.Errorf("no targets specified")
	}
	return targets, nil
}
