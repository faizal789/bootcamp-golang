package main

import (
	"bootcamp-go/api/action"
	"bootcamp-go/api/model"
	"fmt"
)

func main() {
	var result int
	result, _ = action.HitungLuas(10, 10)
	fmt.Printf("hasil = %d ", result)
	action.HitungLuasPointer(20, 30, &result)

	fmt.Print(result)

    var nilai1 model.Point
    var nilai2 model.Point

    nilai1.X = 10
    nilai1.Y = 20

    nilai2.X = 30
    nilai2.Y = 40

    // nilai1 := model.Point{ X : 0, Y : 0}
    // nilai2 := model.Point{ X: 3, Y : 0}

    hasil,_ := action.HitungPanjang(nilai1,nilai2)
    fmt.Printf("\n panjangnya adalah %f ", hasil)

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
