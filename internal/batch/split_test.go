package batch

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		line string
		want []string
	}{
		{"", nil},
		{"   ", nil},
		{"ga user@example.com", []string{"ga", "user@example.com"}},
		{"  ma  user@example.com\tzimbraMailQuota 0 ", []string{"ma", "user@example.com", "zimbraMailQuota", "0"}},
		{`ma u@d displayName "John Doe"`, []string{"ma", "u@d", "displayName", "John Doe"}},
		{`ma u@d displayName 'John "J" Doe'`, []string{"ma", "u@d", "displayName", `John "J" Doe`}},
		{`ma u@d zimbraNotes ""`, []string{"ma", "u@d", "zimbraNotes", ""}},
		{`ma u@d zimbraNotes ''`, []string{"ma", "u@d", "zimbraNotes", ""}},
		{`ma u@d zimbraNotes "say \"hi\" \\ bye"`, []string{"ma", "u@d", "zimbraNotes", `say "hi" \ bye`}},
		{`ma u@d zimbraNotes "a\nb"`, []string{"ma", "u@d", "zimbraNotes", `a\nb`}},
		{`ma u@d zimbraNotes John\ Doe`, []string{"ma", "u@d", "zimbraNotes", "John Doe"}},
		{`ma u@d zimbraNotes pre"quoted part"post`, []string{"ma", "u@d", "zimbraNotes", "prequoted partpost"}},
		{"ma u@d displayName Élodie", []string{"ma", "u@d", "displayName", "Élodie"}},
	}

	for _, tt := range tests {
		got, err := Split(tt.line)

		if err != nil {
			t.Errorf("Split(%q) returned error %v", tt.line, err)
			continue
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Split(%q) = %q, want %q", tt.line, got, tt.want)
		}
	}
}

func TestSplitErrors(t *testing.T) {
	for _, line := range []string{
		`ma u@d displayName "John`,
		`ma u@d displayName 'John`,
		`ma u@d displayName John\`,
	} {
		if _, err := Split(line); err == nil {
			t.Errorf("Split(%q) returned no error", line)
		}
	}
}
