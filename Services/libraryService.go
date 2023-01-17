package services

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type LibraryService interface {
	BookReturn(dueDate, submitDate string) (int, error)
}

type libraryService struct {
}

func NewLibraryService() *libraryService {
	return &libraryService{}
}

func (s *libraryService) BookReturn(dueDate, submitDate string) (int, error) {
	result := 0
	dueDateSplit := strings.Split(dueDate, "-")
	if len(dueDateSplit) != 3 {
		return result, errors.New("pastikan format duedate adalah dd-mm-yyyy")
	}
	submitDateSplit := strings.Split(submitDate, "-")
	if len(submitDateSplit) != 3 {
		return result, errors.New("pastikan format submitdate adalah dd-mm-yyyy")
	}

	dueYear, err := strconv.Atoi(dueDateSplit[2])
	if err != nil {
		return result, errors.New("tahun harus berupa angka")
	}
	dueMonth, err := strconv.Atoi(dueDateSplit[1])
	if err != nil {
		return result, errors.New("bulan harus berupa angka")
	}
	dueDay, err := strconv.Atoi(dueDateSplit[0])
	if err != nil {
		return result, errors.New("tanggal harus berupa angka")
	}
	submitYear, err := strconv.Atoi(submitDateSplit[2])
	if err != nil {
		return result, errors.New("tahun harus berupa angka")
	}
	submitMonth, err := strconv.Atoi(submitDateSplit[1])
	if err != nil {
		return result, errors.New("bulan harus berupa angka")
	}
	submitDay, err := strconv.Atoi(submitDateSplit[0])
	if err != nil {
		return result, errors.New("tanggal harus berupa angka")
	}

	submitDateTime := time.Date(submitYear, time.Month(submitMonth), submitDay, 1, 1, 1, 1, time.Local)
	dueDateTime := time.Date(dueYear, time.Month(dueMonth), dueDay, 1, 1, 1, 1, time.Local)
	if submitDateTime.Year() > dueDateTime.Year() {
		result = 12000
		return result, nil
	} else if (submitDateTime.Year() == dueDateTime.Year()) && (submitDateTime.Month() > dueDateTime.Month()) {
		result = (int(submitDateTime.Month()) - int(dueDateTime.Month())) * 500
		return result, nil
	} else if (submitDateTime.Year() == dueDateTime.Year()) && (submitDateTime.Month() == dueDateTime.Month()) && (submitDateTime.Day() > dueDateTime.Day()) {
		result = 15 * (submitDateTime.Day() - dueDateTime.Day())
		return result, nil
	}
	return result, nil
}
