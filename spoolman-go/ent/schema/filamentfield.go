package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FilamentField holds the schema definition for the FilamentField entity.
type FilamentField struct {
	ent.Schema
}

// Fields of the FilamentField.
func (FilamentField) Fields() []ent.Field {
	return []ent.Field{
		field.Int("filament_id"),
		field.String("key").MaxLen(64),
		field.Text("value"),
	}
}

// Edges of the FilamentField.
func (FilamentField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("filament", Filament.Type).Ref("extra").Unique().Required().Field("filament_id"),
	}
}
