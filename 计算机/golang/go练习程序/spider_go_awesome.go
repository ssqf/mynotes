package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"util"

	"github.com/PuerkitoBio/goquery"
)

var goAwesome = "https://raw.githubusercontent.com/avelino/awesome-go/master/README.md"
var saveName = "./go-awesome.md"
var newFileName = "./go-star-awesome.md"
var itmeRegExp = `\*\s+\[.+]\((http[\S]+)\).+`

func getFile(url string, outfile string, perm os.FileMode) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(outfile, data, perm)
	return err
}

func getGitHubStarAndLastTime(url string) (int, time.Time, error) {
	//log.Printf("getGitHubStarAndLastTime:%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, time.Now(), err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("status code error: %d %s", resp.StatusCode, resp.Status)
		return 0, time.Now(), fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("goquery error: %v", err)
		return 0, time.Now(), fmt.Errorf("goquery error: %v", err)
	}

	starNum := doc.Find("a.social-count.js-social-count").Text()
	num, err := strconv.Atoi(strings.TrimSpace(starNum))
	//fmt.Printf("star num:%d\n", num)

	lastTime, isExist := doc.Find(".commit-tease.js-details-container.Details").Find("relative-time").Attr("datetime")
	if isExist {
		ltime, err := time.Parse(time.RFC3339, lastTime)
		if err != nil {
			return num, time.Now(), err
		}
		return num, ltime, nil
	}

	lastTimeSrc, isExist := doc.Find("include-fragment.commit-tease.commit-loader").Attr("src")
	if !isExist {
		log.Printf("goquery not found lastTime src\n")
		return 0, time.Now(), fmt.Errorf("goquery not found lastTime src")
	}

	respLast, err := http.Get("https://github.com" + lastTimeSrc)
	if err != nil {
		return 0, time.Now(), err
	}
	defer respLast.Body.Close()
	if respLast.StatusCode != 200 {
		log.Printf("status code error: %d %s", respLast.StatusCode, respLast.Status)
		return 0, time.Now(), fmt.Errorf("status code error: %d %s", respLast.StatusCode, respLast.Status)
	}

	doc, err = goquery.NewDocumentFromReader(respLast.Body)
	if err != nil {
		log.Printf("goquery error: %v", err)
		return 0, time.Now(), fmt.Errorf("goquery error: %v", err)
	}

	lastTime, isExist = doc.Find("relative-time").Attr("datetime")
	if !isExist {
		log.Printf("goquery not found lastTime\n")
		return 0, time.Now(), fmt.Errorf("goquery not found lastTime")
	}
	//fmt.Printf("lastTime:%s\n", lastTime)
	ltime, err := time.Parse(time.RFC3339, lastTime)
	if err != nil {
		return num, time.Now(), err
	}
	return num, ltime, nil
}

func main() {
	log.SetFlags(log.Lshortfile)
	// err := getFile(goAwesome, saveName, 0666)
	// if err != nil {
	// 	log.Fatalf("get go-awesome error: % v ", err)
	// }
	f, err := os.Open(saveName)
	if err != nil {
		log.Fatalf("open %s error: % v ", saveName, err)
	}
	defer f.Close()

	reg, err := regexp.Compile(itmeRegExp)
	if err != nil {
		log.Fatalf("regexp compile %s error: % v ", itmeRegExp, err)
	}

	nf, err := os.OpenFile(newFileName, os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("open %s error: % v ", saveName, err)
	}

	defer nf.Close()
	bread := bufio.NewReader(f)
	line, err := util.ReadLine(bread)
	for err == nil {
		//fmt.Println(line)
		if reg.MatchString(line) {
			url := reg.FindStringSubmatch(line)[1]
			starNum, lastTime, err := getGitHubStarAndLastTime(url)
			if err != nil {
				log.Printf("%s get starNum last time err:%s\n", url, err)
			} else {
				str := strings.SplitAfterN(line, ")", 2)
				line = fmt.Sprintf("%s%s%d%s  *%s*", str[0], ":star:", starNum, str[1], lastTime.Format(time.RFC3339))
			}
		}
		io.WriteString(nf, line+"\n")
		line, err = util.ReadLine(bread)
	}
	//log.Printf("err:%v\n", err)
}
