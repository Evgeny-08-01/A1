package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	dataInput := strings.Split(data, ",")
	if len(dataInput) != 2 {
		err := errors.New("incorrect data entry")
		log.Println(err)
		return 0, 0, err
	}
	stepsNumbers, err := strconv.Atoi(dataInput[0])
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if stepsNumbers <= 0 {
		err := errors.New("incorrect data entry on the number of steps")
		log.Println(err)
		return 0, 0, err
	}
	t, err := time.ParseDuration(dataInput[1])
	if t <= 0 {
		err := errors.New("incorrect time data entry")
		log.Println(err)
		return 0, 0, err

	}
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	return stepsNumbers, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	stepsNumbers, t, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if stepsNumbers <= 0 {
		return ""
	}
	distance := float64(stepsNumbers) * stepLength / mInKm

	calories, err := spentcalories.WalkingSpentCalories(stepsNumbers, weight, height, t)
	if err != nil {
		return ""
	}
	stringReturn := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepsNumbers, distance, calories)

	return stringReturn
}
