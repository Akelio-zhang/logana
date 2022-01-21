package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

type dayStat struct {
	Day              string
	ClassStatMap     map[string]*classStat
	ExceptionStatMap map[string]*exceptionStat
}

type classStat struct {
	class       string
	levelCntMap map[string]uint32
}

type exceptionStat struct {
	exception string
	cnt       uint32
}

/**
统计结果
*/
func output(logInfos []LogInfo) {
	dayStatMap := make(map[string]*dayStat)
	for _, logInfo := range logInfos {
		day := getDayString(logInfo.Time)
		val, ok := dayStatMap[day]
		if !ok {
			val = &dayStat{Day: day, ClassStatMap: make(map[string]*classStat), ExceptionStatMap: make(map[string]*exceptionStat)}
			dayStatMap[day] = val
		}
		stat(val, &logInfo)
	}

	logInfosOutput(logInfos)
	for _, dayStat := range dayStatMap {
		classStatOutput(dayStat)
	}

	for _, dayStat := range dayStatMap {
		exceptionOutput(dayStat)
	}
}

func exceptionOutput(dayStat *dayStat) {
	table := tablewriter.NewWriter(os.Stdout)
	header := make([]string, 0)
	header = append(header, "Exception")
	header = append(header, "Count")
	table.SetHeader(header)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT})

	data := make([][]string, 0)
	dayStat.ExceptionStatMap = sortExceptionMap(dayStat.ExceptionStatMap)
	for _, stat := range dayStat.ExceptionStatMap {
		data = append(data, []string{stat.exception, fmt.Sprintf("%d", stat.cnt)})
	}

	for _, v := range data {
		table.Append(v)
	}
	fmt.Printf("EXCEPTION-INFO-STAT|%s\n", dayStat.Day)
	table.Render()
}

func sortExceptionMap(m map[string]*exceptionStat) map[string]*exceptionStat {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	newMap := make(map[string]*exceptionStat)
	for _, k := range keys {
		newMap[k] = m[k]
	}
	return newMap
}

var levelList = []string{"ERROR", "WARN", "INFO", "DEBUG"}

func classStatOutput(dayStat *dayStat) {
	table := tablewriter.NewWriter(os.Stdout)
	header := make([]string, 0)
	header = append(header, "CLASS")
	header = append(header, levelList...)
	header = append(header, "COUNT")
	table.SetHeader(header)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT})

	data := make([][]string, 0)
	dayStat.ClassStatMap = sortClassMap(dayStat.ClassStatMap)
	for _, stat := range dayStat.ClassStatMap {
		levelCntArray := make([]string, 0)
		sum := uint32(0)
		for _, level := range levelList {
			levelCnt, ok := stat.levelCntMap[level]
			if !ok {
				levelCnt = 0
			}
			levelCntArray = append(levelCntArray, fmt.Sprintf("%d", levelCnt))
			sum += levelCnt
		}
		data = append(data, []string{stat.class, levelCntArray[0], levelCntArray[1], levelCntArray[2], fmt.Sprintf("%d", sum)})
	}

	for _, v := range data {
		table.Append(v)
	}
	fmt.Printf("CLASS-INFO-STAT|%s\n", dayStat.Day)
	table.Render()
}

func sortClassMap(m map[string]*classStat) map[string]*classStat {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	newMap := make(map[string]*classStat)
	for _, k := range keys {
		newMap[k] = m[k]
	}
	return newMap
}

func logInfosOutput(logInfos []LogInfo) {
	f, _ := os.Create("statOutput/output.json")
	if f != nil {
		defer f.Close()
		for _, logInfo := range logInfos {
			b, _ := json.Marshal(logInfo)
			fmt.Fprintln(f, string(b))
		}
	} else {
		log.Println("create file failed")
	}
}

func stat(dayStat *dayStat, logInfo *LogInfo) {
	if logInfo.Class != "" {
		val, ok := dayStat.ClassStatMap[logInfo.Class]
		if !ok {
			val = &classStat{class: logInfo.Class, levelCntMap: make(map[string]uint32)}
			dayStat.ClassStatMap[logInfo.Class] = val
		}
		levelCnt, ok := val.levelCntMap[logInfo.Level]
		if !ok {
			levelCnt = 0
		}
		val.levelCntMap[logInfo.Level] = levelCnt + 1
	}

	if logInfo.Exception != "" {
		val, ok := dayStat.ExceptionStatMap[logInfo.Exception]
		if !ok {
			val = &exceptionStat{exception: logInfo.Exception}
			dayStat.ExceptionStatMap[logInfo.Exception] = val
		}
		val.cnt++
	}
}

func getDayString(time string) string {
	return time[0:10]
}
