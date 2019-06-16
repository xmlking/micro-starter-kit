package util

import (
	"time"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
)

func TestTimeToTimestamp(t *testing.T) {
	assert := assert.New(t)
	tm := time.Now()
	ts := TimeToTimestamp(tm)
	assert.Equal(tm.Unix(), ts.Seconds)
}

func TestTimestampToTime(t *testing.T) {
	assert := assert.New(t)
	ts, err := ptypes.TimestampProto(time.Now())
	assert.NoError(err)

	tm := TimestampToTime(ts)
	assert.Equal(tm.Unix(), ts.Seconds)

	tm = TimestampToTime(nil)
	assert.True(tm.IsZero())
}