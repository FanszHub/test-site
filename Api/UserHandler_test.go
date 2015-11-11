package Api

import (
	"net/http/httptest"
	"fmt"
	"testing"
	"net/http"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/FanszHub/test-site/Models"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/FanszHub/test-site/Env"
)

var (
	server *httptest.Server
	usersUrl string
	users []*Models.User
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Registration")
}

type mockDB struct{}

func (mdb *mockDB) AllUsers() ([]*Models.User, error) {
	users = make([]*Models.User, 0)
	users = append(users, &Models.User{"Emma"})
	users = append(users, &Models.User{"Steve"})

	return users, nil
}

func (mdb *mockDB) AddUser(user *Models.User) (error) {
	users = make([]*Models.User, 0)
	users = append(users, user)

	return nil
}

var _ = Describe("UserRegister", func() {

	BeforeEach(func() {

		env := Env.Env{UserRepository: &mockDB{}}

		server = httptest.NewServer(Routes(env))

		usersUrl = fmt.Sprintf("%s/users", server.URL)
	})

	AfterEach(func(){
		server.Close()
		server = nil
	})


	Describe("User registration api", func() {
		Context("Get users", func() {
			It("should find a user", func() {

				request, err := http.NewRequest("GET", usersUrl, nil)

				res, err := http.DefaultClient.Do(request)

				Expect(err).ToNot(HaveOccurred())
				Expect(res.StatusCode).To(Equal(200))

				content, _ := ioutil.ReadAll(res.Body)

				var users []Models.User

				json.Unmarshal(content,&users)

				Expect(len(users)).To(Equal(2))
				Expect(users[0].Username).To(Equal("Emma"))
				Expect(users[1].Username).To(Equal("Steve"))
			})

		})

		Context("Post users", func() {
			It("should save a user", func() {

				var user Models.User
				user.Username = "Thing"

				json, _ := json.Marshal(user)

				var body = bytes.NewReader(json)

				request, err := http.NewRequest("POST", usersUrl, body)
				request.Header.Set("Content-Type", "application/json")

				res, err := http.DefaultClient.Do(request)

				Expect(err).ToNot(HaveOccurred())
				Expect(res.StatusCode).To(Equal(200))

				Expect(len(users)).To(Equal(1))
				Expect(users[0].Username).To(Equal("Thing"))
			})
		})

	})
})