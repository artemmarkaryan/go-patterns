package main

import (
	"log"
	os "os"
)

type IDataPrinter interface {
	GetData() string

	// PrintData — шаблонный метод
	PrintData()
}

type DataPrinter struct {
	// Базовый класс. Тут определён только PrintData
	dp IDataPrinter
}

func (receiver DataPrinter) PrintData() {
	info := receiver.dp.GetData()
	log.Print(info)
}

type FileDataPrinter struct {
	// Расширяющий класс класс. Тут определён GetData
	DataPrinter
	Filename string
}

func (receiver FileDataPrinter) GetData() string {
	bytes, err := os.ReadFile(receiver.Filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

type StringDataPrinter struct {
	// Расширяющий класс класс. Тут определён GetData
	DataPrinter
	Filename string
}

func (receiver StringDataPrinter) GetData() string {
	return "string"
}



func main() {
	fileDataPrinter := FileDataPrinter{
		Filename: "/Users/artemmarkaryan/Docs/Work/web/pickles/server/.gitignore",
	}
	stringDataPrinter := StringDataPrinter{}
	printer := DataPrinter{dp: fileDataPrinter}
	printer2 := DataPrinter{dp: stringDataPrinter}
	printer.PrintData()
	printer2.PrintData()
}
