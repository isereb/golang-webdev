package services

type BaseService struct {
	Name string
}

var makes = make([]Make, 0, 3)

func init() {

	makes = append(makes, initToyota())
	makes = append(makes, initMercedes())
	makes = append(makes, initChevrolet())

}

func initToyota() Make {
	toyota := Make{
		Id:   1,
		Name: "Toyota",
	}
	models := []Model{
		{
			Id:   1,
			Name: "RAV-4",
			Make: &toyota,
		},
		{
			Id:   2,
			Name: "Corolla",
			Make: &toyota,
		},
		{
			Id:   3,
			Name: "Camry",
			Make: &toyota,
		},
	}
	toyota.Models = models
	return toyota
}

func initMercedes() Make {
	mercedes := Make{
		Id:   2,
		Name: "Mercedes",
	}
	models := []Model{
		{
			Id:   1,
			Name: "C-class",
			Make: &mercedes,
		},
		{
			Id:   2,
			Name: "E-class",
			Make: &mercedes,
		},
		{
			Id:   3,
			Name: "S-class",
			Make: &mercedes,
		},
	}
	mercedes.Models = models
	return mercedes
}

func initChevrolet() Make {
	chevrolet := Make{
		Id:   3,
		Name: "Chevrolet",
	}
	models := []Model{
		{
			Id:   1,
			Name: "Camaro",
			Make: &chevrolet,
		},
		{
			Id:   2,
			Name: "Corvette",
			Make: &chevrolet,
		},
		{
			Id:   3,
			Name: "Cruze",
			Make: &chevrolet,
		},
	}
	chevrolet.Models = models
	return chevrolet
}
