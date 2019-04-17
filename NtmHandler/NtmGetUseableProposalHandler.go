package NtmHandler

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
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

	//token, err := NtmMiddleware.NtmCheckToken.CheckTokenFormFunction(w, r)
	//if err != nil {
	//	panic(fmt.Sprintf(err.Error()))
	//	return 1
	//}

	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	toUrl := strings.Replace(r.URL.Path, "GetUseableProposals", h.Args[0], -1) + "?account-id=5c4552455ee2dd7c36a94a9e" //+ token.UserID
	useableProposalURL := strings.Join([]string{scheme, r.Host, toUrl}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", useableProposalURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)
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
