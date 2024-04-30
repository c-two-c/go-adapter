package models

type Module struct {
	Id            string            `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string            `json:"name,omitempty" bson:"name,omitempty"`
	CoverImage    string            `json:"coverImage,omitempty" bson:"coverImage,omitempty"`
	Description   string            `json:"description,omitempty" bson:"description,omitempty"`
	MetaData      map[string]string `json:"metaData,omitempty" bson:"metaData,omitempty"`
	InstructorIds []string          `json:"instructorIds,omitempty" bson:"instructorIds,omitempty"`
	TopicIds      []string          `json:"topicIds,omitempty" bson:"topics,omitempty"`
}
