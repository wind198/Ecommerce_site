package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//defind a slice of ProductPerCategory type
	sliceProductPerCategory := make([]ProductPerCategory, 0)
	//defind a slice of category ID
	sliceCategory := make(map[int]string)
	//select category from database and put it into the slice
	db := newDbHanlder.db
	q := `select categoryID, categoryName from category`
	rows, err := db.Query(q)
	if err != nil {
		log.Printf("Error during query category table, %v\n", err)
	}
	for rows.Next() {
		var id int
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			log.Printf("Error during scanning, %v\n", err)
		} else {
			sliceCategory[id] = name
		}
	}

	//for every categoryID, generate a ProductPerCategory struct
	for categoryID, categoryName := range sliceCategory {
		var aProductPerCategory = ProductPerCategory{
			CategoryName: categoryName,
		}
		//query products with that categoryID and update the struct
		q = fmt.Sprintf("select productID,productName,productPrice, productImageSource from product where categoryID=%v", categoryID)
		rows, err := db.Query(q)
		if err != nil {
			log.Printf("Error during query product table, %v, for categoryID: %v\n", err, categoryID)
		}
		for rows.Next() {
			var aProduct Product
			err = rows.Scan(&aProduct.ProductID, &aProduct.ProductName, &aProduct.ProductPrice, &aProduct.ProductImageSource)
			if err != nil {
				log.Printf("Error during scanning product with category id %v, %v", categoryID, err)
			}
			aProductPerCategory.ProductArray = append(aProductPerCategory.ProductArray, aProduct)
		}
		//append the struct to slice
		sliceProductPerCategory = append(sliceProductPerCategory, aProductPerCategory)
	}
	jsonOutput, err := json.Marshal(sliceProductPerCategory)
	if err != nil {
		log.Printf("Error during constructing json for ouput")
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonOutput)

}
