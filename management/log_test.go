package management

import (
	"testing"

	"github.com/auth0/go-auth0"
)

func TestLog(t *testing.T) {
	// Save a log from the listing so we can look it up by ID later
	var firstLog *Log

	t.Run("List", func(t *testing.T) {
		// Limit results to 5 entries, starting from the first page
		logs, err := m.Log.List(Page(1), PerPage(5))
		if err != nil {
			t.Error(err)
		}

		for i, log := range logs {
			t.Logf("%v\n", log)

			// Save the first log for reading later
			if i == 0 {
				firstLog = log
			}
		}
	})

	t.Run("Read", func(t *testing.T) {
		log, err := m.Log.Read(auth0.StringValue(firstLog.ID))
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", log)
	})

	t.Run("Search", func(t *testing.T) {
		// Search by type "Success Exchange" and limit results to 5 entries
		logs, err := m.Log.List(Parameter("q", `type:"seacft"`), PerPage(5))
		if err != nil {
			t.Error(err)
		}
		for _, log := range logs {
			t.Logf("%v\n", log)
			if auth0.StringValue(log.Type) != "seacft" {
				t.Error("unexpected log type found")
			}
		}
	})
}
