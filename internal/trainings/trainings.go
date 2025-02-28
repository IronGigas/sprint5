package trainings

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	val := strings.Split(datastring, ",")
	if len(val) != 3 {return errors.New("В строке не 3 элемента!\n")}

	steps, err := strconv.Atoi(val[0])
    if err != nil {
        return errors.New("Не удалось получить количество шагов из строки\n")
    }
	t.Steps = steps

	if val[1] != "Бег" && val[1] != "Ходьба" {
		return errors.New("Не удалось получить известный тип тренировки из строки\n")
	}
	t.TrainingType = val[1]

	duration, err := time.ParseDuration(val[2])
	if err != nil {
        return errors.New("Не удалось получить время тренировки из строки\n")
    }
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() string {
	distance:=spentenergy.Distance(t.Steps)
	meanSpeed:=spentenergy.MeanSpeed(t.Steps, t.Duration)
	spentCalories:=0.0
	switch t.TrainingType {
	case "Бег":
		spentCalories = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		spentCalories = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "Неизвестный тип тренировки"
	}

	str := `Тип тренировки: ` + t.TrainingType + "\n" +
    	`Длительность: ` + fmt.Sprintf("%.2f", t.Duration.Hours())  + " ч.\n" +
    	`Дистанция: ` + fmt.Sprintf("%.2f", distance) + " км.\n" +
    	`Скорость: ` + fmt.Sprintf("%.2f", meanSpeed) + " км/ч\n" +
    	`Сожгли калорий: ` + fmt.Sprintf("%.2f", spentCalories)	+ "\n"	
	return str
}

