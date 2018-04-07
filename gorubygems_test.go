package gorubygems

import (
	"testing"
)

func TestRubyGemsApiJustUpdated(t *testing.T) {
	gems := JustUpdated()

	if len(gems) != 50 {
		t.Errorf("Expected a response of recently updated gems")
	}
}

func TestRubyGemsSearch(t *testing.T) {
	term := "brakeman"
	gems := Search(term)

	if len(gems) < 1 {
		t.Errorf("Expected results for %+v", term)
	}
}

func TestRubyGemsLatest(t *testing.T) {
	gems := Latest()

	if len(gems) != 50 {
		t.Errorf("Expected a response of recently added gems")
	}
}

func TestRubyGemsGemInfo(t *testing.T) {
	gem := GemInfo("rails")

	if gem.Name != "rails" {
		t.Errorf("Expected gem info for the specified gem")
	}
}
