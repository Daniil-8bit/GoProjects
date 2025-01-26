package sentences

import (
	"fmt"
	"testing"
)

/*func TestTwoWords(t *testing.T) {
	//t.Error("Nothing here")
	words := []string{"one", "two"}
	received := JoinSomeWords(words)
	expected := "one and two"
	if received != expected {
		t.Errorf("Expected value of JoinSomeWords(%v): '%s', but value: '%s'", words, expected, received)
	}
}

func TestThreeWords(t *testing.T) {
	//t.Error("Nothing here too")
	words := []string{"one", "two", "three"}
	received := JoinSomeWords(words)
	expected := "one, two, and three"
	if received != expected {
		t.Errorf("Expected value of JoinSomeWords(%v): '%s', but value: '%s'", words, expected, received)
	}
}*/

/*func TestSingleWord(t *testing.T) {
	words := []string{"one"}
	received := JoinSomeWords(words)
	expected := "one"
	if received != expected {
		t.Errorf(showError(words, expected, received))
	}
}

func TestTwoWords(t *testing.T) {
	//t.Error("Nothing here")
	words := []string{"one", "two"}
	received := JoinSomeWords(words)
	expected := "one and two"
	if received != expected {
		t.Errorf(showError(words, expected, received))
	}
}

func TestThreeWords(t *testing.T) {
	//t.Error("Nothing here too")
	words := []string{"one", "two", "three"}
	received := JoinSomeWords(words)
	expected := "one, two, and three"
	if received != expected {
		t.Errorf(showError(words, expected, received))
	}
}*/

func showError(w []string, e string, r string) string {
	return fmt.Sprintf("Expected value of JoinSomeWords(%v): '%s', but value: '%s'", w, e, r)
}

type testTable struct {
	wordsList   []string
	expectedRes string
}

func TestJoinWords(t *testing.T) {
	ruleTable := []testTable{
		{wordsList: []string{}, expectedRes: ""},
		{wordsList: []string{"one"}, expectedRes: "one"},
		{wordsList: []string{"one", "two"}, expectedRes: "one and two"},
		{wordsList: []string{"one", "two", "three"}, expectedRes: "one, two, and three"},
	}

	for _, val := range ruleTable {
		received := JoinSomeWords(val.wordsList)
		if received != val.expectedRes {
			t.Errorf(showError(val.wordsList, val.expectedRes, received))
		}
	}
}
