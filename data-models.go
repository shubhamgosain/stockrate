package stockrate

type (
	stockURLValue struct {
		Sector  string
		Company string
		Symbol  string
	}

	// StocksInfo holds the information related to datasource for companies
	StocksInfo map[string]stockURLValue

	// StockPrice holds current price, previous close, open, variation, percentage, volume for stock for BSE and NSE
	StockPrice struct {
		BSE symbolPriceValue
		NSE symbolPriceValue
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

	// StockTechnicals holds stock technical analysis
	StockTechnicals map[string]technicalValue

	movingAverageValue struct {
		SMA        float64
		Indication string
	}

	// StockMovingAverage holds Moving average for 5,10,15,20,50,100,200 days
	StockMovingAverage map[int]movingAverageValue

	pivotPointsValue struct {
		R1    float64
		R2    float64
		R3    float64
		Pivot float64
		S1    float64
		S2    float64
		S3    float64
	}

	// StockPivotLevels stores pivote levels for various types
	StockPivotLevels map[string]pivotPointsValue
)
