package main

import (
	"fmt"
	"net/http"

	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"os"
	"github.com/PharbersDeveloper/NtmPods/BmFactory"
)

func main() {
	version := "v2"
	fmt.Println("pod archi begins")

	fac := BmFactory.BmTable{}
	var pod = BmPodsDefine.Pod{ Name: "alfred test", Factory:fac }
	bmHome := os.Getenv("PHARBERS_HOME")
	pod.RegisterSerFromYAML(bmHome + "/resource/service-def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig("PHARBERS_HOME")

	addr := bmRouter.Host + ":" + bmRouter.Port
	fmt.Println("Listening on ", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)

	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)

	handler := api.Handler().(*httprouter.Router)
	pod.RegisterPanicHandler(handler)
	http.ListenAndServe(":"+bmRouter.Port, handler)

	fmt.Println("pod archi ends")
}
