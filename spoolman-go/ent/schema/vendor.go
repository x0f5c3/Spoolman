package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vendor holds the schema definition for the Vendor entity.
type Vendor struct {
	ent.Schema
}

// Fields of the Vendor.
func (Vendor) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable().Unique(),
		field.Time("registered"),
		field.String("name").MaxLen(64),
		field.Float("empty_spool_weight").Optional(),
		field.String("comment").MaxLen(1024).Optional(),
		field.String("external_id").MaxLen(256).Optional(),
	}
}

// Edges of the Vendor.
func (Vendor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("filaments", Filament.Type),
		edge.To("extra", VendorField.Type),
	}
}
