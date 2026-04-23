package main

import (
	"fmt"
	"math/rand"
)

type Islem struct {
	ID        int
	Kullanıcı int
	Miktar    float64
	Durum     string
}

func veriUret(id int) Islem {
	rasgeleMiktar := 100 + rand.Float64()*(5000-100)
	durum := "Başarılı"

	if rand.Float64() < 0.05 { // %5 oranında şüpheli işlem
		durum = "Supheli"
	}

	return Islem{
		ID:        id,
		Kullanıcı: rand.Intn(1000),
		Miktar:    rasgeleMiktar,
		Durum:     durum,
	}
}

func main() {
	fmt.Println("Veri üretimi başlıyor...")

	for i := 1; i <= 10; i++ {
		islem := veriUret(i)
		fmt.Printf(
			"ID: %-2d | Kullanıcı: %-3d | Miktar: %8.2f | Durum: %s\n",
			islem.ID,
			islem.Kullanıcı,
			islem.Miktar,
			islem.Durum,
		)
	}
}
