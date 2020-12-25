package main

import (
	"github.com/sorokinkir/bgo4/pkg/card"
	"github.com/sorokinkir/bgo4/pkg/transfer"
)

func main() {
	bank1 := card.NewService("Tinkoff bank")
	bank2 := card.NewService("SberBank")
	client1 := bank1.IssueCard("Master Card", "RUB", "0020", 50_000)
	// card2 := bank1.IssueCard("МИР", "RUB", "0020", 10_000)
	client3 := bank2.IssueCard("VISA", "RUB", "0030", 15_000)

	t := transfer.NewService(bank1, 1.5, 10)
	t.Card2Card(client1.Number, client1.Number, 50)
	t.Card2Card(client1.Number, client3.Number, 50)
	t.Card2Card(client3.Number, client3.Number, 300)

	bank1.SearchCard("0020")
}
