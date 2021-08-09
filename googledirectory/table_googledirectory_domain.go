package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleDirectroryDomain(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_domain",
		Description: "Domains defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryDomains,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("domain_name"),
			Hydrate:    getDirectoryDomain,
		},
		Columns: []*plugin.Column{
			{
				Name:        "domain_name",
				Description: "The domain name of the customer.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "Specifies the creation time of the domain.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreationTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "is_primary",
				Description: "Indicates if the domain is a primary domain, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "verified",
				Description: "Indicates the verification state of a domain.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "customer_id",
				Description: "The customer ID to retrieve all account roles.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("customer_id"),
			},
			{
				Name:        "etag",
				Description: "A hash of the metadata, used to ensure there were no concurrent modifications to the resource when attempting an update.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "kind",
				Description: "The type of the API resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "domain_aliases",
				Description: "A list of domain alias objects.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDirectoryDomains(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Set default value to my_customer, to represent current account
	customerID := "my_customer"
	if d.KeyColumnQuals["customer_id"] != nil {
		customerID = d.KeyColumnQuals["customer_id"].GetStringValue()
	}

	resp, err := service.Domains.List(customerID).Do()
	if err != nil {
		return nil, err
	}
	for _, user := range resp.Domains {
		d.StreamListItem(ctx, user)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getDirectoryDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryDomain")

	// Create service
	service, err := AdminService(ctx, d)
	if err != nil {
		return nil, err
	}

	domainName := d.KeyColumnQuals["domain_name"].GetStringValue()

	// Return nil, if no input provided
	if domainName == "" {
		return nil, nil
	}

	resp, err := service.Domains.Get("my_customer", domainName).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
