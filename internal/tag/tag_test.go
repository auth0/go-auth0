package tag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseTag struct {
	Foo bool  `scope:"foo"`
	Bar *bool `scope:"bar"`
	Baz *bool `scope:"baz"`
}

func TestScopes(t *testing.T) {
	testCase := &testCaseTag{
		Foo: true,
		Bar: func(b bool) *bool { return &b }(true),
		Baz: func(b bool) *bool { return &b }(false),
	}

	assert.Equal(t, []string{"foo", "bar"}, Scopes(testCase))
}

func TestSetScopes(t *testing.T) {
	testCase := &testCaseTag{}
	SetScopes(testCase, true, "foo", "bar")

	assert.True(t, testCase.Foo)
	assert.Equal(t, func(b bool) *bool { return &b }(true), testCase.Bar)
}
