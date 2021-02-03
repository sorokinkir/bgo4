package main

import (
	"fmt"

	"github.com/sorokinkir/bgo4/pkg/card"
	"github.com/sorokinkir/bgo4/pkg/transfer"
)

func main() {
	bank1 := card.NewService("Tinkoff bank")

	client1 := bank1.IssueCard(1, "Master Card", "RUB", "5106 2176 6556 4334", 50_000)
	client2 := bank1.IssueCard(2, "VISA", "RUB", "5106 2145 2663 3929", 500)

	fmt.Println("--- Первый платеж ---")
	ownBank := transfer.NewService(bank1, 0, 10)
	res, err := ownBank.Card2Card(client1.Number, client2.Number, 50_000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println("--- Второй платеж ---")
	fromOwnBank := transfer.NewService(bank1, 0.5, 10)
	res, err = fromOwnBank.Card2Card(client2.Number, "5522 0087 3742 2494", 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println("--- Третий платеж ---")
	otherBanks := transfer.NewService(bank1, 1.5, 30)
	res, err = otherBanks.Card2Card("5137 4624 0402 5527", "5369 9028 8434 8708", 100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	// Служит для проверки номера карты
	transfer.CheckCardNumber("79927398713")
}
