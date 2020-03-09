package main

import (
	"Pornolizer7/pornoliser"
	"Pornolizer7/support"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, support.DefaultHomepage())
}

func engineHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["url"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param is missing")
		return
	}
	targetUrl := keys[0]
	fmt.Printf("Target is %s", targetUrl)

	// parse the url
	destURL, _ := url.Parse(targetUrl)
	params := url.Values{}
	//params.Add(inputName, randomdata.RandStringRunes(rand.Intn(4)+8))
	body := strings.NewReader(params.Encode())
	proxyRequest, _ := http.NewRequest(r.Method, destURL.String(), body)
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	proxyRequest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36 dxnpw/pz7")
	resp, err := http.DefaultClient.Do(proxyRequest)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		s.SetText(pornoliser.Pornolise(s.Text(), 80))
	})
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		s.SetText(pornoliser.Pornolise(s.Text(), 70))

	})
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		s.SetText(pornoliser.Pornolise(s.Text(), 60))

	})
	doc.Find("h4").Each(func(i int, s *goquery.Selection) {
		s.SetText(pornoliser.Pornolise(s.Text(), 50))

	})
	doc.Find("h5").Each(func(i int, s *goquery.Selection) {
		s.SetText(pornoliser.Pornolise(s.Text(), 50))

	})
	doc.Find("h6").Each(func(i int, s *goquery.Selection) {
		s.SetText(pornoliser.Pornolise(s.Text(), 50))

	})
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		s.SetText(pornoliser.Pornolise(s.Text(), 30))

	})
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("rel", "noreferrer nofollow")
		href, _ := s.Attr("href")
		if strings.HasPrefix(href, "/") {
			s.SetAttr("href", "http://"+r.Host+"/pornolize?url="+resp.Request.URL.Scheme+"://"+resp.Request.URL.Host+href)
		} else {
			s.SetAttr("href", "http://"+r.Host+"/pornolize?url="+href)
		}
	})

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		source, _ := s.Attr("src")
		if strings.HasPrefix(source, "/") && !strings.HasPrefix(source, "//") {
			s.SetAttr("src", resp.Request.URL.Scheme+"://"+resp.Request.URL.Host+source)
		}
		sourceSet, _ := s.Attr("srcset")
		if strings.HasPrefix(source, "/") && !strings.HasPrefix(sourceSet, "//") {
			s.SetAttr("srcset", resp.Request.URL.Scheme+"://"+resp.Request.URL.Host+sourceSet)
		}

	})
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		source, _ := s.Attr("href")
		if strings.HasPrefix(source, "/") && !strings.HasPrefix(source, "//") {
			s.SetAttr("href", resp.Request.URL.Scheme+"://"+resp.Request.URL.Host+source)
		}

	})

	doc.Find("head").AppendHtml("<base href='" + resp.Request.URL.Scheme + "://" + resp.Request.URL.Host + "'>")
	returnContent, err := doc.Html()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, returnContent)
	defer resp.Body.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/pornolize/", engineHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
