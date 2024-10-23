package types

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Username string `bun:"username"`
	Password string `bun:"password,unique,notnull"`
	ID       uint8  `bun:",pk,autoincrement"`
}
