package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func ParserJson(r io.Reader) (t Template, content string, err error) {
	err = json.NewDecoder(r).Decode(&t)
	content = "To Do"
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
	return t, content, err
}

/*
func Json(jsonStruct interface{}, ifacePtr ...interface{}) martini.Handler {
	return func(context martini.Context, req *http.Request) {
		ensureNotPointer(jsonStruct)
		jsonStruct := reflect.New(reflect.TypeOf(jsonStruct))
		errors := newErrors()

		if req.Body != nil {
			defer req.Body.Close()
		}

		if err := json.NewDecoder(req.Body).Decode(jsonStruct.Interface()); err != nil && err != io.EOF {
			errors.Overall[DeserializationError] = err.Error()
		}

		validateAndMap(jsonStruct, context, errors, ifacePtr...)
	}
}
*/
