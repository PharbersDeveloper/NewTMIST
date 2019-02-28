package NtmHandler

import (
	"fmt"
	"github.com/PharbersDeveloper/NtmPods/NtmPanic"
	"net/http"
)

type CommonPanicHandle struct {
}

func (ctm CommonPanicHandle) NewCommonPanicHandle(args ...interface{}) CommonPanicHandle {
	return CommonPanicHandle{}
}

func (ctm CommonPanicHandle) HandlePanic(rw http.ResponseWriter, r *http.Request, p interface{}) {
	fmt.Println("CommonHandlePanic接收到", p)
	NtmPanic.ErrInstance().ErrorReval(p.(string), rw)
}
