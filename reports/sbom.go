package reports

import "net/url"

const (
	// ReportGetSBOMEndpoint is the endpoint for generating SBOMs
	ReportGetSBOMEndpoint = "v1/report/getSBOM"
)

// SBOMFormat is a string enum for the accepted SBOM formats that we can export
type SBOMFormat string

const (
	// SBOMFormatSPDX is the enum value for the SPDX SBOM format
	SBOMFormatSPDX SBOMFormat = "SPDX"
	// SBOMFormatCycloneDX is the enum value for the CycloneDX SBOM format
	SBOMFormatCycloneDX SBOMFormat = "CycloneDX"
)

// SBOMExportOptions represents all of the different settings a user can specify for how the SBOM is exported.
type SBOMExportOptions struct {
	Format SBOMFormat
}

// Params converts an SBOMExportOptions object into a URL param object for use in making an API request
func (options SBOMExportOptions) Params() *url.Values {
	params := &url.Values{}
	params.Set("sbom_type", string(options.Format))

	return params
}