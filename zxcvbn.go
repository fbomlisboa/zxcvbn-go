package zxcvbn

import (
	"github.com/fbomlisboa/zxcvbn-go/matching"
	"github.com/fbomlisboa/zxcvbn-go/scoring"
	"github.com/fbomlisboa/zxcvbn-go/utils/math"
	"time"
)

func PasswordStrength(password string, userInputs []string) scoring.MinEntropyMatch {
	start := time.Now()
	matches := matching.Omnimatch(password, userInputs)
	result := scoring.MinimumEntropyMatchSequence(password, matches)
	end := time.Now()

	calcTime := end.Nanosecond() - start.Nanosecond()
	result.CalcTime = zxcvbn_math.Round(float64(calcTime)*time.Nanosecond.Seconds(), .5, 3)
	return result
}
