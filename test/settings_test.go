package test

import (
	"testing"

	"github.com/trietmn/go-wiki"
	"github.com/trietmn/go-wiki/cache"
	"github.com/trietmn/go-wiki/utils"
)

func TestSetLanguage(t *testing.T) {
	new_lang := "fr"
	gowiki.SetLanguage(new_lang)
	if new_lang != utils.WikiLanguage {
		t.Errorf("got %v, expect %v", utils.WikiLanguage, new_lang)
	}
	new_lang = "en"
	gowiki.SetLanguage(new_lang)
	if new_lang != utils.WikiLanguage {
		t.Errorf("got %v, expect %v", utils.WikiLanguage, new_lang)
	}
}

func TestSetUserAgent(t *testing.T) {
	new_user := "testing"
	gowiki.SetUserAgent(new_user)
	if new_user != utils.UserAgent {
		t.Errorf("got %v, expect %v", utils.UserAgent, new_user)
	}
}

func TestCache(t *testing.T) {
	// utils.WikiRequester = utils.RequestWikiApi
	old := utils.Cache.GetLen()
	_, _, err := gowiki.Search("Porsche", 3, false)
	if err != nil {
		t.Errorf("%v", err)
	}
	if utils.Cache.GetLen() <= old && old < cache.MaxCacheMemory {
		t.Errorf("expect request got added to the cache")
	}
}
