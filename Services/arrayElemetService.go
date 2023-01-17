package services

import (
	"errors"
	"strconv"
)

type ArrayElementService interface {
	SameSumElement(arrLen int, arr []string) (string, error)
}

type arrayElementService struct {
}

func NewArrayElementService() *arrayElementService {
	return &arrayElementService{}
}

func (s *arrayElementService) SameSumElement(arrLen int, arr []string) (string, error) {
	result := ""
	if arrLen != len(arr) {
		return result, errors.New("arrLen tidak sesuai dengan panjang array")
	}
	for index := range arr {
		left := 0
		right := 0

		for i := 0; i < arrLen; i++ {
			temp, err := strconv.Atoi(arr[i])
			if err != nil {
				return result, err
			}
			if i > index {
				right += temp
			}
			if i < index {
				left += temp
			}
		}
		if right == left {
			result = "YES"
			return result, nil
		}
	}
	result = "NO"
	return result, nil
}
