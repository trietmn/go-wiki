package test

import (
	"errors"
	"github.com/trietmn/go-wiki"
	"github.com/trietmn/go-wiki/page"
	"github.com/trietmn/go-wiki/utils"
	"strconv"
	"testing"
)

var (
	isSetUp bool = false
	Celtuce      = page.WikipediaPage{}
	Cyclone      = page.WikipediaPage{}
	Gwoc         = page.WikipediaPage{}
)

func TestMissingPage(t *testing.T) {
	utils.WikiRequester = MockRequester
	_, err := gowiki.GetPage("purpleberry", -1, false, true)
	if err == nil || err.Error() != errors.New("missing").Error() {
		t.Errorf("got %v, expect 'missing'", err)
	}
}

func TestRedirectPage(t *testing.T) {
	utils.WikiRequester = MockRequester
	res, err := gowiki.GetPage("Menlo Park, New Jersey", -1, false, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	et := "Edison, New Jersey"
	eu := "http://en.wikipedia.org/wiki/Edison,_New_Jersey"
	if res.Title != et {
		t.Errorf("got %v, expect %v", res.Title, et)

	}
	if res.URL != eu {
		t.Errorf("got %v, expect %v", res.URL, eu)
	}
}

func TestRedirectPageFalse(t *testing.T) {
	utils.WikiRequester = MockRequester
	_, err := page.MakeWikipediaPage(-1, "Menlo Park, New Jersey", "", false)
	if err == nil || err.Error() != errors.New("set the redirect argument to true to allow automatic redirects").Error() {
		t.Errorf("expect raise redirect error, but get %v", err)
	}
}

func TestRedirectNoNormalization(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := "Communist party"
	res, err := gowiki.GetPage("Communist Party", -1, false, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if res.Title != expectation {
		t.Errorf("got %v, expect %v", res.Title, expectation)
	}
}

func TestRedirectWithNormalization(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := "Communist party"
	res, err := gowiki.GetPage("communist Party", -1, false, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if res.Title != expectation {
		t.Errorf("got %v, expect %v", res.Title, expectation)
	}
}

func TestRedirectNormalization(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := "Communist party"
	lower, err := gowiki.GetPage("communist Party", -1, false, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	upper, err := gowiki.GetPage("Communist Party", -1, false, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if upper.Title != expectation {
		t.Errorf("got %v, expect %v", upper.Title, expectation)
	}
	if upper.Title != lower.Title {
		t.Errorf("got %v, expect %v", upper.Title, lower.Title)
	}
}

func TestDisambiguation(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := []string{"Dodge Ramcharger", "Dodge Ram Van", "Dodge Mini Ram", "Dodge Caravan C/V", "Ram C/V", "Dodge Ram 50", "Dodge D-Series", "Dodge Rampage", "Ram Trucks"}
	res, err := gowiki.GetPage("Dodge Ram (disambiguation)", -1, false, false)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(res.Disambiguation, expectation) {
		t.Errorf("got %v, expect %v", res.Disambiguation, expectation)
	}
}

func TestAutoSuggest(t *testing.T) {
	utils.WikiRequester = MockRequester
	expectation := "butterfly"
	e_url := "http://en.wikipedia.org/wiki/Butterfly"
	res, err := gowiki.GetPage("butteryfly", -1, true, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if res.Title != expectation {
		t.Errorf("got %v, expect %v", res.Title, expectation)
	}
	if res.URL != e_url {
		t.Errorf("got %v, expect %v", res.URL, e_url)
	}
}

func SetUpPageTest() error {
	if isSetUp {
		return nil
	}
	utils.WikiRequester = MockRequester
	var err error
	Celtuce, err = gowiki.GetPage("Celtuce", -1, true, true)
	if err != nil {
		return err
	}
	Cyclone, err = gowiki.GetPage("Tropical Depression Ten (2005)", -1, true, true)
	if err != nil {
		return err
	}
	Gwoc, err = gowiki.GetPage("Great Wall of China", -1, true, true)
	if err != nil {
		return err
	}
	isSetUp = true
	return nil
}

func TestFromPageID(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	res, err := gowiki.GetPage("", 1868108, true, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !Celtuce.Equal(res) {
		t.Errorf("different page")
	}
}

func TestTitle(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	exp1 := "Celtuce"
	exp2 := "Tropical Depression Ten (2005)"
	if Celtuce.Title != exp1 {
		t.Errorf("got %v, expect %v", Celtuce.Title, exp1)
	}
	if Cyclone.Title != exp2 {
		t.Errorf("got %v, expect %v", Cyclone.Title, exp2)
	}
}

func TestURL(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	exp1 := "http://en.wikipedia.org/wiki/Celtuce"
	exp2 := "http://en.wikipedia.org/wiki/Tropical_Depression_Ten_(2005)"
	if Celtuce.URL != exp1 {
		t.Errorf("got %v, expect %v", Celtuce.Title, exp1)
	}
	if Cyclone.URL != exp2 {
		t.Errorf("got %v, expect %v", Cyclone.Title, exp2)
	}
}

func TestContent(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetContent()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["celtuce.content"].(string) {
		t.Errorf("different content")
	}
	content, err = Cyclone.GetContent()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["cyclone.content"].(string) {
		t.Errorf("different content")
	}
}

func TestRevisionID(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetRevisionID()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["celtuce.revid"].(float64) {
		t.Errorf("different revid")
	}
	content, err = Cyclone.GetRevisionID()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["cyclone.revid"].(float64) {
		t.Errorf("different revid")
	}
}

func TestParentID(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetParentID()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["celtuce.parentid"].(float64) {
		t.Errorf("different parentid, got %v, expect %v", content, MockData["celtuce.parentid"].(float64))
	}
	content, err = Cyclone.GetParentID()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["cyclone.parentid"].(float64) {
		t.Errorf("different parentid, got %v, expect %v", content, MockData["cyclone.parentid"].(float64))
	}
}

func TestSummary(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetSummary()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["celtuce.summary"].(string) {
		t.Errorf("different summary")
	}
	content, err = Cyclone.GetSummary()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["cyclone.summary"].(string) {
		t.Errorf("different summary")
	}
}

func TestImages(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetImagesURL()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["celtuce.images"].([]interface{}))) {
		t.Errorf("different images")
	}
	content, err = Cyclone.GetImagesURL()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["cyclone.images"].([]interface{}))) {
		t.Errorf("different images")
	}
}

func TestReferences(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetReference()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["celtuce.references"].([]interface{}))) {
		t.Errorf("different references")
	}
	content, err = Cyclone.GetReference()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["cyclone.references"].([]interface{}))) {
		t.Errorf("different references")
	}
}

func TestLinks(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetLink()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["celtuce.links"].([]interface{}))) {
		t.Errorf("different links")
	}
	content, err = Cyclone.GetLink()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["cyclone.links"].([]interface{}))) {
		t.Errorf("different links")
	}
}

func TestCategory(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetCategory()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["celtuce.categories"].([]interface{}))) {
		t.Errorf("different categories")
	}
	content, err = Cyclone.GetCategory()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["cyclone.categories"].([]interface{}))) {
		t.Errorf("different categories")
	}
}

func TestHTML(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Celtuce.GetHTML()
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["celtuce.html"].(string) {
		t.Errorf("different html")
	}
}

func TestSectionList(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Cyclone.GetSectionList()
	if err != nil {
		t.Errorf("%v", err)
	}
	if !utils.CompareSlice(content, utils.TurnSliceOfString(MockData["cyclone.sections"].([]interface{}))) {
		t.Errorf("different sections, got %v", content)
	}
}

func TestSection(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Cyclone.GetSection("Impact")
	if err != nil {
		t.Errorf("%v", err)
	}
	if content != MockData["cyclone.section.impact"] {
		t.Errorf("different section content")
	}
	content, err = Cyclone.GetSection("history")
	if err == nil {
		t.Errorf("expect section not exist")
	}
}

func TestCoordinates(t *testing.T) {
	err := SetUpPageTest()
	if err != nil {
		t.Errorf("%v", err)
	}
	utils.WikiRequester = MockRequester
	content, err := Gwoc.GetCoordinate()
	if err != nil {
		t.Errorf("%v", err)
	}
	abs := func(x float64) float64 {
		if x < 0 {
			return (-1) * x
		}
		return x
	}
	v, _ := strconv.ParseFloat(MockData["great_wall_of_china.coordinates.lat"].(string), 64)
	if abs(content[0]-v) > 0.003 {
		t.Errorf("different latitude, got %v, expect %v", content[0], v)
	}
	v, _ = strconv.ParseFloat(MockData["great_wall_of_china.coordinates.lon"].(string), 64)
	if abs(content[1]-v) > 0.003 {
		t.Errorf("different longitude, got %v, expect %v", content[1], v)
	}
}

func TestGetFinalSection(t *testing.T) {
	page, err := gowiki.GetPage("Rafael Nadal", -1, false, true)
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = page.GetSection("External links")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestGetEmptyCoordinate(t *testing.T) {
	page, err := gowiki.GetPage("Rafael Nadal", -1, false, true)
	if err != nil {
		t.Errorf("%v", err)
	}

	_, err = page.GetCoordinate()
	if err != nil {
		t.Errorf("%v", err)
	}
}
