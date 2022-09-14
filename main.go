package main

import (
	// "bootcamp-go/api/action"
	"log"

	"github.com/jmoiron/sqlx"

	// "reflect"
	// "bootcamp-go/api/model"
	"bootcamp-go/api/model"
	// "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// conn := "postgres://postgres:rahasia123@localhost/bootcampgo?sslmode=disable"
	// db, err := sql.Open("postgres", conn)
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	"localhost","5432","postgres","rahasia123","bootcampgo")
	dbx,err := sqlx.Connect("postgres",connection)
	if err != nil {
		// fmt.Print(err.Error())
		log.Fatalln(err)
	} else {
		print("berhasil koneksi")
		// var userId int
		// err = db.QueryRow(`INSERT INTO tabletesting(nama) values ('Faizal') RETURNING id`).Scan(&userId)
		// resultKategori := PanggilKategori(db)
		// for i := range resultKategori{
		// 	fmt.Printf("\n %d | %d |%s| %s",resultKategori[i].KatId,resultKategori[i].DeptId,resultKategori[i].Name,resultKategori[i].Desc)
		// }

		// resultKategoriX := PanggilKategoridbx(dbx)
		// for i := range resultKategoriX{
		// 	fmt.Printf("\n %d | %d |%s| %s",resultKategoriX[i].KatId,resultKategoriX[i].DeptId,resultKategoriX[i].Name,resultKategoriX[i].Desc)
		// }


		resultProdukX := PanggilProdukdbx(dbx)
		for i := range resultProdukX{
			fmt.Printf("\n %s | %d |%s",resultProdukX[i].Nama,resultProdukX[i].KatId,resultProdukX[i].NameKategori)
		}

		// resultProduk := action.PanggilProdukAndKategori(db)
		// for i := range resultProduk{
		// 	fmt.Printf("\n | %s | %d | %s",resultProduk[i].Nama,resultProduk[i].Kategori.KatId,resultProduk[i].Kategori.Name)
		// }
	}

	
		



	// var result int
	// result, _ = action.HitungLuas(10, 10)
	// fmt.Printf("hasil = %d ", result)
	// action.HitungLuasPointer(20, 30, &result)

	// fmt.Print(result)

	// var nilai1 model.Point
	// var nilai2 model.Point

	// nilai1.X = 10
	// nilai1.Y = 20

	// nilai2.X = 30
	// nilai2.Y = 40

	// nilai1 := model.Point{ X : 0, Y : 0}
	// nilai2 := model.Point{ X: 3, Y : 0}

	// hasil,_ := action.HitungPanjang(nilai1,nilai2)
	// fmt.Printf("\n panjangnya adalah %f ", hasil)

	// action.HitungPanjangPointer(&nilai1,&nilai2,&hasil)
	// fmt.Printf("\n panjangnya adalah %f ", hasil)

	// var PanjangVar  model.Panjang
	// PanjangVar.K1 = nilai1
	// PanjangVar.K2 = nilai2
	// action.HitungPanjangPointerByAggregation(&PanjangVar,&hasil)
	// fmt.Printf("\n panjangnya adalah %f ", hasil)

	// var PanjangArr model.Panjangaar
	// PanjangArr.Kord[0] = nilai1
	// PanjangArr.Kord[1] = nilai2
	// action.HitungPanjangPointerByArray(&PanjangArr,&hasil)
	// fmt.Printf("\n panjangnya adalah %f ", hasil)

	// var Luasaar model.Luasaar
	// Luasaar.Kord[0] = nilai1
	// Luasaar.Kord[1] = nilai2
	// Luasaar.Kord[2] = model.Point{ X: 3, Y : 3}
	// action.HitungLuasPointerByArray(&Luasaar,&hasil)
	// fmt.Printf("\n luasnya adalah %f ", hasil)

	// action.Isimatrix()
}
	// func PanggilKategori(Koneksi *sql.DB) []model.Kategori {
	// 	rows , err := Koneksi.Query(`SELECT category_id, department_id, name, description
	// 		FROM public.category;`)
	// 	defer rows.Close()
	// 	var kategoriSlice []model.Kategori
	// 	for rows.Next(){
	// 		var Kategori model.Kategori
	// 		err = rows.Scan(&Kategori.KatId,&Kategori.DeptId,&Kategori.Name,&Kategori.Desc)
	// 		if err != nil {
	// 			fmt.Print(err.Error())
	// 		}
	// 		kategoriSlice = append(kategoriSlice, Kategori)
	// 	}

	// 	return kategoriSlice
		
	// }


	// func PanggilKategoridbx(dbCon *sqlx.DB) []model.Kategori  {
	// 	var kategoriSlice []model.Kategori
	// 	err := dbCon.Select(&kategoriSlice,`SELECT category_id as KatId, department_id as DeptId,
	// 	 name as Name, description as Desc
	// 	FROM public.category;`)
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 	}
	// 	return kategoriSlice
	// }


	func PanggilProdukdbx(dbCon *sqlx.DB) []model.Produk {
		var ProdukSlice []model.Produk
		err := dbCon.Select(&ProdukSlice,`
		SELECT DISTINCT product.name as Nama,product_category.category_id as KatId,category.name as NameKategori from product
		inner join product_category
		on product_category.product_id = product.product_id
		inner join category
		on category.category_id = product_category.category_id`)
		if err != nil {
			fmt.Print(err.Error())
		}
		return ProdukSlice
	}


