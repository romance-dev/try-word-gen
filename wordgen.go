package wordgen

import golinker "github.com/romance-dev/golinker"

// Generator produces strings containing random words that conform to a provided pattern.
type Generator struct {
	parts  []any
	minLen int
}

// New returns a Generator that is designed to be used repeatedly for a provided pattern.
//
// The pattern can contain '#' (base-10 digit), 'h'/'H' (hexadecimal digit), 'n'/'N'/'Ñ' (nouns), 'a'/'A'/'Ã' (adjectives),
// 'v'/'V'/'Ṽ' (verbs) and '*' (any word). '\' can be used to escape the special characters. Upper-case characters
// generate upper-case words. Tilde characters generate title case words.
//
// Example
//
//	`学中文:n-v-a-n` could return "学中文:dog-ran-red-banana"
func New(pattern string) Generator {
	return (*(*func(string) Generator)(golinker.SymbolPtr("github.com/romance-dev/try-word-gen~.New", ____________________plugin_builder())))(pattern)
}

func findWord(words *map[string]struct{}, list []string) string {
	return (*(*func(*map[string]struct{}, []string) string)(golinker.SymbolPtr("github.com/romance-dev/try-word-gen~.findWord", ____________________plugin_builder())))(words, list)
}

func (g Generator) string() string {
	return (*(*func(*Generator) string)(golinker.SymbolPtr("github.com/romance-dev/try-word-gen~.xVGenerator_string_plugin_builder", ____________________plugin_builder())))(&g)
}

// String generates a random string that conforms to the provided pattern.
//
// NOTE: When setting an optional max string length, it should be substantially
// larger than the theoretical minimum to prevent a potential infinite loop. A
// context can be supplied as the first argument to provide a deadline.
// If the deadline is reached, an empty string is returned.
func (g Generator) String(max ...any) string {
	return (*(*func(*Generator, ...any) string)(golinker.SymbolPtr("github.com/romance-dev/try-word-gen~.XVGenerator_String_plugin_builder", ____________________plugin_builder())))(&g, max...)
}

// Error must not be used! It is only defined so that the fmt
// package can "pretty print" the Generator since the Stringer
// interface can't be implemented.
//
// NOTE: Do not use Generator as an error.
func (g Generator) Error() string {
	return (*(*func(*Generator) string)(golinker.SymbolPtr("github.com/romance-dev/try-word-gen~.XVGenerator_Error_plugin_builder", ____________________plugin_builder())))(&g)
}
