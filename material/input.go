package material

type InputMaterial struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}
