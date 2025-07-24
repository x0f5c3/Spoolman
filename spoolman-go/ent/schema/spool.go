package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Spool holds the schema definition for the Spool entity.
type Spool struct {
	ent.Schema
}

// Fields of the Spool.
func (Spool) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable().Unique(),
		field.Time("registered"),
		field.Time("first_used").Optional(),
		field.Time("last_used").Optional(),
		field.Float("price").Optional(),
		field.Int("filament_id"),
		field.Float("initial_weight").Optional(),
		field.Float("spool_weight").Optional(),
		field.Float("used_weight"),
		field.String("location").MaxLen(64).Optional(),
		field.String("lot_nr").MaxLen(64).Optional(),
		field.String("comment").MaxLen(1024).Optional(),
		field.Bool("archived").Optional(),
	}
}

// Edges of the Spool.
func (Spool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("filament", Filament.Type).Ref("spools").Unique().Required().Field("filament_id"),
		edge.To("extra", SpoolField.Type),
	}
}
