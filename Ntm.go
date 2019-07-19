package main

import (
	"Ntm/NtmFactory"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
)

func main() {

	// 本地调试使用，部署时注释
	//os.Setenv("BM_KAFKA_CONF_HOME", fmt.Sprint(os.Getenv("BM_KAFKA_CONF_HOME"), "NtmServiceDeploy/dev-config/resource/kafkaconfig.json"))
	//os.Setenv("BM_XMPP_CONF_HOME", fmt.Sprint(os.Getenv("BM_XMPP_CONF_HOME"), "NtmServiceDeploy/dev-config/resource/xmppconfig.json"))

	version := "v0"
	prodEnv := "NTM_HOME"
	fmt.Println("NTM pods archi begins, version =", version)

	fac := NtmFactory.NtmTable{}
	var pod = BmPodsDefine.Pod{Name: "new TMIST", Factory: fac}
	ntmHome := os.Getenv(prodEnv)
	pod.RegisterSerFromYAML(ntmHome + "/resource/service-def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig(prodEnv)

	addr := bmRouter.Host + ":" + bmRouter.Port
	fmt.Println("Listening on ", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)

	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)

	handler := api.Handler().(*httprouter.Router)
	pod.RegisterPanicHandler(handler)
	http.ListenAndServe(":"+bmRouter.Port, handler)
}
