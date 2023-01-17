package services

import "strconv"

type CandyService interface {
	WhoGetSourCandy(student, candies, firstStudent string) (int, error)
}

type candyService struct {
}

func NewCandyService() *candyService {
	return &candyService{}
}

func (s *candyService) WhoGetSourCandy(student, candies, firstStudent string) (int, error) {
	result := 0
	studentInt, err := strconv.Atoi(student)
	if err != nil {
		return result, err
	}
	candiesInt, err := strconv.Atoi(candies)
	if err != nil {
		return result, err
	}
	firstStudentInt, err := strconv.Atoi(firstStudent)
	if err != nil {
		return result, err
	}
	result = (candiesInt + firstStudentInt - 1) % studentInt
	if result == 0 {
		result = studentInt
	}
	return result, nil
}
