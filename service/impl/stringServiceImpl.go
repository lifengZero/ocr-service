package impl

import (
	"errors"
	"fmt"
	"ocr-service/config"
	"ocr-service/entity"
	"strings"
)

type StringService struct{}

var ErrEmpty = errors.New("empty string")

func (StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	var user = entity.ParentMerchantKey{}
	config.Db.First(&user)
	fmt.Println(user.Id)
	fmt.Println(user.ParentMerchantNo)
	return strings.ToUpper(s), nil
}

func (StringService) Count(s string) int {
	return len(s)
}
