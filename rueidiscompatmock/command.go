package rueidiscompatmock

import (
	"net"
	"sort"
	"strconv"
	"time"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/mock"
	"github.com/redis/rueidis/rueidiscompat"
)

type ExpectedString struct{ exp *expectation }

func (e *ExpectedString) SetVal(v string) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisString(v))
}

func (e *ExpectedString) SetErr(err error) {
	e.exp.setErr(err)
}

func (e *ExpectedString) RedisNil() {
	e.exp.setRedisNil()
}

type ExpectedStatus = ExpectedString

type ExpectedError = ExpectedStatus

type ExpectedBool struct {
	exp    *expectation
	encode func(bool) rueidis.RedisMessage
}

func (e *ExpectedBool) SetVal(v bool) {
	if e.encode != nil {
		e.exp.resultSet = true
		e.exp.result = mock.Result(e.encode(v))
		return
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisBool(v))
}

func (e *ExpectedBool) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedInt struct{ exp *expectation }

func (e *ExpectedInt) SetVal(v int64) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisInt64(v))
}

func (e *ExpectedInt) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedDuration struct {
	exp       *expectation
	precision time.Duration
}

func (e *ExpectedDuration) SetVal(v time.Duration) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisInt64(int64(v / e.precision)))
}

func (e *ExpectedDuration) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedFloat struct{ exp *expectation }

func (e *ExpectedFloat) SetVal(v float64) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisFloat64(v))
}

func (e *ExpectedFloat) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedStringSlice struct{ exp *expectation }

func (e *ExpectedStringSlice) SetVal(v []string) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, s := range v {
		msgs = append(msgs, mock.RedisString(s))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedStringSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedIntSlice struct{ exp *expectation }

func (e *ExpectedIntSlice) SetVal(v []int64) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, n := range v {
		msgs = append(msgs, mock.RedisInt64(n))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedIntSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedBoolSlice struct{ exp *expectation }

func (e *ExpectedBoolSlice) SetVal(v []bool) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, b := range v {
		if b {
			msgs = append(msgs, mock.RedisInt64(1))
		} else {
			msgs = append(msgs, mock.RedisInt64(0))
		}
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedBoolSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedFloatSlice struct{ exp *expectation }

func (e *ExpectedFloatSlice) SetVal(v []float64) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, f := range v {
		msgs = append(msgs, mock.RedisFloat64(f))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedFloatSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedSlice struct{ exp *expectation }

func (e *ExpectedSlice) SetVal(v []any) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, item := range v {
		switch x := item.(type) {
		case nil:
			msgs = append(msgs, mock.RedisNil())
		case string:
			msgs = append(msgs, mock.RedisString(x))
		case int64:
			msgs = append(msgs, mock.RedisInt64(x))
		case bool:
			msgs = append(msgs, mock.RedisBool(x))
		default:
			msgs = append(msgs, mock.RedisString(str(item)))
		}
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedSlice) SetErr(err error) {
	e.exp.setErr(err)
}

func redisMessage(v any) rueidis.RedisMessage {
	switch x := v.(type) {
	case nil:
		return mock.RedisNil()
	case string:
		return mock.RedisString(x)
	case int:
		return mock.RedisInt64(int64(x))
	case int64:
		return mock.RedisInt64(x)
	case uint64:
		return mock.RedisString(itoa(int64(x)))
	case bool:
		return mock.RedisBool(x)
	case float64:
		return mock.RedisFloat64(x)
	case []any:
		msgs := make([]rueidis.RedisMessage, 0, len(x))
		for _, item := range x {
			msgs = append(msgs, redisMessage(item))
		}
		return mock.RedisArray(msgs...)
	case []string:
		msgs := make([]rueidis.RedisMessage, 0, len(x))
		for _, item := range x {
			msgs = append(msgs, mock.RedisString(item))
		}
		return mock.RedisArray(msgs...)
	default:
		return mock.RedisString(str(v))
	}
}

type ExpectedStringStringMap struct{ exp *expectation }

type ExpectedMapStringString = ExpectedStringStringMap

func (e *ExpectedStringStringMap) SetVal(v map[string]string) {
	kv := make(map[string]rueidis.RedisMessage, len(v))
	for k, val := range v {
		kv[k] = mock.RedisString(val)
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisMap(kv))
}

func (e *ExpectedStringStringMap) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedStringIntMap struct{ exp *expectation }

type ExpectedMapStringInt = ExpectedStringIntMap

func (e *ExpectedStringIntMap) SetVal(v map[string]int64) {
	kv := make(map[string]rueidis.RedisMessage, len(v))
	for k, val := range v {
		kv[k] = mock.RedisInt64(val)
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisMap(kv))
}

func (e *ExpectedStringIntMap) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedStringStructMap struct{ exp *expectation }

func (e *ExpectedStringStructMap) SetVal(v []string) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, s := range v {
		msgs = append(msgs, mock.RedisString(s))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedStringStructMap) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedKeyValueSlice struct{ exp *expectation }

func (e *ExpectedKeyValueSlice) SetVal(v []rueidiscompat.KeyValue) {
	pairs := make([]rueidis.RedisMessage, 0, len(v))
	for _, kv := range v {
		pairs = append(pairs, mock.RedisArray(mock.RedisString(kv.Key), mock.RedisString(kv.Value)))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(pairs...))
}

func (e *ExpectedKeyValueSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedKeyValues struct{ exp *expectation }

func (e *ExpectedKeyValues) SetVal(key string, vals []string) {
	msgs := make([]rueidis.RedisMessage, 0, len(vals))
	for _, v := range vals {
		msgs = append(msgs, mock.RedisString(v))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(mock.RedisString(key), mock.RedisArray(msgs...)))
}

func (e *ExpectedKeyValues) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedZSlice struct{ exp *expectation }

func (e *ExpectedZSlice) SetVal(v []rueidiscompat.Z) {
	msgs := make([]rueidis.RedisMessage, 0, len(v)*2)
	for _, z := range v {
		msgs = append(msgs, mock.RedisString(str(z.Member)), mock.RedisString(str(z.Score)))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedZSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedZWithKey struct{ exp *expectation }

func (e *ExpectedZWithKey) SetVal(v *rueidiscompat.ZWithKey) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(
		mock.RedisString(v.Key),
		mock.RedisString(str(v.Member)),
		mock.RedisString(str(v.Score)),
	))
}

func (e *ExpectedZWithKey) SetErr(err error) {
	e.exp.setErr(err)
}

func (e *ExpectedZWithKey) RedisNil() {
	e.exp.setRedisNil()
}

type ExpectedRankWithScore struct{ exp *expectation }

func (e *ExpectedRankWithScore) SetVal(v rueidiscompat.RankScore) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(
		mock.RedisInt64(v.Rank),
		mock.RedisFloat64(v.Score),
	))
}

func (e *ExpectedRankWithScore) SetErr(err error) {
	e.exp.setErr(err)
}

func (e *ExpectedRankWithScore) RedisNil() {
	e.exp.setRedisNil()
}

type ExpectedTime struct{ exp *expectation }

func (e *ExpectedTime) SetVal(v time.Time) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(
		mock.RedisString(itoa(v.Unix())),
		mock.RedisString(itoa(int64(v.Nanosecond()/1000))),
	))
}

func (e *ExpectedTime) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedScan struct{ exp *expectation }

func (e *ExpectedScan) SetVal(keys []string, cursor uint64) {
	msgs := make([]rueidis.RedisMessage, 0, len(keys))
	for _, k := range keys {
		msgs = append(msgs, mock.RedisString(k))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(
		mock.RedisString(itoa(int64(cursor))),
		mock.RedisArray(msgs...),
	))
}

func (e *ExpectedScan) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedGeoPos struct{ exp *expectation }

func (e *ExpectedGeoPos) SetVal(v []*rueidiscompat.GeoPos) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, p := range v {
		if p == nil {
			msgs = append(msgs, mock.RedisNil())
			continue
		}
		msgs = append(msgs, mock.RedisArray(
			mock.RedisFloat64(p.Longitude),
			mock.RedisFloat64(p.Latitude),
		))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedGeoPos) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedGeoLocation struct{ exp *expectation }

type ExpectedGeoSearchLocation = ExpectedGeoLocation

func (e *ExpectedGeoLocation) SetVal(v []rueidiscompat.GeoLocation) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, l := range v {
		parts := []rueidis.RedisMessage{mock.RedisString(l.Name)}
		if l.Dist != 0 {
			parts = append(parts, mock.RedisString(str(l.Dist)))
		}
		if l.GeoHash != 0 {
			parts = append(parts, mock.RedisInt64(l.GeoHash))
		}
		if l.Longitude != 0 || l.Latitude != 0 {
			parts = append(parts, mock.RedisArray(
				mock.RedisFloat64(l.Longitude),
				mock.RedisFloat64(l.Latitude),
			))
		}
		if len(parts) == 1 {
			msgs = append(msgs, parts[0])
			continue
		}
		msgs = append(msgs, mock.RedisArray(parts...))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedGeoLocation) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXMessageSlice struct{ exp *expectation }

func (e *ExpectedXMessageSlice) SetVal(v []rueidiscompat.XMessage) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, m := range v {
		msgs = append(msgs, xMessageEntry(m))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedXMessageSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXStreamSlice struct{ exp *expectation }

func (e *ExpectedXStreamSlice) SetVal(v []rueidiscompat.XStream) {
	streams := make([]rueidis.RedisMessage, 0, len(v))
	for _, s := range v {
		entries := make([]rueidis.RedisMessage, 0, len(s.Messages))
		for _, m := range s.Messages {
			entries = append(entries, xMessageEntry(m))
		}
		streams = append(streams, mock.RedisArray(
			mock.RedisString(s.Stream),
			mock.RedisArray(entries...),
		))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(streams...))
}

func (e *ExpectedXStreamSlice) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedCmd struct{ exp *expectation }

func (e *ExpectedCmd) SetVal(v any) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(redisMessage(v))
	e.exp.rawSet = true
	e.exp.rawResult = v
}

func (e *ExpectedCmd) SetValInt(v int64) *ExpectedCmd {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisInt64(v))
	return e
}

func (e *ExpectedCmd) SetErr(err error) {
	e.exp.setErr(err)
}

func (e *ExpectedCmd) RedisNil() {
	e.exp.setRedisNil()
}

type ExpectedCommandsInfo struct{ exp *expectation }

func (e *ExpectedCommandsInfo) SetVal(v []*rueidiscompat.CommandInfo) {
	entries := make([]rueidis.RedisMessage, 0, len(v))
	for _, info := range v {
		flags := info.Flags
		if flags == nil {
			flags = []string{}
		}
		flagMsgs := make([]rueidis.RedisMessage, len(flags))
		for i, f := range flags {
			flagMsgs[i] = mock.RedisString(f)
		}
		inner := []rueidis.RedisMessage{
			mock.RedisString(info.Name),
			mock.RedisInt64(info.Arity),
			mock.RedisArray(flagMsgs...),
			mock.RedisInt64(info.FirstKeyPos),
			mock.RedisInt64(info.LastKeyPos),
			mock.RedisInt64(info.StepCount),
		}
		if info.ACLFlags != nil {
			aclMsgs := make([]rueidis.RedisMessage, len(info.ACLFlags))
			for i, f := range info.ACLFlags {
				aclMsgs[i] = mock.RedisString(f)
			}
			inner = append(inner, mock.RedisArray(aclMsgs...))
		}
		entries = append(entries, mock.RedisArray(inner...))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(entries...))
}

func (e *ExpectedCommandsInfo) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedKeyFlags struct{ exp *expectation }

func (e *ExpectedKeyFlags) SetVal(v []rueidiscompat.KeyFlags) {
	entries := make([]rueidis.RedisMessage, 0, len(v))
	for _, kf := range v {
		flagMsgs := make([]rueidis.RedisMessage, len(kf.Flags))
		for i, f := range kf.Flags {
			flagMsgs[i] = mock.RedisString(f)
		}
		entries = append(entries, mock.RedisArray(
			mock.RedisString(kf.Key),
			mock.RedisArray(flagMsgs...),
		))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(entries...))
}

func (e *ExpectedKeyFlags) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedSlowLog struct{ exp *expectation }

func (e *ExpectedSlowLog) SetVal(v []rueidiscompat.SlowLog) {
	entries := make([]rueidis.RedisMessage, 0, len(v))
	for _, entry := range v {
		argMsgs := make([]rueidis.RedisMessage, len(entry.Args))
		for i, a := range entry.Args {
			argMsgs[i] = mock.RedisString(a)
		}
		inner := []rueidis.RedisMessage{
			mock.RedisInt64(entry.ID),
			mock.RedisInt64(entry.Time.Unix()),
			mock.RedisInt64(int64(entry.Duration / time.Microsecond)),
			mock.RedisArray(argMsgs...),
		}
		if entry.ClientAddr != "" {
			inner = append(inner, mock.RedisString(entry.ClientAddr))
		}
		if entry.ClientName != "" {
			if entry.ClientAddr == "" {
				inner = append(inner, mock.RedisString(""))
			}
			inner = append(inner, mock.RedisString(entry.ClientName))
		}
		entries = append(entries, mock.RedisArray(inner...))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(entries...))
}

func (e *ExpectedSlowLog) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedOKBool struct{ exp *expectation }

func (e *ExpectedOKBool) SetVal(v bool) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(okBoolMessage(v))
}

func (e *ExpectedOKBool) SetErr(err error) {
	e.exp.setErr(err)
}

func okBoolMessage(v bool) rueidis.RedisMessage {
	if v {
		return mock.RedisString("OK")
	}
	return mock.RedisString("")
}

type ExpectedClusterSlots struct{ exp *expectation }

func (e *ExpectedClusterSlots) SetVal(v []rueidiscompat.ClusterSlot) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(clusterSlotMessages(v)...))
}

func (e *ExpectedClusterSlots) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedClusterShards struct{ exp *expectation }

func (e *ExpectedClusterShards) SetVal(v []rueidiscompat.ClusterShard) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, shard := range v {
		msgs = append(msgs, clusterShardMessage(shard))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedClusterShards) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedClusterLinks struct{ exp *expectation }

func (e *ExpectedClusterLinks) SetVal(v []rueidiscompat.ClusterLink) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, link := range v {
		msgs = append(msgs, clusterLinkMessage(link))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedClusterLinks) SetErr(err error) {
	e.exp.setErr(err)
}

func clusterSlotMessages(slots []rueidiscompat.ClusterSlot) []rueidis.RedisMessage {
	out := make([]rueidis.RedisMessage, 0, len(slots))
	for _, slot := range slots {
		elems := []rueidis.RedisMessage{
			mock.RedisInt64(slot.Start),
			mock.RedisInt64(slot.End),
		}
		for _, node := range slot.Nodes {
			elems = append(elems, clusterNodeMessage(node))
		}
		out = append(out, mock.RedisArray(elems...))
	}
	return out
}

func clusterNodeMessage(node rueidiscompat.ClusterNode) rueidis.RedisMessage {
	host, portStr, err := net.SplitHostPort(node.Addr)
	if err != nil {
		host = node.Addr
		portStr = "0"
	}
	port, _ := strconv.ParseInt(portStr, 10, 64)
	if node.ID != "" {
		return mock.RedisArray(
			mock.RedisString(host),
			mock.RedisInt64(port),
			mock.RedisString(node.ID),
		)
	}
	return mock.RedisArray(
		mock.RedisString(host),
		mock.RedisInt64(port),
	)
}

func clusterShardMessage(shard rueidiscompat.ClusterShard) rueidis.RedisMessage {
	slotPairs := make([]rueidis.RedisMessage, 0, len(shard.Slots)*2)
	for _, sr := range shard.Slots {
		slotPairs = append(slotPairs, mock.RedisInt64(sr.Start), mock.RedisInt64(sr.End))
	}
	nodes := make([]rueidis.RedisMessage, 0, len(shard.Nodes))
	for _, n := range shard.Nodes {
		kv := map[string]rueidis.RedisMessage{}
		if n.ID != "" {
			kv["id"] = mock.RedisString(n.ID)
		}
		if n.Endpoint != "" {
			kv["endpoint"] = mock.RedisString(n.Endpoint)
		}
		if n.IP != "" {
			kv["ip"] = mock.RedisString(n.IP)
		}
		if n.Hostname != "" {
			kv["hostname"] = mock.RedisString(n.Hostname)
		}
		if n.Port != 0 {
			kv["port"] = mock.RedisInt64(n.Port)
		}
		if n.TLSPort != 0 {
			kv["tls-port"] = mock.RedisInt64(n.TLSPort)
		}
		if n.Role != "" {
			kv["role"] = mock.RedisString(n.Role)
		}
		if n.ReplicationOffset != 0 {
			kv["replication-offset"] = mock.RedisInt64(n.ReplicationOffset)
		}
		if n.Health != "" {
			kv["health"] = mock.RedisString(n.Health)
		}
		nodes = append(nodes, mock.RedisMap(kv))
	}
	return mock.RedisMap(map[string]rueidis.RedisMessage{
		"slots": mock.RedisArray(slotPairs...),
		"nodes": mock.RedisArray(nodes...),
	})
}

func clusterLinkMessage(link rueidiscompat.ClusterLink) rueidis.RedisMessage {
	return mock.RedisMap(map[string]rueidis.RedisMessage{
		"direction":             mock.RedisString(link.Direction),
		"node":                  mock.RedisString(link.Node),
		"create-time":           mock.RedisInt64(link.CreateTime),
		"events":                mock.RedisString(link.Events),
		"send-buffer-allocated": mock.RedisInt64(link.SendBufferAllocated),
		"send-buffer-used":      mock.RedisInt64(link.SendBufferUsed),
	})
}

type ExpectedLCS struct {
	exp          *expectation
	readType     uint8
	withMatchLen bool
}

func (e *ExpectedLCS) SetVal(v *rueidiscompat.LCSMatch) {
	switch e.readType {
	case 2:
		e.exp.resultSet = true
		e.exp.result = mock.Result(mock.RedisInt64(v.Len))
	case 3:
		matches := make([]rueidis.RedisMessage, len(v.Matches))
		for i, mp := range v.Matches {
			inner := []rueidis.RedisMessage{
				mock.RedisArray(mock.RedisInt64(mp.Key1.Start), mock.RedisInt64(mp.Key1.End)),
				mock.RedisArray(mock.RedisInt64(mp.Key2.Start), mock.RedisInt64(mp.Key2.End)),
			}
			if e.withMatchLen {
				inner = append(inner, mock.RedisInt64(mp.MatchLen))
			}
			matches[i] = mock.RedisArray(inner...)
		}
		e.exp.resultSet = true
		e.exp.result = mock.Result(mock.RedisMap(map[string]rueidis.RedisMessage{
			"matches": mock.RedisArray(matches...),
			"len":     mock.RedisInt64(v.Len),
		}))
	default:
		e.exp.resultSet = true
		e.exp.result = mock.Result(mock.RedisString(v.MatchString))
	}
}

func (e *ExpectedLCS) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedZSliceWithKey struct{ exp *expectation }

func (e *ExpectedZSliceWithKey) SetVal(key string, vals []rueidiscompat.Z) {
	msgs := make([]rueidis.RedisMessage, 0, len(vals)*2)
	for _, z := range vals {
		msgs = append(msgs, mock.RedisString(str(z.Member)), mock.RedisString(str(z.Score)))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(mock.RedisString(key), mock.RedisArray(msgs...)))
}

func (e *ExpectedZSliceWithKey) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedScriptExists struct{ exp *expectation }

func (e *ExpectedScriptExists) SetVal(v []bool) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, b := range v {
		if b {
			msgs = append(msgs, mock.RedisInt64(1))
		} else {
			msgs = append(msgs, mock.RedisInt64(0))
		}
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedScriptExists) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedFunctionList struct{ exp *expectation }

func (e *ExpectedFunctionList) SetVal(v []rueidiscompat.Library) {
	entries := make([]rueidis.RedisMessage, 0, len(v))
	for _, lib := range v {
		kv := map[string]rueidis.RedisMessage{
			"library_name": mock.RedisString(lib.Name),
			"engine":       mock.RedisString(lib.Engine),
		}
		if lib.Code != "" {
			kv["library_code"] = mock.RedisString(lib.Code)
		}
		fns := make([]rueidis.RedisMessage, 0, len(lib.Functions))
		for _, fn := range lib.Functions {
			flagMsgs := make([]rueidis.RedisMessage, len(fn.Flags))
			for i, f := range fn.Flags {
				flagMsgs[i] = mock.RedisString(f)
			}
			fkv := map[string]rueidis.RedisMessage{
				"name":        mock.RedisString(fn.Name),
				"description": mock.RedisString(fn.Description),
				"flags":       mock.RedisArray(flagMsgs...),
			}
			fns = append(fns, mock.RedisMap(fkv))
		}
		kv["functions"] = mock.RedisArray(fns...)
		entries = append(entries, mock.RedisMap(kv))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(entries...))
}

func (e *ExpectedFunctionList) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXAutoClaim struct{ exp *expectation }

func (e *ExpectedXAutoClaim) SetVal(messages []rueidiscompat.XMessage, start string) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(
		mock.RedisString(start),
		xMessagesRedisArray(messages),
	))
}

func (e *ExpectedXAutoClaim) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXAutoClaimJustID struct{ exp *expectation }

func (e *ExpectedXAutoClaimJustID) SetVal(ids []string, start string) {
	msgs := make([]rueidis.RedisMessage, 0, len(ids))
	for _, id := range ids {
		msgs = append(msgs, mock.RedisString(id))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(
		mock.RedisString(start),
		mock.RedisArray(msgs...),
	))
}

func (e *ExpectedXAutoClaimJustID) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXInfoConsumers struct{ exp *expectation }

func (e *ExpectedXInfoConsumers) SetVal(v []rueidiscompat.XInfoConsumer) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, c := range v {
		kv := map[string]rueidis.RedisMessage{}
		if c.Name != "" {
			kv["name"] = mock.RedisString(c.Name)
		}
		if c.Pending != 0 {
			kv["pending"] = mock.RedisInt64(c.Pending)
		}
		if c.Idle != 0 {
			kv["idle"] = mock.RedisInt64(int64(c.Idle / time.Millisecond))
		}
		msgs = append(msgs, mock.RedisMap(kv))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedXInfoConsumers) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXInfoGroups struct{ exp *expectation }

func (e *ExpectedXInfoGroups) SetVal(v []rueidiscompat.XInfoGroup) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, g := range v {
		msgs = append(msgs, xInfoGroupMessage(g))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedXInfoGroups) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXInfoStream struct{ exp *expectation }

func (e *ExpectedXInfoStream) SetVal(v *rueidiscompat.XInfoStream) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(xInfoStreamMessage(*v))
}

func (e *ExpectedXInfoStream) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXInfoStreamFull struct{ exp *expectation }

func (e *ExpectedXInfoStreamFull) SetVal(v *rueidiscompat.XInfoStreamFull) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(xInfoStreamFullMessage(*v))
}

func (e *ExpectedXInfoStreamFull) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXPending struct{ exp *expectation }

func (e *ExpectedXPending) SetVal(v *rueidiscompat.XPending) {
	e.exp.resultSet = true
	e.exp.result = mock.Result(xPendingMessage(*v))
}

func (e *ExpectedXPending) SetErr(err error) {
	e.exp.setErr(err)
}

type ExpectedXPendingExt struct{ exp *expectation }

func (e *ExpectedXPendingExt) SetVal(v []rueidiscompat.XPendingExt) {
	msgs := make([]rueidis.RedisMessage, 0, len(v))
	for _, p := range v {
		msgs = append(msgs, mock.RedisArray(
			mock.RedisString(p.ID),
			mock.RedisString(p.Consumer),
			mock.RedisInt64(int64(p.Idle/time.Millisecond)),
			mock.RedisInt64(p.RetryCount),
		))
	}
	e.exp.resultSet = true
	e.exp.result = mock.Result(mock.RedisArray(msgs...))
}

func (e *ExpectedXPendingExt) SetErr(err error) {
	e.exp.setErr(err)
}

func xMessagesRedisArray(msgs []rueidiscompat.XMessage) rueidis.RedisMessage {
	entries := make([]rueidis.RedisMessage, 0, len(msgs))
	for _, m := range msgs {
		entries = append(entries, xMessageEntry(m))
	}
	return mock.RedisArray(entries...)
}

func xMessageEntry(m rueidiscompat.XMessage) rueidis.RedisMessage {
	keys := make([]string, 0, len(m.Values))
	for k := range m.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fields := make([]rueidis.RedisMessage, 0, len(keys)*2)
	for _, k := range keys {
		fields = append(fields, mock.RedisString(k), mock.RedisString(str(m.Values[k])))
	}
	return mock.RedisArray(mock.RedisString(m.ID), mock.RedisArray(fields...))
}

func xPendingMessage(v rueidiscompat.XPending) rueidis.RedisMessage {
	names := make([]string, 0, len(v.Consumers))
	for name := range v.Consumers {
		names = append(names, name)
	}
	sort.Strings(names)
	consumers := make([]rueidis.RedisMessage, 0, len(names))
	for _, name := range names {
		consumers = append(consumers, mock.RedisArray(
			mock.RedisString(name),
			mock.RedisInt64(v.Consumers[name]),
		))
	}
	return mock.RedisArray(
		mock.RedisInt64(v.Count),
		mock.RedisString(v.Lower),
		mock.RedisString(v.Higher),
		mock.RedisArray(consumers...),
	)
}

func xInfoGroupMessage(g rueidiscompat.XInfoGroup) rueidis.RedisMessage {
	kv := map[string]rueidis.RedisMessage{}
	if g.Name != "" {
		kv["name"] = mock.RedisString(g.Name)
	}
	if g.Consumers != 0 {
		kv["consumers"] = mock.RedisInt64(g.Consumers)
	}
	if g.Pending != 0 {
		kv["pending"] = mock.RedisInt64(g.Pending)
	}
	if g.EntriesRead != 0 {
		kv["entries-read"] = mock.RedisInt64(g.EntriesRead)
	}
	if g.Lag != 0 {
		kv["lag"] = mock.RedisInt64(g.Lag)
	}
	if g.LastDeliveredID != "" {
		kv["last-delivered-id"] = mock.RedisString(g.LastDeliveredID)
	}
	return mock.RedisMap(kv)
}

func xInfoStreamMessage(v rueidiscompat.XInfoStream) rueidis.RedisMessage {
	kv := map[string]rueidis.RedisMessage{}
	if v.Length != 0 {
		kv["length"] = mock.RedisInt64(v.Length)
	}
	if v.RadixTreeKeys != 0 {
		kv["radix-tree-keys"] = mock.RedisInt64(v.RadixTreeKeys)
	}
	if v.RadixTreeNodes != 0 {
		kv["radix-tree-nodes"] = mock.RedisInt64(v.RadixTreeNodes)
	}
	if v.Groups != 0 {
		kv["groups"] = mock.RedisInt64(v.Groups)
	}
	if v.EntriesAdded != 0 {
		kv["entries-added"] = mock.RedisInt64(v.EntriesAdded)
	}
	if v.IDMPDuration != 0 {
		kv["idmp-duration"] = mock.RedisInt64(v.IDMPDuration)
	}
	if v.IDMPMaxSize != 0 {
		kv["idmp-maxsize"] = mock.RedisInt64(v.IDMPMaxSize)
	}
	if v.PIDsTracked != 0 {
		kv["pids-tracked"] = mock.RedisInt64(v.PIDsTracked)
	}
	if v.IIDsTracked != 0 {
		kv["iids-tracked"] = mock.RedisInt64(v.IIDsTracked)
	}
	if v.IIDsAdded != 0 {
		kv["iids-added"] = mock.RedisInt64(v.IIDsAdded)
	}
	if v.IIDsDuplicates != 0 {
		kv["iids-duplicates"] = mock.RedisInt64(v.IIDsDuplicates)
	}
	if v.LastGeneratedID != "" {
		kv["last-generated-id"] = mock.RedisString(v.LastGeneratedID)
	}
	if v.MaxDeletedEntryID != "" {
		kv["max-deleted-entry-id"] = mock.RedisString(v.MaxDeletedEntryID)
	}
	if v.RecordedFirstEntryID != "" {
		kv["recorded-first-entry-id"] = mock.RedisString(v.RecordedFirstEntryID)
	}
	if v.FirstEntry.ID != "" || len(v.FirstEntry.Values) > 0 {
		kv["first-entry"] = xMessageEntry(v.FirstEntry)
	}
	if v.LastEntry.ID != "" || len(v.LastEntry.Values) > 0 {
		kv["last-entry"] = xMessageEntry(v.LastEntry)
	}
	return mock.RedisMap(kv)
}

func xInfoStreamFullMessage(v rueidiscompat.XInfoStreamFull) rueidis.RedisMessage {
	kv := map[string]rueidis.RedisMessage{}
	if v.Length != 0 {
		kv["length"] = mock.RedisInt64(v.Length)
	}
	if v.RadixTreeKeys != 0 {
		kv["radix-tree-keys"] = mock.RedisInt64(v.RadixTreeKeys)
	}
	if v.RadixTreeNodes != 0 {
		kv["radix-tree-nodes"] = mock.RedisInt64(v.RadixTreeNodes)
	}
	if v.EntriesAdded != 0 {
		kv["entries-added"] = mock.RedisInt64(v.EntriesAdded)
	}
	if v.IDMPDuration != 0 {
		kv["idmp-duration"] = mock.RedisInt64(v.IDMPDuration)
	}
	if v.IDMPMaxSize != 0 {
		kv["idmp-maxsize"] = mock.RedisInt64(v.IDMPMaxSize)
	}
	if v.PIDsTracked != 0 {
		kv["pids-tracked"] = mock.RedisInt64(v.PIDsTracked)
	}
	if v.IIDsTracked != 0 {
		kv["iids-tracked"] = mock.RedisInt64(v.IIDsTracked)
	}
	if v.IIDsAdded != 0 {
		kv["iids-added"] = mock.RedisInt64(v.IIDsAdded)
	}
	if v.IIDsDuplicates != 0 {
		kv["iids-duplicates"] = mock.RedisInt64(v.IIDsDuplicates)
	}
	if v.LastGeneratedID != "" {
		kv["last-generated-id"] = mock.RedisString(v.LastGeneratedID)
	}
	if v.MaxDeletedEntryID != "" {
		kv["max-deleted-entry-id"] = mock.RedisString(v.MaxDeletedEntryID)
	}
	if v.RecordedFirstEntryID != "" {
		kv["recorded-first-entry-id"] = mock.RedisString(v.RecordedFirstEntryID)
	}
	if len(v.Entries) > 0 {
		kv["entries"] = xMessagesRedisArray(v.Entries)
	}
	if len(v.Groups) > 0 {
		groupMsgs := make([]rueidis.RedisMessage, 0, len(v.Groups))
		for _, g := range v.Groups {
			groupMsgs = append(groupMsgs, xInfoStreamFullGroupMessage(g))
		}
		kv["groups"] = mock.RedisArray(groupMsgs...)
	}
	return mock.RedisMap(kv)
}

func xInfoStreamFullGroupMessage(g rueidiscompat.XInfoStreamGroup) rueidis.RedisMessage {
	kv := map[string]rueidis.RedisMessage{}
	if g.Name != "" {
		kv["name"] = mock.RedisString(g.Name)
	}
	if g.LastDeliveredID != "" {
		kv["last-delivered-id"] = mock.RedisString(g.LastDeliveredID)
	}
	if g.EntriesRead != 0 {
		kv["entries-read"] = mock.RedisInt64(g.EntriesRead)
	}
	if g.Lag != 0 {
		kv["lag"] = mock.RedisInt64(g.Lag)
	}
	if g.PelCount != 0 {
		kv["pel-count"] = mock.RedisInt64(g.PelCount)
	}
	if len(g.Pending) > 0 {
		pending := make([]rueidis.RedisMessage, 0, len(g.Pending))
		for _, p := range g.Pending {
			pending = append(pending, mock.RedisArray(
				mock.RedisString(p.ID),
				mock.RedisString(p.Consumer),
				mock.RedisInt64(p.DeliveryTime.UnixMilli()),
				mock.RedisInt64(p.DeliveryCount),
			))
		}
		kv["pending"] = mock.RedisArray(pending...)
	}
	if len(g.Consumers) > 0 {
		consumers := make([]rueidis.RedisMessage, 0, len(g.Consumers))
		for _, c := range g.Consumers {
			consumerKV := map[string]rueidis.RedisMessage{}
			if c.Name != "" {
				consumerKV["name"] = mock.RedisString(c.Name)
			}
			if !c.SeenTime.IsZero() {
				consumerKV["seen-time"] = mock.RedisInt64(c.SeenTime.UnixMilli())
			}
			if c.PelCount != 0 {
				consumerKV["pel-count"] = mock.RedisInt64(c.PelCount)
			}
			if len(c.Pending) > 0 {
				pending := make([]rueidis.RedisMessage, 0, len(c.Pending))
				for _, p := range c.Pending {
					pending = append(pending, mock.RedisArray(
						mock.RedisString(p.ID),
						mock.RedisInt64(p.DeliveryTime.UnixMilli()),
						mock.RedisInt64(p.DeliveryCount),
					))
				}
				consumerKV["pending"] = mock.RedisArray(pending...)
			}
			consumers = append(consumers, mock.RedisMap(consumerKV))
		}
		kv["consumers"] = mock.RedisArray(consumers...)
	}
	return mock.RedisMap(kv)
}

func itoa(n int64) string {
	return strconv.FormatInt(n, 10)
}
