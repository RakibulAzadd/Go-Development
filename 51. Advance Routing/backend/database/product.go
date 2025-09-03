package database

var ProductList []Products

type Products struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}


func init() {

	pro1 := Products{
		ID:          1,
		Title:       "Orange",
		Description: "Fresh Orange",
		Price:       1.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQsa3K1PkfEgVzc6JeayE-uGwejpsBDBbVBUw&s",
	}
	pro2 := Products{
		ID:          2,
		Title:       "Apple",
		Description: "Green Apple",
		Price:       2.00,
		ImgUrl:      "https://static.vecteezy.com/system/resources/previews/012/086/172/non_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}

	pro3 := Products{
		ID:          3,
		Title:       "Banana",
		Description: "Yellow Banana",
		Price:       0.50,
		ImgUrl:      "https://cdn.jwplayer.com/v2/media/0QcqxCCL/poster.jpg?width=1280",
	}
	pro4 := Products{
		ID:          4,
		Title:       "Mango",
		Description: "Sweet Mango",
		Price:       1.20,
		ImgUrl:      "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	}
	pro5 := Products{
		ID:          5,
		Title:       "Grapes",
		Description: "Fresh Grapes",
		Price:       2.50,
		ImgUrl:      "https://ds.rokomari.store/rokomari110/ProductNew20190903/260X372/Graps_Fruit_Seed_Black_-Non_Brand-96fe8-359969.png",
	}
	pro6 := Products{
		ID:          6,
		Title:       "Pineapple",
		Description: "Sweet Pineapple",
		Price:       3.00,
		ImgUrl:      "https://www.healthxchange.sg/adobe/dynamicmedia/deliver/dm-aid--c06c2aed-90cf-4360-a423-7f053b2a44d9/pineapple-health-benefits-and-ways-to-enjoy.jpg?preferwebp=true",
	}
	ProductList = append(ProductList, pro1)
	ProductList = append(ProductList, pro2, pro3, pro4, pro5, pro6)

 
}
