package stockrate

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var (
	stocksURL = make(StocksInfo)
	baseURL   = "https://www.moneycontrol.com/technical-analysis"
	sourceURL = "https://www.moneycontrol.com/india/stockpricequote/"
)

// GetCompanyList returns the list of all the companies tracked via stockrate
func GetCompanyList() (list []string) {
	for key := range stocksURL {
		list = append(list, key)
	}
	return
}

// GetPrice returns current price, previous close, open, variation, percentage and volume for a company
func GetPrice(company string) (StockPrice, error) {
	var stockPrice StockPrice
	url, err := getURL(company)
	if err != nil {
		return stockPrice, err
	}
	doc, err := getStockQuote(url)
	if err != nil {
		return stockPrice, fmt.Errorf("Error in reading stock Price")
	}
	doc.Find(".bsedata_bx").Each(func(i int, s *goquery.Selection) {
		stockPrice.BSE.Price, _ = strconv.ParseFloat(s.Find(".span_price_wrap").Text(), 64)
		stockPrice.BSE.PreviousClose, _ = strconv.ParseFloat(s.Find(".priceprevclose").Text(), 64)
		stockPrice.BSE.Open, _ = strconv.ParseFloat(s.Find(".priceopen").Text(), 64)
		stockPrice.BSE.Variation, _ = strconv.ParseFloat(strings.Split(s.Find(".span_price_change_prcnt").Text(), " ")[0], 64)
		stockPrice.BSE.Percentage, _ = strconv.ParseFloat(strings.Split(strings.Split(s.Find(".span_price_change_prcnt").Text(), "%")[0], "(")[1], 64)
		stockPrice.BSE.Volume, _ = strconv.ParseInt(strings.ReplaceAll(s.Find(".volume_data").Text(), ",", ""), 10, 64)
	})
	doc.Find(".nsedata_bx").Each(func(i int, s *goquery.Selection) {
		stockPrice.NSE.Price, _ = strconv.ParseFloat(s.Find(".span_price_wrap").Text(), 64)
		stockPrice.NSE.PreviousClose, _ = strconv.ParseFloat(s.Find(".priceprevclose").Text(), 64)
		stockPrice.NSE.Open, _ = strconv.ParseFloat(s.Find(".priceopen").Text(), 64)
		stockPrice.NSE.Variation, _ = strconv.ParseFloat(strings.Split(s.Find(".span_price_change_prcnt").Text(), " ")[0], 64)
		stockPrice.NSE.Percentage, _ = strconv.ParseFloat(strings.Split(strings.Split(s.Find(".span_price_change_prcnt").Text(), "%")[0], "(")[1], 64)
		stockPrice.NSE.Volume, _ = strconv.ParseInt(strings.ReplaceAll(s.Find(".volume_data").Text(), ",", ""), 10, 64)
	})
	return stockPrice, nil
}

// GetTechnicals returns the technical valuations of a company with indications
func GetTechnicals(company string) (StockTechnicals, error) {
	stockTechnicals := make(StockTechnicals)
	url, err := getURL(company)
	if err != nil {
		return nil, err
	}
	doc, err := getStockQuote(url)
	if err != nil {
		return nil, fmt.Errorf("Error in reading stock Technicals %v", err.Error())
	}
	doc.Find("#techindd").Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		symbol := strings.Split(strings.Split(s.Find("td").First().Text(), "(")[0], "%")[0]
		level, _ := strconv.ParseFloat(strings.ReplaceAll(s.Find("td").Find("strong").First().Text(), ",", ""), 64)
		indication := s.Find("td").Find("strong").Last().Text()
		if symbol != "" && symbol != "Bollinger Band(20,2)" {
			stockTechnicals[symbol] = technicalValue{level, indication}
		}
	})
	return stockTechnicals, nil
}

// GetMovingAverage returns the 5, 10, 20, 50, 100, 200 days moving average respectively
func GetMovingAverage(company string) (StockMovingAverage, error) {
	stockMovingAverage := make(StockMovingAverage)
	url, err := getURL(company)
	if err != nil {
		return nil, err
	}
	doc, err := getStockQuote(url)
	if err != nil {
		return nil, fmt.Errorf("Error in reading stock Moving Averages %v", err.Error())
	}
	doc.Find("#movingavgd").Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		period, _ := strconv.Atoi(s.Find("td").First().Text())
		sma, _ := strconv.ParseFloat(strings.ReplaceAll(s.Find("td").Find("strong").First().Text(), ",", ""), 64)
		indication := s.Find("td").Find("strong").Last().Text()
		if period != 0 {
			stockMovingAverage[period] = movingAverageValue{sma, indication}
		}
	})
	return stockMovingAverage, nil
}

// GetPivotLevels returns the important pivot levels of a stock given in order R1, R2, R3, Pivot, S1, S2, S3
func GetPivotLevels(company string) (StockPivotLevels, error) {
	stockPivotLevels := make(StockPivotLevels)
	url, err := getURL(company)
	if err != nil {
		return nil, err
	}
	doc, err := getStockQuote(url)
	if err != nil {
		return nil, fmt.Errorf("Error in reading stock Pivot Levels %v", err.Error())
	}
	doc.Find("#pevotld").Find("table").First().Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		pivotType := s.Find("td").First().Text()
		if pivotType != "" {
			var levels []float64
			s.Find("td").Next().Each(func(i int, s *goquery.Selection) {
				level, _ := strconv.ParseFloat(strings.ReplaceAll(s.Text(), ",", ""), 64)
				levels = append(levels, level)
			})
			stockPivotLevels[pivotType] = pivotPointsValue{
				levels[0], levels[1], levels[2], levels[3], levels[4], levels[5], levels[6],
			}
		}
	})
	return stockPivotLevels, nil
}

// getStockQuote creates and returns the web document from a web URL
func getStockQuote(URL string) (*goquery.Document, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// getURL checks whether we can read data for company and returns its data source URL
func getURL(company string) (URL string, err error) {
	if val, found := stocksURL[strings.ToLower(company)]; found {
		URL = baseURL + "/" + val.Company + "/" + val.Symbol + "/daily"
		return
	}
	return "", fmt.Errorf("Company Not Found")
}

// Here stocks information necessary is saved and stored, which is calculated everytime package is imported
func init() {
	fmt.Println("Reading stocks")
	capAlphabets := []string{"A", "B", "C", "D", "E", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for _, char := range capAlphabets {
		doc, err := getStockQuote(sourceURL + char)
		if err != nil {
			log.Panic("Error in fetching stock URLs ", err.Error())
		}
		doc.Find(".bl_12").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			stockName := s.Text()
			if match, _ := regexp.MatchString(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`, link); match {
				stockURLSplit := strings.Split(link, "/")
				stocksURL[strings.ToLower(stockName)] = stockURLValue{stockURLSplit[5], stockURLSplit[6], stockURLSplit[7]}
			}
		})
	}
	fmt.Println("Stocks Read Succesfull")
}
