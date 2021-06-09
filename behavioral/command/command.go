package main

import "log"

type Computer struct{}

func (r *Computer) ToggleOn() string {
	return "Toggle On"
}

func (r *Computer) ToggleOff() string {
	return "Toggle Off"
}

type Command interface {
	/*
		Вынесем исполнение команд компутера в отдельный интерфейс,
		чтобы их можно было удобно хранить и вызывать
	*/
	Execute() string
	Unexecute() string
}

// ToggleOnCommand — команда включения компутера
type ToggleOnCommand struct {
	receiver *Computer
}

func (c ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}
func (c ToggleOnCommand) Unexecute() string {
	return c.receiver.ToggleOff()
}

// ToggleOffCommand — команда выключения компутера
type ToggleOffCommand struct {
	receiver *Computer
}

func (c ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}
func (c ToggleOffCommand) Unexecute() string {
	return c.receiver.ToggleOn()
}

type InvokerCommand struct {
	command  Command
	executed bool
}

type Invoker struct {
	// Инвокер хранит и выполняет команды
	index    int
	commands []InvokerCommand
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, InvokerCommand{command, false})
}

func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) ExecuteOne() string {
	defer func() { i.index ++ }()
	i.commands[i.index].executed = true
	return i.commands[i.index].command.Execute()
}

func (i *Invoker) UndoOne() string {
	i.index--
	i.commands[i.index].executed = false
	return i.commands[i.index].command.Unexecute()
}

func main() {
	computer := Computer{}
	invoker := Invoker{}
	invoker.StoreCommand(ToggleOnCommand{&computer})
	invoker.StoreCommand(ToggleOnCommand{&computer})
	invoker.StoreCommand(ToggleOnCommand{&computer})

	log.Print(invoker.commands)
	log.Print(invoker.ExecuteOne())
	log.Print(invoker.ExecuteOne())
	log.Print(invoker.commands)
	log.Print(invoker.UndoOne())
	log.Print(invoker.commands)

}