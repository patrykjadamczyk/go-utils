package base

type FeatureFlags struct {
	IntFlags    map[string]int64
	UIntFlags   map[string]uint64
	StringFlags map[string]string
	BoolFlags   map[string]bool
	FloatFlags  map[string]float64
	AnyFlags    map[string]any
}

func (f FeatureFlags) Dev() Result[string] {
	return CheckAllResults[string](
		EnsureObjectImplementsInterface[FeatureFlags, EnsuredObject](f),
		EnsureObjectImplementsInterface[*FeatureFlags, DefaultableInterface[*FeatureFlags]](&f),
	)
}

func (f *FeatureFlags) SetIntFlag(key string, value int64) {
	f.IntFlags[key] = value
}

func (f *FeatureFlags) SetUintFlag(key string, value uint64) {
	f.UIntFlags[key] = value
}

func (f *FeatureFlags) SetStringFlag(key string, value string) {
	f.StringFlags[key] = value
}

func (f *FeatureFlags) SetBoolFlag(key string, value bool) {
	f.BoolFlags[key] = value
}

func (f *FeatureFlags) SetFloatFlag(key string, value float64) {
	f.FloatFlags[key] = value
}

func (f *FeatureFlags) SetAnyFlag(key string, value any) {
	f.AnyFlags[key] = value
}

func (f *FeatureFlags) GetIntFlag(key string) int64 {
	return f.IntFlags[key]
}

func (f *FeatureFlags) GetUintFlag(key string) uint64 {
	return f.UIntFlags[key]
}

func (f *FeatureFlags) GetStringFlag(key string) string {
	return f.StringFlags[key]
}

func (f *FeatureFlags) GetBoolFlag(key string) bool {
	return f.BoolFlags[key]
}

func (f *FeatureFlags) GetFloatFlag(key string) float64 {
	return f.FloatFlags[key]
}

func (f *FeatureFlags) GetAnyFlag(key string) any {
	return f.AnyFlags[key]
}

func (f *FeatureFlags) GetDefault() *FeatureFlags {
	f.IntFlags = make(map[string]int64)
	f.UIntFlags = make(map[string]uint64)
	f.StringFlags = make(map[string]string)
	f.BoolFlags = make(map[string]bool)
	f.FloatFlags = make(map[string]float64)
	f.AnyFlags = make(map[string]any)
	return f
}

func MakeFeatureFlags() *FeatureFlags {
	return &FeatureFlags{}
}

var featureFlags FeatureFlags = *MakeFeatureFlags().GetDefault()

func GetFeatureFlags() *FeatureFlags {
	return &featureFlags
}
