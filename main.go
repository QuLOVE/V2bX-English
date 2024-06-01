package main

import (
	//"net/http"
	//_ "net/http/pprof"

	"github.com/QuLOVE/V2bX-English/cmd"
)

func main() {
	//RAM leak detection
	//go func() {
	//	http.ListenAndServe("127.0.0.1:6060", nil)
	//}()
	cmd.Run()
}
