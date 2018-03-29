package main

import (
	"encoding/json"
	"os"
	"time"
)

func main() {
	path := getPath()
	cc := newCompletedConfig()
	cc.Configs = getConfigs(newCrawler(3 * time.Second))
	b, err := json.Marshal(cc)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(path+"/ishadowx.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	f.Write(b)
}
