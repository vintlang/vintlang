#!/usr/bin/env node

/**
 * Script to regenerate docs registry and validate all markdown files
 * Usage: node scripts/update-docs.js
 */

import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Import the generated functions (this is a Node.js script)
async function runDocUpdate() {
  try {
    const docsDir = path.join(__dirname, '..', '..', 'docs');
    const files = fs.readdirSync(docsDir);
    const markdownFiles = files.filter(file => file.endsWith('.md'));
    
    console.log(`Found ${markdownFiles.length} markdown files in docs/`);
    console.log('Files:', markdownFiles.map(f => `  - ${f}`).join('\n'));
    
    // Validate each file can be read
    let validFiles = 0;
    let errorFiles = [];
    
    for (const file of markdownFiles) {
      try {
        const filePath = path.join(docsDir, file);
        const content = fs.readFileSync(filePath, 'utf-8');
        
        if (content.length > 0) {
          validFiles++;
        } else {
          errorFiles.push(`${file}: Empty file`);
        }
      } catch (error) {
        errorFiles.push(`${file}: ${error.message}`);
      }
    }
    
    console.log(`\n✅ ${validFiles} valid markdown files`);
    
    if (errorFiles.length > 0) {
      console.log(`❌ ${errorFiles.length} files with issues:`);
      errorFiles.forEach(error => console.log(`  - ${error}`));
    }
    
    console.log('\n🚀 Docs registry will be automatically generated when the website loads.');
    console.log('All markdown files will be available at /docs/learn/{filename}');
    
  } catch (error) {
    console.error('Error updating docs:', error);
    process.exit(1);
  }
}

// Run the script if called directly
runDocUpdate();