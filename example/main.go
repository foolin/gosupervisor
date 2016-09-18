package main

import (
	"github.com/foolin/gosupervisor"
	"log"
	"time"
)

func main(){
	rpcUrl := "http://127.0.0.1:9001/RPC2"
	rpc := gosupervisor.New(rpcUrl)

	version, err := rpc.GetAPIVersion()
	println("GetAPIVersion")
	printf("ret: %v, error: %v", version, err)

	stopProgress, err := rpc.StopProcess("servername1", true)
	println("StopProcess")
	printf("ret: %v, error: %v", stopProgress, err)

	startProcess, err := rpc.StartProcess("servername1", true)
	println("StartProcess")
	printf("ret: %v, error: %v", startProcess, err)

	list, _ := rpc.GetAllProcessInfo()
	println("GetAllProcessInfo")
	printf("name\t\tstart\t\tdescription")
	for _, v := range list{
		printf("%v\t\t%v\t\t%v", v.Get("name"), time.Unix(v.Int64("start", 0), 0), v.Get("description"))
	}
}


func println(v string)  {
	log.Printf("------------- %v -------------", v)
}

func printf(format string, v ...interface{})  {
	log.Printf(format, v...)
}