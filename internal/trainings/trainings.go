package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	dataSlice := strings.Split(datastring, ",")
	if len(dataSlice) != 3 {
		return fmt.Errorf("invalid input format: expected 3 values, got %d: %s", len(dataSlice), datastring)
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil {
		return fmt.Errorf("invalid step count: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("step count must be greater than zero")
	}
	t.Steps = steps

	t.TrainingType = dataSlice[1]

	duration, err := time.ParseDuration(dataSlice[2])
	if err != nil {
		return fmt.Errorf("invalid duration format: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}
