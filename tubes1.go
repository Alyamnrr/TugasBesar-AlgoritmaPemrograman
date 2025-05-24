package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Proyek struct {
	id        int
	nama      string
	kategori  string
	target    float64
	terkumpul float64
	donatur   int
}
var IDberikutnya = 1

func main() {
	var proyek [100]Proyek
	var jumlahproyek int
	var pilihan int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n====== APLIKASI CROWDFUNDING ======")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tambah Donasi")
		fmt.Println("3. Hapus Proyek")
		fmt.Println("4. Cari Proyek (Sequential)")
		fmt.Println("5. Cari Proyek (Binary Search)")
		fmt.Println("6. Urutkan Dana (Selection Sort)")
		fmt.Println("7. Urutkan Donatur (Insertion Sort)")
		fmt.Println("8. Tampilkan Semua Proyek")
		fmt.Println("9. Proyek Tercapai")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")

		pilihan = inputInteger()


		switch pilihan {
		case 1:
			if jumlahproyek >= 100 {
				fmt.Println("Maaf, Sudah Penuh!.")
				break
			}

			proyek[jumlahproyek].id = IDberikutnya
			IDberikutnya++

			fmt.Print("Nama Proyek: ")
			scanner.Scan()
			nama := scanner.Text()

			if !ceknama(nama, proyek[:jumlahproyek]) {
				fmt.Println("Nama sudah terpakai. Coba yang lain.")
				IDberikutnya--
				break
			}
			proyek[jumlahproyek].nama = nama

			fmt.Println("Kategori:")
			fmt.Println("1. Kesehatan, 2. Pendidikan, 3. Lingkungan, 4. Teknologi, 5. Sosial")
			fmt.Print("Pilih (1-5): ")
			kat := inputIntegerRange(1, 5)
			kategoriList := []string{"Kesehatan", "Pendidikan", "Lingkungan", "Teknologi", "Sosial"}
			proyek[jumlahproyek].kategori = kategoriList[kat-1]

			fmt.Print("Target Dana (Rp): ")
			proyek[jumlahproyek].target = inputdanavalid()
			proyek[jumlahproyek].terkumpul = 0
			proyek[jumlahproyek].donatur = 0

			tampilkandetail(proyek[jumlahproyek])
			jumlahproyek++
			fmt.Println("Proyek berhasil ditambahkan!")

		case 2:
			fmt.Print("Nama Proyek: ")
			scanner.Scan()
			nama := scanner.Text()
			found := false
			for i := 0; i < jumlahproyek; i++ {
				if proyek[i].nama == nama {
					fmt.Print("Jumlah Donasi (Rp): ")
					donasi := inputdanavalid()
					proyek[i].terkumpul += donasi
					proyek[i].donatur++
					fmt.Println("Donasi berhasil. Terima Kasih atas donasimu!")
					tampilkandetail(proyek[i])
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Proyek tidak ditemukan, silahkan tambahkan proyek terlebih dahulu.")
			}

		case 3:
			fmt.Print("Nama Proyek yang ingin dihapus: ")
			scanner.Scan()
			nama := scanner.Text()
			dihapus := false
			for i := 0; i < jumlahproyek; i++ {
				if proyek[i].nama == nama {
					if konfirmasihapus(nama) {
						for j := i; j < jumlahproyek-1; j++ {
							proyek[j] = proyek[j+1]
						}
						jumlahproyek--
						fmt.Println("Proyek berhasil dihapus.")
					} else {
						fmt.Println("Hapus proyek dibatalkan.")
					}
					dihapus = true
					break
				}
			}
			if !dihapus {
				fmt.Println("Proyek tidak ditemukan.")
			}

		case 4:
			fmt.Print("Nama Proyek yang dicari: ")
			scanner.Scan()
			nama := scanner.Text()
			ditemukan := false
			for i := 0; i < jumlahproyek; i++ {
				if proyek[i].nama == nama {
					tampilkandetail(proyek[i])
					ditemukan = true
					break
				}
			}
			if !ditemukan {
				fmt.Println("Proyek tidak ditemukan.")
			}

		case 5:
			if jumlahproyek == 0 {
				fmt.Println("Belum ada proyek.")
				break
			}
			urutkanmanual(&proyek, jumlahproyek)
			fmt.Print("Nama Proyek (binary search): ")
			scanner.Scan()
			nama := scanner.Text()

			low := 0
			high := jumlahproyek - 1
			found := false
			for low <= high {
				mid := (low + high) / 2
				if proyek[mid].nama == nama {
					tampilkandetail(proyek[mid])
					found = true
					break
				} else if proyek[mid].nama < nama {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			if !found {
				fmt.Println("Tidak ditemukan.")
			}

		case 6:
			for i := 0; i < jumlahproyek-1; i++ {
				max := i
				for j := i + 1; j < jumlahproyek; j++ {
					if proyek[j].terkumpul > proyek[max].terkumpul {
						max = j
					}
				}
				proyek[i], proyek[max] = proyek[max], proyek[i]
			}
			fmt.Println("Proyek diurutkan berdasarkan dana terkumpul.")

			for i := 0; i < jumlahproyek; i++ {
				tampilkandetail(proyek[i])
				fmt.Println("----------------------------------")
			}

		case 7:
			for i := 1; i < jumlahproyek; i++ {
				key := proyek[i]
				j := i - 1
				for j >= 0 && proyek[j].donatur < key.donatur {
					proyek[j+1] = proyek[j]
					j--
				}
				proyek[j+1] = key
			}
			fmt.Println("Proyek diurutkan berdasarkan donatur.")

			for i := 0; i < jumlahproyek; i++ {
				tampilkandetail(proyek[i])
				fmt.Println("----------------------------------")
			}

		case 8:
			if jumlahproyek == 0 {
				fmt.Println("Belum ada proyek yang terdaftar.")
			} else {
				for i := 0; i < jumlahproyek; i++ {
					tampilkandetail(proyek[i])
					fmt.Println("----------------------------------")
				}
			}

		case 9:
			ada := false
			fmt.Println("Proyek yang sudah capai target:")
			for i := 0; i < jumlahproyek; i++ {
				if proyek[i].terkumpul >= proyek[i].target {
					tampilkandetail(proyek[i])
					fmt.Println("----------------------------------")
					ada = true
				}
			}
			if !ada {
				fmt.Println("Belum ada proyek yang sudah capai target.")
			}

		case 0:
			fmt.Println("Terima kasih sudah memakai aplikasi ini.")
			return

		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}

func tampilkandetail(p Proyek) {
	fmt.Printf("\nID: %d\n", p.id)
	fmt.Printf("Nama: %s\n", p.nama)
	fmt.Printf("Kategori: %s\n", p.kategori)
	fmt.Printf("Target: Rp%.2f\n", p.target)
	fmt.Printf("Terkumpul: Rp%.2f\n", p.terkumpul)
	fmt.Printf("Donatur: %d orang\n", p.donatur)

	persen := (p.terkumpul / p.target) * 100
	if persen > 100 {
		persen = 100
	}
	tampilkanProgressBar(persen)
}

func tampilkanProgressBar(p float64) {
	bar := "["
	for i := 0; i < 20; i++ {
		if i < int(p/5) {
			bar += "="
		} else {
			bar += " "
		}
	}
	bar += fmt.Sprintf("] %.1f%%", p)
	fmt.Println(bar)
}

func inputInteger() int {
	var i int
	for {
		_, err := fmt.Scanln(&i)
		if err == nil {
			return i
		}
		fmt.Print("Masukkan angka yang bener: ")
		var dummy string
		fmt.Scanln(&dummy)
	}
}

func inputIntegerRange(min, max int) int {
	for {
		val := inputInteger()
		if val >= min && val <= max {
			return val
		}
		fmt.Printf("Harus antara %d dan %d: ", min, max)
	}
}

func inputdanavalid() float64 {
	var dana float64
	for {
		_, err := fmt.Scanln(&dana)
		if err == nil && dana > 0 {
			return dana
		}
		fmt.Print("Masukkan nominal yang valid: Rp")
		var dummy string
		fmt.Scanln(&dummy)
	}
}

func ceknama(nama string, proyek []Proyek) bool {
	for _, p := range proyek {
		if strings.EqualFold(p.nama, nama) {
			return false
		}
	}
	return true
}

func konfirmasihapus(nama string) bool {
	var konfirmasi string
	fmt.Printf("Proyek akan dihapus. Apakah Anda yakin ingin menghapus '%s'? (ya/tidak): ", nama)
	fmt.Scanln(&konfirmasi)
	return strings.ToLower(konfirmasi) == "ya"
}

func urutkanmanual(proyek *[100]Proyek, jumlah int) {
	for i := 0; i < jumlah-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlah; j++ {
			if proyek[j].nama < proyek[minIdx].nama {
				minIdx = j
			}
		}
		proyek[i], proyek[minIdx] = proyek[minIdx], proyek[i]
	}
}
