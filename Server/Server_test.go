package Server

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sclevine/agouti/matchers"
	"github.com/sclevine/agouti"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = Describe("UserRegister", func() {
	var page *agouti.Page

	BeforeEach(func() {
		agoutiDriver = agouti.PhantomJS()

		Expect(agoutiDriver.Start()).To(Succeed())

		go StartMyApp(3131)

		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	Describe("User Registration page", func() {
		Context("when the user registration is reached",func() {
			It("should see the page", func() {
				Expect(page.Navigate("http://localhost:3131")).To(Succeed())

				Expect(page).To(HaveURL("http://localhost:3131/"))
			})
		})
	})
})
