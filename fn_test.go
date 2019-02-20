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
}

func TestName(t *testing.T) {
	setupTest()

	fn := New()
	assert.Equal(t, "190560519-010203-1309178b", fn.Name())
}

func TestPrefix(t *testing.T) {
	setupTest()

	fn := New()
	fn.Prefix = "foo"
	assert.Equal(t, "foo-190560519-010203-1309178b", fn.Name())
}

func TestPostfix(t *testing.T) {
	setupTest()

	fn := New()
	fn.Postfix = "bar"
	assert.Equal(t, "190560519-010203-1309178b-bar", fn.Name())
}

func TestNameWithFileType(t *testing.T) {
	setupTest()

	fn := New()
	assert.Equal(t, "190560519-010203-1309178b.png", fn.NameWithFileType("png"))
}
