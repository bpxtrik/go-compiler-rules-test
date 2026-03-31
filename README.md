# Go Semantic Error Test Suite

This repository contains a set of Go code examples designed to test **semantic error detection** by the **compiler**. The goal is to create a reference showing which constructs are caught by the compiler.

---

## Tested Semantic and Quality Issues

The following categories are covered:

| Category                     | Description                                                  | Compiler | Notes                                                     |
| ---------------------------- | ------------------------------------------------------------ | -------- | --------------------------------------------------------- |
| **Unused Variables**         | Variables declared but never referenced within their scope   | ✅        | `declared and not used: x`. Also prevents code execution. |
| **Ignored Error Values**     | Function calls that return errors which are neither assigned nor checked | ❌        | Allows code compilation and execution.                    |
| **Variable Shadowing**       | Redeclarations of variables in inner scopes that shadow outer-scope variables, especially error variables | ❌        | Allows code compilation and execution.                    |
| **Duplicate/Unused Imports** | Detects imports that are never used or repeated              | ✅        | Prevents code compilation and execution.                  |
| **Unreachable Code**         | Statements after `return`, `panic`, or similar terminating statements | ❌        | Allows code compilation and execution.                    |
| **Empty Code Blocks**        | Empty blocks in constructs like `if`, `for`, or `switch`     | ❌        | Allows code compilation and execution.                    |

> **Legend:**  
> ❌ = Compiler does not flag  
> ⚠️ = Flagged by the compiler  
> ✅ = Disallows file execution

---

## Repository Structure

- Each semantic error detection has its own folder.

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