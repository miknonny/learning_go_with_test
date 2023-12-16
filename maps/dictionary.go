package dictionary

const ErrNotFound = DictionaryErr("could not find the word we were looking for")
const ErrWordExists = DictionaryErr("cannnot add word because it already exists")
const ErrWordDoesNotExist = DictionaryErr("word does not exits")

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil

}

func (d Dictionary) Search(word string) (string, error) {
	found, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return found, nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
