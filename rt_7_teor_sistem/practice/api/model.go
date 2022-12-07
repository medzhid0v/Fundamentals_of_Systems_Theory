package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetData() Film {
	resp, err := http.Get("http://www.omdbapi.com/?i=tt3896198&apikey=722bcb5b")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}


	var data Film

	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		log.Fatal(err)
	}
	
	return data




}