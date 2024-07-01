package db

type Table string

func (t Table) WithSuffix(suffix string) string {
	return string(t) + "." + suffix
}
