package action

import (
	"bootcamp-go/api/model"
	"fmt"
	"math"
)

func HitungLuas(panjang int, lebar int) (int, error) {
	return panjang * lebar, nil
}

func HitungLuasPointer(panjang int, lebar int, result *int) {
	*result = panjang * lebar
}

func HitungPanjang(kordinatSatu model.Point,kordinatDua model.Point)(float64, error) {
	sumX := (kordinatDua.X - kordinatSatu.X)
	sumY := (kordinatDua.Y - kordinatSatu.Y)
	return math.Sqrt((sumX * sumX) + (sumY *sumY)), nil
}

func HitungPanjangPointer(kordinatSatu *model.Point,kordinatDua *model.Point, hasil *float64){
	sumX := (kordinatDua.X - kordinatSatu.X)
	sumY := (kordinatDua.Y - kordinatSatu.Y)
	*hasil = math.Sqrt((sumX * sumX) + (sumY *sumY))
}

func HitungPanjangPointerByAggregation(Panjang *model.Panjang, hasil *float64){
	kordinatSatu := Panjang.K1
	kordinatDua := Panjang.K2
	sumX := (kordinatDua.X - kordinatSatu.X)
	sumY := (kordinatDua.Y - kordinatSatu.Y)
	*hasil = math.Sqrt((sumX * sumX) + (sumY *sumY))
}


func LatihanArray(){
	var arrX [2]float64
	arrX[0] = 2.3
	arrX[1] = 3.4

	for i := range arrX{
		fmt.Println(arrX[i])
	}


	SliceX := []float64{}

	for i := 0; i < len(SliceX); i++ {
		fmt.Println(SliceX[i])
	}
	for i := range SliceX {
		fmt.Print(SliceX[i])
	}
}

	func HitungPanjangPointerByArray(Panjangaar *model.Panjangaar, hasil *float64){
		panjangSatu := Panjangaar.Kord[0]
		panjangDua := Panjangaar.Kord[1]
		HitungPanjangPointer(&panjangSatu,&panjangDua,hasil)
	}

	func HitungLuasPointerByArray(panjang *model.Luasaar, hasil *float64){
		p1,_ := HitungPanjang(panjang.Kord[0],panjang.Kord[1])
		p2,_ := HitungPanjang(panjang.Kord[1],panjang.Kord[2])
		*hasil = p1 * p2
		print(hasil)
		
	}

	func Isimatrix()  {
		var x,y int
		fmt.Print("Masukkan ukuran kordinat x y :")
		fmt.Scanln(&x,&y)
		fmt.Printf("\n x = %d , y = %d ", x , y)
		matrix := make([][] int, x)
		for i := range matrix {
			matrix[i] = make([]int, y)
			for i2 := range matrix[i]{
				fmt.Printf("Masukkan data ke [%d][%d]", i,i2)
				fmt.Scanln(&matrix[i][i2])
			}
		}

		print("Hasil \n")
		for i := range matrix{
			for i2 := range matrix[i]{
				fmt.Println(matrix[i][i2])
			}
		}
		
	}
