package main

import (
	"fmt"
)

/*
	Struktur Utama Program

*/

// Maksimum jumlah mahasiswa dan mata kuliah
const maxMahasiswa = 100
const maxMatkul = 10

// Array untuk menyimpan data mahasiswa
var daftarMahasiswa [maxMahasiswa]Mahasiswa
var jumlahMahasiswa int

// Struct untuk menyimpan Data Setiap Mahasiswa
type Mahasiswa struct {
	Nama  string
	NIM   string
	Prodi string
	MataKuliah [maxMatkul]MataKuliah
	JumlahMK int
}

// Struct untuk menyimpan nilai setiap mahasiswa
type MataKuliah struct {
	Nim string
	NamaMK string
	UTS   float64
	UAS   float64
	Quiz  float64
	Total float64
	Grade string
}

/*
	Mahasiswa

 */

// Fungsi untuk menambahkan Mahasiswa
func tambahMahasiswa(nama, nim, prodi string) {
	if jumlahMahasiswa >= maxMahasiswa {
		fmt.Println("Kapasitas maksimum mahasiswa tercapai!")
		return
	}
	mahasiswa := Mahasiswa{
		Nama:  nama,
		NIM:   nim,
		Prodi: prodi,
	}
	daftarMahasiswa[jumlahMahasiswa] = mahasiswa
	jumlahMahasiswa++
	fmt.Println("Mahasiswa berhasil ditambahkan!")
}

func editMahasiswa(nim string, namaBaru, prodiBaru string) {
	for i := 0; i < jumlahMahasiswa; i++ {
		if daftarMahasiswa[i].NIM == nim {
			
			daftarMahasiswa[i].Nama = namaBaru
			daftarMahasiswa[i].Prodi = prodiBaru
			fmt.Println("Data Mahasiswa dengan NIM", nim, "berhasil diperbarui!")
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan!")
}

func hapusMahasiswa(nim string) {
	for i := 0; i < jumlahMahasiswa; i++ {
		if daftarMahasiswa[i].NIM == nim {
			
			for j := i; j < jumlahMahasiswa-1; j++ {
				daftarMahasiswa[j] = daftarMahasiswa[j+1]
			}
			
			jumlahMahasiswa--
			fmt.Println("Mahasiswa dengan NIM", nim, "berhasil dihapus!")
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan!")
}

// Fungsi untuk menghitung total nilai dan grade
func hitungTotalDanGrade(uts, uas, quiz float64)(float64, string) {
	total := uts*0.3 + uas*0.4 + quiz*0.3
	var grade string
	switch {
	case total >= 85:
		grade = "A"
	case total >= 70:
		grade = "B"
	case total >= 60:
		grade = "C"
	case total >= 50:
		grade = "D"
	default:
		grade = "E"
	}
	return total, grade
}

/*

	Mata Kuliah

 */

// Fungsi untuk Menambahkan Mata Kuliah ke Mahasiswa
func tambahMataKuliah(nim, namaMK string, uts, uas, quiz float64) {
	for i := 0; i < jumlahMahasiswa; i++ {
		if daftarMahasiswa[i].NIM == nim {
			if daftarMahasiswa[i].JumlahMK >= maxMatkul {
				fmt.Println("Mahasiswa dengan NIM", nim, "telah mencapai jumlah maksimum mata kuliah!")
				return
			}
			total, grade := hitungTotalDanGrade(uts, uas, quiz)
			matkul := MataKuliah{
				NamaMK: namaMK,
				UTS:    uts,
				UAS:    uas,
				Quiz:   quiz,
				Total:  total,
				Grade:  grade,
			}
			daftarMahasiswa[i].MataKuliah[daftarMahasiswa[i].JumlahMK] = matkul
			daftarMahasiswa[i].JumlahMK++
			fmt.Println("Mata Kuliah berhasil ditambahkan!")
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan!")
}

func editMataKuliah(nim, namaMK string, utsBaru, uasBaru, quizBaru float64) {
    for i, mhs := range daftarMahasiswa {
        if mhs.NIM == nim {
            for j, mk := range mhs.MataKuliah {
                if mk.NamaMK == namaMK {
                    total, grade := hitungTotalDanGrade(utsBaru, uasBaru, quizBaru)
                    daftarMahasiswa[i].MataKuliah[j].UTS = utsBaru
                    daftarMahasiswa[i].MataKuliah[j].UAS = uasBaru
                    daftarMahasiswa[i].MataKuliah[j].Quiz = quizBaru
                    daftarMahasiswa[i].MataKuliah[j].Total = total
                    daftarMahasiswa[i].MataKuliah[j].Grade = grade
                    fmt.Println("Data Mata Kuliah", namaMK, "berhasil diperbarui untuk Mahasiswa dengan NIM", nim)
                    return
                }
            }
            fmt.Println("Mata Kuliah", namaMK, "tidak ditemukan untuk Mahasiswa dengan NIM", nim)
            return
        }
    }
    fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan!")
}

func hapusMataKuliah(nim, namaMK string) {
	for i := 0; i < jumlahMahasiswa; i++ {
		if daftarMahasiswa[i].NIM == nim {
			for j := 0; j < daftarMahasiswa[i].JumlahMK; j++ {
				if daftarMahasiswa[i].MataKuliah[j].NamaMK == namaMK {
					
					for k := j; k < daftarMahasiswa[i].JumlahMK-1; k++ {
						daftarMahasiswa[i].MataKuliah[k] = daftarMahasiswa[i].MataKuliah[k+1]
					}
					
					daftarMahasiswa[i].JumlahMK--
					fmt.Println("Mata Kuliah", namaMK, "berhasil dihapus dari Mahasiswa dengan NIM", nim)
					return
				}
			}
			fmt.Println("Mata Kuliah", namaMK, "tidak ditemukan untuk Mahasiswa dengan NIM", nim)
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan!")
}

// Fungsi Menampilkan Semua Mahasiswa
func tampilkanMahasiswa() {
	if jumlahMahasiswa == 0 {
		fmt.Println("Data Mahasiswa Kosong.")
		return
	}

	for i := 0; i < jumlahMahasiswa; i++ {
		mhs := daftarMahasiswa[i]
		fmt.Printf("Nama: %s, NIM: %s, Program Studi: %s\n", mhs.Nama, mhs.NIM, mhs.Prodi)
		for j := 0; j < mhs.JumlahMK; j++ {
			mk := mhs.MataKuliah[j]
			fmt.Printf("\tMata Kuliah: %s, UTS: %.2f, UAS: %.2f, Quiz: %.2f, Total: %.2f, Grade: %s\n",
				mk.NamaMK, mk.UTS, mk.UAS, mk.Quiz, mk.Total, mk.Grade)
		}
	}
}

func main() {

	fmt.Print("LOG1: \n")
	tambahMahasiswa("Dafa", "12345", "Teknik Informatika")
	tambahMataKuliah("12345", "Algoritma", 80.5, 86.4,90)
	tampilkanMahasiswa()

	fmt.Print("\n")
	fmt.Print("LOG2: \n")

	tambahMahasiswa("Kevin", "22345", "Teknik Industri")
	tambahMataKuliah("22345", "Kalkulus", 65, 75, 90)
	tampilkanMahasiswa()
	
	fmt.Print("\n")
	fmt.Print("LOG3: \n")
	
	hapusMahasiswa("12345")
	tampilkanMahasiswa()
	
	fmt.Print("\n")
	fmt.Print("LOG4: \n")

	tambahMahasiswa("Dafa", "12345", "Teknik Informatika")
	tambahMataKuliah("12345", "Algoritma", 80.5, 86.4,90)
	editMataKuliah("22345", "Algoritma", 80,80,70)
	hapusMataKuliah("12345", "Algoritma")
	tampilkanMahasiswa()


}

