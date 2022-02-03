package tag

import (
	"testing"

	"github.com/auth0/go-auth0/internal/testing/expect"
)

type test struct {
	Foo bool  `scope:"foo"`
	Bar *bool `scope:"bar"`
	Baz *bool `scope:"baz"`
	Bam *bool `scope:"bam"`
}

func TestScopes(t *testing.T) {
	c := &test{
		Foo: true,
		Bar: func(b bool) *bool { return &b }(true),
		Baz: func(b bool) *bool { return &b }(false),
	}
	expect.Expect(t, Scopes(c), []string{"foo", "bar"})
}

func TestSetScopes(t *testing.T) {
	c := &test{}
	SetScopes(c, true, "foo", "bar")
	expect.Expect(t, c.Foo, true)
	expect.Expect(t, c.Bar, func(b bool) *bool { return &b }(true))
}
