package Api

import (
	"net/http/httptest"
	"fmt"
	"testing"
	"net/http"
	"github.com/FanszHub/test-site/Models"
	"encoding/json"
	"io/ioutil"
	"github.com/FanszHub/test-site/Env"
	. "github.com/smartystreets/goconvey/convey"
	"bytes"
)

var (
	server *httptest.Server
	usersUrl string
	users []*Models.User
)

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

func TestSpec(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {

		var env = Env.Env{UserRepository: &mockDB{}}

		server = httptest.NewServer(Routes(env))

		usersUrl = fmt.Sprintf("%s/users", server.URL)

		Convey("User registration api", func() {
			Convey("Get users", func() {
				Convey("should find a user", func() {

					request, err := http.NewRequest("GET", usersUrl, nil)

					res, err := http.DefaultClient.Do(request)

					So(err, ShouldBeNil)
					So(res.StatusCode, ShouldEqual, 200)

					content, _ := ioutil.ReadAll(res.Body)

					var users []Models.User

					json.Unmarshal(content, &users)

					So(len(users), ShouldEqual, 2)
					So(users[0].Username, ShouldEqual, "Emma")
					So(users[1].Username, ShouldEqual, "Steve")
				})
			})
		})

		Convey("Post users", func() {
			Convey("should save a user", func() {

				var user Models.User
				user.Username = "Thing"

				json, _ := json.Marshal(user)

				var body = bytes.NewReader(json)

				request, err := http.NewRequest("POST", usersUrl, body)
				request.Header.Set("Content-Type", "application/json")

				res, err := http.DefaultClient.Do(request)

				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)
				So(len(users), ShouldEqual, 1)
				So(users[0].Username, ShouldEqual, "Thing")
			})
		})

	})
}