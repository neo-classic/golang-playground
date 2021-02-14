package main

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

type Store interface {
	Open(string) (string, error)
}

func NewStorage(t StorageType) Store {
	switch t {
	case DiskStorage:
		return
	}
}

func newDiskStorage() Store {

}
