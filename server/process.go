package main

import (
	"fmt"
	"math"
)

type Subset struct {
	parent, rank int
}

// Calculating the distance between 2 nodes
func processing(x1, y1, x2, y2 float64) float64 {

	return (math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1))))
}

// fungsi untuk mencari set dari element agar untuk tidak muncul 2 kali
func find(subset []Subset, i int) int {
	// Mencari akar dari root
	if subset[i].parent != i {
		subset[i].parent = find(subset, subset[i].parent)
	}
	return subset[i].parent
}

// fungsi untuk union 2 set dari x dan y
func Union(subset []Subset, x, y int) []Subset {
	xroot := find(subset, x)
	yroot := find(subset, y)

	if subset[xroot].rank < subset[yroot].rank {
		subset[xroot].parent = yroot
	} else if subset[xroot].rank > subset[yroot].rank {
		subset[yroot].parent = xroot
	} else {
		subset[yroot].parent = xroot
		subset[xroot].rank++
	}
	return subset
}

// Making the Kruskal Algorithm

func Kruskal(accept []KumpulID, jmlVertice int) []KumpulID {

	// Asumsi sudah terurut array kumpul ID

	result := []KumpulID{}
	// Buat array untuk subset
	subset := []Subset{}

	// Inisialisasi parent dan ranknya
	for i := 0; i < jmlVertice; i++ {
		sub := &Subset{}
		sub.parent = i
		sub.rank = 0
		subset = append(subset, *sub)
	}
	e := 0
	count := 0
	// Edges yang dimabil sama dengan V-1
	for e < (jmlVertice - 1) {
		// Edge terkecil dipilih
		nxtedge := accept[count]
		x := find(subset, int(nxtedge.ID1-1))
		y := find(subset, int(nxtedge.ID2-1))
		count++
		// Jika menambahakn edge tersebut tidak menyebabkan cycle, maka ditambahkan pada result dan increment e
		if x != y {
			result = append(result, nxtedge)
			subset = Union(subset, x, y)
			e++
		}
		fmt.Println(count)
	}

	return result
}

// The clustering part

func clustering(graph []KumpulID, colors []string) []KumpulID {

	for i := 0; i < len(colors)-1; i++ {
		find := maxDistance(graph)
		fmt.Print("Index with smallest distance: ")
		fmt.Println(find)
		removedID1 := graph[find].ID2
		removedID2 := graph[find].ID1
		graph = removeIndex(graph, find)
		graph = setColor(graph, removedID1, colors[i])
		graph = setColor(graph, removedID2, colors[i+1])
	}
	if len(colors) == 1 {
		graph = setColor(graph, 1, colors[0])
	}
	return graph
}

// The maximum distance find

func maxDistance(graph []KumpulID) int {
	maxIndex := 0
	for idx, j := range graph {
		if graph[maxIndex].Jarak < j.Jarak {
			maxIndex = idx
		}
	}
	return maxIndex
}

func setColor(graph []KumpulID, id float64, color string) []KumpulID {
	count := 0
	for i := range graph {
		if graph[i].ID1 == id && graph[i].Color == "" {
			graph[i].Color = color
			// fmt.Println("masuk 1")
			// fmt.Println(graph[i].Color)
			setColor(graph, graph[i].ID2, color)
		} else if graph[i].ID2 == id && graph[i].Color == "" {
			graph[i].Color = color
			//fmt.Println("masuk 1")
			setColor(graph, graph[i].ID1, color)
		}
		// fmt.Println("here it is")
		// fmt.Println(graph)
		count++
	}

	return graph
}

func removeIndex(graph []KumpulID, index int) []KumpulID {
	test := []KumpulID{}
	for i := range graph {
		if i == index {
		} else {
			test = append(test, graph[i])
		}
	}
	return test
}

func removeIndexString(graph []string, index int) []string {
	test := []string{}
	for i := range graph {
		if i == index {
		} else {
			test = append(test, graph[i])
		}
	}
	return test
}

func settingColors(nodes []Node, graph []KumpulID, colors []string) []Node {

	for _, objects := range graph {
		nodes[int(objects.ID1)].Color = objects.Color
		nodes[int(objects.ID2)].Color = objects.Color

		for i, other := range colors {
			if other == objects.Color {
				colors = removeIndexString(colors, i)
				break
			}
		}
	}
	cnt := 0
	for i, nObject := range nodes {
		if len(colors) == 0 {
			break
		} else if nObject.Color == "" && i != 0 {
			nodes[i].Color = colors[cnt]
			cnt++
		}
	}
	return nodes
}
