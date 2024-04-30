package models

type Topic struct {
	Id          string            `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string            `json:"name,omitempty" bson:"name,omitempty"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"`
	MetaData    map[string]string `json:"metaData,omitempty" bson:"metaData,omitempty"`
	UnitIds     []string          `json:"unitIds,omitempty" bson:"unitIds,omitempty"`
}