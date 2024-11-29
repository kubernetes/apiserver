/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"os"
	"runtime"

	"k8s.io/klog/v2"
)

// RequestShutdown emulates a received event that is considered as shutdown signal (SIGTERM/SIGINT)
// This returns whether a handler was notified
func RequestShutdown() bool {
	if shutdownHandler != nil {
		select {
		case shutdownHandler <- shutdownSignals[0]:
			return true
		default:
		}
	}

	return false
}

func dumpStacks(writeToFile bool, path string) {
	var (
		buf       []byte
		stackSize int
	)
	bufferLen := 16384
	for stackSize == len(buf) {
		buf = make([]byte, bufferLen)
		stackSize = runtime.Stack(buf, true)
		bufferLen *= 2
	}
	buf = buf[:stackSize]
	klog.InfoS("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)

	if writeToFile {
		f, err := os.Create(path)
		if err != nil {
			return
		}
		defer f.Close()
		f.WriteString(string(buf))
		klog.InfoS("goroutine stack dump written to %s", path)
	}
}
