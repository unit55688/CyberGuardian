package common

import (
	"bufio"
	"os"
	"strings"

	"CyberGuardian/logger"

	"gopkg.in/ini.v1"
)

func GetHosts(szHostsFile string) []string {
	if _, err := os.Stat(szHostsFile); err != nil {
		szFile, _ := os.Create(szHostsFile)
		defer szFile.Close()
		logger.Log(szHostsFile, "192.168.1.1")
		logger.Log(szHostsFile, "192.168.1.2")
	}

	szHosts := []string{}
	file, _ := os.Open(szHostsFile)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	// fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		szLines := fileScanner.Text()
		szHosts = append(szHosts, strings.TrimSpace(szLines))

	}
	return szHosts
}

func GetMailConfig(szConfigFile string) map[string]string {
	if _, err := os.Stat(szConfigFile); err != nil {
		szFile, _ := os.Create(szConfigFile)
		defer szFile.Close()
		logger.Log(szConfigFile, "[SMTP]")
		logger.Log(szConfigFile, "server = smtp.gmail.com")
		logger.Log(szConfigFile, "port = 587")
		logger.Log(szConfigFile, "sender = ")
		logger.Log(szConfigFile, "password = ")
		logger.Log(szConfigFile, "receiver = ")
		logger.Log(szConfigFile, "")
		logger.Log(szConfigFile, "[Mail]")
		logger.Log(szConfigFile, "subject = \"設備斷線告警\"")
		logger.Log(szConfigFile, "body = \"設備無回應，請檢查網絡連接！!\"")
		logger.Log(szConfigFile, "")
		logger.Log(szConfigFile, "[Ping]")
		logger.Log(szConfigFile, "round = 30")
		logger.Log(szConfigFile, "count = 3")
		logger.Log(szConfigFile, "timeout = 5")
	}

	szResult := map[string]string{}
	cfg, err := ini.Load(szConfigFile)
	if err != nil {
		logger.ERROR("無法讀取 config.cfg: " + err.Error())
	}

	szMailServer := cfg.Section("SMTP").Key("server").String()
	szMailPort := cfg.Section("SMTP").Key("port").String()
	szMailSender := cfg.Section("SMTP").Key("sender").String()
	szMailPassword := cfg.Section("SMTP").Key("password").String()
	szMailReceiver := cfg.Section("SMTP").Key("receiver").String()

	szMailSubject := cfg.Section("Mail").Key("subject").String()
	szMailBody := cfg.Section("Mail").Key("body").String()

	szPingRound := cfg.Section("Ping").Key("round").String()
	szPingCount := cfg.Section("Ping").Key("count").String()
	szPingTimeout := cfg.Section("Ping").Key("timeout").String()

	szResult["server"] = szMailServer
	szResult["port"] = szMailPort
	szResult["sender"] = szMailSender
	szResult["password"] = szMailPassword
	szResult["receiver"] = szMailReceiver
	szResult["subject"] = szMailSubject
	szResult["body"] = szMailBody
	szResult["round"] = szPingRound
	szResult["count"] = szPingCount
	szResult["timeout"] = szPingTimeout

	return szResult
}
