# Pointer

Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer memiliki nilai 2, maka yang dimaksud pointer adalah alamat memori dimana nilai 2 disimpan, bukan nilai 2 itu sendiri.

Variabel variabel yang memiliki alamat memori yang sama akan terhubung satu sama lain dan nilai nya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang alamat memori nya sama) yaitu nilai nya akan ikut berubah.

Ada 2 hal penting yang perlu diketahui mengenai pointer:
- Variabel biasa bisa diambil nilai pointer nya, caranya dengan menambahkan tanda ampersand(&) tepat sebelum nama variabel. Metode ini disebut dengan referencing
- Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk(*) tepat sebelum nama variabel. Metode ini disebut dengan deferencing