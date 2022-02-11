package management

import (
	"fmt"
	"testing"
)

func TestAnomaly(t *testing.T) {
	t.Run("CheckIP", func(t *testing.T) {
		isBlocked, err := m.Anomaly.CheckIP("1.1.1.1")
		if err != nil {
			t.Error(err)
		}
		if isBlocked {
			t.Error(fmt.Errorf("IP should not be blocked"))
		}
	})

	t.Run("UnblockIP", func(t *testing.T) {
		err := m.Anomaly.UnblockIP("1.1.1.1")
		if err != nil {
			t.Error(err)
		}
	})
}
