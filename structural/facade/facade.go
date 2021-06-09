package main

import "log"

type UserService struct{}

func (s UserService) Delete(id string) error {
	/*
		Сложный код, делающий много проверок в базе данных.
		Если произошла ошибка, вернёт ошибку. Если нет — вернет nil (null)
	*/

	return nil
}

type UserView struct{}

func (view UserView) Delete(id string) {
	/*
		Простой код, который удаляет объект из бд, не заморачиваясь о реализации
	*/

	err := UserService{}.Delete(id)
	if err != nil {
		log.Print(err.Error())
	}
}

func main() {
	view := UserView{}
	view.Delete("10")
}
