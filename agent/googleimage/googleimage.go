package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sync"

	"../../models/image"

	"github.com/pkg/errors"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/googleapi/transport"
)

func main() {

	sortAllFlag := flag.Bool("all", false, "Get default sort images")
	flag.Parse()

	resultNum := int64(10)
	nextStart := int64(1)
	maxPage := 10
	wg := new(sync.WaitGroup)
	for i := 0; i < maxPage; i++ {
		wg.Add(1)
		go getGoogleImages(resultNum, nextStart, *sortAllFlag, wg)
		nextStart += resultNum
	}
	wg.Wait()
}

func getGoogleImages(resultNum int64, nextStart int64, sortAllFlag bool, wg *sync.WaitGroup) error {
	defer wg.Done()
	developerKey := os.Getenv("DEVELOPER_KEY")
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	customsearchService, err := customsearch.New(client)
	if err != nil {
		return errors.Wrap(err, "client error")
	}
	searchEngineID := os.Getenv("SEARCH_ENGINE_ID")
	customsearchQuery := customsearchService.Cse.List("佐倉綾音").Cx(searchEngineID).SearchType("image").Num(resultNum)
	if !sortAllFlag {
		// sortについてはこの辺
		// https://developers.google.com/custom-search/docs/structured_search#sort_by_attribute
		customsearchQuery = customsearchQuery.Sort("date")
	}
	results, err := customsearchQuery.Start(nextStart).Do()
	if err != nil {
		return errors.Wrap(err, "customsearch error")
	}
	for _, item := range results.Items {
		log.Println(item.Link)
		img := image.NewImage(0, item.Link)
		err := img.Save()
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
