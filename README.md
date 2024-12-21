# N-gram index

[![lint](https://github.com/michurin/ngramindex/actions/workflows/lint.yaml/badge.svg)](https://github.com/michurin/ngramindex/actions/workflows/lint.yaml)
[![test](https://github.com/michurin/ngramindex/actions/workflows/test.yaml/badge.svg)](https://github.com/michurin/ngramindex/actions/workflows/test.yaml)
[![codecov](https://github.com/michurin/ngramindex/actions/workflows/codecov.yaml/badge.svg)](https://github.com/michurin/ngramindex/actions/workflows/codecov.yaml)
[![codecov](https://codecov.io/gh/michurin/ngramindex/graph/badge.svg?token=E9VIN7R58G)](https://codecov.io/gh/michurin/ngramindex)
[![Go Report Card](https://goreportcard.com/badge/github.com/michurin/ngramindex)](https://goreportcard.com/report/github.com/michurin/ngramindex)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/michurin/ngramindex)
[![go.dev play](https://shields.io/badge/go.dev-play-089?logo=go&logoColor=white&style=flat)](https://go.dev/play/p/yegNQQ9riTD)

## Motivation

Package offers advantages:

- Document type agnostic, thanks to generics.
- Rune based and Unicode friendly.
- Adjustable text normalization to manage things case sensibility, spaces and punctuation handling, extra typos tolerance etc.

## Examples

- [Life example](https://go.dev/play/p/yegNQQ9riTD).
- [Examples in documentation](https://pkg.go.dev/github.com/michurin/ngramindex), [the same examples right in this repository](https://github.com/michurin/ngramindex/blob/master/example_test.go).

## Known issues

- Beware: index modification is not thread safe.

## Related links

- [Russ Cox - Regular Expression Matching with a Trigram Index or How Google Code Search Worked](https://swtch.com/~rsc/regexp/regexp4.html).
