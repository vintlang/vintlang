package bundler

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

// StringProcessor handles string-based processing of VintLang files
type StringProcessor struct {
	bundle *FileBundle
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

	// First, process all dependency files to extract package definitions
	importedPackages := make(map[string]bool)
	
	for filename, content := range sp.bundle.Files {
		if filename != sp.bundle.MainFile {
			packageName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
			importedPackages[packageName] = true
			
			// Process the dependency file and add it to the result
			processedContent := sp.processDependencyFile(content, packageName)
			result.WriteString(processedContent)
			result.WriteString("\n\n")
		}
	}

	// Then, process the main file
	mainContent, exists := sp.bundle.Files[sp.bundle.MainFile]
	if !exists {
		return "", fmt.Errorf("main file not found in bundle")
	}

	// Remove import statements for packages we've already included
	processedMain := sp.processMainFile(mainContent, importedPackages)
	result.WriteString(processedMain)

	return result.String(), nil
}

// processDependencyFile processes a dependency file to ensure it has proper package structure
func (sp *StringProcessor) processDependencyFile(content, packageName string) string {
	content = strings.TrimSpace(content)
	
	// Check if the content already has a package declaration
	packageRegex := regexp.MustCompile(`(?m)^package\s+` + regexp.QuoteMeta(packageName) + `\s*\{`)
	if packageRegex.MatchString(content) {
		// Already has package declaration, return as-is
		return content
	}
	
	// Remove any existing import statements from the dependency
	// since they should be handled at the main level
	content = sp.removeImportStatements(content)
	
	// If no package declaration, wrap the content in a package
	return fmt.Sprintf("package %s {\n%s\n}", packageName, content)
}

// processMainFile processes the main file to remove import statements for bundled packages
func (sp *StringProcessor) processMainFile(content string, bundledPackages map[string]bool) string {
	lines := strings.Split(content, "\n")
	var result strings.Builder
	
	importRegex := regexp.MustCompile(`^\s*import\s+(\w+)`)
	
	for _, line := range lines {
		// Check if this line is an import statement
		matches := importRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			importedModule := matches[1]
			// If this module is one of our bundled packages, skip the import
			if bundledPackages[importedModule] {
				continue
			}
		}
		
		// Add the line to the result
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