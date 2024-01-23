package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	test_s := "catsanddog"
	test_word_dict := []string{"cats","dog","sand","and","cat"}
	wanted := true

	test_result := wordBreak(test_s, test_word_dict)

	if test_result != wanted {
		t.Fatalf("test_result: %t, wanted: %t", test_result, wanted)
	}
}

func TestCheckBeginsWithWord(t *testing.T) {
	test_s := "thisisatest"
	test_word := "this"
	wanted := true 

	test_result := checkBeginsWithWord(test_word, test_s)

	if test_result != wanted {
		t.Fatalf("test_result: %t, wanted: %t", test_result, wanted)
	}
}


func TestRemoveWord(t *testing.T) {
	test_s := "thisisatest"
	test_word := "this"
	wanted := "isatest"
	
	test_result := removeWord(test_word, test_s)

	if test_result != wanted {
		t.Fatalf("test_result: %s, wanted: %s", test_result, wanted)
	}
}


func TestReduceWordDict(t *testing.T) {
	test_word_dict := []string{"a", "aa", "ab", "b"}
	wanted := []string{"a", "b"}

	test_result := reduceWordDict(test_word_dict)

	if !slicesEqual(test_result, wanted) {
		t.Fatalf("test_result: %s, wanted: %s", test_result, wanted)
	}
}


func slicesEqual(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return false
			}
		}
	}

	return true
}