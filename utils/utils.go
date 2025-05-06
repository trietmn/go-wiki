package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/trietmn/go-wiki/cache"
	"github.com/trietmn/go-wiki/models"
)

const (
	ReqPerSec = 199
	ApiGap    = time.Second / ReqPerSec
)

var (
	UserAgent        string          = "go-wiki"
	WikiLanguage     string          = "en"
	WikiURL          string          = "http://%v.wikipedia.org/w/api.php"
	LastCall         time.Time       = time.Now()
	Cache            cache.WikiCache = cache.MakeWikiCache()
	WikiRequester                    = RequestWikiApi
	WikiRequesterRaw                 = RequestWikiApiRaw
)

func TurnSliceOfString(s []interface{}) []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = v.(string)
	}
	return res
}

/*
Return true if 2 slices are the same
*/
func CompareSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !Isin(b, a[i]) {
			return false
		}
	}
	return true
}

/*
Return true if string s is in list
*/
func Isin(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

/*
Update the last time we call the API (API should)
*/
func UpdateLastCall(now time.Time) {
	LastCall = now
}

/*
Make a request to the Wikipedia API using the given search parameters.

Returns a RequestResult (You can see the model in the models.go file)
*/
func RequestWikiApi(args map[string]string) (models.RequestResult, error) {
	resultBody, err := RequestWikiApiRaw(args)
	if err != nil {
		return models.RequestResult{}, err
	}

	// Parse
	var result models.RequestResult
	err = json.Unmarshal([]byte(resultBody), &result)
	if err != nil {
		return models.RequestResult{}, err
	}
	return result, nil
}

/*GenerateWikiApiRequest Generates a request object to query the Wikipedia API */
func GenerateWikiApiRequest(args map[string]string) (*http.Request, error) {
	url := fmt.Sprintf(WikiURL, WikiLanguage)
	// Make new request object
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Add header
	request.Header.Set("User-Agent", UserAgent)
	q := request.URL.Query()
	// Add parameters
	if args["format"] == "" {
		args["format"] = "json"
	}
	if args["action"] == "" {
		args["action"] = "query"
	}
	for k, v := range args {
		q.Add(k, v)
	}
	request.URL.RawQuery = q.Encode()
	now := time.Now()
	if now.Sub(LastCall) < ApiGap {
		wait := LastCall.Add(ApiGap).Sub(now)
		time.Sleep(wait)
	}
	return request, nil
}

/*RequestWikiApiRaw Queries the Wikipedia API without parsing */
func RequestWikiApiRaw(args map[string]string) ([]byte, error) {
	request, err := GenerateWikiApiRequest(args)
	now := time.Now()
	if err != nil {
		return []byte{}, nil
	}

	// Check in cache
	fullUrl := request.URL.String()
	r, err := Cache.Get(fullUrl)
	if err == nil {
		return r, nil
	}

	// Make GET request
	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(request)
	defer UpdateLastCall(now)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return []byte{}, errors.New("unable to fetch the results")
	}

	// Read body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	// Add to cache
	Cache.Add(fullUrl, body)

	return body, nil
}

/*
Make a deep copy of a map[string]string
*/
func CopyMap(source map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range source {
		res[k] = v
	}
	return res
}

/*
Update map a using map b
*/
func UpdateMap(a map[string]string, b map[string]interface{}) {
	for k, v := range b {
		switch t := v.(type) {
		case int:
			a[k] = strconv.Itoa(t)
		case string:
			a[k] = t
		}
	}
}

func HelpAddURL(s string) string {
	if s[0:4] == "http" {
		return s
	}
	return "http:" + s
}
