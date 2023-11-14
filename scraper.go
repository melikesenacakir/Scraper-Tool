package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"github.com/mbndr/figlet4go"
)

func GetValue(GetUrl string, dateset bool, descset bool) {

	res, err := http.Get(GetUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if GetUrl == "https://thehackernews.com" {
		doc.Find(" .home-right ").Each(func(i int, s *goquery.Selection) {
			title := s.Find("h2").Text()
			color.Magenta("%d. Haber: ", i+1)
			fmt.Printf("%s\n", title)
			if !descset {
				description := s.Find(".home-desc").Text()
				color.Cyan("Açıklama: ")
				fmt.Printf("%s\n", description)
			}
			if !dateset {
				date := s.Find(".item-label .h-datetime").Text()
				color.HiRed("Tarih: ")
				fmt.Printf("%s\n", date)
			}
			fmt.Print("\n")

		})
	} else if GetUrl == "https://www.developer-tech.com/" {
		doc.Find(".main article section").Each(func(i int, s *goquery.Selection) {
			title := s.Find("header h3 a").Text()
			color.Magenta("%d. Haber: ", i+1)
			fmt.Printf("%s\n", title)
			if !descset {
				description := s.Find(".grid-x .cell p").Text()
				color.Cyan("Açıklama: ")
				fmt.Printf("%s\n", description)
			}
			if !dateset {
				date := s.Find(".grid-x .cell .byline .content").Text()
				datev := strings.Index(date, "|")
				date = strings.TrimSpace(date[:datev])
				color.HiRed("Tarih: ")
				fmt.Printf("%s\n", date)
			}
			fmt.Print("\n")

		})

	} else if GetUrl == "https://cybersecuritynews.com/" {
		doc.Find(".td_module_10 .item-details").Each(func(i int, s *goquery.Selection) {
			title := s.Find("h3 a").Text()
			color.Magenta("%d. Haber: ", i+1)
			fmt.Printf("%s\n", title)
			if !descset {
				description := s.Find(".td-excerpt").Text()
				color.Cyan("Açıklama: ")
				description = strings.TrimSpace(description)
				fmt.Printf("%s\n", description)
			}
			if !dateset {
				date := s.Find(".td-post-date time").Text()
				color.HiRed("Tarih: ")
				fmt.Printf("%s\n", date)
			}
			fmt.Print("\n")

		})
	}

}

func main() {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorCyan,
		figlet4go.ColorMagenta,
	}
	options.FontName = "larry3d"
	renderStr, _ := ascii.RenderOpts("Yavuzlar", options)
	fmt.Println(renderStr)
	renderStr2, _ := ascii.RenderOpts("WS TOOL", options)
	fmt.Println(renderStr2)

	first := flag.Bool("1", false, "displays the first news site")
	second := flag.Bool("2", false, "displays the second news site")
	third := flag.Bool("3", false, "displays the third news site")
	date := flag.Bool("date", false, "filters the date part")
	description := flag.Bool("description", false, "filters the description part")
	flag.Parse()
	wg := sync.WaitGroup{}

	firsturl := "https://thehackernews.com"
	secondurl := "https://www.developer-tech.com/"
	thirdurl := "https://cybersecuritynews.com/"

	if *first {
		wg.Add(1)
		go func() {
			defer fmt.Println("----------------thehackernews website news----------------\n\n")
			defer wg.Done()
			GetValue(firsturl, *date, *description)
		}()
		if *second {
			wg.Add(1)
			go func() {
				defer fmt.Println("--------------developer-tech website news----------------\n\n")
				defer wg.Done()
				GetValue(secondurl, *date, *description)
			}()
		}
		if *third {
			wg.Add(1)
			go func() {
				defer fmt.Println("--------------cybersecuritynews website news----------------\n\n")
				defer wg.Done()
				GetValue(thirdurl, *date, *description)
			}()
		}
		wg.Wait()
	} else if *second {
		wg.Add(1)
		go func() {
			defer fmt.Println("--------------developer-tech website news----------------\n\n")
			defer wg.Done()
			GetValue(secondurl, *date, *description)
		}()
		if *first {
			wg.Add(1)
			go func() {
				defer fmt.Println("----------------thehackernews website news----------------\n\n")
				defer wg.Done()
				GetValue(firsturl, *date, *description)
			}()
		}
		if *third {
			wg.Add(1)
			go func() {
				defer fmt.Println("--------------cybersecuritynews website news----------------\n\n")
				defer wg.Done()
				GetValue(thirdurl, *date, *description)
			}()
		}
		wg.Wait()
	} else if *third {
		wg.Add(1)
		go func() {
			defer fmt.Println("--------------cybersecuritynews website news----------------\n\n")
			defer wg.Done()
			GetValue(thirdurl, *date, *description)
		}()
		if *first {
			wg.Add(1)
			go func() {
				defer fmt.Println("----------------thehackernews website news----------------\n\n")
				defer wg.Done()
				GetValue(firsturl, *date, *description)
			}()
		}
		if *second {
			wg.Add(1)
			go func() {
				defer fmt.Println("--------------developer-tech website news----------------\n\n")
				defer wg.Done()
				GetValue(secondurl, *date, *description)
			}()
		}
		wg.Wait()
	}
}
