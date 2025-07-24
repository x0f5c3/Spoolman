package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VendorField holds the schema definition for the VendorField entity.
type VendorField struct {
	ent.Schema
}

// Fields of the VendorField.
func (VendorField) Fields() []ent.Field {
	return []ent.Field{
		field.Int("vendor_id"),
		field.String("key").MaxLen(64),
		field.Text("value"),
	}
}

// Edges of the VendorField.
func (VendorField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("vendor", Vendor.Type).Ref("extra").Unique().Required().Field("vendor_id"),
	}
}
