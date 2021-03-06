package canonicalizing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanonicalizeVenueName(t *testing.T) {
	testcases := []struct {
		source   string
		expected string
	}{
		{"メルメルショップフォトスポット（４Ｆ）", "メルメルショップフォトスポット(4F)"},
		{"ビレッジ（１Ｆ）", "ピューロビレッジ(1F)"},
	}

	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, CanonicalizeVenueName(testcase.source))
	}
}
