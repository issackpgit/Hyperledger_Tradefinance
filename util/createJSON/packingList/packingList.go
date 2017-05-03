package main

import (
	"encoding/json"
	"fmt"
)

//PackingList ...
type PackingList struct {
	CONSIGNEE_NAME            string
	CONSIGNEE_ADDRESS         string
	PACKING_LIST_NO           string
	DATE                      string
	Rows                      []Row
	TOTAL_QUANTITY_MTONS      int
	TOTAL_NET_WEIGHT_KGS      int
	TOTAL_GROSS_WEIGHT_KGS    int
	DELIVERY_TERMS            string
	DOCUMENTARY_CREDIT_NUMBER string
	METHOD_OF_LOADING         string
	CONTAINER_NUMBER          string
	PORT_OF_LOADING           string
	PORT_OF_DISCHARGE         string
	LC_NUMBER                 string
	DATE_OF_PRESENTATION      string
}

//Row ...
type Row struct {
	DESCRIPTION_OF_GOODS string
	QUANTITY_MTONS       int
	NET_WEIGHT_KGS       int
	GROSS_WEIGHT_KGS     int
}

func main() {
	packingList := PackingList{
		CONSIGNEE_NAME:    "North American Coating and Painting Co. Ltd.",
		CONSIGNEE_ADDRESS: "Richmond, BC, Canada",
		PACKING_LIST_NO:   "PL-14072014",
		DATE:              "12/09/2014",
		Rows: []Row{
			{DESCRIPTION_OF_GOODS: "Pure Polyyester Powder Coating", QUANTITY_MTONS: 20, NET_WEIGHT_KGS: 20, GROSS_WEIGHT_KGS: 22},
		},
		TOTAL_QUANTITY_MTONS:      20,
		TOTAL_NET_WEIGHT_KGS:      20,
		TOTAL_GROSS_WEIGHT_KGS:    22,
		DELIVERY_TERMS:            "CIF Port Metro Vancouver, Canada Incoterms 2010",
		DOCUMENTARY_CREDIT_NUMBER: "LC140050",
		METHOD_OF_LOADING:         "1X40 HQ CNTR (S)",
		CONTAINER_NUMBER:          "MSCU 120870-8",
		PORT_OF_LOADING:           "PORT OF SHANGHAI, CHINA",
		PORT_OF_DISCHARGE:         "PORT METRO VANCOUVER, CANADA",
		LC_NUMBER:                 "L960477",
		DATE_OF_PRESENTATION:      "08/16/2016",
	}

	//fmt.Println(invoice)

	b, err := json.Marshal(packingList)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
