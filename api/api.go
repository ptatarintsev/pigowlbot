package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	GetPacksUrl = "http://pigowl.com:8080/getPacks"
	GetPacksStatUrl = "http://pigowl.com:8090/getPacksStat"
)

type PackPhrase struct {
    Phrase      string         `json:"phrase"`
    Complexity  float32        `json:"complexity"`
    Description string         `json:"description"`
    Reviews     map[string]int `json:"reviews"`
}

type Pack struct {
    ID          int          `json:"id"`
    Language    string       `json:"language"`
    Name        string       `json:"name"`
    Description string       `json:"description"`
    Phrases     []PackPhrase `json:"phrases"`
    Version     int          `json:"version"`
    Paid        bool         `json:"paid"`
}

type PackResponse struct {
    Pack  Pack `json:"pack"`
    Count int  `json:"count"`
}

type PackStatResponse struct {
    Timestamp int64  `json:"timestamp"`
    ID int  `json:"id"`
}

type GetPacksResponse struct {
    Packs  []PackResponse `json:"packs"`
}

type GetPacksStatResponse struct {
    PacksStat  []PackStatResponse
}

func GetPackages() *GetPacksResponse {
    res, err := http.Get(GetPacksUrl)
    if err != nil {
	log.Fatal(err)
    }
    defer res.Body.Close()

    response := new(GetPacksResponse) 
    json.NewDecoder(res.Body).Decode(response)

    return response
}

func GetPackagesStatistics() *GetPacksStatResponse {
    res, err := http.Get(GetPacksStatUrl)
    if err != nil {
	log.Fatal(err)
    }
    defer res.Body.Close()

    result := new(GetPacksStatResponse)
    json.NewDecoder(res.Body).Decode(&result.PacksStat)

    return result
}

