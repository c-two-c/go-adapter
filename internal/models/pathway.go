package models

type Pathway struct {
	Id          string            `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string            `json:"name,omitempty" bson:"name,omitempty"`
	CoverImage  string            `json:"coverImage,omitempty" bson:"coverImage,omitempty"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"`
	ModuleIds   []string          `json:"moduleIds,omitempty" bson:"moduleIds,omitempty"`
	MetaData    map[string]string `json:"metaData,omitempty" bson:"metaData,omitempty"`
}
