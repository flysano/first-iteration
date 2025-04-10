package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"personaldata"
	"spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 2 {
		return fmt.Errorf("invalid input format: expected 2 values, got %d: %s", len(dataSlice), datastring)
	}
	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return fmt.Errorf("invalid step count: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("step count must be greater than zero")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(dataSlice[1])
	if err != nil {
		return fmt.Errorf("invalid duration format: %v", err)
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d\nДистанция составила: %.2f\nВы сожгли: %.2f", ds.Steps, distance, calories), nil
}
