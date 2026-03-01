package main

func solveDFS(graph Graph, currNode, endNode Node) bool {
	if graph.Visited[currNode.ID] {
		return false
	}

	graph.History = append(graph.History, currNode)
	graph.Visited[currNode.ID] = true

	if currNode == endNode {
		return true
	}

	for _, node := range graph.AdjacencyList[currNode.ID] {
		var endNodeFound bool

		// skip visited nodes to prevent infinite cycles
		// for each new node, record its parent
		if graph.Visited[node.ID] {
			continue
		}
		graph.Parents[node.ID] = currNode

		// goes to first neighbor always (depth)
		// in this case - down -> right -> left -> up (based on directions array)
		endNodeFound = solveDFS(graph, node, endNode)
		if endNodeFound {
			return true
		}
	}

	return false
}
