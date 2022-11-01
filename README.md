# Go-wiki

<img align="right" width="159px" src="https://upload.wikimedia.org/wikipedia/en/8/80/Wikipedia-logo-v2.svg">

[![go report card](https://goreportcard.com/badge/github.com/trietmn/go-wiki "go report card")](https://goreportcard.com/report/github.com/trietmn/go-wiki)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/trietmn/go-wiki?status.svg)](https://pkg.go.dev/github.com/trietmn/go-wiki)

This is a Golang Wikipedia API wrapper - The Golang module that makes it easy to access and parse data from Wikipedia. You can use this module to crawl data for your data warehouse or use it for your "Know-it-all" AI Chatbot.

## Contents

- [Go-wiki](#go-wiki)
  - [Contents](#contents)
  - [Instalation](#instalation)
  - [Documentation](#documentation)
  - [Quick start](#quick-start)
  - [Functions Examples](#functions-examples)
    - [1. Search](#1-search)
    - [2. GetPage (There are multiple page methods in the Wikipedia Page Methods)](#2-getpage-there-are-multiple-page-methods-in-the-wikipedia-page-methods)
    - [3. Suggest](#3-suggest)
    - [4. Geosearch](#4-geosearch)
    - [5. GetRandom](#5-getrandom)
    - [6. Summary](#6-summary)
  - [Wikipedia Page Methods](#wikipedia-page-methods)
  - [License](#license)
  - [About me](#about-me)
  - [Credit](#credit)

## Instalation

To install Go-Wiki package, you need to install Go and set your Go workspace first.
1. You first need [Go](https://golang.org/) installed, then you can use the below Go command to install Go-wiki.
```sh
go get -u github.com/trietmn/go-wiki
```
2. Import it in your code.
```go
import "github.com/trietmn/go-wiki"
```

## Documentation

You can read the documentation at: <https://pkg.go.dev/github.com/trietmn/go-wiki> 

I will update a full tutorial article on some popular blog as soon as possiple.

## Quick start

```sh
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import (
    "fmt"
    "github.com/trietmn/go-wiki"
)

func main() {
    // Search for the Wikipedia page title
    search_result, _, err := gowiki.Search("Why is the sky blue", 3, false)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("This is your search result: %v\n", search_result)

    // Get the page
    page, err := gowiki.GetPage("Rafael Nadal", -1, false, true)
    if err != nil {
        fmt.Println(err)
    }

    // Get the content of the page
    content, err := page.GetContent()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("This is the page content: %v\n", content)
}
```

```sh
# run example.go
$ go run example.go
```

## Functions Examples

Note: The functions below are functions that you would usually use. Read the [document](https://pkg.go.dev/github.com/trietmn/go-wiki) to see all the functions.

### 1. Search
```go
search_result, suggestion, err := gowiki.Search("Why is the sky blue", 3, true)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("Search result: %v\n", search_result)
fmt.Printf("Suggestion: %v\n", suggestion)
```

### 2. GetPage (There are multiple page methods in the [Wikipedia Page Methods](#wikipedia-page-methods))
```go
page, err := gowiki.GetPage("Rafael Nadal", -1, false, true)
if err != nil {
    fmt.Println(err)
}
// Then now you can use the page methods
```

### 3. Suggest
```go
suggestion, err := gowiki.Suggest("nadal")
if err != nil {
    fmt.Println(err)
}
fmt.Printf("Suggestion: %v\n", suggestion)
```

### 4. Geosearch
```go
res, err := gowiki.GeoSearch(40.67693, 117.23193, -1, "", -1)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("Geosearch result: %v\n", res)
```

### 5. GetRandom
```go
res, err := gowiki.GetRandom(5)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("Random titles: %v\n", res)
```

### 6. Summary
```go
res, err := gowiki.Summary("Rafael Nadal", 5, -1, false, true)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("Summary: %v\n", res)
```

## Wikipedia Page Methods

| Methods        | Description                                          | Example                    |
| -------------- | :--------------------------------------------------- | :------------------------- |
| Equal          | Check if 2 pages are equal to each other             | page1.Equal(page2)         |
| GetContent     | Get the page content                                 | page.GetContent()          |
| GetHTML        | Get the page HTML                                    | page.GetHTML()             |
| GetRevisionID  | Get revid field of a page                            | page.GetRevisionID()       |
| GetParentID    | Get parentid field of a page                         | page.GetParentID()         |
| GetSummary     | Get the summary of the page                          | page.GetSummary()          |
| GetImagesURL   | Get all of the image URL appear in the page          | page.GetImageURL()         |
| GetCoordinate  | Get the page coordinate if exist                     | page.GetCoordinate()       |
| GetReference   | Get all of the extenal links in the page             | page.GetReference()        |
| GetLink        | Get all the titles of Wikipedia page links on a page | page.GetLink()             |
| GetCategory    | Get all of the categories of a page                  | page.GetCategory()         |
| GetSectionList | Get all of the sections of the page                  | page.GetSectionList()      |
| GetSection     | Get the content of a specific section in the page    | page.GetSection("History") |

## License

MIT licensed. See the LICENSE file for full details.

## About me

Connect with me on Linkedin: <https://www.linkedin.com/in/triet-m-nguyen-a94b4b20b/>

## Credit

- [Python Wikipedia](https://github.com/goldsmith/Wikipedia) by goldsmith as inspiration.
- The [Wikimedia Foundation](https://wikimediafoundation.org/) for giving the world free access to data.
