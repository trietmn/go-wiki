package test

import (
	"github.com/trietmn/go-wiki"
	"github.com/trietmn/go-wiki/utils"
	"testing"
)

func TestSearch(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["barack.search"].([]interface{}))
	res, _, err := gowiki.Search("Barack Obama", -1, false)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

func TestSearchLimit(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["porsche.search"].([]interface{}))
	res, _, err := gowiki.Search("Porsche", 3, false)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

func TestSuggestion(t *testing.T) {
	utils.WikiRequester = MockRequester
	res, sug, err := gowiki.Search("hallelulejah", -1, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(res) > 0 {
		t.Errorf("result length should be 0")
	}
	if sug != "hallelujah" {
		t.Errorf("got %v, expect hallelujah", sug)
	}
}

func TestNoSuggestion(t *testing.T) {
	utils.WikiRequester = MockRequester
	res, sug, err := gowiki.Search("qmxjsudek", -1, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(res) > 0 {
		t.Errorf("result length should be 0")
	}
	if sug != "" {
		t.Errorf("got %v, expect nothing", sug)
	}
}
