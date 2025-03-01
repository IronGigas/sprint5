package actioninfo

import (
	"fmt"
)

type DataParser interface {
    Parse(data string) error
	ActionInfo() string
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			fmt.Println("Parsing error:", err)
			continue 
		}
		fmt.Println(dp.ActionInfo())
	}
}
