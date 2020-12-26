package conf

type Options struct {
	LogPath     string
	LogName     string
	LogLevel    string
	MaxSize     int
	MaxAge      int
	Stacktrace  string
	IsStdOut    string
	ProjectName string
}

type Option func(*Options)

func WithLogPath(logPath string) Option {
	return func(o *Options) {
		o.LogPath = logPath
	}
}

func WithLogName(logname string) Option {
	return func(o *Options) {
		o.LogName = logname
	}
}

func WithLogLevel(loglevel string) Option {
	return func(o *Options) {
		o.LogLevel = loglevel
	}
}

func WithMaxSize(maxsize int) Option {
	return func(o *Options) {
		o.MaxAge = maxsize
	}
}

func WithMaxAge(maxage int) Option {
	return func(o *Options) {
		o.MaxAge = maxage
	}
}

func WithStacktrace(stacktrace string) Option {
	return func(o *Options) {
		o.Stacktrace = stacktrace
	}
}

func WithIsStdOut(isstdout string) Option {
	return func(o *Options) {
		o.IsStdOut = isstdout
	}
}

func WithProjectName(projectname string) Option {
	return func(o *Options) {
		o.ProjectName = projectname
	}
}
