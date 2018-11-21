// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

// Fast Binary Encoding float64 final model class
type FinalModelDouble struct {
    buffer Buffer // Final model buffer
    offset int    // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelDouble) FBEAllocationSize() int { return fm.FBESize() }

// Get the final size
func (fm FinalModelDouble) FBESize() int { return 8 }
// Get the final extra size
func (fm FinalModelDouble) FBEExtra() int { return 0 }

// Get the final offset
func (fm FinalModelDouble) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelDouble) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelDouble) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelDouble) FBEUnshift(size int) { fm.offset -= size }

func NewFinalModelDouble(buffer Buffer, offset int) *FinalModelDouble {
    return &FinalModelDouble{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelDouble) Verify() int {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm FinalModelDouble) Get() (float64, int) {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return 0.0, 0
    }

    return ReadDouble(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize()
}

// Set the value
func (fm *FinalModelDouble) Set(value float64) int {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return 0
    }

    WriteDouble(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize()
}