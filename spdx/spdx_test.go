package spdx

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/spdx/tools-golang/spdx"
)

const (
	testToken = "token"
)

func TestSPDX(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("SPDX v2.1", func() {
		g.It("should return the top-level project when no dependencies requested", func() {
			spdxDocumentRef := spdx.MakeDocElementID("", "DOCUMENT")
			spdxRef := spdx.MakeDocElementID("", "some-cool-pkg")
			spdxDependencyRef := spdx.MakeDocElementID("", "some-dep")
			spdxPackage := spdx.Package2_1{
				PackageName:                 "some-cool-pkg",
				PackageSPDXIdentifier:       spdxRef.ElementRefID,
				PackageVersion:              "1.2.3",
				PackageSupplierOrganization: "The Org",
				PackageDownloadLocation:     "https://github.com/some-org/some-cool-pkg.git@main",
				PackageDescription:          "Some description",
			}
			spdxDependencyPackage := spdx.Package2_1{
				PackageName:                 "some-dep",
				PackageSPDXIdentifier:       spdxDependencyRef.ElementRefID,
				PackageVersion:              "3.2.1",
				PackageSupplierOrganization: "The Org",
				PackageDownloadLocation:     "https://github.com/some-org/some-dep.git",
				PackageDescription:          "Some dep description",
			}
			packages := make(map[spdx.ElementID]*spdx.Package2_1)
			packages[spdxRef.ElementRefID] = &spdxPackage
			packages[spdxDependencyRef.ElementRefID] = &spdxDependencyPackage

			relationships := []*spdx.Relationship2_1{{
				RefA:         spdxDocumentRef,
				Relationship: "DESCRIBES",
				RefB:         spdxRef,
			}, {
				RefA:         spdxRef,
				Relationship: "DEPENDS_ON",
				RefB:         spdxDependencyRef,
			}}

			doc := spdx.Document2_1{
				CreationInfo: &spdx.CreationInfo2_1{
					DocumentName:      "SPDX SBOM",
					DocumentNamespace: "http://ionchannel.io",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment:    "some cool package SBOM",
				},
				Packages:      packages,
				Relationships: relationships,
			}

			p, err := ProjectsFromSPDX(&doc, false)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(1))
			Expect(*p[0].Name).To(Equal(spdxPackage.PackageName))
			Expect(*p[0].Type).To(Equal("git"))
			Expect(*p[0].Description).To(Equal(spdxPackage.PackageDescription))
			Expect(len(p[0].Aliases)).To(Equal(1))
			Expect(p[0].Aliases[0].Version).To(Equal(spdxPackage.PackageVersion))
			Expect(*p[0].Branch).To(Equal("main"))
			Expect(*p[0].Source).To(Equal("https://github.com/some-org/some-cool-pkg.git"))
		})

		g.It("should return all projects when dependencies requested", func() {
			spdxDocumentRef := spdx.MakeDocElementID("", "DOCUMENT")
			spdxRef := spdx.MakeDocElementID("", "some-cool-pkg")
			spdxDependencyRef := spdx.MakeDocElementID("", "some-dep")
			spdxPackage := spdx.Package2_1{
				PackageName:                 "some-cool-pkg",
				PackageSPDXIdentifier:       spdxRef.ElementRefID,
				PackageVersion:              "1.2.3",
				PackageSupplierOrganization: "The Org",
				PackageDownloadLocation:     "https://github.com/some-org/some-cool-pkg.git@my-cool-branch",
				PackageDescription:          "Some description",
			}
			spdxDependencyPackage := spdx.Package2_1{
				PackageName:                 "some-dep",
				PackageSPDXIdentifier:       spdxDependencyRef.ElementRefID,
				PackageVersion:              "3.2.1",
				PackageSupplierOrganization: "The Org",
				PackageDescription:          "Some dep description",
			}
			packages := make(map[spdx.ElementID]*spdx.Package2_1)
			packages[spdxRef.ElementRefID] = &spdxPackage
			packages[spdxDependencyRef.ElementRefID] = &spdxDependencyPackage

			relationships := []*spdx.Relationship2_1{{
				RefA:         spdxDocumentRef,
				Relationship: "DESCRIBES",
				RefB:         spdxRef,
			}, {
				RefA:         spdxRef,
				Relationship: "DEPENDS_ON",
				RefB:         spdxDependencyRef,
			}}

			doc := spdx.Document2_1{
				CreationInfo: &spdx.CreationInfo2_1{
					DocumentName:      "SPDX SBOM",
					DocumentNamespace: "http://ionchannel.io",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment:    "some cool package SBOM",
				},
				Packages:      packages,
				Relationships: relationships,
			}

			p, err := ProjectsFromSPDX(&doc, true)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(2))
			Expect(*p[0].Name).To(Equal(spdxPackage.PackageName))
			Expect(*p[0].Type).To(Equal("git"))
			Expect(*p[0].Description).To(Equal(spdxPackage.PackageDescription))
			Expect(len(p[0].Aliases)).To(Equal(1))
			Expect(p[0].Aliases[0].Version).To(Equal(spdxPackage.PackageVersion))
			Expect(p[0].Aliases[0].Org).To(Equal(spdxPackage.PackageSupplierOrganization))
			Expect(*p[0].Branch).To(Equal("my-cool-branch"))
			Expect(*p[1].Name).To(Equal(spdxDependencyPackage.PackageName))
			Expect(*p[1].Type).To(Equal("source_unavailable"))
			Expect(*p[1].Description).To(Equal(spdxDependencyPackage.PackageDescription))
			Expect(len(p[1].Aliases)).To(Equal(1))
			Expect(p[1].Aliases[0].Version).To(Equal(spdxDependencyPackage.PackageVersion))
			Expect(p[1].Aliases[0].Org).To(Equal(spdxPackage.PackageSupplierOrganization))
		})
	})
	g.Describe("SPDX v2.2", func() {
		g.It("should return the top-level project when no dependencies requested", func() {
			spdxDocumentRef := spdx.MakeDocElementID("", "DOCUMENT")
			spdxRef := spdx.MakeDocElementID("", "some-cool-pkg")
			spdxDependencyRef := spdx.MakeDocElementID("", "some-dep")
			spdxPackage := spdx.Package2_2{
				PackageName:                 "some-cool-pkg",
				PackageSPDXIdentifier:       spdxRef.ElementRefID,
				PackageVersion:              "1.2.3",
				PackageSupplierOrganization: "The Org",
				PackageDownloadLocation:     "https://github.com/some-org/some-cool-pkg.git@ian/some-branch",
				PackageDescription:          "Some description",
			}
			spdxDependencyPackage := spdx.Package2_2{
				PackageName:                 "some-dep",
				PackageSPDXIdentifier:       spdxDependencyRef.ElementRefID,
				PackageVersion:              "3.2.1",
				PackageSupplierOrganization: "The Org",
				PackageDownloadLocation:     "https://github.com/some-org/some-dep.git",
				PackageDescription:          "Some dep description",
			}
			packages := make(map[spdx.ElementID]*spdx.Package2_2)
			packages[spdxRef.ElementRefID] = &spdxPackage
			packages[spdxDependencyRef.ElementRefID] = &spdxDependencyPackage

			relationships := []*spdx.Relationship2_2{{
				RefA:         spdxDocumentRef,
				Relationship: "DESCRIBES",
				RefB:         spdxRef,
			}, {
				RefA:         spdxRef,
				Relationship: "DEPENDS_ON",
				RefB:         spdxDependencyRef,
			}}

			doc := spdx.Document2_2{
				CreationInfo: &spdx.CreationInfo2_2{
					DocumentName:      "SPDX SBOM",
					DocumentNamespace: "http://ionchannel.io",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment:    "some cool package SBOM",
				},
				Packages:      packages,
				Relationships: relationships,
			}

			p, err := ProjectsFromSPDX(&doc, false)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(1))
			Expect(*p[0].Name).To(Equal(spdxPackage.PackageName))
			Expect(*p[0].Type).To(Equal("git"))
			Expect(*p[0].Description).To(Equal(spdxPackage.PackageDescription))
			Expect(len(p[0].Aliases)).To(Equal(1))
			Expect(p[0].Aliases[0].Version).To(Equal(spdxPackage.PackageVersion))
			Expect(*p[0].Branch).To(Equal("ian/some-branch"))
		})

		g.It("should return all projects when dependencies requested", func() {
			spdxDocumentRef := spdx.MakeDocElementID("", "DOCUMENT")
			spdxRef := spdx.MakeDocElementID("", "some-cool-pkg")
			spdxDependencyRef := spdx.MakeDocElementID("", "some-dep")
			spdxPackage := spdx.Package2_2{
				PackageName:                 "some-cool-pkg",
				PackageSPDXIdentifier:       spdxRef.ElementRefID,
				PackageVersion:              "1.2.3",
				PackageSupplierOrganization: "The Org",
				PackageDownloadLocation:     "https://github.com/some-org/some-cool-pkg.git",
				PackageDescription:          "Some description",
			}
			spdxDependencyPackage := spdx.Package2_2{
				PackageName:                 "some-dep",
				PackageSPDXIdentifier:       spdxDependencyRef.ElementRefID,
				PackageVersion:              "3.2.1",
				PackageSupplierOrganization: "The Org",
				PackageDescription:          "Some dep description",
			}
			packages := make(map[spdx.ElementID]*spdx.Package2_2)
			packages[spdxRef.ElementRefID] = &spdxPackage
			packages[spdxDependencyRef.ElementRefID] = &spdxDependencyPackage

			relationships := []*spdx.Relationship2_2{{
				RefA:         spdxDocumentRef,
				Relationship: "DESCRIBES",
				RefB:         spdxRef,
			}, {
				RefA:         spdxRef,
				Relationship: "DEPENDS_ON",
				RefB:         spdxDependencyRef,
			}}

			doc := spdx.Document2_2{
				CreationInfo: &spdx.CreationInfo2_2{
					DocumentName:      "SPDX SBOM",
					DocumentNamespace: "http://ionchannel.io",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment:    "some cool package SBOM",
				},
				Packages:      packages,
				Relationships: relationships,
			}

			p, err := ProjectsFromSPDX(&doc, true)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(2))
			Expect(*p[0].Name).To(Equal(spdxPackage.PackageName))
			Expect(*p[0].Type).To(Equal("git"))
			Expect(*p[0].Description).To(Equal(spdxPackage.PackageDescription))
			Expect(len(p[0].Aliases)).To(Equal(1))
			Expect(p[0].Aliases[0].Version).To(Equal(spdxPackage.PackageVersion))
			Expect(p[0].Aliases[0].Org).To(Equal(spdxPackage.PackageSupplierOrganization))
			Expect(*p[0].Branch).To(Equal("HEAD"))
			Expect(*p[1].Name).To(Equal(spdxDependencyPackage.PackageName))
			Expect(*p[1].Type).To(Equal("source_unavailable"))
			Expect(*p[1].Description).To(Equal(spdxDependencyPackage.PackageDescription))
			Expect(len(p[1].Aliases)).To(Equal(1))
			Expect(p[1].Aliases[0].Version).To(Equal(spdxDependencyPackage.PackageVersion))
			Expect(p[1].Aliases[0].Org).To(Equal(spdxPackage.PackageSupplierOrganization))
			Expect(*p[1].Branch).To(Equal(""))
		})
	})

	g.Describe("parse emails from SPDX creator information", func() {
		g.It("should return an email if present", func() {
			creatorInfo := "My Name (myemail@mail.net)"
			email := parseCreatorEmail([]string{creatorInfo})

			Expect(email).ToNot(BeNil())
			Expect(email).To(Equal("myemail@mail.net"))
		})

		g.It("should handle a missing email", func() {
			creatorInfo := "My Name"

			email := parseCreatorEmail([]string{creatorInfo})
			Expect(email).ToNot(BeNil())
			Expect(email).To(Equal(""))

			email = parseCreatorEmail([]string{})
			Expect(email).ToNot(BeNil())
			Expect(email).To(Equal(""))
		})
	})
}
