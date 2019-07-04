package sn

import (
	"bytes"
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

var cSeqInSecondForShort int64	// 某一秒内的临时计数器，反映某一秒内产生sn的序号
var preTimestampForShort string

func init() {
	cSeqInSecond = 0
	cSeqInSecondForShort = 0
	cSeqInSecondForLight = 0
}

func getSequenceNumForShort() (string, error) {
	if preTimestampForShort != getTimeStamp() {
		cSeqInSecondForShort = 0
	}
	sSeqInSecond := ""  // 接受cSeqInSecond
	sCAPACITY := ""	// 接受SECOND_CAPACITY
	if cSeqInSecondForShort < SECOND_CAPACITY {
		atomic.AddInt64(&cSeqInSecondForShort, 1)
		sSeqInSecond = fmt.Sprintf("%-v", cSeqInSecondForShort)
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
生成分布式场景下15位的唯一序列号(可以使用到2022年5月21日)
terminal: 终端识别号
bizCode: 业务编号
machineCode: 机器编号
*/
func GenShortCode(terminal string, bizCode string, machineCode string) (string, error) {
	var code bytes.Buffer
	code.WriteString(terminal) // 终端(1位)
	code.WriteString(bizCode) // 业务号(1位)
	code.WriteString(machineCode) // 机器编号容量99（2位）
	code.WriteString(getTimeStampAndTagForShort()) // 时间戳8位
	//code.WriteString(fmt.Sprintf("%-v", getPID())) // 进程号（3位数字）
	sQInSec, err := getSequenceNumForShort()
	if err != nil {
		return "", err
	}
	code.WriteString(fmt.Sprintf("%-v", sQInSec))// 一个时间单位的容量999（3位数字）
	//code.WriteString(fmt.Sprintf("%-v", genRnd2()))// 随机数(3位数)
	return code.String(), nil
}

func getTimeStampAndTagForShort() string {
	surplusTime := fmt.Sprintf("%-v", (time.Date(2022,5,21,0,0,0,0,&time.Location{}).Unix() - time.Now().Unix()))
	for len(surplusTime) < 8 {
		surplusTime = "0" + surplusTime
	}

	preTimestampForShort = surplusTime
	return preTimestampForShort
}