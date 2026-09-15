# 🐹 Belajar Golang

Repository ini berisi catatan dan latihan saya dalam mempelajari **Golang dari dasar**.

Fokus utama repository ini adalah memahami fundamental Go terlebih dahulu sebelum masuk ke konsep yang lebih kompleks.

---

## 🎯 Learning Goal

Mempelajari Golang secara bertahap mulai dari fundamental:

```text
Basic Syntax
    ↓
Variable & Constant
    ↓
Type & Type Conversion
    ↓
Operator
    ↓
Array
    ↓
Slice
    ↓
Next Topic...
```

Saya ingin memahami konsep dan cara kerja kode, bukan hanya membuat program dengan bantuan AI.

---

# 📚 Materi yang Sudah Dipelajari

## 1. `len`

`len` digunakan untuk mengetahui jumlah atau panjang suatu data.

Contoh:

```go
name := "Golang"

fmt.Println(len(name))
```

---

## 2. Variable

Variable adalah tempat untuk menyimpan data yang nilainya dapat berubah.

### Short Variable Declaration

Menggunakan `:=`.

```go
name := "Rifki"
age := 25
```

Go akan menentukan tipe data berdasarkan nilai yang diberikan.

> `:=` hanya dapat digunakan untuk deklarasi variable di dalam function.

### Menggunakan `var`

Variable juga dapat dibuat menggunakan `var`.

```go
var name string
var age int
var isActive bool
```

Nilai dapat diberikan kemudian:

```go
var name string

name = "Rifki"
```

### Multiple Variable Declaration

Beberapa variable dapat dideklarasikan sekaligus.

```go
var (
    name string
    age  int
    city string
)
```

---

## 3. Constant

Constant adalah nilai yang tidak dapat diubah setelah dideklarasikan.

```go
const pi = 3.14
const appName = "Belajar Golang"
```

Contoh:

```go
const age = 25

// age = 26 // Error
```

Constant cocok digunakan untuk nilai yang memang tidak seharusnya berubah.

---

## 4. Type Conversion

Type conversion digunakan untuk mengubah nilai dari satu tipe data ke tipe data lainnya.

Contoh:

```go
var number int32 = 100

var result int16 = int16(number)
```

Contoh lainnya:

```go
var number int = 100

var result float64 = float64(number)
```

Type conversion berguna ketika sebuah nilai perlu digunakan dalam operasi yang membutuhkan tipe data berbeda.

---

## 5. Type Declaration

Type declaration digunakan untuk membuat tipe data baru berdasarkan tipe data yang sudah ada.

Contoh:

```go
type Age int

var age Age = 25
```

`Age` sekarang merupakan tipe yang didefinisikan sendiri berdasarkan `int`.

---

## 6. Operator Matematika

Operator matematika digunakan untuk melakukan operasi perhitungan.

Operator yang dipelajari:

| Operator | Fungsi         |
| -------- | -------------- |
| `+`      | Penjumlahan    |
| `-`      | Pengurangan    |
| `*`      | Perkalian      |
| `/`      | Pembagian      |
| `%`      | Sisa pembagian |

Contoh:

```go
a := 10
b := 3

fmt.Println(a + b)
fmt.Println(a - b)
fmt.Println(a * b)
fmt.Println(a / b)
fmt.Println(a % b)
```

---

## 7. Operator Perbandingan

Operator perbandingan digunakan untuk membandingkan dua nilai.

Hasil dari operator perbandingan selalu berupa `boolean`:

```text
true
false
```

Operator yang dipelajari:

| Operator | Fungsi                       |
| -------- | ---------------------------- |
| `==`     | Sama dengan                  |
| `!=`     | Tidak sama dengan            |
| `>`      | Lebih besar                  |
| `<`      | Lebih kecil                  |
| `>=`     | Lebih besar atau sama dengan |
| `<=`     | Lebih kecil atau sama dengan |

Contoh:

```go
age := 25

fmt.Println(age == 25)
fmt.Println(age > 20)
fmt.Println(age < 20)
```

---

## 8. Operator Boolean

Operator boolean digunakan untuk menggabungkan atau membalik kondisi `true` dan `false`.

### `&&` — AND

Hasilnya `true` hanya jika **kedua kondisi bernilai true**.

```go
isAdult := true
hasID := true

result := isAdult && hasID
```

Secara sederhana:

```text
true  && true  = true
true  && false = false
false && true  = false
false && false = false
```

### `||` — OR

Hasilnya `true` jika **minimal salah satu kondisi bernilai true**.

```go
isAdmin := false
isOwner := true

result := isAdmin || isOwner
```

Secara sederhana:

```text
true  || true  = true
true  || false = true
false || true  = true
false || false = false
```

### `!` — NOT

Digunakan untuk membalik nilai boolean.

```go
rusak := false

if !rusak {
    fmt.Println("Barang masuk ke gudang")
}
```

Artinya:

```text
!false = true
!true  = false
```

---

## 9. Array

Array adalah kumpulan data dengan **jumlah elemen yang sudah ditentukan**.

Semua elemen dalam array harus memiliki tipe data yang sama.

### Array dengan jumlah tertentu

```go
numbers := [4]int{10, 20, 30, 40}
```

Array tersebut memiliki kapasitas:

```text
4 elemen
```

Tidak dapat menambahkan elemen ke-5:

```go
numbers[4] = 50 // Error
```

### Array dengan jumlah berdasarkan isi awal

Go juga dapat menentukan jumlah elemen berdasarkan jumlah data yang diberikan menggunakan `[...]`.

```go
numbers := [...]int{10, 20, 30, 40}
```

Go akan menentukan ukuran array menjadi `4`.

> Array tetap memiliki ukuran tetap. `[...]` hanya membuat compiler menghitung jumlah elemennya secara otomatis.

---

## 10. Slice

Slice adalah struktur data yang ukurannya dapat berubah.

Slice dapat dianggap sebagai struktur yang digunakan untuk mengakses sebagian atau seluruh array dan dapat berkembang menggunakan `append`.

Contoh:

```go
numbers := []int{10, 20, 30}

numbers = append(numbers, 40)
```

Hasil:

```text
[10 20 30 40]
```

Berbeda dengan array:

```go
numbers := [3]int{10, 20, 30}
```

Array memiliki ukuran tetap.

Sedangkan slice:

```go
numbers := []int{10, 20, 30}

numbers = append(numbers, 40)
```

dapat bertambah.

### Contoh Slice

```go
names := []string{
    "Rifki",
    "Budi",
    "Andi",
}

names = append(names, "Doni")
```

Hasil:

```text
[Rifki Budi Andi Doni]
```

---

# 📌 Progress

| Materi                | Status        |
| --------------------- | ------------- |
| `len`                 | 🟢 Completed  |
| Variable              | 🟢 Completed  |
| Constant              | 🟢 Completed  |
| Type Conversion       | 🟢 Completed  |
| Type Declaration      | 🟢 Completed  |
| Operator Matematika   | 🟢 Completed  |
| Operator Perbandingan | 🟢 Completed  |
| Operator Boolean      | 🟢 Completed  |
| Array                 | 🟢 Completed  |
| Slice                 | 🟢 Completed  |
| Struct                | ⚪ Not Started |
| Map                   | ⚪ Not Started |
| Function              | ⚪ Not Started |
| Pointer               | ⚪ Not Started |
| Interface             | ⚪ Not Started |
| Error Handling        | ⚪ Not Started |
| Concurrency           | ⚪ Not Started |
| HTTP                  | ⚪ Not Started |
| REST API              | ⚪ Not Started |
| Database              | ⚪ Not Started |

---

# 🧠 Learning Principle

Saya ingin memahami Go dari fundamental sebelum menggunakan framework atau library yang lebih kompleks.

Prinsip belajar:

```text
Learn
  ↓
Practice
  ↓
Make Mistakes
  ↓
Debug
  ↓
Understand
  ↓
Repeat
```

AI digunakan sebagai **learning assistant**, bukan sebagai pengganti proses memahami kode.

Targetnya adalah mampu menjelaskan kembali kode yang dibuat dan menulis ulang konsep dasar tanpa bergantung sepenuhnya kepada AI.

---

# 📂 Repository Structure

Struktur repository akan mengikuti perkembangan materi.

```text
Belajar-Golang/
│
├── 01-basic/
│   ├── len/
│   ├── variable/
│   ├── constant/
│   ├── type-conversion/
│   ├── type-declaration/
│   ├── operator/
│   ├── comparison/
│   ├── boolean/
│   ├── array/
│   └── slice/
│
└── README.md
```

Struktur akan bertambah seiring materi baru dipelajari.

---

## 👨‍💻 Author

**Rifki Malaika**

Learning Golang from the fundamentals.

```text
Understand the code,
don't just make the code work.
```
