package function

import (
	"math/rand"
	"time"
)

func RandomNickname() string {
	var adjectives = []string{"Happy", "Silly", "Funny", "Crazy", "Smart", "Brave", "Gentle", "Lucky", "Sleepy", "Charming"}
	var nouns = []string{"Penguin", "Kangaroo", "Elephant", "Giraffe", "Tiger", "Monkey", "Lion", "Koala", "Panda", "Zebra"}

	rand.Seed(time.Now().UnixNano())
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return adj + noun
}
