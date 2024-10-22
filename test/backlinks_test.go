package test

import (
	"testing"

	gowiki "github.com/trietmn/go-wiki"
	"github.com/trietmn/go-wiki/utils"
)

func TestBacklinks(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.backlinks"].([]interface{}))
	res, err := gowiki.GetBacklinks("Great Wall of China")
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}

}
