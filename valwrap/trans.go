package valwrap

func StringPtr(a string) *string {
	return &a
}

func String(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}

func IntPtr(a int) *int {
	return &a
}

func Int(a *int) int {
	if a == nil {
		return 0
	}
	return *a
}

func Int8Ptr(a int8) *int8 {
	return &a
}

func Int8(a *int8) int8 {
	if a == nil {
		return 0
	}
	return *a
}

func Int16Ptr(a int16) *int16 {
	return &a
}

func Int16(a *int16) int16 {
	if a == nil {
		return 0
	}
	return *a
}

func Int32Ptr(a int32) *int32 {
	return &a
}

func Int32(a *int32) int32 {
	if a == nil {
		return 0
	}
	return *a
}

func Int64Ptr(a int64) *int64 {
	return &a
}

func Int64(a *int64) int64 {
	if a == nil {
		return 0
	}
	return *a
}

func BoolPtr(a bool) *bool {
	return &a
}

func Bool(a *bool) bool {
	if a == nil {
		return false
	}
	return *a
}

func UintPtr(a uint) *uint {
	return &a
}

func Uint(a *uint) uint {
	if a == nil {
		return 0
	}
	return *a
}

func Uint8Ptr(a uint8) *uint8 {
	return &a
}

func Uint8(a *uint8) uint8 {
	if a == nil {
		return 0
	}
	return *a
}

func Uint16Ptr(a uint16) *uint16 {
	return &a
}

func Uint16(a *uint16) uint16 {
	if a == nil {
		return 0
	}
	return *a
}

func Uint32Ptr(a uint32) *uint32 {
	return &a
}

func Uint32(a *uint32) uint32 {
	if a == nil {
		return 0
	}
	return *a
}

func Uint64Ptr(a uint64) *uint64 {
	return &a
}

func Uint64(a *uint64) uint64 {
	if a == nil {
		return 0
	}
	return *a
}

func Float32Ptr(a float32) *float32 {
	return &a
}

func Float32(a *float32) float32 {
	if a == nil {
		return 0
	}
	return *a
}

func Float64Ptr(a float64) *float64 {
	return &a
}

func Float64(a *float64) float64 {
	if a == nil {
		return 0
	}
	return *a
}

func IntSlicePtr(a []int) []*int {
	if a == nil {
		return nil
	}
	res := make([]*int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func IntSlice(a []*int) []int {
	if a == nil {
		return nil
	}
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Int8SlicePtr(a []int8) []*int8 {
	if a == nil {
		return nil
	}
	res := make([]*int8, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Int8Slice(a []*int8) []int8 {
	if a == nil {
		return nil
	}
	res := make([]int8, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Int16SlicePtr(a []int16) []*int16 {
	if a == nil {
		return nil
	}
	res := make([]*int16, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Int16Slice(a []*int16) []int16 {
	if a == nil {
		return nil
	}
	res := make([]int16, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Int32SlicePtr(a []int32) []*int32 {
	if a == nil {
		return nil
	}
	res := make([]*int32, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Int32Slice(a []*int32) []int32 {
	if a == nil {
		return nil
	}
	res := make([]int32, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Int64SlicePtr(a []int64) []*int64 {
	if a == nil {
		return nil
	}
	res := make([]*int64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Int64Slice(a []*int64) []int64 {
	if a == nil {
		return nil
	}
	res := make([]int64, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func UintSlicePtr(a []uint) []*uint {
	if a == nil {
		return nil
	}
	res := make([]*uint, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func UintSlice(a []*uint) []uint {
	if a == nil {
		return nil
	}
	res := make([]uint, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Uint8SlicePtr(a []uint8) []*uint8 {
	if a == nil {
		return nil
	}
	res := make([]*uint8, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Uint8Slice(a []*uint8) []uint8 {
	if a == nil {
		return nil
	}
	res := make([]uint8, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Uint16SlicePtr(a []uint16) []*uint16 {
	if a == nil {
		return nil
	}
	res := make([]*uint16, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Uint16Slice(a []*uint16) []uint16 {
	if a == nil {
		return nil
	}
	res := make([]uint16, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Uint32SlicePtr(a []uint32) []*uint32 {
	if a == nil {
		return nil
	}
	res := make([]*uint32, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Uint32Slice(a []*uint32) []uint32 {
	if a == nil {
		return nil
	}
	res := make([]uint32, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Uint64SlicePtr(a []uint64) []*uint64 {
	if a == nil {
		return nil
	}
	res := make([]*uint64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Uint64Slice(a []*uint64) []uint64 {
	if a == nil {
		return nil
	}
	res := make([]uint64, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Float32SlicePtr(a []float32) []*float32 {
	if a == nil {
		return nil
	}
	res := make([]*float32, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Float32Slice(a []*float32) []float32 {
	if a == nil {
		return nil
	}
	res := make([]float32, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func Float64SlicePtr(a []float64) []*float64 {
	if a == nil {
		return nil
	}
	res := make([]*float64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func Float64Slice(a []*float64) []float64 {
	if a == nil {
		return nil
	}
	res := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func StringSlicePtr(a []string) []*string {
	if a == nil {
		return nil
	}
	res := make([]*string, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func StringSlice(a []*string) []string {
	if a == nil {
		return nil
	}
	res := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}

func BoolSlicePtr(a []bool) []*bool {
	if a == nil {
		return nil
	}
	res := make([]*bool, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = &a[i]
	}
	return res
}

func BoolSlice(a []*bool) []bool {
	if a == nil {
		return nil
	}
	res := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			res[i] = *a[i]
		}
	}
	return res
}
