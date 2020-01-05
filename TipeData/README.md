# Tipe Data
Tipe Data adalah kumpulan dari beberapa tipe data

Antara lain:
- int (integer)
- string
- boolean

## Tipe Data Numerik Non Desimal

Tipe Data yang bukan nilai desimal

Secara umum ada 2 tipe data kategori ini yang perlu diketahui yaitu:
- uint = tipe data untuk bilangan positif
- int = tipe data untuk bilangan negatif dan positif

```
uint8 -> 0 - 255
uint16 -> 0 - 65535
uint32 -> 0 - 4294967295
uint64 -> 0 - 18446744073709551615
uint -> sama dengan uint32 atau uint64 (tergantung nilainya)
byte -> sama dengan uint8
int8 -> -128 - 127
int16 -> -32768 - 32767
int32 -> -2147483648 - 2147483647
int64 -> -9223372036854775808 - 9223372036854775807
int -> sama dengan int32 atau int64 (tergantung nilainya)
```

## Tipe Data Numerik Desimal

Tipe data numerik desimal yang perlu diketahui ada 2 yaitu:
- float32
- float64

Perbedaan nya berada di lebar cakupan nilai desimal yang di bisa di tampung

## Tipe Data bool (Boolean)

Tipe data bool hanya berisikan 2 nilai yaitu:
- true
- false

Tipe data ini biasanya dimanfaatkan dalam seleksi kondisi(if else/switch) dan perulangan(looping)

## Tipe Data string

Ciri khas tipe data string yaitu dimulai dengan tanda kutip dan diakhiri dengan tanda kutip ("...")

## Tipe Data nil & Zero value

Sebenarnya nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isinya nil berarti memiliki nilai kosong

Semua tipe data memiliki nilai zero value / nilai default.

```
- Zero value dari tipe data string adalah "" (string kosong)
- Zero value dari tipe data bool adalah false
- Zero value dari tipe data numerik non desimal adalah 0
- Zero value dari tipe data numerik desimal adalah 0.0
```

Zero value tidak sama dengan nil. Nil adalah nilai kosong yang berarti benar benar kosong. Beberapa tipe data yang bisa di set nilai nya dengan nil yaitu:
- pointer
- tipe data fungsi
- slice
- map
- channel
- interface kosong / interface{}