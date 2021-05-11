package example

type User struct {
	ID   int64 `bun:",autoincrement"`
	Name string
}
