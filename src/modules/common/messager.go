package common

type Messager interface {
	Source() string
	Text() string
}
