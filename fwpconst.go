package win

///////////////////////////////////////////////////////////////////////////////
//
// GUIDs for built-in layers.
//
///////////////////////////////////////////////////////////////////////////////

// ac4a9833-f69d-4648-b261-6dc84835ef39
var FWPM_LAYER_INBOUND_TRANSPORT_V4_DISCARD = DEFINE_GUID(
	0xac4a9833, 0xf69d, 0x4648,
	0xb2, 0x61, 0x6d, 0xc8, 0x48, 0x35, 0xef, 0x39,
)

// 2a6ff955-3b2b-49d2-9848-ad9d72dcaab7
var FWPM_LAYER_INBOUND_TRANSPORT_V6_DISCARD = DEFINE_GUID(
	0x2a6ff955, 0x3b2b, 0x49d2,
	0x98, 0x48, 0xad, 0x9d, 0x72, 0xdc, 0xaa, 0xb7,
)

// 41390100-564c-4b32-bc1d-718048354d7c
var FWPM_LAYER_OUTBOUND_ICMP_ERROR_V4 = DEFINE_GUID(
	0x41390100, 0x564c, 0x4b32,
	0xbc, 0x1d, 0x71, 0x80, 0x48, 0x35, 0x4d, 0x7c,
)

// 7fb03b60-7b8d-4dfa-badd-980176fc4e12
var FWPM_LAYER_OUTBOUND_ICMP_ERROR_V6 = DEFINE_GUID(
	0x7fb03b60, 0x7b8d, 0x4dfa,
	0xba, 0xdd, 0x98, 0x01, 0x76, 0xfc, 0x4e, 0x12,
)

///////////////////////////////////////////////////////////////////////////////
//
// GUIDs for built-in sublayers.
//
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
//
// GUIDs for built-in conditions.
//
///////////////////////////////////////////////////////////////////////////////

// 0c1ba1af-5765-453f-af22-a8f791ac775b
var FWPM_CONDITION_IP_LOCAL_PORT = DEFINE_GUID(
	0x0c1ba1af, 0x5765, 0x453f,
	0xaf, 0x22, 0xa8, 0xf7, 0x91, 0xac, 0x77, 0x5b,
)
var FWPM_CONDITION_ICMP_TYPE = FWPM_CONDITION_IP_LOCAL_PORT

// 632ce23b-5167-435c-86d7-e903684aa80c
var FWPM_CONDITION_FLAGS = DEFINE_GUID(
	0x632ce23b, 0x5167, 0x435c,
	0x86, 0xd7, 0xe9, 0x03, 0x68, 0x4a, 0xa8, 0x0c,
)

///////////////////////////////////////////////////////////////////////////////
//
// GUIDs for built-in providers.
//
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
//
// GUIDs for built-in callouts.
//
///////////////////////////////////////////////////////////////////////////////

// eda08606-2494-4d78-89bc-67837c03b969
var FWPM_CALLOUT_WFP_TRANSPORT_LAYER_V4_SILENT_DROP = DEFINE_GUID(
	0xeda08606, 0x2494, 0x4d78,
	0x89, 0xbc, 0x67, 0x83, 0x7c, 0x03, 0xb9, 0x69,
)

// 8693cc74-a075-4156-b476-9286eece814e
var FWPM_CALLOUT_WFP_TRANSPORT_LAYER_V6_SILENT_DROP = DEFINE_GUID(
	0x8693cc74, 0xa075, 0x4156,
	0xb4, 0x76, 0x92, 0x86, 0xee, 0xce, 0x81, 0x4e,
)

///////////////////////////////////////////////////////////////////////////////
//
// GUIDs for built-in provider contexts.
//
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
//
// GUIDs for built-in keying modules.
//
///////////////////////////////////////////////////////////////////////////////
