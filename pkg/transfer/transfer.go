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
	ErrInvalidCardNumber            = errors.New("ошибка в номере карты")
	ErrInvalidCard                  = errors.New("введеный номер карты не нашего банка")
	ErrOwnToOwnCardTransfer         = errors.New("недостаточно денег для перевода или необходимо минимум 10 руб")
	ErrOwnToUnknownCardTransfer     = errors.New("сумма не должна быть меньше 10 руб. и баланс должен быть больше или равен сумме перевода")
	ErrUnknownToUnknownCardTransfer = errors.New("сумма должна быть больше или равен 30 руб. для перевода")
	ErrMoneyTrasnfer                = errors.New("недостаточно денег для перевода")
	ErrUnknown                      = errors.New("unknown error")
)

// NewService transfer package
func NewService(cardsvc *card.Service, commission float64, rubMin int64) *Service {
	return &Service{CardSvc: cardsvc, Commission: commission, RubMin: rubMin}
}

func isValid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	numberCard := strings.Split(number, "")
	numbersSlice := make([]int, len(numberCard))

	for i, row := range numberCard {
		r, err := strconv.Atoi(row)
		if err != nil {
			return false
		}
		numbersSlice[i] = r
	}

	var num int
	if len(numbersSlice)%2 != 0 {
		num = 1
	}

	for i := num; i < len(numbersSlice); i += 2 {
		numbersSlice[i] *= 2
		if numbersSlice[i] > 9 {
			numbersSlice[i] -= 9
		}
	}

	var sumValues int
	// Считаем сумму чисел
	for _, i := range numbersSlice {
		sumValues += i
	}

	if sumValues%10 == 0 {
		return true
	}
	return false
}

// CheckCardNumber print true of false
func CheckCardNumber(num string) {
	fmt.Println(isValid(num))
}

// Card2Card method
func (s *Service) Card2Card(from, to string, amount int64) (total int64, err error) {
	if !isValid(from) || !isValid(to) {
		return amount, ErrInvalidCardNumber
	}

	resultProcent := float64(amount) * (s.Commission / 100)
	finalSumAmount := float64(amount) + resultProcent

	fromCard, fromCardErr := s.CardSvc.SearchCard(from)
	toCard, toCardErr := s.CardSvc.SearchCard(to)

	// Перевод с карты на карту не нашего банка
	if fromCardErr != nil && toCardErr != nil {
		return int64(finalSumAmount), nil
	}

	// Если обе карты наши
	if fromCard != nil && toCard != nil {
		if fromCard.Balance < amount || amount < s.RubMin {
			return amount, ErrOwnToOwnCardTransfer
		}
		fromCard.Balance -= amount
		toCard.Balance += amount
		return int64(finalSumAmount), nil
	}

	// From карта наша, перевод на чужую
	if toCardErr != nil {
		if amount < s.RubMin || fromCard.Balance <= amount {
			return int64(finalSumAmount), ErrOwnToUnknownCardTransfer
		}
		fromCard.Balance -= int64(finalSumAmount)
		return int64(finalSumAmount), nil
	}

	// Перевод на нашу карту
	if fromCardErr != nil {
		toCard.Balance += int64(finalSumAmount)
		return int64(finalSumAmount), nil
	}

	return amount, ErrUnknown
}
