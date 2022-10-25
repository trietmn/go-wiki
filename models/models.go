package models

type RequestError struct {
	Code  string `json:"code"`
	Info  string `json:"info"`
	Aster string `json:"*"`
}

type InnerBasic struct {
	Aster string `json:"*"`
}

type InnerSearchInfo struct {
	TotalHits         int    `json:"totalhits"`
	Suggestion        string `json:"suggestion"`
	SuggestionSnippet string `json:"suggestionsnippet"`
}

type InnerSearch struct {
	Ns        int    `json:"ns"`
	Title     string `json:"title"`
	PageID    int    `json:"pageid"`
	Size      int    `json:"size"`
	Wordcount int    `json:"wordcount"`
	Snippet   string `json:"snippet"`
	Timestamp string `json:"timestamp"`
}

type InnerPage struct {
	Ns                  int                      `json:"ns"`
	Title               string                   `json:"title"`
	PageID              int                      `json:"pageid"`
	ContentModel        string                   `json:"contentmodel"`
	PageLanguage        string                   `json:"pagelanguage"`
	PageLanguageTmlCode string                   `json:"pagelanguagetmlcode"`
	PageLanguageDir     string                   `json:"pagelanguagedir"`
	Touched             string                   `json:"touched"`
	LastRevid           int                      `json:"lastrevid"`
	Length              int                      `json:"length"`
	FullURL             string                   `json:"fullurl"`
	EditURL             string                   `json:"editurl"`
	CanonicalURL        string                   `json:"canonicalurl"`
	PageProps           map[string]string        `json:"pageprops"`
	Missing             string                   `json:"missing"`
	Extract             string                   `json:"extract"`
	Revision            []map[string]interface{} `json:"revisions"`
	Extlink             []map[string]string      `json:"extlinks"`
	Link                []map[string]interface{} `json:"links"`
	Category            []map[string]interface{} `json:"categories"`
	ImageInfo           []map[string]string      `json:"imageinfo"`
	Coordinate          []map[string]interface{} `json:"coordinates"`
}

type InnerGeoSearch struct {
	PageID    int     `json:"pageid"`
	Ns        int     `json:"ns"`
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lom"`
	Distance  float32 `json:"dist"`
	Primary   string  `json:"primary"`
}

type InnerNormalize struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type RequestQuery struct {
	SearchInfo InnerSearchInfo      `json:"searchinfo"`
	Normalize  []InnerNormalize     `json:"normalized"`
	Redirect   []InnerNormalize     `json:"redirects"`
	Search     []InnerSearch        `json:"search"`
	GeoSearch  []InnerGeoSearch     `json:"geosearch"`
	Page       map[string]InnerPage `json:"pages"`
	Random     []InnerSearch        `json:"random"`
	Language   []map[string]string  `json:"languages"`
}

/*
	The result of calling the Wikipedia API
*/
type RequestResult struct {
	Error         RequestError           `json:"error"`
	Warning       map[string]InnerBasic  `json:"warnings"`
	Batchcomplete string                 `json:"batchcomplete"`
	Query         RequestQuery           `json:"query"`
	Servedby      string                 `json:"servedby"`
	Continue      map[string]interface{} `json:"continue"`
	Parse         map[string]interface{} `json:"parse"`
}
