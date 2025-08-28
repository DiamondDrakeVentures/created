package list

import (
	"fmt"
	"slices"
	"strings"

	"github.com/diamonddrakeventures/created/tools/packdiff/crashassistant"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type ModList struct {
	mods *orderedmap.OrderedMap[string, Mod]
}

func NewModList() *ModList {
	return &ModList{
		mods: orderedmap.New[string, Mod](),
	}
}

func (list *ModList) Add(mod Mod) {
	if mod.ModID == "" {
		mod.ModID = mod.Jar
	}

	list.mods.Set(mod.ModID, mod)
}

func (list *ModList) AddFromCrashAssistant(mod crashassistant.Mod) {
	list.Add(ModFromCrashAssistant(mod))
}

func (list *ModList) Get(modID string) (mod Mod, found bool) {
	return list.mods.Get(modID)
}

func (listA *ModList) Diff(listB *ModList) (diff *Diff) {
	diff = NewDiff()

	for pairA := listA.mods.Oldest(); pairA != nil; pairA = pairA.Next() {
		diff.Store(pairA.Key, pairA.Value.Version)
	}

	diff.StartTracking()
	for pairB := listB.mods.Oldest(); pairB != nil; pairB = pairB.Next() {
		diff.Store(pairB.Key, pairB.Value.Version)
	}
	diff.StopTracking()

	return
}

func (list *ModList) Print() {
	printQ := make([]Mod, list.mods.Len())

	i := 0
	for pair := list.mods.Oldest(); pair != nil; pair = pair.Next() {
		printQ[i] = pair.Value
		i++
	}

	slices.SortStableFunc(printQ, func(a Mod, b Mod) int {
		return strings.Compare(a.ModID, b.ModID)
	})

	for _, mod := range printQ {
		fmt.Printf("%s", mod.ModID)
		if mod.Version != "" {
			fmt.Printf(": %s", mod.Version)
		}
		fmt.Println()
	}
}
