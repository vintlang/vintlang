import fs from 'fs';
import path from 'path';

interface DocItem {
  title: string;
  href: string;
  description: string;
  filename: string;
}

/**
 * Extracts the title and description from a markdown file
 */
function extractDocInfo(filename: string, content: string): DocItem {
  const lines = content.split('\n');
  
  // Find the first h1 header (# Title)
  let title = filename.replace('.md', '').replace(/_/g, ' ');
  let description = '';
  
  for (let i = 0; i < lines.length; i++) {
    const line = lines[i].trim();
    
    // Extract title from first h1 header
    if (line.startsWith('# ') && !title) {
      title = line.substring(2).trim();
      continue;
    }
    
    // Extract description from first paragraph after title
    if (title && line && !line.startsWith('#') && !line.startsWith('```') && line.length > 20) {
      description = line.length > 150 ? line.substring(0, 150) + '...' : line;
      break;
    }
  }
  
  // Fallback description if none found
  if (!description) {
    if (title.includes('HTTP')) {
      description = `Learn about ${title.toLowerCase()} functionality in VintLang.`;
    } else if (title.toLowerCase().includes('function')) {
      description = `Learn about functions and function handling in VintLang.`;
    } else if (title.toLowerCase().includes('array')) {
      description = `Learn about arrays and array manipulation in VintLang.`;
    } else if (title.toLowerCase().includes('string')) {
      description = `Learn about string manipulation and functions in VintLang.`;
    } else {
      description = `Learn about ${title.toLowerCase()} in VintLang.`;
    }
  }
  
  return {
    title: formatTitle(title),
    href: `/docs/learn/${filename.replace('.md', '')}`,
    description,
    filename: filename.replace('.md', '')
  };
}

/**
 * Formats the title to be more readable
 */
function formatTitle(title: string): string {
  // Remove "in Vint", "in VintLang", etc. from titles
  title = title.replace(/\s+(in\s+)?(vint|vintlang)(\s+.*)?$/i, '');
  
  // Capitalize each word properly
  return title
    .split(' ')
    .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
    .join(' ')
    .replace(/\bHttp\b/g, 'HTTP')
    .replace(/\bApi\b/g, 'API')
    .replace(/\bSql\b/g, 'SQL')
    .replace(/\bJson\b/g, 'JSON')
    .replace(/\bXml\b/g, 'XML')
    .replace(/\bUuid\b/g, 'UUID')
    .replace(/\bCsv\b/g, 'CSV')
    .replace(/\bOs\b/g, 'OS')
    .replace(/\bIf\b/g, 'If')
    .replace(/\bFor\b/g, 'For');
}

/**
 * Generates the docs registry from markdown files
 */
export async function generateDocsRegistry(): Promise<DocItem[]> {
  const docsDir = path.join(process.cwd(), '..', 'docs');
  
  try {
    const files = fs.readdirSync(docsDir);
    const markdownFiles = files.filter(file => file.endsWith('.md'));
    
    const docItems: DocItem[] = [];
    
    for (const file of markdownFiles) {
      const filePath = path.join(docsDir, file);
      const content = fs.readFileSync(filePath, 'utf-8');
      const docItem = extractDocInfo(file, content);
      docItems.push(docItem);
    }
    
    // Sort by title for consistent ordering
    return docItems.sort((a, b) => a.title.localeCompare(b.title));
    
  } catch (error) {
    console.error('Error generating docs registry:', error);
    // Fallback to existing hardcoded items if file reading fails
    return [
      {
        title: "Strings",
        href: "/docs/learn/strings",
        description: "Learn about string manipulation and functions in VintLang.",
        filename: "strings"
      }
    ];
  }
}

/**
 * Generates categories for better organization
 */
export function categorizeDocItems(items: DocItem[]): { [category: string]: DocItem[] } {
  const categories = {
    'Language Basics': [] as DocItem[],
    'Data Types': [] as DocItem[],
    'Control Flow': [] as DocItem[],
    'Functions & Modules': [] as DocItem[],
    'Built-in Modules': [] as DocItem[],
    'Database': [] as DocItem[],
    'Web & HTTP': [] as DocItem[],
    'Development Tools': [] as DocItem[],
    'Advanced Features': [] as DocItem[]
  };
  
  items.forEach(item => {
    const title = item.title.toLowerCase();
    const filename = item.filename.toLowerCase();
    
    if (['strings', 'numbers', 'bool', 'arrays', 'dictionaries', 'null'].some(type => title.includes(type) || filename.includes(type))) {
      categories['Data Types'].push(item);
    } else if (['if', 'for', 'while', 'switch', 'defer'].some(flow => title.includes(flow) || filename.includes(flow))) {
      categories['Control Flow'].push(item);
    } else if (['function', 'modules', 'packages', 'include'].some(func => title.includes(func) || filename.includes(func))) {
      categories['Functions & Modules'].push(item);
    } else if (['mysql', 'postgres', 'sqlite'].some(db => title.includes(db) || filename.includes(db))) {
      categories['Database'].push(item);
    } else if (['http', 'net', 'url', 'email'].some(web => title.includes(web) || filename.includes(web))) {
      categories['Web & HTTP'].push(item);
    } else if (['bundler', 'cli', 'debug', 'editor', 'tooling'].some(tool => title.includes(tool) || filename.includes(tool))) {
      categories['Development Tools'].push(item);
    } else if (['async', 'pointers', 'reflect', 'llm', 'filewatcher'].some(adv => title.includes(adv) || filename.includes(adv))) {
      categories['Advanced Features'].push(item);
    } else if (['os', 'files', 'path', 'shell', 'term', 'dotenv', 'sysinfo', 'time', 'datetime', 'math', 'crypto', 'hash', 'random', 'uuid', 'regex', 'json', 'xml', 'csv', 'encoding', 'logger', 'schedule'].some(builtin => title.includes(builtin) || filename.includes(builtin))) {
      categories['Built-in Modules'].push(item);
    } else {
      categories['Language Basics'].push(item);
    }
  });
  
  // Remove empty categories
  Object.keys(categories).forEach(key => {
    const categoryKey = key as keyof typeof categories;
    if (categories[categoryKey].length === 0) {
      delete categories[categoryKey];
    }
  });
  
  return categories;
}