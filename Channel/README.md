# Channel

Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain nya. Mekanisme yang dilakukan adalah serah terima data lewat channel tersebut. Dalam komunikasinya, sebuah channel difungsikan sebagai pengirim di sebuah goroutine, dan juga sebagai penerima di goroutine lainnya.
Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous.

Channel merupakan sebuah variabel, dibuat dengan menggunakan kombinasi keyword make dan chan. Variabel channel memiliki satu tugas, menjadi pengirim, atau penerima data.