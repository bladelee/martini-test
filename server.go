package main

import (
	"github.com/go-martini/martini"
	//   "github.com/codegangsta/martini-contrib/binding"
	"fmt"
	"github.com/martini-contrib/render"
	"io"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/hello", HandleHello)

	m.Post("/template", PostTemplate2)
	m.Get("/templates", GetTemplates)
	m.Get("/template/:id", GetTemplate)

	m.Run()
}

func HandleHello(r render.Render) {
	r.JSON(200, map[string]interface{}{"name": "Hello World"})
	// r.JSON(200, `{"name": "Hello World"}`)
}

func PostTemplate(r render.Render) {
	r.JSON(200, nil)
}

func GetTemplate(r render.Render) {

}

func GetTemplates(r render.Render) {
	r.JSON(200, map[string]interface{}{"name": "Hello World"})
	// r.JSON(200, `{"name": "Hello World"}`)
}

func PostTemplate2(w http.ResponseWriter, req *http.Request) (int, string) {

	if req.Body != nil {
		defer req.Body.Close()
	}

	t, tmplcontent, err := ParserJson(req.Body)

	if err != nil && err != io.EOF {
		//		errors.Overall[DeserializationError] = err.Error()
		fmt.Println(err)
	} else {
		ts := CreateTemplateResource(tmplcontent, &t)
		id, err := db.Add(ts)
		switch err {
		case ErrAlreadyExists:
			// Duplicate
			return http.StatusConflict, fmt.Sprintf("the template '%s' name '%s' already exists", ts.Id, t.Name)
		case nil:
			// TODO : Location is expected to be an absolute URI, as per the RFC2616
			w.Header().Set("Location", fmt.Sprintf("/templates/%s", id))
			return http.StatusCreated, "OK"
		default:
			panic(err)
		}

		fmt.Println("template =", t)
	}

	return 0, err.Error()

	/*al := getPostAlbum(r)
	id, err := db.Add(al)
	switch err {
	case ErrAlreadyExists:
		// Duplicate
		return http.StatusConflict, Must(enc.Encode(
			NewError(ErrCodeAlreadyExists, fmt.Sprintf("the template '%s' from '%s' already exists", al.Title, al.Band))))
	case nil:
		// TODO : Location is expected to be an absolute URI, as per the RFC2616
		w.Header().Set("Location", fmt.Sprintf("/albums/%d", id))
		return http.StatusCreated, Must(enc.Encode(al))
	default:
		panic(err)
	}*/
}
