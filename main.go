package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Islem struct {
	ID     int
	Miktar float64
	Durum  string
}

func isci(id int, ch chan Islem) {
	miktar := 100 + rand.Float64()*(5000-100)
	durum := "Basarili"
	if rand.Float64() < 0.1 {
		durum = "Supheli"
	}

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	ch <- Islem{ID: id, Miktar: miktar, Durum: durum}
}

func main() {
	islemKanali := make(chan Islem)

	fmt.Println("Sistem başlatılıyor...")

	for i := 1; i <= 10; i++ {
		go isci(i, islemKanali)
	}

	fmt.Println("İşlemler bekleniyor...")

	for i := 1; i <= 10; i++ {
		gelenVeri := <-islemKanali
		fmt.Printf("İşlem Alındı - ID: %d | Tutar: %.2f | Durum: %s\n",
			gelenVeri.ID, gelenVeri.Miktar, gelenVeri.Durum)
	}

	fmt.Println("Tüm işlemler başarıyla tamamlandı.")
}
