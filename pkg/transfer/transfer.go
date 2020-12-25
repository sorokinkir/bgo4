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
func NewService(c *card.Service, commission float64, rubMin int64) *Service {
	return &Service{CardSvc: c, Commission: commission, RubMin: rubMin}
}

// Card2Card method
func (s *Service) Card2Card(from, to string, amount int64) (total int, ok bool) {
	// TODO Всегда проверяем баланс для совершения операции
	// Между картами тиньков, коммиссии нет
	if from == "0020" {
		if to == "0020" {
			fmt.Println("Карты для перевода выпущены одним банком")
			return
		}
	}
	// На карту тиньков, коммиссии нет
	if from != "0020" {
		if to == "0020" {
			fmt.Println("Перевод с неизвестного банка на карту тиньков")
			return
		}
	}

	// Между картами других банков 1.5% или минимум 10 руб
	if from != "0020" {
		if to != "0020" {
			fmt.Println("Переводы между картами других банков, имеем с них копейки наши")
			return
		}
	}

	return
}
