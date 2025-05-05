package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// RunShell executes a shell command and returns any error
func RunShell(cmd string) error {
	shell := []string{"sh", "-c", cmd}
	if os.Getenv("OS") == "Windows_NT" {
		shell = []string{"cmd", "/C", cmd}
	}
	c := exec.Command(shell[0], shell[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

// FileExists checks if a file exists at the given path
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDirectory checks if the given path is a directory
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// GetFileExtension returns the extension of a file (without the dot)
func GetFileExtension(path string) string {
	return strings.TrimPrefix(filepath.Ext(path), ".")
}

// GetCurrentDirectory returns the current working directory
func GetCurrentDirectory() (string, error) {
	return os.Getwd()
}

// CreateDirectory creates a directory at the given path
func CreateDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

// GetFileModTime returns the last modification time of a file
func GetFileModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// GetFileSize returns the size of a file in bytes
func GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// IsExecutable checks if a file is executable
func IsExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode()&0111 != 0
}

// GetHomeDirectory returns the user's home directory
func GetHomeDirectory() (string, error) {
	return os.UserHomeDir()
}

// GetTempDirectory returns the system's temporary directory
func GetTempDirectory() string {
	return os.TempDir()
}

// GetEnvironmentVariable returns the value of an environment variable
func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

// SetEnvironmentVariable sets an environment variable
func SetEnvironmentVariable(key, value string) error {
	return os.Setenv(key, value)
}

// GetExecutablePath returns the path of the current executable
func GetExecutablePath() (string, error) {
	return os.Executable()
}

// GetHostname returns the system's hostname
func GetHostname() (string, error) {
	return os.Hostname()
}

// GetUsername returns the current user's username
func GetUsername() (string, error) {
	return os.UserHomeDir()
}

// GetProcessID returns the current process ID
func GetProcessID() int {
	return os.Getpid()
}

// GetParentProcessID returns the parent process ID
func GetParentProcessID() int {
	return os.Getppid()
}

// IsRoot checks if the current process is running as root/admin
func IsRoot() bool {
	return os.Geteuid() == 0
}

// GetSystemInfo returns basic system information
func GetSystemInfo() map[string]string {
	info := make(map[string]string)

	hostname, _ := os.Hostname()
	info["hostname"] = hostname

	username, _ := os.UserHomeDir()
	info["username"] = username

	info["pid"] = strconv.Itoa(os.Getpid())
	info["ppid"] = strconv.Itoa(os.Getppid())

	info["temp_dir"] = os.TempDir()

	return info
}
