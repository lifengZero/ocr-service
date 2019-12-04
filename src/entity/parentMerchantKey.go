package entity

import "time"

type ParentMerchantKey struct {
	Id               uint16
	ParentMerchantNo string

	aesKey string

	bizSysNo string

	status int

	createDate time.Time

	pdateDate time.Time

	ips string
}

func (ParentMerchantKey) TableName() string  {
	return "tbl_parent_merchant_key"
}