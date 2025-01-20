//go:generate goversioninfo
package main

import (
	"fmt"
	"sync"
	"time"

	"CyberGuardian/common"
	"CyberGuardian/logger"
	"CyberGuardian/network"
	"CyberGuardian/notify"
)

func checkHost(szHost string, szMailConfig map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	if !network.ICMP(1, szHost) {
		szCurrentTime := time.Now()
		szNow := szCurrentTime.Format("2006-01-02 15:04:05")

		szMessage := szHost + " 設備無回應，請檢查網絡連接！"
		logger.WARN(szMessage)

		szBody := fmt.Sprintf("時間 : %s\n設備IP : %s\n使用檢測工具 : ping\n", szNow, szHost)
		szBody += szMailConfig["body"]

		szResult := szHost + " 設備警告電子郵件發送成功"
		szResultFail := szHost + " 設備警告電子郵件發送失敗"
		notify.MailNotify(szMailConfig["server"], szMailConfig["port"], szMailConfig["sender"], szMailConfig["password"], szMailConfig["receiver"], szMailConfig["subject"], szBody, szResult, szResultFail)
	} else {
		szMessage := szHost + " 設備正常在線"
		logger.INFO(szMessage)
	}
}

func main() {
	szHosts := common.GetHosts("hosts.txt")
	szMailConfig := common.GetMailConfig("config.cfg")
	for {
		var wg sync.WaitGroup
		for _, szHost := range szHosts {
			wg.Add(1)
			go checkHost(szHost, szMailConfig, &wg)
		}
		wg.Wait()
		logger.INFO("-------------------- 分隔線 --------------------")
		time.Sleep(30 * time.Second)

	}
}
