package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Joaovitordebrito/wex-purchase-service/src/model"
)

type purchase struct {
	db *sql.DB
}

func NewPurchaseRepo(db *sql.DB) *purchase {
	return &purchase{db}
}

func (repo purchase) Create(purchase model.Purchase) (uint64, error) {
	statement, err := repo.db.Prepare(
		"insert into purchase (uuid, description, purchase_amount, transaction_date) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(purchase.UUID, purchase.Description, purchase.PurchaseAmount, purchase.TransactionDate)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(rows), nil
}

func (repo purchase) GetConvertedCurrency(uuid string, country string) (model.ConvertedPurchase, error) {

	conn, err := repo.db.Query(
		"select uuid, purchase_amount, transaction_date from purchase where uuid = ?",
		uuid,
	)
	if err != nil {
		return model.ConvertedPurchase{}, err
	}

	defer conn.Close()
	var purchase model.Purchase
	if conn.Next() {
		if err = conn.Scan(
			&purchase.UUID,
			&purchase.PurchaseAmount,
			&purchase.TransactionDate,
		); err != nil {
			return model.ConvertedPurchase{}, err
		}
	}
	convertedCurrency, err := convertCurrency(purchase, country)
	if err != nil {
		return model.ConvertedPurchase{}, err
	}

	return convertedCurrency, nil
}

func convertCurrency(purchase model.Purchase, country string) (model.ConvertedPurchase, error) {
	var response struct {
		Data []struct {
			RecordDate     string `json:"record_date"`
			Country        string `json:"country"`
			Currency       string `json:"currency"`
			CountryDesc    string `json:"country_currency_desc"`
			ExchangeRate   string `json:"exchange_rate"`
			EffectiveDate  string `json:"effective_date"`
			SrcLineNbr     string `json:"src_line_nbr"`
			RecordFiscalYr string `json:"record_fiscal_year"`
		} `json:"data"`
	}

	monthsToSearch := 6
	exchangeRate := 0.0
	for i := 0; i <= monthsToSearch; i++ {
		date, err := time.Parse("2006-01-02", purchase.TransactionDate)
		if err != nil {
			panic(err)
		}
		searchDate := date.AddDate(0, -i, 0)
		lastDayOfMonth := time.Date(
			searchDate.Year(),
			searchDate.Month()+1,
			0,
			0, 0, 0, 0,
			searchDate.Location(),
		)
		searchDateStr := lastDayOfMonth.Format("2006-01-02")

		fmt.Println(searchDateStr)

		requestURL := fmt.Sprintf("https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?filter=country:in:(%s),record_date:eq:%s", country, searchDateStr)
		res, err := http.Get(requestURL)
		if err != nil {
			continue
		}

		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return model.ConvertedPurchase{}, err
		}
		if len(response.Data) > 0 {
			for _, item := range response.Data {
				if item.RecordDate == searchDateStr {
					fmt.Printf("Record Date: %s, Exchange Rate: %s\n", item.RecordDate, item.ExchangeRate)
					floatExchangeRate, err := strconv.ParseFloat(item.ExchangeRate, 64)
					if err != nil {
						return model.ConvertedPurchase{}, err
					}
					exchangeRate = floatExchangeRate
				}
			}
			break
		}
	}

	if len(response.Data) == 0 {
		return model.ConvertedPurchase{}, fmt.Errorf("the purchase cannot be converted to the target currency")
	}

	convertedAmount := purchase.PurchaseAmount * exchangeRate
	formattedConvertedAmount := fmt.Sprintf("%.2f", convertedAmount)
	parsedAndFormattedConvertedAmount, err := strconv.ParseFloat(formattedConvertedAmount, 64)
	if err != nil {
		return model.ConvertedPurchase{}, err
	}

	formattedExchangeRate := fmt.Sprintf("%.2f", exchangeRate)
	parsedAndFormattedExchangeRate, err := strconv.ParseFloat(formattedExchangeRate, 64)
	if err != nil {
		return model.ConvertedPurchase{}, err
	}

	return model.ConvertedPurchase{
		PurchaseAmount:  purchase.PurchaseAmount,
		TargetCurrency:  parsedAndFormattedExchangeRate,
		ConvertedAmount: parsedAndFormattedConvertedAmount,
	}, nil
}
