package ztpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	msg "ztp"
)

func (zc *ZtpClient) SendEvents() (err error) {
	events := &msg.EventMsg{}
	zc.Device.SendEvents(events)
	if len(events.Event) == 0 {
		return
	}
	b, err := json.Marshal(events)
	if err != nil {
		log.Println(err)
		return
	}

	r, err := http.NewRequest("POST", zc.urlPrefix+"event", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		return
	}
	r.SetBasicAuth(zc.login, zc.password)
	r.Header.Add("Content-type", "application/json")
	if DEBUG {
		msg.DumpJson(r)
		msg.DumpJson(events)
	}

eventRetry:
	for retry := 0; retry < 4; retry++ {
		resp, _, err := zc.SendRequest(r)

		if DEBUG {
			log.Println("HTTP Response code", resp.StatusCode)
		}
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * time.Duration(zc.eventRetry))
			continue
		}
		switch resp.StatusCode {
		case http.StatusOK:
			break eventRetry
		}
		continue
	}
	zc.Device.SendEventsComplete()
	return nil
}
