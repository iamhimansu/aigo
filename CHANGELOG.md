# Changelog

## Unreleased

### [DOCS] Added comprehensive architecture documentation
- Created `docs/architecture.md` with 5 detailed Mermaid diagrams.
- Includes: step-by-step architecture, mindmap, tree view, data flow sequence, and function dependency graph.
- Covers all 16 lessons across 7 phases (Core Math → Forward Pass → Loss → Backprop → Optimization → Persistence → NLP).

### [REFACTOR] Restructured project into ordered lesson directories
- Moved 16 standalone Go files into numbered subdirectories (`01_dot_product/` → `16_text_generator/`).
- Each lesson is now an isolated `main.go` — eliminates all duplicate-function compilation errors.
- Moved `brain.json` into `13_save_model/` where it belongs.
- Added `go.mod` (module `aigo`).
- Added `README.md` with lesson table and usage instructions.
