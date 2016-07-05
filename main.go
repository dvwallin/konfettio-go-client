package konfettio

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetJSONFromUrl(user_id int, pool_slug string, hash_key string) string {
	url := fmt.Sprintf("https://konfett.io/c/%d/%s/%s/json", user_id, pool_slug, hash_key)
	r, err := http.Get(url)
	if err != nil {
		log.Printf("error getting value from url: %s\n", err.Error())
	}
	defer r.Body.Close()
	var target string
	json.NewDecoder(r.Body).Decode(target)
	return target
}
