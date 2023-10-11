package domain

import "errors"

const (
	TagTypeTask tagType = 1 << iota
	TagTypeProject
)

type tagType int

var (
	ErrTagTypeAlreadyExisted = errors.New("tag type already existed")
	ErrTagNotExisted         = errors.New("tag type not exist in target tag")
	ErrTagOutOfRange         = errors.New("tag number out of range, or tag not existed in the system")
)

type Tag struct {
	ID      int
	Color   string
	Name    string
	TagType int
}

func (t *Tag) AddType(i tagType) error {
	if (int(i) & t.TagType) == int(i) {
		return ErrTagTypeAlreadyExisted
	}
	if i > TagTypeProject {
		return ErrTagOutOfRange
	}
	t.TagType = t.TagType | int(i)
	return nil
}

func (t *Tag) RemoveType(i tagType) error {
	// placeholder for type conversion of i to int
	ii := int(i)
	if (ii & t.TagType) != ii {
		return ErrTagNotExisted
	}
	t.TagType = t.TagType ^ ii
	return nil
}
