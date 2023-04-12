package logging

import (
	"fmt"
	"github.com/golang/glog"
)

func Info() {
	glog.Infoln("[]")
}

func Recovery(err interface{}, stack []byte) {
	fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack)
}
