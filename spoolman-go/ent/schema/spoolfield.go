package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SpoolField holds the schema definition for the SpoolField entity.
type SpoolField struct {
	ent.Schema
}

// Fields of the SpoolField.
func (SpoolField) Fields() []ent.Field {
	return []ent.Field{
		field.Int("spool_id"),
		field.String("key").MaxLen(64),
		field.Text("value"),
	}
}

// Edges of the SpoolField.
func (SpoolField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spool", Spool.Type).Ref("extra").Unique().Required().Field("spool_id"),
	}
}
