package tests

import (
	"github.com/anoopsivarajan/go-agouti-pom/pages"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Google Search Suite")
}

var _ = Describe("Google Search", func() {

	var driver *agouti.WebDriver
	var page *agouti.Page

	BeforeSuite(func() {
		driver = agouti.ChromeDriver()
		Expect(driver.Start()).To(Succeed())
	})

	BeforeEach(func() {
		var err error
		page, err = driver.NewPage()
		Expect(err).NotTo(HaveOccurred())

		Expect(page.Navigate("https://www.google.com")).To(Succeed())
	})

	AfterSuite(func() {
		Expect(driver.Stop()).To(Succeed())
	})

	It("Golang search", func() {
		search := pages.Search(page)
		By("Search with text", func() {
			Eventually(search.SearchInput).Should(BeFound())

			search.SearchWith("Golang")
			Eventually(page.FindByLink("The Go Programming Language")).Should(BeFound())
		})

	})
})
