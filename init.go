package wordgen

import golinker "github.com/romance-dev/golinker"

import (
	_ "embed"
)

// https://www.listdiff.com/compare-2-lists-difference-tool

// https://gist.github.com/creikey/42d23d1eec6d764e8a1d9fe7e56915c6
// https://institute.callensaxen.com/top-1500-nouns-list-english-vocabulary/
// https://github.com/datmt/English-Words-Updated/blob/master/verbsList (also includes common words)
// https://eslforums.com/adjective-examples/#Adjectives_Starting_with_B
// https://www.desiquintans.com/nounlist

var _nouns *[]byte

var _verbs *[]byte

var _ads *[]byte

var (
	nouns      *[]string // 6517
	verbs      *[]string // 5492
	adjectives *[]string // 2397
)

func parse(raw *[]byte, parsed *[]string) {
	(*(*func(*[]byte, *[]string))(golinker.SymbolPtr("github.com/romance-dev/try-word-gen~.parse", ____________________plugin_builder())))(raw, parsed)
}
