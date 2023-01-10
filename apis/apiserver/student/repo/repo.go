package repo

// Repo defines the storage interface set to combine multiple resource repos.
type Repo interface {
	StudentRepo() StudentRepo
	Close() error
}

var client Repo

// Client return the store client instance.
func Client() Repo {
	return client
}

// SetClient set the store client.
func SetClient(c Repo) {
	client = c
}