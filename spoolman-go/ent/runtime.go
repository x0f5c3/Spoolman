// Code generated by ent, DO NOT EDIT.

package ent

import (
	"spoolman-go/ent/filament"
	"spoolman-go/ent/filamentfield"
	"spoolman-go/ent/schema"
	"spoolman-go/ent/setting"
	"spoolman-go/ent/spool"
	"spoolman-go/ent/spoolfield"
	"spoolman-go/ent/vendor"
	"spoolman-go/ent/vendorfield"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	filamentFields := schema.Filament{}.Fields()
	_ = filamentFields
	// filamentDescName is the schema descriptor for name field.
	filamentDescName := filamentFields[2].Descriptor()
	// filament.NameValidator is a validator for the "name" field. It is called by the builders before save.
	filament.NameValidator = filamentDescName.Validators[0].(func(string) error)
	// filamentDescMaterial is the schema descriptor for material field.
	filamentDescMaterial := filamentFields[4].Descriptor()
	// filament.MaterialValidator is a validator for the "material" field. It is called by the builders before save.
	filament.MaterialValidator = filamentDescMaterial.Validators[0].(func(string) error)
	// filamentDescArticleNumber is the schema descriptor for article_number field.
	filamentDescArticleNumber := filamentFields[10].Descriptor()
	// filament.ArticleNumberValidator is a validator for the "article_number" field. It is called by the builders before save.
	filament.ArticleNumberValidator = filamentDescArticleNumber.Validators[0].(func(string) error)
	// filamentDescComment is the schema descriptor for comment field.
	filamentDescComment := filamentFields[11].Descriptor()
	// filament.CommentValidator is a validator for the "comment" field. It is called by the builders before save.
	filament.CommentValidator = filamentDescComment.Validators[0].(func(string) error)
	// filamentDescColorHex is the schema descriptor for color_hex field.
	filamentDescColorHex := filamentFields[14].Descriptor()
	// filament.ColorHexValidator is a validator for the "color_hex" field. It is called by the builders before save.
	filament.ColorHexValidator = filamentDescColorHex.Validators[0].(func(string) error)
	// filamentDescMultiColorHexes is the schema descriptor for multi_color_hexes field.
	filamentDescMultiColorHexes := filamentFields[15].Descriptor()
	// filament.MultiColorHexesValidator is a validator for the "multi_color_hexes" field. It is called by the builders before save.
	filament.MultiColorHexesValidator = filamentDescMultiColorHexes.Validators[0].(func(string) error)
	// filamentDescMultiColorDirection is the schema descriptor for multi_color_direction field.
	filamentDescMultiColorDirection := filamentFields[16].Descriptor()
	// filament.MultiColorDirectionValidator is a validator for the "multi_color_direction" field. It is called by the builders before save.
	filament.MultiColorDirectionValidator = filamentDescMultiColorDirection.Validators[0].(func(string) error)
	// filamentDescExternalID is the schema descriptor for external_id field.
	filamentDescExternalID := filamentFields[17].Descriptor()
	// filament.ExternalIDValidator is a validator for the "external_id" field. It is called by the builders before save.
	filament.ExternalIDValidator = filamentDescExternalID.Validators[0].(func(string) error)
	// filamentDescID is the schema descriptor for id field.
	filamentDescID := filamentFields[0].Descriptor()
	// filament.IDValidator is a validator for the "id" field. It is called by the builders before save.
	filament.IDValidator = filamentDescID.Validators[0].(func(int) error)
	filamentfieldFields := schema.FilamentField{}.Fields()
	_ = filamentfieldFields
	// filamentfieldDescKey is the schema descriptor for key field.
	filamentfieldDescKey := filamentfieldFields[1].Descriptor()
	// filamentfield.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	filamentfield.KeyValidator = filamentfieldDescKey.Validators[0].(func(string) error)
	settingFields := schema.Setting{}.Fields()
	_ = settingFields
	// settingDescKey is the schema descriptor for key field.
	settingDescKey := settingFields[0].Descriptor()
	// setting.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	setting.KeyValidator = func() func(string) error {
		validators := settingDescKey.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(key string) error {
			for _, fn := range fns {
				if err := fn(key); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	spoolFields := schema.Spool{}.Fields()
	_ = spoolFields
	// spoolDescLocation is the schema descriptor for location field.
	spoolDescLocation := spoolFields[9].Descriptor()
	// spool.LocationValidator is a validator for the "location" field. It is called by the builders before save.
	spool.LocationValidator = spoolDescLocation.Validators[0].(func(string) error)
	// spoolDescLotNr is the schema descriptor for lot_nr field.
	spoolDescLotNr := spoolFields[10].Descriptor()
	// spool.LotNrValidator is a validator for the "lot_nr" field. It is called by the builders before save.
	spool.LotNrValidator = spoolDescLotNr.Validators[0].(func(string) error)
	// spoolDescComment is the schema descriptor for comment field.
	spoolDescComment := spoolFields[11].Descriptor()
	// spool.CommentValidator is a validator for the "comment" field. It is called by the builders before save.
	spool.CommentValidator = spoolDescComment.Validators[0].(func(string) error)
	// spoolDescID is the schema descriptor for id field.
	spoolDescID := spoolFields[0].Descriptor()
	// spool.IDValidator is a validator for the "id" field. It is called by the builders before save.
	spool.IDValidator = spoolDescID.Validators[0].(func(int) error)
	spoolfieldFields := schema.SpoolField{}.Fields()
	_ = spoolfieldFields
	// spoolfieldDescKey is the schema descriptor for key field.
	spoolfieldDescKey := spoolfieldFields[1].Descriptor()
	// spoolfield.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	spoolfield.KeyValidator = spoolfieldDescKey.Validators[0].(func(string) error)
	vendorFields := schema.Vendor{}.Fields()
	_ = vendorFields
	// vendorDescName is the schema descriptor for name field.
	vendorDescName := vendorFields[2].Descriptor()
	// vendor.NameValidator is a validator for the "name" field. It is called by the builders before save.
	vendor.NameValidator = vendorDescName.Validators[0].(func(string) error)
	// vendorDescComment is the schema descriptor for comment field.
	vendorDescComment := vendorFields[4].Descriptor()
	// vendor.CommentValidator is a validator for the "comment" field. It is called by the builders before save.
	vendor.CommentValidator = vendorDescComment.Validators[0].(func(string) error)
	// vendorDescExternalID is the schema descriptor for external_id field.
	vendorDescExternalID := vendorFields[5].Descriptor()
	// vendor.ExternalIDValidator is a validator for the "external_id" field. It is called by the builders before save.
	vendor.ExternalIDValidator = vendorDescExternalID.Validators[0].(func(string) error)
	// vendorDescID is the schema descriptor for id field.
	vendorDescID := vendorFields[0].Descriptor()
	// vendor.IDValidator is a validator for the "id" field. It is called by the builders before save.
	vendor.IDValidator = vendorDescID.Validators[0].(func(int) error)
	vendorfieldFields := schema.VendorField{}.Fields()
	_ = vendorfieldFields
	// vendorfieldDescKey is the schema descriptor for key field.
	vendorfieldDescKey := vendorfieldFields[1].Descriptor()
	// vendorfield.KeyValidator is a validator for the "key" field. It is called by the builders before save.
	vendorfield.KeyValidator = vendorfieldDescKey.Validators[0].(func(string) error)
}
