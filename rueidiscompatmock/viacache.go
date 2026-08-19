package rueidiscompatmock

// ViaCache marks this expectation so it is only satisfied by a call routed
// through DoCache (i.e. rueidiscompat.Cache(ttl).<Cmd>()). A plain, non-cached
// call for the same command will not match it. Expectations that do not call
// ViaCache are unaffected and continue to match calls from either origin.

func (e *ExpectedString) ViaCache() *ExpectedString {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedBool) ViaCache() *ExpectedBool {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedInt) ViaCache() *ExpectedInt {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedDuration) ViaCache() *ExpectedDuration {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedFloat) ViaCache() *ExpectedFloat {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedStringSlice) ViaCache() *ExpectedStringSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedIntSlice) ViaCache() *ExpectedIntSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedBoolSlice) ViaCache() *ExpectedBoolSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedFloatSlice) ViaCache() *ExpectedFloatSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedSlice) ViaCache() *ExpectedSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedStringStringMap) ViaCache() *ExpectedStringStringMap {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedStringIntMap) ViaCache() *ExpectedStringIntMap {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedStringStructMap) ViaCache() *ExpectedStringStructMap {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedKeyValueSlice) ViaCache() *ExpectedKeyValueSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedKeyValues) ViaCache() *ExpectedKeyValues {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedZSlice) ViaCache() *ExpectedZSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedZWithKey) ViaCache() *ExpectedZWithKey {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedRankWithScore) ViaCache() *ExpectedRankWithScore {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedTime) ViaCache() *ExpectedTime {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedScan) ViaCache() *ExpectedScan {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedGeoPos) ViaCache() *ExpectedGeoPos {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedGeoLocation) ViaCache() *ExpectedGeoLocation {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXMessageSlice) ViaCache() *ExpectedXMessageSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXStreamSlice) ViaCache() *ExpectedXStreamSlice {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedCmd) ViaCache() *ExpectedCmd {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedCommandsInfo) ViaCache() *ExpectedCommandsInfo {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedKeyFlags) ViaCache() *ExpectedKeyFlags {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedSlowLog) ViaCache() *ExpectedSlowLog {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedOKBool) ViaCache() *ExpectedOKBool {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedClusterSlots) ViaCache() *ExpectedClusterSlots {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedClusterShards) ViaCache() *ExpectedClusterShards {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedClusterLinks) ViaCache() *ExpectedClusterLinks {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedLCS) ViaCache() *ExpectedLCS {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedZSliceWithKey) ViaCache() *ExpectedZSliceWithKey {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedScriptExists) ViaCache() *ExpectedScriptExists {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedFunctionList) ViaCache() *ExpectedFunctionList {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXAutoClaim) ViaCache() *ExpectedXAutoClaim {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXAutoClaimJustID) ViaCache() *ExpectedXAutoClaimJustID {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXInfoConsumers) ViaCache() *ExpectedXInfoConsumers {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXInfoGroups) ViaCache() *ExpectedXInfoGroups {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXInfoStream) ViaCache() *ExpectedXInfoStream {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXInfoStreamFull) ViaCache() *ExpectedXInfoStreamFull {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXPending) ViaCache() *ExpectedXPending {
	e.exp.requireCache = true
	return e
}

func (e *ExpectedXPendingExt) ViaCache() *ExpectedXPendingExt {
	e.exp.requireCache = true
	return e
}
