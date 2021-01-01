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
		if fromCard.Balance < amount || amount < s.RubMin {
			// fmt.Println("Недостаточно денег для перевода или необходимо минимум 10 руб.")
			return amount, false
		}

		fromCard.Balance -= amount
		toCard.Balance += amount
		return toCard.Balance, true
	}
	// From карта наша, перевод на чужую
	if fromCard != nil && toCard == nil {
		resultProcent := float64(amount) * (s.Commission / 100)
		finalSumAmount := float64(amount) + resultProcent

		if amount < s.RubMin || fromCard.Balance <= amount {
			// fmt.Println("Сумма не должна быть меньше 10 руб. и баланс должен быть больше или равен сумме перевода")
			return int64(finalSumAmount), false
		}

		fromCard.Balance -= int64(finalSumAmount)
		return int64(finalSumAmount), true

	}

	// Перевод на нашу карту
	if fromCard == nil && toCard != nil {
		resultProcent := float64(amount) * (s.Commission / 100)
		finalSumAmount := float64(amount) + resultProcent
		// Зачисляем на карту итоговую сумму + комиссию
		toCard.Balance += int64(finalSumAmount)
		return int64(finalSumAmount), true
	}

	// Перевод с карты на карту не нашего банка
	if fromCard == nil && toCard == nil {
		resultProcent := float64(amount) * (s.Commission / 100)
		finalSumAmount := float64(amount) + resultProcent
		if amount <= s.RubMin {
			fmt.Println("Сумма должна быть больше или равен 30 руб. для перевода.")
			return amount, false
		}
		return int64(finalSumAmount), true
	}

	return amount, false
}
