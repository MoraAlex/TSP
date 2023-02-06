package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type City struct {
	Posx int
	Posy int
}

type Res struct {
	Cities    map[int]City `json:"Cities"`
	Paths     []path       `json:"Paths"`
	Distances [100][]int   `json:"Distances"`
}

type Cities map[int]City

type path []int

const iterations = 100

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	// poblacion con funcion de aptitud
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pob [100][]int
	cities := make(Cities)
	citiesA := [20]City{
		// 0
		{
			Posx: 1,
			Posy: 3,
		},
		// 1
		{
			Posx: 2,
			Posy: 5,
		},
		// 2
		{
			Posx: 2,
			Posy: 7,
		},
		// 3
		{
			Posx: 4,
			Posy: 2,
		},
		// 4
		{
			Posx: 4,
			Posy: 4,
		},
		// 5
		{
			Posx: 4,
			Posy: 7,
		},
		// 6
		{
			Posx: 4,
			Posy: 8,
		},
		// 7
		{
			Posx: 5,
			Posy: 3,
		},
		// 8
		{
			Posx: 6,
			Posy: 1,
		},
		// 9
		{
			Posx: 6,
			Posy: 6,
		},
		// 10
		{
			Posx: 7,
			Posy: 8,
		},
		// 11
		{
			Posx: 8,
			Posy: 2,
		},
		// 12
		{
			Posx: 8,
			Posy: 7,
		},
		// 13
		{
			Posx: 9,
			Posy: 3,
		},
		// 14
		{
			Posx: 10,
			Posy: 7,
		},
		// 15
		{
			Posx: 11,
			Posy: 1,
		},
		// 16
		{
			Posx: 11,
			Posy: 4,
		},
		// 17
		{
			Posx: 11,
			Posy: 6,
		},
		// 18
		{
			Posx: 12,
			Posy: 7,
		},
		// 19
		{
			Posx: 13,
			Posy: 5,
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

	var chartsMost []path
	for i := 0; i < iterations; i++ {
		var mostOpt path
		pob, mostOpt = TSP(pob, cities)
		chartsMost = append(chartsMost, mostOpt)
	}
	var distances [100][]int
	for i, v := range chartsMost {
		if i == 0 {
			distances[i] = append(distances[0], v[len(v)-1])
		} else {
			distances[i] = append(distances[i], distances[i-1]...)
			distances[i] = append(distances[i], v[len(v)-1])
		}
	}
	res := &Res{
		Paths:     chartsMost,
		Cities:    cities,
		Distances: distances,
	}
	resjson, err := json.Marshal(res)
	_, err = w.Write(resjson)
	if err != nil {
		fmt.Println("err", err)
	}

}

func calculateDist(path []int, cities map[int]City) int {
	distance := 0.0
	for j := 0; j < len(path)-1; j++ {
		nextCity := j + 1
		xpos := cities[path[nextCity]].Posx - cities[path[j]].Posx
		ypos := cities[path[nextCity]].Posy - cities[path[j]].Posy
		distance += math.Sqrt(float64((xpos * xpos) + (ypos * ypos)))
	}
	return int(distance)
}

func TSP(pob [100][]int, cities map[int]City) ([100][]int, []int) {
	var childs [][]int
	var mostOpt []int
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
		if len(mostOpt) == 0 {
			mostOpt = winner
		} else if mostOpt[len(mostOpt)-1] > winner[len(winner)-1] {
			mostOpt = winner
		}
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
	return r, mostOpt
}
