package rubygemsapi

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
