package main

import (
	"filcoin/utils"
	"flag"
	"fmt"
)

const (
	pageSize          int    = 20
	znyk_workerWallet string = "f3rhypkcrozpohaievn7bq7rfg45e2exyvnzxtexmzp3qkhzieynlcqie7j5f2jlj2ocnrdnqjsf6k22vbbw5a"
	znyk_ownerWallet  string = "f3uw6a3htjf5vblz7fhn6fm3wc56bujxnp2lngndjaq2snbtqz44qewcfxmcktvjby2sku3n5gb5yxhjip3upa"
	xz_workerWallet   string = "f3qv5cvsho6q5csqwsdf2lw4thecm2bpwtm5kvqhtsv2b6okslrsciuhwczgmufjlzbi3hshl32w465e4jx2kq"
	xz_ownerWallet    string = "f3vk6gavogn2fjeltteeskhi35zabytehiysipxntnohm7c6nqkvx2h4qipkfunze65przue46m3rujqih7kqa"
)

func Play(m utils.Miner, pageIndex int) {
	var ms []utils.Miner = make([]utils.Miner, 0)
	for i := 1; i <= pageIndex; i++ {
		m.PageIndex = i
		ms = append(ms, m)
	}
	rs := utils.PostMiner(ms)
	utils.CreateExcel(rs)
}
func main() {

	var pageIndex int
	var walletName string
	flag.StringVar(&walletName, "n", "znyk_owner", "默认中南云矿owner,也可以是znyk_worker")
	flag.IntVar(&pageIndex, "p", 10, "默认10页")
	flag.Parse()
	switch {
	case walletName == "znyk_owner":
		m := utils.Miner{Address: znyk_ownerWallet, BlockCid: "", IdAddress: "", Method: "Send", PageIndex: pageIndex, PageSize: pageSize}
		Play(m, pageIndex)
	case walletName == "znyk_worker":
		m := utils.Miner{Address: znyk_workerWallet, BlockCid: "", IdAddress: "", Method: "Send", PageIndex: pageIndex, PageSize: pageSize}
		Play(m, pageIndex)
	case walletName == "xz_worker":
		m := utils.Miner{Address: xz_workerWallet, BlockCid: "", IdAddress: "", Method: "Send", PageIndex: pageIndex, PageSize: pageSize}
		Play(m, pageIndex)
	case walletName == "xz_owner":
		m := utils.Miner{Address: xz_ownerWallet, BlockCid: "", IdAddress: "", Method: "Send", PageIndex: pageIndex, PageSize: pageSize}
		Play(m, pageIndex)
	default:
		fmt.Println("输入参数有误!!")
	}
}
