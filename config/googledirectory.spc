connection "googledirectory" {
  plugin    = "googledirectory"

  # `credential_file` (optional) -  - The path to a JSON credential file that contains 
  # Google application credentials.  If `credential_file` is not specified in a connection,
  # credentials will be loaded from:
  #   - The path specified in the `GOOGLE_DIRECTORY_CREDENTIALS` environment variable, if set; otherwise
  #   - The standard location (`~/.config/gcloud/directory_default_credentials.json`)
  #credential_file    = "~/.config/gcloud/directory_default_credentials.json"        
}
