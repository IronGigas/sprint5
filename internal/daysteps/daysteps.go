package daysteps

import (
	"time"
	"strconv"
	"errors"
	"strings"
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	val := strings.Split(datastring, ",")
	if len(val) != 2 {return errors.New("В строке не 2 элемента \n")}
	
    steps, err := strconv.Atoi(val[0])
    if err != nil {
        return errors.New("Не удалось получить количество шагов из строки\n")
    }
	ds.Steps=steps

	duration, err := time.ParseDuration(val[1])
	if err != nil {
        return errors.New("Не удалось получить время тренировки из строки\n")
    }
	ds.Duration=duration
    
    return nil
}

func (ds DaySteps) ActionInfo() string {
	steps := strconv.Itoa(ds.Steps)
	distance:=spentenergy.Distance(ds.Steps)
	spentCalories:= spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	str := `Количество шагов: ` + steps + ".\n" +
		`Дистанция составила ` + fmt.Sprintf("%.2f", distance) + " км.\n" +
		`Вы сожгли: ` + fmt.Sprintf("%.2f", spentCalories)	+ "\n"	
	return str
}
