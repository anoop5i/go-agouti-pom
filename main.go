// package main

// import (
// 	"github.com/sclevine/agouti"
// 	"log"
// )

// func main() {
// 	driver := agouti.ChromeDriver()

// 	if err := driver.Start(); err != nil {
// 		log.Fatal("Failed to start driver:", err)
// 	}

// 	page, err := driver.NewPage()
// 	if err != nil {
// 		log.Fatal("Failed to open page:", err)
// 	}
// 	page.Session().SetImplicitWait(30000)

// 	if err := page.Navigate("https://agouti.org/"); err != nil {
// 		log.Fatal("Failed to navigate:", err)
// 	}

// 	sectionTitle, err := page.FindByID("getting-agouti").Text()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(sectionTitle)

// 	page.FindByID("getting-agouti").Click()

// 	if err := driver.Stop(); err != nil {
// 		log.Fatal("Failed to close pages and stop WebDriver:", err)
// 	}

// }
