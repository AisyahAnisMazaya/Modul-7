*NOMOR 1*
Deskripsi Program
Program Go ini dirancang untuk menentukan posisi suatu titik relatif terhadap dua lingkaran. Program ini akan meminta pengguna untuk memasukkan koordinat pusat dan jari-jari dari dua lingkaran, serta koordinat sebuah titik. Setelah itu, program akan menganalisis posisi titik tersebut dan memberikan output berupa pesan yang menginformasikan apakah titik tersebut berada di dalam satu atau kedua lingkaran, atau di luar keduanya.
Penjelasan Kode
Deklarasi Struktur Data
titik: Struktur ini merepresentasikan sebuah titik dalam koordinat kartesian 2D dengan atribut x dan y yang bertipe integer.
lingkaran: Struktur ini merepresentasikan sebuah lingkaran dengan atribut tengah yang merupakan sebuah struktur titik yang menandai pusat lingkaran, dan radius yang merupakan jari-jari lingkaran.
Fungsi Utama (main)
Deklarasi Variabel: Mendeklarasikan variabel c1, c2 (bertipe lingkaran) untuk menyimpan data dua lingkaran, dan variabel p (bertipe titik) untuk menyimpan data titik yang akan diuji.
Input Data: Meminta pengguna untuk memasukkan nilai-nilai x, y, dan radius untuk kedua lingkaran, serta koordinat titik p.
Pemeriksaan:
dalamLingkaran: Memanggil fungsi dalamLingkaran untuk memeriksa apakah titik p berada di dalam kedua lingkaran.
dL12: Memanggil fungsi dL12 untuk memeriksa jarak antara titik p dan pusat masing-masing lingkaran.
Output: Mencetak hasil pemeriksaan ke layar.
Fungsi Pendukung
dalamLingkaran:
Fungsi ini menerima tiga argumen: dua lingkaran (c1, c2) dan satu titik (p).
Fungsi ini mengembalikan nilai boolean true jika titik p berada di dalam kedua lingkaran, dan false jika tidak.
Logika: Memanggil fungsi dL12 untuk kedua lingkaran dan mengembalikan true jika kedua pemanggilan mengembalikan true.
dL12:
Fungsi ini menerima dua argumen: satu lingkaran (c) dan satu titik (p).
Fungsi ini mengembalikan nilai boolean true jika titik p berada di dalam lingkaran c, dan false jika tidak.
Logika:
Menghitung jarak antara titik p dan pusat lingkaran c menggunakan rumus jarak Euclidean.
Membandingkan jarak tersebut dengan jari-jari lingkaran.
Mengembalikan true jika jarak kurang dari atau sama dengan jari-jari, yang berarti titik berada di dalam lingkaran.
Konsep Penting
Struktur Data: Penggunaan struktur titik dan lingkaran untuk merepresentasikan objek geometri.
Fungsi: Pemisahan tugas menjadi fungsi-fungsi yang lebih kecil untuk meningkatkan modularitas dan readability.
Percabangan: Penggunaan if-else untuk membuat keputusan berdasarkan hasil perhitungan.
Matematika: Penggunaan rumus jarak Euclidean untuk menghitung jarak antara dua titik.
Intinya, program ini menggunakan konsep geometri sederhana untuk menentukan posisi suatu titik relatif terhadap beberapa lingkaran.
Penjelasan Lebih Lanjut
Fungsi math.Pow digunakan untuk menghitung pangkat dari suatu bilangan.
Penggunaan float64 untuk perhitungan jarak memberikan akurasi yang lebih baik dibandingkan dengan tipe data integer.
Logika dalam fungsi dalamLingkaran dan dL12 mengimplementasikan definisi matematis dari titik yang berada di dalam lingkaran.

*NOMOR 2*
Deskripsi Program
Program Go ini dirancang untuk melakukan berbagai operasi pada sebuah array bilangan bulat. Program ini memungkinkan pengguna untuk:
Memasukkan data: Meminta pengguna untuk memasukkan jumlah elemen array dan nilai-nilai elemen tersebut.
Menampilkan data: Menampilkan seluruh isi array, elemen-elemen dengan indeks ganjil, genap, atau kelipatan suatu bilangan.
Manipulasi data: Menghapus elemen pada indeks tertentu.
Analisis data: Menghitung rata-rata, standar deviasi, dan frekuensi kemunculan suatu bilangan dalam array.
Penjelasan Kode
Deklarasi Variabel
N: Menyimpan jumlah elemen dalam array.
x: Menyimpan bilangan untuk mencari indeks kelipatannya.
arr: Array untuk menyimpan bilangan-bilangan yang dimasukkan pengguna.
bilangan: Variabel sementara untuk menyimpan bilangan yang sedang dibaca.
index: Menyimpan indeks elemen yang akan dihapus.
sum, mean, sd, freq: Variabel untuk perhitungan rata-rata, standar deviasi, dan frekuensi.
Fungsi Utama (main)
Input Data: Meminta pengguna untuk memasukkan jumlah elemen array dan nilai-nilai elemen tersebut.
Operasi Array: Melakukan berbagai operasi pada array sesuai dengan pilihan, seperti menampilkan elemen, menghapus elemen, menghitung rata-rata, standar deviasi, dan frekuensi.
Output: Mencetak hasil setiap operasi ke layar.
Fungsi Pendukung
sqrt: Fungsi ini digunakan untuk menghitung akar kuadrat dari sebuah bilangan menggunakan metode Newton-Raphson.
Konsep Penting
Array: Struktur data yang digunakan untuk menyimpan kumpulan data dengan tipe yang sama.
Loop: Digunakan untuk mengiterasi elemen-elemen dalam array.
Kondisi: Digunakan untuk membuat keputusan berdasarkan nilai tertentu (misalnya, indeks ganjil, genap, atau kelipatan).
Fungsi: Digunakan untuk membagi kode menjadi bagian-bagian yang lebih kecil dan dapat digunakan kembali.
Algoritma: Digunakan untuk menyelesaikan masalah, seperti menghitung rata-rata, standar deviasi, dan akar kuadrat.
Analisis Fitur
Fleksibilitas: Program ini sangat fleksibel karena memungkinkan pengguna untuk melakukan berbagai operasi pada array.
Efisiensi: Penggunaan loop dan kondisi memungkinkan program berjalan dengan efisien.
Modularitas: Fungsi sqrt dipisahkan untuk meningkatkan keterbacaan dan memungkinkan penggunaan ulang.
Kelengkapan: Program ini mencakup berbagai operasi yang umum dilakukan pada array.
