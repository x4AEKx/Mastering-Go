package main

import (
	"bytes"
	"testing"
)

func TestSearch(t *testing.T) {
	phoneBook := []User{{Name: "Dmitry", SurName: "Vlasov", Tel: +79772725854}}

	out := &bytes.Buffer{}
	Search(phoneBook, "Vlasov", out)

	got := out.String()
	expected := "{Dmitry Vlasov 79772725854}"

	t.Logf("got: %q", got)
	t.Logf("expected: %q", expected)

	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestList(t *testing.T) {
	phoneBook := []User{{Name: "Mihalis", SurName: "Tsoukalos", Tel: 2109416471}, {Name: "Mary", SurName: "Doe", Tel: 2109416471}}

	out := &bytes.Buffer{}
	List(phoneBook, out)

	got := out.String()
	expected := "{Mihalis Tsoukalos 2109416471}\n{Mary Doe 2109416471}\n"

	t.Logf("got: %q", got)
	t.Logf("expected: %q", expected)

	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}
