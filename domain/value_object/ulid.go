package valueobject

import (
	"math/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

type ID string

func GenerateNewID() ID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return ID(id.String())
}

func NewID(idStr string) (ID, error) {
	id, err := ulid.Parse(idStr)
	if err != nil {
		return "", err
	}
	return ID(id.String()), nil
}

func (id ID) String() string { return string(id) }
