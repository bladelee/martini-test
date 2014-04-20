package main

import (
	"encoding/json"
	"fmt"
	//	"io/ioutil"
)

var file []byte = []byte(`{
  "author": [{
    "id": 123,
    "email": "attila@attilaolah.eu"
  },
  {
    "id": 456,
    "email": "zym@126.eu"
  }],
  "title":  "My Blog",
  "url":    "http://attilaolah.eu"
}`)

type Record struct {
	Author Author `json:"author"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}

type Author []struct {
	ID    float64 `json:"id"`
	Email string  `json:"email"`
}

func testmain() {
	//    var fileName = "F:\\test.json"

	//    file, err := ioutil.ReadFile(fileName)
	template := CreateSampleTemplate("firsttemplate")
	body, err := json.Marshal(template)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to marshal vmlist")
	} else {
		fmt.Println("content =", body)
	}

	/*
		var parsedMap Record
		err := json.Unmarshal(file, &parsedMap)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(parsedMap)
		}*/
}
