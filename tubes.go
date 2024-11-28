package main

import "fmt"
const NMAX int = 5555

type dataAdmin struct {
	fasilitas string
	wahana    string
}

type dataGuest struct {
	tempat             dataAdmin
	nama               string
	jarak, biaya, kode int
}
type arrData [NMAX - 1]dataGuest

type arrTempat [NMAX - 1]dataAdmin

func isiData(guest *arrData, tempat *arrTempat, jumlahGuest *int, k *int) { //Data Dummy yang telah Disediakan
	*k = 1
	guest[0].nama = "Tangkuban Perahu"
	guest[0].jarak = 30
	guest[0].biaya = 200000
	guest[0].kode = *k
	tempat[0].fasilitas = "Area parkir, toilet, warung makan, area pemandian air panas."
	tempat[0].wahana = "menikmati pemandangan kawah, trekking di sekitar kawah"
	*k++
	guest[1].nama = "Kawah Putih"
	guest[1].jarak = 50
	guest[1].biaya = 30000
	guest[1].kode = *k 
	tempat[1].fasilitas = "Area parkir, toilet, warung makan."
	tempat[1].wahana = "Menikmati pemandangan danau kawah, berjalan-jalan di sekitar danau."
	*k++
	guest[2].nama = "Braga"
	guest[2].jarak = 1
	guest[2].biaya = 5000
	guest[2].kode = *k
	tempat[2].fasilitas = "Restoran, kafe, toko-toko, area pejalan kaki."
	tempat[2].wahana = "Menikmati suasana kota, menjelajahi gedung-gedung bersejarah."
	*k++
	guest[3].nama = "Farmhouse"
	guest[3].jarak = 11
	guest[3].biaya = 25000
	guest[3].kode = *k
	tempat[3].fasilitas = "Taman, area parkir, toilet, restoran, kafe."
	tempat[3].wahana = "Berfoto di berbagai spot Instagramable, berinteraksi dengan hewan di peternakan mini."
	*k++
	guest[4].nama = "Floating Market"
	guest[4].jarak = 14
	guest[4].biaya = 20000
	guest[4].kode = *k
	tempat[4].fasilitas = "Area parkir, toilet, warung makan, toko suvenir."
	tempat[4].wahana = "Berbelanja di perahu terapung, menikmati makanan dan minuman lokal, menaiki perahu tradisional."
	*k++
	guest[5].nama = "Trans Studio Bandung"
	guest[5].jarak = 2
	guest[5].biaya = 250000
	guest[5].kode = *k
	tempat[5].fasilitas = "Berbagai wahana, toilet, restoran, kafe, area bermain anak."
	tempat[5].wahana = "Roller coaster, wahana air, simulator, taman bermain indoor."
	*k++
	guest[6].nama = "Dusun Bambu"
	guest[6].jarak = 14
	guest[6].biaya = 25000
	guest[6].kode = *k
	tempat[6].fasilitas = "Restoran, kafe, area parkir, taman."
	tempat[6].wahana = "Berjalan-jalan di sekitar danau, bermain di taman bermain anak, berfoto di spot foto yang indah."
	*k++
	guest[7].nama = "Taman Hutan Raya"
	guest[7].jarak = 10
	guest[7].biaya = 0
	guest[7].kode = *k
	tempat[7].fasilitas = "Area parkir, toilet, area piknik, jalur hiking."
	tempat[7].wahana = "Menjelajahi hutan, trekking, menikmati pemandangan air terjun Curug Dago."
	*k++
	guest[8].nama = "Saung Angklung"
	guest[8].jarak = 6
	guest[8].biaya = 100000
	guest[8].kode = *k
	tempat[8].fasilitas = "Teater, ruang pertunjukan, toilet, toko suvenir."
	tempat[8].wahana = "Menyaksikan pertunjukan musik angklung, belajar memainkan angklung."
	*k++
	guest[9].nama = "Gedung Sate"
	guest[9].jarak = 0
	guest[9].biaya = 10000
	guest[9].kode = *k
	tempat[9].fasilitas = "Area parkir, area pejalan kaki."
	tempat[9].wahana = "Menikmati arsitektur bangunan, mengambil foto"
	*k++
	*jumlahGuest = 10
	
}

func headerLogin() {
	fmt.Println(":===================================================:")
	fmt.Println(":               APLIKASI PARIWISATA 				 :")
	fmt.Println(":                   Kelompok 17                     :")
	fmt.Println(":===================================================:")
	fmt.Println("Admin / Guest ?")
	fmt.Println("1. Admin")
	fmt.Println("2. Guest")
	fmt.Println("3. Keluar (-1)")
}

func main(){
	var guest arrData
	var tempat arrTempat
	var kode int
	var jumlahGuest int
	var status1, status2, cari int

	isiData(&guest, &tempat, &jumlahGuest,&kode)
	headerLogin()
	fmt.Print("Select : ")
	fmt.Scan(&status1)
	if status1 == 1  || status1 == 2{ // MEMILIH ANTARA ADMIN / GUEST
		if status1 == 1 { // jika memilih admin
			fmt.Println(":---------- DASHBOARD ADMIN ----------:")
			fmt.Println(": Kamu Adalah Admin:")
			fmt.Println("1. Tampilkan Data")
			fmt.Println("2. Tambah Data")
			fmt.Println("3. Ubah Data")
			fmt.Println("4. Hapus Data")
			fmt.Println("5. Kembali")
			fmt.Print("Pilihan : ")
			fmt.Scan(&status2)
				if status2 == 1 {
					tampilkanData(&guest,tempat,jumlahGuest)
				} else if status2 == 2 {
					tambahData(&guest, &tempat, &jumlahGuest,&kode)
				} else if status2 == 3 {
					ubahData(&guest, tempat)
				} else if status2 == 4 {
					hapusData(&guest)
				} else {
					main()
				}

		} else if status1 == 2 { // jika memilih guest
			fmt.Println(":---------- DASHBOARD GUEST ----------:")
			fmt.Println("Mau Kemana Hari ini ? Kami akan memberikan beberapa rekomendasi")
			fmt.Println()
			for i := 0; i < jumlahGuest; i++ {
			fmt.Println("Nama : ", guest[i].nama, "Jarak : ", guest[i].jarak, "KM", "Biaya : Rp.", guest[i].biaya)
			}
			fmt.Println("SELECT ONE")
			fmt.Println("1. Terdekat")
			fmt.Println("2. Jarak Tertentu")
			fmt.Println("3. Termurah")
			fmt.Println("4. Harga Tertentu ")
			fmt.Print("Select : ")
			fmt.Scan(&status2)
			if status2 != 5 {
				if status2 == 1 {
					terdekatSemua(&guest, tempat, jumlahGuest)
				} else if status2 == 2 {
					fmt.Print("Masukkan Jarak :")
					fmt.Scan(&cari)
					jarakTertentu(&guest, jumlahGuest, cari)
				} else if status2 == 3 {
					termurahSemua(&guest,tempat,jumlahGuest)
				} else if status2 == 4 {
					fmt.Print("Masukkan Biaya :")
					fmt.Scan(&cari)
					biayaTertentu(&guest,jumlahGuest,cari)
				} else {
					main()
				}
			}	
		} 
	}
}

func tambahData(guest *arrData, tempat *arrTempat, jumlahGuest *int, kode *int) {//Menambah Data
	var nama, wahana, fasilitas, kembali string
	var jarak, biaya int

	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	fmt.Print("Jarak : ")
	fmt.Scan(&jarak)
	fmt.Print("Biaya : ")
	fmt.Scan(&biaya)
	fmt.Print("Fasilitas : ")
	fmt.Scan(&fasilitas)
	fmt.Print("Wahana : ")
	fmt.Scan(&wahana)
	guest[*jumlahGuest].nama = nama
	guest[*jumlahGuest].jarak = jarak
	guest[*jumlahGuest].biaya = biaya
	guest[*jumlahGuest].kode = *kode
	tempat[*jumlahGuest].fasilitas = fasilitas
	tempat[*jumlahGuest].wahana = wahana
	*jumlahGuest++	
	*kode++
	

	for i := 0; i < *jumlahGuest; i++ {
		fmt.Println("Kode : ",guest[i].kode, "Nama : ",guest[i].nama, "Jarak : ", guest[i].jarak, "Biaya : ", guest[i].biaya) // Menampilkan data setelah diinput
		fmt.Println("Fasilitas : ",tempat[i].fasilitas)
		fmt.Println("Wahana", tempat[i].wahana)
	}
	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}

func tampilkanData(A *arrData, tempat arrTempat, jumlahGuest int) { // MENAMPILKAN DATA
	var kembali string

	fmt.Println("     Menampilkan Data     ")
	for i := 0; i < jumlahGuest; i++ {
		fmt.Println("Nama : ", A[i].nama, " | Jarak : ", A[i].jarak, "KM", " | Biaya : Rp.", A[i].biaya)
		fmt.Println("	Fasilitas : ",tempat[i].fasilitas)
		fmt.Println("	Wahana", tempat[i].wahana)
		fmt.Println()
	}
	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}

func ubahData(guest *arrData, tempat arrTempat) { // Mengubah Data
	var jumlahGuest int
	var kode int
	var nama string
	var opsi string
	var jarak, biaya int
	var wahana, fasilitas, kembali string

	isiData(guest, &tempat, &jumlahGuest, &kode)
	for i := 0; i < jumlahGuest; i++ {
		fmt.Println("Kode :", guest[i].kode, "Nama : ", guest[i].nama, "Jarak : ", guest[i].jarak, "KM", "Biaya : Rp.", guest[i].biaya)
	}
	fmt.Print("Masukan Kode :")
	fmt.Scan(&kode)
	for i := 0; i < jumlahGuest; i++ {
		if guest[i].kode == kode {
			fmt.Print("Masukan Opsi :")
			fmt.Scan(&opsi)
			if opsi == "Nama" {
				fmt.Print("Masukan Perubahannya :")
				fmt.Scan(&nama)
				guest[i].nama = nama
			} else if opsi == "Jarak" {
				fmt.Print("Masukan Perubahannya :")
				fmt.Scan(&jarak)
				guest[i].jarak = jarak
			} else if opsi == "Biaya" {
				fmt.Print("Masukan Perubahannya :")
				fmt.Scan(&biaya)
				guest[i].biaya = biaya
			} else if opsi == "Fasilitas" {
				fmt.Print("Masukkan Perubahannya :")
				fmt.Scan(&fasilitas)
				tempat[i].fasilitas = fasilitas
			} else if opsi == "Wahana" {
				fmt.Print("Masukkan Perubahannya :")
				fmt.Scan(&wahana)
				tempat[i].wahana = wahana
			}
		}
	}
	for i := 0; i < jumlahGuest; i++ {
		fmt.Println("Kode :", guest[i].kode, "Nama : ", guest[i].nama, "Jarak : ", guest[i].jarak, "KM", "Biaya : Rp.", guest[i].biaya)
	}
	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}

func hapusData(guest *arrData) { // Menghapus Data
	var tempat arrTempat
	var jumlahGuest int
	var kode int
	var kembali string

	isiData(guest, &tempat, &jumlahGuest,&kode)
	for i := 0; i < jumlahGuest; i++ {
		fmt.Println("Kode :", guest[i].kode, "Nama : ", guest[i].nama, "Jarak : ", guest[i].jarak, "KM", "Biaya : Rp.", guest[i].biaya)
	}
	fmt.Print("Masukan Kode :")
	fmt.Scan(&kode)
	for i := kode - 1; i < jumlahGuest; i++ {
		guest[i].nama = guest[i+1].nama
		guest[i].jarak = guest[i+1].jarak
		guest[i].biaya = guest[i+1].biaya
	}
	guest[jumlahGuest-1].kode = 0
	guest[jumlahGuest-1].nama = ""
	guest[jumlahGuest-1].jarak = 0
	guest[jumlahGuest-1].biaya = 0
	jumlahGuest--
	for i := 0; i < jumlahGuest; i++ {
		fmt.Println("Kode :", guest[i].kode, "Nama : ", guest[i].nama, "Jarak : ", guest[i].jarak, "KM", "Biaya : Rp.", guest[i].biaya)
	}
	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}

// Admin bisa undo ke header
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func terdekatSemua(guest *arrData, tempat arrTempat, jumlahGuest int) { //Menampilkan Tempat yang Terdekat 
	var kembali string

	// menggunakan Selection Sort
	for i := 0; i < jumlahGuest; i++ { 
		var minIdx = i
		for j := i; j < jumlahGuest; j++ {
			if guest[j].jarak < guest[minIdx].jarak {
				minIdx = j
			}
		}
		guest[i].biaya, guest[minIdx].biaya = guest[minIdx].biaya, guest[i].biaya
		guest[i].jarak, guest[minIdx].jarak = guest[minIdx].jarak, guest[i].jarak
		guest[i].nama, guest[minIdx].nama = guest[minIdx].nama, guest[i].nama
	}
	
	fmt.Println("=== WISATA TERDEKAT ===") // menampilkan urutan semua data dari terdekat
	fmt.Println("Nama : ", guest[0].nama)
	fmt.Println("	Jarak : ", guest[0].jarak, "KM")
	fmt.Println("	Biaya : Rp.", guest[0].biaya)
	fmt.Println("	Fasilitas :", tempat[0].fasilitas)
	fmt.Println("	Wahana : ", tempat[0].wahana)
	fmt.Println("=======================")
	fmt.Println("Pilihan Lain : ")
	for i := 1; i < jumlahGuest; i++ {
		fmt.Println(guest[i].nama)
		fmt.Println("	Jarak : ", guest[i].jarak, "KM")
		fmt.Println("	Biaya : Rp.", guest[i].biaya)
		fmt.Println("	Fasilitas :", tempat[i].fasilitas)
		fmt.Println("	Wahana : ", tempat[i].wahana)
	}
	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}

func termurahSemua(guest *arrData, tempat arrTempat, jumlahGuest int) { //Menampilkan Tempat yang Termurah
	var kembali string
	var i int
	var j int

	// menggunakan Insertion Sort
	for i := 1; i < jumlahGuest; i++ {
		key := guest[i]
		j := i - 1
		for j >= 0 && guest[j].biaya > key.biaya {
			guest[j+1] = guest[j]
			j = j - 1
		}
		guest[j+1] = key
	}
	guest[i].biaya, guest[j].biaya = guest[j].biaya, guest[i].biaya
	guest[i].jarak, guest[j].jarak = guest[j].jarak, guest[i].jarak
	guest[i].nama, guest[j].nama = guest[j].nama, guest[i].nama
	

	fmt.Println("=== WISATA TERMURAH ===") // menampilkan urutan semua data dari terdekat
	fmt.Println("Nama : ", guest[0].nama)
	fmt.Println("	Jarak : ", guest[0].jarak, "KM | Biaya : Rp.", guest[0].biaya)
	fmt.Println("	Fasilitas :", tempat[0].fasilitas)
	fmt.Println("	Wahana : ", tempat[0].wahana)
	fmt.Println("=======================")
	fmt.Println("Pilihan Lain : ")
	for i := 1; i < jumlahGuest; i++ {
		fmt.Println(guest[i].nama)
		fmt.Println("	Jarak : ", guest[i].jarak, "KM | Biaya : Rp.", guest[i].biaya)
		fmt.Println("	Fasilitas :", tempat[i].fasilitas)
		fmt.Println("	Wahana : ", tempat[i].wahana)
	}
	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}

func biayaTertentu(guest *arrData, jumlahGuest int, cari int) { //Menampilkan tempat dengan Biaya yang ditentukan oleh user
	var kembali string
	
	// Menggunakan Sequential Search
	fmt.Print()
	for i := 0; i < jumlahGuest; i++ {
		if guest[i].biaya == cari {
			fmt.Println("Wisata dengan Biaya", cari, "adalah")
			fmt.Println("Nama :", guest[i].nama, " | Jarak : ", guest[i].jarak, "KM", " | Biaya : Rp.", guest[i].biaya)
		}
	}
	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}

func jarakTertentu(guest *arrData, jumlahGuest int, cari int) {//Menampilkan tempat dengan Jarak yang ditentukan oleh user
	var left, right, mid int
	var pass, i int
	var idx int
	var kembali string
	var temp dataGuest

	// Menggunakan BinarySearch
	pass = 1
	for pass <= jumlahGuest-1 {
		i = pass
		temp = guest[pass]
		for i > 0 && temp.jarak < guest[i-1].jarak {
			guest[i] = guest[i-1]
			i = i - 1
		}
		guest[i] = temp
		pass = pass + 1
	}
	left = 0
	right = jumlahGuest - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if guest[mid].jarak > cari {
			right = mid - 1
		} else if guest[mid].jarak < cari {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	fmt.Println("Wisata dengan Jarak", cari, "adalah")
	fmt.Println("Nama :", guest[idx].nama, " | Jarak : ", guest[idx].jarak, " | Biaya : ", guest[idx].biaya)

	fmt.Println(":----------------------------------------------------------:")
	fmt.Println("9. Kembali")
	fmt.Print("Select : ")
	fmt.Scan(&kembali)
	if kembali == "9" {
		fmt.Println(":----------------------------------------------------------:")
		main()
	}
}


