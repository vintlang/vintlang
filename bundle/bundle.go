package bundle

// bundledFiles stores files that are bundled into the binary
var bundledFiles map[string]string

// SetBundledFiles sets the bundled files for import resolution
func SetBundledFiles(files map[string]string) {
	bundledFiles = files
}

// GetBundledFiles returns the bundled files
func GetBundledFiles() map[string]string {
	return bundledFiles
}