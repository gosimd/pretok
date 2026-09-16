//go:build !goexperiment.simd || (!arm64 && !amd64)

package scan

const runeSIMDBuild = false

// Materializing runes without SIMD can regress against the byte scanner.
func runeDecodingEnabled() bool { return false }
