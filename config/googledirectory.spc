connection "googledirectory" {
  plugin = "googledirectory"

  # `credential_file` (required) - The path to a JSON credential file that contains service account credentials.
  # If `credential_file` is not specified in a connection, credentials will be loaded from:
  #  - The path specified in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable, if set
  #credential_file = "/path/to/<public_key_fingerprint>-privatekey.json"

  # `impersonate_user` (required) - The email (string) of the user which should be impersonated. Needs permissions to access the Admin APIs.
  #impersonate_user = "username@domain.com"
}