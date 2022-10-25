package test

import (
	"testing"

	"github.com/trietmn/go-wiki"
	"github.com/trietmn/go-wiki/utils"
)

func TestGeosearch(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_seach"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, -1, "", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

func TestGeosearchRadius(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_seach_with_radius"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, 10000, "", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

func TestGeosearchTitle(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_seach_with_existing_article_name"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, -1, "Great Wall of China", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

func TestGeosearchNotExistTitle(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_seach_with_non_existing_article_name"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, -1, "Test", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}
