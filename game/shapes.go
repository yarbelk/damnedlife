/* Shapes for the game.  pass in an offset and the
board you want to add the shape to.
*/
package game

// Glider
// X..
// .XX
// XX.
func Glider(b *Board, x, y int) {
	b.SetAlive(x+0, y+0)
	b.SetAlive(x+1, y+1)
	b.SetAlive(x+2, y+1)
	b.SetAlive(x+0, y+2)
	b.SetAlive(x+1, y+2)
}

// LWSS Lightweight Space Ship
// .XX..
// XXXX.
// XX.XX
// ..XX.
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

// GliderGun
func GliderGun(b *Board, x, y int) {
	b.SetAlive(x+0, y+4)
	b.SetAlive(x+1, y+4)
	b.SetAlive(x+0, y+5)
	b.SetAlive(x+1, y+5)

	b.SetAlive(x+13, y+2)
	b.SetAlive(x+12, y+2)
	b.SetAlive(x+11, y+3)
	b.SetAlive(x+10, y+4)
	b.SetAlive(x+10, y+5)
	b.SetAlive(x+10, y+6)
	b.SetAlive(x+11, y+7)
	b.SetAlive(x+12, y+8)
	b.SetAlive(x+13, y+8)

	b.SetAlive(x+14, y+5)

	b.SetAlive(x+15, y+3)
	b.SetAlive(x+16, y+4)
	b.SetAlive(x+16, y+5)
	b.SetAlive(x+17, y+5)
	b.SetAlive(x+15, y+7)
	b.SetAlive(x+16, y+6)

	b.SetAlive(x+20, y+2)
	b.SetAlive(x+21, y+2)
	b.SetAlive(x+20, y+3)
	b.SetAlive(x+21, y+3)
	b.SetAlive(x+20, y+4)
	b.SetAlive(x+21, y+4)
	b.SetAlive(x+22, y+1)
	b.SetAlive(x+22, y+5)

	b.SetAlive(x+24, y+0)
	b.SetAlive(x+24, y+1)

	b.SetAlive(x+24, y+5)
	b.SetAlive(x+24, y+6)

	b.SetAlive(x+34, y+2)
	b.SetAlive(x+34, y+3)
	b.SetAlive(x+35, y+2)
	b.SetAlive(x+35, y+3)

}
