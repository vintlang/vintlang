package toolkit

import (
	"archive/tar"
	// "archive/zip"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/vintlang/vintlang/internal/config"
	// "strings"
)

var CLI_ARGS []string = []string{}

func GetCliArgs() []string { // Returns the CLI_ARGS
	return CLI_ARGS
}

type ReleaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type Release struct {
	Assets []ReleaseAsset `json:"assets"`
}

func detectPlatform() string {
	switch runtime.GOOS {
	case "linux":
		if _, err := os.Stat("/system/build.prop"); err == nil {
			return "android"
		}
		return "linux"
	case "darwin":
		return "macos"
	case "windows":
		return "windows"
	default:
		return "unsupported"
	}
}

func getBinaryName(platform string) string {
	switch platform {
	case "linux":
		return "vintpm_linux_amd64.tar.gz"
	case "macos":
		return "vintpm_macos_amd64.tar.gz"
	case "android":
		return "vintpm_android_arm64.tar.gz"
	case "windows":
		return "vintpm_windows_amd64.zip"
	default:
		return ""
	}
}

func fetchLatestReleaseURL(binaryName string) (string, error) {
	resp, err := http.Get("https://api.github.com/repos/ekilie/vintpm/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", err
	}

	for _, asset := range release.Assets {
		if asset.Name == binaryName {
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("no suitable binary found for platform")
}

func downloadFile(url, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func installBinary(binaryName, platform string) error {
	if platform == "windows" {
		cmd := exec.Command("unzip", "-o", binaryName, "-d", "C:/usr/local/bin")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	file, err := os.Open(binaryName)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if header.Typeflag == tar.TypeReg {
			destFile := "/usr/local/bin/" + header.Name
			outFile, err := os.Create(destFile)
			if err != nil {
				return err
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}

			if err := os.Chmod(destFile, 0755); err != nil {
				return err
			}
		}
	}

	return nil
}

func InstallVintpm() {
	platform := detectPlatform()
	if platform == "unsupported" {
		fmt.Println("Unsupported platform. Exiting.")
		return
	}

	binaryName := getBinaryName(platform)
	if binaryName == "" {
		fmt.Println("No binary name mapping found for platform. Exiting.")
		return
	}

	fmt.Println("Fetching the latest release information...")
	assetURL, err := fetchLatestReleaseURL(binaryName)
	if err != nil {
		fmt.Printf("❌ Error fetching release: %v\n", err)
		return
	}

	fmt.Println("Downloading the latest release...")
	if err := downloadFile(assetURL, binaryName); err != nil {
		fmt.Printf("❌ Error downloading binary: %v\n", err)
		return
	}

	fmt.Println("Installing vintpm...")
	if err := installBinary(binaryName, platform); err != nil {
		fmt.Printf("❌ Error installing binary: %v\n", err)
		return
	}

	fmt.Println("Cleaning up...")
	if err := os.Remove(binaryName); err != nil {
		fmt.Printf("❌ Error cleaning up: %v\n", err)
	}

	fmt.Println("Installation complete!")
}

func Get(pkg string) {
	switch pkg {
	case "vintpm":
		InstallVintpm()
	}
}

type VintConfig struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	VintVersion string `json:"vint,omitempty"`
	Description string `json:"description"`
}

const sampleReadme = `# VintLang Starter

A modern VintLang project using typed declarations, :: builtin prefix, and a custom package.

## Run

` + "```" + `bash
vint main.vint
` + "```" + `

## Structure

- ` + "`main.vint`" + ` — entry point with ` + "`main()`" + ` function
- ` + "`vintconfig.json`" + ` — project metadata
`

const sampleReadmeMinimal = `# VintLang Minimal Starter

A bare-bones VintLang project.

## Run

` + "```" + `bash
vint main.vint
` + "```" + `
`

const sampleVintCode = `// VintLang starter project
import greetings_module, time

let main = func() {
    ::println("Hello from VintLang!");
    ::println("Current time:", time.now());

    greetings_module.greet("World");
    ::println("Module info:", greetings_module.getInfo());

    ::println(greetings_module.process("sample data"));

    ::println("Greeted count:", greetings_module.getInfo()["counter"]);
};
`

const sampleVintCodeMinimal = `// VintLang minimal starter
import cli

let main = func() {
    ::println("Hello from VintLang!");
};
`

const sampleGreetingsCode = `package greetings_module {
    const VERSION: string = "2.0.0";
    const AUTHOR: string = "VintLang Team";
    const DESCRIPTION: string = "A demo package with all features";

    let greeting: string = "Hello";
    let publicCounter: int = 0;

    let _privateSecret: string = "shhh";
    const _PRIVATE_KEY: string = "abc123";
    let _internalCounter: int = 0;

    let init = func() {
        @.greeting = "Welcome";
        @.publicCounter = 1;
        ::print("greetings_module v" + VERSION + " initialized!");
    };

    let greet = func(name: string) {
        ::print(greeting + ", " + name + "!");
        publicCounter = publicCounter + 1;
    };

    let getInfo = func() {
        return {
            "version": VERSION,
            "author": AUTHOR,
            "desc": DESCRIPTION,
            "counter": publicCounter
        };
    };

    let _logDebug = func(msg: string) {
        ::print("[DEBUG] " + msg);
    };

    let process = func(data: string): string {
        _logDebug("Processing: " + data);
        return "Processed: " + data;
    };
};
`

// createProject scaffolds a new Vint project in the given directory.
// If minimal is true, only main.vint and vintconfig.json are created.
func createProject(projectName string, minimal bool) {
	var vintConfig = VintConfig{
		Name:        projectName,
		Version:     "1.0.0",
		VintVersion: config.VINT_VERSION,
		Description: "A modern VintLang starter project",
	}

	os.Mkdir(projectName, 0755)
	os.Chdir(projectName)

	// README.md
	readmeContent := sampleReadme
	if minimal {
		readmeContent = sampleReadmeMinimal
	}
	fmt.Println("Creating README.md...")
	readmeFile, err := os.Create("README.md")
	if err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		return
	}
	defer readmeFile.Close()
	if _, err := readmeFile.WriteString(readmeContent); err != nil {
		fmt.Printf("Error writing to README.md: %v\n", err)
		return
	}
	fmt.Println("README.md created")

	// vintconfig.json
	if minimal {
		vintConfig.Description = "A minimal VintLang project"
	}
	fmt.Println("Creating vintconfig.json...")
	vintFile, err := os.Create("vintconfig.json")
	if err != nil {
		fmt.Printf("Error creating vintconfig.json: %v\n", err)
		return
	}
	defer vintFile.Close()
	vintData, err := json.MarshalIndent(vintConfig, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling vintconfig.json: %v\n", err)
		return
	}
	if _, err := vintFile.Write(vintData); err != nil {
		fmt.Printf("Error writing to vintconfig.json: %v\n", err)
		return
	}
	fmt.Println("vintconfig.json created")

	// main.vint
	mainContent := sampleVintCode
	if minimal {
		mainContent = sampleVintCodeMinimal
	}
	fmt.Println("Creating main.vint...")
	mainFile, err := os.Create("main.vint")
	if err != nil {
		fmt.Printf("Error creating main.vint: %v\n", err)
		return
	}
	defer mainFile.Close()
	if _, err := mainFile.WriteString(mainContent); err != nil {
		fmt.Printf("Error writing to main.vint: %v\n", err)
		return
	}
	fmt.Println("main.vint created")

	// greetings_module.vint (only in full mode)
	if !minimal {
		fmt.Println("Creating greetings_module.vint...")
		greetingsModuleFile, err := os.Create("greetings_module.vint")
		if err != nil {
			fmt.Printf("Error creating greetings_module.vint: %v\n", err)
			return
		}
		defer greetingsModuleFile.Close()
		if _, err := greetingsModuleFile.WriteString(sampleGreetingsCode); err != nil {
			fmt.Printf("Error writing to greetings_module.vint: %v\n", err)
		}
		fmt.Println("greetings_module.vint created")
	}

	fmt.Printf("Project '%s' initialized successfully!\n", projectName)
}

// Init is the project initializer for 'vint init'.
// Accepts --minimal for a bare-bones scaffold.
func Init(args []string) {
	projectName, minimal := parseInitArgs(args)
	createProject(projectName, minimal)
}

// New is the project initializer for 'vint new'.
// Accepts --minimal for a bare-bones scaffold.
func New(args []string) {
	projectName, minimal := parseInitArgs(args)
	createProject(projectName, minimal)
}

// parseInitArgs extracts the project name and --minimal flag from CLI args.
// args[0] is the program name, args[1] is the command (init/new),
// subsequent args may be the project name and/or flags.
func parseInitArgs(args []string) (string, bool) {
	projectName := "vint-project"
	minimal := false

	for _, arg := range args[2:] {
		if arg == "--minimal" || arg == "-m" {
			minimal = true
		} else if arg != "" && arg[0] != '-' {
			projectName = arg
		}
	}

	return projectName, minimal
}
