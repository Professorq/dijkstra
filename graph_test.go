
package dijkstra

import(
    "testing"
)

func Equal(a, b interface{}, t *testing.T) {
    if a != b {
        t.Logf("%v == %v", a, b)
        t.Fail()
    }
}

func TestAssertedLengthofDataFile(t *testing.T) {
    const expected = 200
    g := NewGraphFromFile("dijkstraData.txt")
    result := g.Len()
    Equal(expected, result, t)
}
