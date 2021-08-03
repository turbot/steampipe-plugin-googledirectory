package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

//// TABLE DEFINITION

func tableGoogleDirectroryOrgUnit(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_org_unit",
		Description: "OrgUnits defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryOrgUnits,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"org_unit_id", "org_unit_path"}),
			Hydrate:    getDirectoryOrgUnit,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The organizational unit's path name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "org_unit_id",
				Description: "The unique ID of the organizational unit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "org_unit_path",
				Description: "The full path to the organizational unit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "block_inheritance",
				Description: "Determines if a sub-organizational unit can inherit the settings of the parent organization.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "description",
				Description: "A short description of the organizational unit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "etag",
				Description: "Specifies the etag of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent_org_unit_id",
				Description: "The unique ID of the parent organizational unit.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent_org_unit_path",
				Description: "The organizational unit's parent path.",
				Type:        proto.ColumnType_STRING,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryOrgUnits(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Orgunits.List("my_customer").Do()
	if err != nil {
		return nil, err
	}

	for _, orgUnit := range resp.OrganizationUnits {
		d.StreamListItem(ctx, orgUnit)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getDirectoryOrgUnit(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryOrgUnit")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	orgUnitID := d.KeyColumnQuals["org_unit_id"].GetStringValue()
	orgUnitPath := d.KeyColumnQuals["org_unit_path"].GetStringValue()

	// Return nil, if no input provided
	if orgUnitID == "" && orgUnitPath == "" {
		return nil, nil
	}

	var inputStr string
	if orgUnitID == "" {
		inputStr = orgUnitPath
	} else {
		inputStr = orgUnitID
	}

	resp, err := service.Orgunits.Get("my_customer", inputStr).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
