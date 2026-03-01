package main

// reconstructs path from vertex's parents (latest search)
//
//	0  1  2  3  4  5  (array indexes)
//
// [-1, 0, 0, 1, 0, 3] (array values, -1 is root vertex)
// if path needs to be found from vertex 0 to vertex 3, then returned array
// will be [0, 1, 3] as a graph it will look like 0 -> 1 -> 3
func reconstructPath(prnts []Node, startNode, endNode Node) []Node {
	//for i, nd := range prnts {
	//	if nd.ID == -1 {
	//		fmt.Printf(" %d ", nd.ID)
	//	} else {
	//		fmt.Printf("█%d█", nd.ID)
	//	}
	//	if i == 0 {
	//		continue
	//	}
	//	if i%9 == 0 {
	//		fmt.Printf("\n")
	//	}
	//}
	//fmt.Printf("\n")
	var path []Node
	for i := endNode; i.ID != NoParent; i = prnts[i.ID] {
		path = append(path, i)
	}

	reversedPath := reverse(path) //reverse for readability
	if reversedPath[0] == startNode {
		return reversedPath
	}

	return []Node{}
}

func reverse(arr []Node) []Node {
	newArr := make([]Node, len(arr))

	for i := 0; i < len(arr); i++ {
		newArr[i] = arr[len(arr)-1-i]
	}

	return newArr
}
