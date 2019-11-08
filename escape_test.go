package schema

import "testing"

// Tests for internal methods.

func TestEscapeWithDoubleQuotes(t *testing.T) {

	tests := []struct {
		input  string
		expect string
	}{
		{input: `foo`, expect: `"foo"`},
		{input: `fo"o`, expect: `"fo""o"`},
		{input: `foo"`, expect: `"foo"""`},
	}

	for _, x := range tests {
		res := escapeWithDoubleQuotes(x.input)
		if res != x.expect {
			t.Errorf("Failed, got: %s, want: %s.", res, x.expect)
		}
	}
}
func TestEscapeWithBackticks(t *testing.T) {

	tests := []struct {
		input  string
		expect string
	}{
		{input: "foo", expect: "`foo`"},
		{input: "fo`o", expect: "`fo``o`"},
		{input: "foo`", expect: "`foo```"},
	}

	for _, x := range tests {
		res := escapeWithBackticks(x.input)
		if res != x.expect {
			t.Errorf("Failed, got: %s, want: %s.", res, x.expect)
		}
	}
}

func TestEscapeWithBrackets(t *testing.T) {

	tests := []struct {
		input  string
		expect string
	}{
		{input: "foo", expect: "[foo]"},
		{input: "fo]o", expect: "[fo]]o]"},
		{input: "foo]", expect: "[foo]]]"},
	}

	for _, x := range tests {
		res := escapeWithBrackets(x.input)
		if res != x.expect {
			t.Errorf("Failed, got: %s, want: %s.", res, x.expect)
		}
	}
}

func TestEscapeWithBraces(t *testing.T) {

	tests := []struct {
		input  string
		expect string
	}{
		{input: "foo", expect: "{foo}"},
		{input: "fo}o", expect: "{fo}}o}"},
		{input: "foo}", expect: "{foo}}}"},
	}

	for _, x := range tests {
		res := escapeWithBraces(x.input)
		if res != x.expect {
			t.Errorf("Failed, got: %s, want: %s.", res, x.expect)
		}
	}
}
