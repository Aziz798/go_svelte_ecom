package product

import (
	"database/sql"
	"fmt"
	"go_ecom/internal/models"
)

func createProduct(product models.Product, db *sql.DB) (string, error) {
	res, err := db.Exec("INSERT INTO Products (Title,Description,Price,Quantity,ImageURL) VALUES(?,?,?,?,?);", product.Title, product.Description, product.Price, product.Quantity, product.ImageURL)
	if err != nil {
		return string(""), err
	}
	productID, err := res.LastInsertId()
	if err != nil {
		return string(""), err
	}
	return fmt.Sprintf("%d", productID), nil
}
