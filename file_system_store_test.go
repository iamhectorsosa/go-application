package poker

import (
	"log"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
	defer cleanDatabase()
	store, err := NewFileSystemPlayerStore(database)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	t.Run("league from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(t, got, want)

		// assert again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		got := store.GetPlayerScore("Chris")
		want := 33
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Hector")
		got := store.GetPlayerScore("Hector")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
				{"Name": "Cleo", "Wins": 10},
				{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)

	})
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
