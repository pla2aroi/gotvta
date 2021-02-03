package gotvta

import (
	"testing"
)

func TestGetAnalysisBTCUSDT(t *testing.T) {
	// Symbol = Bitcoin/TetherUSD
	// Exchange = Binance
	// Screener = Cryptocurrency (crypto)
	// Interval = 1 day (1d)
	GetAnalysis("BTCUSDT", "binance", "crypto", "1d")
}
