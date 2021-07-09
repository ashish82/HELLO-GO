package constant

// Environment type represents various kind of environments an application can run in.
type Environment string

const (
	Production  = Environment("prod")
	Development = Environment("dev")
	QA          = Environment("qa")
	Staging     = Environment("staging")
)

const (
	LogFilePath = "/opt/logs/hello-go/"
)

const (
	APPConfigAPI = "APP_CONFIG_DETAIL"
	CommentAPI   = "COMMENT_API"
)

const (
	DT_FRMT_SEARCH_CNTXT = "2006-01-02"
)
