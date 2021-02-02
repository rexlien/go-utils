package raft

import (
	"fmt"
)

type LogEntry struct {

	Command interface{}
	Term int

}


type Logs struct {

	Log []interface{}

}

func NewLogs() *Logs {
	logs := &Logs{Log: make([]interface{}, 0)}
	return logs
}

func LogEntryToString(entries []*LogEntry) string {

	res := ""
	for _, log := range entries {
		res += fmt.Sprintf("[Term: %d, %+v]", log.Term, log.Command)
	}
	return res
}

func ToString(entries []interface{}) string {

	res := ""
	for _, log := range entries {
		logEntry := log.(*LogEntry)

		res += fmt.Sprintf("[Term: %d, %+v]", logEntry.Term, logEntry.Command)
	}

	return res
}



func (lgs *Logs) GetEntries(from int, to int) []interface{} {

	from = from - 1

	if to == -1 {
		to = len(lgs.Log)
	} else {
		to = to - 1
	}
	slice := lgs.Log[from: to]
	return slice
}

func (lgs* Logs) GetLogEntries(from int, to int) []*LogEntry{

	slice := lgs.GetEntries(from, to)

	entries := make([]*LogEntry, len(slice))
	for i, entry :=  range slice {
		entries[i] = entry.(*LogEntry)
	}
	return entries
}

func (lgs *Logs) GetLogFromLast(offset int) (int, *LogEntry, int) {

	len := len(lgs.Log)
	index := len - offset - 1
	term := 0
	var result *LogEntry = nil

	if index >= 0 {
		result = lgs.Log[index].(*LogEntry)
		if result != nil {
			term = result.Term
		}
	}
	return index + 1, result, term

}

func (lgs *Logs) GetLogFromIndex(index int) (int, *LogEntry, int) {

	if index <= 0 {
		return 0, nil, 0
	}

	len := len(lgs.Log)
	term := 0
	index--
	var result *LogEntry = nil

	if index < len {
		result = lgs.Log[index].(*LogEntry)
		if result != nil {
			term = result.Term
		}
	}
	return index + 1, result, term

}


func (lgs *Logs) ToString() string {

	return ToString(lgs.Log)
}