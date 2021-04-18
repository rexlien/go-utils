package raft

import (
	//"encoding/gob"
	"bytes"
	"encoding/gob"
	"fmt"
)

type LogEntry struct {
	Command interface{}
	Term int
}

type Logs struct {
	Log          []*LogEntry
	offset       int
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

	arrayIndex := LogIndexToArrayIndex(logIndex - lgs.offset)
	return arrayIndex
}

func (lgs *Logs) GetEntriesByIndex(from int, to int) []*LogEntry {

	from = from - 1 - lgs.offset
	//if from < 0

	if to == -1 {
		to = len(lgs.Log)
	} else {
		to = to - 1 - lgs.offset
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
	return index + 1 + lgs.offset, result, term

}

func (lgs *Logs) GetEntryByIndex(index int) (int, *LogEntry, int) {

	if index < lgs.FirstIndex() {
		return 0,nil, 0
	}
	arrayIndex := index - lgs.offset - 1
	if arrayIndex < 0 {
		return 0, nil, 0
	}

	term := 0
	var result *LogEntry = nil
	result = lgs.Log[arrayIndex]
	term = result.Term

	return index, result, term

}


func (lgs *Logs) ToString() string {

	return toString(lgs.Log, lgs.offset)
}

func (lgs *Logs) AppendEntries(entries ...*LogEntry) {
	lgs.Log = append(lgs.Log, entries...)
}

func (lgs *Logs) FirstEntry() (int, *LogEntry, int) {
	if len(lgs.Log) > 0 {
		return lgs.offset + 1, lgs.Log[0], lgs.Log[0].Term
	}

	return lgs.offset + 1, nil, -1
}

func (lgs *Logs) PrevEntry(index int) (int, *LogEntry, int) {

	prevIndex := index - 1
	if prevIndex < 0 {
		panic("index out of bound")
	}

	if prevIndex == lgs.offset {
		return lgs.offset, nil, lgs.IncludedTerm
	}


	arrayIndex := lgs.LogIndexToArrayIndex(prevIndex)
	return prevIndex, lgs.Log[arrayIndex], lgs.Log[arrayIndex].Term

}

func (lgs *Logs) FirstIndex() int {
	//if len(lgs.Log) > 0 {
	return lgs.offset + 1
	//}
	//return lgs.offset
}

func (lgs *Logs) FirstTerm() int {

	if len(lgs.Log) > 0 {
		return lgs.Log[0].Term
	}
	return -1
}

func (lgs *Logs) Offset() int {
	return lgs.offset
}

func (lgs *Logs) SetOffset(offset int) {
	lgs.offset = offset
}
/*
func (lgs *Logs) AppendInterface(interfaces ...interface{}) {

	lgs.Log = append(lgs.Log, interfaces)
}
*/
func (lgs *Logs) ReplaceEntriesFrom(entries []*LogEntry, offset int,  discard bool) {

	if discard {
		lgs.Log = append(lgs.Log[0:offset-lgs.offset], entries...)
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
	return lgs.Len() + lgs.offset
}

func (lgs* Logs) GetEntry(index int) *LogEntry{
	return lgs.Log[index - lgs.offset]
}

func (lgs* Logs) GetEntryByLogIndex(index int) *LogEntry{
	arrayIndex := lgs.LogIndexToArrayIndex(index)
	return lgs.Log[arrayIndex]
}


//return false index already get discarded
func (lgs* Logs)SnapShot(index int) (bool, []byte, int) {

	if index > lgs.offset  {
		arrayIndex := lgs.LogIndexToArrayIndex(index)
		lgs.offset = index
		lgs.IncludedTerm = lgs.Log[arrayIndex].Term
		lgs.Log = lgs.Log[arrayIndex + 1:len(lgs.Log)]

		return true, lgs.Encode(), lgs.IncludedTerm
	} else {
		return false, nil, -1
	}
}

func (lgs *Logs)Encode() []byte {

	w := new(bytes.Buffer)
	e := gob.NewEncoder(w)

	err := e.Encode(lgs.Offset())//e.Encode(rf.logs.Log)
	if err != nil {
		panic("persist error")
	}
	err = e.Encode(lgs.IncludedTerm)
	if err != nil {
		panic("persist error")
	}
	err = e.Encode(lgs.Log)//e.Encode(rf.logs.Log)
	if err != nil {
		panic("persist error")
	}

	data := w.Bytes()
	return data
}

func (lgs *Logs)Decode(byte []byte) {

	//clear log
	lgs.Log = make([]*LogEntry, 0)

	d := gob.NewDecoder(bytes.NewBuffer(byte))
	err := d.Decode(&lgs.offset)
	if err != nil {
		panic("decode error")
	}
	err = d.Decode(&lgs.IncludedTerm)
	if err != nil {
		panic("decode error")
	}

	err = d.Decode(&lgs.Log)
	if err != nil {
		panic("decode error")
	}


}




