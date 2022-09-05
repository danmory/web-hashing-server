package storages

type Storage interface {
	Find(string) (string, error)
	Store(string) (string, error)
	Close() error
}

type Type int8

const (
	Memory Type = iota
	Database
)

func Get(stype Type) (Storage, error) {
	switch stype {
	case Memory:
		memStorage := createMemStorage()
		return memStorage, nil
	case Database:
		DBStorage := createDBStorage()
		return DBStorage, nil
	default:
		return nil, &storageError{reason: "Incorrect storage type was provided"}
	}
}
