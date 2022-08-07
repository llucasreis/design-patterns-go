package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

const dataPath = "./internal/capitals.txt"

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData(dataPath)
		db := singletonDatabase{caps}
		if e == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	result := map[string]int{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line) != 2 {
			continue
		}
		city, population := line[0], line[1]

		p, err := strconv.Atoi(population)
		if err != nil {
			continue
		}
		result[city] = p
	}

	return result, nil
}
