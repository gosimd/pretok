//go:build goexperiment.simd && (arm64 || amd64)

package scan

import (
	"runtime"
	"simd/archsimd"
)

const runeSIMDBuild = true

// Use the buffered algorithm only where Unicode can vectorize its conversion.
// Unicode owns the kernel dispatch (NEON, AVX2 or AVX-512). This eligibility
// check mirrors its scalar fallback condition and runs once during plan setup.
func runeDecodingEnabled() bool {
	return runtime.GOARCH == "arm64" || archsimd.X86.AVX2()
}
