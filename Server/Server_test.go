package Server

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sclevine/agouti/matchers"
	"github.com/sclevine/agouti"
	"testing"
	"gopkg.in/mgo.v2"
	"log"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = Describe("UserRegister", func() {
	var page *agouti.Page

	BeforeSuite(func() {

		session, _ := mgo.Dial("localhost")

		session.DB("TESTGoNuts").DropDatabase()

		agoutiDriver = agouti.PhantomJS()

		Expect(agoutiDriver.Start()).To(Succeed())

		log.Println("Starting")

		go StartMyApp(3232, "TESTGoNuts")

		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		Expect(page.Destroy()).To(Succeed())
		Expect(agoutiDriver.Stop()).To(Succeed())
	})

	Describe("User Registration page", func() {
		Context("when the user registration is reached",func() {
			It("should see the page", func() {
				Expect(page.Navigate("http://localhost:3232")).To(Succeed())

				Expect(page).To(HaveURL("http://localhost:3232/"))
			})
		})
	})
})
