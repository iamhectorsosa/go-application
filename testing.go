package poker

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %d, want %d", got, want)
	}
}

func AssertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, want string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := store.winCalls[0]

	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}
