package services

type CandyService interface {
	WhoGetSourCandy(student, candies, firstStudent int) int
}

type candyService struct {
}

func NewCandyService() *candyService {
	return &candyService{}
}

func (s *candyService) WhoGetSourCandy(student, candies, firstStudent int) (result int) {
	result = (candies + firstStudent - 1) % student
	if result == 0 {
		result = student
	}
	return
}
