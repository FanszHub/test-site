package Server

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
	"net/http"
)

var _ = Describe("UserRegister", func() {

	It("string", func(){

		go StartMyApp(3131)

		userJson := `{"username": "dennis"}`

		reader := strings.NewReader(userJson) //Convert string to reader

		request, _ := http.NewRequest("POST", "http://localhost:3131/users", reader) //Create request with JSON body

		res, _ := http.DefaultClient.Do(request)

		Expect(res.StatusCode).To(Equal(200))
	})
})
