package internal

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

const dataPath = "./internal/capitals.txt"

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

type dummyDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func (db *dummyDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

func GetTotalPopulation(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
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

func GetDummyDatabase() *dummyDatabase {
	db := dummyDatabase{}
	db.capitals = map[string]int{
		"Seoul": 5000000,
		"Tokyo": 2000000,
		"Delhi": 10000000,
	}

	return &db
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
