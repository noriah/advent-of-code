package day17

import (
	"bytes"
	"fmt"
)

type Grid struct {
	cubes  [][][][]*Point
	points []*Point

	cycle int

	startX int
	sizeX  int

	startY int
	sizeY  int

	startZ int
	sizeZ  int

	hasW   bool
	startW int
	sizeW  int
}

type Point struct {
	state     bool
	nextState bool
	x         int
	y         int
	z         int
	w         int
	neighbors []*Point
}

// panic if x y and z are all 0
// unnecessary but neat
// func neighborIndex(x, y, z, w int) int {
// 	id := 40 + (w * 27) + (z * 9) + (y * 3) + x

// 	if id == 40 {
// 		panic("self is not neighbor")
// 	}

// 	if id > 40 {
// 		return id - 1
// 	}

// 	return id
// }

func (g *Grid) GetNumActive() int {
	active := 0

	for _, p := range g.points {

		if p.state {
			active++
		}
	}

	return active
}

func (g *Grid) Init(data []byte, cycles int, hasW bool) {
	splits := bytes.Split(data, []byte("\n"))
	initialX := len(splits[0])
	initialY := len(splits)

	totalX := initialX + 2*cycles
	totalY := initialY + 2*cycles
	totalZ := 1 + 2*cycles
	totalW := 1

	g.startX = cycles
	g.sizeX = initialX
	g.startY = cycles
	g.sizeY = initialY
	g.startZ = cycles
	g.sizeZ = 1
	g.startW = 0
	g.sizeW = 1

	neighborCount := 26

	if hasW {
		g.hasW = true
		totalW = totalZ
		g.startW = cycles
		neighborCount = 80
	}

	totalPoints := totalX * totalY * totalZ * totalW

	g.points = make([]*Point, 0, totalPoints)
	g.cubes = make([][][][]*Point, totalW)

	for w := 0; w < totalW; w++ {
		g.cubes[w] = make([][][]*Point, totalZ)

		for z := 0; z < totalZ; z++ {
			g.cubes[w][z] = make([][]*Point, totalY)

			for y := 0; y < totalY; y++ {
				g.cubes[w][z][y] = make([]*Point, totalX)

				for x := 0; x < totalX; x++ {
					p := &Point{
						x:         x,
						y:         y,
						z:         z,
						w:         w,
						neighbors: make([]*Point, 0, neighborCount),
					}

					g.cubes[w][z][y][x] = p

					g.points = append(g.points, p)
				}
			}
		}
	}

	wMax := 0
	wStart := 0

	if hasW {
		wMax = 1
		wStart = -1
	}

	for _, p := range g.points {

		for w := wStart; w <= wMax; w++ {
			aW := w + p.w

			if aW < 0 || aW >= totalW {
				continue
			}

			for z := -1; z <= 1; z++ {
				aZ := z + p.z

				if aZ < 0 || aZ >= totalZ {
					continue
				}

				for y := -1; y <= 1; y++ {
					aY := y + p.y

					if aY < 0 || aY >= totalY {
						continue
					}

					for x := -1; x <= 1; x++ {
						aX := x + p.x

						if aX < 0 || aX >= totalX {
							continue
						}

						if w == 0 && z == 0 && y == 0 && x == 0 {
							continue
						}

						p.neighbors = append(p.neighbors, g.cubes[aW][aZ][aY][aX])
					}
				}
			}
		}
	}

	// load our initial plane
	for y, s := range splits {

		for x, v := range s {
			g.cubes[g.startW][g.startZ][g.startY+y][g.startX+x].state = v == '#'
		}
	}

	// g.PrintGrids(g.startX, g.sizeX, g.startY, g.sizeY, g.startZ, g.sizeZ, g.startW, g.sizeW)

}

func (g *Grid) Cycle() {
	g.startZ--
	g.sizeZ += 2

	if g.hasW {
		g.startW--
		g.sizeW += 2
	}

	maxX := g.startX + g.sizeX
	maxY := g.startY + g.sizeY
	maxZ := g.startZ + g.sizeZ
	maxW := g.startW + g.sizeW

	// Check State

	for w := g.startW; w < maxW; w++ {

		for z := g.startZ; z < maxZ; z++ {

			for y := g.startY; y < maxY; y++ {

				for x := g.startX; x < maxX; x++ {
					p := g.cubes[w][z][y][x]
					numActive := p.ActiveNeighbors()

					if p.state {

						if numActive == 2 || numActive == 3 {
							p.nextState = true
							continue
						}

						p.nextState = false
						continue
					}

					if numActive == 3 {
						p.nextState = true
					}
				}
			}
		}
	}

	for _, p := range g.points {
		p.state = p.nextState
	}

	// g.PrintGrids(g.startX, g.sizeX, g.startY, g.sizeY, g.startZ, g.sizeZ, g.startW, g.sizeW)

	g.cycle++
	g.startX--
	g.sizeX += 2
	g.startY--
	g.sizeY += 2
}

func (g *Grid) PrintGrids(x, sizeX, y, sizeY, z, sizeZ, w, sizeW int) {

	for iW, maxW := w, w+sizeW; iW < maxW; iW++ {

		for iZ, maxZ := z, z+sizeZ; iZ < maxZ; iZ++ {
			fmt.Printf("z=%d, w=%d\n", iZ, iW)

			for iY, maxY := y, y+sizeY; iY < maxY; iY++ {

				for iX, maxX := x, x+sizeX; iX < maxX; iX++ {

					if g.cubes[iW][iZ][iY][iX].state {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}

				fmt.Println()
			}

			fmt.Println()
		}
	}
}

func (p *Point) ActiveNeighbors() int {
	active := 0

	for _, n := range p.neighbors {

		if n == nil {
			continue
		}

		if n.state {
			active++
		}
	}

	return active
}
