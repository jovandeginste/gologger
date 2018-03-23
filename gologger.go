package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"strconv"
	"strings"
)

func main() {
	var priorityFlag, programName string
	var priority syslog.Priority

	flag.StringVar(&priorityFlag, "priority", "a", "syslog priority")
	flag.StringVar(&priorityFlag, "p", "USER.NOTICE", "syslog priority")
	flag.StringVar(&programName, "tag", "gologger", "syslog tag")
	flag.StringVar(&programName, "t", "gologger", "syslog tag")
	flag.Parse()

	if value, err := strconv.Atoi(priorityFlag); err == nil {
		priority = syslog.Priority(value)
	} else {
		facilityPriority, severityPriority, err := parsePriority(priorityFlag)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		priority = facilityPriority | severityPriority
	}

	argsWithoutProg := strings.Join(flag.Args(), " ")

	log.SetFlags(0)
	logwriter, e := syslog.New(priority, programName)
	if e == nil {
		log.SetOutput(logwriter)
	}

	if argsWithoutProg != "" {
		log.Print(argsWithoutProg)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		log.Print(text)
	}
}

func parsePriority(priorityFlag string) (syslog.Priority, syslog.Priority, error) {
	var facilityFlag, severityFlag string
	splitPriorityFlag := strings.SplitN(priorityFlag, ".", 2)
	if len(splitPriorityFlag) > 0 {
		facilityFlag = strings.ToUpper(splitPriorityFlag[0])
	} else {
		facilityFlag = "USER"
	}
	if len(splitPriorityFlag) > 1 {
		severityFlag = strings.ToUpper(splitPriorityFlag[1])
	} else {
		severityFlag = "NOTICE"
	}

	priorityMap := map[string]syslog.Priority{
		"EMERG":   syslog.LOG_EMERG,
		"ALERT":   syslog.LOG_ALERT,
		"CRIT":    syslog.LOG_CRIT,
		"ERR":     syslog.LOG_ERR,
		"WARNING": syslog.LOG_WARNING,
		"NOTICE":  syslog.LOG_NOTICE,
		"INFO":    syslog.LOG_INFO,
		"DEBUG":   syslog.LOG_DEBUG,

		"KERN":     syslog.LOG_KERN,
		"USER":     syslog.LOG_USER,
		"MAIL":     syslog.LOG_MAIL,
		"DAEMON":   syslog.LOG_DAEMON,
		"AUTH":     syslog.LOG_AUTH,
		"SYSLOG":   syslog.LOG_SYSLOG,
		"LPR":      syslog.LOG_LPR,
		"NEWS":     syslog.LOG_NEWS,
		"UUCP":     syslog.LOG_UUCP,
		"CRON":     syslog.LOG_CRON,
		"AUTHPRIV": syslog.LOG_AUTHPRIV,
		"FTP":      syslog.LOG_FTP,

		"LOCAL0": syslog.LOG_LOCAL0,
		"LOCAL1": syslog.LOG_LOCAL1,
		"LOCAL2": syslog.LOG_LOCAL2,
		"LOCAL3": syslog.LOG_LOCAL3,
		"LOCAL4": syslog.LOG_LOCAL4,
		"LOCAL5": syslog.LOG_LOCAL5,
		"LOCAL6": syslog.LOG_LOCAL6,
		"LOCAL7": syslog.LOG_LOCAL7,
	}

	facilityPriority, ok := priorityMap[facilityFlag]
	if !ok {
		return 0, 0, fmt.Errorf("could not find facility %s", facilityFlag)
	}
	severityPriority, ok := priorityMap[severityFlag]
	if !ok {
		return 0, 0, fmt.Errorf("could not find severity %s", severityFlag)
	}

	return facilityPriority, severityPriority, nil
}
