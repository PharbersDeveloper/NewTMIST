package NtmHandler

import (
	"fmt"
	"github.com/PharbersDeveloper/NtmPods/AuthDaemon"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
)

type NtmUserAgentHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
	au		   *AuthDaemon.AuthClient
}

func (h NtmUserAgentHandler) NewUserAgentHandler(args ...interface{}) NtmUserAgentHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var a *AuthDaemon.AuthClient
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
				if tm.Name() == "AuthClient" {
					a = dm.(*AuthDaemon.AuthClient)
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

	return NtmUserAgentHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r, au: a}
}

func (h NtmUserAgentHandler) GenerateUserAgent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {

	config := h.au.ConfigFromURIParameter(r)
	// 数据用ClientID生成State
	hexStr := fmt.Sprintf("%x", config.ClientID)
	fmt.Println(hexStr)

	u := config.AuthCodeURL("xyz")
	http.Redirect(w, r, u, http.StatusFound)
	return 0
}

func (h NtmUserAgentHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h NtmUserAgentHandler) GetHandlerMethod() string {
	return h.Method
}

