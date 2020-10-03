package stockrate

type (
	// StocksInfo holds the information related to datasource of stocks
	StocksInfo map[string]stockURLValue

	// StockPrice holds current price, previous close, open, variation, percentage, volume of stocks from BSE and NSE
	StockPrice struct {
		BSE symbolPriceValue
		NSE symbolPriceValue
	}

	// StockTechnicals holds stock technical valuations
	StockTechnicals map[string]technicalValue

	// StockMovingAverage holds Moving average for 5, 10, 15, 20, 50, 100, 200 days respectively
	StockMovingAverage map[int]movingAverageValue

	// StockPivotLevels stores a stock pivote levels
	StockPivotLevels map[string]pivotPointsValue

	stockURLValue struct {
		Sector  string
		Company string
		Symbol  string
	}

	symbolPriceValue struct {
		Price         float64
		PreviousClose float64
		Open          float64
		Variation     float64
		Percentage    float64
		Volume        int64
	}

	technicalValue struct {
		Level      float64
		Indication string
	}

	movingAverageValue struct {
		SMA        float64
		Indication string
	}

	pivotPointsValue struct {
		R1    float64
		R2    float64
		R3    float64
		Pivot float64
		S1    float64
		S2    float64
		S3    float64
	}
)
