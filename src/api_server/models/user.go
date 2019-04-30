package models

type User struct {
    Id int64 `db:"id" json:"id"`
    Name string `sql:"size:60" db:"name" json:"name"`
    Password string `sql:"size:60" db:"password" json:"password"`
    Language string `sql:"size:60" db:"language" json:"language"`
}
