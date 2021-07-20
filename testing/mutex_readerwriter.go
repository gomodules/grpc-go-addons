/*
Copyright AppsCode Inc. and Contributors

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

package grpc_testing

import (
	"io"
	"sync"
)

// MutexReadWriter is a io.ReadWriter that can be read and worked on from multiple go routines.
type MutexReadWriter struct {
	sync.Mutex
	rw io.ReadWriter
}

// NewMutexReadWriter creates a new thread-safe io.ReadWriter.
func NewMutexReadWriter(rw io.ReadWriter) *MutexReadWriter {
	return &MutexReadWriter{rw: rw}
}

// Write implements the io.Writer interface.
func (m *MutexReadWriter) Write(p []byte) (int, error) {
	m.Lock()
	defer m.Unlock()
	return m.rw.Write(p)
}

// Read implements the io.Reader interface.
func (m *MutexReadWriter) Read(p []byte) (int, error) {
	m.Lock()
	defer m.Unlock()
	return m.rw.Read(p)
}
