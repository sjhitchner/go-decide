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
	GraphWalk(graph, nil, []*Node{node})
	return graph.Output(w)
}

func GraphWalk(graph *graphviz.Graph, parent *Node, nodes []*Node) {
	for _, node := range nodes {
		graph.AddNode(
			fmt.Sprintf("Node%p", nodes),
			map[string]string{
				//"label": node.Expression.String(),
				"label": fmt.Sprintf(
					"{<f0>%s|<f1>%s}",
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
						for _, parentTrue := range parent.True {
							if node == parentTrue {
								return "T"
							}
						}
						return "F"
					}(),
				},
			)
		}

		if len(node.True) > 0 {
			GraphWalk(graph, node, node.True)
		}
		if node.False != nil {
			GraphWalk(graph, node, []*Node{node.False})
		}
	}
}
