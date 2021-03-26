package raft

import (
	//"encoding/gob"
	"fmt"
)

type LogEntry struct {
	Command interface{}
	Term int
}

type Logs struct {
	Log []*LogEntry
	Offset int
	IncludedTerm int
}

func NewLogs() *Logs {
	logs := &Logs{Log: make([]*LogEntry, 0)}
	return logs
}

//return -1 if invalid
func LogIndexToArrayIndex(logIndex int) int {
	arrayIndex := logIndex - 1
	return arrayIndex
}

func toString(entries []*LogEntry, offset int) string {

	res := ""
	for i, log := range entries {
		res += fmt.Sprintf("[Index: %d Term: %d, %+v]",i + 1 + offset , log.Term, log.Command)
	}

	return res
}

func (lgs *Logs) LogIndexToArrayIndex(logIndex int) int {

	arrayIndex := LogIndexToArrayIndex(logIndex - lgs.Offset)
	return arrayIndex
}

func (lgs *Logs) GetEntriesByIndex(from int, to int) []*LogEntry {

	from = from - 1 - lgs.Offset
	//if from < 0

	if to == -1 {
		to = len(lgs.Log)
	} else {
		to = to - 1 - lgs.Offset
	}
	slice := lgs.Log[from: to]
	return slice
}
/*
func (lgs* Logs) GetLogEntries(from int, to int) []*LogEntry{

	slice := lgs.GetEntriesByIndex(from, to)

	//entries := make([]*LogEntry, len(slice))
	//for i, entry :=  range slice {
	//	entries[i] = entry
	//}
	return slice///entries
}
*/
func (lgs *Logs) GetEntriesFromLastByIndex(offset int) (int, *LogEntry, int) {

	len := len(lgs.Log)
	index := len - offset - 1
	term := lgs.IncludedTerm
	var result *LogEntry = nil

	if index >= 0 {
		result = lgs.Log[index]
		if result != nil {
			term = result.Term
		}
	}
	return index + 1 + lgs.Offset, result, term

}

func (lgs *Logs) GetEntryByIndex(index int) (int, *LogEntry, int) {

	index -= lgs.Offset

	if index <= 0 {
		return 0, nil, 0
	}

	len := len(lgs.Log)
	term := 0
	index--
	var result *LogEntry = nil

	if index < len {
		result = lgs.Log[index]
		if result != nil {
			term = result.Term
		}
	}
	return index + 1, result, term

}


func (lgs *Logs) ToString() string {

	return toString(lgs.Log, lgs.Offset)
}

func (lgs *Logs) AppendEntries(entries ...*LogEntry) {
	lgs.Log = append(lgs.Log, entries...)
}

func (lgs *Logs) FirstEntry() (int, *LogEntry, int) {
	if len(lgs.Log) > 0 {
		return lgs.Offset, lgs.Log[0], lgs.Log[0].Term
	}

	return lgs.Offset, nil, lgs.IncludedTerm
}

func (lgs *Logs) FirstIndex() int {

	return lgs.Offset
}
/*
func (lgs *Logs) AppendInterface(interfaces ...interface{}) {

	lgs.Log = append(lgs.Log, interfaces)
}
*/
func (lgs *Logs) ReplaceEntriesFrom(entries []*LogEntry, offset int,  discard bool) {
	if discard {
		lgs.Log = append(lgs.Log[0:offset-lgs.Offset], entries...)
	} else {

	}
}

func (lgs *Logs) GetEntriesFromByIndex(from int) []*LogEntry {

	arrayIndex := lgs.LogIndexToArrayIndex(from)
	if arrayIndex < 0 {
		panic("array index negative")
	}
	return lgs.Log[arrayIndex:len(lgs.Log)]
}

func (lgs *Logs) Clear() {
	lgs.Log = lgs.Log[:0]
}

func (lgs* Logs) Len() int {

	return len(lgs.Log)
}

func (lgs* Logs) LastIndex() int {
	return lgs.Len() + lgs.Offset
}

func (lgs* Logs) GetEntry(index int) *LogEntry{
	return lgs.Log[index - lgs.Offset]
}

func (lgs* Logs) GetEntryByLogIndex(index int) *LogEntry{
	arrayIndex := lgs.LogIndexToArrayIndex(index)
	return lgs.Log[arrayIndex]
}


//return false index already get discarded
func (lgs* Logs)DiscardBefore(index int) bool {

	if lgs.Offset < index {
		arrayIndex := lgs.LogIndexToArrayIndex(index)
		lgs.Offset = index
		lgs.IncludedTerm = lgs.Log[arrayIndex].Term
		lgs.Log = lgs.Log[arrayIndex + 1:len(lgs.Log)]

		return true
	} else {
		return false
	}
}



