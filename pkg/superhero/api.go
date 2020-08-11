package superhero

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RemoteSuper struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	Powerstats struct {
		Intelligence string `json:"intelligence"`
		Power        string `json:"power"`
	} `json:"powerstats"`

	Biography struct {
		FullName  string `json:"full-name"`
		Alignment string `json:"alignment"`
	} `json:"biography"`

	Work struct {
		Occupation string `json:"occupation"`
	} `json:"work"`

	Image struct {
		URL string `json:"url"`
	} `json:"image"`
}

type ResultQuerySuper struct {
	Response   string         `json:"response"`
	ResultsFor string         `json:"results-for"`
	Results    []*RemoteSuper `json:"results"`
}

type SuperHeroApi struct {
	apiURL string
	token  string
}

func NewSuperHeroApi(apiURL string, token string) *SuperHeroApi {
	return &SuperHeroApi{apiURL, token}
}

func (api *SuperHeroApi) buildRequestPath(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", api.apiURL, api.token, endpoint)
}

func (api *SuperHeroApi) FindByName(name string) ([]*RemoteSuper, error) {
	tgt := api.buildRequestPath("search/" + name)
	resp, err := http.Get(tgt)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ResultQuerySuper
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}
