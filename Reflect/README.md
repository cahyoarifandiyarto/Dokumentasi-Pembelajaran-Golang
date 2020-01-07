# Reflect

Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut, atau bahkan memanipulasinya. Cakupan informasi yang bisa di dapat sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan masih banyak lagi.

Dari banyak fungsi yang tersedia di dalam package tersebut, ada 2 fungsi yang paling penting untuk diketahui, yaitu reflect.Valueof() dan reflect.Typeof()

- Fungsi reflect.Valueof() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi yang berhubungan dengan nilai pada variabel yang dicari
- Sedangkan reflect.Typeof() akan mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi yang berhubungan dengan tipe data variabel yang dicari