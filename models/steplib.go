package models

// StepGroupInfo ...
type StepGroupInfo struct {
	RemovalDate    string            `json:"removal_date,omitempty" yaml:"removal_date,omitempty"`
	DeprecateNotes string            `json:"deprecate_notes,omitempty" yaml:"deprecate_notes,omitempty"`
	AssetURLs      map[string]string `json:"asset_urls,omitempty" yaml:"asset_urls,omitempty"`
}

// StepGroup ...
type StepGroup struct {
	Info                StepGroupInfo   `json:"info,omitempty" yaml:"info,omitempty"`
	LatestVersionNumber string          `json:"latest_version_number,omitempty" yaml:"latest_version_number,omitempty"`
	Versions            map[string]Step `json:"versions,omitempty" yaml:"versions,omitempty"`
}

// StepGroupByID ...
type StepGroupByID map[string]StepGroup

// StepDownloadLocation ...
type StepDownloadLocation struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

// StepLibrary ...
type StepLibrary struct {
	FormatVersion         string                 `json:"format_version" yaml:"format_version"`
	GeneratedAtTimeStamp  int64                  `json:"generated_at_timestamp" yaml:"generated_at_timestamp"`
	SteplibSource         string                 `json:"steplib_source" yaml:"steplib_source"`
	DownloadLocations     []StepDownloadLocation `json:"download_locations" yaml:"download_locations"`
	AssetsDownloadBaseURI string                 `json:"assets_download_base_uri" yaml:"assets_download_base_uri"`
	Steps                 StepGroupByID          `json:"steps" yaml:"steps"`
}
