package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

type jsonRsp struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func main() {
	//target := &url.URL{Scheme: "http", Host: "www.google.com"}
	target, _ := url.Parse(getEnv("BACKEND_URL", "http://localhost:8080"))
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = rewriteBdy

	http.Handle("/", proxy)
	log.Fatal(http.ListenAndServe(":8088", nil))

}

//func reverseBytes(a []byte) []byte {
//	for i := len(a)/2-1; i >= 0; i-- {
//		opp := len(a)-1-i
//		a[i], a[opp] = a[opp], a[i]
//	}
//	return a
//}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func reverseMessage(m []byte) []byte {
	//var dat map[string]interface{}
	var j jsonRsp
	log.Println("income string:" + string(m))
	if err := json.Unmarshal(m, &j); err != nil {
		panic(err)
	}

	j = jsonRsp{j.Id, reverseString(j.Message)}
	outJson, err := json.Marshal(j)
	if err != nil {
		return nil
	}
	log.Println("reverse string:" + string(outJson))
	return outJson

}
func rewriteBdy(r *http.Response) (err error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = r.Body.Close()
	if err != nil {
		return err
	}
	b = reverseMessage(b)
	body := ioutil.NopCloser(bytes.NewReader(b))
	r.Body = body
	r.ContentLength = int64(len(b))
	r.Header.Set("Content-Length", strconv.Itoa(len(b)))
	return nil
}
