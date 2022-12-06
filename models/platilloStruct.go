package models

type Platillo struct {
	Id          int
	Nombre      string
	Descripcion string
	Precio      int
}

func ReturnPlatillo() Platillo {
	return Platillo{}
}
