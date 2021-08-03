package googledirectory

import (
	"context"
	"errors"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	admin "google.golang.org/api/admin/directory/v1"
)

func AdminService(ctx context.Context, d *plugin.QueryData) (*admin.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googledirectory.admin"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*admin.Service), nil
	}

	// so it was not in cache - create service
	ts, err := getTokenSource(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func getTokenSource(ctx context.Context, d *plugin.QueryData) (oauth2.TokenSource, error) {
	// NOTE: based on https://developers.google.com/admin-sdk/directory/v1/guides/delegation#go

	// have we already created and cached the token?
	cacheKey := "googledirectory.token_source"
	if ts, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return ts.(oauth2.TokenSource), nil
	}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var credentialPath, impersonateUser string
	googledirectoryConfig := GetConfig(d.Connection)
	if &googledirectoryConfig != nil {
		if googledirectoryConfig.CredentialFile != nil {
			credentialPath = *googledirectoryConfig.CredentialFile
		}
		if googledirectoryConfig.ImpersonateUser != nil {
			impersonateUser = *googledirectoryConfig.ImpersonateUser
		}
	}

	// Credentials not set
	if credentialPath == "" {
		return nil, errors.New("credential_file must be configured")
	}

	// Read credential file
	jsonCredentials, err := ioutil.ReadFile(credentialPath)
	if err != nil {
		return nil, err
	}

	// Authorize the request
	config, err := google.JWTConfigFromJSON(
		jsonCredentials,
		admin.AdminDirectoryDomainReadonlyScope,
		admin.AdminDirectoryGroupReadonlyScope,
		admin.AdminDirectoryOrgunitReadonlyScope,
		admin.AdminDirectoryRolemanagementReadonlyScope,
		admin.AdminDirectoryUserReadonlyScope,
	)
	if err != nil {
		return nil, err
	}
	config.Subject = impersonateUser

	ts := config.TokenSource(ctx)

	// cache the token source
	d.ConnectionManager.Cache.Set(cacheKey, ts)

	return ts, nil
}
