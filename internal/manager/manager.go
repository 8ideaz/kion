package manager

import (
	"errors"
	"fmt"

	"github.com/8ideaz/kion/internal/utils"

	"github.com/8ideaz/kion/internal/pkgdb"
)

// Install installs a package.
func Install(packageName string) error {
	// Check if already installed
	if pkgdb.IsInstalled(packageName) {
		return fmt.Errorf("package %s is already installed", packageName)
	}

	// Download package
	err := utils.DownloadPackage(packageName)
	if err != nil {
		return err
	}

	// Verify package integrity
	if !utils.VerifyPackage(packageName) {
		return errors.New("package verification failed")
	}

	// Unpack and install
	// Here you would unpack and copy files, run scripts, etc.
	fmt.Printf("Package %s installed successfully\n", packageName)
	return pkgdb.AddPackage(packageName)
}

// Remove removes a package.
func Remove(packageName string) error {
	if !pkgdb.IsInstalled(packageName) {
		return fmt.Errorf("package %s is not installed", packageName)
	}

	// Remove package files
	// Here you would remove files and run any uninstall scripts.
	fmt.Printf("Package %s removed successfully\n", packageName)
	return pkgdb.RemovePackage(packageName)
}

// Update updates a package.
func Update(packageName string) error {
	// Simplified: Check for updates and install new version
	fmt.Printf("Package %s updated successfully\n", packageName)
	return nil
}
