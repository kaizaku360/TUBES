package main

import (
	"fmt"
	"strings"
	"time"
)

type Peminjam struct {
	Nama           string
	JumlahPinjaman float64
	Tenor          time.Time
	Status         string
	TotalBunga     float64
	TotalBayar     float64
	CicilanBulanan float64
}

var dataPeminjam []Peminjam = []Peminjam{
	{
		Nama:           "Riku Asakura",
		JumlahPinjaman: 5000000,
		Tenor:          time.Date(2025, 12, 15, 23, 59, 59, 0, time.Local),
		Status:         "Sukses",
		TotalBunga:     0,
		TotalBayar:     0,
		CicilanBulanan: 0,
	},
}

func main() {
	fmt.Println("===== Selamat Datang di Aplikasi Pinjaman Online =====")
	fmt.Println("1. Mulai Aplikasi")
	fmt.Println("0. Keluar")

	var pilihan int
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		mainMenu()
	case 0:
		fmt.Println("Program selesai.")
	default:
		fmt.Println("Pilihan tidak tersedia.")
	}
}

func mainMenu() {
	for {
		fmt.Println("\n===== Menu Utama =====")
		fmt.Println("1. Manajemen Data Peminjam")
		fmt.Println("2. Perhitungan Pinjaman")
		fmt.Println("3. Pencarian Data Peminjam")
		fmt.Println("4. Pengurutan Daftar Peminjam")
		fmt.Println("5. Laporan Pinjaman")
		fmt.Println("0. Keluar")

		var pilih int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			manajemenData()
		case 2:
			perhitunganPinjaman()
		case 3:
			pencarianData()
		case 4:
			pengurutanData()
		case 5:
			laporanPinjaman()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func manajemenData() {
	for {
		fmt.Println("\n1. Tambah Data")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("0. Kembali")

		var pilih int
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData()
		case 2:
			ubahData()
		case 3:
			hapusData()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahData() {
	var p Peminjam

	fmt.Print("Nama Peminjam (ketik 0 untuk batal): ")
	fmt.Scan(&p.Nama)
	if p.Nama == "0" {
		fmt.Println("Batal menambah data.")
		return
	}

	fmt.Print("Jumlah Pinjaman: ")
	fmt.Scan(&p.JumlahPinjaman)

	fmt.Println("Status: 1. Sukses | 2. Belum Selesai")
	fmt.Print("Masukkan pilihan status (1/2): ")
	var statusInt int
	fmt.Scan(&statusInt)
	if statusInt == 1 {
		p.Status = "Sukses"
	} else {
		p.Status = "Belum Selesai"
	}

	p.Tenor = inputTanggal()
	dataPeminjam = append(dataPeminjam, p)
	fmt.Println("Data berhasil ditambahkan.")
}

func inputTanggal() time.Time {
	var d, m, y int
	fmt.Println("Masukkan tanggal jatuh tempo (ketik 0 untuk batal):")
	fmt.Print("Tanggal: ")
	fmt.Scan(&d)
	if d == 0 {
		return time.Now()
	}
	fmt.Print("Bulan: ")
	fmt.Scan(&m)
	if m == 0 {
		return time.Now()
	}
	fmt.Print("Tahun: ")
	fmt.Scan(&y)
	if y == 0 {
		return time.Now()
	}

	if d < 1 || d > 31 || m < 1 || m > 12 || y < 2024 {
		fmt.Println("Tanggal tidak valid, menggunakan hari ini.")
		return time.Now()
	}
	return time.Date(y, time.Month(m), d, 23, 59, 59, 0, time.Local)
}

func ubahData() {
	tampilkanData()

	var idx int
	fmt.Print("Pilih nomor data yang diubah (0 untuk kembali): ")
	fmt.Scan(&idx)
	if idx == 0 {
		return
	}
	if idx < 1 || idx > len(dataPeminjam) {
		fmt.Println("Index tidak valid.")
		return
	}
	idx--

	fmt.Print("Nama baru: ")
	fmt.Scan(&dataPeminjam[idx].Nama)

	fmt.Print("Jumlah Pinjaman baru: ")
	fmt.Scan(&dataPeminjam[idx].JumlahPinjaman)

	fmt.Print("Status baru: ")
	fmt.Scan(&dataPeminjam[idx].Status)

	fmt.Print("Ubah tanggal jatuh tempo? (yes/no): ")
	var ubah string
	fmt.Scan(&ubah)
	if strings.ToLower(ubah) == "yes" {
		dataPeminjam[idx].Tenor = inputTanggal()
	}
	fmt.Println("Data berhasil diubah.")
}

func hapusData() {
	tampilkanData()

	var idx int
	fmt.Print("Pilih nomor data yang ingin dihapus (0 untuk kembali): ")
	fmt.Scan(&idx)
	if idx == 0 {
		return
	}
	if idx < 1 || idx > len(dataPeminjam) {
		fmt.Println("Index tidak valid.")
		return
	}
	idx--
	dataPeminjam = append(dataPeminjam[:idx], dataPeminjam[idx+1:]...)
	fmt.Println("Data berhasil dihapus.")
}

func tampilkanData() {
	if len(dataPeminjam) == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i, p := range dataPeminjam {
		fmt.Printf("%d. Nama: %s | Pinjaman: Rp%.0f | Tenor: %s | Status: %s\n",
			i+1, p.Nama, p.JumlahPinjaman, p.Tenor.Format("2006-01-02"), p.Status)
	}
}

func perhitunganPinjaman() {
	tampilkanData()

	var pilih int
	fmt.Print("Pilih data untuk dihitung (0 untuk kembali): ")
	fmt.Scan(&pilih)
	if pilih == 0 {
		return
	}
	if pilih < 1 || pilih > len(dataPeminjam) {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	idx := pilih - 1
	const bunga = 0.05
	tenorBulan := int(dataPeminjam[idx].Tenor.Sub(time.Now()).Hours() / 24 / 30)

	if tenorBulan <= 0 {
		fmt.Println("Tenor sudah lewat atau kurang dari 1 bulan.")
		return
	}

	dataPeminjam[idx].TotalBunga = dataPeminjam[idx].JumlahPinjaman * bunga * float64(tenorBulan)
	dataPeminjam[idx].TotalBayar = dataPeminjam[idx].JumlahPinjaman + dataPeminjam[idx].TotalBunga
	dataPeminjam[idx].CicilanBulanan = dataPeminjam[idx].TotalBayar / float64(tenorBulan)

	fmt.Printf("\nNama: %s\nTotal Bunga: Rp%.0f\nTotal Bayar: Rp%.0f\nCicilan per Bulan: Rp%.0f\n",
		dataPeminjam[idx].Nama, dataPeminjam[idx].TotalBunga, dataPeminjam[idx].TotalBayar, dataPeminjam[idx].CicilanBulanan)
}

func pencarianData() {
	fmt.Print("Masukkan nama yang dicari (ketik 0 untuk kembali): ")
	var cari string
	fmt.Scan(&cari)
	if cari == "0" {
		return
	}

	ditemukan := false
	for i, p := range dataPeminjam {
		if strings.Contains(strings.ToLower(p.Nama), strings.ToLower(cari)) {
			fmt.Printf("Data ke-%d | Nama: %s | Pinjaman: Rp%.0f | Status: %s\n", i+1, p.Nama, p.JumlahPinjaman, p.Status)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

func pengurutanData() {
	fmt.Println("\nUrutkan berdasarkan:")
	fmt.Println("1. Jumlah Pinjaman")
	fmt.Println("2. Tanggal Jatuh Tempo")
	fmt.Println("0. Kembali")
	var kriteria int
	fmt.Print("Pilihan: ")
	fmt.Scan(&kriteria)
	if kriteria == 0 {
		return
	}

	fmt.Println("Metode sorting:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Println("0. Kembali")
	var metode int
	fmt.Print("Pilihan: ")
	fmt.Scan(&metode)
	if metode == 0 {
		return
	}

	switch {
	case kriteria == 1 && metode == 1:
		selectionSortJumlah()
	case kriteria == 1 && metode == 2:
		insertionSortJumlah()
	case kriteria == 2 && metode == 1:
		selectionSortTenor()
	case kriteria == 2 && metode == 2:
		insertionSortTenor()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
	fmt.Println("Data berhasil diurutkan.")
}

func selectionSortJumlah() {
	for i := 0; i < len(dataPeminjam)-1; i++ {
		min := i
		for j := i + 1; j < len(dataPeminjam); j++ {
			if dataPeminjam[j].JumlahPinjaman < dataPeminjam[min].JumlahPinjaman {
				min = j
			}
		}
		dataPeminjam[i], dataPeminjam[min] = dataPeminjam[min], dataPeminjam[i]
	}
}

func insertionSortJumlah() {
	for i := 1; i < len(dataPeminjam); i++ {
		key := dataPeminjam[i]
		j := i - 1
		for j >= 0 && dataPeminjam[j].JumlahPinjaman > key.JumlahPinjaman {
			dataPeminjam[j+1] = dataPeminjam[j]
			j--
		}
		dataPeminjam[j+1] = key
	}
}

func selectionSortTenor() {
	for i := 0; i < len(dataPeminjam)-1; i++ {
		min := i
		for j := i + 1; j < len(dataPeminjam); j++ {
			if dataPeminjam[j].Tenor.Before(dataPeminjam[min].Tenor) {
				min = j
			}
		}
		dataPeminjam[i], dataPeminjam[min] = dataPeminjam[min], dataPeminjam[i]
	}
}

func insertionSortTenor() {
	for i := 1; i < len(dataPeminjam); i++ {
		key := dataPeminjam[i]
		j := i - 1
		for j >= 0 && dataPeminjam[j].Tenor.After(key.Tenor) {
			dataPeminjam[j+1] = dataPeminjam[j]
			j--
		}
		dataPeminjam[j+1] = key
	}
}

func laporanPinjaman() {
	if len(dataPeminjam) == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	fmt.Println("\n===== Laporan Data Peminjam =====")
	for i, p := range dataPeminjam {
		fmt.Printf("%d. Nama: %s | Pinjaman: Rp%.0f | Tenor: %s | Status: %s\n",
			i+1, p.Nama, p.JumlahPinjaman, p.Tenor.Format("2006-01-02"), p.Status)
	}
}

