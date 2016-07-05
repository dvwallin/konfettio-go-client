package konfettio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

func GetJSON(user_id int, pool_slug string, hash_key string) (target Response) {
	url := fmt.Sprintf("https://konfett.io/c/%d/%s/%s/jsonverbose", user_id, pool_slug, hash_key)
	fmt.Println(url)
	r, err := http.Get(url)
	if err != nil {
		log.Printf("error getting value from url: %s\n", err.Error())
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&target)
	return target
}

func GetText(user_id int, pool_slug string, hash_key string) (target string) {
	url := fmt.Sprintf("https://konfett.io/c/%d/%s/%s/text", user_id, pool_slug, hash_key)
	fmt.Println(url)
	r, err := http.Get(url)
	if err != nil {
		log.Printf("error getting value from url: %s\n", err.Error())
	}
	defer r.Body.Close()
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading value from url: %s\n", err.Error())
	}
	return string(bs)
}
