package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").MaxLen(64).NotEmpty().Immutable(),
		field.Text("value"),
		field.Time("last_updated"),
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}
