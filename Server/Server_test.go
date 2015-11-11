package Server

import (
	"github.com/sclevine/agouti"
	"testing"
	"gopkg.in/mgo.v2"
	"log"
	. "github.com/smartystreets/goconvey/convey"
)


func TestSpec(t *testing.T) {
	Convey("UserRegister", t, func() {
		var agoutiDriver *agouti.WebDriver
		var page *agouti.Page

		session, _ := mgo.Dial("localhost")

		session.DB("TESTGoNuts").DropDatabase()

		agoutiDriver = agouti.PhantomJS()

		So(agoutiDriver.Start(), ShouldBeNil)

		log.Println("Starting")

		go StartMyApp(3232, "TESTGoNuts")

		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
		So(err, ShouldBeNil)

		Convey("User Registration page", func() {
			Convey("when the user registration is reached", func() {
				Convey("should see the page", func() {
					So(page.Navigate("http://localhost:3232"), ShouldBeNil)

					url, _ := page.URL()
					So(url, ShouldEqual, "http://localhost:3232/")
				})
			})
		})


		Reset(func() {
			So(page.Destroy(), ShouldBeNil)
			So(agoutiDriver.Stop(), ShouldBeNil)
		})
	})
}