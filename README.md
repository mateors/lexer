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

## Testing our lexer
* `go test ./lexer/`
* `go test ./ast/`
* `go test ./parser/`
* `go test -run TestOperatorPrecedenceParsing ./parser/`
* `go test -v -run TestOperatorPrecedenceParsing2 ./parser/`