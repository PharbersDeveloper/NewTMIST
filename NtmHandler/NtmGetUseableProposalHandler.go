package NtmHandler

import (
	"net/http"
	"reflect"
	"strings"
	"io/ioutil"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/julienschmidt/httprouter"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
)

type NtmGetUseableProposalsHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h NtmGetUseableProposalsHandler) NewGetUseableProposalsHandler(args ...interface{}) NtmGetUseableProposalsHandler {
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

	return NtmGetUseableProposalsHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r}
}

func (h NtmGetUseableProposalsHandler) GetUseableProposals(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	// 验证token
	auth := r.Header.Get("Authorization")
	arr := strings.Split(auth, " ")
	if len(arr) < 2 || arr[0] != "bearer" {
		panic("Auth Failed!")
	}
	token := arr[1]
	err := h.rd.CheckToken(token)
	if err != nil {
		panic(err.Error())
	}

	// 取出uid
	redisDriver := h.rd.GetRedisClient()
	defer redisDriver.Close()
	uid, _ := redisDriver.HGet(token+"_info", "uid").Result()

	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	toUrl := strings.Replace(r.URL.Path, "GetUseableProposals", h.Args[0], -1) + "?account-id=" + uid
	useableProposalURL := strings.Join([]string{scheme, r.Host, toUrl}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", useableProposalURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, _ := client.Do(req)

	body, err := ioutil.ReadAll(response.Body)
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
	return 0
}

func (h NtmGetUseableProposalsHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h NtmGetUseableProposalsHandler) GetHandlerMethod() string {
	return h.Method
}
