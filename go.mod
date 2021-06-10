module kubeform.dev/tests

go 1.16

require (
	github.com/appscode/go v0.0.0-20200323182826-54e98e09185a
	github.com/onsi/ginkgo v1.15.0
	github.com/onsi/gomega v1.10.5
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
	kmodules.xyz/client-go v0.0.0-20210514054158-27e164b43474
	kmodules.xyz/constants v0.0.0-20210218100002-2c304bfda278
	kubeform.dev/apimachinery v0.0.0-20210522083809-de09a8decb97
	kubeform.dev/provider-aws-api v0.0.0-20210527141620-c6a08d78b50e
	kubeform.dev/provider-azurerm-api v0.0.0-20210527141752-c119586d67ac
	kubeform.dev/provider-digitalocean-api v0.0.0-20210610093257-9199f2455636
	kubeform.dev/provider-google-api v0.0.0-20210527141831-3e6468db0d81
	kubeform.dev/provider-linode-api v0.0.0-20210610063430-55cfaa15099b
)

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.21.1-rc.0.0.20210405112358-ad4c2289ba4c

replace github.com/json-iterator/go => github.com/gomodules/json-iterator v1.1.12-0.20210506053207-2a3ea71074bc
