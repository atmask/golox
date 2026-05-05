# golox

A Go implementation of the Lox programming language from [Crafting Interpreters](https://craftinginterpreters.com/) by Robert Nystrom.

## Usage

```bash
# Run a script
golox script.lox

# Start the REPL
golox
```

## Build

Requires Go 1.26+ and [just](https://github.com/casey/just).

```bash
just build
just run [script]
just test
just clean
```
