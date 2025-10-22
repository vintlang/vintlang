# styled module

The `styled` module provides a small collection of terminal styling helpers exposed to VintLang programs. It is implemented in Go (see `module/styled.go`) and wraps the `github.com/fatih/color` package to print colored and styled text to the terminal.

This module is intended for CLI/REPL applications, demos, and examples where readable, colorized output improves developer experience.

## Overview

- Package path: `module` (internal runtime module)
- Purpose: Provide simple functions to print colored or styled output from Vint code.
- Behavior: Each function joins its arguments with spaces, converts them to strings via `Inspect()`, and prints the result with the corresponding style. Functions return `nil` on success or an `Error` object if passed `nil` values.

## Exported functions

The module exposes the following function names (available via the `module`'s runtime registration as `StyledFunctions`):

- red(...) — print in red
- green(...) — print in green
- yellow(...) — print in yellow
- blue(...) — print in blue

- header(...) — header/title style (bold green)
- error(...) — error style (bold red)
- success(...) — success style (bright green)
- warning(...) — warning style (bold yellow)
- info(...) — informational style (bold cyan)
- debug(...) — debug style (bold magenta)
- inputPrompt(...) — input prompt style (blue, bold)
- highlight(...) — highlight style (bright blue, bold)
- dim(...) — dim/faint text
- inverted(...) — inverted foreground/background

## Function contract

- Inputs: Any number of Vint values. Each argument is converted to string via its `Inspect()` method and joined with a space.
- Output: `nil` on success, or an `Error` object when an argument is `nil`.
- Side effects: Prints a single newline-terminated line to stdout using the selected color/style.

## Usage examples

From Vint code you can call these helpers once the module is registered and exposed by the runtime. Example (in a script):

```vint
// Example usage of styled functions
let name = "Vint"
// print a header
module.header("Starting", name)

// informational
module.info("Loaded module:", name)

// success and error messages
module.success("Operation completed successfully")
module.error("Unable to open connection")

// colored text
module.red("This is red text")
module.green("All good")

// prompt
module.inputPrompt("Enter value:")
```

Note: the exact symbol used from VintLang to call these functions depends on how the host runtime exposes `StyledFunctions` to Vint programs — some runtimes register them under `module.styled` or attach them to a `styled` object. Check your runtime's module registration code or REPL to see the exact name used at runtime.

## Error handling

If any argument is `nil` in the interpreter, the function will return an `Error` object with message `Operation cannot be performed on nil` and nothing will be printed.

## Implementation notes (for maintainers)

- The module is implemented in `module/styled.go`. It uses `fatih/color` to define reusable `*color.Color` instances and a `printStyled` helper.
- `StyledFunctions` is a `map[string]object.ModuleFunction` populated in `init()` with function names -> Go function handlers.
- Each named function delegates to `printStyled(args, style)` which converts args to strings using `Inspect()`.

## Tests and suggestions

- Consider adding unit tests that stub `object.VintObject` values to assert the printed output and error behavior.
- Consider exposing an API to return styled strings (instead of directly printing) for environments where printing is undesirable (e.g., logging systems or tests).

---

Docs created automatically by a maintainer tool.
