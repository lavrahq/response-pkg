package sdk

import "fmt"

// Package is the expected values from every Response package.
type Package interface {
	GetAuthor() string      // GetAuthor should return the author of the package.
	GetName() string        // GetName should return the name of the package.
	GetVersion() string     // GetVersion should return the version of the package.
	Install(sdk *SDK) error // Install is used to register all of the packages.
}

// RegisteredPackages are all of the packages registered with the Response Package Service.
var RegisteredPackages []Package

// SDKs is the slice of SDKs registered.
var SDKs []*SDK

// SDK is the SDK exposed during package install to add items to be installed.
type SDK struct {
	RegisteredApps []*App
}

// RegisterPackage registers a package to be installed in the Response instance.
func RegisterPackage(pkg Package) {
	RegisteredPackages = append(RegisteredPackages, pkg)
}

// FetchPackages fetches the packages registered.
func FetchPackages() {

	// get the packages
	for _, p := range RegisteredPackages {

		// notify that we found the package
		fmt.Printf("Found the `%s` package by `%s`. Beginning package installation... \n", p.GetName(), p.GetAuthor())

		pkgSDK := &SDK{}
		p.Install(pkgSDK)

		SDKs = append(SDKs, pkgSDK)
	}
}
