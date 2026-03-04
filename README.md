# AIGO — Build AI From Scratch in Go

A step-by-step learning project that builds a neural network from the ground up. Each lesson is a standalone Go program you can run independently.

## Lessons

| # | Directory | Concept |
|---|---|---|
| 01 | `01_dot_product/` | Dot product — the atomic math operation |
| 02 | `02_activation/` | ReLU activation — non-linear gate |
| 03 | `03_predict/` | Linear prediction — dot product + bias |
| 04 | `04_neuron/` | Single neuron — predict → activate |
| 05 | `05_layer/` | Dense layer — array of neurons |
| 06 | `06_network/` | Forward pass — stacked layers |
| 07 | `07_loss/` | MSE loss — measuring error |
| 08 | `08_backprop/` | MSE derivative — error slope |
| 09 | `09_derivatives/` | ReLU derivative — activation slope |
| 10 | `10_gradients/` | Chain rule — weight-specific blame |
| 11 | `11_optimizer/` | SGD — weight update step |
| 12 | `12_train/` | Training loop — full learning cycle |
| 13 | `13_save_model/` | Persistence — save/load brain to JSON |
| 14 | `14_tokenizer_encode/` | Tokenization — words to integer IDs |
| 15 | `15_tokenizer_decode/` | Detokenization — IDs back to words |
| 16 | `16_text_generator/` | Bigram text generation |

## Architecture

See the full architecture breakdown with detailed Mermaid diagrams:

👉 **[docs/architecture.md](docs/architecture.md)**

Includes:
- **Step-by-step flowcharts** — internal logic of every module
- **Mindmap** — bird's-eye view of all 7 phases and functions
- **Tree view** — complete project file/function hierarchy
- **Data flow sequence** — how data moves through the training cycle
- **Function dependency graph** — which lessons reuse which functions

## Usage

Run any lesson individually:

```bash
go run ./01_dot_product/
go run ./04_neuron/
go run ./12_train/
```

Run all lessons at once:

```bash
for dir in 0*/ 1*/; do echo "=== $dir ===" && go run ./$dir; done
```

## Dependencies

- Go 1.18+ (standard library only, no external packages)

## Structure

Each lesson lives in its own directory with a `main.go` file. This isolation prevents Go from complaining about duplicate function names across lessons.
