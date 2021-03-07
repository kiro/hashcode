package task

import (
	"fmt"
	"io"
	"os"
)

type Intersection struct {
	Id IntersectionId
	In []StreetId
	T  []int
}

func (i *Intersection) Clone() *Intersection {
	in := make([]StreetId, len(i.In))
	copy(in, i.In)
	t := make([]int, len(i.T))
	copy(t, i.T)

	return &Intersection{
		i.Id,
		in,
		t,
	}
}

func (i *Intersection) Used() int {
	count := 0
	for _, v := range i.T {
		if v != 0 {
			count++
		}
	}
	return count
}

type State struct {
	Task *Task
	G    []Intersection
}

func MkState(task *Task) *State {
	used := make(map[StreetId]bool)

	for _, car := range task.Cars {
		for _, id := range car.Path {
			used[id] = true
		}
	}

	c := make(map[IntersectionId]int)
	for _, street := range task.Streets {
		if used[street.Id] {
			c[street.To] = c[street.To] + 1
		}
	}

	s := &State{task, make([]Intersection, task.N)}
	for id, n := range c {
		s.G[id].In = make([]StreetId, 0, n)
		s.G[id].T = make([]int, n)
	}

	for _, street := range task.Streets {
		if used[street.Id] {
			s.G[street.To].In = append(s.G[street.To].In, street.Id)
		}
	}

	return s
}

func (s *State) Clone() *State {
	res := &State{s.Task, make([]Intersection, 0, len(s.G))}
	for _, i := range s.G {
		res.G = append(res.G, *i.Clone())
	}

	return res
}

func (s *State) Write(file string) {
	count := 0
	for _, i := range s.G {
		if i.Used() != 0 {
			count++
		}
	}

	out, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	println(out, count)

	for _, i := range s.G {
		if c := i.Used(); c != 0 {
			println(out, i.Id)
			println(out, c)

			for j := 0; j < len(i.In); j++ {
				if i.T[j] != 0 {
					street := s.Task.Streets[i.In[j]].Name
					println(out, street, i.T[j])
				}
			}
		}
	}

	err = out.Close()
	if err != nil {
		panic(err)
	}
}

func println(out io.Writer, args ...interface{}) {
	_, err := fmt.Fprintln(out, args...)
	if err != nil {
		panic(err)
	}
}
