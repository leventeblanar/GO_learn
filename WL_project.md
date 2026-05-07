# Go Reference – Weather Logger projekt

---

## net/http csomag

### `http.Get(url string)`
HTTP GET kérést küld egy URL-re. Visszaad egy `*http.Response`-t és egy `error`-t.
```go
resp, err := http.Get("https://api.open-meteo.com/v1/forecast?...")
```
A `resp.Body` tartalmazza a szerver válaszát – ezt kell majd beolvasni és lezárni.

---

### `resp.Body.Close()`
Lezárja a response body-t. Mindig `defer`-rel kell meghívni a GET után, hogy a kapcsolat felszabaduljon akkor is, ha hiba történik.
```go
defer resp.Body.Close()
```

---

### `http.HandleFunc(pattern string, handler func)`
Regisztrál egy handler függvényt egy URL path-hoz. Ha jön egy kérés az adott path-ra, meghívja a handler függvényt.
```go
http.HandleFunc("/readings", server.ReadingsHandler(database))
```

---

### `http.ListenAndServe(addr string, handler http.Handler)`
Elindítja a HTTP szervert a megadott porton. Végtelen loopban figyel, soha nem tér vissza normálisan – csak hiba esetén.
```go
log.Fatal(http.ListenAndServe(":8080", nil))
```
A `nil` azt jelenti hogy a default router-t használja (amit `HandleFunc`-kal töltöttünk fel).

---

### `http.ResponseWriter`
Egy **interface** – ez reprezentálja a HTTP választ amit a szerver visszaküld a kliensnek. Három dolgot lehet rajta csinálni:
```go
w.Header().Set("Content-Type", "application/json") // fejléc beállítása
w.WriteHeader(http.StatusInternalServerError)       // státusz kód küldése
w.Write([]byte("hello"))                            // body írása
```

---

### `http.Error(w, message, statusCode)`
Hibát küld vissza a kliensnek – beállítja a státusz kódot és a hibaüzenetet.
```go
http.Error(w, "Failed to get readings", http.StatusInternalServerError)
```

---

### `http.HandlerFunc`
Egy típus ami egy függvényt handler-ré alakít. A `ReadingsHandler` closure ezt adja vissza.
```go
func ReadingsHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) { ... }
}
```

---

## encoding/json csomag

### `json.NewDecoder(r io.Reader).Decode(&v)`
Létrehoz egy decoder-t ami egy `io.Reader`-ből (pl. `resp.Body`) olvas, és a JSON-t beletölti a `v` változóba. A `&` azért kell mert pointer kell, hogy módosítani tudja a változót.
```go
var result WeatherResponse
err = json.NewDecoder(resp.Body).Decode(&result)
```

---

### `json.NewEncoder(w io.Writer).Encode(v)`
Létrehoz egy encoder-t ami egy `io.Writer`-be (pl. `http.ResponseWriter`) ír, és a `v` változót JSON-ná alakítja és beleírja.
```go
json.NewEncoder(w).Encode(readings)
```

---

## database/sql csomag

### `sql.Open(driverName, dataSourceName string)`
Előkészíti a DB kapcsolatot – de még **nem csatlakozik**. Az első tényleges műveletig vár.
```go
db, err := sql.Open("postgres", connStr)
```

---

### `db.Exec(query string, args ...interface{})`
SQL utasítást futtat ami nem ad vissza sorokat (INSERT, UPDATE, DELETE, CREATE TABLE).
```go
_, err := db.Exec(`INSERT INTO readings (timestamp, temp) VALUES ($1, $2)`,
    time.Now().UTC(), 18.5)
```
Az első visszatérési érték (`_`) egy `Result` – INSERT/UPDATE esetén megmondja hány sort érintett, de nekünk nem kellett.

---

### `db.Query(query string, args ...interface{})`
SQL lekérdezést futtat ami sorokat ad vissza. Visszaad egy `*sql.Rows`-t.
```go
rows, err := db.Query("SELECT temp, windspeed FROM readings LIMIT $1", 20)
```

---

### `rows.Next()`
Lépteti a kurzort a következő sorra. `true`-t ad vissza amíg van sor, `false`-t ha vége.
```go
for rows.Next() {
    // feldolgozás
}
```

---

### `rows.Scan(&v1, &v2, ...)`
Az aktuális sort beletölti a megadott változókba. A sorrend egyezzen a SELECT-ben lévő oszlopok sorrendjével.
```go
var w api.CurrentWeather
err := rows.Scan(&w.Temperature, &w.Windspeed, &w.Weathercode)
```

---

### `rows.Close()`
Lezárja a rows objektumot. Mindig `defer`-rel kell meghívni a Query után.
```go
defer rows.Close()
```

---

### `db.Close()`
Lezárja a DB kapcsolatot. `main`-ben `defer`-rel hívjuk.
```go
defer database.Close()
```

---

## time csomag

### `time.Now().UTC()`
Az aktuális időt adja vissza UTC timezone-ban mint `time.Time` típus.
```go
timestamp := time.Now().UTC()
```

---

### `time.NewTicker(d time.Duration)`
Létrehoz egy Ticker-t ami minden `d` időközönként küld egy jelzést a `C` channel-jén.
```go
ticker := time.NewTicker(10 * time.Minute)
for range ticker.C {
    // lefut minden 10 percben
}
```

---

### `time.Minute`, `time.Second`, `time.Hour`
Előre definiált `time.Duration` konstansok.
```go
10 * time.Minute  // 10 perc
30 * time.Second  // 30 másodperc
2 * time.Hour     // 2 óra
```

---

## fmt csomag

### `fmt.Printf(format string, args ...interface{})`
Formázott kiírás a stdout-ra. Nem tesz újsor karaktert a végére.
```go
fmt.Printf("Temp: %.1f°C, Wind: %.1f km/h\n", weather.Temperature, weather.Windspeed)
```
Fontosabb formátumok: `%d` int, `%f` float, `%.1f` egy tizedesjegy, `%s` string, `%v` bármilyen érték.

---

### `fmt.Println(args ...interface{})`
Kiír és sort emel. Szóközöket tesz az argumentumok közé automatikusan.
```go
fmt.Println("Saved to DB!")
```

---

## log csomag

### `log.Fatal(args ...interface{})`
Kiírja a hibaüzenetet és leállítja a programot (`os.Exit(1)`). Akkor használjuk ha a hiba nem kezelhető és a program nem tud tovább futni.
```go
log.Fatal(err)
```

### `log.Println(args ...interface{})`
Kiír egy üzenetet timestamp-pel – nem állítja le a programot. Goroutine-ban hasznos ahol nem akarunk Fatal-t.
```go
log.Println("Fetch error:", err)
```

---

## Nyelvi elemek

### `defer`
Elhalasztja a függvényhívást addig amíg a körülötte lévő függvény visszatér. Akkor fut le, ha a függvény lefut – akár normálisan, akár hibával.
```go
defer resp.Body.Close()  // lefut amikor FetchWeather() visszatér
defer rows.Close()       // lefut amikor GetReadings() visszatér
```

---

### `_` blank identifier
Két különböző használat:

1. **Nem használt visszatérési érték eldobása:**
```go
_, err := db.Exec(...)  // a Result-ot nem akarjuk
```

2. **Side-effect import** – importálja a csomagot (lefut az `init()`-je) de a nevét nem teszi elérhetővé:
```go
_ "github.com/lib/pq"  // regisztrálja a postgres drivert
```

---

### Closure
Egy függvény ami "bezár" egy külső változót magába. Akkor hasznos ha egy függvénynek extra paramétereket kell átadni de a hívó szigorú szignatúrát vár.
```go
func ReadingsHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // a db itt elérhető, "be van zárva" a belső függvénybe
        readings, _ := dbpackage.GetReadings(db, 20)
    }
}
```

---

### Goroutine
Könnyű párhuzamos szál. A `go` kulcsszóval indítjuk – a program nem vár rá, tovább megy.
```go
go func() {
    // ez párhuzamosan fut
    ticker := time.NewTicker(10 * time.Minute)
    for range ticker.C {
        api.FetchWeather()
    }
}()  // az ()-al rögtön el is indítjuk
```

---

### Pointer (`*` és `&`)
- `*sql.DB` – pointer típus, a memóriacímre mutat nem másolatot tartalmaz
- `&result` – a változó memóriacímét adja vissza

```go
func InitDB(connStr string) (*sql.DB, error) { ... }  // pointert ad vissza
json.NewDecoder(resp.Body).Decode(&result)             // result címét adjuk át
```

---

### Struct tag
Metadata a struct mezőhöz – megmondja a JSON decoder-nek melyik JSON key-t töltse ebbe a mezőbe.
```go
type CurrentWeather struct {
    Temperature float64 `json:"temperature"`  // a "temperature" key ide kerül
    Windspeed   float64 `json:"windspeed"`
    Weathercode int     `json:"weathercode"`
}
```

---

## Interface-ek

### `io.Reader`
Bármi ami rendelkezik `Read(p []byte) (n int, err error)` metódussal. A JSON decoder ebből olvas – nem érdekli mi a forrás (fájl, HTTP response, buffer), csak hogy tud-e belőle olvasni.

### `io.Writer`
Bármi ami rendelkezik `Write(p []byte) (n int, err error)` metódussal. A JSON encoder ebbe ír.

### `http.ResponseWriter`
A HTTP válasz interface – fejléc beállítás, státusz kód, body írás.

### Összefoglalva
```
io.Reader  →  forrás (olvasunk belőle)  →  resp.Body, fájl, buffer
io.Writer  →  cél   (írunk bele)        →  http.ResponseWriter, fájl, buffer
```
Ezért működik ugyanaz a `json.NewDecoder` / `json.NewEncoder` minden forrással/céllal – az interface a közös nyelv.