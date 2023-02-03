package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type city struct {
	posx int
	posy int
}

const iterations = 100

func main() {
	// poblacion con funcion de aptitud
	var pob [100][]int
	cities := make(map[int]city)
	citiesA := [20]city{
		// 0
		{
			posx: 1,
			posy: 3,
		},
		// 1
		{
			posx: 2,
			posy: 5,
		},
		// 2
		{
			posx: 2,
			posy: 7,
		},
		// 3
		{
			posx: 4,
			posy: 2,
		},
		// 4
		{
			posx: 4,
			posy: 4,
		},
		// 5
		{
			posx: 4,
			posy: 7,
		},
		// 6
		{
			posx: 4,
			posy: 8,
		},
		// 7
		{
			posx: 5,
			posy: 3,
		},
		// 8
		{
			posx: 6,
			posy: 1,
		},
		// 9
		{
			posx: 6,
			posy: 6,
		},
		// 10
		{
			posx: 7,
			posy: 8,
		},
		// 11
		{
			posx: 8,
			posy: 2,
		},
		// 12
		{
			posx: 8,
			posy: 7,
		},
		// 13
		{
			posx: 9,
			posy: 3,
		},
		// 14
		{
			posx: 10,
			posy: 7,
		},
		// 15
		{
			posx: 11,
			posy: 1,
		},
		// 16
		{
			posx: 11,
			posy: 4,
		},
		// 17
		{
			posx: 11,
			posy: 6,
		},
		// 18
		{
			posx: 12,
			posy: 7,
		},
		// 19
		{
			posx: 13,
			posy: 5,
		},
	}
	for i, v := range citiesA {
		cities[i+1] = v
	}

	for i := range pob {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		pob[i] = a
		distance := calculateDist(pob[i], cities)
		pob[i] = append(pob[i], int(distance))
	}
	for i := 0; i < iterations; i++ {
		pob = TSP(pob, cities)
	}
	for i, v := range pob {
		fmt.Println("response: ", i, v)
	}
}

func calculateDist(path []int, cities map[int]city) int {
	distance := 0.0
	for j := 0; j < len(path)-1; j++ {
		nextCity := j + 1
		xpos := cities[path[nextCity]].posx - cities[path[j]].posx
		ypos := cities[path[nextCity]].posy - cities[path[j]].posy
		distance += math.Sqrt(float64((xpos * xpos) + (ypos * ypos)))
	}
	return int(distance)
}

func TSP(pob [100][]int, cities map[int]city) [100][]int {
	var childs [][]int
	for i := 0; i < iterations; i++ {
		var winner []int
		distWinner := 0
		// concurso y seleccion aleatoria de los cinco contendientes
		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().UnixNano())
			min := 0
			max := 99
			// seleccionamos un contendiente aleatoriamente
			contender := rand.Intn(max-min+1) + min
			// obtenemos la distancia del contendiente
			distContender := pob[contender][len(pob[contender])-1]
			// si es la primera iteracion asignamos ese contendiente como ganador
			if len(winner) == 0 {
				winner = pob[contender]
				distWinner = pob[contender][len(pob[contender])-1]
				// si la distancia del contendiente es mas chica que la del ganador tenemos un nuevo ganador
			} else if distWinner > distContender {
				winner = pob[contender]
				distWinner = pob[contender][len(pob[contender])-1]
			}
		}

		// operacion reproductiva
		// quitamos la distancia pre guardada
		var ch []int
		ch = append(ch, winner[0:len(winner)-1]...)
		min := 0
		rand.Seed(time.Now().UnixNano())
		start := rand.Intn(19-min+1) + min
		maxLen := len(ch) - start - 1
		rand.Seed(time.Now().UnixNano())
		l := (rand.Intn(maxLen-min+1) + min) + start
		for i, j := start, l; i < j; i, j = i+1, j-1 {
			ch[i], ch[j] = ch[j], ch[i]
		}
		distance := calculateDist(ch, cities)
		ch = append(ch, distance)
		childs = append(childs, ch)
	}
	var r [100][]int
	copy(r[:], childs[:])
	return r
}
