import { DocItem, CategorizedDocs, DocsRegistry } from "./docs-types";

// This will be generated at build time by scripts/build-docs.js
let _docsRegistry: DocsRegistry | null = null;

/**
 * Load the static docs registry (with fallback for development)
 */
async function loadDocsRegistry(): Promise<DocsRegistry> {
  if (_docsRegistry) {
    return _docsRegistry;
  }

  try {
    const registry = await import("./docs-generated.json");
    _docsRegistry = registry.default || registry;
    return _docsRegistry as DocsRegistry;
  } catch (error) {
    console.warn(
      "Failed to load generated docs registry, using fallback:",
      error
    );

    // Fallback for development when the file might not exist yet
    const fallbackRegistry: DocsRegistry = {
      items: [
        {
          title: "Strings",
          href: "/docs/learn/strings",
          description:
            "Learn about string manipulation and functions in VintLang.",
          filename: "strings",
        },
        {
          title: "Numbers",
          href: "/docs/learn/numbers",
          description: "Understand number types and operations in VintLang.",
          filename: "numbers",
        },
      ],
      categorized: {
        "Data Types": [
          {
            title: "Strings",
            href: "/docs/learn/strings",
            description:
              "Learn about string manipulation and functions in VintLang.",
            filename: "strings",
          },
          {
            title: "Numbers",
            href: "/docs/learn/numbers",
            description: "Understand number types and operations in VintLang.",
            filename: "numbers",
          },
        ],
      },
      generatedAt: new Date().toISOString(),
    };

    _docsRegistry = fallbackRegistry;
    return _docsRegistry;
  }
}

/**
 * Get all documentation items
 */
export async function getLearnItems(): Promise<DocItem[]> {
  const registry = await loadDocsRegistry();
  return registry.items;
}

/**
 * Get categorized documentation items
 */
export async function getCategorizedDocs(): Promise<CategorizedDocs> {
  const registry = await loadDocsRegistry();
  return registry.categorized;
}

/**
 * For backwards compatibility, export a function that returns all items
 */
export async function getAllDocItems(): Promise<DocItem[]> {
  return await getLearnItems();
}

/**
 * Legacy export for existing code - now dynamically loaded from static data
 */
export const learnItems: DocItem[] = [];
