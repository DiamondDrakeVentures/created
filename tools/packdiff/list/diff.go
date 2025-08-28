package list

import (
	"fmt"
	"slices"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

const (
	ChangeAdd    Change = "++"
	ChangeRemove Change = "--"
	ChangeUpdate Change = "^^"
	ChangeNone   Change = "  "
)

type Change string

type Diff struct {
	tracking bool
	adds     int
	rems     int
	upds     int
	none     int
	diffs    *orderedmap.OrderedMap[string, DiffEntry]
}

func (diff *Diff) StartTracking() {
	for pair := diff.diffs.Oldest(); pair != nil; pair = pair.Next() {
		pair.Value.Tracked = false
	}
	diff.tracking = true
}

func (diff *Diff) StopTracking() {
	for pair := diff.diffs.Oldest(); pair != nil; pair = pair.Next() {
		if !pair.Value.Tracked {
			pair.Value.Change = ChangeRemove
			diff.rems++
		} else {
			diff.none++
		}
	}
	diff.tracking = false
}

func (diff *Diff) Store(modID, version string) {
	if !diff.tracking {
		diff.diffs.Set(modID, DiffEntry{
			Change:   ChangeNone,
			VersionA: version,
		})
	} else {
		entry, found := diff.diffs.Get(modID)
		if found {
			if entry.VersionA != version {
				entry.Change = ChangeUpdate
				entry.VersionB = version
				diff.upds++
			}
		} else {
			entry.Change = ChangeAdd
			entry.VersionA = version
			diff.adds++
		}
		entry.Tracked = true

		diff.diffs.Set(modID, entry)
	}
}

func (diff *Diff) Results() (results Results) {
	results = make([]ModDiff, diff.diffs.Len())

	i := 0
	for pair := diff.diffs.Oldest(); pair != nil; pair = pair.Next() {
		results[i] = ModDiff{
			Change:   pair.Value.Change,
			ModID:    pair.Key,
			VersionA: pair.Value.VersionA,
			VersionB: pair.Value.VersionB,
		}
		i++
	}

	slices.SortStableFunc(results, func(a, b ModDiff) int {
		return strings.Compare(a.ModID, b.ModID)
	})

	return
}

func (diff *Diff) Rems() (results Results) {
	return slices.DeleteFunc(diff.Results(), func(modDiff ModDiff) bool {
		return modDiff.Change != ChangeRemove
	})
}

func (diff *Diff) Adds() (results Results) {
	return slices.DeleteFunc(diff.Results(), func(modDiff ModDiff) bool {
		return modDiff.Change != ChangeAdd
	})
}

func (diff *Diff) Updates() (results Results) {
	return slices.DeleteFunc(diff.Results(), func(modDiff ModDiff) bool {
		return modDiff.Change != ChangeUpdate
	})
}

func (diff *Diff) Unchanged() (results Results) {
	return slices.DeleteFunc(diff.Results(), func(modDiff ModDiff) bool {
		return modDiff.Change != ChangeNone
	})
}

func (diff *Diff) PrintStats() {
	fmt.Printf("%d Removal(s)\t%d Addition(s)\t%d Update(s)\t%d Unchanged\n",
		diff.rems, diff.adds, diff.upds, diff.none)
}

func NewDiff() *Diff {
	return &Diff{
		diffs: orderedmap.New[string, DiffEntry](),
	}
}

type DiffEntry struct {
	Change   Change
	VersionA string
	VersionB string
	Tracked  bool
}

type Results []ModDiff

func (results Results) Print() {
	for _, mod := range results {
		fmt.Printf("%s %s", mod.Change, mod.ModID)
		if mod.VersionA != "" {
			fmt.Printf(": %s", mod.VersionA)
		}
		if mod.Change == ChangeUpdate {
			fmt.Printf(" => %s", mod.VersionB)
		}
		fmt.Println()
	}
}

type ModDiff struct {
	Change   Change
	ModID    string
	VersionA string
	VersionB string
}
