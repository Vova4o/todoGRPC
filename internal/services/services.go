package services

import (
	"time"

	"github.com/Vova4o/todogrpc/internal/models"
	"github.com/Vova4o/todogrpc/pkg/datecalc"
)

func NextDateRequest(nowRequest string, task models.DBTask) (string, error) {
	// move to service module?
	timeNow, err := time.Parse("20060102", nowRequest)
	if err != nil {
		return "", err
	}

	newDate, err := datecalc.NextDate(timeNow, task.Date, task.Repeat)
	if err != nil {
		return "", err
	}

	return newDate, nil
}
