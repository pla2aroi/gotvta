package gotvta

var (
	buy = "BUY"
	strongBuy = "STRONG_BUY"
	sell = "SELL"
    strongSell = "STRONG_SELL"
    neutral = "NEUTRAL"
)

// ComputeMA returns technical analysis for moving averages.
func ComputeMA(MA float64, close float64) string {
	if MA < close {
		return buy
	} else if MA > close {
		return sell
	} else {
		return neutral
	}
}

// ComputeRSI returns technical analysis for relative strength index.
func ComputeRSI(RSI float64, RSI1 float64) string {
	if RSI < 30 && RSI1 > RSI {
		return buy
	} else if RSI > 70 && RSI1 < RSI {
		return sell
	} else {
		return neutral
	}
}

// ComputeStoch returns technical analysis for stochastic.
func ComputeStoch(stochK float64, stochD float64, stochK1 float64, stochD1 float64) string {
	if stochK < 20 && stochD < 20 && stochK > stochD && stochK1 < stochD1 {
		return buy
	} else if stochK > 80 && stochD > 80 && stochK < stochD && stochK1 > stochD1 {
		return sell
	} else {
		return neutral
	}
}

// ComputeCCI20 returns technical analysis for commodity channel index 20.
func ComputeCCI20(CCI20 float64, CCI201 float64) string {
	if CCI20 < -100 && CCI20 > CCI201 {
		return buy
	} else if CCI20 > 100 && CCI20 < CCI201 {
        return sell
	} else {
		return neutral
	}
}

// ComputeADX returns technical analysis for average directional index.
func ComputeADX(ADX float64, ADXpDI float64, ADXnDI float64, ADXpDI1 float64, ADXnDI1 float64) string {
	if ADX > 20 && ADXpDI1 < ADXnDI1 && ADXpDI > ADXnDI {
		return buy
	} else if ADX > 20 && ADXpDI1 > ADXnDI1 && ADXpDI < ADXnDI {
		return sell
	} else {
		return neutral
	}
}

// ComputeAO returns technical analysis for awesome oscillator.
func ComputeAO(AO float64, AO1 float64) string {
	if AO > 0 && AO1 < 0 || AO > 0 && AO1 > 0 && AO > AO1 {
        return buy
	} else if (AO < 0 && AO1 > 0 || AO < 0 && AO1 < 0 && AO < AO1) {
        return sell
	} else {
		return neutral
	}
}

// ComputeMom returns technical analysis for momentum.
func ComputeMom(mom float64, mom1 float64) string {
	if mom < mom1 {
        return sell
	} else if mom > mom1 {
		return buy
	} else {
		return neutral
	}
}

// ComputeMACD returns technical analysis for moving average convergence/divergence.
func ComputeMACD(MACD float64, signal float64) string {
	if MACD > signal {
		return buy
	} else if MACD < signal {
		return sell
	} else {
		return neutral
	}
}

// ComputeBBBuy returns technical analysis for bull bear buy.
func ComputeBBBuy(close float64, BBLower float64) string {
	if close < BBLower {
		return buy
	}
	return neutral
}

// ComputeBBSell returns technical analysis for bull bear sell.
func ComputeBBSell(close float64, BBUpper float64) string {
	if close > BBUpper {
		return sell
	}
	return neutral
}

// ComputePSAR returns technical analysis for parabolic stop-and-reverse.
func ComputePSAR(PSAR float64, open float64) string {
	if PSAR < open {
		return buy
	} else if PSAR > open {
		return sell
	} else {
		return neutral
	}
}

// ComputeRecommend returns technical analysis for recommend.
func ComputeRecommend(value float64) string {
	if value >= -1 && value < -.5 {
		return strongSell
	} else if value >= -.5 && value < 0 {
		return sell
	} else if value > 0 && value <= .5 {
		return buy
	} else if value > .5 && value <= 1 {
		return strongBuy
	} else {
		return neutral
	}
}

// ComputeSimple returns technical analysis for simple.
func ComputeSimple(value float64) string {
	if value == -1 {
		return sell
	} else if value == 1 {
		return buy
	} else {
		return neutral
	}
}

