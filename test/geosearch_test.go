package test

import (
	"testing"
	"time"

	"github.com/trietmn/go-wiki"
	"github.com/trietmn/go-wiki/utils"
)

// Test the Geosearch function
func TestGeosearch(t *testing.T) {
	time.Sleep(time.Second / 4)
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_search"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, -1, "", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

// Test Geosearch with radius
func TestGeosearchRadius(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_search_with_radius"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, 10000, "", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

// Test Geosearch with title
func TestGeosearchTitle(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_search_with_existing_article_name"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, -1, "Great Wall of China", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}

// Test Geosearch with a not existed title
func TestGeosearchNotExistTitle(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := utils.TurnSliceOfString(MockData["great_wall_of_china.geo_search_with_non_existing_article_name"].([]interface{}))
	res, err := gowiki.GeoSearch(40.67693, 117.23193, -1, "Test", -1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res, expectation) {
		t.Errorf("got %v, expect %v", res, expectation)
	}
}
