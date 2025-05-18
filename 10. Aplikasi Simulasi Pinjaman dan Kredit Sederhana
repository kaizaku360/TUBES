package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// -----------STRUCT DATA PEMINJAM-----------
type sistem struct {
	Nama            string
	Jumlah_Pinjaman float64
	Tenor           time.Time
	Status          string
	TotalBunga      float64
	TotalBayar      float64
	CicilanBulanan  float64
}

var sistems []sistem

// -----------MENU UTAMA-----------
func mainMenu() {
	for {
		fmt.Println("Pilih menu: ")
		fmt.Println("1. Manajemen Data Peminjam")
		fmt.Println("2. Perhitungan Pinjaman")
		fmt.Println("3. Pencarian Data Peminjam")
		fmt.Println("4. Pengurutan Daftar Peminjam")
		fmt.Println("5. Laporan Pinjaman")
		fmt.Println("0. Keluar")

		var pilih int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilih)

		switch {
		case pilih == 1:
			ManajenDataPinjaman()
		case pilih == 2:
			PerhitunganPinjaman()
		case pilih == 3:
			Pencariandatapinjaman()
		case pilih == 4:
			PengurutanDatapeminjam()
		case pilih == 5:
			LaporanPinjaman()
		case pilih == 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			os.Exit(0)
		}
	}
}

// -----------FITUR CRUD-----------
func ManajenDataPinjaman() {
	fmt.Println("Pilih menu: ")
	fmt.Println("1. Tambah Data Peminjam")
	fmt.Println("2. Hapus Data Peminjam")
	fmt.Println("3. Ubah Data Peminjam")

	var pilih int
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilih)
	switch {
	case pilih == 1:
		tambahData()
	case pilih == 2:
		hapusData()
	case pilih == 3:
		ubahData()
	}
}

// -----------INPUT TANGGAL JATUH TEMPO-----------
func jatuhTempo() time.Time {
	var tanggal, bulan, tahun int
	fmt.Scanln()
	fmt.Println("Masukkan tanggal jatuh tempo (hari, bulan, tahun): ")
	fmt.Print("Tanggal: ")
	fmt.Scan(&tanggal)
	fmt.Print("Bulan: ")
	fmt.Scan(&bulan)
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	if tanggal < 1 || tanggal > 31 || bulan < 1 || bulan > 12 || tahun < 2024 {
		fmt.Println("Tanggal Tidak Valid.")
		return time.Now()
	}

	deadline := time.Date(tahun, time.Month(bulan), tanggal, 23, 59, 59, 0, time.Local)
	return deadline
}

// -----------HITUNG SISA HARI-----------
func sisaHari(deadline time.Time) int {
	hariIni := time.Now()
	durasi := deadline.Sub(hariIni)
	return int(durasi.Hours() / 24)
}

// -----------TAMBAH DATA-----------
func tambahData() {
	var data sistem
	fmt.Print("Masukkan Nama Peminjam: ")
	fmt.Scan(&data.Nama)

	fmt.Print("Masukan Jumlah Pinjaman: ")
	fmt.Scan(&data.Jumlah_Pinjaman)

	fmt.Print("Status Pembayaran (Sukses/Belum Selesai): ")
	fmt.Scan(&data.Status)

	data.Tenor = jatuhTempo()
	sistems = append(sistems, data)

	fmt.Println("\nData berhasil ditambahkan:")
	fmt.Printf("Nama: %s\n", data.Nama)
	fmt.Printf("Jumlah Pinjaman: Rp.%.2f\n", data.Jumlah_Pinjaman)
	fmt.Printf("Status: %s\n", data.Status)
	fmt.Printf("Tenor: %s\n", data.Tenor.Format("02 January 2006"))
}

// -----------HAPUS DATA-----------
func hapusData() {
	if len(sistems) == 0 {
		fmt.Println("Data tidak ada")
		return
	}
	fmt.Println("\nHapus Data:")
	for i, sistem := range sistems {
		fmt.Printf("%d. nama peminjam: %s\nJumlah Pinjaman: RP.%.2f\nTenor: %s\nStatus: %s\n",
			i+1,
			sistem.Nama,
			sistem.Jumlah_Pinjaman,
			sistem.Tenor.Format("2006-01-02"),
			sistem.Status)
	}
	var idx int
	fmt.Print("Pilih Data yang ingin dihapus: ")
	fmt.Scan(&idx)
	if idx < 1 || idx > len(sistems) {
		fmt.Println("Data tidak ada")
		return
	}
	idx--
	sistems = append(sistems[:idx], sistems[idx+1:]...)
	fmt.Println("Data Berhasil Dihapus")
}

// -----------UBAH DATA-----------
func ubahData() {
	if len(sistems) == 0 {
		fmt.Println("Data tidak ada")
		return
	}
	fmt.Println("\nUbah Data:")
	for i, sistem := range sistems {
		fmt.Printf("%d. nama peminjam: %s\nJumlah Pinjaman: %.2f\nTenor: %s\nStatus: %s\n", i+1, sistem.Nama, sistem.Jumlah_Pinjaman, sistem.Tenor, sistem.Status)
	}
	var idx int
	fmt.Print("Pilih Data yang ingin diubah: ")
	fmt.Scan(&idx)
	if idx < 1 || idx > len(sistems) {
		fmt.Println("Data tidak ada")
		return
	}
	idx--

	fmt.Println("\nData yang akan diubah:")
	fmt.Print("Nama Peminjam Baru: ")
	fmt.Scan(&sistems[idx].Nama)
	fmt.Print("Jumlah Pinjaman Baru: ")
	fmt.Scan(&sistems[idx].Jumlah_Pinjaman)
	fmt.Print("Masukan Status Terbaru: ")
	fmt.Scan(&sistems[idx].Status)
	fmt.Print("Ubah Deadline (yes/no): ")

	var ubahdeadline string
	fmt.Scan(&ubahdeadline)
	if strings.ToLower(ubahdeadline) == "yes" {
		sistems[idx].Tenor = jatuhTempo()
	}
	fmt.Println("Data Berhasil Diubah")
}

// -----------PERHITUNGAN PINJAMAN-----------
func PerhitunganPinjaman() {
	if len(sistems) == 0 {
		fmt.Println("data belum ada, silakan tambah dulu")
		return
	}

	fmt.Println("\npilih data peminjam yang mau dihitung:")
	for i, s := range sistems {
		fmt.Printf("%d. nama: %s, pinjaman: rp%.0f, tenor: %s\n", i+1, s.Nama, s.Jumlah_Pinjaman, s.Tenor.Format("2006-01-02"))
	}

	var pilih int
	fmt.Print("masukkan nomor: ")
	fmt.Scan(&pilih)

	if pilih < 1 || pilih > len(sistems) {
		fmt.Println("nomor tidak valid")
		return
	}

	idx := pilih - 1
	const bunga = 0.05
	tenorBulan := int(sistems[idx].Tenor.Sub(time.Now()).Hours() / 24 / 30)
	if tenorBulan <= 0 {
		fmt.Println("tenor sudah lewat atau kurang dari 1 bulan")
		return
	}

	sistems[idx].TotalBunga = sistems[idx].Jumlah_Pinjaman * bunga * float64(tenorBulan)
	sistems[idx].TotalBayar = sistems[idx].Jumlah_Pinjaman + sistems[idx].TotalBunga
	sistems[idx].CicilanBulanan = sistems[idx].TotalBayar / float64(tenorBulan)

	fmt.Println("\n=== hasil perhitungan ===")
	fmt.Printf("nama: %s\n", sistems[idx].Nama)
	fmt.Printf("jumlah pinjaman: rp%.0f\n", sistems[idx].Jumlah_Pinjaman)
	fmt.Printf("tenor: %d bulan\n", tenorBulan)
	fmt.Printf("total bunga: rp%.0f\n", sistems[idx].TotalBunga)
	fmt.Printf("total bayar: rp%.0f\n", sistems[idx].TotalBayar)
	fmt.Printf("cicilan per bulan: rp%.0f\n", sistems[idx].CicilanBulanan)
}

// -----------PENCARIAN DATA-----------
func Pencariandatapinjaman() {
	if len(sistems) == 0 {
		fmt.Println("Data tidak ada")
		return
	}
	var keyword string
	fmt.Println("\n===== Pencarian Data Peminjam =====")
	fmt.Print("Masukkan nama peminjam yang ingin dicari: ")
	fmt.Scan(&keyword)

	ditemukan := false
	fmt.Println("")

	for i, sistem := range sistems {
		if strings.Contains(strings.ToLower(sistem.Nama), strings.ToLower(keyword)) {
			ditemukan = true
			fmt.Printf("Data ke-%d\n", i+1)
			fmt.Printf("Nama Peminjam   : %s\n", sistem.Nama)
			fmt.Printf("Jumlah Pinjaman : RP.%.2f\n", sistem.Jumlah_Pinjaman)
			fmt.Printf("Tenor           : %s\n", sistem.Tenor.Format("2006-01-02 15:04:05"))
			fmt.Printf("Status          : %s\n", sistem.Status)
			fmt.Println("--------------------------------")
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

// -----------PENGURUTAN DATA PEMINJAM-----------
func PengurutanDatapeminjam() {
	if len(sistems) == 0 {
		fmt.Println("data masih kosong")
		return
	}

	fmt.Println("\n=== pilih urutan berdasarkan apa ===")
	fmt.Println("1. jumlah pinjaman")
	fmt.Println("2. tenor (tanggal jatuh tempo)")
	var pilihan int
	fmt.Print("masukkan pilihan: ")
	fmt.Scan(&pilihan)

	fmt.Println("\n=== pilih metode sorting ===")
	fmt.Println("1. selection sort")
	fmt.Println("2. insertion sort")
	var metode int
	fmt.Print("masukkan metode: ")
	fmt.Scan(&metode)

	if pilihan == 1 {
		if metode == 1 {
			selectionSortJumlahPinjaman()
		} else {
			insertionSortJumlahPinjaman()
		}
	} else if pilihan == 2 {
		if metode == 1 {
			selectionSortTenor()
		} else {
			insertionSortTenor()
		}
	} else {
		fmt.Println("pilihan tidak valid")
		return
	}

	fmt.Println("\n=== data setelah diurutkan ===")
	for i, s := range sistems {
		fmt.Printf("%d. nama: %s, pinjaman: rp%.0f, tenor: %s\n", i+1, s.Nama, s.Jumlah_Pinjaman, s.Tenor.Format("2006-01-02"))
	}
}

// -----------SELECTION SORT JUMLAH PINJAMAN-----------
func selectionSortJumlahPinjaman() {
	for i := 0; i < len(sistems)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(sistems); j++ {
			if sistems[j].Jumlah_Pinjaman < sistems[minIdx].Jumlah_Pinjaman {
				minIdx = j
			}
		}
		sistems[i], sistems[minIdx] = sistems[minIdx], sistems[i]
	}
}

// -----------INSERTION SORT JUMLAH PINJAMAN-----------
func insertionSortJumlahPinjaman() {
	for i := 1; i < len(sistems); i++ {
		temp := sistems[i]
		j := i - 1
		for j >= 0 && sistems[j].Jumlah_Pinjaman > temp.Jumlah_Pinjaman {
			sistems[j+1] = sistems[j]
			j--
		}
		sistems[j+1] = temp
	}
}

// -----------SELECTION SORT TENOR-----------
func selectionSortTenor() {
	for i := 0; i < len(sistems)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(sistems); j++ {
			if sistems[j].Tenor.Before(sistems[minIdx].Tenor) {
				minIdx = j
			}
		}
		sistems[i], sistems[minIdx] = sistems[minIdx], sistems[i]
	}
}

// -----------INSERTION SORT TENOR-----------
func insertionSortTenor() {
	for i := 1; i < len(sistems); i++ {
		temp := sistems[i]
		j := i - 1
		for j >= 0 && sistems[j].Tenor.After(temp.Tenor) {
			sistems[j+1] = sistems[j]
			j--
		}
		sistems[j+1] = temp
	}
}

// -----------LAPORAN PINJAMAN-----------
func LaporanPinjaman() {
	if len(sistems) == 0 {
		fmt.Println("Data tidak ada")
		return
	}

	fmt.Println("\n===== Laporan Pinjaman =====")
	for i, sistem := range sistems {
		fmt.Printf("Data ke-%d\n", i+1)
		fmt.Printf("Nama Peminjam   : %s\n", sistem.Nama)
		fmt.Printf("Jumlah Pinjaman : RP.%.2f\n", sistem.Jumlah_Pinjaman)
		fmt.Printf("Tenor           : %s\n", sistem.Tenor.Format("2006-01-02 15:04:05"))
		fmt.Printf("Status          : %s\n", sistem.Status)
		fmt.Println("-------------------------------")
	}
}

// -----------MAIN PROGRAM-----------
func main() {
	fmt.Println("=====Selamat Datang Di Aplikasi Pinjaman Online======")
	fmt.Println("Pilih 1 Untuk Melanjutkan")
	fmt.Println("Pilihn 0 Untuk Mengakhiri")

	var pilih int
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		mainMenu()
	} else if pilih == 0 {
		fmt.Println("Program Telah Berakhir")
		fmt.Println()
	} else {
		fmt.Println("Pilihan Tidak Tersedia")
	}
}
