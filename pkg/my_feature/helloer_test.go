package my_feature_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"yngvark.com/go-cmd-template/pkg/my_feature"
)

func TestSayHello(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		expect string
	}{
		{
			name:   "Should work",
			expect: "Hello",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expect, my_feature.Hello())
		})
	}
}
