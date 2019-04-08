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

	config := h.au.ConfigFromURIParameter(r)


	token := &oauth2.Token{}
	tokenResult, err := h.RdGetValueByKey("RefreshToken_" + refreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}
	json.Unmarshal([]byte(*tokenResult), &token)

	token.Expiry = time.Now()

	token, err = config.TokenSource(context.Background(), token).Token()

	tokenUUID, _ := h.RdGetValueByKey(token.AccessToken)
	result, _ := h.RdGetValueByKey(*tokenUUID)

	var oauthObject map[string]interface{}
	json.Unmarshal([]byte(*result), &oauthObject)

	phToken := AuthDaemon.PhToken{
		Scope: token.Extra("scope").(string),
		AccountID: oauthObject["UserID"].(string),
	}
	phToken.AccessToken = token.AccessToken
	phToken.RefreshToken = token.RefreshToken
	phToken.Expiry = token.Expiry
	phToken.TokenType = token.TokenType

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 1
	}

	h.RdPushToken("RefreshToken_" + phToken.RefreshToken, &phToken)

	defer h.RdDeleteToken("RefreshToken_" + refreshToken)

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(phToken)

	return 0
}

func (h RefreshTokenHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h RefreshTokenHandler) GetHandlerMethod() string {
	return h.Method
}

func (h RefreshTokenHandler) RdGetValueByKey(key string) (*string, error){
	client := h.rd.GetRedisClient()
	defer client.Close()

	result, err := client.Get(key).Result()

	if err != nil {
		return nil ,err
	}
	return &result, nil
}

func (h RefreshTokenHandler) RdPushToken(key string, token *AuthDaemon.PhToken) error {
	jsonToken, _ := json.Marshal(token)

	client := h.rd.GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	//pipe
	pipe.Append(key, string(jsonToken))

	_, err := pipe.Exec()
	return err
}

func (h RefreshTokenHandler) RdDeleteToken(key string) int {
	client := h.rd.GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Del(key)

	pipe.Exec()

	return 0
}