package game

func Glider(b *Board, x, y int) {
	b.SetAlive(x+0, y+0)
	b.SetAlive(x+1, y+1)
	b.SetAlive(x+2, y+1)
	b.SetAlive(x+0, y+2)
	b.SetAlive(x+1, y+2)
}

func LWSS(b *Board, x, y int) {
	b.SetAlive(x+1, y+0)
	b.SetAlive(x+2, y+0)

	b.SetAlive(x+0, y+1)
	b.SetAlive(x+1, y+1)
	b.SetAlive(x+2, y+1)
	b.SetAlive(x+3, y+1)

	b.SetAlive(x+0, y+2)
	b.SetAlive(x+1, y+2)
	b.SetAlive(x+3, y+2)
	b.SetAlive(x+4, y+2)

	b.SetAlive(x+2, y+3)
	b.SetAlive(x+3, y+3)
}
