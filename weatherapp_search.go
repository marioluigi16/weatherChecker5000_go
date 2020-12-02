package main

import (
	"fmt"
	"sort"
	"strings"
)

//entry point function for returning cities by search param
func doSearch(searchVal string) []city {
	ret := make([]city, 0)

	idx := sort.Search(len(lstCity), func(idx int) bool { return searchStringsStartsWith(searchVal, idx) })
	fmt.Println(idx)

	//if no match, search returns ending index of search range
	if idx != len(lstCity) {
		fmt.Println(lstCity[idx].Name)
		ret = append(ret, lstCity[idx])
	}

	//traverse the next 4 results and add them if needed
	for n := 0; n < 4; n++ {
		idx++
		if checkNextResult(searchVal, idx, lstCity) {
			ret = append(ret, lstCity[idx])
		} else {
			return ret
		}
	}
	return ret
}

//subroutine used by binary search
func searchStringsStartsWith(val string, idx int) bool {
	sliceLen := 0
	if len(lstCity[idx].Name) < len(val) {
		sliceLen = len(lstCity[idx].Name)
	} else {
		sliceLen = len(val)
	}

	//ToLower both sides for a case insensitive compare
	return strings.ToLower(val) <= strings.ToLower(lstCity[idx].Name[:sliceLen])
}

//function for gathering any additional matches
func checkNextResult(val string, idx int, lstCity []city) bool {
	//return searchStringsStartsWith(searchVal, idx)
	sliceLen := 0
	if len(lstCity[idx].Name) < len(val) {
		sliceLen = len(lstCity[idx].Name)
	} else {
		sliceLen = len(val)
	}

	return strings.ToLower(val) == strings.ToLower(lstCity[idx].Name[:sliceLen])
}
