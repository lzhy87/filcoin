package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	pageSize          int    = 20
	znyk_workerWallet string = "f3rhypkcrozpohaievn7bq7rfg45e2exyvnzxtexmzp3qkhzieynlcqie7j5f2jlj2ocnrdnqjsf6k22vbbw5a"
	znyk_ownerWallet  string = "f3uw6a3htjf5vblz7fhn6fm3wc56bujxnp2lngndjaq2snbtqz44qewcfxmcktvjby2sku3n5gb5yxhjip3upa"
	xz_workerWallet   string = "f3qv5cvsho6q5csqwsdf2lw4thecm2bpwtm5kvqhtsv2b6okslrsciuhwczgmufjlzbi3hshl32w465e4jx2kq"
	xz_ownerWallet    string = "f3vk6gavogn2fjeltteeskhi35zabytehiysipxntnohm7c6nqkvx2h4qipkfunze65przue46m3rujqih7kqa"
)

var sum float64 = 0.0
var sum1 float64 = 0.0

func revMinerWalletInfo(r []RespWallet) {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("充值记录")
	// 设置单元格的值
	f.SetSheetName("Sheet1", "转账记录")
	//设置转账地址(到)

	f.SetCellValue("充值记录", "C1", "到：")
	f.SetCellValue("转账记录", "C1", "到：")
	num2 := 1
	num := 0
	for _, v := range r {
		for _, v := range v.Data {
			if v.To == znyk_workerWallet || v.To == xz_workerWallet {
				num++
				f.SetCellValue("充值记录", "C"+strconv.Itoa(num+1), v.To)
			} else {
				f.SetCellValue("转账记录", "C"+strconv.Itoa(num2+1), v.To)
				num2++
			}

		}
		f.SetCellValue("充值记录", "C"+strconv.Itoa(num+1), "")
		f.SetCellValue("转账记录", "C"+strconv.Itoa(num2+1), "")
	}
	//设置转账地址(从)
	num = 0
	num2 = 1
	f.SetCellValue("充值记录", "B1", "从:")
	f.SetCellValue("转账记录", "B1", "从:")
	for _, v := range r {
		for _, v := range v.Data {
			if v.To == znyk_workerWallet || v.To == xz_workerWallet {
				num++
				f.SetCellValue("充值记录", "B"+strconv.Itoa(num+1), v.From)
			} else {
				f.SetCellValue("转账记录", "B"+strconv.Itoa(num2+1), v.From)
				num2++
			}
			//fmt.Println(i, "B"+strconv.Itoa(i+1))
		}
		f.SetCellValue("充值记录", "B"+strconv.Itoa(num+1), "累计：")
		f.SetCellValue("转账记录", "B"+strconv.Itoa(num2+1), "累计：")
	}
	//设置金额
	num = 0
	num2 = 1
	f.SetCellValue("充值记录", "D1", "金额")
	f.SetCellValue("转账记录", "D1", "金额")
	for _, v := range r {
		for _, v := range v.Data {
			if v.To == znyk_workerWallet || v.To == xz_workerWallet {
				num++
				vTotal := strings.Split(v.Value, "FIL")

				for _, v := range vTotal {
					v = strings.TrimSpace(v)
					res, _ := strconv.ParseFloat(v, 64)
					sum += res

				}

				f.SetCellValue("充值记录", "D"+strconv.Itoa(num+1), v.Value)
			} else {
				vTotal := strings.Split(v.Value, "FIL")

				for _, v := range vTotal {
					v = strings.TrimSpace(v)
					res, _ := strconv.ParseFloat(v, 64)
					sum1 += res

				}

				f.SetCellValue("转账记录", "D"+strconv.Itoa(num2+1), v.Value)
				num2++
			}
		}

		f.SetCellValue("充值记录", "D"+strconv.Itoa(num+1), sum)
		f.SetCellValue("转账记录", "D"+strconv.Itoa(num2+1), sum1)

	}
	//设置时间：
	num = 0
	num2 = 1
	f.SetCellValue("充值记录", "E1", "时间：")
	f.SetCellValue("转账记录", "E1", "时间：")
	for _, v := range r {
		for _, v := range v.Data {
			if v.To == znyk_workerWallet || v.To == xz_workerWallet {
				num++
				f.SetCellValue("充值记录", "E"+strconv.Itoa(num+1), v.TimeFormat)
				//	fmt.Println(i, "B"+strconv.Itoa(i+1))
			} else {
				f.SetCellValue("转账记录", "E"+strconv.Itoa(num2+1), v.TimeFormat)
				num2++
			}
		}
		f.SetCellValue("充值记录", "E"+strconv.Itoa(num+1), "")
		f.SetCellValue("转账记录", "E"+strconv.Itoa(num2+1), "")
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件

	if err := f.SaveAs(r[0].Name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

func CreateExcel(r []RespWallet) {
	revMinerWalletInfo(r)

}
