import { generateDocsRegistry, categorizeDocItems } from "./generate-docs";

// Cache for the docs registry to avoid regenerating on every request
let _docsCache: Awaited<ReturnType<typeof generateDocsRegistry>> | null = null;
let _categorizedCache: ReturnType<typeof categorizeDocItems> | null = null;

export async function getLearnItems() {
  if (!_docsCache) {
    _docsCache = await generateDocsRegistry();
  }
  return _docsCache;
}

export async function getCategorizedDocs() {
  if (!_categorizedCache) {
    const items = await getLearnItems();
    _categorizedCache = categorizeDocItems(items);
  }
  return _categorizedCache;
}

// For backwards compatibility, export a function that returns all items
export async function getAllDocItems() {
  return await getLearnItems();
}

// Legacy export for existing code (will be dynamically populated)
export const learnItems = [
  {
    title: "Strings",
    href: "/docs/learn/strings",
    description: "Learn about string manipulation and functions in VintLang.",
  },
  {
    title: "Numbers",
    href: "/docs/learn/numbers",
    description: "Understand number types and operations in VintLang.",
  },
];
