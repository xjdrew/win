package win

import (
	"syscall"
	"unsafe"
)

const (
	RPC_C_AUTHN_WINNT   = 10
	RPC_C_AUTHN_DEFAULT = 0xFFFFFFFF
)

const (
	FWPM_SESSION_FLAG_DYNAMIC = 1
)

const (
	FWPM_TXN_READ_ONLY = 1
)

// FWP_DATA_TYPE
type FWP_DATA_TYPE uint32

const (
	FWP_EMPTY FWP_DATA_TYPE = iota
	FWP_UINT8
	FWP_UINT16
	FWP_UINT32
	FWP_UINT64
	FWP_INT8
	FWP_INT16
	FWP_INT32
	FWP_INT64
	FWP_FLOAT
	FWP_DOUBLE
	FWP_BYTE_ARRAY16_TYPE
	FWP_BYTE_BLOB_TYPE
	FWP_SID
	FWP_SECURITY_DESCRIPTOR_TYPE
	FWP_TOKEN_INFORMATION_TYPE
	FWP_TOKEN_ACCESS_INFORMATION_TYPE
	FWP_UNICODE_STRING_TYPE
	FWP_BYTE_ARRAY6_TYPE
	FWP_SINGLE_DATA_TYPE_MAX = 0xFF
	FWP_V4_ADDR_MASK         = FWP_SINGLE_DATA_TYPE_MAX + 1
	FWP_V6_ADDR_MASK         = FWP_SINGLE_DATA_TYPE_MAX + 2
	FWP_RANGE_TYPE           = FWP_SINGLE_DATA_TYPE_MAX + 3
	FWP_DATA_TYPE_MAX        = FWP_SINGLE_DATA_TYPE_MAX + 4
)

type FWP_VALUE0 struct {
	Type FWP_DATA_TYPE
	Data uintptr
}

type FWPM_DISPLAY_DATA0 struct {
	Name        *uint16
	Description *uint16
}

type FWP_BYTE_BLOB struct {
	Size uint32
	Data uintptr
}

type FWPM_SESSION0 struct {
	SessionKey           GUID
	DisplayData          FWPM_DISPLAY_DATA0
	Flags                uint32
	TxnWaitTimeoutInMSec uint32
	ProcessId            uint32
	Sid                  uintptr
	username             uintptr
	kernelMode           BOOL
}

type FWPM_PROVIDER0 struct {
	ProviderKey  GUID
	DisplayData  FWPM_DISPLAY_DATA0
	Flags        uint32
	ProviderData FWP_BYTE_BLOB
	ServiceName  uintptr
}

type FWPM_SUBLAYER0 struct {
	SubLayerKey  GUID
	DisplayData  FWPM_DISPLAY_DATA0
	Flags        uint32
	ProviderKey  uintptr
	ProviderData FWP_BYTE_BLOB
	Weight       uint16
}

type FWPM_SESSION_ENUM_TEMPLATE0 struct {
	Reserved uint64
}

var (
	// Library
	libfwpuclnt uintptr

	// Shared Functions
	fwpmFreeMemory0 uintptr

	// Session Management
	fwpmEngineClose0              uintptr
	fwpmEngineGetOption0          uintptr
	fwpmEngineGetSecurityInfo0    uintptr
	fwpmEngineOpen0               uintptr
	fwpmEngineSetOption0          uintptr
	fwpmEngineSetSecurityInfo0    uintptr
	fwpmSessionCreateEnumHandle0  uintptr
	fwpmSessionDestroyEnumHandle0 uintptr
	fwpmSessionEnum0              uintptr

	// Provider Management
	fwpmProviderAdd0                  uintptr
	fwpmProviderCreateEnumHandle0     uintptr
	fwpmProviderDeleteByKey0          uintptr
	fwpmProviderDestroyEnumHandle0    uintptr
	fwpmProviderEnum0                 uintptr
	fwpmProviderGetByKey0             uintptr
	fwpmProviderGetSecurityInfoByKey0 uintptr
	fwpmProviderSetSecurityInfoByKey0 uintptr
	fwpmProviderSubscribeChanges0     uintptr
	fwpmProviderSubscriptionsGet0     uintptr
	fwpmProviderUnsubscribeChanges0   uintptr

	// Sublayer Management
	fwpmSubLayerAdd0                  uintptr
	fwpmSubLayerCreateEnumHandle0     uintptr
	fwpmSubLayerDeleteByKey0          uintptr
	fwpmSubLayerDestroyEnumHandle0    uintptr
	fwpmSubLayerEnum0                 uintptr
	fwpmSubLayerGetByKey0             uintptr
	fwpmSubLayerGetSecurityInfoByKey0 uintptr
	fwpmSubLayerSetSecurityInfoByKey0 uintptr
	fwpmSubLayerSubscribeChanges0     uintptr
	fwpmSubLayerSubscriptionsGet0     uintptr
	fwpmSubLayerUnsubscribeChanges0   uintptr

	// Transaction Management
	fwpmTransactionAbort0  uintptr
	fwpmTransactionBegin0  uintptr
	fwpmTransactionCommit0 uintptr
)

func init() {
	// Library
	libfwpuclnt = MustLoadLibrary("fwpuclnt.dll")

	// Shared Functions
	fwpmFreeMemory0 = MustGetProcAddress(libfwpuclnt, "FwpmFreeMemory0")

	// Session Management
	fwpmEngineClose0 = MustGetProcAddress(libfwpuclnt, "FwpmEngineClose0")
	fwpmEngineGetOption0 = MustGetProcAddress(libfwpuclnt, "FwpmEngineGetOption0")
	fwpmEngineGetSecurityInfo0 = MustGetProcAddress(libfwpuclnt, "FwpmEngineGetSecurityInfo0")
	fwpmEngineOpen0 = MustGetProcAddress(libfwpuclnt, "FwpmEngineOpen0")
	fwpmEngineSetOption0 = MustGetProcAddress(libfwpuclnt, "FwpmEngineSetOption0")
	fwpmEngineSetSecurityInfo0 = MustGetProcAddress(libfwpuclnt, "FwpmEngineSetSecurityInfo0")
	fwpmSessionCreateEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmSessionCreateEnumHandle0")
	fwpmSessionDestroyEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmSessionDestroyEnumHandle0")
	fwpmSessionEnum0 = MustGetProcAddress(libfwpuclnt, "FwpmSessionEnum0")

	// Provider Management
	fwpmProviderAdd0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderAdd0")
	fwpmProviderCreateEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderCreateEnumHandle0")
	fwpmProviderDeleteByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderDeleteByKey0")
	fwpmProviderDestroyEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderDestroyEnumHandle0")
	fwpmProviderEnum0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderEnum0")
	fwpmProviderGetByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderGetByKey0")
	fwpmProviderGetSecurityInfoByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderGetSecurityInfoByKey0")
	fwpmProviderSetSecurityInfoByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderSetSecurityInfoByKey0")
	fwpmProviderSubscribeChanges0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderSubscribeChanges0")
	fwpmProviderSubscriptionsGet0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderSubscriptionsGet0")
	fwpmProviderUnsubscribeChanges0 = MustGetProcAddress(libfwpuclnt, "FwpmProviderUnsubscribeChanges0")

	// Sublayer Management
	fwpmSubLayerAdd0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerAdd0")
	fwpmSubLayerCreateEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerCreateEnumHandle0")
	fwpmSubLayerDeleteByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerDeleteByKey0")
	fwpmSubLayerDestroyEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerDestroyEnumHandle0")
	fwpmSubLayerEnum0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerEnum0")
	fwpmSubLayerGetByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerGetByKey0")
	fwpmSubLayerGetSecurityInfoByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerGetSecurityInfoByKey0")
	fwpmSubLayerSetSecurityInfoByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerSetSecurityInfoByKey0")
	fwpmSubLayerSubscribeChanges0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerSubscribeChanges0")
	fwpmSubLayerSubscriptionsGet0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerSubscriptionsGet0")
	fwpmSubLayerUnsubscribeChanges0 = MustGetProcAddress(libfwpuclnt, "FwpmSubLayerUnsubscribeChanges0")

	// Transaction Management
	fwpmTransactionAbort0 = MustGetProcAddress(libfwpuclnt, "FwpmTransactionAbort0")
	fwpmTransactionBegin0 = MustGetProcAddress(libfwpuclnt, "FwpmTransactionBegin0")
	fwpmTransactionCommit0 = MustGetProcAddress(libfwpuclnt, "FwpmTransactionCommit0")
}

// Shared Functions
func FwpmFreeMemory0(p uintptr) {
	syscall.Syscall(fwpmFreeMemory0, 1,
		p,
		0,
		0)
}

// Session Management
func FwpmEngineClose0(engineHandle HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmEngineClose0, 1,
		uintptr(engineHandle),
		0,
		0)
	return uint32(ret)
}

func FwpmEngineGetOption0(engineHandle HANDLE, option uint32, value *uint32) uint32 {
	var fwpValue *FWP_VALUE0
	ret, _, _ := syscall.Syscall(fwpmEngineGetOption0, 3,
		uintptr(engineHandle),
		uintptr(option),
		uintptr(unsafe.Pointer(&fwpValue)))

	if ret == 0 {
		if fwpValue.Type == FWP_UINT32 {
			*value = *(*uint32)(unsafe.Pointer((&fwpValue.Data)))
		}
		FwpmFreeMemory0(uintptr(unsafe.Pointer(&fwpValue)))
	}
	return uint32(ret)
}

func FwpmEngineGetSecurityInfo0() {
}

func FwpmEngineOpen0(serverName *uint16, authnService uint32, authIdentity uintptr, session *FWPM_SESSION0, engineHandle *HANDLE) uint32 {
	ret, _, _ := syscall.Syscall6(fwpmEngineOpen0, 5,
		uintptr(unsafe.Pointer(serverName)),
		uintptr(authnService),
		authIdentity,
		uintptr(unsafe.Pointer(session)),
		uintptr(unsafe.Pointer(engineHandle)),
		0,
	)
	return uint32(ret)
}

func FwpmEngineSetOption0(engineHandle HANDLE, option uint32, value uint32) uint32 {
	var fwpValue FWP_VALUE0
	fwpValue.Type = FWP_UINT32
	*(*uint32)(unsafe.Pointer((&fwpValue.Data))) = value
	ret, _, _ := syscall.Syscall(fwpmEngineSetOption0, 3,
		uintptr(engineHandle),
		uintptr(option),
		uintptr(unsafe.Pointer(&fwpValue)))
	return uint32(ret)
}

func FwpmEngineSetSecurityInfo0() {
}

func FwpmSessionCreateEnumHandle0(engineHandle HANDLE, enumTemplate *FWPM_SESSION_ENUM_TEMPLATE0, enumHandle *HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmSessionCreateEnumHandle0, 3,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(enumTemplate)),
		uintptr(unsafe.Pointer(enumHandle)))
	return uint32(ret)
}

func FwpmSessionDestroyEnumHandle0(engineHandle HANDLE, enumHandle HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmSessionDestroyEnumHandle0, 2,
		uintptr(engineHandle),
		uintptr(enumHandle),
		0)
	return uint32(ret)
}

func FwpmSessionEnum0(engineHandle HANDLE, enumHandle HANDLE, numEntriesRequested uint32, sessions ***FWPM_SESSION0, numEntriesReturned *uint32) uint32 {
	ret, _, _ := syscall.Syscall6(fwpmSessionEnum0, 5,
		uintptr(engineHandle),
		uintptr(enumHandle),
		uintptr(numEntriesRequested),
		uintptr(unsafe.Pointer(sessions)),
		uintptr(unsafe.Pointer(numEntriesReturned)),
		0,
	)
	return uint32(ret)
}

// Transaction Management
func FwpmTransactionAbort0(engineHandle HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmTransactionAbort0, 1,
		uintptr(engineHandle),
		0,
		0)
	return uint32(ret)
}
func FwpmTransactionBegin0(engineHandle HANDLE, flags uint32) uint32 {
	ret, _, _ := syscall.Syscall(fwpmTransactionBegin0, 2,
		uintptr(engineHandle),
		uintptr(flags),
		0)
	return uint32(ret)
}
func FwpmTransactionCommit0(engineHandle HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmTransactionCommit0, 1,
		uintptr(engineHandle),
		0,
		0)
	return uint32(ret)
}

// Provider Management
func FwpmProviderAdd0(engineHandle HANDLE, provider *FWPM_PROVIDER0, sd uintptr) uint32 {
	ret, _, _ := syscall.Syscall(fwpmProviderAdd0, 3,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(provider)),
		uintptr(sd))
	return uint32(ret)
}

func FwpmProviderCreateEnumHandle0(engineHandle HANDLE, key *GUID) uint32 {
	ret, _, _ := syscall.Syscall(fwpmProviderCreateEnumHandle0, 2,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(key)),
		0)
	return uint32(ret)
}

func FwpmProviderDeleteByKey0() {
}
func FwpmProviderDestroyEnumHandle0() {
}
func FwpmProviderEnum0() {
}
func FwpmProviderGetByKey0() {
}
func FwpmProviderGetSecurityInfoByKey0() {
}
func FwpmProviderSetSecurityInfoByKey0() {
}
func FwpmProviderSubscribeChanges0() {
}
func FwpmProviderSubscriptionsGet0() {
}
func FwpmProviderUnsubscribeChanges0() {
}

// Sublayer Management
func FwpmSubLayerAdd0(engineHandle HANDLE, subLayer *FWPM_SUBLAYER0, sd uintptr) uint32 {
	ret, _, _ := syscall.Syscall(fwpmSubLayerAdd0, 3,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(subLayer)),
		uintptr(sd))
	return uint32(ret)
}

func FwpmSubLayerCreateEnumHandle0() {
}
func FwpmSubLayerDeleteByKey0() {
}
func FwpmSubLayerDestroyEnumHandle0() {
}
func FwpmSubLayerEnum0() {
}
func FwpmSubLayerGetByKey0() {
}
func FwpmSubLayerGetSecurityInfoByKey0() {
}
func FwpmSubLayerSetSecurityInfoByKey0() {
}
func FwpmSubLayerSubscribeChanges0() {
}
func FwpmSubLayerSubscriptionsGet0() {
}
func FwpmSubLayerUnsubscribeChanges0() {
}
