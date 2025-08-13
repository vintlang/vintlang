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

	// First, process all dependency files
	importedPackages := make(map[string]bool)
	includedFiles := make(map[string]bool)
	
	for filename, content := range sp.bundle.Files {
		if filename != sp.bundle.MainFile {
			if sp.bundle.IncludeFiles[filename] {
				// This is an included file - embed directly
				processedContent := sp.processIncludedFile(content)
				result.WriteString(processedContent)
				result.WriteString("\n\n")
				includedFiles[filename] = true
			} else {
				// This is an imported module - wrap in package
				packageName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
				importedPackages[packageName] = true
				
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
// and include statements for bundled files
func (sp *StringProcessor) processMainFile(content string, bundledPackages map[string]bool, includedFiles map[string]bool) string {
	lines := strings.Split(content, "\n")
	var result strings.Builder
	
	importRegex := regexp.MustCompile(`^\s*import\s+(\w+)`)
	includeRegex := regexp.MustCompile(`^\s*include\s+"([^"]+)"`)
	
	for _, line := range lines {
		skipLine := false
		
		// Check if this line is an import statement
		importMatches := importRegex.FindStringSubmatch(line)
		if len(importMatches) > 1 {
			importedModule := importMatches[1]
			// If this module is one of our bundled packages, skip the import
			if bundledPackages[importedModule] {
				skipLine = true
			}
		}
		
		// Check if this line is an include statement
		includeMatches := includeRegex.FindStringSubmatch(line)
		if len(includeMatches) > 1 {
			// Always skip include statements since we've embedded the content
			skipLine = true
		}
		
		if !skipLine {
			result.WriteString(line)
			result.WriteString("\n")
		}
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