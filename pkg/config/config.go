package config

type CommonConfig struct {
	// SubscriptionId specifies the user's Azure subscription id.
	SubscriptionId string
	// OutputDir specifies the Terraform working directory for aztfy to import resources and generate TF configs.
	OutputDir string
	// Append specifies whether this run is in append mode, in which case aztfy will generate some "safe" file name to avoid conflicts to usre's existing files.
	Append bool
	// DevProvider specifies whether to use a development provider built locally rather than using a version pinned provider from official Terraform registry.
	DevProvider bool
	// ContinueOnError specifies whether continue the progress even hit an import error.
	ContinueOnError bool
	// BackendType specifies the Terraform backend type.
	BackendType string
	// BackendConfig specifies an array of Terraform backend configs.
	BackendConfig []string
	// FullConfig specifies whether to export all (non computed-only) Terarform properties when generating TF configs.
	FullConfig bool
	// Parallelism specifies the parallelism for the process
	Parallelism int
	// ModulePath specifies the path of the module (e.g. "module1.module2") where the resources will be imported and config generated.
	// Note that only modules whose "source" is local path is supported. By default, it is the root module.
	ModulePath string
	// HCLOnly is a strange field, which is only used internally by aztfy to indicate whether to remove other files other than TF config at the end.
	// External Go modules shoudl just ignore it.
	HCLOnly bool
}

type Config struct {
	CommonConfig

	// Exactly one of below is specified

	// ResourceId specifies the Azure resource id, this indicates the resource mode.
	ResourceId string
	// ResourceGroupName specifies the name of the resource group, this indicates the resource group mode.
	ResourceGroupName string
	// ARGPredicate specifies the ARG where predicate, this indicates the query mode.
	ARGPredicate string
	// MappingFile specifies the path of mapping file, this indicates the map file mode.
	MappingFile string

	// ResourceNamePattern specifies the resource name pattern, this only applies to resource group mode and query mode.
	ResourceNamePattern string

	// RecursiveQuery specifies whether to recursively list the child/proxy resources of the ARG resulted resource list, this only applies to query mode.
	RecursiveQuery bool

	// TFResourceName specifies the TF resource name, this only applies to resource mode.
	TFResourceName string
	// TFResourceName specifies the TF resource type (if empty, aztfy will deduce the type), this only applies to resource mode.
	TFResourceType string
}
