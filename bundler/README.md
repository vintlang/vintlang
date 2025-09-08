# VintLang Bundler - Quick Reference

## What is it?
The VintLang Bundler transforms multi-file VintLang projects into self-contained Go binaries.

## Simple Usage
```bash
vint bundle main.vint
```

## How it Works (Simplified)
```
Your VintLang Files → Dependency Analysis → Code Processing → Go Code Generation → Compiled Binary
```

## Key Features
- 🔍 **Auto-discovery**: Finds all imported/included files automatically
- 📦 **Multi-file support**: Handles complex projects with imports and includes  
- 🏗️ **Package system**: Proper module resolution and package wrapping
- 🚀 **Self-contained**: No external dependencies needed to run the binary
- 🔧 **Cross-platform**: Build for different OS/architectures

## Import vs Include
- `import module_name` → Wraps content in packages (modular approach)
- `include "file.vint"` → Embeds content directly (simple embedding)

## Example Project Structure
```
project/
├── main.vint           # import utils; include "config.vint"
├── utils.vint          # package utils { ... }
└── config.vint         # let appName = "My App"
```

Run `vint bundle main.vint` and get a single executable that contains all three files!

## More Details
See the full documentation at `docs/bundler.md` for complete technical details, architecture diagrams, and advanced usage examples.