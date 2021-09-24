package model

import "errors"

type Money struct {
	amount   float64
	currency string
}

func NewMoney(amount float64, currency string) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (m *Money) Add(arg Money) (*Money, error) {
	if m.currency != arg.currency {
		return nil, errors.New("通過単位が異なります。")
	}
	return NewMoney(m.amount+arg.amount, m.currency), nil
}

func (m *Money) Amount() float64 {
	return m.amount
}
