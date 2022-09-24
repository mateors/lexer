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
 A prefix operator is an operator "in front of" its operand.
 * -5
 * !true
 * !false

### Postfix operators
 A postfix operator is an operator "after" its operand

 ### Infix operators (or "binary operators")
 An infix operator sits between its operands
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

## Pratt Parsing
* [Simple-but-powerful-pratt-parsing](https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html)
* [Parsing-made-easy](https://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy)
* [Recursive-descent-and-pratt-parsing](https://chidiwilliams.com/post/on-recursive-descent-and-pratt-parsing)
* [Handwritten a Parser](https://segmentfault.com/a/1190000041457544/en)

*"Pratt Parsing is very simple to understand, trivial to implement, easy to use, extremely effcient in practice if not in theory, yet flexible enough to meet most reasonable syntactic needs of users"*

Top Down Operator Precedence Parsing, or Pratt parsing, was invented as `an alternative to parsers` based on context-free grammars and the Backus-Naur-Form.

And that is also the `main difference`: instead of associating parsing functions with `grammar rules (defined in BNF or EBNF)`, Pratt associates these functions (which he calls `"semantic code"`) with single token types.

A crucial part of this idea is that each token type can have two parsing functions associated with it, depending on the token's position - `infix` or `prefix`.

* The first thing we need to do for expression parsing is to prepare our AST.

> Let statement -> `let x = 5 + 5;`

> Expression statement -> `x + 5;`


## Implementing the Pratt Parser
A Pratt parser’s main idea is the association of parsing functions *(which Pratt calls "semantic code")* with token types.

Each token type can have up to two parsing functions associated with it, depending on whether the token is found in a prefix or an infix position.

## Prefix operator syntax
> `<prefix operator><expression>;`
*unary expressions*

## Infix operator syntax
> `<expression> <infix operator> <expression>;`
*Because of the two operands (left and right) these expressions are sometimes called "binary expressions"*


## Table-driven testing approach
* [Deep-dive-into-table-driven-testing](https://engineering.mercari.com/en/blog/entry/20211221-a-deep-dive-into-table-driven-testing-in-golang)
* [gotests makes writing Go tests easy](https://github.com/cweill/gotests)

### Here is what happens when we parse `1 + 2 + 3;`

![parse 1](./screens/parse_1.png)

![parse 2](./screens/parse_2.png)

![parse 3](./screens/parse_3.png)

![parse 4](./screens/parse_4.png)

### when parsing the expression statement `-1 * 2 + 3`

```
BEGIN parseExpressionStatement
        BEGIN parseExpression
                BEGIN parsePrefixExpression
                        BEGIN parseExpression
                                BEGIN parseIntegerLiteral
                                END parseIntegerLiteral
                        END parseExpression
                END parsePrefixExpression
                BEGIN parseInfixExpression
                        BEGIN parseExpression
                                BEGIN parseIntegerLiteral
                                END parseIntegerLiteral
                        END parseExpression
                END parseInfixExpression
                BEGIN parseInfixExpression
                        BEGIN parseExpression
                                BEGIN parseIntegerLiteral
                                END parseIntegerLiteral
                        END parseExpression
                END parseInfixExpression
        END parseExpression
END parseExpressionStatement
```

## Block statements 
Block statements are a series of statements enclosed by an opening `{` and a closing `}`

## if-else conditional statement structure
> `if (<condition>) <consequence> else <alternative>`

> *parsing is prone to off-by-one errors*

## Structure of function literals

### Function literals look like this:
```
fn(x, y) {
return x + y;
}
```

It starts with the keyword fn, followed by a list of parameters, followed by a block statement, which is the function’s body, that gets executed when the function is called. The abstract structure of a function literal is this:

> `fn <parameters> <block statement>`

### The parameters in function
They are just a list of identifiers that are comma-separated and surrounded by parentheses:

> `(<parameter one>, <parameter two>, <parameter three>, ...)`

This list can also be empty:
```
fn() {
return foobar + barfoo;
}
```

### the two main parts of a function literals
1. The list of parameters 
2. The block statement that is the function's body. 

That’s all we need to keep in mind when defining the AST node:

## Call Expressions structure
* `<expression>( <comma separated expressions> )`

## Evaluation
A treewalking interpreter that recursively evaluates an AST is probably the slowest of all approaches, but easy to build, extend, reason about and as portable as the language it’s implemented in.

An interpreter that compiles to bytecode and uses a virtual machine to evaluate said bytecode is going to be a lot faster. But more complicated and harder to build, too.

## A Tree-Walking Interpreter
> Executing the AST while traversing it. 

We are going to build is a `tree-walking interpreter`. We are going to `take the AST` our parser builds for us `and interpret it` `"on the fly"`, without any preprocessing or compilation step.

The design we are going to use is heavily inspired by the interpreter presented in `The Structure and Interpretation of Computer Programs` (SICP), especially its usage of environments.

> Becase it's the easiest way to get started, it's easy to understand and to extend later on.

### We only need two things really: 
1. A tree-walking evaluator 
2. A way to represent values in Golang (host language).

* If you want to build a fast interpreter you can't get away with a slow and bloated object system.
* And if you are going to write your own garbage collector, you need to think about how it will keep track of the values in the system.
* If you don’t care about performance, then it does make sense to keep things simple and easy to understand until further requirements arise.


## Structure of Eval (first version)
> `func Eval(node ast.Node) object.Object`

Eval will take an `ast.Node` as input and return an `object.Object`.

## Environment
> the environment is a hash map that associates strings with objects.

## Functions
1. Define an internal representation of functions in our object system 
2. Add support for function calls to Eval.

> Extending the environment

Closures are functions that “close over” the environment they were defined in. They carry their own environment around and whenever they’re called they can access it.

## Higher-order functions
Higher-order functions are functions that either return other functions or receive them as arguments.

## Garbage collection (GC)
In short: keep track of object allocations and references to objects, make enough memory available for future object allocations and give memory back when it’s not needed anymore. This last point is what garbage collection is all about. Without it the programs would “leak” and finally run out of memory

### mark and sweep algorithm.
> allocating and freeing memory

### system calls
> Talking to the kernel is normally done via something called system calls.

### Array
elements in an array literal can be any type of expression. Integer literals, function literals, infix or prefix expressions.

### Index operator expressions
* myArray[0];
* [1, 2, 3, 4][2];
* let myArray = [1, 2, 3, 4]; myArray[2];
* myArray[2 + 1];
* returnsArray()[1];

### Basic Structure of Array expression
> `<expression>[<expression>]`

## Testing our lexer
* `go test ./lexer/`
* `go test ./ast/`
* `go test ./parser/`
* `go test -run TestOperatorPrecedenceParsing ./parser/`
* `go test -v -run TestOperatorPrecedenceParsing2 ./parser/`