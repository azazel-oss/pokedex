package pokedex

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/azazel-oss/pokedex/internal/pokecache"
)

type locationBodyJson struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

func RunPokedex(urlToVisit string) locationBodyJson {
	res, err := http.Get(urlToVisit)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	response := locationBodyJson{}
	json.Unmarshal(body, &response)
	return response
}

func GetNextLocations(cache pokecache.Cache, url string) locationBodyJson {
	if value, ok := cache.Get(url); ok {
		response := locationBodyJson{}
		json.Unmarshal(value, &response)
		return response
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	response := locationBodyJson{}
	json.Unmarshal(body, &response)
	cache.Add(url, body)
	return response
}

func GetPreviousLocations(cache pokecache.Cache, url string) locationBodyJson {
	if value, ok := cache.Get(url); ok {
		response := locationBodyJson{}
		json.Unmarshal(value, &response)
		return response
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	response := locationBodyJson{}
	json.Unmarshal(body, &response)
	cache.Add(url, body)
	return response
}
