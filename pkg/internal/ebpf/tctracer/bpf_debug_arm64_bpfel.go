// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package tctracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_debugCallProtocolArgsT struct {
	PidConn    bpf_debugPidConnectionInfoT
	SmallBuf   [24]uint8
	U_buf      uint64
	BytesLen   int32
	Ssl        uint8
	Direction  uint8
	OrigDport  uint16
	PacketType uint8
	_          [7]byte
}

type bpf_debugConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_debugCpSupportDataT struct {
	T_key      bpf_debugTraceKeyT
	RealClient uint8
}

type bpf_debugEgressKeyT struct {
	S_port uint16
	D_port uint16
}

type bpf_debugGoAddrKeyT struct {
	Pid  uint64
	Addr uint64
}

type bpf_debugGrpcFramesCtxT struct {
	PrevInfo        bpf_debugHttp2GrpcRequestT
	HasPrevInfo     uint8
	_               [3]byte
	Pos             int32
	SavedBufPos     int32
	SavedStreamId   uint32
	FoundDataFrame  uint8
	Iterations      uint8
	TerminateSearch uint8
	_               [1]byte
	Stream          bpf_debugHttp2ConnStreamT
	Args            bpf_debugCallProtocolArgsT
}

type bpf_debugHttp2ConnInfoDataT struct {
	Id    uint64
	Flags uint8
	_     [7]byte
}

type bpf_debugHttp2ConnStreamT struct {
	PidConn  bpf_debugPidConnectionInfoT
	StreamId uint32
}

type bpf_debugHttp2GrpcRequestT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_debugConnectionInfoT
	Data            [256]uint8
	RetData         [64]uint8
	Type            uint8
	_               [3]byte
	Len             int32
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Ssl       uint8
	_         [3]byte
	NewConnId uint64
	Tp        struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_debugHttpConnectionMetadataT struct {
	Pid struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Type uint8
}

type bpf_debugHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	Tp              struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_debugHttpInfoT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_debugConnectionInfoT
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	Len             uint32
	RespLen         uint32
	Status          uint16
	Type            uint8
	Ssl             uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	ExtraId uint64
	TaskTid uint32
	_       [4]byte
}

type bpf_debugMsgBufferT struct {
	Buf [256]uint8
	Pos uint16
}

type bpf_debugMsgDataT struct{ Buf [1024]uint8 }

type bpf_debugPidConnectionInfoT struct {
	Conn bpf_debugConnectionInfoT
	Pid  uint32
}

type bpf_debugPidKeyT struct {
	Pid uint32
	Ns  uint32
}

type bpf_debugSslArgsT struct {
	Ssl    uint64
	Buf    uint64
	LenPtr uint64
}

type bpf_debugSslPidConnectionInfoT struct {
	P_conn    bpf_debugPidConnectionInfoT
	OrigDport uint16
	_         [2]byte
}

type bpf_debugTcHttpCtx struct {
	Offset  uint32
	Seen    uint32
	Written uint32
}

type bpf_debugTcpReqT struct {
	Flags           uint8
	_               [3]byte
	ConnInfo        bpf_debugConnectionInfoT
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	Rbuf            [128]uint8
	Len             uint32
	RespLen         uint32
	Ssl             uint8
	Direction       uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_  [2]byte
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	ExtraId uint64
}

type bpf_debugTpBufDataT struct{ Buf [256]uint8 }

type bpf_debugTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid     uint32
	Valid   uint8
	ReqType uint8
	_       [2]byte
}

type bpf_debugTraceKeyT struct {
	P_key   bpf_debugPidKeyT
	ExtraId uint64
}

type bpf_debugTraceMapKeyT struct {
	Conn bpf_debugConnectionInfoT
	Type uint32
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
	BeylaAppEgress                         *ebpf.ProgramSpec `ebpf:"beyla_app_egress"`
	BeylaAppIngress                        *ebpf.ProgramSpec `ebpf:"beyla_app_ingress"`
	BeylaAsyncReset                        *ebpf.ProgramSpec `ebpf:"beyla_async_reset"`
	BeylaEmitAsyncInit                     *ebpf.ProgramSpec `ebpf:"beyla_emit_async_init"`
	BeylaPacketExtender                    *ebpf.ProgramSpec `ebpf:"beyla_packet_extender"`
	BeylaProtocolHttp                      *ebpf.ProgramSpec `ebpf:"beyla_protocol_http"`
	BeylaProtocolHttp2                     *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2"`
	BeylaProtocolHttp2GrpcFrames           *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_frames"`
	BeylaProtocolHttp2GrpcHandleEndFrame   *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_handle_end_frame"`
	BeylaProtocolHttp2GrpcHandleStartFrame *ebpf.ProgramSpec `ebpf:"beyla_protocol_http2_grpc_handle_start_frame"`
	BeylaProtocolTcp                       *ebpf.ProgramSpec `ebpf:"beyla_protocol_tcp"`
	BeylaSockmapTracker                    *ebpf.ProgramSpec `ebpf:"beyla_sockmap_tracker"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	ActiveNodejsIds           *ebpf.MapSpec `ebpf:"active_nodejs_ids"`
	ActiveSslConnections      *ebpf.MapSpec `ebpf:"active_ssl_connections"`
	ActiveSslReadArgs         *ebpf.MapSpec `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs        *ebpf.MapSpec `ebpf:"active_ssl_write_args"`
	ActiveUnixSocks           *ebpf.MapSpec `ebpf:"active_unix_socks"`
	AsyncResetArgs            *ebpf.MapSpec `ebpf:"async_reset_args"`
	BufMem                    *ebpf.MapSpec `ebpf:"buf_mem"`
	CloneMap                  *ebpf.MapSpec `ebpf:"clone_map"`
	ConnectionMetaMem         *ebpf.MapSpec `ebpf:"connection_meta_mem"`
	CpSupportConnectInfo      *ebpf.MapSpec `ebpf:"cp_support_connect_info"`
	DebugEvents               *ebpf.MapSpec `ebpf:"debug_events"`
	Events                    *ebpf.MapSpec `ebpf:"events"`
	GrpcFramesCtxMem          *ebpf.MapSpec `ebpf:"grpc_frames_ctx_mem"`
	Http2InfoMem              *ebpf.MapSpec `ebpf:"http2_info_mem"`
	HttpInfoMem               *ebpf.MapSpec `ebpf:"http_info_mem"`
	IncomingTraceMap          *ebpf.MapSpec `ebpf:"incoming_trace_map"`
	IovecMem                  *ebpf.MapSpec `ebpf:"iovec_mem"`
	JumpTable                 *ebpf.MapSpec `ebpf:"jump_table"`
	MsgBuffers                *ebpf.MapSpec `ebpf:"msg_buffers"`
	NodejsParentMap           *ebpf.MapSpec `ebpf:"nodejs_parent_map"`
	OngoingGoHttp             *ebpf.MapSpec `ebpf:"ongoing_go_http"`
	OngoingHttp               *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttp2Connections   *ebpf.MapSpec `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc          *ebpf.MapSpec `ebpf:"ongoing_http2_grpc"`
	OngoingHttpClientRequests *ebpf.MapSpec `ebpf:"ongoing_http_client_requests"`
	OngoingHttpFallback       *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq             *ebpf.MapSpec `ebpf:"ongoing_tcp_req"`
	OutgoingTraceMap          *ebpf.MapSpec `ebpf:"outgoing_trace_map"`
	PidCache                  *ebpf.MapSpec `ebpf:"pid_cache"`
	PidTidToConn              *ebpf.MapSpec `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem           *ebpf.MapSpec `ebpf:"protocol_args_mem"`
	ServerTraces              *ebpf.MapSpec `ebpf:"server_traces"`
	SockDir                   *ebpf.MapSpec `ebpf:"sock_dir"`
	SslToConn                 *ebpf.MapSpec `ebpf:"ssl_to_conn"`
	SslToPidTid               *ebpf.MapSpec `ebpf:"ssl_to_pid_tid"`
	TcHttpCtxMap              *ebpf.MapSpec `ebpf:"tc_http_ctx_map"`
	TcpReqMem                 *ebpf.MapSpec `ebpf:"tcp_req_mem"`
	TpBufMemory               *ebpf.MapSpec `ebpf:"tp_buf_memory"`
	TpCharBufMem              *ebpf.MapSpec `ebpf:"tp_char_buf_mem"`
	TpInfoMem                 *ebpf.MapSpec `ebpf:"tp_info_mem"`
	TraceMap                  *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids                 *ebpf.MapSpec `ebpf:"valid_pids"`
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
	ActiveNodejsIds           *ebpf.Map `ebpf:"active_nodejs_ids"`
	ActiveSslConnections      *ebpf.Map `ebpf:"active_ssl_connections"`
	ActiveSslReadArgs         *ebpf.Map `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs        *ebpf.Map `ebpf:"active_ssl_write_args"`
	ActiveUnixSocks           *ebpf.Map `ebpf:"active_unix_socks"`
	AsyncResetArgs            *ebpf.Map `ebpf:"async_reset_args"`
	BufMem                    *ebpf.Map `ebpf:"buf_mem"`
	CloneMap                  *ebpf.Map `ebpf:"clone_map"`
	ConnectionMetaMem         *ebpf.Map `ebpf:"connection_meta_mem"`
	CpSupportConnectInfo      *ebpf.Map `ebpf:"cp_support_connect_info"`
	DebugEvents               *ebpf.Map `ebpf:"debug_events"`
	Events                    *ebpf.Map `ebpf:"events"`
	GrpcFramesCtxMem          *ebpf.Map `ebpf:"grpc_frames_ctx_mem"`
	Http2InfoMem              *ebpf.Map `ebpf:"http2_info_mem"`
	HttpInfoMem               *ebpf.Map `ebpf:"http_info_mem"`
	IncomingTraceMap          *ebpf.Map `ebpf:"incoming_trace_map"`
	IovecMem                  *ebpf.Map `ebpf:"iovec_mem"`
	JumpTable                 *ebpf.Map `ebpf:"jump_table"`
	MsgBuffers                *ebpf.Map `ebpf:"msg_buffers"`
	NodejsParentMap           *ebpf.Map `ebpf:"nodejs_parent_map"`
	OngoingGoHttp             *ebpf.Map `ebpf:"ongoing_go_http"`
	OngoingHttp               *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttp2Connections   *ebpf.Map `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc          *ebpf.Map `ebpf:"ongoing_http2_grpc"`
	OngoingHttpClientRequests *ebpf.Map `ebpf:"ongoing_http_client_requests"`
	OngoingHttpFallback       *ebpf.Map `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq             *ebpf.Map `ebpf:"ongoing_tcp_req"`
	OutgoingTraceMap          *ebpf.Map `ebpf:"outgoing_trace_map"`
	PidCache                  *ebpf.Map `ebpf:"pid_cache"`
	PidTidToConn              *ebpf.Map `ebpf:"pid_tid_to_conn"`
	ProtocolArgsMem           *ebpf.Map `ebpf:"protocol_args_mem"`
	ServerTraces              *ebpf.Map `ebpf:"server_traces"`
	SockDir                   *ebpf.Map `ebpf:"sock_dir"`
	SslToConn                 *ebpf.Map `ebpf:"ssl_to_conn"`
	SslToPidTid               *ebpf.Map `ebpf:"ssl_to_pid_tid"`
	TcHttpCtxMap              *ebpf.Map `ebpf:"tc_http_ctx_map"`
	TcpReqMem                 *ebpf.Map `ebpf:"tcp_req_mem"`
	TpBufMemory               *ebpf.Map `ebpf:"tp_buf_memory"`
	TpCharBufMem              *ebpf.Map `ebpf:"tp_char_buf_mem"`
	TpInfoMem                 *ebpf.Map `ebpf:"tp_info_mem"`
	TraceMap                  *ebpf.Map `ebpf:"trace_map"`
	ValidPids                 *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.ActiveNodejsIds,
		m.ActiveSslConnections,
		m.ActiveSslReadArgs,
		m.ActiveSslWriteArgs,
		m.ActiveUnixSocks,
		m.AsyncResetArgs,
		m.BufMem,
		m.CloneMap,
		m.ConnectionMetaMem,
		m.CpSupportConnectInfo,
		m.DebugEvents,
		m.Events,
		m.GrpcFramesCtxMem,
		m.Http2InfoMem,
		m.HttpInfoMem,
		m.IncomingTraceMap,
		m.IovecMem,
		m.JumpTable,
		m.MsgBuffers,
		m.NodejsParentMap,
		m.OngoingGoHttp,
		m.OngoingHttp,
		m.OngoingHttp2Connections,
		m.OngoingHttp2Grpc,
		m.OngoingHttpClientRequests,
		m.OngoingHttpFallback,
		m.OngoingTcpReq,
		m.OutgoingTraceMap,
		m.PidCache,
		m.PidTidToConn,
		m.ProtocolArgsMem,
		m.ServerTraces,
		m.SockDir,
		m.SslToConn,
		m.SslToPidTid,
		m.TcHttpCtxMap,
		m.TcpReqMem,
		m.TpBufMemory,
		m.TpCharBufMem,
		m.TpInfoMem,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	BeylaAppEgress                         *ebpf.Program `ebpf:"beyla_app_egress"`
	BeylaAppIngress                        *ebpf.Program `ebpf:"beyla_app_ingress"`
	BeylaAsyncReset                        *ebpf.Program `ebpf:"beyla_async_reset"`
	BeylaEmitAsyncInit                     *ebpf.Program `ebpf:"beyla_emit_async_init"`
	BeylaPacketExtender                    *ebpf.Program `ebpf:"beyla_packet_extender"`
	BeylaProtocolHttp                      *ebpf.Program `ebpf:"beyla_protocol_http"`
	BeylaProtocolHttp2                     *ebpf.Program `ebpf:"beyla_protocol_http2"`
	BeylaProtocolHttp2GrpcFrames           *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_frames"`
	BeylaProtocolHttp2GrpcHandleEndFrame   *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_handle_end_frame"`
	BeylaProtocolHttp2GrpcHandleStartFrame *ebpf.Program `ebpf:"beyla_protocol_http2_grpc_handle_start_frame"`
	BeylaProtocolTcp                       *ebpf.Program `ebpf:"beyla_protocol_tcp"`
	BeylaSockmapTracker                    *ebpf.Program `ebpf:"beyla_sockmap_tracker"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.BeylaAppEgress,
		p.BeylaAppIngress,
		p.BeylaAsyncReset,
		p.BeylaEmitAsyncInit,
		p.BeylaPacketExtender,
		p.BeylaProtocolHttp,
		p.BeylaProtocolHttp2,
		p.BeylaProtocolHttp2GrpcFrames,
		p.BeylaProtocolHttp2GrpcHandleEndFrame,
		p.BeylaProtocolHttp2GrpcHandleStartFrame,
		p.BeylaProtocolTcp,
		p.BeylaSockmapTracker,
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
//go:embed bpf_debug_arm64_bpfel.o
var _Bpf_debugBytes []byte
