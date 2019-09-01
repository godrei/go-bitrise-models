package models

// StepByListID ...
type StepByListID map[string]Step

// Workflow ...
type Workflow struct {
	Title        string                 `json:"title,omitempty" yaml:"title,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	BeforeRun    []string               `json:"before_run,omitempty" yaml:"before_run,omitempty"`
	AfterRun     []string               `json:"after_run,omitempty" yaml:"after_run,omitempty"`
	Environments []Variable             `json:"envs,omitempty" yaml:"envs,omitempty"`
	Steps        []StepByListID         `json:"steps,omitempty" yaml:"steps,omitempty"`
	Meta         map[string]interface{} `json:"meta,omitempty" yaml:"meta,omitempty"`
}

// App ...
type App struct {
	Title        string                 `json:"title,omitempty" yaml:"title,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Environments []Variable             `json:"envs,omitempty" yaml:"envs,omitempty"`
	Meta         map[string]interface{} `json:"meta,omitempty" yaml:"meta,omitempty"`
}

// Trigger ...
type Trigger struct {
	PushBranch              string `json:"push_branch,omitempty" yaml:"push_branch,omitempty"`
	PullRequestSourceBranch string `json:"pull_request_source_branch,omitempty" yaml:"pull_request_source_branch,omitempty"`
	PullRequestTargetBranch string `json:"pull_request_target_branch,omitempty" yaml:"pull_request_target_branch,omitempty"`
	Tag                     string `json:"tag,omitempty" yaml:"tag,omitempty"`
	WorkflowID              string `json:"workflow,omitempty" yaml:"workflow,omitempty"`

	// deprecated
	Pattern              string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	IsPullRequestAllowed bool   `json:"is_pull_request_allowed,omitempty" yaml:"is_pull_request_allowed,omitempty"`
}

// Triggers ...
type Triggers []Trigger

// Config ...
type Config struct {
	FormatVersion        string `json:"format_version" yaml:"format_version"`
	DefaultStepLibSource string `json:"default_step_lib_source,omitempty" yaml:"default_step_lib_source,omitempty"`
	ProjectType          string `json:"project_type" yaml:"project_type"`
	//
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Summary     string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	//
	App        App                 `json:"app,omitempty" yaml:"app,omitempty"`
	TriggerMap Triggers            `json:"trigger_map,omitempty" yaml:"trigger_map,omitempty"`
	Workflows  map[string]Workflow `json:"workflows,omitempty" yaml:"workflows,omitempty"`
}
