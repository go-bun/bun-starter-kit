package example

type Org struct {
	ID      int64 `bun:",autoincrement"`
	Name    string
	OwnerID int64
	Owner   *User `bun:"rel:has-one"`
}
