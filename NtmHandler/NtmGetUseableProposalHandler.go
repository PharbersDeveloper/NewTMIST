package NtmHandler

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/julienschmidt/httprouter"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"strings"
	"fmt"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"gopkg.in/mgo.v2/bson"
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
	w.Header().Add("Content-Type", "application/json")

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

	redisDriver := h.rd.GetRedisClient()
	uid, _ := redisDriver.HGet(token+"_info", "uid").Result()
	fmt.Println(uid)

	var out NtmModel.UseableProposal
	cond := bson.M{"account-id": uid}
	err = h.db.FindOneByCondition(&out, &out, cond)

	//response := map[string]interface{}{
	//	"status": "",
	//	"result": nil,
	//	"error":  nil,
	//}

	//if err == nil && out.ID != "" {
	//	hex := md5.New()
	//	io.WriteString(hex, out.ID)
	//	out.Password = ""
	//	token := fmt.Sprintf("%x", hex.Sum(nil))
	//	err = h.rd.PushToken(token, time.Hour*24*365)
	//	out.Token = token
	//
	//	response["status"] = "ok"
	//	response["result"] = out
	//	response["error"] = err
	//
	//	//reval, _ := json.Marshal(response)
	//	enc := json.NewEncoder(w)
	//	enc.Encode(response)
	//	return 0
	//} else {
	//	response["status"] = "error"
	//	response["error"] = "账户或密码错误！"
	//	enc := json.NewEncoder(w)
	//	enc.Encode(response)
	//	return 1
	//}

	data := []string{
		"北京市",
	}
	enc := json.NewEncoder(w)
	enc.Encode(data)
	return 0
}

func (h NtmGetUseableProposalsHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h NtmGetUseableProposalsHandler) GetHandlerMethod() string {
	return h.Method
}
