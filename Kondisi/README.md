# Kondisi

Kondisi digunakan untuk mengontrol alur dari sebuah program.

Yang jadi acuan oleh kondisi adalah nilai bertipe data bool (boolean)

Go memiliki 2 macam keyword untuk melakukan seleksi kondisi, yaitu if else dan switch.

## if, else if, else

Cara penerapan if-else di Go kurang lebih sama seperti pada bahasa pemrograman yang lainnya. Yang membedakan hanya tanda kurungnya, di Go tidak perlu di tulis.

### Variabel Temporary pada if - else

Variabel Temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi dimana ia ditempatkan saja. Manfaat penggunaan variabel temporary antara lain:
- Scope atau cakupan variabel jelas, hanya bisa digunakan pada blok seleksi kondisi itu saja
- Kode menjadi lebih rapi

Deklarasi variabel temporary hanya bisa dilakukan lewat metode type inference yang menggunakan tanda :=

## switch - case

Switch merupakan seleksi kondisi yang sifatnya fokus pada satu nilai/variabel, lalu kemudian di cek nilainya satu persatu.

Jika suatu nilai/variabel belum terpenuhi maka blok kondisi DEFAULT dipanggil. Bisa di bilang DEFAULT merupakan ELSE dalam sebuah switch.

Switch pada pemrograman Go memiliki perbedaan dengan bahasa pemrograman lain, Ketika sebuah CASE terpenuhi, tidak akan dilanjut ke case berikutnya, meskipun tidak ada keyword BREAK disitu.

### Switch dengan gaya if - else

di pemrograman Go, switch bisa digunakan dengan gaya ala if - else. Nilai yang akan dibandingkan tidak dituliskan setelah keyword SWITCH, melainkan akan ditulis langsung dalam bentuk perbandingan dalam keyword CASE.

### Keyword fallthrough dalam switch

Keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan ke case selanjutnya dengan tanpa menghiraukan nilai kondisinya, jadi case di pengecekan selanjutnya tersebut selalu dianggap benar (meskipun aslinya adalah salah)

## Kondisi bersarang

Kondisi bersarang adalah seleksi kondisi, yang berada dalam seleksi kondisi, yang mungkin juga berada dalam seleksi kondisi, dan seterusnya. Seleksi kondisi bersarang bisa dilakukan pada if - else, switch, atau kombinasi keduanya.