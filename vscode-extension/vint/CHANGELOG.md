# Change Log

All notable changes to the "vint" extension will be documented in this file.

Check [Keep a Changelog](http://keepachangelog.com/) for recommendations on how to structure this file.

## [0.1.0] - 2024-12-19

### Fixed
- Updated function snippets to use correct VintLang syntax (`let name = func()` instead of `function`)
- Fixed import snippets to match VintLang syntax (removed `from` keyword)
- Removed unnecessary semicolons from snippets to match VintLang style
- Fixed package.json structure (moved misplaced commands and keybindings)
- Added snippets contribution to package.json

### Added
- `println` snippet for print with newline
- `convert` snippet for type conversion
- Updated package snippet to use block syntax

### Changed
- Updated VSCode engine requirement to ^1.74.0 for better compatibility
- Bumped extension version to 0.1.0
- Enhanced language aliases to include "VintLang"

## [0.0.1] - Initial

- Initial release