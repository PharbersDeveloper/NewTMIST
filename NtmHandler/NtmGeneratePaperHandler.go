package NtmHandler

import (
		"fmt"
	"github.com/PharbersDeveloper/NtmPods/NtmMiddleware"
	"github.com/PharbersDeveloper/NtmPods/NtmModel"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
		"net/http"
	"reflect"
	"strings"
	"github.com/PharbersDeveloper/NtmPods/NtmDataStorage"
	"time"
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
	mdb := []BmDaemons.BmDaemon{h.db}
	w.Header().Add("Content-Type", "application/json")

	token, err := NtmMiddleware.NtmCheckToken.CheckTokenFormFunction(w, r)
	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	proposalId := r.FormValue("proposal-id")
	proposalModel, err := NtmDataStorage.NtmProposalStorage{}.NewProposalStorage(mdb).GetOne(proposalId)
	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	paperModel := NtmModel.Paper{
		AccountID: token.UserID,
		ProposalID: proposalModel.ID,
		Name: proposalModel.Name,
		Describe: proposalModel.Describe,
		StartTime: time.Now().Unix(),
		EndTime: 0,
		InputState: "未开始",
		InputIDs: proposalModel.InputIDs,
		ReportIDs: proposalModel.ReportIDs,
	}

	paperId := NtmDataStorage.NtmPaperStorage{}.NewPaperStorage(mdb).Insert(paperModel)

	//拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	toUrl := strings.Replace(r.URL.Path, "GeneratePaper", h.Args[0], -1) + "/" + paperId
	paperURL := strings.Join([]string{scheme, r.Host, toUrl}, "")

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("GET", paperURL, nil)
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