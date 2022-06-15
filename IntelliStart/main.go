package main

import (
	"fmt"
)

const (
	ApplePrice = 5.99
	PearPrice  = 7.
	HaveMoney  = 23.0
)

func main() {
	AppleCost := ApplePrice * 9                 //Стоимость 9-ти яблок
	PearCost := PearPrice * 8                   //Стоимость 8-ми груш
	TotalCost := AppleCost + PearCost           //Стоимость 9-ти яблок и 8-ми груш
	AmountPear := HaveMoney / PearPrice         //Сколько груш мы можем купить?
	AmountApple := HaveMoney / ApplePrice       //Сколько яблок мы можем купить?
	ApplePearCost := ApplePrice*2 + PearPrice*2 //Стоимость 2-х яблок и 2-х груш

	fmt.Println("1). Сколько денег надо потратить чтобы купить 9 яблок и 8 груш?")
	fmt.Println("Надо потратить", TotalCost, "Грн.")
	fmt.Println("2). Сколько груш мы можем купить?")
	fmt.Println("Мы можем купить", int(AmountPear), "груши")
	fmt.Println("3). Сколько яблок мы можем купить?")
	fmt.Println("Мы можем купить", int(AmountApple), "яблока")
	fmt.Println("4). Можем ли мы купить 2 груши и 2 яблока")
	if ApplePearCost < HaveMoney {
		fmt.Println("Можем")
	} else {
		fmt.Println("Не можем")
	}

}
