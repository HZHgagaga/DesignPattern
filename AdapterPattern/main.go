package main

import "fmt"

type DataLog interface {
	Data(string) string
}

type GameDataLog struct{}

func (g *GameDataLog) Data(str string) string {
	log := fmt.Sprintf("GameDataLog:%s", str)
	fmt.Println(log)
	return log
}

type InfoLog interface {
	Info(string) string
}

type GameInfoLog struct{}

func (g *GameInfoLog) Info(str string) string {
	log := fmt.Sprintf("GameInfoLog:%s", str)
	fmt.Println(log)
	return log
}

//对象适配器
type DataLogAdapter struct {
	dataLog *GameDataLog
}

func NewDataLogAdapter(log *GameDataLog) InfoLog {
	return &DataLogAdapter{
		dataLog: log,
	}
}

func (d *DataLogAdapter) Info(str string) string {
	return d.dataLog.Data(str)
}

//类适配器
type DataLogAdapterC struct {
	GameDataLog
	GameInfoLog
}

func (d *DataLogAdapterC) Data(str string) string {
	return d.GameDataLog.Data(str)
}

func (d *DataLogAdapterC) Info(str string) string {
	return d.GameDataLog.Data(str)
}

func SendDataCenter(log InfoLog, str string) {
	fmt.Println("Send info log center:", log.Info(str))
}

func main() {
	iLog := &GameInfoLog{}
	SendDataCenter(iLog, "ilog")
	dLog := &GameDataLog{}
	dLogAdapter := NewDataLogAdapter(dLog)
	SendDataCenter(dLogAdapter, "dlog")
	dLogAdapterC := &DataLogAdapterC{}
	SendDataCenter(dLogAdapterC, "dlogC")
}
