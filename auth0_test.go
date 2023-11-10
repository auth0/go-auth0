package auth0

import (
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	for _, test := range []struct {
		in       *bool
		expected bool
	}{
		{nil, false},
		{Bool(false), false},
		{Bool(true), true},
	} {
		have := BoolValue(test.in)
		if have != test.expected {
			t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
		}
	}
}

func TestInt(t *testing.T) {
	for _, test := range []struct {
		in       *int
		expected int
	}{
		{nil, 0},
		{Int(0), 0},
		{Int(1), 1},
		{Int(-1), -1},
	} {
		have := IntValue(test.in)
		if have != test.expected {
			t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
		}
	}
}

func TestFloat64(t *testing.T) {
	for _, test := range []struct {
		in       *float64
		expected float64
	}{
		{nil, 0},
		{Float64(0), 0},
		{Float64(1), 1},
		{Float64(-1), -1},
	} {
		have := Float64Value(test.in)
		if have != test.expected {
			t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
		}
	}
}

func TestString(t *testing.T) {
	for _, test := range []struct {
		in       *string
		expected string
	}{
		{nil, ""},
		{String(""), ""},
		{String("foo"), "foo"},
		{String("bar"), "bar"},
	} {
		have := StringValue(test.in)
		if have != test.expected {
			t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
		}
	}
}

func TestStringf(t *testing.T) {
	for _, test := range []struct {
		in       *string
		expected string
	}{
		{nil, ""},
		{Stringf(""), ""},
		{Stringf("%s", "foo"), "foo"},
		{Stringf("%s %d", "bar", 1), "bar 1"},
	} {
		have := StringValue(test.in)
		if have != test.expected {
			t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
		}
	}
}

func TestTime(t *testing.T) {
	for _, test := range []struct {
		in       *time.Time
		expected time.Time
	}{
		{nil, time.Time{}},
		{Time(time.Time{}), time.Time{}},
	} {
		have := TimeValue(test.in)
		if have != test.expected {
			t.Errorf("unexpected output. have %v, expected %v", have, test.expected)
		}
	}
}
