package fn

import (
	"crypto/sha256"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Fn in a file namer
type Fn struct {
	Prefix      string
	Postfix     string
	Delimiter   string
	GitShaSize  int
	ProcShaSize int
}

// New creates a new Fn with default settings.
func New() *Fn {
	return &Fn{Delimiter: "-", GitShaSize: 7, ProcShaSize: 8}
}

func (fn *Fn) getFormattedTime() string {
	now := time.Now()
	fmtStr := fmt.Sprintf("20160102%s150405", fn.Delimiter)
	return now.Format(fmtStr)
}

func (fn *Fn) getProcTimeHash() string {
	h := sha256.New()
	content := fmt.Sprintf("%d:%d", time.Now().UnixNano(), os.Getppid())
	h.Write([]byte(content))
	digest := fmt.Sprintf("%x", h.Sum(nil))

	return digest[:fn.ProcShaSize]
}

func getGitHash() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"rev-parse", "--verify", "HEAD"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		return ""
	}
	return string(cmdOut)
}

func (fn *Fn) getGitHash() string {
	return getGitHash()[:fn.GitShaSize]
}

// Name generates a file name (without file extension) based on the current settings
func (fn *Fn) Name() string {
	components := []string{
		fn.getFormattedTime(),
		fn.getGitHash(),
		fn.getProcTimeHash(),
	}
	if fn.Prefix != "" {
		components = append([]string{fn.Prefix}, components...)
	}
	if fn.Postfix != "" {
		components = append(components, fn.Postfix)
	}
	return strings.Join(components, fn.Delimiter)
}

// NameWithFileType is identical to `Name()`, except that it appends a file extension
func (fn *Fn) NameWithFileType(fType string) string {
	return fmt.Sprintf("%s.%s", fn.Name(), fType)
}
