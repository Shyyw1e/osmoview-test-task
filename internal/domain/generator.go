package domain

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var names = []string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
var tagPool = []string{"red", "green", "blue", "yellow", "white", "black"}

func RandomData(id int) Data {
	tags := make([]string, rand.Intn(3) + 1)
	for i := range tags {
		tags[i] = tagPool[rand.Intn(len(tagPool))]
	}

	attr := make(map[string]int)
	for i := 0; i < rand.Intn(3) + 1; i++ {
		key := tagPool[rand.Intn(len(tagPool))]
		attr[key] = rand.Intn(100)
	}
	
	return Data{
		ID: id,
		Name: names[rand.Intn(len(names))],
		Tags: tags,
		Attributes: attr,
		Active: rand.Intn(2) == 0,
		Score: rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}