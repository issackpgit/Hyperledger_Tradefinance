package main

import (
	"encoding/json"
	"fmt"
)

//BL ...
type BL struct {
	SCAC                                 string
	BL_NO                                int
	BOOKING_NO                           int
	EXPORT_REFERENCES                    string
	SVC_CONTRACT                         string
	ONWARD_INLAND_ROUTING                string
	SHIPPER_NAME_ADDRESS                 string
	CONSIGNEE_NAME_ADDRESS               string
	VESSEL                               string
	VOYAGE_NO                            int
	PORT_OF_LOADING                      string
	PORT_OF_DISCHARGE                    string
	PLACE_OF_RECEIPT                     string
	PLACE_OF_DELIVERY                    string
	Rows                                 []Row
	FREIGHT_AND_CHARGES                  int
	RATE                                 int
	UNIT                                 int
	CURRENCY                             string
	PREPAID                              string
	TOTAL_CONTAINERS_RECEIVED_BY_CARRIER int
	CONTAINER_NUMBER                     string
	PLACE_OF_ISSUE_OF_BL                 string
	NUMBER_AND_SEQUENCE_OF_ORIGINAL_BLS  string
	DATE_OF_ISSUE_OF_BL                  string
	DECLARED_VALUE                       int
	SHIPPER_ON_BOARD_DATE                string
	SIGNED_BY                            string
	LC_NUMBER                            string
	DATE_OF_PRESENTATION                 string
}

//Row ...
type Row struct {
	DESCRIPTION_OF_GOODS string
	WEIGHT               int
	MEASUREMENT          int
}

func main() {
	bl := BL{
		SCAC:                   "A",
		BL_NO:                  101,
		BOOKING_NO:             112,
		EXPORT_REFERENCES:      "B",
		SVC_CONTRACT:           "C",
		ONWARD_INLAND_ROUTING:  "D",
		SHIPPER_NAME_ADDRESS:   "E",
		CONSIGNEE_NAME_ADDRESS: "F",
		VESSEL:                 "G",
		VOYAGE_NO:              32,
		PORT_OF_LOADING:        "H",
		PORT_OF_DISCHARGE:      "I",
		PLACE_OF_RECEIPT:       "J",
		PLACE_OF_DELIVERY:      "K",
		Rows: []Row{
			{DESCRIPTION_OF_GOODS: "NUCLEAR_REACTOR", WEIGHT: 90, MEASUREMENT: 102},
		},
		FREIGHT_AND_CHARGES: 4500,
		RATE:                51,
		UNIT:                60,
		CURRENCY:            "USD",
		PREPAID:             "Y",
		TOTAL_CONTAINERS_RECEIVED_BY_CARRIER: 1,
		CONTAINER_NUMBER:                     "MSCU 120870-8",
		PLACE_OF_ISSUE_OF_BL:                 "SINGAPORE",
		NUMBER_AND_SEQUENCE_OF_ORIGINAL_BLS:  "3",
		DATE_OF_ISSUE_OF_BL:                  "21/02/2015",
		DECLARED_VALUE:                       8000,
		SHIPPER_ON_BOARD_DATE:                "24/02/2015",
		SIGNED_BY:                            "MAERSK",
		LC_NUMBER:                            "L960477",
		DATE_OF_PRESENTATION:                 "08/16/2016",
	}

	//fmt.Println(invoice)

	b, err := json.Marshal(bl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
