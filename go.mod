module github.com/object88/shipbot

go 1.13

require (
	github.com/nlopes/slack v0.6.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.3.2
	helm.sh/helm/v3 v3.0.2
	k8s.io/cli-runtime v0.17.0
	k8s.io/kubectl v0.17.0 // indirect
)

// github.com/Azure/go-autorest/autorest has different versions for the Go
// modules than it does for releases on the repository. Note the correct
// version when updating.
// github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.0
replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
