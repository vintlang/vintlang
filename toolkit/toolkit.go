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
	Description string `json:"description"`
}

const sampleVintCode = `// Simple string manipulation and message printing
import greetings_module
import time

// Print a greeting
print("Hello, VintLang World! it currently",time.now())

// Demonstrate string splitting
let phrase = "VintLang"
let letters = phrase.split("")
for letter in letters {
    print(letter)
}

//from the greetings_module
greetings_module.greet("Developer")`

const sampleGreetingsCode = `
package greetings_module{
	// Demonstrate a simple function from a package
	let greet = func(name) {
		print("Hello, " + name + "!")
	}
}
`

func Init(args []string) {
	projectName := "vint-project"
	if len(args) >= 2 {
		projectName = args[2]
	}

	// Structure for vintconfig.json
	var vintConfig = VintConfig{
		Name:        projectName,
		Version:     "1.0.0",
		Description: "I love VintLang",
	}

	// creating the project directory
	os.Mkdir(projectName, 0755)
	os.Chdir(projectName)

	// Creating vintconfig.json
	fmt.Println("🫠 Creating vintconfig.json...")
	vintFile, err := os.Create("vintconfig.json")
	if err != nil {
		fmt.Printf("❌ Error creating vintconfig.json: %v\n", err)
		return
	}
	defer vintFile.Close()

	vintData, err := json.MarshalIndent(vintConfig, "", "  ")
	if err != nil {
		fmt.Printf("❌ Error marshalling vintconfig.json: %v\n", err)
		return
	}
	if _, err := vintFile.Write(vintData); err != nil {
		fmt.Printf("❌ Error writing to vintconfig.json: %v\n", err)
		return
	}
	fmt.Println("✅ vintconfig.json created successfully!")

	// Creating main.vint
	fmt.Println("🫠 Creating main.vint...")
	mainFile, err := os.Create("main.vint")
	if err != nil {
		fmt.Printf("❌ Error creating main.vint: %v\n", err)
		return
	}
	defer mainFile.Close()

	if _, err := mainFile.WriteString(sampleVintCode); err != nil {
		fmt.Printf("❌ Error writing to main.vint: %v\n", err)
		return
	}
	fmt.Println("✅ main.vint created successfully!")

	//creating greetings_module
	greetings_module_file, err := os.Create("greetings_module.vint")
	if err != nil {
		fmt.Printf("❌ Error creating greetings_module.vint: %v\n", err)
		return
	}
	defer greetings_module_file.Close()
	if _, err := greetings_module_file.WriteString(sampleGreetingsCode); err != nil {
		fmt.Printf("❌ Error writing to greetings_module.vint: %v\n", err)
	}
	fmt.Println("✅ greetings_module.vint creates succesfully")

	// Success message
	fmt.Printf("🚀 Project '%s' initialized successfully!\n", projectName)
}
