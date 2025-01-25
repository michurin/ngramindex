package ngramindex //nolint: testpackage

import (
	"testing"
)

func TestIndex_tooShort(t *testing.T) { // major part of code coverage comes from examples
	t.Parallel()

	ngIdx := Index[int]()

	ngIdx.Add(1, []rune("ab"))
	ngIdx.Add(2, []rune("abc"))

	if len(ngIdx.docs) != 1 {
		t.Error("Expected 1 document, but got:", len(ngIdx.docs))
	}

	result := ngIdx.Search([]rune("ab"))
	if len(result) > 0 {
		t.Error("Expected empty result, but got:", result)
	}
}
