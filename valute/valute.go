package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv" // Добавили для конвертации строки в число
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var currencyToCountry = map[string]string{
	"USD": "US", "EUR": "EU", "GBP": "GB", "JPY": "JP", "AUD": "AU",
	"CAD": "CA", "CHF": "CH", "CNY": "CN", "HKD": "HK", "NZD": "NZ",
	"SEK": "SE", "KRW": "KR", "SGD": "SG", "NOK": "NO", "MXN": "MX",
	"INR": "IN", "RUB": "RU", "ZAR": "ZA", "TRY": "TR", "BRL": "BR",
	"TWD": "TW", "DKK": "DK", "PLN": "PL", "THB": "TH", "IDR": "ID",
	"HUF": "HU", "CZK": "CZ", "ILS": "IL", "PHP": "PH", "AED": "AE",
}

func getFlag(currency string) string {
	countryCode, ok := currencyToCountry[currency]
	if !ok {
		return "🏳️"
	}
	var flag strings.Builder
	for _, r := range strings.ToUpper(countryCode) {
		flag.WriteRune(r + 127397)
	}
	return flag.String()
}

type CurrencyResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Конвертер валют Pro")
	window.Resize(fyne.NewSize(400, 500))

	currentRates, _ := getRatesData("USD")

	var options []string
	for code := range currencyToCountry {
		options = append(options, fmt.Sprintf("%s %s", getFlag(code), code))
	}
	sort.Strings(options)

	labelResult := widget.NewLabel("Введите сумму и нажмите Рассчитать")
	labelResult.Alignment = fyne.TextAlignCenter

	
	inputAmount := widget.NewEntry()
	inputAmount.SetPlaceHolder("Например: 100")
	inputAmount.SetText("1") 
	selectFrom := widget.NewSelect(options, nil)
	selectTo := widget.NewSelect(options, nil)

	selectBase := widget.NewSelect(options, func(selected string) {
		baseCode := selected[len(selected)-3:]
		newRates, err := getRatesData(baseCode)
		if err == nil {
			currentRates = newRates
			labelResult.SetText("Данные обновлены")
		} else {
			labelResult.SetText("Ошибка обновления данных")
		}
	})

	selectBase.SetSelected("🇺🇸 USD")
	selectFrom.SetSelected("🇪🇺 EUR")
	selectTo.SetSelected("🇺🇸 USD")

	doConvert := func() {
		amountStr := inputAmount.Text
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			labelResult.SetText("Ошибка: введите число")
			return
		}

		val1 := selectFrom.Selected[len(selectFrom.Selected)-3:]
		val2 := selectTo.Selected[len(selectTo.Selected)-3:]

		course1, ok1 := currentRates[val1]
		course2, ok2 := currentRates[val2]

		if !ok1 || !ok2 {
			labelResult.SetText("Ошибка: валюта не найдена")
			return
		}

		
		finalResult := (course2 / course1) * amount

		
		labelResult.SetText(fmt.Sprintf("%.2f %s = %.2f %s", amount, val1, finalResult, val2))
	}

	btnConvert := widget.NewButton("Рассчитать", doConvert)

	window.SetContent(container.NewVBox(
		widget.NewLabel("Сумма:"),
		inputAmount,
		widget.NewSeparator(),
		widget.NewLabel("Из валюты:"),
		selectFrom,
		widget.NewLabel("В валюту:"),
		selectTo,
		widget.NewSeparator(),
		btnConvert,
		labelResult,
	))

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
	if data.Rates == nil {
		data.Rates = make(map[string]float64)
	}
	data.Rates[baseCurrency] = 1.0
	return data.Rates, nil
}

