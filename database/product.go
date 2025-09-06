package database

var productsList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productsList) + 1
	productsList = append(productsList, p)
	return p
}

func List() []Product {
	return productsList
}

func Get(productID int) *Product {
	for _, product := range productsList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productsList {
		if p.ID == product.ID {
			productsList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product

	for idx, p := range productsList {
		if p.ID != productID {
			tempList[idx] = p
		}
	}
	productsList = tempList
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

	productsList = append(productsList, prd1)
	productsList = append(productsList, prd2)
	productsList = append(productsList, prd3)
}
