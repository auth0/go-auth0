package management

import (
	"errors"
	"testing"
	"time"

	"github.com/auth0/go-auth0"
)

func ensureActionBuilt(a *Action) (err error) {
	i := 1
	var r *Action
	for i < 20 {
		r, err = m.Action.Read(a.GetID())
		if err != nil {
			return
		}
		if r.GetStatus() == ActionStatusBuilt {
			break
		}
		time.Sleep(1 * time.Second)
		i++
	}
	if r.GetStatus() != ActionStatusBuilt {
		err = errors.New("action failed to build")
	}
	return
}

func TestActions(t *testing.T) {
	r := &Action{
		Name: auth0.String("test-action"),
		Code: auth0.String("exports.onExecutePostLogin = async (event, api) =\u003e {}"),
		SupportedTriggers: []*ActionTrigger{
			{
				ID:      auth0.String(ActionTriggerPostLogin),
				Version: auth0.String("v2"),
			},
		},
		Dependencies: []*ActionDependency{
			{
				Name:        auth0.String("lodash"),
				Version:     auth0.String("4.0.0"),
				RegistryURL: auth0.String("https://www.npmjs.com/package/lodash"),
			},
		},
		Secrets: []*ActionSecret{
			{
				Name:  auth0.String("mySecretName"),
				Value: auth0.String("mySecretValue"),
			},
		},
	}

	var err error
	var v *ActionVersion
	var vl *ActionVersionList

	t.Run("Triggers", func(t *testing.T) {
		l, err := m.Action.Triggers()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", l)
	})

	t.Run("Create", func(t *testing.T) {
		err = m.Action.Create(r)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Read", func(t *testing.T) {
		r, err = m.Action.Read(r.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("List", func(t *testing.T) {
		r, err := m.Action.List()
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("Deploy", func(t *testing.T) {
		err = ensureActionBuilt(r)
		if err != nil {
			t.Fatal(err)
		}
		v, err = m.Action.Deploy(r.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", v)
	})

	t.Run("Update", func(t *testing.T) {
		id := r.GetID()

		r.ID = nil        // read-only
		r.UpdatedAt = nil // read-only
		r.CreatedAt = nil // read-only
		r.Status = nil    // read-only
		r.Code = auth0.String("exports.onExecutePostLogin = async (event, api) => { api.user.setUserMetadata('myParam', 'foo'); };")

		err = m.Action.Update(id, r)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", r)
	})

	t.Run("DeployAgain", func(t *testing.T) {
		err = ensureActionBuilt(r)
		if err != nil {
			t.Fatal(err)
		}
		v, err = m.Action.Deploy(r.GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", v)
	})

	t.Run("Version", func(t *testing.T) {
		v, err := m.Action.Version(r.GetID(), v.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", v)
	})

	t.Run("Versions", func(t *testing.T) {
		vl, err = m.Action.Versions(r.GetID())
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", vl)
	})

	t.Run("DeployVersion", func(t *testing.T) {
		v, err = m.Action.DeployVersion(r.GetID(), vl.Versions[0].GetID())
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", v)
	})

	t.Run("UpdateBindings", func(t *testing.T) {
		b := []*ActionBinding{
			{
				Ref: &ActionBindingReference{
					Type:  auth0.String(ActionBindingReferenceByName),
					Value: r.Name,
				},
				DisplayName: auth0.String("My test action Binding"),
			},
		}
		err = m.Action.UpdateBindings(ActionTriggerPostLogin, b)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", b)
	})

	t.Run("Bindings", func(t *testing.T) {
		bl, err := m.Action.Bindings(ActionTriggerPostLogin)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", bl)
	})

	t.Run("Test", func(t *testing.T) {
		p := &ActionTestPayload{
			"event": ActionTestPayload{
				"user": ActionTestPayload{
					"email":         "j+smith@example.com",
					"emailVerified": true,
					"id":            "auth0|5f7c8ec7c33c6c004bbafe82",
				},
			},
		}
		err = m.Action.Test(r.GetID(), p)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", p)
	})

	t.Run("Execution", func(t *testing.T) {
		_, err := m.Action.Execution("M9IqRp9wQLaYNrSwz6YPTTIwMjEwNDA0")
		if err != nil {
			mgmtError, _ := err.(*managementError)
			if mgmtError.StatusCode != 404 {
				t.Fatal(err)
			}
			// Expect a 404 as we can't get execution ID via API
			t.Log(err)
			return
		}
		t.Fatal(errors.New("read execution unexpectedly succeeded"))
	})

	t.Run("ClearBindings", func(t *testing.T) {
		b := make([]*ActionBinding, 0)
		err = m.Action.UpdateBindings(ActionTriggerPostLogin, b)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%v\n", b)
	})

	t.Run("Delete", func(t *testing.T) {
		err = m.Action.Delete(r.GetID())
		if err != nil {
			t.Fatal(err)
		}
	})
}
