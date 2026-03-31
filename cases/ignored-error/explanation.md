## Ignored error explanation

# Issue
Does the compiler catch ignored error during function call?

# Response
The compiler does not catch ignored error
Allows both compilation and execution

# Potential consequences
It can cause unintended bugs, such as nil pointer runtime error.
In this case, the file is <nil> which causes unintended behaviour. Any other call on the `file` instance will result in nil. This makes debugging even harder, because no runtime error will be thrown.
