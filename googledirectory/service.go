package googledirectory

import (
	"context"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	admin "google.golang.org/api/admin/directory/v1"
)

func AdminService(ctx context.Context, d *plugin.QueryData) (*admin.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "AdminService"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*admin.Service), nil
	}

	// so it was not in cache - create service
	svc, err := createDirectoryService(ctx, d.Connection)
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)
	return svc, nil
}

func createDirectoryService(ctx context.Context, connection *plugin.Connection) (*admin.Service, error) {
	// Get credential file path, and user to impersonate from config (if mentioned)
	var credentialPath, impersonateUser string
	googledirectoryConfig := GetConfig(connection)
	if &googledirectoryConfig != nil {
		if googledirectoryConfig.CredentialFile != nil {
			credentialPath = *googledirectoryConfig.CredentialFile
		}
		if googledirectoryConfig.ImpersonateUser != nil {
			impersonateUser = *googledirectoryConfig.ImpersonateUser
		}
	}

	// Read credential file
	jsonCredentials, err := ioutil.ReadFile(credentialPath)
	if err != nil {
		return nil, err
	}

	// Authorize the request
	config, err := google.JWTConfigFromJSON(jsonCredentials, admin.AdminDirectoryUserScope)
	if err != nil {
		return nil, err
	}
	config.Subject = impersonateUser

	ts := config.TokenSource(ctx)

	// Create service
	srv, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	return srv, nil
}
