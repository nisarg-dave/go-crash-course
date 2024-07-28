# Go Crash Course

1. Go is a statically typed.
2. Go is strongly typed. Meaning that the type of a variable is determined at compile time. You can't add a string to an integer, unlike JavaScript.
3. Go is compiled. Meaning that the code is converted to machine code before it is run. It produces a binary executable.
4. Go has fast compilation times and also supports cross-compilation, meaning that you can compile code for multiple platforms.
5. Go has built in concurrency.
6. Go is known for its simplicity. This simplicity reduces the cognitive load on developers, making it easier to write, review, and maintain code.
7. Go comes with a standard library. This library contains many useful functions and packages that can be used to build applications.
8. Go includes garbage collection. This means that the programmer doesn't have to worry about memory management, as the language takes care of it for you.
9. Go is scalable. This means that it can handle large amounts of data and can be used to build large-scale applications.

## Modules and Packages

Package is a folder of Go files, a collection of these packages is called a module. When initialising a new project, we are really initialising a module (go mod init).

## How to compile and run a Go program

```bash
go run main.go
```
