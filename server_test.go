package main_test

import (
	"encoding/json"

	"bytes"
	"github.com/martini-contrib/binding"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	. "mytest"
)

var _ = Describe("Mytest", func() {

	var (
		body []byte
		err  error
		//todos []Todo
	)

	Context(" Hello World Json Test", func() {
		It("returns a 200 Status Code", func() {
			Request("GET", "/templates", HandleHello)
			log.Println(response)
			log.Println(response.Body)
			Expect(response.Code).To(Equal(200))
			Expect(response.Body).To(MatchJSON(`{"name": "Hello World"}`))

		})
	})

	Context("提交模板", func() {

		BeforeEach(func() {
			template := CreateSampleTemplate("firsttemplate")
			body, err = json.Marshal(template)
			if err != nil {
				log.Println(err)
				log.Println("Unable to marshal vmlist")
			} else {
				log.Println("content =", body)
			}
		})

		It("returns a 200 Status Code", func() {
			PostRequest("POST", "/template", binding.Json(Template{}), PostTemplate, bytes.NewReader(body))
			Expect(response.Code).To(Equal(200))
		})

		It("returns a 200 Status Code", func() {
			PostRequestCommon("POST", "/template", PostTemplate2, bytes.NewReader(body))
			Expect(response.Code).To(Equal(201))
		})
	})

})
