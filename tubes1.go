package main

import (
	"bufio"
	"fmt"
	"os"
)

type Proyek struct {
	nama      string
	kategori  string
	target    float64
	terkumpul float64
	donatur   int
}

func main() {
	var proyek [100]Proyek
	var jumlahproyek int
	var pilihan int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n======================================")
		fmt.Println("          APLIKASI CROWDFUNDING       ")
		fmt.Println("======================================")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Tambah Donasi")
		fmt.Println("3. Hapus Proyek")
		fmt.Println("4. Cari Proyek (Sequential)")
		fmt.Println("5. Cari Proyek (Binary)")
		fmt.Println("6. Urutkan (Selection Sort)")
		fmt.Println("7. Urutkan (Insertion Sort)")
		fmt.Println("8. Tampilkan Semua Proyek")
		fmt.Println("9. Proyek Tercapai")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			if jumlahproyek >= 100 {
				fmt.Println("Batas proyek sudah penuh!")
				break
			}

			fmt.Print("Nama Proyek: ")
			scanner.Scan()
			proyek[jumlahproyek].nama = scanner.Text()

			fmt.Println("Pilih Kategori Proyek:")
			fmt.Println("1. Kesehatan")
			fmt.Println("2. Pendidikan")
			fmt.Println("3. Lingkungan")
			fmt.Println("4. Teknologi")
			fmt.Println("5. Sosial")
			fmt.Print("Pilih Kategori (1-5): ")
			var pilihanKategori int
			fmt.Scanln(&pilihanKategori)

			switch pilihanKategori {
			case 1:
				proyek[jumlahproyek].kategori = "Kesehatan"
			case 2:
				proyek[jumlahproyek].kategori = "Pendidikan"
			case 3:
				proyek[jumlahproyek].kategori = "Lingkungan"
			case 4:
				proyek[jumlahproyek].kategori = "Teknologi"
			case 5:
				proyek[jumlahproyek].kategori = "Sosial"
			default:
				fmt.Println("Kategori tidak valid, menggunakan kategori default: Umum")
				proyek[jumlahproyek].kategori = "Umum"
			}

			fmt.Print("Target Dana: ")
			fmt.Scanln(&proyek[jumlahproyek].target)
			proyek[jumlahproyek].terkumpul = 0
			proyek[jumlahproyek].donatur = 0

			tampilkandetail(proyek[jumlahproyek])
			jumlahproyek++
			fmt.Println("Proyek berhasil ditambahkan.")

		case 2:
			var nama string
			var donasi float64
			var ditemukan bool = false
			fmt.Print("Nama Proyek: ")
			scanner.Scan()
			nama = scanner.Text()

			for i := 0; i < jumlahproyek; i++ {
				if proyek[i].nama == nama {
					fmt.Print("Jumlah Donasi: ")
					fmt.Scanln(&donasi)
					proyek[i].terkumpul += donasi
					proyek[i].donatur++
					fmt.Println("Donasi berhasil dicatat.")
					tampilkandetail(proyek[i])
					ditemukan = true
					break
				}
			}

			if !ditemukan {
				fmt.Println("Proyek tidak ditemukan.")
			}

		case 3:
			var nama string
			var ditemukan bool = false
			fmt.Print("Nama Proyek yang akan dihapus: ")
			scanner.Scan()
			nama = scanner.Text()

			for i := 0; i < jumlahproyek; i++ {
				if proyek[i].nama == nama {
					for j := i; j < jumlahproyek-1; j++ {
						proyek[j] = proyek[j+1]
					}
					jumlahproyek--
					fmt.Println("Proyek berhasil dihapus.")
					ditemukan = true
					break
				}
			}

			if !ditemukan {
				fmt.Println("Proyek tidak ditemukan.")
			}

		case 4:
			var nama string
			var ditemukan bool = false
			fmt.Print("Nama Proyek yang dicari: ")
			scanner.Scan()
			nama = scanner.Text()

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

			var nama string
			fmt.Print("Nama Proyek yang dicari (binary search): ")
			scanner.Scan()
			nama = scanner.Text()

			low := 0
			high := jumlahproyek - 1
			ditemukan := false

			for low <= high {
				mid := (low + high) / 2
				if proyek[mid].nama == nama {
					tampilkandetail(proyek[mid])
					ditemukan = true
					break
				} else if proyek[mid].nama < nama {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}

			if !ditemukan {
				fmt.Println("Proyek tidak ditemukan.")
			}

		case 6:
			for i := 0; i < jumlahproyek-1; i++ {
				maxIdx := i
				for j := i + 1; j < jumlahproyek; j++ {
					if proyek[j].terkumpul > proyek[maxIdx].terkumpul {
						maxIdx = j
					}
				}
				proyek[i], proyek[maxIdx] = proyek[maxIdx], proyek[i]
			}
			fmt.Println("Proyek telah diurutkan berdasarkan dana terkumpul (descending).")

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
			fmt.Println("Proyek telah diurutkan berdasarkan jumlah donatur (descending).")

		case 8:
			if jumlahproyek == 0 {
				fmt.Println("Belum ada proyek yang terdaftar.")
			} else {
				fmt.Println("\nDAFTAR SEMUA PROYEK:")
				for i := 0; i < jumlahproyek; i++ {
					tampilkandetail(proyek[i])
					fmt.Println("----------------------")
				}
			}

		case 9:
			var adaProyek bool = false
			fmt.Println("\nPROYEK YANG SUDAH MENCAPAI TARGET:")
			for i := 0; i < jumlahproyek; i++ {
				if proyek[i].terkumpul >= proyek[i].target {
					tampilkandetail(proyek[i])
					fmt.Println("----------------------")
					adaProyek = true
				}
			}
			if !adaProyek {
				fmt.Println("Belum ada proyek yang mencapai target.")
			}

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi crowdfunding.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkandetail(p Proyek) {
	fmt.Printf("\nNama: %s\n", p.nama)
	fmt.Printf("Kategori: %s\n", p.kategori)
	fmt.Printf("Target: Rp%.2f\n", p.target)
	fmt.Printf("Terkumpul: Rp%.2f\n", p.terkumpul)
	fmt.Printf("Jumlah Donatur: %d\n", p.donatur)

	persen := (p.terkumpul / p.target) * 100
	if persen > 100 {
		persen = 100
	}
	fmt.Printf("Progress: %.2f%%\n", persen)
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
