import fs from "fs";
import path from "path";

/**
 * Build-time script to generate static docs registry from markdown files
 * This runs during build to create a static JSON file that can be imported
 */

/**
 * Extracts the title and description from a markdown file
 * @param {string} filename - The markdown filename
 * @param {string} content - The file content
 * @returns {Object} DocItem object with title, href, description, filename
 */
function extractDocInfo(filename, content) {
  const lines = content.split("\n");

  // Find the first h1 header (# Title)
  let title = filename.replace(".md", "").replace(/_/g, " ");
  let description = "";

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();

    // Extract title from first h1 header
    if (line.startsWith("# ") && !title) {
      title = line.substring(2).trim();
      continue;
    }

    // Extract description from first paragraph after title
    if (
      title &&
      line &&
      !line.startsWith("#") &&
      !line.startsWith("```") &&
      !line.startsWith("<!--") &&
      line.length > 20
    ) {
      description = line.length > 150 ? line.substring(0, 150) + "..." : line;
      break;
    }
  }

  // Fallback description if none found
  if (!description) {
    if (title.includes("HTTP")) {
      description = `Learn about ${title.toLowerCase()} functionality in VintLang.`;
    } else if (title.toLowerCase().includes("function")) {
      description = `Learn about functions and function handling in VintLang.`;
    } else if (title.toLowerCase().includes("array")) {
      description = `Learn about arrays and array manipulation in VintLang.`;
    } else if (title.toLowerCase().includes("string")) {
      description = `Learn about string manipulation and functions in VintLang.`;
    } else {
      description = `Learn about ${title.toLowerCase()} in VintLang.`;
    }
  }

  return {
    title: formatTitle(title),
    href: `/docs/learn/${filename.replace(".md", "")}`,
    description,
    filename: filename.replace(".md", ""),
  };
}

/**
 * Formats the title to be more readable
 * @param {string} title - The title to format
 * @returns {string} The formatted title
 */
function formatTitle(title) {
  // Remove "in Vint", "in VintLang", etc. from titles
  title = title.replace(/\s+(in\s+)?(vint|vintlang)(\s+.*)?$/i, "");

  // Capitalize each word properly
  return title
    .split(" ")
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
    .join(" ")
    .replace(/\bHttp\b/g, "HTTP")
    .replace(/\bApi\b/g, "API")
    .replace(/\bSql\b/g, "SQL")
    .replace(/\bJson\b/g, "JSON")
    .replace(/\bXml\b/g, "XML")
    .replace(/\bUuid\b/g, "UUID")
    .replace(/\bCsv\b/g, "CSV")
    .replace(/\bOs\b/g, "OS")
    .replace(/\bIf\b/g, "If")
    .replace(/\bFor\b/g, "For");
}

/**
 * Generates categories for better organization
 * @param {Array} items - Array of doc items to categorize
 * @returns {Object} Categorized docs object
 */
function categorizeDocItems(items) {
  const categories = {
    "Language Basics": [],
    "Data Types": [],
    "Control Flow": [],
    "Functions & Modules": [],
    "Built-in Modules": [],
    Database: [],
    "Web & HTTP": [],
    "Development Tools": [],
    "Advanced Features": [],
  };

  items.forEach((item) => {
    const title = item.title.toLowerCase();
    const filename = item.filename.toLowerCase();

    if (
      ["strings", "numbers", "bool", "arrays", "dictionaries", "null"].some(
        (type) => title.includes(type) || filename.includes(type)
      )
    ) {
      categories["Data Types"].push(item);
    } else if (
      ["if", "for", "while", "switch", "defer"].some(
        (flow) => title.includes(flow) || filename.includes(flow)
      )
    ) {
      categories["Control Flow"].push(item);
    } else if (
      ["function", "modules", "packages", "include"].some(
        (func) => title.includes(func) || filename.includes(func)
      )
    ) {
      categories["Functions & Modules"].push(item);
    } else if (
      ["mysql", "postgres", "sqlite"].some(
        (db) => title.includes(db) || filename.includes(db)
      )
    ) {
      categories["Database"].push(item);
    } else if (
      ["http", "net", "url", "email"].some(
        (web) => title.includes(web) || filename.includes(web)
      )
    ) {
      categories["Web & HTTP"].push(item);
    } else if (
      ["bundler", "cli", "debug", "editor", "tooling"].some(
        (tool) => title.includes(tool) || filename.includes(tool)
      )
    ) {
      categories["Development Tools"].push(item);
    } else if (
      ["async", "pointers", "reflect", "llm", "filewatcher"].some(
        (adv) => title.includes(adv) || filename.includes(adv)
      )
    ) {
      categories["Advanced Features"].push(item);
    } else if (
      [
        "os",
        "files",
        "path",
        "shell",
        "term",
        "dotenv",
        "sysinfo",
        "time",
        "datetime",
        "math",
        "crypto",
        "hash",
        "random",
        "uuid",
        "regex",
        "json",
        "xml",
        "csv",
        "encoding",
        "logger",
        "schedule",
      ].some((builtin) => title.includes(builtin) || filename.includes(builtin))
    ) {
      categories["Built-in Modules"].push(item);
    } else {
      categories["Language Basics"].push(item);
    }
  });

  // Remove empty categories
  Object.keys(categories).forEach((key) => {
    if (categories[key].length === 0) {
      delete categories[key];
    }
  });

  return categories;
}

/**
 * Main function to generate the docs registry
 * @returns {Promise<void>} Promise that resolves when generation is complete
 */
async function generateDocsRegistry() {
  try {
    // Path to docs directory (relative to project root)
    const docsDir = path.join(process.cwd(), "..", "docs");
    const outputPath = path.join(process.cwd(), "lib", "docs-generated.json");
    const typesOutputPath = path.join(process.cwd(), "lib", "docs-types.ts");

    console.log("Generating docs registry...");
    console.log(`Reading from: ${docsDir}`);
    console.log(`Writing to: ${outputPath}`);

    // Check if docs directory exists
    if (!fs.existsSync(docsDir)) {
      console.error(`Docs directory not found: ${docsDir}`);
      process.exit(1);
    }

    // Read all markdown files
    const files = fs.readdirSync(docsDir);
    const markdownFiles = files.filter((file) => file.endsWith(".md"));

    console.log(`Found ${markdownFiles.length} markdown files`);

    if (markdownFiles.length === 0) {
      console.warn("WARNING: No markdown files found in docs directory");
    }

    // Process each markdown file
    const docItems = [];

    for (const file of markdownFiles) {
      const filePath = path.join(docsDir, file);
      const content = fs.readFileSync(filePath, "utf-8");
      const docItem = extractDocInfo(file, content);
      docItems.push(docItem);
      console.log(`  ✓ Processed: ${file} -> "${docItem.title}"`);
    }

    // Sort by title for consistent ordering
    docItems.sort((a, b) => a.title.localeCompare(b.title));

    // Categorize the items
    const categorized = categorizeDocItems(docItems);

    // Create the registry object with warning header
    const timestamp = new Date().toISOString();
    const registry = {
      items: docItems,
      categorized,
      generatedAt: timestamp,
    };

    // Ensure lib directory exists
    const libDir = path.dirname(outputPath);
    if (!fs.existsSync(libDir)) {
      fs.mkdirSync(libDir, { recursive: true });
    }

    // Write the JSON file with proper formatting and warning
    const jsonContent = `{
  "_warning": "AUTO-GENERATED FILE - DO NOT EDIT MANUALLY",
  "_note": "This file is automatically generated at build time by scripts/build-docs.js",
  "_instruction": "To modify docs data, edit the markdown files in ../docs/ instead",
  "_generatedAt": "${timestamp}",
${JSON.stringify(registry, null, 2)
  .slice(1, -1)
  .split("\n")
  .slice(1)
  .join("\n")}
}`;

    fs.writeFileSync(outputPath, jsonContent);

    // Generate TypeScript types
    const typesContent = `// AUTO-GENERATED FILE - DO NOT EDIT MANUALLY
// This file is automatically generated at build time by scripts/build-docs.js
// To modify types, edit the script instead - your changes will be overwritten
// Generated at: ${timestamp}

export interface DocItem {
  title: string;
  href: string;
  description: string;
  filename: string;
}

export interface CategorizedDocs {
  [category: string]: DocItem[];
}

export interface DocsRegistry {
  items: DocItem[];
  categorized: CategorizedDocs;
  generatedAt: string;
}
`;

    fs.writeFileSync(typesOutputPath, typesContent);

    console.log(
      `Successfully generated docs registry with ${docItems.length} items`
    );
    console.log(`Categories: ${Object.keys(categorized).join(", ")}`);
    console.log(`Types generated: ${typesOutputPath}`);
    console.log(`Registry saved: ${outputPath}`);
  } catch (error) {
    console.error("Error generating docs registry:", error);
    process.exit(1);
  }
}

// Run the script when executed directly
if (import.meta.url === `file://${process.argv[1]}`) {
  generateDocsRegistry();
}

export { generateDocsRegistry };
