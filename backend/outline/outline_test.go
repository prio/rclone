// Test Outline filesystem interface
package outline_test

import (
	"testing"

	"github.com/rclone/rclone/backend/outline"
	"github.com/rclone/rclone/fstest"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	if *fstest.RemoteName == "" {
		*fstest.RemoteName = "TestOutline:"
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: *fstest.RemoteName,
		NilObject:  (*outline.Object)(nil),
	})
}
