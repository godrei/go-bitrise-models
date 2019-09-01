package models

import (
	"time"
)

// StepSource ...
type StepSource struct {
	Git    string `json:"git,omitempty" yaml:"git,omitempty"`
	Commit string `json:"commit,omitempty" yaml:"commit,omitempty"`
}

// StepDependency ...
type StepDependency struct {
	Manager string `json:"manager,omitempty" yaml:"manager,omitempty"`
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
}

// BrewDep ...
type BrewDep struct {
	// Name is the package name for Brew
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// BinName is the binary's name, if it doesn't match the package's name.
	// Can be used for e.g. calling `which`.
	// E.g. in case of "AWS CLI" the package is `awscli` and the binary is `aws`.
	// If BinName is empty Name will be used as BinName too.
	BinName string `json:"bin_name,omitempty" yaml:"bin_name,omitempty"`
}

// AptGetDep ...
type AptGetDep struct {
	// Name is the package name for Apt-get
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// BinName is the binary's name, if it doesn't match the package's name.
	// Can be used for e.g. calling `which`.
	// E.g. in case of "AWS CLI" the package is `awscli` and the binary is `aws`.
	// If BinName is empty Name will be used as BinName too.
	BinName string `json:"bin_name,omitempty" yaml:"bin_name,omitempty"`
}

// CheckOnlyDep ...
type CheckOnlyDep struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Deps ...
type Deps struct {
	Brew      []BrewDep      `json:"brew,omitempty" yaml:"brew,omitempty"`
	AptGet    []AptGetDep    `json:"apt_get,omitempty" yaml:"apt_get,omitempty"`
	CheckOnly []CheckOnlyDep `json:"check_only,omitempty" yaml:"check_only,omitempty"`
}

// BashToolkit ...
type BashToolkit struct {
	EntryFile string `json:"entry_file,omitempty" yaml:"entry_file,omitempty"`
}

// GoToolkit ...
type GoToolkit struct {
	// PackageName - required
	PackageName string `json:"package_name" yaml:"package_name"`
}

// StepToolkit ...
type StepToolkit struct {
	Bash *BashToolkit `json:"bash,omitempty" yaml:"bash,omitempty"`
	Go   *GoToolkit   `json:"go,omitempty" yaml:"go,omitempty"`
}

// Step ...
type Step struct {
	Title       *string `json:"title,omitempty" yaml:"title,omitempty"`
	Summary     *string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	//
	Website       *string `json:"website,omitempty" yaml:"website,omitempty"`
	SourceCodeURL *string `json:"source_code_url,omitempty" yaml:"source_code_url,omitempty"`
	SupportURL    *string `json:"support_url,omitempty" yaml:"support_url,omitempty"`
	// auto-generated at share
	PublishedAt *time.Time        `json:"published_at,omitempty" yaml:"published_at,omitempty"`
	Source      *StepSource       `json:"source,omitempty" yaml:"source,omitempty"`
	AssetURLs   map[string]string `json:"asset_urls,omitempty" yaml:"asset_urls,omitempty"`
	//
	HostOsTags          []string         `json:"host_os_tags,omitempty" yaml:"host_os_tags,omitempty"`
	ProjectTypeTags     []string         `json:"project_type_tags,omitempty" yaml:"project_type_tags,omitempty"`
	TypeTags            []string         `json:"type_tags,omitempty" yaml:"type_tags,omitempty"`
	Dependencies        []StepDependency `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Toolkit             *StepToolkit     `json:"toolkit,omitempty" yaml:"toolkit,omitempty"`
	Deps                *Deps            `json:"deps,omitempty" yaml:"deps,omitempty"`
	IsRequiresAdminUser *bool            `json:"is_requires_admin_user,omitempty" yaml:"is_requires_admin_user,omitempty"`
	// IsAlwaysRun : if true then this step will always run,
	//  even if a previous step fails.
	IsAlwaysRun *bool `json:"is_always_run,omitempty" yaml:"is_always_run,omitempty"`
	// IsSkippable : if true and this step fails the build will still continue.
	//  If false then the build will be marked as failed and only those
	//  steps will run which are marked with IsAlwaysRun.
	IsSkippable *bool `json:"is_skippable,omitempty" yaml:"is_skippable,omitempty"`
	// RunIf : only run the step if the template example evaluates to true
	RunIf   *string                `json:"run_if,omitempty" yaml:"run_if,omitempty"`
	Timeout *int                   `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty" yaml:"meta,omitempty"`
	//
	Inputs  []Variable `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs []Variable `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}
