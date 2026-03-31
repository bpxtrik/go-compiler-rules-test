## Duplicate/unused imports explanation

# Issue
Does the compiler catch duplicate/unused imports?

# Response
The compiler does catch both of the cases.
It prevents code compilation and execution with messages `fmt redeclared in this block` and `"fmt" imported and not used` respectively.
