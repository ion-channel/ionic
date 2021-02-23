package spdx

import (
	"testing"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/spdx/tools-golang/spdx"
)

const (
	testToken = "token"
)

func TestSPDX(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("create Project from SPDX", func() {
		g.It("should return no error, and a Project if spdx file is valid format 2.1", func() {
			doc := spdx.Document2_1{
				CreationInfo: &spdx.CreationInfo2_1{
					DocumentName:      "some-cool-pkg",
					DocumentNamespace: "http://%v:%v/goodurl",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment:    "some cool package",
				},
			}

			p, err := ProjectFromSPDX2_1(&doc)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(*p.Name).To(Equal(doc.CreationInfo.DocumentName))
			Expect(*p.Type).To(Equal("artifact"))
			Expect(*p.Description).To(Equal(doc.CreationInfo.CreatorComment))
		})

		g.It("should return no error, and a Project if spdx file is valid format 2.2", func() {
			doc := spdx.Document2_2{
				CreationInfo: &spdx.CreationInfo2_2{
					DocumentName:      "some-cool-pkg",
					DocumentNamespace: "http://github.com/my/repo",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment:    "some cool package",
				},
			}

			p, err := ProjectFromSPDX2_2(&doc)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(*p.Name).To(Equal(doc.CreationInfo.DocumentName))
			Expect(*p.Type).To(Equal("git"))
			Expect(*p.Description).To(Equal(doc.CreationInfo.CreatorComment))
		})

		g.It("should return an error if a created Project is invalid due to missing field for 2.2", func() {
			doc := spdx.Document2_2{
				CreationInfo: &spdx.CreationInfo2_2{
					DocumentName:   "some-cool-pkg",
					CreatorPersons: []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment: "some cool package",
				},
			}

			_, err := ProjectFromSPDX2_2(&doc)

			Expect(err).ToNot(BeNil())
			errStr := err.Error()
			Expect(errStr).To(ContainSubstring("DocumentNamespace"))
		})

		g.It("should return an error if a created Project is invalid due to missing field for 2.1", func() {
			doc := spdx.Document2_1{
				CreationInfo: &spdx.CreationInfo2_1{
					DocumentName:   "some-cool-pkg",
					CreatorPersons: []string{"Monsieur Package Creator (mpc@mail.com)"},
					CreatorComment: "some cool package",
				},
			}

			_, err := ProjectFromSPDX2_1(&doc)

			Expect(err).ToNot(BeNil())
			errStr := err.Error()
			Expect(errStr).To(ContainSubstring("DocumentNamespace"))
		})
	})

	g.Describe("create Project from an SPDX package", func() {
		g.It("should return no error if a created Project from package is valid for 2.1", func() {
			pkg := make(map[spdx.ElementID]*spdx.Package2_1)
			packageName := "some-cool-package"
			pkg["id"] = &spdx.Package2_1{
				PackageName:             packageName,
				PackageDownloadLocation: "http://%v:%v/goodurl",
				PackageDescription:      "some cool package",
			}
			doc := spdx.Document2_1{
				CreationInfo: &spdx.CreationInfo2_1{
					DocumentNamespace: "http://%v:%v/goodurl",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
				},
				Packages: pkg,
			}

			p, err := ProjectPackageFromSPDX2_1(&doc, packageName)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(1))
			first := p[0]
			Expect(*first.Name).To(Equal(packageName))
			Expect(*first.Type).To(Equal("artifact"))
			Expect(*first.Description).To(Equal(pkg["id"].PackageDescription))
		})

		g.It("should return no error if a created Project from package is valid for 2.2", func() {
			pkg := make(map[spdx.ElementID]*spdx.Package2_2)
			packageName := "some-cool-package"
			pkg["id"] = &spdx.Package2_2{
				PackageName:             packageName,
				PackageDownloadLocation: "http://github.com/my/repo",
				PackageDescription:      "some cool package",
			}
			doc := spdx.Document2_2{
				CreationInfo: &spdx.CreationInfo2_2{
					DocumentNamespace: "http://%v:%v/goodurl",
					CreatorPersons:    []string{"Monsieur Package Creator (mpc@mail.com)"},
				},
				Packages: pkg,
			}

			p, err := ProjectPackageFromSPDX2_2(&doc, packageName)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(1))
			first := p[0]
			Expect(*first.Name).To(Equal(packageName))
			Expect(*first.Type).To(Equal("git"))
			Expect(*first.Description).To(Equal(pkg["id"].PackageDescription))
		})

		g.It("should an no error if a package is invalid (due to missing download location) for 2.2", func() {
			pkg := make(map[spdx.ElementID]*spdx.Package2_2)
			packageName := "some-cool-package"
			pkg["id"] = &spdx.Package2_2{
				PackageName:             packageName,
				PackageDownloadLocation: "",
				PackageDescription:      "some cool package",
			}
			doc := spdx.Document2_2{
				CreationInfo: &spdx.CreationInfo2_2{},
				Packages:     pkg,
			}

			p, err := ProjectPackageFromSPDX2_2(&doc, packageName)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(1))
			first := p[0]
			Expect(*first.Name).To(Equal(packageName))
			Expect(*first.Type).To(Equal("source_unavailable"))
		})

		g.It("should an no error if a package is invalid (due to missing download location) for 2.1", func() {
			pkg := make(map[spdx.ElementID]*spdx.Package2_1)
			packageName := "some-cool-package"
			pkg["id"] = &spdx.Package2_1{
				PackageName:        packageName,
				PackageDescription: "some cool package",
			}
			doc := spdx.Document2_1{
				CreationInfo: &spdx.CreationInfo2_1{},
				Packages:     pkg,
			}

			p, err := ProjectPackageFromSPDX2_1(&doc, packageName)

			Expect(err).To(BeNil())
			Expect(p).NotTo(BeNil())
			Expect(len(p)).To(Equal(1))
			first := p[0]
			Expect(*first.Name).To(Equal(packageName))
			Expect(*first.Type).To(Equal("source_unavailable"))
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
