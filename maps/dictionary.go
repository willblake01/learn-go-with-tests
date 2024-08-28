package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound           = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists       = DictionaryErr("word already exists in the dictionary")
	ErrorWordDoesNotExist = DictionaryErr("word does not exist in the dictionary")
)

// Search finds a word in a dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a word and its definition into a dictionary.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

// Update changes the definition of a word in a dictionary.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete removes a word from a dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
