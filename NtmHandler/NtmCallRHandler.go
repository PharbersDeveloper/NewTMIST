package NtmHandler

import (
	"Ntm/NtmDaemons/NtmXmpp"
	"Ntm/NtmDataStorage"
	"Ntm/NtmModel"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

var NtmCallR NtmCallRHandler

type NtmCallRHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
	xmpp	   *NtmXmpp.NtmXmpp
	kafka 	   *bmkafka.BmKafkaConfig
}

type calcStruct struct {
	Header     map[string]string      `json:"header"`
	Account    string                 `json:"account"`
	Proposal   string                 `json:"proposal"`
	PaperInput string                 `json:"paperInput"`
	CurrentScenario   map[string]interface{} `json:"currentScenario"`
	Scenarios		[]interface{} `json:"scenarios"`
	Body       map[string]interface{} `json:"body"`
}

type resultStruct struct {
	Header     map[string]string      `json:"header"`
	Account    string                 `json:"account"`
	Proposal   string                 `json:"proposal"`
	PaperInput string                 `json:"paperInput"`
	Scenario   string 				  `json:"scenario"`
	Body       map[string]interface{} `json:"body"`
	Error 	   map[string]interface{} `json:"error"`
}

type assessmentResult struct {
	Level int `json:"level" mapstructure:"level"`
	Code  int `json:"code" mapstructure:"code"`
}

func (h NtmCallRHandler) NewCallRHandler(args ...interface{}) NtmCallRHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var x *NtmXmpp.NtmXmpp
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
				if tm.Name() == "NtmXmpp" {
					x = dm.(*NtmXmpp.NtmXmpp)
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

	kafka, _ := bmkafka.GetConfigInstance()
	NtmCallR = NtmCallRHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r , xmpp: x, kafka: kafka}

	go func() {
		topic := []string{kafka.Topics[len(kafka.Topics) -1]}
		fmt.Println(topic)
		kafka.SubscribeTopics(topic, subscriptionFunc)
	}()

	return NtmCallR
}

func (h NtmCallRHandler) CallRCalculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	mdb := []BmDaemons.BmDaemon{h.db}
	req := getApi2goRequest(r, r.Header)
	w.Header().Add("Content-Type", "application/json")
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	result := map[string]interface{}{}
	enc := json.NewEncoder(w)
	json.Unmarshal(res, &params)

	proposalId, pok := params["proposal-id"]
	accountId, aok := params["account-id"]
	scenarioId, sok := params["scenario-id"]

	if !pok {
		result["status"] = "error"
		result["msg"] = "计算失败，proposal-id参数缺失"
		enc.Encode(result)
		return 1
	} else if !sok {
		result["status"] = "error"
		result["msg"] = "计算失败，scenario-id参数缺失"
		enc.Encode(result)
		return 1
	} else if !aok {
		result["status"] = "error"
		result["msg"] = "计算失败，account-id参数缺失"
		enc.Encode(result)
		return 1
	}

	var (
		inputs []map[string]interface{}
		histories map[string]interface{}
		historyProducts []map[string]interface{}
		historyRepresentatives []map[string]interface{}
		historyManager map[string]interface{}
		historyHospitals []map[string]interface{}
		goodsConfigMapping map[string]string
	)

	histories = make(map[string]interface{})
	goodsConfigMapping = make(map[string]string)
	historyManager = make(map[string]interface{})

	scenarioStorage := NtmDataStorage.NtmScenarioStorage{}.NewScenarioStorage(mdb)
	paperStorage := NtmDataStorage.NtmPaperStorage{}.NewPaperStorage(mdb)
	paperInputStorage := NtmDataStorage.NtmPaperinputStorage{}.NewPaperinputStorage(mdb)
	salesReportStorage := NtmDataStorage.NtmSalesReportStorage{}.NewSalesReportStorage(mdb)
	hospitalSalesReportStorage := NtmDataStorage.NtmHospitalSalesReportStorage{}.NewHospitalSalesReportStorage(mdb)
	destConfigStorage := NtmDataStorage.NtmDestConfigStorage{}.NewDestConfigStorage(mdb)
	hospitalConfigStorage := NtmDataStorage.NtmHospitalConfigStorage{}.NewHospitalConfigStorage(mdb)
	hospitalStorage := NtmDataStorage.NtmHospitalStorage{}.NewHospitalStorage(mdb)
	productSalesReportStorage := NtmDataStorage.NtmProductSalesReportStorage{}.NewProductSalesReportStorage(mdb)
	goodsConfigStorage := NtmDataStorage.NtmGoodsConfigStorage{}.NewGoodsConfigStorage(mdb)
	productConfigStorage := NtmDataStorage.NtmProductConfigStorage{}.NewProductConfigStorage(mdb)
	productStorage := NtmDataStorage.NtmProductStorage{}.NewProductStorage(mdb)
	resourceConfigStorage := NtmDataStorage.NtmResourceConfigStorage{}.NewResourceConfigStorage(mdb)
	representativeConfigStorage := NtmDataStorage.NtmRepresentativeConfigStorage{}.NewRepresentativeConfigStorage(mdb)
	representativeStorage := NtmDataStorage.NtmRepresentativeStorage{}.NewRepresentativeStorage(mdb)
	managerConfigStorage := NtmDataStorage.NtmManagerConfigStorage{}.NewManagerConfigStorage(mdb)
	businessInputStorage := NtmDataStorage.NtmBusinessinputStorage{}.NewBusinessinputStorage(mdb)
	representativeInputStorage := NtmDataStorage.NtmRepresentativeinputStorage{}.NewRepresentativeinputStorage(mdb)
	managerInputStorage := NtmDataStorage.NtmManagerinputStorage{}.NewManagerinputStorage(mdb)
	personnelAssessmentStorage := NtmDataStorage.NtmPersonnelAssessmentStorage{}.NewPersonnelAssessmentStorage(mdb)
	representativeAbilitStorage := NtmDataStorage.NtmRepresentativeAbilityStorage{}.NewRepresentativeAbilityStorage(mdb)
	actionKpiStorage := NtmDataStorage.NtmActionKpiStorage{}.NewActionKpiStorage(mdb)


	// 查询当前的Paper
	req.QueryParams["account-id"] = []string{accountId}
	req.QueryParams["proposal-id"] = []string{proposalId}
	papers := paperStorage.GetAll(req, -1, -1)
	if len(papers) == 0 {
		result["status"] = "error"
		result["msg"] = "计算失败，Paper Is Null"
		enc.Encode(result)
		return 1
	}
	// 当前周期的所有产品
	cleanQueryParams(&req)
	req.QueryParams["scenario-id"] = []string{scenarioId}
	for _, goodsConfig := range goodsConfigStorage.GetAll(req, -1,-1) {
		productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
		product, _ := productStorage.GetOne(productConfig.ProductID)
		goodsConfigMapping[goodsConfig.ID] = product.Name
	}

	paper := papers[0]

	// 最新的输入
	paperInput, _ := paperInputStorage.GetOne(paper.InputIDs[len(paper.InputIDs) - 1])

	cleanQueryParams(&req)

	// 最新的BusinessInputs Inputs
	req.QueryParams["ids"] = paperInput.BusinessinputIDs
	businessInputs := businessInputStorage.GetAll(req, -1,-1)
	req.QueryParams["ids"] = paperInput.RepresentativeinputIDs
	representativeInputs :=  representativeInputStorage.GetAll(req, -1, -1)
	req.QueryParams["ids"] = paperInput.ManagerinputIDs
	managerInputs := managerInputStorage.GetAll(req, -1, -1)

	for _, businessInput := range businessInputs {
		var (
			products []map[string]interface{}
			manager map[string]interface{}
		)

		manager = make(map[string]interface{})

		hospitalMap := map[string]interface{}{}
		representativeMap := map[string]interface{}{}

		destConfig, _ := destConfigStorage.GetOne(businessInput.DestConfigId)
		hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
		hospital, _ := hospitalStorage.GetOne(hospitalConfig.HospitalID)
		resourceConfig, _ := resourceConfigStorage.GetOne(businessInput.ResourceConfigId)
		representativeConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
		representative, _ := representativeStorage.GetOne(representativeConfig.RepresentativeID)


		representativeMap["resource-config-id"] = resourceConfig.ID
		representativeMap["representative-id"] = representative.ID
		representativeMap["representative-name"] = representative.Name
		for _, representativeInput := range representativeInputs {
			if representativeInput.ResourceConfigId == resourceConfig.ID {
				representativeMap["product-knowledge-training"] = representativeInput.ProductKnowledgeTraining
				representativeMap["sales-ability-training"] = representativeInput.SalesAbilityTraining
				representativeMap["region-training"] = representativeInput.RegionTraining
				representativeMap["performance-training"] = representativeInput.PerformanceTraining
				representativeMap["vocational-development"] = representativeInput.VocationalDevelopment
				representativeMap["assist-access-time"] = representativeInput.AssistAccessTime
				representativeMap["ability-coach"] = representativeInput.AbilityCoach
			}
		}


		products = append(products, map[string]interface{}{
			"goods-config-id": businessInput.GoodsConfigId,
			"product-name": goodsConfigMapping[businessInput.GoodsConfigId],
			"sales-target": businessInput.SalesTarget,
			"budget": businessInput.Budget,
			"meeting-places": businessInput.MeetingPlaces,
			"visit-time": businessInput.VisitTime,
		})

		for _, managerInput := range managerInputs {
			manager["strategy-analysis-time"] = managerInput.StrategyAnalysisTime
			manager["admin-work-time"] = managerInput.AdminWorkTime
			manager["client-management-time"] = managerInput.ClientManagementTime
			manager["KPI-analysis-time"] = managerInput.KpiAnalysisTime
			manager["team-meeting-time"] = managerInput.TeamMeetingTime
		}

		hospitalMap["dest-config-id"] = destConfig.ID
		hospitalMap["hospital-name"] = hospital.Name
		hospitalMap["hospital-level"] = hospital.HospitalLevel
		hospitalMap["representative"] = representativeMap
		hospitalMap["products"] = products
		hospitalMap["manager"] = manager
		inputs = append(inputs, hospitalMap)
	}


	// 历史前一个周期
	salesReport, _ := salesReportStorage.GetOne(paper.SalesReportIDs[len(paper.SalesReportIDs) - 1])

	cleanQueryParams(&req)
	req.QueryParams["ids"] = salesReport.HospitalSalesReportIDs
	hospitalSalesReports := hospitalSalesReportStorage.GetAll(req, -1,-1)
	req.QueryParams["ids"] = salesReport.ProductSalesReportIDs
	productSalesReports := productSalesReportStorage.GetAll(req, -1, -1)
	cleanQueryParams(&req)
	req.QueryParams["scenario-id"] = []string{salesReport.ScenarioID}
	resourceConfigs := resourceConfigStorage.GetAll(req, -1, -1)

	cleanQueryParams(&req)
	req.QueryParams["ids"] = paper.PersonnelAssessmentIDs[len(paper.PersonnelAssessmentIDs) - 1 :]
	personnelAssessments := personnelAssessmentStorage.GetAll(req, -1, -1)

	// 历史前一个中期的产品
	cleanQueryParams(&req)
	req.QueryParams["scenario-id"] = []string{salesReport.ScenarioID}

	// hospitals
	for _, hospitalSalesReport := range hospitalSalesReports {
		hospitalMap := map[string]interface{}{}

		destConfig, _ := destConfigStorage.GetOne(hospitalSalesReport.DestConfigID)
		hospitalConfig, _ := hospitalConfigStorage.GetOne(destConfig.DestID)
		hospital, _ := hospitalStorage.GetOne(hospitalConfig.HospitalID)

		resourceConfig, _ := resourceConfigStorage.GetOne(hospitalSalesReport.ResourceConfigID)
		representativeConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
		representative, _ := representativeStorage.GetOne(representativeConfig.RepresentativeID)

		goodsConfig, _ := goodsConfigStorage.GetOne(hospitalSalesReport.GoodsConfigID)
		productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
		product, _ := productStorage.GetOne(productConfig.ProductID)

		hospitalMap["hospital-name"] = hospital.Name
		hospitalMap["representative-name"] = representative.Name
		hospitalMap["product-name"] = product.Name
		hospitalMap["potential"] = hospitalSalesReport.Potential
		hospitalMap["sales"] = hospitalSalesReport.Sales
		hospitalMap["sales-quota"] = hospitalSalesReport.SalesQuota
		hospitalMap["share"] = hospitalSalesReport.Share
		historyHospitals = append(historyHospitals, hospitalMap)
	}

	// products
	for _, productSalesReport := range productSalesReports {
		goodsConfig, _ := goodsConfigStorage.GetOne(productSalesReport.GoodsConfigID)
		productConfig, _ := productConfigStorage.GetOne(goodsConfig.GoodsID)
		product, _ := productStorage.GetOne(productConfig.ProductID)
		historyProducts = append(historyProducts, map[string]interface{}{
			"goods-config-id": goodsConfig.ID,
			"product-name": product.Name,
			"life-cycle": productConfig.LifeCycle,
			"share": productSalesReport.Share,
		})
	}

	// representatives
	cleanQueryParams(&req)
	for _, personnelAssessment := range personnelAssessments {
		req.QueryParams["ids"] = personnelAssessment.RepresentativeAbilityIDs
		representativeAbilits := representativeAbilitStorage.GetAll(req, -1, -1)
		req.QueryParams["ids"] = personnelAssessment.ActionKpiIDs
		actionKpis := actionKpiStorage.GetAll(req, -1, -1)

		repConfigMap := map[string]interface{}{}

		for _, resourceConfig := range resourceConfigs {
			repConfig, _ := representativeConfigStorage.GetOne(resourceConfig.ResourceID)
			repConfigMap[repConfig.RepresentativeID] = repConfig.TotalTime
		}

		for _, representativeAbilit := range representativeAbilits {
			for _, actionKpi := range actionKpis {
				if representativeAbilit.RepresentativeID == actionKpi.RepresentativeID {
					representative, _ := representativeStorage.GetOne(representativeAbilit.RepresentativeID)
					historyRepresentatives = append(historyRepresentatives, map[string]interface{}{
						"representative-name": representative.Name,
						"product-knowledge": representativeAbilit.ProductKnowledge,
						"sales-ability": representativeAbilit.SalesAbility,
						"regional-management-ability": representativeAbilit.RegionalManagementAbility,
						"job-enthusiasm": representativeAbilit.JobEnthusiasm,
						"behavior-validity": representativeAbilit.BehaviorValidity,
						"total-time": repConfigMap[representativeAbilit.RepresentativeID],
						"target-number": actionKpi.TargetNumber,
						"target-coverage": actionKpi.TargetCoverage,
						"high-level-frequency": actionKpi.HighLevelFrequency,
						"middle-level-frequency": actionKpi.MiddleLevelFrequency,
						"low-level-frequency": actionKpi.LowLevelFrequency,
					})
				}
			}
		}

	}

	// manager
	for _, resourceConfig := range resourceConfigs {
		if resourceConfig.ResourceType != 0 {
			managerConfig, _ := managerConfigStorage.GetOne(resourceConfigs[0].ResourceID)
			historyManager["total-business-indicators"] = managerConfig.TotalBusinessIndicators
			historyManager["total-budgets"] = managerConfig.TotalBudgets
			historyManager["total-meeting-places"] = managerConfig.TotalMeetingPlaces
			historyManager["manager-kpi"] = managerConfig.ManagerKPI
			historyManager["manager-time"] = managerConfig.ManagerTime
		}
	}


	histories["hospitals"] = historyHospitals
	histories["products"] = historyProducts
	histories["representatives"] = historyRepresentatives
	histories["manager"] = historyManager

	header := map[string]string{}
	currentScenario := map[string]interface{}{}
	var scenarios []interface{}
	body := map[string]interface{}{}

	header["application"] = "tmist"
	header["contentType"] = "json"

	sm, _ := scenarioStorage.GetOne(scenarioId)
	currentScenario["id"] = sm.ID
	currentScenario["phase"] = sm.Phase

	body["inputs"] = inputs
	body["histories"] = histories

	// 查询所有的周期
	cleanQueryParams(&req)
	req.QueryParams["proposal-id"] = []string{proposalId}
	for _, v := range scenarioStorage.GetAll(req, -1, -1) {
		if v.Phase > 0 {
			scenarios = append(scenarios, map[string]interface{}{
				"id": v.ID,
				"phase" : v.Phase,
			})
		}
	}

	cs := &calcStruct {
		Header:     header,
		Account:    accountId,
		Proposal:   proposalId,
		PaperInput: paperInput.ID,
		CurrentScenario:   currentScenario,
		Scenarios:		scenarios,
		Body:       body,
	}

	c, _ := json.Marshal(cs)

	fmt.Println(string(c))

	topic := h.kafka.Topics[0]
	fmt.Println(topic)
	h.kafka.Produce(&topic, c)

	result["status"] = "ok"
	result["msg"] = "正在计算"
	enc.Encode(result)

	return 0
}

func (h NtmCallRHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h NtmCallRHandler) GetHandlerMethod() string {
	return h.Method
}

func cleanQueryParams(r *api2go.Request) {
	r.QueryParams = map[string][]string{}
}

func getApi2goRequest(r *http.Request, header http.Header) api2go.Request{
	return api2go.Request{
		PlainRequest: r,
		Header: header,
		QueryParams: map[string][]string{},
	}
}

func subscriptionFunc(content interface{}) {
	//c := content.([]byte)
	//fmt.Println(string(c))

	h := NtmCallR
	c := content.([]byte)
	fmt.Println(string(c))

	var result resultStruct

	err := json.Unmarshal(c, &result)

	if result.Header["application"] == "tmist" {
		ctx := map[string]string {
			"client-id": "5cbd9f94f4ce4352ecb082a0",
			"type": "calc",
			"account-id": result.Account,
			"proposal-id": result.Proposal,
			"paperInput-id": result.PaperInput,
			"scenario-id": result.Scenario,
			"time": strconv.FormatInt(time.Now().UnixNano() / 1e6, 10),
		}

		if err != nil ||  result.Error != nil{
			ctx["status"] = "no"
			ctx["msg"] = "计算失败"
			r, _ := json.Marshal(ctx)
			fmt.Println(string(r))
			_ = h.xmpp.SendGroupMsg(h.Args[0], string(r))
			return
		}

		if len(c) > 2 {
			var (
				hospitalSalesReport NtmModel.HospitalSalesReport
				productSalesReport NtmModel.ProductSalesReport
				representativeSalesReport NtmModel.RepresentativeSalesReport
				representativeAbility NtmModel.RepresentativeAbility
				actionKpi NtmModel.ActionKpi

				hospitalSalesReportIDs []string
				productSalesReportIDs []string
				representativeSalesReportIDs []string
				representativeAbilityIDs []string
				actionKpiIDs []string

				assessmentReportID string
			)
			mdb := []BmDaemons.BmDaemon{h.db}

			scenarioStorage := NtmDataStorage.NtmScenarioStorage{}.NewScenarioStorage(mdb)
			hospitalSalesReportStorage := NtmDataStorage.NtmHospitalSalesReportStorage{}.NewHospitalSalesReportStorage(mdb)
			productSalesReportStorage := NtmDataStorage.NtmProductSalesReportStorage{}.NewProductSalesReportStorage(mdb)
			representativeSalesReportStorage := NtmDataStorage.NtmRepresentativeSalesReportStorage{}.NewRepresentativeSalesReportStorage(mdb)
			representativeAbilityStorage := NtmDataStorage.NtmRepresentativeAbilityStorage{}.NewRepresentativeAbilityStorage(mdb)
			actionKpiStorage := NtmDataStorage.NtmActionKpiStorage{}.NewActionKpiStorage(mdb)
			salesReportStorage := NtmDataStorage.NtmSalesReportStorage{}.NewSalesReportStorage(mdb)
			personnelAssessmentStorage := NtmDataStorage.NtmPersonnelAssessmentStorage{}.NewPersonnelAssessmentStorage(mdb)
			paperStorage := NtmDataStorage.NtmPaperStorage{}.NewPaperStorage(mdb)

			assessmentReportStorage := NtmDataStorage.NtmAssessmentReportStorage{}.NewAssessmentReportStorage(mdb)
			regionalDivisionResultStorage := NtmDataStorage.NtmRegionalDivisionResultStorage{}.NewRegionalDivisionResultStorage(mdb)
			targetAssignsResultStorage := NtmDataStorage.NtmTargetAssignsResultStorage{}.NewTargetAssignsResultStorage(mdb)
			resourceAssignsResultStorage := NtmDataStorage.NtmResourceAssignsResultStorage{}.NewResourceAssignsResultStorage(mdb)
			manageTimeResultStorage := NtmDataStorage.NtmManageTimeResultStorage{}.NewManageTimeResultStorage(mdb)
			manageTeamResultStorage := NtmDataStorage.NtmManageTeamResultStorage{}.NewManageTeamResultStorage(mdb)
			generalPerformanceResultStorage := NtmDataStorage.NtmGeneralPerformanceResultStorage{}.NewGeneralPerformanceResultStorage(mdb)
			assessmentReportDescribeStorage := NtmDataStorage.NtmAssessmentReportDescribeStorage{}.NewAssessmentReportDescribeStorage(mdb)
			levelStorage := NtmDataStorage.NtmLevelStorage{}.NewLevelStorage(mdb)
			levelConfigStorage := NtmDataStorage.NtmLevelConfigStorage{}.NewLevelConfigStorage(mdb)

			req := api2go.Request{
				QueryParams: map[string][]string{},
			}

			req.QueryParams["proposal-id"] = []string{result.Proposal}
			req.QueryParams["account-id"] = []string{result.Account}
			req.QueryParams["orderby"] = []string{"time"}

			papers := paperStorage.GetAll(req, -1, -1)

			if len(papers) > 0 {
				paper := papers[len(papers) - 1]
				body := result.Body

				hospitalSalesReports := body["hospitalSalesReports"].([]interface{})
				productSalesReports := body["productSalesReports"].([]interface{})
				representativeSalesReports := body["representativeSalesReports"].([]interface{})
				representativeAbilities := body["representativeAbility"].([]interface{})
				actionKpis := body["actionKpi"].([]interface{})

				for _, v := range hospitalSalesReports {
					mapstructure.Decode(v, &hospitalSalesReport)
					hospitalSalesReportIDs = append(hospitalSalesReportIDs, hospitalSalesReportStorage.Insert(hospitalSalesReport))
				}

				for _, v := range productSalesReports {
					mapstructure.Decode(v, &productSalesReport)
					productSalesReportIDs = append(productSalesReportIDs, productSalesReportStorage.Insert(productSalesReport))
				}

				for _, v := range representativeSalesReports {
					mapstructure.Decode(v, &representativeSalesReport)
					representativeSalesReportIDs = append(representativeSalesReportIDs, representativeSalesReportStorage.Insert(representativeSalesReport))
				}

				for _, v := range representativeAbilities {
					mapstructure.Decode(v, &representativeAbility)
					representativeAbilityIDs = append(representativeAbilityIDs, representativeAbilityStorage.Insert(representativeAbility))
				}

				for _, v := range actionKpis {
					mapstructure.Decode(v, &actionKpi)
					actionKpiIDs = append(actionKpiIDs, actionKpiStorage.Insert(actionKpi))
				}

				salesReportID := salesReportStorage.Insert(NtmModel.SalesReport{
					ScenarioID: result.Scenario,
					PaperInputID: result.PaperInput,
					Time: time.Now().UnixNano() / 1e6,
					HospitalSalesReportIDs: hospitalSalesReportIDs,
					ProductSalesReportIDs: productSalesReportIDs,
					RepresentativeSalesReportIDs: representativeSalesReportIDs,
				})

				paper.SalesReportIDs = append(paper.SalesReportIDs, salesReportID)

				personnelAssessmentID := personnelAssessmentStorage.Insert(NtmModel.PersonnelAssessment{
					ScenarioID: result.Scenario,
					PaperInputID: result.PaperInput,
					RepresentativeAbilityIDs: representativeAbilityIDs,
					ActionKpiIDs: actionKpiIDs,
					Time: time.Now().UnixNano() / 1e6,
				})

				paper.PersonnelAssessmentIDs = append(paper.PersonnelAssessmentIDs, personnelAssessmentID)

				req.QueryParams["proposal-id"] = []string{result.Proposal}
				scenarios := scenarioStorage.GetAll(req, -1,-1)

				if s := scenarios[len(scenarios)-1]; s.ID == result.Scenario {
					var (
						regionDivisionResult assessmentResult
						targetAssignsResult assessmentResult
						resourceAssignsResult assessmentResult
						manageTimeResult assessmentResult
						manageTeamResult assessmentResult
						generalPerformanceResult assessmentResult
						regionDivisionResultDescribeIDs []string
						targetAssignsResultDescribeIDs []string
						resourceAssignsResultDescribeIDs []string
						manageTimeResultDescribeIDs []string
						manageTeamResultDescribeIDs []string
						generalPerformanceResultDescribeIDs []string
					)
					mapstructure.Decode(body["regionDivisionResult"], &regionDivisionResult)
					mapstructure.Decode(body["targetAssignsResult"], &targetAssignsResult)
					mapstructure.Decode(body["resourceAssignsResult"], &resourceAssignsResult)
					mapstructure.Decode(body["manageTimeResult"], &manageTimeResult)
					mapstructure.Decode(body["manageTeamResult"], &manageTeamResult)
					mapstructure.Decode(body["generalPerformanceResult"], &generalPerformanceResult)

					// TODO 测试结束后拉出去
					// regionDivisionResult
					req.QueryParams = map[string][]string{}
					req.QueryParams["level-code"] = []string{strconv.Itoa(regionDivisionResult.Level)}
					regionDivisionResultLevels := levelStorage.GetAll(req, -1, -1)
					req.QueryParams = map[string][]string{}
					req.QueryParams["code"] = []string{strconv.Itoa(regionDivisionResult.Code)}
					for _, av := range assessmentReportDescribeStorage.GetAll(req, -1, -1) {
						regionDivisionResultDescribeIDs = append(regionDivisionResultDescribeIDs, av.ID)
					}
					req.QueryParams["level-id"] = []string{regionDivisionResultLevels[0].ID}
					regionDivisionResultLevelConfigs := levelConfigStorage.GetAll(req, -1, -1)
					regionalDivisionResultID := regionalDivisionResultStorage.Insert(NtmModel.RegionalDivisionResult{
						LevelConfigID: regionDivisionResultLevelConfigs[0].ID,
						AssessmentReportDescribeIDs: regionDivisionResultDescribeIDs,
					})

					// targetAssignsResult
					req.QueryParams = map[string][]string{}
					req.QueryParams["level-code"] = []string{strconv.Itoa(targetAssignsResult.Level)}
					targetAssignsResultLevels := levelStorage.GetAll(req, -1, -1)
					req.QueryParams = map[string][]string{}
					req.QueryParams["code"] = []string{strconv.Itoa(targetAssignsResult.Code)}
					for _, av := range assessmentReportDescribeStorage.GetAll(req, -1, -1) {
						targetAssignsResultDescribeIDs = append(targetAssignsResultDescribeIDs, av.ID)
					}
					req.QueryParams["level-id"] = []string{targetAssignsResultLevels[0].ID}
					targetAssignsResultLevelConfigs := levelConfigStorage.GetAll(req, -1, -1)
					targetAssignsResultID := targetAssignsResultStorage.Insert(NtmModel.TargetAssignsResult{
						LevelConfigID: targetAssignsResultLevelConfigs[0].ID,
						AssessmentReportDescribeIDs: targetAssignsResultDescribeIDs,
					})


					// resourceAssignsResult
					req.QueryParams = map[string][]string{}
					req.QueryParams["level-code"] = []string{strconv.Itoa(resourceAssignsResult.Level)}
					resourceAssignsResultLevels := levelStorage.GetAll(req, -1, -1)
					req.QueryParams = map[string][]string{}
					req.QueryParams["code"] = []string{strconv.Itoa(resourceAssignsResult.Code)}
					for _, av := range assessmentReportDescribeStorage.GetAll(req, -1, -1) {
						resourceAssignsResultDescribeIDs = append(resourceAssignsResultDescribeIDs, av.ID)
					}
					req.QueryParams["level-id"] = []string{resourceAssignsResultLevels[0].ID}
					resourceAssignsResultLevelConfigs := levelConfigStorage.GetAll(req, -1, -1)
					resourceAssignsResultID := resourceAssignsResultStorage.Insert(NtmModel.ResourceAssignsResult{
						LevelConfigID: resourceAssignsResultLevelConfigs[0].ID,
						AssessmentReportDescribeIDs: resourceAssignsResultDescribeIDs,
					})


					// manageTimeResult
					req.QueryParams = map[string][]string{}
					req.QueryParams["level-code"] = []string{strconv.Itoa(manageTimeResult.Level)}
					manageTimeResultLevels := levelStorage.GetAll(req, -1, -1)
					req.QueryParams = map[string][]string{}
					req.QueryParams["code"] = []string{strconv.Itoa(manageTimeResult.Code)}
					for _, av := range assessmentReportDescribeStorage.GetAll(req, -1, -1) {
						manageTimeResultDescribeIDs = append(manageTimeResultDescribeIDs, av.ID)
					}
					req.QueryParams["level-id"] = []string{manageTimeResultLevels[0].ID}
					manageTimeResultLevelConfigs := levelConfigStorage.GetAll(req, -1, -1)
					manageTimeResultID := manageTimeResultStorage.Insert(NtmModel.ManageTimeResult{
						LevelConfigID: manageTimeResultLevelConfigs[0].ID,
						AssessmentReportDescribeIDs: manageTimeResultDescribeIDs,
					})



					// manageTeamResult
					req.QueryParams = map[string][]string{}
					req.QueryParams["level-code"] = []string{strconv.Itoa(manageTeamResult.Level)}
					manageTeamResultLevels := levelStorage.GetAll(req, -1, -1)
					req.QueryParams = map[string][]string{}
					req.QueryParams["code"] = []string{strconv.Itoa(manageTeamResult.Code)}
					for _, av := range assessmentReportDescribeStorage.GetAll(req, -1, -1) {
						manageTeamResultDescribeIDs = append(manageTeamResultDescribeIDs, av.ID)
					}
					req.QueryParams["level-id"] = []string{manageTeamResultLevels[0].ID}
					manageTeamResultLevelConfigs := levelConfigStorage.GetAll(req, -1, -1)
					manageTeamResultID := manageTeamResultStorage.Insert(NtmModel.ManageTeamResult{
						LevelConfigID: manageTeamResultLevelConfigs[0].ID,
						AssessmentReportDescribeIDs: manageTeamResultDescribeIDs,
					})


					// generalPerformanceResult
					req.QueryParams = map[string][]string{}
					req.QueryParams["level-code"] = []string{strconv.Itoa(generalPerformanceResult.Level)}
					generalPerformanceResultLevels := levelStorage.GetAll(req, -1, -1)
					req.QueryParams = map[string][]string{}
					req.QueryParams["code"] = []string{strconv.Itoa(generalPerformanceResult.Code)}
					for _, av := range assessmentReportDescribeStorage.GetAll(req, -1, -1) {
						generalPerformanceResultDescribeIDs = append(generalPerformanceResultDescribeIDs, av.ID)
					}
					req.QueryParams["level-id"] = []string{generalPerformanceResultLevels[0].ID}
					generalPerformanceResultLevelConfigs := levelConfigStorage.GetAll(req, -1, -1)
					generalPerformanceResultID := generalPerformanceResultStorage.Insert(NtmModel.GeneralPerformanceResult{
						LevelConfigID: generalPerformanceResultLevelConfigs[0].ID,
						AssessmentReportDescribeIDs: generalPerformanceResultDescribeIDs,
					})

					assessmentReport := NtmModel.AssessmentReport{
						ScenarioID:   result.Scenario,
						Time:         time.Now().UnixNano() / 1e6,
						PaperInputID: result.PaperInput,
						RegionalDivisionResultID: regionalDivisionResultID,
						TargetAssignsResultID: targetAssignsResultID,
						ResourceAssignsResultID: resourceAssignsResultID,
						ManageTimeResultID: manageTimeResultID,
						ManageTeamResultID: manageTeamResultID,
						GeneralPerformanceResultID: generalPerformanceResultID,
					}

					assessmentReportID = assessmentReportStorage.Insert(assessmentReport)
				}


				if len(assessmentReportID) > 0 {
					paper.AssessmentReportIDs = append(paper.AssessmentReportIDs, assessmentReportID)
				}

				// TODO: @Alex自己留，这面等重构
				var state int
				for _, scenario := range scenarios {
					if scenario.ID == result.Scenario {
						if paper.TotalPhase == scenario.Phase {
							state = 3
						} else {
							state = 2
						}
					}
				}
				paper.InputState = state
				// TODO: @Alex 时间有问题存在UTC转CST问题，因为服务器的都是UTC，Golang默认也是读UTC，等Bug改完后整体做转换
				paper.EndTime = time.Now().UnixNano() / 1e6

				err = paperStorage.Update(*paper)

				if err != nil {
					panic("更新Paper失败")
				}

				ctx["status"] = "ok"
				ctx["msg"] = "计算成功"

				r, _ := json.Marshal(ctx)
				fmt.Println(string(r))
				_ = h.xmpp.SendGroupMsg(h.Args[0], string(r))
			}  else {
				ctx["status"] = "no"
				ctx["msg"] = "计算失败，出现异常！"

				r, _ := json.Marshal(ctx)
				fmt.Println(string(r))
				_ = h.xmpp.SendGroupMsg(h.Args[0], string(r))
			}
		}
	}
}