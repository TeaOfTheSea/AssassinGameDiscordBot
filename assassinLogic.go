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

                    	 Game Operations

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

			personsChain, personsNode, err := FindElementInChains(chains, person)
			if fmt.Sprint(err) == "Desired string was not an element in any given list" ||
				fmt.Sprint(err) == "Slice recieved is empty" {
				personsChain = list.New()
				chains = append(chains, personsChain)
				personsNode = personsChain.PushBack(person)
			} else if err != nil {
				return nil, err
			}

			// Now that we have eather created a new chain or
			// found our place in an existing chain, we can
			// perform the hat draw.

			targetIndex := rand.Intn(len(papers))
			// If the last person remaining draws themself, we
			// can't fix the situation without restarting, so this
			// if does just that.
			if len(papers) == 1 && papers[targetIndex] == person {
				break
			}
			// If, however, there are multiple left, we can just
			// redraw until we've gotten a different person.
			for papers[targetIndex] == person {
				targetIndex = rand.Intn(len(papers))
			}
			target := papers[targetIndex]
			// Once we have a target, we need to see if they are
			// already in a chain to avoid duplicating users
			// across chains.

			targetsChain, _, err := FindElementInChains(chains, target)
			if fmt.Sprint(err) == "Desired string was not an element in any given list" {
				// If the target was not found in an existing chain,
				// we can add them to the person's chain with no
				// hassle.
				personsChain.InsertAfter(target, personsNode)
			} else if err == nil {
				// If the target was found in an existing chain,
				// we need to merge the two chains together, unless
				// the target is already in the same chain as the
				// person.
				targetsChainIndex := slices.Index(chains, targetsChain)
				if personsChain.Front().Value != target {
					personsChain.PushBackList(targetsChain)
					chains = slices.Delete(chains, targetsChainIndex, targetsChainIndex+1)
				}
			} else {
				return nil, err
			}
			papers = slices.Delete(papers, targetIndex, targetIndex+1)
		}
		if len(papers) == 0 {
			solutionFound = true
		}
	}
	return chains, nil
}

// Passed a slice of all of the current chains, this
// function will remove the player from the slice they are
// in and return the slice, the hunter of our player, and
// the target of our player. hunter -> player -> target
func PlayerKilled(chains []*list.List, player string) (personsList *list.List, hunter string, target string, err error) {
	if chains == nil {
		return nil, "", "", errors.New("Input slice is nil")
	}
	if len(chains) == 0 {
		return nil, "", "", errors.New("Input slice is empty")
	}
	personsList, personsElement, err := FindElementInChains(chains, player)
	if err != nil {
		return nil, "", "", err
	}
	hunter = ""
	target = ""
	if personsElement.Prev() != nil {
		hunter = personsElement.Prev().Value.(string)
	}
	if personsElement.Next() != nil {
		target = personsElement.Next().Value.(string)
	}
	return personsList, hunter, target, nil
}

/*
############################################################

                    Linked List Operations

############################################################
*/

// If given a linked list from container/list, this function
// will return the first element e which has a value s.
func FindElementInChain(chain *list.List, s string) (*list.Element, error) {
	if chain.Front() == (*list.Element)(nil) {
		return nil, errors.New("Input linked list empty")
	}
	for e := chain.Front(); e != nil; e = e.Next() {
		if e.Value == s {
			return e, nil
		}
	}
	return nil, errors.New("Desired string was not an element in this list")
}

// If given a slice of linked lists from container/list,
// this function will return the first list and element of
// that list which contains the value s.
func FindElementInChains(chains []*list.List, s string) (*list.List, *list.Element, error) {
	if chains == nil {
		return nil, nil, errors.New("Slice recieved is nil")
	}
	if len(chains) == 0 {
		return nil, nil, errors.New("Slice recieved is empty")
	}
	for _, chain := range chains {
		element, err := FindElementInChain(chain, s)
		if err != nil && fmt.Sprint(err) != "Desired string was not an element in this list" {
			return nil, nil, err
		} else if err == nil {
			return chain, element, nil
		}
	}
	return nil, nil, errors.New("Desired string was not an element in any given list")
}
