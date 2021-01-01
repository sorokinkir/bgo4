package transfer

import (
	"testing"

	"github.com/sorokinkir/bgo4/pkg/card"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc    *card.Service
		Commission float64
		RubMin     int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	cardSvc := card.NewService("Tinkoff bank")
	cardSvc.Add(
		&card.Card{
			ID:      1,
			Balance: 500,
			Number:  "0001",
		},
		&card.Card{
			ID:      2,
			Balance: 400,
			Number:  "0002",
		},
		&card.Card{
			ID:      3,
			Balance: 300,
			Number:  "0003",
		},
		&card.Card{
			ID:      4,
			Balance: 100,
			Number:  "0004",
		},
		&card.Card{
			ID:      5,
			Balance: 0,
			Number:  "0005",
		},
	)
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		{
			name:      "Карта своего банка -> Карта своего банка (денег достаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0, RubMin: 10},
			args:      args{from: "0001", to: "0002", amount: 100},
			wantTotal: 500,
			wantOk:    true,
		},
		{
			name:      "Карта своего банка -> Карта своего банка (денег недостаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0, RubMin: 10},
			args:      args{from: "0001", to: "0002", amount: 600},
			wantTotal: 600,
			wantOk:    false,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег достаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "0003", to: "0010", amount: 200},
			wantTotal: 201,
			wantOk:    true,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег недостаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "0004", to: "0011", amount: 200},
			wantTotal: 201,
			wantOk:    false,
		},
		{
			name:      "Карта чужого банка -> Карта своего банка",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "0012", to: "0005", amount: 200},
			wantTotal: 201,
			wantOk:    true,
		},
		{
			name:      "Карта чужого банка -> Карта чужого банка",
			fields:    fields{CardSvc: cardSvc, Commission: 1.5, RubMin: 30},
			args:      args{from: "0013", to: "0014", amount: 200},
			wantTotal: 203,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:    tt.fields.CardSvc,
				Commission: tt.fields.Commission,
				RubMin:     tt.fields.RubMin,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Service.Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Service.Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
