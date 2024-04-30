package models

type Unit struct {
	// H(HLS)P(PPT)D(Docs) - Eg - H-topicId-randomString
	Id          string            `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string            `json:"name,omitempty" bson:"name,omitempty"`
	Description string            `json:"description,omitempty" bson:"description,omitempty"`
	MetaData    map[string]string `json:"metaData,omitempty" bson:"metaData,omitempty"`
	ModuleIds   []string          `json:"moduleIds,omitempty" bson:"moduleIds,omitempty"`
}
