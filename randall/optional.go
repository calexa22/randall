package randall

import (
	"time"

	"github.com/shopspring/decimal"
)

func OptionalUInt(v uint) *uint {
	return &v
}

func OptionalInt(v int) *int {
	return &v
}

func OptionalDecimal(v decimal.Decimal) *decimal.Decimal {
	return &v
}

func OptionalBool(v bool) *bool {
	return &v
}

func OptionalString(v string) *string {
	return &v
}

func OptionalTime(v time.Time) *time.Time {
	return &v
}

func Optional[T any](v T) *T {
	return &v
}
