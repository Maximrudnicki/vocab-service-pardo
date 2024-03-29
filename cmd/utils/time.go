package utils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTimestamp(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
