// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

// Fast Binary Encoding int64 final model class
type FinalModelInt64 struct {
    buffer Buffer // Final model buffer
    offset int    // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelInt64) FBEAllocationSize() int { return fm.FBESize() }

// Get the final size
func (fm FinalModelInt64) FBESize() int { return 8 }
// Get the final extra size
func (fm FinalModelInt64) FBEExtra() int { return 0 }

// Get the final offset
func (fm FinalModelInt64) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelInt64) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelInt64) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelInt64) FBEUnshift(size int) { fm.offset -= size }

func NewFinalModelInt64(buffer Buffer, offset int) *FinalModelInt64 {
    return &FinalModelInt64{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelInt64) Verify() int {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm FinalModelInt64) Get() (int64, int) {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return 0, 0
    }

    return ReadInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize()
}

// Set the value
func (fm *FinalModelInt64) Set(value int64) int {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return 0
    }

    WriteInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize()
}