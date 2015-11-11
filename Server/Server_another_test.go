package Server

import (
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/FanszHub/test-site/Models"
	"encoding/json"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	//"gopkg.in/mgo.v2"
	"log"
)

func TestStuff(t *testing.T) {
	Convey("UserRegister", t, func() {

		//session, _ := mgo.Dial("localhost")

		//session.DB("TESTGoNuts").DropDatabase()

		log.Println("Starting")

		go StartMyApp(3333, "TESTGoNuts")

		Convey("Testing", func() {
			userJson := `{"username": "dennis"}`

			reader := strings.NewReader(userJson) //Convert string to reader

			request, _ := http.NewRequest("POST", "http://localhost:3333/users", reader) //Create request with JSON body
			request.Header.Set("Content-Type", "application/json")

			res, _ := http.DefaultClient.Do(request)

			So(res.StatusCode, ShouldEqual, 200)

			request, _ = http.NewRequest("GET", "http://localhost:3333/users", nil) //Create request with JSON body

			res, _ = http.DefaultClient.Do(request)

			So(res.StatusCode, ShouldEqual, 200)

			content, _ := ioutil.ReadAll(res.Body)

			var users []Models.User

			json.Unmarshal(content, &users)

			So(len(users), ShouldEqual, 1)
			So(users[0].Username, ShouldEqual, "dennis")

		})
	})
}