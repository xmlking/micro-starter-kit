package time

import (
    "testing"
    "time"

    "github.com/golang/protobuf/ptypes"
    "github.com/stretchr/testify/assert"
)

func TestTimeToTimestamp(t *testing.T) {
    assert := assert.New(t)
    tm := time.Now()
    ts := ToTimestamp(tm)
    assert.Equal(tm.Unix(), ts.Seconds)
}

func TestTimestampToTime(t *testing.T) {
    assert := assert.New(t)
    ts, err := ptypes.TimestampProto(time.Now())
    assert.NoError(err)

    tm := ToTime(ts)
    assert.Equal(tm.Unix(), ts.Seconds)

    tm = ToTime(nil)
    assert.True(tm.IsZero())
}
