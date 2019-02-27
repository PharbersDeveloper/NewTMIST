package BmHandler

import (
	"fmt"
	"net/http"
	"github.com/PharbersDeveloper/NtmPods/BmPanic"
)

type CommonPanicHandle struct {
}

func (ctm CommonPanicHandle) NewCommonPanicHandle(args ...interface{}) CommonPanicHandle {
	return CommonPanicHandle{}
}

func (ctm CommonPanicHandle) HandlePanic(rw http.ResponseWriter, r *http.Request, p interface{}) {
	fmt.Println("CommonHandlePanic接收到", p)
	BmPanic.ErrInstance().ErrorReval(p.(string), rw)
}
