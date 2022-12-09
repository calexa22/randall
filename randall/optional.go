package randall

import (
	"time"
)

func OptionalUInt(v uint) *uint {
	return &v
}

func OptionalFloat32(v float32) *float32 {
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
