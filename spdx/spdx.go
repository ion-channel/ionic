package spdx

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/ion-channel/ionic/aliases"
	"github.com/ion-channel/ionic/projects"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdxlib"
)

type packageInfo struct {
	Name             string
	Version          string
	DownloadLocation string
	Description      string
	Organization     string
}

// packageInfoFromPackage takes either an spdx.Package2_1 or spdx.Package2_2 and returns a packageInfo object.
// This is used to convert SPDX packages to version-agnostic representations of the data we need.
func packageInfoFromPackage(spdxPackage interface{}) packageInfo {
	var name, version, downloadLocation, description, organization string

	switch spdxPackage.(type) {
	case spdx.Package2_1:
		packageTyped := spdxPackage.(spdx.Package2_1)
		name = packageTyped.PackageName
		version = packageTyped.PackageVersion
		downloadLocation = packageTyped.PackageDownloadLocation
		description = packageTyped.PackageDescription
		organization = packageTyped.PackageSupplierOrganization
	case spdx.Package2_2:
		packageTyped := spdxPackage.(spdx.Package2_2)
		name = packageTyped.PackageName
		version = packageTyped.PackageVersion
		downloadLocation = packageTyped.PackageDownloadLocation
		description = packageTyped.PackageDescription
		organization = packageTyped.PackageSupplierOrganization
	}

	return packageInfo{
		Name:             name,
		Version:          version,
		DownloadLocation: downloadLocation,
		Description:      description,
		Organization:     organization,
	}
}

// ProjectsFromSPDX parses packages from an SPDX Document (v2.1 or v2.2) into Projects.
// The given document must be of the type *spdx.Document2_1 or *spdx.Document2_2.
// A package in the document must have a valid, resolveable PackageDownloadLocation in order to create a project
func ProjectsFromSPDX(doc interface{}, includeDependencies bool) ([]projects.Project, error) {
	// use a SPDX-version-agnostic container for tracking package info
	packageInfos := []packageInfo{}

	switch doc.(type) {
	case *spdx.Document2_1:
		docTyped := doc.(*spdx.Document2_1)
		pkgIDs := []spdx.ElementID{}

		if includeDependencies {
			// just get all of the packages
			for _, spdxPackage := range docTyped.Packages {
				pkgIDs = append(pkgIDs, spdxPackage.PackageSPDXIdentifier)
			}
			pkgIDs = spdxlib.SortElementIDs(pkgIDs)
		} else {
			// get only the top-level packages (already sorted for us)
			topLevelPkgIDs, err := spdxlib.GetDescribedPackageIDs2_1(docTyped)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve described packages from SPDX 2.1 document: %s", err.Error())
			}

			pkgIDs = topLevelPkgIDs
		}

		for _, pkgID := range pkgIDs {
			if pkg := docTyped.Packages[pkgID]; pkg != nil {
				packageInfos = append(packageInfos, packageInfoFromPackage(*pkg))
			}
		}
	case *spdx.Document2_2:
		docTyped := doc.(*spdx.Document2_2)
		pkgIDs := []spdx.ElementID{}

		if includeDependencies {
			// just get all of the packages
			for _, spdxPackage := range docTyped.Packages {
				pkgIDs = append(pkgIDs, spdxPackage.PackageSPDXIdentifier)
			}
			pkgIDs = spdxlib.SortElementIDs(pkgIDs)
		} else {
			// get only the top-level packages (already sorted for us)
			topLevelPkgIDs, err := spdxlib.GetDescribedPackageIDs2_2(docTyped)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve described packages from SPDX 2.2 document: %s", err.Error())
			}

			pkgIDs = topLevelPkgIDs
		}

		for _, pkgID := range pkgIDs {
			if pkg := docTyped.Packages[pkgID]; pkg != nil {
				packageInfos = append(packageInfos, packageInfoFromPackage(*pkg))
			}
		}
	default:
		return nil, fmt.Errorf("wrong document type given, need *spdx.Document2_1 or *spdx.Document2_2")
	}

	projs := []projects.Project{}
	for ii := range packageInfos {
		pkg := packageInfos[ii]
		// info we need to parse out of the SBOM
		var ptype, source, branch string

		tmpID := uuid.New().String()

		if pkg.DownloadLocation == "" || pkg.DownloadLocation == "NOASSERTION" || pkg.DownloadLocation == "NONE" {
			ptype = "source_unavailable"
		} else if strings.Contains(pkg.DownloadLocation, "git") {
			ptype = "git"

			// SPDX spec says that git URLs can look like "git+https://github.com/..."
			// we need to strip off the "git+"
			if strings.Index(pkg.DownloadLocation, "git+") == 0 {
				source = pkg.DownloadLocation[4:]
			} else {
				source = pkg.DownloadLocation
			}

			// try to figure out which branch to monitor. The branch name will be after the last '@'.
			var foundBranchName bool
			branchDelimiterIndex := strings.LastIndex(source, "@")
			if branchDelimiterIndex != -1 {
				// if there is a ':' after the last '@', we know we do not have a branch name,
				// because branch names cannot contain colons. A git URL with an '@' that does not denote a branch will
				// always also contain a colon.
				possibleBranch := source[branchDelimiterIndex+1:]
				// this thing could be a branch name, or it could be a commit hash.
				// To determine which it is, check if it looks like a commit hash (exactly 40 lower-case hex characters)
				commitHashRegex := regexp.MustCompile("[a-f0-9]{40}")
				if !strings.Contains(possibleBranch, ":") && !commitHashRegex.MatchString(possibleBranch) {
					branch = possibleBranch
					foundBranchName = true
					// now we need to remove the branch name from the source
					source = source[0:branchDelimiterIndex]
				}
			}
			if !foundBranchName {
				// use the remote's default branch
				branch = "HEAD"
			}
		} else {
			source = pkg.DownloadLocation
			ptype = "artifact"
		}

		proj := projects.Project{
			ID:          &tmpID,
			Branch:      &branch,
			Description: &pkg.Description,
			Type:        &ptype,
			Source:      &source,
			Name:        &pkg.Name,
			Active:      true,
			Monitor:     true,
		}

		// check if version, org, or name are not empty strings
		if len(pkg.Version) > 0 || len(pkg.Organization) > 0 || len(pkg.Name) > 0 {
			proj.Aliases = []aliases.Alias{{
				Name:    pkg.Name,
				Org:     pkg.Organization,
				Version: pkg.Version,
			}}
		}

		projs = append(projs, proj)

	}

	return projs, nil
}

// Helper function to parse email from SPDX Creator info
// SPDX email comes in the form Creator: Person: My Name (myname@mail.com)
// returns empty string if no email is found
func parseCreatorEmail(creatorPersons []string) string {
	if len(creatorPersons) > 0 {
		re := regexp.MustCompile(`\((.*?)\)`)
		email := re.FindStringSubmatch(creatorPersons[0])
		if len(email) > 0 && email != nil {
			return email[1]
		}
	}
	return ""
}

// Pretty print functions for SPDX file license/package information.

func spdxV2_2(doc *spdx.Document2_2) {
	// print the struct containing the SPDX file's Creation Info section data
	fmt.Printf("==============\n")
	fmt.Printf("Creation info:\n")
	fmt.Printf("==============\n")
	fmt.Printf("%#v\n\n", doc.CreationInfo)
	pkgIDs, err := spdxlib.GetDescribedPackageIDs2_2(doc)
	if err != nil {
		fmt.Printf("Unable to get describe packages from SPDX document: %v\n", err)
		return
	}

	// SPDX Document does contain packages, so we'll go through each one
	for _, pkgID := range pkgIDs {
		pkg, ok := doc.Packages[pkgID]
		if !ok {
			fmt.Printf("Package %s has described relationship but ID not found\n", string(pkgID))
			continue
		}

		// check whether the package had its files analyzed
		if !pkg.FilesAnalyzed {
			fmt.Printf("Package %s (%s) had FilesAnalyzed: false\n", string(pkgID), pkg.PackageName)
			continue
		}

		// also check whether the package has any files present
		if pkg.Files == nil || len(pkg.Files) < 1 {
			fmt.Printf("Package %s (%s) has no Files\n", string(pkgID), pkg.PackageName)
			continue
		}

		// if we got here, there's at least one file
		// print the filename and license info for the first 50
		fmt.Printf("============================\n")
		fmt.Printf("Package %s (%s)\n", string(pkgID), pkg.PackageName)
		fmt.Printf("File info (up to first 50):\n")
		i := 1
		for _, f := range pkg.Files {
			// note that these will be in random order, since we're pulling
			// from a map. if we care about order, we should first pull the
			// IDs into a slice, sort it, and then print the ordered files.
			fmt.Printf("- File %d: %s\n", i, f.FileName)
			fmt.Printf("    License from file: %v\n", f.LicenseInfoInFile)
			fmt.Printf("    License concluded: %v\n", f.LicenseConcluded)
			i++
			if i > 50 {
				break
			}
		}
	}
}

func spdxV2_1(doc *spdx.Document2_1) {
	// we can now take a look at its contents via the various data
	// structures representing the SPDX document's sections.

	// print the struct containing the SPDX file's Creation Info section data
	fmt.Printf("==============\n")
	fmt.Printf("Creation info:\n")
	fmt.Printf("==============\n")
	fmt.Printf("%#v\n\n", doc.CreationInfo)
	pkgIDs, err := spdxlib.GetDescribedPackageIDs2_1(doc)
	if err != nil {
		fmt.Printf("Unable to get describe packages from SPDX document: %v\n", err)
		return
	}
	// SPDX Document does contain packages, so we'll go through each one
	for _, pkgID := range pkgIDs {
		pkg, ok := doc.Packages[pkgID]
		if !ok {
			fmt.Printf("Package %s has described relationship but ID not found\n", string(pkgID))
			continue
		}

		// check whether the package had its files analyzed
		if !pkg.FilesAnalyzed {
			fmt.Printf("Package %s (%s) had FilesAnalyzed: false\n", string(pkgID), pkg.PackageName)
			continue
		}

		// also check whether the package has any files present
		if pkg.Files == nil || len(pkg.Files) < 1 {
			fmt.Printf("Package %s (%s) has no Files\n", string(pkgID), pkg.PackageName)
			continue
		}

		// if we got here, there's at least one file
		// print the filename and license info for the first 50
		fmt.Printf("============================\n")
		fmt.Printf("Package %s (%s)\n", string(pkgID), pkg.PackageName)
		fmt.Printf("File info (up to first 50):\n")
		i := 1
		for _, f := range pkg.Files {
			// note that these will be in random order, since we're pulling
			// from a map. if we care about order, we should first pull the
			// IDs into a slice, sort it, and then print the ordered files.
			fmt.Printf("- File %d: %s\n", i, f.FileName)
			fmt.Printf("    License from file: %v\n", f.LicenseInfoInFile)
			fmt.Printf("    License concluded: %v\n", f.LicenseConcluded)
			i++
			if i > 50 {
				break
			}
		}
	}
}
