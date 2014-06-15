
package dijkstra

import(
    "testing"
)

func AssertEqual(a, b interface{}, t *testing.T) {
    if a != b {
        t.Logf("%v == %v", a, b)
        t.Fail()
    }
}

func TestAssertedLengthofDataFile(t *testing.T) {
    const expected = 200
    g := NewGraphFromFile("dijkstraData.txt")
    result := g.Len()
    AssertEqual(expected, result, t)
}

func TestShortestPathIsLowestSingleArc(t *testing.T) {
    const expected = 15
    v := map[int]Vertex{
        1 : {
            id: 1,
            arcs: map[int]int{
                2: 14,
                3: 15,
            },
         },
         2: {
             id : 2,
             arcs: map[int]int{
                 3: 8,
             },
         },
         3: {
             id: 3,
             arcs: map[int]int{},
         },
     }
    g := NewGraph(v)
    result := g.ShortestPath(1, 3)
    AssertEqual(result, expected, t)
    if t.Failed() {
        t.Log(g)
    }
}

func TestShortestTakesAHop(t *testing.T) {
    const expected = 10
    v := map[int]Vertex{
        1 : {
            id: 1,
            arcs: map[int]int{
                2: 2,
                3: 15,
            },
         },
         2: {
             id : 2,
             arcs: map[int]int{
                 3: 8,
             },
         },
         3: {
             id: 3,
             arcs: map[int]int{},
         },
     }
    g := NewGraph(v)
    result := g.ShortestPath(1, 3)
    AssertEqual(result, expected, t)
    if t.Failed() {
        t.Log(g)
    }
}

func TestUnnonnectedNodeReturnsMillion(t *testing.T) {
    const expected = 1000000
    v := map[int]Vertex{
        1: {
            id: 1,
            arcs: map[int]int{},
        },
        2: {
            id: 2,
            arcs: map[int]int{},
        },
    }
    g := NewGraph(v)
    result := g.ShortestPath(1, 3)
    AssertEqual(result, expected, t)
    if t.Failed() {
        t.Log(g)
    }
}
