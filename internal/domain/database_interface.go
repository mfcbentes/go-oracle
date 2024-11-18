package domain

type Database interface {
	Connect() error
	Close() error
}
