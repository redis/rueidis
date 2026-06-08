// Code generated DO NOT EDIT

package cmds

import "strconv"

type Arcount Incomplete

func (b Builder) Arcount() (c Arcount) {
	c = Arcount{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARCOUNT")
	return c
}

func (c Arcount) Key(key string) ArcountKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArcountKey)(c)
}

type ArcountKey Incomplete

func (c ArcountKey) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type Ardel Incomplete

func (b Builder) Ardel() (c Ardel) {
	c = Ardel{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "ARDEL")
	return c
}

func (c Ardel) Key(key string) ArdelKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArdelKey)(c)
}

type ArdelIndex Incomplete

func (c ArdelIndex) Index(index ...int64) ArdelIndex {
	for _, n := range index {
		c.cs.s = append(c.cs.s, strconv.FormatInt(n, 10))
	}
	return c
}

func (c ArdelIndex) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArdelKey Incomplete

func (c ArdelKey) Index(index ...int64) ArdelIndex {
	for _, n := range index {
		c.cs.s = append(c.cs.s, strconv.FormatInt(n, 10))
	}
	return (ArdelIndex)(c)
}

type Ardelrange Incomplete

func (b Builder) Ardelrange() (c Ardelrange) {
	c = Ardelrange{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "ARDELRANGE")
	return c
}

func (c Ardelrange) Key(key string) ArdelrangeKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArdelrangeKey)(c)
}

type ArdelrangeKey Incomplete

func (c ArdelrangeKey) Start(start int64) ArdelrangeRangeStart {
	c.cs.s = append(c.cs.s, strconv.FormatInt(start, 10))
	return (ArdelrangeRangeStart)(c)
}

type ArdelrangeRangeEnd Incomplete

func (c ArdelrangeRangeEnd) Start(start int64) ArdelrangeRangeStart {
	c.cs.s = append(c.cs.s, strconv.FormatInt(start, 10))
	return (ArdelrangeRangeStart)(c)
}

func (c ArdelrangeRangeEnd) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArdelrangeRangeStart Incomplete

func (c ArdelrangeRangeStart) End(end int64) ArdelrangeRangeEnd {
	c.cs.s = append(c.cs.s, strconv.FormatInt(end, 10))
	return (ArdelrangeRangeEnd)(c)
}

type Arget Incomplete

func (b Builder) Arget() (c Arget) {
	c = Arget{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARGET")
	return c
}

func (c Arget) Key(key string) ArgetKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArgetKey)(c)
}

type ArgetIndex Incomplete

func (c ArgetIndex) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgetKey Incomplete

func (c ArgetKey) Index(index int64) ArgetIndex {
	c.cs.s = append(c.cs.s, strconv.FormatInt(index, 10))
	return (ArgetIndex)(c)
}

type Argetrange Incomplete

func (b Builder) Argetrange() (c Argetrange) {
	c = Argetrange{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARGETRANGE")
	return c
}

func (c Argetrange) Key(key string) ArgetrangeKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArgetrangeKey)(c)
}

type ArgetrangeEnd Incomplete

func (c ArgetrangeEnd) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgetrangeKey Incomplete

func (c ArgetrangeKey) Start(start int64) ArgetrangeStart {
	c.cs.s = append(c.cs.s, strconv.FormatInt(start, 10))
	return (ArgetrangeStart)(c)
}

type ArgetrangeStart Incomplete

func (c ArgetrangeStart) End(end int64) ArgetrangeEnd {
	c.cs.s = append(c.cs.s, strconv.FormatInt(end, 10))
	return (ArgetrangeEnd)(c)
}

type Argrep Incomplete

func (b Builder) Argrep() (c Argrep) {
	c = Argrep{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARGREP")
	return c
}

func (c Argrep) Key(key string) ArgrepKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArgrepKey)(c)
}

type ArgrepEnd Incomplete

func (c ArgrepEnd) Exact() ArgrepPredicateExactExact {
	c.cs.s = append(c.cs.s, "EXACT")
	return (ArgrepPredicateExactExact)(c)
}

func (c ArgrepEnd) Match() ArgrepPredicateMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (ArgrepPredicateMatchMatch)(c)
}

func (c ArgrepEnd) Glob() ArgrepPredicateGlobGlob {
	c.cs.s = append(c.cs.s, "GLOB")
	return (ArgrepPredicateGlobGlob)(c)
}

func (c ArgrepEnd) Re() ArgrepPredicateReRe {
	c.cs.s = append(c.cs.s, "RE")
	return (ArgrepPredicateReRe)(c)
}

type ArgrepKey Incomplete

func (c ArgrepKey) Start(start string) ArgrepStart {
	c.cs.s = append(c.cs.s, start)
	return (ArgrepStart)(c)
}

type ArgrepOptionsAnd Incomplete

func (c ArgrepOptionsAnd) Exact() ArgrepPredicateExactExact {
	c.cs.s = append(c.cs.s, "EXACT")
	return (ArgrepPredicateExactExact)(c)
}

func (c ArgrepOptionsAnd) Match() ArgrepPredicateMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (ArgrepPredicateMatchMatch)(c)
}

func (c ArgrepOptionsAnd) Glob() ArgrepPredicateGlobGlob {
	c.cs.s = append(c.cs.s, "GLOB")
	return (ArgrepPredicateGlobGlob)(c)
}

func (c ArgrepOptionsAnd) Re() ArgrepPredicateReRe {
	c.cs.s = append(c.cs.s, "RE")
	return (ArgrepPredicateReRe)(c)
}

type ArgrepOptionsLimit Incomplete

func (c ArgrepOptionsLimit) And() ArgrepOptionsAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (ArgrepOptionsAnd)(c)
}

func (c ArgrepOptionsLimit) Or() ArgrepOptionsOr {
	c.cs.s = append(c.cs.s, "OR")
	return (ArgrepOptionsOr)(c)
}

func (c ArgrepOptionsLimit) Limit(limit int64) ArgrepOptionsLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return c
}

func (c ArgrepOptionsLimit) Withvalues() ArgrepOptionsWithvalues {
	c.cs.s = append(c.cs.s, "WITHVALUES")
	return (ArgrepOptionsWithvalues)(c)
}

func (c ArgrepOptionsLimit) Nocase() ArgrepOptionsNocase {
	c.cs.s = append(c.cs.s, "NOCASE")
	return (ArgrepOptionsNocase)(c)
}

func (c ArgrepOptionsLimit) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgrepOptionsNocase Incomplete

func (c ArgrepOptionsNocase) And() ArgrepOptionsAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (ArgrepOptionsAnd)(c)
}

func (c ArgrepOptionsNocase) Or() ArgrepOptionsOr {
	c.cs.s = append(c.cs.s, "OR")
	return (ArgrepOptionsOr)(c)
}

func (c ArgrepOptionsNocase) Limit(limit int64) ArgrepOptionsLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return (ArgrepOptionsLimit)(c)
}

func (c ArgrepOptionsNocase) Withvalues() ArgrepOptionsWithvalues {
	c.cs.s = append(c.cs.s, "WITHVALUES")
	return (ArgrepOptionsWithvalues)(c)
}

func (c ArgrepOptionsNocase) Nocase() ArgrepOptionsNocase {
	c.cs.s = append(c.cs.s, "NOCASE")
	return c
}

func (c ArgrepOptionsNocase) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgrepOptionsOr Incomplete

func (c ArgrepOptionsOr) Exact() ArgrepPredicateExactExact {
	c.cs.s = append(c.cs.s, "EXACT")
	return (ArgrepPredicateExactExact)(c)
}

func (c ArgrepOptionsOr) Match() ArgrepPredicateMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (ArgrepPredicateMatchMatch)(c)
}

func (c ArgrepOptionsOr) Glob() ArgrepPredicateGlobGlob {
	c.cs.s = append(c.cs.s, "GLOB")
	return (ArgrepPredicateGlobGlob)(c)
}

func (c ArgrepOptionsOr) Re() ArgrepPredicateReRe {
	c.cs.s = append(c.cs.s, "RE")
	return (ArgrepPredicateReRe)(c)
}

type ArgrepOptionsWithvalues Incomplete

func (c ArgrepOptionsWithvalues) And() ArgrepOptionsAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (ArgrepOptionsAnd)(c)
}

func (c ArgrepOptionsWithvalues) Or() ArgrepOptionsOr {
	c.cs.s = append(c.cs.s, "OR")
	return (ArgrepOptionsOr)(c)
}

func (c ArgrepOptionsWithvalues) Limit(limit int64) ArgrepOptionsLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return (ArgrepOptionsLimit)(c)
}

func (c ArgrepOptionsWithvalues) Withvalues() ArgrepOptionsWithvalues {
	c.cs.s = append(c.cs.s, "WITHVALUES")
	return c
}

func (c ArgrepOptionsWithvalues) Nocase() ArgrepOptionsNocase {
	c.cs.s = append(c.cs.s, "NOCASE")
	return (ArgrepOptionsNocase)(c)
}

func (c ArgrepOptionsWithvalues) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgrepPredicateExactExact Incomplete

func (c ArgrepPredicateExactExact) String(string string) ArgrepPredicateExactString {
	c.cs.s = append(c.cs.s, string)
	return (ArgrepPredicateExactString)(c)
}

type ArgrepPredicateExactString Incomplete

func (c ArgrepPredicateExactString) Exact() ArgrepPredicateExactExact {
	c.cs.s = append(c.cs.s, "EXACT")
	return (ArgrepPredicateExactExact)(c)
}

func (c ArgrepPredicateExactString) Match() ArgrepPredicateMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (ArgrepPredicateMatchMatch)(c)
}

func (c ArgrepPredicateExactString) Glob() ArgrepPredicateGlobGlob {
	c.cs.s = append(c.cs.s, "GLOB")
	return (ArgrepPredicateGlobGlob)(c)
}

func (c ArgrepPredicateExactString) Re() ArgrepPredicateReRe {
	c.cs.s = append(c.cs.s, "RE")
	return (ArgrepPredicateReRe)(c)
}

func (c ArgrepPredicateExactString) And() ArgrepOptionsAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (ArgrepOptionsAnd)(c)
}

func (c ArgrepPredicateExactString) Or() ArgrepOptionsOr {
	c.cs.s = append(c.cs.s, "OR")
	return (ArgrepOptionsOr)(c)
}

func (c ArgrepPredicateExactString) Limit(limit int64) ArgrepOptionsLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return (ArgrepOptionsLimit)(c)
}

func (c ArgrepPredicateExactString) Withvalues() ArgrepOptionsWithvalues {
	c.cs.s = append(c.cs.s, "WITHVALUES")
	return (ArgrepOptionsWithvalues)(c)
}

func (c ArgrepPredicateExactString) Nocase() ArgrepOptionsNocase {
	c.cs.s = append(c.cs.s, "NOCASE")
	return (ArgrepOptionsNocase)(c)
}

func (c ArgrepPredicateExactString) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgrepPredicateGlobGlob Incomplete

func (c ArgrepPredicateGlobGlob) Pattern(pattern string) ArgrepPredicateGlobPattern {
	c.cs.s = append(c.cs.s, pattern)
	return (ArgrepPredicateGlobPattern)(c)
}

type ArgrepPredicateGlobPattern Incomplete

func (c ArgrepPredicateGlobPattern) Exact() ArgrepPredicateExactExact {
	c.cs.s = append(c.cs.s, "EXACT")
	return (ArgrepPredicateExactExact)(c)
}

func (c ArgrepPredicateGlobPattern) Match() ArgrepPredicateMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (ArgrepPredicateMatchMatch)(c)
}

func (c ArgrepPredicateGlobPattern) Glob() ArgrepPredicateGlobGlob {
	c.cs.s = append(c.cs.s, "GLOB")
	return (ArgrepPredicateGlobGlob)(c)
}

func (c ArgrepPredicateGlobPattern) Re() ArgrepPredicateReRe {
	c.cs.s = append(c.cs.s, "RE")
	return (ArgrepPredicateReRe)(c)
}

func (c ArgrepPredicateGlobPattern) And() ArgrepOptionsAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (ArgrepOptionsAnd)(c)
}

func (c ArgrepPredicateGlobPattern) Or() ArgrepOptionsOr {
	c.cs.s = append(c.cs.s, "OR")
	return (ArgrepOptionsOr)(c)
}

func (c ArgrepPredicateGlobPattern) Limit(limit int64) ArgrepOptionsLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return (ArgrepOptionsLimit)(c)
}

func (c ArgrepPredicateGlobPattern) Withvalues() ArgrepOptionsWithvalues {
	c.cs.s = append(c.cs.s, "WITHVALUES")
	return (ArgrepOptionsWithvalues)(c)
}

func (c ArgrepPredicateGlobPattern) Nocase() ArgrepOptionsNocase {
	c.cs.s = append(c.cs.s, "NOCASE")
	return (ArgrepOptionsNocase)(c)
}

func (c ArgrepPredicateGlobPattern) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgrepPredicateMatchMatch Incomplete

func (c ArgrepPredicateMatchMatch) String(string string) ArgrepPredicateMatchString {
	c.cs.s = append(c.cs.s, string)
	return (ArgrepPredicateMatchString)(c)
}

type ArgrepPredicateMatchString Incomplete

func (c ArgrepPredicateMatchString) Exact() ArgrepPredicateExactExact {
	c.cs.s = append(c.cs.s, "EXACT")
	return (ArgrepPredicateExactExact)(c)
}

func (c ArgrepPredicateMatchString) Match() ArgrepPredicateMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (ArgrepPredicateMatchMatch)(c)
}

func (c ArgrepPredicateMatchString) Glob() ArgrepPredicateGlobGlob {
	c.cs.s = append(c.cs.s, "GLOB")
	return (ArgrepPredicateGlobGlob)(c)
}

func (c ArgrepPredicateMatchString) Re() ArgrepPredicateReRe {
	c.cs.s = append(c.cs.s, "RE")
	return (ArgrepPredicateReRe)(c)
}

func (c ArgrepPredicateMatchString) And() ArgrepOptionsAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (ArgrepOptionsAnd)(c)
}

func (c ArgrepPredicateMatchString) Or() ArgrepOptionsOr {
	c.cs.s = append(c.cs.s, "OR")
	return (ArgrepOptionsOr)(c)
}

func (c ArgrepPredicateMatchString) Limit(limit int64) ArgrepOptionsLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return (ArgrepOptionsLimit)(c)
}

func (c ArgrepPredicateMatchString) Withvalues() ArgrepOptionsWithvalues {
	c.cs.s = append(c.cs.s, "WITHVALUES")
	return (ArgrepOptionsWithvalues)(c)
}

func (c ArgrepPredicateMatchString) Nocase() ArgrepOptionsNocase {
	c.cs.s = append(c.cs.s, "NOCASE")
	return (ArgrepOptionsNocase)(c)
}

func (c ArgrepPredicateMatchString) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgrepPredicateRePattern Incomplete

func (c ArgrepPredicateRePattern) Exact() ArgrepPredicateExactExact {
	c.cs.s = append(c.cs.s, "EXACT")
	return (ArgrepPredicateExactExact)(c)
}

func (c ArgrepPredicateRePattern) Match() ArgrepPredicateMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (ArgrepPredicateMatchMatch)(c)
}

func (c ArgrepPredicateRePattern) Glob() ArgrepPredicateGlobGlob {
	c.cs.s = append(c.cs.s, "GLOB")
	return (ArgrepPredicateGlobGlob)(c)
}

func (c ArgrepPredicateRePattern) Re() ArgrepPredicateReRe {
	c.cs.s = append(c.cs.s, "RE")
	return (ArgrepPredicateReRe)(c)
}

func (c ArgrepPredicateRePattern) And() ArgrepOptionsAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (ArgrepOptionsAnd)(c)
}

func (c ArgrepPredicateRePattern) Or() ArgrepOptionsOr {
	c.cs.s = append(c.cs.s, "OR")
	return (ArgrepOptionsOr)(c)
}

func (c ArgrepPredicateRePattern) Limit(limit int64) ArgrepOptionsLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return (ArgrepOptionsLimit)(c)
}

func (c ArgrepPredicateRePattern) Withvalues() ArgrepOptionsWithvalues {
	c.cs.s = append(c.cs.s, "WITHVALUES")
	return (ArgrepOptionsWithvalues)(c)
}

func (c ArgrepPredicateRePattern) Nocase() ArgrepOptionsNocase {
	c.cs.s = append(c.cs.s, "NOCASE")
	return (ArgrepOptionsNocase)(c)
}

func (c ArgrepPredicateRePattern) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArgrepPredicateReRe Incomplete

func (c ArgrepPredicateReRe) Pattern(pattern string) ArgrepPredicateRePattern {
	c.cs.s = append(c.cs.s, pattern)
	return (ArgrepPredicateRePattern)(c)
}

type ArgrepStart Incomplete

func (c ArgrepStart) End(end string) ArgrepEnd {
	c.cs.s = append(c.cs.s, end)
	return (ArgrepEnd)(c)
}

type Arinfo Incomplete

func (b Builder) Arinfo() (c Arinfo) {
	c = Arinfo{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARINFO")
	return c
}

func (c Arinfo) Key(key string) ArinfoKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArinfoKey)(c)
}

type ArinfoFull Incomplete

func (c ArinfoFull) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArinfoKey Incomplete

func (c ArinfoKey) Full() ArinfoFull {
	c.cs.s = append(c.cs.s, "FULL")
	return (ArinfoFull)(c)
}

func (c ArinfoKey) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type Arinsert Incomplete

func (b Builder) Arinsert() (c Arinsert) {
	c = Arinsert{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "ARINSERT")
	return c
}

func (c Arinsert) Key(key string) ArinsertKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArinsertKey)(c)
}

type ArinsertKey Incomplete

func (c ArinsertKey) Value(value ...string) ArinsertValue {
	c.cs.s = append(c.cs.s, value...)
	return (ArinsertValue)(c)
}

type ArinsertValue Incomplete

func (c ArinsertValue) Value(value ...string) ArinsertValue {
	c.cs.s = append(c.cs.s, value...)
	return c
}

func (c ArinsertValue) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type Arlastitems Incomplete

func (b Builder) Arlastitems() (c Arlastitems) {
	c = Arlastitems{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARLASTITEMS")
	return c
}

func (c Arlastitems) Key(key string) ArlastitemsKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArlastitemsKey)(c)
}

type ArlastitemsCount Incomplete

func (c ArlastitemsCount) Rev() ArlastitemsRev {
	c.cs.s = append(c.cs.s, "REV")
	return (ArlastitemsRev)(c)
}

func (c ArlastitemsCount) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArlastitemsKey Incomplete

func (c ArlastitemsKey) Count(count int64) ArlastitemsCount {
	c.cs.s = append(c.cs.s, strconv.FormatInt(count, 10))
	return (ArlastitemsCount)(c)
}

type ArlastitemsRev Incomplete

func (c ArlastitemsRev) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type Arlen Incomplete

func (b Builder) Arlen() (c Arlen) {
	c = Arlen{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARLEN")
	return c
}

func (c Arlen) Key(key string) ArlenKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArlenKey)(c)
}

type ArlenKey Incomplete

func (c ArlenKey) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type Armget Incomplete

func (b Builder) Armget() (c Armget) {
	c = Armget{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARMGET")
	return c
}

func (c Armget) Key(key string) ArmgetKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArmgetKey)(c)
}

type ArmgetIndex Incomplete

func (c ArmgetIndex) Index(index ...int64) ArmgetIndex {
	for _, n := range index {
		c.cs.s = append(c.cs.s, strconv.FormatInt(n, 10))
	}
	return c
}

func (c ArmgetIndex) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArmgetKey Incomplete

func (c ArmgetKey) Index(index ...int64) ArmgetIndex {
	for _, n := range index {
		c.cs.s = append(c.cs.s, strconv.FormatInt(n, 10))
	}
	return (ArmgetIndex)(c)
}

type Armset Incomplete

func (b Builder) Armset() (c Armset) {
	c = Armset{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "ARMSET")
	return c
}

func (c Armset) Key(key string) ArmsetKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArmsetKey)(c)
}

type ArmsetDataIndex Incomplete

func (c ArmsetDataIndex) Value(value string) ArmsetDataValue {
	c.cs.s = append(c.cs.s, value)
	return (ArmsetDataValue)(c)
}

type ArmsetDataValue Incomplete

func (c ArmsetDataValue) Index(index int64) ArmsetDataIndex {
	c.cs.s = append(c.cs.s, strconv.FormatInt(index, 10))
	return (ArmsetDataIndex)(c)
}

func (c ArmsetDataValue) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArmsetKey Incomplete

func (c ArmsetKey) Index(index int64) ArmsetDataIndex {
	c.cs.s = append(c.cs.s, strconv.FormatInt(index, 10))
	return (ArmsetDataIndex)(c)
}

type Arnext Incomplete

func (b Builder) Arnext() (c Arnext) {
	c = Arnext{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARNEXT")
	return c
}

func (c Arnext) Key(key string) ArnextKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArnextKey)(c)
}

type ArnextKey Incomplete

func (c ArnextKey) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type Arop Incomplete

func (b Builder) Arop() (c Arop) {
	c = Arop{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "AROP")
	return c
}

func (c Arop) Key(key string) AropKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (AropKey)(c)
}

type AropEnd Incomplete

func (c AropEnd) Sum() AropOperationSum {
	c.cs.s = append(c.cs.s, "SUM")
	return (AropOperationSum)(c)
}

func (c AropEnd) Min() AropOperationMin {
	c.cs.s = append(c.cs.s, "MIN")
	return (AropOperationMin)(c)
}

func (c AropEnd) Max() AropOperationMax {
	c.cs.s = append(c.cs.s, "MAX")
	return (AropOperationMax)(c)
}

func (c AropEnd) And() AropOperationAnd {
	c.cs.s = append(c.cs.s, "AND")
	return (AropOperationAnd)(c)
}

func (c AropEnd) Or() AropOperationOr {
	c.cs.s = append(c.cs.s, "OR")
	return (AropOperationOr)(c)
}

func (c AropEnd) Xor() AropOperationXor {
	c.cs.s = append(c.cs.s, "XOR")
	return (AropOperationXor)(c)
}

func (c AropEnd) Match() AropOperationMatchMatch {
	c.cs.s = append(c.cs.s, "MATCH")
	return (AropOperationMatchMatch)(c)
}

func (c AropEnd) Used() AropOperationUsed {
	c.cs.s = append(c.cs.s, "USED")
	return (AropOperationUsed)(c)
}

type AropKey Incomplete

func (c AropKey) Start(start int64) AropStart {
	c.cs.s = append(c.cs.s, strconv.FormatInt(start, 10))
	return (AropStart)(c)
}

type AropOperationAnd Incomplete

func (c AropOperationAnd) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropOperationMatchMatch Incomplete

func (c AropOperationMatchMatch) Value(value string) AropOperationMatchValue {
	c.cs.s = append(c.cs.s, value)
	return (AropOperationMatchValue)(c)
}

type AropOperationMatchValue Incomplete

func (c AropOperationMatchValue) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropOperationMax Incomplete

func (c AropOperationMax) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropOperationMin Incomplete

func (c AropOperationMin) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropOperationOr Incomplete

func (c AropOperationOr) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropOperationSum Incomplete

func (c AropOperationSum) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropOperationUsed Incomplete

func (c AropOperationUsed) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropOperationXor Incomplete

func (c AropOperationXor) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type AropStart Incomplete

func (c AropStart) End(end int64) AropEnd {
	c.cs.s = append(c.cs.s, strconv.FormatInt(end, 10))
	return (AropEnd)(c)
}

type Arring Incomplete

func (b Builder) Arring() (c Arring) {
	c = Arring{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "ARRING")
	return c
}

func (c Arring) Key(key string) ArringKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArringKey)(c)
}

type ArringKey Incomplete

func (c ArringKey) Size(size int64) ArringSize {
	c.cs.s = append(c.cs.s, strconv.FormatInt(size, 10))
	return (ArringSize)(c)
}

type ArringSize Incomplete

func (c ArringSize) Value(value ...string) ArringValue {
	c.cs.s = append(c.cs.s, value...)
	return (ArringValue)(c)
}

type ArringValue Incomplete

func (c ArringValue) Value(value ...string) ArringValue {
	c.cs.s = append(c.cs.s, value...)
	return c
}

func (c ArringValue) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type Arscan Incomplete

func (b Builder) Arscan() (c Arscan) {
	c = Arscan{cs: get(), ks: b.ks, cf: int16(readonly)}
	c.cs.s = append(c.cs.s, "ARSCAN")
	return c
}

func (c Arscan) Key(key string) ArscanKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArscanKey)(c)
}

type ArscanEnd Incomplete

func (c ArscanEnd) Limit(limit int64) ArscanLimit {
	c.cs.s = append(c.cs.s, "LIMIT", strconv.FormatInt(limit, 10))
	return (ArscanLimit)(c)
}

func (c ArscanEnd) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArscanKey Incomplete

func (c ArscanKey) Start(start int64) ArscanStart {
	c.cs.s = append(c.cs.s, strconv.FormatInt(start, 10))
	return (ArscanStart)(c)
}

type ArscanLimit Incomplete

func (c ArscanLimit) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArscanStart Incomplete

func (c ArscanStart) End(end int64) ArscanEnd {
	c.cs.s = append(c.cs.s, strconv.FormatInt(end, 10))
	return (ArscanEnd)(c)
}

type Arseek Incomplete

func (b Builder) Arseek() (c Arseek) {
	c = Arseek{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "ARSEEK")
	return c
}

func (c Arseek) Key(key string) ArseekKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArseekKey)(c)
}

type ArseekIndex Incomplete

func (c ArseekIndex) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}

type ArseekKey Incomplete

func (c ArseekKey) Index(index int64) ArseekIndex {
	c.cs.s = append(c.cs.s, strconv.FormatInt(index, 10))
	return (ArseekIndex)(c)
}

type Arset Incomplete

func (b Builder) Arset() (c Arset) {
	c = Arset{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "ARSET")
	return c
}

func (c Arset) Key(key string) ArsetKey {
	if c.ks&NoSlot == NoSlot {
		c.ks = NoSlot | slot(key)
	} else {
		c.ks = check(c.ks, slot(key))
	}
	c.cs.s = append(c.cs.s, key)
	return (ArsetKey)(c)
}

type ArsetIndex Incomplete

func (c ArsetIndex) Value(value ...string) ArsetValue {
	c.cs.s = append(c.cs.s, value...)
	return (ArsetValue)(c)
}

type ArsetKey Incomplete

func (c ArsetKey) Index(index int64) ArsetIndex {
	c.cs.s = append(c.cs.s, strconv.FormatInt(index, 10))
	return (ArsetIndex)(c)
}

type ArsetValue Incomplete

func (c ArsetValue) Value(value ...string) ArsetValue {
	c.cs.s = append(c.cs.s, value...)
	return c
}

func (c ArsetValue) Build() Completed {
	c.cs.Build()
	return Completed{cs: c.cs, cf: uint16(c.cf), ks: c.ks}
}
