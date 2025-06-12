package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	dataInput := strings.Split(data, ",")
	if len(dataInput) != 3 {
		err := errors.New("ошибка ввода")
		return 0, "", 0, err
	}
	stepsNumbers, err := strconv.Atoi(dataInput[0])
	if err != nil {
		return 0, "", 0, err
	}
	if stepsNumbers <= 0 {
		return 0, "", 0, errors.New("")
	}

	t, err := time.ParseDuration(dataInput[2])
	if t <= 0 {
		return 0, "", 0, errors.New("")

	}
	if err != nil {
		return 0, "", 0, err
	}
	return stepsNumbers, dataInput[1], t, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activityMode, t, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	runMode, err := RunningSpentCalories(steps, weight, height, t)
	if err != nil {
		return "", err
	}
	walkMode, err := WalkingSpentCalories(steps, weight, height, t)
	if err != nil {
		return "", err
	}
	switch activityMode {
	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityMode, float64(t.Hours()), distance(steps, height), meanSpeed(steps, height, t), runMode), nil
	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityMode, float64(t.Hours()), distance(steps, height), meanSpeed(steps, height, t), walkMode), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		var errInput = errors.New("")

		return 0, errInput
	}
	durationInMinutes := duration.Minutes()
	calcCalories := (weight * meanSpeed(steps, height, duration) * durationInMinutes) / minInH
	return calcCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("")
	}
	durationInMinutes := duration.Minutes()
	calcCalories := (weight * meanSpeed(steps, height, duration) * durationInMinutes) / minInH * walkingCaloriesCoefficient
	return calcCalories, nil
}
