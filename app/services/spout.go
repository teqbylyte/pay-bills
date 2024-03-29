package services

import "martpay/app/utils"

type Spout struct {
}

func (s Spout) GetName() string {
	return "Spout"
}

func (s Spout) PurchaseAirtime(amount string, network string, phone string, reference string) utils.MyResponse {
	return utils.MyResponse{}
}
