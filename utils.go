// util package contains a cross section of utility functions
package ztp

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
    "net/http/httputil"
)

type UIErrorType struct {
	Errors []UIErrors `json:"errors"`
}

// This is the common error reporting format for all HTTP error reporting
type UIErrors struct {
	ErrorMessage string `json:"errorMessage"`
	Resource     string `json:"resource,omitempty"`
	ErrorCode    int    `json:"errorCode"`
	ReasonCode   int    `json:"reasonCode,omitempty"`
}

// Environment variable utilDEBUG=1 sets the package DEBUG=true
const EnvDebug = "utilDEBUG"

//DEBUG can be enabled for this package via:
//  testutil.DEBUG = true
//  export utilDEBUG=1 environment variable
var DEBUG = false

func init() {
	envSetup()
}
func envSetup() {
	if d, ok := os.LookupEnv(EnvDebug); ok {
		switch d {
		case "0":
			DEBUG = false
		case "1":
			DEBUG = true
		}
	}
}

// DumpJson() pretty prints any data structure into formatted JSON.
// When DEBUG == false (default), no output is produced. Code can leave DumpJson(<struct>) calls in place without concert that output will be cluttered with volumes of data.
// Output is logged instead of printed to stdout
func DumpJson(n interface{}) {
	// set util.DEBUG = true in your application
	if !DEBUG {
		return
	}

	_, fname, lineno, _ := runtime.Caller(1)

	j, err := json.MarshalIndent(n, "", "    ")
	if err != nil {
		log.Printf("%s:%d %+v\n", fname, lineno, err)
		log.Printf("%s:%d %+v\n", fname, lineno, n)
	} else {
		log.Printf("%s:%d %+v\n", fname, lineno, string(j))
	}
}

func DumpReqURL(r *http.Request) {
    if !DEBUG {
        return
    }
    log.Println(r.Method, r.URL.String())
}
func DumpRequest(r *http.Request) {
    if !DEBUG {
        return
    }
    x,_ := httputil.DumpRequest(r, true)
    log.Println(string(x))
}
func DumpResponse(resp *http.Response) {
    if !DEBUG {
        return
    }
    x,_ := httputil.DumpResponse(resp, true)
    log.Println(string(x))
}

// UIError - used for creating UI error responses by producting a standardized error reporting format
// It is a direct replacement for http.Error(w, <string>, code)
func UIError(w http.ResponseWriter, msg string, code int) {
	uiErr := UIErrorType{Errors: []UIErrors{
		UIErrors{
			ErrorCode:    code,
			ErrorMessage: msg,
			Resource:     "",
		},
	},
	}
	DumpJson(uiErr)
	j, _ := json.Marshal(&uiErr)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

// UIErrorWithReason - used for creating UI error responses by producting a standardized error reporting format with the addition of a reason code
// It is not a direct replacement for http.Error(w, <string>, code)
// Primarily used during JWT auth to provide additional reasons why a 401 response occurs.
func UIErrorWithReason(w http.ResponseWriter, msg string, code int, reason int) {
	uiErr := UIErrorType{Errors: []UIErrors{
		UIErrors{
			ErrorCode:    code,
			ErrorMessage: msg,
			Resource:     "",
			ReasonCode:   reason,
		},
	},
	}
	DumpJson(uiErr)
	j, _ := json.Marshal(&uiErr)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(j)
}

// SendGzipJson - Format an HTTP response by encoding JSON into a gzipped payload
// the HTTP header is set to Content-Encoding: gzip
func SendGzipJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	g := gzip.NewWriter(w)
	json.NewEncoder(g).Encode(data)
	g.Close()
}
