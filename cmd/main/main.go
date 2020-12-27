package main

import (
	"github.com/sorokinkir/bgo4/pkg/card"
	"github.com/sorokinkir/bgo4/pkg/transfer"
)

func main() {
	bank1 := card.NewService("Tinkoff bank")

	client1 := bank1.IssueCard(1, "Master Card", "RUB", "1020", 50_000)
	client2 := bank1.IssueCard(2, "VISA", "RUB", "1030", 500)

	t := transfer.NewService(bank1, 1.5, 10)
	t.Card2Card(client1.Number, client2.Number, 50_000)
	t.Card2Card(client2.Number, "0001", 100)
}
