package poker

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, nil}
}

type InMemoryPlayerStore struct {
	store    map[string]int
	winCalls []string
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
	i.winCalls = append(i.winCalls, name)
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() League {
	var league League
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
