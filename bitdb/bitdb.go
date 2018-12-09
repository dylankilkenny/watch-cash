package bitdb

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url string = "https://bitdb.network/q/"

type BitDbRequest struct {
	api   string
	BitDB BitDB
}

type Query struct {
	Find interface{} `json:"find"`
}

type Response struct {
	Function string `json:"f"`
}

type BitDB struct {
	Version  int      `json:"v"`
	Query    Query    `json:"q"`
	Response Response `json:"r"`
}

type TxHash struct {
	Hash string `json:"tx.h"`
}

type BitDdResponse struct {
	Confirmed   string   `json:"c"`
	Unconfirmed []string `json:"u"`
}

func New(version int, api string) *BitDbRequest {
	request := new(BitDbRequest)
	request.api = api
	request.BitDB.Version = version
	return request
}

func (b *BitDbRequest) Request(query interface{}, jq string) (string, error) {
	b.BitDB.Query.Find = query
	b.BitDB.Response.Function = jq
	j, err := json.Marshal(b.BitDB)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	b64Query := b64.StdEncoding.EncodeToString([]byte(j))
	req, err := http.NewRequest("GET", url+b64Query, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("key", b.api)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body), nil
}
