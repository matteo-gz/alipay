package alipay

type SubMerchant struct {
	MerchantId   string `json:"merchant_id,omitempty"`
	MerchantType string `json:"merchant_type,omitempty"`
}
type SettleInfo struct {
	SettleDetailInfos []*SettleDetailInfo `json:"settle_detail_infos,omitempty"`
}
type SettleDetailInfo struct {
	TransInType string `json:"trans_in_type,omitempty"`
	Amount      string `json:"amount,omitempty"`
}

// 2022-9-26
