# go-eq-gen

[![test](https://github.com/github/docs/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/mwillfox/go-eq-gen/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mwillfox/go-eq-gen)](https://goreportcard.com/report/github.com/mwillfox/go-eq-gen)

A golang demo which explores evolutionary algorithms to generate an equation which
solves to a given integer.

![Example](https://github.com/mwillfox/go-eq-gen/blob/main/doc/example.gif)

## Why?

I originally started this project in 2015 to learn the go programming language and a
little about generative AI at the same time. At the time, I followed a tutorial on
genetic (evolutionary) algorithms with all code examples written in Java and ported
the code to go. I've since revisited the project and refactored the original code to
be more go idiomatic. The entire project is simply for fun and education.

## Known Limitations

### Small Target Size

This demo is limited to generating equations which solve to relatively small numbers
due to the chromosome length and the current algorithm only allowing single digit
operands.

### Solving Errors

The sub-operations in the generative equations solve to
integers. This leads to potential situations where `9 / 2 = 4`. If you encounter
a generative equation which is incorrect, this is most likely the cause.

## Future Improvements

I have begun work porting the original `string` gene encoding implementation to a `byte`
encoding implementation. I plan to benchmark the two implemtnations to explore the suspected
performance enhancements of the bytes solution.
