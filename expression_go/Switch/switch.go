package Switch

import "fmt"

func Switch() {
	grade := 'B'
	marks := 90

	switch marks {
	case 90:
		grade = 'A'
	case 80:
		grade = 'B'
	case 70, 60, 50:
		grade = 'C'
	default:
		grade = 'D'
	}

	switch {
	case grade == 'A':
		fmt.Println("优秀")
	case grade == 'B':
		fmt.Println("良好")
	case grade == 'C':
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	fmt.Println("你的等级是：", grade)
}

func TypeSwitch() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型: %T\n", i)
	case int:
		fmt.Println("x是int型")
	case float64:
		fmt.Println("x是float64型")
	case func(int) float64:
		fmt.Println("x是func(int)型")
	case bool, string:
		fmt.Println("x是bool或string型")
	default:
		fmt.Println("未知型")
	}
}

func Fallthrough() {
	switch {
	case false:
		fmt.Println("1.case条件语句为false")
		fallthrough
	case true:
		fmt.Println("2.case条件语句为true")
		fallthrough
	case false:
		fmt.Println("3.case条件语句为false")
		//fallthrough
	case true:
		fmt.Println("4.case条件语句为true")
		//fallthrough
	case false:
		fmt.Println("5.case条件语句为false")
		//fallthrough
	default:
		fmt.Println("6.默认case")
	}
}
