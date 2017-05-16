package win

import (
	"syscall"
	"unsafe"
)

// WFP Error Codes
const (
	FWP_E_CALLOUT_NOT_FOUND                     = 0x80320001
	FWP_E_CONDITION_NOT_FOUND                   = 0x80320002
	FWP_E_FILTER_NOT_FOUND                      = 0x80320003
	FWP_E_LAYER_NOT_FOUND                       = 0x80320004
	FWP_E_PROVIDER_NOT_FOUND                    = 0x80320005
	FWP_E_PROVIDER_CONTEXT_NOT_FOUND            = 0x80320006
	FWP_E_SUBLAYER_NOT_FOUND                    = 0x80320007
	FWP_E_NOT_FOUND                             = 0x80320008
	FWP_E_ALREADY_EXISTS                        = 0x80320009
	FWP_E_IN_USE                                = 0x8032000A
	FWP_E_DYNAMIC_SESSION_IN_PROGRESS           = 0x8032000B
	FWP_E_WRONG_SESSION                         = 0x8032000C
	FWP_E_NO_TXN_IN_PROGRESS                    = 0x8032000D
	FWP_E_TXN_IN_PROGRESS                       = 0x8032000E
	FWP_E_TXN_ABORTED                           = 0x8032000F
	FWP_E_SESSION_ABORTED                       = 0x80320010
	FWP_E_INCOMPATIBLE_TXN                      = 0x80320011
	FWP_E_TIMEOUT                               = 0x80320012
	FWP_E_NET_EVENTS_DISABLED                   = 0x80320013
	FWP_E_INCOMPATIBLE_LAYER                    = 0x80320014
	FWP_E_KM_CLIENTS_ONLY                       = 0x80320015
	FWP_E_LIFETIME_MISMATCH                     = 0x80320016
	FWP_E_BUILTIN_OBJECT                        = 0x80320017
	FWP_E_TOO_MANY_CALLOUTS                     = 0x80320018
	FWP_E_NOTIFICATION_DROPPED                  = 0x80320019
	FWP_E_TRAFFIC_MISMATCH                      = 0x8032001A
	FWP_E_INCOMPATIBLE_SA_STATE                 = 0x8032001B
	FWP_E_NULL_POINTER                          = 0x8032001C
	FWP_E_INVALID_ENUMERATOR                    = 0x8032001D
	FWP_E_INVALID_FLAGS                         = 0x8032001E
	FWP_E_INVALID_NET_MASK                      = 0x8032001F
	FWP_E_INVALID_RANGE                         = 0x80320020
	FWP_E_INVALID_INTERVAL                      = 0x80320021
	FWP_E_ZERO_LENGTH_ARRAY                     = 0x80320022
	FWP_E_NULL_DISPLAY_NAME                     = 0x80320023
	FWP_E_INVALID_ACTION_TYPE                   = 0x80320024
	FWP_E_INVALID_WEIGHT                        = 0x80320025
	FWP_E_MATCH_TYPE_MISMATCH                   = 0x80320026
	FWP_E_TYPE_MISMATCH                         = 0x80320027
	FWP_E_OUT_OF_BOUNDS                         = 0x80320028
	FWP_E_RESERVED                              = 0x80320029
	FWP_E_DUPLICATE_CONDITION                   = 0x8032002A
	FWP_E_DUPLICATE_KEYMOD                      = 0x8032002B
	FWP_E_ACTION_INCOMPATIBLE_WITH_LAYER        = 0x8032002C
	FWP_E_ACTION_INCOMPATIBLE_WITH_SUBLAYER     = 0x8032002D
	FWP_E_CONTEXT_INCOMPATIBLE_WITH_LAYER       = 0x8032002E
	FWP_E_CONTEXT_INCOMPATIBLE_WITH_CALLOUT     = 0x8032002F
	FWP_E_INCOMPATIBLE_AUTH_METHOD              = 0x80320030
	FWP_E_INCOMPATIBLE_DH_GROUP                 = 0x80320031
	FWP_E_EM_NOT_SUPPORTED                      = 0x80320032
	FWP_E_NEVER_MATCH                           = 0x80320033
	FWP_E_PROVIDER_CONTEXT_MISMATCH             = 0x80320034
	FWP_E_INVALID_PARAMETER                     = 0x80320035
	FWP_E_TOO_MANY_SUBLAYERS                    = 0x80320036
	FWP_E_CALLOUT_NOTIFICATION_FAILED           = 0x80320037
	FWP_E_INVALID_AUTH_TRANSFORM                = 0x80320038
	FWP_E_INVALID_CIPHER_TRANSFORM              = 0x80320039
	FWP_E_INCOMPATIBLE_CIPHER_TRANSFORM         = 0x8032003A
	FWP_E_INVALID_TRANSFORM_COMBINATION         = 0x8032003B
	FWP_E_DUPLICATE_AUTH_METHOD                 = 0x8032003C
	FWP_E_INVALID_TUNNEL_ENDPOINT               = 0x8032003D
	FWP_E_L2_DRIVER_NOT_READY                   = 0x8032003E
	FWP_E_KEY_DICTATOR_ALREADY_REGISTERED       = 0x8032003F
	FWP_E_KEY_DICTATION_INVALID_KEYING_MATERIAL = 0x80320040
	FWP_E_CONNECTIONS_DISABLED                  = 0x80320041
	FWP_E_INVALID_DNS_NAME                      = 0x80320042
	FWP_E_STILL_ON                              = 0x80320043
	FWP_E_IKEEXT_NOT_RUNNING                    = 0x80320044
	FWP_E_DROP_NOICMP                           = 0x80320104
)

const (
	RPC_C_AUTHN_WINNT   = 10
	RPC_C_AUTHN_DEFAULT = 0xFFFFFFFF
)

const (
	FWPM_PROVIDER_FLAG_PERSISTENT = 0x00000001
	FWPM_PROVIDER_FLAG_DISABLED   = 0x00000010
)

const (
	FWPM_SESSION_FLAG_DYNAMIC = 1
)

const (
	FWPM_TXN_READ_ONLY = 1
)

const (
	FWP_ACTION_FLAG_TERMINATING     = 0x00001000
	FWP_ACTION_FLAG_NON_TERMINATING = 0x00002000
	FWP_ACTION_FLAG_CALLOUT         = 0x00004000
)

type FWP_ACTION_TYPE uint32

const (
	FWP_ACTION_BLOCK               FWP_ACTION_TYPE = 0x00000001 | FWP_ACTION_FLAG_TERMINATING
	FWP_ACTION_PERMIT                              = 0x00000002 | FWP_ACTION_FLAG_TERMINATING
	FWP_ACTION_CALLOUT_TERMINATING                 = 0x00000003 | FWP_ACTION_FLAG_CALLOUT | FWP_ACTION_FLAG_TERMINATING
	FWP_ACTION_CALLOUT_INSPECTION                  = 0x00000004 | FWP_ACTION_FLAG_CALLOUT | FWP_ACTION_FLAG_NON_TERMINATING
	FWP_ACTION_CALLOUT_UNKNOWN                     = 0x00000005 | FWP_ACTION_FLAG_CALLOUT
	FWP_ACTION_CONTINUE                            = 0x00000006 | FWP_ACTION_FLAG_NON_TERMINATING
)

const (
	FWP_CONDITION_FLAG_IS_LOOPBACK              = 0x00000001
	FWP_CONDITION_FLAG_IS_IPSEC_SECURED         = 0x00000002
	FWP_CONDITION_FLAG_IS_REAUTHORIZE           = 0x00000004
	FWP_CONDITION_FLAG_IS_WILDCARD_BIND         = 0x00000008
	FWP_CONDITION_FLAG_IS_RAW_ENDPOINT          = 0x00000010
	FWP_CONDITION_FLAG_IS_FRAGMENT              = 0x00000020
	FWP_CONDITION_FLAG_IS_FRAGMENT_GROUP        = 0x00000040
	FWP_CONDITION_FLAG_IS_IPSEC_NATT_RECLASSIFY = 0x00000080
	FWP_CONDITION_FLAG_REQUIRES_ALE_CLASSIFY    = 0x00000100
	FWP_CONDITION_FLAG_IS_IMPLICIT_BIND         = 0x00000200
)

type FWP_MATCH_TYPE uint32

const (
	FWP_MATCH_EQUAL FWP_MATCH_TYPE = iota
	FWP_MATCH_GREATER
	FWP_MATCH_LESS
	FWP_MATCH_GREATER_OR_EQUAL
	FWP_MATCH_LESS_OR_EQUAL
	FWP_MATCH_RANGE
	FWP_MATCH_FLAGS_ALL_SET
	FWP_MATCH_FLAGS_ANY_SET
	FWP_MATCH_FLAGS_NONE_SET
	FWP_MATCH_EQUAL_CASE_INSENSITIVE
	FWP_MATCH_NOT_EQUAL
	FWP_MATCH_TYPE_MAX
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

const (
	FWP_FILTER_ENUM_FLAG_BEST_TERMINATING_MATCH = 0x00000001
	FWP_FILTER_ENUM_FLAG_SORTED                 = 0x00000002
	FWP_FILTER_ENUM_FLAG_BOOTTIME_ONLY          = 0x00000004
	FWP_FILTER_ENUM_FLAG_INCLUDE_BOOTTIME       = 0x00000008
	FWP_FILTER_ENUM_FLAG_INCLUDE_DISABLED       = 0x00000010
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

const (
	FWPM_SUBLAYER_FLAG_PERSISTENT = 1
)

type FWPM_SUBLAYER0 struct {
	SubLayerKey  GUID
	DisplayData  FWPM_DISPLAY_DATA0
	Flags        uint32
	ProviderKey  *GUID
	ProviderData FWP_BYTE_BLOB
	Weight       uint16
}

type FWPM_SESSION_ENUM_TEMPLATE0 struct {
	Reserved uint64
}

type FWPM_ACTION0 struct {
	Type FWP_ACTION_TYPE
	/*
	  union {
	    GUID filterType;
	    GUID calloutKey;
	  };
	*/
	CalloutKey GUID
}

type FWP_CONDITION_VALUE0 struct {
	Type FWP_DATA_TYPE
	Data uintptr
}

type FWPM_FILTER_CONDITION0 struct {
	FieldKey       GUID
	MatchType      FWP_MATCH_TYPE
	ConditionValue FWP_CONDITION_VALUE0
}

type FWPM_FILTER0 struct {
	FilterKey           GUID
	DisplayData         FWPM_DISPLAY_DATA0
	Flags               uint32
	ProviderKey         *GUID
	ProviderData        FWP_BYTE_BLOB
	LayerKey            GUID
	SubLayerKey         GUID
	Weight              FWP_VALUE0
	NumFilterConditions uint32
	FilterCondition     *FWPM_FILTER_CONDITION0
	Action              FWPM_ACTION0
	/*
		union {
		    UINT64 rawContext;
		    GUID   providerContextKey;
		};
	*/
	ProviderContextKey GUID
	reserved           *GUID
	FilterId           uint64
	EffectiveWeight    FWP_VALUE0
}

type FWP_FILTER_ENUM_TYPE uint32

const (
	FWP_FILTER_ENUM_FULLY_CONTAINED FWP_FILTER_ENUM_TYPE = iota
	FWP_FILTER_ENUM_OVERLAPPING
	FWP_FILTER_ENUM_TYPE_MAX
)

type FWPM_FILTER_ENUM_TEMPLATE0 struct {
	ProviderKey             *GUID
	LayerKey                GUID
	EnumType                FWP_FILTER_ENUM_TYPE
	Flags                   uint32
	ProviderContextTemplate uintptr
	NumFilterConditions     uint32
	FilterCondition         *FWPM_FILTER_CONDITION0
	ActionMask              uint32
	CalloutKey              *GUID
}

type FWPM_LAYER0 struct {
	LayerKey           GUID
	DisplayData        FWPM_DISPLAY_DATA0
	Flags              uint32
	NumFields          uint32
	Field              uintptr
	DefaultSubLayerKey GUID
	LayerId            uint16
}

type FWPM_CALLOUT0 struct {
	CalloutKey      GUID
	DisplayData     FWPM_DISPLAY_DATA0
	Flags           uint32
	ProviderKey     *GUID
	ProviderData    FWP_BYTE_BLOB
	ApplicableLayer GUID
	CalloutId       uint32
}

type FWPM_CALLOUT_ENUM_TEMPLATE0 struct {
	ProviderKey *GUID
	LayerKey    GUID
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

	// Filter Management
	fwpmFilterAdd0                  uintptr
	fwpmFilterCreateEnumHandle0     uintptr
	fwpmFilterDeleteById0           uintptr
	fwpmFilterDeleteByKey0          uintptr
	fwpmFilterDestroyEnumHandle0    uintptr
	fwpmFilterEnum0                 uintptr
	fwpmFilterGetById0              uintptr
	fwpmFilterGetByKey0             uintptr
	fwpmFilterGetSecurityInfoByKey0 uintptr
	fwpmFilterSetSecurityInfoByKey0 uintptr
	fwpmFilterSubscribeChanges0     uintptr
	fwpmFilterSubscriptionsGet0     uintptr
	fwpmFilterUnsubscribeChanges0   uintptr
	fwpmGetAppIdFromFileName0       uintptr

	// Layer Management
	fwpmLayerCreateEnumHandle0  uintptr
	fwpmLayerDestroyEnumHandle0 uintptr
	fwpmLayerEnum0              uintptr
	fwpmLayerGetById0           uintptr
	fwpmLayerGetByKey0          uintptr

	// Callout Management
	fwpmCalloutAdd0               uintptr
	fwpmCalloutCreateEnumHandle0  uintptr
	fwpmCalloutDeleteById0        uintptr
	fwpmCalloutDeleteByKey0       uintptr
	fwpmCalloutDestroyEnumHandle0 uintptr
	fwpmCalloutEnum0              uintptr
	fwpmCalloutGetById0           uintptr
	fwpmCalloutGetByKey0          uintptr

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

	// Filter Managment
	fwpmFilterAdd0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterAdd0")
	fwpmFilterCreateEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterCreateEnumHandle0")
	fwpmFilterDeleteById0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterDeleteById0")
	fwpmFilterDeleteByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterDeleteByKey0")
	fwpmFilterDestroyEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterDestroyEnumHandle0")
	fwpmFilterEnum0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterEnum0")
	fwpmFilterGetById0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterGetById0")
	fwpmFilterGetByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterGetByKey0")
	fwpmFilterGetSecurityInfoByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterGetSecurityInfoByKey0")
	fwpmFilterSetSecurityInfoByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterSetSecurityInfoByKey0")
	fwpmFilterSubscribeChanges0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterSubscribeChanges0")
	fwpmFilterSubscriptionsGet0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterSubscriptionsGet0")
	fwpmFilterUnsubscribeChanges0 = MustGetProcAddress(libfwpuclnt, "FwpmFilterUnsubscribeChanges0")
	fwpmGetAppIdFromFileName0 = MustGetProcAddress(libfwpuclnt, "FwpmGetAppIdFromFileName0")

	// Layer Management
	fwpmLayerCreateEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmLayerCreateEnumHandle0")
	fwpmLayerDestroyEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmLayerDestroyEnumHandle0")
	fwpmLayerEnum0 = MustGetProcAddress(libfwpuclnt, "FwpmLayerEnum0")
	fwpmLayerGetById0 = MustGetProcAddress(libfwpuclnt, "FwpmLayerGetById0")
	fwpmLayerGetByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmLayerGetByKey0")

	// Callout Management
	fwpmCalloutAdd0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutAdd0")
	fwpmCalloutCreateEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutCreateEnumHandle0")
	fwpmCalloutDeleteById0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutDeleteById0")
	fwpmCalloutDeleteByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutDeleteByKey0")
	fwpmCalloutDestroyEnumHandle0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutDestroyEnumHandle0")
	fwpmCalloutEnum0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutEnum0")
	fwpmCalloutGetById0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutGetById0")
	fwpmCalloutGetByKey0 = MustGetProcAddress(libfwpuclnt, "FwpmCalloutGetByKey0")

	// Transaction Management
	fwpmTransactionAbort0 = MustGetProcAddress(libfwpuclnt, "FwpmTransactionAbort0")
	fwpmTransactionBegin0 = MustGetProcAddress(libfwpuclnt, "FwpmTransactionBegin0")
	fwpmTransactionCommit0 = MustGetProcAddress(libfwpuclnt, "FwpmTransactionCommit0")
}

// Shared Functions
func FwpmFreeMemory0(p unsafe.Pointer) {
	syscall.Syscall(fwpmFreeMemory0, 1,
		uintptr(p),
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
		FwpmFreeMemory0(unsafe.Pointer(&fwpValue))
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

func FwpmProviderDeleteByKey0(engineHandle HANDLE, key *GUID) uint32 {
	ret, _, _ := syscall.Syscall(fwpmProviderDeleteByKey0, 2,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(key)),
		0,
	)
	return uint32(ret)
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
func FwpmSubLayerDeleteByKey0(engineHandle HANDLE, key *GUID) uint32 {
	ret, _, _ := syscall.Syscall(fwpmSubLayerDeleteByKey0, 2,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(key)),
		0,
	)
	return uint32(ret)
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

// Filter Management
func FwpmFilterAdd0(engineHandle HANDLE, filter *FWPM_FILTER0, sd uintptr, id *uint64) uint32 {
	ret, _, _ := syscall.Syscall6(fwpmFilterAdd0, 4,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(filter)),
		sd,
		uintptr(unsafe.Pointer(id)),
		0,
		0,
	)
	return uint32(ret)
}

func FwpmFilterCreateEnumHandle0(engineHandle HANDLE, enumTemplate *FWPM_FILTER_ENUM_TEMPLATE0, enumHandle *HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmFilterCreateEnumHandle0, 3,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(enumTemplate)),
		uintptr(unsafe.Pointer(enumHandle)),
	)
	return uint32(ret)
}

func FwpmFilterDeleteById0() {}

func FwpmFilterDeleteByKey0(engineHandle HANDLE, key *GUID) uint32 {
	ret, _, _ := syscall.Syscall(fwpmFilterDeleteByKey0, 2,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(key)),
		0,
	)
	return uint32(ret)
}

func FwpmFilterDestroyEnumHandle0(engineHandle HANDLE, enumHandle HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmFilterDestroyEnumHandle0, 2,
		uintptr(engineHandle),
		uintptr(enumHandle),
		0,
	)
	return uint32(ret)
}

func FwpmFilterEnum0(engineHandle HANDLE, enumHandle HANDLE, numEntriesRequested uint32, entries ***FWPM_FILTER0, numEntriesReturned *uint32) uint32 {
	ret, _, _ := syscall.Syscall6(fwpmFilterEnum0, 5,
		uintptr(engineHandle),
		uintptr(enumHandle),
		uintptr(numEntriesRequested),
		uintptr(unsafe.Pointer(entries)),
		uintptr(unsafe.Pointer(numEntriesReturned)),
		0,
	)
	return uint32(ret)
}

func FwpmFilterGetById0()              {}
func FwpmFilterGetByKey0()             {}
func FwpmFilterGetSecurityInfoByKey0() {}
func FwpmFilterSetSecurityInfoByKey0() {}
func FwpmFilterSubscribeChanges0()     {}
func FwpmFilterSubscriptionsGet0()     {}
func FwpmFilterUnsubscribeChanges0()   {}
func FwpmGetAppIdFromFileName0()       {}

// Layer Management
func FwpmLayerCreateEnumHandle0(engineHandle HANDLE, enumTemplate uintptr, enumHandle *HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmLayerCreateEnumHandle0, 3,
		uintptr(engineHandle),
		enumTemplate,
		uintptr(unsafe.Pointer(enumHandle)),
	)
	return uint32(ret)
}

func FwpmLayerDestroyEnumHandle0(engineHandle HANDLE, enumHandle HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmLayerDestroyEnumHandle0, 2,
		uintptr(engineHandle),
		uintptr(enumHandle),
		0,
	)
	return uint32(ret)
}

func FwpmLayerEnum0(engineHandle HANDLE, enumHandle HANDLE, numEntriesRequested uint32, entries ***FWPM_LAYER0, numEntriesReturned *uint32) uint32 {
	ret, _, _ := syscall.Syscall6(fwpmLayerEnum0, 5,
		uintptr(engineHandle),
		uintptr(enumHandle),
		uintptr(numEntriesRequested),
		uintptr(unsafe.Pointer(entries)),
		uintptr(unsafe.Pointer(numEntriesReturned)),
		0,
	)
	return uint32(ret)
}

func FwpmLayerGetById0(engineHandle HANDLE, id uint16, layer **FWPM_LAYER0) uint32 {
	ret, _, _ := syscall.Syscall(fwpmLayerGetById0, 3,
		uintptr(engineHandle),
		uintptr(id),
		uintptr(unsafe.Pointer(layer)),
	)
	return uint32(ret)
}

func FwpmLayerGetByKey0(engineHandle HANDLE, key *GUID, layer **FWPM_LAYER0) uint32 {
	ret, _, _ := syscall.Syscall(fwpmLayerGetByKey0, 3,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(key)),
		uintptr(unsafe.Pointer(layer)),
	)
	return uint32(ret)
}

// Callout Management
func FwpmCalloutAdd0(engineHandle HANDLE, callout *FWPM_CALLOUT0, sd uintptr, id *uint32) uint32 {
	ret, _, _ := syscall.Syscall6(fwpmCalloutAdd0, 4,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(callout)),
		sd,
		uintptr(unsafe.Pointer(id)),
		0,
		0,
	)
	return uint32(ret)
}
func FwpmCalloutCreateEnumHandle0(engineHandle HANDLE, enumTemplate *FWPM_CALLOUT_ENUM_TEMPLATE0, enumHandle *HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmCalloutCreateEnumHandle0, 3,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(enumTemplate)),
		uintptr(unsafe.Pointer(enumHandle)),
	)
	return uint32(ret)
}

func FwpmCalloutDeleteById0(engineHandle HANDLE, id uint32) uint32 {
	ret, _, _ := syscall.Syscall(fwpmCalloutDeleteById0, 2,
		uintptr(engineHandle),
		uintptr(id),
		0,
	)
	return uint32(ret)
}

func FwpmCalloutDeleteByKey0(engineHandle HANDLE, key *GUID) uint32 {
	ret, _, _ := syscall.Syscall(fwpmCalloutDeleteById0, 2,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(key)),
		0,
	)
	return uint32(ret)
}

func FwpmCalloutDestroyEnumHandle0(engineHandle HANDLE, enumHandle HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(fwpmCalloutDestroyEnumHandle0, 2,
		uintptr(engineHandle),
		uintptr(enumHandle),
		0,
	)
	return uint32(ret)
}

func FwpmCalloutEnum0(engineHandle HANDLE, enumHandle HANDLE, numEntriesRequested uint32, entries ***FWPM_CALLOUT0, numEntriesReturned *uint32) uint32 {
	ret, _, _ := syscall.Syscall6(fwpmCalloutEnum0, 5,
		uintptr(engineHandle),
		uintptr(enumHandle),
		uintptr(numEntriesRequested),
		uintptr(unsafe.Pointer(entries)),
		uintptr(unsafe.Pointer(numEntriesReturned)),
		0,
	)
	return uint32(ret)
}

func FwpmCalloutGetById0(engineHandle HANDLE, id uint32, callout **FWPM_CALLOUT0) uint32 {
	ret, _, _ := syscall.Syscall(fwpmCalloutGetById0, 3,
		uintptr(engineHandle),
		uintptr(id),
		uintptr(unsafe.Pointer(callout)),
	)
	return uint32(ret)
}

func FwpmCalloutGetByKey0(engineHandle HANDLE, key *GUID, callout **FWPM_CALLOUT0) uint32 {
	ret, _, _ := syscall.Syscall(fwpmCalloutGetByKey0, 3,
		uintptr(engineHandle),
		uintptr(unsafe.Pointer(key)),
		uintptr(unsafe.Pointer(callout)),
	)
	return uint32(ret)
}
