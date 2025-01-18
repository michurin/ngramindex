package ngramindex

import (
	"sort"
)

const defaultNgramLen = 3

type docInfo[T any] struct {
	doc        T
	ngramCount int
}

// NgramIndex low-level rune-based implementation. Trigram by default (see [Index] and [WithNgramLen] option).
//
// You might want to use [StringNgramIndex] that works with strings
// and provides string normalization.
type NgramIndex[T any] struct {
	idx          map[string][]int
	docs         []docInfo[T]
	ngramLen     int
	cutOffWeight int
}

func Index[T any](opts ...Option) *NgramIndex[T] {
	cfg := indexConfig{ngramLen: defaultNgramLen}
	for _, opt := range opts {
		opt(&cfg)
	}

	return &NgramIndex[T]{
		idx:          make(map[string][]int),
		docs:         nil,
		ngramLen:     cfg.ngramLen,
		cutOffWeight: 0, // TODO: make option? or it has to be an Search's argument?
	}
}

// Add establish association between text and document.
// Text can be any text representation of document.
// For example, concatenation of title and body.
// The doc can be any pointer to "document":
// index in DB, path in a filesystem, [io/fs.DirEntry], or anything else.
func (ni *NgramIndex[T]) Add(text []rune, doc T) {
	if len(text) < ni.ngramLen {
		return
	}

	idx := len(ni.docs)
	ni.docs = append(ni.docs, docInfo[T]{
		doc:        doc,
		ngramCount: len(text) - ni.ngramLen,
	})

	for i := 0; i < len(text)-ni.ngramLen+1; i++ {
		ngram := string(text[i : i+ni.ngramLen])
		ni.idx[ngram] = append(ni.idx[ngram], idx)
	}
}

// Search is slightly oversimplified method. It returns sorted results.
func (ni *NgramIndex[T]) Search(text []rune) []T {
	filtered := ni.lookup(text)

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].MatchedNgrams > filtered[j].MatchedNgrams
	})

	res := make([]T, len(filtered))
	for i, v := range filtered {
		res[i] = v.Document
	}

	return res
}

type Match[T any] struct {
	Document      T       // Document itself
	TotalNgrams   int     // Total ngrams in document
	MatchedNgrams int     // Matched
	MatchRate     float64 // Literally MatchedNgrams/TotalNgrams
}

func (ni *NgramIndex[T]) Lookup(text []rune) []Match[T] {
	return ni.lookup(text)
}

func (ni *NgramIndex[T]) lookup(text []rune) []Match[T] {
	if len(text) < ni.ngramLen {
		return nil
	}

	cfg := make([]int, len(ni.docs))

	for i := 0; i < len(text)-ni.ngramLen+1; i++ {
		for _, idx := range ni.idx[string(text[i:i+ni.ngramLen])] {
			cfg[idx]++
		}
	}

	info := []struct {
		doc   docInfo[T]
		count int
	}(nil)

	for i, v := range cfg {
		if v > ni.cutOffWeight {
			info = append(info, struct {
				doc   docInfo[T]
				count int
			}{ni.docs[i], v})
		}
	}

	res := make([]Match[T], len(info))
	for i, v := range info {
		res[i] = Match[T]{
			Document:      v.doc.doc,
			TotalNgrams:   v.doc.ngramCount,
			MatchedNgrams: v.count,
			MatchRate:     float64(v.count) / float64(v.doc.ngramCount),
		}
	}

	return res
}

type indexConfig struct {
	ngramLen int
}

type Option func(*indexConfig)

func WithNgramLen(ngramLen int) Option {
	return func(ic *indexConfig) { ic.ngramLen = ngramLen }
}

// StringNgramIndex is convenient wrapper around [NgramIndex].
type StringNgramIndex[T any] struct {
	idx        *NgramIndex[T]
	normolizer func(string) []rune
}

// StringIndex is convenient wrapper around [Index], it
// based on strings and allows to perform custom strings normalization (see [WithNormolizer]).
func StringIndex[T any](opts ...StringIndexOption[T]) *StringNgramIndex[T] {
	idx := &StringNgramIndex[T]{
		idx:        Index[T](),
		normolizer: func(s string) []rune { return []rune(s) },
	}
	for _, opt := range opts {
		opt(idx)
	}

	return idx
}

// Add normalizes text and send it to [NgramIndex.Add].
func (si *StringNgramIndex[T]) Add(text string, doc T) {
	si.idx.Add(si.normolizer(text), doc)
}

// Search is simple wrapper around [NgramIndex.Search] with text normalization.
func (si *StringNgramIndex[T]) Search(text string) []T {
	return si.idx.Search(si.normolizer(text))
}

// Lookup is simple wrapper around [NgramIndex.Lookup] with text normalization.
func (si *StringNgramIndex[T]) Lookup(text string) []Match[T] {
	return si.idx.Lookup(si.normolizer(text))
}

type StringIndexOption[T any] func(*StringNgramIndex[T])

func WithNormolizer[T any](f func(string) []rune) StringIndexOption[T] {
	return func(si *StringNgramIndex[T]) { si.normolizer = f }
}

func WithNgramIndex[T any](idx *NgramIndex[T]) StringIndexOption[T] {
	return func(si *StringNgramIndex[T]) { si.idx = idx }
}
