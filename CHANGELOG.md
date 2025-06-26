# Changelog

## [Unreleased] - Last 7 Days

### Major Features & Enhancements
- **Compiler & VM Foundation**
  - Introduced a new bytecode compiler and virtual machine for VintLang.
  - Implemented support for integer arithmetic, boolean literals, and comparison operators (`+`, `-`, `*`, `/`, `<`, `>`, `==`, `!=`).
  - Added comprehensive tests for the compiler and VM covering arithmetic and boolean logic.
  - Added a `test_vm.go` utility for manual testing of expression evaluation.

- **New Modules**
  - Added `path`, `random`, and `csv` modules with documentation and examples.
    - `path`: File path manipulation utilities.
    - `random`: Random number, string, and choice generation.
    - `csv`: Read and write CSV files easily.

- **Database Modules**
  - Implemented and registered new `mysql` and `postgres` modules.
  - Added connection, query, and fetch functions for both.
  - Provided documentation and usage examples.

### Refactoring & Code Quality
- **Evaluator Refactor**
  - Reduced code duplication in print-related builtins (`print`, `println`, `printErr`, `printlnErr`).
  - Improved error messages for clarity and consistency.

- **MySQL/Postgres Refactor**
  - Renamed MySQL functions for clarity and consistency.
  - Enhanced parameter conversion and connection management.

### Documentation
- **Docs Improvements**
  - Streamlined and expanded documentation for math, builtins, and new modules.
  - Added detailed usage examples for new modules and features.

### Bug Fixes
- Fixed error message formatting in the evaluator and math module.
- Updated error handling for nil/null values.

### Miscellaneous
- Commented out main function in `test_vm.go` for testing purposes.
- Cleaned up obsolete files and improved project structure.

---

**Highlights:**
You've made significant progress toward a modern, performant VintLang runtime with a new compiler/VM, expanded the standard library, and improved code quality and documentation across the board. 