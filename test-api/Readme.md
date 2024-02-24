# Mini-Proyek Go: Sistem Informasi Perpustakaan Sederhana

## Deskripsi Tugas:
Buatlah sistem informasi sederhana untuk perpustakaan menggunakan bahasa pemrograman Go yang memanfaatkan konsep variabel, tipe data, struktur data, control flow, pointer, fungsi, map, iterasi, logika, dan operasi matematika. Sistem ini harus dapat melakukan beberapa operasi dasar seperti menambahkan buku, mencari buku, dan meminjam buku.

## Spesifikasi Tugas:

## Struktur Data Buku:

- Definisikan sebuah struct bernama Book yang memiliki field Title (string), Author (string), Year (int), dan Available (bool).

- Tambahkan metode Describe() pada struct Book yang mencetak deskripsi buku dalam format: "Title: [Title], Author: [Author], Year: [Year], Available: [Yes/No]".
Koleksi Buku:

- Gunakan slice untuk menyimpan koleksi buku dalam perpustakaan.

- Buat fungsi AddBook untuk menambahkan buku baru ke dalam koleksi.

### Pencarian Buku:

- Buat fungsi FindBook yang mencari buku berdasarkan judul dan penulis, lalu cetak deskripsi buku jika ditemukan. Gunakan loop dan kondisi untuk pencarian.


### Meminjam Buku:

- Implementasikan fungsi BorrowBook yang mengubah status Available menjadi false jika buku tersedia untuk dipinjam. Pastikan untuk menggunakan pointer dalam parameter fungsi ini untuk memodifikasi slice buku secara langsung.


### Operasi Hari:

Tambahkan konsep hari menggunakan switch case untuk menentukan hari apa buku dapat dipinjam (misalnya, hanya hari kerja).


### Operasi Matematika dan Logika:

Implementasikan sistem penalty untuk keterlambatan pengembalian buku. Gunakan operasi matematika untuk menghitung denda dan logika boolean untuk menentukan apakah denda diberlakukan.

### Tampilkan Koleksi Buku:

Buat fungsi DisplayLibrary untuk menampilkan semua buku yang ada di perpustakaan, termasuk status ketersediaannya.


### Main Program:

Dalam fungsi main, buatlah beberapa buku dan tambahkan ke dalam koleksi. Implementasikan logika untuk menampilkan menu operasi (tambah buku, cari buku, pinjam buku, tampilkan koleksi) dan terima input dari pengguna untuk melakukan operasi yang diinginkan.

### Catatan:

Pastikan untuk memberikan dokumentasi yang jelas pada kode Anda, menjelaskan fungsi dan logika yang digunakan.
Coba untuk membuat kode Anda seefisien dan seefektif mungkin, menggunakan konsep seperti loop, kondisi, dan struktur data secara bijaksana.