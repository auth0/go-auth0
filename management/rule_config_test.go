package management

import (
	"testing"

	"github.com/auth0/go-auth0"
)

func TestRuleConfig(t *testing.T) {
	key := "foo"
	r := &RuleConfig{Value: auth0.String("bar")}

	var err error

	t.Run("Upsert", func(t *testing.T) {
		err = m.RuleConfig.Upsert(key, r)
		if err != nil {
			t.Error(err)
		}
		rkey := auth0.StringValue(r.Key)
		if rkey != key {
			t.Errorf("key mismatch %q != %q", rkey, key)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.RuleConfig.Read(key)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Upsert", func(t *testing.T) {
		r.Key = nil // read-only
		r.Value = auth0.String("baz")
		err = m.RuleConfig.Upsert(key, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.RuleConfig.Delete(key)
		if err != nil {
			t.Error(err)
		}
	})
}
