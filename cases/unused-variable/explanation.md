## Unused variable explanation

# Issue
Does the compiler detect unused variable inside a scope?

# Response
The compiler does catch the unused variable with the message `declared and not used: x`.
Also, prevents code execution via `go run main.go`.
