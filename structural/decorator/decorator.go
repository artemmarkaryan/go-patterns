package main

import (
	"fmt"
	"log"
	"time"
)

type Printable interface {
	/*
		Целевой оборачиваемый интерфейс
	*/
	Print() string
}

type LogMessage struct {} // Целевой оборачиваемый класс

func (m LogMessage) Print() string {
	return "log"
}

type LogMessageWithData struct {
	// Декторатор. Добавляет функционал, не меняя интерфейс
	payload Printable
}

func (m LogMessageWithData) Print() string {
	return fmt.Sprintf("(%v) -> %v", time.Now().Format(time.ANSIC), m.payload.Print())
}

type LogMessageWithEmoji struct {
	// Декторатор. Добавляет функционал, не меняя интерфейс
	payload Printable
}

func (m LogMessageWithEmoji) Print() string {
	return fmt.Sprintf("(😎) -> %v", m.payload.Print())
}



func main() {
	message := LogMessage{}

	datedMessage := LogMessageWithData{payload: message}
	datedEmojiMessage := LogMessageWithEmoji{payload: datedMessage}
	log.Print(datedEmojiMessage.Print())
}
