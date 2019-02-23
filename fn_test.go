package fn

import (
	"os"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
)

func setupTest() {
	testTime := time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
	monkey.Patch(time.Now, func() time.Time { return testTime })
	monkey.Patch(os.Getppid, func() int { return 12345 })
	monkey.Patch(getGitHash, func() string { return "711de72aacdd0fda1c95486f9e67e477df2cee9a" })
}

func TestName(t *testing.T) {
	setupTest()

	fn := New()
	assert.Equal(t, "19740519-010203-711de72-1309178b", fn.Name())
}

func TestPrefix(t *testing.T) {
	setupTest()

	fn := New()
	fn.Prefix = "foo"
	assert.Equal(t, "foo-19740519-010203-711de72-1309178b", fn.Name())
}

func TestPostfix(t *testing.T) {
	setupTest()

	fn := New()
	fn.Postfix = "bar"
	assert.Equal(t, "19740519-010203-711de72-1309178b-bar", fn.Name())
}

func TestNameWithFileType(t *testing.T) {
	setupTest()

	fn := New()
	assert.Equal(t, "19740519-010203-711de72-1309178b.png", fn.NameWithFileType("png"))
}
