package services

import (
	"martpay/app/request"
	"martpay/app/utils"
)

type Spout struct {
}

func (s Spout) GetName() string {
	return "Spout"
}

func (s Spout) PurchaseAirtime(txRef string, data request.NewTnxInterface) utils.MyResponse {
	return utils.MyResponse{
		Success: true,
		Message: "",
	}
}
