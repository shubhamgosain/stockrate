# Stockrate

Stockrate is a golang module for reading price and necessary valuations for symbols listed in indian stock markets NSE and BSE. It can provide the information for all the stocks listed at BSE and NSE.
 
Module can provides the list of companies, price of a symbol, moving averages, pivot Levels and technical valuations. For more information, visit its documentation https://pkg.go.dev/github.com/shubhamgosain/stockrate

---
# Functions

Import
--
    import "github.com/shubhamgosain/stockrate"

GetCompanyList 
--
GetCompanyList returns the list of all the companies tracked via stockrate. It include companies listed at BSE and NSE markets and is re tracked whenever package is imported.

Eg. 

    fmt.Println(stockrate.GetCompanyList())
   
    [--sunstar soft acewin agriteck avikem resin bharat agri dhan..............]
    
GetPrice
--
GetPrice returns current price, previous close, open, variation, percentage and volume for a particular stock. It returns value of type  (StockPrice, error)

Eg. 

    fmt.Println(stockrate.GetPrice("bpcl"))
   
    {353.3 353.05 358 0.25 0.07 16212565}
    
GetTechnicals
--
StockTechnicals holds stock technical analysis like ADX, ATR, StockTechnicals holds stock technical analysis, MACD, MFI, ROC, RSC, RSI etc for a particular stock.

Eg. 

    fmt.Println(stockrate.GetTechnicals("bpcl"))
   
    map[ADX:{18.72 Weak Trend} ATR:{16.83 High Volatility} Bollinger Band:{0 - -} CCI:{-177.14 Bearish} MACD:{-12.48 Bearish} MFI:{15.69 Oversold} ROC:{-14.54 Bearish} RSC :{84.31 Underperformer} RSI:{28.5 Bearish} Stochastic:{6.55 Oversold} Williamson:{-92.93 Oversold}]
    
GetMovingAverage
--
GetMovingAverage returns the 5,10,20,50,100,200 days moving average for a particular stock.

Eg. 

    fmt.Println(stockrate.GetMovingAverage("bpcl"))
   
    map[5:{370.76 Bearish} 10:{382.44 Bearish} 20:{398.26 Bearish} 50:{410.52 Bearish} 100:{389.87 Bearish} 200:{401.21 Bearish}]
  
GetPivotLevels
--
GetPivotLevels returns the important pivot levels of company given in order R1, R2, R3, Pivot, S1, S2, S3 for a particular stock.

Eg. 

    fmt.Println(stockrate.GetPivotLevels("bpcl"))
   
    map[Camarilla:{354.59 355.88 357.16 354.12 352.01 350.72 349.44} Classic:{360.73 368.17 374.78 354.12 346.68 340.07 332.63} Fibonacci:{359.48 362.8 368.17 354.12 348.75 345.43 340.07}]
    
---

Note : Stockrate web scrape the requested symbol information from website https://www.moneycontrol.com/. The returned valuations are very much dependent on the symbols listings from https://www.moneycontrol.com/india/stockpricequote/ and hence cannot claim on its preciseness. 
