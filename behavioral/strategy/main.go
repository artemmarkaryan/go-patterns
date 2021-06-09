/*
Стратегия — это поведенческий паттерн проектирования, который определяет
семейство схожих алгоритмов и помещает каждый из них в собственный класс, после
чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
*/

package main

import (
	"fmt"
	"time"
)

type Logger interface {
	log(message string)
}

type LogWithTime struct{}

func (l LogWithTime) log(message string) {
	fmt.Printf("%v: %v\n", time.Now().Format(time.Kitchen), message)
}

type LogWithEmoji struct{}

func (l LogWithEmoji) log(message string) {
	fmt.Printf("😎: %v\n", message)
}

type Process struct {
	// В logger лежит заменяемая логика
	// В зависимости от переданного объекта, логирование происходит по-разному
	logger Logger
}

func (p *Process) SetLogger(l Logger) {
	p.logger = l
}

func (p *Process) DoSth() {
	p.logger.log("I did something")
}

func main() {
	someProcess := Process{}

	someProcess.SetLogger(LogWithTime{})
	someProcess.DoSth()

	someProcess.SetLogger(LogWithEmoji{})
	someProcess.DoSth()
}
