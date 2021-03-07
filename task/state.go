package task

type Intersection struct {
	Id IntersectionId
	In []StreetId
	T  []int
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
	res := &State{s.Task, make([]Intersection, len(s.G))}
	for i := 0; i < len(s.G); i++ {
		in := make([]StreetId, len(s.G[i].In))
		copy(in, s.G[i].In)
		t := make([]int, len(s.G[i].T))
		copy(t, s.G[i].T)

		res.G[i] = Intersection{
			s.G[i].Id,
			in,
			t,
		}
	}

	return res
}

func (s *State) Write(file string) {

}
