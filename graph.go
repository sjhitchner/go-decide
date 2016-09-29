package decide

import (
	"fmt"
	graphviz "github.com/sjhitchner/go-graphviz"
	"io"
	"sort"
	"strings"
)

// Generate a Graphviz representation of the node tree
func Graph(w io.Writer, node *Node) error {
	graph := graphviz.NewGraph("DecisisonTree")
	GraphWalk(graph, nil, node)
	return graph.Output(w)
}

func GraphWalk(graph *graphviz.Graph, parent, node *Node) {
	graph.AddNode(
		fmt.Sprintf("Node%p", node),
		map[string]string{
			//"label": node.Expression.String(),
			"label": fmt.Sprintf(
				"{<f0>%s|<f1>%s\n\n\n}",
				node.Expression.String(),
				func() string {
					list := make([]string, len(node.Payload))
					for i := range node.Payload {
						list[i] = string(node.Payload[i])
					}
					sort.Strings(list)
					return strings.Join(list, "|")
				}()),
			"shape": "Mrecord",
		})

	if parent != nil {
		graph.AddEdge(
			fmt.Sprintf("Node%p", parent),
			fmt.Sprintf("Node%p", node),
			true,
			map[string]string{
				"label": func() string {
					if node == parent.True {
						return "T"
					}
					return "F"
				}(),
			},
		)
	}

	if node.True != nil {
		GraphWalk(graph, node, node.True)
	}
	if node.False != nil {
		GraphWalk(graph, node, node.False)
	}
}
