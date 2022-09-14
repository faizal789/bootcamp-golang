package action

import (
	"bootcamp-go/api/model"
	"database/sql"
	"fmt"
)

func PanggilProdukAndKategori(Koneksi *sql.DB) []model.Produk {
	query, err := Koneksi.Query(`
	SELECT product.name,product_category.category_id,category.name from product
	LEFT join product_category
	on product_category.product_id = product.product_id
	left join category
	on category.category_id = product_category.category_id`)
	defer query.Close()
	var produkSlice []model.Produk
	for query.Next(){
		var produk model.Produk
		err = query.Scan(&produk.Nama,&produk.Kategori.KatId,&produk.Kategori.Name)
		if err != nil {
			fmt.Print(err.Error())
		}
		produkSlice = append(produkSlice, produk)
	}
	return produkSlice
}