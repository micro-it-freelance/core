package core_entity

import "github.com/goccy/go-json"

type ImplementedEntity struct{}

func (e ImplementedEntity) ToJSON() ([]byte, error) {
	return json.MarshalIndent(e, "", "\t")
}
