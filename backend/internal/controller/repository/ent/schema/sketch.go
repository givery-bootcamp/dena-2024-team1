package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sketch holds the schema definition for the Sketch entity.
type Sketch struct {
	ent.Schema
}

// Mixin of the Sketch.
func (Sketch) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
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
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("sketches").
			Unique().
			Required().
			Field("user_id"),
	}
}
