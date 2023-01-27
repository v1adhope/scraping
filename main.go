package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
)

const uri = "https://skysmart.ru/articles/english/populyarnye-anglijskie-slova-s-perevodom"

type word struct {
	en string
	ru string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("skysmart.ru", "www.skysmart.ru"),
	)

	words := make([]word, 0, 904)

	c.OnHTML("tbody > tr", func(e *colly.HTMLElement) {
		selector := e.DOM

		en := selector.Find("td:nth-child(2) > p").Text()
		ru := selector.Find("td:nth-child(4) > p").Text()

		words = append(words, word{
			en: en,
			ru: ru,
		})

		// func() {
		// 	siteCount := selector.Find("td:nth-child(1) > p").Text()
		//
		// 	siteIntCount, _ := strconv.Atoi(siteCount)
		//
		// 	if siteIntCount != len(words) {
		// 		fmt.Println(siteCount)
		// 	}
		// }()
	})

	c.Visit(uri)

	file := excelize.NewFile()
	defer file.Close()

	file.SetSheetName("Sheet1", "Info")
	file.SetCellValue("Info", "A1", "Source:")
	file.SetCellValue("Info", "B1", uri)
	file.SetCellValue("Info", "A2", "Word count:")
	file.SetCellValue("Info", "B2", len(words))

	if _, err := file.NewSheet("Words"); err != nil {
		fmt.Println(err)
	}

	for k, v := range words {
		file.SetCellValue("Words", fmt.Sprintf("A%d", k+1), v.en)
		file.SetCellValue("Words", fmt.Sprintf("B%d", k+1), v.ru)
	}

	if err := file.SaveAs("popular-en-ru-words.xlsx"); err != nil {
		fmt.Println(err)
	}
}
