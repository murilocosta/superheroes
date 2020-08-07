package superhero

import (
	"testing"
)

const (
	apiURL = "https://superheroapi.com/api"
	token  = "3185322551504722"
)

func TestFindByName(t *testing.T) {
	api := NewSuperHeroApi(apiURL, token)
	rsp, err := api.FindByName("superman")
	if err != nil {
		t.Errorf("Could not make request to API:\n%s", err)
	}

	// Need to have something on the response
	if rsp == nil || len(rsp) == 0 {
		t.Error("Failed to fetch values from API")
	}
}
