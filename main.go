package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Islem struct {
	ID    int
	Durum string
}

func isci(id int, bant chan Islem) {
	// Her işçi rastgele bir sürede işini bitirsin (0-5 saniye arası)
	// Bazen çok yavaş kalabilirler
	gecikme := time.Duration(rand.Intn(5000)) * time.Millisecond
	time.Sleep(gecikme)

	bant <- Islem{ID: id, Durum: "Tamamlandi"}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	islemBandı := make(chan Islem)

	fmt.Println("Sistem başlatıldı. 10 işlem bekleniyor...")

	for i := 1; i <= 10; i++ {
		go isci(i, islemBandı)
	}

	for i := 1; i <= 10; i++ {
		// select ifadesi ile aynı anda birden fazla kanal operasyonu dinlenir:
		// 1. İşlem bandından gelen veri
		// 2. Belirlenen zaman aşımı süresinin dolması
		select {
		case veri := <-islemBandı:
			fmt.Printf("İşlem verisi alındı: ID %d\n", veri.ID)

		case <-time.After(3 * time.Second):
			// Eğer 3 saniye içinde işlem bandından veri gelmezse zaman aşımı tetiklenir.
			fmt.Println("HATA: İşlem hattı zaman aşımına uğradı, sistem sonlandırılıyor!")
			return // Programı tamamen kapat
		}
	}

	fmt.Println("Tüm işlemler başarıyla tamamlandı.")
}
