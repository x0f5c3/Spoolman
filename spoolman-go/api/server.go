package api

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"spoolman-go/ent/filament"
	setting2 "spoolman-go/ent/setting"
	"spoolman-go/ent/spool"
	"spoolman-go/ent/spoolvendor"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
	"spoolman-go/ent"
)

// convStoInts is a helper function to convert a comma-separated string to a slice of integers.
// It returns the slice of integers and a boolean indicating if any integers were found.
func convStoInts(s string) ([]int, bool) {
	q := strings.Split(s, ",")
	if len(q) > 0 {
		var ret []int
		for _, v := range q {
			if vv, err := strconv.Atoi(v); err == nil {
				ret = append(ret, vv)
			}
		}
		return ret, len(ret) > 0
	}
	return nil, false
}

// Server is a struct that implements the ServerInterface interface.
type Server struct {
	DB *ent.Client
}

// FindArticleNumbersArticleNumberGet retrieves article numbers.
func (s *Server) FindArticleNumbersArticleNumberGet(ctx iris.Context) {
	// Example implementation
	articleNumbers, err := s.DB.Filament.Query().Select("article_number").Strings(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}
	if err := ctx.JSON(articleNumbers); err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString("Failed to write response")
		return
	}
}

// BackupBackupPost handles backup creation.
func (s *Server) BackupBackupPost(ctx iris.Context) {
	// Placeholder for backup logic
	ctx.JSON(map[string]string{"message": "Backup created successfully"})
}

// ExportFilamentsExportFilamentsGet exports filaments.
func (s *Server) ExportFilamentsExportFilamentsGet(ctx iris.Context, params ExportFilamentsExportFilamentsGetParams) {
	filaments, err := s.DB.Filament.Query().All(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}
	if params.Fmt == Csv {
		ctx.ContentType("text/csv")
		w := ctx.ResponseWriter()
		cw := csv.NewWriter(w)
		if err := cw.Write([]string{"ID", "Name", "VendorID", "Material", "Price", "Density", "Diameter", "Weight"}); err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			ctx.JSON(map[string]string{"error": "Failed to write CSV header: " + err.Error()})
			return
		}
		for _, f := range filaments {
			err = cw.Write([]string{fmt.Sprint(f.ID), f.Name, fmt.Sprint(f.VendorID), f.Material, fmt.Sprint(f.Price), fmt.Sprint(f.Density), fmt.Sprint(f.Diameter), fmt.Sprint(f.Weight)})
			if err != nil {
				ctx.StopWithError(500, err)
				return
			}
		}
	} else {
		err = ctx.JSON(filaments)
		if err != nil {
			ctx.StopWithError(500, err)
		}
	}
}

// ExportSpoolsExportSpoolsGet exports spools.
func (s *Server) ExportSpoolsExportSpoolsGet(ctx iris.Context, params ExportSpoolsExportSpoolsGetParams) {
	spools, err := s.DB.Spool.Query().All(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}
	ctx.JSON(spools)
}

// ExportVendorsExportVendorsGet exports vendors.
func (s *Server) ExportVendorsExportVendorsGet(ctx iris.Context, params ExportVendorsExportVendorsGetParams) {
	vendors, err := s.DB.SpoolVendor.Query().All(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}
	ctx.JSON(vendors)
}

// GetAllExternalFilamentsExternalFilamentGet retrieves all external filaments.
func (s *Server) GetAllExternalFilamentsExternalFilamentGet(ctx iris.Context) {
	// Placeholder for external filaments logic
	ctx.JSON(map[string]string{"message": "External filaments retrieved successfully"})
}

// GetAllExternalMaterialsExternalMaterialGet retrieves all external materials.
func (s *Server) GetAllExternalMaterialsExternalMaterialGet(ctx iris.Context) {
	// Placeholder for external materials logic
	ctx.JSON(map[string]string{"message": "External materials retrieved successfully"})
}

// AddFilamentFilamentPost adds a new filament.
func (s *Server) AddFilamentFilamentPost(ctx iris.Context) {
	var params FilamentParameters
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid input: " + err.Error()})
		return
	}

	filament, err := s.DB.Filament.Create().
		SetName(*params.Name).
		SetVendorID(*params.VendorId).
		SetMaterial(*params.Material).
		SetPrice(*params.Price).
		SetDensity(params.Density).
		SetDiameter(params.Diameter).
		SetWeight(*params.Weight).
		Save(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(filament)
}

// FindFilamentsFilamentGet retrieves filaments based on query parameters.
func (s *Server) FindFilamentsFilamentGet(ctx iris.Context, params FindFilamentsFilamentGetParams) {
	query := s.DB.Filament.Query()

	if params.Name != nil {
		query = query.Where(filament.NameEQ(*params.Name))
	}
	if params.Material != nil {
		query = query.Where(filament.MaterialEQ(*params.Material))
	}
	if params.VendorId != nil {
		if qq, ok := convStoInts(*params.VendorId); ok {
			query = query.Where(filament.VendorIDIn(qq...))
		}
	}

	filaments, err := query.All(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(filaments)
}

// AddSpoolSpoolPost adds a new spool.
func (s *Server) AddSpoolSpoolPost(ctx iris.Context) {
	var params SpoolParameters
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid input: " + err.Error()})
		return
	}

	spool, err := s.DB.Spool.Create().
		SetFilamentID(params.FilamentId).
		SetInitialWeight(*params.InitialWeight).
		SetLocation(*params.Location).
		SetLotNr(*params.LotNr).
		SetPrice(*params.Price).
		SetRemainingWeight(*params.RemainingWeight).
		SetInitialWeight(*params.SpoolWeight).
		SetUsedWeight(*params.UsedWeight).
		Save(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(spool)
}

// FindSpoolSpoolGet retrieves spools based on query parameters.
func (s *Server) FindSpoolSpoolGet(ctx iris.Context, params FindSpoolSpoolGetParams) {
	query := s.DB.Spool.Query()

	if params.FilamentId != nil {
		if qq, ok := convStoInts(*params.FilamentId); ok {
			query = query.Where(spool.FilamentIDIn(qq...))
		}
	}
	if params.Location != nil {
		query = query.Where(spool.LocationEQ(*params.Location))
	}
	if params.LotNr != nil {
		query = query.Where(spool.LotNrEQ(*params.LotNr))
	}

	spools, err := query.All(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(spools)
}

// AddVendorVendorPost adds a new vendor.
func (s *Server) AddVendorVendorPost(ctx iris.Context) {
	var params VendorParameters
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid input: " + err.Error()})
		return
	}

	vendor, err := s.DB.SpoolVendor.Create().
		SetName(params.Name).
		SetComment(*params.Comment).
		SetEmptySpoolWeight(*params.EmptySpoolWeight).
		SetExternalID(*params.ExternalId).
		Save(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(vendor)
}

// FindVendorVendorGet retrieves vendors based on query parameters.
func (s *Server) FindVendorVendorGet(ctx iris.Context, params FindVendorVendorGetParams) {
	query := s.DB.SpoolVendor.Query()

	if params.Name != nil {
		query = query.Where(spoolvendor.NameEQ(*params.Name))
	}
	if params.ExternalId != nil {
		query = query.Where(spoolvendor.ExternalIDEQ(*params.ExternalId))
	}

	vendors, err := query.All(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(vendors)
}

// GetSettingSettingKeyGet retrieves a specific setting by key.
func (s *Server) GetSettingSettingKeyGet(ctx iris.Context, key string) {
	setting, err := s.DB.Setting.Query().Where(setting2.KeyEQ(key)).Only(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(map[string]string{"error": "Setting not found: " + err.Error()})
		return
	}

	ctx.JSON(setting)
}

// SetSettingSettingKeyPost updates or creates a setting by key.
func (s *Server) SetSettingSettingKeyPost(ctx iris.Context, key string) {
	var value string
	if err := ctx.ReadJSON(&value); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid input: " + err.Error()})
		return
	}

	setting, err := s.DB.Setting.Create().
		SetKey(key).
		SetValue(value).
		//OnConflict().
		//UpdateNewValues().
		Save(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(setting)
}

// DeleteFilamentFilamentFilamentIdDelete deletes a filament by its ID.
func (s *Server) DeleteFilamentFilamentFilamentIdDelete(ctx iris.Context, filamentId int) {
	err := s.DB.Filament.DeleteOneID(filamentId).Exec(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(map[string]string{"error": "Filament not found: " + err.Error()})
		return
	}

	ctx.JSON(map[string]string{"message": "Filament deleted successfully"})
}

// DeleteSpoolSpoolSpoolIdDelete deletes a spool by its ID.
func (s *Server) DeleteSpoolSpoolSpoolIdDelete(ctx iris.Context, spoolId int) {
	err := s.DB.Spool.DeleteOneID(spoolId).Exec(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(map[string]string{"error": "Spool not found: " + err.Error()})
		return
	}

	ctx.JSON(map[string]string{"message": "Spool deleted successfully"})
}

// DeleteVendorVendorVendorIdDelete deletes a vendor by its ID.
func (s *Server) DeleteVendorVendorVendorIdDelete(ctx iris.Context, vendorId int) {
	err := s.DB.SpoolVendor.DeleteOneID(vendorId).Exec(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(map[string]string{"error": "Vendor not found: " + err.Error()})
		return
	}

	ctx.JSON(map[string]string{"message": "Vendor deleted successfully"})
}

// UpdateFilamentFilamentFilamentIdPatch updates a filament by its ID.
func (s *Server) UpdateFilamentFilamentFilamentIdPatch(ctx iris.Context, filamentId int) {
	var params FilamentUpdateParameters
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid input: " + err.Error()})
		return
	}

	update := s.DB.Filament.UpdateOneID(filamentId)
	if params.Name != nil {
		update = update.SetName(*params.Name)
	}
	if params.Material != nil {
		update = update.SetMaterial(*params.Material)
	}
	if params.Price != nil {
		update = update.SetPrice(*params.Price)
	}

	save, err := update.Save(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}

	ctx.JSON(save)
}

// UpdateSpoolSpoolSpoolIdPatch updates a spool by its ID.
func (s *Server) UpdateSpoolSpoolSpoolIdPatch(ctx iris.Context, spoolId int) {
	var params SpoolUpdateParameters
	if err := ctx.ReadJSON(&params); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(map[string]string{"error": "Invalid input: " + err.Error()})
		return
	}

	update := s.DB.Spool.UpdateOneID(spoolId)
	if params.Location != nil {
		update = update.SetLocation(*params.Location)
	}
	if params.LotNr != nil {
		update = update.SetLotNr(*params.LotNr)
	}
	if params.Price != nil {
		update = update.SetPrice(*params.Price)
	}
	save, err := update.Save(context.Background())
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(map[string]string{"error": err.Error()})
		return
	}
	ctx.JSON(save)
}

// Additional methods to satisfy the ServerInterface will be implemented here.
