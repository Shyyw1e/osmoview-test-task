package domain

import (
	"time"
)

type Data struct {
	ID int `yaml:"id"`
	Name string `yaml:"name"`
	Tags []string `yaml:"tags"`
	Attributes map[string]int `yaml:"attributes"`
	Active bool `yaml:"active"`
	Score float64 `yaml:"score"`
	Timestamp time.Time `yaml:"timestamp"`
}

