## Method decleration
Method được decle theo kiểu biến thể của function nhưng được gắn thêm 1 param nữa đứng trước tên function. Cái param đó gắn function vào kiểu của cái prram đó

**Ví Dụ**
```
package geometry
import "math"
type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

- Cái extra param `p` gọi là method receiver

- `p.Distance` được gọi là selector