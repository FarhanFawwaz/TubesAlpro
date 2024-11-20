package main

import (
	"fmt"
)

type historiTransaksi struct {
	lembagaSembako  tbarang
	hargaTotal      int
	jumlahPembelian int
	waktuTransaksi  waktu
	jenis           string
}

type tbarang struct {
	jenis      string
	harga      int
	kodeBarang string
	kwalitas   string
	stok       int
}

type waktu struct {
	tanggal int
	bulan   int
	tahun   int
}

type barang struct {
	sBarang [1000]tbarang
	nBarang int
}

type transaksiRecord struct {
	hTransaksi [4000]historiTransaksi
	nTransaksi int
}

func main() {
	menu()
	// func main hanya akan berisi procedure menu saja
}

func menu() {
	var pilih int
	var benda barang
	var transaksi transaksiRecord

	// var pelanggan customerData
	pilih = 1

	for pilih != 8 {
		header()
		fmt.Println("Pilihan Anda: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			menuInputData(&benda)
		} else if pilih == 2 {
			menuUbahData(&benda)
		} else if pilih == 3 {
			menuHapusData(&benda)
		} else if pilih == 4 {
			catatTransaksi(&benda, &transaksi)
			tampilkanTransaksi(&transaksi)
			tampilkanData(benda)
		} else if pilih == 5 {
			menuCariBarang(&benda)
		} else if pilih == 6 {
			menuUrutkanData(&benda)
		} else if pilih == 7 {
			menuCariEkstrem(&benda)
		} else if pilih == 8 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
		} else {
			fmt.Println("Pilihan Anda Tidak Sesuai")
		}

	}
}

func header() {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Aplikasi Inventori Barang			")
	fmt.Println("=========================================================")
	fmt.Println("		     Kelompok 14			        ")
	fmt.Println("	  Farhan Muamar Fawwaz 103032300076		")
	fmt.Println("	  David Nathan Honggo Kusumo 1030323000197		")
	fmt.Println("	 	      IT-47-01	 		 	")
	fmt.Println("=======================================================")
	fmt.Println(" 			MENU			   	")
	fmt.Println("=======================================================")
	fmt.Println("Silahkan Pilih Pilihan Diantara Berikut  	  	")
	fmt.Println("Pilihan :						")
	fmt.Println("1. Input Data")
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Transaksi Masuk/Keluar")
	fmt.Println("5. Cari Data Barang")
	fmt.Println("6. Urutkan Data Barang")
	fmt.Println("7. Menu Cari Nilai Ekstrem")
	fmt.Println("8. Selesai")
	fmt.Println("=========================================================")
}

func menuInputData(benda *barang) {
	var pilih int
	var cek bool
	cek = true
	for cek {
		header1()
		fmt.Println("Pilihan Anda: ")
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 {
			switch pilih {
			case 1:
				inputBarang(benda)
			case 2:
				fmt.Println("Kembali ke Menu")
				cek = false
			}
		} else {
			fmt.Println("Pilihan Anda Tidak Sesuai")
		}
	}
}
func seqSearch(benda barang, kode string) int {
	ketemu := -1
	k := 0
	for ketemu == -1 && k < benda.nBarang {
		if benda.sBarang[k].kodeBarang == kode {
			ketemu = k
		}
		k++
	}
	return ketemu
}
func inputBarang(benda *barang) {
	var jenis, kodeBarang string
	fmt.Print("Jenis Barang: ")
	fmt.Scan(&jenis)
	fmt.Print("Kode Barang: ")
	fmt.Scan(&kodeBarang)
	for kodeBarang != "XXX" {
		if seqSearch(*benda, kodeBarang) == -1 {
			benda.sBarang[benda.nBarang].jenis = jenis
			benda.sBarang[benda.nBarang].kodeBarang = kodeBarang
			fmt.Print("Harga: ")
			fmt.Scan(&benda.sBarang[benda.nBarang].harga)
			fmt.Print("Kwalitas: ")
			fmt.Scan(&benda.sBarang[benda.nBarang].kwalitas)
			fmt.Print("Stok: ")
			fmt.Scan(&benda.sBarang[benda.nBarang].stok)
			benda.nBarang++
			fmt.Println("Data barang berhasil ditambahkan.")
		} else {
			fmt.Println("Barang dengan kode tersebut sudah ada.")
		}
		fmt.Print("Jenis Barang: ")
		fmt.Scan(&jenis)
		fmt.Print("Kode Barang: ")
		fmt.Scan(&kodeBarang)
	}
	tampilkanData1(benda)
}
func tampilkanData1(benda *barang) {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Data Barang			")
	fmt.Println("=========================================================")
	for i := 0; i < benda.nBarang; i++ {
		fmt.Printf("Kode Barang: %s, Jenis: %s, Harga: %d, Kwalitas: %s, Stok: %d\n",
			benda.sBarang[i].kodeBarang, benda.sBarang[i].jenis, benda.sBarang[i].harga, benda.sBarang[i].kwalitas, benda.sBarang[i].stok)
	}
}

func menuUbahData(benda *barang) {
	var pilih int
	var cek bool
	cek = true
	for cek {
		header2()
		fmt.Println("Pilihan Anda: ")
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 {
			switch pilih {
			case 1:
				ubahData(benda)
			case 2:
				fmt.Println("Kembali ke Menu")
				cek = false
			}
		} else {
			fmt.Println("Pilihan Anda Tidak Sesuai")
		}
	}
}

func ubahData(benda *barang) {
	var index int
	var kodeBarang string
	fmt.Print("Masukkan kode barang yang ingin diubah: ")
	fmt.Scan(&kodeBarang)

	index = seqSearch(*benda, kodeBarang)
	if index != -1 {
		fmt.Println("Data lama barang:")
		fmt.Printf("Jenis: %s,Harga: %d ,Kwalitas: %s, Stok: %d\n", benda.sBarang[index].jenis, benda.sBarang[index].harga, benda.sBarang[index].kwalitas, benda.sBarang[index].stok)

		fmt.Print("Jenis Barang baru: ")
		fmt.Scan(&benda.sBarang[index].jenis)
		fmt.Print("Harga Barang baru: ")
		fmt.Scan(&benda.sBarang[index].harga)
		fmt.Print("Kwalitas baru: ")
		fmt.Scan(&benda.sBarang[index].kwalitas)
		fmt.Print("Stok baru: ")
		fmt.Scan(&benda.sBarang[index].stok)

		fmt.Println("Data barang berhasil diubah.")
	} else {
		fmt.Println("Barang dengan kode tersebut tidak ditemukan.")
	}
	tampilkanData1(benda)
}
func menuHapusData(benda *barang) {
	var pilih int
	var cek bool
	cek = true
	for cek {
		header4()
		fmt.Println("Pilihan Anda: ")
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 {
			switch pilih {
			case 1:
				hapusData(benda)
			case 2:
				fmt.Println("Kembali ke Menu")
				cek = false
			}
		} else {
			fmt.Println("Pilihan Anda Tidak Sesuai")
		}
	}
}
func hapusData(benda *barang) {
	var index int
	var kodeBarang string
	fmt.Print("Masukkan kode barang yang ingin dihapus: ")
	fmt.Scan(&kodeBarang)

	index = binarySearch(benda, kodeBarang)
	if index != -1 {
		for i := index; i < benda.nBarang-1; i++ {
			benda.sBarang[i] = benda.sBarang[i+1]
		}
		benda.nBarang--
		fmt.Println("Data barang berhasil dihapus.")
	} else {
		fmt.Println("Barang dengan kode tersebut tidak ditemukan.")
	}
	tampilkanData1(benda)
}

func binarySearch(benda *barang, kode string) int {

	left := 0
	right := benda.nBarang - 1
	found := -1
	for left <= right && found == -1 {
		mid := left + (right-left)/2
		if benda.sBarang[mid].kodeBarang == kode {
			found = mid
		} else if benda.sBarang[mid].kodeBarang < kode {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}

func header1() {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Input Data Barang			")
	fmt.Println("=========================================================")
	fmt.Println("1. Tambah Barang")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Println("=========================================================")
}
func header2() {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Edit Data Barang			")
	fmt.Println("=========================================================")
	fmt.Println("1. Edit Barang")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Println("=========================================================")
}
func header4() {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Hapus Data Barang			")
	fmt.Println("=========================================================")
	fmt.Println("1. Hapus Barang")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Println("=========================================================")
}
func catatTransaksi(benda *barang, transaksi *transaksiRecord) {
	var jumlah, pilihan int
	var tanggal, bulan, tahun int
	var time waktu
	var kodeBarang string
	fmt.Print("Masukkan kode barang: ")
	fmt.Scan(&kodeBarang)
	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&jumlah)
	fmt.Println("=========================================================")
	fmt.Println("Pilih tipe transaksi: ")
	fmt.Println("1. Masuk")
	fmt.Println("2. Keluar")
	fmt.Println("=========================================================")
	fmt.Scan(&pilihan)

	var tipe string
	if pilihan == 1 || pilihan == 2 {
		if pilihan == 1 {
			tipe = "Masuk"
		} else if pilihan == 2 {
			tipe = "Keluar"
		}
		index := seqSearch(*benda, kodeBarang)
		if index == -1 {
			fmt.Println("Barang dengan kode tersebut tidak ditemukan.")
		} else {
			fmt.Print("Masukkan tanggal (dd): ")
			fmt.Scan(&tanggal)
			fmt.Print("Masukkan bulan (mm): ")
			fmt.Scan(&bulan)
			fmt.Print("Masukkan tahun (yyyy): ")
			fmt.Scan(&tahun)

			if tipe == "Masuk" {
				benda.sBarang[index].stok += jumlah
			} else if tipe == "Keluar" {
				if benda.sBarang[index].stok >= jumlah {
					benda.sBarang[index].stok -= jumlah
				} else {
					fmt.Println("Stok barang tidak mencukupi untuk transaksi keluar.")
					return
				}
			}

			transaksi.hTransaksi[transaksi.nTransaksi].jenis = tipe
			transaksi.hTransaksi[transaksi.nTransaksi].lembagaSembako = benda.sBarang[index]
			transaksi.hTransaksi[transaksi.nTransaksi].jumlahPembelian = jumlah
			time.tanggal = tanggal
			time.bulan = bulan
			time.tahun = tahun
			transaksi.hTransaksi[transaksi.nTransaksi].waktuTransaksi = time
			transaksi.hTransaksi[transaksi.nTransaksi].hargaTotal = jumlah * benda.sBarang[index].harga
			transaksi.nTransaksi++
			fmt.Println("Transaksi berhasil dicatat.")
		}
	} else {
		fmt.Print("Transaksi Tidak Valid")
	}
}

func tampilkanTransaksi(transaksi *transaksiRecord) {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Riwayat Transaksi			")
	for i := 0; i < transaksi.nTransaksi; i++ {
		t := transaksi.hTransaksi[i]
		fmt.Println("Jenis:", t.jenis)
		fmt.Println("Kode Barang:", t.lembagaSembako.kodeBarang)
		fmt.Println("Jumlah Pembelian:", t.jumlahPembelian)
		fmt.Printf("Tanggal: %02d-%02d-%04d\n", t.waktuTransaksi.tanggal, t.waktuTransaksi.bulan, t.waktuTransaksi.tahun)
	}
}
func cariBarangBerdasarkanJenis(benda *barang, jenis string) {
	var ditemukan bool
	fmt.Println("Hasil pencarian berdasarkan jenis barang:", jenis)
	ditemukan = false
	for i := 0; i < benda.nBarang; i++ {
		if benda.sBarang[i].jenis == jenis {
			fmt.Printf("Kode Barang: %s, Jenis: %s, Kwalitas: %s, Stok: %d\n",
				benda.sBarang[i].kodeBarang, benda.sBarang[i].jenis, benda.sBarang[i].kwalitas, benda.sBarang[i].stok)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Barang dengan jenis tersebut tidak ditemukan.")
	}
}

func cariBarangBerdasarkanKode(benda *barang, kode string) {
	fmt.Println("Hasil pencarian berdasarkan kode barang:", kode)
	index := seqSearch(*benda, kode)
	if index != -1 {
		fmt.Printf("Kode Barang: %s, Jenis: %s, Kwalitas: %s, Stok: %d\n",
			benda.sBarang[index].kodeBarang, benda.sBarang[index].jenis, benda.sBarang[index].kwalitas, benda.sBarang[index].stok)
	} else {
		fmt.Println("Barang dengan kode tersebut tidak ditemukan.")
	}
}
func menuCariBarang(benda *barang) {
	var pilih int
	for {
		header3()
		fmt.Println("Pilihan Anda: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			var jenis string
			fmt.Print("Masukkan jenis barang: ")
			fmt.Scan(&jenis)
			cariBarangBerdasarkanJenis(benda, jenis)
		} else if pilih == 2 {
			var kode string
			fmt.Print("Masukkan kode barang: ")
			fmt.Scan(&kode)
			cariBarangBerdasarkanKode(benda, kode)
		} else if pilih == 3 {
			fmt.Println("Kembali ke Menu Utama")
			return
		} else {
			fmt.Println("Pilihan Anda Tidak Sesuai")
		}
	}
}

func header3() {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Cari Barang			")
	fmt.Println("=========================================================")
	fmt.Println("1. Cari berdasarkan jenis barang")
	fmt.Println("2. Cari berdasarkan kode barang")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println("=========================================================")
}
func menuUrutkanData(benda *barang) {
	var pilih int
	fmt.Println("=========================================================")
	fmt.Println("	       Urutkan Data Barang")
	fmt.Println("=========================================================")
	fmt.Println("1. Urutkan Berdasarkan Stok")
	fmt.Println("2. Urutkan Berdasarkan Harga")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println("=========================================================")
	fmt.Print("Pilihan Anda: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		urutkanBerdasarkanStok(benda)
		tampilkanData1(benda)
	case 2:
		urutkanBerdasarkanHarga(benda)
		tampilkanData1(benda)
	case 3:
		return
	default:
		fmt.Println("Pilihan Anda Tidak Sesuai")
	}
}
func urutkanBerdasarkanStok(benda *barang) {
	// teurut secara descdending(menurun) menggunakan selection sort
	n := benda.nBarang
	for pass := 1; pass <= n; pass++ {
		idx := pass - 1
		i := pass
		for i < n {
			if benda.sBarang[idx].stok < benda.sBarang[i].stok {
				idx = i
			}
			i++
		}
		temp := benda.sBarang[pass-1]
		benda.sBarang[pass-1] = benda.sBarang[idx]
		benda.sBarang[idx] = temp
	}
	fmt.Println("Data barang berhasil diurutkan berdasarkan stok.")
}

func tampilkanData(benda barang) {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Data Barang			")
	fmt.Println("=========================================================")
	for i := 0; i < benda.nBarang; i++ {
		fmt.Printf("Kode Barang: %s, Jenis: %s, Harga: %d, Kwalitas: %s, Stok: %d\n",
			benda.sBarang[i].kodeBarang, benda.sBarang[i].jenis, benda.sBarang[i].harga, benda.sBarang[i].kwalitas, benda.sBarang[i].stok)
	}
}

func urutkanBerdasarkanHarga(benda *barang) {
	//terurut secara ascending
	var pass int
	var i int
	pass = 1
	for pass <= benda.nBarang-1 {
		i = pass
		temp := benda.sBarang[pass]
		for i > 0 && temp.harga < benda.sBarang[i-1].harga {
			benda.sBarang[i] = benda.sBarang[i-1]
			i = i - 1
		}
		benda.sBarang[i] = temp
		pass = pass + 1
	}
	fmt.Println("Data barang berhasil diurutkan berdasarkan harga.")
}

func menuCariEkstrem(benda *barang) {
	var pilih int
	var cek bool
	cek = true
	for cek {
		header6()
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilih)
		if pilih == 1 || pilih == 2 || pilih == 3 {
			switch pilih {
			case 1:
				findMaxStok(benda)
			case 2:
				findMinHarga(benda)
			case 3:
				cek = false
			}
		} else {
			fmt.Println("Pilihan Anda Tidak Sesuai")
		}
	}
}
func findMaxStok(benda *barang) {
	if benda.nBarang == 0 {
		fmt.Println("Tidak ada barang yang tersedia.")
		return
	}
	maxIndex := 0
	for i := 1; i < benda.nBarang; i++ {
		if benda.sBarang[i].stok > benda.sBarang[maxIndex].stok {
			maxIndex = i
		}
	}
	fmt.Printf("Barang dengan stok terbanyak adalah Kode Barang: %s, Jenis: %s, Harga: %d, Kwalitas: %s, Stok: %d\n",
		benda.sBarang[maxIndex].kodeBarang, benda.sBarang[maxIndex].jenis, benda.sBarang[maxIndex].harga, benda.sBarang[maxIndex].kwalitas, benda.sBarang[maxIndex].stok)
}

func findMinHarga(benda *barang) {
	if benda.nBarang == 0 {
		fmt.Println("Tidak ada barang yang tersedia.")
		return
	}
	minIndex := 0
	for i := 1; i < benda.nBarang; i++ {
		if benda.sBarang[i].harga < benda.sBarang[minIndex].harga {
			minIndex = i
		}
	}
	fmt.Printf("Barang dengan harga termurah adalah Kode Barang: %s, Jenis: %s, Harga: %d, Kwalitas: %s, Stok: %d\n",
		benda.sBarang[minIndex].kodeBarang, benda.sBarang[minIndex].jenis, benda.sBarang[minIndex].harga, benda.sBarang[minIndex].kwalitas, benda.sBarang[minIndex].stok)
}
func header6() {
	fmt.Println("=========================================================")
	fmt.Println(" 	       Menu Cari Barang Ekstrem			")
	fmt.Println("=========================================================")
	fmt.Println("1. Cari Barang dengan Stok Terbanyak")
	fmt.Println("2. Cari Barang dengan Harga Termurah")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println("=========================================================")
}
