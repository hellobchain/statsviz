package plot

import (
	"fmt"
	"image/color"
)

func RGBString(r, g, b uint8) string {
	return fmt.Sprintf(`"rgb(%d,%d,%d,0)"`, r, g, b)
}

type WeightedColor struct {
	Value float64
	Color color.RGBA
}

func (c WeightedColor) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`[%f,"rgb(%d,%d,%d,%f)"]`,
		c.Value, c.Color.R, c.Color.G, c.Color.B, float64(c.Color.A)/255)
	return []byte(str), nil
}

// BlueShades https://mdigi.tools/color-shades/
var BlueShades = []WeightedColor{ // 蓝色
	{Value: 0.0, Color: color.RGBA{R: 0xea, G: 0xf8, B: 0xfd, A: 1}},
	{Value: 0.1, Color: color.RGBA{R: 0xbf, G: 0xeb, B: 0xfa, A: 1}},
	{Value: 0.2, Color: color.RGBA{R: 0x94, G: 0xdd, B: 0xf6, A: 1}},
	{Value: 0.3, Color: color.RGBA{R: 0x69, G: 0xd0, B: 0xf2, A: 1}},
	{Value: 0.4, Color: color.RGBA{R: 0x3f, G: 0xc2, B: 0xef, A: 1}},
	{Value: 0.5, Color: color.RGBA{R: 0x14, G: 0xb5, B: 0xeb, A: 1}},
	{Value: 0.6, Color: color.RGBA{R: 0x10, G: 0x94, B: 0xc0, A: 1}},
	{Value: 0.7, Color: color.RGBA{R: 0x0d, G: 0x73, B: 0x96, A: 1}},
	{Value: 0.8, Color: color.RGBA{R: 0x09, G: 0x52, B: 0x6b, A: 1}},
	{Value: 0.9, Color: color.RGBA{R: 0x05, G: 0x31, B: 0x40, A: 1}},
	{Value: 1.0, Color: color.RGBA{R: 0x02, G: 0x10, B: 0x15, A: 1}},
}

var PinkShades = []WeightedColor{ // 粉色
	{Value: 0.0, Color: color.RGBA{R: 0xfe, G: 0xe7, B: 0xf3, A: 1}},
	{Value: 0.1, Color: color.RGBA{R: 0xfc, G: 0xb6, B: 0xdc, A: 1}},
	{Value: 0.2, Color: color.RGBA{R: 0xf9, G: 0x85, B: 0xc5, A: 1}},
	{Value: 0.3, Color: color.RGBA{R: 0xf7, G: 0x55, B: 0xae, A: 1}},
	{Value: 0.4, Color: color.RGBA{R: 0xf5, G: 0x24, B: 0x96, A: 1}},
	{Value: 0.5, Color: color.RGBA{R: 0xdb, G: 0x0a, B: 0x7d, A: 1}},
	{Value: 0.6, Color: color.RGBA{R: 0xaa, G: 0x08, B: 0x61, A: 1}},
	{Value: 0.7, Color: color.RGBA{R: 0x7a, G: 0x06, B: 0x45, A: 1}},
	{Value: 0.8, Color: color.RGBA{R: 0x49, G: 0x03, B: 0x2a, A: 1}},
	{Value: 0.9, Color: color.RGBA{R: 0x18, G: 0x01, B: 0x0e, A: 1}},
	{Value: 1.0, Color: color.RGBA{A: 1}},
}

var GreenShades = []WeightedColor{ // 绿色
	{Value: 0.0, Color: color.RGBA{R: 0xed, G: 0xf7, B: 0xf2}},
	{Value: 0.1, Color: color.RGBA{R: 0xc9, G: 0xe8, B: 0xd7}},
	{Value: 0.2, Color: color.RGBA{R: 0xa5, G: 0xd9, B: 0xbc}},
	{Value: 0.3, Color: color.RGBA{R: 0x81, G: 0xca, B: 0xa2}},
	{Value: 0.4, Color: color.RGBA{R: 0x5e, G: 0xbb, B: 0x87}},
	{Value: 0.5, Color: color.RGBA{R: 0x44, G: 0xa1, B: 0x6e}},
	{Value: 0.6, Color: color.RGBA{R: 0x35, G: 0x7e, B: 0x55}},
	{Value: 0.7, Color: color.RGBA{R: 0x26, G: 0x5a, B: 0x3d}},
	{Value: 0.8, Color: color.RGBA{R: 0x17, G: 0x36, B: 0x25}},
	{Value: 0.9, Color: color.RGBA{R: 0x08, G: 0x12, B: 0x0c}},
	{Value: 1.0, Color: color.RGBA{}},
}
