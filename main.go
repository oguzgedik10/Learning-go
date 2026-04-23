package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Islem struct {
	ID     int
	Miktar float64
	Durum  string
}

func veriUret(id int, wg *sync.WaitGroup) {
	defer wg.Done() // iş bitince bitirdim mesaji gönderir

	miktar := 100 + rand.Float64()*(5000-100)
	durum := "Başarılı"

	if rand.Float64() <= 0.1 {
		durum = "Şüpheli"
	}

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) //gecikme

	fmt.Printf("ID: %d | Miktar: %.2f | Durum: %s\n", id, miktar, durum)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Eşzamanlı üretim hattı başlatıldı...")

	for i := 1; i <= 50; i++ {
		wg.Add(1)           // Yeni bir işçi yola çıkıyor
		go veriUret(i, &wg) // Başına 'go' koyduk, artık arka planda çalışıyor
	}

	wg.Wait() // Tüm işçiler (50 tane) 'Done' diyene kadar bekle
	fmt.Println("Tüm işlemler tamamlandı.")
}
