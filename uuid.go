package uuid

import (
	"encoding"

	"github.com/globalsign/mgo/bson"
	satori "github.com/satori/go.uuid"
)

const (
	Size = satori.Size
)

type UUID satori.UUID

var Nil = UUID{}

var (
	_ bson.Getter              = (*UUID)(nil)
	_ bson.Setter              = (*UUID)(nil)
	_ encoding.TextMarshaler   = (*UUID)(nil)
	_ encoding.TextUnmarshaler = (*UUID)(nil)
)

func FromString(raw string) (UUID, error) {
	id, err := satori.FromString(raw)
	return UUID(id), err
}

func FromBytes(raw []byte) (UUID, error) {
	id, err := satori.FromBytes(raw)
	return UUID(id), err
}

func (uuid UUID) String() string {
	return satori.UUID(uuid).String()
}

func (uuid UUID) Bytes() []byte {
	return satori.UUID(uuid).Bytes()
}

func NewV4() UUID {
	id := satori.Must(satori.NewV4())
	return UUID(id)
}

func (uuid UUID) GetBSON() (interface{}, error) {
	return uuid.String(), nil
}

func IsNil(id UUID) bool {
	return id == Nil
}

func (uuid UUID) IsNil() bool {
	return IsNil(uuid)
}

func (uuid *UUID) SetBSON(raw bson.Raw) error {
	var str string
	err := raw.Unmarshal(&str)
	if err != nil {
		return err
	}

	id, err := FromString(str)
	if err != nil {
		return err
	}

	*uuid = id

	return nil
}

func (uuid UUID) MarshalText() ([]byte, error) {
	return satori.UUID(uuid).MarshalText()
}

func (uuid *UUID) UnmarshalText(data []byte) error {
	id := satori.UUID(*uuid)
	err := id.UnmarshalText(data)
	if err != nil {
		return err
	}

	*uuid = UUID(id)
	return nil
}
