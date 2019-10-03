package solvingmaze

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

type client struct {
	apiKey      string
	warehouseId string
	httpClient  *http.Client
	debug       bool
}

func NewClient(apiKey string, warehouseId string, timeout int, debug bool) (c *client) {
	return &client{apiKey, warehouseId, &http.Client{Timeout: time.Second * time.Duration(timeout)}, debug}
}

func (c client) dumpRequest(r *http.Request) {
	if r == nil {
		log.Print("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (c client) dumpResponse(r *http.Response) {
	if r == nil {
		log.Print("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

func (c *client) Post(resource string, body []byte) (response []byte, err error) {
	var url string
	if strings.HasPrefix(resource, "http") {
		url = fmt.Sprintf("%s/key/%s/warehouse%s", resource, c.apiKey, c.warehouseId)
	} else {
		url = fmt.Sprintf("%s/%s/key/%s/warehouse/%s", API_BASE, resource, c.apiKey, c.warehouseId)
	}

	if c.debug {
		log.Printf("POST to %s", url)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))

	if err != nil {
		return
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Accept", "application/json")

	if c.debug {
		c.dumpRequest(req)
	}

	resp, err := c.httpClient.Do(req)

	if c.debug {
		c.dumpResponse(resp)
	}

	if err != nil {
		return
	}

	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}

	return
}
