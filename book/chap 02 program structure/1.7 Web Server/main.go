package main 

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)
var palette = []color.Color{color.White, color.Black}
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)
	
var mu sync.Mutex
var count int


func main() {
	// nếu có đường dẫn "/" thì hàm handler sẽ được gọi
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles := r.URL.Query().Get("cycles")
		// cycles := queryData[0]
		cyclesTransforrm, err := strconv.Atoi(cycles)
		if err != nil {
			fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
		}
		lissajous(w, float64(cyclesTransforrm))
	}	
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// nếu có 2 request đồng thời muốn thay đổi nó thì nó sẽ bị lỗi RACE CONDITTION và để tránh lỗi này
// chúng ta phải chắc chắn là sẽ chỉ có 1 goroutine được access vào biến trong 1 thười điểm
// và để làm điều này chúng ta dùng mu.Lock() và mu.Unlock()
// func handler(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }


// // cái Fprintf in các giá trị ra web browser
// // cái ParseForm lấy các giá trị ở trong phần query ra
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
// 	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
// 	if err := r.ParseForm(); err != nil {
// 		log.Println(err)
// 	}
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
// 	}
// }


// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}	
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
	

// server chạy cái handler cho mỗi request ở mỗi goroutine riêng biệt 
// nên nó có thể thực hiện nhiều request rất nhanh và nhẹ 1 cách đồng thời

