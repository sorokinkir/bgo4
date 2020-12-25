package card

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

// NewService card
func NewService(BankName string) *Service {
	return &Service{BankName: BankName}
}

// IssueCard ....
func (s *Service) IssueCard(issue, currency, number string, balance int64) *Card {
	card := &Card{
		Issuer:   issue,
		Balance:  balance,
		Currency: currency,
		Number:   number,
	}
	s.Cards = append(s.Cards, card)
	return card
}

// SearchCard by number. ex: 0020
func (s *Service) SearchCard(cardNum string) *Card {
	for _, row := range s.Cards {
		if row.Number == cardNum {
			return row
		}
	}
	// Ничего не возвращаем тогда
	return nil
}
