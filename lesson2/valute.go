package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type CurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

func main1() {
	// 1. Сначала просто загружаем данные, ничего не спрашивая у пользователя
	rates, err := getRatesData("USD")
	if err != nil {
		fmt.Printf("Ошибка при загрузке данных: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("💱  КОНВЕРТЕР ВАЛЮТ (База: USD)")
	fmt.Println("==================")
	fmt.Println("Выберите действие:")
	fmt.Println("1. Узнать курс определённой валюты в USD")
	fmt.Println("2. Показать список всех валют")
	fmt.Println("3. Узнать курс определенной валюты в выбранной валюте")
	
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		//Спрашиваем валюту ТОЛЬКО здесь
		fmt.Print("Выберите валюту: ")
		var val string
		fmt.Scanln(&val)
		val = strings.ToUpper(val)

		// Используем переменную rates, которую получили в начале main
		course, ok := rates[val]
		if !ok {
			fmt.Printf("Валюта '%s' не найдена.\n", val)
		} else {
			fmt.Printf("Курс USD к %s: %.4f\n", val, course)
		}

	case 2:
		fmt.Println("\nДоступные валюты:")
		for currency := range rates {
			fmt.Printf("- %s\n", currency)
		}
	case 3:
		fmt.Println("Выберите первую валюту:")
		var val1 string 
		fmt.Scanln(&val1)
		val1 = strings.ToUpper(val1)
		course1, ok := rates[val1]
		if !ok {
			fmt.Printf("Валюта '%s' не найдена.\n", val1)
		}
		fmt.Println("Выберите вторую валюту:")
		var val2 string 
		fmt.Scanln(&val2)
		val2 = strings.ToUpper(val2)
		course2, ok := rates[val2]
		if !ok {
			fmt.Printf("Валюта '%s' не найдена.\n", val2)
		}
		
		fmt.Printf("Курс %s к %s: %.4f\n", val1, val2, course1,course2)
	default:
		fmt.Println("Неверный выбор")
	}
}

// Эта функция только забирает данные из API и возвращает их в main
func getRatesData(baseCurrency string) (map[string]float64, error) {
	url := fmt.Sprintf("https://api.frankfurter.app/latest?from=%s", baseCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data CurrencyResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data.Rates, nil
}
