package nexmo_test

import (
	"testing"

	"github.com/ammario/nexmo"
	"github.com/stretchr/testify/require"
)

func TestNumberInsight(t *testing.T) {
	t.Parallel()
	c := client()
	t.Run("Basic", func(t *testing.T) {
		resp, err := c.NumberInsightBasic("2562893333")
		require.Nil(t, err)
		require.Equal(t, "UG", resp.CountryCode)
		t.Logf("%+v", resp)

		resp, err = c.NumberInsightBasic("2562893333", nexmo.Country("US"))
		require.Nil(t, err)
		require.Equal(t, "US", resp.CountryCode)
		t.Logf("%+v", resp)
	})

}
