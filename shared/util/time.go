package util

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
    "github.com/rs/zerolog/log"
)

// TimeToTimestamp returns a protobuf Timestamp from a Time object
func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	if nil != err {
		log.Error().Err(err).Msg("Time to Timestamp error")
	}
	return ts
}

// TimestampToTime returns a Time object from a protobuf Timestamp
func TimestampToTime(ts *timestamp.Timestamp) time.Time {
	if nil == ts {
		return time.Time{}
	}
	t, err := ptypes.Timestamp(ts)
	if nil != err {
        log.Error().Err(err).Msg("Timestamp to Times error")
	}
	return t
}
