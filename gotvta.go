package gotvta

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	// Interval1Minute is for 1 minute interval.
	Interval1Minute = "1m"
	// Interval5Minutes is for 5 minutes interval.
	Interval5Minutes = "5m"
	// Interval15Minutes is for 15 minutes interval.
	Interval15Minutes = "15m"
	// Interval1Hour is for 1 hour interval.
	Interval1Hour = "1h"
	// Interval4Hours is for 4 hours interval.
	Interval4Hours = "4h"
	// Interval1Day is for 1 day interval.
	Interval1Day = "1d"
	// Interval1Week is for 1 week interval.
	Interval1Week = "1W"
	// Interval1Month is for 1 month interval.
	Interval1Month = "1M"
)

var (
	// Forex stores exchange for forex.
	Forex = "FX_IDC"
	// CFD stores exchange for CFD.
	CFD = "TVC"
)

var (
	scanURL    = "https://scanner.tradingview.com/%s/scan"
	indicators = []string{"Recommend.Other{}", "Recommend.All{}", "Recommend.MA{}", "RSI{}", "RSI[1]{}", "Stoch.K{}", "Stoch.D{}", "Stoch.K[1]{}", "Stoch.D[1]{}", "CCI20{}", "CCI20[1]{}", "ADX{}", "ADX+DI{}", "ADX-DI{}", "ADX+DI[1]{}", "ADX-DI[1]{}", "AO{}", "AO[1]{}", "Mom{}", "Mom[1]{}", "MACD.macd{}", "MACD.signal{}", "Rec.Stoch.RSI{}", "Stoch.RSI.K{}", "Rec.WR{}", "W.R{}", "Rec.BBPower{}", "BBPower{}", "Rec.UO{}", "UO{}", "close{}", "EMA5{}", "SMA5{}", "EMA10{}", "SMA10{}", "EMA20{}", "SMA20{}", "EMA30{}", "SMA30{}", "EMA50{}", "SMA50{}", "EMA100{}", "SMA100{}", "EMA200{}", "SMA200{}", "Rec.Ichimoku{}", "Ichimoku.BLine{}", "Rec.VWMA{}", "VWMA{}", "Rec.HullMA9{}", "HullMA9{}", "Pivot.M.Classic.S3{}", "Pivot.M.Classic.S2{}", "Pivot.M.Classic.S1{}", "Pivot.M.Classic.Middle{}", "Pivot.M.Classic.R1{}", "Pivot.M.Classic.R2{}", "Pivot.M.Classic.R3{}", "Pivot.M.Fibonacci.S3{}", "Pivot.M.Fibonacci.S2{}", "Pivot.M.Fibonacci.S1{}", "Pivot.M.Fibonacci.Middle{}", "Pivot.M.Fibonacci.R1{}", "Pivot.M.Fibonacci.R2{}", "Pivot.M.Fibonacci.R3{}", "Pivot.M.Camarilla.S3{}", "Pivot.M.Camarilla.S2{}", "Pivot.M.Camarilla.S1{}", "Pivot.M.Camarilla.Middle{}", "Pivot.M.Camarilla.R1{}", "Pivot.M.Camarilla.R2{}", "Pivot.M.Camarilla.R3{}", "Pivot.M.Woodie.S3{}", "Pivot.M.Woodie.S2{}", "Pivot.M.Woodie.S1{}", "Pivot.M.Woodie.Middle{}", "Pivot.M.Woodie.R1{}", "Pivot.M.Woodie.R2{}", "Pivot.M.Woodie.R3{}", "Pivot.M.Demark.S1{}", "Pivot.M.Demark.Middle{}", "Pivot.M.Demark.R1{}"}
)

type dictionary map[string]interface{}

// Analysis stores technical analysis information.
type Analysis struct {
	Exchange       string
	Symbol         string
	Screener       string
	Interval       string
	Time           time.Time
	Summary        dictionary
	Oscillators    dictionary
	MovingAverages dictionary
	Indicators     dictionary
}

//createData returns data to be POSTed to TradingView.
func createData(symbol string, interval string) string {
	data := ""
	switch interval {
	case Interval1Minute:
		data = "|1"
	case Interval5Minutes:
		data = "|5"
	case Interval15Minutes:
		data = "|15"
	case Interval1Hour:
		data = "|60"
	case Interval4Hours:
		data = "|240"
	case Interval1Week:
		data = "|1W"
	case Interval1Month:
		data = "|1M"
	default:
		if interval != "1d" {
			log.Print("Warning: Interval is empty or not valid, defaulting to 1 day.")
		}
	}

	reqColumns := []string{}
	for _, indicator := range indicators {
		reqColumns = append(reqColumns, strings.ReplaceAll(indicator, "{}", data))
	}

	x, err := json.Marshal(reqColumns)
	if err != nil {
		print("Error when marshalling reqColumns.")
	}
	jsonData := fmt.Sprintf("{\"symbols\":{\"tickers\":[\"%s\"],\"query\":{\"types\":[]}},\"columns\":%s}", strings.ToUpper(symbol), string(x))
	return jsonData
}

// GetAnalysis retrieve and compute indicators from TradingView.
func GetAnalysis(symbol string, exchange string, screener string, interval string) (Analysis, error) {
	if screener == "" {
		return Analysis{}, errors.New("screener is empty or not valid")
	} else if exchange == "" {
		return Analysis{}, errors.New("exchange is empty or not valid")
	} else if symbol == "" {
		return Analysis{}, errors.New("symbol is empty or not valid")
	}

	exchangeSymbol := strings.ToUpper(fmt.Sprintf("%s:%s", exchange, symbol))
	data := createData(exchangeSymbol, interval)
	tvScanner := fmt.Sprintf(scanURL, screener)

	client := &http.Client{}
	req, err := http.NewRequest("POST", tvScanner, strings.NewReader(data))
	if err != nil {
		print(err)
	}

	req.Header.Set("User-Agent", "GoTVTA/1.0")

	resp, err := client.Do(req)
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	fmt.Print(string(body))

	// Currently Working: Retrieve data from TradingView.
	// TODO: Unmarshal body, get indicators, compute indicators, remove some log.

	return Analysis{}, nil
}
