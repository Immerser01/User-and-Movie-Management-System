package DatabaseStruct

import (
	_ "github.com/lib/pq"
)

type UserData struct {
	Name  string
	Email string
	//DOB      string
	Password string
	//Movies   []string
}
