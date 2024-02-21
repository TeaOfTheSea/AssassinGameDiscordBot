package assassinlogic

import (
	"container/list"
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

/*
############################################################

              LL / String Conversion

############################################################
*/

// This function returns a linked list if fed a string
// in the format "A -> B -> C"
func StringToLL(s string) (*list.List, error) {
	// Return empty list and error if input string is empty
	if s == "" {
		return list.New(), errors.New("Input string is empty")
	}
	var sValues []string = strings.Split(s, " -> ")
	output := list.New()
	for _, v := range sValues {
		output.PushBack(v)
	}
	return output, nil
}

// This function returns a string in the format
// "A -> B -> C" if fed a linked list
func LLToString(l *list.List) (string, error) {
	firstElement := l.Front()
	if firstElement == nil {
		return "", errors.New("Input list is empty")
	}
	output := ""
	for e := l.Front(); e != nil; e = e.Next() {
		output += fmt.Sprint(e.Value)
		if e.Next() != nil {
			output += " -> "
		}
	}
	return output, nil
}

/*
############################################################

                    Linked List Operations

############################################################
*/

func BuildLL(s []string) ([]*list.List, error) {
	if len(s) == 0 {
		return []*list.List{}, errors.New("Input slice is empty")
	}
	if len(s) == 1 {
		chain := list.New()
		chain.PushFront(s[0])
		return []*list.List{chain}, errors.New("Input slice has only one element")
	}

	// This segment builds a single linked list, but this func
	// has been altered to return multiple to more closely
	// resemble the way the game is played irl.

	// rand.Shuffle(len(s), func(i, j int) {
	// 	s[i], s[j] = s[j], s[i]
	// })
	// output := list.New()
	// for _, v := range s {
	// 	output.PushBack(v)
	// }
	// return output, nil

	// When playing assassin in real life, the way my friends
	// and I have played, we each write our names on some
	// sort of paper and then drop it into a pot. We would
	// each grab a paper without looking into the pot and then
	// look at the paper we had drawn. If the paper we had
	// drawn had our own name on it, we would put our paper
	// back and draw again until we drew someone else's name.

	// A unique consequence of this approach is that it can
	// lead to multiple chains of players, instead of a
	// single complete chain. This code has been made to
	// replicate the process which results in that outcome as
	// closely as possible.
	var chains []*list.List

	// Due to the random chance involved in this process, it
	// is possible that we may have to start again from the
	// beginning. This solutionFound loop means that we can
	// simply break to start over.
	solutionFound := false
	for solutionFound == false {
		papers := make([]string, len(s))
		copy(papers, s)
		chains = []*list.List{}
		for _, person := range s {
			// For each person walking up to the bucket, we need
			// to know if this person is in a cycle. To see why
			// we need to know this, consider a case where our
			// person has been drawn before. That would mean they
			// are already in a linked list. In this case, we want
			// to make sure that when our person draws their
			// target, we place the targer into the existing
			// linked list "chain."
			var personsNode *list.Element = nil
			var personsChain *list.List = nil
			for _, chain := range chains {
				for e := chain.Front(); e != nil; e = e.Next() {
					if person == e.Value {
						personsNode = e
						personsChain = chain
					}
				}
			}
			if personsNode == nil {
				personsChain = list.New()
				chains = append(chains, personsChain)
				personsNode = personsChain.PushBack(person)
			}

			// Now that we have eather created a new chain or
			// found our place in an existing chain, we can
			// perform the hat draw.

			targetIndex := rand.Intn(len(papers))
			if len(papers) == 1 && papers[targetIndex] == person {
				break
			} else {
				for papers[targetIndex] == person {
					targetIndex = rand.Intn(len(papers))
				}
				foundTargetInChain := 0
				for e := personsChain.Front(); e != nil; e = e.Next() {
					if papers[targetIndex] != e.Value {
						personsChain.InsertAfter(papers[targetIndex], personsNode)
					}
				}
				papers = slices.Delete(papers, targetIndex, targetIndex+1)
			}
		}
		if len(papers) == 0 {
			solutionFound = true
		}
	}
	return chains, nil
}
