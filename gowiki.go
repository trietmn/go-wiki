package gowiki

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/trietmn/go-wiki/cache"
	"github.com/trietmn/go-wiki/page"
	"github.com/trietmn/go-wiki/utils"
)

/*
Change the user-agent that you use to crawl Wikipedia data
*/
func SetUserAgent(user string) {

	utils.UserAgent = user
}

/*
Change the language of the API being requested.
Set `prefix` to one of the two letter prefixes found on the
list of all Wikipedia <http://meta.wikimedia.org/wiki/List_of_Wikipedias>.
Then clear all of the cache
*/
func SetLanguage(lang string) {

	utils.WikiLanguage = lang
	utils.Cache.Clear()
}

/*
Change the language of the API being requested.
Set `prefix` to one of the two letter prefixes found on the
list of all Wikipedia <http://meta.wikimedia.org/wiki/List_of_Wikipedias>.
Then clear all of the cache
*/
func SetURL(url string) {

	utils.WikiURL = url
	utils.Cache.Clear()
}

/*
Change the max number of the request responses stored in the Cache
*/
func SetMaxCacheMemory(n int) {

	cache.MaxCacheMemory = n
}

/*
Change the max duration of the request responses exist in the Cache
*/
func SetCacheDuration(x time.Duration) {

	cache.CacheExpiration = x
}

/*
List all the currently supported language prefixes (usually ISO language code).
Can be inputted to `set_lang` to change the Mediawiki that `wikipedia` requests results from.

Returns: Map of <prefix>: <local_lang_name> pairs.
*/
func GetAvailableLanguage() (map[string]string, error) {
	args := map[string]string{
		"action": "query",
		"meta":   "siteinfo",
		"siprop": "languages",
	}
	res, err := utils.WikiRequester(args)
	if err != nil {
		return map[string]string{}, err
	}
	if res.Error.Code != "" {
		return map[string]string{}, errors.New(res.Error.Info)
	}
	result := map[string]string{}
	for _, v := range res.Query.Language {
		result[v["code"]] = v["*"]
	}
	return result, nil
}

/*
Do a Wikipedia search for `query`.

Keyword arguments:

* _input: The query used to search Ex:"Who invented the lightbulb"

* limit: The maxmimum number of results returned. Use -1 to use default setting

* suggest: If True, return results and suggestion (if any) in a tuple. Fasle is defalt

Return:

* A slice of Wikipedia titles from the search engine

* Suggestion if `suggest` is being set True

* Error
*/
func Search(_input string, limit int, suggest bool) ([]string, string, error) {
	if limit < 0 {
		limit = 10
	}
	args := map[string]string{
		"action":   "query",
		"list":     "search",
		"srprop":   "",
		"srlimit":  strconv.Itoa(limit),
		"srsearch": _input,
	}

	res, err := utils.WikiRequester(args)
	if err != nil {
		return []string{}, "", err
	}
	if res.Error.Code != "" {
		return []string{}, "", errors.New(res.Error.Info)
	}

	result := make([]string, 0, len(res.Query.Search))
	for _, s := range res.Query.Search {
		result = append(result, s.Title)
	}
	if suggest {
		return result, res.Query.SearchInfo.Suggestion, nil
	}
	return result, "", nil
}

/*
Get a Wikipedia search suggestion for `_input`.

Returns a string or "" if no suggestion was found.
*/
func Suggest(_input string) (string, error) {
	args := map[string]string{
		"action":   "query",
		"list":     "search",
		"srlimit":  "1",
		"srprop":   "",
		"srinfo":   "suggestion",
		"srsearch": _input,
	}

	res, err := utils.WikiRequester(args)
	if err != nil {
		return "", err
	}
	if res.Error.Code != "" {
		return "", errors.New(res.Error.Info)
	}
	return res.Query.SearchInfo.Suggestion, nil
}

/*
Do a wikipedia geo search for `latitude` and `longitude`
using HTTP API described in http://www.mediawiki.org/wiki/Extension:GeoData

Arguments:

* latitude: Latitude of the searched place

* longitude: longitude of the searched place

* title(optional): The title of an article to search for. Use "" to use the default setting

* limit(optional): The maximum number of results returned. Use -1 to use the default setting

* radius(optional): Search radius in meters. The value must be between 10 and 10000. Use -1 to use the default setting

Return:

* A slice of geosearch titles

* Error
*/
func GeoSearch(latitude float32, longitude float32, radius float32, title string, limit int) ([]string, error) {
	if radius <= 0 {
		radius = 1000
	}
	if limit < 0 {
		limit = 10
	}
	args := map[string]string{
		"action":   "query",
		"list":     "geosearch",
		"gsradius": fmt.Sprintf("%v", radius),
		"gscoord":  fmt.Sprintf("%v|%v", latitude, longitude),
		"gslimit":  strconv.Itoa(limit),
	}
	if title != "" {
		args["titles"] = title
	}
	res, err := utils.WikiRequester(args)
	if err != nil {
		return []string{}, err
	}
	if res.Error.Code != "" {
		return []string{}, errors.New(res.Error.Info)
	}

	result := make([]string, 0, len(res.Query.GeoSearch))
	if len(res.Query.Page) > 0 {
		for k, v := range res.Query.Page {
			if k != "-1" {
				result = append(result, v.Title)
			}
		}
	} else {
		for _, s := range res.Query.GeoSearch {
			result = append(result, s.Title)
		}
	}
	return result, nil
}

/*
Get a list of random Wikipedia article titles.

**Note:: Random only gets articles from namespace 0, meaning no Category, User talk, or other meta-Wikipedia pages.

Keyword arguments:

* limit: The number of random pages returned (max of 10)
*/
func GetRandom(limit int) ([]string, error) {
	if limit < 0 {
		limit = 5
	}
	args := map[string]string{
		"action":      "query",
		"list":        "random",
		"rnnamespace": "0",
		"rnlimit":     strconv.Itoa(limit),
	}
	res, err := utils.WikiRequester(args)
	if err != nil {
		return []string{}, err
	}
	if res.Error.Code != "" {
		return []string{}, errors.New(res.Error.Info)
	}
	result := make([]string, 0, len(res.Query.Random))
	for _, s := range res.Query.Random {
		result = append(result, s.Title)
	}
	return result, nil
}

/*
Get a WikipediaPage object for the page with title `title` or the pageid `pageid` (mutually exclusive).

Keyword arguments:

* title: The title of the page to load

* pageid: The numeric pageid of the page to load

* auto_suggest: Let Wikipedia find a valid page title for the query. Default should be False

* redirect: Allow redirection. Default should be True

Return:

* A WikipediaPage object

* Error
*/
func GetPage(title string, pageid int, suggest bool, redirect bool) (page.WikipediaPage, error) {
	if pageid >= 0 {
		return page.MakeWikipediaPage(pageid, "", "", redirect)
	}
	if title != "" {
		titles, suggestion, err := Search(title, 1, suggest)
		if err != nil {
			return page.MakeWikipediaPage(-1, title, "", redirect)
		}
		var pagetitle string
		if suggest {
			pagetitle = suggestion
		}
		if pagetitle == "" && len(titles) > 0 {
			pagetitle = titles[0]
		}
		if pagetitle == "" {
			return page.WikipediaPage{}, errors.New("page not exist")
		}
		return page.MakeWikipediaPage(-1, pagetitle, "", redirect)
	}
	return page.WikipediaPage{}, errors.New("must have either title or pageid to work")
}

/*
Return a string summary of a page

**Note:: This is a convenience wrapper - auto_suggest and redirect are enabled by default\

Keyword arguments:

* title: Title of the page you want to get the summary

* numsentence: If set, return the first `numsentence` sentences (can be no greater than 10).

* numchar: If set, return only the first `numchar` characters (actual text returned may be slightly longer).

* auto_suggest: Let Wikipedia find a valid page title for the query. Default is False

* redirect: Allow redirection without raising RedirectError. Defalt is True
*/
func Summary(title string, numsentence int, numchar int, suggest bool, redirect bool) (string, error) {
	page, err := GetPage(title, -1, suggest, redirect)
	if err != nil {
		return "", err
	}
	args := map[string]string{
		"prop":        "extracts",
		"explaintext": "",
		"titles":      page.Title,
	}
	if numsentence > 0 {
		args["exsentences"] = strconv.Itoa(numsentence)
	} else if numchar > 0 {
		args["exchars"] = strconv.Itoa(numchar)
	} else {
		args["exintro"] = ""
	}

	res, err := utils.WikiRequester(args)
	if err != nil {
		return "", err
	}
	if res.Error.Code != "" {
		return "", errors.New(res.Error.Info)
	}

	return res.Query.Page[strconv.Itoa(page.PageID)].Extract, nil
}
