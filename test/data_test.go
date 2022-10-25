package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/trietmn/go-wiki/models"
	"github.com/trietmn/go-wiki/utils"
)

const (
	mockDataPath        string = "./mock_data.json"
	mockWikiRequestPath string = "./mock_wiki_request.json"
)

var (
	MockData        map[string]interface{}          = MakeMockData(mockDataPath)
	MockWikiRequest map[string]models.RequestResult = MakeMockWikiRequest(mockWikiRequestPath)
	MockRequestMap  map[string]map[string]string    = MakeMockRequestMap(MockWikiRequest)
)

func MakeMockData(filepath string) map[string]interface{} {
	var res map[string]interface{}
	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return res
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &res)

	return res
}

func MakeMockWikiRequest(filepath string) map[string]models.RequestResult {
	var res map[string]models.RequestResult
	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return res
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(byteValue), &res)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func MakeMockRequestMap(mockrequest map[string]models.RequestResult) map[string]map[string]string {
	res := make(map[string]map[string]string)
	for key, _ := range mockrequest {
		args := strings.Split(key, ";")
		temp := make(map[string]string)
		for _, arg := range args {
			kv := strings.Split(arg, ":")
			temp[kv[0]] = kv[1]
		}
		res[key] = temp
	}
	return res
}

func MockRequester(args map[string]string) (models.RequestResult, error) {
OuterLoop:
	for key, value := range MockRequestMap {
		for k, v := range args {
			if k == "action" && v == "query" {
				continue
			}
			if _, ok := value[k]; !ok {
				continue OuterLoop
			}
			if value[k] != v {
				continue OuterLoop
			}
		}
		utils.Cache.Add(key, MockWikiRequest[key])
		return MockWikiRequest[key], nil
	}
	return models.RequestResult{}, errors.New("mock request not exist")
}
