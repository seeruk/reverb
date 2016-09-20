package storage

import "github.com/SeerUK/reverb/model"

type Driver interface {
	FindAll(*[]model.Request) error
	Find(int, *model.Request) error
	Persist(*model.Request) error
}
