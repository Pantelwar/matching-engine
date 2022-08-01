package util

import (
	d "github.com/shopspring/decimal"
)

// StandardBigDecimal initializes the standard big decimcal
type StandardBigDecimal struct {
	d.Decimal
}

// NewLogger initializes the standard logger
func NewDecimalFromString(str string) (*StandardBigDecimal, error) {
	var err error
	s := &StandardBigDecimal{d.Decimal{}}
	s.Decimal, err = d.NewFromString(str)
	return s, err
}

// NewLogger initializes the standard logger
func NewDecimalFromFloat(str float64) *StandardBigDecimal {
	s := &StandardBigDecimal{d.Decimal{}}
	s.Decimal = d.NewFromFloat(str)
	return s
}

func (s *StandardBigDecimal) Add(other *StandardBigDecimal) *StandardBigDecimal {
	bigDecimal := s.Decimal.Add(other.Decimal)
	return &StandardBigDecimal{bigDecimal}
}

func (s *StandardBigDecimal) Sub(other *StandardBigDecimal) *StandardBigDecimal {
	bigDecimal := s.Decimal.Sub(other.Decimal)
	return &StandardBigDecimal{bigDecimal}
}

func (s *StandardBigDecimal) Mul(other *StandardBigDecimal) *StandardBigDecimal {
	bigDecimal := s.Decimal.Mul(other.Decimal)
	return &StandardBigDecimal{bigDecimal}
}

func (s *StandardBigDecimal) Div(other *StandardBigDecimal) *StandardBigDecimal {
	bigDecimal := s.Decimal.Div(other.Decimal)
	return &StandardBigDecimal{bigDecimal}
}

func (s *StandardBigDecimal) Cmp(other *StandardBigDecimal) int {
	return s.Decimal.Cmp(other.Decimal)
}

func (s *StandardBigDecimal) Neg() *StandardBigDecimal {
	bigDecimal := s.Decimal.Neg()
	return &StandardBigDecimal{bigDecimal}
}

func (s *StandardBigDecimal) String() string {
	return s.Decimal.String()
}

func (s *StandardBigDecimal) Float64() float64 {
	floatValue, _ := s.Decimal.Float64()
	return floatValue
}
