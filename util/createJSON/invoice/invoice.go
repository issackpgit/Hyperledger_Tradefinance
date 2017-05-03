package main

import (
	"encoding/json"
	"fmt"
)

//State ...
type Invoice struct {
	PAYER                string
	PAYEE                string
	TAX_REGISTRY_NO      int
	INVOICE_CODE         int
	INVOICE_NUMBER       int
	PRINTING_NO          int
	Rows                 []Row
	TOTAL_IN_WORDS       string
	TOTAL_IN_FIGURES     int
	PRINT_NO             int
	ANTI_FORGERY_CODE    string
	DATE_ISSUED          string
	LC_NUMBER            string
	DATE_OF_PRESENTATION string
	CURRENCY             string
}

//Row ...
type Row struct {
	//ID             string //`json:"id" bson:"id"`
	SERVICE        string
	ITEM           int
	AMOUNT_CHARGED int
	REMARKS        string
}

func main() {
	invoice := Invoice{
		PAYER:           "A",
		PAYEE:           "B",
		TAX_REGISTRY_NO: 1,
		INVOICE_CODE:    2,
		INVOICE_NUMBER:  3,
		PRINTING_NO:     4,
		Rows: []Row{
			{SERVICE: "A", ITEM: 2, AMOUNT_CHARGED: 40, REMARKS: "A"},
			{SERVICE: "B", ITEM: 21, AMOUNT_CHARGED: 60, REMARKS: "A"},
		},
		TOTAL_IN_WORDS:       "ONE HUNDRED",
		TOTAL_IN_FIGURES:     100,
		PRINT_NO:             5,
		ANTI_FORGERY_CODE:    "ABCD",
		DATE_ISSUED:          "21/12/2014",
		LC_NUMBER:            "L960477",
		DATE_OF_PRESENTATION: "08/16/2016",
	}

	//fmt.Println(invoice)

	b, err := json.Marshal(invoice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
