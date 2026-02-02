package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JournalEntry struct {
	Date  string `json:"date"`
	Mood  string `json:"mood"`
	Notes string `json:"notes,omitempty"`
}

func main() {
	//блок вывода текста

	fmt.Println("📖  МОЙ ДНЕВНИК НАСТРОЕНИЯ")
	fmt.Println("========================")
	fmt.Println("1. Добавить запись")
	fmt.Println("2. Просмотреть историю")
	fmt.Println("3. Выход")
	//блок создание переменных
	var choice int
	fmt.Print("Выбери действие: ")
	fmt.Scan(&choice)
	// блок логики программы
	switch choice {
	case 1:
		addEntry()
	case 2:
		viewHistory()
	case 3:
		fmt.Println("До свидания!")
		os.Exit(0)
	default:
		fmt.Println("Неверный выбор")
	}
}

// создает файл формата .json и сохраняет его
func addEntry() {
    filename := "journal.json"
    var entries []JournalEntry

    // Читаем старые записи
    fileData, _ := os.ReadFile(filename)
    json.Unmarshal(fileData, &entries)

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Какое у тебя сегодня настроение? ")
    if scanner.Scan() {
        mood := scanner.Text()
        if mood == "" {
             scanner.Scan()
             mood = scanner.Text()
        }

        fmt.Print("Как прошел день? ")
        scanner.Scan()
        notes := scanner.Text()

        newEntry := JournalEntry{
            Date:  time.Now().Format("2006-01-02"),
            Notes: notes,
            Mood:  mood,
        }

        entries = append(entries, newEntry)
        newData, _ := json.MarshalIndent(entries, "", "  ")
        os.WriteFile(filename, newData, 0644)
        fmt.Println("Запись сохранена!")
    }
}

// позволяет просмотреть историю прошлых записей
func viewHistory() {
	data, err := os.ReadFile("journal.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("История пуста. Сделай первую запись!")
			return
		}
		panic(err)
	}
	var entries []JournalEntry
	json.Unmarshal(data, &entries)
	for i, entry := range entries {
			fmt.Printf("%d. %s\n", i+1, &entry)
		}

	fmt.Println("\n=== ВСЕ ЗАПИСИ ===")
	fmt.Println(string(data))
}
