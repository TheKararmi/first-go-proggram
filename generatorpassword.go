package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Генератор паролей ===")

	var passwordLength int
	fmt.Print("Введите желаемую длину пароля: ")
	fmt.Scan(&passwordLength)

	if passwordLength < 6 {
		fmt.Println("Предупреждение: Слишком короткий пароль! Установлена длина 6.")
		passwordLength = 6
	}

	chars := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%^&*-_=+"

	password := make([]byte, passwordLength)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	fmt.Printf("Твой новый пароль: %s\n", string(password))
}
