# lexer

I truly wanted to understand how interpreters work and that included understanding how lexers and parsers work.

## What is an interpreter?
Interpreter: take source code and evaluate it.

> Interpreters that parse the source code, build an *abstract syntax tree (AST)* out of it and then evaluate this tree. This type of interpreter is sometimes called “tree-walking” interpreter, because it “walks” the AST and interprets it.

## What we want?
We build a tree-walking interpreter.

We're going to build our own **lexer**, our own parser, our own tree representation and our own evaluator. We'll see what **tokens** are, what an abstract syntax tree is, how to build such a tree, how to evaluate it and how to extend our language with new data structures and built-in functions.


## Why its important?
Without a compiler or an interpreter a programming language is nothing more than an idea or a specification.

## Testing our lexer
> `go test ./lexer/`
