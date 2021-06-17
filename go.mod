module kubeform.dev/tests

go 1.16

require (
	github.com/appscode/go v0.0.0-20200323182826-54e98e09185a
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
	kmodules.xyz/client-go v0.0.0-20210514054158-27e164b43474
	kmodules.xyz/constants v0.0.0-20210218100002-2c304bfda278
	kubeform.dev/apimachinery v0.0.0-20210615065315-df2f79694cbe
	kubeform.dev/provider-aws-api v0.0.0-20210527141620-c6a08d78b50e
	kubeform.dev/provider-azurerm-api v0.0.0-20210527141752-c119586d67ac
	kubeform.dev/provider-digitalocean-api v0.0.0-20210610093257-9199f2455636
	kubeform.dev/provider-google-api v0.0.0-20210527141831-3e6468db0d81
	kubeform.dev/provider-linode-api v0.0.0-20210617050532-e448227cc491
	sigs.k8s.io/cli-utils v0.25.0
)

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.21.1-rc.0.0.20210405112358-ad4c2289ba4c

replace github.com/json-iterator/go => github.com/gomodules/json-iterator v1.1.12-0.20210506053207-2a3ea71074bc
