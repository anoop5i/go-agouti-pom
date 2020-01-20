package main

import (
	"github.com/anoopsivarajan/go-agouti-pom/pages"
	"github.com/sclevine/agouti"
	"log"
	"time"
)

func main() {
	// driver := agouti.ChromeDriver(agouti.ChromeOptions("args", []string{"--headless"}))
	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open page:", err)
	}
	window, err := page.Session().GetWindow()
	window.SetSize(1090, 768)
	page.Session().SetImplicitWait(30000)

	if err := page.Navigate("https://google.com"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}

	se := pages.Search(page)

	se.SearchWith("Golang")

	// all := page.AllByXPath("//h3")
	// elements, err := all.Selection.Elements()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, item := range elements {
	// 	log.Println(item.GetText())
	// }

	time.Sleep(2 * time.Second)

	if err := driver.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}

}
