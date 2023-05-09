package main

import(
	"fmt"
	"time"
	"sort"
	"text/tabwriter"
	"os"
	"flag"
)

type Track struct{
	Title,Artist,Album string
	Year int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"I got smoke", "V is Burning", "LTDZ", 2022, length("3m32s")},
}

func length(s string)time.Duration{
	d, err := time.ParseDuration(s)
	if err != nil{
		panic(s)
	}
	return d
}

func PrintTracks(track []*Track){
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range track{
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type trackSorter struct{
	tracks []*Track
	by func(t1, t2 *Track)bool
}

func(s *trackSorter)Len()int{return len(s.tracks)}
func(s *trackSorter)Swap(i, j int){s.tracks[i], s.tracks[j] = s.tracks[j], s.tracks[i]}
func(s *trackSorter)Less(i, j int)bool{return s.by(s.tracks[i], s.tracks[j])}

func sortBy(t []*Track, by func(t1, t2 *Track)bool, reverseFlag bool){
	sorter := &trackSorter{
		tracks : t,
		by : by,
	}
	if !reverseFlag{
		sort.Sort(sorter)
	}else{
		sort.Sort(sort.Reverse(sorter))
	}
}

func byYear(t1, t2 *Track)bool{
	return t1.Year < t2.Year
}
func byArtist(t1, t2 *Track)bool{
	return t1.Artist < t2.Artist
}
func byAlbum(t1, t2 *Track)bool{
	return t1.Album < t2.Album
}
func byLength(t1, t2 *Track)bool{
	return t1.Length < t2.Length
}
func byTitle(t1, t2 *Track)bool{
	return t1.Title < t2.Title
}

func main(){
	var yearFlag = flag.Bool("year", false, "sort tracks by year")
	var artistFlag = flag.Bool("artist", false, "sort tracks by artist")
	var albumFlag = flag.Bool("album", false, "sort tracks by album")
	var titleFlag = flag.Bool("title", false, "sort tracks by title")
	var lengthFlag = flag.Bool("length", false, "sort tracks by length")
	var reverseFlag = flag.Bool("r", false, "sort reverse")
	flag.Parse()
	if *yearFlag{
		sortBy(tracks, byYear, *reverseFlag)
	}else if *artistFlag{
		sortBy(tracks, byArtist, *reverseFlag)
	}else if *albumFlag{
		sortBy(tracks, byAlbum, *reverseFlag)
	}else if *titleFlag{
		sortBy(tracks, byTitle, *reverseFlag)
	}else if *lengthFlag{
		sortBy(tracks, byLength, *reverseFlag)
	}else{
		fmt.Println("use -h for help")
		return
	}
	PrintTracks(tracks)
}