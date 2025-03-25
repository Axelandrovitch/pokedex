package pokeapi

import (
	"reflect"
	"testing"

	"github.com/Axelandrovitch/pokedex/internal/pokecache"
)

func TestFetchLocations(t *testing.T) {
	const url = "https://pokeapi.co/api/v2/location-area/"
	expectedResults := []string{
		"canalave-city-area",
		"eterna-city-area",
		"pastoria-city-area",
		"sunyshore-city-area",
		"sinnoh-pokemon-league-area",
		"oreburgh-mine-1f",
		"oreburgh-mine-b1f",
		"valley-windworks-area",
		"eterna-forest-area",
		"fuego-ironworks-area",
		"mt-coronet-1f-route-207",
		"mt-coronet-2f",
		"mt-coronet-3f",
		"mt-coronet-exterior-snowfall",
		"mt-coronet-exterior-blizzard",
		"mt-coronet-4f",
		"mt-coronet-4f-small-room",
		"mt-coronet-5f",
		"mt-coronet-6f",
		"mt-coronet-1f-from-exterior",
	}

	cache := pokecache.NewCache(5)
	data, err := FetchLocations(cache, url)
	if err != nil {
		t.Fatalf("expected to fetch locations, got error: %v", err)
	}

	var fetchedResults []string
	for _, location := range data.Results {
		fetchedResults = append(fetchedResults, location.Name)
	}

	if !reflect.DeepEqual(fetchedResults, expectedResults) {
		t.Errorf("fetched results do not match expected results.\nGot: %v\nWant: %v", fetchedResults, expectedResults)
	}
}
