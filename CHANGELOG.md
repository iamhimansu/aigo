# Changelog

## Unreleased

### [REFACTOR] Restructured project into ordered lesson directories
- Moved 16 standalone Go files into numbered subdirectories (`01_dot_product/` → `16_text_generator/`).
- Each lesson is now an isolated `main.go` — eliminates all duplicate-function compilation errors.
- Moved `brain.json` into `13_save_model/` where it belongs.
- Added `go.mod` (module `aigo`).
- Added `README.md` with lesson table and usage instructions.
