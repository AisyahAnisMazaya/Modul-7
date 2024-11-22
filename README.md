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

*NOMOR 3*
Deskripsi Program
Program Go ini dirancang untuk mencatat hasil pertandingan antara dua klub olahraga. Program ini akan meminta pengguna untuk memasukkan nama kedua klub dan kemudian secara berulang meminta skor untuk setiap pertandingan. Program akan terus berjalan hingga pengguna memasukkan skor negatif sebagai tanda akhir input. Setelah semua data pertandingan dimasukkan, program akan menampilkan hasil setiap pertandingan dan klub mana yang memenangkan pertandingan tersebut.
Penjelasan Kode
Deklarasi Variabel
club1, club2: String untuk menyimpan nama kedua klub.
skor1, skor2: Integer untuk menyimpan skor masing-masing klub dalam setiap pertandingan.
clubMenang: Slice of string untuk menyimpan nama klub yang memenangkan setiap pertandingan.
hasil: Slice of string untuk menyimpan hasil setiap pertandingan (menang, kalah, atau seri).
Fungsi Utama (main)
Input Nama Klub: Meminta pengguna untuk memasukkan nama kedua klub yang akan bertanding.
Loop Pertandingan:
Meminta pengguna memasukkan skor untuk setiap pertandingan.
Kondisi Berhenti: Jika salah satu skor negatif, loop akan berhenti, menandakan akhir input.
Menentukan Pemenang: Membandingkan skor dan menambahkan nama klub yang menang ke slice clubMenang dan hasil pertandingan ke slice hasil. Jika seri, maka kata "Draw" ditambahkan ke slice hasil.
Output Hasil: Mencetak hasil setiap pertandingan ke layar.
Konsep Penting
Slice: Struktur data yang dinamis dan fleksibel untuk menyimpan kumpulan data dengan tipe yang sama. Dalam program ini, slice digunakan untuk menyimpan nama klub yang menang dan hasil setiap pertandingan.
Loop: Digunakan untuk mengulang proses input dan pengolahan data pertandingan.
Kondisi: Digunakan untuk membuat keputusan, seperti menentukan pemenang atau mengakhiri loop.
Analisis Fitur
Fleksibilitas: Program dapat mencatat hasil pertandingan sebanyak yang diinginkan pengguna.
Kemudahan Penggunaan: Program memiliki antarmuka yang sederhana dan mudah dipahami.
Kelengkapan: Program dapat mencatat hasil pertandingan, menentukan pemenang, dan menampilkan hasil secara terorganisir.
Penggunaan Slice
Slice clubMenang digunakan untuk menyimpan nama klub yang memenangkan setiap pertandingan. Setiap kali ada pertandingan yang selesai, nama klub yang menang akan ditambahkan ke akhir slice menggunakan fungsi append. Slice hasil digunakan untuk menyimpan hasil setiap pertandingan (menang, kalah, atau seri).
Kelebihan dan Kekurangan
Kelebihan:
Sederhana dan mudah dipahami.
Fleksibel untuk jumlah pertandingan yang tidak terbatas.
Memberikan output yang jelas dan terorganisir.
Kekurangan:
Tidak ada validasi input untuk memastikan pengguna memasukkan data yang valid (misalnya, skor harus bilangan bulat positif).
Tidak ada fitur tambahan seperti menghitung total kemenangan, kekalahan, atau seri untuk setiap klub.
Potensi Pengembangan
Validasi Input: Menambahkan validasi untuk memastikan pengguna memasukkan data yang valid.
Fitur Tambahan: Menambahkan fitur untuk menghitung statistik pertandingan, seperti total kemenangan, kekalahan, dan seri untuk setiap klub.
Antarmuka Pengguna: Membuat antarmuka pengguna yang lebih interaktif, misalnya menggunakan library seperti fmt.Println untuk menampilkan hasil dalam format yang lebih menarik.
Persistensi Data: Menyimpan data pertandingan ke dalam file atau database untuk digunakan di kemudian hari.

*NOMOR 4*
Deskripsi Program
Program Go ini dirancang untuk memproses teks yang diinputkan oleh pengguna. Fungsi-fungsi utama dalam program ini meliputi:
Membaca Input: Membaca teks dari input pengguna, membatasi maksimal 127 karakter.
Membalik Teks: Membalik urutan karakter dalam teks.
Mendeteksi Palindrom: Memeriksa apakah teks tersebut merupakan palindrom.
Menampilkan Hasil: Menampilkan teks asli, teks terbalik, dan hasil deteksi palindrom.
Penjelasan Kode
Deklarasi Variabel dan Konstanta
NMAX: Konstanta yang menentukan ukuran maksimum array.
tabel: Tipe alias untuk array rune dengan ukuran NMAX, digunakan untuk menyimpan karakter-karakter teks.
m: Integer untuk menyimpan panjang teks yang diinputkan.
Fungsi-Fungsi
isiArray:
Meminta pengguna untuk memasukkan teks.
Membaca karakter per karakter hingga mencapai titik ('.') atau batas maksimum.
Menyimpan karakter-karakter yang valid ke dalam array t.
cetakArray:
Mencetak setiap karakter dalam array t ke layar.
balikanArray:
Membalik urutan karakter dalam array t menggunakan teknik swapping.
palindrom:
Membuat salinan array t ke array temp.
Membalik array temp.
Membandingkan elemen-elemen kedua array, jika ada perbedaan, maka teks bukan palindrom.
Fungsi Utama (main)
Deklarasi Variabel: Mendeklarasikan variabel tab dan m.
Membaca Input: Memanggil fungsi isiArray untuk membaca teks dari pengguna dan menyimpannya ke dalam array tab.
Membalik Teks: Memanggil fungsi balikanArray untuk membalik urutan karakter dalam array tab.
Menampilkan Teks Terbalik: Memanggil fungsi cetakArray untuk mencetak teks terbalik.
Mendeteksi Palindrom: Memanggil fungsi palindrom untuk memeriksa apakah teks tersebut merupakan palindrom.
Menampilkan Hasil Palindrom: Mencetak hasil deteksi palindrom ke layar.
Konsep Penting
Array: Digunakan untuk menyimpan karakter-karakter teks.
Loop: Digunakan untuk mengiterasi karakter-karakter dalam array.
Kondisi: Digunakan untuk membuat keputusan, seperti memeriksa akhir input atau membandingkan karakter.
Fungsi: Digunakan untuk membagi kode menjadi bagian-bagian yang lebih kecil dan dapat digunakan kembali.
Algoritma Pembalik Array: Digunakan untuk membalik urutan karakter dalam array.
Algoritma Deteksi Palindrom: Digunakan untuk memeriksa apakah teks tersebut merupakan palindrom.
