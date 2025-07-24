package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Filament holds the schema definition for the Filament entity.
type Filament struct {
	ent.Schema
}

// Fields of the Filament.
func (Filament) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable().Unique(),
		field.Time("registered"),
		field.String("name").MaxLen(64).Optional(),
		field.Int("vendor_id").Optional(),
		field.String("material").MaxLen(64).Optional(),
		field.Float("price").Optional(),
		field.Float("density"),
		field.Float("diameter"),
		field.Float("weight").Optional(),
		field.Float("spool_weight").Optional(),
		field.String("article_number").MaxLen(64).Optional(),
		field.String("comment").MaxLen(1024).Optional(),
		field.Int("settings_extruder_temp").Optional(),
		field.Int("settings_bed_temp").Optional(),
		field.String("color_hex").MaxLen(8).Optional(),
		field.String("multi_color_hexes").MaxLen(128).Optional(),
		field.String("multi_color_direction").MaxLen(16).Optional(),
		field.String("external_id").MaxLen(256).Optional(),
	}
}

// Edges of the Filament.
func (Filament) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("vendor", Vendor.Type).Ref("filaments").Unique().Field("vendor_id"),
		edge.To("spools", Spool.Type),
		edge.To("extra", FilamentField.Type),
	}
}
