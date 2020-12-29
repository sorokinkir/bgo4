package transfer

import (
	"fmt"

	"github.com/sorokinkir/bgo4/pkg/card"
)

// Service transfer
type Service struct {
	CardSvc    *card.Service
	Commission float64
	RubMin     int64
}

// NewService transfer package
func NewService(cardsvc *card.Service, commission float64, rubMin int64) *Service {
	return &Service{CardSvc: cardsvc, Commission: commission, RubMin: rubMin}
}

// Card2Card method
func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	fromCard := s.CardSvc.SearchCard(from)
	toCard := s.CardSvc.SearchCard(to)

	// Если обе карты наши
	if fromCard != nil && toCard != nil {
		if fromCard.Balance < amount {
			fmt.Println("Недостаточно денег для перевода")
			return amount, false
		}

		fromCard.Balance -= amount
		toCard.Balance += amount
		return toCard.Balance, true
	}
	// From карта наша, перевод на чужую
	if fromCard != nil && toCard == nil {
		//fmt.Println("Перевод с нашей канты на другой банк")
		if amount < 10 && fromCard.Balance <= amount {
			fmt.Println("Сумма не должна быть меньше 10 руб. и баланс должен быть больше или равен сумме перевода")
			return fromCard.Balance, false
		}

		resultAmount := float64(amount) * (1 - 0.5/100)
		fromCard.Balance -= int64(resultAmount)
		return fromCard.Balance, true

	}

	// Перевод на нашу карту
	if fromCard == nil && toCard != nil {
		// Баланс не проверяем, т.к. перевод на карту нашего банка
		toCard.Balance += amount
		return toCard.Balance, true
	}

	// Перевод с карты на карту не нашего банка
	if fromCard == nil && toCard == nil {
		if amount < 30 {
			fmt.Println("Сумма должна быть больше 30 руб. для перевода.")
		} else {
			resultAmount := float64(amount) * (1 - 1.5/100)
			resultAmount += float64(amount)
			return int64(resultAmount), true
		}
	}

	return 0, false
}
