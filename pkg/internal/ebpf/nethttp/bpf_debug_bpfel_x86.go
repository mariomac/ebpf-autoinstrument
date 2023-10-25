// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package nethttp

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_debugFuncInvocation struct {
	StartMonotimeNs uint64
	Regs            struct {
		R15    uint64
		R14    uint64
		R13    uint64
		R12    uint64
		Bp     uint64
		Bx     uint64
		R11    uint64
		R10    uint64
		R9     uint64
		R8     uint64
		Ax     uint64
		Cx     uint64
		Dx     uint64
		Si     uint64
		Di     uint64
		OrigAx uint64
		Ip     uint64
		Cs     uint64
		Flags  uint64
		Sp     uint64
		Ss     uint64
	}
}

type bpf_debugGoroutineMetadata struct {
	Parent    uint64
	Timestamp uint64
}

// loadBpf_debug returns the embedded CollectionSpec for bpf_debug.
func loadBpf_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_debug: %w", err)
	}

	return spec, err
}

// loadBpf_debugObjects loads bpf_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_debugObjects
//	*bpf_debugPrograms
//	*bpf_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugSpecs struct {
	bpf_debugProgramSpecs
	bpf_debugMapSpecs
}

// bpf_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugProgramSpecs struct {
	UprobeServeHTTP           *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTP"`
	UprobeWriteHeader         *ebpf.ProgramSpec `ebpf:"uprobe_WriteHeader"`
	UprobeRoundTrip           *ebpf.ProgramSpec `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn     *ebpf.ProgramSpec `ebpf:"uprobe_roundTripReturn"`
	UprobeStartBackgroundRead *ebpf.ProgramSpec `ebpf:"uprobe_startBackgroundRead"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	Events                    *ebpf.MapSpec `ebpf:"events"`
	GolangMapbucketStorageMap *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	Newproc1                  *ebpf.MapSpec `ebpf:"newproc1"`
	OngoingGoroutines         *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests *ebpf.MapSpec `ebpf:"ongoing_http_client_requests"`
	OngoingServerRequests     *ebpf.MapSpec `ebpf:"ongoing_server_requests"`
}

// bpf_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugObjects struct {
	bpf_debugPrograms
	bpf_debugMaps
}

func (o *bpf_debugObjects) Close() error {
	return _Bpf_debugClose(
		&o.bpf_debugPrograms,
		&o.bpf_debugMaps,
	)
}

// bpf_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugMaps struct {
	Events                    *ebpf.Map `ebpf:"events"`
	GolangMapbucketStorageMap *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	Newproc1                  *ebpf.Map `ebpf:"newproc1"`
	OngoingGoroutines         *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests *ebpf.Map `ebpf:"ongoing_http_client_requests"`
	OngoingServerRequests     *ebpf.Map `ebpf:"ongoing_server_requests"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.Events,
		m.GolangMapbucketStorageMap,
		m.Newproc1,
		m.OngoingGoroutines,
		m.OngoingHttpClientRequests,
		m.OngoingServerRequests,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	UprobeServeHTTP           *ebpf.Program `ebpf:"uprobe_ServeHTTP"`
	UprobeWriteHeader         *ebpf.Program `ebpf:"uprobe_WriteHeader"`
	UprobeRoundTrip           *ebpf.Program `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn     *ebpf.Program `ebpf:"uprobe_roundTripReturn"`
	UprobeStartBackgroundRead *ebpf.Program `ebpf:"uprobe_startBackgroundRead"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.UprobeServeHTTP,
		p.UprobeWriteHeader,
		p.UprobeRoundTrip,
		p.UprobeRoundTripReturn,
		p.UprobeStartBackgroundRead,
	)
}

func _Bpf_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_debug_bpfel_x86.o
var _Bpf_debugBytes []byte
