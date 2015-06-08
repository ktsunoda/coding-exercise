// http://www.careercup.com/question?id=6303093824159744
// http://repl.it/r47/3
package main

import "fmt"
import "math/rand"
import "sort"

type Kid struct {
    Name string
    Weight int
}

func (k Kid) String() string {
    return fmt.Sprintf("%s: %d", k.Name, k.Weight)
}


type ByWeight []Kid

func (w ByWeight) Len() int           { return len(w) }
func (w ByWeight) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w ByWeight) Less(i, j int) bool { return w[i].Weight < w[j].Weight }

func shuffle(kids []Kid) []Kid {
	ord := make([]Kid, 0, len(kids))
	for _, n := range rand.Perm(len(kids)) {
		ord = append(ord, kids[n])
	}
	return ord
}


func GetCanoes(kids []Kid) int {
    sort.Sort(ByWeight(kids))
    
    front := 0
    back := len(kids) - 1
    canoes := 0

    for back >= front {
        if (kids[len(kids) - 1].Weight > 150) {
            back--
            continue
        }
        
        canoes++
        
        if (kids[front].Weight + kids[back].Weight <= 150) {
            front++
        }
        
        back--
    }

    return canoes
}


func main() {
    kids := []Kid{}
    for i := 50; i < 102; i++ {
        kids = append(kids, Kid{fmt.Sprintf("Kid %d", i), i})
    }
    
    kids = shuffle(kids)
	sort.Sort(ByWeight(kids))
	fmt.Println(GetCanoes(kids))
}

