package NtmHandler

import (
	"encoding/json"
	"github.com/PharbersDeveloper/NtmPods/AuthDaemon"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

type AuthHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
	au		   *AuthDaemon.AuthClient
}

func (h AuthHandler) NewAuthHandler(args ...interface{}) AuthHandler {
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

	return AuthHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r, au: a}
}

func (h AuthHandler) GenerateAccessToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	queryForm, _ := url.ParseQuery(r.URL.RawQuery)

	// 替AuthServer过滤掉一部分减轻压力
	state := queryForm["state"][0]
	if len(state) <= 0 {
		http.Error(w, "State invalid", http.StatusBadRequest)
		return 1
	}
	code := queryForm["code"][0]
	if len(code) <= 0 {
		http.Error(w, "code invalid", http.StatusBadRequest)
		return 1
	}

	clientId := queryForm["client_id"][0]
	if len(clientId) <= 0 {
		http.Error(w, "client_id invalid", http.StatusBadRequest)
		return 1
	}

	clientSecret := queryForm["client_secret"][0]
	if len(clientSecret) <= 0 {
		http.Error(w, "client_secret invalid", http.StatusBadRequest)
		return 1
	}

	config := h.au.ConfigFromURIParameter(r)

	token, err := config.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}

	scope := token.Extra("scope")

	phToken := AuthDaemon.PhToken{
		Scope: scope.(string),
	}
	phToken.AccessToken = token.AccessToken
	phToken.RefreshToken = token.RefreshToken
	phToken.Expiry = token.Expiry
	phToken.TokenType = token.TokenType

	// 存入Redis RefreshToken
	err = h.RdPushRefreshToken(token.RefreshToken, &phToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err = e.Encode(phToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}

	return 0
}

func (h AuthHandler) PasswordLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	body, err := ioutil.ReadAll(r.Body)

	var parameter map[string]interface{}
	json.Unmarshal(body, &parameter)

	config := h.au.ConfigFromURIParameter(r)
	token, err := config.PasswordCredentialsToken(context.Background(), parameter["username"].(string), parameter["password"].(string))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}

	scope := token.Extra("scope")

	phToken := AuthDaemon.PhToken{
		Scope: scope.(string),
	}
	phToken.AccessToken = token.AccessToken
	phToken.RefreshToken = token.RefreshToken
	phToken.Expiry = token.Expiry
	phToken.TokenType = token.TokenType

	err = h.RdPushRefreshToken(token.RefreshToken, &phToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	err = e.Encode(phToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}

	return 0
}

func (h AuthHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h AuthHandler) GetHandlerMethod() string {
	return h.Method
}

func (h AuthHandler) RdPushRefreshToken(key string, token *AuthDaemon.PhToken) error {
	jsonToken, _ := json.Marshal(token)

	client := h.rd.GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Append(key, string(jsonToken))

	_, err := pipe.Exec()
	return err
}

