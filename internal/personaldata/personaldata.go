package personaldata

import ("fmt")

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	str := `Имя: ` + p.Name + "\n" +
    	`Вес:  ` + fmt.Sprintf("%.2f", p.Weight) + "\n" +
    	`Рост: ` + fmt.Sprintf("%.2f", p.Height)
	fmt.Println(str)
}

