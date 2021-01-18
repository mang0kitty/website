package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/mang0kitty/website/src/helpers"
	"github.com/mang0kitty/website/src/models"
)

func SearchBooks() {

	r, err := os.OpenFile("store/books.json", os.O_RDWR, os.ModePerm)
	helpers.CheckError(err)
	fmt.Println("Successfully Opened books.json")

	defer r.Close()

	var books []*models.Book

	helpers.CheckError(json.NewDecoder(r).Decode(&books))

	for _, book := range books {
		grb, err := GoodReads(book.ISBN)
		if err != nil {
			fmt.Println("Failed to get info for ISBN", book.ISBN)
			log.Println(err)
			continue
		}
		fmt.Printf(" Receieved book from GoodReads API %+v\n", grb)

		book.Description = grb.Description
		book.Rating = grb.AverageRating
	}

	_, err = r.Seek(0, os.SEEK_SET)
	helpers.CheckError(err)
	helpers.CheckError(r.Truncate(0))
	helpers.CheckError(json.NewEncoder(r).Encode(&books))
}

func GoodReads(query string) (*XMLBook, error) {
	str := "https://www.goodreads.com/book/isbn/" + query
	u, _ := url.Parse(str)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("key", os.Getenv("GOODREADS_API_KEY"))
	u.RawQuery = q.Encode()

	if xmlStream, err := FetchXML(u.String()); err != nil {
		return nil, err
	} else {
		defer xmlStream.Close()
		var decoder = xml.NewDecoder(xmlStream)
		decoder.Strict = false
		var result XMLBook
		if err := decoder.Decode(&result); err != nil {
			return nil, err
		} else {
			return &result, nil
		}
	}
}
