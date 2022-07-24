package main

import (
	"fmt"
)

type node struct {
	x int
	y int
}

var v [410][410]int
var m int
var n int
var q []node
var xx = []int{1, 1, 2, 2, -1, -1, -2, -2}
var yy = []int{2, -2, 1, -1, 2, -1, 1, -1}

func main() {
	var x int
	var y int
	var o node
	var u node
	for i := 0; i < 410; i++ {
		for j := 0; j < 410; j++ {
			v[i][j] = -1
		}
	}
	fmt.Scanf("%d %d %d %d", &n, &m, &x, &y)
	o.x = x
	o.y = y
	q = append(q, o)
	v[x][y] = 0
	for len(q) != 0 {
		u = q[0]
		q = q[1:]
		for i := 0; i < 8; i++ {
			dx := u.x + xx[i]
			dy := u.y + yy[i]
			if dx >= 1 && dx < n && dy >= 1 && dy <= m && v[dx][dy] == -1 {
				o.x = dx
				o.y = dy
				q = append(q, o)
				v[dx][dy] = v[u.x][u.y] + 1
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			fmt.Printf("%-5d", v[i][j])

		}
		fmt.Println()
	}
}
