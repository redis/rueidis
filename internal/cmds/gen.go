// Code generated DO NOT EDIT

package cmds

import "strconv"

type AclCat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclCat) Categoryname(Categoryname string) AclCatCategoryname {
	return AclCatCategoryname{cf: c.cf, cs: append(c.cs, Categoryname)}
}

func (c AclCat) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclCat() (c AclCat) {
	c.cs = append(b.get(), "ACL", "CAT")
	return
}

type AclCatCategoryname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclCatCategoryname) Build() Completed {
	return Completed(c)
}

type AclDeluser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclDeluser) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (b *Builder) AclDeluser() (c AclDeluser) {
	c.cs = append(b.get(), "ACL", "DELUSER")
	return
}

type AclDeluserUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclDeluserUsername) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (c AclDeluserUsername) Build() Completed {
	return Completed(c)
}

type AclGenpass struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclGenpass) Bits(Bits int64) AclGenpassBits {
	return AclGenpassBits{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bits, 10))}
}

func (c AclGenpass) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclGenpass() (c AclGenpass) {
	c.cs = append(b.get(), "ACL", "GENPASS")
	return
}

type AclGenpassBits struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclGenpassBits) Build() Completed {
	return Completed(c)
}

type AclGetuser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclGetuser) Username(Username string) AclGetuserUsername {
	return AclGetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (b *Builder) AclGetuser() (c AclGetuser) {
	c.cs = append(b.get(), "ACL", "GETUSER")
	return
}

type AclGetuserUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclGetuserUsername) Build() Completed {
	return Completed(c)
}

type AclHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclHelp) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclHelp() (c AclHelp) {
	c.cs = append(b.get(), "ACL", "HELP")
	return
}

type AclList struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclList) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclList() (c AclList) {
	c.cs = append(b.get(), "ACL", "LIST")
	return
}

type AclLoad struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclLoad) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclLoad() (c AclLoad) {
	c.cs = append(b.get(), "ACL", "LOAD")
	return
}

type AclLog struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclLog) CountOrReset(CountOrReset string) AclLogCountOrReset {
	return AclLogCountOrReset{cf: c.cf, cs: append(c.cs, CountOrReset)}
}

func (c AclLog) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclLog() (c AclLog) {
	c.cs = append(b.get(), "ACL", "LOG")
	return
}

type AclLogCountOrReset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclLogCountOrReset) Build() Completed {
	return Completed(c)
}

type AclSave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclSave) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclSave() (c AclSave) {
	c.cs = append(b.get(), "ACL", "SAVE")
	return
}

type AclSetuser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclSetuser) Username(Username string) AclSetuserUsername {
	return AclSetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (b *Builder) AclSetuser() (c AclSetuser) {
	c.cs = append(b.get(), "ACL", "SETUSER")
	return
}

type AclSetuserRule struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclSetuserRule) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c AclSetuserRule) Build() Completed {
	return Completed(c)
}

type AclSetuserUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclSetuserUsername) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c AclSetuserUsername) Build() Completed {
	return Completed(c)
}

type AclUsers struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclUsers) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclUsers() (c AclUsers) {
	c.cs = append(b.get(), "ACL", "USERS")
	return
}

type AclWhoami struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AclWhoami) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclWhoami() (c AclWhoami) {
	c.cs = append(b.get(), "ACL", "WHOAMI")
	return
}

type Append struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Append) Key(Key string) AppendKey {
	return AppendKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Append() (c Append) {
	c.cs = append(b.get(), "APPEND")
	return
}

type AppendKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AppendKey) Value(Value string) AppendValue {
	return AppendValue{cf: c.cf, cs: append(c.cs, Value)}
}

type AppendValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AppendValue) Build() Completed {
	return Completed(c)
}

type Asking struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Asking) Build() Completed {
	return Completed(c)
}

func (b *Builder) Asking() (c Asking) {
	c.cs = append(b.get(), "ASKING")
	return
}

type Auth struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Auth) Username(Username string) AuthUsername {
	return AuthUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (c Auth) Password(Password string) AuthPassword {
	return AuthPassword{cf: c.cf, cs: append(c.cs, Password)}
}

func (b *Builder) Auth() (c Auth) {
	c.cs = append(b.get(), "AUTH")
	return
}

type AuthPassword struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AuthPassword) Build() Completed {
	return Completed(c)
}

type AuthUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c AuthUsername) Password(Password string) AuthPassword {
	return AuthPassword{cf: c.cf, cs: append(c.cs, Password)}
}

type Bgrewriteaof struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bgrewriteaof) Build() Completed {
	return Completed(c)
}

func (b *Builder) Bgrewriteaof() (c Bgrewriteaof) {
	c.cs = append(b.get(), "BGREWRITEAOF")
	return
}

type Bgsave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bgsave) Schedule() BgsaveScheduleSchedule {
	return BgsaveScheduleSchedule{cf: c.cf, cs: append(c.cs, "SCHEDULE")}
}

func (c Bgsave) Build() Completed {
	return Completed(c)
}

func (b *Builder) Bgsave() (c Bgsave) {
	c.cs = append(b.get(), "BGSAVE")
	return
}

type BgsaveScheduleSchedule struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BgsaveScheduleSchedule) Build() Completed {
	return Completed(c)
}

type Bitcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bitcount) Key(Key string) BitcountKey {
	return BitcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Bitcount() (c Bitcount) {
	c.cs = append(b.get(), "BITCOUNT")
	return
}

type BitcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitcountKey) StartEnd(Start int64, End int64) BitcountStartEnd {
	return BitcountStartEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10), strconv.FormatInt(End, 10))}
}

func (c BitcountKey) Build() Completed {
	return Completed(c)
}

func (c BitcountKey) Cache() Cacheable {
	return Cacheable(c)
}

type BitcountStartEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitcountStartEnd) Build() Completed {
	return Completed(c)
}

func (c BitcountStartEnd) Cache() Cacheable {
	return Cacheable(c)
}

type Bitfield struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bitfield) Key(Key string) BitfieldKey {
	return BitfieldKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Bitfield() (c Bitfield) {
	c.cs = append(b.get(), "BITFIELD")
	return
}

type BitfieldFail struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldFail) Build() Completed {
	return Completed(c)
}

type BitfieldGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldGet) Set(Type string, Offset int64, Value int64) BitfieldSet {
	return BitfieldSet{cf: c.cf, cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10))}
}

func (c BitfieldGet) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cf: c.cf, cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c BitfieldGet) Wrap() BitfieldWrap {
	return BitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c BitfieldGet) Sat() BitfieldSat {
	return BitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c BitfieldGet) Fail() BitfieldFail {
	return BitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c BitfieldGet) Build() Completed {
	return Completed(c)
}

type BitfieldIncrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldIncrby) Wrap() BitfieldWrap {
	return BitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c BitfieldIncrby) Sat() BitfieldSat {
	return BitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c BitfieldIncrby) Fail() BitfieldFail {
	return BitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c BitfieldIncrby) Build() Completed {
	return Completed(c)
}

type BitfieldKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldKey) Get(Type string, Offset int64) BitfieldGet {
	return BitfieldGet{cf: c.cf, cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

func (c BitfieldKey) Set(Type string, Offset int64, Value int64) BitfieldSet {
	return BitfieldSet{cf: c.cf, cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10))}
}

func (c BitfieldKey) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cf: c.cf, cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c BitfieldKey) Wrap() BitfieldWrap {
	return BitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c BitfieldKey) Sat() BitfieldSat {
	return BitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c BitfieldKey) Fail() BitfieldFail {
	return BitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c BitfieldKey) Build() Completed {
	return Completed(c)
}

type BitfieldRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldRo) Key(Key string) BitfieldRoKey {
	return BitfieldRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) BitfieldRo() (c BitfieldRo) {
	c.cs = append(b.get(), "BITFIELD_RO")
	return
}

type BitfieldRoGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldRoGet) Build() Completed {
	return Completed(c)
}

func (c BitfieldRoGet) Cache() Cacheable {
	return Cacheable(c)
}

type BitfieldRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldRoKey) Get(Type string, Offset int64) BitfieldRoGet {
	return BitfieldRoGet{cf: c.cf, cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

func (c BitfieldRoKey) Cache() Cacheable {
	return Cacheable(c)
}

type BitfieldSat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldSat) Build() Completed {
	return Completed(c)
}

type BitfieldSet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldSet) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cf: c.cf, cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c BitfieldSet) Wrap() BitfieldWrap {
	return BitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c BitfieldSet) Sat() BitfieldSat {
	return BitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c BitfieldSet) Fail() BitfieldFail {
	return BitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c BitfieldSet) Build() Completed {
	return Completed(c)
}

type BitfieldWrap struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitfieldWrap) Build() Completed {
	return Completed(c)
}

type Bitop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bitop) Operation(Operation string) BitopOperation {
	return BitopOperation{cf: c.cf, cs: append(c.cs, Operation)}
}

func (b *Builder) Bitop() (c Bitop) {
	c.cs = append(b.get(), "BITOP")
	return
}

type BitopDestkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitopDestkey) Key(Key ...string) BitopKey {
	return BitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BitopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitopKey) Key(Key ...string) BitopKey {
	return BitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c BitopKey) Build() Completed {
	return Completed(c)
}

type BitopOperation struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitopOperation) Destkey(Destkey string) BitopDestkey {
	return BitopDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

type Bitpos struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bitpos) Key(Key string) BitposKey {
	return BitposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Bitpos() (c Bitpos) {
	c.cs = append(b.get(), "BITPOS")
	return
}

type BitposBit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitposBit) Start(Start int64) BitposIndexStart {
	return BitposIndexStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c BitposBit) Cache() Cacheable {
	return Cacheable(c)
}

type BitposIndexEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitposIndexEnd) Build() Completed {
	return Completed(c)
}

func (c BitposIndexEnd) Cache() Cacheable {
	return Cacheable(c)
}

type BitposIndexStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitposIndexStart) End(End int64) BitposIndexEnd {
	return BitposIndexEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c BitposIndexStart) Build() Completed {
	return Completed(c)
}

func (c BitposIndexStart) Cache() Cacheable {
	return Cacheable(c)
}

type BitposKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BitposKey) Bit(Bit int64) BitposBit {
	return BitposBit{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bit, 10))}
}

func (c BitposKey) Cache() Cacheable {
	return Cacheable(c)
}

type Blmove struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Blmove) Source(Source string) BlmoveSource {
	return BlmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Blmove() (c Blmove) {
	c.cs = append(b.get(), "BLMOVE")
	c.cf = blockTag
	return
}

type BlmoveDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmoveDestination) Left() BlmoveWherefromLeft {
	return BlmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveDestination) Right() BlmoveWherefromRight {
	return BlmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmoveSource) Destination(Destination string) BlmoveDestination {
	return BlmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type BlmoveTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmoveTimeout) Build() Completed {
	return Completed(c)
}

type BlmoveWherefromLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmoveWherefromLeft) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromLeft) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveWherefromRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmoveWherefromRight) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromRight) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveWheretoLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmoveWheretoLeft) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BlmoveWheretoRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmoveWheretoRight) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type Blmpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Blmpop) Timeout(Timeout float64) BlmpopTimeout {
	return BlmpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (b *Builder) Blmpop() (c Blmpop) {
	c.cs = append(b.get(), "BLMPOP")
	c.cf = blockTag
	return
}

type BlmpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmpopCount) Build() Completed {
	return Completed(c)
}

type BlmpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmpopKey) Left() BlmpopWhereLeft {
	return BlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmpopKey) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c BlmpopKey) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BlmpopNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmpopNumkeys) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c BlmpopNumkeys) Left() BlmpopWhereLeft {
	return BlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmpopNumkeys) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmpopTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmpopTimeout) Numkeys(Numkeys int64) BlmpopNumkeys {
	return BlmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type BlmpopWhereLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmpopWhereLeft) Count(Count int64) BlmpopCount {
	return BlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type BlmpopWhereRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlmpopWhereRight) Count(Count int64) BlmpopCount {
	return BlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Blpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Blpop) Key(Key ...string) BlpopKey {
	return BlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Blpop() (c Blpop) {
	c.cs = append(b.get(), "BLPOP")
	c.cf = blockTag
	return
}

type BlpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlpopKey) Timeout(Timeout float64) BlpopTimeout {
	return BlpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BlpopKey) Key(Key ...string) BlpopKey {
	return BlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BlpopTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BlpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Brpop) Key(Key ...string) BrpopKey {
	return BrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Brpop() (c Brpop) {
	c.cs = append(b.get(), "BRPOP")
	c.cf = blockTag
	return
}

type BrpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BrpopKey) Timeout(Timeout float64) BrpopTimeout {
	return BrpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BrpopKey) Key(Key ...string) BrpopKey {
	return BrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BrpopTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BrpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpoplpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Brpoplpush) Source(Source string) BrpoplpushSource {
	return BrpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Brpoplpush() (c Brpoplpush) {
	c.cs = append(b.get(), "BRPOPLPUSH")
	c.cf = blockTag
	return
}

type BrpoplpushDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BrpoplpushDestination) Timeout(Timeout float64) BrpoplpushTimeout {
	return BrpoplpushTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BrpoplpushSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BrpoplpushSource) Destination(Destination string) BrpoplpushDestination {
	return BrpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type BrpoplpushTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BrpoplpushTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bzpopmax) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmax() (c Bzpopmax) {
	c.cs = append(b.get(), "BZPOPMAX")
	c.cf = blockTag
	return
}

type BzpopmaxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BzpopmaxKey) Timeout(Timeout float64) BzpopmaxTimeout {
	return BzpopmaxTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopmaxKey) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BzpopmaxTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BzpopmaxTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Bzpopmin) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmin() (c Bzpopmin) {
	c.cs = append(b.get(), "BZPOPMIN")
	c.cf = blockTag
	return
}

type BzpopminKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BzpopminKey) Timeout(Timeout float64) BzpopminTimeout {
	return BzpopminTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopminKey) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BzpopminTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c BzpopminTimeout) Build() Completed {
	return Completed(c)
}

type ClientCaching struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientCaching) Yes() ClientCachingModeYes {
	return ClientCachingModeYes{cf: c.cf, cs: append(c.cs, "YES")}
}

func (c ClientCaching) No() ClientCachingModeNo {
	return ClientCachingModeNo{cf: c.cf, cs: append(c.cs, "NO")}
}

func (b *Builder) ClientCaching() (c ClientCaching) {
	c.cs = append(b.get(), "CLIENT", "CACHING")
	return
}

type ClientCachingModeNo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientCachingModeNo) Build() Completed {
	return Completed(c)
}

type ClientCachingModeYes struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientCachingModeYes) Build() Completed {
	return Completed(c)
}

type ClientGetname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientGetname) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientGetname() (c ClientGetname) {
	c.cs = append(b.get(), "CLIENT", "GETNAME")
	return
}

type ClientGetredir struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientGetredir) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientGetredir() (c ClientGetredir) {
	c.cs = append(b.get(), "CLIENT", "GETREDIR")
	return
}

type ClientId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientId) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientId() (c ClientId) {
	c.cs = append(b.get(), "CLIENT", "ID")
	return
}

type ClientInfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientInfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientInfo() (c ClientInfo) {
	c.cs = append(b.get(), "CLIENT", "INFO")
	return
}

type ClientKill struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKill) IpPort(IpPort string) ClientKillIpPort {
	return ClientKillIpPort{cf: c.cf, cs: append(c.cs, IpPort)}
}

func (c ClientKill) Id(ClientId int64) ClientKillId {
	return ClientKillId{cf: c.cf, cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10))}
}

func (c ClientKill) Normal() ClientKillNormal {
	return ClientKillNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c ClientKill) Master() ClientKillMaster {
	return ClientKillMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c ClientKill) Slave() ClientKillSlave {
	return ClientKillSlave{cf: c.cf, cs: append(c.cs, "slave")}
}

func (c ClientKill) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c ClientKill) User(Username string) ClientKillUser {
	return ClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c ClientKill) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKill) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKill) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKill) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientKill() (c ClientKill) {
	c.cs = append(b.get(), "CLIENT", "KILL")
	return
}

type ClientKillAddr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillAddr) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillAddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillAddr) Build() Completed {
	return Completed(c)
}

type ClientKillId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillId) Normal() ClientKillNormal {
	return ClientKillNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c ClientKillId) Master() ClientKillMaster {
	return ClientKillMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c ClientKillId) Slave() ClientKillSlave {
	return ClientKillSlave{cf: c.cf, cs: append(c.cs, "slave")}
}

func (c ClientKillId) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c ClientKillId) User(Username string) ClientKillUser {
	return ClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c ClientKillId) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillId) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillId) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillId) Build() Completed {
	return Completed(c)
}

type ClientKillIpPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillIpPort) Id(ClientId int64) ClientKillId {
	return ClientKillId{cf: c.cf, cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10))}
}

func (c ClientKillIpPort) Normal() ClientKillNormal {
	return ClientKillNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c ClientKillIpPort) Master() ClientKillMaster {
	return ClientKillMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c ClientKillIpPort) Slave() ClientKillSlave {
	return ClientKillSlave{cf: c.cf, cs: append(c.cs, "slave")}
}

func (c ClientKillIpPort) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c ClientKillIpPort) User(Username string) ClientKillUser {
	return ClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c ClientKillIpPort) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillIpPort) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillIpPort) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillIpPort) Build() Completed {
	return Completed(c)
}

type ClientKillLaddr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillLaddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillLaddr) Build() Completed {
	return Completed(c)
}

type ClientKillMaster struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillMaster) User(Username string) ClientKillUser {
	return ClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c ClientKillMaster) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillMaster) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillMaster) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillMaster) Build() Completed {
	return Completed(c)
}

type ClientKillNormal struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillNormal) User(Username string) ClientKillUser {
	return ClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c ClientKillNormal) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillNormal) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillNormal) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillNormal) Build() Completed {
	return Completed(c)
}

type ClientKillPubsub struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillPubsub) User(Username string) ClientKillUser {
	return ClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c ClientKillPubsub) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillPubsub) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillPubsub) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillPubsub) Build() Completed {
	return Completed(c)
}

type ClientKillSkipme struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillSkipme) Build() Completed {
	return Completed(c)
}

type ClientKillSlave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillSlave) User(Username string) ClientKillUser {
	return ClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c ClientKillSlave) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillSlave) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillSlave) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillSlave) Build() Completed {
	return Completed(c)
}

type ClientKillUser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientKillUser) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillUser) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillUser) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillUser) Build() Completed {
	return Completed(c)
}

type ClientList struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientList) Normal() ClientListNormal {
	return ClientListNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c ClientList) Master() ClientListMaster {
	return ClientListMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c ClientList) Replica() ClientListReplica {
	return ClientListReplica{cf: c.cf, cs: append(c.cs, "replica")}
}

func (c ClientList) Pubsub() ClientListPubsub {
	return ClientListPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c ClientList) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (b *Builder) ClientList() (c ClientList) {
	c.cs = append(b.get(), "CLIENT", "LIST")
	return
}

type ClientListIdClientId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientListIdClientId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cf: c.cf, cs: c.cs}
}

func (c ClientListIdClientId) Build() Completed {
	return Completed(c)
}

type ClientListIdId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientListIdId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cf: c.cf, cs: c.cs}
}

type ClientListMaster struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientListMaster) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientListNormal struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientListNormal) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientListPubsub struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientListPubsub) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientListReplica struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientListReplica) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientNoEvict struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientNoEvict) On() ClientNoEvictEnabledOn {
	return ClientNoEvictEnabledOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c ClientNoEvict) Off() ClientNoEvictEnabledOff {
	return ClientNoEvictEnabledOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (b *Builder) ClientNoEvict() (c ClientNoEvict) {
	c.cs = append(b.get(), "CLIENT", "NO-EVICT")
	return
}

type ClientNoEvictEnabledOff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientNoEvictEnabledOff) Build() Completed {
	return Completed(c)
}

type ClientNoEvictEnabledOn struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientNoEvictEnabledOn) Build() Completed {
	return Completed(c)
}

type ClientPause struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientPause) Timeout(Timeout int64) ClientPauseTimeout {
	return ClientPauseTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

func (b *Builder) ClientPause() (c ClientPause) {
	c.cs = append(b.get(), "CLIENT", "PAUSE")
	c.cf = blockTag
	return
}

type ClientPauseModeAll struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientPauseModeAll) Build() Completed {
	return Completed(c)
}

type ClientPauseModeWrite struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientPauseModeWrite) Build() Completed {
	return Completed(c)
}

type ClientPauseTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientPauseTimeout) Write() ClientPauseModeWrite {
	return ClientPauseModeWrite{cf: c.cf, cs: append(c.cs, "WRITE")}
}

func (c ClientPauseTimeout) All() ClientPauseModeAll {
	return ClientPauseModeAll{cf: c.cf, cs: append(c.cs, "ALL")}
}

func (c ClientPauseTimeout) Build() Completed {
	return Completed(c)
}

type ClientReply struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientReply) On() ClientReplyReplyModeOn {
	return ClientReplyReplyModeOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c ClientReply) Off() ClientReplyReplyModeOff {
	return ClientReplyReplyModeOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (c ClientReply) Skip() ClientReplyReplyModeSkip {
	return ClientReplyReplyModeSkip{cf: c.cf, cs: append(c.cs, "SKIP")}
}

func (b *Builder) ClientReply() (c ClientReply) {
	c.cs = append(b.get(), "CLIENT", "REPLY")
	return
}

type ClientReplyReplyModeOff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientReplyReplyModeOff) Build() Completed {
	return Completed(c)
}

type ClientReplyReplyModeOn struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientReplyReplyModeOn) Build() Completed {
	return Completed(c)
}

type ClientReplyReplyModeSkip struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientReplyReplyModeSkip) Build() Completed {
	return Completed(c)
}

type ClientSetname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientSetname) ConnectionName(ConnectionName string) ClientSetnameConnectionName {
	return ClientSetnameConnectionName{cf: c.cf, cs: append(c.cs, ConnectionName)}
}

func (b *Builder) ClientSetname() (c ClientSetname) {
	c.cs = append(b.get(), "CLIENT", "SETNAME")
	return
}

type ClientSetnameConnectionName struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientSetnameConnectionName) Build() Completed {
	return Completed(c)
}

type ClientTracking struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTracking) On() ClientTrackingStatusOn {
	return ClientTrackingStatusOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c ClientTracking) Off() ClientTrackingStatusOff {
	return ClientTrackingStatusOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (b *Builder) ClientTracking() (c ClientTracking) {
	c.cs = append(b.get(), "CLIENT", "TRACKING")
	return
}

type ClientTrackingBcastBcast struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingBcastBcast) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingBcastBcast) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingBcastBcast) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingBcastBcast) Build() Completed {
	return Completed(c)
}

type ClientTrackingNoloopNoloop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingNoloopNoloop) Build() Completed {
	return Completed(c)
}

type ClientTrackingOptinOptin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingOptinOptin) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingOptinOptin) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingOptinOptin) Build() Completed {
	return Completed(c)
}

type ClientTrackingOptoutOptout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingOptoutOptout) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingOptoutOptout) Build() Completed {
	return Completed(c)
}

type ClientTrackingPrefix struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingPrefix) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingPrefix) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingPrefix) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingPrefix) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingPrefix) Prefix(Prefix ...string) ClientTrackingPrefix {
	return ClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingPrefix) Build() Completed {
	return Completed(c)
}

type ClientTrackingRedirect struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingRedirect) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingRedirect) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingRedirect) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingRedirect) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingRedirect) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingRedirect) Build() Completed {
	return Completed(c)
}

type ClientTrackingStatusOff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingStatusOff) Redirect(ClientId int64) ClientTrackingRedirect {
	return ClientTrackingRedirect{cf: c.cf, cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10))}
}

func (c ClientTrackingStatusOff) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingStatusOff) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingStatusOff) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingStatusOff) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingStatusOff) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingStatusOff) Build() Completed {
	return Completed(c)
}

type ClientTrackingStatusOn struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackingStatusOn) Redirect(ClientId int64) ClientTrackingRedirect {
	return ClientTrackingRedirect{cf: c.cf, cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10))}
}

func (c ClientTrackingStatusOn) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingStatusOn) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingStatusOn) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingStatusOn) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingStatusOn) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingStatusOn) Build() Completed {
	return Completed(c)
}

type ClientTrackinginfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientTrackinginfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientTrackinginfo() (c ClientTrackinginfo) {
	c.cs = append(b.get(), "CLIENT", "TRACKINGINFO")
	return
}

type ClientUnblock struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientUnblock) ClientId(ClientId int64) ClientUnblockClientId {
	return ClientUnblockClientId{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ClientId, 10))}
}

func (b *Builder) ClientUnblock() (c ClientUnblock) {
	c.cs = append(b.get(), "CLIENT", "UNBLOCK")
	return
}

type ClientUnblockClientId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientUnblockClientId) Timeout() ClientUnblockUnblockTypeTimeout {
	return ClientUnblockUnblockTypeTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT")}
}

func (c ClientUnblockClientId) Error() ClientUnblockUnblockTypeError {
	return ClientUnblockUnblockTypeError{cf: c.cf, cs: append(c.cs, "ERROR")}
}

func (c ClientUnblockClientId) Build() Completed {
	return Completed(c)
}

type ClientUnblockUnblockTypeError struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientUnblockUnblockTypeError) Build() Completed {
	return Completed(c)
}

type ClientUnblockUnblockTypeTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientUnblockUnblockTypeTimeout) Build() Completed {
	return Completed(c)
}

type ClientUnpause struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClientUnpause) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientUnpause() (c ClientUnpause) {
	c.cs = append(b.get(), "CLIENT", "UNPAUSE")
	return
}

type ClusterAddslots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterAddslots) Slot(Slot ...int64) ClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterAddslotsSlot{cf: c.cf, cs: c.cs}
}

func (b *Builder) ClusterAddslots() (c ClusterAddslots) {
	c.cs = append(b.get(), "CLUSTER", "ADDSLOTS")
	return
}

type ClusterAddslotsSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterAddslotsSlot) Slot(Slot ...int64) ClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterAddslotsSlot{cf: c.cf, cs: c.cs}
}

func (c ClusterAddslotsSlot) Build() Completed {
	return Completed(c)
}

type ClusterBumpepoch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterBumpepoch) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterBumpepoch() (c ClusterBumpepoch) {
	c.cs = append(b.get(), "CLUSTER", "BUMPEPOCH")
	return
}

type ClusterCountFailureReports struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterCountFailureReports) NodeId(NodeId string) ClusterCountFailureReportsNodeId {
	return ClusterCountFailureReportsNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterCountFailureReports() (c ClusterCountFailureReports) {
	c.cs = append(b.get(), "CLUSTER", "COUNT-FAILURE-REPORTS")
	return
}

type ClusterCountFailureReportsNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterCountFailureReportsNodeId) Build() Completed {
	return Completed(c)
}

type ClusterCountkeysinslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterCountkeysinslot) Slot(Slot int64) ClusterCountkeysinslotSlot {
	return ClusterCountkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *Builder) ClusterCountkeysinslot() (c ClusterCountkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "COUNTKEYSINSLOT")
	return
}

type ClusterCountkeysinslotSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterCountkeysinslotSlot) Build() Completed {
	return Completed(c)
}

type ClusterDelslots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterDelslots) Slot(Slot ...int64) ClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterDelslotsSlot{cf: c.cf, cs: c.cs}
}

func (b *Builder) ClusterDelslots() (c ClusterDelslots) {
	c.cs = append(b.get(), "CLUSTER", "DELSLOTS")
	return
}

type ClusterDelslotsSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterDelslotsSlot) Slot(Slot ...int64) ClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterDelslotsSlot{cf: c.cf, cs: c.cs}
}

func (c ClusterDelslotsSlot) Build() Completed {
	return Completed(c)
}

type ClusterFailover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterFailover) Force() ClusterFailoverOptionsForce {
	return ClusterFailoverOptionsForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c ClusterFailover) Takeover() ClusterFailoverOptionsTakeover {
	return ClusterFailoverOptionsTakeover{cf: c.cf, cs: append(c.cs, "TAKEOVER")}
}

func (c ClusterFailover) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterFailover() (c ClusterFailover) {
	c.cs = append(b.get(), "CLUSTER", "FAILOVER")
	return
}

type ClusterFailoverOptionsForce struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterFailoverOptionsForce) Build() Completed {
	return Completed(c)
}

type ClusterFailoverOptionsTakeover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterFailoverOptionsTakeover) Build() Completed {
	return Completed(c)
}

type ClusterFlushslots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterFlushslots) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterFlushslots() (c ClusterFlushslots) {
	c.cs = append(b.get(), "CLUSTER", "FLUSHSLOTS")
	return
}

type ClusterForget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterForget) NodeId(NodeId string) ClusterForgetNodeId {
	return ClusterForgetNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterForget() (c ClusterForget) {
	c.cs = append(b.get(), "CLUSTER", "FORGET")
	return
}

type ClusterForgetNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterForgetNodeId) Build() Completed {
	return Completed(c)
}

type ClusterGetkeysinslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterGetkeysinslot) Slot(Slot int64) ClusterGetkeysinslotSlot {
	return ClusterGetkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *Builder) ClusterGetkeysinslot() (c ClusterGetkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "GETKEYSINSLOT")
	return
}

type ClusterGetkeysinslotCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterGetkeysinslotCount) Build() Completed {
	return Completed(c)
}

type ClusterGetkeysinslotSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterGetkeysinslotSlot) Count(Count int64) ClusterGetkeysinslotCount {
	return ClusterGetkeysinslotCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type ClusterInfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterInfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterInfo() (c ClusterInfo) {
	c.cs = append(b.get(), "CLUSTER", "INFO")
	return
}

type ClusterKeyslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterKeyslot) Key(Key string) ClusterKeyslotKey {
	return ClusterKeyslotKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) ClusterKeyslot() (c ClusterKeyslot) {
	c.cs = append(b.get(), "CLUSTER", "KEYSLOT")
	return
}

type ClusterKeyslotKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterKeyslotKey) Build() Completed {
	return Completed(c)
}

type ClusterMeet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterMeet) Ip(Ip string) ClusterMeetIp {
	return ClusterMeetIp{cf: c.cf, cs: append(c.cs, Ip)}
}

func (b *Builder) ClusterMeet() (c ClusterMeet) {
	c.cs = append(b.get(), "CLUSTER", "MEET")
	return
}

type ClusterMeetIp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterMeetIp) Port(Port int64) ClusterMeetPort {
	return ClusterMeetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type ClusterMeetPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterMeetPort) Build() Completed {
	return Completed(c)
}

type ClusterMyid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterMyid) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterMyid() (c ClusterMyid) {
	c.cs = append(b.get(), "CLUSTER", "MYID")
	return
}

type ClusterNodes struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterNodes) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterNodes() (c ClusterNodes) {
	c.cs = append(b.get(), "CLUSTER", "NODES")
	return
}

type ClusterReplicas struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterReplicas) NodeId(NodeId string) ClusterReplicasNodeId {
	return ClusterReplicasNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterReplicas() (c ClusterReplicas) {
	c.cs = append(b.get(), "CLUSTER", "REPLICAS")
	return
}

type ClusterReplicasNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterReplicasNodeId) Build() Completed {
	return Completed(c)
}

type ClusterReplicate struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterReplicate) NodeId(NodeId string) ClusterReplicateNodeId {
	return ClusterReplicateNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterReplicate() (c ClusterReplicate) {
	c.cs = append(b.get(), "CLUSTER", "REPLICATE")
	return
}

type ClusterReplicateNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterReplicateNodeId) Build() Completed {
	return Completed(c)
}

type ClusterReset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterReset) Hard() ClusterResetResetTypeHard {
	return ClusterResetResetTypeHard{cf: c.cf, cs: append(c.cs, "HARD")}
}

func (c ClusterReset) Soft() ClusterResetResetTypeSoft {
	return ClusterResetResetTypeSoft{cf: c.cf, cs: append(c.cs, "SOFT")}
}

func (c ClusterReset) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterReset() (c ClusterReset) {
	c.cs = append(b.get(), "CLUSTER", "RESET")
	return
}

type ClusterResetResetTypeHard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterResetResetTypeHard) Build() Completed {
	return Completed(c)
}

type ClusterResetResetTypeSoft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterResetResetTypeSoft) Build() Completed {
	return Completed(c)
}

type ClusterSaveconfig struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSaveconfig) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterSaveconfig() (c ClusterSaveconfig) {
	c.cs = append(b.get(), "CLUSTER", "SAVECONFIG")
	return
}

type ClusterSetConfigEpoch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetConfigEpoch) ConfigEpoch(ConfigEpoch int64) ClusterSetConfigEpochConfigEpoch {
	return ClusterSetConfigEpochConfigEpoch{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ConfigEpoch, 10))}
}

func (b *Builder) ClusterSetConfigEpoch() (c ClusterSetConfigEpoch) {
	c.cs = append(b.get(), "CLUSTER", "SET-CONFIG-EPOCH")
	return
}

type ClusterSetConfigEpochConfigEpoch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetConfigEpochConfigEpoch) Build() Completed {
	return Completed(c)
}

type ClusterSetslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetslot) Slot(Slot int64) ClusterSetslotSlot {
	return ClusterSetslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *Builder) ClusterSetslot() (c ClusterSetslot) {
	c.cs = append(b.get(), "CLUSTER", "SETSLOT")
	return
}

type ClusterSetslotNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetslotNodeId) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetslotSlot) Importing() ClusterSetslotSubcommandImporting {
	return ClusterSetslotSubcommandImporting{cf: c.cf, cs: append(c.cs, "IMPORTING")}
}

func (c ClusterSetslotSlot) Migrating() ClusterSetslotSubcommandMigrating {
	return ClusterSetslotSubcommandMigrating{cf: c.cf, cs: append(c.cs, "MIGRATING")}
}

func (c ClusterSetslotSlot) Stable() ClusterSetslotSubcommandStable {
	return ClusterSetslotSubcommandStable{cf: c.cf, cs: append(c.cs, "STABLE")}
}

func (c ClusterSetslotSlot) Node() ClusterSetslotSubcommandNode {
	return ClusterSetslotSubcommandNode{cf: c.cf, cs: append(c.cs, "NODE")}
}

type ClusterSetslotSubcommandImporting struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetslotSubcommandImporting) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandImporting) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandMigrating struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetslotSubcommandMigrating) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandMigrating) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandNode struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetslotSubcommandNode) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandNode) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandStable struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSetslotSubcommandStable) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandStable) Build() Completed {
	return Completed(c)
}

type ClusterSlaves struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSlaves) NodeId(NodeId string) ClusterSlavesNodeId {
	return ClusterSlavesNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterSlaves() (c ClusterSlaves) {
	c.cs = append(b.get(), "CLUSTER", "SLAVES")
	return
}

type ClusterSlavesNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSlavesNodeId) Build() Completed {
	return Completed(c)
}

type ClusterSlots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ClusterSlots) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterSlots() (c ClusterSlots) {
	c.cs = append(b.get(), "CLUSTER", "SLOTS")
	return
}

type Command struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Command) Build() Completed {
	return Completed(c)
}

func (b *Builder) Command() (c Command) {
	c.cs = append(b.get(), "COMMAND")
	return
}

type CommandCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CommandCount) Build() Completed {
	return Completed(c)
}

func (b *Builder) CommandCount() (c CommandCount) {
	c.cs = append(b.get(), "COMMAND", "COUNT")
	return
}

type CommandGetkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CommandGetkeys) Build() Completed {
	return Completed(c)
}

func (b *Builder) CommandGetkeys() (c CommandGetkeys) {
	c.cs = append(b.get(), "COMMAND", "GETKEYS")
	return
}

type CommandInfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CommandInfo) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (b *Builder) CommandInfo() (c CommandInfo) {
	c.cs = append(b.get(), "COMMAND", "INFO")
	return
}

type CommandInfoCommandName struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CommandInfoCommandName) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (c CommandInfoCommandName) Build() Completed {
	return Completed(c)
}

type ConfigGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ConfigGet) Parameter(Parameter string) ConfigGetParameter {
	return ConfigGetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
}

func (b *Builder) ConfigGet() (c ConfigGet) {
	c.cs = append(b.get(), "CONFIG", "GET")
	return
}

type ConfigGetParameter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ConfigGetParameter) Build() Completed {
	return Completed(c)
}

type ConfigResetstat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ConfigResetstat) Build() Completed {
	return Completed(c)
}

func (b *Builder) ConfigResetstat() (c ConfigResetstat) {
	c.cs = append(b.get(), "CONFIG", "RESETSTAT")
	return
}

type ConfigRewrite struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ConfigRewrite) Build() Completed {
	return Completed(c)
}

func (b *Builder) ConfigRewrite() (c ConfigRewrite) {
	c.cs = append(b.get(), "CONFIG", "REWRITE")
	return
}

type ConfigSet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ConfigSet) Parameter(Parameter string) ConfigSetParameter {
	return ConfigSetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
}

func (b *Builder) ConfigSet() (c ConfigSet) {
	c.cs = append(b.get(), "CONFIG", "SET")
	return
}

type ConfigSetParameter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ConfigSetParameter) Value(Value string) ConfigSetValue {
	return ConfigSetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type ConfigSetValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ConfigSetValue) Build() Completed {
	return Completed(c)
}

type Copy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Copy) Source(Source string) CopySource {
	return CopySource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Copy() (c Copy) {
	c.cs = append(b.get(), "COPY")
	return
}

type CopyDb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CopyDb) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c CopyDb) Build() Completed {
	return Completed(c)
}

type CopyDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CopyDestination) Db(DestinationDb int64) CopyDb {
	return CopyDb{cf: c.cf, cs: append(c.cs, "DB", strconv.FormatInt(DestinationDb, 10))}
}

func (c CopyDestination) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c CopyDestination) Build() Completed {
	return Completed(c)
}

type CopyReplaceReplace struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CopyReplaceReplace) Build() Completed {
	return Completed(c)
}

type CopySource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c CopySource) Destination(Destination string) CopyDestination {
	return CopyDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Dbsize struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Dbsize) Build() Completed {
	return Completed(c)
}

func (b *Builder) Dbsize() (c Dbsize) {
	c.cs = append(b.get(), "DBSIZE")
	return
}

type DebugObject struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DebugObject) Key(Key string) DebugObjectKey {
	return DebugObjectKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) DebugObject() (c DebugObject) {
	c.cs = append(b.get(), "DEBUG", "OBJECT")
	return
}

type DebugObjectKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DebugObjectKey) Build() Completed {
	return Completed(c)
}

type DebugSegfault struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DebugSegfault) Build() Completed {
	return Completed(c)
}

func (b *Builder) DebugSegfault() (c DebugSegfault) {
	c.cs = append(b.get(), "DEBUG", "SEGFAULT")
	return
}

type Decr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Decr) Key(Key string) DecrKey {
	return DecrKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Decr() (c Decr) {
	c.cs = append(b.get(), "DECR")
	return
}

type DecrKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DecrKey) Build() Completed {
	return Completed(c)
}

type Decrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Decrby) Key(Key string) DecrbyKey {
	return DecrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Decrby() (c Decrby) {
	c.cs = append(b.get(), "DECRBY")
	return
}

type DecrbyDecrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DecrbyDecrement) Build() Completed {
	return Completed(c)
}

type DecrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DecrbyKey) Decrement(Decrement int64) DecrbyDecrement {
	return DecrbyDecrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Decrement, 10))}
}

type Del struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Del) Key(Key ...string) DelKey {
	return DelKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Del() (c Del) {
	c.cs = append(b.get(), "DEL")
	return
}

type DelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DelKey) Key(Key ...string) DelKey {
	return DelKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c DelKey) Build() Completed {
	return Completed(c)
}

type Discard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Discard) Build() Completed {
	return Completed(c)
}

func (b *Builder) Discard() (c Discard) {
	c.cs = append(b.get(), "DISCARD")
	return
}

type Dump struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Dump) Key(Key string) DumpKey {
	return DumpKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Dump() (c Dump) {
	c.cs = append(b.get(), "DUMP")
	return
}

type DumpKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c DumpKey) Build() Completed {
	return Completed(c)
}

type Echo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Echo) Message(Message string) EchoMessage {
	return EchoMessage{cf: c.cf, cs: append(c.cs, Message)}
}

func (b *Builder) Echo() (c Echo) {
	c.cs = append(b.get(), "ECHO")
	return
}

type EchoMessage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EchoMessage) Build() Completed {
	return Completed(c)
}

type Eval struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Eval) Script(Script string) EvalScript {
	return EvalScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *Builder) Eval() (c Eval) {
	c.cs = append(b.get(), "EVAL")
	return
}

type EvalArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalArg) Arg(Arg ...string) EvalArg {
	return EvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalArg) Build() Completed {
	return Completed(c)
}

type EvalKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalKey) Arg(Arg ...string) EvalArg {
	return EvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalKey) Key(Key ...string) EvalKey {
	return EvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalKey) Build() Completed {
	return Completed(c)
}

type EvalNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalNumkeys) Key(Key ...string) EvalKey {
	return EvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalNumkeys) Arg(Arg ...string) EvalArg {
	return EvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalNumkeys) Build() Completed {
	return Completed(c)
}

type EvalRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalRo) Script(Script string) EvalRoScript {
	return EvalRoScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *Builder) EvalRo() (c EvalRo) {
	c.cs = append(b.get(), "EVAL_RO")
	return
}

type EvalRoArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalRoArg) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalRoArg) Build() Completed {
	return Completed(c)
}

type EvalRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalRoKey) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalRoKey) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalRoNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalRoNumkeys) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalRoScript struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalRoScript) Numkeys(Numkeys int64) EvalRoNumkeys {
	return EvalRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalScript struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalScript) Numkeys(Numkeys int64) EvalNumkeys {
	return EvalNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type Evalsha struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Evalsha) Sha1(Sha1 string) EvalshaSha1 {
	return EvalshaSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *Builder) Evalsha() (c Evalsha) {
	c.cs = append(b.get(), "EVALSHA")
	return
}

type EvalshaArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaArg) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaArg) Build() Completed {
	return Completed(c)
}

type EvalshaKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaKey) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaKey) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalshaKey) Build() Completed {
	return Completed(c)
}

type EvalshaNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaNumkeys) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalshaNumkeys) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaNumkeys) Build() Completed {
	return Completed(c)
}

type EvalshaRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaRo) Sha1(Sha1 string) EvalshaRoSha1 {
	return EvalshaRoSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *Builder) EvalshaRo() (c EvalshaRo) {
	c.cs = append(b.get(), "EVALSHA_RO")
	return
}

type EvalshaRoArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaRoArg) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaRoArg) Build() Completed {
	return Completed(c)
}

type EvalshaRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaRoKey) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaRoKey) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalshaRoNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaRoNumkeys) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalshaRoSha1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaRoSha1) Numkeys(Numkeys int64) EvalshaRoNumkeys {
	return EvalshaRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalshaSha1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c EvalshaSha1) Numkeys(Numkeys int64) EvalshaNumkeys {
	return EvalshaNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type Exec struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Exec) Build() Completed {
	return Completed(c)
}

func (b *Builder) Exec() (c Exec) {
	c.cs = append(b.get(), "EXEC")
	return
}

type Exists struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Exists) Key(Key ...string) ExistsKey {
	return ExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Exists() (c Exists) {
	c.cs = append(b.get(), "EXISTS")
	return
}

type ExistsKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExistsKey) Key(Key ...string) ExistsKey {
	return ExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ExistsKey) Build() Completed {
	return Completed(c)
}

type Expire struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Expire) Key(Key string) ExpireKey {
	return ExpireKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Expire() (c Expire) {
	c.cs = append(b.get(), "EXPIRE")
	return
}

type ExpireConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireConditionGt) Build() Completed {
	return Completed(c)
}

type ExpireConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireConditionLt) Build() Completed {
	return Completed(c)
}

type ExpireConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireConditionNx) Build() Completed {
	return Completed(c)
}

type ExpireConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireConditionXx) Build() Completed {
	return Completed(c)
}

type ExpireKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireKey) Seconds(Seconds int64) ExpireSeconds {
	return ExpireSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type ExpireSeconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireSeconds) Nx() ExpireConditionNx {
	return ExpireConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c ExpireSeconds) Xx() ExpireConditionXx {
	return ExpireConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c ExpireSeconds) Gt() ExpireConditionGt {
	return ExpireConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c ExpireSeconds) Lt() ExpireConditionLt {
	return ExpireConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c ExpireSeconds) Build() Completed {
	return Completed(c)
}

type Expireat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Expireat) Key(Key string) ExpireatKey {
	return ExpireatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Expireat() (c Expireat) {
	c.cs = append(b.get(), "EXPIREAT")
	return
}

type ExpireatConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireatConditionGt) Build() Completed {
	return Completed(c)
}

type ExpireatConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireatConditionLt) Build() Completed {
	return Completed(c)
}

type ExpireatConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireatConditionNx) Build() Completed {
	return Completed(c)
}

type ExpireatConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireatConditionXx) Build() Completed {
	return Completed(c)
}

type ExpireatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireatKey) Timestamp(Timestamp int64) ExpireatTimestamp {
	return ExpireatTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timestamp, 10))}
}

type ExpireatTimestamp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpireatTimestamp) Nx() ExpireatConditionNx {
	return ExpireatConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c ExpireatTimestamp) Xx() ExpireatConditionXx {
	return ExpireatConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c ExpireatTimestamp) Gt() ExpireatConditionGt {
	return ExpireatConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c ExpireatTimestamp) Lt() ExpireatConditionLt {
	return ExpireatConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c ExpireatTimestamp) Build() Completed {
	return Completed(c)
}

type Expiretime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Expiretime) Key(Key string) ExpiretimeKey {
	return ExpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Expiretime() (c Expiretime) {
	c.cs = append(b.get(), "EXPIRETIME")
	return
}

type ExpiretimeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ExpiretimeKey) Build() Completed {
	return Completed(c)
}

func (c ExpiretimeKey) Cache() Cacheable {
	return Cacheable(c)
}

type Failover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Failover) To() FailoverTargetTo {
	return FailoverTargetTo{cf: c.cf, cs: append(c.cs, "TO")}
}

func (c Failover) Abort() FailoverAbort {
	return FailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c Failover) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (b *Builder) Failover() (c Failover) {
	c.cs = append(b.get(), "FAILOVER")
	return
}

type FailoverAbort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FailoverAbort) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverAbort) Build() Completed {
	return Completed(c)
}

type FailoverTargetForce struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FailoverTargetForce) Abort() FailoverAbort {
	return FailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c FailoverTargetForce) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverTargetForce) Build() Completed {
	return Completed(c)
}

type FailoverTargetHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FailoverTargetHost) Port(Port int64) FailoverTargetPort {
	return FailoverTargetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type FailoverTargetPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FailoverTargetPort) Force() FailoverTargetForce {
	return FailoverTargetForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c FailoverTargetPort) Abort() FailoverAbort {
	return FailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c FailoverTargetPort) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverTargetPort) Build() Completed {
	return Completed(c)
}

type FailoverTargetTo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FailoverTargetTo) Host(Host string) FailoverTargetHost {
	return FailoverTargetHost{cf: c.cf, cs: append(c.cs, Host)}
}

type FailoverTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FailoverTimeout) Build() Completed {
	return Completed(c)
}

type Flushall struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Flushall) Async() FlushallAsyncAsync {
	return FlushallAsyncAsync{cf: c.cf, cs: append(c.cs, "ASYNC")}
}

func (c Flushall) Sync() FlushallAsyncSync {
	return FlushallAsyncSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c Flushall) Build() Completed {
	return Completed(c)
}

func (b *Builder) Flushall() (c Flushall) {
	c.cs = append(b.get(), "FLUSHALL")
	return
}

type FlushallAsyncAsync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FlushallAsyncAsync) Build() Completed {
	return Completed(c)
}

type FlushallAsyncSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FlushallAsyncSync) Build() Completed {
	return Completed(c)
}

type Flushdb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Flushdb) Async() FlushdbAsyncAsync {
	return FlushdbAsyncAsync{cf: c.cf, cs: append(c.cs, "ASYNC")}
}

func (c Flushdb) Sync() FlushdbAsyncSync {
	return FlushdbAsyncSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c Flushdb) Build() Completed {
	return Completed(c)
}

func (b *Builder) Flushdb() (c Flushdb) {
	c.cs = append(b.get(), "FLUSHDB")
	return
}

type FlushdbAsyncAsync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FlushdbAsyncAsync) Build() Completed {
	return Completed(c)
}

type FlushdbAsyncSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c FlushdbAsyncSync) Build() Completed {
	return Completed(c)
}

type Geoadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Geoadd) Key(Key string) GeoaddKey {
	return GeoaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geoadd() (c Geoadd) {
	c.cs = append(b.get(), "GEOADD")
	return
}

type GeoaddChangeCh struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoaddChangeCh) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type GeoaddConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoaddConditionNx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c GeoaddConditionNx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type GeoaddConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoaddConditionXx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c GeoaddConditionXx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type GeoaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoaddKey) Nx() GeoaddConditionNx {
	return GeoaddConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c GeoaddKey) Xx() GeoaddConditionXx {
	return GeoaddConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c GeoaddKey) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c GeoaddKey) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type GeoaddLongitudeLatitudeMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member)}
}

func (c GeoaddLongitudeLatitudeMember) Build() Completed {
	return Completed(c)
}

type Geodist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Geodist) Key(Key string) GeodistKey {
	return GeodistKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geodist() (c Geodist) {
	c.cs = append(b.get(), "GEODIST")
	return
}

type GeodistKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeodistKey) Member1(Member1 string) GeodistMember1 {
	return GeodistMember1{cf: c.cf, cs: append(c.cs, Member1)}
}

func (c GeodistKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistMember1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeodistMember1) Member2(Member2 string) GeodistMember2 {
	return GeodistMember2{cf: c.cf, cs: append(c.cs, Member2)}
}

func (c GeodistMember1) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistMember2 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeodistMember2) M() GeodistUnitM {
	return GeodistUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeodistMember2) Km() GeodistUnitKm {
	return GeodistUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeodistMember2) Ft() GeodistUnitFt {
	return GeodistUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeodistMember2) Mi() GeodistUnitMi {
	return GeodistUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c GeodistMember2) Build() Completed {
	return Completed(c)
}

func (c GeodistMember2) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeodistUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeodistUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeodistUnitM) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeodistUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type Geohash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Geohash) Key(Key string) GeohashKey {
	return GeohashKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geohash() (c Geohash) {
	c.cs = append(b.get(), "GEOHASH")
	return
}

type GeohashKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeohashKey) Member(Member ...string) GeohashMember {
	return GeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeohashKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeohashMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeohashMember) Member(Member ...string) GeohashMember {
	return GeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeohashMember) Build() Completed {
	return Completed(c)
}

func (c GeohashMember) Cache() Cacheable {
	return Cacheable(c)
}

type Geopos struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Geopos) Key(Key string) GeoposKey {
	return GeoposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geopos() (c Geopos) {
	c.cs = append(b.get(), "GEOPOS")
	return
}

type GeoposKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoposKey) Member(Member ...string) GeoposMember {
	return GeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeoposKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeoposMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoposMember) Member(Member ...string) GeoposMember {
	return GeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeoposMember) Build() Completed {
	return Completed(c)
}

func (c GeoposMember) Cache() Cacheable {
	return Cacheable(c)
}

type Georadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Georadius) Key(Key string) GeoradiusKey {
	return GeoradiusKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Georadius() (c Georadius) {
	c.cs = append(b.get(), "GEORADIUS")
	return
}

type GeoradiusCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusCountAnyAny) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusCountAnyAny) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusCountAnyAny) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusCountAnyAny) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeoradiusCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusCountCount) Any() GeoradiusCountAnyAny {
	return GeoradiusCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c GeoradiusCountCount) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusCountCount) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusCountCount) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusCountCount) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusCountCount) Build() Completed {
	return Completed(c)
}

type GeoradiusKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusKey) Longitude(Longitude float64) GeoradiusLongitude {
	return GeoradiusLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type GeoradiusLatitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusLatitude) Radius(Radius float64) GeoradiusRadius {
	return GeoradiusRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusLongitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusLongitude) Latitude(Latitude float64) GeoradiusLatitude {
	return GeoradiusLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type GeoradiusOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusOrderAsc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusOrderAsc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusOrderAsc) Build() Completed {
	return Completed(c)
}

type GeoradiusOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusOrderDesc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusOrderDesc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusOrderDesc) Build() Completed {
	return Completed(c)
}

type GeoradiusRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRadius) M() GeoradiusUnitM {
	return GeoradiusUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeoradiusRadius) Km() GeoradiusUnitKm {
	return GeoradiusUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeoradiusRadius) Ft() GeoradiusUnitFt {
	return GeoradiusUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeoradiusRadius) Mi() GeoradiusUnitMi {
	return GeoradiusUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type GeoradiusRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRo) Key(Key string) GeoradiusRoKey {
	return GeoradiusRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) GeoradiusRo() (c GeoradiusRo) {
	c.cs = append(b.get(), "GEORADIUS_RO")
	return
}

type GeoradiusRoCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoCountAnyAny) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoCountAnyAny) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoCountAnyAny) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoCountAnyAny) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoCountAnyAny) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoCountCount) Any() GeoradiusRoCountAnyAny {
	return GeoradiusRoCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c GeoradiusRoCountCount) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoCountCount) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoCountCount) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoCountCount) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoCountCount) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoKey) Longitude(Longitude float64) GeoradiusRoLongitude {
	return GeoradiusRoLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

func (c GeoradiusRoKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoLatitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoLatitude) Radius(Radius float64) GeoradiusRoRadius {
	return GeoradiusRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeoradiusRoLatitude) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoLongitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoLongitude) Latitude(Latitude float64) GeoradiusRoLatitude {
	return GeoradiusRoLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeoradiusRoLongitude) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoOrderAsc) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoOrderDesc) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoRadius) M() GeoradiusRoUnitM {
	return GeoradiusRoUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeoradiusRoRadius) Km() GeoradiusRoUnitKm {
	return GeoradiusRoUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeoradiusRoRadius) Ft() GeoradiusRoUnitFt {
	return GeoradiusRoUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeoradiusRoRadius) Mi() GeoradiusRoUnitMi {
	return GeoradiusRoUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c GeoradiusRoRadius) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoStoredist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoStoredist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoUnitFt) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusRoUnitFt) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusRoUnitFt) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusRoUnitFt) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusRoUnitFt) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoUnitFt) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoUnitFt) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoUnitKm) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusRoUnitKm) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusRoUnitKm) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusRoUnitKm) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusRoUnitKm) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoUnitKm) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoUnitKm) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoUnitM) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusRoUnitM) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusRoUnitM) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusRoUnitM) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusRoUnitM) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoUnitM) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoUnitM) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoUnitMi) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusRoUnitMi) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusRoUnitMi) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusRoUnitMi) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusRoUnitMi) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoUnitMi) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoUnitMi) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoWithcoordWithcoord) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusRoWithcoordWithcoord) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusRoWithcoordWithcoord) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusRoWithcoordWithcoord) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoWithcoordWithcoord) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoWithcoordWithcoord) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoWithdistWithdist) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusRoWithdistWithdist) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusRoWithdistWithdist) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoWithdistWithdist) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoWithdistWithdist) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusRoWithhashWithhash) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusRoWithhashWithhash) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusRoWithhashWithhash) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusRoWithhashWithhash) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusStore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusStore) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusStore) Build() Completed {
	return Completed(c)
}

type GeoradiusStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusStoredist) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusUnitFt) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitFt) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitFt) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitFt) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitFt) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitFt) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitFt) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitFt) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusUnitKm) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitKm) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitKm) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitKm) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitKm) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitKm) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitKm) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitKm) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusUnitM) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitM) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitM) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitM) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitM) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitM) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitM) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitM) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusUnitMi) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitMi) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitMi) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitMi) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitMi) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitMi) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitMi) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitMi) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusWithcoordWithcoord) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusWithcoordWithcoord) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusWithcoordWithcoord) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusWithcoordWithcoord) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusWithcoordWithcoord) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusWithcoordWithcoord) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusWithcoordWithcoord) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusWithdistWithdist) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusWithdistWithdist) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusWithdistWithdist) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusWithdistWithdist) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusWithdistWithdist) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusWithdistWithdist) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusWithhashWithhash) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusWithhashWithhash) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusWithhashWithhash) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusWithhashWithhash) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusWithhashWithhash) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type Georadiusbymember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Georadiusbymember) Key(Key string) GeoradiusbymemberKey {
	return GeoradiusbymemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Georadiusbymember() (c Georadiusbymember) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER")
	return
}

type GeoradiusbymemberCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberCountAnyAny) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberCountAnyAny) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberCountAnyAny) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberCountAnyAny) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberCountCount) Any() GeoradiusbymemberCountAnyAny {
	return GeoradiusbymemberCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c GeoradiusbymemberCountCount) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberCountCount) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberCountCount) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberCountCount) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberCountCount) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberKey) Member(Member string) GeoradiusbymemberMember {
	return GeoradiusbymemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

type GeoradiusbymemberMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberMember) Radius(Radius float64) GeoradiusbymemberRadius {
	return GeoradiusbymemberRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusbymemberOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberOrderAsc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberOrderAsc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberOrderAsc) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberOrderDesc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberOrderDesc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberOrderDesc) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRadius) M() GeoradiusbymemberUnitM {
	return GeoradiusbymemberUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeoradiusbymemberRadius) Km() GeoradiusbymemberUnitKm {
	return GeoradiusbymemberUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeoradiusbymemberRadius) Ft() GeoradiusbymemberUnitFt {
	return GeoradiusbymemberUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeoradiusbymemberRadius) Mi() GeoradiusbymemberUnitMi {
	return GeoradiusbymemberUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type GeoradiusbymemberRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRo) Key(Key string) GeoradiusbymemberRoKey {
	return GeoradiusbymemberRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) GeoradiusbymemberRo() (c GeoradiusbymemberRo) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER_RO")
	return
}

type GeoradiusbymemberRoCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoCountAnyAny) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoCountAnyAny) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoCountAnyAny) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoCountAnyAny) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoCountAnyAny) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoCountCount) Any() GeoradiusbymemberRoCountAnyAny {
	return GeoradiusbymemberRoCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c GeoradiusbymemberRoCountCount) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoCountCount) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoCountCount) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoCountCount) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoCountCount) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoKey) Member(Member string) GeoradiusbymemberRoMember {
	return GeoradiusbymemberRoMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c GeoradiusbymemberRoKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoMember) Radius(Radius float64) GeoradiusbymemberRoRadius {
	return GeoradiusbymemberRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeoradiusbymemberRoMember) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoOrderAsc) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoOrderDesc) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoRadius) M() GeoradiusbymemberRoUnitM {
	return GeoradiusbymemberRoUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeoradiusbymemberRoRadius) Km() GeoradiusbymemberRoUnitKm {
	return GeoradiusbymemberRoUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeoradiusbymemberRoRadius) Ft() GeoradiusbymemberRoUnitFt {
	return GeoradiusbymemberRoUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeoradiusbymemberRoRadius) Mi() GeoradiusbymemberRoUnitMi {
	return GeoradiusbymemberRoUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c GeoradiusbymemberRoRadius) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoStoredist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoStoredist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoUnitFt) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberRoUnitFt) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberRoUnitFt) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberRoUnitFt) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberRoUnitFt) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoUnitFt) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoUnitFt) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoUnitKm) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberRoUnitKm) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberRoUnitKm) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberRoUnitKm) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberRoUnitKm) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoUnitKm) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoUnitKm) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoUnitM) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberRoUnitM) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberRoUnitM) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberRoUnitM) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberRoUnitM) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoUnitM) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoUnitM) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoUnitMi) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberRoUnitMi) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberRoUnitMi) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberRoUnitMi) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberRoUnitMi) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoUnitMi) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoUnitMi) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoWithdistWithdist) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberRoWithdistWithdist) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberRoWithdistWithdist) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoWithdistWithdist) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoWithdistWithdist) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberRoWithhashWithhash) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberRoWithhashWithhash) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberRoWithhashWithhash) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberRoWithhashWithhash) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberStore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberStore) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberStore) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberStoredist) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberUnitFt) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitFt) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitFt) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitFt) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitFt) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitFt) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitFt) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitFt) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberUnitKm) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitKm) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitKm) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitKm) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitKm) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitKm) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitKm) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitKm) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberUnitM) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitM) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitM) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitM) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitM) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitM) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitM) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitM) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberUnitMi) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitMi) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitMi) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitMi) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitMi) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitMi) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitMi) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitMi) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberWithcoordWithcoord) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberWithcoordWithcoord) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberWithcoordWithcoord) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberWithdistWithdist) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberWithdistWithdist) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberWithdistWithdist) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberWithdistWithdist) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberWithdistWithdist) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberWithdistWithdist) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeoradiusbymemberWithhashWithhash) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberWithhashWithhash) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberWithhashWithhash) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberWithhashWithhash) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberWithhashWithhash) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type Geosearch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Geosearch) Key(Key string) GeosearchKey {
	return GeosearchKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geosearch() (c Geosearch) {
	c.cs = append(b.get(), "GEOSEARCH")
	return
}

type GeosearchBoxBybox struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchBoxBybox) Height(Height float64) GeosearchBoxHeight {
	return GeosearchBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

func (c GeosearchBoxBybox) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxHeight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchBoxHeight) M() GeosearchBoxUnitM {
	return GeosearchBoxUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeosearchBoxHeight) Km() GeosearchBoxUnitKm {
	return GeosearchBoxUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeosearchBoxHeight) Ft() GeosearchBoxUnitFt {
	return GeosearchBoxUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeosearchBoxHeight) Mi() GeosearchBoxUnitMi {
	return GeosearchBoxUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c GeosearchBoxHeight) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchBoxUnitFt) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitFt) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitFt) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitFt) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitFt) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitFt) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchBoxUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchBoxUnitKm) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitKm) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitKm) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitKm) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitKm) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitKm) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchBoxUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchBoxUnitM) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitM) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitM) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitM) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitM) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitM) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchBoxUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchBoxUnitMi) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitMi) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitMi) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitMi) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitMi) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitMi) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchBoxUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleByradius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchCircleByradius) M() GeosearchCircleUnitM {
	return GeosearchCircleUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeosearchCircleByradius) Km() GeosearchCircleUnitKm {
	return GeosearchCircleUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeosearchCircleByradius) Ft() GeosearchCircleUnitFt {
	return GeosearchCircleUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeosearchCircleByradius) Mi() GeosearchCircleUnitMi {
	return GeosearchCircleUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c GeosearchCircleByradius) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchCircleUnitFt) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitFt) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitFt) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitFt) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitFt) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitFt) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitFt) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCircleUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchCircleUnitKm) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitKm) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitKm) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitKm) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitKm) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitKm) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitKm) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCircleUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchCircleUnitM) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitM) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitM) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitM) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitM) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitM) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitM) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCircleUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchCircleUnitMi) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitMi) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitMi) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitMi) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitMi) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitMi) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitMi) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCircleUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchCountAnyAny) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCountAnyAny) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCountAnyAny) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCountAnyAny) Build() Completed {
	return Completed(c)
}

func (c GeosearchCountAnyAny) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchCountCount) Any() GeosearchCountAnyAny {
	return GeosearchCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c GeosearchCountCount) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCountCount) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCountCount) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCountCount) Build() Completed {
	return Completed(c)
}

func (c GeosearchCountCount) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchFromlonlat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchFromlonlat) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchFromlonlat) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchFromlonlat) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchFromlonlat) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchFromlonlat) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchFromlonlat) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchFromlonlat) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchFromlonlat) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchFromlonlat) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchFrommember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchFrommember) Fromlonlat(Longitude float64, Latitude float64) GeosearchFromlonlat {
	return GeosearchFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchFrommember) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchFrommember) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchFrommember) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchFrommember) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchFrommember) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchFrommember) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchFrommember) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchFrommember) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchFrommember) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchKey) Frommember(Member string) GeosearchFrommember {
	return GeosearchFrommember{cf: c.cf, cs: append(c.cs, "FROMMEMBER", Member)}
}

func (c GeosearchKey) Fromlonlat(Longitude float64, Latitude float64) GeosearchFromlonlat {
	return GeosearchFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchKey) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchKey) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchKey) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchKey) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchKey) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchKey) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchKey) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchKey) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchOrderAsc) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchOrderAsc) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchOrderAsc) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchOrderAsc) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchOrderDesc) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchOrderDesc) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchOrderDesc) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchOrderDesc) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchWithcoordWithcoord) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchWithcoordWithcoord) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

func (c GeosearchWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchWithdistWithdist) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchWithdistWithdist) Build() Completed {
	return Completed(c)
}

func (c GeosearchWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchWithhashWithhash) Build() Completed {
	return Completed(c)
}

func (c GeosearchWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type Geosearchstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Geosearchstore) Destination(Destination string) GeosearchstoreDestination {
	return GeosearchstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Geosearchstore() (c Geosearchstore) {
	c.cs = append(b.get(), "GEOSEARCHSTORE")
	return
}

type GeosearchstoreBoxBybox struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreBoxBybox) Height(Height float64) GeosearchstoreBoxHeight {
	return GeosearchstoreBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type GeosearchstoreBoxHeight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreBoxHeight) M() GeosearchstoreBoxUnitM {
	return GeosearchstoreBoxUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeosearchstoreBoxHeight) Km() GeosearchstoreBoxUnitKm {
	return GeosearchstoreBoxUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeosearchstoreBoxHeight) Ft() GeosearchstoreBoxUnitFt {
	return GeosearchstoreBoxUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeosearchstoreBoxHeight) Mi() GeosearchstoreBoxUnitMi {
	return GeosearchstoreBoxUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type GeosearchstoreBoxUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreBoxUnitFt) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitFt) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitFt) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitFt) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreBoxUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreBoxUnitKm) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitKm) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitKm) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitKm) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreBoxUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreBoxUnitM) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitM) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitM) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitM) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreBoxUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreBoxUnitMi) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitMi) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitMi) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitMi) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleByradius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreCircleByradius) M() GeosearchstoreCircleUnitM {
	return GeosearchstoreCircleUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c GeosearchstoreCircleByradius) Km() GeosearchstoreCircleUnitKm {
	return GeosearchstoreCircleUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c GeosearchstoreCircleByradius) Ft() GeosearchstoreCircleUnitFt {
	return GeosearchstoreCircleUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c GeosearchstoreCircleByradius) Mi() GeosearchstoreCircleUnitMi {
	return GeosearchstoreCircleUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type GeosearchstoreCircleUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreCircleUnitFt) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitFt) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitFt) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitFt) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitFt) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreCircleUnitKm) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitKm) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitKm) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitKm) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitKm) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreCircleUnitM) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitM) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitM) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitM) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitM) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreCircleUnitMi) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitMi) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitMi) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitMi) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitMi) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreCountAnyAny) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreCountCount) Any() GeosearchstoreCountAnyAny {
	return GeosearchstoreCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c GeosearchstoreCountCount) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreCountCount) Build() Completed {
	return Completed(c)
}

type GeosearchstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreDestination) Source(Source string) GeosearchstoreSource {
	return GeosearchstoreSource{cf: c.cf, cs: append(c.cs, Source)}
}

type GeosearchstoreFromlonlat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreFromlonlat) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchstoreFromlonlat) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreFromlonlat) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreFromlonlat) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreFromlonlat) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreFromlonlat) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreFrommember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreFrommember) Fromlonlat(Longitude float64, Latitude float64) GeosearchstoreFromlonlat {
	return GeosearchstoreFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchstoreFrommember) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchstoreFrommember) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreFrommember) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreFrommember) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreFrommember) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreFrommember) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreOrderAsc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderAsc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreOrderDesc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderDesc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreSource) Frommember(Member string) GeosearchstoreFrommember {
	return GeosearchstoreFrommember{cf: c.cf, cs: append(c.cs, "FROMMEMBER", Member)}
}

func (c GeosearchstoreSource) Fromlonlat(Longitude float64, Latitude float64) GeosearchstoreFromlonlat {
	return GeosearchstoreFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchstoreSource) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchstoreSource) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreSource) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreSource) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreSource) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreSource) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreStoredistStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GeosearchstoreStoredistStoredist) Build() Completed {
	return Completed(c)
}

type Get struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Get) Key(Key string) GetKey {
	return GetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Get() (c Get) {
	c.cs = append(b.get(), "GET")
	return
}

type GetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetKey) Build() Completed {
	return Completed(c)
}

func (c GetKey) Cache() Cacheable {
	return Cacheable(c)
}

type Getbit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Getbit) Key(Key string) GetbitKey {
	return GetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getbit() (c Getbit) {
	c.cs = append(b.get(), "GETBIT")
	return
}

type GetbitKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetbitKey) Offset(Offset int64) GetbitOffset {
	return GetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

func (c GetbitKey) Cache() Cacheable {
	return Cacheable(c)
}

type GetbitOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetbitOffset) Build() Completed {
	return Completed(c)
}

func (c GetbitOffset) Cache() Cacheable {
	return Cacheable(c)
}

type Getdel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Getdel) Key(Key string) GetdelKey {
	return GetdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getdel() (c Getdel) {
	c.cs = append(b.get(), "GETDEL")
	return
}

type GetdelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetdelKey) Build() Completed {
	return Completed(c)
}

type Getex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Getex) Key(Key string) GetexKey {
	return GetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getex() (c Getex) {
	c.cs = append(b.get(), "GETEX")
	return
}

type GetexExpirationEx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetexExpirationEx) Build() Completed {
	return Completed(c)
}

type GetexExpirationExat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetexExpirationExat) Build() Completed {
	return Completed(c)
}

type GetexExpirationPersist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetexExpirationPersist) Build() Completed {
	return Completed(c)
}

type GetexExpirationPx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetexExpirationPx) Build() Completed {
	return Completed(c)
}

type GetexExpirationPxat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetexExpirationPxat) Build() Completed {
	return Completed(c)
}

type GetexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetexKey) Ex(Seconds int64) GetexExpirationEx {
	return GetexExpirationEx{cf: c.cf, cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10))}
}

func (c GetexKey) Px(Milliseconds int64) GetexExpirationPx {
	return GetexExpirationPx{cf: c.cf, cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10))}
}

func (c GetexKey) Exat(Timestamp int64) GetexExpirationExat {
	return GetexExpirationExat{cf: c.cf, cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10))}
}

func (c GetexKey) Pxat(Millisecondstimestamp int64) GetexExpirationPxat {
	return GetexExpirationPxat{cf: c.cf, cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10))}
}

func (c GetexKey) Persist() GetexExpirationPersist {
	return GetexExpirationPersist{cf: c.cf, cs: append(c.cs, "PERSIST")}
}

func (c GetexKey) Build() Completed {
	return Completed(c)
}

type Getrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Getrange) Key(Key string) GetrangeKey {
	return GetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getrange() (c Getrange) {
	c.cs = append(b.get(), "GETRANGE")
	return
}

type GetrangeEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetrangeEnd) Build() Completed {
	return Completed(c)
}

func (c GetrangeEnd) Cache() Cacheable {
	return Cacheable(c)
}

type GetrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetrangeKey) Start(Start int64) GetrangeStart {
	return GetrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c GetrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type GetrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetrangeStart) End(End int64) GetrangeEnd {
	return GetrangeEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c GetrangeStart) Cache() Cacheable {
	return Cacheable(c)
}

type Getset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Getset) Key(Key string) GetsetKey {
	return GetsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getset() (c Getset) {
	c.cs = append(b.get(), "GETSET")
	return
}

type GetsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetsetKey) Value(Value string) GetsetValue {
	return GetsetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type GetsetValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c GetsetValue) Build() Completed {
	return Completed(c)
}

type Hdel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hdel) Key(Key string) HdelKey {
	return HdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hdel() (c Hdel) {
	c.cs = append(b.get(), "HDEL")
	return
}

type HdelField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HdelField) Field(Field ...string) HdelField {
	return HdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c HdelField) Build() Completed {
	return Completed(c)
}

type HdelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HdelKey) Field(Field ...string) HdelField {
	return HdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

type Hello struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hello) Protover(Protover int64) HelloArgumentsProtover {
	return HelloArgumentsProtover{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Protover, 10))}
}

func (b *Builder) Hello() (c Hello) {
	c.cs = append(b.get(), "HELLO")
	return
}

type HelloArgumentsAuth struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HelloArgumentsAuth) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
}

func (c HelloArgumentsAuth) Build() Completed {
	return Completed(c)
}

type HelloArgumentsProtover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HelloArgumentsProtover) Auth(Username string, Password string) HelloArgumentsAuth {
	return HelloArgumentsAuth{cf: c.cf, cs: append(c.cs, "AUTH", Username, Password)}
}

func (c HelloArgumentsProtover) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
}

func (c HelloArgumentsProtover) Build() Completed {
	return Completed(c)
}

type HelloArgumentsSetname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HelloArgumentsSetname) Build() Completed {
	return Completed(c)
}

type Hexists struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hexists) Key(Key string) HexistsKey {
	return HexistsKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hexists() (c Hexists) {
	c.cs = append(b.get(), "HEXISTS")
	return
}

type HexistsField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HexistsField) Build() Completed {
	return Completed(c)
}

func (c HexistsField) Cache() Cacheable {
	return Cacheable(c)
}

type HexistsKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HexistsKey) Field(Field string) HexistsField {
	return HexistsField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c HexistsKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hget) Key(Key string) HgetKey {
	return HgetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hget() (c Hget) {
	c.cs = append(b.get(), "HGET")
	return
}

type HgetField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HgetField) Build() Completed {
	return Completed(c)
}

func (c HgetField) Cache() Cacheable {
	return Cacheable(c)
}

type HgetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HgetKey) Field(Field string) HgetField {
	return HgetField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c HgetKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hgetall struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hgetall) Key(Key string) HgetallKey {
	return HgetallKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hgetall() (c Hgetall) {
	c.cs = append(b.get(), "HGETALL")
	return
}

type HgetallKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HgetallKey) Build() Completed {
	return Completed(c)
}

func (c HgetallKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hincrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hincrby) Key(Key string) HincrbyKey {
	return HincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hincrby() (c Hincrby) {
	c.cs = append(b.get(), "HINCRBY")
	return
}

type HincrbyField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HincrbyField) Increment(Increment int64) HincrbyIncrement {
	return HincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type HincrbyIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HincrbyIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HincrbyKey) Field(Field string) HincrbyField {
	return HincrbyField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hincrbyfloat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hincrbyfloat) Key(Key string) HincrbyfloatKey {
	return HincrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hincrbyfloat() (c Hincrbyfloat) {
	c.cs = append(b.get(), "HINCRBYFLOAT")
	return
}

type HincrbyfloatField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HincrbyfloatField) Increment(Increment float64) HincrbyfloatIncrement {
	return HincrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type HincrbyfloatIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HincrbyfloatIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyfloatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HincrbyfloatKey) Field(Field string) HincrbyfloatField {
	return HincrbyfloatField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hkeys) Key(Key string) HkeysKey {
	return HkeysKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hkeys() (c Hkeys) {
	c.cs = append(b.get(), "HKEYS")
	return
}

type HkeysKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HkeysKey) Build() Completed {
	return Completed(c)
}

func (c HkeysKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hlen) Key(Key string) HlenKey {
	return HlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hlen() (c Hlen) {
	c.cs = append(b.get(), "HLEN")
	return
}

type HlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HlenKey) Build() Completed {
	return Completed(c)
}

func (c HlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hmget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hmget) Key(Key string) HmgetKey {
	return HmgetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hmget() (c Hmget) {
	c.cs = append(b.get(), "HMGET")
	return
}

type HmgetField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HmgetField) Field(Field ...string) HmgetField {
	return HmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c HmgetField) Build() Completed {
	return Completed(c)
}

func (c HmgetField) Cache() Cacheable {
	return Cacheable(c)
}

type HmgetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HmgetKey) Field(Field ...string) HmgetField {
	return HmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c HmgetKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hmset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hmset) Key(Key string) HmsetKey {
	return HmsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hmset() (c Hmset) {
	c.cs = append(b.get(), "HMSET")
	return
}

type HmsetFieldValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HmsetFieldValue) FieldValue(Field string, Value string) HmsetFieldValue {
	return HmsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c HmsetFieldValue) Build() Completed {
	return Completed(c)
}

type HmsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HmsetKey) FieldValue() HmsetFieldValue {
	return HmsetFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type Hrandfield struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hrandfield) Key(Key string) HrandfieldKey {
	return HrandfieldKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hrandfield() (c Hrandfield) {
	c.cs = append(b.get(), "HRANDFIELD")
	return
}

type HrandfieldKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HrandfieldKey) Count(Count int64) HrandfieldOptionsCount {
	return HrandfieldOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type HrandfieldOptionsCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HrandfieldOptionsCount) Withvalues() HrandfieldOptionsWithvaluesWithvalues {
	return HrandfieldOptionsWithvaluesWithvalues{cf: c.cf, cs: append(c.cs, "WITHVALUES")}
}

func (c HrandfieldOptionsCount) Build() Completed {
	return Completed(c)
}

type HrandfieldOptionsWithvaluesWithvalues struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HrandfieldOptionsWithvaluesWithvalues) Build() Completed {
	return Completed(c)
}

type Hscan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hscan) Key(Key string) HscanKey {
	return HscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hscan() (c Hscan) {
	c.cs = append(b.get(), "HSCAN")
	return
}

type HscanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HscanCount) Build() Completed {
	return Completed(c)
}

type HscanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HscanCursor) Match(Pattern string) HscanMatch {
	return HscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c HscanCursor) Count(Count int64) HscanCount {
	return HscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c HscanCursor) Build() Completed {
	return Completed(c)
}

type HscanKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HscanKey) Cursor(Cursor int64) HscanCursor {
	return HscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type HscanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HscanMatch) Count(Count int64) HscanCount {
	return HscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c HscanMatch) Build() Completed {
	return Completed(c)
}

type Hset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hset) Key(Key string) HsetKey {
	return HsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hset() (c Hset) {
	c.cs = append(b.get(), "HSET")
	return
}

type HsetFieldValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HsetFieldValue) FieldValue(Field string, Value string) HsetFieldValue {
	return HsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c HsetFieldValue) Build() Completed {
	return Completed(c)
}

type HsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HsetKey) FieldValue() HsetFieldValue {
	return HsetFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type Hsetnx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hsetnx) Key(Key string) HsetnxKey {
	return HsetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hsetnx() (c Hsetnx) {
	c.cs = append(b.get(), "HSETNX")
	return
}

type HsetnxField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HsetnxField) Value(Value string) HsetnxValue {
	return HsetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type HsetnxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HsetnxKey) Field(Field string) HsetnxField {
	return HsetnxField{cf: c.cf, cs: append(c.cs, Field)}
}

type HsetnxValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HsetnxValue) Build() Completed {
	return Completed(c)
}

type Hstrlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hstrlen) Key(Key string) HstrlenKey {
	return HstrlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hstrlen() (c Hstrlen) {
	c.cs = append(b.get(), "HSTRLEN")
	return
}

type HstrlenField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HstrlenField) Build() Completed {
	return Completed(c)
}

func (c HstrlenField) Cache() Cacheable {
	return Cacheable(c)
}

type HstrlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HstrlenKey) Field(Field string) HstrlenField {
	return HstrlenField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c HstrlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hvals struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Hvals) Key(Key string) HvalsKey {
	return HvalsKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hvals() (c Hvals) {
	c.cs = append(b.get(), "HVALS")
	return
}

type HvalsKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c HvalsKey) Build() Completed {
	return Completed(c)
}

func (c HvalsKey) Cache() Cacheable {
	return Cacheable(c)
}

type Incr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Incr) Key(Key string) IncrKey {
	return IncrKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Incr() (c Incr) {
	c.cs = append(b.get(), "INCR")
	return
}

type IncrKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c IncrKey) Build() Completed {
	return Completed(c)
}

type Incrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Incrby) Key(Key string) IncrbyKey {
	return IncrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Incrby() (c Incrby) {
	c.cs = append(b.get(), "INCRBY")
	return
}

type IncrbyIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c IncrbyIncrement) Build() Completed {
	return Completed(c)
}

type IncrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c IncrbyKey) Increment(Increment int64) IncrbyIncrement {
	return IncrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type Incrbyfloat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Incrbyfloat) Key(Key string) IncrbyfloatKey {
	return IncrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Incrbyfloat() (c Incrbyfloat) {
	c.cs = append(b.get(), "INCRBYFLOAT")
	return
}

type IncrbyfloatIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c IncrbyfloatIncrement) Build() Completed {
	return Completed(c)
}

type IncrbyfloatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c IncrbyfloatKey) Increment(Increment float64) IncrbyfloatIncrement {
	return IncrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type Info struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Info) Section(Section string) InfoSection {
	return InfoSection{cf: c.cf, cs: append(c.cs, Section)}
}

func (c Info) Build() Completed {
	return Completed(c)
}

func (b *Builder) Info() (c Info) {
	c.cs = append(b.get(), "INFO")
	return
}

type InfoSection struct {
	cs []string
	cf uint16
	ks uint16
}

func (c InfoSection) Build() Completed {
	return Completed(c)
}

type Keys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Keys) Pattern(Pattern string) KeysPattern {
	return KeysPattern{cf: c.cf, cs: append(c.cs, Pattern)}
}

func (b *Builder) Keys() (c Keys) {
	c.cs = append(b.get(), "KEYS")
	return
}

type KeysPattern struct {
	cs []string
	cf uint16
	ks uint16
}

func (c KeysPattern) Build() Completed {
	return Completed(c)
}

type Lastsave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lastsave) Build() Completed {
	return Completed(c)
}

func (b *Builder) Lastsave() (c Lastsave) {
	c.cs = append(b.get(), "LASTSAVE")
	return
}

type LatencyDoctor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyDoctor) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyDoctor() (c LatencyDoctor) {
	c.cs = append(b.get(), "LATENCY", "DOCTOR")
	return
}

type LatencyGraph struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyGraph) Event(Event string) LatencyGraphEvent {
	return LatencyGraphEvent{cf: c.cf, cs: append(c.cs, Event)}
}

func (b *Builder) LatencyGraph() (c LatencyGraph) {
	c.cs = append(b.get(), "LATENCY", "GRAPH")
	return
}

type LatencyGraphEvent struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyGraphEvent) Build() Completed {
	return Completed(c)
}

type LatencyHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyHelp) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyHelp() (c LatencyHelp) {
	c.cs = append(b.get(), "LATENCY", "HELP")
	return
}

type LatencyHistory struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyHistory) Event(Event string) LatencyHistoryEvent {
	return LatencyHistoryEvent{cf: c.cf, cs: append(c.cs, Event)}
}

func (b *Builder) LatencyHistory() (c LatencyHistory) {
	c.cs = append(b.get(), "LATENCY", "HISTORY")
	return
}

type LatencyHistoryEvent struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyHistoryEvent) Build() Completed {
	return Completed(c)
}

type LatencyLatest struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyLatest) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyLatest() (c LatencyLatest) {
	c.cs = append(b.get(), "LATENCY", "LATEST")
	return
}

type LatencyReset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyReset) Event(Event ...string) LatencyResetEvent {
	return LatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
}

func (c LatencyReset) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyReset() (c LatencyReset) {
	c.cs = append(b.get(), "LATENCY", "RESET")
	return
}

type LatencyResetEvent struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LatencyResetEvent) Event(Event ...string) LatencyResetEvent {
	return LatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
}

func (c LatencyResetEvent) Build() Completed {
	return Completed(c)
}

type Lindex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lindex) Key(Key string) LindexKey {
	return LindexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lindex() (c Lindex) {
	c.cs = append(b.get(), "LINDEX")
	return
}

type LindexIndex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LindexIndex) Build() Completed {
	return Completed(c)
}

func (c LindexIndex) Cache() Cacheable {
	return Cacheable(c)
}

type LindexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LindexKey) Index(Index int64) LindexIndex {
	return LindexIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

func (c LindexKey) Cache() Cacheable {
	return Cacheable(c)
}

type Linsert struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Linsert) Key(Key string) LinsertKey {
	return LinsertKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Linsert() (c Linsert) {
	c.cs = append(b.get(), "LINSERT")
	return
}

type LinsertElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LinsertElement) Build() Completed {
	return Completed(c)
}

type LinsertKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LinsertKey) Before() LinsertWhereBefore {
	return LinsertWhereBefore{cf: c.cf, cs: append(c.cs, "BEFORE")}
}

func (c LinsertKey) After() LinsertWhereAfter {
	return LinsertWhereAfter{cf: c.cf, cs: append(c.cs, "AFTER")}
}

type LinsertPivot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LinsertPivot) Element(Element string) LinsertElement {
	return LinsertElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LinsertWhereAfter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LinsertWhereAfter) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type LinsertWhereBefore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LinsertWhereBefore) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type Llen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Llen) Key(Key string) LlenKey {
	return LlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Llen() (c Llen) {
	c.cs = append(b.get(), "LLEN")
	return
}

type LlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LlenKey) Build() Completed {
	return Completed(c)
}

func (c LlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Lmove struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lmove) Source(Source string) LmoveSource {
	return LmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Lmove() (c Lmove) {
	c.cs = append(b.get(), "LMOVE")
	return
}

type LmoveDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmoveDestination) Left() LmoveWherefromLeft {
	return LmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveDestination) Right() LmoveWherefromRight {
	return LmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmoveSource) Destination(Destination string) LmoveDestination {
	return LmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type LmoveWherefromLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmoveWherefromLeft) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromLeft) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveWherefromRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmoveWherefromRight) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromRight) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveWheretoLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmoveWheretoLeft) Build() Completed {
	return Completed(c)
}

type LmoveWheretoRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmoveWheretoRight) Build() Completed {
	return Completed(c)
}

type Lmpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lmpop) Numkeys(Numkeys int64) LmpopNumkeys {
	return LmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Lmpop() (c Lmpop) {
	c.cs = append(b.get(), "LMPOP")
	return
}

type LmpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmpopCount) Build() Completed {
	return Completed(c)
}

type LmpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmpopKey) Left() LmpopWhereLeft {
	return LmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmpopKey) Right() LmpopWhereRight {
	return LmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c LmpopKey) Key(Key ...string) LmpopKey {
	return LmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type LmpopNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmpopNumkeys) Key(Key ...string) LmpopKey {
	return LmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c LmpopNumkeys) Left() LmpopWhereLeft {
	return LmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmpopNumkeys) Right() LmpopWhereRight {
	return LmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmpopWhereLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmpopWhereLeft) Count(Count int64) LmpopCount {
	return LmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type LmpopWhereRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LmpopWhereRight) Count(Count int64) LmpopCount {
	return LmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Lolwut struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lolwut) Version(Version int64) LolwutVersion {
	return LolwutVersion{cf: c.cf, cs: append(c.cs, "VERSION", strconv.FormatInt(Version, 10))}
}

func (c Lolwut) Build() Completed {
	return Completed(c)
}

func (b *Builder) Lolwut() (c Lolwut) {
	c.cs = append(b.get(), "LOLWUT")
	return
}

type LolwutVersion struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LolwutVersion) Build() Completed {
	return Completed(c)
}

type Lpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lpop) Key(Key string) LpopKey {
	return LpopKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lpop() (c Lpop) {
	c.cs = append(b.get(), "LPOP")
	return
}

type LpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LpopCount) Build() Completed {
	return Completed(c)
}

type LpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LpopKey) Count(Count int64) LpopCount {
	return LpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c LpopKey) Build() Completed {
	return Completed(c)
}

type Lpos struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lpos) Key(Key string) LposKey {
	return LposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lpos() (c Lpos) {
	c.cs = append(b.get(), "LPOS")
	return
}

type LposCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LposCount) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c LposCount) Build() Completed {
	return Completed(c)
}

func (c LposCount) Cache() Cacheable {
	return Cacheable(c)
}

type LposElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LposElement) Rank(Rank int64) LposRank {
	return LposRank{cf: c.cf, cs: append(c.cs, "RANK", strconv.FormatInt(Rank, 10))}
}

func (c LposElement) Count(NumMatches int64) LposCount {
	return LposCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10))}
}

func (c LposElement) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c LposElement) Build() Completed {
	return Completed(c)
}

func (c LposElement) Cache() Cacheable {
	return Cacheable(c)
}

type LposKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LposKey) Element(Element string) LposElement {
	return LposElement{cf: c.cf, cs: append(c.cs, Element)}
}

func (c LposKey) Cache() Cacheable {
	return Cacheable(c)
}

type LposMaxlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LposMaxlen) Build() Completed {
	return Completed(c)
}

func (c LposMaxlen) Cache() Cacheable {
	return Cacheable(c)
}

type LposRank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LposRank) Count(NumMatches int64) LposCount {
	return LposCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10))}
}

func (c LposRank) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c LposRank) Build() Completed {
	return Completed(c)
}

func (c LposRank) Cache() Cacheable {
	return Cacheable(c)
}

type Lpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lpush) Key(Key string) LpushKey {
	return LpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lpush() (c Lpush) {
	c.cs = append(b.get(), "LPUSH")
	return
}

type LpushElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LpushElement) Element(Element ...string) LpushElement {
	return LpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c LpushElement) Build() Completed {
	return Completed(c)
}

type LpushKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LpushKey) Element(Element ...string) LpushElement {
	return LpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Lpushx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lpushx) Key(Key string) LpushxKey {
	return LpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lpushx() (c Lpushx) {
	c.cs = append(b.get(), "LPUSHX")
	return
}

type LpushxElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LpushxElement) Element(Element ...string) LpushxElement {
	return LpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c LpushxElement) Build() Completed {
	return Completed(c)
}

type LpushxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LpushxKey) Element(Element ...string) LpushxElement {
	return LpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Lrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lrange) Key(Key string) LrangeKey {
	return LrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lrange() (c Lrange) {
	c.cs = append(b.get(), "LRANGE")
	return
}

type LrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LrangeKey) Start(Start int64) LrangeStart {
	return LrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c LrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type LrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LrangeStart) Stop(Stop int64) LrangeStop {
	return LrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

func (c LrangeStart) Cache() Cacheable {
	return Cacheable(c)
}

type LrangeStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LrangeStop) Build() Completed {
	return Completed(c)
}

func (c LrangeStop) Cache() Cacheable {
	return Cacheable(c)
}

type Lrem struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lrem) Key(Key string) LremKey {
	return LremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lrem() (c Lrem) {
	c.cs = append(b.get(), "LREM")
	return
}

type LremCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LremCount) Element(Element string) LremElement {
	return LremElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LremElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LremElement) Build() Completed {
	return Completed(c)
}

type LremKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LremKey) Count(Count int64) LremCount {
	return LremCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type Lset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Lset) Key(Key string) LsetKey {
	return LsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lset() (c Lset) {
	c.cs = append(b.get(), "LSET")
	return
}

type LsetElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LsetElement) Build() Completed {
	return Completed(c)
}

type LsetIndex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LsetIndex) Element(Element string) LsetElement {
	return LsetElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LsetKey) Index(Index int64) LsetIndex {
	return LsetIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type Ltrim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Ltrim) Key(Key string) LtrimKey {
	return LtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Ltrim() (c Ltrim) {
	c.cs = append(b.get(), "LTRIM")
	return
}

type LtrimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LtrimKey) Start(Start int64) LtrimStart {
	return LtrimStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type LtrimStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LtrimStart) Stop(Stop int64) LtrimStop {
	return LtrimStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type LtrimStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c LtrimStop) Build() Completed {
	return Completed(c)
}

type MemoryDoctor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryDoctor) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryDoctor() (c MemoryDoctor) {
	c.cs = append(b.get(), "MEMORY", "DOCTOR")
	return
}

type MemoryHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryHelp) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryHelp() (c MemoryHelp) {
	c.cs = append(b.get(), "MEMORY", "HELP")
	return
}

type MemoryMallocStats struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryMallocStats) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryMallocStats() (c MemoryMallocStats) {
	c.cs = append(b.get(), "MEMORY", "MALLOC-STATS")
	return
}

type MemoryPurge struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryPurge) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryPurge() (c MemoryPurge) {
	c.cs = append(b.get(), "MEMORY", "PURGE")
	return
}

type MemoryStats struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryStats) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryStats() (c MemoryStats) {
	c.cs = append(b.get(), "MEMORY", "STATS")
	return
}

type MemoryUsage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryUsage) Key(Key string) MemoryUsageKey {
	return MemoryUsageKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) MemoryUsage() (c MemoryUsage) {
	c.cs = append(b.get(), "MEMORY", "USAGE")
	return
}

type MemoryUsageKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryUsageKey) Samples(Count int64) MemoryUsageSamples {
	return MemoryUsageSamples{cf: c.cf, cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10))}
}

func (c MemoryUsageKey) Build() Completed {
	return Completed(c)
}

type MemoryUsageSamples struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MemoryUsageSamples) Build() Completed {
	return Completed(c)
}

type Mget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Mget) Key(Key ...string) MgetKey {
	return MgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Mget() (c Mget) {
	c.cs = append(b.get(), "MGET")
	return
}

type MgetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MgetKey) Key(Key ...string) MgetKey {
	return MgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MgetKey) Build() Completed {
	return Completed(c)
}

type Migrate struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Migrate) Host(Host string) MigrateHost {
	return MigrateHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *Builder) Migrate() (c Migrate) {
	c.cs = append(b.get(), "MIGRATE")
	c.cf = blockTag
	return
}

type MigrateAuth struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateAuth) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateAuth) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MigrateAuth) Build() Completed {
	return Completed(c)
}

type MigrateAuth2 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateAuth2) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MigrateAuth2) Build() Completed {
	return Completed(c)
}

type MigrateCopyCopy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateCopyCopy) Replace() MigrateReplaceReplace {
	return MigrateReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c MigrateCopyCopy) Auth(Password string) MigrateAuth {
	return MigrateAuth{cf: c.cf, cs: append(c.cs, "AUTH", Password)}
}

func (c MigrateCopyCopy) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateCopyCopy) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MigrateCopyCopy) Build() Completed {
	return Completed(c)
}

type MigrateDestinationDb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateDestinationDb) Timeout(Timeout int64) MigrateTimeout {
	return MigrateTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type MigrateHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateHost) Port(Port string) MigratePort {
	return MigratePort{cf: c.cf, cs: append(c.cs, Port)}
}

type MigrateKeyEmpty struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateKeyEmpty) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateKeyKey) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateKeys) Keys(Keys ...string) MigrateKeys {
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Keys...)}
}

func (c MigrateKeys) Build() Completed {
	return Completed(c)
}

type MigratePort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigratePort) Key() MigrateKeyKey {
	return MigrateKeyKey{cf: c.cf, cs: append(c.cs, "key")}
}

func (c MigratePort) Empty() MigrateKeyEmpty {
	return MigrateKeyEmpty{cf: c.cf, cs: append(c.cs, "\"\"")}
}

type MigrateReplaceReplace struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateReplaceReplace) Auth(Password string) MigrateAuth {
	return MigrateAuth{cf: c.cf, cs: append(c.cs, "AUTH", Password)}
}

func (c MigrateReplaceReplace) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateReplaceReplace) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MigrateReplaceReplace) Build() Completed {
	return Completed(c)
}

type MigrateTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MigrateTimeout) Copy() MigrateCopyCopy {
	return MigrateCopyCopy{cf: c.cf, cs: append(c.cs, "COPY")}
}

func (c MigrateTimeout) Replace() MigrateReplaceReplace {
	return MigrateReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c MigrateTimeout) Auth(Password string) MigrateAuth {
	return MigrateAuth{cf: c.cf, cs: append(c.cs, "AUTH", Password)}
}

func (c MigrateTimeout) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateTimeout) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MigrateTimeout) Build() Completed {
	return Completed(c)
}

type ModuleList struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ModuleList) Build() Completed {
	return Completed(c)
}

func (b *Builder) ModuleList() (c ModuleList) {
	c.cs = append(b.get(), "MODULE", "LIST")
	return
}

type ModuleLoad struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ModuleLoad) Path(Path string) ModuleLoadPath {
	return ModuleLoadPath{cf: c.cf, cs: append(c.cs, Path)}
}

func (b *Builder) ModuleLoad() (c ModuleLoad) {
	c.cs = append(b.get(), "MODULE", "LOAD")
	return
}

type ModuleLoadArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ModuleLoadArg) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c ModuleLoadArg) Build() Completed {
	return Completed(c)
}

type ModuleLoadPath struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ModuleLoadPath) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c ModuleLoadPath) Build() Completed {
	return Completed(c)
}

type ModuleUnload struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ModuleUnload) Name(Name string) ModuleUnloadName {
	return ModuleUnloadName{cf: c.cf, cs: append(c.cs, Name)}
}

func (b *Builder) ModuleUnload() (c ModuleUnload) {
	c.cs = append(b.get(), "MODULE", "UNLOAD")
	return
}

type ModuleUnloadName struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ModuleUnloadName) Build() Completed {
	return Completed(c)
}

type Monitor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Monitor) Build() Completed {
	return Completed(c)
}

func (b *Builder) Monitor() (c Monitor) {
	c.cs = append(b.get(), "MONITOR")
	return
}

type Move struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Move) Key(Key string) MoveKey {
	return MoveKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Move() (c Move) {
	c.cs = append(b.get(), "MOVE")
	return
}

type MoveDb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MoveDb) Build() Completed {
	return Completed(c)
}

type MoveKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MoveKey) Db(Db int64) MoveDb {
	return MoveDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Db, 10))}
}

type Mset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Mset) KeyValue() MsetKeyValue {
	return MsetKeyValue{cf: c.cf, cs: append(c.cs, )}
}

func (b *Builder) Mset() (c Mset) {
	c.cs = append(b.get(), "MSET")
	return
}

type MsetKeyValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MsetKeyValue) KeyValue(Key string, Value string) MsetKeyValue {
	return MsetKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c MsetKeyValue) Build() Completed {
	return Completed(c)
}

type Msetnx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Msetnx) KeyValue() MsetnxKeyValue {
	return MsetnxKeyValue{cf: c.cf, cs: append(c.cs, )}
}

func (b *Builder) Msetnx() (c Msetnx) {
	c.cs = append(b.get(), "MSETNX")
	return
}

type MsetnxKeyValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c MsetnxKeyValue) KeyValue(Key string, Value string) MsetnxKeyValue {
	return MsetnxKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c MsetnxKeyValue) Build() Completed {
	return Completed(c)
}

type Multi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Multi) Build() Completed {
	return Completed(c)
}

func (b *Builder) Multi() (c Multi) {
	c.cs = append(b.get(), "MULTI")
	return
}

type Object struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Object) Subcommand(Subcommand string) ObjectSubcommand {
	return ObjectSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *Builder) Object() (c Object) {
	c.cs = append(b.get(), "OBJECT")
	return
}

type ObjectArguments struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ObjectArguments) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c ObjectArguments) Build() Completed {
	return Completed(c)
}

type ObjectSubcommand struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ObjectSubcommand) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c ObjectSubcommand) Build() Completed {
	return Completed(c)
}

type Persist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Persist) Key(Key string) PersistKey {
	return PersistKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Persist() (c Persist) {
	c.cs = append(b.get(), "PERSIST")
	return
}

type PersistKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PersistKey) Build() Completed {
	return Completed(c)
}

type Pexpire struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pexpire) Key(Key string) PexpireKey {
	return PexpireKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pexpire() (c Pexpire) {
	c.cs = append(b.get(), "PEXPIRE")
	return
}

type PexpireConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireConditionGt) Build() Completed {
	return Completed(c)
}

type PexpireConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireConditionLt) Build() Completed {
	return Completed(c)
}

type PexpireConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireConditionNx) Build() Completed {
	return Completed(c)
}

type PexpireConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireConditionXx) Build() Completed {
	return Completed(c)
}

type PexpireKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireKey) Milliseconds(Milliseconds int64) PexpireMilliseconds {
	return PexpireMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PexpireMilliseconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireMilliseconds) Nx() PexpireConditionNx {
	return PexpireConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c PexpireMilliseconds) Xx() PexpireConditionXx {
	return PexpireConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c PexpireMilliseconds) Gt() PexpireConditionGt {
	return PexpireConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c PexpireMilliseconds) Lt() PexpireConditionLt {
	return PexpireConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c PexpireMilliseconds) Build() Completed {
	return Completed(c)
}

type Pexpireat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pexpireat) Key(Key string) PexpireatKey {
	return PexpireatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pexpireat() (c Pexpireat) {
	c.cs = append(b.get(), "PEXPIREAT")
	return
}

type PexpireatConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireatConditionGt) Build() Completed {
	return Completed(c)
}

type PexpireatConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireatConditionLt) Build() Completed {
	return Completed(c)
}

type PexpireatConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireatConditionNx) Build() Completed {
	return Completed(c)
}

type PexpireatConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireatConditionXx) Build() Completed {
	return Completed(c)
}

type PexpireatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireatKey) MillisecondsTimestamp(MillisecondsTimestamp int64) PexpireatMillisecondsTimestamp {
	return PexpireatMillisecondsTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10))}
}

type PexpireatMillisecondsTimestamp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpireatMillisecondsTimestamp) Nx() PexpireatConditionNx {
	return PexpireatConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c PexpireatMillisecondsTimestamp) Xx() PexpireatConditionXx {
	return PexpireatConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c PexpireatMillisecondsTimestamp) Gt() PexpireatConditionGt {
	return PexpireatConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c PexpireatMillisecondsTimestamp) Lt() PexpireatConditionLt {
	return PexpireatConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c PexpireatMillisecondsTimestamp) Build() Completed {
	return Completed(c)
}

type Pexpiretime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pexpiretime) Key(Key string) PexpiretimeKey {
	return PexpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pexpiretime() (c Pexpiretime) {
	c.cs = append(b.get(), "PEXPIRETIME")
	return
}

type PexpiretimeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PexpiretimeKey) Build() Completed {
	return Completed(c)
}

func (c PexpiretimeKey) Cache() Cacheable {
	return Cacheable(c)
}

type Pfadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pfadd) Key(Key string) PfaddKey {
	return PfaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pfadd() (c Pfadd) {
	c.cs = append(b.get(), "PFADD")
	return
}

type PfaddElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PfaddElement) Element(Element ...string) PfaddElement {
	return PfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c PfaddElement) Build() Completed {
	return Completed(c)
}

type PfaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PfaddKey) Element(Element ...string) PfaddElement {
	return PfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c PfaddKey) Build() Completed {
	return Completed(c)
}

type Pfcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pfcount) Key(Key ...string) PfcountKey {
	return PfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Pfcount() (c Pfcount) {
	c.cs = append(b.get(), "PFCOUNT")
	return
}

type PfcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PfcountKey) Key(Key ...string) PfcountKey {
	return PfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c PfcountKey) Build() Completed {
	return Completed(c)
}

type Pfmerge struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pfmerge) Destkey(Destkey string) PfmergeDestkey {
	return PfmergeDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

func (b *Builder) Pfmerge() (c Pfmerge) {
	c.cs = append(b.get(), "PFMERGE")
	return
}

type PfmergeDestkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PfmergeDestkey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

type PfmergeSourcekey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PfmergeSourcekey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

func (c PfmergeSourcekey) Build() Completed {
	return Completed(c)
}

type Ping struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Ping) Message(Message string) PingMessage {
	return PingMessage{cf: c.cf, cs: append(c.cs, Message)}
}

func (c Ping) Build() Completed {
	return Completed(c)
}

func (b *Builder) Ping() (c Ping) {
	c.cs = append(b.get(), "PING")
	return
}

type PingMessage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PingMessage) Build() Completed {
	return Completed(c)
}

type Psetex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Psetex) Key(Key string) PsetexKey {
	return PsetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Psetex() (c Psetex) {
	c.cs = append(b.get(), "PSETEX")
	return
}

type PsetexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PsetexKey) Milliseconds(Milliseconds int64) PsetexMilliseconds {
	return PsetexMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PsetexMilliseconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PsetexMilliseconds) Value(Value string) PsetexValue {
	return PsetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type PsetexValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PsetexValue) Build() Completed {
	return Completed(c)
}

type Psubscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Psubscribe) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (b *Builder) Psubscribe() (c Psubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	c.cf = noRetTag
	return
}

type PsubscribePattern struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PsubscribePattern) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c PsubscribePattern) Build() Completed {
	return Completed(c)
}

type Psync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Psync) Replicationid(Replicationid int64) PsyncReplicationid {
	return PsyncReplicationid{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Replicationid, 10))}
}

func (b *Builder) Psync() (c Psync) {
	c.cs = append(b.get(), "PSYNC")
	return
}

type PsyncOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PsyncOffset) Build() Completed {
	return Completed(c)
}

type PsyncReplicationid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PsyncReplicationid) Offset(Offset int64) PsyncOffset {
	return PsyncOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type Pttl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pttl) Key(Key string) PttlKey {
	return PttlKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pttl() (c Pttl) {
	c.cs = append(b.get(), "PTTL")
	return
}

type PttlKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PttlKey) Build() Completed {
	return Completed(c)
}

func (c PttlKey) Cache() Cacheable {
	return Cacheable(c)
}

type Publish struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Publish) Channel(Channel string) PublishChannel {
	return PublishChannel{cf: c.cf, cs: append(c.cs, Channel)}
}

func (b *Builder) Publish() (c Publish) {
	c.cs = append(b.get(), "PUBLISH")
	return
}

type PublishChannel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PublishChannel) Message(Message string) PublishMessage {
	return PublishMessage{cf: c.cf, cs: append(c.cs, Message)}
}

type PublishMessage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PublishMessage) Build() Completed {
	return Completed(c)
}

type Pubsub struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Pubsub) Subcommand(Subcommand string) PubsubSubcommand {
	return PubsubSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *Builder) Pubsub() (c Pubsub) {
	c.cs = append(b.get(), "PUBSUB")
	return
}

type PubsubArgument struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PubsubArgument) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c PubsubArgument) Build() Completed {
	return Completed(c)
}

type PubsubSubcommand struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PubsubSubcommand) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c PubsubSubcommand) Build() Completed {
	return Completed(c)
}

type Punsubscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Punsubscribe) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c Punsubscribe) Build() Completed {
	return Completed(c)
}

func (b *Builder) Punsubscribe() (c Punsubscribe) {
	c.cs = append(b.get(), "PUNSUBSCRIBE")
	c.cf = noRetTag
	return
}

type PunsubscribePattern struct {
	cs []string
	cf uint16
	ks uint16
}

func (c PunsubscribePattern) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c PunsubscribePattern) Build() Completed {
	return Completed(c)
}

type Quit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Quit) Build() Completed {
	return Completed(c)
}

func (b *Builder) Quit() (c Quit) {
	c.cs = append(b.get(), "QUIT")
	return
}

type Randomkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Randomkey) Build() Completed {
	return Completed(c)
}

func (b *Builder) Randomkey() (c Randomkey) {
	c.cs = append(b.get(), "RANDOMKEY")
	return
}

type Readonly struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Readonly) Build() Completed {
	return Completed(c)
}

func (b *Builder) Readonly() (c Readonly) {
	c.cs = append(b.get(), "READONLY")
	return
}

type Readwrite struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Readwrite) Build() Completed {
	return Completed(c)
}

func (b *Builder) Readwrite() (c Readwrite) {
	c.cs = append(b.get(), "READWRITE")
	return
}

type Rename struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Rename) Key(Key string) RenameKey {
	return RenameKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Rename() (c Rename) {
	c.cs = append(b.get(), "RENAME")
	return
}

type RenameKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RenameKey) Newkey(Newkey string) RenameNewkey {
	return RenameNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type RenameNewkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RenameNewkey) Build() Completed {
	return Completed(c)
}

type Renamenx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Renamenx) Key(Key string) RenamenxKey {
	return RenamenxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Renamenx() (c Renamenx) {
	c.cs = append(b.get(), "RENAMENX")
	return
}

type RenamenxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RenamenxKey) Newkey(Newkey string) RenamenxNewkey {
	return RenamenxNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type RenamenxNewkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RenamenxNewkey) Build() Completed {
	return Completed(c)
}

type Replicaof struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Replicaof) Host(Host string) ReplicaofHost {
	return ReplicaofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *Builder) Replicaof() (c Replicaof) {
	c.cs = append(b.get(), "REPLICAOF")
	return
}

type ReplicaofHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ReplicaofHost) Port(Port string) ReplicaofPort {
	return ReplicaofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type ReplicaofPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ReplicaofPort) Build() Completed {
	return Completed(c)
}

type Reset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Reset) Build() Completed {
	return Completed(c)
}

func (b *Builder) Reset() (c Reset) {
	c.cs = append(b.get(), "RESET")
	return
}

type Restore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Restore) Key(Key string) RestoreKey {
	return RestoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Restore() (c Restore) {
	c.cs = append(b.get(), "RESTORE")
	return
}

type RestoreAbsttlAbsttl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RestoreAbsttlAbsttl) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c RestoreAbsttlAbsttl) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreAbsttlAbsttl) Build() Completed {
	return Completed(c)
}

type RestoreFreq struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RestoreFreq) Build() Completed {
	return Completed(c)
}

type RestoreIdletime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RestoreIdletime) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreIdletime) Build() Completed {
	return Completed(c)
}

type RestoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RestoreKey) Ttl(Ttl int64) RestoreTtl {
	return RestoreTtl{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Ttl, 10))}
}

type RestoreReplaceReplace struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RestoreReplaceReplace) Absttl() RestoreAbsttlAbsttl {
	return RestoreAbsttlAbsttl{cf: c.cf, cs: append(c.cs, "ABSTTL")}
}

func (c RestoreReplaceReplace) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c RestoreReplaceReplace) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreReplaceReplace) Build() Completed {
	return Completed(c)
}

type RestoreSerializedValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RestoreSerializedValue) Replace() RestoreReplaceReplace {
	return RestoreReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c RestoreSerializedValue) Absttl() RestoreAbsttlAbsttl {
	return RestoreAbsttlAbsttl{cf: c.cf, cs: append(c.cs, "ABSTTL")}
}

func (c RestoreSerializedValue) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c RestoreSerializedValue) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreSerializedValue) Build() Completed {
	return Completed(c)
}

type RestoreTtl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RestoreTtl) SerializedValue(SerializedValue string) RestoreSerializedValue {
	return RestoreSerializedValue{cf: c.cf, cs: append(c.cs, SerializedValue)}
}

type Role struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Role) Build() Completed {
	return Completed(c)
}

func (b *Builder) Role() (c Role) {
	c.cs = append(b.get(), "ROLE")
	return
}

type Rpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Rpop) Key(Key string) RpopKey {
	return RpopKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Rpop() (c Rpop) {
	c.cs = append(b.get(), "RPOP")
	return
}

type RpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpopCount) Build() Completed {
	return Completed(c)
}

type RpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpopKey) Count(Count int64) RpopCount {
	return RpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c RpopKey) Build() Completed {
	return Completed(c)
}

type Rpoplpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Rpoplpush) Source(Source string) RpoplpushSource {
	return RpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Rpoplpush() (c Rpoplpush) {
	c.cs = append(b.get(), "RPOPLPUSH")
	return
}

type RpoplpushDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpoplpushDestination) Build() Completed {
	return Completed(c)
}

type RpoplpushSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpoplpushSource) Destination(Destination string) RpoplpushDestination {
	return RpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Rpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Rpush) Key(Key string) RpushKey {
	return RpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Rpush() (c Rpush) {
	c.cs = append(b.get(), "RPUSH")
	return
}

type RpushElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpushElement) Element(Element ...string) RpushElement {
	return RpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c RpushElement) Build() Completed {
	return Completed(c)
}

type RpushKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpushKey) Element(Element ...string) RpushElement {
	return RpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Rpushx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Rpushx) Key(Key string) RpushxKey {
	return RpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Rpushx() (c Rpushx) {
	c.cs = append(b.get(), "RPUSHX")
	return
}

type RpushxElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpushxElement) Element(Element ...string) RpushxElement {
	return RpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c RpushxElement) Build() Completed {
	return Completed(c)
}

type RpushxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c RpushxKey) Element(Element ...string) RpushxElement {
	return RpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Sadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sadd) Key(Key string) SaddKey {
	return SaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sadd() (c Sadd) {
	c.cs = append(b.get(), "SADD")
	return
}

type SaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SaddKey) Member(Member ...string) SaddMember {
	return SaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SaddMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SaddMember) Member(Member ...string) SaddMember {
	return SaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SaddMember) Build() Completed {
	return Completed(c)
}

type Save struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Save) Build() Completed {
	return Completed(c)
}

func (b *Builder) Save() (c Save) {
	c.cs = append(b.get(), "SAVE")
	return
}

type Scan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Scan) Cursor(Cursor int64) ScanCursor {
	return ScanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

func (b *Builder) Scan() (c Scan) {
	c.cs = append(b.get(), "SCAN")
	return
}

type ScanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScanCount) Type(Type string) ScanType {
	return ScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c ScanCount) Build() Completed {
	return Completed(c)
}

type ScanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScanCursor) Match(Pattern string) ScanMatch {
	return ScanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c ScanCursor) Count(Count int64) ScanCount {
	return ScanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ScanCursor) Type(Type string) ScanType {
	return ScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c ScanCursor) Build() Completed {
	return Completed(c)
}

type ScanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScanMatch) Count(Count int64) ScanCount {
	return ScanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ScanMatch) Type(Type string) ScanType {
	return ScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c ScanMatch) Build() Completed {
	return Completed(c)
}

type ScanType struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScanType) Build() Completed {
	return Completed(c)
}

type Scard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Scard) Key(Key string) ScardKey {
	return ScardKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Scard() (c Scard) {
	c.cs = append(b.get(), "SCARD")
	return
}

type ScardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScardKey) Build() Completed {
	return Completed(c)
}

func (c ScardKey) Cache() Cacheable {
	return Cacheable(c)
}

type ScriptDebug struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptDebug) Yes() ScriptDebugModeYes {
	return ScriptDebugModeYes{cf: c.cf, cs: append(c.cs, "YES")}
}

func (c ScriptDebug) Sync() ScriptDebugModeSync {
	return ScriptDebugModeSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c ScriptDebug) No() ScriptDebugModeNo {
	return ScriptDebugModeNo{cf: c.cf, cs: append(c.cs, "NO")}
}

func (b *Builder) ScriptDebug() (c ScriptDebug) {
	c.cs = append(b.get(), "SCRIPT", "DEBUG")
	return
}

type ScriptDebugModeNo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptDebugModeNo) Build() Completed {
	return Completed(c)
}

type ScriptDebugModeSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptDebugModeSync) Build() Completed {
	return Completed(c)
}

type ScriptDebugModeYes struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptDebugModeYes) Build() Completed {
	return Completed(c)
}

type ScriptExists struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptExists) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (b *Builder) ScriptExists() (c ScriptExists) {
	c.cs = append(b.get(), "SCRIPT", "EXISTS")
	return
}

type ScriptExistsSha1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptExistsSha1) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (c ScriptExistsSha1) Build() Completed {
	return Completed(c)
}

type ScriptFlush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptFlush) Async() ScriptFlushAsyncAsync {
	return ScriptFlushAsyncAsync{cf: c.cf, cs: append(c.cs, "ASYNC")}
}

func (c ScriptFlush) Sync() ScriptFlushAsyncSync {
	return ScriptFlushAsyncSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c ScriptFlush) Build() Completed {
	return Completed(c)
}

func (b *Builder) ScriptFlush() (c ScriptFlush) {
	c.cs = append(b.get(), "SCRIPT", "FLUSH")
	return
}

type ScriptFlushAsyncAsync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptFlushAsyncAsync) Build() Completed {
	return Completed(c)
}

type ScriptFlushAsyncSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptFlushAsyncSync) Build() Completed {
	return Completed(c)
}

type ScriptKill struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptKill) Build() Completed {
	return Completed(c)
}

func (b *Builder) ScriptKill() (c ScriptKill) {
	c.cs = append(b.get(), "SCRIPT", "KILL")
	return
}

type ScriptLoad struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptLoad) Script(Script string) ScriptLoadScript {
	return ScriptLoadScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *Builder) ScriptLoad() (c ScriptLoad) {
	c.cs = append(b.get(), "SCRIPT", "LOAD")
	return
}

type ScriptLoadScript struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ScriptLoadScript) Build() Completed {
	return Completed(c)
}

type Sdiff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sdiff) Key(Key ...string) SdiffKey {
	return SdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sdiff() (c Sdiff) {
	c.cs = append(b.get(), "SDIFF")
	return
}

type SdiffKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SdiffKey) Key(Key ...string) SdiffKey {
	return SdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SdiffKey) Build() Completed {
	return Completed(c)
}

type Sdiffstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sdiffstore) Destination(Destination string) SdiffstoreDestination {
	return SdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Sdiffstore() (c Sdiffstore) {
	c.cs = append(b.get(), "SDIFFSTORE")
	return
}

type SdiffstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SdiffstoreDestination) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SdiffstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SdiffstoreKey) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SdiffstoreKey) Build() Completed {
	return Completed(c)
}

type Select struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Select) Index(Index int64) SelectIndex {
	return SelectIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

func (b *Builder) Select() (c Select) {
	c.cs = append(b.get(), "SELECT")
	return
}

type SelectIndex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SelectIndex) Build() Completed {
	return Completed(c)
}

type Set struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Set) Key(Key string) SetKey {
	return SetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Set() (c Set) {
	c.cs = append(b.get(), "SET")
	return
}

type SetConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetConditionNx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetConditionNx) Build() Completed {
	return Completed(c)
}

type SetConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetConditionXx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetConditionXx) Build() Completed {
	return Completed(c)
}

type SetExpirationEx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetExpirationEx) Nx() SetConditionNx {
	return SetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SetExpirationEx) Xx() SetConditionXx {
	return SetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SetExpirationEx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetExpirationEx) Build() Completed {
	return Completed(c)
}

type SetExpirationExat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetExpirationExat) Nx() SetConditionNx {
	return SetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SetExpirationExat) Xx() SetConditionXx {
	return SetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SetExpirationExat) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetExpirationExat) Build() Completed {
	return Completed(c)
}

type SetExpirationKeepttl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetExpirationKeepttl) Nx() SetConditionNx {
	return SetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SetExpirationKeepttl) Xx() SetConditionXx {
	return SetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SetExpirationKeepttl) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetExpirationKeepttl) Build() Completed {
	return Completed(c)
}

type SetExpirationPx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetExpirationPx) Nx() SetConditionNx {
	return SetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SetExpirationPx) Xx() SetConditionXx {
	return SetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SetExpirationPx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetExpirationPx) Build() Completed {
	return Completed(c)
}

type SetExpirationPxat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetExpirationPxat) Nx() SetConditionNx {
	return SetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SetExpirationPxat) Xx() SetConditionXx {
	return SetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SetExpirationPxat) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetExpirationPxat) Build() Completed {
	return Completed(c)
}

type SetGetGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetGetGet) Build() Completed {
	return Completed(c)
}

type SetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetKey) Value(Value string) SetValue {
	return SetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetValue) Ex(Seconds int64) SetExpirationEx {
	return SetExpirationEx{cf: c.cf, cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10))}
}

func (c SetValue) Px(Milliseconds int64) SetExpirationPx {
	return SetExpirationPx{cf: c.cf, cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10))}
}

func (c SetValue) Exat(Timestamp int64) SetExpirationExat {
	return SetExpirationExat{cf: c.cf, cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10))}
}

func (c SetValue) Pxat(Millisecondstimestamp int64) SetExpirationPxat {
	return SetExpirationPxat{cf: c.cf, cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10))}
}

func (c SetValue) Keepttl() SetExpirationKeepttl {
	return SetExpirationKeepttl{cf: c.cf, cs: append(c.cs, "KEEPTTL")}
}

func (c SetValue) Nx() SetConditionNx {
	return SetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SetValue) Xx() SetConditionXx {
	return SetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SetValue) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetValue) Build() Completed {
	return Completed(c)
}

type Setbit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Setbit) Key(Key string) SetbitKey {
	return SetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setbit() (c Setbit) {
	c.cs = append(b.get(), "SETBIT")
	return
}

type SetbitKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetbitKey) Offset(Offset int64) SetbitOffset {
	return SetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetbitOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetbitOffset) Value(Value int64) SetbitValue {
	return SetbitValue{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Value, 10))}
}

type SetbitValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetbitValue) Build() Completed {
	return Completed(c)
}

type Setex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Setex) Key(Key string) SetexKey {
	return SetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setex() (c Setex) {
	c.cs = append(b.get(), "SETEX")
	return
}

type SetexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetexKey) Seconds(Seconds int64) SetexSeconds {
	return SetexSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SetexSeconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetexSeconds) Value(Value string) SetexValue {
	return SetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetexValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetexValue) Build() Completed {
	return Completed(c)
}

type Setnx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Setnx) Key(Key string) SetnxKey {
	return SetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setnx() (c Setnx) {
	c.cs = append(b.get(), "SETNX")
	return
}

type SetnxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetnxKey) Value(Value string) SetnxValue {
	return SetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetnxValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetnxValue) Build() Completed {
	return Completed(c)
}

type Setrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Setrange) Key(Key string) SetrangeKey {
	return SetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setrange() (c Setrange) {
	c.cs = append(b.get(), "SETRANGE")
	return
}

type SetrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetrangeKey) Offset(Offset int64) SetrangeOffset {
	return SetrangeOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetrangeOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetrangeOffset) Value(Value string) SetrangeValue {
	return SetrangeValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetrangeValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SetrangeValue) Build() Completed {
	return Completed(c)
}

type Shutdown struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Shutdown) Nosave() ShutdownSaveModeNosave {
	return ShutdownSaveModeNosave{cf: c.cf, cs: append(c.cs, "NOSAVE")}
}

func (c Shutdown) Save() ShutdownSaveModeSave {
	return ShutdownSaveModeSave{cf: c.cf, cs: append(c.cs, "SAVE")}
}

func (c Shutdown) Build() Completed {
	return Completed(c)
}

func (b *Builder) Shutdown() (c Shutdown) {
	c.cs = append(b.get(), "SHUTDOWN")
	return
}

type ShutdownSaveModeNosave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ShutdownSaveModeNosave) Build() Completed {
	return Completed(c)
}

type ShutdownSaveModeSave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ShutdownSaveModeSave) Build() Completed {
	return Completed(c)
}

type Sinter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sinter) Key(Key ...string) SinterKey {
	return SinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sinter() (c Sinter) {
	c.cs = append(b.get(), "SINTER")
	return
}

type SinterKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SinterKey) Key(Key ...string) SinterKey {
	return SinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SinterKey) Build() Completed {
	return Completed(c)
}

type Sintercard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sintercard) Key(Key ...string) SintercardKey {
	return SintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sintercard() (c Sintercard) {
	c.cs = append(b.get(), "SINTERCARD")
	return
}

type SintercardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SintercardKey) Key(Key ...string) SintercardKey {
	return SintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SintercardKey) Build() Completed {
	return Completed(c)
}

type Sinterstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sinterstore) Destination(Destination string) SinterstoreDestination {
	return SinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Sinterstore() (c Sinterstore) {
	c.cs = append(b.get(), "SINTERSTORE")
	return
}

type SinterstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SinterstoreDestination) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SinterstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SinterstoreKey) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SinterstoreKey) Build() Completed {
	return Completed(c)
}

type Sismember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sismember) Key(Key string) SismemberKey {
	return SismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sismember() (c Sismember) {
	c.cs = append(b.get(), "SISMEMBER")
	return
}

type SismemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SismemberKey) Member(Member string) SismemberMember {
	return SismemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c SismemberKey) Cache() Cacheable {
	return Cacheable(c)
}

type SismemberMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SismemberMember) Build() Completed {
	return Completed(c)
}

func (c SismemberMember) Cache() Cacheable {
	return Cacheable(c)
}

type Slaveof struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Slaveof) Host(Host string) SlaveofHost {
	return SlaveofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *Builder) Slaveof() (c Slaveof) {
	c.cs = append(b.get(), "SLAVEOF")
	return
}

type SlaveofHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SlaveofHost) Port(Port string) SlaveofPort {
	return SlaveofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type SlaveofPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SlaveofPort) Build() Completed {
	return Completed(c)
}

type Slowlog struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Slowlog) Subcommand(Subcommand string) SlowlogSubcommand {
	return SlowlogSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *Builder) Slowlog() (c Slowlog) {
	c.cs = append(b.get(), "SLOWLOG")
	return
}

type SlowlogArgument struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SlowlogArgument) Build() Completed {
	return Completed(c)
}

type SlowlogSubcommand struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SlowlogSubcommand) Argument(Argument string) SlowlogArgument {
	return SlowlogArgument{cf: c.cf, cs: append(c.cs, Argument)}
}

func (c SlowlogSubcommand) Build() Completed {
	return Completed(c)
}

type Smembers struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Smembers) Key(Key string) SmembersKey {
	return SmembersKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Smembers() (c Smembers) {
	c.cs = append(b.get(), "SMEMBERS")
	return
}

type SmembersKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SmembersKey) Build() Completed {
	return Completed(c)
}

func (c SmembersKey) Cache() Cacheable {
	return Cacheable(c)
}

type Smismember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Smismember) Key(Key string) SmismemberKey {
	return SmismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Smismember() (c Smismember) {
	c.cs = append(b.get(), "SMISMEMBER")
	return
}

type SmismemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SmismemberKey) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SmismemberKey) Cache() Cacheable {
	return Cacheable(c)
}

type SmismemberMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SmismemberMember) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SmismemberMember) Build() Completed {
	return Completed(c)
}

func (c SmismemberMember) Cache() Cacheable {
	return Cacheable(c)
}

type Smove struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Smove) Source(Source string) SmoveSource {
	return SmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Smove() (c Smove) {
	c.cs = append(b.get(), "SMOVE")
	return
}

type SmoveDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SmoveDestination) Member(Member string) SmoveMember {
	return SmoveMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SmoveMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SmoveMember) Build() Completed {
	return Completed(c)
}

type SmoveSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SmoveSource) Destination(Destination string) SmoveDestination {
	return SmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Sort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sort) Key(Key string) SortKey {
	return SortKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sort() (c Sort) {
	c.cs = append(b.get(), "SORT")
	return
}

type SortBy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortBy) Limit(Offset int64, Count int64) SortLimit {
	return SortLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortBy) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SortBy) Asc() SortOrderAsc {
	return SortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortBy) Desc() SortOrderDesc {
	return SortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortBy) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortBy) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortBy) Build() Completed {
	return Completed(c)
}

type SortGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortGet) Asc() SortOrderAsc {
	return SortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortGet) Desc() SortOrderDesc {
	return SortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortGet) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortGet) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortGet) Get(Get ...string) SortGet {
	return SortGet{cf: c.cf, cs: append(c.cs, Get...)}
}

func (c SortGet) Build() Completed {
	return Completed(c)
}

type SortKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortKey) By(Pattern string) SortBy {
	return SortBy{cf: c.cf, cs: append(c.cs, "BY", Pattern)}
}

func (c SortKey) Limit(Offset int64, Count int64) SortLimit {
	return SortLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortKey) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SortKey) Asc() SortOrderAsc {
	return SortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortKey) Desc() SortOrderDesc {
	return SortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortKey) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortKey) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortKey) Build() Completed {
	return Completed(c)
}

type SortLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortLimit) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SortLimit) Asc() SortOrderAsc {
	return SortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortLimit) Desc() SortOrderDesc {
	return SortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortLimit) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortLimit) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortLimit) Build() Completed {
	return Completed(c)
}

type SortOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortOrderAsc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortOrderAsc) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortOrderAsc) Build() Completed {
	return Completed(c)
}

type SortOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortOrderDesc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortOrderDesc) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortOrderDesc) Build() Completed {
	return Completed(c)
}

type SortRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRo) Key(Key string) SortRoKey {
	return SortRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) SortRo() (c SortRo) {
	c.cs = append(b.get(), "SORT_RO")
	return
}

type SortRoBy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRoBy) Limit(Offset int64, Count int64) SortRoLimit {
	return SortRoLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortRoBy) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SortRoBy) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortRoBy) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortRoBy) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoBy) Build() Completed {
	return Completed(c)
}

func (c SortRoBy) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRoGet) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortRoGet) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortRoGet) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoGet) Get(Get ...string) SortRoGet {
	return SortRoGet{cf: c.cf, cs: append(c.cs, Get...)}
}

func (c SortRoGet) Build() Completed {
	return Completed(c)
}

func (c SortRoGet) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRoKey) By(Pattern string) SortRoBy {
	return SortRoBy{cf: c.cf, cs: append(c.cs, "BY", Pattern)}
}

func (c SortRoKey) Limit(Offset int64, Count int64) SortRoLimit {
	return SortRoLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortRoKey) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SortRoKey) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortRoKey) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortRoKey) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoKey) Build() Completed {
	return Completed(c)
}

func (c SortRoKey) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRoLimit) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SortRoLimit) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SortRoLimit) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SortRoLimit) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoLimit) Build() Completed {
	return Completed(c)
}

func (c SortRoLimit) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRoOrderAsc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c SortRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRoOrderDesc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c SortRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoSortingAlpha struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortRoSortingAlpha) Build() Completed {
	return Completed(c)
}

func (c SortRoSortingAlpha) Cache() Cacheable {
	return Cacheable(c)
}

type SortSortingAlpha struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortSortingAlpha) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortSortingAlpha) Build() Completed {
	return Completed(c)
}

type SortStore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SortStore) Build() Completed {
	return Completed(c)
}

type Spop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Spop) Key(Key string) SpopKey {
	return SpopKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Spop() (c Spop) {
	c.cs = append(b.get(), "SPOP")
	return
}

type SpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SpopCount) Build() Completed {
	return Completed(c)
}

type SpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SpopKey) Count(Count int64) SpopCount {
	return SpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SpopKey) Build() Completed {
	return Completed(c)
}

type Srandmember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Srandmember) Key(Key string) SrandmemberKey {
	return SrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Srandmember() (c Srandmember) {
	c.cs = append(b.get(), "SRANDMEMBER")
	return
}

type SrandmemberCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SrandmemberCount) Build() Completed {
	return Completed(c)
}

type SrandmemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SrandmemberKey) Count(Count int64) SrandmemberCount {
	return SrandmemberCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SrandmemberKey) Build() Completed {
	return Completed(c)
}

type Srem struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Srem) Key(Key string) SremKey {
	return SremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Srem() (c Srem) {
	c.cs = append(b.get(), "SREM")
	return
}

type SremKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SremKey) Member(Member ...string) SremMember {
	return SremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SremMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SremMember) Member(Member ...string) SremMember {
	return SremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SremMember) Build() Completed {
	return Completed(c)
}

type Sscan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sscan) Key(Key string) SscanKey {
	return SscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sscan() (c Sscan) {
	c.cs = append(b.get(), "SSCAN")
	return
}

type SscanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SscanCount) Build() Completed {
	return Completed(c)
}

type SscanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SscanCursor) Match(Pattern string) SscanMatch {
	return SscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SscanCursor) Count(Count int64) SscanCount {
	return SscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SscanCursor) Build() Completed {
	return Completed(c)
}

type SscanKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SscanKey) Cursor(Cursor int64) SscanCursor {
	return SscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SscanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SscanMatch) Count(Count int64) SscanCount {
	return SscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SscanMatch) Build() Completed {
	return Completed(c)
}

type Stralgo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Stralgo) Lcs() StralgoAlgorithmLcs {
	return StralgoAlgorithmLcs{cf: c.cf, cs: append(c.cs, "LCS")}
}

func (b *Builder) Stralgo() (c Stralgo) {
	c.cs = append(b.get(), "STRALGO")
	return
}

type StralgoAlgoSpecificArgument struct {
	cs []string
	cf uint16
	ks uint16
}

func (c StralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

func (c StralgoAlgoSpecificArgument) Build() Completed {
	return Completed(c)
}

type StralgoAlgorithmLcs struct {
	cs []string
	cf uint16
	ks uint16
}

func (c StralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

type Strlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Strlen) Key(Key string) StrlenKey {
	return StrlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Strlen() (c Strlen) {
	c.cs = append(b.get(), "STRLEN")
	return
}

type StrlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c StrlenKey) Build() Completed {
	return Completed(c)
}

func (c StrlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Subscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Subscribe) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (b *Builder) Subscribe() (c Subscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	c.cf = noRetTag
	return
}

type SubscribeChannel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SubscribeChannel) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c SubscribeChannel) Build() Completed {
	return Completed(c)
}

type Sunion struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sunion) Key(Key ...string) SunionKey {
	return SunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sunion() (c Sunion) {
	c.cs = append(b.get(), "SUNION")
	return
}

type SunionKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SunionKey) Key(Key ...string) SunionKey {
	return SunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SunionKey) Build() Completed {
	return Completed(c)
}

type Sunionstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sunionstore) Destination(Destination string) SunionstoreDestination {
	return SunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Sunionstore() (c Sunionstore) {
	c.cs = append(b.get(), "SUNIONSTORE")
	return
}

type SunionstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SunionstoreDestination) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SunionstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SunionstoreKey) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SunionstoreKey) Build() Completed {
	return Completed(c)
}

type Swapdb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Swapdb) Index1(Index1 int64) SwapdbIndex1 {
	return SwapdbIndex1{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index1, 10))}
}

func (b *Builder) Swapdb() (c Swapdb) {
	c.cs = append(b.get(), "SWAPDB")
	return
}

type SwapdbIndex1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SwapdbIndex1) Index2(Index2 int64) SwapdbIndex2 {
	return SwapdbIndex2{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index2, 10))}
}

type SwapdbIndex2 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SwapdbIndex2) Build() Completed {
	return Completed(c)
}

type Sync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Sync) Build() Completed {
	return Completed(c)
}

func (b *Builder) Sync() (c Sync) {
	c.cs = append(b.get(), "SYNC")
	return
}

type Time struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Time) Build() Completed {
	return Completed(c)
}

func (b *Builder) Time() (c Time) {
	c.cs = append(b.get(), "TIME")
	return
}

type Touch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Touch) Key(Key ...string) TouchKey {
	return TouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Touch() (c Touch) {
	c.cs = append(b.get(), "TOUCH")
	return
}

type TouchKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c TouchKey) Key(Key ...string) TouchKey {
	return TouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c TouchKey) Build() Completed {
	return Completed(c)
}

type Ttl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Ttl) Key(Key string) TtlKey {
	return TtlKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Ttl() (c Ttl) {
	c.cs = append(b.get(), "TTL")
	return
}

type TtlKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c TtlKey) Build() Completed {
	return Completed(c)
}

func (c TtlKey) Cache() Cacheable {
	return Cacheable(c)
}

type Type struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Type) Key(Key string) TypeKey {
	return TypeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Type() (c Type) {
	c.cs = append(b.get(), "TYPE")
	return
}

type TypeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c TypeKey) Build() Completed {
	return Completed(c)
}

func (c TypeKey) Cache() Cacheable {
	return Cacheable(c)
}

type Unlink struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Unlink) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Unlink() (c Unlink) {
	c.cs = append(b.get(), "UNLINK")
	return
}

type UnlinkKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c UnlinkKey) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c UnlinkKey) Build() Completed {
	return Completed(c)
}

type Unsubscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Unsubscribe) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c Unsubscribe) Build() Completed {
	return Completed(c)
}

func (b *Builder) Unsubscribe() (c Unsubscribe) {
	c.cs = append(b.get(), "UNSUBSCRIBE")
	c.cf = noRetTag
	return
}

type UnsubscribeChannel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c UnsubscribeChannel) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c UnsubscribeChannel) Build() Completed {
	return Completed(c)
}

type Unwatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Unwatch) Build() Completed {
	return Completed(c)
}

func (b *Builder) Unwatch() (c Unwatch) {
	c.cs = append(b.get(), "UNWATCH")
	return
}

type Wait struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Wait) Numreplicas(Numreplicas int64) WaitNumreplicas {
	return WaitNumreplicas{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numreplicas, 10))}
}

func (b *Builder) Wait() (c Wait) {
	c.cs = append(b.get(), "WAIT")
	c.cf = blockTag
	return
}

type WaitNumreplicas struct {
	cs []string
	cf uint16
	ks uint16
}

func (c WaitNumreplicas) Timeout(Timeout int64) WaitTimeout {
	return WaitTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type WaitTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c WaitTimeout) Build() Completed {
	return Completed(c)
}

type Watch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Watch) Key(Key ...string) WatchKey {
	return WatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Watch() (c Watch) {
	c.cs = append(b.get(), "WATCH")
	return
}

type WatchKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c WatchKey) Key(Key ...string) WatchKey {
	return WatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c WatchKey) Build() Completed {
	return Completed(c)
}

type Xack struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xack) Key(Key string) XackKey {
	return XackKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xack() (c Xack) {
	c.cs = append(b.get(), "XACK")
	return
}

type XackGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XackGroup) Id(Id ...string) XackId {
	return XackId{cf: c.cf, cs: append(c.cs, Id...)}
}

type XackId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XackId) Id(Id ...string) XackId {
	return XackId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XackId) Build() Completed {
	return Completed(c)
}

type XackKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XackKey) Group(Group string) XackGroup {
	return XackGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type Xadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xadd) Key(Key string) XaddKey {
	return XaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xadd() (c Xadd) {
	c.cs = append(b.get(), "XADD")
	return
}

type XaddFieldValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddFieldValue) FieldValue(Field string, Value string) XaddFieldValue {
	return XaddFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c XaddFieldValue) Build() Completed {
	return Completed(c)
}

type XaddId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddId) FieldValue() XaddFieldValue {
	return XaddFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type XaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddKey) Nomkstream() XaddNomkstream {
	return XaddNomkstream{cf: c.cf, cs: append(c.cs, "NOMKSTREAM")}
}

func (c XaddKey) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c XaddKey) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

func (c XaddKey) Id(Id string) XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type XaddNomkstream struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddNomkstream) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c XaddNomkstream) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

func (c XaddNomkstream) Id(Id string) XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type XaddTrimLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddTrimLimit) Id(Id string) XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type XaddTrimOperatorAlmost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddTrimOperatorAlmost) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimOperatorExact struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddTrimOperatorExact) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimStrategyMaxlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddTrimStrategyMaxlen) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XaddTrimStrategyMaxlen) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XaddTrimStrategyMaxlen) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimStrategyMinid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddTrimStrategyMinid) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XaddTrimStrategyMinid) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XaddTrimStrategyMinid) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimThreshold struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XaddTrimThreshold) Limit(Count int64) XaddTrimLimit {
	return XaddTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XaddTrimThreshold) Id(Id string) XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type Xautoclaim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xautoclaim) Key(Key string) XautoclaimKey {
	return XautoclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xautoclaim() (c Xautoclaim) {
	c.cs = append(b.get(), "XAUTOCLAIM")
	return
}

type XautoclaimConsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XautoclaimConsumer) MinIdleTime(MinIdleTime string) XautoclaimMinIdleTime {
	return XautoclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type XautoclaimCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XautoclaimCount) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XautoclaimCount) Build() Completed {
	return Completed(c)
}

type XautoclaimGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XautoclaimGroup) Consumer(Consumer string) XautoclaimConsumer {
	return XautoclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type XautoclaimJustidJustid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XautoclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XautoclaimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XautoclaimKey) Group(Group string) XautoclaimGroup {
	return XautoclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type XautoclaimMinIdleTime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XautoclaimMinIdleTime) Start(Start string) XautoclaimStart {
	return XautoclaimStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XautoclaimStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XautoclaimStart) Count(Count int64) XautoclaimCount {
	return XautoclaimCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XautoclaimStart) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XautoclaimStart) Build() Completed {
	return Completed(c)
}

type Xclaim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xclaim) Key(Key string) XclaimKey {
	return XclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xclaim() (c Xclaim) {
	c.cs = append(b.get(), "XCLAIM")
	return
}

type XclaimConsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimConsumer) MinIdleTime(MinIdleTime string) XclaimMinIdleTime {
	return XclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type XclaimForceForce struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimForceForce) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimForceForce) Build() Completed {
	return Completed(c)
}

type XclaimGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimGroup) Consumer(Consumer string) XclaimConsumer {
	return XclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type XclaimId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimId) Idle(Ms int64) XclaimIdle {
	return XclaimIdle{cf: c.cf, cs: append(c.cs, "IDLE", strconv.FormatInt(Ms, 10))}
}

func (c XclaimId) Time(MsUnixTime int64) XclaimTime {
	return XclaimTime{cf: c.cf, cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10))}
}

func (c XclaimId) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cf: c.cf, cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c XclaimId) Force() XclaimForceForce {
	return XclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c XclaimId) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimId) Id(Id ...string) XclaimId {
	return XclaimId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XclaimId) Build() Completed {
	return Completed(c)
}

type XclaimIdle struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimIdle) Time(MsUnixTime int64) XclaimTime {
	return XclaimTime{cf: c.cf, cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10))}
}

func (c XclaimIdle) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cf: c.cf, cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c XclaimIdle) Force() XclaimForceForce {
	return XclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c XclaimIdle) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimIdle) Build() Completed {
	return Completed(c)
}

type XclaimJustidJustid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XclaimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimKey) Group(Group string) XclaimGroup {
	return XclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type XclaimMinIdleTime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimMinIdleTime) Id(Id ...string) XclaimId {
	return XclaimId{cf: c.cf, cs: append(c.cs, Id...)}
}

type XclaimRetrycount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimRetrycount) Force() XclaimForceForce {
	return XclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c XclaimRetrycount) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimRetrycount) Build() Completed {
	return Completed(c)
}

type XclaimTime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XclaimTime) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cf: c.cf, cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c XclaimTime) Force() XclaimForceForce {
	return XclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c XclaimTime) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimTime) Build() Completed {
	return Completed(c)
}

type Xdel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xdel) Key(Key string) XdelKey {
	return XdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xdel() (c Xdel) {
	c.cs = append(b.get(), "XDEL")
	return
}

type XdelId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XdelId) Id(Id ...string) XdelId {
	return XdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XdelId) Build() Completed {
	return Completed(c)
}

type XdelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XdelKey) Id(Id ...string) XdelId {
	return XdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

type Xgroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xgroup) Create(Key string, Groupname string) XgroupCreateCreate {
	return XgroupCreateCreate{cf: c.cf, cs: append(c.cs, "CREATE", Key, Groupname)}
}

func (c Xgroup) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c Xgroup) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c Xgroup) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c Xgroup) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (b *Builder) Xgroup() (c Xgroup) {
	c.cs = append(b.get(), "XGROUP")
	return
}

type XgroupCreateCreate struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupCreateCreate) Id(Id string) XgroupCreateId {
	return XgroupCreateId{cf: c.cf, cs: append(c.cs, Id)}
}

type XgroupCreateId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupCreateId) Mkstream() XgroupCreateMkstream {
	return XgroupCreateMkstream{cf: c.cf, cs: append(c.cs, "MKSTREAM")}
}

func (c XgroupCreateId) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c XgroupCreateId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupCreateId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type XgroupCreateMkstream struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupCreateMkstream) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c XgroupCreateMkstream) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupCreateMkstream) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateMkstream) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type XgroupCreateconsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateconsumer) Build() Completed {
	return Completed(c)
}

type XgroupDelconsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupDelconsumer) Build() Completed {
	return Completed(c)
}

type XgroupDestroy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupDestroy) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupDestroy) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupDestroy) Build() Completed {
	return Completed(c)
}

type XgroupSetidId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupSetidId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupSetidId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidId) Build() Completed {
	return Completed(c)
}

type XgroupSetidSetid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XgroupSetidSetid) Id(Id string) XgroupSetidId {
	return XgroupSetidId{cf: c.cf, cs: append(c.cs, Id)}
}

type Xinfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xinfo) Consumers(Key string, Groupname string) XinfoConsumers {
	return XinfoConsumers{cf: c.cf, cs: append(c.cs, "CONSUMERS", Key, Groupname)}
}

func (c Xinfo) Groups(Key string) XinfoGroups {
	return XinfoGroups{cf: c.cf, cs: append(c.cs, "GROUPS", Key)}
}

func (c Xinfo) Stream(Key string) XinfoStream {
	return XinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c Xinfo) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c Xinfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) Xinfo() (c Xinfo) {
	c.cs = append(b.get(), "XINFO")
	return
}

type XinfoConsumers struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XinfoConsumers) Groups(Key string) XinfoGroups {
	return XinfoGroups{cf: c.cf, cs: append(c.cs, "GROUPS", Key)}
}

func (c XinfoConsumers) Stream(Key string) XinfoStream {
	return XinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c XinfoConsumers) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c XinfoConsumers) Build() Completed {
	return Completed(c)
}

type XinfoGroups struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XinfoGroups) Stream(Key string) XinfoStream {
	return XinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c XinfoGroups) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c XinfoGroups) Build() Completed {
	return Completed(c)
}

type XinfoHelpHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XinfoHelpHelp) Build() Completed {
	return Completed(c)
}

type XinfoStream struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XinfoStream) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c XinfoStream) Build() Completed {
	return Completed(c)
}

type Xlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xlen) Key(Key string) XlenKey {
	return XlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xlen() (c Xlen) {
	c.cs = append(b.get(), "XLEN")
	return
}

type XlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XlenKey) Build() Completed {
	return Completed(c)
}

type Xpending struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xpending) Key(Key string) XpendingKey {
	return XpendingKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xpending() (c Xpending) {
	c.cs = append(b.get(), "XPENDING")
	return
}

type XpendingFiltersConsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XpendingFiltersConsumer) Build() Completed {
	return Completed(c)
}

type XpendingFiltersCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XpendingFiltersCount) Consumer(Consumer string) XpendingFiltersConsumer {
	return XpendingFiltersConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

func (c XpendingFiltersCount) Build() Completed {
	return Completed(c)
}

type XpendingFiltersEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XpendingFiltersEnd) Count(Count int64) XpendingFiltersCount {
	return XpendingFiltersCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type XpendingFiltersIdle struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XpendingFiltersIdle) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XpendingFiltersStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XpendingFiltersStart) End(End string) XpendingFiltersEnd {
	return XpendingFiltersEnd{cf: c.cf, cs: append(c.cs, End)}
}

type XpendingGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XpendingGroup) Idle(MinIdleTime int64) XpendingFiltersIdle {
	return XpendingFiltersIdle{cf: c.cf, cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10))}
}

func (c XpendingGroup) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XpendingKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XpendingKey) Group(Group string) XpendingGroup {
	return XpendingGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type Xrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xrange) Key(Key string) XrangeKey {
	return XrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xrange() (c Xrange) {
	c.cs = append(b.get(), "XRANGE")
	return
}

type XrangeCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrangeCount) Build() Completed {
	return Completed(c)
}

type XrangeEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrangeEnd) Count(Count int64) XrangeCount {
	return XrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrangeEnd) Build() Completed {
	return Completed(c)
}

type XrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrangeKey) Start(Start string) XrangeStart {
	return XrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrangeStart) End(End string) XrangeEnd {
	return XrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type Xread struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xread) Count(Count int64) XreadCount {
	return XreadCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c Xread) Block(Milliseconds int64) XreadBlock {
	c.cf = blockTag
	return XreadBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c Xread) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

func (b *Builder) Xread() (c Xread) {
	c.cs = append(b.get(), "XREAD")
	return
}

type XreadBlock struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadBlock) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadCount) Block(Milliseconds int64) XreadBlock {
	c.cf = blockTag
	return XreadBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadCount) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadId) Id(Id ...string) XreadId {
	return XreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadId) Build() Completed {
	return Completed(c)
}

type XreadKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadKey) Id(Id ...string) XreadId {
	return XreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadKey) Key(Key ...string) XreadKey {
	return XreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type XreadStreamsStreams struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadStreamsStreams) Key(Key ...string) XreadKey {
	return XreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Xreadgroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xreadgroup) Group(Group string, Consumer string) XreadgroupGroup {
	return XreadgroupGroup{cf: c.cf, cs: append(c.cs, "GROUP", Group, Consumer)}
}

func (b *Builder) Xreadgroup() (c Xreadgroup) {
	c.cs = append(b.get(), "XREADGROUP")
	return
}

type XreadgroupBlock struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadgroupBlock) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c XreadgroupBlock) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadgroupCount) Block(Milliseconds int64) XreadgroupBlock {
	c.cf = blockTag
	return XreadgroupBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadgroupCount) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c XreadgroupCount) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadgroupGroup) Count(Count int64) XreadgroupCount {
	return XreadgroupCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XreadgroupGroup) Block(Milliseconds int64) XreadgroupBlock {
	c.cf = blockTag
	return XreadgroupBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadgroupGroup) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c XreadgroupGroup) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadgroupId) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadgroupId) Build() Completed {
	return Completed(c)
}

type XreadgroupKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadgroupKey) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadgroupKey) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type XreadgroupNoackNoack struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadgroupNoackNoack) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupStreamsStreams struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XreadgroupStreamsStreams) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Xrevrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xrevrange) Key(Key string) XrevrangeKey {
	return XrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xrevrange() (c Xrevrange) {
	c.cs = append(b.get(), "XREVRANGE")
	return
}

type XrevrangeCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrevrangeCount) Build() Completed {
	return Completed(c)
}

type XrevrangeEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrevrangeEnd) Start(Start string) XrevrangeStart {
	return XrevrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XrevrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrevrangeKey) End(End string) XrevrangeEnd {
	return XrevrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type XrevrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XrevrangeStart) Count(Count int64) XrevrangeCount {
	return XrevrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrevrangeStart) Build() Completed {
	return Completed(c)
}

type Xtrim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Xtrim) Key(Key string) XtrimKey {
	return XtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xtrim() (c Xtrim) {
	c.cs = append(b.get(), "XTRIM")
	return
}

type XtrimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XtrimKey) Maxlen() XtrimTrimStrategyMaxlen {
	return XtrimTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c XtrimKey) Minid() XtrimTrimStrategyMinid {
	return XtrimTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

type XtrimTrimLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XtrimTrimLimit) Build() Completed {
	return Completed(c)
}

type XtrimTrimOperatorAlmost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XtrimTrimOperatorAlmost) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimOperatorExact struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XtrimTrimOperatorExact) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimStrategyMaxlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XtrimTrimStrategyMaxlen) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XtrimTrimStrategyMaxlen) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XtrimTrimStrategyMaxlen) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimStrategyMinid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XtrimTrimStrategyMinid) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XtrimTrimStrategyMinid) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XtrimTrimStrategyMinid) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimThreshold struct {
	cs []string
	cf uint16
	ks uint16
}

func (c XtrimTrimThreshold) Limit(Count int64) XtrimTrimLimit {
	return XtrimTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XtrimTrimThreshold) Build() Completed {
	return Completed(c)
}

type Zadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zadd) Key(Key string) ZaddKey {
	return ZaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zadd() (c Zadd) {
	c.cs = append(b.get(), "ZADD")
	return
}

type ZaddChangeCh struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddChangeCh) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddChangeCh) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddComparisonGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddComparisonGt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c ZaddComparisonGt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddComparisonGt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddComparisonLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddComparisonLt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c ZaddComparisonLt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddComparisonLt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddConditionNx) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c ZaddConditionNx) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c ZaddConditionNx) Ch() ZaddChangeCh {
	return ZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c ZaddConditionNx) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddConditionNx) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddConditionXx) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c ZaddConditionXx) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c ZaddConditionXx) Ch() ZaddChangeCh {
	return ZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c ZaddConditionXx) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddConditionXx) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddIncrementIncr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddIncrementIncr) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddKey) Nx() ZaddConditionNx {
	return ZaddConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c ZaddKey) Xx() ZaddConditionXx {
	return ZaddConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c ZaddKey) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c ZaddKey) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c ZaddKey) Ch() ZaddChangeCh {
	return ZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c ZaddKey) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddKey) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddScoreMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZaddScoreMember) ScoreMember(Score float64, Member string) ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member)}
}

func (c ZaddScoreMember) Build() Completed {
	return Completed(c)
}

type Zcard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zcard) Key(Key string) ZcardKey {
	return ZcardKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zcard() (c Zcard) {
	c.cs = append(b.get(), "ZCARD")
	return
}

type ZcardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZcardKey) Build() Completed {
	return Completed(c)
}

func (c ZcardKey) Cache() Cacheable {
	return Cacheable(c)
}

type Zcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zcount) Key(Key string) ZcountKey {
	return ZcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zcount() (c Zcount) {
	c.cs = append(b.get(), "ZCOUNT")
	return
}

type ZcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZcountKey) Min(Min float64) ZcountMin {
	return ZcountMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c ZcountKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZcountMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZcountMax) Build() Completed {
	return Completed(c)
}

func (c ZcountMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZcountMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZcountMin) Max(Max float64) ZcountMax {
	return ZcountMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c ZcountMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zdiff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zdiff) Numkeys(Numkeys int64) ZdiffNumkeys {
	return ZdiffNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zdiff() (c Zdiff) {
	c.cs = append(b.get(), "ZDIFF")
	return
}

type ZdiffKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZdiffKey) Withscores() ZdiffWithscoresWithscores {
	return ZdiffWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZdiffKey) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZdiffKey) Build() Completed {
	return Completed(c)
}

type ZdiffNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZdiffNumkeys) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZdiffWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZdiffWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zdiffstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zdiffstore) Destination(Destination string) ZdiffstoreDestination {
	return ZdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Zdiffstore() (c Zdiffstore) {
	c.cs = append(b.get(), "ZDIFFSTORE")
	return
}

type ZdiffstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZdiffstoreDestination) Numkeys(Numkeys int64) ZdiffstoreNumkeys {
	return ZdiffstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZdiffstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZdiffstoreKey) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZdiffstoreKey) Build() Completed {
	return Completed(c)
}

type ZdiffstoreNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZdiffstoreNumkeys) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Zincrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zincrby) Key(Key string) ZincrbyKey {
	return ZincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zincrby() (c Zincrby) {
	c.cs = append(b.get(), "ZINCRBY")
	return
}

type ZincrbyIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZincrbyIncrement) Member(Member string) ZincrbyMember {
	return ZincrbyMember{cf: c.cf, cs: append(c.cs, Member)}
}

type ZincrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZincrbyKey) Increment(Increment int64) ZincrbyIncrement {
	return ZincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type ZincrbyMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZincrbyMember) Build() Completed {
	return Completed(c)
}

type Zinter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zinter) Numkeys(Numkeys int64) ZinterNumkeys {
	return ZinterNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zinter() (c Zinter) {
	c.cs = append(b.get(), "ZINTER")
	return
}

type ZinterAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterAggregateMax) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMax) Build() Completed {
	return Completed(c)
}

type ZinterAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterAggregateMin) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMin) Build() Completed {
	return Completed(c)
}

type ZinterAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterAggregateSum) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateSum) Build() Completed {
	return Completed(c)
}

type ZinterKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterKey) Weights(Weight ...int64) ZinterWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterWeights{cf: c.cf, cs: c.cs}
}

func (c ZinterKey) Sum() ZinterAggregateSum {
	return ZinterAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZinterKey) Min() ZinterAggregateMin {
	return ZinterAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZinterKey) Max() ZinterAggregateMax {
	return ZinterAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZinterKey) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterKey) Key(Key ...string) ZinterKey {
	return ZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZinterKey) Build() Completed {
	return Completed(c)
}

type ZinterNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterNumkeys) Key(Key ...string) ZinterKey {
	return ZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZinterWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterWeights) Sum() ZinterAggregateSum {
	return ZinterAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZinterWeights) Min() ZinterAggregateMin {
	return ZinterAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZinterWeights) Max() ZinterAggregateMax {
	return ZinterAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZinterWeights) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterWeights) Weights(Weights ...int64) ZinterWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterWeights{cf: c.cf, cs: c.cs}
}

func (c ZinterWeights) Build() Completed {
	return Completed(c)
}

type ZinterWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zintercard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zintercard) Numkeys(Numkeys int64) ZintercardNumkeys {
	return ZintercardNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zintercard() (c Zintercard) {
	c.cs = append(b.get(), "ZINTERCARD")
	return
}

type ZintercardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZintercardKey) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZintercardKey) Build() Completed {
	return Completed(c)
}

type ZintercardNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZintercardNumkeys) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Zinterstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zinterstore) Destination(Destination string) ZinterstoreDestination {
	return ZinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Zinterstore() (c Zinterstore) {
	c.cs = append(b.get(), "ZINTERSTORE")
	return
}

type ZinterstoreAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterstoreAggregateMax) Build() Completed {
	return Completed(c)
}

type ZinterstoreAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterstoreAggregateMin) Build() Completed {
	return Completed(c)
}

type ZinterstoreAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterstoreAggregateSum) Build() Completed {
	return Completed(c)
}

type ZinterstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterstoreDestination) Numkeys(Numkeys int64) ZinterstoreNumkeys {
	return ZinterstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZinterstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterstoreKey) Weights(Weight ...int64) ZinterstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterstoreWeights{cf: c.cf, cs: c.cs}
}

func (c ZinterstoreKey) Sum() ZinterstoreAggregateSum {
	return ZinterstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZinterstoreKey) Min() ZinterstoreAggregateMin {
	return ZinterstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZinterstoreKey) Max() ZinterstoreAggregateMax {
	return ZinterstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZinterstoreKey) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZinterstoreKey) Build() Completed {
	return Completed(c)
}

type ZinterstoreNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterstoreNumkeys) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZinterstoreWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZinterstoreWeights) Sum() ZinterstoreAggregateSum {
	return ZinterstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZinterstoreWeights) Min() ZinterstoreAggregateMin {
	return ZinterstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZinterstoreWeights) Max() ZinterstoreAggregateMax {
	return ZinterstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZinterstoreWeights) Weights(Weights ...int64) ZinterstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterstoreWeights{cf: c.cf, cs: c.cs}
}

func (c ZinterstoreWeights) Build() Completed {
	return Completed(c)
}

type Zlexcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zlexcount) Key(Key string) ZlexcountKey {
	return ZlexcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zlexcount() (c Zlexcount) {
	c.cs = append(b.get(), "ZLEXCOUNT")
	return
}

type ZlexcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZlexcountKey) Min(Min string) ZlexcountMin {
	return ZlexcountMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZlexcountKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZlexcountMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZlexcountMax) Build() Completed {
	return Completed(c)
}

func (c ZlexcountMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZlexcountMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZlexcountMin) Max(Max string) ZlexcountMax {
	return ZlexcountMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZlexcountMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zmscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zmscore) Key(Key string) ZmscoreKey {
	return ZmscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zmscore() (c Zmscore) {
	c.cs = append(b.get(), "ZMSCORE")
	return
}

type ZmscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZmscoreKey) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c ZmscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZmscoreMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZmscoreMember) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c ZmscoreMember) Build() Completed {
	return Completed(c)
}

func (c ZmscoreMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zpopmax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zpopmax) Key(Key string) ZpopmaxKey {
	return ZpopmaxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zpopmax() (c Zpopmax) {
	c.cs = append(b.get(), "ZPOPMAX")
	return
}

type ZpopmaxCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZpopmaxCount) Build() Completed {
	return Completed(c)
}

type ZpopmaxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZpopmaxKey) Count(Count int64) ZpopmaxCount {
	return ZpopmaxCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopmaxKey) Build() Completed {
	return Completed(c)
}

type Zpopmin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zpopmin) Key(Key string) ZpopminKey {
	return ZpopminKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zpopmin() (c Zpopmin) {
	c.cs = append(b.get(), "ZPOPMIN")
	return
}

type ZpopminCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZpopminCount) Build() Completed {
	return Completed(c)
}

type ZpopminKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZpopminKey) Count(Count int64) ZpopminCount {
	return ZpopminCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopminKey) Build() Completed {
	return Completed(c)
}

type Zrandmember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrandmember) Key(Key string) ZrandmemberKey {
	return ZrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrandmember() (c Zrandmember) {
	c.cs = append(b.get(), "ZRANDMEMBER")
	return
}

type ZrandmemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrandmemberKey) Count(Count int64) ZrandmemberOptionsCount {
	return ZrandmemberOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type ZrandmemberOptionsCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrandmemberOptionsCount) Withscores() ZrandmemberOptionsWithscoresWithscores {
	return ZrandmemberOptionsWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrandmemberOptionsCount) Build() Completed {
	return Completed(c)
}

type ZrandmemberOptionsWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrandmemberOptionsWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrange) Key(Key string) ZrangeKey {
	return ZrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrange() (c Zrange) {
	c.cs = append(b.get(), "ZRANGE")
	return
}

type ZrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeKey) Min(Min string) ZrangeMin {
	return ZrangeMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeLimit) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangeLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeMax) Byscore() ZrangeSortbyByscore {
	return ZrangeSortbyByscore{cf: c.cf, cs: append(c.cs, "BYSCORE")}
}

func (c ZrangeMax) Bylex() ZrangeSortbyBylex {
	return ZrangeSortbyBylex{cf: c.cf, cs: append(c.cs, "BYLEX")}
}

func (c ZrangeMax) Rev() ZrangeRevRev {
	return ZrangeRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangeMax) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeMax) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeMax) Build() Completed {
	return Completed(c)
}

func (c ZrangeMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeMin) Max(Max string) ZrangeMax {
	return ZrangeMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZrangeMin) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeRevRev struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeRevRev) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeRevRev) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeRevRev) Build() Completed {
	return Completed(c)
}

func (c ZrangeRevRev) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeSortbyBylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeSortbyBylex) Rev() ZrangeRevRev {
	return ZrangeRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangeSortbyBylex) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeSortbyBylex) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeSortbyBylex) Build() Completed {
	return Completed(c)
}

func (c ZrangeSortbyBylex) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeSortbyByscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeSortbyByscore) Rev() ZrangeRevRev {
	return ZrangeRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangeSortbyByscore) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeSortbyByscore) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeSortbyByscore) Build() Completed {
	return Completed(c)
}

func (c ZrangeSortbyByscore) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangeWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrangeWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangebylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrangebylex) Key(Key string) ZrangebylexKey {
	return ZrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrangebylex() (c Zrangebylex) {
	c.cs = append(b.get(), "ZRANGEBYLEX")
	return
}

type ZrangebylexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebylexKey) Min(Min string) ZrangebylexMin {
	return ZrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZrangebylexKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebylexLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangebylexLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebylexMax) Limit(Offset int64, Count int64) ZrangebylexLimit {
	return ZrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebylexMax) Build() Completed {
	return Completed(c)
}

func (c ZrangebylexMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebylexMin) Max(Max string) ZrangebylexMax {
	return ZrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZrangebylexMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangebyscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrangebyscore) Key(Key string) ZrangebyscoreKey {
	return ZrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrangebyscore() (c Zrangebyscore) {
	c.cs = append(b.get(), "ZRANGEBYSCORE")
	return
}

type ZrangebyscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebyscoreKey) Min(Min float64) ZrangebyscoreMin {
	return ZrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c ZrangebyscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebyscoreLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebyscoreMax) Withscores() ZrangebyscoreWithscoresWithscores {
	return ZrangebyscoreWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangebyscoreMax) Limit(Offset int64, Count int64) ZrangebyscoreLimit {
	return ZrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebyscoreMax) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebyscoreMin) Max(Max float64) ZrangebyscoreMax {
	return ZrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c ZrangebyscoreMin) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrangebyscoreLimit {
	return ZrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebyscoreWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangestore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrangestore) Dst(Dst string) ZrangestoreDst {
	return ZrangestoreDst{cf: c.cf, cs: append(c.cs, Dst)}
}

func (b *Builder) Zrangestore() (c Zrangestore) {
	c.cs = append(b.get(), "ZRANGESTORE")
	return
}

type ZrangestoreDst struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreDst) Src(Src string) ZrangestoreSrc {
	return ZrangestoreSrc{cf: c.cf, cs: append(c.cs, Src)}
}

type ZrangestoreLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreLimit) Build() Completed {
	return Completed(c)
}

type ZrangestoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreMax) Byscore() ZrangestoreSortbyByscore {
	return ZrangestoreSortbyByscore{cf: c.cf, cs: append(c.cs, "BYSCORE")}
}

func (c ZrangestoreMax) Bylex() ZrangestoreSortbyBylex {
	return ZrangestoreSortbyBylex{cf: c.cf, cs: append(c.cs, "BYLEX")}
}

func (c ZrangestoreMax) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangestoreMax) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreMax) Build() Completed {
	return Completed(c)
}

type ZrangestoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreMin) Max(Max string) ZrangestoreMax {
	return ZrangestoreMax{cf: c.cf, cs: append(c.cs, Max)}
}

type ZrangestoreRevRev struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreRevRev) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreRevRev) Build() Completed {
	return Completed(c)
}

type ZrangestoreSortbyBylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreSortbyBylex) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangestoreSortbyBylex) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreSortbyBylex) Build() Completed {
	return Completed(c)
}

type ZrangestoreSortbyByscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreSortbyByscore) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangestoreSortbyByscore) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreSortbyByscore) Build() Completed {
	return Completed(c)
}

type ZrangestoreSrc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrangestoreSrc) Min(Min string) ZrangestoreMin {
	return ZrangestoreMin{cf: c.cf, cs: append(c.cs, Min)}
}

type Zrank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrank) Key(Key string) ZrankKey {
	return ZrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrank() (c Zrank) {
	c.cs = append(b.get(), "ZRANK")
	return
}

type ZrankKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrankKey) Member(Member string) ZrankMember {
	return ZrankMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c ZrankKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrankMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrankMember) Build() Completed {
	return Completed(c)
}

func (c ZrankMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zrem struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrem) Key(Key string) ZremKey {
	return ZremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrem() (c Zrem) {
	c.cs = append(b.get(), "ZREM")
	return
}

type ZremKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremKey) Member(Member ...string) ZremMember {
	return ZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type ZremMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremMember) Member(Member ...string) ZremMember {
	return ZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c ZremMember) Build() Completed {
	return Completed(c)
}

type Zremrangebylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zremrangebylex) Key(Key string) ZremrangebylexKey {
	return ZremrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebylex() (c Zremrangebylex) {
	c.cs = append(b.get(), "ZREMRANGEBYLEX")
	return
}

type ZremrangebylexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebylexKey) Min(Min string) ZremrangebylexMin {
	return ZremrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

type ZremrangebylexMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebylexMax) Build() Completed {
	return Completed(c)
}

type ZremrangebylexMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebylexMin) Max(Max string) ZremrangebylexMax {
	return ZremrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

type Zremrangebyrank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zremrangebyrank) Key(Key string) ZremrangebyrankKey {
	return ZremrangebyrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebyrank() (c Zremrangebyrank) {
	c.cs = append(b.get(), "ZREMRANGEBYRANK")
	return
}

type ZremrangebyrankKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebyrankKey) Start(Start int64) ZremrangebyrankStart {
	return ZremrangebyrankStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type ZremrangebyrankStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebyrankStart) Stop(Stop int64) ZremrangebyrankStop {
	return ZremrangebyrankStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type ZremrangebyrankStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebyrankStop) Build() Completed {
	return Completed(c)
}

type Zremrangebyscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zremrangebyscore) Key(Key string) ZremrangebyscoreKey {
	return ZremrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebyscore() (c Zremrangebyscore) {
	c.cs = append(b.get(), "ZREMRANGEBYSCORE")
	return
}

type ZremrangebyscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebyscoreKey) Min(Min float64) ZremrangebyscoreMin {
	return ZremrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZremrangebyscoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebyscoreMax) Build() Completed {
	return Completed(c)
}

type ZremrangebyscoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZremrangebyscoreMin) Max(Max float64) ZremrangebyscoreMax {
	return ZremrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type Zrevrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrevrange) Key(Key string) ZrevrangeKey {
	return ZrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrange() (c Zrevrange) {
	c.cs = append(b.get(), "ZREVRANGE")
	return
}

type ZrevrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangeKey) Start(Start int64) ZrevrangeStart {
	return ZrevrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c ZrevrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangeStart) Stop(Stop int64) ZrevrangeStop {
	return ZrevrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

func (c ZrevrangeStart) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangeStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangeStop) Withscores() ZrevrangeWithscoresWithscores {
	return ZrevrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrevrangeStop) Build() Completed {
	return Completed(c)
}

func (c ZrevrangeStop) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangeWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangeWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrevrangeWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrangebylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrevrangebylex) Key(Key string) ZrevrangebylexKey {
	return ZrevrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrangebylex() (c Zrevrangebylex) {
	c.cs = append(b.get(), "ZREVRANGEBYLEX")
	return
}

type ZrevrangebylexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebylexKey) Max(Max string) ZrevrangebylexMax {
	return ZrevrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZrevrangebylexKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebylexLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebylexLimit) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebylexLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebylexMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebylexMax) Min(Min string) ZrevrangebylexMin {
	return ZrevrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZrevrangebylexMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebylexMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebylexMin) Limit(Offset int64, Count int64) ZrevrangebylexLimit {
	return ZrevrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebylexMin) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebylexMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrangebyscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrevrangebyscore) Key(Key string) ZrevrangebyscoreKey {
	return ZrevrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrangebyscore() (c Zrevrangebyscore) {
	c.cs = append(b.get(), "ZREVRANGEBYSCORE")
	return
}

type ZrevrangebyscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebyscoreKey) Max(Max float64) ZrevrangebyscoreMax {
	return ZrevrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c ZrevrangebyscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebyscoreLimit) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebyscoreMax) Min(Min float64) ZrevrangebyscoreMin {
	return ZrevrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c ZrevrangebyscoreMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebyscoreMin) Withscores() ZrevrangebyscoreWithscoresWithscores {
	return ZrevrangebyscoreWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrevrangebyscoreMin) Limit(Offset int64, Count int64) ZrevrangebyscoreLimit {
	return ZrevrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebyscoreMin) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreMin) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrevrangebyscoreLimit {
	return ZrevrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebyscoreWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zrevrank) Key(Key string) ZrevrankKey {
	return ZrevrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrank() (c Zrevrank) {
	c.cs = append(b.get(), "ZREVRANK")
	return
}

type ZrevrankKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrankKey) Member(Member string) ZrevrankMember {
	return ZrevrankMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c ZrevrankKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrankMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZrevrankMember) Build() Completed {
	return Completed(c)
}

func (c ZrevrankMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zscan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zscan) Key(Key string) ZscanKey {
	return ZscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zscan() (c Zscan) {
	c.cs = append(b.get(), "ZSCAN")
	return
}

type ZscanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZscanCount) Build() Completed {
	return Completed(c)
}

type ZscanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZscanCursor) Match(Pattern string) ZscanMatch {
	return ZscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c ZscanCursor) Count(Count int64) ZscanCount {
	return ZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ZscanCursor) Build() Completed {
	return Completed(c)
}

type ZscanKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZscanKey) Cursor(Cursor int64) ZscanCursor {
	return ZscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type ZscanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZscanMatch) Count(Count int64) ZscanCount {
	return ZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ZscanMatch) Build() Completed {
	return Completed(c)
}

type Zscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zscore) Key(Key string) ZscoreKey {
	return ZscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zscore() (c Zscore) {
	c.cs = append(b.get(), "ZSCORE")
	return
}

type ZscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZscoreKey) Member(Member string) ZscoreMember {
	return ZscoreMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c ZscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZscoreMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZscoreMember) Build() Completed {
	return Completed(c)
}

func (c ZscoreMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zunion struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zunion) Numkeys(Numkeys int64) ZunionNumkeys {
	return ZunionNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zunion() (c Zunion) {
	c.cs = append(b.get(), "ZUNION")
	return
}

type ZunionAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionAggregateMax) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMax) Build() Completed {
	return Completed(c)
}

type ZunionAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionAggregateMin) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMin) Build() Completed {
	return Completed(c)
}

type ZunionAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionAggregateSum) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateSum) Build() Completed {
	return Completed(c)
}

type ZunionKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionKey) Weights(Weight ...int64) ZunionWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionWeights{cf: c.cf, cs: c.cs}
}

func (c ZunionKey) Sum() ZunionAggregateSum {
	return ZunionAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZunionKey) Min() ZunionAggregateMin {
	return ZunionAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZunionKey) Max() ZunionAggregateMax {
	return ZunionAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZunionKey) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionKey) Key(Key ...string) ZunionKey {
	return ZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZunionKey) Build() Completed {
	return Completed(c)
}

type ZunionNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionNumkeys) Key(Key ...string) ZunionKey {
	return ZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZunionWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionWeights) Sum() ZunionAggregateSum {
	return ZunionAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZunionWeights) Min() ZunionAggregateMin {
	return ZunionAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZunionWeights) Max() ZunionAggregateMax {
	return ZunionAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZunionWeights) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionWeights) Weights(Weights ...int64) ZunionWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionWeights{cf: c.cf, cs: c.cs}
}

func (c ZunionWeights) Build() Completed {
	return Completed(c)
}

type ZunionWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zunionstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c Zunionstore) Destination(Destination string) ZunionstoreDestination {
	return ZunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Zunionstore() (c Zunionstore) {
	c.cs = append(b.get(), "ZUNIONSTORE")
	return
}

type ZunionstoreAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionstoreAggregateMax) Build() Completed {
	return Completed(c)
}

type ZunionstoreAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionstoreAggregateMin) Build() Completed {
	return Completed(c)
}

type ZunionstoreAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionstoreAggregateSum) Build() Completed {
	return Completed(c)
}

type ZunionstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionstoreDestination) Numkeys(Numkeys int64) ZunionstoreNumkeys {
	return ZunionstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZunionstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionstoreKey) Weights(Weight ...int64) ZunionstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionstoreWeights{cf: c.cf, cs: c.cs}
}

func (c ZunionstoreKey) Sum() ZunionstoreAggregateSum {
	return ZunionstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZunionstoreKey) Min() ZunionstoreAggregateMin {
	return ZunionstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZunionstoreKey) Max() ZunionstoreAggregateMax {
	return ZunionstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZunionstoreKey) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZunionstoreKey) Build() Completed {
	return Completed(c)
}

type ZunionstoreNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionstoreNumkeys) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZunionstoreWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c ZunionstoreWeights) Sum() ZunionstoreAggregateSum {
	return ZunionstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c ZunionstoreWeights) Min() ZunionstoreAggregateMin {
	return ZunionstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c ZunionstoreWeights) Max() ZunionstoreAggregateMax {
	return ZunionstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c ZunionstoreWeights) Weights(Weights ...int64) ZunionstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionstoreWeights{cf: c.cf, cs: c.cs}
}

func (c ZunionstoreWeights) Build() Completed {
	return Completed(c)
}

type SAclCat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclCat) Categoryname(Categoryname string) SAclCatCategoryname {
	return SAclCatCategoryname{cf: c.cf, cs: append(c.cs, Categoryname)}
}

func (c SAclCat) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclCat() (c SAclCat) {
	c.cs = append(b.get(), "ACL", "CAT")
	c.ks = initSlot
	return
}

type SAclCatCategoryname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclCatCategoryname) Build() SCompleted {
	return SCompleted(c)
}

type SAclDeluser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclDeluser) Username(Username ...string) SAclDeluserUsername {
	return SAclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (b *SBuilder) AclDeluser() (c SAclDeluser) {
	c.cs = append(b.get(), "ACL", "DELUSER")
	c.ks = initSlot
	return
}

type SAclDeluserUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclDeluserUsername) Username(Username ...string) SAclDeluserUsername {
	return SAclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (c SAclDeluserUsername) Build() SCompleted {
	return SCompleted(c)
}

type SAclGenpass struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclGenpass) Bits(Bits int64) SAclGenpassBits {
	return SAclGenpassBits{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bits, 10))}
}

func (c SAclGenpass) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclGenpass() (c SAclGenpass) {
	c.cs = append(b.get(), "ACL", "GENPASS")
	c.ks = initSlot
	return
}

type SAclGenpassBits struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclGenpassBits) Build() SCompleted {
	return SCompleted(c)
}

type SAclGetuser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclGetuser) Username(Username string) SAclGetuserUsername {
	return SAclGetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (b *SBuilder) AclGetuser() (c SAclGetuser) {
	c.cs = append(b.get(), "ACL", "GETUSER")
	c.ks = initSlot
	return
}

type SAclGetuserUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclGetuserUsername) Build() SCompleted {
	return SCompleted(c)
}

type SAclHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclHelp) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclHelp() (c SAclHelp) {
	c.cs = append(b.get(), "ACL", "HELP")
	c.ks = initSlot
	return
}

type SAclList struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclList) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclList() (c SAclList) {
	c.cs = append(b.get(), "ACL", "LIST")
	c.ks = initSlot
	return
}

type SAclLoad struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclLoad) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclLoad() (c SAclLoad) {
	c.cs = append(b.get(), "ACL", "LOAD")
	c.ks = initSlot
	return
}

type SAclLog struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclLog) CountOrReset(CountOrReset string) SAclLogCountOrReset {
	return SAclLogCountOrReset{cf: c.cf, cs: append(c.cs, CountOrReset)}
}

func (c SAclLog) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclLog() (c SAclLog) {
	c.cs = append(b.get(), "ACL", "LOG")
	c.ks = initSlot
	return
}

type SAclLogCountOrReset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclLogCountOrReset) Build() SCompleted {
	return SCompleted(c)
}

type SAclSave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclSave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclSave() (c SAclSave) {
	c.cs = append(b.get(), "ACL", "SAVE")
	c.ks = initSlot
	return
}

type SAclSetuser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclSetuser) Username(Username string) SAclSetuserUsername {
	return SAclSetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (b *SBuilder) AclSetuser() (c SAclSetuser) {
	c.cs = append(b.get(), "ACL", "SETUSER")
	c.ks = initSlot
	return
}

type SAclSetuserRule struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclSetuserRule) Rule(Rule ...string) SAclSetuserRule {
	return SAclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c SAclSetuserRule) Build() SCompleted {
	return SCompleted(c)
}

type SAclSetuserUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclSetuserUsername) Rule(Rule ...string) SAclSetuserRule {
	return SAclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c SAclSetuserUsername) Build() SCompleted {
	return SCompleted(c)
}

type SAclUsers struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclUsers) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclUsers() (c SAclUsers) {
	c.cs = append(b.get(), "ACL", "USERS")
	c.ks = initSlot
	return
}

type SAclWhoami struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAclWhoami) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclWhoami() (c SAclWhoami) {
	c.cs = append(b.get(), "ACL", "WHOAMI")
	c.ks = initSlot
	return
}

type SAppend struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAppend) Key(Key string) SAppendKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SAppendKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Append() (c SAppend) {
	c.cs = append(b.get(), "APPEND")
	c.ks = initSlot
	return
}

type SAppendKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAppendKey) Value(Value string) SAppendValue {
	return SAppendValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SAppendValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAppendValue) Build() SCompleted {
	return SCompleted(c)
}

type SAsking struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAsking) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Asking() (c SAsking) {
	c.cs = append(b.get(), "ASKING")
	c.ks = initSlot
	return
}

type SAuth struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAuth) Username(Username string) SAuthUsername {
	return SAuthUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (c SAuth) Password(Password string) SAuthPassword {
	return SAuthPassword{cf: c.cf, cs: append(c.cs, Password)}
}

func (b *SBuilder) Auth() (c SAuth) {
	c.cs = append(b.get(), "AUTH")
	c.ks = initSlot
	return
}

type SAuthPassword struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAuthPassword) Build() SCompleted {
	return SCompleted(c)
}

type SAuthUsername struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SAuthUsername) Password(Password string) SAuthPassword {
	return SAuthPassword{cf: c.cf, cs: append(c.cs, Password)}
}

type SBgrewriteaof struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBgrewriteaof) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Bgrewriteaof() (c SBgrewriteaof) {
	c.cs = append(b.get(), "BGREWRITEAOF")
	c.ks = initSlot
	return
}

type SBgsave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBgsave) Schedule() SBgsaveScheduleSchedule {
	return SBgsaveScheduleSchedule{cf: c.cf, cs: append(c.cs, "SCHEDULE")}
}

func (c SBgsave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Bgsave() (c SBgsave) {
	c.cs = append(b.get(), "BGSAVE")
	c.ks = initSlot
	return
}

type SBgsaveScheduleSchedule struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBgsaveScheduleSchedule) Build() SCompleted {
	return SCompleted(c)
}

type SBitcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitcount) Key(Key string) SBitcountKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBitcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Bitcount() (c SBitcount) {
	c.cs = append(b.get(), "BITCOUNT")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SBitcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitcountKey) StartEnd(Start int64, End int64) SBitcountStartEnd {
	return SBitcountStartEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10), strconv.FormatInt(End, 10))}
}

func (c SBitcountKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitcountKey) Cache() SCacheable {
	return SCacheable(c)
}

type SBitcountStartEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitcountStartEnd) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitcountStartEnd) Cache() SCacheable {
	return SCacheable(c)
}

type SBitfield struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfield) Key(Key string) SBitfieldKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBitfieldKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Bitfield() (c SBitfield) {
	c.cs = append(b.get(), "BITFIELD")
	c.ks = initSlot
	return
}

type SBitfieldFail struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldFail) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldGet) Set(Type string, Offset int64, Value int64) SBitfieldSet {
	return SBitfieldSet{cf: c.cf, cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10))}
}

func (c SBitfieldGet) Incrby(Type string, Offset int64, Increment int64) SBitfieldIncrby {
	return SBitfieldIncrby{cf: c.cf, cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c SBitfieldGet) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c SBitfieldGet) Sat() SBitfieldSat {
	return SBitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c SBitfieldGet) Fail() SBitfieldFail {
	return SBitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c SBitfieldGet) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldIncrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldIncrby) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c SBitfieldIncrby) Sat() SBitfieldSat {
	return SBitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c SBitfieldIncrby) Fail() SBitfieldFail {
	return SBitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c SBitfieldIncrby) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldKey) Get(Type string, Offset int64) SBitfieldGet {
	return SBitfieldGet{cf: c.cf, cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

func (c SBitfieldKey) Set(Type string, Offset int64, Value int64) SBitfieldSet {
	return SBitfieldSet{cf: c.cf, cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10))}
}

func (c SBitfieldKey) Incrby(Type string, Offset int64, Increment int64) SBitfieldIncrby {
	return SBitfieldIncrby{cf: c.cf, cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c SBitfieldKey) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c SBitfieldKey) Sat() SBitfieldSat {
	return SBitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c SBitfieldKey) Fail() SBitfieldFail {
	return SBitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c SBitfieldKey) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldRo) Key(Key string) SBitfieldRoKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBitfieldRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) BitfieldRo() (c SBitfieldRo) {
	c.cs = append(b.get(), "BITFIELD_RO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SBitfieldRoGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldRoGet) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitfieldRoGet) Cache() SCacheable {
	return SCacheable(c)
}

type SBitfieldRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldRoKey) Get(Type string, Offset int64) SBitfieldRoGet {
	return SBitfieldRoGet{cf: c.cf, cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

func (c SBitfieldRoKey) Cache() SCacheable {
	return SCacheable(c)
}

type SBitfieldSat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldSat) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldSet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldSet) Incrby(Type string, Offset int64, Increment int64) SBitfieldIncrby {
	return SBitfieldIncrby{cf: c.cf, cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c SBitfieldSet) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cf: c.cf, cs: append(c.cs, "WRAP")}
}

func (c SBitfieldSet) Sat() SBitfieldSat {
	return SBitfieldSat{cf: c.cf, cs: append(c.cs, "SAT")}
}

func (c SBitfieldSet) Fail() SBitfieldFail {
	return SBitfieldFail{cf: c.cf, cs: append(c.cs, "FAIL")}
}

func (c SBitfieldSet) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldWrap struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitfieldWrap) Build() SCompleted {
	return SCompleted(c)
}

type SBitop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitop) Operation(Operation string) SBitopOperation {
	return SBitopOperation{cf: c.cf, cs: append(c.cs, Operation)}
}

func (b *SBuilder) Bitop() (c SBitop) {
	c.cs = append(b.get(), "BITOP")
	c.ks = initSlot
	return
}

type SBitopDestkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitopDestkey) Key(Key ...string) SBitopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBitopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitopKey) Key(Key ...string) SBitopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SBitopKey) Build() SCompleted {
	return SCompleted(c)
}

type SBitopOperation struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitopOperation) Destkey(Destkey string) SBitopDestkey {
	s := slot(Destkey)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBitopDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

type SBitpos struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitpos) Key(Key string) SBitposKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBitposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Bitpos() (c SBitpos) {
	c.cs = append(b.get(), "BITPOS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SBitposBit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitposBit) Start(Start int64) SBitposIndexStart {
	return SBitposIndexStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c SBitposBit) Cache() SCacheable {
	return SCacheable(c)
}

type SBitposIndexEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitposIndexEnd) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitposIndexEnd) Cache() SCacheable {
	return SCacheable(c)
}

type SBitposIndexStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitposIndexStart) End(End int64) SBitposIndexEnd {
	return SBitposIndexEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c SBitposIndexStart) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitposIndexStart) Cache() SCacheable {
	return SCacheable(c)
}

type SBitposKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBitposKey) Bit(Bit int64) SBitposBit {
	return SBitposBit{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bit, 10))}
}

func (c SBitposKey) Cache() SCacheable {
	return SCacheable(c)
}

type SBlmove struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmove) Source(Source string) SBlmoveSource {
	s := slot(Source)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBlmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Blmove() (c SBlmove) {
	c.cs = append(b.get(), "BLMOVE")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SBlmoveDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmoveDestination) Left() SBlmoveWherefromLeft {
	return SBlmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmoveDestination) Right() SBlmoveWherefromRight {
	return SBlmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmoveSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmoveSource) Destination(Destination string) SBlmoveDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBlmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SBlmoveTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmoveTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBlmoveWherefromLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmoveWherefromLeft) Left() SBlmoveWheretoLeft {
	return SBlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmoveWherefromLeft) Right() SBlmoveWheretoRight {
	return SBlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmoveWherefromRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmoveWherefromRight) Left() SBlmoveWheretoLeft {
	return SBlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmoveWherefromRight) Right() SBlmoveWheretoRight {
	return SBlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmoveWheretoLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmoveWheretoLeft) Timeout(Timeout float64) SBlmoveTimeout {
	return SBlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type SBlmoveWheretoRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmoveWheretoRight) Timeout(Timeout float64) SBlmoveTimeout {
	return SBlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type SBlmpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmpop) Timeout(Timeout float64) SBlmpopTimeout {
	return SBlmpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (b *SBuilder) Blmpop() (c SBlmpop) {
	c.cs = append(b.get(), "BLMPOP")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SBlmpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SBlmpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmpopKey) Left() SBlmpopWhereLeft {
	return SBlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmpopKey) Right() SBlmpopWhereRight {
	return SBlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c SBlmpopKey) Key(Key ...string) SBlmpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBlmpopNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmpopNumkeys) Key(Key ...string) SBlmpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SBlmpopNumkeys) Left() SBlmpopWhereLeft {
	return SBlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmpopNumkeys) Right() SBlmpopWhereRight {
	return SBlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmpopTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmpopTimeout) Numkeys(Numkeys int64) SBlmpopNumkeys {
	return SBlmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SBlmpopWhereLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmpopWhereLeft) Count(Count int64) SBlmpopCount {
	return SBlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SBlmpopWhereLeft) Build() SCompleted {
	return SCompleted(c)
}

type SBlmpopWhereRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlmpopWhereRight) Count(Count int64) SBlmpopCount {
	return SBlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SBlmpopWhereRight) Build() SCompleted {
	return SCompleted(c)
}

type SBlpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlpop) Key(Key ...string) SBlpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Blpop() (c SBlpop) {
	c.cs = append(b.get(), "BLPOP")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SBlpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlpopKey) Timeout(Timeout float64) SBlpopTimeout {
	return SBlpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBlpopKey) Key(Key ...string) SBlpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBlpopTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBlpopTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBrpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBrpop) Key(Key ...string) SBrpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Brpop() (c SBrpop) {
	c.cs = append(b.get(), "BRPOP")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SBrpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBrpopKey) Timeout(Timeout float64) SBrpopTimeout {
	return SBrpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBrpopKey) Key(Key ...string) SBrpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBrpopTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBrpopTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBrpoplpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBrpoplpush) Source(Source string) SBrpoplpushSource {
	s := slot(Source)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBrpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Brpoplpush() (c SBrpoplpush) {
	c.cs = append(b.get(), "BRPOPLPUSH")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SBrpoplpushDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBrpoplpushDestination) Timeout(Timeout float64) SBrpoplpushTimeout {
	return SBrpoplpushTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type SBrpoplpushSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBrpoplpushSource) Destination(Destination string) SBrpoplpushDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SBrpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SBrpoplpushTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBrpoplpushTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBzpopmax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBzpopmax) Key(Key ...string) SBzpopmaxKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Bzpopmax() (c SBzpopmax) {
	c.cs = append(b.get(), "BZPOPMAX")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SBzpopmaxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBzpopmaxKey) Timeout(Timeout float64) SBzpopmaxTimeout {
	return SBzpopmaxTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBzpopmaxKey) Key(Key ...string) SBzpopmaxKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBzpopmaxTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBzpopmaxTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBzpopmin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBzpopmin) Key(Key ...string) SBzpopminKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Bzpopmin() (c SBzpopmin) {
	c.cs = append(b.get(), "BZPOPMIN")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SBzpopminKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBzpopminKey) Timeout(Timeout float64) SBzpopminTimeout {
	return SBzpopminTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBzpopminKey) Key(Key ...string) SBzpopminKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SBzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBzpopminTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SBzpopminTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientCaching struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientCaching) Yes() SClientCachingModeYes {
	return SClientCachingModeYes{cf: c.cf, cs: append(c.cs, "YES")}
}

func (c SClientCaching) No() SClientCachingModeNo {
	return SClientCachingModeNo{cf: c.cf, cs: append(c.cs, "NO")}
}

func (b *SBuilder) ClientCaching() (c SClientCaching) {
	c.cs = append(b.get(), "CLIENT", "CACHING")
	c.ks = initSlot
	return
}

type SClientCachingModeNo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientCachingModeNo) Build() SCompleted {
	return SCompleted(c)
}

type SClientCachingModeYes struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientCachingModeYes) Build() SCompleted {
	return SCompleted(c)
}

type SClientGetname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientGetname) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientGetname() (c SClientGetname) {
	c.cs = append(b.get(), "CLIENT", "GETNAME")
	c.ks = initSlot
	return
}

type SClientGetredir struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientGetredir) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientGetredir() (c SClientGetredir) {
	c.cs = append(b.get(), "CLIENT", "GETREDIR")
	c.ks = initSlot
	return
}

type SClientId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientId) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientId() (c SClientId) {
	c.cs = append(b.get(), "CLIENT", "ID")
	c.ks = initSlot
	return
}

type SClientInfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientInfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientInfo() (c SClientInfo) {
	c.cs = append(b.get(), "CLIENT", "INFO")
	c.ks = initSlot
	return
}

type SClientKill struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKill) IpPort(IpPort string) SClientKillIpPort {
	return SClientKillIpPort{cf: c.cf, cs: append(c.cs, IpPort)}
}

func (c SClientKill) Id(ClientId int64) SClientKillId {
	return SClientKillId{cf: c.cf, cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10))}
}

func (c SClientKill) Normal() SClientKillNormal {
	return SClientKillNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c SClientKill) Master() SClientKillMaster {
	return SClientKillMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c SClientKill) Slave() SClientKillSlave {
	return SClientKillSlave{cf: c.cf, cs: append(c.cs, "slave")}
}

func (c SClientKill) Pubsub() SClientKillPubsub {
	return SClientKillPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c SClientKill) User(Username string) SClientKillUser {
	return SClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c SClientKill) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKill) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKill) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKill) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientKill() (c SClientKill) {
	c.cs = append(b.get(), "CLIENT", "KILL")
	c.ks = initSlot
	return
}

type SClientKillAddr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillAddr) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillAddr) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillAddr) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillId) Normal() SClientKillNormal {
	return SClientKillNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c SClientKillId) Master() SClientKillMaster {
	return SClientKillMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c SClientKillId) Slave() SClientKillSlave {
	return SClientKillSlave{cf: c.cf, cs: append(c.cs, "slave")}
}

func (c SClientKillId) Pubsub() SClientKillPubsub {
	return SClientKillPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c SClientKillId) User(Username string) SClientKillUser {
	return SClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c SClientKillId) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKillId) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillId) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillId) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillIpPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillIpPort) Id(ClientId int64) SClientKillId {
	return SClientKillId{cf: c.cf, cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10))}
}

func (c SClientKillIpPort) Normal() SClientKillNormal {
	return SClientKillNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c SClientKillIpPort) Master() SClientKillMaster {
	return SClientKillMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c SClientKillIpPort) Slave() SClientKillSlave {
	return SClientKillSlave{cf: c.cf, cs: append(c.cs, "slave")}
}

func (c SClientKillIpPort) Pubsub() SClientKillPubsub {
	return SClientKillPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c SClientKillIpPort) User(Username string) SClientKillUser {
	return SClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c SClientKillIpPort) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKillIpPort) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillIpPort) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillIpPort) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillLaddr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillLaddr) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillLaddr) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillMaster struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillMaster) User(Username string) SClientKillUser {
	return SClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c SClientKillMaster) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKillMaster) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillMaster) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillMaster) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillNormal struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillNormal) User(Username string) SClientKillUser {
	return SClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c SClientKillNormal) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKillNormal) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillNormal) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillNormal) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillPubsub struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillPubsub) User(Username string) SClientKillUser {
	return SClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c SClientKillPubsub) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKillPubsub) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillPubsub) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillPubsub) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillSkipme struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillSkipme) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillSlave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillSlave) User(Username string) SClientKillUser {
	return SClientKillUser{cf: c.cf, cs: append(c.cs, "USER", Username)}
}

func (c SClientKillSlave) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKillSlave) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillSlave) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillSlave) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillUser struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientKillUser) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cf: c.cf, cs: append(c.cs, "ADDR", IpPort)}
}

func (c SClientKillUser) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillUser) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillUser) Build() SCompleted {
	return SCompleted(c)
}

type SClientList struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientList) Normal() SClientListNormal {
	return SClientListNormal{cf: c.cf, cs: append(c.cs, "normal")}
}

func (c SClientList) Master() SClientListMaster {
	return SClientListMaster{cf: c.cf, cs: append(c.cs, "master")}
}

func (c SClientList) Replica() SClientListReplica {
	return SClientListReplica{cf: c.cf, cs: append(c.cs, "replica")}
}

func (c SClientList) Pubsub() SClientListPubsub {
	return SClientListPubsub{cf: c.cf, cs: append(c.cs, "pubsub")}
}

func (c SClientList) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (b *SBuilder) ClientList() (c SClientList) {
	c.cs = append(b.get(), "CLIENT", "LIST")
	c.ks = initSlot
	return
}

type SClientListIdClientId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientListIdClientId) ClientId(ClientId ...int64) SClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClientListIdClientId{cf: c.cf, cs: c.cs}
}

func (c SClientListIdClientId) Build() SCompleted {
	return SCompleted(c)
}

type SClientListIdId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientListIdId) ClientId(ClientId ...int64) SClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClientListIdClientId{cf: c.cf, cs: c.cs}
}

type SClientListMaster struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientListMaster) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type SClientListNormal struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientListNormal) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type SClientListPubsub struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientListPubsub) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type SClientListReplica struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientListReplica) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type SClientNoEvict struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientNoEvict) On() SClientNoEvictEnabledOn {
	return SClientNoEvictEnabledOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c SClientNoEvict) Off() SClientNoEvictEnabledOff {
	return SClientNoEvictEnabledOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (b *SBuilder) ClientNoEvict() (c SClientNoEvict) {
	c.cs = append(b.get(), "CLIENT", "NO-EVICT")
	c.ks = initSlot
	return
}

type SClientNoEvictEnabledOff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientNoEvictEnabledOff) Build() SCompleted {
	return SCompleted(c)
}

type SClientNoEvictEnabledOn struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientNoEvictEnabledOn) Build() SCompleted {
	return SCompleted(c)
}

type SClientPause struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientPause) Timeout(Timeout int64) SClientPauseTimeout {
	return SClientPauseTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

func (b *SBuilder) ClientPause() (c SClientPause) {
	c.cs = append(b.get(), "CLIENT", "PAUSE")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SClientPauseModeAll struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientPauseModeAll) Build() SCompleted {
	return SCompleted(c)
}

type SClientPauseModeWrite struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientPauseModeWrite) Build() SCompleted {
	return SCompleted(c)
}

type SClientPauseTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientPauseTimeout) Write() SClientPauseModeWrite {
	return SClientPauseModeWrite{cf: c.cf, cs: append(c.cs, "WRITE")}
}

func (c SClientPauseTimeout) All() SClientPauseModeAll {
	return SClientPauseModeAll{cf: c.cf, cs: append(c.cs, "ALL")}
}

func (c SClientPauseTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientReply struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientReply) On() SClientReplyReplyModeOn {
	return SClientReplyReplyModeOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c SClientReply) Off() SClientReplyReplyModeOff {
	return SClientReplyReplyModeOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (c SClientReply) Skip() SClientReplyReplyModeSkip {
	return SClientReplyReplyModeSkip{cf: c.cf, cs: append(c.cs, "SKIP")}
}

func (b *SBuilder) ClientReply() (c SClientReply) {
	c.cs = append(b.get(), "CLIENT", "REPLY")
	c.ks = initSlot
	return
}

type SClientReplyReplyModeOff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientReplyReplyModeOff) Build() SCompleted {
	return SCompleted(c)
}

type SClientReplyReplyModeOn struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientReplyReplyModeOn) Build() SCompleted {
	return SCompleted(c)
}

type SClientReplyReplyModeSkip struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientReplyReplyModeSkip) Build() SCompleted {
	return SCompleted(c)
}

type SClientSetname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientSetname) ConnectionName(ConnectionName string) SClientSetnameConnectionName {
	return SClientSetnameConnectionName{cf: c.cf, cs: append(c.cs, ConnectionName)}
}

func (b *SBuilder) ClientSetname() (c SClientSetname) {
	c.cs = append(b.get(), "CLIENT", "SETNAME")
	c.ks = initSlot
	return
}

type SClientSetnameConnectionName struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientSetnameConnectionName) Build() SCompleted {
	return SCompleted(c)
}

type SClientTracking struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTracking) On() SClientTrackingStatusOn {
	return SClientTrackingStatusOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c SClientTracking) Off() SClientTrackingStatusOff {
	return SClientTrackingStatusOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (b *SBuilder) ClientTracking() (c SClientTracking) {
	c.cs = append(b.get(), "CLIENT", "TRACKING")
	c.ks = initSlot
	return
}

type SClientTrackingBcastBcast struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingBcastBcast) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c SClientTrackingBcastBcast) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c SClientTrackingBcastBcast) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingBcastBcast) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingNoloopNoloop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingNoloopNoloop) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingOptinOptin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingOptinOptin) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c SClientTrackingOptinOptin) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingOptinOptin) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingOptoutOptout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingOptoutOptout) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingOptoutOptout) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingPrefix struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingPrefix) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c SClientTrackingPrefix) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c SClientTrackingPrefix) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c SClientTrackingPrefix) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingPrefix) Prefix(Prefix ...string) SClientTrackingPrefix {
	return SClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c SClientTrackingPrefix) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingRedirect struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingRedirect) Prefix(Prefix ...string) SClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return SClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c SClientTrackingRedirect) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c SClientTrackingRedirect) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c SClientTrackingRedirect) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c SClientTrackingRedirect) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingRedirect) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingStatusOff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingStatusOff) Redirect(ClientId int64) SClientTrackingRedirect {
	return SClientTrackingRedirect{cf: c.cf, cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10))}
}

func (c SClientTrackingStatusOff) Prefix(Prefix ...string) SClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return SClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c SClientTrackingStatusOff) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c SClientTrackingStatusOff) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c SClientTrackingStatusOff) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c SClientTrackingStatusOff) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingStatusOff) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingStatusOn struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackingStatusOn) Redirect(ClientId int64) SClientTrackingRedirect {
	return SClientTrackingRedirect{cf: c.cf, cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10))}
}

func (c SClientTrackingStatusOn) Prefix(Prefix ...string) SClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return SClientTrackingPrefix{cf: c.cf, cs: append(c.cs, Prefix...)}
}

func (c SClientTrackingStatusOn) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cf: c.cf, cs: append(c.cs, "BCAST")}
}

func (c SClientTrackingStatusOn) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cf: c.cf, cs: append(c.cs, "OPTIN")}
}

func (c SClientTrackingStatusOn) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c SClientTrackingStatusOn) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingStatusOn) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackinginfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientTrackinginfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientTrackinginfo() (c SClientTrackinginfo) {
	c.cs = append(b.get(), "CLIENT", "TRACKINGINFO")
	c.ks = initSlot
	return
}

type SClientUnblock struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientUnblock) ClientId(ClientId int64) SClientUnblockClientId {
	return SClientUnblockClientId{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ClientId, 10))}
}

func (b *SBuilder) ClientUnblock() (c SClientUnblock) {
	c.cs = append(b.get(), "CLIENT", "UNBLOCK")
	c.ks = initSlot
	return
}

type SClientUnblockClientId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientUnblockClientId) Timeout() SClientUnblockUnblockTypeTimeout {
	return SClientUnblockUnblockTypeTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT")}
}

func (c SClientUnblockClientId) Error() SClientUnblockUnblockTypeError {
	return SClientUnblockUnblockTypeError{cf: c.cf, cs: append(c.cs, "ERROR")}
}

func (c SClientUnblockClientId) Build() SCompleted {
	return SCompleted(c)
}

type SClientUnblockUnblockTypeError struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientUnblockUnblockTypeError) Build() SCompleted {
	return SCompleted(c)
}

type SClientUnblockUnblockTypeTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientUnblockUnblockTypeTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientUnpause struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClientUnpause) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientUnpause() (c SClientUnpause) {
	c.cs = append(b.get(), "CLIENT", "UNPAUSE")
	c.ks = initSlot
	return
}

type SClusterAddslots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterAddslots) Slot(Slot ...int64) SClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterAddslotsSlot{cf: c.cf, cs: c.cs}
}

func (b *SBuilder) ClusterAddslots() (c SClusterAddslots) {
	c.cs = append(b.get(), "CLUSTER", "ADDSLOTS")
	c.ks = initSlot
	return
}

type SClusterAddslotsSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterAddslotsSlot) Slot(Slot ...int64) SClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterAddslotsSlot{cf: c.cf, cs: c.cs}
}

func (c SClusterAddslotsSlot) Build() SCompleted {
	return SCompleted(c)
}

type SClusterBumpepoch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterBumpepoch) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterBumpepoch() (c SClusterBumpepoch) {
	c.cs = append(b.get(), "CLUSTER", "BUMPEPOCH")
	c.ks = initSlot
	return
}

type SClusterCountFailureReports struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterCountFailureReports) NodeId(NodeId string) SClusterCountFailureReportsNodeId {
	return SClusterCountFailureReportsNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *SBuilder) ClusterCountFailureReports() (c SClusterCountFailureReports) {
	c.cs = append(b.get(), "CLUSTER", "COUNT-FAILURE-REPORTS")
	c.ks = initSlot
	return
}

type SClusterCountFailureReportsNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterCountFailureReportsNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterCountkeysinslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterCountkeysinslot) Slot(Slot int64) SClusterCountkeysinslotSlot {
	return SClusterCountkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *SBuilder) ClusterCountkeysinslot() (c SClusterCountkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "COUNTKEYSINSLOT")
	c.ks = initSlot
	return
}

type SClusterCountkeysinslotSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterCountkeysinslotSlot) Build() SCompleted {
	return SCompleted(c)
}

type SClusterDelslots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterDelslots) Slot(Slot ...int64) SClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterDelslotsSlot{cf: c.cf, cs: c.cs}
}

func (b *SBuilder) ClusterDelslots() (c SClusterDelslots) {
	c.cs = append(b.get(), "CLUSTER", "DELSLOTS")
	c.ks = initSlot
	return
}

type SClusterDelslotsSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterDelslotsSlot) Slot(Slot ...int64) SClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterDelslotsSlot{cf: c.cf, cs: c.cs}
}

func (c SClusterDelslotsSlot) Build() SCompleted {
	return SCompleted(c)
}

type SClusterFailover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterFailover) Force() SClusterFailoverOptionsForce {
	return SClusterFailoverOptionsForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c SClusterFailover) Takeover() SClusterFailoverOptionsTakeover {
	return SClusterFailoverOptionsTakeover{cf: c.cf, cs: append(c.cs, "TAKEOVER")}
}

func (c SClusterFailover) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterFailover() (c SClusterFailover) {
	c.cs = append(b.get(), "CLUSTER", "FAILOVER")
	c.ks = initSlot
	return
}

type SClusterFailoverOptionsForce struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterFailoverOptionsForce) Build() SCompleted {
	return SCompleted(c)
}

type SClusterFailoverOptionsTakeover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterFailoverOptionsTakeover) Build() SCompleted {
	return SCompleted(c)
}

type SClusterFlushslots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterFlushslots) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterFlushslots() (c SClusterFlushslots) {
	c.cs = append(b.get(), "CLUSTER", "FLUSHSLOTS")
	c.ks = initSlot
	return
}

type SClusterForget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterForget) NodeId(NodeId string) SClusterForgetNodeId {
	return SClusterForgetNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *SBuilder) ClusterForget() (c SClusterForget) {
	c.cs = append(b.get(), "CLUSTER", "FORGET")
	c.ks = initSlot
	return
}

type SClusterForgetNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterForgetNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterGetkeysinslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterGetkeysinslot) Slot(Slot int64) SClusterGetkeysinslotSlot {
	return SClusterGetkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *SBuilder) ClusterGetkeysinslot() (c SClusterGetkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "GETKEYSINSLOT")
	c.ks = initSlot
	return
}

type SClusterGetkeysinslotCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterGetkeysinslotCount) Build() SCompleted {
	return SCompleted(c)
}

type SClusterGetkeysinslotSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterGetkeysinslotSlot) Count(Count int64) SClusterGetkeysinslotCount {
	return SClusterGetkeysinslotCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type SClusterInfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterInfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterInfo() (c SClusterInfo) {
	c.cs = append(b.get(), "CLUSTER", "INFO")
	c.ks = initSlot
	return
}

type SClusterKeyslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterKeyslot) Key(Key string) SClusterKeyslotKey {
	return SClusterKeyslotKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) ClusterKeyslot() (c SClusterKeyslot) {
	c.cs = append(b.get(), "CLUSTER", "KEYSLOT")
	c.ks = initSlot
	return
}

type SClusterKeyslotKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterKeyslotKey) Build() SCompleted {
	return SCompleted(c)
}

type SClusterMeet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterMeet) Ip(Ip string) SClusterMeetIp {
	return SClusterMeetIp{cf: c.cf, cs: append(c.cs, Ip)}
}

func (b *SBuilder) ClusterMeet() (c SClusterMeet) {
	c.cs = append(b.get(), "CLUSTER", "MEET")
	c.ks = initSlot
	return
}

type SClusterMeetIp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterMeetIp) Port(Port int64) SClusterMeetPort {
	return SClusterMeetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type SClusterMeetPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterMeetPort) Build() SCompleted {
	return SCompleted(c)
}

type SClusterMyid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterMyid) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterMyid() (c SClusterMyid) {
	c.cs = append(b.get(), "CLUSTER", "MYID")
	c.ks = initSlot
	return
}

type SClusterNodes struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterNodes) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterNodes() (c SClusterNodes) {
	c.cs = append(b.get(), "CLUSTER", "NODES")
	c.ks = initSlot
	return
}

type SClusterReplicas struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterReplicas) NodeId(NodeId string) SClusterReplicasNodeId {
	return SClusterReplicasNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *SBuilder) ClusterReplicas() (c SClusterReplicas) {
	c.cs = append(b.get(), "CLUSTER", "REPLICAS")
	c.ks = initSlot
	return
}

type SClusterReplicasNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterReplicasNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterReplicate struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterReplicate) NodeId(NodeId string) SClusterReplicateNodeId {
	return SClusterReplicateNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *SBuilder) ClusterReplicate() (c SClusterReplicate) {
	c.cs = append(b.get(), "CLUSTER", "REPLICATE")
	c.ks = initSlot
	return
}

type SClusterReplicateNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterReplicateNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterReset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterReset) Hard() SClusterResetResetTypeHard {
	return SClusterResetResetTypeHard{cf: c.cf, cs: append(c.cs, "HARD")}
}

func (c SClusterReset) Soft() SClusterResetResetTypeSoft {
	return SClusterResetResetTypeSoft{cf: c.cf, cs: append(c.cs, "SOFT")}
}

func (c SClusterReset) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterReset() (c SClusterReset) {
	c.cs = append(b.get(), "CLUSTER", "RESET")
	c.ks = initSlot
	return
}

type SClusterResetResetTypeHard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterResetResetTypeHard) Build() SCompleted {
	return SCompleted(c)
}

type SClusterResetResetTypeSoft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterResetResetTypeSoft) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSaveconfig struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSaveconfig) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterSaveconfig() (c SClusterSaveconfig) {
	c.cs = append(b.get(), "CLUSTER", "SAVECONFIG")
	c.ks = initSlot
	return
}

type SClusterSetConfigEpoch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetConfigEpoch) ConfigEpoch(ConfigEpoch int64) SClusterSetConfigEpochConfigEpoch {
	return SClusterSetConfigEpochConfigEpoch{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ConfigEpoch, 10))}
}

func (b *SBuilder) ClusterSetConfigEpoch() (c SClusterSetConfigEpoch) {
	c.cs = append(b.get(), "CLUSTER", "SET-CONFIG-EPOCH")
	c.ks = initSlot
	return
}

type SClusterSetConfigEpochConfigEpoch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetConfigEpochConfigEpoch) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetslot) Slot(Slot int64) SClusterSetslotSlot {
	return SClusterSetslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *SBuilder) ClusterSetslot() (c SClusterSetslot) {
	c.cs = append(b.get(), "CLUSTER", "SETSLOT")
	c.ks = initSlot
	return
}

type SClusterSetslotNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetslotNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSlot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetslotSlot) Importing() SClusterSetslotSubcommandImporting {
	return SClusterSetslotSubcommandImporting{cf: c.cf, cs: append(c.cs, "IMPORTING")}
}

func (c SClusterSetslotSlot) Migrating() SClusterSetslotSubcommandMigrating {
	return SClusterSetslotSubcommandMigrating{cf: c.cf, cs: append(c.cs, "MIGRATING")}
}

func (c SClusterSetslotSlot) Stable() SClusterSetslotSubcommandStable {
	return SClusterSetslotSubcommandStable{cf: c.cf, cs: append(c.cs, "STABLE")}
}

func (c SClusterSetslotSlot) Node() SClusterSetslotSubcommandNode {
	return SClusterSetslotSubcommandNode{cf: c.cf, cs: append(c.cs, "NODE")}
}

type SClusterSetslotSubcommandImporting struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetslotSubcommandImporting) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandImporting) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandMigrating struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetslotSubcommandMigrating) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandMigrating) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandNode struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetslotSubcommandNode) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandNode) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandStable struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSetslotSubcommandStable) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandStable) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSlaves struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSlaves) NodeId(NodeId string) SClusterSlavesNodeId {
	return SClusterSlavesNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (b *SBuilder) ClusterSlaves() (c SClusterSlaves) {
	c.cs = append(b.get(), "CLUSTER", "SLAVES")
	c.ks = initSlot
	return
}

type SClusterSlavesNodeId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSlavesNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSlots struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SClusterSlots) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterSlots() (c SClusterSlots) {
	c.cs = append(b.get(), "CLUSTER", "SLOTS")
	c.ks = initSlot
	return
}

type SCommand struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCommand) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Command() (c SCommand) {
	c.cs = append(b.get(), "COMMAND")
	c.ks = initSlot
	return
}

type SCommandCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCommandCount) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) CommandCount() (c SCommandCount) {
	c.cs = append(b.get(), "COMMAND", "COUNT")
	c.ks = initSlot
	return
}

type SCommandGetkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCommandGetkeys) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) CommandGetkeys() (c SCommandGetkeys) {
	c.cs = append(b.get(), "COMMAND", "GETKEYS")
	c.ks = initSlot
	return
}

type SCommandInfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCommandInfo) CommandName(CommandName ...string) SCommandInfoCommandName {
	return SCommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (b *SBuilder) CommandInfo() (c SCommandInfo) {
	c.cs = append(b.get(), "COMMAND", "INFO")
	c.ks = initSlot
	return
}

type SCommandInfoCommandName struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCommandInfoCommandName) CommandName(CommandName ...string) SCommandInfoCommandName {
	return SCommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (c SCommandInfoCommandName) Build() SCompleted {
	return SCompleted(c)
}

type SConfigGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SConfigGet) Parameter(Parameter string) SConfigGetParameter {
	return SConfigGetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
}

func (b *SBuilder) ConfigGet() (c SConfigGet) {
	c.cs = append(b.get(), "CONFIG", "GET")
	c.ks = initSlot
	return
}

type SConfigGetParameter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SConfigGetParameter) Build() SCompleted {
	return SCompleted(c)
}

type SConfigResetstat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SConfigResetstat) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ConfigResetstat() (c SConfigResetstat) {
	c.cs = append(b.get(), "CONFIG", "RESETSTAT")
	c.ks = initSlot
	return
}

type SConfigRewrite struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SConfigRewrite) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ConfigRewrite() (c SConfigRewrite) {
	c.cs = append(b.get(), "CONFIG", "REWRITE")
	c.ks = initSlot
	return
}

type SConfigSet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SConfigSet) Parameter(Parameter string) SConfigSetParameter {
	return SConfigSetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
}

func (b *SBuilder) ConfigSet() (c SConfigSet) {
	c.cs = append(b.get(), "CONFIG", "SET")
	c.ks = initSlot
	return
}

type SConfigSetParameter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SConfigSetParameter) Value(Value string) SConfigSetValue {
	return SConfigSetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SConfigSetValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SConfigSetValue) Build() SCompleted {
	return SCompleted(c)
}

type SCopy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCopy) Source(Source string) SCopySource {
	s := slot(Source)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SCopySource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Copy() (c SCopy) {
	c.cs = append(b.get(), "COPY")
	c.ks = initSlot
	return
}

type SCopyDb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCopyDb) Replace() SCopyReplaceReplace {
	return SCopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c SCopyDb) Build() SCompleted {
	return SCompleted(c)
}

type SCopyDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCopyDestination) Db(DestinationDb int64) SCopyDb {
	return SCopyDb{cf: c.cf, cs: append(c.cs, "DB", strconv.FormatInt(DestinationDb, 10))}
}

func (c SCopyDestination) Replace() SCopyReplaceReplace {
	return SCopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c SCopyDestination) Build() SCompleted {
	return SCompleted(c)
}

type SCopyReplaceReplace struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCopyReplaceReplace) Build() SCompleted {
	return SCompleted(c)
}

type SCopySource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SCopySource) Destination(Destination string) SCopyDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SCopyDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SDbsize struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDbsize) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Dbsize() (c SDbsize) {
	c.cs = append(b.get(), "DBSIZE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SDebugObject struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDebugObject) Key(Key string) SDebugObjectKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SDebugObjectKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) DebugObject() (c SDebugObject) {
	c.cs = append(b.get(), "DEBUG", "OBJECT")
	c.ks = initSlot
	return
}

type SDebugObjectKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDebugObjectKey) Build() SCompleted {
	return SCompleted(c)
}

type SDebugSegfault struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDebugSegfault) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) DebugSegfault() (c SDebugSegfault) {
	c.cs = append(b.get(), "DEBUG", "SEGFAULT")
	c.ks = initSlot
	return
}

type SDecr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDecr) Key(Key string) SDecrKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SDecrKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Decr() (c SDecr) {
	c.cs = append(b.get(), "DECR")
	c.ks = initSlot
	return
}

type SDecrKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDecrKey) Build() SCompleted {
	return SCompleted(c)
}

type SDecrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDecrby) Key(Key string) SDecrbyKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SDecrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Decrby() (c SDecrby) {
	c.cs = append(b.get(), "DECRBY")
	c.ks = initSlot
	return
}

type SDecrbyDecrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDecrbyDecrement) Build() SCompleted {
	return SCompleted(c)
}

type SDecrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDecrbyKey) Decrement(Decrement int64) SDecrbyDecrement {
	return SDecrbyDecrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Decrement, 10))}
}

type SDel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDel) Key(Key ...string) SDelKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SDelKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Del() (c SDel) {
	c.cs = append(b.get(), "DEL")
	c.ks = initSlot
	return
}

type SDelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDelKey) Key(Key ...string) SDelKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SDelKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SDelKey) Build() SCompleted {
	return SCompleted(c)
}

type SDiscard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDiscard) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Discard() (c SDiscard) {
	c.cs = append(b.get(), "DISCARD")
	c.ks = initSlot
	return
}

type SDump struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDump) Key(Key string) SDumpKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SDumpKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Dump() (c SDump) {
	c.cs = append(b.get(), "DUMP")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SDumpKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SDumpKey) Build() SCompleted {
	return SCompleted(c)
}

type SEcho struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEcho) Message(Message string) SEchoMessage {
	return SEchoMessage{cf: c.cf, cs: append(c.cs, Message)}
}

func (b *SBuilder) Echo() (c SEcho) {
	c.cs = append(b.get(), "ECHO")
	c.ks = initSlot
	return
}

type SEchoMessage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEchoMessage) Build() SCompleted {
	return SCompleted(c)
}

type SEval struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEval) Script(Script string) SEvalScript {
	return SEvalScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *SBuilder) Eval() (c SEval) {
	c.cs = append(b.get(), "EVAL")
	c.ks = initSlot
	return
}

type SEvalArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalArg) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalKey) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalKey) Key(Key ...string) SEvalKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalKey) Build() SCompleted {
	return SCompleted(c)
}

type SEvalNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalNumkeys) Key(Key ...string) SEvalKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalNumkeys) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalNumkeys) Build() SCompleted {
	return SCompleted(c)
}

type SEvalRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalRo) Script(Script string) SEvalRoScript {
	return SEvalRoScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *SBuilder) EvalRo() (c SEvalRo) {
	c.cs = append(b.get(), "EVAL_RO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SEvalRoArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalRoArg) Arg(Arg ...string) SEvalRoArg {
	return SEvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalRoArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalRoKey) Arg(Arg ...string) SEvalRoArg {
	return SEvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalRoKey) Key(Key ...string) SEvalRoKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalRoNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalRoNumkeys) Key(Key ...string) SEvalRoKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalRoScript struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalRoScript) Numkeys(Numkeys int64) SEvalRoNumkeys {
	return SEvalRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SEvalScript struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalScript) Numkeys(Numkeys int64) SEvalNumkeys {
	return SEvalNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SEvalsha struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalsha) Sha1(Sha1 string) SEvalshaSha1 {
	return SEvalshaSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *SBuilder) Evalsha() (c SEvalsha) {
	c.cs = append(b.get(), "EVALSHA")
	c.ks = initSlot
	return
}

type SEvalshaArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaArg) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaKey) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaKey) Key(Key ...string) SEvalshaKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalshaKey) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaNumkeys) Key(Key ...string) SEvalshaKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalshaNumkeys) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaNumkeys) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaRo) Sha1(Sha1 string) SEvalshaRoSha1 {
	return SEvalshaRoSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *SBuilder) EvalshaRo() (c SEvalshaRo) {
	c.cs = append(b.get(), "EVALSHA_RO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SEvalshaRoArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaRoArg) Arg(Arg ...string) SEvalshaRoArg {
	return SEvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaRoArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaRoKey) Arg(Arg ...string) SEvalshaRoArg {
	return SEvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaRoKey) Key(Key ...string) SEvalshaRoKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalshaRoNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaRoNumkeys) Key(Key ...string) SEvalshaRoKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SEvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalshaRoSha1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaRoSha1) Numkeys(Numkeys int64) SEvalshaRoNumkeys {
	return SEvalshaRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SEvalshaSha1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SEvalshaSha1) Numkeys(Numkeys int64) SEvalshaNumkeys {
	return SEvalshaNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SExec struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExec) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Exec() (c SExec) {
	c.cs = append(b.get(), "EXEC")
	c.ks = initSlot
	return
}

type SExists struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExists) Key(Key ...string) SExistsKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Exists() (c SExists) {
	c.cs = append(b.get(), "EXISTS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SExistsKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExistsKey) Key(Key ...string) SExistsKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SExistsKey) Build() SCompleted {
	return SCompleted(c)
}

type SExpire struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpire) Key(Key string) SExpireKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SExpireKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Expire() (c SExpire) {
	c.cs = append(b.get(), "EXPIRE")
	c.ks = initSlot
	return
}

type SExpireConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireKey) Seconds(Seconds int64) SExpireSeconds {
	return SExpireSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SExpireSeconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireSeconds) Nx() SExpireConditionNx {
	return SExpireConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SExpireSeconds) Xx() SExpireConditionXx {
	return SExpireConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SExpireSeconds) Gt() SExpireConditionGt {
	return SExpireConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c SExpireSeconds) Lt() SExpireConditionLt {
	return SExpireConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c SExpireSeconds) Build() SCompleted {
	return SCompleted(c)
}

type SExpireat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireat) Key(Key string) SExpireatKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SExpireatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Expireat() (c SExpireat) {
	c.cs = append(b.get(), "EXPIREAT")
	c.ks = initSlot
	return
}

type SExpireatConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireatConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireatConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireatConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireatConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireatKey) Timestamp(Timestamp int64) SExpireatTimestamp {
	return SExpireatTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timestamp, 10))}
}

type SExpireatTimestamp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpireatTimestamp) Nx() SExpireatConditionNx {
	return SExpireatConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SExpireatTimestamp) Xx() SExpireatConditionXx {
	return SExpireatConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SExpireatTimestamp) Gt() SExpireatConditionGt {
	return SExpireatConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c SExpireatTimestamp) Lt() SExpireatConditionLt {
	return SExpireatConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c SExpireatTimestamp) Build() SCompleted {
	return SCompleted(c)
}

type SExpiretime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpiretime) Key(Key string) SExpiretimeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SExpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Expiretime() (c SExpiretime) {
	c.cs = append(b.get(), "EXPIRETIME")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SExpiretimeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SExpiretimeKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SExpiretimeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SFailover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFailover) To() SFailoverTargetTo {
	return SFailoverTargetTo{cf: c.cf, cs: append(c.cs, "TO")}
}

func (c SFailover) Abort() SFailoverAbort {
	return SFailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c SFailover) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (b *SBuilder) Failover() (c SFailover) {
	c.cs = append(b.get(), "FAILOVER")
	c.ks = initSlot
	return
}

type SFailoverAbort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFailoverAbort) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c SFailoverAbort) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetForce struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFailoverTargetForce) Abort() SFailoverAbort {
	return SFailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c SFailoverTargetForce) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c SFailoverTargetForce) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFailoverTargetHost) Port(Port int64) SFailoverTargetPort {
	return SFailoverTargetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type SFailoverTargetPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFailoverTargetPort) Force() SFailoverTargetForce {
	return SFailoverTargetForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c SFailoverTargetPort) Abort() SFailoverAbort {
	return SFailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c SFailoverTargetPort) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c SFailoverTargetPort) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetTo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFailoverTargetTo) Host(Host string) SFailoverTargetHost {
	return SFailoverTargetHost{cf: c.cf, cs: append(c.cs, Host)}
}

type SFailoverTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFailoverTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SFlushall struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFlushall) Async() SFlushallAsyncAsync {
	return SFlushallAsyncAsync{cf: c.cf, cs: append(c.cs, "ASYNC")}
}

func (c SFlushall) Sync() SFlushallAsyncSync {
	return SFlushallAsyncSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c SFlushall) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Flushall() (c SFlushall) {
	c.cs = append(b.get(), "FLUSHALL")
	c.ks = initSlot
	return
}

type SFlushallAsyncAsync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFlushallAsyncAsync) Build() SCompleted {
	return SCompleted(c)
}

type SFlushallAsyncSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFlushallAsyncSync) Build() SCompleted {
	return SCompleted(c)
}

type SFlushdb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFlushdb) Async() SFlushdbAsyncAsync {
	return SFlushdbAsyncAsync{cf: c.cf, cs: append(c.cs, "ASYNC")}
}

func (c SFlushdb) Sync() SFlushdbAsyncSync {
	return SFlushdbAsyncSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c SFlushdb) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Flushdb() (c SFlushdb) {
	c.cs = append(b.get(), "FLUSHDB")
	c.ks = initSlot
	return
}

type SFlushdbAsyncAsync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFlushdbAsyncAsync) Build() SCompleted {
	return SCompleted(c)
}

type SFlushdbAsyncSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SFlushdbAsyncSync) Build() SCompleted {
	return SCompleted(c)
}

type SGeoadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoadd) Key(Key string) SGeoaddKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geoadd() (c SGeoadd) {
	c.cs = append(b.get(), "GEOADD")
	c.ks = initSlot
	return
}

type SGeoaddChangeCh struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoaddChangeCh) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type SGeoaddConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoaddConditionNx) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SGeoaddConditionNx) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type SGeoaddConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoaddConditionXx) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SGeoaddConditionXx) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type SGeoaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoaddKey) Nx() SGeoaddConditionNx {
	return SGeoaddConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SGeoaddKey) Xx() SGeoaddConditionXx {
	return SGeoaddConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SGeoaddKey) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SGeoaddKey) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type SGeoaddLongitudeLatitudeMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member)}
}

func (c SGeoaddLongitudeLatitudeMember) Build() SCompleted {
	return SCompleted(c)
}

type SGeodist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodist) Key(Key string) SGeodistKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeodistKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geodist() (c SGeodist) {
	c.cs = append(b.get(), "GEODIST")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGeodistKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodistKey) Member1(Member1 string) SGeodistMember1 {
	return SGeodistMember1{cf: c.cf, cs: append(c.cs, Member1)}
}

func (c SGeodistKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistMember1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodistMember1) Member2(Member2 string) SGeodistMember2 {
	return SGeodistMember2{cf: c.cf, cs: append(c.cs, Member2)}
}

func (c SGeodistMember1) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistMember2 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodistMember2) M() SGeodistUnitM {
	return SGeodistUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeodistMember2) Km() SGeodistUnitKm {
	return SGeodistUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeodistMember2) Ft() SGeodistUnitFt {
	return SGeodistUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeodistMember2) Mi() SGeodistUnitMi {
	return SGeodistUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c SGeodistMember2) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistMember2) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodistUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodistUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodistUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeodistUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeohash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeohash) Key(Key string) SGeohashKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeohashKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geohash() (c SGeohash) {
	c.cs = append(b.get(), "GEOHASH")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGeohashKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeohashKey) Member(Member ...string) SGeohashMember {
	return SGeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SGeohashKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeohashMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeohashMember) Member(Member ...string) SGeohashMember {
	return SGeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SGeohashMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeohashMember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeopos struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeopos) Key(Key string) SGeoposKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geopos() (c SGeopos) {
	c.cs = append(b.get(), "GEOPOS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGeoposKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoposKey) Member(Member ...string) SGeoposMember {
	return SGeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SGeoposKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoposMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoposMember) Member(Member ...string) SGeoposMember {
	return SGeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SGeoposMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoposMember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradius) Key(Key string) SGeoradiusKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Georadius() (c SGeoradius) {
	c.cs = append(b.get(), "GEORADIUS")
	c.ks = initSlot
	return
}

type SGeoradiusCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusCountAnyAny) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusCountAnyAny) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusCountAnyAny) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusCountAnyAny) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusCountCount) Any() SGeoradiusCountAnyAny {
	return SGeoradiusCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c SGeoradiusCountCount) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusCountCount) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusCountCount) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusCountCount) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusKey) Longitude(Longitude float64) SGeoradiusLongitude {
	return SGeoradiusLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type SGeoradiusLatitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusLatitude) Radius(Radius float64) SGeoradiusRadius {
	return SGeoradiusRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type SGeoradiusLongitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusLongitude) Latitude(Latitude float64) SGeoradiusLatitude {
	return SGeoradiusLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type SGeoradiusOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusOrderAsc) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusOrderAsc) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusOrderDesc) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusOrderDesc) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRadius) M() SGeoradiusUnitM {
	return SGeoradiusUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeoradiusRadius) Km() SGeoradiusUnitKm {
	return SGeoradiusUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeoradiusRadius) Ft() SGeoradiusUnitFt {
	return SGeoradiusUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeoradiusRadius) Mi() SGeoradiusUnitMi {
	return SGeoradiusUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type SGeoradiusRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRo) Key(Key string) SGeoradiusRoKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) GeoradiusRo() (c SGeoradiusRo) {
	c.cs = append(b.get(), "GEORADIUS_RO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGeoradiusRoCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoCountAnyAny) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoCountAnyAny) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoCountAnyAny) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoCountCount) Any() SGeoradiusRoCountAnyAny {
	return SGeoradiusRoCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c SGeoradiusRoCountCount) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoCountCount) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoCountCount) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoKey) Longitude(Longitude float64) SGeoradiusRoLongitude {
	return SGeoradiusRoLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

func (c SGeoradiusRoKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoLatitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoLatitude) Radius(Radius float64) SGeoradiusRoRadius {
	return SGeoradiusRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeoradiusRoLatitude) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoLongitude struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoLongitude) Latitude(Latitude float64) SGeoradiusRoLatitude {
	return SGeoradiusRoLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c SGeoradiusRoLongitude) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoOrderAsc) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoOrderDesc) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoRadius) M() SGeoradiusRoUnitM {
	return SGeoradiusRoUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeoradiusRoRadius) Km() SGeoradiusRoUnitKm {
	return SGeoradiusRoUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeoradiusRoRadius) Ft() SGeoradiusRoUnitFt {
	return SGeoradiusRoUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeoradiusRoRadius) Mi() SGeoradiusRoUnitMi {
	return SGeoradiusRoUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c SGeoradiusRoRadius) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoStoredist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoStoredist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoUnitFt) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusRoUnitFt) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusRoUnitFt) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusRoUnitFt) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusRoUnitFt) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoUnitFt) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoUnitFt) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoUnitKm) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusRoUnitKm) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusRoUnitKm) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusRoUnitKm) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusRoUnitKm) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoUnitKm) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoUnitKm) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoUnitM) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusRoUnitM) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusRoUnitM) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusRoUnitM) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusRoUnitM) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoUnitM) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoUnitM) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoUnitMi) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusRoUnitMi) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusRoUnitMi) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusRoUnitMi) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusRoUnitMi) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoUnitMi) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoUnitMi) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoWithcoordWithcoord) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusRoWithcoordWithcoord) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusRoWithcoordWithcoord) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusRoWithcoordWithcoord) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoWithcoordWithcoord) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoWithcoordWithcoord) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoWithdistWithdist) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusRoWithdistWithdist) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusRoWithdistWithdist) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoWithdistWithdist) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoWithdistWithdist) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusRoWithhashWithhash) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusRoWithhashWithhash) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoWithhashWithhash) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoWithhashWithhash) Storedist(Key string) SGeoradiusRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoWithhashWithhash) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusStore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusStore) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusStore) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusStoredist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusUnitFt) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusUnitFt) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusUnitFt) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusUnitFt) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusUnitFt) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusUnitFt) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusUnitFt) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitFt) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusUnitKm) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusUnitKm) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusUnitKm) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusUnitKm) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusUnitKm) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusUnitKm) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusUnitKm) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitKm) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusUnitM) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusUnitM) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusUnitM) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusUnitM) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusUnitM) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusUnitM) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusUnitM) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitM) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusUnitMi) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusUnitMi) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusUnitMi) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusUnitMi) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusUnitMi) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusUnitMi) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusUnitMi) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitMi) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusWithcoordWithcoord) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusWithcoordWithcoord) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusWithcoordWithcoord) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusWithcoordWithcoord) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusWithcoordWithcoord) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusWithcoordWithcoord) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusWithcoordWithcoord) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusWithdistWithdist) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusWithdistWithdist) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusWithdistWithdist) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusWithdistWithdist) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusWithdistWithdist) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusWithdistWithdist) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusWithhashWithhash) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusWithhashWithhash) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusWithhashWithhash) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusWithhashWithhash) Store(Key string) SGeoradiusStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusWithhashWithhash) Storedist(Key string) SGeoradiusStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusbymember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymember) Key(Key string) SGeoradiusbymemberKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Georadiusbymember() (c SGeoradiusbymember) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER")
	c.ks = initSlot
	return
}

type SGeoradiusbymemberCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberCountAnyAny) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberCountAnyAny) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberCountAnyAny) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberCountAnyAny) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberCountCount) Any() SGeoradiusbymemberCountAnyAny {
	return SGeoradiusbymemberCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c SGeoradiusbymemberCountCount) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberCountCount) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberCountCount) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberCountCount) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberKey) Member(Member string) SGeoradiusbymemberMember {
	return SGeoradiusbymemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SGeoradiusbymemberMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberMember) Radius(Radius float64) SGeoradiusbymemberRadius {
	return SGeoradiusbymemberRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type SGeoradiusbymemberOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberOrderAsc) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberOrderAsc) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberOrderDesc) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberOrderDesc) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRadius) M() SGeoradiusbymemberUnitM {
	return SGeoradiusbymemberUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeoradiusbymemberRadius) Km() SGeoradiusbymemberUnitKm {
	return SGeoradiusbymemberUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeoradiusbymemberRadius) Ft() SGeoradiusbymemberUnitFt {
	return SGeoradiusbymemberUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeoradiusbymemberRadius) Mi() SGeoradiusbymemberUnitMi {
	return SGeoradiusbymemberUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type SGeoradiusbymemberRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRo) Key(Key string) SGeoradiusbymemberRoKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) GeoradiusbymemberRo() (c SGeoradiusbymemberRo) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER_RO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGeoradiusbymemberRoCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoCountAnyAny) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoCountAnyAny) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoCountAnyAny) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoCountCount) Any() SGeoradiusbymemberRoCountAnyAny {
	return SGeoradiusbymemberRoCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c SGeoradiusbymemberRoCountCount) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoCountCount) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoCountCount) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoKey) Member(Member string) SGeoradiusbymemberRoMember {
	return SGeoradiusbymemberRoMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c SGeoradiusbymemberRoKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoMember) Radius(Radius float64) SGeoradiusbymemberRoRadius {
	return SGeoradiusbymemberRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeoradiusbymemberRoMember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoOrderAsc) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoOrderDesc) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoRadius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoRadius) M() SGeoradiusbymemberRoUnitM {
	return SGeoradiusbymemberRoUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeoradiusbymemberRoRadius) Km() SGeoradiusbymemberRoUnitKm {
	return SGeoradiusbymemberRoUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeoradiusbymemberRoRadius) Ft() SGeoradiusbymemberRoUnitFt {
	return SGeoradiusbymemberRoUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeoradiusbymemberRoRadius) Mi() SGeoradiusbymemberRoUnitMi {
	return SGeoradiusbymemberRoUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c SGeoradiusbymemberRoRadius) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoStoredist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoStoredist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoUnitFt) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberRoUnitFt) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberRoUnitFt) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberRoUnitFt) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberRoUnitFt) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoUnitFt) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoUnitFt) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoUnitKm) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberRoUnitKm) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberRoUnitKm) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberRoUnitKm) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberRoUnitKm) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoUnitKm) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoUnitKm) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoUnitM) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberRoUnitM) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberRoUnitM) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberRoUnitM) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberRoUnitM) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoUnitM) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoUnitM) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoUnitMi) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberRoUnitMi) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberRoUnitMi) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberRoUnitMi) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberRoUnitMi) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoUnitMi) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoUnitMi) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoWithdistWithdist) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberRoWithhashWithhash) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberStore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberStore) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberStore) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberStoredist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberUnitFt) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberUnitFt) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberUnitFt) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberUnitFt) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberUnitFt) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberUnitFt) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberUnitFt) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitFt) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusbymemberUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberUnitKm) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberUnitKm) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberUnitKm) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberUnitKm) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberUnitKm) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberUnitKm) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberUnitKm) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitKm) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusbymemberUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberUnitM) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberUnitM) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberUnitM) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberUnitM) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberUnitM) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberUnitM) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberUnitM) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitM) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusbymemberUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberUnitMi) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeoradiusbymemberUnitMi) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberUnitMi) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberUnitMi) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberUnitMi) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberUnitMi) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberUnitMi) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitMi) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusbymemberWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberWithcoordWithcoord) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusbymemberWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberWithdistWithdist) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeoradiusbymemberWithdistWithdist) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberWithdistWithdist) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberWithdistWithdist) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberWithdistWithdist) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberWithdistWithdist) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeoradiusbymemberWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeoradiusbymemberWithhashWithhash) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeoradiusbymemberWithhashWithhash) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberWithhashWithhash) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberWithhashWithhash) Store(Key string) SGeoradiusbymemberStore {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberWithhashWithhash) Storedist(Key string) SGeoradiusbymemberStoredist {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

type SGeosearch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearch) Key(Key string) SGeosearchKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeosearchKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geosearch() (c SGeosearch) {
	c.cs = append(b.get(), "GEOSEARCH")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGeosearchBoxBybox struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchBoxBybox) Height(Height float64) SGeosearchBoxHeight {
	return SGeosearchBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

func (c SGeosearchBoxBybox) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxHeight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchBoxHeight) M() SGeosearchBoxUnitM {
	return SGeosearchBoxUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeosearchBoxHeight) Km() SGeosearchBoxUnitKm {
	return SGeosearchBoxUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeosearchBoxHeight) Ft() SGeosearchBoxUnitFt {
	return SGeosearchBoxUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeosearchBoxHeight) Mi() SGeosearchBoxUnitMi {
	return SGeosearchBoxUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c SGeosearchBoxHeight) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchBoxUnitFt) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchBoxUnitFt) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchBoxUnitFt) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchBoxUnitFt) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchBoxUnitFt) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchBoxUnitFt) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchBoxUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchBoxUnitKm) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchBoxUnitKm) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchBoxUnitKm) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchBoxUnitKm) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchBoxUnitKm) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchBoxUnitKm) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchBoxUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchBoxUnitM) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchBoxUnitM) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchBoxUnitM) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchBoxUnitM) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchBoxUnitM) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchBoxUnitM) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchBoxUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchBoxUnitMi) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchBoxUnitMi) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchBoxUnitMi) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchBoxUnitMi) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchBoxUnitMi) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchBoxUnitMi) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchBoxUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleByradius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchCircleByradius) M() SGeosearchCircleUnitM {
	return SGeosearchCircleUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeosearchCircleByradius) Km() SGeosearchCircleUnitKm {
	return SGeosearchCircleUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeosearchCircleByradius) Ft() SGeosearchCircleUnitFt {
	return SGeosearchCircleUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeosearchCircleByradius) Mi() SGeosearchCircleUnitMi {
	return SGeosearchCircleUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

func (c SGeosearchCircleByradius) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchCircleUnitFt) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchCircleUnitFt) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchCircleUnitFt) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchCircleUnitFt) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchCircleUnitFt) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchCircleUnitFt) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchCircleUnitFt) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchCircleUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchCircleUnitKm) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchCircleUnitKm) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchCircleUnitKm) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchCircleUnitKm) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchCircleUnitKm) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchCircleUnitKm) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchCircleUnitKm) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchCircleUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchCircleUnitM) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchCircleUnitM) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchCircleUnitM) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchCircleUnitM) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchCircleUnitM) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchCircleUnitM) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchCircleUnitM) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchCircleUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchCircleUnitMi) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchCircleUnitMi) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchCircleUnitMi) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchCircleUnitMi) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchCircleUnitMi) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchCircleUnitMi) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchCircleUnitMi) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchCircleUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchCountAnyAny) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchCountAnyAny) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchCountAnyAny) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchCountCount) Any() SGeosearchCountAnyAny {
	return SGeosearchCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c SGeosearchCountCount) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchCountCount) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchCountCount) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchFromlonlat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchFromlonlat) Byradius(Radius float64) SGeosearchCircleByradius {
	return SGeosearchCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeosearchFromlonlat) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchFromlonlat) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchFromlonlat) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchFromlonlat) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchFromlonlat) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchFromlonlat) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchFromlonlat) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchFromlonlat) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchFrommember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchFrommember) Fromlonlat(Longitude float64, Latitude float64) SGeosearchFromlonlat {
	return SGeosearchFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c SGeosearchFrommember) Byradius(Radius float64) SGeosearchCircleByradius {
	return SGeosearchCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeosearchFrommember) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchFrommember) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchFrommember) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchFrommember) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchFrommember) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchFrommember) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchFrommember) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchFrommember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchKey) Frommember(Member string) SGeosearchFrommember {
	return SGeosearchFrommember{cf: c.cf, cs: append(c.cs, "FROMMEMBER", Member)}
}

func (c SGeosearchKey) Fromlonlat(Longitude float64, Latitude float64) SGeosearchFromlonlat {
	return SGeosearchFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c SGeosearchKey) Byradius(Radius float64) SGeosearchCircleByradius {
	return SGeosearchCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeosearchKey) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchKey) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchKey) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchKey) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchKey) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchKey) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchKey) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchOrderAsc) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchOrderAsc) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchOrderAsc) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchOrderAsc) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchOrderDesc) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchOrderDesc) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cf: c.cf, cs: append(c.cs, "WITHCOORD")}
}

func (c SGeosearchOrderDesc) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchOrderDesc) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchWithcoordWithcoord struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchWithcoordWithcoord) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cf: c.cf, cs: append(c.cs, "WITHDIST")}
}

func (c SGeosearchWithcoordWithcoord) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchWithdistWithdist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchWithdistWithdist) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
}

func (c SGeosearchWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchWithhashWithhash struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchWithhashWithhash) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstore) Destination(Destination string) SGeosearchstoreDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeosearchstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Geosearchstore() (c SGeosearchstore) {
	c.cs = append(b.get(), "GEOSEARCHSTORE")
	c.ks = initSlot
	return
}

type SGeosearchstoreBoxBybox struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreBoxBybox) Height(Height float64) SGeosearchstoreBoxHeight {
	return SGeosearchstoreBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type SGeosearchstoreBoxHeight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreBoxHeight) M() SGeosearchstoreBoxUnitM {
	return SGeosearchstoreBoxUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeosearchstoreBoxHeight) Km() SGeosearchstoreBoxUnitKm {
	return SGeosearchstoreBoxUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeosearchstoreBoxHeight) Ft() SGeosearchstoreBoxUnitFt {
	return SGeosearchstoreBoxUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeosearchstoreBoxHeight) Mi() SGeosearchstoreBoxUnitMi {
	return SGeosearchstoreBoxUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type SGeosearchstoreBoxUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreBoxUnitFt) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreBoxUnitFt) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreBoxUnitFt) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreBoxUnitFt) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreBoxUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreBoxUnitKm) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreBoxUnitKm) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreBoxUnitKm) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreBoxUnitKm) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreBoxUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreBoxUnitM) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreBoxUnitM) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreBoxUnitM) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreBoxUnitM) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreBoxUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreBoxUnitMi) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreBoxUnitMi) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreBoxUnitMi) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreBoxUnitMi) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreCircleByradius struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreCircleByradius) M() SGeosearchstoreCircleUnitM {
	return SGeosearchstoreCircleUnitM{cf: c.cf, cs: append(c.cs, "m")}
}

func (c SGeosearchstoreCircleByradius) Km() SGeosearchstoreCircleUnitKm {
	return SGeosearchstoreCircleUnitKm{cf: c.cf, cs: append(c.cs, "km")}
}

func (c SGeosearchstoreCircleByradius) Ft() SGeosearchstoreCircleUnitFt {
	return SGeosearchstoreCircleUnitFt{cf: c.cf, cs: append(c.cs, "ft")}
}

func (c SGeosearchstoreCircleByradius) Mi() SGeosearchstoreCircleUnitMi {
	return SGeosearchstoreCircleUnitMi{cf: c.cf, cs: append(c.cs, "mi")}
}

type SGeosearchstoreCircleUnitFt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreCircleUnitFt) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchstoreCircleUnitFt) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreCircleUnitFt) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreCircleUnitFt) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreCircleUnitFt) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreCircleUnitKm struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreCircleUnitKm) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchstoreCircleUnitKm) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreCircleUnitKm) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreCircleUnitKm) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreCircleUnitKm) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreCircleUnitM struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreCircleUnitM) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchstoreCircleUnitM) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreCircleUnitM) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreCircleUnitM) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreCircleUnitM) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreCircleUnitMi struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreCircleUnitMi) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchstoreCircleUnitMi) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreCircleUnitMi) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreCircleUnitMi) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreCircleUnitMi) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreCountAnyAny struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreCountAnyAny) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c SGeosearchstoreCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCountCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreCountCount) Any() SGeosearchstoreCountAnyAny {
	return SGeosearchstoreCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c SGeosearchstoreCountCount) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c SGeosearchstoreCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreDestination) Source(Source string) SGeosearchstoreSource {
	s := slot(Source)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGeosearchstoreSource{cf: c.cf, cs: append(c.cs, Source)}
}

type SGeosearchstoreFromlonlat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreFromlonlat) Byradius(Radius float64) SGeosearchstoreCircleByradius {
	return SGeosearchstoreCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeosearchstoreFromlonlat) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchstoreFromlonlat) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreFromlonlat) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreFromlonlat) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreFromlonlat) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreFrommember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreFrommember) Fromlonlat(Longitude float64, Latitude float64) SGeosearchstoreFromlonlat {
	return SGeosearchstoreFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c SGeosearchstoreFrommember) Byradius(Radius float64) SGeosearchstoreCircleByradius {
	return SGeosearchstoreCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeosearchstoreFrommember) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchstoreFrommember) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreFrommember) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreFrommember) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreFrommember) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreOrderAsc) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreOrderAsc) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreOrderDesc) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreOrderDesc) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreSource) Frommember(Member string) SGeosearchstoreFrommember {
	return SGeosearchstoreFrommember{cf: c.cf, cs: append(c.cs, "FROMMEMBER", Member)}
}

func (c SGeosearchstoreSource) Fromlonlat(Longitude float64, Latitude float64) SGeosearchstoreFromlonlat {
	return SGeosearchstoreFromlonlat{cf: c.cf, cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c SGeosearchstoreSource) Byradius(Radius float64) SGeosearchstoreCircleByradius {
	return SGeosearchstoreCircleByradius{cf: c.cf, cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c SGeosearchstoreSource) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cf: c.cf, cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c SGeosearchstoreSource) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeosearchstoreSource) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeosearchstoreSource) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreSource) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type SGeosearchstoreStoredistStoredist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGeosearchstoreStoredistStoredist) Build() SCompleted {
	return SCompleted(c)
}

type SGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGet) Key(Key string) SGetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Get() (c SGet) {
	c.cs = append(b.get(), "GET")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SGetKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGetbit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetbit) Key(Key string) SGetbitKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Getbit() (c SGetbit) {
	c.cs = append(b.get(), "GETBIT")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGetbitKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetbitKey) Offset(Offset int64) SGetbitOffset {
	return SGetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

func (c SGetbitKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGetbitOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetbitOffset) Build() SCompleted {
	return SCompleted(c)
}

func (c SGetbitOffset) Cache() SCacheable {
	return SCacheable(c)
}

type SGetdel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetdel) Key(Key string) SGetdelKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGetdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Getdel() (c SGetdel) {
	c.cs = append(b.get(), "GETDEL")
	c.ks = initSlot
	return
}

type SGetdelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetdelKey) Build() SCompleted {
	return SCompleted(c)
}

type SGetex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetex) Key(Key string) SGetexKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Getex() (c SGetex) {
	c.cs = append(b.get(), "GETEX")
	c.ks = initSlot
	return
}

type SGetexExpirationEx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetexExpirationEx) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationExat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetexExpirationExat) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationPersist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetexExpirationPersist) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationPx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetexExpirationPx) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationPxat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetexExpirationPxat) Build() SCompleted {
	return SCompleted(c)
}

type SGetexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetexKey) Ex(Seconds int64) SGetexExpirationEx {
	return SGetexExpirationEx{cf: c.cf, cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10))}
}

func (c SGetexKey) Px(Milliseconds int64) SGetexExpirationPx {
	return SGetexExpirationPx{cf: c.cf, cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10))}
}

func (c SGetexKey) Exat(Timestamp int64) SGetexExpirationExat {
	return SGetexExpirationExat{cf: c.cf, cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10))}
}

func (c SGetexKey) Pxat(Millisecondstimestamp int64) SGetexExpirationPxat {
	return SGetexExpirationPxat{cf: c.cf, cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10))}
}

func (c SGetexKey) Persist() SGetexExpirationPersist {
	return SGetexExpirationPersist{cf: c.cf, cs: append(c.cs, "PERSIST")}
}

func (c SGetexKey) Build() SCompleted {
	return SCompleted(c)
}

type SGetrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetrange) Key(Key string) SGetrangeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Getrange() (c SGetrange) {
	c.cs = append(b.get(), "GETRANGE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SGetrangeEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetrangeEnd) Build() SCompleted {
	return SCompleted(c)
}

func (c SGetrangeEnd) Cache() SCacheable {
	return SCacheable(c)
}

type SGetrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetrangeKey) Start(Start int64) SGetrangeStart {
	return SGetrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c SGetrangeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGetrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetrangeStart) End(End int64) SGetrangeEnd {
	return SGetrangeEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c SGetrangeStart) Cache() SCacheable {
	return SCacheable(c)
}

type SGetset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetset) Key(Key string) SGetsetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SGetsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Getset() (c SGetset) {
	c.cs = append(b.get(), "GETSET")
	c.ks = initSlot
	return
}

type SGetsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetsetKey) Value(Value string) SGetsetValue {
	return SGetsetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SGetsetValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SGetsetValue) Build() SCompleted {
	return SCompleted(c)
}

type SHdel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHdel) Key(Key string) SHdelKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hdel() (c SHdel) {
	c.cs = append(b.get(), "HDEL")
	c.ks = initSlot
	return
}

type SHdelField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHdelField) Field(Field ...string) SHdelField {
	return SHdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c SHdelField) Build() SCompleted {
	return SCompleted(c)
}

type SHdelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHdelKey) Field(Field ...string) SHdelField {
	return SHdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

type SHello struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHello) Protover(Protover int64) SHelloArgumentsProtover {
	return SHelloArgumentsProtover{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Protover, 10))}
}

func (b *SBuilder) Hello() (c SHello) {
	c.cs = append(b.get(), "HELLO")
	c.ks = initSlot
	return
}

type SHelloArgumentsAuth struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHelloArgumentsAuth) Setname(Clientname string) SHelloArgumentsSetname {
	return SHelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
}

func (c SHelloArgumentsAuth) Build() SCompleted {
	return SCompleted(c)
}

type SHelloArgumentsProtover struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHelloArgumentsProtover) Auth(Username string, Password string) SHelloArgumentsAuth {
	return SHelloArgumentsAuth{cf: c.cf, cs: append(c.cs, "AUTH", Username, Password)}
}

func (c SHelloArgumentsProtover) Setname(Clientname string) SHelloArgumentsSetname {
	return SHelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
}

func (c SHelloArgumentsProtover) Build() SCompleted {
	return SCompleted(c)
}

type SHelloArgumentsSetname struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHelloArgumentsSetname) Build() SCompleted {
	return SCompleted(c)
}

type SHexists struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHexists) Key(Key string) SHexistsKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHexistsKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hexists() (c SHexists) {
	c.cs = append(b.get(), "HEXISTS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHexistsField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHexistsField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHexistsField) Cache() SCacheable {
	return SCacheable(c)
}

type SHexistsKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHexistsKey) Field(Field string) SHexistsField {
	return SHexistsField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c SHexistsKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHget) Key(Key string) SHgetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHgetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hget() (c SHget) {
	c.cs = append(b.get(), "HGET")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHgetField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHgetField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHgetField) Cache() SCacheable {
	return SCacheable(c)
}

type SHgetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHgetKey) Field(Field string) SHgetField {
	return SHgetField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c SHgetKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHgetall struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHgetall) Key(Key string) SHgetallKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHgetallKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hgetall() (c SHgetall) {
	c.cs = append(b.get(), "HGETALL")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHgetallKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHgetallKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHgetallKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHincrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrby) Key(Key string) SHincrbyKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hincrby() (c SHincrby) {
	c.cs = append(b.get(), "HINCRBY")
	c.ks = initSlot
	return
}

type SHincrbyField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrbyField) Increment(Increment int64) SHincrbyIncrement {
	return SHincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type SHincrbyIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrbyIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SHincrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrbyKey) Field(Field string) SHincrbyField {
	return SHincrbyField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHincrbyfloat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrbyfloat) Key(Key string) SHincrbyfloatKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHincrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hincrbyfloat() (c SHincrbyfloat) {
	c.cs = append(b.get(), "HINCRBYFLOAT")
	c.ks = initSlot
	return
}

type SHincrbyfloatField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrbyfloatField) Increment(Increment float64) SHincrbyfloatIncrement {
	return SHincrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type SHincrbyfloatIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrbyfloatIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SHincrbyfloatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHincrbyfloatKey) Field(Field string) SHincrbyfloatField {
	return SHincrbyfloatField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHkeys) Key(Key string) SHkeysKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHkeysKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hkeys() (c SHkeys) {
	c.cs = append(b.get(), "HKEYS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHkeysKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHkeysKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHkeysKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHlen) Key(Key string) SHlenKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hlen() (c SHlen) {
	c.cs = append(b.get(), "HLEN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHlenKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHlenKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHmget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHmget) Key(Key string) SHmgetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHmgetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hmget() (c SHmget) {
	c.cs = append(b.get(), "HMGET")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHmgetField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHmgetField) Field(Field ...string) SHmgetField {
	return SHmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c SHmgetField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHmgetField) Cache() SCacheable {
	return SCacheable(c)
}

type SHmgetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHmgetKey) Field(Field ...string) SHmgetField {
	return SHmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c SHmgetKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHmset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHmset) Key(Key string) SHmsetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHmsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hmset() (c SHmset) {
	c.cs = append(b.get(), "HMSET")
	c.ks = initSlot
	return
}

type SHmsetFieldValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHmsetFieldValue) FieldValue(Field string, Value string) SHmsetFieldValue {
	return SHmsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c SHmsetFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SHmsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHmsetKey) FieldValue() SHmsetFieldValue {
	return SHmsetFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type SHrandfield struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHrandfield) Key(Key string) SHrandfieldKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHrandfieldKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hrandfield() (c SHrandfield) {
	c.cs = append(b.get(), "HRANDFIELD")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHrandfieldKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHrandfieldKey) Count(Count int64) SHrandfieldOptionsCount {
	return SHrandfieldOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type SHrandfieldOptionsCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHrandfieldOptionsCount) Withvalues() SHrandfieldOptionsWithvaluesWithvalues {
	return SHrandfieldOptionsWithvaluesWithvalues{cf: c.cf, cs: append(c.cs, "WITHVALUES")}
}

func (c SHrandfieldOptionsCount) Build() SCompleted {
	return SCompleted(c)
}

type SHrandfieldOptionsWithvaluesWithvalues struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHrandfieldOptionsWithvaluesWithvalues) Build() SCompleted {
	return SCompleted(c)
}

type SHscan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHscan) Key(Key string) SHscanKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hscan() (c SHscan) {
	c.cs = append(b.get(), "HSCAN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHscanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHscanCount) Build() SCompleted {
	return SCompleted(c)
}

type SHscanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHscanCursor) Match(Pattern string) SHscanMatch {
	return SHscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SHscanCursor) Count(Count int64) SHscanCount {
	return SHscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SHscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SHscanKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHscanKey) Cursor(Cursor int64) SHscanCursor {
	return SHscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SHscanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHscanMatch) Count(Count int64) SHscanCount {
	return SHscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SHscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SHset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHset) Key(Key string) SHsetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hset() (c SHset) {
	c.cs = append(b.get(), "HSET")
	c.ks = initSlot
	return
}

type SHsetFieldValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHsetFieldValue) FieldValue(Field string, Value string) SHsetFieldValue {
	return SHsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c SHsetFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SHsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHsetKey) FieldValue() SHsetFieldValue {
	return SHsetFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type SHsetnx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHsetnx) Key(Key string) SHsetnxKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHsetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hsetnx() (c SHsetnx) {
	c.cs = append(b.get(), "HSETNX")
	c.ks = initSlot
	return
}

type SHsetnxField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHsetnxField) Value(Value string) SHsetnxValue {
	return SHsetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SHsetnxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHsetnxKey) Field(Field string) SHsetnxField {
	return SHsetnxField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHsetnxValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHsetnxValue) Build() SCompleted {
	return SCompleted(c)
}

type SHstrlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHstrlen) Key(Key string) SHstrlenKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHstrlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hstrlen() (c SHstrlen) {
	c.cs = append(b.get(), "HSTRLEN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHstrlenField struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHstrlenField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHstrlenField) Cache() SCacheable {
	return SCacheable(c)
}

type SHstrlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHstrlenKey) Field(Field string) SHstrlenField {
	return SHstrlenField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c SHstrlenKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHvals struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHvals) Key(Key string) SHvalsKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SHvalsKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hvals() (c SHvals) {
	c.cs = append(b.get(), "HVALS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SHvalsKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SHvalsKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHvalsKey) Cache() SCacheable {
	return SCacheable(c)
}

type SIncr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncr) Key(Key string) SIncrKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SIncrKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Incr() (c SIncr) {
	c.cs = append(b.get(), "INCR")
	c.ks = initSlot
	return
}

type SIncrKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncrKey) Build() SCompleted {
	return SCompleted(c)
}

type SIncrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncrby) Key(Key string) SIncrbyKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SIncrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Incrby() (c SIncrby) {
	c.cs = append(b.get(), "INCRBY")
	c.ks = initSlot
	return
}

type SIncrbyIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncrbyIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SIncrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncrbyKey) Increment(Increment int64) SIncrbyIncrement {
	return SIncrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type SIncrbyfloat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncrbyfloat) Key(Key string) SIncrbyfloatKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SIncrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Incrbyfloat() (c SIncrbyfloat) {
	c.cs = append(b.get(), "INCRBYFLOAT")
	c.ks = initSlot
	return
}

type SIncrbyfloatIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncrbyfloatIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SIncrbyfloatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SIncrbyfloatKey) Increment(Increment float64) SIncrbyfloatIncrement {
	return SIncrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type SInfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SInfo) Section(Section string) SInfoSection {
	return SInfoSection{cf: c.cf, cs: append(c.cs, Section)}
}

func (c SInfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Info() (c SInfo) {
	c.cs = append(b.get(), "INFO")
	c.ks = initSlot
	return
}

type SInfoSection struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SInfoSection) Build() SCompleted {
	return SCompleted(c)
}

type SKeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SKeys) Pattern(Pattern string) SKeysPattern {
	return SKeysPattern{cf: c.cf, cs: append(c.cs, Pattern)}
}

func (b *SBuilder) Keys() (c SKeys) {
	c.cs = append(b.get(), "KEYS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SKeysPattern struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SKeysPattern) Build() SCompleted {
	return SCompleted(c)
}

type SLastsave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLastsave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Lastsave() (c SLastsave) {
	c.cs = append(b.get(), "LASTSAVE")
	c.ks = initSlot
	return
}

type SLatencyDoctor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyDoctor) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyDoctor() (c SLatencyDoctor) {
	c.cs = append(b.get(), "LATENCY", "DOCTOR")
	c.ks = initSlot
	return
}

type SLatencyGraph struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyGraph) Event(Event string) SLatencyGraphEvent {
	return SLatencyGraphEvent{cf: c.cf, cs: append(c.cs, Event)}
}

func (b *SBuilder) LatencyGraph() (c SLatencyGraph) {
	c.cs = append(b.get(), "LATENCY", "GRAPH")
	c.ks = initSlot
	return
}

type SLatencyGraphEvent struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyGraphEvent) Build() SCompleted {
	return SCompleted(c)
}

type SLatencyHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyHelp) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyHelp() (c SLatencyHelp) {
	c.cs = append(b.get(), "LATENCY", "HELP")
	c.ks = initSlot
	return
}

type SLatencyHistory struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyHistory) Event(Event string) SLatencyHistoryEvent {
	return SLatencyHistoryEvent{cf: c.cf, cs: append(c.cs, Event)}
}

func (b *SBuilder) LatencyHistory() (c SLatencyHistory) {
	c.cs = append(b.get(), "LATENCY", "HISTORY")
	c.ks = initSlot
	return
}

type SLatencyHistoryEvent struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyHistoryEvent) Build() SCompleted {
	return SCompleted(c)
}

type SLatencyLatest struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyLatest) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyLatest() (c SLatencyLatest) {
	c.cs = append(b.get(), "LATENCY", "LATEST")
	c.ks = initSlot
	return
}

type SLatencyReset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyReset) Event(Event ...string) SLatencyResetEvent {
	return SLatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
}

func (c SLatencyReset) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyReset() (c SLatencyReset) {
	c.cs = append(b.get(), "LATENCY", "RESET")
	c.ks = initSlot
	return
}

type SLatencyResetEvent struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLatencyResetEvent) Event(Event ...string) SLatencyResetEvent {
	return SLatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
}

func (c SLatencyResetEvent) Build() SCompleted {
	return SCompleted(c)
}

type SLindex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLindex) Key(Key string) SLindexKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLindexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lindex() (c SLindex) {
	c.cs = append(b.get(), "LINDEX")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SLindexIndex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLindexIndex) Build() SCompleted {
	return SCompleted(c)
}

func (c SLindexIndex) Cache() SCacheable {
	return SCacheable(c)
}

type SLindexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLindexKey) Index(Index int64) SLindexIndex {
	return SLindexIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

func (c SLindexKey) Cache() SCacheable {
	return SCacheable(c)
}

type SLinsert struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLinsert) Key(Key string) SLinsertKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLinsertKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Linsert() (c SLinsert) {
	c.cs = append(b.get(), "LINSERT")
	c.ks = initSlot
	return
}

type SLinsertElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLinsertElement) Build() SCompleted {
	return SCompleted(c)
}

type SLinsertKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLinsertKey) Before() SLinsertWhereBefore {
	return SLinsertWhereBefore{cf: c.cf, cs: append(c.cs, "BEFORE")}
}

func (c SLinsertKey) After() SLinsertWhereAfter {
	return SLinsertWhereAfter{cf: c.cf, cs: append(c.cs, "AFTER")}
}

type SLinsertPivot struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLinsertPivot) Element(Element string) SLinsertElement {
	return SLinsertElement{cf: c.cf, cs: append(c.cs, Element)}
}

type SLinsertWhereAfter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLinsertWhereAfter) Pivot(Pivot string) SLinsertPivot {
	return SLinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type SLinsertWhereBefore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLinsertWhereBefore) Pivot(Pivot string) SLinsertPivot {
	return SLinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type SLlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLlen) Key(Key string) SLlenKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Llen() (c SLlen) {
	c.cs = append(b.get(), "LLEN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SLlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLlenKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SLlenKey) Cache() SCacheable {
	return SCacheable(c)
}

type SLmove struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmove) Source(Source string) SLmoveSource {
	s := slot(Source)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Lmove() (c SLmove) {
	c.cs = append(b.get(), "LMOVE")
	c.ks = initSlot
	return
}

type SLmoveDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmoveDestination) Left() SLmoveWherefromLeft {
	return SLmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmoveDestination) Right() SLmoveWherefromRight {
	return SLmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SLmoveSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmoveSource) Destination(Destination string) SLmoveDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SLmoveWherefromLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmoveWherefromLeft) Left() SLmoveWheretoLeft {
	return SLmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmoveWherefromLeft) Right() SLmoveWheretoRight {
	return SLmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SLmoveWherefromRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmoveWherefromRight) Left() SLmoveWheretoLeft {
	return SLmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmoveWherefromRight) Right() SLmoveWheretoRight {
	return SLmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SLmoveWheretoLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmoveWheretoLeft) Build() SCompleted {
	return SCompleted(c)
}

type SLmoveWheretoRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmoveWheretoRight) Build() SCompleted {
	return SCompleted(c)
}

type SLmpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmpop) Numkeys(Numkeys int64) SLmpopNumkeys {
	return SLmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Lmpop() (c SLmpop) {
	c.cs = append(b.get(), "LMPOP")
	c.ks = initSlot
	return
}

type SLmpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SLmpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmpopKey) Left() SLmpopWhereLeft {
	return SLmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmpopKey) Right() SLmpopWhereRight {
	return SLmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c SLmpopKey) Key(Key ...string) SLmpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SLmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SLmpopNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmpopNumkeys) Key(Key ...string) SLmpopKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SLmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SLmpopNumkeys) Left() SLmpopWhereLeft {
	return SLmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmpopNumkeys) Right() SLmpopWhereRight {
	return SLmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SLmpopWhereLeft struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmpopWhereLeft) Count(Count int64) SLmpopCount {
	return SLmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SLmpopWhereLeft) Build() SCompleted {
	return SCompleted(c)
}

type SLmpopWhereRight struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLmpopWhereRight) Count(Count int64) SLmpopCount {
	return SLmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SLmpopWhereRight) Build() SCompleted {
	return SCompleted(c)
}

type SLolwut struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLolwut) Version(Version int64) SLolwutVersion {
	return SLolwutVersion{cf: c.cf, cs: append(c.cs, "VERSION", strconv.FormatInt(Version, 10))}
}

func (c SLolwut) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Lolwut() (c SLolwut) {
	c.cs = append(b.get(), "LOLWUT")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SLolwutVersion struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLolwutVersion) Build() SCompleted {
	return SCompleted(c)
}

type SLpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpop) Key(Key string) SLpopKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLpopKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lpop() (c SLpop) {
	c.cs = append(b.get(), "LPOP")
	c.ks = initSlot
	return
}

type SLpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SLpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpopKey) Count(Count int64) SLpopCount {
	return SLpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SLpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SLpos struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpos) Key(Key string) SLposKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lpos() (c SLpos) {
	c.cs = append(b.get(), "LPOS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SLposCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLposCount) Maxlen(Len int64) SLposMaxlen {
	return SLposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c SLposCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposCount) Cache() SCacheable {
	return SCacheable(c)
}

type SLposElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLposElement) Rank(Rank int64) SLposRank {
	return SLposRank{cf: c.cf, cs: append(c.cs, "RANK", strconv.FormatInt(Rank, 10))}
}

func (c SLposElement) Count(NumMatches int64) SLposCount {
	return SLposCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10))}
}

func (c SLposElement) Maxlen(Len int64) SLposMaxlen {
	return SLposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c SLposElement) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposElement) Cache() SCacheable {
	return SCacheable(c)
}

type SLposKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLposKey) Element(Element string) SLposElement {
	return SLposElement{cf: c.cf, cs: append(c.cs, Element)}
}

func (c SLposKey) Cache() SCacheable {
	return SCacheable(c)
}

type SLposMaxlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLposMaxlen) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposMaxlen) Cache() SCacheable {
	return SCacheable(c)
}

type SLposRank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLposRank) Count(NumMatches int64) SLposCount {
	return SLposCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10))}
}

func (c SLposRank) Maxlen(Len int64) SLposMaxlen {
	return SLposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c SLposRank) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposRank) Cache() SCacheable {
	return SCacheable(c)
}

type SLpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpush) Key(Key string) SLpushKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lpush() (c SLpush) {
	c.cs = append(b.get(), "LPUSH")
	c.ks = initSlot
	return
}

type SLpushElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpushElement) Element(Element ...string) SLpushElement {
	return SLpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SLpushElement) Build() SCompleted {
	return SCompleted(c)
}

type SLpushKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpushKey) Element(Element ...string) SLpushElement {
	return SLpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SLpushx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpushx) Key(Key string) SLpushxKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lpushx() (c SLpushx) {
	c.cs = append(b.get(), "LPUSHX")
	c.ks = initSlot
	return
}

type SLpushxElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpushxElement) Element(Element ...string) SLpushxElement {
	return SLpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SLpushxElement) Build() SCompleted {
	return SCompleted(c)
}

type SLpushxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLpushxKey) Element(Element ...string) SLpushxElement {
	return SLpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SLrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLrange) Key(Key string) SLrangeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lrange() (c SLrange) {
	c.cs = append(b.get(), "LRANGE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SLrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLrangeKey) Start(Start int64) SLrangeStart {
	return SLrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c SLrangeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SLrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLrangeStart) Stop(Stop int64) SLrangeStop {
	return SLrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

func (c SLrangeStart) Cache() SCacheable {
	return SCacheable(c)
}

type SLrangeStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLrangeStop) Build() SCompleted {
	return SCompleted(c)
}

func (c SLrangeStop) Cache() SCacheable {
	return SCacheable(c)
}

type SLrem struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLrem) Key(Key string) SLremKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lrem() (c SLrem) {
	c.cs = append(b.get(), "LREM")
	c.ks = initSlot
	return
}

type SLremCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLremCount) Element(Element string) SLremElement {
	return SLremElement{cf: c.cf, cs: append(c.cs, Element)}
}

type SLremElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLremElement) Build() SCompleted {
	return SCompleted(c)
}

type SLremKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLremKey) Count(Count int64) SLremCount {
	return SLremCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type SLset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLset) Key(Key string) SLsetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lset() (c SLset) {
	c.cs = append(b.get(), "LSET")
	c.ks = initSlot
	return
}

type SLsetElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLsetElement) Build() SCompleted {
	return SCompleted(c)
}

type SLsetIndex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLsetIndex) Element(Element string) SLsetElement {
	return SLsetElement{cf: c.cf, cs: append(c.cs, Element)}
}

type SLsetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLsetKey) Index(Index int64) SLsetIndex {
	return SLsetIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type SLtrim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLtrim) Key(Key string) SLtrimKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SLtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Ltrim() (c SLtrim) {
	c.cs = append(b.get(), "LTRIM")
	c.ks = initSlot
	return
}

type SLtrimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLtrimKey) Start(Start int64) SLtrimStart {
	return SLtrimStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type SLtrimStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLtrimStart) Stop(Stop int64) SLtrimStop {
	return SLtrimStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type SLtrimStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SLtrimStop) Build() SCompleted {
	return SCompleted(c)
}

type SMemoryDoctor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryDoctor) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryDoctor() (c SMemoryDoctor) {
	c.cs = append(b.get(), "MEMORY", "DOCTOR")
	c.ks = initSlot
	return
}

type SMemoryHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryHelp) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryHelp() (c SMemoryHelp) {
	c.cs = append(b.get(), "MEMORY", "HELP")
	c.ks = initSlot
	return
}

type SMemoryMallocStats struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryMallocStats) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryMallocStats() (c SMemoryMallocStats) {
	c.cs = append(b.get(), "MEMORY", "MALLOC-STATS")
	c.ks = initSlot
	return
}

type SMemoryPurge struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryPurge) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryPurge() (c SMemoryPurge) {
	c.cs = append(b.get(), "MEMORY", "PURGE")
	c.ks = initSlot
	return
}

type SMemoryStats struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryStats) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryStats() (c SMemoryStats) {
	c.cs = append(b.get(), "MEMORY", "STATS")
	c.ks = initSlot
	return
}

type SMemoryUsage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryUsage) Key(Key string) SMemoryUsageKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SMemoryUsageKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) MemoryUsage() (c SMemoryUsage) {
	c.cs = append(b.get(), "MEMORY", "USAGE")
	c.ks = initSlot
	return
}

type SMemoryUsageKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryUsageKey) Samples(Count int64) SMemoryUsageSamples {
	return SMemoryUsageSamples{cf: c.cf, cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10))}
}

func (c SMemoryUsageKey) Build() SCompleted {
	return SCompleted(c)
}

type SMemoryUsageSamples struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMemoryUsageSamples) Build() SCompleted {
	return SCompleted(c)
}

type SMget struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMget) Key(Key ...string) SMgetKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SMgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Mget() (c SMget) {
	c.cs = append(b.get(), "MGET")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SMgetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMgetKey) Key(Key ...string) SMgetKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SMgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMgetKey) Build() SCompleted {
	return SCompleted(c)
}

type SMigrate struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrate) Host(Host string) SMigrateHost {
	return SMigrateHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *SBuilder) Migrate() (c SMigrate) {
	c.cs = append(b.get(), "MIGRATE")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SMigrateAuth struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateAuth) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c SMigrateAuth) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateAuth) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateAuth2 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateAuth2) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateAuth2) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateCopyCopy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateCopyCopy) Replace() SMigrateReplaceReplace {
	return SMigrateReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c SMigrateCopyCopy) Auth(Password string) SMigrateAuth {
	return SMigrateAuth{cf: c.cf, cs: append(c.cs, "AUTH", Password)}
}

func (c SMigrateCopyCopy) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c SMigrateCopyCopy) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateCopyCopy) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateDestinationDb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateDestinationDb) Timeout(Timeout int64) SMigrateTimeout {
	return SMigrateTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type SMigrateHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateHost) Port(Port string) SMigratePort {
	return SMigratePort{cf: c.cf, cs: append(c.cs, Port)}
}

type SMigrateKeyEmpty struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateKeyEmpty) DestinationDb(DestinationDb int64) SMigrateDestinationDb {
	return SMigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type SMigrateKeyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateKeyKey) DestinationDb(DestinationDb int64) SMigrateDestinationDb {
	return SMigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type SMigrateKeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateKeys) Keys(Keys ...string) SMigrateKeys {
	for _, k := range Keys {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Keys...)}
}

func (c SMigrateKeys) Build() SCompleted {
	return SCompleted(c)
}

type SMigratePort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigratePort) Key() SMigrateKeyKey {
	return SMigrateKeyKey{cf: c.cf, cs: append(c.cs, "key")}
}

func (c SMigratePort) Empty() SMigrateKeyEmpty {
	return SMigrateKeyEmpty{cf: c.cf, cs: append(c.cs, "\"\"")}
}

type SMigrateReplaceReplace struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateReplaceReplace) Auth(Password string) SMigrateAuth {
	return SMigrateAuth{cf: c.cf, cs: append(c.cs, "AUTH", Password)}
}

func (c SMigrateReplaceReplace) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c SMigrateReplaceReplace) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateReplaceReplace) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMigrateTimeout) Copy() SMigrateCopyCopy {
	return SMigrateCopyCopy{cf: c.cf, cs: append(c.cs, "COPY")}
}

func (c SMigrateTimeout) Replace() SMigrateReplaceReplace {
	return SMigrateReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c SMigrateTimeout) Auth(Password string) SMigrateAuth {
	return SMigrateAuth{cf: c.cf, cs: append(c.cs, "AUTH", Password)}
}

func (c SMigrateTimeout) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c SMigrateTimeout) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SModuleList struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SModuleList) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ModuleList() (c SModuleList) {
	c.cs = append(b.get(), "MODULE", "LIST")
	c.ks = initSlot
	return
}

type SModuleLoad struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SModuleLoad) Path(Path string) SModuleLoadPath {
	return SModuleLoadPath{cf: c.cf, cs: append(c.cs, Path)}
}

func (b *SBuilder) ModuleLoad() (c SModuleLoad) {
	c.cs = append(b.get(), "MODULE", "LOAD")
	c.ks = initSlot
	return
}

type SModuleLoadArg struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SModuleLoadArg) Arg(Arg ...string) SModuleLoadArg {
	return SModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SModuleLoadArg) Build() SCompleted {
	return SCompleted(c)
}

type SModuleLoadPath struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SModuleLoadPath) Arg(Arg ...string) SModuleLoadArg {
	return SModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SModuleLoadPath) Build() SCompleted {
	return SCompleted(c)
}

type SModuleUnload struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SModuleUnload) Name(Name string) SModuleUnloadName {
	return SModuleUnloadName{cf: c.cf, cs: append(c.cs, Name)}
}

func (b *SBuilder) ModuleUnload() (c SModuleUnload) {
	c.cs = append(b.get(), "MODULE", "UNLOAD")
	c.ks = initSlot
	return
}

type SModuleUnloadName struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SModuleUnloadName) Build() SCompleted {
	return SCompleted(c)
}

type SMonitor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMonitor) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Monitor() (c SMonitor) {
	c.cs = append(b.get(), "MONITOR")
	c.ks = initSlot
	return
}

type SMove struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMove) Key(Key string) SMoveKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SMoveKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Move() (c SMove) {
	c.cs = append(b.get(), "MOVE")
	c.ks = initSlot
	return
}

type SMoveDb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMoveDb) Build() SCompleted {
	return SCompleted(c)
}

type SMoveKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMoveKey) Db(Db int64) SMoveDb {
	return SMoveDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Db, 10))}
}

type SMset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMset) KeyValue() SMsetKeyValue {
	return SMsetKeyValue{cf: c.cf, cs: append(c.cs, )}
}

func (b *SBuilder) Mset() (c SMset) {
	c.cs = append(b.get(), "MSET")
	c.ks = initSlot
	return
}

type SMsetKeyValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMsetKeyValue) KeyValue(Key string, Value string) SMsetKeyValue {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SMsetKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c SMsetKeyValue) Build() SCompleted {
	return SCompleted(c)
}

type SMsetnx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMsetnx) KeyValue() SMsetnxKeyValue {
	return SMsetnxKeyValue{cf: c.cf, cs: append(c.cs, )}
}

func (b *SBuilder) Msetnx() (c SMsetnx) {
	c.cs = append(b.get(), "MSETNX")
	c.ks = initSlot
	return
}

type SMsetnxKeyValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMsetnxKeyValue) KeyValue(Key string, Value string) SMsetnxKeyValue {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SMsetnxKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c SMsetnxKeyValue) Build() SCompleted {
	return SCompleted(c)
}

type SMulti struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SMulti) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Multi() (c SMulti) {
	c.cs = append(b.get(), "MULTI")
	c.ks = initSlot
	return
}

type SObject struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SObject) Subcommand(Subcommand string) SObjectSubcommand {
	return SObjectSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *SBuilder) Object() (c SObject) {
	c.cs = append(b.get(), "OBJECT")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SObjectArguments struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SObjectArguments) Arguments(Arguments ...string) SObjectArguments {
	return SObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c SObjectArguments) Build() SCompleted {
	return SCompleted(c)
}

type SObjectSubcommand struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SObjectSubcommand) Arguments(Arguments ...string) SObjectArguments {
	return SObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c SObjectSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SPersist struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPersist) Key(Key string) SPersistKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPersistKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Persist() (c SPersist) {
	c.cs = append(b.get(), "PERSIST")
	c.ks = initSlot
	return
}

type SPersistKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPersistKey) Build() SCompleted {
	return SCompleted(c)
}

type SPexpire struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpire) Key(Key string) SPexpireKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPexpireKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Pexpire() (c SPexpire) {
	c.cs = append(b.get(), "PEXPIRE")
	c.ks = initSlot
	return
}

type SPexpireConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireKey) Milliseconds(Milliseconds int64) SPexpireMilliseconds {
	return SPexpireMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type SPexpireMilliseconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireMilliseconds) Nx() SPexpireConditionNx {
	return SPexpireConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SPexpireMilliseconds) Xx() SPexpireConditionXx {
	return SPexpireConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SPexpireMilliseconds) Gt() SPexpireConditionGt {
	return SPexpireConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c SPexpireMilliseconds) Lt() SPexpireConditionLt {
	return SPexpireConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c SPexpireMilliseconds) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireat) Key(Key string) SPexpireatKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPexpireatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Pexpireat() (c SPexpireat) {
	c.cs = append(b.get(), "PEXPIREAT")
	c.ks = initSlot
	return
}

type SPexpireatConditionGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireatConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatConditionLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireatConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireatConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireatConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireatKey) MillisecondsTimestamp(MillisecondsTimestamp int64) SPexpireatMillisecondsTimestamp {
	return SPexpireatMillisecondsTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10))}
}

type SPexpireatMillisecondsTimestamp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpireatMillisecondsTimestamp) Nx() SPexpireatConditionNx {
	return SPexpireatConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SPexpireatMillisecondsTimestamp) Xx() SPexpireatConditionXx {
	return SPexpireatConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SPexpireatMillisecondsTimestamp) Gt() SPexpireatConditionGt {
	return SPexpireatConditionGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c SPexpireatMillisecondsTimestamp) Lt() SPexpireatConditionLt {
	return SPexpireatConditionLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c SPexpireatMillisecondsTimestamp) Build() SCompleted {
	return SCompleted(c)
}

type SPexpiretime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpiretime) Key(Key string) SPexpiretimeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPexpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Pexpiretime() (c SPexpiretime) {
	c.cs = append(b.get(), "PEXPIRETIME")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SPexpiretimeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPexpiretimeKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SPexpiretimeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SPfadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfadd) Key(Key string) SPfaddKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPfaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Pfadd() (c SPfadd) {
	c.cs = append(b.get(), "PFADD")
	c.ks = initSlot
	return
}

type SPfaddElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfaddElement) Element(Element ...string) SPfaddElement {
	return SPfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SPfaddElement) Build() SCompleted {
	return SCompleted(c)
}

type SPfaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfaddKey) Element(Element ...string) SPfaddElement {
	return SPfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SPfaddKey) Build() SCompleted {
	return SCompleted(c)
}

type SPfcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfcount) Key(Key ...string) SPfcountKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SPfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Pfcount() (c SPfcount) {
	c.cs = append(b.get(), "PFCOUNT")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SPfcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfcountKey) Key(Key ...string) SPfcountKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SPfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SPfcountKey) Build() SCompleted {
	return SCompleted(c)
}

type SPfmerge struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfmerge) Destkey(Destkey string) SPfmergeDestkey {
	s := slot(Destkey)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPfmergeDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

func (b *SBuilder) Pfmerge() (c SPfmerge) {
	c.cs = append(b.get(), "PFMERGE")
	c.ks = initSlot
	return
}

type SPfmergeDestkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfmergeDestkey) Sourcekey(Sourcekey ...string) SPfmergeSourcekey {
	for _, k := range Sourcekey {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SPfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

type SPfmergeSourcekey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPfmergeSourcekey) Sourcekey(Sourcekey ...string) SPfmergeSourcekey {
	for _, k := range Sourcekey {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SPfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

func (c SPfmergeSourcekey) Build() SCompleted {
	return SCompleted(c)
}

type SPing struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPing) Message(Message string) SPingMessage {
	return SPingMessage{cf: c.cf, cs: append(c.cs, Message)}
}

func (c SPing) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Ping() (c SPing) {
	c.cs = append(b.get(), "PING")
	c.ks = initSlot
	return
}

type SPingMessage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPingMessage) Build() SCompleted {
	return SCompleted(c)
}

type SPsetex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsetex) Key(Key string) SPsetexKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPsetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Psetex() (c SPsetex) {
	c.cs = append(b.get(), "PSETEX")
	c.ks = initSlot
	return
}

type SPsetexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsetexKey) Milliseconds(Milliseconds int64) SPsetexMilliseconds {
	return SPsetexMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type SPsetexMilliseconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsetexMilliseconds) Value(Value string) SPsetexValue {
	return SPsetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SPsetexValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsetexValue) Build() SCompleted {
	return SCompleted(c)
}

type SPsubscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsubscribe) Pattern(Pattern ...string) SPsubscribePattern {
	return SPsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (b *SBuilder) Psubscribe() (c SPsubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	c.cf = noRetTag
	c.ks = initSlot
	return
}

type SPsubscribePattern struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsubscribePattern) Pattern(Pattern ...string) SPsubscribePattern {
	return SPsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SPsubscribePattern) Build() SCompleted {
	return SCompleted(c)
}

type SPsync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsync) Replicationid(Replicationid int64) SPsyncReplicationid {
	return SPsyncReplicationid{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Replicationid, 10))}
}

func (b *SBuilder) Psync() (c SPsync) {
	c.cs = append(b.get(), "PSYNC")
	c.ks = initSlot
	return
}

type SPsyncOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsyncOffset) Build() SCompleted {
	return SCompleted(c)
}

type SPsyncReplicationid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPsyncReplicationid) Offset(Offset int64) SPsyncOffset {
	return SPsyncOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SPttl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPttl) Key(Key string) SPttlKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SPttlKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Pttl() (c SPttl) {
	c.cs = append(b.get(), "PTTL")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SPttlKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPttlKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SPttlKey) Cache() SCacheable {
	return SCacheable(c)
}

type SPublish struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPublish) Channel(Channel string) SPublishChannel {
	return SPublishChannel{cf: c.cf, cs: append(c.cs, Channel)}
}

func (b *SBuilder) Publish() (c SPublish) {
	c.cs = append(b.get(), "PUBLISH")
	c.ks = initSlot
	return
}

type SPublishChannel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPublishChannel) Message(Message string) SPublishMessage {
	return SPublishMessage{cf: c.cf, cs: append(c.cs, Message)}
}

type SPublishMessage struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPublishMessage) Build() SCompleted {
	return SCompleted(c)
}

type SPubsub struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPubsub) Subcommand(Subcommand string) SPubsubSubcommand {
	return SPubsubSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *SBuilder) Pubsub() (c SPubsub) {
	c.cs = append(b.get(), "PUBSUB")
	c.ks = initSlot
	return
}

type SPubsubArgument struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPubsubArgument) Argument(Argument ...string) SPubsubArgument {
	return SPubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c SPubsubArgument) Build() SCompleted {
	return SCompleted(c)
}

type SPubsubSubcommand struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPubsubSubcommand) Argument(Argument ...string) SPubsubArgument {
	return SPubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c SPubsubSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SPunsubscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPunsubscribe) Pattern(Pattern ...string) SPunsubscribePattern {
	return SPunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SPunsubscribe) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Punsubscribe() (c SPunsubscribe) {
	c.cs = append(b.get(), "PUNSUBSCRIBE")
	c.cf = noRetTag
	c.ks = initSlot
	return
}

type SPunsubscribePattern struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SPunsubscribePattern) Pattern(Pattern ...string) SPunsubscribePattern {
	return SPunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SPunsubscribePattern) Build() SCompleted {
	return SCompleted(c)
}

type SQuit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SQuit) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Quit() (c SQuit) {
	c.cs = append(b.get(), "QUIT")
	c.ks = initSlot
	return
}

type SRandomkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRandomkey) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Randomkey() (c SRandomkey) {
	c.cs = append(b.get(), "RANDOMKEY")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SReadonly struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SReadonly) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Readonly() (c SReadonly) {
	c.cs = append(b.get(), "READONLY")
	c.ks = initSlot
	return
}

type SReadwrite struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SReadwrite) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Readwrite() (c SReadwrite) {
	c.cs = append(b.get(), "READWRITE")
	c.ks = initSlot
	return
}

type SRename struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRename) Key(Key string) SRenameKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRenameKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Rename() (c SRename) {
	c.cs = append(b.get(), "RENAME")
	c.ks = initSlot
	return
}

type SRenameKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRenameKey) Newkey(Newkey string) SRenameNewkey {
	s := slot(Newkey)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRenameNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type SRenameNewkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRenameNewkey) Build() SCompleted {
	return SCompleted(c)
}

type SRenamenx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRenamenx) Key(Key string) SRenamenxKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRenamenxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Renamenx() (c SRenamenx) {
	c.cs = append(b.get(), "RENAMENX")
	c.ks = initSlot
	return
}

type SRenamenxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRenamenxKey) Newkey(Newkey string) SRenamenxNewkey {
	s := slot(Newkey)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRenamenxNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type SRenamenxNewkey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRenamenxNewkey) Build() SCompleted {
	return SCompleted(c)
}

type SReplicaof struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SReplicaof) Host(Host string) SReplicaofHost {
	return SReplicaofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *SBuilder) Replicaof() (c SReplicaof) {
	c.cs = append(b.get(), "REPLICAOF")
	c.ks = initSlot
	return
}

type SReplicaofHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SReplicaofHost) Port(Port string) SReplicaofPort {
	return SReplicaofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type SReplicaofPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SReplicaofPort) Build() SCompleted {
	return SCompleted(c)
}

type SReset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SReset) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Reset() (c SReset) {
	c.cs = append(b.get(), "RESET")
	c.ks = initSlot
	return
}

type SRestore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestore) Key(Key string) SRestoreKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRestoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Restore() (c SRestore) {
	c.cs = append(b.get(), "RESTORE")
	c.ks = initSlot
	return
}

type SRestoreAbsttlAbsttl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestoreAbsttlAbsttl) Idletime(Seconds int64) SRestoreIdletime {
	return SRestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c SRestoreAbsttlAbsttl) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c SRestoreAbsttlAbsttl) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreFreq struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestoreFreq) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreIdletime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestoreIdletime) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c SRestoreIdletime) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestoreKey) Ttl(Ttl int64) SRestoreTtl {
	return SRestoreTtl{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Ttl, 10))}
}

type SRestoreReplaceReplace struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestoreReplaceReplace) Absttl() SRestoreAbsttlAbsttl {
	return SRestoreAbsttlAbsttl{cf: c.cf, cs: append(c.cs, "ABSTTL")}
}

func (c SRestoreReplaceReplace) Idletime(Seconds int64) SRestoreIdletime {
	return SRestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c SRestoreReplaceReplace) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c SRestoreReplaceReplace) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreSerializedValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestoreSerializedValue) Replace() SRestoreReplaceReplace {
	return SRestoreReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c SRestoreSerializedValue) Absttl() SRestoreAbsttlAbsttl {
	return SRestoreAbsttlAbsttl{cf: c.cf, cs: append(c.cs, "ABSTTL")}
}

func (c SRestoreSerializedValue) Idletime(Seconds int64) SRestoreIdletime {
	return SRestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c SRestoreSerializedValue) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c SRestoreSerializedValue) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreTtl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRestoreTtl) SerializedValue(SerializedValue string) SRestoreSerializedValue {
	return SRestoreSerializedValue{cf: c.cf, cs: append(c.cs, SerializedValue)}
}

type SRole struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRole) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Role() (c SRole) {
	c.cs = append(b.get(), "ROLE")
	c.ks = initSlot
	return
}

type SRpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpop) Key(Key string) SRpopKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRpopKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Rpop() (c SRpop) {
	c.cs = append(b.get(), "RPOP")
	c.ks = initSlot
	return
}

type SRpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SRpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpopKey) Count(Count int64) SRpopCount {
	return SRpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SRpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SRpoplpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpoplpush) Source(Source string) SRpoplpushSource {
	s := slot(Source)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Rpoplpush() (c SRpoplpush) {
	c.cs = append(b.get(), "RPOPLPUSH")
	c.ks = initSlot
	return
}

type SRpoplpushDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpoplpushDestination) Build() SCompleted {
	return SCompleted(c)
}

type SRpoplpushSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpoplpushSource) Destination(Destination string) SRpoplpushDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SRpush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpush) Key(Key string) SRpushKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Rpush() (c SRpush) {
	c.cs = append(b.get(), "RPUSH")
	c.ks = initSlot
	return
}

type SRpushElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpushElement) Element(Element ...string) SRpushElement {
	return SRpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SRpushElement) Build() SCompleted {
	return SCompleted(c)
}

type SRpushKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpushKey) Element(Element ...string) SRpushElement {
	return SRpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SRpushx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpushx) Key(Key string) SRpushxKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SRpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Rpushx() (c SRpushx) {
	c.cs = append(b.get(), "RPUSHX")
	c.ks = initSlot
	return
}

type SRpushxElement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpushxElement) Element(Element ...string) SRpushxElement {
	return SRpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SRpushxElement) Build() SCompleted {
	return SCompleted(c)
}

type SRpushxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SRpushxKey) Element(Element ...string) SRpushxElement {
	return SRpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SSadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSadd) Key(Key string) SSaddKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Sadd() (c SSadd) {
	c.cs = append(b.get(), "SADD")
	c.ks = initSlot
	return
}

type SSaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSaddKey) Member(Member ...string) SSaddMember {
	return SSaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SSaddMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSaddMember) Member(Member ...string) SSaddMember {
	return SSaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SSaddMember) Build() SCompleted {
	return SCompleted(c)
}

type SSave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Save() (c SSave) {
	c.cs = append(b.get(), "SAVE")
	c.ks = initSlot
	return
}

type SScan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScan) Cursor(Cursor int64) SScanCursor {
	return SScanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

func (b *SBuilder) Scan() (c SScan) {
	c.cs = append(b.get(), "SCAN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SScanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScanCount) Type(Type string) SScanType {
	return SScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c SScanCount) Build() SCompleted {
	return SCompleted(c)
}

type SScanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScanCursor) Match(Pattern string) SScanMatch {
	return SScanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SScanCursor) Count(Count int64) SScanCount {
	return SScanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SScanCursor) Type(Type string) SScanType {
	return SScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c SScanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SScanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScanMatch) Count(Count int64) SScanCount {
	return SScanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SScanMatch) Type(Type string) SScanType {
	return SScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c SScanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SScanType struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScanType) Build() SCompleted {
	return SCompleted(c)
}

type SScard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScard) Key(Key string) SScardKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SScardKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Scard() (c SScard) {
	c.cs = append(b.get(), "SCARD")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SScardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScardKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SScardKey) Cache() SCacheable {
	return SCacheable(c)
}

type SScriptDebug struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptDebug) Yes() SScriptDebugModeYes {
	return SScriptDebugModeYes{cf: c.cf, cs: append(c.cs, "YES")}
}

func (c SScriptDebug) Sync() SScriptDebugModeSync {
	return SScriptDebugModeSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c SScriptDebug) No() SScriptDebugModeNo {
	return SScriptDebugModeNo{cf: c.cf, cs: append(c.cs, "NO")}
}

func (b *SBuilder) ScriptDebug() (c SScriptDebug) {
	c.cs = append(b.get(), "SCRIPT", "DEBUG")
	c.ks = initSlot
	return
}

type SScriptDebugModeNo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptDebugModeNo) Build() SCompleted {
	return SCompleted(c)
}

type SScriptDebugModeSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptDebugModeSync) Build() SCompleted {
	return SCompleted(c)
}

type SScriptDebugModeYes struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptDebugModeYes) Build() SCompleted {
	return SCompleted(c)
}

type SScriptExists struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptExists) Sha1(Sha1 ...string) SScriptExistsSha1 {
	return SScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (b *SBuilder) ScriptExists() (c SScriptExists) {
	c.cs = append(b.get(), "SCRIPT", "EXISTS")
	c.ks = initSlot
	return
}

type SScriptExistsSha1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptExistsSha1) Sha1(Sha1 ...string) SScriptExistsSha1 {
	return SScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (c SScriptExistsSha1) Build() SCompleted {
	return SCompleted(c)
}

type SScriptFlush struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptFlush) Async() SScriptFlushAsyncAsync {
	return SScriptFlushAsyncAsync{cf: c.cf, cs: append(c.cs, "ASYNC")}
}

func (c SScriptFlush) Sync() SScriptFlushAsyncSync {
	return SScriptFlushAsyncSync{cf: c.cf, cs: append(c.cs, "SYNC")}
}

func (c SScriptFlush) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ScriptFlush() (c SScriptFlush) {
	c.cs = append(b.get(), "SCRIPT", "FLUSH")
	c.ks = initSlot
	return
}

type SScriptFlushAsyncAsync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptFlushAsyncAsync) Build() SCompleted {
	return SCompleted(c)
}

type SScriptFlushAsyncSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptFlushAsyncSync) Build() SCompleted {
	return SCompleted(c)
}

type SScriptKill struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptKill) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ScriptKill() (c SScriptKill) {
	c.cs = append(b.get(), "SCRIPT", "KILL")
	c.ks = initSlot
	return
}

type SScriptLoad struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptLoad) Script(Script string) SScriptLoadScript {
	return SScriptLoadScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *SBuilder) ScriptLoad() (c SScriptLoad) {
	c.cs = append(b.get(), "SCRIPT", "LOAD")
	c.ks = initSlot
	return
}

type SScriptLoadScript struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SScriptLoadScript) Build() SCompleted {
	return SCompleted(c)
}

type SSdiff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSdiff) Key(Key ...string) SSdiffKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Sdiff() (c SSdiff) {
	c.cs = append(b.get(), "SDIFF")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSdiffKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSdiffKey) Key(Key ...string) SSdiffKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSdiffKey) Build() SCompleted {
	return SCompleted(c)
}

type SSdiffstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSdiffstore) Destination(Destination string) SSdiffstoreDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Sdiffstore() (c SSdiffstore) {
	c.cs = append(b.get(), "SDIFFSTORE")
	c.ks = initSlot
	return
}

type SSdiffstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSdiffstoreDestination) Key(Key ...string) SSdiffstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SSdiffstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSdiffstoreKey) Key(Key ...string) SSdiffstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSdiffstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSelect struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSelect) Index(Index int64) SSelectIndex {
	return SSelectIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

func (b *SBuilder) Select() (c SSelect) {
	c.cs = append(b.get(), "SELECT")
	c.ks = initSlot
	return
}

type SSelectIndex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSelectIndex) Build() SCompleted {
	return SCompleted(c)
}

type SSet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSet) Key(Key string) SSetKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Set() (c SSet) {
	c.cs = append(b.get(), "SET")
	c.ks = initSlot
	return
}

type SSetConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetConditionNx) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SSetConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetConditionXx) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationEx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetExpirationEx) Nx() SSetConditionNx {
	return SSetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SSetExpirationEx) Xx() SSetConditionXx {
	return SSetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SSetExpirationEx) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetExpirationEx) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationExat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetExpirationExat) Nx() SSetConditionNx {
	return SSetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SSetExpirationExat) Xx() SSetConditionXx {
	return SSetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SSetExpirationExat) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetExpirationExat) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationKeepttl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetExpirationKeepttl) Nx() SSetConditionNx {
	return SSetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SSetExpirationKeepttl) Xx() SSetConditionXx {
	return SSetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SSetExpirationKeepttl) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetExpirationKeepttl) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationPx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetExpirationPx) Nx() SSetConditionNx {
	return SSetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SSetExpirationPx) Xx() SSetConditionXx {
	return SSetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SSetExpirationPx) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetExpirationPx) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationPxat struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetExpirationPxat) Nx() SSetConditionNx {
	return SSetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SSetExpirationPxat) Xx() SSetConditionXx {
	return SSetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SSetExpirationPxat) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetExpirationPxat) Build() SCompleted {
	return SCompleted(c)
}

type SSetGetGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetGetGet) Build() SCompleted {
	return SCompleted(c)
}

type SSetKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetKey) Value(Value string) SSetValue {
	return SSetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetValue) Ex(Seconds int64) SSetExpirationEx {
	return SSetExpirationEx{cf: c.cf, cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10))}
}

func (c SSetValue) Px(Milliseconds int64) SSetExpirationPx {
	return SSetExpirationPx{cf: c.cf, cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10))}
}

func (c SSetValue) Exat(Timestamp int64) SSetExpirationExat {
	return SSetExpirationExat{cf: c.cf, cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10))}
}

func (c SSetValue) Pxat(Millisecondstimestamp int64) SSetExpirationPxat {
	return SSetExpirationPxat{cf: c.cf, cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10))}
}

func (c SSetValue) Keepttl() SSetExpirationKeepttl {
	return SSetExpirationKeepttl{cf: c.cf, cs: append(c.cs, "KEEPTTL")}
}

func (c SSetValue) Nx() SSetConditionNx {
	return SSetConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SSetValue) Xx() SSetConditionXx {
	return SSetConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SSetValue) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetbit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetbit) Key(Key string) SSetbitKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setbit() (c SSetbit) {
	c.cs = append(b.get(), "SETBIT")
	c.ks = initSlot
	return
}

type SSetbitKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetbitKey) Offset(Offset int64) SSetbitOffset {
	return SSetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SSetbitOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetbitOffset) Value(Value int64) SSetbitValue {
	return SSetbitValue{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Value, 10))}
}

type SSetbitValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetbitValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetex) Key(Key string) SSetexKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setex() (c SSetex) {
	c.cs = append(b.get(), "SETEX")
	c.ks = initSlot
	return
}

type SSetexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetexKey) Seconds(Seconds int64) SSetexSeconds {
	return SSetexSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SSetexSeconds struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetexSeconds) Value(Value string) SSetexValue {
	return SSetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetexValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetexValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetnx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetnx) Key(Key string) SSetnxKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setnx() (c SSetnx) {
	c.cs = append(b.get(), "SETNX")
	c.ks = initSlot
	return
}

type SSetnxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetnxKey) Value(Value string) SSetnxValue {
	return SSetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetnxValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetnxValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetrange) Key(Key string) SSetrangeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setrange() (c SSetrange) {
	c.cs = append(b.get(), "SETRANGE")
	c.ks = initSlot
	return
}

type SSetrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetrangeKey) Offset(Offset int64) SSetrangeOffset {
	return SSetrangeOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SSetrangeOffset struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetrangeOffset) Value(Value string) SSetrangeValue {
	return SSetrangeValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetrangeValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSetrangeValue) Build() SCompleted {
	return SCompleted(c)
}

type SShutdown struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SShutdown) Nosave() SShutdownSaveModeNosave {
	return SShutdownSaveModeNosave{cf: c.cf, cs: append(c.cs, "NOSAVE")}
}

func (c SShutdown) Save() SShutdownSaveModeSave {
	return SShutdownSaveModeSave{cf: c.cf, cs: append(c.cs, "SAVE")}
}

func (c SShutdown) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Shutdown() (c SShutdown) {
	c.cs = append(b.get(), "SHUTDOWN")
	c.ks = initSlot
	return
}

type SShutdownSaveModeNosave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SShutdownSaveModeNosave) Build() SCompleted {
	return SCompleted(c)
}

type SShutdownSaveModeSave struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SShutdownSaveModeSave) Build() SCompleted {
	return SCompleted(c)
}

type SSinter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSinter) Key(Key ...string) SSinterKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Sinter() (c SSinter) {
	c.cs = append(b.get(), "SINTER")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSinterKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSinterKey) Key(Key ...string) SSinterKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSinterKey) Build() SCompleted {
	return SCompleted(c)
}

type SSintercard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSintercard) Key(Key ...string) SSintercardKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Sintercard() (c SSintercard) {
	c.cs = append(b.get(), "SINTERCARD")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSintercardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSintercardKey) Key(Key ...string) SSintercardKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSintercardKey) Build() SCompleted {
	return SCompleted(c)
}

type SSinterstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSinterstore) Destination(Destination string) SSinterstoreDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Sinterstore() (c SSinterstore) {
	c.cs = append(b.get(), "SINTERSTORE")
	c.ks = initSlot
	return
}

type SSinterstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSinterstoreDestination) Key(Key ...string) SSinterstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SSinterstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSinterstoreKey) Key(Key ...string) SSinterstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSinterstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSismember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSismember) Key(Key string) SSismemberKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Sismember() (c SSismember) {
	c.cs = append(b.get(), "SISMEMBER")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSismemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSismemberKey) Member(Member string) SSismemberMember {
	return SSismemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c SSismemberKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSismemberMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSismemberMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SSismemberMember) Cache() SCacheable {
	return SCacheable(c)
}

type SSlaveof struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSlaveof) Host(Host string) SSlaveofHost {
	return SSlaveofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *SBuilder) Slaveof() (c SSlaveof) {
	c.cs = append(b.get(), "SLAVEOF")
	c.ks = initSlot
	return
}

type SSlaveofHost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSlaveofHost) Port(Port string) SSlaveofPort {
	return SSlaveofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type SSlaveofPort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSlaveofPort) Build() SCompleted {
	return SCompleted(c)
}

type SSlowlog struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSlowlog) Subcommand(Subcommand string) SSlowlogSubcommand {
	return SSlowlogSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *SBuilder) Slowlog() (c SSlowlog) {
	c.cs = append(b.get(), "SLOWLOG")
	c.ks = initSlot
	return
}

type SSlowlogArgument struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSlowlogArgument) Build() SCompleted {
	return SCompleted(c)
}

type SSlowlogSubcommand struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSlowlogSubcommand) Argument(Argument string) SSlowlogArgument {
	return SSlowlogArgument{cf: c.cf, cs: append(c.cs, Argument)}
}

func (c SSlowlogSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SSmembers struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmembers) Key(Key string) SSmembersKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSmembersKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Smembers() (c SSmembers) {
	c.cs = append(b.get(), "SMEMBERS")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSmembersKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmembersKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SSmembersKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSmismember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmismember) Key(Key string) SSmismemberKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSmismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Smismember() (c SSmismember) {
	c.cs = append(b.get(), "SMISMEMBER")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSmismemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmismemberKey) Member(Member ...string) SSmismemberMember {
	return SSmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SSmismemberKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSmismemberMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmismemberMember) Member(Member ...string) SSmismemberMember {
	return SSmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SSmismemberMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SSmismemberMember) Cache() SCacheable {
	return SCacheable(c)
}

type SSmove struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmove) Source(Source string) SSmoveSource {
	s := slot(Source)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Smove() (c SSmove) {
	c.cs = append(b.get(), "SMOVE")
	c.ks = initSlot
	return
}

type SSmoveDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmoveDestination) Member(Member string) SSmoveMember {
	return SSmoveMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SSmoveMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmoveMember) Build() SCompleted {
	return SCompleted(c)
}

type SSmoveSource struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSmoveSource) Destination(Destination string) SSmoveDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SSort struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSort) Key(Key string) SSortKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Sort() (c SSort) {
	c.cs = append(b.get(), "SORT")
	c.ks = initSlot
	return
}

type SSortBy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortBy) Limit(Offset int64, Count int64) SSortLimit {
	return SSortLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SSortBy) Get(Pattern ...string) SSortGet {
	c.cs = append(c.cs, "GET")
	return SSortGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SSortBy) Asc() SSortOrderAsc {
	return SSortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortBy) Desc() SSortOrderDesc {
	return SSortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortBy) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortBy) Store(Destination string) SSortStore {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortBy) Build() SCompleted {
	return SCompleted(c)
}

type SSortGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortGet) Asc() SSortOrderAsc {
	return SSortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortGet) Desc() SSortOrderDesc {
	return SSortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortGet) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortGet) Store(Destination string) SSortStore {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortGet) Get(Get ...string) SSortGet {
	return SSortGet{cf: c.cf, cs: append(c.cs, Get...)}
}

func (c SSortGet) Build() SCompleted {
	return SCompleted(c)
}

type SSortKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortKey) By(Pattern string) SSortBy {
	return SSortBy{cf: c.cf, cs: append(c.cs, "BY", Pattern)}
}

func (c SSortKey) Limit(Offset int64, Count int64) SSortLimit {
	return SSortLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SSortKey) Get(Pattern ...string) SSortGet {
	c.cs = append(c.cs, "GET")
	return SSortGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SSortKey) Asc() SSortOrderAsc {
	return SSortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortKey) Desc() SSortOrderDesc {
	return SSortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortKey) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortKey) Store(Destination string) SSortStore {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortKey) Build() SCompleted {
	return SCompleted(c)
}

type SSortLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortLimit) Get(Pattern ...string) SSortGet {
	c.cs = append(c.cs, "GET")
	return SSortGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SSortLimit) Asc() SSortOrderAsc {
	return SSortOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortLimit) Desc() SSortOrderDesc {
	return SSortOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortLimit) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortLimit) Store(Destination string) SSortStore {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortLimit) Build() SCompleted {
	return SCompleted(c)
}

type SSortOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortOrderAsc) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortOrderAsc) Store(Destination string) SSortStore {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SSortOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortOrderDesc) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortOrderDesc) Store(Destination string) SSortStore {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SSortRo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRo) Key(Key string) SSortRoKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) SortRo() (c SSortRo) {
	c.cs = append(b.get(), "SORT_RO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSortRoBy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRoBy) Limit(Offset int64, Count int64) SSortRoLimit {
	return SSortRoLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SSortRoBy) Get(Pattern ...string) SSortRoGet {
	c.cs = append(c.cs, "GET")
	return SSortRoGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SSortRoBy) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortRoBy) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortRoBy) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortRoBy) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoBy) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoGet struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRoGet) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortRoGet) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortRoGet) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortRoGet) Get(Get ...string) SSortRoGet {
	return SSortRoGet{cf: c.cf, cs: append(c.cs, Get...)}
}

func (c SSortRoGet) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoGet) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRoKey) By(Pattern string) SSortRoBy {
	return SSortRoBy{cf: c.cf, cs: append(c.cs, "BY", Pattern)}
}

func (c SSortRoKey) Limit(Offset int64, Count int64) SSortRoLimit {
	return SSortRoLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SSortRoKey) Get(Pattern ...string) SSortRoGet {
	c.cs = append(c.cs, "GET")
	return SSortRoGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SSortRoKey) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortRoKey) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortRoKey) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortRoKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRoLimit) Get(Pattern ...string) SSortRoGet {
	c.cs = append(c.cs, "GET")
	return SSortRoGet{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SSortRoLimit) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SSortRoLimit) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SSortRoLimit) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortRoLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoOrderAsc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRoOrderAsc) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortRoOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoOrderDesc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRoOrderDesc) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoSortingAlpha struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortRoSortingAlpha) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoSortingAlpha) Cache() SCacheable {
	return SCacheable(c)
}

type SSortSortingAlpha struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortSortingAlpha) Store(Destination string) SSortStore {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortSortingAlpha) Build() SCompleted {
	return SCompleted(c)
}

type SSortStore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSortStore) Build() SCompleted {
	return SCompleted(c)
}

type SSpop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSpop) Key(Key string) SSpopKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSpopKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Spop() (c SSpop) {
	c.cs = append(b.get(), "SPOP")
	c.ks = initSlot
	return
}

type SSpopCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SSpopKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSpopKey) Count(Count int64) SSpopCount {
	return SSpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SSpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SSrandmember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSrandmember) Key(Key string) SSrandmemberKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Srandmember() (c SSrandmember) {
	c.cs = append(b.get(), "SRANDMEMBER")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSrandmemberCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSrandmemberCount) Build() SCompleted {
	return SCompleted(c)
}

type SSrandmemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSrandmemberKey) Count(Count int64) SSrandmemberCount {
	return SSrandmemberCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SSrandmemberKey) Build() SCompleted {
	return SCompleted(c)
}

type SSrem struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSrem) Key(Key string) SSremKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Srem() (c SSrem) {
	c.cs = append(b.get(), "SREM")
	c.ks = initSlot
	return
}

type SSremKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSremKey) Member(Member ...string) SSremMember {
	return SSremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SSremMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSremMember) Member(Member ...string) SSremMember {
	return SSremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SSremMember) Build() SCompleted {
	return SCompleted(c)
}

type SSscan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSscan) Key(Key string) SSscanKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Sscan() (c SSscan) {
	c.cs = append(b.get(), "SSCAN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSscanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSscanCount) Build() SCompleted {
	return SCompleted(c)
}

type SSscanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSscanCursor) Match(Pattern string) SSscanMatch {
	return SSscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SSscanCursor) Count(Count int64) SSscanCount {
	return SSscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SSscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SSscanKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSscanKey) Cursor(Cursor int64) SSscanCursor {
	return SSscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SSscanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSscanMatch) Count(Count int64) SSscanCount {
	return SSscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SSscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SStralgo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SStralgo) Lcs() SStralgoAlgorithmLcs {
	return SStralgoAlgorithmLcs{cf: c.cf, cs: append(c.cs, "LCS")}
}

func (b *SBuilder) Stralgo() (c SStralgo) {
	c.cs = append(b.get(), "STRALGO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SStralgoAlgoSpecificArgument struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SStralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) SStralgoAlgoSpecificArgument {
	return SStralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

func (c SStralgoAlgoSpecificArgument) Build() SCompleted {
	return SCompleted(c)
}

type SStralgoAlgorithmLcs struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SStralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) SStralgoAlgoSpecificArgument {
	return SStralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

type SStrlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SStrlen) Key(Key string) SStrlenKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SStrlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Strlen() (c SStrlen) {
	c.cs = append(b.get(), "STRLEN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SStrlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SStrlenKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SStrlenKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSubscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSubscribe) Channel(Channel ...string) SSubscribeChannel {
	return SSubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (b *SBuilder) Subscribe() (c SSubscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	c.cf = noRetTag
	c.ks = initSlot
	return
}

type SSubscribeChannel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSubscribeChannel) Channel(Channel ...string) SSubscribeChannel {
	return SSubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c SSubscribeChannel) Build() SCompleted {
	return SCompleted(c)
}

type SSunion struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSunion) Key(Key ...string) SSunionKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Sunion() (c SSunion) {
	c.cs = append(b.get(), "SUNION")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SSunionKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSunionKey) Key(Key ...string) SSunionKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSunionKey) Build() SCompleted {
	return SCompleted(c)
}

type SSunionstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSunionstore) Destination(Destination string) SSunionstoreDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SSunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Sunionstore() (c SSunionstore) {
	c.cs = append(b.get(), "SUNIONSTORE")
	c.ks = initSlot
	return
}

type SSunionstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSunionstoreDestination) Key(Key ...string) SSunionstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SSunionstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSunionstoreKey) Key(Key ...string) SSunionstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SSunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSunionstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSwapdb struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSwapdb) Index1(Index1 int64) SSwapdbIndex1 {
	return SSwapdbIndex1{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index1, 10))}
}

func (b *SBuilder) Swapdb() (c SSwapdb) {
	c.cs = append(b.get(), "SWAPDB")
	c.ks = initSlot
	return
}

type SSwapdbIndex1 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSwapdbIndex1) Index2(Index2 int64) SSwapdbIndex2 {
	return SSwapdbIndex2{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index2, 10))}
}

type SSwapdbIndex2 struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSwapdbIndex2) Build() SCompleted {
	return SCompleted(c)
}

type SSync struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SSync) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Sync() (c SSync) {
	c.cs = append(b.get(), "SYNC")
	c.ks = initSlot
	return
}

type STime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c STime) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Time() (c STime) {
	c.cs = append(b.get(), "TIME")
	c.ks = initSlot
	return
}

type STouch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c STouch) Key(Key ...string) STouchKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return STouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Touch() (c STouch) {
	c.cs = append(b.get(), "TOUCH")
	c.cf = readonly
	c.ks = initSlot
	return
}

type STouchKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c STouchKey) Key(Key ...string) STouchKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return STouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c STouchKey) Build() SCompleted {
	return SCompleted(c)
}

type STtl struct {
	cs []string
	cf uint16
	ks uint16
}

func (c STtl) Key(Key string) STtlKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return STtlKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Ttl() (c STtl) {
	c.cs = append(b.get(), "TTL")
	c.cf = readonly
	c.ks = initSlot
	return
}

type STtlKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c STtlKey) Build() SCompleted {
	return SCompleted(c)
}

func (c STtlKey) Cache() SCacheable {
	return SCacheable(c)
}

type SType struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SType) Key(Key string) STypeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return STypeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Type() (c SType) {
	c.cs = append(b.get(), "TYPE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type STypeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c STypeKey) Build() SCompleted {
	return SCompleted(c)
}

func (c STypeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SUnlink struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SUnlink) Key(Key ...string) SUnlinkKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SUnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Unlink() (c SUnlink) {
	c.cs = append(b.get(), "UNLINK")
	c.ks = initSlot
	return
}

type SUnlinkKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SUnlinkKey) Key(Key ...string) SUnlinkKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SUnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SUnlinkKey) Build() SCompleted {
	return SCompleted(c)
}

type SUnsubscribe struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SUnsubscribe) Channel(Channel ...string) SUnsubscribeChannel {
	return SUnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c SUnsubscribe) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Unsubscribe() (c SUnsubscribe) {
	c.cs = append(b.get(), "UNSUBSCRIBE")
	c.cf = noRetTag
	c.ks = initSlot
	return
}

type SUnsubscribeChannel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SUnsubscribeChannel) Channel(Channel ...string) SUnsubscribeChannel {
	return SUnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c SUnsubscribeChannel) Build() SCompleted {
	return SCompleted(c)
}

type SUnwatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SUnwatch) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Unwatch() (c SUnwatch) {
	c.cs = append(b.get(), "UNWATCH")
	c.ks = initSlot
	return
}

type SWait struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SWait) Numreplicas(Numreplicas int64) SWaitNumreplicas {
	return SWaitNumreplicas{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numreplicas, 10))}
}

func (b *SBuilder) Wait() (c SWait) {
	c.cs = append(b.get(), "WAIT")
	c.cf = blockTag
	c.ks = initSlot
	return
}

type SWaitNumreplicas struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SWaitNumreplicas) Timeout(Timeout int64) SWaitTimeout {
	return SWaitTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type SWaitTimeout struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SWaitTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SWatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SWatch) Key(Key ...string) SWatchKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SWatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Watch() (c SWatch) {
	c.cs = append(b.get(), "WATCH")
	c.ks = initSlot
	return
}

type SWatchKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SWatchKey) Key(Key ...string) SWatchKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SWatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SWatchKey) Build() SCompleted {
	return SCompleted(c)
}

type SXack struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXack) Key(Key string) SXackKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXackKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xack() (c SXack) {
	c.cs = append(b.get(), "XACK")
	c.ks = initSlot
	return
}

type SXackGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXackGroup) Id(Id ...string) SXackId {
	return SXackId{cf: c.cf, cs: append(c.cs, Id...)}
}

type SXackId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXackId) Id(Id ...string) SXackId {
	return SXackId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXackId) Build() SCompleted {
	return SCompleted(c)
}

type SXackKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXackKey) Group(Group string) SXackGroup {
	return SXackGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXadd) Key(Key string) SXaddKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xadd() (c SXadd) {
	c.cs = append(b.get(), "XADD")
	c.ks = initSlot
	return
}

type SXaddFieldValue struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddFieldValue) FieldValue(Field string, Value string) SXaddFieldValue {
	return SXaddFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c SXaddFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SXaddId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddId) FieldValue() SXaddFieldValue {
	return SXaddFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type SXaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddKey) Nomkstream() SXaddNomkstream {
	return SXaddNomkstream{cf: c.cf, cs: append(c.cs, "NOMKSTREAM")}
}

func (c SXaddKey) Maxlen() SXaddTrimStrategyMaxlen {
	return SXaddTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c SXaddKey) Minid() SXaddTrimStrategyMinid {
	return SXaddTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

func (c SXaddKey) Id(Id string) SXaddId {
	return SXaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXaddNomkstream struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddNomkstream) Maxlen() SXaddTrimStrategyMaxlen {
	return SXaddTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c SXaddNomkstream) Minid() SXaddTrimStrategyMinid {
	return SXaddTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

func (c SXaddNomkstream) Id(Id string) SXaddId {
	return SXaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXaddTrimLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddTrimLimit) Id(Id string) SXaddId {
	return SXaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXaddTrimOperatorAlmost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddTrimOperatorAlmost) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimOperatorExact struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddTrimOperatorExact) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimStrategyMaxlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddTrimStrategyMaxlen) Exact() SXaddTrimOperatorExact {
	return SXaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXaddTrimStrategyMaxlen) Almost() SXaddTrimOperatorAlmost {
	return SXaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXaddTrimStrategyMaxlen) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimStrategyMinid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddTrimStrategyMinid) Exact() SXaddTrimOperatorExact {
	return SXaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXaddTrimStrategyMinid) Almost() SXaddTrimOperatorAlmost {
	return SXaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXaddTrimStrategyMinid) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimThreshold struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXaddTrimThreshold) Limit(Count int64) SXaddTrimLimit {
	return SXaddTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c SXaddTrimThreshold) Id(Id string) SXaddId {
	return SXaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXautoclaim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaim) Key(Key string) SXautoclaimKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXautoclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xautoclaim() (c SXautoclaim) {
	c.cs = append(b.get(), "XAUTOCLAIM")
	c.ks = initSlot
	return
}

type SXautoclaimConsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaimConsumer) MinIdleTime(MinIdleTime string) SXautoclaimMinIdleTime {
	return SXautoclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type SXautoclaimCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaimCount) Justid() SXautoclaimJustidJustid {
	return SXautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXautoclaimCount) Build() SCompleted {
	return SCompleted(c)
}

type SXautoclaimGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaimGroup) Consumer(Consumer string) SXautoclaimConsumer {
	return SXautoclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type SXautoclaimJustidJustid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaimJustidJustid) Build() SCompleted {
	return SCompleted(c)
}

type SXautoclaimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaimKey) Group(Group string) SXautoclaimGroup {
	return SXautoclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXautoclaimMinIdleTime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaimMinIdleTime) Start(Start string) SXautoclaimStart {
	return SXautoclaimStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXautoclaimStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXautoclaimStart) Count(Count int64) SXautoclaimCount {
	return SXautoclaimCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXautoclaimStart) Justid() SXautoclaimJustidJustid {
	return SXautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXautoclaimStart) Build() SCompleted {
	return SCompleted(c)
}

type SXclaim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaim) Key(Key string) SXclaimKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xclaim() (c SXclaim) {
	c.cs = append(b.get(), "XCLAIM")
	c.ks = initSlot
	return
}

type SXclaimConsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimConsumer) MinIdleTime(MinIdleTime string) SXclaimMinIdleTime {
	return SXclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type SXclaimForceForce struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimForceForce) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXclaimForceForce) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimGroup) Consumer(Consumer string) SXclaimConsumer {
	return SXclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type SXclaimId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimId) Idle(Ms int64) SXclaimIdle {
	return SXclaimIdle{cf: c.cf, cs: append(c.cs, "IDLE", strconv.FormatInt(Ms, 10))}
}

func (c SXclaimId) Time(MsUnixTime int64) SXclaimTime {
	return SXclaimTime{cf: c.cf, cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10))}
}

func (c SXclaimId) Retrycount(Count int64) SXclaimRetrycount {
	return SXclaimRetrycount{cf: c.cf, cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c SXclaimId) Force() SXclaimForceForce {
	return SXclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c SXclaimId) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXclaimId) Id(Id ...string) SXclaimId {
	return SXclaimId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXclaimId) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimIdle struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimIdle) Time(MsUnixTime int64) SXclaimTime {
	return SXclaimTime{cf: c.cf, cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10))}
}

func (c SXclaimIdle) Retrycount(Count int64) SXclaimRetrycount {
	return SXclaimRetrycount{cf: c.cf, cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c SXclaimIdle) Force() SXclaimForceForce {
	return SXclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c SXclaimIdle) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXclaimIdle) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimJustidJustid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimJustidJustid) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimKey) Group(Group string) SXclaimGroup {
	return SXclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXclaimMinIdleTime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimMinIdleTime) Id(Id ...string) SXclaimId {
	return SXclaimId{cf: c.cf, cs: append(c.cs, Id...)}
}

type SXclaimRetrycount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimRetrycount) Force() SXclaimForceForce {
	return SXclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c SXclaimRetrycount) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXclaimRetrycount) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimTime struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXclaimTime) Retrycount(Count int64) SXclaimRetrycount {
	return SXclaimRetrycount{cf: c.cf, cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c SXclaimTime) Force() SXclaimForceForce {
	return SXclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c SXclaimTime) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXclaimTime) Build() SCompleted {
	return SCompleted(c)
}

type SXdel struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXdel) Key(Key string) SXdelKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xdel() (c SXdel) {
	c.cs = append(b.get(), "XDEL")
	c.ks = initSlot
	return
}

type SXdelId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXdelId) Id(Id ...string) SXdelId {
	return SXdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXdelId) Build() SCompleted {
	return SCompleted(c)
}

type SXdelKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXdelKey) Id(Id ...string) SXdelId {
	return SXdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

type SXgroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroup) Create(Key string, Groupname string) SXgroupCreateCreate {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupCreateCreate{cf: c.cf, cs: append(c.cs, "CREATE", Key, Groupname)}
}

func (c SXgroup) Setid(Key string, Groupname string) SXgroupSetidSetid {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c SXgroup) Destroy(Key string, Groupname string) SXgroupDestroy {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroup) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroup) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (b *SBuilder) Xgroup() (c SXgroup) {
	c.cs = append(b.get(), "XGROUP")
	c.ks = initSlot
	return
}

type SXgroupCreateCreate struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupCreateCreate) Id(Id string) SXgroupCreateId {
	return SXgroupCreateId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXgroupCreateId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupCreateId) Mkstream() SXgroupCreateMkstream {
	return SXgroupCreateMkstream{cf: c.cf, cs: append(c.cs, "MKSTREAM")}
}

func (c SXgroupCreateId) Setid(Key string, Groupname string) SXgroupSetidSetid {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c SXgroupCreateId) Destroy(Key string, Groupname string) SXgroupDestroy {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroupCreateId) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupCreateId) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type SXgroupCreateMkstream struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupCreateMkstream) Setid(Key string, Groupname string) SXgroupSetidSetid {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c SXgroupCreateMkstream) Destroy(Key string, Groupname string) SXgroupDestroy {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroupCreateMkstream) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupCreateMkstream) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type SXgroupCreateconsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupCreateconsumer) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupDelconsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupDelconsumer) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupDestroy struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupDestroy) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupDestroy) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupDestroy) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupSetidId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupSetidId) Destroy(Key string, Groupname string) SXgroupDestroy {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroupSetidId) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupSetidId) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupSetidId) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupSetidSetid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXgroupSetidSetid) Id(Id string) SXgroupSetidId {
	return SXgroupSetidId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXinfo struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXinfo) Consumers(Key string, Groupname string) SXinfoConsumers {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXinfoConsumers{cf: c.cf, cs: append(c.cs, "CONSUMERS", Key, Groupname)}
}

func (c SXinfo) Groups(Key string) SXinfoGroups {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXinfoGroups{cf: c.cf, cs: append(c.cs, "GROUPS", Key)}
}

func (c SXinfo) Stream(Key string) SXinfoStream {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c SXinfo) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c SXinfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Xinfo() (c SXinfo) {
	c.cs = append(b.get(), "XINFO")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SXinfoConsumers struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXinfoConsumers) Groups(Key string) SXinfoGroups {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXinfoGroups{cf: c.cf, cs: append(c.cs, "GROUPS", Key)}
}

func (c SXinfoConsumers) Stream(Key string) SXinfoStream {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c SXinfoConsumers) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c SXinfoConsumers) Build() SCompleted {
	return SCompleted(c)
}

type SXinfoGroups struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXinfoGroups) Stream(Key string) SXinfoStream {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c SXinfoGroups) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c SXinfoGroups) Build() SCompleted {
	return SCompleted(c)
}

type SXinfoHelpHelp struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXinfoHelpHelp) Build() SCompleted {
	return SCompleted(c)
}

type SXinfoStream struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXinfoStream) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c SXinfoStream) Build() SCompleted {
	return SCompleted(c)
}

type SXlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXlen) Key(Key string) SXlenKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xlen() (c SXlen) {
	c.cs = append(b.get(), "XLEN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SXlenKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXlenKey) Build() SCompleted {
	return SCompleted(c)
}

type SXpending struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpending) Key(Key string) SXpendingKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXpendingKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xpending() (c SXpending) {
	c.cs = append(b.get(), "XPENDING")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SXpendingFiltersConsumer struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpendingFiltersConsumer) Build() SCompleted {
	return SCompleted(c)
}

type SXpendingFiltersCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpendingFiltersCount) Consumer(Consumer string) SXpendingFiltersConsumer {
	return SXpendingFiltersConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

func (c SXpendingFiltersCount) Build() SCompleted {
	return SCompleted(c)
}

type SXpendingFiltersEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpendingFiltersEnd) Count(Count int64) SXpendingFiltersCount {
	return SXpendingFiltersCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type SXpendingFiltersIdle struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpendingFiltersIdle) Start(Start string) SXpendingFiltersStart {
	return SXpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXpendingFiltersStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpendingFiltersStart) End(End string) SXpendingFiltersEnd {
	return SXpendingFiltersEnd{cf: c.cf, cs: append(c.cs, End)}
}

type SXpendingGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpendingGroup) Idle(MinIdleTime int64) SXpendingFiltersIdle {
	return SXpendingFiltersIdle{cf: c.cf, cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10))}
}

func (c SXpendingGroup) Start(Start string) SXpendingFiltersStart {
	return SXpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXpendingKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXpendingKey) Group(Group string) SXpendingGroup {
	return SXpendingGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrange) Key(Key string) SXrangeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xrange() (c SXrange) {
	c.cs = append(b.get(), "XRANGE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SXrangeCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrangeCount) Build() SCompleted {
	return SCompleted(c)
}

type SXrangeEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrangeEnd) Count(Count int64) SXrangeCount {
	return SXrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXrangeEnd) Build() SCompleted {
	return SCompleted(c)
}

type SXrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrangeKey) Start(Start string) SXrangeStart {
	return SXrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrangeStart) End(End string) SXrangeEnd {
	return SXrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type SXread struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXread) Count(Count int64) SXreadCount {
	return SXreadCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXread) Block(Milliseconds int64) SXreadBlock {
	c.cf = blockTag
	return SXreadBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c SXread) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

func (b *SBuilder) Xread() (c SXread) {
	c.cs = append(b.get(), "XREAD")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SXreadBlock struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadBlock) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadCount) Block(Milliseconds int64) SXreadBlock {
	c.cf = blockTag
	return SXreadBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c SXreadCount) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadId) Id(Id ...string) SXreadId {
	return SXreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadId) Build() SCompleted {
	return SCompleted(c)
}

type SXreadKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadKey) Id(Id ...string) SXreadId {
	return SXreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadKey) Key(Key ...string) SXreadKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SXreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXreadStreamsStreams struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadStreamsStreams) Key(Key ...string) SXreadKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SXreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXreadgroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroup) Group(Group string, Consumer string) SXreadgroupGroup {
	return SXreadgroupGroup{cf: c.cf, cs: append(c.cs, "GROUP", Group, Consumer)}
}

func (b *SBuilder) Xreadgroup() (c SXreadgroup) {
	c.cs = append(b.get(), "XREADGROUP")
	c.ks = initSlot
	return
}

type SXreadgroupBlock struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroupBlock) Noack() SXreadgroupNoackNoack {
	return SXreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c SXreadgroupBlock) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadgroupCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroupCount) Block(Milliseconds int64) SXreadgroupBlock {
	c.cf = blockTag
	return SXreadgroupBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c SXreadgroupCount) Noack() SXreadgroupNoackNoack {
	return SXreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c SXreadgroupCount) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadgroupGroup struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroupGroup) Count(Count int64) SXreadgroupCount {
	return SXreadgroupCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXreadgroupGroup) Block(Milliseconds int64) SXreadgroupBlock {
	c.cf = blockTag
	return SXreadgroupBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c SXreadgroupGroup) Noack() SXreadgroupNoackNoack {
	return SXreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c SXreadgroupGroup) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadgroupId struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroupId) Id(Id ...string) SXreadgroupId {
	return SXreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadgroupId) Build() SCompleted {
	return SCompleted(c)
}

type SXreadgroupKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroupKey) Id(Id ...string) SXreadgroupId {
	return SXreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadgroupKey) Key(Key ...string) SXreadgroupKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SXreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXreadgroupNoackNoack struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroupNoackNoack) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadgroupStreamsStreams struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXreadgroupStreamsStreams) Key(Key ...string) SXreadgroupKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SXreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXrevrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrevrange) Key(Key string) SXrevrangeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xrevrange() (c SXrevrange) {
	c.cs = append(b.get(), "XREVRANGE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SXrevrangeCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrevrangeCount) Build() SCompleted {
	return SCompleted(c)
}

type SXrevrangeEnd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrevrangeEnd) Start(Start string) SXrevrangeStart {
	return SXrevrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXrevrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrevrangeKey) End(End string) SXrevrangeEnd {
	return SXrevrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type SXrevrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXrevrangeStart) Count(Count int64) SXrevrangeCount {
	return SXrevrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXrevrangeStart) Build() SCompleted {
	return SCompleted(c)
}

type SXtrim struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrim) Key(Key string) SXtrimKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SXtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xtrim() (c SXtrim) {
	c.cs = append(b.get(), "XTRIM")
	c.ks = initSlot
	return
}

type SXtrimKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrimKey) Maxlen() SXtrimTrimStrategyMaxlen {
	return SXtrimTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c SXtrimKey) Minid() SXtrimTrimStrategyMinid {
	return SXtrimTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

type SXtrimTrimLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrimTrimLimit) Build() SCompleted {
	return SCompleted(c)
}

type SXtrimTrimOperatorAlmost struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrimTrimOperatorAlmost) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimOperatorExact struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrimTrimOperatorExact) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimStrategyMaxlen struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrimTrimStrategyMaxlen) Exact() SXtrimTrimOperatorExact {
	return SXtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXtrimTrimStrategyMaxlen) Almost() SXtrimTrimOperatorAlmost {
	return SXtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXtrimTrimStrategyMaxlen) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimStrategyMinid struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrimTrimStrategyMinid) Exact() SXtrimTrimOperatorExact {
	return SXtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXtrimTrimStrategyMinid) Almost() SXtrimTrimOperatorAlmost {
	return SXtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXtrimTrimStrategyMinid) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimThreshold struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SXtrimTrimThreshold) Limit(Count int64) SXtrimTrimLimit {
	return SXtrimTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c SXtrimTrimThreshold) Build() SCompleted {
	return SCompleted(c)
}

type SZadd struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZadd) Key(Key string) SZaddKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zadd() (c SZadd) {
	c.cs = append(b.get(), "ZADD")
	c.ks = initSlot
	return
}

type SZaddChangeCh struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddChangeCh) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddChangeCh) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type SZaddComparisonGt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddComparisonGt) Ch() SZaddChangeCh {
	return SZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SZaddComparisonGt) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddComparisonGt) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type SZaddComparisonLt struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddComparisonLt) Ch() SZaddChangeCh {
	return SZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SZaddComparisonLt) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddComparisonLt) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type SZaddConditionNx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddConditionNx) Gt() SZaddComparisonGt {
	return SZaddComparisonGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c SZaddConditionNx) Lt() SZaddComparisonLt {
	return SZaddComparisonLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c SZaddConditionNx) Ch() SZaddChangeCh {
	return SZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SZaddConditionNx) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddConditionNx) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type SZaddConditionXx struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddConditionXx) Gt() SZaddComparisonGt {
	return SZaddComparisonGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c SZaddConditionXx) Lt() SZaddComparisonLt {
	return SZaddComparisonLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c SZaddConditionXx) Ch() SZaddChangeCh {
	return SZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SZaddConditionXx) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddConditionXx) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type SZaddIncrementIncr struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddIncrementIncr) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type SZaddKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddKey) Nx() SZaddConditionNx {
	return SZaddConditionNx{cf: c.cf, cs: append(c.cs, "NX")}
}

func (c SZaddKey) Xx() SZaddConditionXx {
	return SZaddConditionXx{cf: c.cf, cs: append(c.cs, "XX")}
}

func (c SZaddKey) Gt() SZaddComparisonGt {
	return SZaddComparisonGt{cf: c.cf, cs: append(c.cs, "GT")}
}

func (c SZaddKey) Lt() SZaddComparisonLt {
	return SZaddComparisonLt{cf: c.cf, cs: append(c.cs, "LT")}
}

func (c SZaddKey) Ch() SZaddChangeCh {
	return SZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SZaddKey) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddKey) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type SZaddScoreMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZaddScoreMember) ScoreMember(Score float64, Member string) SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member)}
}

func (c SZaddScoreMember) Build() SCompleted {
	return SCompleted(c)
}

type SZcard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZcard) Key(Key string) SZcardKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZcardKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zcard() (c SZcard) {
	c.cs = append(b.get(), "ZCARD")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZcardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZcardKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SZcardKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZcount) Key(Key string) SZcountKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zcount() (c SZcount) {
	c.cs = append(b.get(), "ZCOUNT")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZcountKey) Min(Min float64) SZcountMin {
	return SZcountMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c SZcountKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZcountMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZcountMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZcountMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZcountMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZcountMin) Max(Max float64) SZcountMax {
	return SZcountMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c SZcountMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZdiff struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiff) Numkeys(Numkeys int64) SZdiffNumkeys {
	return SZdiffNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Zdiff() (c SZdiff) {
	c.cs = append(b.get(), "ZDIFF")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZdiffKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiffKey) Withscores() SZdiffWithscoresWithscores {
	return SZdiffWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZdiffKey) Key(Key ...string) SZdiffKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZdiffKey) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiffNumkeys) Key(Key ...string) SZdiffKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZdiffWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiffWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiffstore) Destination(Destination string) SZdiffstoreDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Zdiffstore() (c SZdiffstore) {
	c.cs = append(b.get(), "ZDIFFSTORE")
	c.ks = initSlot
	return
}

type SZdiffstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiffstoreDestination) Numkeys(Numkeys int64) SZdiffstoreNumkeys {
	return SZdiffstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SZdiffstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiffstoreKey) Key(Key ...string) SZdiffstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZdiffstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffstoreNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZdiffstoreNumkeys) Key(Key ...string) SZdiffstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZincrby struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZincrby) Key(Key string) SZincrbyKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zincrby() (c SZincrby) {
	c.cs = append(b.get(), "ZINCRBY")
	c.ks = initSlot
	return
}

type SZincrbyIncrement struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZincrbyIncrement) Member(Member string) SZincrbyMember {
	return SZincrbyMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SZincrbyKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZincrbyKey) Increment(Increment int64) SZincrbyIncrement {
	return SZincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type SZincrbyMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZincrbyMember) Build() SCompleted {
	return SCompleted(c)
}

type SZinter struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinter) Numkeys(Numkeys int64) SZinterNumkeys {
	return SZinterNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Zinter() (c SZinter) {
	c.cs = append(b.get(), "ZINTER")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZinterAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterAggregateMax) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZinterAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZinterAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterAggregateMin) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZinterAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZinterAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterAggregateSum) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZinterAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZinterKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterKey) Weights(Weight ...int64) SZinterWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterWeights{cf: c.cf, cs: c.cs}
}

func (c SZinterKey) Sum() SZinterAggregateSum {
	return SZinterAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZinterKey) Min() SZinterAggregateMin {
	return SZinterAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZinterKey) Max() SZinterAggregateMax {
	return SZinterAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZinterKey) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZinterKey) Key(Key ...string) SZinterKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZinterKey) Build() SCompleted {
	return SCompleted(c)
}

type SZinterNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterNumkeys) Key(Key ...string) SZinterKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZinterWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterWeights) Sum() SZinterAggregateSum {
	return SZinterAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZinterWeights) Min() SZinterAggregateMin {
	return SZinterAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZinterWeights) Max() SZinterAggregateMax {
	return SZinterAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZinterWeights) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZinterWeights) Weights(Weights ...int64) SZinterWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterWeights{cf: c.cf, cs: c.cs}
}

func (c SZinterWeights) Build() SCompleted {
	return SCompleted(c)
}

type SZinterWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZintercard struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZintercard) Numkeys(Numkeys int64) SZintercardNumkeys {
	return SZintercardNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Zintercard() (c SZintercard) {
	c.cs = append(b.get(), "ZINTERCARD")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZintercardKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZintercardKey) Key(Key ...string) SZintercardKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZintercardKey) Build() SCompleted {
	return SCompleted(c)
}

type SZintercardNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZintercardNumkeys) Key(Key ...string) SZintercardKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZinterstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstore) Destination(Destination string) SZinterstoreDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Zinterstore() (c SZinterstore) {
	c.cs = append(b.get(), "ZINTERSTORE")
	c.ks = initSlot
	return
}

type SZinterstoreAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstoreAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstoreAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstoreAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstoreDestination) Numkeys(Numkeys int64) SZinterstoreNumkeys {
	return SZinterstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SZinterstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstoreKey) Weights(Weight ...int64) SZinterstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterstoreWeights{cf: c.cf, cs: c.cs}
}

func (c SZinterstoreKey) Sum() SZinterstoreAggregateSum {
	return SZinterstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZinterstoreKey) Min() SZinterstoreAggregateMin {
	return SZinterstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZinterstoreKey) Max() SZinterstoreAggregateMax {
	return SZinterstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZinterstoreKey) Key(Key ...string) SZinterstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZinterstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstoreNumkeys) Key(Key ...string) SZinterstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZinterstoreWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZinterstoreWeights) Sum() SZinterstoreAggregateSum {
	return SZinterstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZinterstoreWeights) Min() SZinterstoreAggregateMin {
	return SZinterstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZinterstoreWeights) Max() SZinterstoreAggregateMax {
	return SZinterstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZinterstoreWeights) Weights(Weights ...int64) SZinterstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterstoreWeights{cf: c.cf, cs: c.cs}
}

func (c SZinterstoreWeights) Build() SCompleted {
	return SCompleted(c)
}

type SZlexcount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZlexcount) Key(Key string) SZlexcountKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZlexcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zlexcount() (c SZlexcount) {
	c.cs = append(b.get(), "ZLEXCOUNT")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZlexcountKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZlexcountKey) Min(Min string) SZlexcountMin {
	return SZlexcountMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c SZlexcountKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZlexcountMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZlexcountMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZlexcountMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZlexcountMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZlexcountMin) Max(Max string) SZlexcountMax {
	return SZlexcountMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c SZlexcountMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZmscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZmscore) Key(Key string) SZmscoreKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZmscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zmscore() (c SZmscore) {
	c.cs = append(b.get(), "ZMSCORE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZmscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZmscoreKey) Member(Member ...string) SZmscoreMember {
	return SZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SZmscoreKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZmscoreMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZmscoreMember) Member(Member ...string) SZmscoreMember {
	return SZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SZmscoreMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZmscoreMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZpopmax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZpopmax) Key(Key string) SZpopmaxKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZpopmaxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zpopmax() (c SZpopmax) {
	c.cs = append(b.get(), "ZPOPMAX")
	c.ks = initSlot
	return
}

type SZpopmaxCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZpopmaxCount) Build() SCompleted {
	return SCompleted(c)
}

type SZpopmaxKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZpopmaxKey) Count(Count int64) SZpopmaxCount {
	return SZpopmaxCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SZpopmaxKey) Build() SCompleted {
	return SCompleted(c)
}

type SZpopmin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZpopmin) Key(Key string) SZpopminKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZpopminKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zpopmin() (c SZpopmin) {
	c.cs = append(b.get(), "ZPOPMIN")
	c.ks = initSlot
	return
}

type SZpopminCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZpopminCount) Build() SCompleted {
	return SCompleted(c)
}

type SZpopminKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZpopminKey) Count(Count int64) SZpopminCount {
	return SZpopminCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SZpopminKey) Build() SCompleted {
	return SCompleted(c)
}

type SZrandmember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrandmember) Key(Key string) SZrandmemberKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrandmember() (c SZrandmember) {
	c.cs = append(b.get(), "ZRANDMEMBER")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrandmemberKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrandmemberKey) Count(Count int64) SZrandmemberOptionsCount {
	return SZrandmemberOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type SZrandmemberOptionsCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrandmemberOptionsCount) Withscores() SZrandmemberOptionsWithscoresWithscores {
	return SZrandmemberOptionsWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrandmemberOptionsCount) Build() SCompleted {
	return SCompleted(c)
}

type SZrandmemberOptionsWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrandmemberOptionsWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrange) Key(Key string) SZrangeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrange() (c SZrange) {
	c.cs = append(b.get(), "ZRANGE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeKey) Min(Min string) SZrangeMin {
	return SZrangeMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c SZrangeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeLimit) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrangeLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeMax) Byscore() SZrangeSortbyByscore {
	return SZrangeSortbyByscore{cf: c.cf, cs: append(c.cs, "BYSCORE")}
}

func (c SZrangeMax) Bylex() SZrangeSortbyBylex {
	return SZrangeSortbyBylex{cf: c.cf, cs: append(c.cs, "BYLEX")}
}

func (c SZrangeMax) Rev() SZrangeRevRev {
	return SZrangeRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangeMax) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangeMax) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrangeMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeMin) Max(Max string) SZrangeMax {
	return SZrangeMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c SZrangeMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeRevRev struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeRevRev) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangeRevRev) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrangeRevRev) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeRevRev) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeSortbyBylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeSortbyBylex) Rev() SZrangeRevRev {
	return SZrangeRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangeSortbyBylex) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangeSortbyBylex) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrangeSortbyBylex) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeSortbyBylex) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeSortbyByscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeSortbyByscore) Rev() SZrangeRevRev {
	return SZrangeRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangeSortbyByscore) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangeSortbyByscore) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrangeSortbyByscore) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeSortbyByscore) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangeWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebylex) Key(Key string) SZrangebylexKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrangebylex() (c SZrangebylex) {
	c.cs = append(b.get(), "ZRANGEBYLEX")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrangebylexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebylexKey) Min(Min string) SZrangebylexMin {
	return SZrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c SZrangebylexKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylexLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebylexLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebylexLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylexMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebylexMax) Limit(Offset int64, Count int64) SZrangebylexLimit {
	return SZrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangebylexMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebylexMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylexMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebylexMin) Max(Max string) SZrangebylexMax {
	return SZrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c SZrangebylexMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebyscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebyscore) Key(Key string) SZrangebyscoreKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrangebyscore() (c SZrangebyscore) {
	c.cs = append(b.get(), "ZRANGEBYSCORE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrangebyscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebyscoreKey) Min(Min float64) SZrangebyscoreMin {
	return SZrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c SZrangebyscoreKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebyscoreLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebyscoreLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebyscoreLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebyscoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebyscoreMax) Withscores() SZrangebyscoreWithscoresWithscores {
	return SZrangebyscoreWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrangebyscoreMax) Limit(Offset int64, Count int64) SZrangebyscoreLimit {
	return SZrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangebyscoreMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebyscoreMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebyscoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebyscoreMin) Max(Max float64) SZrangebyscoreMax {
	return SZrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c SZrangebyscoreMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebyscoreWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) SZrangebyscoreLimit {
	return SZrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangebyscoreWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebyscoreWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangestore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestore) Dst(Dst string) SZrangestoreDst {
	s := slot(Dst)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrangestoreDst{cf: c.cf, cs: append(c.cs, Dst)}
}

func (b *SBuilder) Zrangestore() (c SZrangestore) {
	c.cs = append(b.get(), "ZRANGESTORE")
	c.ks = initSlot
	return
}

type SZrangestoreDst struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreDst) Src(Src string) SZrangestoreSrc {
	s := slot(Src)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrangestoreSrc{cf: c.cf, cs: append(c.cs, Src)}
}

type SZrangestoreLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreLimit) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreMax) Byscore() SZrangestoreSortbyByscore {
	return SZrangestoreSortbyByscore{cf: c.cf, cs: append(c.cs, "BYSCORE")}
}

func (c SZrangestoreMax) Bylex() SZrangestoreSortbyBylex {
	return SZrangestoreSortbyBylex{cf: c.cf, cs: append(c.cs, "BYLEX")}
}

func (c SZrangestoreMax) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangestoreMax) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangestoreMax) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreMin) Max(Max string) SZrangestoreMax {
	return SZrangestoreMax{cf: c.cf, cs: append(c.cs, Max)}
}

type SZrangestoreRevRev struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreRevRev) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangestoreRevRev) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSortbyBylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreSortbyBylex) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangestoreSortbyBylex) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangestoreSortbyBylex) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSortbyByscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreSortbyByscore) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangestoreSortbyByscore) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangestoreSortbyByscore) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSrc struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrangestoreSrc) Min(Min string) SZrangestoreMin {
	return SZrangestoreMin{cf: c.cf, cs: append(c.cs, Min)}
}

type SZrank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrank) Key(Key string) SZrankKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrank() (c SZrank) {
	c.cs = append(b.get(), "ZRANK")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrankKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrankKey) Member(Member string) SZrankMember {
	return SZrankMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c SZrankKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrankMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrankMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrankMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZrem struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrem) Key(Key string) SZremKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrem() (c SZrem) {
	c.cs = append(b.get(), "ZREM")
	c.ks = initSlot
	return
}

type SZremKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremKey) Member(Member ...string) SZremMember {
	return SZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SZremMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremMember) Member(Member ...string) SZremMember {
	return SZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SZremMember) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebylex) Key(Key string) SZremrangebylexKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZremrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zremrangebylex() (c SZremrangebylex) {
	c.cs = append(b.get(), "ZREMRANGEBYLEX")
	c.ks = initSlot
	return
}

type SZremrangebylexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebylexKey) Min(Min string) SZremrangebylexMin {
	return SZremrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

type SZremrangebylexMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebylexMax) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebylexMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebylexMin) Max(Max string) SZremrangebylexMax {
	return SZremrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

type SZremrangebyrank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyrank) Key(Key string) SZremrangebyrankKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZremrangebyrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zremrangebyrank() (c SZremrangebyrank) {
	c.cs = append(b.get(), "ZREMRANGEBYRANK")
	c.ks = initSlot
	return
}

type SZremrangebyrankKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyrankKey) Start(Start int64) SZremrangebyrankStart {
	return SZremrangebyrankStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type SZremrangebyrankStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyrankStart) Stop(Stop int64) SZremrangebyrankStop {
	return SZremrangebyrankStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type SZremrangebyrankStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyrankStop) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebyscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyscore) Key(Key string) SZremrangebyscoreKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZremrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zremrangebyscore() (c SZremrangebyscore) {
	c.cs = append(b.get(), "ZREMRANGEBYSCORE")
	c.ks = initSlot
	return
}

type SZremrangebyscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyscoreKey) Min(Min float64) SZremrangebyscoreMin {
	return SZremrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type SZremrangebyscoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyscoreMax) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebyscoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZremrangebyscoreMin) Max(Max float64) SZremrangebyscoreMax {
	return SZremrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type SZrevrange struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrange) Key(Key string) SZrevrangeKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrange() (c SZrevrange) {
	c.cs = append(b.get(), "ZREVRANGE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrevrangeKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangeKey) Start(Start int64) SZrevrangeStart {
	return SZrevrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c SZrevrangeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangeStart struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangeStart) Stop(Stop int64) SZrevrangeStop {
	return SZrevrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

func (c SZrevrangeStart) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangeStop struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangeStop) Withscores() SZrevrangeWithscoresWithscores {
	return SZrevrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrevrangeStop) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangeStop) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangeWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangeWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangeWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebylex struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebylex) Key(Key string) SZrevrangebylexKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrevrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrangebylex() (c SZrevrangebylex) {
	c.cs = append(b.get(), "ZREVRANGEBYLEX")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrevrangebylexKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebylexKey) Max(Max string) SZrevrangebylexMax {
	return SZrevrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c SZrevrangebylexKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebylexLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebylexLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebylexLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebylexMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebylexMax) Min(Min string) SZrevrangebylexMin {
	return SZrevrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c SZrevrangebylexMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebylexMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebylexMin) Limit(Offset int64, Count int64) SZrevrangebylexLimit {
	return SZrevrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrevrangebylexMin) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebylexMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebyscore) Key(Key string) SZrevrangebyscoreKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrevrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrangebyscore() (c SZrevrangebyscore) {
	c.cs = append(b.get(), "ZREVRANGEBYSCORE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrevrangebyscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebyscoreKey) Max(Max float64) SZrevrangebyscoreMax {
	return SZrevrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c SZrevrangebyscoreKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscoreLimit struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebyscoreLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebyscoreLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscoreMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebyscoreMax) Min(Min float64) SZrevrangebyscoreMin {
	return SZrevrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c SZrevrangebyscoreMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscoreMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebyscoreMin) Withscores() SZrevrangebyscoreWithscoresWithscores {
	return SZrevrangebyscoreWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrevrangebyscoreMin) Limit(Offset int64, Count int64) SZrevrangebyscoreLimit {
	return SZrevrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrevrangebyscoreMin) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebyscoreMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscoreWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) SZrevrangebyscoreLimit {
	return SZrevrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrevrangebyscoreWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebyscoreWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrank struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrank) Key(Key string) SZrevrankKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZrevrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrank() (c SZrevrank) {
	c.cs = append(b.get(), "ZREVRANK")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZrevrankKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrankKey) Member(Member string) SZrevrankMember {
	return SZrevrankMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c SZrevrankKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrankMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZrevrankMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrankMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZscan struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscan) Key(Key string) SZscanKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zscan() (c SZscan) {
	c.cs = append(b.get(), "ZSCAN")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZscanCount struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscanCount) Build() SCompleted {
	return SCompleted(c)
}

type SZscanCursor struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscanCursor) Match(Pattern string) SZscanMatch {
	return SZscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SZscanCursor) Count(Count int64) SZscanCount {
	return SZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SZscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SZscanKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscanKey) Cursor(Cursor int64) SZscanCursor {
	return SZscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SZscanMatch struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscanMatch) Count(Count int64) SZscanCount {
	return SZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SZscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SZscore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscore) Key(Key string) SZscoreKey {
	s := slot(Key)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zscore() (c SZscore) {
	c.cs = append(b.get(), "ZSCORE")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZscoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscoreKey) Member(Member string) SZscoreMember {
	return SZscoreMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c SZscoreKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZscoreMember struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZscoreMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZscoreMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZunion struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunion) Numkeys(Numkeys int64) SZunionNumkeys {
	return SZunionNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Zunion() (c SZunion) {
	c.cs = append(b.get(), "ZUNION")
	c.cf = readonly
	c.ks = initSlot
	return
}

type SZunionAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionAggregateMax) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZunionAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZunionAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionAggregateMin) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZunionAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZunionAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionAggregateSum) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZunionAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZunionKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionKey) Weights(Weight ...int64) SZunionWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionWeights{cf: c.cf, cs: c.cs}
}

func (c SZunionKey) Sum() SZunionAggregateSum {
	return SZunionAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZunionKey) Min() SZunionAggregateMin {
	return SZunionAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZunionKey) Max() SZunionAggregateMax {
	return SZunionAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZunionKey) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZunionKey) Key(Key ...string) SZunionKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZunionKey) Build() SCompleted {
	return SCompleted(c)
}

type SZunionNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionNumkeys) Key(Key ...string) SZunionKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZunionWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionWeights) Sum() SZunionAggregateSum {
	return SZunionAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZunionWeights) Min() SZunionAggregateMin {
	return SZunionAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZunionWeights) Max() SZunionAggregateMax {
	return SZunionAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZunionWeights) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZunionWeights) Weights(Weights ...int64) SZunionWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionWeights{cf: c.cf, cs: c.cs}
}

func (c SZunionWeights) Build() SCompleted {
	return SCompleted(c)
}

type SZunionWithscoresWithscores struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstore struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstore) Destination(Destination string) SZunionstoreDestination {
	s := slot(Destination)
	if c.ks == initSlot {
		c.ks = s
	} else if c.ks != s {
		panic(multiKeySlotErr)
	}
	return SZunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Zunionstore() (c SZunionstore) {
	c.cs = append(b.get(), "ZUNIONSTORE")
	c.ks = initSlot
	return
}

type SZunionstoreAggregateMax struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstoreAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreAggregateMin struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstoreAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreAggregateSum struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstoreAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreDestination struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstoreDestination) Numkeys(Numkeys int64) SZunionstoreNumkeys {
	return SZunionstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SZunionstoreKey struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstoreKey) Weights(Weight ...int64) SZunionstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionstoreWeights{cf: c.cf, cs: c.cs}
}

func (c SZunionstoreKey) Sum() SZunionstoreAggregateSum {
	return SZunionstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZunionstoreKey) Min() SZunionstoreAggregateMin {
	return SZunionstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZunionstoreKey) Max() SZunionstoreAggregateMax {
	return SZunionstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZunionstoreKey) Key(Key ...string) SZunionstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZunionstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreNumkeys struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstoreNumkeys) Key(Key ...string) SZunionstoreKey {
	for _, k := range Key {
		s := slot(k)
		if c.ks == initSlot {
			c.ks = s
		} else if c.ks != s {
			panic(multiKeySlotErr)
		}
	}
	return SZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZunionstoreWeights struct {
	cs []string
	cf uint16
	ks uint16
}

func (c SZunionstoreWeights) Sum() SZunionstoreAggregateSum {
	return SZunionstoreAggregateSum{cf: c.cf, cs: append(c.cs, "SUM")}
}

func (c SZunionstoreWeights) Min() SZunionstoreAggregateMin {
	return SZunionstoreAggregateMin{cf: c.cf, cs: append(c.cs, "MIN")}
}

func (c SZunionstoreWeights) Max() SZunionstoreAggregateMax {
	return SZunionstoreAggregateMax{cf: c.cf, cs: append(c.cs, "MAX")}
}

func (c SZunionstoreWeights) Weights(Weights ...int64) SZunionstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionstoreWeights{cf: c.cf, cs: c.cs}
}

func (c SZunionstoreWeights) Build() SCompleted {
	return SCompleted(c)
}

