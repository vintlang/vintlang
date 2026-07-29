# Vint CLI Tooling

Vint provides several CLI tools to help you manage, format, and scaffold your projects:

---

## Formatter

Automatically formats your Vint code for consistency.

**Usage:**
```sh
vint fmt <file.vint>
```
This will overwrite the file with a pretty-printed version.

---

## Linter (Planned)

A linter will analyze your code for common mistakes and style issues.

**Planned Usage:**
```sh
vint lint <file.vint>
```

---

## Project Scaffolding

Quickly create a new Vint project with the recommended structure and sample files.

**Usage:**
```sh
vint init <project-name>
vint new <project-name>
```
This creates a new directory with a `main.vint`, `greetings_module.vint`, and a `vintconfig.json`.

---

## Package Manager

Install and manage Vint packages (currently supports installing `vintpm`).

**Usage:**
```sh
vint get <package>
```

---

For more information, run `vint help` or see the README. 