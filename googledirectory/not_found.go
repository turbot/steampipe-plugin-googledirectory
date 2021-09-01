package googledirectory

import (
	"regexp"

	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"google.golang.org/api/googleapi"
)

// function which returns an IsNotFoundErrorPredicate for Google Directory API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if gerr, ok := err.(*googleapi.Error); ok {
			if types.ToString(gerr.Code) == "403" {
				// return true, if service API is disabled
				regexExp := regexp.MustCompile(`googleapi: Error 403: [^\s]+ API has not been used in project [0-9]+ before or it is disabled\.`)
				return regexExp.MatchString(err.Error())
			}
			return helpers.StringSliceContains(notFoundErrors, types.ToString(gerr.Code))
		}
		return false
	}
}
