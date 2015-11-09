package Server

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/FanszHub/test-site/Models"
	"encoding/json"
)

var _ = Describe("UserRegister", func() {

	It("string", func(){

		go StartMyApp(3131)

		userJson := `{"username": "dennis"}`

		reader := strings.NewReader(userJson) //Convert string to reader

		request, _ := http.NewRequest("POST", "http://localhost:3131/users", reader) //Create request with JSON body

		res, _ := http.DefaultClient.Do(request)

		Expect(res.StatusCode).To(Equal(200))

		request, _ = http.NewRequest("GET", "http://localhost:3131/users", nil) //Create request with JSON body

		res, _ = http.DefaultClient.Do(request)

		Expect(res.StatusCode).To(Equal(200))

		content, _ := ioutil.ReadAll(res.Body)

		var users []Models.User

		json.Unmarshal(content,&users)

		Expect(len(users)).To(Equal(1))
		Expect(users[0].Username).To(Equal("dennis"))

	})
})
