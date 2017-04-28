package win

import (
	"syscall"
	"unsafe"
)

const (
	WM_IME_STARTCOMPOSITION = 0x010D
)

const (
	CFS_DEFAULT        = 0x0000
	CFS_RECT           = 0x0001
	CFS_POINT          = 0x0002
	CFS_FORCE_POSITION = 0x0020
	CFS_CANDIDATEPOS   = 0x0040
	CFS_EXCLUDE        = 0x0080
)

type CANDIDATEFORM struct {
	Index      uint32
	Style      uint32
	CurrentPos POINT
	Area       RECT
}

type COMPOSITIONFORM struct {
	Style      uint32
	CurrentPos POINT
	Area       RECT
}

type HIMC uint32

var (
	immGetContext           uintptr
	immGetCandidateWindow   uintptr
	immSetCandidateWindow   uintptr
	immReleaseContext       uintptr
	immSetCompositionWindow uintptr
)

func init() {
	libimm32 := MustLoadLibrary("imm32.dll")

	immGetContext = MustGetProcAddress(uintptr(libimm32), "ImmGetContext")
	immGetCandidateWindow = MustGetProcAddress(uintptr(libimm32), "ImmGetCandidateWindow")
	immSetCandidateWindow = MustGetProcAddress(uintptr(libimm32), "ImmSetCandidateWindow")
	immReleaseContext = MustGetProcAddress(uintptr(libimm32), "ImmReleaseContext")
	immSetCompositionWindow = MustGetProcAddress(uintptr(libimm32), "ImmSetCompositionWindow")
}

func ImmGetContext(hwnd HWND) HIMC {
	ret, _, _ := syscall.Syscall(immGetContext, 1, uintptr(hwnd), 0, 0)
	return HIMC(ret)
}
func ImmGetCandidateWindow(himc HIMC, index uint32, form *CANDIDATEFORM) bool {
	ret, _, _ := syscall.Syscall(immGetCandidateWindow, 3, uintptr(himc), uintptr(index), uintptr(unsafe.Pointer(form)))
	return ret != 0
}

func ImmSetCandidateWindow(himc HIMC, form *CANDIDATEFORM) bool {
	ret, _, _ := syscall.Syscall(immSetCandidateWindow, 2, uintptr(himc), uintptr(unsafe.Pointer(form)), 0)
	return ret != 0
}

func ImmReleaseContext(hwnd HWND, himc HIMC) bool {
	ret, _, _ := syscall.Syscall(immReleaseContext, 1, uintptr(hwnd), uintptr(himc), 0)
	return ret != 0
}

func ImmSetCompositionWindow(himc HIMC, form *COMPOSITIONFORM) bool {
	ret, _, _ := syscall.Syscall(immSetCompositionWindow, 2, uintptr(himc), uintptr(unsafe.Pointer(form)), 0)
	return ret != 0
}
