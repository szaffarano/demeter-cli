package demeter

import (
	"github.com/c-bata/go-prompt"
)

// Completer autocomplete the prompt options
func Completer(d prompt.Document) []prompt.Suggest {
	if len(d.TextBeforeCursor()) == 0 {
		return []prompt.Suggest{}
	}

	suggestions := make([]prompt.Suggest, len(commands), len(commands))
	idx := 0
	for _, cmd := range commands {
		suggestions[idx] = prompt.Suggest{Text: cmd.Name, Description: cmd.Description}
		idx++
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
