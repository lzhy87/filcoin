package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Miner struct {
	Address   string `json:"address"`
	BlockCid  string `json:"blockCid"`
	IdAddress string `json:"idAddress"`
	Method    string `json:"method"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}
type RespWallet struct {
	Name      string
	Code      int              `json:"code"`
	Total     int              `json:"total"`
	PageIndex int              `json:"pageIndex"`
	PageSize  int              `json:"pageSize"`
	Message   string           `json:"Message"`
	Data      []WalletSendInfo `json:"data"`
}
type WalletSendInfo struct {
	Cid          string `json:"cid"`
	TimeFormat   string `json:"timeFormat"`
	Time         int    `json:"time"`
	Fee          string `json:"fee"`
	Height       int    `json:"height"`
	From         string `json:"from"`
	To           string `json:"to"`
	Value        string `json:"value"`
	Method       string `json:"method"`
	ToType       int    `json:"toType"`
	ExitCodeName string `json:"exitCodeName"`
}

func PostMiner(m []Miner) []RespWallet {

	client := &http.Client{}
	rs := []RespWallet{}
	for _, v := range m {
		data, _ := json.Marshal(v)
		fmt.Println(string(data))
		req, _ := http.NewRequest("POST", "https://api2.filscout.com/api/v1/message", bytes.NewReader(data))
		resp, _ := client.Do(req)
		defer req.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		r := RespWallet{}
		err := json.Unmarshal(body, &r)
		switch v.Address {
		case "f3rhypkcrozpohaievn7bq7rfg45e2exyvnzxtexmzp3qkhzieynlcqie7j5f2jlj2ocnrdnqjsf6k22vbbw5a":
			r.Name = "znyk_worker"
		case "f3uw6a3htjf5vblz7fhn6fm3wc56bujxnp2lngndjaq2snbtqz44qewcfxmcktvjby2sku3n5gb5yxhjip3upa":
			r.Name = "znyk_onwer"
		case "f3qv5cvsho6q5csqwsdf2lw4thecm2bpwtm5kvqhtsv2b6okslrsciuhwczgmufjlzbi3hshl32w465e4jx2kq":
			r.Name = "xz_worker"
		case "f3vk6gavogn2fjeltteeskhi35zabytehiysipxntnohm7c6nqkvx2h4qipkfunze65przue46m3rujqih7kqa":
			r.Name = "xz_owner"
		default:
			r.Name = "未知"
		}
		if err != nil {
			log.Fatal(err)
		}
		rs = append(rs, r)
	}
	return rs
	// fmt.Println(string(body))

}
