package parser

type Parser interface {
	Parse() error
	GetString(...string) string
	GetInt(...string) int
}
