package NtmHandler

import (
	"encoding/json"
	"github.com/PharbersDeveloper/NtmPods/AuthDaemon"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

type RefreshTokenHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
	au         *AuthDaemon.AuthClient
}

func (h RefreshTokenHandler) NewRefreshTokenHandler(args ...interface{}) RefreshTokenHandler {
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

	return RefreshTokenHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r, au: a}
}

func (h RefreshTokenHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)

	refreshToken := queryForm["refresh_token"][0]
	if len(refreshToken) <= 0 {
		http.Error(w, "refresh_token invalid", http.StatusBadRequest)
		return 1
	}

	token := h.RdGetToken(refreshToken)
	token.Expiry = time.Now()

	config := h.au.ConfigFromURIParameter(r)

	token, _ = config.TokenSource(context.Background(), token).Token()

	h.RdPushToken(token.RefreshToken, token)

	h.RdDeleteToken(refreshToken)

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(token)

	return 0
}

func (h RefreshTokenHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h RefreshTokenHandler) GetHandlerMethod() string {
	return h.Method
}

func (h RefreshTokenHandler) RdGetToken(key string) *oauth2.Token {
	client := h.rd.GetRedisClient()
	defer client.Close()

	result, _ := client.Get(key).Result()
	token := oauth2.Token{}
	json.Unmarshal([]byte(result), &token)

	return &token
}

func (h RefreshTokenHandler) RdPushToken(key string, token *oauth2.Token) int {
	jsonToken, _ := json.Marshal(token)

	client := h.rd.GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Append(key, string(jsonToken))

	pipe.Expire(key, time.Until(token.Expiry))

	pipe.Exec()
	return 0
}

func (h RefreshTokenHandler) RdDeleteToken(key string) int {
	client := h.rd.GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Del(key)

	pipe.Exec()

	return 0
}