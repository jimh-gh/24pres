package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const token = "REPLACE ME"
const user = "REPLACE ME"

// define title or else hostname is used.

type PushResp struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

func push(t string, m string) {
	if t == "" {
		t, _ = os.Hostname()
	}
	var mes PushResp
	params := url.Values{}
	params.Add("token", token)
	params.Add("user", user)
	params.Add("title", t)
	params.Add("monospace", "1")
	params.Add("message", m)

	resp, err := http.PostForm("https://api.pushover.net/1/messages.json", params)
	if err != nil {
		log.Println(err)
		return
	} else {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
			}
		}(resp.Body)
		//if resp.StatusCode >=400 && resp.StatusCode <=499 {
		//	//goto End
		//}
		//else if resp.StatusCode == 500 {
		//	time.Sleep(5 * time.Second)
		//}

		body, _ := io.ReadAll(resp.Body)
		err := json.Unmarshal(body, &mes)
		if err != nil {
			log.Println(err)
		}
		if mes.Status != 1 {
			log.Println(mes)
		}
		//log.Println(params)
	}

}
