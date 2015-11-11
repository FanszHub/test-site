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
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
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
	})
}