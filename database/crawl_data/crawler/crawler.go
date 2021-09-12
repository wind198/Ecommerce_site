package crawler

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"example.com/database/todb"
	"github.com/gocolly/colly"
)

type Product struct {
	productCategory    string
	productName        string
	productPrice       float64
	productImageSource string
}

func (p *Product) String() string {
	return fmt.Sprintf("%s,%s,%v", strings.Trim(p.productCategory, " \n"),
		strings.Trim(p.productName, " \n"),
		p.productPrice)
}

var validPriceString = regexp.MustCompile(`(.*)\$(\d+),?(\d*\.\d+)(.*)`)

func Crawl(db *sql.DB) {
	//declare collectors
	c1 := colly.NewCollector(colly.Async(true))
	c2 := colly.NewCollector(colly.Async(true))

	//defind OnHTML behavior of c1, direct c1 what it should look for
	c1.OnHTML("li.desktop-3.tablet-fourth", func(h1 *colly.HTMLElement) {
		log.Println("hello")
		if strings.TrimSpace(h1.ChildText("h4>a")) == "Home Decor" {
			log.Println("hello2")
			h1.ForEach("ul.mega-stack", func(i int, h2 *colly.HTMLElement) {
				log.Println("hello3")
				h2.ForEach("li>a", func(i int, h3 *colly.HTMLElement) {
					log.Println("hello4")
					link := h3.Attr("href")
					c2.Visit(h3.Request.AbsoluteURL(link))
				})
			})
		}
	})
	//c2 will collect all necessary information and insert them to database
	c2.OnHTML("body", func(h *colly.HTMLElement) {
		var category string
		var name string
		var priceText string
		var price float64
		var imageSource string
		category = strings.Trim(h.ChildText("h1.nope"), " \n")
		if category == "" {
			category = h.ChildText("h1:nth-child(1)")
		}
		h.ForEach("div.productindexinner", func(i int, e *colly.HTMLElement) {
			name = e.ChildText("h3")
			priceText = e.ChildText("div.price div.prod-price")

			var err error
			m := validPriceString.FindStringSubmatch(priceText)
			// log.Println(strings.Join(m, ", "))
			if m != nil {
				price, err = strconv.ParseFloat(m[2]+m[3], 64)
				if err != nil {
					log.Printf("Error during extracting price value, %v\n", err)
				}
			}

			imageSource = e.ChildAttr(".prod-image  img", "src")
			aProduct := Product{
				productCategory:    category,
				productName:        name,
				productPrice:       price,
				productImageSource: imageSource,
			}

			rows, err := todb.InsertToDB(db, aProduct.productCategory, aProduct.productName, aProduct.productPrice, aProduct.productImageSource)
			if rows != 0 {
				log.Printf("Insert for product %s suceed, %v affected\n", aProduct.String(), rows)
			} else {
				log.Printf("Insert for product %s failed, %v\n", aProduct.String(), err)
			}
		})

	})
	//find the ul that immediately follow it
	//for each li > a, make new request
	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("C1 visiting...", r.URL.String())
		r.Headers.Add("Accept-Language", "en-US")
	})
	c1.OnError(func(r *colly.Response, e error) {
		log.Println("Error ", e)
	})
	c2.OnError(func(r *colly.Response, e error) {
		log.Println("Error ", e)
	})
	c1.Visit("https://interiorsonline.com.au/")
	c1.Wait()
	c2.Wait()
}
