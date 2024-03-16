package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Get(url string, token string, i interface{}) error {
	r, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	w, err := client.Do(r)

	if err != nil {
		return err
	}

	defer w.Body.Close()

	return json.NewDecoder(w.Body).Decode(&i)
}
