package config

import (
	env "github.com/mauhftw/go-guitarists/helpers"
)

// TODO: Add a prefix env variable
// Define environment variables here
var (
	GuitaristReleaseVersion = env.GetEnvVar("GUITARISTS_RELEASE_VERSION", "development").(string)
)
