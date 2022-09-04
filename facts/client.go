package facts

import (
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	BaseURL string
}

const ErrFactAPIProblem = "oh no"

func (c Client) Get() string {
	res, _ := http.Get(c.BaseURL)

	if res.StatusCode != http.StatusOK {
		return ErrFactAPIProblem
	}
	defer res.Body.Close()

	facts, err := factsFromResponse(res.Body)
	if err != nil {
		return ErrFactAPIProblem
	}

	if len(facts) == 0 {
		return ErrFactAPIProblem
	}

	return facts[0].Fact
}

func factsFromResponse(res io.Reader) (facts, error) {
	resBody, _ := io.ReadAll(res)
	var facts facts
	err := json.Unmarshal(resBody, &facts)
	if err != nil {
		return nil, err
	}
	return facts, nil
}

type factDTO struct {
	Fact string `json:"fact"`
}

type facts []factDTO
