package NtmHandler

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type NtmUserAgentHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h NtmUserAgentHandler) NewUserAgentHandler(args ...interface{}) NtmUserAgentHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
				if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}

	return NtmUserAgentHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r }
}

func (h NtmUserAgentHandler) GenerateUserAgent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {

	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	version := strings.Split(r.URL.Path, "/")[1]
	resource := fmt.Sprint(h.Args[0], "/"+version+"/", "ThirdParty", "?", r.URL.RawQuery)
	mergeURL := strings.Join([]string{scheme, resource}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", mergeURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Fuck Error")
	}
	result, err := ioutil.ReadAll(response.Body)
	data := map[string]string {}
	json.Unmarshal(result, &data)

	res, tok := data["type"]
	if tok && res == "url" {
		http.Redirect(w, r, data["redirect-uri"], http.StatusFound)
	} else {

	}

	//http.Redirect(w, r, data["redirect-uri"], http.StatusFound)

	return 0
}

func (h NtmUserAgentHandler) AuthUserAgent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	file, err := os.Open(h.Args[0])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return 1
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)
	return 0
}

func (h NtmUserAgentHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h NtmUserAgentHandler) GetHandlerMethod() string {
	return h.Method
}
