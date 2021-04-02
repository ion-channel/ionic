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

var (
	spdxVersion string
	rulesetID   string
	packageName string
	teamID      string
)

// ProjectFromSPDX2_2 assumes parsing on individual package with package source at DocumentNameSpace, from a top level SPDX file (v2.2)
// (DocumentNameSpace must be a resolvable URL)
func ProjectFromSPDX2_2(doc *spdx.Document2_2) (projects.Project, error) {
	p := projects.Project{}
	// need DocumentNameSpace for source
	if doc.CreationInfo.DocumentNamespace == "" {
		err := fmt.Errorf("Error while creating project SPDX file %s must contain DocumentNamespace field", doc.CreationInfo.DocumentName)
		return p, err
	}

	// create our project
	var name, ptype, source, branch, description string
	active := true
	monitor := true

	tmpID := uuid.New().String()
	name = doc.CreationInfo.DocumentName
	if strings.Contains(doc.CreationInfo.DocumentNamespace, "git") {
		source = doc.CreationInfo.DocumentNamespace
		branch = "HEAD" // use the remote's default branch
		ptype = "git"
	} else {
		source = doc.CreationInfo.DocumentNamespace
		ptype = "artifact"
	}

	if doc.CreationInfo.CreatorComment != "" {
		description = doc.CreationInfo.CreatorComment
	}

	proj := projects.Project{
		ID:          &tmpID,
		Branch:      &branch,
		Description: &description,
		TeamID:      &teamID,
		Type:        &ptype,
		Source:      &source,
		Name:        &name,
		RulesetID:   &rulesetID,
		Active:      active,
		Monitor:     monitor,
	}

	return proj, nil

}

// ProjectFromSPDX2_1 assumes parsing on individual package with package source at DocumentNamespace, from a top level SPDX file (v2.1)
// (DocumentNamespace must be a resolvable URL
func ProjectFromSPDX2_1(doc *spdx.Document2_1) (projects.Project, error) {
	p := projects.Project{}
	// need DocumentNamespace for source
	if doc.CreationInfo.DocumentNamespace == "" {
		err := fmt.Errorf("Error while creating project SPDX file %s must contain DocumentNamespace field", doc.CreationInfo.DocumentName)
		return p, err
	}

	// create our project
	var name, ptype, source, branch, description string
	active := true
	monitor := true

	tmpID := uuid.New().String()
	name = doc.CreationInfo.DocumentName
	if strings.Contains(doc.CreationInfo.DocumentNamespace, "git") {
		source = doc.CreationInfo.DocumentNamespace
		branch = "HEAD" // use the remote's default branch
		ptype = "git"
	} else {
		source = doc.CreationInfo.DocumentNamespace
		ptype = "artifact"
	}

	if doc.CreationInfo.CreatorComment != "" {
		description = doc.CreationInfo.CreatorComment
	}

	proj := projects.Project{
		ID:          &tmpID,
		Branch:      &branch,
		Description: &description,
		TeamID:      &teamID,
		Type:        &ptype,
		Source:      &source,
		Name:        &name,
		RulesetID:   &rulesetID,
		Active:      active,
		Monitor:     monitor,
	}

	return proj, nil
}

// ProjectPackageFromSPDX2_2 parses packages from an SPDX file (v2.2) with package source at PackageDownloadLocation
// PackageDownloadLocation must be a resolveable URL to create a project
func ProjectPackageFromSPDX2_2(doc *spdx.Document2_2, packageName string) ([]projects.Project, error) {
	projs := make([]projects.Project, 0)
	pkgIDs, err := spdxlib.GetDescribedPackageIDs2_2(doc)
	if err != nil {
		return projs, fmt.Errorf("unable to get describe packages from SPDX document: %v", err)
	}

	// SPDX Document does contain packages, so we'll go through each one
	for _, pkgID := range pkgIDs {
		pkg, ok := doc.Packages[pkgID]
		if !ok {
			fmt.Printf("Package %s has described relationship but ID not found - skipping package\n", string(pkgID))
			continue
		}

		// create our project
		var name, ptype, source, branch, description string
		active := true
		monitor := true

		tmpID := uuid.New().String()
		name = pkg.PackageName

		if pkg.PackageDownloadLocation == "" || pkg.PackageDownloadLocation == "NOASSERTION" {
			ptype = "source_unavailable"
		} else if strings.Contains(pkg.PackageDownloadLocation, "git") {
			source = pkg.PackageDownloadLocation
			branch = "HEAD" // use the remote's default branch
			ptype = "git"
		} else {
			source = pkg.PackageDownloadLocation
			ptype = "artifact"
		}

		if pkg.PackageDescription != "" {
			description = pkg.PackageDescription
		}

		proj := projects.Project{
			ID:          &tmpID,
			Branch:      &branch,
			Description: &description,
			TeamID:      &teamID,
			Type:        &ptype,
			Source:      &source,
			Name:        &name,
			RulesetID:   &rulesetID,
			Active:      active,
			Monitor:     monitor,
		}

		pVersion := ""
		if pkg.PackageVersion != "" {
			pVersion = pkg.PackageVersion
		}

		pOrg := ""
		if pkg.PackageSupplierOrganization != "" {
			pOrg = pkg.PackageSupplierOrganization
		}

		pName := ""
		if pkg.PackageName != "" {
			pName = pkg.PackageName
		}

		// check if any of pVersion, pOrg, pName are not empty string
		if len(pVersion) > 0 || len(pOrg) > 0 || len(pName) > 0 {
			alias := aliases.Alias{
				Name:    pName,
				Org:     pOrg,
				Version: pVersion,
			}
			proj.Aliases = []aliases.Alias{alias}
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

// ProjectPackageFromSPDX2_1 parses packages from an SPDX file (v2.1) with package source at PackageDownloadLocation
// PackageDownloadLocation must be a resolveable URL to create a project
func ProjectPackageFromSPDX2_1(doc *spdx.Document2_1, packageName string) ([]projects.Project, error) {
	projs := make([]projects.Project, 0)

	pkgIDs, err := spdxlib.GetDescribedPackageIDs2_1(doc)
	if err != nil {
		return projs, fmt.Errorf("unable to get describe packages from SPDX document: %v", err)
	}

	// SPDX Document does contain packages, so we'll go through each one
	for _, pkgID := range pkgIDs {
		pkg, ok := doc.Packages[pkgID]
		if !ok {
			fmt.Printf("Package %s has described relationship but ID not found - skipping package\n", string(pkgID))
			continue
		}

		var name, ptype, source, branch, description string
		active := true
		monitor := true

		tmpID := uuid.New().String()
		name = pkg.PackageName

		if pkg.PackageDownloadLocation == "" || pkg.PackageDownloadLocation == "NOASSERTION" {
			ptype = "source_unavailable"
		} else if strings.Contains(pkg.PackageDownloadLocation, "git") {
			source = pkg.PackageDownloadLocation
			branch = "HEAD" // use the remote's default branch
			ptype = "git"
		} else {
			source = pkg.PackageDownloadLocation
			ptype = "artifact"
		}

		if pkg.PackageDescription != "" {
			description = pkg.PackageDescription
		}

		proj := projects.Project{
			ID:          &tmpID,
			Branch:      &branch,
			Description: &description,
			TeamID:      &teamID,
			Type:        &ptype,
			Source:      &source,
			Name:        &name,
			RulesetID:   &rulesetID,
			Active:      active,
			Monitor:     monitor,
		}

		pVersion := ""
		if pkg.PackageVersion != "" {
			pVersion = pkg.PackageVersion
		}

		pOrg := ""
		if pkg.PackageSupplierOrganization != "" {
			pOrg = pkg.PackageSupplierOrganization
		}

		pName := ""
		if pkg.PackageName != "" {
			pName = pkg.PackageName
		}

		// check if any of pVersion, pOrg, pName are not empty string
		if len(pVersion) > 0 || len(pOrg) > 0 || len(pName) > 0 {
			alias := aliases.Alias{
				Name:    pName,
				Org:     pOrg,
				Version: pVersion,
			}
			proj.Aliases = []aliases.Alias{alias}
		}

		projs = append(projs, proj)

	}

	return projs, nil
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
