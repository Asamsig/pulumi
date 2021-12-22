// Copyright 2016-2019, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package display

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/pulumi/pulumi/pkg/v3/engine/events"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// We use RFC 5424 timestamps with millisecond precision for displaying time stamps on watch
// entries. Go does not pre-define a format string for this format, though it is similar to
// time.RFC3339Nano.
//
// See https://tools.ietf.org/html/rfc5424#section-6.2.3.
const timeFormat = "15:04:05.000"

// ShowWatchEvents renders incoming engine events for display in Watch Mode.
func ShowWatchEvents(op string, action apitype.UpdateKind, eventsC <-chan events.Event, done chan<- bool, opts Options) {
	// Ensure we close the done channel before exiting.
	defer func() { close(done) }()
	for e := range eventsC {
		// In the event of cancelation, break out of the loop immediately.
		if e.Type == events.CancelEvent {
			break
		}

		// For all other events, use the payload to build up the JSON digest we'll emit later.
		switch e.Type {
		// Events occurring early:
		case events.PreludeEvent, events.SummaryEvent, events.StdoutColorEvent:
			// Ignore it
			continue
		case events.PolicyViolationEvent:
			// At this point in time, we don't handle policy events as part of pulumi watch
			continue
		case events.DiagEvent:
			// Skip any ephemeral or debug messages, and elide all colorization.
			p := e.Payload().(events.DiagEventPayload)
			resourceName := ""
			if p.URN != "" {
				resourceName = string(p.URN.Name())
			}
			PrintfWithWatchPrefix(time.Now(), resourceName,
				"%s", renderDiffDiagEvent(p, opts))
		case events.ResourcePreEvent:
			p := e.Payload().(events.ResourcePreEventPayload)
			if shouldShow(p.Metadata, opts) {
				PrintfWithWatchPrefix(time.Now(), string(p.Metadata.URN.Name()),
					"%s %s\n", p.Metadata.Op, p.Metadata.URN.Type())
			}
		case events.ResourceOutputsEvent:
			p := e.Payload().(events.ResourceOutputsEventPayload)
			if shouldShow(p.Metadata, opts) {
				PrintfWithWatchPrefix(time.Now(), string(p.Metadata.URN.Name()),
					"done %s %s\n", p.Metadata.Op, p.Metadata.URN.Type())
			}
		case events.ResourceOperationFailed:
			p := e.Payload().(events.ResourceOperationFailedPayload)
			if shouldShow(p.Metadata, opts) {
				PrintfWithWatchPrefix(time.Now(), string(p.Metadata.URN.Name()),
					"failed %s %s\n", p.Metadata.Op, p.Metadata.URN.Type())
			}
		default:
			contract.Failf("unknown event type '%s'", e.Type)
		}
	}
}

// Watch output is written from multiple concurrent goroutines.  For now we synchronize Printfs to
// the watch output stream as a simple way to avoid garbled output.
var watchPrintfMutex sync.Mutex

// PrintfWithWatchPrefix wraps fmt.Printf with a watch mode prefixer that adds a timestamp and
// resource metadata.
func PrintfWithWatchPrefix(t time.Time, resourceName string, format string, a ...interface{}) {
	watchPrintfMutex.Lock()
	defer watchPrintfMutex.Unlock()
	prefix := fmt.Sprintf("%12.12s[%20.20s] ", t.Format(timeFormat), resourceName)
	out := &prefixer{os.Stdout, []byte(prefix)}
	_, err := fmt.Fprintf(out, format, a...)
	contract.IgnoreError(err)
}

type prefixer struct {
	writer io.Writer
	prefix []byte
}

var _ io.Writer = (*prefixer)(nil)

func (prefixer *prefixer) Write(p []byte) (int, error) {
	n := 0
	lines := bytes.SplitAfter(p, []byte{'\n'})
	for _, line := range lines {
		// If p ends with a newline, we may see an "" as the last element of lines, which we will skip.
		if len(line) == 0 {
			continue
		}
		_, err := prefixer.writer.Write(prefixer.prefix)
		if err != nil {
			return n, err
		}
		m, err := prefixer.writer.Write(line)
		n += m
		if err != nil {
			return n, err
		}
	}
	return n, nil
}
