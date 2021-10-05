package v1alpha1

type LOG_FORMAT int

const (
	LogFormatDefault LOG_FORMAT = iota
	LogFormatJSON
)

func (l LOG_FORMAT) String() string {
	return []string{"LogFormatDefault", "LogFormatJSON"}[l]
}
