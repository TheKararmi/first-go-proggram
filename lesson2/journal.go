package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	
)

type Entry struct {
	Text string `json:"text"`
}

func main2() {
	fileName := "person.json"
	fmt.Println("=== МОЙ ДНЕВНИК НАСТРОЕНИЯ ===")
	fmt.Println("1. Добавить запись")
	fmt.Println("2. Просмотреть историю")
	fmt.Println("3. Выход")

	var choice int
	fmt.Print("Выбери действие: ")
	fmt.Scanln(&choice)

if choice == 1 {
    var entries []Entry
    fileData, _ := os.ReadFile(fileName)
    json.Unmarshal(fileData, &entries)

    fmt.Println("Напиши свою запись:")
    
 
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        text := scanner.Text() 

        newEntry := Entry{Text: text}
        entries = append(entries, newEntry)

        newData, _ := json.MarshalIndent(entries, "", "  ")
        os.WriteFile(fileName, newData, 0644)
        fmt.Println("Запись сохранена!")
    }
}else if choice == 2 {

		fileData, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("История пока пуста.")
			return
		}

		var entries []Entry
		json.Unmarshal(fileData, &entries)

		fmt.Println("\n--- ТВОЯ ИСТОРИЯ ---")
		for i, entry := range entries {
			fmt.Printf("%d. %s\n", i+1, entry.Text)
		}

	} else if choice == 3 {
		return
	} else {
		fmt.Println("Ошибка!")
	}
}

