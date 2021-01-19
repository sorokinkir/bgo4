package card

import (
	"errors"
	"fmt"
	"strings"
)

// Service card
type Service struct {
	BankName string
	Cards    []*Card
}

// Card ....
type Card struct {
	ID       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

// ErrCardNotFound errors for package cards
var ErrCardNotFound = errors.New("карта не найдена среди наших карт")

// NewService card
func NewService(BankName string) *Service {
	return &Service{BankName: BankName}
}

// IssueCard выпуск карта для нашего банка
func (s *Service) IssueCard(id int64, issue, currency, number string, balance int64) *Card {
	card := &Card{
		ID:       id,
		Issuer:   issue,
		Balance:  balance,
		Currency: currency,
		Number:   number,
	}
	s.Cards = append(s.Cards, card)
	return card
}

// SearchCard Проверяем карту, является ли она нашего банка
func (s *Service) SearchCard(cardNum string) (*Card, error) {
	if strings.HasPrefix(cardNum, "5106 21") == true {
		for _, row := range s.Cards {
			if row.Number == cardNum {
				return row, nil
			}
		}
	}
	return nil, ErrCardNotFound
}

// Add добавлен метод для добавления карт
func (s *Service) Add(cards ...*Card) {
	s.Cards = append(s.Cards, cards...)
}

// CardsNumbers show cards in structure
func (s *Service) CardsNumbers() {
	for _, card := range s.Cards {
		if strings.HasPrefix(card.Number, "5106 21") {
			fmt.Println("Наша карта: ", card.Number)
		} else {
			fmt.Println("В слайсе: ", card.Number, " не нашего банка")
		}
	}
}
