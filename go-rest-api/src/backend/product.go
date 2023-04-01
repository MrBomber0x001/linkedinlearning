package backend

import (
	"database/sql"
	"fmt"
)

type product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	Status      int    `json:"status"`
}

/* GET Products
* Open the connection
* Query against the connection
* return results
* Close the connection
 */
func getAllProducts(db *sql.DB) ([]product, error) {
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		return nil, fmt.Errorf("getAllProducts: error - %s", err.Error())
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product

		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Inventory, &p.Price, &p.Status); err != nil {
			return nil, fmt.Errorf("getAllProducts - scanning: error - %s", err.Error())
		}

		products = append(products, p)
	}
	return products, err
}
