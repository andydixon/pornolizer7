package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"pornolizer-go/engine"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//go:embed html
var homepage embed.FS

type apiResponse struct {
	Pornolised string `json:"pornolised"`
}

type reqObject struct {
	Text string `json:"text"`
	Sweariness int `json:"sweariness"`
}

func main() {
	serverRoot, _ := fs.Sub(homepage, "html")

	http.HandleFunc("/pornolize/", engineHandler)
	http.HandleFunc("/translate/", engineHandler)
	http.HandleFunc("/api/", apiHandler)
    http.Handle("/", http.FileServer(http.FS(serverRoot)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
           log.Printf("Error reading body: %v", err)
           http.Error(w, "can't read body", http.StatusBadRequest)
           return
        }
	var request reqObject
	err = json.Unmarshal(body,&request)
	if err != nil {
		response := &apiResponse{Pornolised: "error"}
		responseJson, _ := json.Marshal(response)
		fmt.Fprint(w, string(responseJson))
	} else {
		language := "en"
		response := &apiResponse{Pornolised: engine.Pornolise(request.Text, request.Sweariness, language, 1337)}
		responseJson, _ := json.Marshal(response)
		fmt.Fprint(w, string(responseJson))
	}

}

func engineHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["url"]
	language := "en"
	
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param is missing")
		return
	}
	lang, ok := r.URL.Query()["lang"]

	if ok && len(lang[0]) < 1 {
		language = lang[0]
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
	proxyRequest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36 dxn/pz7")
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
		s.SetText(engine.Pornolise(s.Text(), 80, language, i))
	})
	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		s.SetText(engine.Pornolise(s.Text(), 70, language, i))

	})
	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		s.SetText(engine.Pornolise(s.Text(), 60, language, i))

	})
	doc.Find("h4").Each(func(i int, s *goquery.Selection) {
		s.SetText(engine.Pornolise(s.Text(), 50, language, i))

	})
	doc.Find("h5").Each(func(i int, s *goquery.Selection) {
		s.SetText(engine.Pornolise(s.Text(), 50, language, i))

	})
	doc.Find("h6").Each(func(i int, s *goquery.Selection) {
		s.SetText(engine.Pornolise(s.Text(), 50, language, i))

	})
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		s.SetText(engine.Pornolise(s.Text(), 30, language, i))

	})
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		s.SetAttr("rel", "noreferrer nofollow")
		href, _ := s.Attr("href")
		if strings.HasPrefix(href, "/") {
			s.SetAttr("href", "http://"+r.Host+"/pornolize?lang="+language+"&url="+resp.Request.URL.Scheme+"://"+resp.Request.URL.Host+href)
		} else {
			s.SetAttr("href", "http://"+r.Host+"/pornolize?lang="+language+"&url="+href)
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
	doc.Find("head").PrependHtml("\n\n<!-- Converted by The Pornoliser (c)2021 Andy Dixon - pornolize.com / andydixon.com -->\n\n");
	
	returnContent, err := doc.Html()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, returnContent)
	defer resp.Body.Close()
}
