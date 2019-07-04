package sn

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync/atomic"
	"time"
)

// 订单涞源、时间戳、随机码、机器码、进程号、每次进程重启随机自增
const (
	TERMINAL_ANDROID = "0"
	TERMINAL_IOS = "1"
	TERMINAL_WECHAT_MINIPROGRAM = "2"
	TERMINAL_WECHAT_MP = "3"

	SECOND_CAPACITY = 999  // 一秒钟的序号容量
)

// 业务编号
const (
	SERVICE_ORDER = "0"
	CROWDSOURCING_ORDER = "1"
	CROWDSOURCING_CONTRACT = "2"
)

var cSeqInSecond int64	// 某一秒内的临时计数器，反映某一秒内产生sn的序号
var preTimestamp string

func init() {
	cSeqInSecond = 0
	cSeqInSecondForShort = 0
	cSeqInSecondForLight = 0
}

func getSequenceNum() (string, error) {
	if preTimestamp != getTimeStamp() {
		cSeqInSecond = 0
	}
	sSeqInSecond := ""  // 接受cSeqInSecond
	sCAPACITY := ""	// 接受SECOND_CAPACITY
	if cSeqInSecond < SECOND_CAPACITY {
		atomic.AddInt64(&cSeqInSecond, 1)
		sSeqInSecond = fmt.Sprintf("%-v", cSeqInSecond)
		sCAPACITY = fmt.Sprintf("%-v", SECOND_CAPACITY)
		for len(sSeqInSecond) < len(sCAPACITY) {
			sSeqInSecond = "0" + sSeqInSecond
		}
	} else {
		return "", errors.New("每Second的容量已达上限")
	}
	return sSeqInSecond, nil
}

/**
生成分布式场景下18位的唯一序列号(可以使用到2022年5月21日)
terminal: 终端识别号
bizCode: 业务编号
machineCode: 机器编号
*/
func GenCode(terminal string, bizCode string, machineCode string) (string, error) {
	var code bytes.Buffer
	code.WriteString(terminal) // 终端(1位)
	code.WriteString(bizCode) // 业务号(1位)
	code.WriteString(machineCode) // 机器编号容量99（2位）
	code.WriteString(getTimeStampAndTag()) // 时间戳8位
	//code.WriteString(fmt.Sprintf("%-v", getPID())) // 进程号（3位数字）
	sQInSec, err := getSequenceNum()
	if err != nil {
		return "", err
	}
	code.WriteString(fmt.Sprintf("%-v", sQInSec))// 一个时间单位的容量999（3位数字）
	code.WriteString(fmt.Sprintf("%-v", genRnd3()))// 随机数(3位数)
	return code.String(), nil
}

func getTimeStampAndTag() string {
	surplusTime := fmt.Sprintf("%-v", (time.Date(2022,5,21,0,0,0,0,&time.Location{}).Unix() - time.Now().Unix()))
	for len(surplusTime) < 8 {
		surplusTime = "0" + surplusTime
	}

	preTimestamp = surplusTime
	return preTimestamp
}

func getTimeStamp() string {
	surplusTime := fmt.Sprintf("%-v", (time.Date(2022,5,21,0,0,0,0,&time.Location{}).Unix() - time.Now().Unix()))
	for len(surplusTime) < 8 {
		surplusTime = "0" + surplusTime
	}

	return surplusTime
}

func GetTimestamp() string {
	return getTimeStamp()
}

// 获取机器的Mac地址
func GetMac() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", nil
	}

	for _, inter := range netInterfaces {
		mac := inter.HardwareAddr
		fmt.Println(mac, "--", inter.Name, "--", inter.Flags)
	}
	return "", nil
}

// 获取3位随机数
func genRnd3() int64 {
	rnd := 0
	for rnd < 100 || rnd > 999 {
		seed := time.Now().Nanosecond() % 799
		if seed > 140 {
			rnd = rand.Intn(seed)
		}
	}
	return int64(rnd)
}

// 获取2位随机数
func genRnd2() int64 {
	rnd := 0
	for rnd < 100 || rnd > 999 {
		seed := time.Now().Nanosecond() % 79
		if seed > 140 {
			rnd = rand.Intn(seed)
		}
	}
	return int64(rnd)
}

func getPID() string {
	spid := fmt.Sprintf("%-v", os.Getegid())
	for len(spid) < 3 {
		spid = "0" + spid
	}
	return spid
}