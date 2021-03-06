package main

import (
	"fmt"

	"github.com/sorokinkir/bgo4/pkg/card"
	"github.com/sorokinkir/bgo4/pkg/transfer"
)

func main() {
	bank1 := card.NewService("Tinkoff bank")

	client1 := bank1.IssueCard(1, "Master Card", "RUB", "1020", 50_000)
	client2 := bank1.IssueCard(2, "VISA", "RUB", "1030", 500)

	ownBank := transfer.NewService(bank1, 0, 10)
	fmt.Println(ownBank.Card2Card(client1.Number, client2.Number, 50_000))

	fromOwnBank := transfer.NewService(bank1, 0.5, 10)
	fmt.Println(fromOwnBank.Card2Card(client2.Number, "0001", 50))

	otherBanks := transfer.NewService(bank1, 1.5, 30)
	fmt.Println(otherBanks.Card2Card("0002", "0003", 100))
}
