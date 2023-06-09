package main

import(
	"fmt"
	"bytes"
)

type IntSet struct{
	words []uint64
}

func (s *IntSet)Has(x int)bool{
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet)Add(x int){
	word, bit := x/64, uint(x%64)
	for word >= len(s.words){
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1<<bit
}

func (s *IntSet)UnionWith(t *IntSet){
	for i, tword := range t.words{
		if i < len(s.words){
			s.words[i] |= tword
		}else{
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet)String()string{
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words{
		if word == 0{
			continue
		}
		for j := 0; j<64; j++{
			if word&(1<<j)!=0{
				if buf.Len() > len("{"){
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func (s *IntSet)Len()int{
	ans := 0
	for _, word := range s.words{
		if word == 0{
			continue
		}
		for j := 0; j<64; j++{
			if word&(1<<j)!=0{
				ans++
			}
		}
	}
	return ans
}

func (s *IntSet)Remove(x int){
	word, bit := x/64, uint(x%64)
	if s.Has(x){
		s.words[word] ^= 1<<bit
	}else{
		fmt.Println("x is not in s")
	}
}

func main(){
	var x,y IntSet
	x.Add(0)
	fmt.Println(x.words)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	y.Add(9)
	y.Add(42)
	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.words)
	fmt.Println(x.Len())
	x.Remove(14)
	x.Remove(144)
	fmt.Println(&x)
}