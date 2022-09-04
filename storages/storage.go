package storages

type Storage interface {
	Find(string) (string, error)
	Store(string) (string, error)
}

type Type int8

const (
	Memory Type = iota
	Database
)

func Get(stype Type) (Storage, error) {
	switch stype {
	case Memory:
		memStorage := new(memoryStorage)
		memStorage.data = make(map[string]string)
		return memStorage, nil
	case Database:
		return nil, nil
	default:
		return nil, &storageError{reason: "Incorrect storage type was provided"}
	}
}
