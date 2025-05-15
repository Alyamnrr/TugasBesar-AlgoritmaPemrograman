package main

import "fmt"

func main() {
	var proyek [100]struct {
		nama      string
		kategori  string
		target    float64
		terkumpul float64
		donatur   int
	}
	var totalProyek int
	var pilihan int

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
			for {
				if totalProyek >= 100 {
					fmt.Println("Batas proyek sudah penuh!")
					break
				}

				fmt.Print("Nama Proyek: ")
				fmt.Scanln(&proyek[totalProyek].nama)

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
					proyek[totalProyek].kategori = "Kesehatan"
				case 2:
					proyek[totalProyek].kategori = "Pendidikan"
				case 3:
					proyek[totalProyek].kategori = "Lingkungan"
				case 4:
					proyek[totalProyek].kategori = "Teknologi"
				case 5:
					proyek[totalProyek].kategori = "Sosial"
				default:
					fmt.Println("Kategori tidak valid, menggunakan kategori default: Umum")
					proyek[totalProyek].kategori = "Umum"
				}

				fmt.Print("Target Dana: ")
				fmt.Scanln(&proyek[totalProyek].target)
				proyek[totalProyek].terkumpul = 0
				proyek[totalProyek].donatur = 0
				totalProyek++
				fmt.Println("Proyek berhasil ditambahkan.")

				var lanjut string
				fmt.Print("Tambah proyek lagi? (y/n): ")
				fmt.Scanln(&lanjut)
				if lanjut != "y" && lanjut != "Y" {
					break
				}
			}

		case 2:
			for {
				var nama string
				var donasi float64
				var ditemukan bool = false
				fmt.Print("Nama Proyek: ")
				fmt.Scanln(&nama)

				for i := 0; i < totalProyek; i++ {
					if proyek[i].nama == nama {
						fmt.Print("Jumlah Donasi: ")
						fmt.Scanln(&donasi)
						proyek[i].terkumpul += donasi
						proyek[i].donatur++
						fmt.Println("Donasi berhasil dicatat.")
						ditemukan = true
						break
					}
				}

				if !ditemukan {
					fmt.Println("Proyek tidak ditemukan.")
				}

				var lanjut string
				fmt.Print("Tambah donasi lagi? (y/n): ")
				fmt.Scanln(&lanjut)
				if lanjut != "y" && lanjut != "Y" {
					break
				}
			}

		case 3:
			for {
				var nama string
				var ditemukan bool = false
				fmt.Print("Nama Proyek yang akan dihapus: ")
				fmt.Scanln(&nama)

				for i := 0; i < totalProyek; i++ {
					if proyek[i].nama == nama {
						for j := i; j < totalProyek-1; j++ {
							proyek[j] = proyek[j+1]
						}
						totalProyek--
						fmt.Println("Proyek berhasil dihapus.")
						ditemukan = true
						break
					}
				}

				if !ditemukan {
					fmt.Println("Proyek tidak ditemukan.")
				}

				var lanjut string
				fmt.Print("Hapus proyek lain? (y/n): ")
				fmt.Scanln(&lanjut)
				if lanjut != "y" && lanjut != "Y" {
					break
				}
			}

		case 4:
			for {
				var nama string
				var ditemukan bool = false
				fmt.Print("Nama Proyek yang dicari: ")
				fmt.Scanln(&nama)

				for i := 0; i < totalProyek; i++ {
					if proyek[i].nama == nama {
						tampilkanDetail(proyek[i])
						ditemukan = true
						break
					}
				}

				if !ditemukan {
					fmt.Println("Proyek tidak ditemukan.")
				}

				var lanjut string
				fmt.Print("Cari proyek lain? (y/n): ")
				fmt.Scanln(&lanjut)
				if lanjut != "y" && lanjut != "Y" {
					break
				}
			}

		case 5:
			fmt.Println("Fitur ini memerlukan data yang telah terurut terlebih dahulu.")

		case 6:
			for i := 0; i < totalProyek-1; i++ {
				maxIdx := i
				for j := i + 1; j < totalProyek; j++ {
					if proyek[j].terkumpul > proyek[maxIdx].terkumpul {
						maxIdx = j
					}
				}
				proyek[i], proyek[maxIdx] = proyek[maxIdx], proyek[i]
			}
			fmt.Println("Proyek telah diurutkan berdasarkan dana terkumpul (descending).")

		case 7:
			for i := 1; i < totalProyek; i++ {
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
			if totalProyek == 0 {
				fmt.Println("Belum ada proyek yang terdaftar.")
			} else {
				fmt.Println("\nDAFTAR SEMUA PROYEK:")
				for i := 0; i < totalProyek; i++ {
					tampilkanDetail(proyek[i])
					fmt.Println("----------------------")
				}
			}

		case 9:
			var adaProyek bool = false
			fmt.Println("\nPROYEK YANG SUDAH MENCAPAI TARGET:")
			for i := 0; i < totalProyek; i++ {
				if proyek[i].terkumpul >= proyek[i].target {
					tampilkanDetail(proyek[i])
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

func tampilkanDetail(p struct {
	nama      string
	kategori  string
	target    float64
	terkumpul float64
	donatur   int
}) {
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
