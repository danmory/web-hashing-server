package storages

type storageError struct {
	reason string
}

func (err *storageError) Error() string {
	return err.reason
}
