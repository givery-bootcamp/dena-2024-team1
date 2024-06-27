package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Sketch holds the schema definition for the Sketch entity.
type Sketch struct {
	ent.Schema
}

// Fields of the Sketch.
func (Sketch) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("image_name"),
	}
}

// Edges of the Sketch.
func (Sketch) Edges() []ent.Edge {
	return nil
}
