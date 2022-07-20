package shop

type Item struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

type ItemUpdateInput struct {
	Name  *string `json:"name" binding:"required"`
	Price *int    `json:"price" binding:"required"`
}
