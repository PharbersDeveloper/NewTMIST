package NtmMiddleware

import (
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/manyminds/api2go"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type NtmCheckTokenMiddleware struct {
	Args []string
	rd   *BmRedis.BmRedis
}

type result struct {
	AllScope	string	`json:"all_scope"`
	ClientID	string	`json:"client_id"`
	Expires		float64	`json:"expires_in"`
	RefreshExpires float64 `json:"refresh_expires_in"`
	Scope 		string 	`json:"scope"`
	UserID 		string 	`json:"user_id"`
}

func (ctm NtmCheckTokenMiddleware) NewCheckTokenMiddleware(args ...interface{}) NtmCheckTokenMiddleware {
	var r *BmRedis.BmRedis
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				}
			}
		} else if i == 1 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}

	return NtmCheckTokenMiddleware{Args: ag, rd: r}
}

func (ctm NtmCheckTokenMiddleware) DoMiddleware(c api2go.APIContexter, w http.ResponseWriter, r *http.Request) {
	if _, err := CheckTokenFormFunction(w, r); err != nil {
		panic(err.Error())
	}
}

// TODO @Alex这块需要重构
func CheckTokenFormFunction(w http.ResponseWriter, r *http.Request) (result, error) {
	w.Header().Add("Content-Type", "application/json")

	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}

	resource := fmt.Sprint("localhost:9096/", "v0/", "TokenValidation")

	mergeURL := strings.Join([]string{scheme, resource}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("POST", mergeURL, nil)

	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, err := client.Do(req)

	body, err := ioutil.ReadAll(response.Body)

	//var temp map[string]interface{}
	temp := result{}
	err = json.Unmarshal(body, &temp)
	//phToken := AuthDaemon.PhToken{Scope: temp["scope"].(string)}
	//phToken.RefreshToken = temp["refresh_token"].(string)
	////phToken.Expiry =  temp["expiry_id"].(string)
	//phToken.TokenType = temp["token_type"].(string)

	return temp, err

}