package management

import "testing"

func TestStat(t *testing.T) {
	t.Run("ActiveUsers", func(t *testing.T) {
		i, err := m.Stat.ActiveUsers()
		if err != nil {
			t.Error(err)
		}
		t.Logf("Active users: %d\n", i)
	})

	t.Run("Daily", func(t *testing.T) {
		s, err := m.Stat.Daily()
		if err != nil {
			t.Error(err)
		}
		for i, sd := range s {
			if i > 2 {
				break
			}
			t.Logf("%v\n", sd)
		}
	})
}
