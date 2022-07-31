# Deskripsi Program

Program ini merupakan backend website untuk tugas MST-Based Clustering yang diprogramkan menggunakan golang dengan database mongodb. Program dapat menerima input file berupa excel yang berisi titik-titik dengan sumbu x dan y serta jumlah yang ingin dicluster.  Program ini dapat memvisualisasi hasil dari clustering menggunakan algoritma kruskal.

# Algoritma Kruskal
Algoritma kruskal merupakan algoritma minimum spanning tree. Algoritma ini menerima grafik yang kemudian menghasilkan suatu pohon. Tiap  titik dipasangkan dan menghasilkan suatu jarak. Jarak tersebut diurutkan dari yang terkecil sampai terbesar dengan pasangan yang unik lalu diambil secara berturut-turut agar membentuk suatu pohon. Apabila suatu pasangan titik pada graf membentuk siklik, maka pasangan tersebut diabaikan. Sehingga hasil akhir berupa grafik asiklik atau pohon.


# Analisis Kompleksitas Algoritma Kruskal
Pertama, setiap titik membentuk pasangan titik yang unik. Kompleksitas sebesar (O(V)). Kemudian edges tersebut diurutkan berdasarkan jaraknya O(ElogE) jika menggunakan mergesort. Ketiga, diambil edge (u,v) yang u tidak sama dengan v agar tidak membentuk siklik. Sehingga iterasi sebesar E dan pengecekkan sebesar logV sehingga kompleksitas adalah E.logV. Sehingga Kompleksitas total adalah O(V) + O(E log E) + O(E log V), karena E lebih kecil dibandingkan V^2, maka O(E log V).

# Cara menjalankan Backend
Prerequisites: Docker
1. Clone repository ini
2. Navigasi ke folder server
3. Jalankan program dengan mengetik pada command line `docker compose up -d terlebih dahulu, apabila tidak bisa jalankan pada command prompt `go run main.go process.go`
4. Baru jalankan program pada link frontend

# Referensi Belajar
https://www.bing.com/videos/search?q=golang+%3a%3d&docid=608039654378579726&mid=5A3E25A54B8F15C464725A3E25A54B8F15C46472&view=detail&FORM=VIRE \
https://www.educba.com/mongodb-commands/ \
https://www.mongodb.com/docs/manual/reference/command/ \
https://pkg.go.dev/github.com/gofiber/fiber/v2/middleware/cors \
