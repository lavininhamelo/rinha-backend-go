package entities

type Person struct {
	ID       string   `db:"id" json:"id"`
	Nickname string   `db:"nickname" json:"apelido"`
	Name     string   `db:"name" json:"nome"`
	Birthday string   `db:"birthday" json:"nascimento"`
	Stack    []string `db:"stack" json:"stack"`
}
