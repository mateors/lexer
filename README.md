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

## Parsing Expressions

 ### Prefix operators
 * -5
 * !true
 * !false

 ### Infix operators (or "binary operators")
 * 5 + 5
 * 5 - 5
 * 5 / 5
 * 5 * 5

### We can use parentheses to group expressions and influence the order of evaluation
* 5 * (5 + 5)
* ((5 + 5) * 5) * 5

### There are call expressions:
* add(2, 3)
* add(add(2, 3), add(5, 10))
* max(5, add(5, (5 * 5)))

### Identifiers are expressions too:
* foo * bar / foobar
* add(foo, bar)

## if expressions
* `let result = if (10 > 5) { true } else { false };`
* `result // => true`

## Testing our lexer
> `go test ./lexer/`
