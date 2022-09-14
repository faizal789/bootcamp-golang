package model

type Point struct{
	X float64
	Y float64
}

type Panjang struct{
	K1 Point
	K2 Point
}

type Panjangaar struct{
	Kord [2]Point
}
type Luasaar struct{
	Kord [3]Point
}


type Siswa struct{
	Id int
	Nama string
	Pelajaran []MataPelajaran
}

type MataPelajaran struct{
	Id int
	Nama string
	Nilai float32
}

type Kategori struct{
	KatId any
	DeptId int
	Name any
	Desc string
}


type Produk struct{
	Kategori Kategori
	ProdukId string	
	Nama string
	KatId any
	NameKategori any
}




