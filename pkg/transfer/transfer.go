package transfer

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/sorokinkir/bgo4/pkg/card"
)

// Service transfer
type Service struct {
	CardSvc    *card.Service
	Commission float64
	RubMin     int64
}

var (
	ErrInvalidCardNumber            = errors.New("Ошибка в номере карты")
	ErrInvalidCard                  = errors.New("введеный номер карты не нашего банка")
	ErrOwnToOwnCardTransfer         = errors.New("недостаточно денег для перевода или необходимо минимум 10 руб")
	ErrOwnToUnknownCardTransfer     = errors.New("сумма не должна быть меньше 10 руб. и баланс должен быть больше или равен сумме перевода")
	ErrUnknownToUnknownCardTransfer = errors.New("сумма должна быть больше или равен 30 руб. для перевода")
	ErrUnknown                      = errors.New("unknown error")
)

// NewService transfer package
func NewService(cardsvc *card.Service, commission float64, rubMin int64) *Service {
	return &Service{CardSvc: cardsvc, Commission: commission, RubMin: rubMin}
}

func isValid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	numberCard := strings.Split(number, "")
	numbersSlice := make([]int, len(number))

	for i, row := range numberCard {
		j, err := strconv.Atoi(row)
		if err != nil {
			return false
		}
		numbersSlice[i] = j
	}

	return checkCardByLuhn(numbersSlice)
}

func checkCardByLuhn(numbers []int) bool {
	var check int
	for i := 0; i < len(numbers); i += 2 {
		cardNumber := numbers[i] * 2

		if cardNumber > 9 {
			cardNumber -= 9
		}
		numbers[i] = cardNumber
	}

	for _, i := range numbers {
		check += i
	}

	if check%10 == 0 {
		return true
	}
	return false
}

// Card2Card method
func (s *Service) Card2Card(from, to string, amount int64) (total int64, err error) {
	if !isValid(from) || !isValid(to) {
		return amount, ErrInvalidCardNumber
	}

	fromCard, err := s.CardSvc.SearchCard(from)
	if err != nil {
		fmt.Println("Платеж на карту: <--", err)
	}
	toCard, err := s.CardSvc.SearchCard(to)
	if err != nil {
		fmt.Println("Отправляем платеж: -->", err)
	}

	// Если обе карты наши
	if fromCard != nil && toCard != nil {
		if fromCard.Balance < amount || amount < s.RubMin {
			return amount, ErrOwnToOwnCardTransfer
		}
		resultProcent := float64(amount) * (s.Commission / 100)
		finalSumAmount := float64(amount) + resultProcent
		fromCard.Balance -= amount
		toCard.Balance += amount
		return int64(finalSumAmount), nil
	}
	// From карта наша, перевод на чужую
	if fromCard != nil && toCard == nil {
		resultProcent := float64(amount) * (s.Commission / 100)
		finalSumAmount := float64(amount) + resultProcent

		if amount < s.RubMin || fromCard.Balance <= amount {
			return int64(finalSumAmount), ErrOwnToUnknownCardTransfer
		}

		fromCard.Balance -= int64(finalSumAmount)
		return int64(finalSumAmount), nil
	}

	// Перевод на нашу карту
	if fromCard == nil && toCard != nil {
		resultProcent := float64(amount) * (s.Commission / 100)
		finalSumAmount := float64(amount) + resultProcent
		// Зачисляем на карту итоговую сумму + комиссию
		toCard.Balance += int64(finalSumAmount)
		return int64(finalSumAmount), nil
	}

	// Перевод с карты на карту не нашего банка
	if fromCard == nil && toCard == nil {
		resultProcent := float64(amount) * (s.Commission / 100)
		finalSumAmount := float64(amount) + resultProcent
		if amount <= s.RubMin {
			return amount, ErrUnknownToUnknownCardTransfer
		}
		return int64(finalSumAmount), nil
	}

	return amount, ErrUnknown
}
