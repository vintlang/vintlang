package bundler

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	// "github.com/vintlang/vintlang/module"
)

// StringProcessor handles string-based processing of VintLang files
type StringProcessor struct {
	bundle          *FileBundle
	bundledPackages map[string]bool
}

// NewStringProcessor creates a new string processor
func NewStringProcessor(bundle *FileBundle) *StringProcessor {
	return &StringProcessor{
		bundle: bundle,
	}
}

// ProcessBundle processes the bundle using string manipulation
func (sp *StringProcessor) ProcessBundle() (string, error) {
	if sp.bundle == nil || len(sp.bundle.Files) == 0 {
		return "", fmt.Errorf("no files in bundle")
	}

	var result strings.Builder

	// Build the full set of bundled package names first so that
	// cross-imports between dependencies can be stripped.
	importedPackages := make(map[string]bool)
	includedFiles := make(map[string]bool)

	for filename := range sp.bundle.Files {
		if filename != sp.bundle.MainFile {
			if sp.bundle.IncludeFiles[filename] {
				includedFiles[filename] = true
			} else {
				packageName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
				importedPackages[packageName] = true
			}
		}
	}

	// Store the bundled package set on the processor so helper methods can use it
	sp.bundledPackages = importedPackages

	// Process all dependency files
	for filename, content := range sp.bundle.Files {
		if filename != sp.bundle.MainFile {
			if includedFiles[filename] {
				// This is an included file - embed directly
				processedContent := sp.processIncludedFile(content)
				result.WriteString(processedContent)
				result.WriteString("\n\n")
			} else {
				// This is an imported module - wrap in package
				packageName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))

				processedContent := sp.processDependencyFile(content, packageName)
				result.WriteString(processedContent)
				result.WriteString("\n\n")
			}
		}
	}

	// Then, process the main file
	mainContent, exists := sp.bundle.Files[sp.bundle.MainFile]
	if !exists {
		return "", fmt.Errorf("main file not found in bundle")
	}

	// Remove import statements for packages we've already included
	// and remove include statements for files we've already embedded
	processedMain := sp.processMainFile(mainContent, importedPackages, includedFiles)
	result.WriteString(processedMain)

	return result.String(), nil
}

// processIncludedFile processes an included file by embedding it directly
func (sp *StringProcessor) processIncludedFile(content string) string {
	content = strings.TrimSpace(content)

	// Remove any import statements from included files since they should be handled at the main level
	content = sp.removeImportStatements(content)

	// Remove any include statements from included files to avoid circular includes
	content = sp.removeIncludeStatements(content)

	return content
}

// processDependencyFile processes a dependency file to ensure it has proper package structure
func (sp *StringProcessor) processDependencyFile(content, packageName string) string {
	content = strings.TrimSpace(content)

	// Check if the content already has a package declaration
	packageRegex := regexp.MustCompile(`(?m)^package\s+` + regexp.QuoteMeta(packageName) + `\s*\{`)
	if packageRegex.MatchString(content) {
		// Already has package declaration — still need to strip imports for
		// bundled packages that are already inlined in the combined code
		content = sp.removeBundledImports(content)
		return content
	}

	// Remove any existing import statements from the dependency
	// since they should be handled at the main level
	content = sp.removeImportStatements(content)

	// If no package declaration, wrap the content in a package
	return fmt.Sprintf("package %s {\n%s\n}", packageName, content)
}

// processMainFile processes the main file to remove import statements for bundled packages
// and include statements for bundled files
func (sp *StringProcessor) processMainFile(content string, bundledPackages map[string]bool, includedFiles map[string]bool) string {
	lines := strings.Split(content, "\n")
	var result strings.Builder

	importRegex := regexp.MustCompile(`^\s*import\s+([\w, ]+)`)
	includeRegex := regexp.MustCompile(`^\s*include\s+"([^"]+)"`)

	for _, line := range lines {
		// Check if this line is an include statement
		includeMatches := includeRegex.FindStringSubmatch(line)
		if len(includeMatches) > 0 {
			// Always skip include statements since we've embedded the content
			continue
		}

		// Check if this line is an import statement
		importMatches := importRegex.FindStringSubmatch(line)
		if len(importMatches) > 1 {
			modules := strings.Split(importMatches[1], ",")
			var modulesToKeep []string

			for _, m := range modules {
				mod := strings.TrimSpace(m)
				// If this module is not a bundled package, keep it
				if !bundledPackages[mod] {
					modulesToKeep = append(modulesToKeep, mod)
				}
			}

			// If there are modules to keep, reconstruct the import statement
			if len(modulesToKeep) > 0 {
				result.WriteString(fmt.Sprintf("import %s\n", strings.Join(modulesToKeep, ", ")))
			} // else, skip the import statement entirely
			continue
		}

		// If it's not an include or import statement, write the line as is
		result.WriteString(line)
		result.WriteString("\n")
	}

	return result.String()
}

// removeIncludeStatements removes include statements from content
func (sp *StringProcessor) removeIncludeStatements(content string) string {
	lines := strings.Split(content, "\n")
	var result strings.Builder

	includeRegex := regexp.MustCompile(`^\s*include\s+`)

	for _, line := range lines {
		// Skip include lines
		if includeRegex.MatchString(line) {
			continue
		}

		result.WriteString(line)
		result.WriteString("\n")
	}

	return result.String()
}

// removeImportStatements removes import statements from content
func (sp *StringProcessor) removeImportStatements(content string) string {
	lines := strings.Split(content, "\n")
	var result strings.Builder

	importRegex := regexp.MustCompile(`^\s*import\s+`)

	for _, line := range lines {
		// Skip import lines
		if importRegex.MatchString(line) {
			continue
		}

		result.WriteString(line)
		result.WriteString("\n")
	}

	return result.String()
}

// removeBundledImports removes only import statements for packages that are
// already bundled (inlined), leaving imports for built-in modules intact.
func (sp *StringProcessor) removeBundledImports(content string) string {
	lines := strings.Split(content, "\n")
	var result strings.Builder

	importRegex := regexp.MustCompile(`^\s*import\s+([\w, ]+)`)

	for _, line := range lines {
		importMatches := importRegex.FindStringSubmatch(line)
		if len(importMatches) > 1 {
			modules := strings.Split(importMatches[1], ",")
			var modulesToKeep []string

			for _, m := range modules {
				mod := strings.TrimSpace(m)
				if !sp.bundledPackages[mod] {
					modulesToKeep = append(modulesToKeep, mod)
				}
			}

			if len(modulesToKeep) > 0 {
				fmt.Fprintf(&result, "import %s\n", strings.Join(modulesToKeep, ", "))
			}
			continue
		}

		result.WriteString(line)
		result.WriteString("\n")
	}

	return result.String()
}
