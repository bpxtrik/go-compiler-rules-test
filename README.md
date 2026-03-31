# Go Semantic Error Test Suite

This repository contains a set of Go code examples designed to test **semantic and code-quality issues** that are **not always detected by the Go compiler**. The goal is to create a reference showing which constructs are caught by the compiler.

---

## Tested Semantic and Quality Issues

The following categories are covered:

| Category                          | Description                                                  | Compiler | Notes                                                        |
| --------------------------------- | ------------------------------------------------------------ | -------- | ------------------------------------------------------------ |
| **Unused Variables**              | Variables declared but never referenced within their scope   | ❌        | Helps catch dead code and potential logic mistakes           |
| **Ignored Error Values**          | Function calls that return errors which are neither assigned nor checked | ❌        | Ignoring errors is a common source of bugs in Go programs    |
| **Variable Shadowing**            | Redeclarations of variables in inner scopes that shadow outer-scope variables, especially error variables | ❌        | Can lead to subtle bugs when outer variable changes are ignored |
| **Defer Statements Inside Loops** | Use of `defer` inside loops which delays resource release    | ❌        | Can lead to unexpected memory or file descriptor retention   |
| **Duplicate/Unused Imports**      | Detects imports that are never used or repeated              | ❌        | Keeps imports clean and avoids unnecessary compilation overhead |
| **Unreachable Code**              | Statements after `return`, `panic`, or similar terminating statements | ❌        | Improves code readability and maintainability                |
| **Empty Code Blocks**             | Empty blocks in constructs like `if`, `for`, or `switch`     | ❌        | Often indicates incomplete logic or placeholders             |

> **Legend:**  
> ❌ = Compiler does not flag  
> ⚠️ = Flagged by the compiler  
> ✅ = Disallows file execution

---

## Repository Structure

/cases/
 /unused-variable/
   main.go
   explanation.md
 /ignored-error/
   main.go
   explanation.md
 /variable-shadowing/
   main.go
   explanation.md
 /defer-in-loop/
   main.go
   explanation.md
 /duplicate-imports/
   main.go
   explanation.md
 /unreachable-code/
   main.go
   explanation.md
 /empty-blocks/
   main.go
   explanation.md

- Each folder contains a **minimal reproducible example**.
- `explanation.md` describes:
  - What the issue is
  - How the compiler responds
  - Potential consequences

---

## How to Test

For each case, run:

```bash
# Compile the code
go build main.go

# Optional: Run the code to observe runtime behavior
go run main.go