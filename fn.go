package fn

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"time"
)

type Fn struct {
	Prefix      string
	Postfix     string
	Delimiter   string
	GitShaSize  int
	ProcShaSize int
}

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

func (fn *Fn) Name() string {
	components := []string{
		fn.getFormattedTime(),
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

func (fn *Fn) NameWithFileType(fType string) string {
	return fmt.Sprintf("%s.%s", fn.Name(), fType)
}
