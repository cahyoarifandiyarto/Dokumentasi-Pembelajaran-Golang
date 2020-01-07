# Interface

Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.

Interface merupakan tipe data. Nilai objek bertipe interface zero value nya adalah nil. Interface mulai bisa digunakan ketika sudah ada isinya, yaitu objek konkret yang memiliki definisi method minimal sama dengan yang ada di interface nya

## Embedded Interface

Interface bisa di embed ke interface lain, sama seperti struct. Cara penerapan nya juga sama, cukup dengan menuliskan nama interface yang ingin di embed ke dalam interface tujuan.