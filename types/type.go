package types

type ConfigType interface {
	Type() string
	IsValid(key string) bool
	Parse(key string, value string)
}