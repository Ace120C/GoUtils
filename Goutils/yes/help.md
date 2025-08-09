## Understanding the code:

`os.Args` is a slice of strings in Go's `os` package that holds everything the user typed after the program's name.

```go
if len(os.Args) > 1 {
    str = os.Args[1]
}
```
the logic behind this code block is we are checking if there is an argument after the programs name or not in the command line the default is
`os.Args[] = ["program's name. go"]` but if you added an argument after it, it will be `os.Args[] = ["program's name.go", "your argument"]` instead
so in short if the slice has more than 1 item the program will use the word at index 1 instead.
