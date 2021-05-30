package decorator

import (
	"fmt"
	"time"
)

type Printable interface {
	Print()
}

type LogMessage struct {
	payload string
}

func (m LogMessage) Print() {
	fmt.Print(m.payload)
}

type LogMessageWithData struct {
	payload LogMessage
}

func (m LogMessageWithData) Print() {
	m.payload = LogMessage{
		payload: fmt.Sprintf("%v: %v", time.Now().Format(time.ANSIC), m.payload),
	}
	fmt.Print(m.payload)
}

func main() {
	message := LogMessage{"example"}
	datedMessage := LogMessageWithData{payload: message}
	datedMessage.Print()
}
