package comic

import (
	"errors"
	"math/rand"
	"time"
)

var (
	firstComic = time.Date(1978, time.June, 19, 0, 0, 0, 0, time.UTC)
)

type TimeRange struct {
	Start time.Time
	End time.Time
}

func (timeRange TimeRange) validateTimeRange() error {
	if timeRange.Start.Before(firstComic) {
		return errors.New("start date is before first comic date")
	}

	if timeRange.End.Before(timeRange.Start) {
		return errors.New("the end date is before the start date")
	}

	return nil
}

func DefaultRange() TimeRange {
	return TimeRange{
		Start: firstComic,
		End: time.Now(),
	}
}

func GenerateRandomDate(timeRange TimeRange) (randomDate time.Time, err error) {
	validationError := timeRange.validateTimeRange()
	if validationError != nil {
		return time.Time{}, validationError
	}

	startTimestamp := timeRange.Start.Unix()
	endTimestamp := timeRange.End.Unix()
	delta := endTimestamp - startTimestamp

	sec := rand.Int63n(delta) + startTimestamp
	return time.Unix(sec, 0), nil
}