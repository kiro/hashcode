package task

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
)

type StreetId int
type IntersectionId int
type CarId int

type Street struct {
	Id     StreetId
	Name   string
	From   IntersectionId
	To     IntersectionId
	Length int
}

type Car struct {
	Id   CarId
	Path []StreetId
}

type Task struct {
	T       int
	N       int
	M       int
	C       int
	F       int
	Cars    []Car
	Streets []Street
}

// TODO(kiro): Remove unused intersections if needed
func Read(file string) *Task {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	b := bufio.NewScanner(bytes.NewBuffer(data))
	b.Split(bufio.ScanWords)

	t := &Task{}

	t.T = nextInt(b)
	t.N = nextInt(b)
	t.M = nextInt(b)
	t.C = nextInt(b)
	t.F = nextInt(b)
	t.Streets = make([]Street, t.M)
	t.Cars = make([]Car, t.C)

	streetId := map[string]StreetId{}

	for i := 0; i < t.M; i++ {
		s := Street{}
		s.From = IntersectionId(nextInt(b))
		s.To = IntersectionId(nextInt(b))
		s.Name = nextString(b)
		s.Id = StreetId(i)
		s.Length = nextInt(b)
		t.Streets[i] = s
		streetId[s.Name] = s.Id
	}

	for i := 0; i < t.C; i++ {
		n := nextInt(b)
		path := make([]StreetId, n)
		for j := 0; j < n; j++ {
			name := nextString(b)
			path[j] = streetId[name]
		}
		t.Cars[i] = Car{CarId(i), path}
	}

	return t
}

func nextInt(b *bufio.Scanner) int {
	b.Scan()
	value, err := strconv.Atoi(b.Text())
	if err != nil {
		panic(err)
	}
	return value
}

func nextString(b *bufio.Scanner) string {
	b.Scan()
	return b.Text()
}
