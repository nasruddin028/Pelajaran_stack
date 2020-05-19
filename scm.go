
package main

import (
	"fmt"
)

type Stack struct {
    slice []int
}

//untuk menambahkan data di atas
func (s *Stack) Push(i int) {
   s.slice = append(s.slice, i) 
} 

//untuk kembali ke atas
func (s *Stack) Peek() int {
 return s.slice[len(s.slice)-1]
}

//Untuk menghapus dan kembali keatas
func (s *Stack) Pop() int {
 var ret int = s.Peek()
 s.slice = s.slice[0:len(s.slice)-1]	
 return ret
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func main () {
 var s *Stack = new(Stack)
 var jumlah,index,mobil,total int
 var pilih string

 fmt.Println()
 fmt.Println()
 fmt.Println("******************************************************")
 fmt.Println("SELAMAT MENGGUNAKAN APLIKASI MONITORING INPUTAN GUDANG")
 fmt.Println("======================================================")
 fmt.Println("                Silahkan Tambahkan Data               ")
 fmt.Println("======================================================")
 fmt.Println()
 fmt.Print("Berapa jumlah Mobil yang datang ?")
 fmt.Scanln(&mobil)
index = 1
 for index != mobil + 1 {
	fmt.Print("Masukkan Jumlah Material Mobil ke ")
	fmt.Print(index)
	fmt.Print(" : ")
	fmt.Scanln(&jumlah)
	s.Push(jumlah)
	

// setelah tersimpan diseleksi apakah data sudah sesuai atau akan dihapus
	fmt.Print("Apakah data ingin dihapus Y/N :")
	fmt.Scanln(&pilih)

	fmt.Println()
	fmt.Println()
	if pilih == "y" {
		s.Pop()
		fmt.Println("Data sudah terhapus")

		fmt.Print("Masukkan Jumlah Material Mobil ke ")
		fmt.Print(index)
		fmt.Print(" : ")
		fmt.Scanln(&jumlah)
		s.Push(jumlah)
	} 
	index = index + 1
    total = total + jumlah	
}


 fmt.Println("Terimakasih telah menginput data")
 fmt.Println("Jumlah Material yang masuk :")
 index = index - 1
 fmt.Println()
 fmt.Println("------------------------------")
 for index >= 1 {
	fmt.Print("| Mobil ke : ")
	fmt.Print(index)
	fmt.Print("       | ") 
	fmt.Print(s.Pop())
	fmt.Println(" Box |")
	fmt.Println("------------------------------")
 index = index - 1
 }
fmt.Print("| Total Barang Masuk | ")
fmt.Print(total)
fmt.Println(" Box |")
fmt.Println("------------------------------")
fmt.Println()


}