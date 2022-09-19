package main

import (
	"github.com/GeeCache/consistenthash"
	"log"
	"strconv"
)

func main() {
	hash := consistenthash.New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	hash.Add("6", "4", "2")

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		result := hash.Get(k)
		log.Printf("k: %s, result: %s", k, result)
		if hash.Get(k) != v {
			log.Printf("Asking for %s, should have yielded %s", k, v)
		}
	}

	// Adds 8, 18, 28
	hash.Add("8")

	// 27 should now map to 8.
	testCases["27"] = "8"

	for k, v := range testCases {
		if k == "27" {
			ret := hash.Get(k)
			log.Printf("ret: %+v", ret)
		}
		if hash.Get(k) != v {
			log.Printf("Asking for %s, should have yielded %s", k, v)
		}
	}

}
