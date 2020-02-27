package util

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/xmlking/logger/log"
)

// TimeToTimestamp returns a protobuf Timestamp from a Time object
func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	ts, err := ptypes.TimestampProto(t)
	if nil != err {
		log.Errorw("Time to Timestamp error", err)
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
		log.Errorw("Timestamp to Times error", err)
	}
	return t
}
