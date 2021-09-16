connection "googledirectory" {
  plugin = "googledirectory"

  # `impersonated_user_email` (required) - The email (string) of the user which should be impersonated. Needs permissions to access the Admin APIs.
  # `impersonated_user_email` must be set, since the service account needs to impersonate a user with Admin API permissions to access the directory.
  #impersonated_user_email = "username@domain.com"

  # `credential_file` (required) - The path to a JSON credential file that contains service account credentials.
  #credential_file = "/path/to/my/creds.json"
}
