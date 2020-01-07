# Buffered Channel

Proses transfer data pada channel secara default dilakukan dengan cara un-buffered, tidak di buffer di memori. Ketika terjadi proses kirim data via channel dari sebuah goroutine, maka harus ada goroutine yang lain yang bertugas untuk menerima data dari channel yang sama, dengan proses serah terima yang bersifat block. Maksudnya, baris kode setelah kode pengiriman dan penerimaan data tidak akan diproses sebelum proses serah terima itu sendiri selesai.

Buffered channel sedikit berbeda. Pada channel jenis ini, ditentukan angka jumlah buffernya. Angka tersebut menjadi penentu jumlah data yang bisa dikirimkan bersamaan. Selama jumlah data yangg dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan asynchronous (non blocking).

Ketika jumlah data yang dikirim sudah melewati batas buffer, maka pengiriman data hanya bisa dilakukan ketika salah satu data yang sudah terkirim sudah terambil dari channel di goroutine penerima, sehingga ada slot channel yang kosong yang bisa digunakan.

Proses pengiriman data buffered pada channel adalah asynchronous ketika jumlah data yang dikirimkan tidak melebihi batas buffer yang sudah ditentukan. Namun pada bagian channel penerimaan data selalu bersifat synchronous.