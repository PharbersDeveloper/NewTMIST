package NtmHandler

import (
	"encoding/json"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type NtmGeneratePaperHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h NtmGeneratePaperHandler) NewGeneratePaperHandler(args ...interface{}) NtmGeneratePaperHandler {
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

	return NtmGeneratePaperHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r}
}

func (h NtmGeneratePaperHandler) GeneratePaper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")

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

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return 1
	}

	// 临时存储post穿过来的json
	var temp map[string]interface{}
	var inputIds []string
	json.Unmarshal(body, &temp)

	a := temp["input-ids"].([]interface{})

	for _, v := range a {
		value := v.(string)
		inputIds = append(inputIds, value)
	}

	res := NtmModel.Paper{}
	json.Unmarshal(body, &res)
	res.AccountID = uid
	res.InputIDs = inputIds // 由于Paper实体中配制的json:"-"所以要手动设置，这块儿会改的

	oid, _ := h.db.InsertBmObject(&res)

	//拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	toUrl := strings.Replace(r.URL.Path, "GeneratePaper", h.Args[0], -1) + "/" + oid
	useableProposalURL := strings.Join([]string{scheme, r.Host, toUrl}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", useableProposalURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, _ := client.Do(req)

	responseBody, err := ioutil.ReadAll(response.Body)
	w.Header().Add("Content-Type", "application/json")
	w.Write(responseBody)

	return 0
}

func (h NtmGeneratePaperHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h NtmGeneratePaperHandler) GetHandlerMethod() string {
	return h.Method
}