package sn

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestGenCode2(t *testing.T) {
	var data []string = []string{}
	fmt.Println(time.Now().Unix())
	for i := 0; i < 999999; i++ {
		unionid, err := GenCode(TERMINAL_WECHAT_MINIPROGRAM, CROWDSOURCING_CONTRACT, "99")
		if err != nil {

		} else {
			data = append(data, unionid)
		}
	}
	fmt.Println(time.Now().Unix())

	hasTheSame := false
	tempData := ""
	for i := 0; i < len(data); i++ {
		if data[i] == tempData {
			hasTheSame = true
			break
		}
		tempData = data[i]
	}
	fmt.Println("是否有重复的：", hasTheSame)
}

func TestGenCode(t *testing.T) {
	var data []string = []string{}
	fmt.Println(time.Now().Unix())
	for i := 0; i < 9999999; i++ {
		unionid, err := GenCode(TERMINAL_WECHAT_MINIPROGRAM, CROWDSOURCING_CONTRACT, "99")
		if err != nil {

		} else {
			data = append(data, unionid)
		}
	}
	fmt.Println(time.Now().Unix())

	fmt.Println("产生的id号数量：", len(data))

	hasTheSame := false
	tempData := ""
	for i := 0; i < len(data); i++ {
		if data[i] == tempData {
			hasTheSame = true
			break
		}
		tempData = data[i]
	}
	fmt.Println("是否有重复的：", hasTheSame)

	lenType := []int {}
	noTrimData := []string {}
	for i := 0; i < len(data); i++ {
		hasSameTrim := false
		for j := 0; j < len(lenType); j++ {
			if len(data[i]) == lenType[j] {
				hasSameTrim = true
				break
			}
		}
		if hasSameTrim == false {
			noTrimData = append(noTrimData, data[i])
			lenType = append(lenType, len(data[i]))
		}
	}
	if len(lenType) > 1 || len(lenType) == 0 {
		fmt.Println("是否整齐：长度有", len(lenType), "种", "分别是：", lenType)
		fmt.Println("发现不同长度的第一个数据为:", noTrimData)
	} else {
		fmt.Println("数据全部整齐,长度为：", lenType[0])
	}

	fmt.Println("第一个是：", data[0])
	fmt.Println("最后一个：", data[len(data) - 1])

	for i := 0; i < 100; i++ {
		fmt.Println(data[i])
	}
}

func TestGenShortCode(t *testing.T) {
	var data []string = []string{}
	fmt.Println(time.Now().Unix())
	for i := 0; i < 99999999; i++ {
		unionid, err := GenShortCode(TERMINAL_WECHAT_MINIPROGRAM, CROWDSOURCING_CONTRACT, "99")
		if err != nil {

		} else {
			data = append(data, unionid)
		}
	}
	fmt.Println(time.Now().Unix())

	fmt.Println("产生的id号数量：", len(data))

	hasTheSame := false
	tempData := ""
	for i := 0; i < len(data); i++ {
		if data[i] == tempData {
			hasTheSame = true
			break
		}
		tempData = data[i]
	}
	fmt.Println("是否有重复的：", hasTheSame)

	lenType := []int {}
	noTrimData := []string {}
	for i := 0; i < len(data); i++ {
		hasSameTrim := false
		for j := 0; j < len(lenType); j++ {
			if len(data[i]) == lenType[j] {
				hasSameTrim = true
				break
			}
		}
		if hasSameTrim == false {
			noTrimData = append(noTrimData, data[i])
			lenType = append(lenType, len(data[i]))
		}
	}
	if len(lenType) > 1 || len(lenType) == 0 {
		fmt.Println("是否整齐：长度有", len(lenType), "种", "分别是：", lenType)
		fmt.Println("发现不同长度的第一个数据为:", noTrimData)
	} else {
		fmt.Println("数据全部整齐,长度为：", lenType[0])
	}

	fmt.Println("第一个是：", data[0])
	fmt.Println("最后一个：", data[len(data) - 1])

	for i := 0; i < 100; i++ {
		fmt.Println("data", data[i])
	}
}

func TestGenShortCode2(t *testing.T) {
	unionid, err := GenShortCode(TERMINAL_WECHAT_MINIPROGRAM, CROWDSOURCING_CONTRACT, "99")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println(unionid)
}


func TestGetMac(t *testing.T) {
	GetMac()
}

func TestGetTimestamp(t *testing.T) {
	fmt.Println(GetTimestamp())
	fmt.Println(len(fmt.Sprintf("%-v", GetTimestamp())))
	fmt.Println(os.Getegid())

	fmt.Println(time.Now().Unix())
	fmt.Println((time.Date(2025,1,1,0,0,0,0,&time.Location{}).Unix() - time.Now().Unix()))
}

//func TestGenRnd(t *testing.T) {
//	fmt.Println(GenRnd())
//}