package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"github.com/docker/go-units"

	"github.com/fatih/color"
)

const (
	unknown        = "unknown"
	unknownVersion = "0.0.0"
)

const (
	// maximum length of commit
	commitLength = 10
)

// base version's information (overwrite via ldflags)
var (
	gitVersion   = unknownVersion
	gitCommit    = unknown
	gitTreeState = unknown
	buildDate    = unknown
)

// store information about build and version
type Info struct {
	GitVersion   string    `json:"gitVersion"`
	GitCommit    string    `json:"gitCommit"`
	GitTreeState string    `json:"gitTreeState"`
	GitTime      time.Time `json:"gitTime"`
	BuildDate    string    `json:"buildDate"`
	GoVersion    string    `json:"goVersion"`
	Compiler     string    `json:"compiler"`
	Platform     string    `json:"platform"`
	Race         bool      `json:"race"`
}

// Get returns initialized Info object
func Get() Info {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		info = &debug.BuildInfo{}
	}

	var vcsTime time.Time
	for _, v := range info.Settings {
		switch v.Key {
		case "vcs.time":
			var err error
			vcsTime, err = time.Parse(time.RFC3339, v.Value)
			if err != nil {
				vcsTime = time.Time{}
			}
		case "vcs.commit":
			gitCommit = v.Value
		}
	}

	// <system>/<arch>
	var p strings.Builder
	p.WriteString(runtime.GOOS)
	p.WriteRune('/')
	p.WriteString(runtime.GOARCH)

	return Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		GitTime:      vcsTime,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     p.String(),
		Race:         isRace,
	}
}

// Version returns version string
func Version() string {
	return gitVersion
}

// String returns Go-like struture as string
func (i Info) String() string {
	return fmt.Sprintf("%#v", i)
}

// Pretty returns pretty string with version and build information
func (i Info) Pretty() string {
	var b strings.Builder
	// Golang version
	b.WriteString(i.GoVersion)
	b.WriteRune(' ')
	// platform
	b.WriteString(i.Platform)
	if i.Race {
		// race
		b.WriteRune(' ')
		b.WriteString("(race detector enabled)")
	}
	if len(i.GitCommit) > commitLength {
		// SHA1 hash
		b.WriteRune(' ')
		b.WriteString(i.GitCommit[:commitLength])
	}
	if i.GitVersion != unknownVersion {
		// version
		b.WriteRune(' ')
		b.WriteString(i.GitVersion)
	}
	return b.String()
}

// Pretty returns pretty string with version and build information with colors
func (i Info) PrettyColorful() string {
	c := i.GitCommit
	if len(c) > commitLength {
		c = c[:commitLength]
	}

	var x strings.Builder
	// version
	x.WriteString(color.New(color.FgGreen, color.Bold).Sprint(i.GitVersion))
	x.WriteRune(' ')
	// Golang version and platform
	x.WriteString(color.New(color.FgCyan, color.Faint).Sprint(fmt.Sprintf("(%s %s)",
		i.GoVersion, i.Platform)))
	x.WriteRune(' ')
	x.WriteString(color.New(color.FgGreen, color.Faint).Sprint(c))

	if !i.GitTime.IsZero() {
		x.WriteRune(' ')
		x.WriteRune('[')
		x.WriteString(color.New(color.FgMagenta, color.Faint).Sprint(
			units.HumanDuration(time.Since(i.GitTime)), " ago"))
		x.WriteRune(']')
	}

	return x.String()
}
