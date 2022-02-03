package management

import (
	"encoding/json"
	"testing"

	"github.com/auth0/go-auth0"
	"github.com/auth0/go-auth0/internal/testing/expect"
)

func TestPrompt(t *testing.T) {
	t.Cleanup(func() {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "classic",
			IdentifierFirst:          auth0.Bool(false),
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Update to the new identifier first experience", func(t *testing.T) {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "new",
			IdentifierFirst:          auth0.Bool(true),
		})
		if err != nil {
			t.Error(err)
		}

		ps, err := m.Prompt.Read()
		if err != nil {
			t.Error(err)
		}
		expect.Expect(t, ps.UniversalLoginExperience, "new")
		expect.Expect(t, ps.IdentifierFirst, auth0.Bool(true))
	})

	t.Run("Update to the classic non identifier first experience", func(t *testing.T) {
		err := m.Prompt.Update(&Prompt{
			UniversalLoginExperience: "classic",
			IdentifierFirst:          auth0.Bool(false),
		})
		if err != nil {
			t.Error(err)
		}

		ps, err := m.Prompt.Read()
		if err != nil {
			t.Error(err)
		}
		expect.Expect(t, ps.UniversalLoginExperience, "classic")
		expect.Expect(t, ps.IdentifierFirst, auth0.Bool(false))
	})
}

func TestPromptCustomText(t *testing.T) {
	t.Cleanup(func() {
		prompt := "login"
		lang := "en"
		body := make(map[string]interface{})
		err := m.Prompt.SetCustomText(prompt, lang, body)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Retrieve custom text", func(t *testing.T) {
		prompt := "login"
		lang := "en"

		texts, err := m.Prompt.CustomText(prompt, lang)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v\n", texts)
	})

	t.Run("Set custom text", func(t *testing.T) {
		prompt := "login"
		lang := "en"

		var body map[string]interface{}
		err := json.Unmarshal([]byte(`{ "login": { "title": "Welcome" } }`), &body)
		if err != nil {
			t.Error(err)
		}

		err = m.Prompt.SetCustomText(prompt, lang, body)
		if err != nil {
			t.Error(err)
		}
		expect.Expect(t, body["login"].(map[string]interface{})["title"], "Welcome")
	})
}
