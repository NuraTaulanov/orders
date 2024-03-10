package models

type Configuration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

// Order type corresponds to the 'orders' table in the DB and defines its schema
type Order struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id           int64  `db:"order_id"`
	CustomerName string `db:"customer_name"`
	Status       string `db:"status"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
}

// OrderProduct type corresponds to the 'order_products' table in the DB and defines its schema.
// Can't specify foreign key constraints as currently, there is no support for foreign keys in gorp.
type OrderProduct struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id           int64  `db:"order_product_id"`
	OrderId      int64  `db:"order_id"`
	ProductId    int64  `db:"product_id"`
	ProductEan   string `db:"product_ean"`
	CustomerName string `db:"customer_name"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
}

// Customer type corresponds to the 'customers' table in the DB and defines its schema
type Customer struct {
	Id           int64  `db:"customer_id"`
	CustomerName string `db:"customer_name"`
	Address      string `db:"address"`
	Phone        int64  `db:"phone"`
	CreatedAt    int64  `db:"created_at"`
	UpdatedAt    int64  `db:"updated_at"`
}

// Product type corresponds to the 'products' table in the DB and defines its schema
type Product struct {
	Id          int64  `db:"product_id" json:"product_id"`
	ProductName string `db:"product_name"`
	EanBarcode  string `db:"product_ean" json:"product_ean"`
	Category    string `db:"category"`
	SubCategory string `db:"sub_category"`
	Description string `db:"description"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
}
