package common

import (
	"bufio"
	"os"
	"strings"

	"CyberGuardian/logger"

	"gopkg.in/ini.v1"
)

func GetHosts(szHostsFile string) []string {
	szFuncName := "GetHosts"
	szHosts := []string{}
	file, err := os.Open(szHostsFile)
	if err != nil {
		logger.ERROR("<" + szFuncName + "> -> " + err.Error())
	}
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
	szResult := map[string]string{}
	cfg, err := ini.Load(szConfigFile)
	if err != nil {
		logger.ERROR("無法讀取 config.cfg: " + err.Error())
	}

	szServer := cfg.Section("SMTP").Key("server").String()
	szPort := cfg.Section("SMTP").Key("port").String()
	szSender := cfg.Section("SMTP").Key("sender").String()
	szPassword := cfg.Section("SMTP").Key("password").String()
	szReceiver := cfg.Section("SMTP").Key("receiver").String()

	szSubject := cfg.Section("Mail").Key("subject").String()
	szBody := cfg.Section("Mail").Key("body").String()

	szResult["server"] = szServer
	szResult["port"] = szPort
	szResult["sender"] = szSender
	szResult["password"] = szPassword
	szResult["receiver"] = szReceiver
	szResult["subject"] = szSubject
	szResult["body"] = szBody

	return szResult
}
