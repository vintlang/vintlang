// Type declarations for JSON imports
declare module "*.json" {
  const value: any;
  export default value;
}

// Specific declaration for our generated docs file
declare module "./docs-generated.json" {
  import { DocsRegistry } from "./docs-types";
  const registry: DocsRegistry;
  export default registry;
}