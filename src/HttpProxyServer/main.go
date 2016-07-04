package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"httpproxy"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func HandleHttpProxy(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handleHttpProxy start deserializer for req")
	pf := thrift.NewTBinaryProtocolFactoryDefault()

	t1 := thrift.NewTDeserializer()
	t1.Protocol = pf.GetProtocol(t1.Transport)
	reqHttpPackage := httpproxy.NewReqHttpPackage()
	if req.ContentLength > 0 {
		body, _ := ioutil.ReadAll(req.Body)
		bodystr := string(body)
		t1.ReadString(reqHttpPackage, bodystr)
		fmt.Println("parse req body :", reqHttpPackage.String())
	}

	fmt.Println("handleHttpProxy start serializer for resp")
	resp := httpproxy.NewRespHttpPackage()
	resp.NCMDID = reqHttpPackage.NCMDID
	resp.Sequence = reqHttpPackage.Sequence
	resp.Result_ = 0
	resp.SMessage = "success zeus "
	fmt.Println("")
	t := thrift.NewTSerializer()
	t.Protocol = pf.GetProtocol(t.Transport)
	b, err := t.Write(resp)
	if err == nil {
		w.Write(b)
	} else {
		fmt.Println("parse resp error")

	}
}

func main() {

	//	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("NumCPU: ", runtime.NumCPU())
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello, world.")
	//})

	http.HandleFunc("/httpproxy", HandleHttpProxy)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
