package normalizing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeCharacter(t *testing.T) {
	testcases := []struct {
		source            string
		expectedCharacter string
		expectedCostume   string
	}{
		{"キティ・ホワイト", "キティ・ホワイト", ""},
		{"キティ(ゴースト)", "キティ・ホワイト", "ゴースト"},
		{"メアリー・ホワイト(ママ)", "メアリー・ホワイト", ""},
		{"ミルク", "みるく", ""},
	}

	for _, testcase := range testcases {
		actualCharacter, actualCostume := NormalizeCharacter(testcase.source)
		assert.Equal(t, testcase.expectedCharacter, actualCharacter)
		assert.Equal(t, testcase.expectedCostume, actualCostume)
	}
}
