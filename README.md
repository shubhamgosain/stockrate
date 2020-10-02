# Stockrate

Stockrate is a golang package for reading price and necessary valuations for symbols listed in indian stock markets NSE and BSE. It can provide the information for all the stocks listed at BSE and NSE.
 
Package can provides the list of companies, price of a symbol, moving averages, pivot Levels and technical valuations. For more information, visit its documentation https://pkg.go.dev/github.com/shubhamgosain/stockrate  

# Functions

# Import

   import "github.com/shubhamgosain/stockrate"

# GetCompanyList

GetCompanyList returns the list of all the companies tracked via stockrate. It include companies listed at BSE and NSE markets and is re tracked whenever package is imported.

Eg. 
   fmt.Println(stockrate.GetCompanyList())
   
   [sunstar soft acewin agriteck avikem resin bharat agri dhan..............]

Note: Stockrate web scrape the requested symbol information from website https://www.moneycontrol.com/. The returned valuations are very much dependent on the symbols listings from https://www.moneycontrol.com/india/stockpricequote/ and thus cannot claim on its preciseness. 
