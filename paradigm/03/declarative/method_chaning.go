package main

func main() {

	p := persistence{}

	p.Init().
		From(User{UserId: "Bong"}).
		Select().
		Where(nil).
		GroupBy(nil)

}

type User struct {
	UserId string `json:"user_id"`
}

type persistence struct {
}

func (c *persistence) Init() *persistence {

	return c
}

func (c *persistence) Select() *persistence {

	return c
}

func (c *persistence) From(persistence interface{}) *persistence {

	return c
}

func (c *persistence) Where(where interface{}) *persistence {

	return c
}

func (c *persistence) GroupBy(groupBy interface{}) *persistence {

	return c
}
