package database

var ProductsList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is a fruit. I love it.",
		Price:       100,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is a fruit. I love it.",
		Price:       150,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is a fruit. I love it.",
		Price:       200,
		ImgUrl:      "https://desime.co.uk/cdn/shop/files/orange.jpg?v=1690113703",
	}

	ProductsList = append(ProductsList, prd1)
	ProductsList = append(ProductsList, prd2)
	ProductsList = append(ProductsList, prd3)
}
