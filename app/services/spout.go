package services

import (
	"fmt"
	"pay-bills/app/request"
	"pay-bills/app/utils"
)

type Spout struct {
}

func (s Spout) GetName() string {
	return "Spout"
}

func (s Spout) PurchaseAirtime(txRef string, data request.NewTnxInterface) utils.MyResponse {
	airtime := data.(request.AirtimeData)

	fmt.Println(airtime)

	// TODO: Call provider api to make purchase and return response on status.

	return utils.MyResponse{
		Success: true,
		Message: "",
	}
}
