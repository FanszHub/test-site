package Api

import (
	"net/http/httptest"
	"fmt"
	"testing"
	"net/http"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/FanszHub/test-site/Models"
	"io/ioutil"
)

var (
	server *httptest.Server
	usersUrl string
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Registration")
}

type mockDB struct{}

func (mdb *mockDB) AllUsers() ([]*Models.User, error) {
	users := make([]*Models.User, 0)
	users = append(users, &Models.User{"Emma"})
	users = append(users, &Models.User{"Steve"})
	return users, nil
}

var _ = Describe("UserRegister", func() {

	BeforeEach(func() {

		env := &Env{Db: &mockDB{}}

		server = httptest.NewServer(Handlers(env))

		usersUrl = fmt.Sprintf("%s/users", server.URL)
	})

	Describe("User registration api", func() {

		Context("Get users", func() {
			It("should find a user", func() {

				request, err := http.NewRequest("GET", usersUrl, nil)

				res, err := http.DefaultClient.Do(request)

				Expect(err).ToNot(HaveOccurred())
				Expect(res.StatusCode).To(Equal(200))

				content, _ := ioutil.ReadAll(res.Body)

				Expect(string(content)).To(Equal("Emma"))
			})
		})
	})
})