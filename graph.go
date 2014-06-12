
package dijkstra

import(
    "bufio"
    "log"
    "os"
    "strconv"
    "strings"
)

type Edge struct{
    tail int
    head int
    weight int
}

type Graph struct {
    vertices map[int]bool
    edges []Edge
}

func NewGraph(e []Edge) *Graph {
    g := new(Graph)
    g.vertices = make(map[int]bool)
    // g.adjacent = make(map[int][]int)
    g.edges = e
    for _, v := range e {
        g.vertices[v.tail] = false
        g.vertices[v.head] = false
    }
    return g
}

func NewGraphFromFile(fn string) *Graph {
    e := EdgeListFromFile(fn)
    return NewGraph(e)
}

func EdgeListFromFile(fn string) (e []Edge) {
    f, err := os.Open(fn)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Split(string(line), "\t")
        t := words[0]
        tail, err := strconv.Atoi(t)
        for i := 1; i < len(words); i++ {
            item := strings.Split(words[i], ",")
            if len(item) < 2 {
                break
            }
            h, w := item[0], item[1]
            head, err := strconv.Atoi(h)
            if err != nil {
                log.Print(err)
            }
            weight, err := strconv.Atoi(w)
            if err != nil {
                log.Print(err)
            }
            e = append(e, Edge{tail: tail, head: head, weight:weight})
        }
        if err != nil {
            log.Print(err)
        }
        err = nil
    }
    return
}

func (g *Graph) Len() int { return len(g.vertices)  }

