// Auto-generated types for docs registry
// Generated at: 2025-09-26T14:06:40.619Z

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
