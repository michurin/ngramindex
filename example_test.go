package ngramindex_test

import (
	"fmt"
	"strings"

	"github.com/michurin/ngramindex"
)

func Example_basic() {
	// Our documents
	docs := map[string]string{
		"Luke_22:35": `Then Jesus asked them, "When I sent you without purse, bag or sandals, did you lack anything?" "Nothing," they answered.`,
		"Luke_22:36": `He said to them, "But now if you have a purse, take it, and also a bag; and if you don’t have a sword, sell your cloak and buy one.`,
		"Luke_22:37": `It is written: 'And he was numbered with the transgressors'; and I tell you that this must be fulfilled in me. Yes, what is written about me is reaching its fulfillment."`,
		"Luke_22:38": `The disciples said, "See, Lord, here are two swords." "That’s enough!" he replied.`,
	}

	// In our "database" the type of index is string (T=string).
	// It could be integer, path string, [os.DirEntry], or anything else.
	ngIdx := ngramindex.StringIndex[string]()

	// Associate texts (v) with document's indexes (k)
	for k, v := range docs {
		ngIdx.Add(v, k)
	}

	// Search for "sword"
	// 22:38 wins because "Lord" is also matching with "ord" (end of "sword")
	results := ngIdx.Search("sword")
	for _, v := range results {
		fmt.Println(v)
	}

	// output:
	// Luke_22:38
	// Luke_22:36
}

func Example_textNormalization() {
	// Our documents
	docs := map[string]string{
		"Luke_22:35": `Then Jesus asked them, "When I sent you without purse, bag or sandals, did you lack anything?" "Nothing," they answered.`,
		"Luke_22:36": `He said to them, "But now if you have a purse, take it, and also a bag; and if you don’t have a sword, sell your cloak and buy one.`,
		"Luke_22:37": `It is written: 'And he was numbered with the transgressors'; and I tell you that this must be fulfilled in me. Yes, what is written about me is reaching its fulfillment."`,
		"Luke_22:38": `The disciples said, "See, Lord, here are two swords." "That’s enough!" he replied.`,
	}

	// Looking up is case insensitive and considers u and ú
	// as the same letter.
	// Obviously, it is useful for things like
	// spaces normalization, punctuation skipping an so on.
	ngIdx := ngramindex.StringIndex(
		ngramindex.WithNormolizer[string](func(s string) []rune {
			return []rune(strings.ToLower(strings.ReplaceAll(s, "ú", "u")))
		}),
	)

	// Associate texts (v) with document's indexes (k)
	for k, v := range docs {
		ngIdx.Add(v, k)
	}

	results := ngIdx.Search("JESÚS") // we will find "Jesus" in 22:35
	for _, v := range results {
		fmt.Println(v)
	}

	// output:
	// Luke_22:35
}

func Example_indexSettings() {
	ngIdx := ngramindex.StringIndex(ngramindex.WithNgramIndex(ngramindex.Index[int](ngramindex.WithNgramLen(2))))

	ngIdx.Add("what", 1)
	ngIdx.Add("that", 2)

	results := ngIdx.Search("with") // "th" is common in "with" and "that"
	fmt.Println(results)

	// output:
	// [2]
}
