package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//key:=e020bd6a
//api search : http://www.omdbapi.com/?apikey=[yourkey]&
// api image :http://img.omdbapi.com/?apikey=[yourkey]&

// type FilmOutput struct {
// 	Title       string "json:Title"
// 	Year        string "json:Year"
// 	Genre       string "json:Genre"
// 	Poster      string "json:Poster"
// 	Description string "json:Plot"
// 	Actors      string "json:Actors"
// }

type FilmOutput struct {
	Title  string
	Year   string
	Genre  string
	Actors string
	Poster string
	Plot   string
}

func DataApi() FilmOutput {
	url := "http://www.omdbapi.com/?apikey=e020bd6a&t=game"
	response, err := http.Get(url)

	var responseObject FilmOutput
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
