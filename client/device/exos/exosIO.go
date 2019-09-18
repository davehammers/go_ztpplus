package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os/exec"
)

type exosIOT interface {
	Get(cmd string, i interface{}) error
	//Set(cmd string) error
}

var exosIO exosIOT

type localEXOS struct{}
type remoteEXOS struct {
	ipAddr string
}

func SetExosIO(ipAddr string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if ipAddr == "" {
		// running locally on an EXOS switch
		exosIO = localEXOS{}
		return
	}
	exosIO = remoteEXOS{
		ipAddr: ipAddr,
	}

}

func (l localEXOS) Get(cmd string, i interface{}) (err error) {
	out, err := exec.Command("/exos/bin/exsh", "-b", "-n", "0", "-c", string(cmd)).Output()
	if out != nil {
		var i interface{}
		// check to see if the results are JSON
		err = json.Unmarshal(out, &i)
	}
	return
}

func (r remoteEXOS) Get(cmd string, i interface{}) (err error) {
	// connect to this socket
	conn, err := net.Dial("tcp", r.ipAddr+":2020")
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	log.Println("sending", cmd)
	fmt.Fprintf(conn, cmd+"\n")
	err = json.NewDecoder(conn).Decode(&i)
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	//x, _ := json.MarshalIndent(i, "", "  ")
	//log.Println(string(x))
	return
}
