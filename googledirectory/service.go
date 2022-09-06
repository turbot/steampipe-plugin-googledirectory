package googledirectory

import (
	"context"
	"errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	admin "google.golang.org/api/admin/directory/v1"
)

func AdminService(ctx context.Context, d *plugin.QueryData) (*admin.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googledirectory.admin"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*admin.Service), nil
	}

	// so it was not in cache - create service
	opts, err := getSessionConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := admin.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func getSessionConfig(ctx context.Context, d *plugin.QueryData) ([]option.ClientOption, error) {
	opts := []option.ClientOption{}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var credentialContent, tokenPath string
	googledirectoryConfig := GetConfig(d.Connection)

	// 'credential_file' in connection config is DEPRECATED, and will be removed in future release
	// use `credentials` instead
	if googledirectoryConfig.Credentials != nil {
		credentialContent = *googledirectoryConfig.Credentials
	} else if googledirectoryConfig.CredentialFile != nil {
		credentialContent = *googledirectoryConfig.CredentialFile
	}

	if googledirectoryConfig.TokenPath != nil {
		tokenPath = *googledirectoryConfig.TokenPath
	}

	// If credential path provided, use domain-wide delegation
	if credentialContent != "" {
		ts, err := getTokenSource(ctx, d)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(ts))
		return opts, nil
	}

	// If token path provided, authenticate using OAuth 2.0
	if tokenPath != "" {
		path, err := expandPath(tokenPath)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithCredentialsFile(path))
		return opts, nil
	}

	return nil, nil
}

// Returns a JWT TokenSource using the configuration and the HTTP client from the provided context
func getTokenSource(ctx context.Context, d *plugin.QueryData) (oauth2.TokenSource, error) {
	// NOTE: based on https://developers.google.com/admin-sdk/directory/v1/guides/delegation#go

	// have we already created and cached the token?
	cacheKey := "googledirectory.token_source"
	if ts, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return ts.(oauth2.TokenSource), nil
	}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var impersonateUser string
	googledirectoryConfig := GetConfig(d.Connection)

	// Read credential from JSON string, or from the given path
	// NOTE: 'credential_file' in connection config is DEPRECATED, and will be removed in future release
	// use `credentials` instead
	var creds string
	if googledirectoryConfig.Credentials != nil {
		creds = *googledirectoryConfig.Credentials
	} else if googledirectoryConfig.CredentialFile != nil {
		creds = *googledirectoryConfig.CredentialFile
	}

	// Read credential
	credentialContent, err := pathOrContents(creds)
	if err != nil {
		return nil, err
	}

	if googledirectoryConfig.ImpersonatedUserEmail != nil {
		impersonateUser = *googledirectoryConfig.ImpersonatedUserEmail
	}

	// Return error, since impersonation required to authenticate using domain-wide delegation
	if impersonateUser == "" {
		return nil, errors.New("impersonated_user_email must be configured")
	}

	// Authorize the request
	config, err := google.JWTConfigFromJSON(
		[]byte(credentialContent),
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
