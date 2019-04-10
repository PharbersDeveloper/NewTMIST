package NtmHandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/NtmPods/AuthDaemon"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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

	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	version := strings.Split(r.URL.Path, "/")[1]
	resource := fmt.Sprint("192.168.100.116:9096", "/"+version+"/", "GenerateAccessToken", "?", r.URL.RawQuery)
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
	data := map[string]interface{}{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println("AccessToken Error")
	}
	enc := json.NewEncoder(w)
	enc.Encode(data)
	return 0
}

func (h AuthHandler) PasswordLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	//body, err := ioutil.ReadAll(r.Body)
	//return 0

	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	version := strings.Split(r.URL.Path, "/")[1]
	resource := fmt.Sprint("192.168.100.116:9096", "/"+version+"/", "PasswordLogin", "?", r.URL.RawQuery)
	mergeURL := strings.Join([]string{scheme, resource}, "")

	body, err := ioutil.ReadAll(r.Body)

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("POST", mergeURL, bytes.NewBuffer(body))
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Fuck Error")
	}
	result, err := ioutil.ReadAll(response.Body)
	data := map[string]interface{}{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println("PasswordAccessToken Error")
	}
	enc := json.NewEncoder(w)
	enc.Encode(data)
	return 0
}

func (h AuthHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	version := strings.Split(r.URL.Path, "/")[1]
	resource := fmt.Sprint("192.168.100.116:9096", "/"+version+"/", "RefreshAccessToken", "?", r.URL.RawQuery)
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
	data := map[string]interface{}{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		fmt.Println("RefreshAccessToken Error")
	}
	enc := json.NewEncoder(w)
	enc.Encode(data)
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

	//pipe
	pipe.Append(key, string(jsonToken))

	_, err := pipe.Exec()
	return err
}

func (h AuthHandler) RdGetValueByKey(key string) (*string, error){
	client := h.rd.GetRedisClient()
	defer client.Close()

	result, err := client.Get(key).Result()

	if err != nil {
		return nil ,err
	}
	return &result, nil
}
