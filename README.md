# N-gram Indexing and Searching

[![lint](https://github.com/michurin/ngramindex/actions/workflows/lint.yaml/badge.svg)](https://github.com/michurin/ngramindex/actions/workflows/lint.yaml)
[![test](https://github.com/michurin/ngramindex/actions/workflows/test.yaml/badge.svg)](https://github.com/michurin/ngramindex/actions/workflows/test.yaml)
[![codecov](https://github.com/michurin/ngramindex/actions/workflows/codecov.yaml/badge.svg)](https://github.com/michurin/ngramindex/actions/workflows/codecov.yaml)
[![codecov](https://codecov.io/gh/michurin/ngramindex/graph/badge.svg?token=E9VIN7R58G)](https://codecov.io/gh/michurin/ngramindex)
[![Go Report Card](https://goreportcard.com/badge/github.com/michurin/ngramindex)](https://goreportcard.com/report/github.com/michurin/ngramindex)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/michurin/ngramindex)
[![go.dev play](https://shields.io/badge/go.dev-play-089?logo=go&logoColor=white&style=flat)](https://go.dev/play/p/QClnrDlruau)

N-gram indexing is a simple and powerful lookup technique. It is based on approximate (fuzzy) string matching.

## Motivation

The package offers advantages:

- Document type agnostic, thanks to generics.
- Rune based and Unicode friendly.
- Adjustable text normalization to manage things like case sensibility, spaces and punctuation handling, extra typos tolerance etc.
- Simple ranking algorithm out of the box.
- Ability to customize ranking algorithm entirely up to your implementation of less-function for sorting.
- Ability to associate one document with several texts and lookup by several texts

## Examples

- [Life example](https://go.dev/play/p/QClnrDlruau).
- [Examples in documentation](https://pkg.go.dev/github.com/michurin/ngramindex), [the same examples right in this repository](https://github.com/michurin/ngramindex/blob/master/example_test.go).

## Known issues

- Beware: index modification is not thread safe.
- It is in-memory implementation.
- There is no way to import/export/save/restore the index.
- It is impossible to remove document from index.

## Related links

- [Russ Cox - Regular Expression Matching with a Trigram Index or How Google Code Search Worked](https://swtch.com/~rsc/regexp/regexp4.html).
