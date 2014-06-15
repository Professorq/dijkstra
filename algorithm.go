
package dijkstra


func (g *Graph) ShortestPath(src, dest int) (x int) {
    g.visit(src)
    v := g.vertices[src]
    h := make(Candidates, len(v.arcs))
    // initialize the heap with out edges from src
    for id, y := range v.arcs {
        v := g.vertices[id]
        v.dist = y
        g.vertices[id] = v
        h.Push(v)
    }
    for src != dest {
        v = h.Pop()
        src = v.id
        g.visit(src)
        for w, d := range v.arcs {
            if g.visited[w] {
                continue
            }
            c := g.vertices[w]
            distance := d + v.dist
            if distance < c.dist {
                c.dist = distance
                g.vertices[w] = c
            }
            h.Push(c)
        }
    }
    v = g.vertices[dest]
    return v.dist
}
