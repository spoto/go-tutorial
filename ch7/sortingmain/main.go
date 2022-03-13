package main

import (
	"fmt"
	"sort"
	"time"
)
import "github.com/spoto/go-tutorial/ch7/sorting"

var tracks = []*sorting.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	sort.Sort(sorting.ByArtist(tracks))
	sorting.PrintTracks(tracks)
	sort.Sort(sort.Reverse(sorting.ByArtist(tracks)))
	sorting.PrintTracks(tracks)
	sort.Sort(sorting.ByYear(tracks))
	sorting.PrintTracks(tracks)

	sort.Sort(sorting.CustomSort{T: tracks, LessFunction: func(x, y *sorting.Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return x.Length < y.Length
	}})
	sorting.PrintTracks(tracks)

	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}
