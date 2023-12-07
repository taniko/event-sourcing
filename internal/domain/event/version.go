package event

type Version int

func (v Version) Next() Version {
	return v + 1
}
