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
			Number:  "5106 2176 6556 4334",
		},
		&card.Card{
			ID:      2,
			Balance: 400,
			Number:  "5106 2145 2663 3929",
		},
		&card.Card{
			ID:      3,
			Balance: 300,
			Number:  "5106 2134 7446 7926",
		},
		&card.Card{
			ID:      4,
			Balance: 100,
			Number:  "5106 2162 8120 3467",
		},
		&card.Card{
			ID:      5,
			Balance: 0,
			Number:  "5106 2162 9813 3749",
		},
	)
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantErr   error
	}{
		{
			name:      "Карта своего банка -> Карта своего банка (денег достаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "5106 2176 6556 4334", to: "5106 2145 2663 3929", amount: 100},
			wantTotal: 100,
			wantErr:   nil,
		},
		{
			name:      "Карта своего банка -> Карта своего банка (денег недостаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "5106 2176 6556 4334", to: "5106 2145 2663 3929", amount: 600},
			wantTotal: 600,
			wantErr:   ErrOwnToOwnCardTransfer,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег достаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "5106 2134 7446 7926", to: "5520 7164 2794 9988", amount: 200},
			wantTotal: 201,
			wantErr:   nil,
		},
		{
			name:      "Карта своего банка -> Карта чужого банка (денег недостаточно на той карте, с которой отправляется)",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "5106 2162 8120 3467", to: "5324 2066 2982 5382", amount: 200},
			wantTotal: 201,
			wantErr:   ErrOwnToUnknownCardTransfer,
		},
		{
			name:      "Карта чужого банка -> Карта своего банка",
			fields:    fields{CardSvc: cardSvc, Commission: 0.5, RubMin: 10},
			args:      args{from: "5323 5571 4095 4965", to: "5106 2162 9813 3749", amount: 200},
			wantTotal: 201,
			wantErr:   nil,
		},
		{
			name:      "Карта чужого банка -> Карта чужого банка",
			fields:    fields{CardSvc: cardSvc, Commission: 1.5, RubMin: 30},
			args:      args{from: "5288 7784 8489 6553", to: "5177 4374 2025 1233", amount: 200},
			wantTotal: 203,
			wantErr:   nil,
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
			if gotOk != tt.wantErr {
				t.Errorf("Service.Card2Card() gotOk = %v, want %v", gotOk, tt.wantErr)
			}
		})
	}
}
