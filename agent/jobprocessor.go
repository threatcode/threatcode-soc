

package agent

import (
  "github.com/threatcode/threatcode-soc/model"
  "io"
  "time"
)

type JobProcessor interface {
  ProcessJob(*model.Job, io.ReadCloser) (io.ReadCloser, error)
  CleanupJob(*model.Job)
  GetDataEpoch() time.Time
}
