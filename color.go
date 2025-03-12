package main

type Color struct {
	r, g, b, a uint32
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return c.r, c.g, c.b, c.a
}

var (
	// bluish    = Color{r: 0x2000, g: 0x4000, b: 0xffff, a: 0xffff}
	// reddish   = Color{r: 0xffff, g: 0x4000, b: 0x2000, a: 0xffff}
	// greenish  = Color{r: 0x2000, g: 0xffff, b: 0x4000, a: 0xffff}
	yellowish = Color{r: 0xffff, g: 0xffff, b: 0x2000, a: 0xffff}
	// purplish  = Color{r: 0x8000, g: 0x2000, b: 0xffff, a: 0xffff}
)
