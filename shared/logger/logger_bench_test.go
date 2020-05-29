package logger_test

import (
    "testing"

    "github.com/rs/zerolog/log"
)

func BenchmarkInfoLog(b *testing.B) {
    b.Run("zerolog", func(b *testing.B) {
        b.ResetTimer()
        b.RunParallel(func(pb *testing.PB) {
            for pb.Next() {
                log.Info().Msg("Benchmarking: InfoZ")
            }
        })
    })
}
