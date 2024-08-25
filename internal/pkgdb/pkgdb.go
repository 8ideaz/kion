package pkgdb

import (
	"encoding/json"
	"os"
)

var dbFile = "pkgdb.json"

// Package represents a package.
type Package struct {
	Name string `json:"name"`
}

// IsInstalled checks if a package is installed.
func IsInstalled(packageName string) bool {
	packages := loadPackages()
	for _, pkg := range packages {
		if pkg.Name == packageName {
			return true
		}
	}
	return false
}

// AddPackage adds a package to the database.
func AddPackage(packageName string) error {
	packages := loadPackages()
	packages = append(packages, Package{Name: packageName})
	return savePackages(packages)
}

// RemovePackage removes a package from the database.
func RemovePackage(packageName string) error {
	packages := loadPackages()
	for i, pkg := range packages {
		if pkg.Name == packageName {
			packages = append(packages[:i], packages[i+1:]...)
			break
		}
	}
	return savePackages(packages)
}

func loadPackages() []Package {
	file, err := os.Open(dbFile)
	if err != nil {
		return []Package{}
	}
	defer file.Close()

	var packages []Package
	json.NewDecoder(file).Decode(&packages)
	return packages
}

func savePackages(packages []Package) error {
	file, err := os.Create(dbFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(packages)
}
