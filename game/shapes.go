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
