package services

import "time"

type LibraryService interface {
	BookReturn(dueDate, submitDate time.Time) int
}

type libraryService struct {
}

func NewLibraryService() *libraryService {
	return &libraryService{}
}

func (s *libraryService) BookReturn(dueDate, submitDate time.Time) (result int) {
	result = 0
	if submitDate.Year() > dueDate.Year() {
		result = 12000
		return
	} else if (submitDate.Year() == dueDate.Year()) && (submitDate.Month() > dueDate.Month()) {
		result = (int(submitDate.Month()) - int(dueDate.Month())) * 500
		return
	} else if (submitDate.Year() == dueDate.Year()) && (submitDate.Month() == dueDate.Month()) && (submitDate.Day() > dueDate.Day()) {
		result = 15 * (submitDate.Day() - dueDate.Day())
		return
	}
	return
}
