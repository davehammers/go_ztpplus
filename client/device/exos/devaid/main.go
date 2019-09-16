package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"os/exec"
	"time"
)

type Error struct {
	Error  string `json:"error,omitempty"`
	Cmd    string `json:"cmd,omitempty"`
	Output string `json:"output,omitempty"`
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	for {
		l, err := net.Listen("tcp", ":2020")
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 20)
			continue
		}
		for {
			// Wait for a connection.
			// each connection is one CLI command
			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
				break
			}
			// Handle the connection in a new goroutine.
			// The loop then returns to accepting, so that
			// multiple connections may be served concurrently.
			go func(c net.Conn) {
				var err error
				defer c.Close()
				cmd, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					log.Println(err)
					e := Error{
						Error: err.Error(),
						Cmd:   cmd,
					}
					json.NewEncoder(c).Encode(e)
					return
				}

				out, err := exec.Command("/exos/bin/exsh", "-b", "-n", "0", "-c", string(cmd)).Output()
				if out != nil {
					var i interface{}
					// check to see if the results are JSON
					err = json.Unmarshal(out, &i)
				}
				// not JSON, format the response in a JSON error format
				if err != nil {
					//log.Println(err)
					e := Error{
						Error:  err.Error(),
						Cmd:    cmd,
						Output: string(out),
					}
					json.NewEncoder(c).Encode(e)
					return
				}
				c.Write(out)
			}(conn)
		}
		defer l.Close()
	}
}
