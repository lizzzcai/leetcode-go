package leetcode

import "testing"

func TestWordDictionary(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		obj := Constructor()
		obj.AddWord("bad")
		obj.AddWord("dad")
		obj.AddWord("mad")
		if obj.Search("pad") != false {
			t.Errorf("want %v, got %v", false, true)
		}

		if obj.Search("bad") != true {
			t.Errorf("want %v, got %v", true, false)
		}

		if obj.Search(".ad") != true {
			t.Errorf("want %v, got %v", true, false)
		}

		if obj.Search("b..") != true {
			t.Errorf("want %v, got %v", true, false)
		}
	})
}
