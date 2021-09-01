package googledirectory

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleDirectoryDomainAlias(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googledirectory_domain_alias",
		Description: "Domain alias defined in the Google Workspace directory.",
		List: &plugin.ListConfig{
			Hydrate: listDirectoryDomainAliases,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
				{
					Name:    "parent_domain_name",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "domain_alias_name",
					Require: plugin.Required,
				},
				{
					Name:    "customer_id",
					Require: plugin.Optional,
				},
			},
			Hydrate: getDirectoryDomainAlias,
		},
		Columns: []*plugin.Column{
			{
				Name:        "domain_alias_name",
				Description: "The domain alias name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "parent_domain_name",
				Description: "The parent domain name that the domain alias is associated with.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "creation_time",
				Description: "The creation time of the domain alias.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("CreationTime").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "verified",
				Description: "Indicates the verification state of a domain alias.",
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
		},
	}
}

//// LIST FUNCTION

func listDirectoryDomainAliases(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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
	var parentDomainName string
	if d.KeyColumnQuals["parent_domain_name"] != nil {
		parentDomainName = d.KeyColumnQuals["parent_domain_name"].GetStringValue()
	}

	resp, err := service.DomainAliases.List(customerID).ParentDomainName(parentDomainName).Do()
	if err != nil {
		return nil, err
	}
	for _, domainAlias := range resp.DomainAliases {
		d.StreamListItem(ctx, domainAlias)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getDirectoryDomainAlias(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDirectoryDomainAlias")

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
	domainAliasName := d.KeyColumnQuals["domain_alias_name"].GetStringValue()

	// Return nil, if no input provided
	if domainAliasName == "" {
		return nil, nil
	}

	resp, err := service.DomainAliases.Get(customerID, domainAliasName).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
