package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Конвертер валют")
	window.Resize(fyne.NewSize(500, 400))

	rates, err := getRatesData("USD")
	if err != nil {
		window.SetContent(widget.NewLabel("Ошибка сети: " + err.Error()))
		window.ShowAndRun()
		return
	}

	var currencies []string
	for k := range rates {
		currencies = append(currencies, k)
	}
	sort.Strings(currencies)

	labelResult := widget.NewLabel("Выберите валюты для расчета")
	labelResult.TextStyle = fyne.TextStyle{Bold: true}

	selectFrom := widget.NewSelectEntry(currencies)
	selectFrom.SetText("EUR")

	selectTo := widget.NewSelectEntry(currencies)
	selectTo.SetText("USD")

	
	doConvert := func() {
		val1 := strings.ToUpper(selectFrom.Text)
		val2 := strings.ToUpper(selectTo.Text)

		course1, ok1 := rates[val1]
		course2, ok2 := rates[val2]

		if !ok1 || !ok2 {
			labelResult.SetText("Ошибка: неверный код валюты")
			return
		}

		finalRate := course2 / course1
		labelResult.SetText(fmt.Sprintf("Курс %s к %s: %.4f", val1, val2, finalRate))
	}


	selectFrom.OnSubmitted = func(s string) {
		doConvert()
	}
	selectTo.OnSubmitted = func(s string) {
		doConvert()
	}


	btnConvert := widget.NewButton("Рассчитать курс", doConvert)

	content := container.NewVBox(
		widget.NewLabel("Из валюты:"),
		selectFrom,
		widget.NewLabel("В валюту:"),
		selectTo,
		btnConvert,
		labelResult,
	)

	window.SetContent(content)
	window.ShowAndRun()
}

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
	data.Rates[baseCurrency] = 1.0
	return data.Rates, nil
}
