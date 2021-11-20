// Code generated DO NOT EDIT

package cmds

import "strconv"

type AclCat Completed

func (c AclCat) Categoryname(Categoryname string) AclCatCategoryname {
	return AclCatCategoryname{cs: append(c.cs, Categoryname), cf: c.cf, ks: c.ks}
}

func (c AclCat) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclCat() (c AclCat) {
	c.cs = append(b.get(), "ACL", "CAT")
	return
}

type AclCatCategoryname Completed

func (c AclCatCategoryname) Build() Completed {
	return Completed(c)
}

type AclDeluser Completed

func (c AclDeluser) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cs: append(c.cs, Username...), cf: c.cf, ks: c.ks}
}

func (b *Builder) AclDeluser() (c AclDeluser) {
	c.cs = append(b.get(), "ACL", "DELUSER")
	return
}

type AclDeluserUsername Completed

func (c AclDeluserUsername) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cs: append(c.cs, Username...), cf: c.cf, ks: c.ks}
}

func (c AclDeluserUsername) Build() Completed {
	return Completed(c)
}

type AclGenpass Completed

func (c AclGenpass) Bits(Bits int64) AclGenpassBits {
	return AclGenpassBits{cs: append(c.cs, strconv.FormatInt(Bits, 10)), cf: c.cf, ks: c.ks}
}

func (c AclGenpass) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclGenpass() (c AclGenpass) {
	c.cs = append(b.get(), "ACL", "GENPASS")
	return
}

type AclGenpassBits Completed

func (c AclGenpassBits) Build() Completed {
	return Completed(c)
}

type AclGetuser Completed

func (c AclGetuser) Username(Username string) AclGetuserUsername {
	return AclGetuserUsername{cs: append(c.cs, Username), cf: c.cf, ks: c.ks}
}

func (b *Builder) AclGetuser() (c AclGetuser) {
	c.cs = append(b.get(), "ACL", "GETUSER")
	return
}

type AclGetuserUsername Completed

func (c AclGetuserUsername) Build() Completed {
	return Completed(c)
}

type AclHelp Completed

func (c AclHelp) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclHelp() (c AclHelp) {
	c.cs = append(b.get(), "ACL", "HELP")
	return
}

type AclList Completed

func (c AclList) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclList() (c AclList) {
	c.cs = append(b.get(), "ACL", "LIST")
	return
}

type AclLoad Completed

func (c AclLoad) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclLoad() (c AclLoad) {
	c.cs = append(b.get(), "ACL", "LOAD")
	return
}

type AclLog Completed

func (c AclLog) CountOrReset(CountOrReset string) AclLogCountOrReset {
	return AclLogCountOrReset{cs: append(c.cs, CountOrReset), cf: c.cf, ks: c.ks}
}

func (c AclLog) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclLog() (c AclLog) {
	c.cs = append(b.get(), "ACL", "LOG")
	return
}

type AclLogCountOrReset Completed

func (c AclLogCountOrReset) Build() Completed {
	return Completed(c)
}

type AclSave Completed

func (c AclSave) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclSave() (c AclSave) {
	c.cs = append(b.get(), "ACL", "SAVE")
	return
}

type AclSetuser Completed

func (c AclSetuser) Username(Username string) AclSetuserUsername {
	return AclSetuserUsername{cs: append(c.cs, Username), cf: c.cf, ks: c.ks}
}

func (b *Builder) AclSetuser() (c AclSetuser) {
	c.cs = append(b.get(), "ACL", "SETUSER")
	return
}

type AclSetuserRule Completed

func (c AclSetuserRule) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cs: append(c.cs, Rule...), cf: c.cf, ks: c.ks}
}

func (c AclSetuserRule) Build() Completed {
	return Completed(c)
}

type AclSetuserUsername Completed

func (c AclSetuserUsername) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cs: append(c.cs, Rule...), cf: c.cf, ks: c.ks}
}

func (c AclSetuserUsername) Build() Completed {
	return Completed(c)
}

type AclUsers Completed

func (c AclUsers) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclUsers() (c AclUsers) {
	c.cs = append(b.get(), "ACL", "USERS")
	return
}

type AclWhoami Completed

func (c AclWhoami) Build() Completed {
	return Completed(c)
}

func (b *Builder) AclWhoami() (c AclWhoami) {
	c.cs = append(b.get(), "ACL", "WHOAMI")
	return
}

type Append Completed

func (c Append) Key(Key string) AppendKey {
	return AppendKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Append() (c Append) {
	c.cs = append(b.get(), "APPEND")
	return
}

type AppendKey Completed

func (c AppendKey) Value(Value string) AppendValue {
	return AppendValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type AppendValue Completed

func (c AppendValue) Build() Completed {
	return Completed(c)
}

type Asking Completed

func (c Asking) Build() Completed {
	return Completed(c)
}

func (b *Builder) Asking() (c Asking) {
	c.cs = append(b.get(), "ASKING")
	return
}

type Auth Completed

func (c Auth) Username(Username string) AuthUsername {
	return AuthUsername{cs: append(c.cs, Username), cf: c.cf, ks: c.ks}
}

func (c Auth) Password(Password string) AuthPassword {
	return AuthPassword{cs: append(c.cs, Password), cf: c.cf, ks: c.ks}
}

func (b *Builder) Auth() (c Auth) {
	c.cs = append(b.get(), "AUTH")
	return
}

type AuthPassword Completed

func (c AuthPassword) Build() Completed {
	return Completed(c)
}

type AuthUsername Completed

func (c AuthUsername) Password(Password string) AuthPassword {
	return AuthPassword{cs: append(c.cs, Password), cf: c.cf, ks: c.ks}
}

type Bgrewriteaof Completed

func (c Bgrewriteaof) Build() Completed {
	return Completed(c)
}

func (b *Builder) Bgrewriteaof() (c Bgrewriteaof) {
	c.cs = append(b.get(), "BGREWRITEAOF")
	return
}

type Bgsave Completed

func (c Bgsave) Schedule() BgsaveScheduleSchedule {
	return BgsaveScheduleSchedule{cs: append(c.cs, "SCHEDULE"), cf: c.cf, ks: c.ks}
}

func (c Bgsave) Build() Completed {
	return Completed(c)
}

func (b *Builder) Bgsave() (c Bgsave) {
	c.cs = append(b.get(), "BGSAVE")
	return
}

type BgsaveScheduleSchedule Completed

func (c BgsaveScheduleSchedule) Build() Completed {
	return Completed(c)
}

type Bitcount Completed

func (c Bitcount) Key(Key string) BitcountKey {
	return BitcountKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Bitcount() (c Bitcount) {
	c.cs = append(b.get(), "BITCOUNT")
	c.cf = readonly
	return
}

type BitcountKey Completed

func (c BitcountKey) StartEnd(Start int64, End int64) BitcountStartEnd {
	return BitcountStartEnd{cs: append(c.cs, strconv.FormatInt(Start, 10), strconv.FormatInt(End, 10)), cf: c.cf, ks: c.ks}
}

func (c BitcountKey) Build() Completed {
	return Completed(c)
}

func (c BitcountKey) Cache() Cacheable {
	return Cacheable(c)
}

type BitcountStartEnd Completed

func (c BitcountStartEnd) Build() Completed {
	return Completed(c)
}

func (c BitcountStartEnd) Cache() Cacheable {
	return Cacheable(c)
}

type Bitfield Completed

func (c Bitfield) Key(Key string) BitfieldKey {
	return BitfieldKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Bitfield() (c Bitfield) {
	c.cs = append(b.get(), "BITFIELD")
	return
}

type BitfieldFail Completed

func (c BitfieldFail) Build() Completed {
	return Completed(c)
}

type BitfieldGet Completed

func (c BitfieldGet) Set(Type string, Offset int64, Value int64) BitfieldSet {
	return BitfieldSet{cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10)), cf: c.cf, ks: c.ks}
}

func (c BitfieldGet) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

func (c BitfieldGet) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c BitfieldGet) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c BitfieldGet) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c BitfieldGet) Build() Completed {
	return Completed(c)
}

type BitfieldIncrby Completed

func (c BitfieldIncrby) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c BitfieldIncrby) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c BitfieldIncrby) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c BitfieldIncrby) Build() Completed {
	return Completed(c)
}

type BitfieldKey Completed

func (c BitfieldKey) Get(Type string, Offset int64) BitfieldGet {
	return BitfieldGet{cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

func (c BitfieldKey) Set(Type string, Offset int64, Value int64) BitfieldSet {
	return BitfieldSet{cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10)), cf: c.cf, ks: c.ks}
}

func (c BitfieldKey) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

func (c BitfieldKey) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c BitfieldKey) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c BitfieldKey) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c BitfieldKey) Build() Completed {
	return Completed(c)
}

type BitfieldRo Completed

func (c BitfieldRo) Key(Key string) BitfieldRoKey {
	return BitfieldRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) BitfieldRo() (c BitfieldRo) {
	c.cs = append(b.get(), "BITFIELD_RO")
	c.cf = readonly
	return
}

type BitfieldRoGet Completed

func (c BitfieldRoGet) Build() Completed {
	return Completed(c)
}

func (c BitfieldRoGet) Cache() Cacheable {
	return Cacheable(c)
}

type BitfieldRoKey Completed

func (c BitfieldRoKey) Get(Type string, Offset int64) BitfieldRoGet {
	return BitfieldRoGet{cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type BitfieldSat Completed

func (c BitfieldSat) Build() Completed {
	return Completed(c)
}

type BitfieldSet Completed

func (c BitfieldSet) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

func (c BitfieldSet) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c BitfieldSet) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c BitfieldSet) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c BitfieldSet) Build() Completed {
	return Completed(c)
}

type BitfieldWrap Completed

func (c BitfieldWrap) Build() Completed {
	return Completed(c)
}

type Bitop Completed

func (c Bitop) Operation(Operation string) BitopOperation {
	return BitopOperation{cs: append(c.cs, Operation), cf: c.cf, ks: c.ks}
}

func (b *Builder) Bitop() (c Bitop) {
	c.cs = append(b.get(), "BITOP")
	return
}

type BitopDestkey Completed

func (c BitopDestkey) Key(Key ...string) BitopKey {
	return BitopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type BitopKey Completed

func (c BitopKey) Key(Key ...string) BitopKey {
	return BitopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c BitopKey) Build() Completed {
	return Completed(c)
}

type BitopOperation Completed

func (c BitopOperation) Destkey(Destkey string) BitopDestkey {
	return BitopDestkey{cs: append(c.cs, Destkey), cf: c.cf, ks: c.ks}
}

type Bitpos Completed

func (c Bitpos) Key(Key string) BitposKey {
	return BitposKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Bitpos() (c Bitpos) {
	c.cs = append(b.get(), "BITPOS")
	c.cf = readonly
	return
}

type BitposBit Completed

func (c BitposBit) Start(Start int64) BitposIndexStart {
	return BitposIndexStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

func (c BitposBit) Build() Completed {
	return Completed(c)
}

func (c BitposBit) Cache() Cacheable {
	return Cacheable(c)
}

type BitposIndexEnd Completed

func (c BitposIndexEnd) Build() Completed {
	return Completed(c)
}

func (c BitposIndexEnd) Cache() Cacheable {
	return Cacheable(c)
}

type BitposIndexStart Completed

func (c BitposIndexStart) End(End int64) BitposIndexEnd {
	return BitposIndexEnd{cs: append(c.cs, strconv.FormatInt(End, 10)), cf: c.cf, ks: c.ks}
}

func (c BitposIndexStart) Build() Completed {
	return Completed(c)
}

func (c BitposIndexStart) Cache() Cacheable {
	return Cacheable(c)
}

type BitposKey Completed

func (c BitposKey) Bit(Bit int64) BitposBit {
	return BitposBit{cs: append(c.cs, strconv.FormatInt(Bit, 10)), cf: c.cf, ks: c.ks}
}

type Blmove Completed

func (c Blmove) Source(Source string) BlmoveSource {
	return BlmoveSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *Builder) Blmove() (c Blmove) {
	c.cs = append(b.get(), "BLMOVE")
	c.cf = blockTag
	return
}

type BlmoveDestination Completed

func (c BlmoveDestination) Left() BlmoveWherefromLeft {
	return BlmoveWherefromLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c BlmoveDestination) Right() BlmoveWherefromRight {
	return BlmoveWherefromRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type BlmoveSource Completed

func (c BlmoveSource) Destination(Destination string) BlmoveDestination {
	return BlmoveDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type BlmoveTimeout Completed

func (c BlmoveTimeout) Build() Completed {
	return Completed(c)
}

type BlmoveWherefromLeft Completed

func (c BlmoveWherefromLeft) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c BlmoveWherefromLeft) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type BlmoveWherefromRight Completed

func (c BlmoveWherefromRight) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c BlmoveWherefromRight) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type BlmoveWheretoLeft Completed

func (c BlmoveWheretoLeft) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type BlmoveWheretoRight Completed

func (c BlmoveWheretoRight) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type Blmpop Completed

func (c Blmpop) Timeout(Timeout float64) BlmpopTimeout {
	return BlmpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Blmpop() (c Blmpop) {
	c.cs = append(b.get(), "BLMPOP")
	c.cf = blockTag
	return
}

type BlmpopCount Completed

func (c BlmpopCount) Build() Completed {
	return Completed(c)
}

type BlmpopKey Completed

func (c BlmpopKey) Left() BlmpopWhereLeft {
	return BlmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c BlmpopKey) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

func (c BlmpopKey) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type BlmpopNumkeys Completed

func (c BlmpopNumkeys) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c BlmpopNumkeys) Left() BlmpopWhereLeft {
	return BlmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c BlmpopNumkeys) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type BlmpopTimeout Completed

func (c BlmpopTimeout) Numkeys(Numkeys int64) BlmpopNumkeys {
	return BlmpopNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type BlmpopWhereLeft Completed

func (c BlmpopWhereLeft) Count(Count int64) BlmpopCount {
	return BlmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c BlmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type BlmpopWhereRight Completed

func (c BlmpopWhereRight) Count(Count int64) BlmpopCount {
	return BlmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c BlmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Blpop Completed

func (c Blpop) Key(Key ...string) BlpopKey {
	return BlpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Blpop() (c Blpop) {
	c.cs = append(b.get(), "BLPOP")
	c.cf = blockTag
	return
}

type BlpopKey Completed

func (c BlpopKey) Timeout(Timeout float64) BlpopTimeout {
	return BlpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c BlpopKey) Key(Key ...string) BlpopKey {
	return BlpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type BlpopTimeout Completed

func (c BlpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpop Completed

func (c Brpop) Key(Key ...string) BrpopKey {
	return BrpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Brpop() (c Brpop) {
	c.cs = append(b.get(), "BRPOP")
	c.cf = blockTag
	return
}

type BrpopKey Completed

func (c BrpopKey) Timeout(Timeout float64) BrpopTimeout {
	return BrpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c BrpopKey) Key(Key ...string) BrpopKey {
	return BrpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type BrpopTimeout Completed

func (c BrpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpoplpush Completed

func (c Brpoplpush) Source(Source string) BrpoplpushSource {
	return BrpoplpushSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *Builder) Brpoplpush() (c Brpoplpush) {
	c.cs = append(b.get(), "BRPOPLPUSH")
	c.cf = blockTag
	return
}

type BrpoplpushDestination Completed

func (c BrpoplpushDestination) Timeout(Timeout float64) BrpoplpushTimeout {
	return BrpoplpushTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type BrpoplpushSource Completed

func (c BrpoplpushSource) Destination(Destination string) BrpoplpushDestination {
	return BrpoplpushDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type BrpoplpushTimeout Completed

func (c BrpoplpushTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmax Completed

func (c Bzpopmax) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Bzpopmax() (c Bzpopmax) {
	c.cs = append(b.get(), "BZPOPMAX")
	c.cf = blockTag
	return
}

type BzpopmaxKey Completed

func (c BzpopmaxKey) Timeout(Timeout float64) BzpopmaxTimeout {
	return BzpopmaxTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c BzpopmaxKey) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type BzpopmaxTimeout Completed

func (c BzpopmaxTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmin Completed

func (c Bzpopmin) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Bzpopmin() (c Bzpopmin) {
	c.cs = append(b.get(), "BZPOPMIN")
	c.cf = blockTag
	return
}

type BzpopminKey Completed

func (c BzpopminKey) Timeout(Timeout float64) BzpopminTimeout {
	return BzpopminTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c BzpopminKey) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type BzpopminTimeout Completed

func (c BzpopminTimeout) Build() Completed {
	return Completed(c)
}

type ClientCaching Completed

func (c ClientCaching) Yes() ClientCachingModeYes {
	return ClientCachingModeYes{cs: append(c.cs, "YES"), cf: c.cf, ks: c.ks}
}

func (c ClientCaching) No() ClientCachingModeNo {
	return ClientCachingModeNo{cs: append(c.cs, "NO"), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClientCaching() (c ClientCaching) {
	c.cs = append(b.get(), "CLIENT", "CACHING")
	return
}

type ClientCachingModeNo Completed

func (c ClientCachingModeNo) Build() Completed {
	return Completed(c)
}

type ClientCachingModeYes Completed

func (c ClientCachingModeYes) Build() Completed {
	return Completed(c)
}

type ClientGetname Completed

func (c ClientGetname) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientGetname() (c ClientGetname) {
	c.cs = append(b.get(), "CLIENT", "GETNAME")
	return
}

type ClientGetredir Completed

func (c ClientGetredir) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientGetredir() (c ClientGetredir) {
	c.cs = append(b.get(), "CLIENT", "GETREDIR")
	return
}

type ClientId Completed

func (c ClientId) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientId() (c ClientId) {
	c.cs = append(b.get(), "CLIENT", "ID")
	return
}

type ClientInfo Completed

func (c ClientInfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientInfo() (c ClientInfo) {
	c.cs = append(b.get(), "CLIENT", "INFO")
	return
}

type ClientKill Completed

func (c ClientKill) IpPort(IpPort string) ClientKillIpPort {
	return ClientKillIpPort{cs: append(c.cs, IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Id(ClientId int64) ClientKillId {
	return ClientKillId{cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Normal() ClientKillNormal {
	return ClientKillNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Master() ClientKillMaster {
	return ClientKillMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Slave() ClientKillSlave {
	return ClientKillSlave{cs: append(c.cs, "slave"), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c ClientKill) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKill) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientKill() (c ClientKill) {
	c.cs = append(b.get(), "CLIENT", "KILL")
	return
}

type ClientKillAddr Completed

func (c ClientKillAddr) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillAddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillAddr) Build() Completed {
	return Completed(c)
}

type ClientKillId Completed

func (c ClientKillId) Normal() ClientKillNormal {
	return ClientKillNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) Master() ClientKillMaster {
	return ClientKillMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) Slave() ClientKillSlave {
	return ClientKillSlave{cs: append(c.cs, "slave"), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillId) Build() Completed {
	return Completed(c)
}

type ClientKillIpPort Completed

func (c ClientKillIpPort) Id(ClientId int64) ClientKillId {
	return ClientKillId{cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Normal() ClientKillNormal {
	return ClientKillNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Master() ClientKillMaster {
	return ClientKillMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Slave() ClientKillSlave {
	return ClientKillSlave{cs: append(c.cs, "slave"), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillIpPort) Build() Completed {
	return Completed(c)
}

type ClientKillLaddr Completed

func (c ClientKillLaddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillLaddr) Build() Completed {
	return Completed(c)
}

type ClientKillMaster Completed

func (c ClientKillMaster) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c ClientKillMaster) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillMaster) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillMaster) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillMaster) Build() Completed {
	return Completed(c)
}

type ClientKillNormal Completed

func (c ClientKillNormal) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c ClientKillNormal) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillNormal) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillNormal) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillNormal) Build() Completed {
	return Completed(c)
}

type ClientKillPubsub Completed

func (c ClientKillPubsub) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c ClientKillPubsub) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillPubsub) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillPubsub) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillPubsub) Build() Completed {
	return Completed(c)
}

type ClientKillSkipme Completed

func (c ClientKillSkipme) Build() Completed {
	return Completed(c)
}

type ClientKillSlave Completed

func (c ClientKillSlave) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c ClientKillSlave) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillSlave) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillSlave) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillSlave) Build() Completed {
	return Completed(c)
}

type ClientKillUser Completed

func (c ClientKillUser) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillUser) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c ClientKillUser) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c ClientKillUser) Build() Completed {
	return Completed(c)
}

type ClientList Completed

func (c ClientList) Normal() ClientListNormal {
	return ClientListNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c ClientList) Master() ClientListMaster {
	return ClientListMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c ClientList) Replica() ClientListReplica {
	return ClientListReplica{cs: append(c.cs, "replica"), cf: c.cf, ks: c.ks}
}

func (c ClientList) Pubsub() ClientListPubsub {
	return ClientListPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c ClientList) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c ClientList) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientList() (c ClientList) {
	c.cs = append(b.get(), "CLIENT", "LIST")
	return
}

type ClientListIdClientId Completed

func (c ClientListIdClientId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ClientListIdClientId) Build() Completed {
	return Completed(c)
}

type ClientListIdId Completed

func (c ClientListIdId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ClientListMaster Completed

func (c ClientListMaster) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c ClientListMaster) Build() Completed {
	return Completed(c)
}

type ClientListNormal Completed

func (c ClientListNormal) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c ClientListNormal) Build() Completed {
	return Completed(c)
}

type ClientListPubsub Completed

func (c ClientListPubsub) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c ClientListPubsub) Build() Completed {
	return Completed(c)
}

type ClientListReplica Completed

func (c ClientListReplica) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c ClientListReplica) Build() Completed {
	return Completed(c)
}

type ClientNoEvict Completed

func (c ClientNoEvict) On() ClientNoEvictEnabledOn {
	return ClientNoEvictEnabledOn{cs: append(c.cs, "ON"), cf: c.cf, ks: c.ks}
}

func (c ClientNoEvict) Off() ClientNoEvictEnabledOff {
	return ClientNoEvictEnabledOff{cs: append(c.cs, "OFF"), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClientNoEvict() (c ClientNoEvict) {
	c.cs = append(b.get(), "CLIENT", "NO-EVICT")
	return
}

type ClientNoEvictEnabledOff Completed

func (c ClientNoEvictEnabledOff) Build() Completed {
	return Completed(c)
}

type ClientNoEvictEnabledOn Completed

func (c ClientNoEvictEnabledOn) Build() Completed {
	return Completed(c)
}

type ClientPause Completed

func (c ClientPause) Timeout(Timeout int64) ClientPauseTimeout {
	return ClientPauseTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClientPause() (c ClientPause) {
	c.cs = append(b.get(), "CLIENT", "PAUSE")
	c.cf = blockTag
	return
}

type ClientPauseModeAll Completed

func (c ClientPauseModeAll) Build() Completed {
	return Completed(c)
}

type ClientPauseModeWrite Completed

func (c ClientPauseModeWrite) Build() Completed {
	return Completed(c)
}

type ClientPauseTimeout Completed

func (c ClientPauseTimeout) Write() ClientPauseModeWrite {
	return ClientPauseModeWrite{cs: append(c.cs, "WRITE"), cf: c.cf, ks: c.ks}
}

func (c ClientPauseTimeout) All() ClientPauseModeAll {
	return ClientPauseModeAll{cs: append(c.cs, "ALL"), cf: c.cf, ks: c.ks}
}

func (c ClientPauseTimeout) Build() Completed {
	return Completed(c)
}

type ClientReply Completed

func (c ClientReply) On() ClientReplyReplyModeOn {
	return ClientReplyReplyModeOn{cs: append(c.cs, "ON"), cf: c.cf, ks: c.ks}
}

func (c ClientReply) Off() ClientReplyReplyModeOff {
	return ClientReplyReplyModeOff{cs: append(c.cs, "OFF"), cf: c.cf, ks: c.ks}
}

func (c ClientReply) Skip() ClientReplyReplyModeSkip {
	return ClientReplyReplyModeSkip{cs: append(c.cs, "SKIP"), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClientReply() (c ClientReply) {
	c.cs = append(b.get(), "CLIENT", "REPLY")
	return
}

type ClientReplyReplyModeOff Completed

func (c ClientReplyReplyModeOff) Build() Completed {
	return Completed(c)
}

type ClientReplyReplyModeOn Completed

func (c ClientReplyReplyModeOn) Build() Completed {
	return Completed(c)
}

type ClientReplyReplyModeSkip Completed

func (c ClientReplyReplyModeSkip) Build() Completed {
	return Completed(c)
}

type ClientSetname Completed

func (c ClientSetname) ConnectionName(ConnectionName string) ClientSetnameConnectionName {
	return ClientSetnameConnectionName{cs: append(c.cs, ConnectionName), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClientSetname() (c ClientSetname) {
	c.cs = append(b.get(), "CLIENT", "SETNAME")
	return
}

type ClientSetnameConnectionName Completed

func (c ClientSetnameConnectionName) Build() Completed {
	return Completed(c)
}

type ClientTracking Completed

func (c ClientTracking) On() ClientTrackingStatusOn {
	return ClientTrackingStatusOn{cs: append(c.cs, "ON"), cf: c.cf, ks: c.ks}
}

func (c ClientTracking) Off() ClientTrackingStatusOff {
	return ClientTrackingStatusOff{cs: append(c.cs, "OFF"), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClientTracking() (c ClientTracking) {
	c.cs = append(b.get(), "CLIENT", "TRACKING")
	return
}

type ClientTrackingBcastBcast Completed

func (c ClientTrackingBcastBcast) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingBcastBcast) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingBcastBcast) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingBcastBcast) Build() Completed {
	return Completed(c)
}

type ClientTrackingNoloopNoloop Completed

func (c ClientTrackingNoloopNoloop) Build() Completed {
	return Completed(c)
}

type ClientTrackingOptinOptin Completed

func (c ClientTrackingOptinOptin) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingOptinOptin) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingOptinOptin) Build() Completed {
	return Completed(c)
}

type ClientTrackingOptoutOptout Completed

func (c ClientTrackingOptoutOptout) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingOptoutOptout) Build() Completed {
	return Completed(c)
}

type ClientTrackingPrefix Completed

func (c ClientTrackingPrefix) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingPrefix) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingPrefix) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingPrefix) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingPrefix) Prefix(Prefix ...string) ClientTrackingPrefix {
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingPrefix) Build() Completed {
	return Completed(c)
}

type ClientTrackingRedirect Completed

func (c ClientTrackingRedirect) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingRedirect) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingRedirect) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingRedirect) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingRedirect) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingRedirect) Build() Completed {
	return Completed(c)
}

type ClientTrackingStatusOff Completed

func (c ClientTrackingStatusOff) Redirect(ClientId int64) ClientTrackingRedirect {
	return ClientTrackingRedirect{cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOff) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOff) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOff) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOff) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOff) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOff) Build() Completed {
	return Completed(c)
}

type ClientTrackingStatusOn Completed

func (c ClientTrackingStatusOn) Redirect(ClientId int64) ClientTrackingRedirect {
	return ClientTrackingRedirect{cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOn) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOn) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOn) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOn) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOn) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c ClientTrackingStatusOn) Build() Completed {
	return Completed(c)
}

type ClientTrackinginfo Completed

func (c ClientTrackinginfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientTrackinginfo() (c ClientTrackinginfo) {
	c.cs = append(b.get(), "CLIENT", "TRACKINGINFO")
	return
}

type ClientUnblock Completed

func (c ClientUnblock) ClientId(ClientId int64) ClientUnblockClientId {
	return ClientUnblockClientId{cs: append(c.cs, strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClientUnblock() (c ClientUnblock) {
	c.cs = append(b.get(), "CLIENT", "UNBLOCK")
	return
}

type ClientUnblockClientId Completed

func (c ClientUnblockClientId) Timeout() ClientUnblockUnblockTypeTimeout {
	return ClientUnblockUnblockTypeTimeout{cs: append(c.cs, "TIMEOUT"), cf: c.cf, ks: c.ks}
}

func (c ClientUnblockClientId) Error() ClientUnblockUnblockTypeError {
	return ClientUnblockUnblockTypeError{cs: append(c.cs, "ERROR"), cf: c.cf, ks: c.ks}
}

func (c ClientUnblockClientId) Build() Completed {
	return Completed(c)
}

type ClientUnblockUnblockTypeError Completed

func (c ClientUnblockUnblockTypeError) Build() Completed {
	return Completed(c)
}

type ClientUnblockUnblockTypeTimeout Completed

func (c ClientUnblockUnblockTypeTimeout) Build() Completed {
	return Completed(c)
}

type ClientUnpause Completed

func (c ClientUnpause) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClientUnpause() (c ClientUnpause) {
	c.cs = append(b.get(), "CLIENT", "UNPAUSE")
	return
}

type ClusterAddslots Completed

func (c ClusterAddslots) Slot(Slot ...int64) ClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterAddslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterAddslots() (c ClusterAddslots) {
	c.cs = append(b.get(), "CLUSTER", "ADDSLOTS")
	return
}

type ClusterAddslotsSlot Completed

func (c ClusterAddslotsSlot) Slot(Slot ...int64) ClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterAddslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ClusterAddslotsSlot) Build() Completed {
	return Completed(c)
}

type ClusterBumpepoch Completed

func (c ClusterBumpepoch) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterBumpepoch() (c ClusterBumpepoch) {
	c.cs = append(b.get(), "CLUSTER", "BUMPEPOCH")
	return
}

type ClusterCountFailureReports Completed

func (c ClusterCountFailureReports) NodeId(NodeId string) ClusterCountFailureReportsNodeId {
	return ClusterCountFailureReportsNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterCountFailureReports() (c ClusterCountFailureReports) {
	c.cs = append(b.get(), "CLUSTER", "COUNT-FAILURE-REPORTS")
	return
}

type ClusterCountFailureReportsNodeId Completed

func (c ClusterCountFailureReportsNodeId) Build() Completed {
	return Completed(c)
}

type ClusterCountkeysinslot Completed

func (c ClusterCountkeysinslot) Slot(Slot int64) ClusterCountkeysinslotSlot {
	return ClusterCountkeysinslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterCountkeysinslot() (c ClusterCountkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "COUNTKEYSINSLOT")
	return
}

type ClusterCountkeysinslotSlot Completed

func (c ClusterCountkeysinslotSlot) Build() Completed {
	return Completed(c)
}

type ClusterDelslots Completed

func (c ClusterDelslots) Slot(Slot ...int64) ClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterDelslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterDelslots() (c ClusterDelslots) {
	c.cs = append(b.get(), "CLUSTER", "DELSLOTS")
	return
}

type ClusterDelslotsSlot Completed

func (c ClusterDelslotsSlot) Slot(Slot ...int64) ClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterDelslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ClusterDelslotsSlot) Build() Completed {
	return Completed(c)
}

type ClusterFailover Completed

func (c ClusterFailover) Force() ClusterFailoverOptionsForce {
	return ClusterFailoverOptionsForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c ClusterFailover) Takeover() ClusterFailoverOptionsTakeover {
	return ClusterFailoverOptionsTakeover{cs: append(c.cs, "TAKEOVER"), cf: c.cf, ks: c.ks}
}

func (c ClusterFailover) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterFailover() (c ClusterFailover) {
	c.cs = append(b.get(), "CLUSTER", "FAILOVER")
	return
}

type ClusterFailoverOptionsForce Completed

func (c ClusterFailoverOptionsForce) Build() Completed {
	return Completed(c)
}

type ClusterFailoverOptionsTakeover Completed

func (c ClusterFailoverOptionsTakeover) Build() Completed {
	return Completed(c)
}

type ClusterFlushslots Completed

func (c ClusterFlushslots) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterFlushslots() (c ClusterFlushslots) {
	c.cs = append(b.get(), "CLUSTER", "FLUSHSLOTS")
	return
}

type ClusterForget Completed

func (c ClusterForget) NodeId(NodeId string) ClusterForgetNodeId {
	return ClusterForgetNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterForget() (c ClusterForget) {
	c.cs = append(b.get(), "CLUSTER", "FORGET")
	return
}

type ClusterForgetNodeId Completed

func (c ClusterForgetNodeId) Build() Completed {
	return Completed(c)
}

type ClusterGetkeysinslot Completed

func (c ClusterGetkeysinslot) Slot(Slot int64) ClusterGetkeysinslotSlot {
	return ClusterGetkeysinslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterGetkeysinslot() (c ClusterGetkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "GETKEYSINSLOT")
	return
}

type ClusterGetkeysinslotCount Completed

func (c ClusterGetkeysinslotCount) Build() Completed {
	return Completed(c)
}

type ClusterGetkeysinslotSlot Completed

func (c ClusterGetkeysinslotSlot) Count(Count int64) ClusterGetkeysinslotCount {
	return ClusterGetkeysinslotCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

type ClusterInfo Completed

func (c ClusterInfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterInfo() (c ClusterInfo) {
	c.cs = append(b.get(), "CLUSTER", "INFO")
	return
}

type ClusterKeyslot Completed

func (c ClusterKeyslot) Key(Key string) ClusterKeyslotKey {
	return ClusterKeyslotKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterKeyslot() (c ClusterKeyslot) {
	c.cs = append(b.get(), "CLUSTER", "KEYSLOT")
	return
}

type ClusterKeyslotKey Completed

func (c ClusterKeyslotKey) Build() Completed {
	return Completed(c)
}

type ClusterMeet Completed

func (c ClusterMeet) Ip(Ip string) ClusterMeetIp {
	return ClusterMeetIp{cs: append(c.cs, Ip), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterMeet() (c ClusterMeet) {
	c.cs = append(b.get(), "CLUSTER", "MEET")
	return
}

type ClusterMeetIp Completed

func (c ClusterMeetIp) Port(Port int64) ClusterMeetPort {
	return ClusterMeetPort{cs: append(c.cs, strconv.FormatInt(Port, 10)), cf: c.cf, ks: c.ks}
}

type ClusterMeetPort Completed

func (c ClusterMeetPort) Build() Completed {
	return Completed(c)
}

type ClusterMyid Completed

func (c ClusterMyid) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterMyid() (c ClusterMyid) {
	c.cs = append(b.get(), "CLUSTER", "MYID")
	return
}

type ClusterNodes Completed

func (c ClusterNodes) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterNodes() (c ClusterNodes) {
	c.cs = append(b.get(), "CLUSTER", "NODES")
	return
}

type ClusterReplicas Completed

func (c ClusterReplicas) NodeId(NodeId string) ClusterReplicasNodeId {
	return ClusterReplicasNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterReplicas() (c ClusterReplicas) {
	c.cs = append(b.get(), "CLUSTER", "REPLICAS")
	return
}

type ClusterReplicasNodeId Completed

func (c ClusterReplicasNodeId) Build() Completed {
	return Completed(c)
}

type ClusterReplicate Completed

func (c ClusterReplicate) NodeId(NodeId string) ClusterReplicateNodeId {
	return ClusterReplicateNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterReplicate() (c ClusterReplicate) {
	c.cs = append(b.get(), "CLUSTER", "REPLICATE")
	return
}

type ClusterReplicateNodeId Completed

func (c ClusterReplicateNodeId) Build() Completed {
	return Completed(c)
}

type ClusterReset Completed

func (c ClusterReset) Hard() ClusterResetResetTypeHard {
	return ClusterResetResetTypeHard{cs: append(c.cs, "HARD"), cf: c.cf, ks: c.ks}
}

func (c ClusterReset) Soft() ClusterResetResetTypeSoft {
	return ClusterResetResetTypeSoft{cs: append(c.cs, "SOFT"), cf: c.cf, ks: c.ks}
}

func (c ClusterReset) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterReset() (c ClusterReset) {
	c.cs = append(b.get(), "CLUSTER", "RESET")
	return
}

type ClusterResetResetTypeHard Completed

func (c ClusterResetResetTypeHard) Build() Completed {
	return Completed(c)
}

type ClusterResetResetTypeSoft Completed

func (c ClusterResetResetTypeSoft) Build() Completed {
	return Completed(c)
}

type ClusterSaveconfig Completed

func (c ClusterSaveconfig) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterSaveconfig() (c ClusterSaveconfig) {
	c.cs = append(b.get(), "CLUSTER", "SAVECONFIG")
	return
}

type ClusterSetConfigEpoch Completed

func (c ClusterSetConfigEpoch) ConfigEpoch(ConfigEpoch int64) ClusterSetConfigEpochConfigEpoch {
	return ClusterSetConfigEpochConfigEpoch{cs: append(c.cs, strconv.FormatInt(ConfigEpoch, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterSetConfigEpoch() (c ClusterSetConfigEpoch) {
	c.cs = append(b.get(), "CLUSTER", "SET-CONFIG-EPOCH")
	return
}

type ClusterSetConfigEpochConfigEpoch Completed

func (c ClusterSetConfigEpochConfigEpoch) Build() Completed {
	return Completed(c)
}

type ClusterSetslot Completed

func (c ClusterSetslot) Slot(Slot int64) ClusterSetslotSlot {
	return ClusterSetslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterSetslot() (c ClusterSetslot) {
	c.cs = append(b.get(), "CLUSTER", "SETSLOT")
	return
}

type ClusterSetslotNodeId Completed

func (c ClusterSetslotNodeId) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSlot Completed

func (c ClusterSetslotSlot) Importing() ClusterSetslotSubcommandImporting {
	return ClusterSetslotSubcommandImporting{cs: append(c.cs, "IMPORTING"), cf: c.cf, ks: c.ks}
}

func (c ClusterSetslotSlot) Migrating() ClusterSetslotSubcommandMigrating {
	return ClusterSetslotSubcommandMigrating{cs: append(c.cs, "MIGRATING"), cf: c.cf, ks: c.ks}
}

func (c ClusterSetslotSlot) Stable() ClusterSetslotSubcommandStable {
	return ClusterSetslotSubcommandStable{cs: append(c.cs, "STABLE"), cf: c.cf, ks: c.ks}
}

func (c ClusterSetslotSlot) Node() ClusterSetslotSubcommandNode {
	return ClusterSetslotSubcommandNode{cs: append(c.cs, "NODE"), cf: c.cf, ks: c.ks}
}

type ClusterSetslotSubcommandImporting Completed

func (c ClusterSetslotSubcommandImporting) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c ClusterSetslotSubcommandImporting) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandMigrating Completed

func (c ClusterSetslotSubcommandMigrating) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c ClusterSetslotSubcommandMigrating) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandNode Completed

func (c ClusterSetslotSubcommandNode) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c ClusterSetslotSubcommandNode) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandStable Completed

func (c ClusterSetslotSubcommandStable) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c ClusterSetslotSubcommandStable) Build() Completed {
	return Completed(c)
}

type ClusterSlaves Completed

func (c ClusterSlaves) NodeId(NodeId string) ClusterSlavesNodeId {
	return ClusterSlavesNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *Builder) ClusterSlaves() (c ClusterSlaves) {
	c.cs = append(b.get(), "CLUSTER", "SLAVES")
	return
}

type ClusterSlavesNodeId Completed

func (c ClusterSlavesNodeId) Build() Completed {
	return Completed(c)
}

type ClusterSlots Completed

func (c ClusterSlots) Build() Completed {
	return Completed(c)
}

func (b *Builder) ClusterSlots() (c ClusterSlots) {
	c.cs = append(b.get(), "CLUSTER", "SLOTS")
	return
}

type Command Completed

func (c Command) Build() Completed {
	return Completed(c)
}

func (b *Builder) Command() (c Command) {
	c.cs = append(b.get(), "COMMAND")
	return
}

type CommandCount Completed

func (c CommandCount) Build() Completed {
	return Completed(c)
}

func (b *Builder) CommandCount() (c CommandCount) {
	c.cs = append(b.get(), "COMMAND", "COUNT")
	return
}

type CommandGetkeys Completed

func (c CommandGetkeys) Build() Completed {
	return Completed(c)
}

func (b *Builder) CommandGetkeys() (c CommandGetkeys) {
	c.cs = append(b.get(), "COMMAND", "GETKEYS")
	return
}

type CommandInfo Completed

func (c CommandInfo) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cs: append(c.cs, CommandName...), cf: c.cf, ks: c.ks}
}

func (b *Builder) CommandInfo() (c CommandInfo) {
	c.cs = append(b.get(), "COMMAND", "INFO")
	return
}

type CommandInfoCommandName Completed

func (c CommandInfoCommandName) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cs: append(c.cs, CommandName...), cf: c.cf, ks: c.ks}
}

func (c CommandInfoCommandName) Build() Completed {
	return Completed(c)
}

type ConfigGet Completed

func (c ConfigGet) Parameter(Parameter string) ConfigGetParameter {
	return ConfigGetParameter{cs: append(c.cs, Parameter), cf: c.cf, ks: c.ks}
}

func (b *Builder) ConfigGet() (c ConfigGet) {
	c.cs = append(b.get(), "CONFIG", "GET")
	return
}

type ConfigGetParameter Completed

func (c ConfigGetParameter) Build() Completed {
	return Completed(c)
}

type ConfigResetstat Completed

func (c ConfigResetstat) Build() Completed {
	return Completed(c)
}

func (b *Builder) ConfigResetstat() (c ConfigResetstat) {
	c.cs = append(b.get(), "CONFIG", "RESETSTAT")
	return
}

type ConfigRewrite Completed

func (c ConfigRewrite) Build() Completed {
	return Completed(c)
}

func (b *Builder) ConfigRewrite() (c ConfigRewrite) {
	c.cs = append(b.get(), "CONFIG", "REWRITE")
	return
}

type ConfigSet Completed

func (c ConfigSet) Parameter(Parameter string) ConfigSetParameter {
	return ConfigSetParameter{cs: append(c.cs, Parameter), cf: c.cf, ks: c.ks}
}

func (b *Builder) ConfigSet() (c ConfigSet) {
	c.cs = append(b.get(), "CONFIG", "SET")
	return
}

type ConfigSetParameter Completed

func (c ConfigSetParameter) Value(Value string) ConfigSetValue {
	return ConfigSetValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type ConfigSetValue Completed

func (c ConfigSetValue) Build() Completed {
	return Completed(c)
}

type Copy Completed

func (c Copy) Source(Source string) CopySource {
	return CopySource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *Builder) Copy() (c Copy) {
	c.cs = append(b.get(), "COPY")
	return
}

type CopyDb Completed

func (c CopyDb) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c CopyDb) Build() Completed {
	return Completed(c)
}

type CopyDestination Completed

func (c CopyDestination) Db(DestinationDb int64) CopyDb {
	return CopyDb{cs: append(c.cs, "DB", strconv.FormatInt(DestinationDb, 10)), cf: c.cf, ks: c.ks}
}

func (c CopyDestination) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c CopyDestination) Build() Completed {
	return Completed(c)
}

type CopyReplaceReplace Completed

func (c CopyReplaceReplace) Build() Completed {
	return Completed(c)
}

type CopySource Completed

func (c CopySource) Destination(Destination string) CopyDestination {
	return CopyDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type Dbsize Completed

func (c Dbsize) Build() Completed {
	return Completed(c)
}

func (b *Builder) Dbsize() (c Dbsize) {
	c.cs = append(b.get(), "DBSIZE")
	c.cf = readonly
	return
}

type DebugObject Completed

func (c DebugObject) Key(Key string) DebugObjectKey {
	return DebugObjectKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) DebugObject() (c DebugObject) {
	c.cs = append(b.get(), "DEBUG", "OBJECT")
	return
}

type DebugObjectKey Completed

func (c DebugObjectKey) Build() Completed {
	return Completed(c)
}

type DebugSegfault Completed

func (c DebugSegfault) Build() Completed {
	return Completed(c)
}

func (b *Builder) DebugSegfault() (c DebugSegfault) {
	c.cs = append(b.get(), "DEBUG", "SEGFAULT")
	return
}

type Decr Completed

func (c Decr) Key(Key string) DecrKey {
	return DecrKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Decr() (c Decr) {
	c.cs = append(b.get(), "DECR")
	return
}

type DecrKey Completed

func (c DecrKey) Build() Completed {
	return Completed(c)
}

type Decrby Completed

func (c Decrby) Key(Key string) DecrbyKey {
	return DecrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Decrby() (c Decrby) {
	c.cs = append(b.get(), "DECRBY")
	return
}

type DecrbyDecrement Completed

func (c DecrbyDecrement) Build() Completed {
	return Completed(c)
}

type DecrbyKey Completed

func (c DecrbyKey) Decrement(Decrement int64) DecrbyDecrement {
	return DecrbyDecrement{cs: append(c.cs, strconv.FormatInt(Decrement, 10)), cf: c.cf, ks: c.ks}
}

type Del Completed

func (c Del) Key(Key ...string) DelKey {
	return DelKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Del() (c Del) {
	c.cs = append(b.get(), "DEL")
	return
}

type DelKey Completed

func (c DelKey) Key(Key ...string) DelKey {
	return DelKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c DelKey) Build() Completed {
	return Completed(c)
}

type Discard Completed

func (c Discard) Build() Completed {
	return Completed(c)
}

func (b *Builder) Discard() (c Discard) {
	c.cs = append(b.get(), "DISCARD")
	return
}

type Dump Completed

func (c Dump) Key(Key string) DumpKey {
	return DumpKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Dump() (c Dump) {
	c.cs = append(b.get(), "DUMP")
	c.cf = readonly
	return
}

type DumpKey Completed

func (c DumpKey) Build() Completed {
	return Completed(c)
}

type Echo Completed

func (c Echo) Message(Message string) EchoMessage {
	return EchoMessage{cs: append(c.cs, Message), cf: c.cf, ks: c.ks}
}

func (b *Builder) Echo() (c Echo) {
	c.cs = append(b.get(), "ECHO")
	return
}

type EchoMessage Completed

func (c EchoMessage) Build() Completed {
	return Completed(c)
}

type Eval Completed

func (c Eval) Script(Script string) EvalScript {
	return EvalScript{cs: append(c.cs, Script), cf: c.cf, ks: c.ks}
}

func (b *Builder) Eval() (c Eval) {
	c.cs = append(b.get(), "EVAL")
	return
}

type EvalArg Completed

func (c EvalArg) Arg(Arg ...string) EvalArg {
	return EvalArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalArg) Build() Completed {
	return Completed(c)
}

type EvalKey Completed

func (c EvalKey) Arg(Arg ...string) EvalArg {
	return EvalArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalKey) Key(Key ...string) EvalKey {
	return EvalKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c EvalKey) Build() Completed {
	return Completed(c)
}

type EvalNumkeys Completed

func (c EvalNumkeys) Key(Key ...string) EvalKey {
	return EvalKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c EvalNumkeys) Arg(Arg ...string) EvalArg {
	return EvalArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalNumkeys) Build() Completed {
	return Completed(c)
}

type EvalRo Completed

func (c EvalRo) Script(Script string) EvalRoScript {
	return EvalRoScript{cs: append(c.cs, Script), cf: c.cf, ks: c.ks}
}

func (b *Builder) EvalRo() (c EvalRo) {
	c.cs = append(b.get(), "EVAL_RO")
	c.cf = readonly
	return
}

type EvalRoArg Completed

func (c EvalRoArg) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalRoArg) Build() Completed {
	return Completed(c)
}

type EvalRoKey Completed

func (c EvalRoKey) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalRoKey) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type EvalRoNumkeys Completed

func (c EvalRoNumkeys) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type EvalRoScript Completed

func (c EvalRoScript) Numkeys(Numkeys int64) EvalRoNumkeys {
	return EvalRoNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type EvalScript Completed

func (c EvalScript) Numkeys(Numkeys int64) EvalNumkeys {
	return EvalNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type Evalsha Completed

func (c Evalsha) Sha1(Sha1 string) EvalshaSha1 {
	return EvalshaSha1{cs: append(c.cs, Sha1), cf: c.cf, ks: c.ks}
}

func (b *Builder) Evalsha() (c Evalsha) {
	c.cs = append(b.get(), "EVALSHA")
	return
}

type EvalshaArg Completed

func (c EvalshaArg) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalshaArg) Build() Completed {
	return Completed(c)
}

type EvalshaKey Completed

func (c EvalshaKey) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalshaKey) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c EvalshaKey) Build() Completed {
	return Completed(c)
}

type EvalshaNumkeys Completed

func (c EvalshaNumkeys) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c EvalshaNumkeys) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalshaNumkeys) Build() Completed {
	return Completed(c)
}

type EvalshaRo Completed

func (c EvalshaRo) Sha1(Sha1 string) EvalshaRoSha1 {
	return EvalshaRoSha1{cs: append(c.cs, Sha1), cf: c.cf, ks: c.ks}
}

func (b *Builder) EvalshaRo() (c EvalshaRo) {
	c.cs = append(b.get(), "EVALSHA_RO")
	c.cf = readonly
	return
}

type EvalshaRoArg Completed

func (c EvalshaRoArg) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalshaRoArg) Build() Completed {
	return Completed(c)
}

type EvalshaRoKey Completed

func (c EvalshaRoKey) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c EvalshaRoKey) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type EvalshaRoNumkeys Completed

func (c EvalshaRoNumkeys) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type EvalshaRoSha1 Completed

func (c EvalshaRoSha1) Numkeys(Numkeys int64) EvalshaRoNumkeys {
	return EvalshaRoNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type EvalshaSha1 Completed

func (c EvalshaSha1) Numkeys(Numkeys int64) EvalshaNumkeys {
	return EvalshaNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type Exec Completed

func (c Exec) Build() Completed {
	return Completed(c)
}

func (b *Builder) Exec() (c Exec) {
	c.cs = append(b.get(), "EXEC")
	return
}

type Exists Completed

func (c Exists) Key(Key ...string) ExistsKey {
	return ExistsKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Exists() (c Exists) {
	c.cs = append(b.get(), "EXISTS")
	c.cf = readonly
	return
}

type ExistsKey Completed

func (c ExistsKey) Key(Key ...string) ExistsKey {
	return ExistsKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ExistsKey) Build() Completed {
	return Completed(c)
}

type Expire Completed

func (c Expire) Key(Key string) ExpireKey {
	return ExpireKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Expire() (c Expire) {
	c.cs = append(b.get(), "EXPIRE")
	return
}

type ExpireConditionGt Completed

func (c ExpireConditionGt) Build() Completed {
	return Completed(c)
}

type ExpireConditionLt Completed

func (c ExpireConditionLt) Build() Completed {
	return Completed(c)
}

type ExpireConditionNx Completed

func (c ExpireConditionNx) Build() Completed {
	return Completed(c)
}

type ExpireConditionXx Completed

func (c ExpireConditionXx) Build() Completed {
	return Completed(c)
}

type ExpireKey Completed

func (c ExpireKey) Seconds(Seconds int64) ExpireSeconds {
	return ExpireSeconds{cs: append(c.cs, strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

type ExpireSeconds Completed

func (c ExpireSeconds) Nx() ExpireConditionNx {
	return ExpireConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c ExpireSeconds) Xx() ExpireConditionXx {
	return ExpireConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c ExpireSeconds) Gt() ExpireConditionGt {
	return ExpireConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c ExpireSeconds) Lt() ExpireConditionLt {
	return ExpireConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c ExpireSeconds) Build() Completed {
	return Completed(c)
}

type Expireat Completed

func (c Expireat) Key(Key string) ExpireatKey {
	return ExpireatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Expireat() (c Expireat) {
	c.cs = append(b.get(), "EXPIREAT")
	return
}

type ExpireatConditionGt Completed

func (c ExpireatConditionGt) Build() Completed {
	return Completed(c)
}

type ExpireatConditionLt Completed

func (c ExpireatConditionLt) Build() Completed {
	return Completed(c)
}

type ExpireatConditionNx Completed

func (c ExpireatConditionNx) Build() Completed {
	return Completed(c)
}

type ExpireatConditionXx Completed

func (c ExpireatConditionXx) Build() Completed {
	return Completed(c)
}

type ExpireatKey Completed

func (c ExpireatKey) Timestamp(Timestamp int64) ExpireatTimestamp {
	return ExpireatTimestamp{cs: append(c.cs, strconv.FormatInt(Timestamp, 10)), cf: c.cf, ks: c.ks}
}

type ExpireatTimestamp Completed

func (c ExpireatTimestamp) Nx() ExpireatConditionNx {
	return ExpireatConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c ExpireatTimestamp) Xx() ExpireatConditionXx {
	return ExpireatConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c ExpireatTimestamp) Gt() ExpireatConditionGt {
	return ExpireatConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c ExpireatTimestamp) Lt() ExpireatConditionLt {
	return ExpireatConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c ExpireatTimestamp) Build() Completed {
	return Completed(c)
}

type Expiretime Completed

func (c Expiretime) Key(Key string) ExpiretimeKey {
	return ExpiretimeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Expiretime() (c Expiretime) {
	c.cs = append(b.get(), "EXPIRETIME")
	c.cf = readonly
	return
}

type ExpiretimeKey Completed

func (c ExpiretimeKey) Build() Completed {
	return Completed(c)
}

func (c ExpiretimeKey) Cache() Cacheable {
	return Cacheable(c)
}

type Failover Completed

func (c Failover) To() FailoverTargetTo {
	return FailoverTargetTo{cs: append(c.cs, "TO"), cf: c.cf, ks: c.ks}
}

func (c Failover) Abort() FailoverAbort {
	return FailoverAbort{cs: append(c.cs, "ABORT"), cf: c.cf, ks: c.ks}
}

func (c Failover) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c Failover) Build() Completed {
	return Completed(c)
}

func (b *Builder) Failover() (c Failover) {
	c.cs = append(b.get(), "FAILOVER")
	return
}

type FailoverAbort Completed

func (c FailoverAbort) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c FailoverAbort) Build() Completed {
	return Completed(c)
}

type FailoverTargetForce Completed

func (c FailoverTargetForce) Abort() FailoverAbort {
	return FailoverAbort{cs: append(c.cs, "ABORT"), cf: c.cf, ks: c.ks}
}

func (c FailoverTargetForce) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c FailoverTargetForce) Build() Completed {
	return Completed(c)
}

type FailoverTargetHost Completed

func (c FailoverTargetHost) Port(Port int64) FailoverTargetPort {
	return FailoverTargetPort{cs: append(c.cs, strconv.FormatInt(Port, 10)), cf: c.cf, ks: c.ks}
}

type FailoverTargetPort Completed

func (c FailoverTargetPort) Force() FailoverTargetForce {
	return FailoverTargetForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c FailoverTargetPort) Abort() FailoverAbort {
	return FailoverAbort{cs: append(c.cs, "ABORT"), cf: c.cf, ks: c.ks}
}

func (c FailoverTargetPort) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c FailoverTargetPort) Build() Completed {
	return Completed(c)
}

type FailoverTargetTo Completed

func (c FailoverTargetTo) Host(Host string) FailoverTargetHost {
	return FailoverTargetHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

type FailoverTimeout Completed

func (c FailoverTimeout) Build() Completed {
	return Completed(c)
}

type Flushall Completed

func (c Flushall) Async() FlushallAsyncAsync {
	return FlushallAsyncAsync{cs: append(c.cs, "ASYNC"), cf: c.cf, ks: c.ks}
}

func (c Flushall) Sync() FlushallAsyncSync {
	return FlushallAsyncSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c Flushall) Build() Completed {
	return Completed(c)
}

func (b *Builder) Flushall() (c Flushall) {
	c.cs = append(b.get(), "FLUSHALL")
	return
}

type FlushallAsyncAsync Completed

func (c FlushallAsyncAsync) Build() Completed {
	return Completed(c)
}

type FlushallAsyncSync Completed

func (c FlushallAsyncSync) Build() Completed {
	return Completed(c)
}

type Flushdb Completed

func (c Flushdb) Async() FlushdbAsyncAsync {
	return FlushdbAsyncAsync{cs: append(c.cs, "ASYNC"), cf: c.cf, ks: c.ks}
}

func (c Flushdb) Sync() FlushdbAsyncSync {
	return FlushdbAsyncSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c Flushdb) Build() Completed {
	return Completed(c)
}

func (b *Builder) Flushdb() (c Flushdb) {
	c.cs = append(b.get(), "FLUSHDB")
	return
}

type FlushdbAsyncAsync Completed

func (c FlushdbAsyncAsync) Build() Completed {
	return Completed(c)
}

type FlushdbAsyncSync Completed

func (c FlushdbAsyncSync) Build() Completed {
	return Completed(c)
}

type Geoadd Completed

func (c Geoadd) Key(Key string) GeoaddKey {
	return GeoaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Geoadd() (c Geoadd) {
	c.cs = append(b.get(), "GEOADD")
	return
}

type GeoaddChangeCh Completed

func (c GeoaddChangeCh) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type GeoaddConditionNx Completed

func (c GeoaddConditionNx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c GeoaddConditionNx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type GeoaddConditionXx Completed

func (c GeoaddConditionXx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c GeoaddConditionXx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type GeoaddKey Completed

func (c GeoaddKey) Nx() GeoaddConditionNx {
	return GeoaddConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c GeoaddKey) Xx() GeoaddConditionXx {
	return GeoaddConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c GeoaddKey) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c GeoaddKey) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type GeoaddLongitudeLatitudeMember Completed

func (c GeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member), cf: c.cf, ks: c.ks}
}

func (c GeoaddLongitudeLatitudeMember) Build() Completed {
	return Completed(c)
}

type Geodist Completed

func (c Geodist) Key(Key string) GeodistKey {
	return GeodistKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Geodist() (c Geodist) {
	c.cs = append(b.get(), "GEODIST")
	c.cf = readonly
	return
}

type GeodistKey Completed

func (c GeodistKey) Member1(Member1 string) GeodistMember1 {
	return GeodistMember1{cs: append(c.cs, Member1), cf: c.cf, ks: c.ks}
}

type GeodistMember1 Completed

func (c GeodistMember1) Member2(Member2 string) GeodistMember2 {
	return GeodistMember2{cs: append(c.cs, Member2), cf: c.cf, ks: c.ks}
}

type GeodistMember2 Completed

func (c GeodistMember2) M() GeodistUnitM {
	return GeodistUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeodistMember2) Km() GeodistUnitKm {
	return GeodistUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeodistMember2) Ft() GeodistUnitFt {
	return GeodistUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeodistMember2) Mi() GeodistUnitMi {
	return GeodistUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

func (c GeodistMember2) Build() Completed {
	return Completed(c)
}

func (c GeodistMember2) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitFt Completed

func (c GeodistUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitKm Completed

func (c GeodistUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitM Completed

func (c GeodistUnitM) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitMi Completed

func (c GeodistUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type Geohash Completed

func (c Geohash) Key(Key string) GeohashKey {
	return GeohashKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Geohash() (c Geohash) {
	c.cs = append(b.get(), "GEOHASH")
	c.cf = readonly
	return
}

type GeohashKey Completed

func (c GeohashKey) Member(Member ...string) GeohashMember {
	return GeohashMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type GeohashMember Completed

func (c GeohashMember) Member(Member ...string) GeohashMember {
	return GeohashMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c GeohashMember) Build() Completed {
	return Completed(c)
}

func (c GeohashMember) Cache() Cacheable {
	return Cacheable(c)
}

type Geopos Completed

func (c Geopos) Key(Key string) GeoposKey {
	return GeoposKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Geopos() (c Geopos) {
	c.cs = append(b.get(), "GEOPOS")
	c.cf = readonly
	return
}

type GeoposKey Completed

func (c GeoposKey) Member(Member ...string) GeoposMember {
	return GeoposMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type GeoposMember Completed

func (c GeoposMember) Member(Member ...string) GeoposMember {
	return GeoposMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c GeoposMember) Build() Completed {
	return Completed(c)
}

func (c GeoposMember) Cache() Cacheable {
	return Cacheable(c)
}

type Georadius Completed

func (c Georadius) Key(Key string) GeoradiusKey {
	return GeoradiusKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Georadius() (c Georadius) {
	c.cs = append(b.get(), "GEORADIUS")
	return
}

type GeoradiusCountAnyAny Completed

func (c GeoradiusCountAnyAny) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountAnyAny) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountAnyAny) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountAnyAny) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeoradiusCountCount Completed

func (c GeoradiusCountCount) Any() GeoradiusCountAnyAny {
	return GeoradiusCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountCount) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountCount) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountCount) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountCount) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusCountCount) Build() Completed {
	return Completed(c)
}

type GeoradiusKey Completed

func (c GeoradiusKey) Longitude(Longitude float64) GeoradiusLongitude {
	return GeoradiusLongitude{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusLatitude Completed

func (c GeoradiusLatitude) Radius(Radius float64) GeoradiusRadius {
	return GeoradiusRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusLongitude Completed

func (c GeoradiusLongitude) Latitude(Latitude float64) GeoradiusLatitude {
	return GeoradiusLatitude{cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusOrderAsc Completed

func (c GeoradiusOrderAsc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusOrderAsc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusOrderAsc) Build() Completed {
	return Completed(c)
}

type GeoradiusOrderDesc Completed

func (c GeoradiusOrderDesc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusOrderDesc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusOrderDesc) Build() Completed {
	return Completed(c)
}

type GeoradiusRadius Completed

func (c GeoradiusRadius) M() GeoradiusUnitM {
	return GeoradiusUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRadius) Km() GeoradiusUnitKm {
	return GeoradiusUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRadius) Ft() GeoradiusUnitFt {
	return GeoradiusUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRadius) Mi() GeoradiusUnitMi {
	return GeoradiusUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeoradiusRo Completed

func (c GeoradiusRo) Key(Key string) GeoradiusRoKey {
	return GeoradiusRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) GeoradiusRo() (c GeoradiusRo) {
	c.cs = append(b.get(), "GEORADIUS_RO")
	c.cf = readonly
	return
}

type GeoradiusRoCountAnyAny Completed

func (c GeoradiusRoCountAnyAny) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoCountAnyAny) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoCountAnyAny) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoCountAnyAny) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoCountAnyAny) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoCountCount Completed

func (c GeoradiusRoCountCount) Any() GeoradiusRoCountAnyAny {
	return GeoradiusRoCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoCountCount) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoCountCount) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoCountCount) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoCountCount) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoCountCount) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoKey Completed

func (c GeoradiusRoKey) Longitude(Longitude float64) GeoradiusRoLongitude {
	return GeoradiusRoLongitude{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusRoLatitude Completed

func (c GeoradiusRoLatitude) Radius(Radius float64) GeoradiusRoRadius {
	return GeoradiusRoRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusRoLongitude Completed

func (c GeoradiusRoLongitude) Latitude(Latitude float64) GeoradiusRoLatitude {
	return GeoradiusRoLatitude{cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusRoOrderAsc Completed

func (c GeoradiusRoOrderAsc) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoOrderDesc Completed

func (c GeoradiusRoOrderDesc) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoRadius Completed

func (c GeoradiusRoRadius) M() GeoradiusRoUnitM {
	return GeoradiusRoUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoRadius) Km() GeoradiusRoUnitKm {
	return GeoradiusRoUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoRadius) Ft() GeoradiusRoUnitFt {
	return GeoradiusRoUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoRadius) Mi() GeoradiusRoUnitMi {
	return GeoradiusRoUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeoradiusRoStoredist Completed

func (c GeoradiusRoStoredist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoStoredist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitFt Completed

func (c GeoradiusRoUnitFt) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitFt) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitFt) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitFt) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitFt) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitFt) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitFt) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitKm Completed

func (c GeoradiusRoUnitKm) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitKm) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitKm) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitKm) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitKm) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitKm) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitKm) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitM Completed

func (c GeoradiusRoUnitM) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitM) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitM) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitM) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitM) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitM) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitM) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitM) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitMi Completed

func (c GeoradiusRoUnitMi) Withcoord() GeoradiusRoWithcoordWithcoord {
	return GeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitMi) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitMi) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitMi) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitMi) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitMi) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitMi) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithcoordWithcoord Completed

func (c GeoradiusRoWithcoordWithcoord) Withdist() GeoradiusRoWithdistWithdist {
	return GeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithcoordWithcoord) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithcoordWithcoord) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithcoordWithcoord) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithcoordWithcoord) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithcoordWithcoord) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithdistWithdist Completed

func (c GeoradiusRoWithdistWithdist) Withhash() GeoradiusRoWithhashWithhash {
	return GeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithdistWithdist) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithdistWithdist) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithdistWithdist) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithdistWithdist) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithdistWithdist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithhashWithhash Completed

func (c GeoradiusRoWithhashWithhash) Count(Count int64) GeoradiusRoCountCount {
	return GeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithhashWithhash) Asc() GeoradiusRoOrderAsc {
	return GeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithhashWithhash) Desc() GeoradiusRoOrderDesc {
	return GeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithhashWithhash) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusRoWithhashWithhash) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusStore Completed

func (c GeoradiusStore) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusStore) Build() Completed {
	return Completed(c)
}

type GeoradiusStoredist Completed

func (c GeoradiusStoredist) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitFt Completed

func (c GeoradiusUnitFt) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitFt) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitKm Completed

func (c GeoradiusUnitKm) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitKm) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitM Completed

func (c GeoradiusUnitM) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitM) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitMi Completed

func (c GeoradiusUnitMi) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusUnitMi) Build() Completed {
	return Completed(c)
}

type GeoradiusWithcoordWithcoord Completed

func (c GeoradiusWithcoordWithcoord) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithcoordWithcoord) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithcoordWithcoord) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithcoordWithcoord) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithcoordWithcoord) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithcoordWithcoord) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithcoordWithcoord) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

type GeoradiusWithdistWithdist Completed

func (c GeoradiusWithdistWithdist) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithdistWithdist) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithdistWithdist) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithdistWithdist) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithdistWithdist) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithdistWithdist) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithdistWithdist) Build() Completed {
	return Completed(c)
}

type GeoradiusWithhashWithhash Completed

func (c GeoradiusWithhashWithhash) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithhashWithhash) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithhashWithhash) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithhashWithhash) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithhashWithhash) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusWithhashWithhash) Build() Completed {
	return Completed(c)
}

type Georadiusbymember Completed

func (c Georadiusbymember) Key(Key string) GeoradiusbymemberKey {
	return GeoradiusbymemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Georadiusbymember() (c Georadiusbymember) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER")
	return
}

type GeoradiusbymemberCountAnyAny Completed

func (c GeoradiusbymemberCountAnyAny) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountAnyAny) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountAnyAny) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountAnyAny) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberCountCount Completed

func (c GeoradiusbymemberCountCount) Any() GeoradiusbymemberCountAnyAny {
	return GeoradiusbymemberCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountCount) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountCount) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountCount) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountCount) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberCountCount) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberKey Completed

func (c GeoradiusbymemberKey) Member(Member string) GeoradiusbymemberMember {
	return GeoradiusbymemberMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type GeoradiusbymemberMember Completed

func (c GeoradiusbymemberMember) Radius(Radius float64) GeoradiusbymemberRadius {
	return GeoradiusbymemberRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusbymemberOrderAsc Completed

func (c GeoradiusbymemberOrderAsc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberOrderAsc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberOrderAsc) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberOrderDesc Completed

func (c GeoradiusbymemberOrderDesc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberOrderDesc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberOrderDesc) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberRadius Completed

func (c GeoradiusbymemberRadius) M() GeoradiusbymemberUnitM {
	return GeoradiusbymemberUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRadius) Km() GeoradiusbymemberUnitKm {
	return GeoradiusbymemberUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRadius) Ft() GeoradiusbymemberUnitFt {
	return GeoradiusbymemberUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRadius) Mi() GeoradiusbymemberUnitMi {
	return GeoradiusbymemberUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeoradiusbymemberRo Completed

func (c GeoradiusbymemberRo) Key(Key string) GeoradiusbymemberRoKey {
	return GeoradiusbymemberRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) GeoradiusbymemberRo() (c GeoradiusbymemberRo) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER_RO")
	c.cf = readonly
	return
}

type GeoradiusbymemberRoCountAnyAny Completed

func (c GeoradiusbymemberRoCountAnyAny) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoCountAnyAny) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoCountAnyAny) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoCountAnyAny) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoCountAnyAny) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoCountCount Completed

func (c GeoradiusbymemberRoCountCount) Any() GeoradiusbymemberRoCountAnyAny {
	return GeoradiusbymemberRoCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoCountCount) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoCountCount) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoCountCount) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoCountCount) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoCountCount) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoKey Completed

func (c GeoradiusbymemberRoKey) Member(Member string) GeoradiusbymemberRoMember {
	return GeoradiusbymemberRoMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type GeoradiusbymemberRoMember Completed

func (c GeoradiusbymemberRoMember) Radius(Radius float64) GeoradiusbymemberRoRadius {
	return GeoradiusbymemberRoRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeoradiusbymemberRoOrderAsc Completed

func (c GeoradiusbymemberRoOrderAsc) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoOrderDesc Completed

func (c GeoradiusbymemberRoOrderDesc) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoRadius Completed

func (c GeoradiusbymemberRoRadius) M() GeoradiusbymemberRoUnitM {
	return GeoradiusbymemberRoUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoRadius) Km() GeoradiusbymemberRoUnitKm {
	return GeoradiusbymemberRoUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoRadius) Ft() GeoradiusbymemberRoUnitFt {
	return GeoradiusbymemberRoUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoRadius) Mi() GeoradiusbymemberRoUnitMi {
	return GeoradiusbymemberRoUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeoradiusbymemberRoStoredist Completed

func (c GeoradiusbymemberRoStoredist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoStoredist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitFt Completed

func (c GeoradiusbymemberRoUnitFt) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitFt) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitFt) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitFt) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitFt) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitFt) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitFt) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitKm Completed

func (c GeoradiusbymemberRoUnitKm) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitKm) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitKm) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitKm) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitKm) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitKm) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitKm) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitM Completed

func (c GeoradiusbymemberRoUnitM) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitM) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitM) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitM) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitM) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitM) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitM) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitM) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitMi Completed

func (c GeoradiusbymemberRoUnitMi) Withcoord() GeoradiusbymemberRoWithcoordWithcoord {
	return GeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitMi) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitMi) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitMi) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitMi) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitMi) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitMi) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithcoordWithcoord Completed

func (c GeoradiusbymemberRoWithcoordWithcoord) Withdist() GeoradiusbymemberRoWithdistWithdist {
	return GeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithdistWithdist Completed

func (c GeoradiusbymemberRoWithdistWithdist) Withhash() GeoradiusbymemberRoWithhashWithhash {
	return GeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithdistWithdist) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithdistWithdist) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithdistWithdist) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithdistWithdist) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithdistWithdist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithhashWithhash Completed

func (c GeoradiusbymemberRoWithhashWithhash) Count(Count int64) GeoradiusbymemberRoCountCount {
	return GeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithhashWithhash) Asc() GeoradiusbymemberRoOrderAsc {
	return GeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithhashWithhash) Desc() GeoradiusbymemberRoOrderDesc {
	return GeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithhashWithhash) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberRoWithhashWithhash) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberStore Completed

func (c GeoradiusbymemberStore) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberStore) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberStoredist Completed

func (c GeoradiusbymemberStoredist) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitFt Completed

func (c GeoradiusbymemberUnitFt) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitFt) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitKm Completed

func (c GeoradiusbymemberUnitKm) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitKm) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitM Completed

func (c GeoradiusbymemberUnitM) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitM) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitMi Completed

func (c GeoradiusbymemberUnitMi) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberUnitMi) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberWithcoordWithcoord Completed

func (c GeoradiusbymemberWithcoordWithcoord) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithcoordWithcoord) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithcoordWithcoord) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithcoordWithcoord) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithcoordWithcoord) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithcoordWithcoord) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithcoordWithcoord) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberWithdistWithdist Completed

func (c GeoradiusbymemberWithdistWithdist) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithdistWithdist) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithdistWithdist) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithdistWithdist) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithdistWithdist) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithdistWithdist) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithdistWithdist) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberWithhashWithhash Completed

func (c GeoradiusbymemberWithhashWithhash) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithhashWithhash) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithhashWithhash) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithhashWithhash) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithhashWithhash) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c GeoradiusbymemberWithhashWithhash) Build() Completed {
	return Completed(c)
}

type Geosearch Completed

func (c Geosearch) Key(Key string) GeosearchKey {
	return GeosearchKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Geosearch() (c Geosearch) {
	c.cs = append(b.get(), "GEOSEARCH")
	c.cf = readonly
	return
}

type GeosearchBoxBybox Completed

func (c GeosearchBoxBybox) Height(Height float64) GeosearchBoxHeight {
	return GeosearchBoxHeight{cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeosearchBoxHeight Completed

func (c GeosearchBoxHeight) M() GeosearchBoxUnitM {
	return GeosearchBoxUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxHeight) Km() GeosearchBoxUnitKm {
	return GeosearchBoxUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxHeight) Ft() GeosearchBoxUnitFt {
	return GeosearchBoxUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxHeight) Mi() GeosearchBoxUnitMi {
	return GeosearchBoxUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeosearchBoxUnitFt Completed

func (c GeosearchBoxUnitFt) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitFt) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitFt) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitFt) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitFt) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitFt) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitKm Completed

func (c GeosearchBoxUnitKm) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitKm) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitKm) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitKm) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitKm) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitKm) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitM Completed

func (c GeosearchBoxUnitM) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitM) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitM) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitM) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitM) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitM) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitM) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitMi Completed

func (c GeosearchBoxUnitMi) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitMi) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitMi) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitMi) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitMi) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitMi) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchBoxUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleByradius Completed

func (c GeosearchCircleByradius) M() GeosearchCircleUnitM {
	return GeosearchCircleUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleByradius) Km() GeosearchCircleUnitKm {
	return GeosearchCircleUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleByradius) Ft() GeosearchCircleUnitFt {
	return GeosearchCircleUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleByradius) Mi() GeosearchCircleUnitMi {
	return GeosearchCircleUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeosearchCircleUnitFt Completed

func (c GeosearchCircleUnitFt) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitFt) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitFt) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitFt) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitFt) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitFt) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitFt) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitKm Completed

func (c GeosearchCircleUnitKm) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitKm) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitKm) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitKm) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitKm) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitKm) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitKm) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitM Completed

func (c GeosearchCircleUnitM) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitM) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitM) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitM) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitM) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitM) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitM) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitM) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitMi Completed

func (c GeosearchCircleUnitMi) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitMi) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitMi) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitMi) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitMi) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitMi) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitMi) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCircleUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCountAnyAny Completed

func (c GeosearchCountAnyAny) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCountAnyAny) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCountAnyAny) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCountAnyAny) Build() Completed {
	return Completed(c)
}

func (c GeosearchCountAnyAny) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCountCount Completed

func (c GeosearchCountCount) Any() GeosearchCountAnyAny {
	return GeosearchCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCountCount) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCountCount) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCountCount) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchCountCount) Build() Completed {
	return Completed(c)
}

func (c GeosearchCountCount) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchFromlonlat Completed

func (c GeosearchFromlonlat) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFromlonlat) Build() Completed {
	return Completed(c)
}

func (c GeosearchFromlonlat) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchFrommember Completed

func (c GeosearchFrommember) Fromlonlat(Longitude float64, Latitude float64) GeosearchFromlonlat {
	return GeosearchFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchFrommember) Build() Completed {
	return Completed(c)
}

func (c GeosearchFrommember) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchKey Completed

func (c GeosearchKey) Frommember(Member string) GeosearchFrommember {
	return GeosearchFrommember{cs: append(c.cs, "FROMMEMBER", Member), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Fromlonlat(Longitude float64, Latitude float64) GeosearchFromlonlat {
	return GeosearchFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchKey) Build() Completed {
	return Completed(c)
}

func (c GeosearchKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchOrderAsc Completed

func (c GeosearchOrderAsc) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderAsc) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderAsc) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderAsc) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeosearchOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchOrderDesc Completed

func (c GeosearchOrderDesc) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderDesc) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderDesc) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderDesc) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeosearchOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchWithcoordWithcoord Completed

func (c GeosearchWithcoordWithcoord) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchWithcoordWithcoord) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

func (c GeosearchWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchWithdistWithdist Completed

func (c GeosearchWithdistWithdist) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c GeosearchWithdistWithdist) Build() Completed {
	return Completed(c)
}

func (c GeosearchWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchWithhashWithhash Completed

func (c GeosearchWithhashWithhash) Build() Completed {
	return Completed(c)
}

func (c GeosearchWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type Geosearchstore Completed

func (c Geosearchstore) Destination(Destination string) GeosearchstoreDestination {
	return GeosearchstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *Builder) Geosearchstore() (c Geosearchstore) {
	c.cs = append(b.get(), "GEOSEARCHSTORE")
	return
}

type GeosearchstoreBoxBybox Completed

func (c GeosearchstoreBoxBybox) Height(Height float64) GeosearchstoreBoxHeight {
	return GeosearchstoreBoxHeight{cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type GeosearchstoreBoxHeight Completed

func (c GeosearchstoreBoxHeight) M() GeosearchstoreBoxUnitM {
	return GeosearchstoreBoxUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxHeight) Km() GeosearchstoreBoxUnitKm {
	return GeosearchstoreBoxUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxHeight) Ft() GeosearchstoreBoxUnitFt {
	return GeosearchstoreBoxUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxHeight) Mi() GeosearchstoreBoxUnitMi {
	return GeosearchstoreBoxUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeosearchstoreBoxUnitFt Completed

func (c GeosearchstoreBoxUnitFt) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitFt) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitFt) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitFt) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitFt) Build() Completed {
	return Completed(c)
}

type GeosearchstoreBoxUnitKm Completed

func (c GeosearchstoreBoxUnitKm) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitKm) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitKm) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitKm) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitKm) Build() Completed {
	return Completed(c)
}

type GeosearchstoreBoxUnitM Completed

func (c GeosearchstoreBoxUnitM) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitM) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitM) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitM) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitM) Build() Completed {
	return Completed(c)
}

type GeosearchstoreBoxUnitMi Completed

func (c GeosearchstoreBoxUnitMi) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitMi) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitMi) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitMi) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreBoxUnitMi) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleByradius Completed

func (c GeosearchstoreCircleByradius) M() GeosearchstoreCircleUnitM {
	return GeosearchstoreCircleUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleByradius) Km() GeosearchstoreCircleUnitKm {
	return GeosearchstoreCircleUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleByradius) Ft() GeosearchstoreCircleUnitFt {
	return GeosearchstoreCircleUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleByradius) Mi() GeosearchstoreCircleUnitMi {
	return GeosearchstoreCircleUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type GeosearchstoreCircleUnitFt Completed

func (c GeosearchstoreCircleUnitFt) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitFt) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitFt) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitFt) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitFt) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitFt) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleUnitKm Completed

func (c GeosearchstoreCircleUnitKm) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitKm) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitKm) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitKm) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitKm) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitKm) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleUnitM Completed

func (c GeosearchstoreCircleUnitM) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitM) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitM) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitM) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitM) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitM) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleUnitMi Completed

func (c GeosearchstoreCircleUnitMi) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitMi) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitMi) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitMi) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitMi) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCircleUnitMi) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCountAnyAny Completed

func (c GeosearchstoreCountAnyAny) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCountCount Completed

func (c GeosearchstoreCountCount) Any() GeosearchstoreCountAnyAny {
	return GeosearchstoreCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCountCount) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreCountCount) Build() Completed {
	return Completed(c)
}

type GeosearchstoreDestination Completed

func (c GeosearchstoreDestination) Source(Source string) GeosearchstoreSource {
	return GeosearchstoreSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

type GeosearchstoreFromlonlat Completed

func (c GeosearchstoreFromlonlat) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFromlonlat) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFromlonlat) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFromlonlat) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFromlonlat) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFromlonlat) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFromlonlat) Build() Completed {
	return Completed(c)
}

type GeosearchstoreFrommember Completed

func (c GeosearchstoreFrommember) Fromlonlat(Longitude float64, Latitude float64) GeosearchstoreFromlonlat {
	return GeosearchstoreFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFrommember) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFrommember) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFrommember) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFrommember) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFrommember) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFrommember) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreFrommember) Build() Completed {
	return Completed(c)
}

type GeosearchstoreOrderAsc Completed

func (c GeosearchstoreOrderAsc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreOrderAsc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreOrderAsc) Build() Completed {
	return Completed(c)
}

type GeosearchstoreOrderDesc Completed

func (c GeosearchstoreOrderDesc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreOrderDesc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreOrderDesc) Build() Completed {
	return Completed(c)
}

type GeosearchstoreSource Completed

func (c GeosearchstoreSource) Frommember(Member string) GeosearchstoreFrommember {
	return GeosearchstoreFrommember{cs: append(c.cs, "FROMMEMBER", Member), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Fromlonlat(Longitude float64, Latitude float64) GeosearchstoreFromlonlat {
	return GeosearchstoreFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c GeosearchstoreSource) Build() Completed {
	return Completed(c)
}

type GeosearchstoreStoredistStoredist Completed

func (c GeosearchstoreStoredistStoredist) Build() Completed {
	return Completed(c)
}

type Get Completed

func (c Get) Key(Key string) GetKey {
	return GetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Get() (c Get) {
	c.cs = append(b.get(), "GET")
	c.cf = readonly
	return
}

type GetKey Completed

func (c GetKey) Build() Completed {
	return Completed(c)
}

func (c GetKey) Cache() Cacheable {
	return Cacheable(c)
}

type Getbit Completed

func (c Getbit) Key(Key string) GetbitKey {
	return GetbitKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Getbit() (c Getbit) {
	c.cs = append(b.get(), "GETBIT")
	c.cf = readonly
	return
}

type GetbitKey Completed

func (c GetbitKey) Offset(Offset int64) GetbitOffset {
	return GetbitOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type GetbitOffset Completed

func (c GetbitOffset) Build() Completed {
	return Completed(c)
}

func (c GetbitOffset) Cache() Cacheable {
	return Cacheable(c)
}

type Getdel Completed

func (c Getdel) Key(Key string) GetdelKey {
	return GetdelKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Getdel() (c Getdel) {
	c.cs = append(b.get(), "GETDEL")
	return
}

type GetdelKey Completed

func (c GetdelKey) Build() Completed {
	return Completed(c)
}

type Getex Completed

func (c Getex) Key(Key string) GetexKey {
	return GetexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Getex() (c Getex) {
	c.cs = append(b.get(), "GETEX")
	return
}

type GetexExpirationEx Completed

func (c GetexExpirationEx) Build() Completed {
	return Completed(c)
}

type GetexExpirationExat Completed

func (c GetexExpirationExat) Build() Completed {
	return Completed(c)
}

type GetexExpirationPersist Completed

func (c GetexExpirationPersist) Build() Completed {
	return Completed(c)
}

type GetexExpirationPx Completed

func (c GetexExpirationPx) Build() Completed {
	return Completed(c)
}

type GetexExpirationPxat Completed

func (c GetexExpirationPxat) Build() Completed {
	return Completed(c)
}

type GetexKey Completed

func (c GetexKey) Ex(Seconds int64) GetexExpirationEx {
	return GetexExpirationEx{cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c GetexKey) Px(Milliseconds int64) GetexExpirationPx {
	return GetexExpirationPx{cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c GetexKey) Exat(Timestamp int64) GetexExpirationExat {
	return GetexExpirationExat{cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c GetexKey) Pxat(Millisecondstimestamp int64) GetexExpirationPxat {
	return GetexExpirationPxat{cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c GetexKey) Persist() GetexExpirationPersist {
	return GetexExpirationPersist{cs: append(c.cs, "PERSIST"), cf: c.cf, ks: c.ks}
}

func (c GetexKey) Build() Completed {
	return Completed(c)
}

type Getrange Completed

func (c Getrange) Key(Key string) GetrangeKey {
	return GetrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Getrange() (c Getrange) {
	c.cs = append(b.get(), "GETRANGE")
	c.cf = readonly
	return
}

type GetrangeEnd Completed

func (c GetrangeEnd) Build() Completed {
	return Completed(c)
}

func (c GetrangeEnd) Cache() Cacheable {
	return Cacheable(c)
}

type GetrangeKey Completed

func (c GetrangeKey) Start(Start int64) GetrangeStart {
	return GetrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type GetrangeStart Completed

func (c GetrangeStart) End(End int64) GetrangeEnd {
	return GetrangeEnd{cs: append(c.cs, strconv.FormatInt(End, 10)), cf: c.cf, ks: c.ks}
}

type Getset Completed

func (c Getset) Key(Key string) GetsetKey {
	return GetsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Getset() (c Getset) {
	c.cs = append(b.get(), "GETSET")
	return
}

type GetsetKey Completed

func (c GetsetKey) Value(Value string) GetsetValue {
	return GetsetValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type GetsetValue Completed

func (c GetsetValue) Build() Completed {
	return Completed(c)
}

type Hdel Completed

func (c Hdel) Key(Key string) HdelKey {
	return HdelKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hdel() (c Hdel) {
	c.cs = append(b.get(), "HDEL")
	return
}

type HdelField Completed

func (c HdelField) Field(Field ...string) HdelField {
	return HdelField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

func (c HdelField) Build() Completed {
	return Completed(c)
}

type HdelKey Completed

func (c HdelKey) Field(Field ...string) HdelField {
	return HdelField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

type Hello Completed

func (c Hello) Protover(Protover int64) HelloArgumentsProtover {
	return HelloArgumentsProtover{cs: append(c.cs, strconv.FormatInt(Protover, 10)), cf: c.cf, ks: c.ks}
}

func (c Hello) Build() Completed {
	return Completed(c)
}

func (b *Builder) Hello() (c Hello) {
	c.cs = append(b.get(), "HELLO")
	return
}

type HelloArgumentsAuth Completed

func (c HelloArgumentsAuth) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cs: append(c.cs, "SETNAME", Clientname), cf: c.cf, ks: c.ks}
}

func (c HelloArgumentsAuth) Build() Completed {
	return Completed(c)
}

type HelloArgumentsProtover Completed

func (c HelloArgumentsProtover) Auth(Username string, Password string) HelloArgumentsAuth {
	return HelloArgumentsAuth{cs: append(c.cs, "AUTH", Username, Password), cf: c.cf, ks: c.ks}
}

func (c HelloArgumentsProtover) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cs: append(c.cs, "SETNAME", Clientname), cf: c.cf, ks: c.ks}
}

func (c HelloArgumentsProtover) Build() Completed {
	return Completed(c)
}

type HelloArgumentsSetname Completed

func (c HelloArgumentsSetname) Build() Completed {
	return Completed(c)
}

type Hexists Completed

func (c Hexists) Key(Key string) HexistsKey {
	return HexistsKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hexists() (c Hexists) {
	c.cs = append(b.get(), "HEXISTS")
	c.cf = readonly
	return
}

type HexistsField Completed

func (c HexistsField) Build() Completed {
	return Completed(c)
}

func (c HexistsField) Cache() Cacheable {
	return Cacheable(c)
}

type HexistsKey Completed

func (c HexistsKey) Field(Field string) HexistsField {
	return HexistsField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type Hget Completed

func (c Hget) Key(Key string) HgetKey {
	return HgetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hget() (c Hget) {
	c.cs = append(b.get(), "HGET")
	c.cf = readonly
	return
}

type HgetField Completed

func (c HgetField) Build() Completed {
	return Completed(c)
}

func (c HgetField) Cache() Cacheable {
	return Cacheable(c)
}

type HgetKey Completed

func (c HgetKey) Field(Field string) HgetField {
	return HgetField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type Hgetall Completed

func (c Hgetall) Key(Key string) HgetallKey {
	return HgetallKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hgetall() (c Hgetall) {
	c.cs = append(b.get(), "HGETALL")
	c.cf = readonly
	return
}

type HgetallKey Completed

func (c HgetallKey) Build() Completed {
	return Completed(c)
}

func (c HgetallKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hincrby Completed

func (c Hincrby) Key(Key string) HincrbyKey {
	return HincrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hincrby() (c Hincrby) {
	c.cs = append(b.get(), "HINCRBY")
	return
}

type HincrbyField Completed

func (c HincrbyField) Increment(Increment int64) HincrbyIncrement {
	return HincrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

type HincrbyIncrement Completed

func (c HincrbyIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyKey Completed

func (c HincrbyKey) Field(Field string) HincrbyField {
	return HincrbyField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type Hincrbyfloat Completed

func (c Hincrbyfloat) Key(Key string) HincrbyfloatKey {
	return HincrbyfloatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hincrbyfloat() (c Hincrbyfloat) {
	c.cs = append(b.get(), "HINCRBYFLOAT")
	return
}

type HincrbyfloatField Completed

func (c HincrbyfloatField) Increment(Increment float64) HincrbyfloatIncrement {
	return HincrbyfloatIncrement{cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type HincrbyfloatIncrement Completed

func (c HincrbyfloatIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyfloatKey Completed

func (c HincrbyfloatKey) Field(Field string) HincrbyfloatField {
	return HincrbyfloatField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type Hkeys Completed

func (c Hkeys) Key(Key string) HkeysKey {
	return HkeysKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hkeys() (c Hkeys) {
	c.cs = append(b.get(), "HKEYS")
	c.cf = readonly
	return
}

type HkeysKey Completed

func (c HkeysKey) Build() Completed {
	return Completed(c)
}

func (c HkeysKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hlen Completed

func (c Hlen) Key(Key string) HlenKey {
	return HlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hlen() (c Hlen) {
	c.cs = append(b.get(), "HLEN")
	c.cf = readonly
	return
}

type HlenKey Completed

func (c HlenKey) Build() Completed {
	return Completed(c)
}

func (c HlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hmget Completed

func (c Hmget) Key(Key string) HmgetKey {
	return HmgetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hmget() (c Hmget) {
	c.cs = append(b.get(), "HMGET")
	c.cf = readonly
	return
}

type HmgetField Completed

func (c HmgetField) Field(Field ...string) HmgetField {
	return HmgetField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

func (c HmgetField) Build() Completed {
	return Completed(c)
}

func (c HmgetField) Cache() Cacheable {
	return Cacheable(c)
}

type HmgetKey Completed

func (c HmgetKey) Field(Field ...string) HmgetField {
	return HmgetField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

type Hmset Completed

func (c Hmset) Key(Key string) HmsetKey {
	return HmsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hmset() (c Hmset) {
	c.cs = append(b.get(), "HMSET")
	return
}

type HmsetFieldValue Completed

func (c HmsetFieldValue) FieldValue(Field string, Value string) HmsetFieldValue {
	return HmsetFieldValue{cs: append(c.cs, Field, Value), cf: c.cf, ks: c.ks}
}

func (c HmsetFieldValue) Build() Completed {
	return Completed(c)
}

type HmsetKey Completed

func (c HmsetKey) FieldValue() HmsetFieldValue {
	return HmsetFieldValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

type Hrandfield Completed

func (c Hrandfield) Key(Key string) HrandfieldKey {
	return HrandfieldKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hrandfield() (c Hrandfield) {
	c.cs = append(b.get(), "HRANDFIELD")
	c.cf = readonly
	return
}

type HrandfieldKey Completed

func (c HrandfieldKey) Count(Count int64) HrandfieldOptionsCount {
	return HrandfieldOptionsCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c HrandfieldKey) Build() Completed {
	return Completed(c)
}

type HrandfieldOptionsCount Completed

func (c HrandfieldOptionsCount) Withvalues() HrandfieldOptionsWithvaluesWithvalues {
	return HrandfieldOptionsWithvaluesWithvalues{cs: append(c.cs, "WITHVALUES"), cf: c.cf, ks: c.ks}
}

func (c HrandfieldOptionsCount) Build() Completed {
	return Completed(c)
}

type HrandfieldOptionsWithvaluesWithvalues Completed

func (c HrandfieldOptionsWithvaluesWithvalues) Build() Completed {
	return Completed(c)
}

type Hscan Completed

func (c Hscan) Key(Key string) HscanKey {
	return HscanKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hscan() (c Hscan) {
	c.cs = append(b.get(), "HSCAN")
	c.cf = readonly
	return
}

type HscanCount Completed

func (c HscanCount) Build() Completed {
	return Completed(c)
}

type HscanCursor Completed

func (c HscanCursor) Match(Pattern string) HscanMatch {
	return HscanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c HscanCursor) Count(Count int64) HscanCount {
	return HscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c HscanCursor) Build() Completed {
	return Completed(c)
}

type HscanKey Completed

func (c HscanKey) Cursor(Cursor int64) HscanCursor {
	return HscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

type HscanMatch Completed

func (c HscanMatch) Count(Count int64) HscanCount {
	return HscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c HscanMatch) Build() Completed {
	return Completed(c)
}

type Hset Completed

func (c Hset) Key(Key string) HsetKey {
	return HsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hset() (c Hset) {
	c.cs = append(b.get(), "HSET")
	return
}

type HsetFieldValue Completed

func (c HsetFieldValue) FieldValue(Field string, Value string) HsetFieldValue {
	return HsetFieldValue{cs: append(c.cs, Field, Value), cf: c.cf, ks: c.ks}
}

func (c HsetFieldValue) Build() Completed {
	return Completed(c)
}

type HsetKey Completed

func (c HsetKey) FieldValue() HsetFieldValue {
	return HsetFieldValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

type Hsetnx Completed

func (c Hsetnx) Key(Key string) HsetnxKey {
	return HsetnxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hsetnx() (c Hsetnx) {
	c.cs = append(b.get(), "HSETNX")
	return
}

type HsetnxField Completed

func (c HsetnxField) Value(Value string) HsetnxValue {
	return HsetnxValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type HsetnxKey Completed

func (c HsetnxKey) Field(Field string) HsetnxField {
	return HsetnxField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type HsetnxValue Completed

func (c HsetnxValue) Build() Completed {
	return Completed(c)
}

type Hstrlen Completed

func (c Hstrlen) Key(Key string) HstrlenKey {
	return HstrlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hstrlen() (c Hstrlen) {
	c.cs = append(b.get(), "HSTRLEN")
	c.cf = readonly
	return
}

type HstrlenField Completed

func (c HstrlenField) Build() Completed {
	return Completed(c)
}

func (c HstrlenField) Cache() Cacheable {
	return Cacheable(c)
}

type HstrlenKey Completed

func (c HstrlenKey) Field(Field string) HstrlenField {
	return HstrlenField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type Hvals Completed

func (c Hvals) Key(Key string) HvalsKey {
	return HvalsKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Hvals() (c Hvals) {
	c.cs = append(b.get(), "HVALS")
	c.cf = readonly
	return
}

type HvalsKey Completed

func (c HvalsKey) Build() Completed {
	return Completed(c)
}

func (c HvalsKey) Cache() Cacheable {
	return Cacheable(c)
}

type Incr Completed

func (c Incr) Key(Key string) IncrKey {
	return IncrKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Incr() (c Incr) {
	c.cs = append(b.get(), "INCR")
	return
}

type IncrKey Completed

func (c IncrKey) Build() Completed {
	return Completed(c)
}

type Incrby Completed

func (c Incrby) Key(Key string) IncrbyKey {
	return IncrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Incrby() (c Incrby) {
	c.cs = append(b.get(), "INCRBY")
	return
}

type IncrbyIncrement Completed

func (c IncrbyIncrement) Build() Completed {
	return Completed(c)
}

type IncrbyKey Completed

func (c IncrbyKey) Increment(Increment int64) IncrbyIncrement {
	return IncrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

type Incrbyfloat Completed

func (c Incrbyfloat) Key(Key string) IncrbyfloatKey {
	return IncrbyfloatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Incrbyfloat() (c Incrbyfloat) {
	c.cs = append(b.get(), "INCRBYFLOAT")
	return
}

type IncrbyfloatIncrement Completed

func (c IncrbyfloatIncrement) Build() Completed {
	return Completed(c)
}

type IncrbyfloatKey Completed

func (c IncrbyfloatKey) Increment(Increment float64) IncrbyfloatIncrement {
	return IncrbyfloatIncrement{cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type Info Completed

func (c Info) Section(Section string) InfoSection {
	return InfoSection{cs: append(c.cs, Section), cf: c.cf, ks: c.ks}
}

func (c Info) Build() Completed {
	return Completed(c)
}

func (b *Builder) Info() (c Info) {
	c.cs = append(b.get(), "INFO")
	return
}

type InfoSection Completed

func (c InfoSection) Build() Completed {
	return Completed(c)
}

type Keys Completed

func (c Keys) Pattern(Pattern string) KeysPattern {
	return KeysPattern{cs: append(c.cs, Pattern), cf: c.cf, ks: c.ks}
}

func (b *Builder) Keys() (c Keys) {
	c.cs = append(b.get(), "KEYS")
	c.cf = readonly
	return
}

type KeysPattern Completed

func (c KeysPattern) Build() Completed {
	return Completed(c)
}

type Lastsave Completed

func (c Lastsave) Build() Completed {
	return Completed(c)
}

func (b *Builder) Lastsave() (c Lastsave) {
	c.cs = append(b.get(), "LASTSAVE")
	return
}

type LatencyDoctor Completed

func (c LatencyDoctor) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyDoctor() (c LatencyDoctor) {
	c.cs = append(b.get(), "LATENCY", "DOCTOR")
	return
}

type LatencyGraph Completed

func (c LatencyGraph) Event(Event string) LatencyGraphEvent {
	return LatencyGraphEvent{cs: append(c.cs, Event), cf: c.cf, ks: c.ks}
}

func (b *Builder) LatencyGraph() (c LatencyGraph) {
	c.cs = append(b.get(), "LATENCY", "GRAPH")
	return
}

type LatencyGraphEvent Completed

func (c LatencyGraphEvent) Build() Completed {
	return Completed(c)
}

type LatencyHelp Completed

func (c LatencyHelp) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyHelp() (c LatencyHelp) {
	c.cs = append(b.get(), "LATENCY", "HELP")
	return
}

type LatencyHistory Completed

func (c LatencyHistory) Event(Event string) LatencyHistoryEvent {
	return LatencyHistoryEvent{cs: append(c.cs, Event), cf: c.cf, ks: c.ks}
}

func (b *Builder) LatencyHistory() (c LatencyHistory) {
	c.cs = append(b.get(), "LATENCY", "HISTORY")
	return
}

type LatencyHistoryEvent Completed

func (c LatencyHistoryEvent) Build() Completed {
	return Completed(c)
}

type LatencyLatest Completed

func (c LatencyLatest) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyLatest() (c LatencyLatest) {
	c.cs = append(b.get(), "LATENCY", "LATEST")
	return
}

type LatencyReset Completed

func (c LatencyReset) Event(Event ...string) LatencyResetEvent {
	return LatencyResetEvent{cs: append(c.cs, Event...), cf: c.cf, ks: c.ks}
}

func (c LatencyReset) Build() Completed {
	return Completed(c)
}

func (b *Builder) LatencyReset() (c LatencyReset) {
	c.cs = append(b.get(), "LATENCY", "RESET")
	return
}

type LatencyResetEvent Completed

func (c LatencyResetEvent) Event(Event ...string) LatencyResetEvent {
	return LatencyResetEvent{cs: append(c.cs, Event...), cf: c.cf, ks: c.ks}
}

func (c LatencyResetEvent) Build() Completed {
	return Completed(c)
}

type Lindex Completed

func (c Lindex) Key(Key string) LindexKey {
	return LindexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lindex() (c Lindex) {
	c.cs = append(b.get(), "LINDEX")
	c.cf = readonly
	return
}

type LindexIndex Completed

func (c LindexIndex) Build() Completed {
	return Completed(c)
}

func (c LindexIndex) Cache() Cacheable {
	return Cacheable(c)
}

type LindexKey Completed

func (c LindexKey) Index(Index int64) LindexIndex {
	return LindexIndex{cs: append(c.cs, strconv.FormatInt(Index, 10)), cf: c.cf, ks: c.ks}
}

type Linsert Completed

func (c Linsert) Key(Key string) LinsertKey {
	return LinsertKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Linsert() (c Linsert) {
	c.cs = append(b.get(), "LINSERT")
	return
}

type LinsertElement Completed

func (c LinsertElement) Build() Completed {
	return Completed(c)
}

type LinsertKey Completed

func (c LinsertKey) Before() LinsertWhereBefore {
	return LinsertWhereBefore{cs: append(c.cs, "BEFORE"), cf: c.cf, ks: c.ks}
}

func (c LinsertKey) After() LinsertWhereAfter {
	return LinsertWhereAfter{cs: append(c.cs, "AFTER"), cf: c.cf, ks: c.ks}
}

type LinsertPivot Completed

func (c LinsertPivot) Element(Element string) LinsertElement {
	return LinsertElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type LinsertWhereAfter Completed

func (c LinsertWhereAfter) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cs: append(c.cs, Pivot), cf: c.cf, ks: c.ks}
}

type LinsertWhereBefore Completed

func (c LinsertWhereBefore) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cs: append(c.cs, Pivot), cf: c.cf, ks: c.ks}
}

type Llen Completed

func (c Llen) Key(Key string) LlenKey {
	return LlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Llen() (c Llen) {
	c.cs = append(b.get(), "LLEN")
	c.cf = readonly
	return
}

type LlenKey Completed

func (c LlenKey) Build() Completed {
	return Completed(c)
}

func (c LlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Lmove Completed

func (c Lmove) Source(Source string) LmoveSource {
	return LmoveSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lmove() (c Lmove) {
	c.cs = append(b.get(), "LMOVE")
	return
}

type LmoveDestination Completed

func (c LmoveDestination) Left() LmoveWherefromLeft {
	return LmoveWherefromLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c LmoveDestination) Right() LmoveWherefromRight {
	return LmoveWherefromRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type LmoveSource Completed

func (c LmoveSource) Destination(Destination string) LmoveDestination {
	return LmoveDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type LmoveWherefromLeft Completed

func (c LmoveWherefromLeft) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c LmoveWherefromLeft) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type LmoveWherefromRight Completed

func (c LmoveWherefromRight) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c LmoveWherefromRight) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type LmoveWheretoLeft Completed

func (c LmoveWheretoLeft) Build() Completed {
	return Completed(c)
}

type LmoveWheretoRight Completed

func (c LmoveWheretoRight) Build() Completed {
	return Completed(c)
}

type Lmpop Completed

func (c Lmpop) Numkeys(Numkeys int64) LmpopNumkeys {
	return LmpopNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lmpop() (c Lmpop) {
	c.cs = append(b.get(), "LMPOP")
	return
}

type LmpopCount Completed

func (c LmpopCount) Build() Completed {
	return Completed(c)
}

type LmpopKey Completed

func (c LmpopKey) Left() LmpopWhereLeft {
	return LmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c LmpopKey) Right() LmpopWhereRight {
	return LmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

func (c LmpopKey) Key(Key ...string) LmpopKey {
	return LmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type LmpopNumkeys Completed

func (c LmpopNumkeys) Key(Key ...string) LmpopKey {
	return LmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c LmpopNumkeys) Left() LmpopWhereLeft {
	return LmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c LmpopNumkeys) Right() LmpopWhereRight {
	return LmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type LmpopWhereLeft Completed

func (c LmpopWhereLeft) Count(Count int64) LmpopCount {
	return LmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c LmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type LmpopWhereRight Completed

func (c LmpopWhereRight) Count(Count int64) LmpopCount {
	return LmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c LmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Lolwut Completed

func (c Lolwut) Version(Version int64) LolwutVersion {
	return LolwutVersion{cs: append(c.cs, "VERSION", strconv.FormatInt(Version, 10)), cf: c.cf, ks: c.ks}
}

func (c Lolwut) Build() Completed {
	return Completed(c)
}

func (b *Builder) Lolwut() (c Lolwut) {
	c.cs = append(b.get(), "LOLWUT")
	c.cf = readonly
	return
}

type LolwutVersion Completed

func (c LolwutVersion) Build() Completed {
	return Completed(c)
}

type Lpop Completed

func (c Lpop) Key(Key string) LpopKey {
	return LpopKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lpop() (c Lpop) {
	c.cs = append(b.get(), "LPOP")
	return
}

type LpopCount Completed

func (c LpopCount) Build() Completed {
	return Completed(c)
}

type LpopKey Completed

func (c LpopKey) Count(Count int64) LpopCount {
	return LpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c LpopKey) Build() Completed {
	return Completed(c)
}

type Lpos Completed

func (c Lpos) Key(Key string) LposKey {
	return LposKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lpos() (c Lpos) {
	c.cs = append(b.get(), "LPOS")
	c.cf = readonly
	return
}

type LposCount Completed

func (c LposCount) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10)), cf: c.cf, ks: c.ks}
}

func (c LposCount) Build() Completed {
	return Completed(c)
}

func (c LposCount) Cache() Cacheable {
	return Cacheable(c)
}

type LposElement Completed

func (c LposElement) Rank(Rank int64) LposRank {
	return LposRank{cs: append(c.cs, "RANK", strconv.FormatInt(Rank, 10)), cf: c.cf, ks: c.ks}
}

func (c LposElement) Count(NumMatches int64) LposCount {
	return LposCount{cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10)), cf: c.cf, ks: c.ks}
}

func (c LposElement) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10)), cf: c.cf, ks: c.ks}
}

func (c LposElement) Build() Completed {
	return Completed(c)
}

func (c LposElement) Cache() Cacheable {
	return Cacheable(c)
}

type LposKey Completed

func (c LposKey) Element(Element string) LposElement {
	return LposElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type LposMaxlen Completed

func (c LposMaxlen) Build() Completed {
	return Completed(c)
}

func (c LposMaxlen) Cache() Cacheable {
	return Cacheable(c)
}

type LposRank Completed

func (c LposRank) Count(NumMatches int64) LposCount {
	return LposCount{cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10)), cf: c.cf, ks: c.ks}
}

func (c LposRank) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10)), cf: c.cf, ks: c.ks}
}

func (c LposRank) Build() Completed {
	return Completed(c)
}

func (c LposRank) Cache() Cacheable {
	return Cacheable(c)
}

type Lpush Completed

func (c Lpush) Key(Key string) LpushKey {
	return LpushKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lpush() (c Lpush) {
	c.cs = append(b.get(), "LPUSH")
	return
}

type LpushElement Completed

func (c LpushElement) Element(Element ...string) LpushElement {
	return LpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c LpushElement) Build() Completed {
	return Completed(c)
}

type LpushKey Completed

func (c LpushKey) Element(Element ...string) LpushElement {
	return LpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type Lpushx Completed

func (c Lpushx) Key(Key string) LpushxKey {
	return LpushxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lpushx() (c Lpushx) {
	c.cs = append(b.get(), "LPUSHX")
	return
}

type LpushxElement Completed

func (c LpushxElement) Element(Element ...string) LpushxElement {
	return LpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c LpushxElement) Build() Completed {
	return Completed(c)
}

type LpushxKey Completed

func (c LpushxKey) Element(Element ...string) LpushxElement {
	return LpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type Lrange Completed

func (c Lrange) Key(Key string) LrangeKey {
	return LrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lrange() (c Lrange) {
	c.cs = append(b.get(), "LRANGE")
	c.cf = readonly
	return
}

type LrangeKey Completed

func (c LrangeKey) Start(Start int64) LrangeStart {
	return LrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type LrangeStart Completed

func (c LrangeStart) Stop(Stop int64) LrangeStop {
	return LrangeStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type LrangeStop Completed

func (c LrangeStop) Build() Completed {
	return Completed(c)
}

func (c LrangeStop) Cache() Cacheable {
	return Cacheable(c)
}

type Lrem Completed

func (c Lrem) Key(Key string) LremKey {
	return LremKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lrem() (c Lrem) {
	c.cs = append(b.get(), "LREM")
	return
}

type LremCount Completed

func (c LremCount) Element(Element string) LremElement {
	return LremElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type LremElement Completed

func (c LremElement) Build() Completed {
	return Completed(c)
}

type LremKey Completed

func (c LremKey) Count(Count int64) LremCount {
	return LremCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

type Lset Completed

func (c Lset) Key(Key string) LsetKey {
	return LsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Lset() (c Lset) {
	c.cs = append(b.get(), "LSET")
	return
}

type LsetElement Completed

func (c LsetElement) Build() Completed {
	return Completed(c)
}

type LsetIndex Completed

func (c LsetIndex) Element(Element string) LsetElement {
	return LsetElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type LsetKey Completed

func (c LsetKey) Index(Index int64) LsetIndex {
	return LsetIndex{cs: append(c.cs, strconv.FormatInt(Index, 10)), cf: c.cf, ks: c.ks}
}

type Ltrim Completed

func (c Ltrim) Key(Key string) LtrimKey {
	return LtrimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Ltrim() (c Ltrim) {
	c.cs = append(b.get(), "LTRIM")
	return
}

type LtrimKey Completed

func (c LtrimKey) Start(Start int64) LtrimStart {
	return LtrimStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type LtrimStart Completed

func (c LtrimStart) Stop(Stop int64) LtrimStop {
	return LtrimStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type LtrimStop Completed

func (c LtrimStop) Build() Completed {
	return Completed(c)
}

type MemoryDoctor Completed

func (c MemoryDoctor) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryDoctor() (c MemoryDoctor) {
	c.cs = append(b.get(), "MEMORY", "DOCTOR")
	return
}

type MemoryHelp Completed

func (c MemoryHelp) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryHelp() (c MemoryHelp) {
	c.cs = append(b.get(), "MEMORY", "HELP")
	return
}

type MemoryMallocStats Completed

func (c MemoryMallocStats) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryMallocStats() (c MemoryMallocStats) {
	c.cs = append(b.get(), "MEMORY", "MALLOC-STATS")
	return
}

type MemoryPurge Completed

func (c MemoryPurge) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryPurge() (c MemoryPurge) {
	c.cs = append(b.get(), "MEMORY", "PURGE")
	return
}

type MemoryStats Completed

func (c MemoryStats) Build() Completed {
	return Completed(c)
}

func (b *Builder) MemoryStats() (c MemoryStats) {
	c.cs = append(b.get(), "MEMORY", "STATS")
	return
}

type MemoryUsage Completed

func (c MemoryUsage) Key(Key string) MemoryUsageKey {
	return MemoryUsageKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) MemoryUsage() (c MemoryUsage) {
	c.cs = append(b.get(), "MEMORY", "USAGE")
	return
}

type MemoryUsageKey Completed

func (c MemoryUsageKey) Samples(Count int64) MemoryUsageSamples {
	return MemoryUsageSamples{cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c MemoryUsageKey) Build() Completed {
	return Completed(c)
}

type MemoryUsageSamples Completed

func (c MemoryUsageSamples) Build() Completed {
	return Completed(c)
}

type Mget Completed

func (c Mget) Key(Key ...string) MgetKey {
	return MgetKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Mget() (c Mget) {
	c.cs = append(b.get(), "MGET")
	c.cf = readonly
	return
}

type MgetKey Completed

func (c MgetKey) Key(Key ...string) MgetKey {
	return MgetKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c MgetKey) Build() Completed {
	return Completed(c)
}

type Migrate Completed

func (c Migrate) Host(Host string) MigrateHost {
	return MigrateHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

func (b *Builder) Migrate() (c Migrate) {
	c.cs = append(b.get(), "MIGRATE")
	c.cf = blockTag
	return
}

type MigrateAuth Completed

func (c MigrateAuth) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c MigrateAuth) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c MigrateAuth) Build() Completed {
	return Completed(c)
}

type MigrateAuth2 Completed

func (c MigrateAuth2) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c MigrateAuth2) Build() Completed {
	return Completed(c)
}

type MigrateCopyCopy Completed

func (c MigrateCopyCopy) Replace() MigrateReplaceReplace {
	return MigrateReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c MigrateCopyCopy) Auth(Password string) MigrateAuth {
	return MigrateAuth{cs: append(c.cs, "AUTH", Password), cf: c.cf, ks: c.ks}
}

func (c MigrateCopyCopy) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c MigrateCopyCopy) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c MigrateCopyCopy) Build() Completed {
	return Completed(c)
}

type MigrateDestinationDb Completed

func (c MigrateDestinationDb) Timeout(Timeout int64) MigrateTimeout {
	return MigrateTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10)), cf: c.cf, ks: c.ks}
}

type MigrateHost Completed

func (c MigrateHost) Port(Port string) MigratePort {
	return MigratePort{cs: append(c.cs, Port), cf: c.cf, ks: c.ks}
}

type MigrateKeyEmpty Completed

func (c MigrateKeyEmpty) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cs: append(c.cs, strconv.FormatInt(DestinationDb, 10)), cf: c.cf, ks: c.ks}
}

type MigrateKeyKey Completed

func (c MigrateKeyKey) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cs: append(c.cs, strconv.FormatInt(DestinationDb, 10)), cf: c.cf, ks: c.ks}
}

type MigrateKeys Completed

func (c MigrateKeys) Keys(Keys ...string) MigrateKeys {
	return MigrateKeys{cs: append(c.cs, Keys...), cf: c.cf, ks: c.ks}
}

func (c MigrateKeys) Build() Completed {
	return Completed(c)
}

type MigratePort Completed

func (c MigratePort) Key() MigrateKeyKey {
	return MigrateKeyKey{cs: append(c.cs, "key"), cf: c.cf, ks: c.ks}
}

func (c MigratePort) Empty() MigrateKeyEmpty {
	return MigrateKeyEmpty{cs: append(c.cs, "\"\""), cf: c.cf, ks: c.ks}
}

type MigrateReplaceReplace Completed

func (c MigrateReplaceReplace) Auth(Password string) MigrateAuth {
	return MigrateAuth{cs: append(c.cs, "AUTH", Password), cf: c.cf, ks: c.ks}
}

func (c MigrateReplaceReplace) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c MigrateReplaceReplace) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c MigrateReplaceReplace) Build() Completed {
	return Completed(c)
}

type MigrateTimeout Completed

func (c MigrateTimeout) Copy() MigrateCopyCopy {
	return MigrateCopyCopy{cs: append(c.cs, "COPY"), cf: c.cf, ks: c.ks}
}

func (c MigrateTimeout) Replace() MigrateReplaceReplace {
	return MigrateReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c MigrateTimeout) Auth(Password string) MigrateAuth {
	return MigrateAuth{cs: append(c.cs, "AUTH", Password), cf: c.cf, ks: c.ks}
}

func (c MigrateTimeout) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c MigrateTimeout) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c MigrateTimeout) Build() Completed {
	return Completed(c)
}

type ModuleList Completed

func (c ModuleList) Build() Completed {
	return Completed(c)
}

func (b *Builder) ModuleList() (c ModuleList) {
	c.cs = append(b.get(), "MODULE", "LIST")
	return
}

type ModuleLoad Completed

func (c ModuleLoad) Path(Path string) ModuleLoadPath {
	return ModuleLoadPath{cs: append(c.cs, Path), cf: c.cf, ks: c.ks}
}

func (b *Builder) ModuleLoad() (c ModuleLoad) {
	c.cs = append(b.get(), "MODULE", "LOAD")
	return
}

type ModuleLoadArg Completed

func (c ModuleLoadArg) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c ModuleLoadArg) Build() Completed {
	return Completed(c)
}

type ModuleLoadPath Completed

func (c ModuleLoadPath) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c ModuleLoadPath) Build() Completed {
	return Completed(c)
}

type ModuleUnload Completed

func (c ModuleUnload) Name(Name string) ModuleUnloadName {
	return ModuleUnloadName{cs: append(c.cs, Name), cf: c.cf, ks: c.ks}
}

func (b *Builder) ModuleUnload() (c ModuleUnload) {
	c.cs = append(b.get(), "MODULE", "UNLOAD")
	return
}

type ModuleUnloadName Completed

func (c ModuleUnloadName) Build() Completed {
	return Completed(c)
}

type Monitor Completed

func (c Monitor) Build() Completed {
	return Completed(c)
}

func (b *Builder) Monitor() (c Monitor) {
	c.cs = append(b.get(), "MONITOR")
	return
}

type Move Completed

func (c Move) Key(Key string) MoveKey {
	return MoveKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Move() (c Move) {
	c.cs = append(b.get(), "MOVE")
	return
}

type MoveDb Completed

func (c MoveDb) Build() Completed {
	return Completed(c)
}

type MoveKey Completed

func (c MoveKey) Db(Db int64) MoveDb {
	return MoveDb{cs: append(c.cs, strconv.FormatInt(Db, 10)), cf: c.cf, ks: c.ks}
}

type Mset Completed

func (c Mset) KeyValue() MsetKeyValue {
	return MsetKeyValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *Builder) Mset() (c Mset) {
	c.cs = append(b.get(), "MSET")
	return
}

type MsetKeyValue Completed

func (c MsetKeyValue) KeyValue(Key string, Value string) MsetKeyValue {
	return MsetKeyValue{cs: append(c.cs, Key, Value), cf: c.cf, ks: c.ks}
}

func (c MsetKeyValue) Build() Completed {
	return Completed(c)
}

type Msetnx Completed

func (c Msetnx) KeyValue() MsetnxKeyValue {
	return MsetnxKeyValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *Builder) Msetnx() (c Msetnx) {
	c.cs = append(b.get(), "MSETNX")
	return
}

type MsetnxKeyValue Completed

func (c MsetnxKeyValue) KeyValue(Key string, Value string) MsetnxKeyValue {
	return MsetnxKeyValue{cs: append(c.cs, Key, Value), cf: c.cf, ks: c.ks}
}

func (c MsetnxKeyValue) Build() Completed {
	return Completed(c)
}

type Multi Completed

func (c Multi) Build() Completed {
	return Completed(c)
}

func (b *Builder) Multi() (c Multi) {
	c.cs = append(b.get(), "MULTI")
	return
}

type Object Completed

func (c Object) Subcommand(Subcommand string) ObjectSubcommand {
	return ObjectSubcommand{cs: append(c.cs, Subcommand), cf: c.cf, ks: c.ks}
}

func (b *Builder) Object() (c Object) {
	c.cs = append(b.get(), "OBJECT")
	c.cf = readonly
	return
}

type ObjectArguments Completed

func (c ObjectArguments) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cs: append(c.cs, Arguments...), cf: c.cf, ks: c.ks}
}

func (c ObjectArguments) Build() Completed {
	return Completed(c)
}

type ObjectSubcommand Completed

func (c ObjectSubcommand) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cs: append(c.cs, Arguments...), cf: c.cf, ks: c.ks}
}

func (c ObjectSubcommand) Build() Completed {
	return Completed(c)
}

type Persist Completed

func (c Persist) Key(Key string) PersistKey {
	return PersistKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Persist() (c Persist) {
	c.cs = append(b.get(), "PERSIST")
	return
}

type PersistKey Completed

func (c PersistKey) Build() Completed {
	return Completed(c)
}

type Pexpire Completed

func (c Pexpire) Key(Key string) PexpireKey {
	return PexpireKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pexpire() (c Pexpire) {
	c.cs = append(b.get(), "PEXPIRE")
	return
}

type PexpireConditionGt Completed

func (c PexpireConditionGt) Build() Completed {
	return Completed(c)
}

type PexpireConditionLt Completed

func (c PexpireConditionLt) Build() Completed {
	return Completed(c)
}

type PexpireConditionNx Completed

func (c PexpireConditionNx) Build() Completed {
	return Completed(c)
}

type PexpireConditionXx Completed

func (c PexpireConditionXx) Build() Completed {
	return Completed(c)
}

type PexpireKey Completed

func (c PexpireKey) Milliseconds(Milliseconds int64) PexpireMilliseconds {
	return PexpireMilliseconds{cs: append(c.cs, strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

type PexpireMilliseconds Completed

func (c PexpireMilliseconds) Nx() PexpireConditionNx {
	return PexpireConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c PexpireMilliseconds) Xx() PexpireConditionXx {
	return PexpireConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c PexpireMilliseconds) Gt() PexpireConditionGt {
	return PexpireConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c PexpireMilliseconds) Lt() PexpireConditionLt {
	return PexpireConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c PexpireMilliseconds) Build() Completed {
	return Completed(c)
}

type Pexpireat Completed

func (c Pexpireat) Key(Key string) PexpireatKey {
	return PexpireatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pexpireat() (c Pexpireat) {
	c.cs = append(b.get(), "PEXPIREAT")
	return
}

type PexpireatConditionGt Completed

func (c PexpireatConditionGt) Build() Completed {
	return Completed(c)
}

type PexpireatConditionLt Completed

func (c PexpireatConditionLt) Build() Completed {
	return Completed(c)
}

type PexpireatConditionNx Completed

func (c PexpireatConditionNx) Build() Completed {
	return Completed(c)
}

type PexpireatConditionXx Completed

func (c PexpireatConditionXx) Build() Completed {
	return Completed(c)
}

type PexpireatKey Completed

func (c PexpireatKey) MillisecondsTimestamp(MillisecondsTimestamp int64) PexpireatMillisecondsTimestamp {
	return PexpireatMillisecondsTimestamp{cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10)), cf: c.cf, ks: c.ks}
}

type PexpireatMillisecondsTimestamp Completed

func (c PexpireatMillisecondsTimestamp) Nx() PexpireatConditionNx {
	return PexpireatConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c PexpireatMillisecondsTimestamp) Xx() PexpireatConditionXx {
	return PexpireatConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c PexpireatMillisecondsTimestamp) Gt() PexpireatConditionGt {
	return PexpireatConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c PexpireatMillisecondsTimestamp) Lt() PexpireatConditionLt {
	return PexpireatConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c PexpireatMillisecondsTimestamp) Build() Completed {
	return Completed(c)
}

type Pexpiretime Completed

func (c Pexpiretime) Key(Key string) PexpiretimeKey {
	return PexpiretimeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pexpiretime() (c Pexpiretime) {
	c.cs = append(b.get(), "PEXPIRETIME")
	c.cf = readonly
	return
}

type PexpiretimeKey Completed

func (c PexpiretimeKey) Build() Completed {
	return Completed(c)
}

func (c PexpiretimeKey) Cache() Cacheable {
	return Cacheable(c)
}

type Pfadd Completed

func (c Pfadd) Key(Key string) PfaddKey {
	return PfaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pfadd() (c Pfadd) {
	c.cs = append(b.get(), "PFADD")
	return
}

type PfaddElement Completed

func (c PfaddElement) Element(Element ...string) PfaddElement {
	return PfaddElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c PfaddElement) Build() Completed {
	return Completed(c)
}

type PfaddKey Completed

func (c PfaddKey) Element(Element ...string) PfaddElement {
	return PfaddElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c PfaddKey) Build() Completed {
	return Completed(c)
}

type Pfcount Completed

func (c Pfcount) Key(Key ...string) PfcountKey {
	return PfcountKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pfcount() (c Pfcount) {
	c.cs = append(b.get(), "PFCOUNT")
	c.cf = readonly
	return
}

type PfcountKey Completed

func (c PfcountKey) Key(Key ...string) PfcountKey {
	return PfcountKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c PfcountKey) Build() Completed {
	return Completed(c)
}

type Pfmerge Completed

func (c Pfmerge) Destkey(Destkey string) PfmergeDestkey {
	return PfmergeDestkey{cs: append(c.cs, Destkey), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pfmerge() (c Pfmerge) {
	c.cs = append(b.get(), "PFMERGE")
	return
}

type PfmergeDestkey Completed

func (c PfmergeDestkey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cs: append(c.cs, Sourcekey...), cf: c.cf, ks: c.ks}
}

type PfmergeSourcekey Completed

func (c PfmergeSourcekey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cs: append(c.cs, Sourcekey...), cf: c.cf, ks: c.ks}
}

func (c PfmergeSourcekey) Build() Completed {
	return Completed(c)
}

type Ping Completed

func (c Ping) Message(Message string) PingMessage {
	return PingMessage{cs: append(c.cs, Message), cf: c.cf, ks: c.ks}
}

func (c Ping) Build() Completed {
	return Completed(c)
}

func (b *Builder) Ping() (c Ping) {
	c.cs = append(b.get(), "PING")
	return
}

type PingMessage Completed

func (c PingMessage) Build() Completed {
	return Completed(c)
}

type Psetex Completed

func (c Psetex) Key(Key string) PsetexKey {
	return PsetexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Psetex() (c Psetex) {
	c.cs = append(b.get(), "PSETEX")
	return
}

type PsetexKey Completed

func (c PsetexKey) Milliseconds(Milliseconds int64) PsetexMilliseconds {
	return PsetexMilliseconds{cs: append(c.cs, strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

type PsetexMilliseconds Completed

func (c PsetexMilliseconds) Value(Value string) PsetexValue {
	return PsetexValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type PsetexValue Completed

func (c PsetexValue) Build() Completed {
	return Completed(c)
}

type Psubscribe Completed

func (c Psubscribe) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Psubscribe() (c Psubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	c.cf = noRetTag
	return
}

type PsubscribePattern Completed

func (c PsubscribePattern) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c PsubscribePattern) Build() Completed {
	return Completed(c)
}

type Psync Completed

func (c Psync) Replicationid(Replicationid int64) PsyncReplicationid {
	return PsyncReplicationid{cs: append(c.cs, strconv.FormatInt(Replicationid, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Psync() (c Psync) {
	c.cs = append(b.get(), "PSYNC")
	return
}

type PsyncOffset Completed

func (c PsyncOffset) Build() Completed {
	return Completed(c)
}

type PsyncReplicationid Completed

func (c PsyncReplicationid) Offset(Offset int64) PsyncOffset {
	return PsyncOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type Pttl Completed

func (c Pttl) Key(Key string) PttlKey {
	return PttlKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pttl() (c Pttl) {
	c.cs = append(b.get(), "PTTL")
	c.cf = readonly
	return
}

type PttlKey Completed

func (c PttlKey) Build() Completed {
	return Completed(c)
}

func (c PttlKey) Cache() Cacheable {
	return Cacheable(c)
}

type Publish Completed

func (c Publish) Channel(Channel string) PublishChannel {
	return PublishChannel{cs: append(c.cs, Channel), cf: c.cf, ks: c.ks}
}

func (b *Builder) Publish() (c Publish) {
	c.cs = append(b.get(), "PUBLISH")
	return
}

type PublishChannel Completed

func (c PublishChannel) Message(Message string) PublishMessage {
	return PublishMessage{cs: append(c.cs, Message), cf: c.cf, ks: c.ks}
}

type PublishMessage Completed

func (c PublishMessage) Build() Completed {
	return Completed(c)
}

type Pubsub Completed

func (c Pubsub) Subcommand(Subcommand string) PubsubSubcommand {
	return PubsubSubcommand{cs: append(c.cs, Subcommand), cf: c.cf, ks: c.ks}
}

func (b *Builder) Pubsub() (c Pubsub) {
	c.cs = append(b.get(), "PUBSUB")
	return
}

type PubsubArgument Completed

func (c PubsubArgument) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cs: append(c.cs, Argument...), cf: c.cf, ks: c.ks}
}

func (c PubsubArgument) Build() Completed {
	return Completed(c)
}

type PubsubSubcommand Completed

func (c PubsubSubcommand) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cs: append(c.cs, Argument...), cf: c.cf, ks: c.ks}
}

func (c PubsubSubcommand) Build() Completed {
	return Completed(c)
}

type Punsubscribe Completed

func (c Punsubscribe) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c Punsubscribe) Build() Completed {
	return Completed(c)
}

func (b *Builder) Punsubscribe() (c Punsubscribe) {
	c.cs = append(b.get(), "PUNSUBSCRIBE")
	c.cf = noRetTag
	return
}

type PunsubscribePattern Completed

func (c PunsubscribePattern) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c PunsubscribePattern) Build() Completed {
	return Completed(c)
}

type Quit Completed

func (c Quit) Build() Completed {
	return Completed(c)
}

func (b *Builder) Quit() (c Quit) {
	c.cs = append(b.get(), "QUIT")
	return
}

type Randomkey Completed

func (c Randomkey) Build() Completed {
	return Completed(c)
}

func (b *Builder) Randomkey() (c Randomkey) {
	c.cs = append(b.get(), "RANDOMKEY")
	c.cf = readonly
	return
}

type Readonly Completed

func (c Readonly) Build() Completed {
	return Completed(c)
}

func (b *Builder) Readonly() (c Readonly) {
	c.cs = append(b.get(), "READONLY")
	return
}

type Readwrite Completed

func (c Readwrite) Build() Completed {
	return Completed(c)
}

func (b *Builder) Readwrite() (c Readwrite) {
	c.cs = append(b.get(), "READWRITE")
	return
}

type Rename Completed

func (c Rename) Key(Key string) RenameKey {
	return RenameKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Rename() (c Rename) {
	c.cs = append(b.get(), "RENAME")
	return
}

type RenameKey Completed

func (c RenameKey) Newkey(Newkey string) RenameNewkey {
	return RenameNewkey{cs: append(c.cs, Newkey), cf: c.cf, ks: c.ks}
}

type RenameNewkey Completed

func (c RenameNewkey) Build() Completed {
	return Completed(c)
}

type Renamenx Completed

func (c Renamenx) Key(Key string) RenamenxKey {
	return RenamenxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Renamenx() (c Renamenx) {
	c.cs = append(b.get(), "RENAMENX")
	return
}

type RenamenxKey Completed

func (c RenamenxKey) Newkey(Newkey string) RenamenxNewkey {
	return RenamenxNewkey{cs: append(c.cs, Newkey), cf: c.cf, ks: c.ks}
}

type RenamenxNewkey Completed

func (c RenamenxNewkey) Build() Completed {
	return Completed(c)
}

type Replicaof Completed

func (c Replicaof) Host(Host string) ReplicaofHost {
	return ReplicaofHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

func (b *Builder) Replicaof() (c Replicaof) {
	c.cs = append(b.get(), "REPLICAOF")
	return
}

type ReplicaofHost Completed

func (c ReplicaofHost) Port(Port string) ReplicaofPort {
	return ReplicaofPort{cs: append(c.cs, Port), cf: c.cf, ks: c.ks}
}

type ReplicaofPort Completed

func (c ReplicaofPort) Build() Completed {
	return Completed(c)
}

type Reset Completed

func (c Reset) Build() Completed {
	return Completed(c)
}

func (b *Builder) Reset() (c Reset) {
	c.cs = append(b.get(), "RESET")
	return
}

type Restore Completed

func (c Restore) Key(Key string) RestoreKey {
	return RestoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Restore() (c Restore) {
	c.cs = append(b.get(), "RESTORE")
	return
}

type RestoreAbsttlAbsttl Completed

func (c RestoreAbsttlAbsttl) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c RestoreAbsttlAbsttl) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c RestoreAbsttlAbsttl) Build() Completed {
	return Completed(c)
}

type RestoreFreq Completed

func (c RestoreFreq) Build() Completed {
	return Completed(c)
}

type RestoreIdletime Completed

func (c RestoreIdletime) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c RestoreIdletime) Build() Completed {
	return Completed(c)
}

type RestoreKey Completed

func (c RestoreKey) Ttl(Ttl int64) RestoreTtl {
	return RestoreTtl{cs: append(c.cs, strconv.FormatInt(Ttl, 10)), cf: c.cf, ks: c.ks}
}

type RestoreReplaceReplace Completed

func (c RestoreReplaceReplace) Absttl() RestoreAbsttlAbsttl {
	return RestoreAbsttlAbsttl{cs: append(c.cs, "ABSTTL"), cf: c.cf, ks: c.ks}
}

func (c RestoreReplaceReplace) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c RestoreReplaceReplace) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c RestoreReplaceReplace) Build() Completed {
	return Completed(c)
}

type RestoreSerializedValue Completed

func (c RestoreSerializedValue) Replace() RestoreReplaceReplace {
	return RestoreReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c RestoreSerializedValue) Absttl() RestoreAbsttlAbsttl {
	return RestoreAbsttlAbsttl{cs: append(c.cs, "ABSTTL"), cf: c.cf, ks: c.ks}
}

func (c RestoreSerializedValue) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c RestoreSerializedValue) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c RestoreSerializedValue) Build() Completed {
	return Completed(c)
}

type RestoreTtl Completed

func (c RestoreTtl) SerializedValue(SerializedValue string) RestoreSerializedValue {
	return RestoreSerializedValue{cs: append(c.cs, SerializedValue), cf: c.cf, ks: c.ks}
}

type Role Completed

func (c Role) Build() Completed {
	return Completed(c)
}

func (b *Builder) Role() (c Role) {
	c.cs = append(b.get(), "ROLE")
	return
}

type Rpop Completed

func (c Rpop) Key(Key string) RpopKey {
	return RpopKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Rpop() (c Rpop) {
	c.cs = append(b.get(), "RPOP")
	return
}

type RpopCount Completed

func (c RpopCount) Build() Completed {
	return Completed(c)
}

type RpopKey Completed

func (c RpopKey) Count(Count int64) RpopCount {
	return RpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c RpopKey) Build() Completed {
	return Completed(c)
}

type Rpoplpush Completed

func (c Rpoplpush) Source(Source string) RpoplpushSource {
	return RpoplpushSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *Builder) Rpoplpush() (c Rpoplpush) {
	c.cs = append(b.get(), "RPOPLPUSH")
	return
}

type RpoplpushDestination Completed

func (c RpoplpushDestination) Build() Completed {
	return Completed(c)
}

type RpoplpushSource Completed

func (c RpoplpushSource) Destination(Destination string) RpoplpushDestination {
	return RpoplpushDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type Rpush Completed

func (c Rpush) Key(Key string) RpushKey {
	return RpushKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Rpush() (c Rpush) {
	c.cs = append(b.get(), "RPUSH")
	return
}

type RpushElement Completed

func (c RpushElement) Element(Element ...string) RpushElement {
	return RpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c RpushElement) Build() Completed {
	return Completed(c)
}

type RpushKey Completed

func (c RpushKey) Element(Element ...string) RpushElement {
	return RpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type Rpushx Completed

func (c Rpushx) Key(Key string) RpushxKey {
	return RpushxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Rpushx() (c Rpushx) {
	c.cs = append(b.get(), "RPUSHX")
	return
}

type RpushxElement Completed

func (c RpushxElement) Element(Element ...string) RpushxElement {
	return RpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c RpushxElement) Build() Completed {
	return Completed(c)
}

type RpushxKey Completed

func (c RpushxKey) Element(Element ...string) RpushxElement {
	return RpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type Sadd Completed

func (c Sadd) Key(Key string) SaddKey {
	return SaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sadd() (c Sadd) {
	c.cs = append(b.get(), "SADD")
	return
}

type SaddKey Completed

func (c SaddKey) Member(Member ...string) SaddMember {
	return SaddMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SaddMember Completed

func (c SaddMember) Member(Member ...string) SaddMember {
	return SaddMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SaddMember) Build() Completed {
	return Completed(c)
}

type Save Completed

func (c Save) Build() Completed {
	return Completed(c)
}

func (b *Builder) Save() (c Save) {
	c.cs = append(b.get(), "SAVE")
	return
}

type Scan Completed

func (c Scan) Cursor(Cursor int64) ScanCursor {
	return ScanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Scan() (c Scan) {
	c.cs = append(b.get(), "SCAN")
	c.cf = readonly
	return
}

type ScanCount Completed

func (c ScanCount) Type(Type string) ScanType {
	return ScanType{cs: append(c.cs, "TYPE", Type), cf: c.cf, ks: c.ks}
}

func (c ScanCount) Build() Completed {
	return Completed(c)
}

type ScanCursor Completed

func (c ScanCursor) Match(Pattern string) ScanMatch {
	return ScanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c ScanCursor) Count(Count int64) ScanCount {
	return ScanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ScanCursor) Type(Type string) ScanType {
	return ScanType{cs: append(c.cs, "TYPE", Type), cf: c.cf, ks: c.ks}
}

func (c ScanCursor) Build() Completed {
	return Completed(c)
}

type ScanMatch Completed

func (c ScanMatch) Count(Count int64) ScanCount {
	return ScanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ScanMatch) Type(Type string) ScanType {
	return ScanType{cs: append(c.cs, "TYPE", Type), cf: c.cf, ks: c.ks}
}

func (c ScanMatch) Build() Completed {
	return Completed(c)
}

type ScanType Completed

func (c ScanType) Build() Completed {
	return Completed(c)
}

type Scard Completed

func (c Scard) Key(Key string) ScardKey {
	return ScardKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Scard() (c Scard) {
	c.cs = append(b.get(), "SCARD")
	c.cf = readonly
	return
}

type ScardKey Completed

func (c ScardKey) Build() Completed {
	return Completed(c)
}

func (c ScardKey) Cache() Cacheable {
	return Cacheable(c)
}

type ScriptDebug Completed

func (c ScriptDebug) Yes() ScriptDebugModeYes {
	return ScriptDebugModeYes{cs: append(c.cs, "YES"), cf: c.cf, ks: c.ks}
}

func (c ScriptDebug) Sync() ScriptDebugModeSync {
	return ScriptDebugModeSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c ScriptDebug) No() ScriptDebugModeNo {
	return ScriptDebugModeNo{cs: append(c.cs, "NO"), cf: c.cf, ks: c.ks}
}

func (b *Builder) ScriptDebug() (c ScriptDebug) {
	c.cs = append(b.get(), "SCRIPT", "DEBUG")
	return
}

type ScriptDebugModeNo Completed

func (c ScriptDebugModeNo) Build() Completed {
	return Completed(c)
}

type ScriptDebugModeSync Completed

func (c ScriptDebugModeSync) Build() Completed {
	return Completed(c)
}

type ScriptDebugModeYes Completed

func (c ScriptDebugModeYes) Build() Completed {
	return Completed(c)
}

type ScriptExists Completed

func (c ScriptExists) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cs: append(c.cs, Sha1...), cf: c.cf, ks: c.ks}
}

func (b *Builder) ScriptExists() (c ScriptExists) {
	c.cs = append(b.get(), "SCRIPT", "EXISTS")
	return
}

type ScriptExistsSha1 Completed

func (c ScriptExistsSha1) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cs: append(c.cs, Sha1...), cf: c.cf, ks: c.ks}
}

func (c ScriptExistsSha1) Build() Completed {
	return Completed(c)
}

type ScriptFlush Completed

func (c ScriptFlush) Async() ScriptFlushAsyncAsync {
	return ScriptFlushAsyncAsync{cs: append(c.cs, "ASYNC"), cf: c.cf, ks: c.ks}
}

func (c ScriptFlush) Sync() ScriptFlushAsyncSync {
	return ScriptFlushAsyncSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c ScriptFlush) Build() Completed {
	return Completed(c)
}

func (b *Builder) ScriptFlush() (c ScriptFlush) {
	c.cs = append(b.get(), "SCRIPT", "FLUSH")
	return
}

type ScriptFlushAsyncAsync Completed

func (c ScriptFlushAsyncAsync) Build() Completed {
	return Completed(c)
}

type ScriptFlushAsyncSync Completed

func (c ScriptFlushAsyncSync) Build() Completed {
	return Completed(c)
}

type ScriptKill Completed

func (c ScriptKill) Build() Completed {
	return Completed(c)
}

func (b *Builder) ScriptKill() (c ScriptKill) {
	c.cs = append(b.get(), "SCRIPT", "KILL")
	return
}

type ScriptLoad Completed

func (c ScriptLoad) Script(Script string) ScriptLoadScript {
	return ScriptLoadScript{cs: append(c.cs, Script), cf: c.cf, ks: c.ks}
}

func (b *Builder) ScriptLoad() (c ScriptLoad) {
	c.cs = append(b.get(), "SCRIPT", "LOAD")
	return
}

type ScriptLoadScript Completed

func (c ScriptLoadScript) Build() Completed {
	return Completed(c)
}

type Sdiff Completed

func (c Sdiff) Key(Key ...string) SdiffKey {
	return SdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sdiff() (c Sdiff) {
	c.cs = append(b.get(), "SDIFF")
	c.cf = readonly
	return
}

type SdiffKey Completed

func (c SdiffKey) Key(Key ...string) SdiffKey {
	return SdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SdiffKey) Build() Completed {
	return Completed(c)
}

type Sdiffstore Completed

func (c Sdiffstore) Destination(Destination string) SdiffstoreDestination {
	return SdiffstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sdiffstore() (c Sdiffstore) {
	c.cs = append(b.get(), "SDIFFSTORE")
	return
}

type SdiffstoreDestination Completed

func (c SdiffstoreDestination) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SdiffstoreKey Completed

func (c SdiffstoreKey) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SdiffstoreKey) Build() Completed {
	return Completed(c)
}

type Select Completed

func (c Select) Index(Index int64) SelectIndex {
	return SelectIndex{cs: append(c.cs, strconv.FormatInt(Index, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Select() (c Select) {
	c.cs = append(b.get(), "SELECT")
	return
}

type SelectIndex Completed

func (c SelectIndex) Build() Completed {
	return Completed(c)
}

type Set Completed

func (c Set) Key(Key string) SetKey {
	return SetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Set() (c Set) {
	c.cs = append(b.get(), "SET")
	return
}

type SetConditionNx Completed

func (c SetConditionNx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetConditionNx) Build() Completed {
	return Completed(c)
}

type SetConditionXx Completed

func (c SetConditionXx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetConditionXx) Build() Completed {
	return Completed(c)
}

type SetExpirationEx Completed

func (c SetExpirationEx) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationEx) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationEx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationEx) Build() Completed {
	return Completed(c)
}

type SetExpirationExat Completed

func (c SetExpirationExat) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationExat) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationExat) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationExat) Build() Completed {
	return Completed(c)
}

type SetExpirationKeepttl Completed

func (c SetExpirationKeepttl) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationKeepttl) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationKeepttl) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationKeepttl) Build() Completed {
	return Completed(c)
}

type SetExpirationPx Completed

func (c SetExpirationPx) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationPx) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationPx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationPx) Build() Completed {
	return Completed(c)
}

type SetExpirationPxat Completed

func (c SetExpirationPxat) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationPxat) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationPxat) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetExpirationPxat) Build() Completed {
	return Completed(c)
}

type SetGetGet Completed

func (c SetGetGet) Build() Completed {
	return Completed(c)
}

type SetKey Completed

func (c SetKey) Value(Value string) SetValue {
	return SetValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SetValue Completed

func (c SetValue) Ex(Seconds int64) SetExpirationEx {
	return SetExpirationEx{cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SetValue) Px(Milliseconds int64) SetExpirationPx {
	return SetExpirationPx{cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SetValue) Exat(Timestamp int64) SetExpirationExat {
	return SetExpirationExat{cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c SetValue) Pxat(Millisecondstimestamp int64) SetExpirationPxat {
	return SetExpirationPxat{cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c SetValue) Keepttl() SetExpirationKeepttl {
	return SetExpirationKeepttl{cs: append(c.cs, "KEEPTTL"), cf: c.cf, ks: c.ks}
}

func (c SetValue) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SetValue) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SetValue) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SetValue) Build() Completed {
	return Completed(c)
}

type Setbit Completed

func (c Setbit) Key(Key string) SetbitKey {
	return SetbitKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Setbit() (c Setbit) {
	c.cs = append(b.get(), "SETBIT")
	return
}

type SetbitKey Completed

func (c SetbitKey) Offset(Offset int64) SetbitOffset {
	return SetbitOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type SetbitOffset Completed

func (c SetbitOffset) Value(Value int64) SetbitValue {
	return SetbitValue{cs: append(c.cs, strconv.FormatInt(Value, 10)), cf: c.cf, ks: c.ks}
}

type SetbitValue Completed

func (c SetbitValue) Build() Completed {
	return Completed(c)
}

type Setex Completed

func (c Setex) Key(Key string) SetexKey {
	return SetexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Setex() (c Setex) {
	c.cs = append(b.get(), "SETEX")
	return
}

type SetexKey Completed

func (c SetexKey) Seconds(Seconds int64) SetexSeconds {
	return SetexSeconds{cs: append(c.cs, strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

type SetexSeconds Completed

func (c SetexSeconds) Value(Value string) SetexValue {
	return SetexValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SetexValue Completed

func (c SetexValue) Build() Completed {
	return Completed(c)
}

type Setnx Completed

func (c Setnx) Key(Key string) SetnxKey {
	return SetnxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Setnx() (c Setnx) {
	c.cs = append(b.get(), "SETNX")
	return
}

type SetnxKey Completed

func (c SetnxKey) Value(Value string) SetnxValue {
	return SetnxValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SetnxValue Completed

func (c SetnxValue) Build() Completed {
	return Completed(c)
}

type Setrange Completed

func (c Setrange) Key(Key string) SetrangeKey {
	return SetrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Setrange() (c Setrange) {
	c.cs = append(b.get(), "SETRANGE")
	return
}

type SetrangeKey Completed

func (c SetrangeKey) Offset(Offset int64) SetrangeOffset {
	return SetrangeOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type SetrangeOffset Completed

func (c SetrangeOffset) Value(Value string) SetrangeValue {
	return SetrangeValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SetrangeValue Completed

func (c SetrangeValue) Build() Completed {
	return Completed(c)
}

type Shutdown Completed

func (c Shutdown) Nosave() ShutdownSaveModeNosave {
	return ShutdownSaveModeNosave{cs: append(c.cs, "NOSAVE"), cf: c.cf, ks: c.ks}
}

func (c Shutdown) Save() ShutdownSaveModeSave {
	return ShutdownSaveModeSave{cs: append(c.cs, "SAVE"), cf: c.cf, ks: c.ks}
}

func (c Shutdown) Build() Completed {
	return Completed(c)
}

func (b *Builder) Shutdown() (c Shutdown) {
	c.cs = append(b.get(), "SHUTDOWN")
	return
}

type ShutdownSaveModeNosave Completed

func (c ShutdownSaveModeNosave) Build() Completed {
	return Completed(c)
}

type ShutdownSaveModeSave Completed

func (c ShutdownSaveModeSave) Build() Completed {
	return Completed(c)
}

type Sinter Completed

func (c Sinter) Key(Key ...string) SinterKey {
	return SinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sinter() (c Sinter) {
	c.cs = append(b.get(), "SINTER")
	c.cf = readonly
	return
}

type SinterKey Completed

func (c SinterKey) Key(Key ...string) SinterKey {
	return SinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SinterKey) Build() Completed {
	return Completed(c)
}

type Sintercard Completed

func (c Sintercard) Key(Key ...string) SintercardKey {
	return SintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sintercard() (c Sintercard) {
	c.cs = append(b.get(), "SINTERCARD")
	c.cf = readonly
	return
}

type SintercardKey Completed

func (c SintercardKey) Key(Key ...string) SintercardKey {
	return SintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SintercardKey) Build() Completed {
	return Completed(c)
}

type Sinterstore Completed

func (c Sinterstore) Destination(Destination string) SinterstoreDestination {
	return SinterstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sinterstore() (c Sinterstore) {
	c.cs = append(b.get(), "SINTERSTORE")
	return
}

type SinterstoreDestination Completed

func (c SinterstoreDestination) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SinterstoreKey Completed

func (c SinterstoreKey) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SinterstoreKey) Build() Completed {
	return Completed(c)
}

type Sismember Completed

func (c Sismember) Key(Key string) SismemberKey {
	return SismemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sismember() (c Sismember) {
	c.cs = append(b.get(), "SISMEMBER")
	c.cf = readonly
	return
}

type SismemberKey Completed

func (c SismemberKey) Member(Member string) SismemberMember {
	return SismemberMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SismemberMember Completed

func (c SismemberMember) Build() Completed {
	return Completed(c)
}

func (c SismemberMember) Cache() Cacheable {
	return Cacheable(c)
}

type Slaveof Completed

func (c Slaveof) Host(Host string) SlaveofHost {
	return SlaveofHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

func (b *Builder) Slaveof() (c Slaveof) {
	c.cs = append(b.get(), "SLAVEOF")
	return
}

type SlaveofHost Completed

func (c SlaveofHost) Port(Port string) SlaveofPort {
	return SlaveofPort{cs: append(c.cs, Port), cf: c.cf, ks: c.ks}
}

type SlaveofPort Completed

func (c SlaveofPort) Build() Completed {
	return Completed(c)
}

type Slowlog Completed

func (c Slowlog) Subcommand(Subcommand string) SlowlogSubcommand {
	return SlowlogSubcommand{cs: append(c.cs, Subcommand), cf: c.cf, ks: c.ks}
}

func (b *Builder) Slowlog() (c Slowlog) {
	c.cs = append(b.get(), "SLOWLOG")
	return
}

type SlowlogArgument Completed

func (c SlowlogArgument) Build() Completed {
	return Completed(c)
}

type SlowlogSubcommand Completed

func (c SlowlogSubcommand) Argument(Argument string) SlowlogArgument {
	return SlowlogArgument{cs: append(c.cs, Argument), cf: c.cf, ks: c.ks}
}

func (c SlowlogSubcommand) Build() Completed {
	return Completed(c)
}

type Smembers Completed

func (c Smembers) Key(Key string) SmembersKey {
	return SmembersKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Smembers() (c Smembers) {
	c.cs = append(b.get(), "SMEMBERS")
	c.cf = readonly
	return
}

type SmembersKey Completed

func (c SmembersKey) Build() Completed {
	return Completed(c)
}

func (c SmembersKey) Cache() Cacheable {
	return Cacheable(c)
}

type Smismember Completed

func (c Smismember) Key(Key string) SmismemberKey {
	return SmismemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Smismember() (c Smismember) {
	c.cs = append(b.get(), "SMISMEMBER")
	c.cf = readonly
	return
}

type SmismemberKey Completed

func (c SmismemberKey) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SmismemberMember Completed

func (c SmismemberMember) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SmismemberMember) Build() Completed {
	return Completed(c)
}

func (c SmismemberMember) Cache() Cacheable {
	return Cacheable(c)
}

type Smove Completed

func (c Smove) Source(Source string) SmoveSource {
	return SmoveSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *Builder) Smove() (c Smove) {
	c.cs = append(b.get(), "SMOVE")
	return
}

type SmoveDestination Completed

func (c SmoveDestination) Member(Member string) SmoveMember {
	return SmoveMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SmoveMember Completed

func (c SmoveMember) Build() Completed {
	return Completed(c)
}

type SmoveSource Completed

func (c SmoveSource) Destination(Destination string) SmoveDestination {
	return SmoveDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type Sort Completed

func (c Sort) Key(Key string) SortKey {
	return SortKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sort() (c Sort) {
	c.cs = append(b.get(), "SORT")
	return
}

type SortBy Completed

func (c SortBy) Limit(Offset int64, Count int64) SortLimit {
	return SortLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SortBy) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SortBy) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortBy) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortBy) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortBy) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SortBy) Build() Completed {
	return Completed(c)
}

type SortGet Completed

func (c SortGet) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortGet) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortGet) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortGet) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SortGet) Get(Get ...string) SortGet {
	return SortGet{cs: append(c.cs, Get...), cf: c.cf, ks: c.ks}
}

func (c SortGet) Build() Completed {
	return Completed(c)
}

type SortKey Completed

func (c SortKey) By(Pattern string) SortBy {
	return SortBy{cs: append(c.cs, "BY", Pattern), cf: c.cf, ks: c.ks}
}

func (c SortKey) Limit(Offset int64, Count int64) SortLimit {
	return SortLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SortKey) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SortKey) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortKey) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortKey) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortKey) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SortKey) Build() Completed {
	return Completed(c)
}

type SortLimit Completed

func (c SortLimit) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SortLimit) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortLimit) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortLimit) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortLimit) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SortLimit) Build() Completed {
	return Completed(c)
}

type SortOrderAsc Completed

func (c SortOrderAsc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortOrderAsc) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SortOrderAsc) Build() Completed {
	return Completed(c)
}

type SortOrderDesc Completed

func (c SortOrderDesc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortOrderDesc) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SortOrderDesc) Build() Completed {
	return Completed(c)
}

type SortRo Completed

func (c SortRo) Key(Key string) SortRoKey {
	return SortRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) SortRo() (c SortRo) {
	c.cs = append(b.get(), "SORT_RO")
	c.cf = readonly
	return
}

type SortRoBy Completed

func (c SortRoBy) Limit(Offset int64, Count int64) SortRoLimit {
	return SortRoLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SortRoBy) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SortRoBy) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortRoBy) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortRoBy) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortRoBy) Build() Completed {
	return Completed(c)
}

func (c SortRoBy) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoGet Completed

func (c SortRoGet) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortRoGet) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortRoGet) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortRoGet) Get(Get ...string) SortRoGet {
	return SortRoGet{cs: append(c.cs, Get...), cf: c.cf, ks: c.ks}
}

func (c SortRoGet) Build() Completed {
	return Completed(c)
}

func (c SortRoGet) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoKey Completed

func (c SortRoKey) By(Pattern string) SortRoBy {
	return SortRoBy{cs: append(c.cs, "BY", Pattern), cf: c.cf, ks: c.ks}
}

func (c SortRoKey) Limit(Offset int64, Count int64) SortRoLimit {
	return SortRoLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SortRoKey) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SortRoKey) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortRoKey) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortRoKey) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortRoKey) Build() Completed {
	return Completed(c)
}

func (c SortRoKey) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoLimit Completed

func (c SortRoLimit) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SortRoLimit) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SortRoLimit) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SortRoLimit) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortRoLimit) Build() Completed {
	return Completed(c)
}

func (c SortRoLimit) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoOrderAsc Completed

func (c SortRoOrderAsc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c SortRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoOrderDesc Completed

func (c SortRoOrderDesc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SortRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c SortRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoSortingAlpha Completed

func (c SortRoSortingAlpha) Build() Completed {
	return Completed(c)
}

func (c SortRoSortingAlpha) Cache() Cacheable {
	return Cacheable(c)
}

type SortSortingAlpha Completed

func (c SortSortingAlpha) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SortSortingAlpha) Build() Completed {
	return Completed(c)
}

type SortStore Completed

func (c SortStore) Build() Completed {
	return Completed(c)
}

type Spop Completed

func (c Spop) Key(Key string) SpopKey {
	return SpopKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Spop() (c Spop) {
	c.cs = append(b.get(), "SPOP")
	return
}

type SpopCount Completed

func (c SpopCount) Build() Completed {
	return Completed(c)
}

type SpopKey Completed

func (c SpopKey) Count(Count int64) SpopCount {
	return SpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SpopKey) Build() Completed {
	return Completed(c)
}

type Srandmember Completed

func (c Srandmember) Key(Key string) SrandmemberKey {
	return SrandmemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Srandmember() (c Srandmember) {
	c.cs = append(b.get(), "SRANDMEMBER")
	c.cf = readonly
	return
}

type SrandmemberCount Completed

func (c SrandmemberCount) Build() Completed {
	return Completed(c)
}

type SrandmemberKey Completed

func (c SrandmemberKey) Count(Count int64) SrandmemberCount {
	return SrandmemberCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SrandmemberKey) Build() Completed {
	return Completed(c)
}

type Srem Completed

func (c Srem) Key(Key string) SremKey {
	return SremKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Srem() (c Srem) {
	c.cs = append(b.get(), "SREM")
	return
}

type SremKey Completed

func (c SremKey) Member(Member ...string) SremMember {
	return SremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SremMember Completed

func (c SremMember) Member(Member ...string) SremMember {
	return SremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SremMember) Build() Completed {
	return Completed(c)
}

type Sscan Completed

func (c Sscan) Key(Key string) SscanKey {
	return SscanKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sscan() (c Sscan) {
	c.cs = append(b.get(), "SSCAN")
	c.cf = readonly
	return
}

type SscanCount Completed

func (c SscanCount) Build() Completed {
	return Completed(c)
}

type SscanCursor Completed

func (c SscanCursor) Match(Pattern string) SscanMatch {
	return SscanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c SscanCursor) Count(Count int64) SscanCount {
	return SscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SscanCursor) Build() Completed {
	return Completed(c)
}

type SscanKey Completed

func (c SscanKey) Cursor(Cursor int64) SscanCursor {
	return SscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

type SscanMatch Completed

func (c SscanMatch) Count(Count int64) SscanCount {
	return SscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SscanMatch) Build() Completed {
	return Completed(c)
}

type Stralgo Completed

func (c Stralgo) Lcs() StralgoAlgorithmLcs {
	return StralgoAlgorithmLcs{cs: append(c.cs, "LCS"), cf: c.cf, ks: c.ks}
}

func (b *Builder) Stralgo() (c Stralgo) {
	c.cs = append(b.get(), "STRALGO")
	c.cf = readonly
	return
}

type StralgoAlgoSpecificArgument Completed

func (c StralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cs: append(c.cs, AlgoSpecificArgument...), cf: c.cf, ks: c.ks}
}

func (c StralgoAlgoSpecificArgument) Build() Completed {
	return Completed(c)
}

type StralgoAlgorithmLcs Completed

func (c StralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cs: append(c.cs, AlgoSpecificArgument...), cf: c.cf, ks: c.ks}
}

type Strlen Completed

func (c Strlen) Key(Key string) StrlenKey {
	return StrlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Strlen() (c Strlen) {
	c.cs = append(b.get(), "STRLEN")
	c.cf = readonly
	return
}

type StrlenKey Completed

func (c StrlenKey) Build() Completed {
	return Completed(c)
}

func (c StrlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Subscribe Completed

func (c Subscribe) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Subscribe() (c Subscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	c.cf = noRetTag
	return
}

type SubscribeChannel Completed

func (c SubscribeChannel) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (c SubscribeChannel) Build() Completed {
	return Completed(c)
}

type Sunion Completed

func (c Sunion) Key(Key ...string) SunionKey {
	return SunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sunion() (c Sunion) {
	c.cs = append(b.get(), "SUNION")
	c.cf = readonly
	return
}

type SunionKey Completed

func (c SunionKey) Key(Key ...string) SunionKey {
	return SunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SunionKey) Build() Completed {
	return Completed(c)
}

type Sunionstore Completed

func (c Sunionstore) Destination(Destination string) SunionstoreDestination {
	return SunionstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *Builder) Sunionstore() (c Sunionstore) {
	c.cs = append(b.get(), "SUNIONSTORE")
	return
}

type SunionstoreDestination Completed

func (c SunionstoreDestination) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SunionstoreKey Completed

func (c SunionstoreKey) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SunionstoreKey) Build() Completed {
	return Completed(c)
}

type Swapdb Completed

func (c Swapdb) Index1(Index1 int64) SwapdbIndex1 {
	return SwapdbIndex1{cs: append(c.cs, strconv.FormatInt(Index1, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Swapdb() (c Swapdb) {
	c.cs = append(b.get(), "SWAPDB")
	return
}

type SwapdbIndex1 Completed

func (c SwapdbIndex1) Index2(Index2 int64) SwapdbIndex2 {
	return SwapdbIndex2{cs: append(c.cs, strconv.FormatInt(Index2, 10)), cf: c.cf, ks: c.ks}
}

type SwapdbIndex2 Completed

func (c SwapdbIndex2) Build() Completed {
	return Completed(c)
}

type Sync Completed

func (c Sync) Build() Completed {
	return Completed(c)
}

func (b *Builder) Sync() (c Sync) {
	c.cs = append(b.get(), "SYNC")
	return
}

type Time Completed

func (c Time) Build() Completed {
	return Completed(c)
}

func (b *Builder) Time() (c Time) {
	c.cs = append(b.get(), "TIME")
	return
}

type Touch Completed

func (c Touch) Key(Key ...string) TouchKey {
	return TouchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Touch() (c Touch) {
	c.cs = append(b.get(), "TOUCH")
	c.cf = readonly
	return
}

type TouchKey Completed

func (c TouchKey) Key(Key ...string) TouchKey {
	return TouchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c TouchKey) Build() Completed {
	return Completed(c)
}

type Ttl Completed

func (c Ttl) Key(Key string) TtlKey {
	return TtlKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Ttl() (c Ttl) {
	c.cs = append(b.get(), "TTL")
	c.cf = readonly
	return
}

type TtlKey Completed

func (c TtlKey) Build() Completed {
	return Completed(c)
}

func (c TtlKey) Cache() Cacheable {
	return Cacheable(c)
}

type Type Completed

func (c Type) Key(Key string) TypeKey {
	return TypeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Type() (c Type) {
	c.cs = append(b.get(), "TYPE")
	c.cf = readonly
	return
}

type TypeKey Completed

func (c TypeKey) Build() Completed {
	return Completed(c)
}

func (c TypeKey) Cache() Cacheable {
	return Cacheable(c)
}

type Unlink Completed

func (c Unlink) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Unlink() (c Unlink) {
	c.cs = append(b.get(), "UNLINK")
	return
}

type UnlinkKey Completed

func (c UnlinkKey) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c UnlinkKey) Build() Completed {
	return Completed(c)
}

type Unsubscribe Completed

func (c Unsubscribe) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (c Unsubscribe) Build() Completed {
	return Completed(c)
}

func (b *Builder) Unsubscribe() (c Unsubscribe) {
	c.cs = append(b.get(), "UNSUBSCRIBE")
	c.cf = noRetTag
	return
}

type UnsubscribeChannel Completed

func (c UnsubscribeChannel) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (c UnsubscribeChannel) Build() Completed {
	return Completed(c)
}

type Unwatch Completed

func (c Unwatch) Build() Completed {
	return Completed(c)
}

func (b *Builder) Unwatch() (c Unwatch) {
	c.cs = append(b.get(), "UNWATCH")
	return
}

type Wait Completed

func (c Wait) Numreplicas(Numreplicas int64) WaitNumreplicas {
	return WaitNumreplicas{cs: append(c.cs, strconv.FormatInt(Numreplicas, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Wait() (c Wait) {
	c.cs = append(b.get(), "WAIT")
	c.cf = blockTag
	return
}

type WaitNumreplicas Completed

func (c WaitNumreplicas) Timeout(Timeout int64) WaitTimeout {
	return WaitTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10)), cf: c.cf, ks: c.ks}
}

type WaitTimeout Completed

func (c WaitTimeout) Build() Completed {
	return Completed(c)
}

type Watch Completed

func (c Watch) Key(Key ...string) WatchKey {
	return WatchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *Builder) Watch() (c Watch) {
	c.cs = append(b.get(), "WATCH")
	return
}

type WatchKey Completed

func (c WatchKey) Key(Key ...string) WatchKey {
	return WatchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c WatchKey) Build() Completed {
	return Completed(c)
}

type Xack Completed

func (c Xack) Key(Key string) XackKey {
	return XackKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xack() (c Xack) {
	c.cs = append(b.get(), "XACK")
	return
}

type XackGroup Completed

func (c XackGroup) Id(Id ...string) XackId {
	return XackId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

type XackId Completed

func (c XackId) Id(Id ...string) XackId {
	return XackId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c XackId) Build() Completed {
	return Completed(c)
}

type XackKey Completed

func (c XackKey) Group(Group string) XackGroup {
	return XackGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type Xadd Completed

func (c Xadd) Key(Key string) XaddKey {
	return XaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xadd() (c Xadd) {
	c.cs = append(b.get(), "XADD")
	return
}

type XaddFieldValue Completed

func (c XaddFieldValue) FieldValue(Field string, Value string) XaddFieldValue {
	return XaddFieldValue{cs: append(c.cs, Field, Value), cf: c.cf, ks: c.ks}
}

func (c XaddFieldValue) Build() Completed {
	return Completed(c)
}

type XaddId Completed

func (c XaddId) FieldValue() XaddFieldValue {
	return XaddFieldValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

type XaddKey Completed

func (c XaddKey) Nomkstream() XaddNomkstream {
	return XaddNomkstream{cs: append(c.cs, "NOMKSTREAM"), cf: c.cf, ks: c.ks}
}

func (c XaddKey) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN"), cf: c.cf, ks: c.ks}
}

func (c XaddKey) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cs: append(c.cs, "MINID"), cf: c.cf, ks: c.ks}
}

func (c XaddKey) Id(Id string) XaddId {
	return XaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type XaddNomkstream Completed

func (c XaddNomkstream) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN"), cf: c.cf, ks: c.ks}
}

func (c XaddNomkstream) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cs: append(c.cs, "MINID"), cf: c.cf, ks: c.ks}
}

func (c XaddNomkstream) Id(Id string) XaddId {
	return XaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type XaddTrimLimit Completed

func (c XaddTrimLimit) Id(Id string) XaddId {
	return XaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type XaddTrimOperatorAlmost Completed

func (c XaddTrimOperatorAlmost) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XaddTrimOperatorExact Completed

func (c XaddTrimOperatorExact) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XaddTrimStrategyMaxlen Completed

func (c XaddTrimStrategyMaxlen) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c XaddTrimStrategyMaxlen) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c XaddTrimStrategyMaxlen) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XaddTrimStrategyMinid Completed

func (c XaddTrimStrategyMinid) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c XaddTrimStrategyMinid) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c XaddTrimStrategyMinid) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XaddTrimThreshold Completed

func (c XaddTrimThreshold) Limit(Count int64) XaddTrimLimit {
	return XaddTrimLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XaddTrimThreshold) Id(Id string) XaddId {
	return XaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type Xautoclaim Completed

func (c Xautoclaim) Key(Key string) XautoclaimKey {
	return XautoclaimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xautoclaim() (c Xautoclaim) {
	c.cs = append(b.get(), "XAUTOCLAIM")
	return
}

type XautoclaimConsumer Completed

func (c XautoclaimConsumer) MinIdleTime(MinIdleTime string) XautoclaimMinIdleTime {
	return XautoclaimMinIdleTime{cs: append(c.cs, MinIdleTime), cf: c.cf, ks: c.ks}
}

type XautoclaimCount Completed

func (c XautoclaimCount) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c XautoclaimCount) Build() Completed {
	return Completed(c)
}

type XautoclaimGroup Completed

func (c XautoclaimGroup) Consumer(Consumer string) XautoclaimConsumer {
	return XautoclaimConsumer{cs: append(c.cs, Consumer), cf: c.cf, ks: c.ks}
}

type XautoclaimJustidJustid Completed

func (c XautoclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XautoclaimKey Completed

func (c XautoclaimKey) Group(Group string) XautoclaimGroup {
	return XautoclaimGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type XautoclaimMinIdleTime Completed

func (c XautoclaimMinIdleTime) Start(Start string) XautoclaimStart {
	return XautoclaimStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type XautoclaimStart Completed

func (c XautoclaimStart) Count(Count int64) XautoclaimCount {
	return XautoclaimCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XautoclaimStart) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c XautoclaimStart) Build() Completed {
	return Completed(c)
}

type Xclaim Completed

func (c Xclaim) Key(Key string) XclaimKey {
	return XclaimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xclaim() (c Xclaim) {
	c.cs = append(b.get(), "XCLAIM")
	return
}

type XclaimConsumer Completed

func (c XclaimConsumer) MinIdleTime(MinIdleTime string) XclaimMinIdleTime {
	return XclaimMinIdleTime{cs: append(c.cs, MinIdleTime), cf: c.cf, ks: c.ks}
}

type XclaimForceForce Completed

func (c XclaimForceForce) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c XclaimForceForce) Build() Completed {
	return Completed(c)
}

type XclaimGroup Completed

func (c XclaimGroup) Consumer(Consumer string) XclaimConsumer {
	return XclaimConsumer{cs: append(c.cs, Consumer), cf: c.cf, ks: c.ks}
}

type XclaimId Completed

func (c XclaimId) Idle(Ms int64) XclaimIdle {
	return XclaimIdle{cs: append(c.cs, "IDLE", strconv.FormatInt(Ms, 10)), cf: c.cf, ks: c.ks}
}

func (c XclaimId) Time(MsUnixTime int64) XclaimTime {
	return XclaimTime{cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10)), cf: c.cf, ks: c.ks}
}

func (c XclaimId) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XclaimId) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c XclaimId) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c XclaimId) Id(Id ...string) XclaimId {
	return XclaimId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c XclaimId) Build() Completed {
	return Completed(c)
}

type XclaimIdle Completed

func (c XclaimIdle) Time(MsUnixTime int64) XclaimTime {
	return XclaimTime{cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10)), cf: c.cf, ks: c.ks}
}

func (c XclaimIdle) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XclaimIdle) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c XclaimIdle) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c XclaimIdle) Build() Completed {
	return Completed(c)
}

type XclaimJustidJustid Completed

func (c XclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XclaimKey Completed

func (c XclaimKey) Group(Group string) XclaimGroup {
	return XclaimGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type XclaimMinIdleTime Completed

func (c XclaimMinIdleTime) Id(Id ...string) XclaimId {
	return XclaimId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

type XclaimRetrycount Completed

func (c XclaimRetrycount) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c XclaimRetrycount) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c XclaimRetrycount) Build() Completed {
	return Completed(c)
}

type XclaimTime Completed

func (c XclaimTime) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XclaimTime) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c XclaimTime) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c XclaimTime) Build() Completed {
	return Completed(c)
}

type Xdel Completed

func (c Xdel) Key(Key string) XdelKey {
	return XdelKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xdel() (c Xdel) {
	c.cs = append(b.get(), "XDEL")
	return
}

type XdelId Completed

func (c XdelId) Id(Id ...string) XdelId {
	return XdelId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c XdelId) Build() Completed {
	return Completed(c)
}

type XdelKey Completed

func (c XdelKey) Id(Id ...string) XdelId {
	return XdelId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

type Xgroup Completed

func (c Xgroup) Create(Key string, Groupname string) XgroupCreateCreate {
	return XgroupCreateCreate{cs: append(c.cs, "CREATE", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c Xgroup) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c Xgroup) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c Xgroup) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c Xgroup) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c Xgroup) Build() Completed {
	return Completed(c)
}

func (b *Builder) Xgroup() (c Xgroup) {
	c.cs = append(b.get(), "XGROUP")
	return
}

type XgroupCreateCreate Completed

func (c XgroupCreateCreate) Id(Id string) XgroupCreateId {
	return XgroupCreateId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type XgroupCreateId Completed

func (c XgroupCreateId) Mkstream() XgroupCreateMkstream {
	return XgroupCreateMkstream{cs: append(c.cs, "MKSTREAM"), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateId) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateId) Build() Completed {
	return Completed(c)
}

type XgroupCreateMkstream Completed

func (c XgroupCreateMkstream) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateMkstream) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateMkstream) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateMkstream) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateMkstream) Build() Completed {
	return Completed(c)
}

type XgroupCreateconsumer Completed

func (c XgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupCreateconsumer) Build() Completed {
	return Completed(c)
}

type XgroupDelconsumer Completed

func (c XgroupDelconsumer) Build() Completed {
	return Completed(c)
}

type XgroupDestroy Completed

func (c XgroupDestroy) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupDestroy) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupDestroy) Build() Completed {
	return Completed(c)
}

type XgroupSetidId Completed

func (c XgroupSetidId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c XgroupSetidId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupSetidId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c XgroupSetidId) Build() Completed {
	return Completed(c)
}

type XgroupSetidSetid Completed

func (c XgroupSetidSetid) Id(Id string) XgroupSetidId {
	return XgroupSetidId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type Xinfo Completed

func (c Xinfo) Consumers(Key string, Groupname string) XinfoConsumers {
	return XinfoConsumers{cs: append(c.cs, "CONSUMERS", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c Xinfo) Groups(Key string) XinfoGroups {
	return XinfoGroups{cs: append(c.cs, "GROUPS", Key), cf: c.cf, ks: c.ks}
}

func (c Xinfo) Stream(Key string) XinfoStream {
	return XinfoStream{cs: append(c.cs, "STREAM", Key), cf: c.cf, ks: c.ks}
}

func (c Xinfo) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c Xinfo) Build() Completed {
	return Completed(c)
}

func (b *Builder) Xinfo() (c Xinfo) {
	c.cs = append(b.get(), "XINFO")
	c.cf = readonly
	return
}

type XinfoConsumers Completed

func (c XinfoConsumers) Groups(Key string) XinfoGroups {
	return XinfoGroups{cs: append(c.cs, "GROUPS", Key), cf: c.cf, ks: c.ks}
}

func (c XinfoConsumers) Stream(Key string) XinfoStream {
	return XinfoStream{cs: append(c.cs, "STREAM", Key), cf: c.cf, ks: c.ks}
}

func (c XinfoConsumers) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c XinfoConsumers) Build() Completed {
	return Completed(c)
}

type XinfoGroups Completed

func (c XinfoGroups) Stream(Key string) XinfoStream {
	return XinfoStream{cs: append(c.cs, "STREAM", Key), cf: c.cf, ks: c.ks}
}

func (c XinfoGroups) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c XinfoGroups) Build() Completed {
	return Completed(c)
}

type XinfoHelpHelp Completed

func (c XinfoHelpHelp) Build() Completed {
	return Completed(c)
}

type XinfoStream Completed

func (c XinfoStream) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c XinfoStream) Build() Completed {
	return Completed(c)
}

type Xlen Completed

func (c Xlen) Key(Key string) XlenKey {
	return XlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xlen() (c Xlen) {
	c.cs = append(b.get(), "XLEN")
	c.cf = readonly
	return
}

type XlenKey Completed

func (c XlenKey) Build() Completed {
	return Completed(c)
}

type Xpending Completed

func (c Xpending) Key(Key string) XpendingKey {
	return XpendingKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xpending() (c Xpending) {
	c.cs = append(b.get(), "XPENDING")
	c.cf = readonly
	return
}

type XpendingFiltersConsumer Completed

func (c XpendingFiltersConsumer) Build() Completed {
	return Completed(c)
}

type XpendingFiltersCount Completed

func (c XpendingFiltersCount) Consumer(Consumer string) XpendingFiltersConsumer {
	return XpendingFiltersConsumer{cs: append(c.cs, Consumer), cf: c.cf, ks: c.ks}
}

func (c XpendingFiltersCount) Build() Completed {
	return Completed(c)
}

type XpendingFiltersEnd Completed

func (c XpendingFiltersEnd) Count(Count int64) XpendingFiltersCount {
	return XpendingFiltersCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

type XpendingFiltersIdle Completed

func (c XpendingFiltersIdle) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type XpendingFiltersStart Completed

func (c XpendingFiltersStart) End(End string) XpendingFiltersEnd {
	return XpendingFiltersEnd{cs: append(c.cs, End), cf: c.cf, ks: c.ks}
}

type XpendingGroup Completed

func (c XpendingGroup) Idle(MinIdleTime int64) XpendingFiltersIdle {
	return XpendingFiltersIdle{cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10)), cf: c.cf, ks: c.ks}
}

func (c XpendingGroup) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type XpendingKey Completed

func (c XpendingKey) Group(Group string) XpendingGroup {
	return XpendingGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type Xrange Completed

func (c Xrange) Key(Key string) XrangeKey {
	return XrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xrange() (c Xrange) {
	c.cs = append(b.get(), "XRANGE")
	c.cf = readonly
	return
}

type XrangeCount Completed

func (c XrangeCount) Build() Completed {
	return Completed(c)
}

type XrangeEnd Completed

func (c XrangeEnd) Count(Count int64) XrangeCount {
	return XrangeCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XrangeEnd) Build() Completed {
	return Completed(c)
}

type XrangeKey Completed

func (c XrangeKey) Start(Start string) XrangeStart {
	return XrangeStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type XrangeStart Completed

func (c XrangeStart) End(End string) XrangeEnd {
	return XrangeEnd{cs: append(c.cs, End), cf: c.cf, ks: c.ks}
}

type Xread Completed

func (c Xread) Count(Count int64) XreadCount {
	return XreadCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c Xread) Block(Milliseconds int64) XreadBlock {
	c.cf = blockTag
	return XreadBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c Xread) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xread() (c Xread) {
	c.cs = append(b.get(), "XREAD")
	c.cf = readonly
	return
}

type XreadBlock Completed

func (c XreadBlock) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type XreadCount Completed

func (c XreadCount) Block(Milliseconds int64) XreadBlock {
	c.cf = blockTag
	return XreadBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c XreadCount) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type XreadId Completed

func (c XreadId) Id(Id ...string) XreadId {
	return XreadId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c XreadId) Build() Completed {
	return Completed(c)
}

type XreadKey Completed

func (c XreadKey) Id(Id ...string) XreadId {
	return XreadId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c XreadKey) Key(Key ...string) XreadKey {
	return XreadKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type XreadStreamsStreams Completed

func (c XreadStreamsStreams) Key(Key ...string) XreadKey {
	return XreadKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type Xreadgroup Completed

func (c Xreadgroup) Group(Group string, Consumer string) XreadgroupGroup {
	return XreadgroupGroup{cs: append(c.cs, "GROUP", Group, Consumer), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xreadgroup() (c Xreadgroup) {
	c.cs = append(b.get(), "XREADGROUP")
	return
}

type XreadgroupBlock Completed

func (c XreadgroupBlock) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cs: append(c.cs, "NOACK"), cf: c.cf, ks: c.ks}
}

func (c XreadgroupBlock) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type XreadgroupCount Completed

func (c XreadgroupCount) Block(Milliseconds int64) XreadgroupBlock {
	c.cf = blockTag
	return XreadgroupBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c XreadgroupCount) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cs: append(c.cs, "NOACK"), cf: c.cf, ks: c.ks}
}

func (c XreadgroupCount) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type XreadgroupGroup Completed

func (c XreadgroupGroup) Count(Count int64) XreadgroupCount {
	return XreadgroupCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XreadgroupGroup) Block(Milliseconds int64) XreadgroupBlock {
	c.cf = blockTag
	return XreadgroupBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c XreadgroupGroup) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cs: append(c.cs, "NOACK"), cf: c.cf, ks: c.ks}
}

func (c XreadgroupGroup) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type XreadgroupId Completed

func (c XreadgroupId) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c XreadgroupId) Build() Completed {
	return Completed(c)
}

type XreadgroupKey Completed

func (c XreadgroupKey) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c XreadgroupKey) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type XreadgroupNoackNoack Completed

func (c XreadgroupNoackNoack) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type XreadgroupStreamsStreams Completed

func (c XreadgroupStreamsStreams) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type Xrevrange Completed

func (c Xrevrange) Key(Key string) XrevrangeKey {
	return XrevrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xrevrange() (c Xrevrange) {
	c.cs = append(b.get(), "XREVRANGE")
	c.cf = readonly
	return
}

type XrevrangeCount Completed

func (c XrevrangeCount) Build() Completed {
	return Completed(c)
}

type XrevrangeEnd Completed

func (c XrevrangeEnd) Start(Start string) XrevrangeStart {
	return XrevrangeStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type XrevrangeKey Completed

func (c XrevrangeKey) End(End string) XrevrangeEnd {
	return XrevrangeEnd{cs: append(c.cs, End), cf: c.cf, ks: c.ks}
}

type XrevrangeStart Completed

func (c XrevrangeStart) Count(Count int64) XrevrangeCount {
	return XrevrangeCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XrevrangeStart) Build() Completed {
	return Completed(c)
}

type Xtrim Completed

func (c Xtrim) Key(Key string) XtrimKey {
	return XtrimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Xtrim() (c Xtrim) {
	c.cs = append(b.get(), "XTRIM")
	return
}

type XtrimKey Completed

func (c XtrimKey) Maxlen() XtrimTrimStrategyMaxlen {
	return XtrimTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN"), cf: c.cf, ks: c.ks}
}

func (c XtrimKey) Minid() XtrimTrimStrategyMinid {
	return XtrimTrimStrategyMinid{cs: append(c.cs, "MINID"), cf: c.cf, ks: c.ks}
}

type XtrimTrimLimit Completed

func (c XtrimTrimLimit) Build() Completed {
	return Completed(c)
}

type XtrimTrimOperatorAlmost Completed

func (c XtrimTrimOperatorAlmost) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XtrimTrimOperatorExact Completed

func (c XtrimTrimOperatorExact) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XtrimTrimStrategyMaxlen Completed

func (c XtrimTrimStrategyMaxlen) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c XtrimTrimStrategyMaxlen) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c XtrimTrimStrategyMaxlen) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XtrimTrimStrategyMinid Completed

func (c XtrimTrimStrategyMinid) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c XtrimTrimStrategyMinid) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c XtrimTrimStrategyMinid) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type XtrimTrimThreshold Completed

func (c XtrimTrimThreshold) Limit(Count int64) XtrimTrimLimit {
	return XtrimTrimLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c XtrimTrimThreshold) Build() Completed {
	return Completed(c)
}

type Zadd Completed

func (c Zadd) Key(Key string) ZaddKey {
	return ZaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zadd() (c Zadd) {
	c.cs = append(b.get(), "ZADD")
	return
}

type ZaddChangeCh Completed

func (c ZaddChangeCh) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c ZaddChangeCh) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ZaddComparisonGt Completed

func (c ZaddComparisonGt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c ZaddComparisonGt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c ZaddComparisonGt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ZaddComparisonLt Completed

func (c ZaddComparisonLt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c ZaddComparisonLt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c ZaddComparisonLt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ZaddConditionNx Completed

func (c ZaddConditionNx) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionNx) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionNx) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionNx) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionNx) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ZaddConditionXx Completed

func (c ZaddConditionXx) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionXx) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionXx) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionXx) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c ZaddConditionXx) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ZaddIncrementIncr Completed

func (c ZaddIncrementIncr) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ZaddKey Completed

func (c ZaddKey) Nx() ZaddConditionNx {
	return ZaddConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c ZaddKey) Xx() ZaddConditionXx {
	return ZaddConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c ZaddKey) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c ZaddKey) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c ZaddKey) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c ZaddKey) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c ZaddKey) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type ZaddScoreMember Completed

func (c ZaddScoreMember) ScoreMember(Score float64, Member string) ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member), cf: c.cf, ks: c.ks}
}

func (c ZaddScoreMember) Build() Completed {
	return Completed(c)
}

type Zcard Completed

func (c Zcard) Key(Key string) ZcardKey {
	return ZcardKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zcard() (c Zcard) {
	c.cs = append(b.get(), "ZCARD")
	c.cf = readonly
	return
}

type ZcardKey Completed

func (c ZcardKey) Build() Completed {
	return Completed(c)
}

func (c ZcardKey) Cache() Cacheable {
	return Cacheable(c)
}

type Zcount Completed

func (c Zcount) Key(Key string) ZcountKey {
	return ZcountKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zcount() (c Zcount) {
	c.cs = append(b.get(), "ZCOUNT")
	c.cf = readonly
	return
}

type ZcountKey Completed

func (c ZcountKey) Min(Min float64) ZcountMin {
	return ZcountMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type ZcountMax Completed

func (c ZcountMax) Build() Completed {
	return Completed(c)
}

func (c ZcountMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZcountMin Completed

func (c ZcountMin) Max(Max float64) ZcountMax {
	return ZcountMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type Zdiff Completed

func (c Zdiff) Numkeys(Numkeys int64) ZdiffNumkeys {
	return ZdiffNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zdiff() (c Zdiff) {
	c.cs = append(b.get(), "ZDIFF")
	c.cf = readonly
	return
}

type ZdiffKey Completed

func (c ZdiffKey) Withscores() ZdiffWithscoresWithscores {
	return ZdiffWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZdiffKey) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ZdiffKey) Build() Completed {
	return Completed(c)
}

type ZdiffNumkeys Completed

func (c ZdiffNumkeys) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type ZdiffWithscoresWithscores Completed

func (c ZdiffWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zdiffstore Completed

func (c Zdiffstore) Destination(Destination string) ZdiffstoreDestination {
	return ZdiffstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zdiffstore() (c Zdiffstore) {
	c.cs = append(b.get(), "ZDIFFSTORE")
	return
}

type ZdiffstoreDestination Completed

func (c ZdiffstoreDestination) Numkeys(Numkeys int64) ZdiffstoreNumkeys {
	return ZdiffstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type ZdiffstoreKey Completed

func (c ZdiffstoreKey) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ZdiffstoreKey) Build() Completed {
	return Completed(c)
}

type ZdiffstoreNumkeys Completed

func (c ZdiffstoreNumkeys) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type Zincrby Completed

func (c Zincrby) Key(Key string) ZincrbyKey {
	return ZincrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zincrby() (c Zincrby) {
	c.cs = append(b.get(), "ZINCRBY")
	return
}

type ZincrbyIncrement Completed

func (c ZincrbyIncrement) Member(Member string) ZincrbyMember {
	return ZincrbyMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type ZincrbyKey Completed

func (c ZincrbyKey) Increment(Increment int64) ZincrbyIncrement {
	return ZincrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

type ZincrbyMember Completed

func (c ZincrbyMember) Build() Completed {
	return Completed(c)
}

type Zinter Completed

func (c Zinter) Numkeys(Numkeys int64) ZinterNumkeys {
	return ZinterNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zinter() (c Zinter) {
	c.cs = append(b.get(), "ZINTER")
	c.cf = readonly
	return
}

type ZinterAggregateMax Completed

func (c ZinterAggregateMax) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZinterAggregateMax) Build() Completed {
	return Completed(c)
}

type ZinterAggregateMin Completed

func (c ZinterAggregateMin) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZinterAggregateMin) Build() Completed {
	return Completed(c)
}

type ZinterAggregateSum Completed

func (c ZinterAggregateSum) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZinterAggregateSum) Build() Completed {
	return Completed(c)
}

type ZinterKey Completed

func (c ZinterKey) Weights(Weight ...int64) ZinterWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZinterKey) Sum() ZinterAggregateSum {
	return ZinterAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZinterKey) Min() ZinterAggregateMin {
	return ZinterAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZinterKey) Max() ZinterAggregateMax {
	return ZinterAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZinterKey) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZinterKey) Key(Key ...string) ZinterKey {
	return ZinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ZinterKey) Build() Completed {
	return Completed(c)
}

type ZinterNumkeys Completed

func (c ZinterNumkeys) Key(Key ...string) ZinterKey {
	return ZinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type ZinterWeights Completed

func (c ZinterWeights) Sum() ZinterAggregateSum {
	return ZinterAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZinterWeights) Min() ZinterAggregateMin {
	return ZinterAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZinterWeights) Max() ZinterAggregateMax {
	return ZinterAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZinterWeights) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZinterWeights) Weights(Weights ...int64) ZinterWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZinterWeights) Build() Completed {
	return Completed(c)
}

type ZinterWithscoresWithscores Completed

func (c ZinterWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zintercard Completed

func (c Zintercard) Numkeys(Numkeys int64) ZintercardNumkeys {
	return ZintercardNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zintercard() (c Zintercard) {
	c.cs = append(b.get(), "ZINTERCARD")
	c.cf = readonly
	return
}

type ZintercardKey Completed

func (c ZintercardKey) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ZintercardKey) Build() Completed {
	return Completed(c)
}

type ZintercardNumkeys Completed

func (c ZintercardNumkeys) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type Zinterstore Completed

func (c Zinterstore) Destination(Destination string) ZinterstoreDestination {
	return ZinterstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zinterstore() (c Zinterstore) {
	c.cs = append(b.get(), "ZINTERSTORE")
	return
}

type ZinterstoreAggregateMax Completed

func (c ZinterstoreAggregateMax) Build() Completed {
	return Completed(c)
}

type ZinterstoreAggregateMin Completed

func (c ZinterstoreAggregateMin) Build() Completed {
	return Completed(c)
}

type ZinterstoreAggregateSum Completed

func (c ZinterstoreAggregateSum) Build() Completed {
	return Completed(c)
}

type ZinterstoreDestination Completed

func (c ZinterstoreDestination) Numkeys(Numkeys int64) ZinterstoreNumkeys {
	return ZinterstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type ZinterstoreKey Completed

func (c ZinterstoreKey) Weights(Weight ...int64) ZinterstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZinterstoreKey) Sum() ZinterstoreAggregateSum {
	return ZinterstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZinterstoreKey) Min() ZinterstoreAggregateMin {
	return ZinterstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZinterstoreKey) Max() ZinterstoreAggregateMax {
	return ZinterstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZinterstoreKey) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ZinterstoreKey) Build() Completed {
	return Completed(c)
}

type ZinterstoreNumkeys Completed

func (c ZinterstoreNumkeys) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type ZinterstoreWeights Completed

func (c ZinterstoreWeights) Sum() ZinterstoreAggregateSum {
	return ZinterstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZinterstoreWeights) Min() ZinterstoreAggregateMin {
	return ZinterstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZinterstoreWeights) Max() ZinterstoreAggregateMax {
	return ZinterstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZinterstoreWeights) Weights(Weights ...int64) ZinterstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZinterstoreWeights) Build() Completed {
	return Completed(c)
}

type Zlexcount Completed

func (c Zlexcount) Key(Key string) ZlexcountKey {
	return ZlexcountKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zlexcount() (c Zlexcount) {
	c.cs = append(b.get(), "ZLEXCOUNT")
	c.cf = readonly
	return
}

type ZlexcountKey Completed

func (c ZlexcountKey) Min(Min string) ZlexcountMin {
	return ZlexcountMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type ZlexcountMax Completed

func (c ZlexcountMax) Build() Completed {
	return Completed(c)
}

func (c ZlexcountMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZlexcountMin Completed

func (c ZlexcountMin) Max(Max string) ZlexcountMax {
	return ZlexcountMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type Zmscore Completed

func (c Zmscore) Key(Key string) ZmscoreKey {
	return ZmscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zmscore() (c Zmscore) {
	c.cs = append(b.get(), "ZMSCORE")
	c.cf = readonly
	return
}

type ZmscoreKey Completed

func (c ZmscoreKey) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type ZmscoreMember Completed

func (c ZmscoreMember) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c ZmscoreMember) Build() Completed {
	return Completed(c)
}

func (c ZmscoreMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zpopmax Completed

func (c Zpopmax) Key(Key string) ZpopmaxKey {
	return ZpopmaxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zpopmax() (c Zpopmax) {
	c.cs = append(b.get(), "ZPOPMAX")
	return
}

type ZpopmaxCount Completed

func (c ZpopmaxCount) Build() Completed {
	return Completed(c)
}

type ZpopmaxKey Completed

func (c ZpopmaxKey) Count(Count int64) ZpopmaxCount {
	return ZpopmaxCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZpopmaxKey) Build() Completed {
	return Completed(c)
}

type Zpopmin Completed

func (c Zpopmin) Key(Key string) ZpopminKey {
	return ZpopminKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zpopmin() (c Zpopmin) {
	c.cs = append(b.get(), "ZPOPMIN")
	return
}

type ZpopminCount Completed

func (c ZpopminCount) Build() Completed {
	return Completed(c)
}

type ZpopminKey Completed

func (c ZpopminKey) Count(Count int64) ZpopminCount {
	return ZpopminCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZpopminKey) Build() Completed {
	return Completed(c)
}

type Zrandmember Completed

func (c Zrandmember) Key(Key string) ZrandmemberKey {
	return ZrandmemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrandmember() (c Zrandmember) {
	c.cs = append(b.get(), "ZRANDMEMBER")
	c.cf = readonly
	return
}

type ZrandmemberKey Completed

func (c ZrandmemberKey) Count(Count int64) ZrandmemberOptionsCount {
	return ZrandmemberOptionsCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrandmemberKey) Build() Completed {
	return Completed(c)
}

type ZrandmemberOptionsCount Completed

func (c ZrandmemberOptionsCount) Withscores() ZrandmemberOptionsWithscoresWithscores {
	return ZrandmemberOptionsWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrandmemberOptionsCount) Build() Completed {
	return Completed(c)
}

type ZrandmemberOptionsWithscoresWithscores Completed

func (c ZrandmemberOptionsWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zrange Completed

func (c Zrange) Key(Key string) ZrangeKey {
	return ZrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrange() (c Zrange) {
	c.cs = append(b.get(), "ZRANGE")
	c.cf = readonly
	return
}

type ZrangeKey Completed

func (c ZrangeKey) Min(Min string) ZrangeMin {
	return ZrangeMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type ZrangeLimit Completed

func (c ZrangeLimit) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrangeLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangeLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeMax Completed

func (c ZrangeMax) Byscore() ZrangeSortbyByscore {
	return ZrangeSortbyByscore{cs: append(c.cs, "BYSCORE"), cf: c.cf, ks: c.ks}
}

func (c ZrangeMax) Bylex() ZrangeSortbyBylex {
	return ZrangeSortbyBylex{cs: append(c.cs, "BYLEX"), cf: c.cf, ks: c.ks}
}

func (c ZrangeMax) Rev() ZrangeRevRev {
	return ZrangeRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c ZrangeMax) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangeMax) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrangeMax) Build() Completed {
	return Completed(c)
}

func (c ZrangeMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeMin Completed

func (c ZrangeMin) Max(Max string) ZrangeMax {
	return ZrangeMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type ZrangeRevRev Completed

func (c ZrangeRevRev) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangeRevRev) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrangeRevRev) Build() Completed {
	return Completed(c)
}

func (c ZrangeRevRev) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeSortbyBylex Completed

func (c ZrangeSortbyBylex) Rev() ZrangeRevRev {
	return ZrangeRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c ZrangeSortbyBylex) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangeSortbyBylex) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrangeSortbyBylex) Build() Completed {
	return Completed(c)
}

func (c ZrangeSortbyBylex) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeSortbyByscore Completed

func (c ZrangeSortbyByscore) Rev() ZrangeRevRev {
	return ZrangeRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c ZrangeSortbyByscore) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangeSortbyByscore) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrangeSortbyByscore) Build() Completed {
	return Completed(c)
}

func (c ZrangeSortbyByscore) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeWithscoresWithscores Completed

func (c ZrangeWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrangeWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangebylex Completed

func (c Zrangebylex) Key(Key string) ZrangebylexKey {
	return ZrangebylexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrangebylex() (c Zrangebylex) {
	c.cs = append(b.get(), "ZRANGEBYLEX")
	c.cf = readonly
	return
}

type ZrangebylexKey Completed

func (c ZrangebylexKey) Min(Min string) ZrangebylexMin {
	return ZrangebylexMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type ZrangebylexLimit Completed

func (c ZrangebylexLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangebylexLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexMax Completed

func (c ZrangebylexMax) Limit(Offset int64, Count int64) ZrangebylexLimit {
	return ZrangebylexLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangebylexMax) Build() Completed {
	return Completed(c)
}

func (c ZrangebylexMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexMin Completed

func (c ZrangebylexMin) Max(Max string) ZrangebylexMax {
	return ZrangebylexMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type Zrangebyscore Completed

func (c Zrangebyscore) Key(Key string) ZrangebyscoreKey {
	return ZrangebyscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrangebyscore() (c Zrangebyscore) {
	c.cs = append(b.get(), "ZRANGEBYSCORE")
	c.cf = readonly
	return
}

type ZrangebyscoreKey Completed

func (c ZrangebyscoreKey) Min(Min float64) ZrangebyscoreMin {
	return ZrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type ZrangebyscoreLimit Completed

func (c ZrangebyscoreLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreMax Completed

func (c ZrangebyscoreMax) Withscores() ZrangebyscoreWithscoresWithscores {
	return ZrangebyscoreWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrangebyscoreMax) Limit(Offset int64, Count int64) ZrangebyscoreLimit {
	return ZrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangebyscoreMax) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreMin Completed

func (c ZrangebyscoreMin) Max(Max float64) ZrangebyscoreMax {
	return ZrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type ZrangebyscoreWithscoresWithscores Completed

func (c ZrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrangebyscoreLimit {
	return ZrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangebyscoreWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangestore Completed

func (c Zrangestore) Dst(Dst string) ZrangestoreDst {
	return ZrangestoreDst{cs: append(c.cs, Dst), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrangestore() (c Zrangestore) {
	c.cs = append(b.get(), "ZRANGESTORE")
	return
}

type ZrangestoreDst Completed

func (c ZrangestoreDst) Src(Src string) ZrangestoreSrc {
	return ZrangestoreSrc{cs: append(c.cs, Src), cf: c.cf, ks: c.ks}
}

type ZrangestoreLimit Completed

func (c ZrangestoreLimit) Build() Completed {
	return Completed(c)
}

type ZrangestoreMax Completed

func (c ZrangestoreMax) Byscore() ZrangestoreSortbyByscore {
	return ZrangestoreSortbyByscore{cs: append(c.cs, "BYSCORE"), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreMax) Bylex() ZrangestoreSortbyBylex {
	return ZrangestoreSortbyBylex{cs: append(c.cs, "BYLEX"), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreMax) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreMax) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreMax) Build() Completed {
	return Completed(c)
}

type ZrangestoreMin Completed

func (c ZrangestoreMin) Max(Max string) ZrangestoreMax {
	return ZrangestoreMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type ZrangestoreRevRev Completed

func (c ZrangestoreRevRev) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreRevRev) Build() Completed {
	return Completed(c)
}

type ZrangestoreSortbyBylex Completed

func (c ZrangestoreSortbyBylex) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreSortbyBylex) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreSortbyBylex) Build() Completed {
	return Completed(c)
}

type ZrangestoreSortbyByscore Completed

func (c ZrangestoreSortbyByscore) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreSortbyByscore) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrangestoreSortbyByscore) Build() Completed {
	return Completed(c)
}

type ZrangestoreSrc Completed

func (c ZrangestoreSrc) Min(Min string) ZrangestoreMin {
	return ZrangestoreMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type Zrank Completed

func (c Zrank) Key(Key string) ZrankKey {
	return ZrankKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrank() (c Zrank) {
	c.cs = append(b.get(), "ZRANK")
	c.cf = readonly
	return
}

type ZrankKey Completed

func (c ZrankKey) Member(Member string) ZrankMember {
	return ZrankMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type ZrankMember Completed

func (c ZrankMember) Build() Completed {
	return Completed(c)
}

func (c ZrankMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zrem Completed

func (c Zrem) Key(Key string) ZremKey {
	return ZremKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrem() (c Zrem) {
	c.cs = append(b.get(), "ZREM")
	return
}

type ZremKey Completed

func (c ZremKey) Member(Member ...string) ZremMember {
	return ZremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type ZremMember Completed

func (c ZremMember) Member(Member ...string) ZremMember {
	return ZremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c ZremMember) Build() Completed {
	return Completed(c)
}

type Zremrangebylex Completed

func (c Zremrangebylex) Key(Key string) ZremrangebylexKey {
	return ZremrangebylexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zremrangebylex() (c Zremrangebylex) {
	c.cs = append(b.get(), "ZREMRANGEBYLEX")
	return
}

type ZremrangebylexKey Completed

func (c ZremrangebylexKey) Min(Min string) ZremrangebylexMin {
	return ZremrangebylexMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type ZremrangebylexMax Completed

func (c ZremrangebylexMax) Build() Completed {
	return Completed(c)
}

type ZremrangebylexMin Completed

func (c ZremrangebylexMin) Max(Max string) ZremrangebylexMax {
	return ZremrangebylexMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type Zremrangebyrank Completed

func (c Zremrangebyrank) Key(Key string) ZremrangebyrankKey {
	return ZremrangebyrankKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zremrangebyrank() (c Zremrangebyrank) {
	c.cs = append(b.get(), "ZREMRANGEBYRANK")
	return
}

type ZremrangebyrankKey Completed

func (c ZremrangebyrankKey) Start(Start int64) ZremrangebyrankStart {
	return ZremrangebyrankStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type ZremrangebyrankStart Completed

func (c ZremrangebyrankStart) Stop(Stop int64) ZremrangebyrankStop {
	return ZremrangebyrankStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type ZremrangebyrankStop Completed

func (c ZremrangebyrankStop) Build() Completed {
	return Completed(c)
}

type Zremrangebyscore Completed

func (c Zremrangebyscore) Key(Key string) ZremrangebyscoreKey {
	return ZremrangebyscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zremrangebyscore() (c Zremrangebyscore) {
	c.cs = append(b.get(), "ZREMRANGEBYSCORE")
	return
}

type ZremrangebyscoreKey Completed

func (c ZremrangebyscoreKey) Min(Min float64) ZremrangebyscoreMin {
	return ZremrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type ZremrangebyscoreMax Completed

func (c ZremrangebyscoreMax) Build() Completed {
	return Completed(c)
}

type ZremrangebyscoreMin Completed

func (c ZremrangebyscoreMin) Max(Max float64) ZremrangebyscoreMax {
	return ZremrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type Zrevrange Completed

func (c Zrevrange) Key(Key string) ZrevrangeKey {
	return ZrevrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrevrange() (c Zrevrange) {
	c.cs = append(b.get(), "ZREVRANGE")
	c.cf = readonly
	return
}

type ZrevrangeKey Completed

func (c ZrevrangeKey) Start(Start int64) ZrevrangeStart {
	return ZrevrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type ZrevrangeStart Completed

func (c ZrevrangeStart) Stop(Stop int64) ZrevrangeStop {
	return ZrevrangeStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type ZrevrangeStop Completed

func (c ZrevrangeStop) Withscores() ZrevrangeWithscoresWithscores {
	return ZrevrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrevrangeStop) Build() Completed {
	return Completed(c)
}

func (c ZrevrangeStop) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangeWithscoresWithscores Completed

func (c ZrevrangeWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrevrangeWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrangebylex Completed

func (c Zrevrangebylex) Key(Key string) ZrevrangebylexKey {
	return ZrevrangebylexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrevrangebylex() (c Zrevrangebylex) {
	c.cs = append(b.get(), "ZREVRANGEBYLEX")
	c.cf = readonly
	return
}

type ZrevrangebylexKey Completed

func (c ZrevrangebylexKey) Max(Max string) ZrevrangebylexMax {
	return ZrevrangebylexMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type ZrevrangebylexLimit Completed

func (c ZrevrangebylexLimit) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebylexLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebylexMax Completed

func (c ZrevrangebylexMax) Min(Min string) ZrevrangebylexMin {
	return ZrevrangebylexMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type ZrevrangebylexMin Completed

func (c ZrevrangebylexMin) Limit(Offset int64, Count int64) ZrevrangebylexLimit {
	return ZrevrangebylexLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrevrangebylexMin) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebylexMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrangebyscore Completed

func (c Zrevrangebyscore) Key(Key string) ZrevrangebyscoreKey {
	return ZrevrangebyscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrevrangebyscore() (c Zrevrangebyscore) {
	c.cs = append(b.get(), "ZREVRANGEBYSCORE")
	c.cf = readonly
	return
}

type ZrevrangebyscoreKey Completed

func (c ZrevrangebyscoreKey) Max(Max float64) ZrevrangebyscoreMax {
	return ZrevrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type ZrevrangebyscoreLimit Completed

func (c ZrevrangebyscoreLimit) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreMax Completed

func (c ZrevrangebyscoreMax) Min(Min float64) ZrevrangebyscoreMin {
	return ZrevrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type ZrevrangebyscoreMin Completed

func (c ZrevrangebyscoreMin) Withscores() ZrevrangebyscoreWithscoresWithscores {
	return ZrevrangebyscoreWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZrevrangebyscoreMin) Limit(Offset int64, Count int64) ZrevrangebyscoreLimit {
	return ZrevrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrevrangebyscoreMin) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreMin) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreWithscoresWithscores Completed

func (c ZrevrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrevrangebyscoreLimit {
	return ZrevrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZrevrangebyscoreWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrank Completed

func (c Zrevrank) Key(Key string) ZrevrankKey {
	return ZrevrankKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zrevrank() (c Zrevrank) {
	c.cs = append(b.get(), "ZREVRANK")
	c.cf = readonly
	return
}

type ZrevrankKey Completed

func (c ZrevrankKey) Member(Member string) ZrevrankMember {
	return ZrevrankMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type ZrevrankMember Completed

func (c ZrevrankMember) Build() Completed {
	return Completed(c)
}

func (c ZrevrankMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zscan Completed

func (c Zscan) Key(Key string) ZscanKey {
	return ZscanKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zscan() (c Zscan) {
	c.cs = append(b.get(), "ZSCAN")
	c.cf = readonly
	return
}

type ZscanCount Completed

func (c ZscanCount) Build() Completed {
	return Completed(c)
}

type ZscanCursor Completed

func (c ZscanCursor) Match(Pattern string) ZscanMatch {
	return ZscanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c ZscanCursor) Count(Count int64) ZscanCount {
	return ZscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZscanCursor) Build() Completed {
	return Completed(c)
}

type ZscanKey Completed

func (c ZscanKey) Cursor(Cursor int64) ZscanCursor {
	return ZscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

type ZscanMatch Completed

func (c ZscanMatch) Count(Count int64) ZscanCount {
	return ZscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c ZscanMatch) Build() Completed {
	return Completed(c)
}

type Zscore Completed

func (c Zscore) Key(Key string) ZscoreKey {
	return ZscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zscore() (c Zscore) {
	c.cs = append(b.get(), "ZSCORE")
	c.cf = readonly
	return
}

type ZscoreKey Completed

func (c ZscoreKey) Member(Member string) ZscoreMember {
	return ZscoreMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type ZscoreMember Completed

func (c ZscoreMember) Build() Completed {
	return Completed(c)
}

func (c ZscoreMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zunion Completed

func (c Zunion) Numkeys(Numkeys int64) ZunionNumkeys {
	return ZunionNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zunion() (c Zunion) {
	c.cs = append(b.get(), "ZUNION")
	c.cf = readonly
	return
}

type ZunionAggregateMax Completed

func (c ZunionAggregateMax) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZunionAggregateMax) Build() Completed {
	return Completed(c)
}

type ZunionAggregateMin Completed

func (c ZunionAggregateMin) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZunionAggregateMin) Build() Completed {
	return Completed(c)
}

type ZunionAggregateSum Completed

func (c ZunionAggregateSum) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZunionAggregateSum) Build() Completed {
	return Completed(c)
}

type ZunionKey Completed

func (c ZunionKey) Weights(Weight ...int64) ZunionWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZunionKey) Sum() ZunionAggregateSum {
	return ZunionAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZunionKey) Min() ZunionAggregateMin {
	return ZunionAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZunionKey) Max() ZunionAggregateMax {
	return ZunionAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZunionKey) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZunionKey) Key(Key ...string) ZunionKey {
	return ZunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ZunionKey) Build() Completed {
	return Completed(c)
}

type ZunionNumkeys Completed

func (c ZunionNumkeys) Key(Key ...string) ZunionKey {
	return ZunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type ZunionWeights Completed

func (c ZunionWeights) Sum() ZunionAggregateSum {
	return ZunionAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZunionWeights) Min() ZunionAggregateMin {
	return ZunionAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZunionWeights) Max() ZunionAggregateMax {
	return ZunionAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZunionWeights) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c ZunionWeights) Weights(Weights ...int64) ZunionWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZunionWeights) Build() Completed {
	return Completed(c)
}

type ZunionWithscoresWithscores Completed

func (c ZunionWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zunionstore Completed

func (c Zunionstore) Destination(Destination string) ZunionstoreDestination {
	return ZunionstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *Builder) Zunionstore() (c Zunionstore) {
	c.cs = append(b.get(), "ZUNIONSTORE")
	return
}

type ZunionstoreAggregateMax Completed

func (c ZunionstoreAggregateMax) Build() Completed {
	return Completed(c)
}

type ZunionstoreAggregateMin Completed

func (c ZunionstoreAggregateMin) Build() Completed {
	return Completed(c)
}

type ZunionstoreAggregateSum Completed

func (c ZunionstoreAggregateSum) Build() Completed {
	return Completed(c)
}

type ZunionstoreDestination Completed

func (c ZunionstoreDestination) Numkeys(Numkeys int64) ZunionstoreNumkeys {
	return ZunionstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type ZunionstoreKey Completed

func (c ZunionstoreKey) Weights(Weight ...int64) ZunionstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZunionstoreKey) Sum() ZunionstoreAggregateSum {
	return ZunionstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZunionstoreKey) Min() ZunionstoreAggregateMin {
	return ZunionstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZunionstoreKey) Max() ZunionstoreAggregateMax {
	return ZunionstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZunionstoreKey) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c ZunionstoreKey) Build() Completed {
	return Completed(c)
}

type ZunionstoreNumkeys Completed

func (c ZunionstoreNumkeys) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type ZunionstoreWeights Completed

func (c ZunionstoreWeights) Sum() ZunionstoreAggregateSum {
	return ZunionstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c ZunionstoreWeights) Min() ZunionstoreAggregateMin {
	return ZunionstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c ZunionstoreWeights) Max() ZunionstoreAggregateMax {
	return ZunionstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c ZunionstoreWeights) Weights(Weights ...int64) ZunionstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c ZunionstoreWeights) Build() Completed {
	return Completed(c)
}

type SAclCat SCompleted

func (c SAclCat) Categoryname(Categoryname string) SAclCatCategoryname {
	return SAclCatCategoryname{cs: append(c.cs, Categoryname), cf: c.cf, ks: c.ks}
}

func (c SAclCat) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclCat() (c SAclCat) {
	c.cs = append(b.get(), "ACL", "CAT")
	c.ks = InitSlot
	return
}

type SAclCatCategoryname SCompleted

func (c SAclCatCategoryname) Build() SCompleted {
	return SCompleted(c)
}

type SAclDeluser SCompleted

func (c SAclDeluser) Username(Username ...string) SAclDeluserUsername {
	return SAclDeluserUsername{cs: append(c.cs, Username...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) AclDeluser() (c SAclDeluser) {
	c.cs = append(b.get(), "ACL", "DELUSER")
	c.ks = InitSlot
	return
}

type SAclDeluserUsername SCompleted

func (c SAclDeluserUsername) Username(Username ...string) SAclDeluserUsername {
	return SAclDeluserUsername{cs: append(c.cs, Username...), cf: c.cf, ks: c.ks}
}

func (c SAclDeluserUsername) Build() SCompleted {
	return SCompleted(c)
}

type SAclGenpass SCompleted

func (c SAclGenpass) Bits(Bits int64) SAclGenpassBits {
	return SAclGenpassBits{cs: append(c.cs, strconv.FormatInt(Bits, 10)), cf: c.cf, ks: c.ks}
}

func (c SAclGenpass) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclGenpass() (c SAclGenpass) {
	c.cs = append(b.get(), "ACL", "GENPASS")
	c.ks = InitSlot
	return
}

type SAclGenpassBits SCompleted

func (c SAclGenpassBits) Build() SCompleted {
	return SCompleted(c)
}

type SAclGetuser SCompleted

func (c SAclGetuser) Username(Username string) SAclGetuserUsername {
	return SAclGetuserUsername{cs: append(c.cs, Username), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) AclGetuser() (c SAclGetuser) {
	c.cs = append(b.get(), "ACL", "GETUSER")
	c.ks = InitSlot
	return
}

type SAclGetuserUsername SCompleted

func (c SAclGetuserUsername) Build() SCompleted {
	return SCompleted(c)
}

type SAclHelp SCompleted

func (c SAclHelp) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclHelp() (c SAclHelp) {
	c.cs = append(b.get(), "ACL", "HELP")
	c.ks = InitSlot
	return
}

type SAclList SCompleted

func (c SAclList) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclList() (c SAclList) {
	c.cs = append(b.get(), "ACL", "LIST")
	c.ks = InitSlot
	return
}

type SAclLoad SCompleted

func (c SAclLoad) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclLoad() (c SAclLoad) {
	c.cs = append(b.get(), "ACL", "LOAD")
	c.ks = InitSlot
	return
}

type SAclLog SCompleted

func (c SAclLog) CountOrReset(CountOrReset string) SAclLogCountOrReset {
	return SAclLogCountOrReset{cs: append(c.cs, CountOrReset), cf: c.cf, ks: c.ks}
}

func (c SAclLog) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclLog() (c SAclLog) {
	c.cs = append(b.get(), "ACL", "LOG")
	c.ks = InitSlot
	return
}

type SAclLogCountOrReset SCompleted

func (c SAclLogCountOrReset) Build() SCompleted {
	return SCompleted(c)
}

type SAclSave SCompleted

func (c SAclSave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclSave() (c SAclSave) {
	c.cs = append(b.get(), "ACL", "SAVE")
	c.ks = InitSlot
	return
}

type SAclSetuser SCompleted

func (c SAclSetuser) Username(Username string) SAclSetuserUsername {
	return SAclSetuserUsername{cs: append(c.cs, Username), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) AclSetuser() (c SAclSetuser) {
	c.cs = append(b.get(), "ACL", "SETUSER")
	c.ks = InitSlot
	return
}

type SAclSetuserRule SCompleted

func (c SAclSetuserRule) Rule(Rule ...string) SAclSetuserRule {
	return SAclSetuserRule{cs: append(c.cs, Rule...), cf: c.cf, ks: c.ks}
}

func (c SAclSetuserRule) Build() SCompleted {
	return SCompleted(c)
}

type SAclSetuserUsername SCompleted

func (c SAclSetuserUsername) Rule(Rule ...string) SAclSetuserRule {
	return SAclSetuserRule{cs: append(c.cs, Rule...), cf: c.cf, ks: c.ks}
}

func (c SAclSetuserUsername) Build() SCompleted {
	return SCompleted(c)
}

type SAclUsers SCompleted

func (c SAclUsers) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclUsers() (c SAclUsers) {
	c.cs = append(b.get(), "ACL", "USERS")
	c.ks = InitSlot
	return
}

type SAclWhoami SCompleted

func (c SAclWhoami) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) AclWhoami() (c SAclWhoami) {
	c.cs = append(b.get(), "ACL", "WHOAMI")
	c.ks = InitSlot
	return
}

type SAppend SCompleted

func (c SAppend) Key(Key string) SAppendKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SAppendKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Append() (c SAppend) {
	c.cs = append(b.get(), "APPEND")
	c.ks = InitSlot
	return
}

type SAppendKey SCompleted

func (c SAppendKey) Value(Value string) SAppendValue {
	return SAppendValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SAppendValue SCompleted

func (c SAppendValue) Build() SCompleted {
	return SCompleted(c)
}

type SAsking SCompleted

func (c SAsking) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Asking() (c SAsking) {
	c.cs = append(b.get(), "ASKING")
	c.ks = InitSlot
	return
}

type SAuth SCompleted

func (c SAuth) Username(Username string) SAuthUsername {
	return SAuthUsername{cs: append(c.cs, Username), cf: c.cf, ks: c.ks}
}

func (c SAuth) Password(Password string) SAuthPassword {
	return SAuthPassword{cs: append(c.cs, Password), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Auth() (c SAuth) {
	c.cs = append(b.get(), "AUTH")
	c.ks = InitSlot
	return
}

type SAuthPassword SCompleted

func (c SAuthPassword) Build() SCompleted {
	return SCompleted(c)
}

type SAuthUsername SCompleted

func (c SAuthUsername) Password(Password string) SAuthPassword {
	return SAuthPassword{cs: append(c.cs, Password), cf: c.cf, ks: c.ks}
}

type SBgrewriteaof SCompleted

func (c SBgrewriteaof) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Bgrewriteaof() (c SBgrewriteaof) {
	c.cs = append(b.get(), "BGREWRITEAOF")
	c.ks = InitSlot
	return
}

type SBgsave SCompleted

func (c SBgsave) Schedule() SBgsaveScheduleSchedule {
	return SBgsaveScheduleSchedule{cs: append(c.cs, "SCHEDULE"), cf: c.cf, ks: c.ks}
}

func (c SBgsave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Bgsave() (c SBgsave) {
	c.cs = append(b.get(), "BGSAVE")
	c.ks = InitSlot
	return
}

type SBgsaveScheduleSchedule SCompleted

func (c SBgsaveScheduleSchedule) Build() SCompleted {
	return SCompleted(c)
}

type SBitcount SCompleted

func (c SBitcount) Key(Key string) SBitcountKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SBitcountKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Bitcount() (c SBitcount) {
	c.cs = append(b.get(), "BITCOUNT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SBitcountKey SCompleted

func (c SBitcountKey) StartEnd(Start int64, End int64) SBitcountStartEnd {
	return SBitcountStartEnd{cs: append(c.cs, strconv.FormatInt(Start, 10), strconv.FormatInt(End, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitcountKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitcountKey) Cache() SCacheable {
	return SCacheable(c)
}

type SBitcountStartEnd SCompleted

func (c SBitcountStartEnd) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitcountStartEnd) Cache() SCacheable {
	return SCacheable(c)
}

type SBitfield SCompleted

func (c SBitfield) Key(Key string) SBitfieldKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SBitfieldKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Bitfield() (c SBitfield) {
	c.cs = append(b.get(), "BITFIELD")
	c.ks = InitSlot
	return
}

type SBitfieldFail SCompleted

func (c SBitfieldFail) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldGet SCompleted

func (c SBitfieldGet) Set(Type string, Offset int64, Value int64) SBitfieldSet {
	return SBitfieldSet{cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitfieldGet) Incrby(Type string, Offset int64, Increment int64) SBitfieldIncrby {
	return SBitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitfieldGet) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldGet) Sat() SBitfieldSat {
	return SBitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldGet) Fail() SBitfieldFail {
	return SBitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldGet) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldIncrby SCompleted

func (c SBitfieldIncrby) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldIncrby) Sat() SBitfieldSat {
	return SBitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldIncrby) Fail() SBitfieldFail {
	return SBitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldIncrby) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldKey SCompleted

func (c SBitfieldKey) Get(Type string, Offset int64) SBitfieldGet {
	return SBitfieldGet{cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitfieldKey) Set(Type string, Offset int64, Value int64) SBitfieldSet {
	return SBitfieldSet{cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitfieldKey) Incrby(Type string, Offset int64, Increment int64) SBitfieldIncrby {
	return SBitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitfieldKey) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldKey) Sat() SBitfieldSat {
	return SBitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldKey) Fail() SBitfieldFail {
	return SBitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldKey) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldRo SCompleted

func (c SBitfieldRo) Key(Key string) SBitfieldRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SBitfieldRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) BitfieldRo() (c SBitfieldRo) {
	c.cs = append(b.get(), "BITFIELD_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SBitfieldRoGet SCompleted

func (c SBitfieldRoGet) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitfieldRoGet) Cache() SCacheable {
	return SCacheable(c)
}

type SBitfieldRoKey SCompleted

func (c SBitfieldRoKey) Get(Type string, Offset int64) SBitfieldRoGet {
	return SBitfieldRoGet{cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type SBitfieldSat SCompleted

func (c SBitfieldSat) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldSet SCompleted

func (c SBitfieldSet) Incrby(Type string, Offset int64, Increment int64) SBitfieldIncrby {
	return SBitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitfieldSet) Wrap() SBitfieldWrap {
	return SBitfieldWrap{cs: append(c.cs, "WRAP"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldSet) Sat() SBitfieldSat {
	return SBitfieldSat{cs: append(c.cs, "SAT"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldSet) Fail() SBitfieldFail {
	return SBitfieldFail{cs: append(c.cs, "FAIL"), cf: c.cf, ks: c.ks}
}

func (c SBitfieldSet) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldWrap SCompleted

func (c SBitfieldWrap) Build() SCompleted {
	return SCompleted(c)
}

type SBitop SCompleted

func (c SBitop) Operation(Operation string) SBitopOperation {
	return SBitopOperation{cs: append(c.cs, Operation), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Bitop() (c SBitop) {
	c.cs = append(b.get(), "BITOP")
	c.ks = InitSlot
	return
}

type SBitopDestkey SCompleted

func (c SBitopDestkey) Key(Key ...string) SBitopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBitopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SBitopKey SCompleted

func (c SBitopKey) Key(Key ...string) SBitopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBitopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SBitopKey) Build() SCompleted {
	return SCompleted(c)
}

type SBitopOperation SCompleted

func (c SBitopOperation) Destkey(Destkey string) SBitopDestkey {
	c.ks = checkSlot(c.ks, slot(Destkey))
	return SBitopDestkey{cs: append(c.cs, Destkey), cf: c.cf, ks: c.ks}
}

type SBitpos SCompleted

func (c SBitpos) Key(Key string) SBitposKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SBitposKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Bitpos() (c SBitpos) {
	c.cs = append(b.get(), "BITPOS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SBitposBit SCompleted

func (c SBitposBit) Start(Start int64) SBitposIndexStart {
	return SBitposIndexStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitposBit) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitposBit) Cache() SCacheable {
	return SCacheable(c)
}

type SBitposIndexEnd SCompleted

func (c SBitposIndexEnd) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitposIndexEnd) Cache() SCacheable {
	return SCacheable(c)
}

type SBitposIndexStart SCompleted

func (c SBitposIndexStart) End(End int64) SBitposIndexEnd {
	return SBitposIndexEnd{cs: append(c.cs, strconv.FormatInt(End, 10)), cf: c.cf, ks: c.ks}
}

func (c SBitposIndexStart) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitposIndexStart) Cache() SCacheable {
	return SCacheable(c)
}

type SBitposKey SCompleted

func (c SBitposKey) Bit(Bit int64) SBitposBit {
	return SBitposBit{cs: append(c.cs, strconv.FormatInt(Bit, 10)), cf: c.cf, ks: c.ks}
}

type SBlmove SCompleted

func (c SBlmove) Source(Source string) SBlmoveSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SBlmoveSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Blmove() (c SBlmove) {
	c.cs = append(b.get(), "BLMOVE")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBlmoveDestination SCompleted

func (c SBlmoveDestination) Left() SBlmoveWherefromLeft {
	return SBlmoveWherefromLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SBlmoveDestination) Right() SBlmoveWherefromRight {
	return SBlmoveWherefromRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SBlmoveSource SCompleted

func (c SBlmoveSource) Destination(Destination string) SBlmoveDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SBlmoveDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type SBlmoveTimeout SCompleted

func (c SBlmoveTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBlmoveWherefromLeft SCompleted

func (c SBlmoveWherefromLeft) Left() SBlmoveWheretoLeft {
	return SBlmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SBlmoveWherefromLeft) Right() SBlmoveWheretoRight {
	return SBlmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SBlmoveWherefromRight SCompleted

func (c SBlmoveWherefromRight) Left() SBlmoveWheretoLeft {
	return SBlmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SBlmoveWherefromRight) Right() SBlmoveWheretoRight {
	return SBlmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SBlmoveWheretoLeft SCompleted

func (c SBlmoveWheretoLeft) Timeout(Timeout float64) SBlmoveTimeout {
	return SBlmoveTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SBlmoveWheretoRight SCompleted

func (c SBlmoveWheretoRight) Timeout(Timeout float64) SBlmoveTimeout {
	return SBlmoveTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SBlmpop SCompleted

func (c SBlmpop) Timeout(Timeout float64) SBlmpopTimeout {
	return SBlmpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Blmpop() (c SBlmpop) {
	c.cs = append(b.get(), "BLMPOP")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBlmpopCount SCompleted

func (c SBlmpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SBlmpopKey SCompleted

func (c SBlmpopKey) Left() SBlmpopWhereLeft {
	return SBlmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SBlmpopKey) Right() SBlmpopWhereRight {
	return SBlmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

func (c SBlmpopKey) Key(Key ...string) SBlmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SBlmpopNumkeys SCompleted

func (c SBlmpopNumkeys) Key(Key ...string) SBlmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SBlmpopNumkeys) Left() SBlmpopWhereLeft {
	return SBlmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SBlmpopNumkeys) Right() SBlmpopWhereRight {
	return SBlmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SBlmpopTimeout SCompleted

func (c SBlmpopTimeout) Numkeys(Numkeys int64) SBlmpopNumkeys {
	return SBlmpopNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SBlmpopWhereLeft SCompleted

func (c SBlmpopWhereLeft) Count(Count int64) SBlmpopCount {
	return SBlmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SBlmpopWhereLeft) Build() SCompleted {
	return SCompleted(c)
}

type SBlmpopWhereRight SCompleted

func (c SBlmpopWhereRight) Count(Count int64) SBlmpopCount {
	return SBlmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SBlmpopWhereRight) Build() SCompleted {
	return SCompleted(c)
}

type SBlpop SCompleted

func (c SBlpop) Key(Key ...string) SBlpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Blpop() (c SBlpop) {
	c.cs = append(b.get(), "BLPOP")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBlpopKey SCompleted

func (c SBlpopKey) Timeout(Timeout float64) SBlpopTimeout {
	return SBlpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SBlpopKey) Key(Key ...string) SBlpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SBlpopTimeout SCompleted

func (c SBlpopTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBrpop SCompleted

func (c SBrpop) Key(Key ...string) SBrpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBrpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Brpop() (c SBrpop) {
	c.cs = append(b.get(), "BRPOP")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBrpopKey SCompleted

func (c SBrpopKey) Timeout(Timeout float64) SBrpopTimeout {
	return SBrpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SBrpopKey) Key(Key ...string) SBrpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBrpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SBrpopTimeout SCompleted

func (c SBrpopTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBrpoplpush SCompleted

func (c SBrpoplpush) Source(Source string) SBrpoplpushSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SBrpoplpushSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Brpoplpush() (c SBrpoplpush) {
	c.cs = append(b.get(), "BRPOPLPUSH")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBrpoplpushDestination SCompleted

func (c SBrpoplpushDestination) Timeout(Timeout float64) SBrpoplpushTimeout {
	return SBrpoplpushTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SBrpoplpushSource SCompleted

func (c SBrpoplpushSource) Destination(Destination string) SBrpoplpushDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SBrpoplpushDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type SBrpoplpushTimeout SCompleted

func (c SBrpoplpushTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBzpopmax SCompleted

func (c SBzpopmax) Key(Key ...string) SBzpopmaxKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBzpopmaxKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Bzpopmax() (c SBzpopmax) {
	c.cs = append(b.get(), "BZPOPMAX")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBzpopmaxKey SCompleted

func (c SBzpopmaxKey) Timeout(Timeout float64) SBzpopmaxTimeout {
	return SBzpopmaxTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SBzpopmaxKey) Key(Key ...string) SBzpopmaxKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBzpopmaxKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SBzpopmaxTimeout SCompleted

func (c SBzpopmaxTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBzpopmin SCompleted

func (c SBzpopmin) Key(Key ...string) SBzpopminKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBzpopminKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Bzpopmin() (c SBzpopmin) {
	c.cs = append(b.get(), "BZPOPMIN")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBzpopminKey SCompleted

func (c SBzpopminKey) Timeout(Timeout float64) SBzpopminTimeout {
	return SBzpopminTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SBzpopminKey) Key(Key ...string) SBzpopminKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBzpopminKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SBzpopminTimeout SCompleted

func (c SBzpopminTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientCaching SCompleted

func (c SClientCaching) Yes() SClientCachingModeYes {
	return SClientCachingModeYes{cs: append(c.cs, "YES"), cf: c.cf, ks: c.ks}
}

func (c SClientCaching) No() SClientCachingModeNo {
	return SClientCachingModeNo{cs: append(c.cs, "NO"), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClientCaching() (c SClientCaching) {
	c.cs = append(b.get(), "CLIENT", "CACHING")
	c.ks = InitSlot
	return
}

type SClientCachingModeNo SCompleted

func (c SClientCachingModeNo) Build() SCompleted {
	return SCompleted(c)
}

type SClientCachingModeYes SCompleted

func (c SClientCachingModeYes) Build() SCompleted {
	return SCompleted(c)
}

type SClientGetname SCompleted

func (c SClientGetname) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientGetname() (c SClientGetname) {
	c.cs = append(b.get(), "CLIENT", "GETNAME")
	c.ks = InitSlot
	return
}

type SClientGetredir SCompleted

func (c SClientGetredir) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientGetredir() (c SClientGetredir) {
	c.cs = append(b.get(), "CLIENT", "GETREDIR")
	c.ks = InitSlot
	return
}

type SClientId SCompleted

func (c SClientId) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientId() (c SClientId) {
	c.cs = append(b.get(), "CLIENT", "ID")
	c.ks = InitSlot
	return
}

type SClientInfo SCompleted

func (c SClientInfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientInfo() (c SClientInfo) {
	c.cs = append(b.get(), "CLIENT", "INFO")
	c.ks = InitSlot
	return
}

type SClientKill SCompleted

func (c SClientKill) IpPort(IpPort string) SClientKillIpPort {
	return SClientKillIpPort{cs: append(c.cs, IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Id(ClientId int64) SClientKillId {
	return SClientKillId{cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Normal() SClientKillNormal {
	return SClientKillNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Master() SClientKillMaster {
	return SClientKillMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Slave() SClientKillSlave {
	return SClientKillSlave{cs: append(c.cs, "slave"), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Pubsub() SClientKillPubsub {
	return SClientKillPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c SClientKill) User(Username string) SClientKillUser {
	return SClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKill) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientKill() (c SClientKill) {
	c.cs = append(b.get(), "CLIENT", "KILL")
	c.ks = InitSlot
	return
}

type SClientKillAddr SCompleted

func (c SClientKillAddr) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillAddr) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillAddr) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillId SCompleted

func (c SClientKillId) Normal() SClientKillNormal {
	return SClientKillNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) Master() SClientKillMaster {
	return SClientKillMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) Slave() SClientKillSlave {
	return SClientKillSlave{cs: append(c.cs, "slave"), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) Pubsub() SClientKillPubsub {
	return SClientKillPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) User(Username string) SClientKillUser {
	return SClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillId) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillIpPort SCompleted

func (c SClientKillIpPort) Id(ClientId int64) SClientKillId {
	return SClientKillId{cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Normal() SClientKillNormal {
	return SClientKillNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Master() SClientKillMaster {
	return SClientKillMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Slave() SClientKillSlave {
	return SClientKillSlave{cs: append(c.cs, "slave"), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Pubsub() SClientKillPubsub {
	return SClientKillPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) User(Username string) SClientKillUser {
	return SClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillIpPort) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillLaddr SCompleted

func (c SClientKillLaddr) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillLaddr) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillMaster SCompleted

func (c SClientKillMaster) User(Username string) SClientKillUser {
	return SClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c SClientKillMaster) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillMaster) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillMaster) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillMaster) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillNormal SCompleted

func (c SClientKillNormal) User(Username string) SClientKillUser {
	return SClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c SClientKillNormal) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillNormal) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillNormal) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillNormal) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillPubsub SCompleted

func (c SClientKillPubsub) User(Username string) SClientKillUser {
	return SClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c SClientKillPubsub) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillPubsub) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillPubsub) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillPubsub) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillSkipme SCompleted

func (c SClientKillSkipme) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillSlave SCompleted

func (c SClientKillSlave) User(Username string) SClientKillUser {
	return SClientKillUser{cs: append(c.cs, "USER", Username), cf: c.cf, ks: c.ks}
}

func (c SClientKillSlave) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillSlave) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillSlave) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillSlave) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillUser SCompleted

func (c SClientKillUser) Addr(IpPort string) SClientKillAddr {
	return SClientKillAddr{cs: append(c.cs, "ADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillUser) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cs: append(c.cs, "LADDR", IpPort), cf: c.cf, ks: c.ks}
}

func (c SClientKillUser) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo), cf: c.cf, ks: c.ks}
}

func (c SClientKillUser) Build() SCompleted {
	return SCompleted(c)
}

type SClientList SCompleted

func (c SClientList) Normal() SClientListNormal {
	return SClientListNormal{cs: append(c.cs, "normal"), cf: c.cf, ks: c.ks}
}

func (c SClientList) Master() SClientListMaster {
	return SClientListMaster{cs: append(c.cs, "master"), cf: c.cf, ks: c.ks}
}

func (c SClientList) Replica() SClientListReplica {
	return SClientListReplica{cs: append(c.cs, "replica"), cf: c.cf, ks: c.ks}
}

func (c SClientList) Pubsub() SClientListPubsub {
	return SClientListPubsub{cs: append(c.cs, "pubsub"), cf: c.cf, ks: c.ks}
}

func (c SClientList) Id() SClientListIdId {
	return SClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c SClientList) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientList() (c SClientList) {
	c.cs = append(b.get(), "CLIENT", "LIST")
	c.ks = InitSlot
	return
}

type SClientListIdClientId SCompleted

func (c SClientListIdClientId) ClientId(ClientId ...int64) SClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClientListIdClientId{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SClientListIdClientId) Build() SCompleted {
	return SCompleted(c)
}

type SClientListIdId SCompleted

func (c SClientListIdId) ClientId(ClientId ...int64) SClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClientListIdClientId{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SClientListMaster SCompleted

func (c SClientListMaster) Id() SClientListIdId {
	return SClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c SClientListMaster) Build() SCompleted {
	return SCompleted(c)
}

type SClientListNormal SCompleted

func (c SClientListNormal) Id() SClientListIdId {
	return SClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c SClientListNormal) Build() SCompleted {
	return SCompleted(c)
}

type SClientListPubsub SCompleted

func (c SClientListPubsub) Id() SClientListIdId {
	return SClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c SClientListPubsub) Build() SCompleted {
	return SCompleted(c)
}

type SClientListReplica SCompleted

func (c SClientListReplica) Id() SClientListIdId {
	return SClientListIdId{cs: append(c.cs, "ID"), cf: c.cf, ks: c.ks}
}

func (c SClientListReplica) Build() SCompleted {
	return SCompleted(c)
}

type SClientNoEvict SCompleted

func (c SClientNoEvict) On() SClientNoEvictEnabledOn {
	return SClientNoEvictEnabledOn{cs: append(c.cs, "ON"), cf: c.cf, ks: c.ks}
}

func (c SClientNoEvict) Off() SClientNoEvictEnabledOff {
	return SClientNoEvictEnabledOff{cs: append(c.cs, "OFF"), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClientNoEvict() (c SClientNoEvict) {
	c.cs = append(b.get(), "CLIENT", "NO-EVICT")
	c.ks = InitSlot
	return
}

type SClientNoEvictEnabledOff SCompleted

func (c SClientNoEvictEnabledOff) Build() SCompleted {
	return SCompleted(c)
}

type SClientNoEvictEnabledOn SCompleted

func (c SClientNoEvictEnabledOn) Build() SCompleted {
	return SCompleted(c)
}

type SClientPause SCompleted

func (c SClientPause) Timeout(Timeout int64) SClientPauseTimeout {
	return SClientPauseTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClientPause() (c SClientPause) {
	c.cs = append(b.get(), "CLIENT", "PAUSE")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SClientPauseModeAll SCompleted

func (c SClientPauseModeAll) Build() SCompleted {
	return SCompleted(c)
}

type SClientPauseModeWrite SCompleted

func (c SClientPauseModeWrite) Build() SCompleted {
	return SCompleted(c)
}

type SClientPauseTimeout SCompleted

func (c SClientPauseTimeout) Write() SClientPauseModeWrite {
	return SClientPauseModeWrite{cs: append(c.cs, "WRITE"), cf: c.cf, ks: c.ks}
}

func (c SClientPauseTimeout) All() SClientPauseModeAll {
	return SClientPauseModeAll{cs: append(c.cs, "ALL"), cf: c.cf, ks: c.ks}
}

func (c SClientPauseTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientReply SCompleted

func (c SClientReply) On() SClientReplyReplyModeOn {
	return SClientReplyReplyModeOn{cs: append(c.cs, "ON"), cf: c.cf, ks: c.ks}
}

func (c SClientReply) Off() SClientReplyReplyModeOff {
	return SClientReplyReplyModeOff{cs: append(c.cs, "OFF"), cf: c.cf, ks: c.ks}
}

func (c SClientReply) Skip() SClientReplyReplyModeSkip {
	return SClientReplyReplyModeSkip{cs: append(c.cs, "SKIP"), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClientReply() (c SClientReply) {
	c.cs = append(b.get(), "CLIENT", "REPLY")
	c.ks = InitSlot
	return
}

type SClientReplyReplyModeOff SCompleted

func (c SClientReplyReplyModeOff) Build() SCompleted {
	return SCompleted(c)
}

type SClientReplyReplyModeOn SCompleted

func (c SClientReplyReplyModeOn) Build() SCompleted {
	return SCompleted(c)
}

type SClientReplyReplyModeSkip SCompleted

func (c SClientReplyReplyModeSkip) Build() SCompleted {
	return SCompleted(c)
}

type SClientSetname SCompleted

func (c SClientSetname) ConnectionName(ConnectionName string) SClientSetnameConnectionName {
	return SClientSetnameConnectionName{cs: append(c.cs, ConnectionName), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClientSetname() (c SClientSetname) {
	c.cs = append(b.get(), "CLIENT", "SETNAME")
	c.ks = InitSlot
	return
}

type SClientSetnameConnectionName SCompleted

func (c SClientSetnameConnectionName) Build() SCompleted {
	return SCompleted(c)
}

type SClientTracking SCompleted

func (c SClientTracking) On() SClientTrackingStatusOn {
	return SClientTrackingStatusOn{cs: append(c.cs, "ON"), cf: c.cf, ks: c.ks}
}

func (c SClientTracking) Off() SClientTrackingStatusOff {
	return SClientTrackingStatusOff{cs: append(c.cs, "OFF"), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClientTracking() (c SClientTracking) {
	c.cs = append(b.get(), "CLIENT", "TRACKING")
	c.ks = InitSlot
	return
}

type SClientTrackingBcastBcast SCompleted

func (c SClientTrackingBcastBcast) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingBcastBcast) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingBcastBcast) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingBcastBcast) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingNoloopNoloop SCompleted

func (c SClientTrackingNoloopNoloop) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingOptinOptin SCompleted

func (c SClientTrackingOptinOptin) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingOptinOptin) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingOptinOptin) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingOptoutOptout SCompleted

func (c SClientTrackingOptoutOptout) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingOptoutOptout) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingPrefix SCompleted

func (c SClientTrackingPrefix) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingPrefix) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingPrefix) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingPrefix) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingPrefix) Prefix(Prefix ...string) SClientTrackingPrefix {
	return SClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingPrefix) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingRedirect SCompleted

func (c SClientTrackingRedirect) Prefix(Prefix ...string) SClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return SClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingRedirect) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingRedirect) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingRedirect) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingRedirect) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingRedirect) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingStatusOff SCompleted

func (c SClientTrackingStatusOff) Redirect(ClientId int64) SClientTrackingRedirect {
	return SClientTrackingRedirect{cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOff) Prefix(Prefix ...string) SClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return SClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOff) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOff) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOff) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOff) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOff) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingStatusOn SCompleted

func (c SClientTrackingStatusOn) Redirect(ClientId int64) SClientTrackingRedirect {
	return SClientTrackingRedirect{cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOn) Prefix(Prefix ...string) SClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return SClientTrackingPrefix{cs: append(c.cs, Prefix...), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOn) Bcast() SClientTrackingBcastBcast {
	return SClientTrackingBcastBcast{cs: append(c.cs, "BCAST"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOn) Optin() SClientTrackingOptinOptin {
	return SClientTrackingOptinOptin{cs: append(c.cs, "OPTIN"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOn) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOn) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP"), cf: c.cf, ks: c.ks}
}

func (c SClientTrackingStatusOn) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackinginfo SCompleted

func (c SClientTrackinginfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientTrackinginfo() (c SClientTrackinginfo) {
	c.cs = append(b.get(), "CLIENT", "TRACKINGINFO")
	c.ks = InitSlot
	return
}

type SClientUnblock SCompleted

func (c SClientUnblock) ClientId(ClientId int64) SClientUnblockClientId {
	return SClientUnblockClientId{cs: append(c.cs, strconv.FormatInt(ClientId, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClientUnblock() (c SClientUnblock) {
	c.cs = append(b.get(), "CLIENT", "UNBLOCK")
	c.ks = InitSlot
	return
}

type SClientUnblockClientId SCompleted

func (c SClientUnblockClientId) Timeout() SClientUnblockUnblockTypeTimeout {
	return SClientUnblockUnblockTypeTimeout{cs: append(c.cs, "TIMEOUT"), cf: c.cf, ks: c.ks}
}

func (c SClientUnblockClientId) Error() SClientUnblockUnblockTypeError {
	return SClientUnblockUnblockTypeError{cs: append(c.cs, "ERROR"), cf: c.cf, ks: c.ks}
}

func (c SClientUnblockClientId) Build() SCompleted {
	return SCompleted(c)
}

type SClientUnblockUnblockTypeError SCompleted

func (c SClientUnblockUnblockTypeError) Build() SCompleted {
	return SCompleted(c)
}

type SClientUnblockUnblockTypeTimeout SCompleted

func (c SClientUnblockUnblockTypeTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientUnpause SCompleted

func (c SClientUnpause) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClientUnpause() (c SClientUnpause) {
	c.cs = append(b.get(), "CLIENT", "UNPAUSE")
	c.ks = InitSlot
	return
}

type SClusterAddslots SCompleted

func (c SClusterAddslots) Slot(Slot ...int64) SClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterAddslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterAddslots() (c SClusterAddslots) {
	c.cs = append(b.get(), "CLUSTER", "ADDSLOTS")
	c.ks = InitSlot
	return
}

type SClusterAddslotsSlot SCompleted

func (c SClusterAddslotsSlot) Slot(Slot ...int64) SClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterAddslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SClusterAddslotsSlot) Build() SCompleted {
	return SCompleted(c)
}

type SClusterBumpepoch SCompleted

func (c SClusterBumpepoch) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterBumpepoch() (c SClusterBumpepoch) {
	c.cs = append(b.get(), "CLUSTER", "BUMPEPOCH")
	c.ks = InitSlot
	return
}

type SClusterCountFailureReports SCompleted

func (c SClusterCountFailureReports) NodeId(NodeId string) SClusterCountFailureReportsNodeId {
	return SClusterCountFailureReportsNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterCountFailureReports() (c SClusterCountFailureReports) {
	c.cs = append(b.get(), "CLUSTER", "COUNT-FAILURE-REPORTS")
	c.ks = InitSlot
	return
}

type SClusterCountFailureReportsNodeId SCompleted

func (c SClusterCountFailureReportsNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterCountkeysinslot SCompleted

func (c SClusterCountkeysinslot) Slot(Slot int64) SClusterCountkeysinslotSlot {
	return SClusterCountkeysinslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterCountkeysinslot() (c SClusterCountkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "COUNTKEYSINSLOT")
	c.ks = InitSlot
	return
}

type SClusterCountkeysinslotSlot SCompleted

func (c SClusterCountkeysinslotSlot) Build() SCompleted {
	return SCompleted(c)
}

type SClusterDelslots SCompleted

func (c SClusterDelslots) Slot(Slot ...int64) SClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterDelslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterDelslots() (c SClusterDelslots) {
	c.cs = append(b.get(), "CLUSTER", "DELSLOTS")
	c.ks = InitSlot
	return
}

type SClusterDelslotsSlot SCompleted

func (c SClusterDelslotsSlot) Slot(Slot ...int64) SClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClusterDelslotsSlot{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SClusterDelslotsSlot) Build() SCompleted {
	return SCompleted(c)
}

type SClusterFailover SCompleted

func (c SClusterFailover) Force() SClusterFailoverOptionsForce {
	return SClusterFailoverOptionsForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c SClusterFailover) Takeover() SClusterFailoverOptionsTakeover {
	return SClusterFailoverOptionsTakeover{cs: append(c.cs, "TAKEOVER"), cf: c.cf, ks: c.ks}
}

func (c SClusterFailover) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterFailover() (c SClusterFailover) {
	c.cs = append(b.get(), "CLUSTER", "FAILOVER")
	c.ks = InitSlot
	return
}

type SClusterFailoverOptionsForce SCompleted

func (c SClusterFailoverOptionsForce) Build() SCompleted {
	return SCompleted(c)
}

type SClusterFailoverOptionsTakeover SCompleted

func (c SClusterFailoverOptionsTakeover) Build() SCompleted {
	return SCompleted(c)
}

type SClusterFlushslots SCompleted

func (c SClusterFlushslots) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterFlushslots() (c SClusterFlushslots) {
	c.cs = append(b.get(), "CLUSTER", "FLUSHSLOTS")
	c.ks = InitSlot
	return
}

type SClusterForget SCompleted

func (c SClusterForget) NodeId(NodeId string) SClusterForgetNodeId {
	return SClusterForgetNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterForget() (c SClusterForget) {
	c.cs = append(b.get(), "CLUSTER", "FORGET")
	c.ks = InitSlot
	return
}

type SClusterForgetNodeId SCompleted

func (c SClusterForgetNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterGetkeysinslot SCompleted

func (c SClusterGetkeysinslot) Slot(Slot int64) SClusterGetkeysinslotSlot {
	return SClusterGetkeysinslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterGetkeysinslot() (c SClusterGetkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "GETKEYSINSLOT")
	c.ks = InitSlot
	return
}

type SClusterGetkeysinslotCount SCompleted

func (c SClusterGetkeysinslotCount) Build() SCompleted {
	return SCompleted(c)
}

type SClusterGetkeysinslotSlot SCompleted

func (c SClusterGetkeysinslotSlot) Count(Count int64) SClusterGetkeysinslotCount {
	return SClusterGetkeysinslotCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

type SClusterInfo SCompleted

func (c SClusterInfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterInfo() (c SClusterInfo) {
	c.cs = append(b.get(), "CLUSTER", "INFO")
	c.ks = InitSlot
	return
}

type SClusterKeyslot SCompleted

func (c SClusterKeyslot) Key(Key string) SClusterKeyslotKey {
	return SClusterKeyslotKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterKeyslot() (c SClusterKeyslot) {
	c.cs = append(b.get(), "CLUSTER", "KEYSLOT")
	c.ks = InitSlot
	return
}

type SClusterKeyslotKey SCompleted

func (c SClusterKeyslotKey) Build() SCompleted {
	return SCompleted(c)
}

type SClusterMeet SCompleted

func (c SClusterMeet) Ip(Ip string) SClusterMeetIp {
	return SClusterMeetIp{cs: append(c.cs, Ip), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterMeet() (c SClusterMeet) {
	c.cs = append(b.get(), "CLUSTER", "MEET")
	c.ks = InitSlot
	return
}

type SClusterMeetIp SCompleted

func (c SClusterMeetIp) Port(Port int64) SClusterMeetPort {
	return SClusterMeetPort{cs: append(c.cs, strconv.FormatInt(Port, 10)), cf: c.cf, ks: c.ks}
}

type SClusterMeetPort SCompleted

func (c SClusterMeetPort) Build() SCompleted {
	return SCompleted(c)
}

type SClusterMyid SCompleted

func (c SClusterMyid) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterMyid() (c SClusterMyid) {
	c.cs = append(b.get(), "CLUSTER", "MYID")
	c.ks = InitSlot
	return
}

type SClusterNodes SCompleted

func (c SClusterNodes) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterNodes() (c SClusterNodes) {
	c.cs = append(b.get(), "CLUSTER", "NODES")
	c.ks = InitSlot
	return
}

type SClusterReplicas SCompleted

func (c SClusterReplicas) NodeId(NodeId string) SClusterReplicasNodeId {
	return SClusterReplicasNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterReplicas() (c SClusterReplicas) {
	c.cs = append(b.get(), "CLUSTER", "REPLICAS")
	c.ks = InitSlot
	return
}

type SClusterReplicasNodeId SCompleted

func (c SClusterReplicasNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterReplicate SCompleted

func (c SClusterReplicate) NodeId(NodeId string) SClusterReplicateNodeId {
	return SClusterReplicateNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterReplicate() (c SClusterReplicate) {
	c.cs = append(b.get(), "CLUSTER", "REPLICATE")
	c.ks = InitSlot
	return
}

type SClusterReplicateNodeId SCompleted

func (c SClusterReplicateNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterReset SCompleted

func (c SClusterReset) Hard() SClusterResetResetTypeHard {
	return SClusterResetResetTypeHard{cs: append(c.cs, "HARD"), cf: c.cf, ks: c.ks}
}

func (c SClusterReset) Soft() SClusterResetResetTypeSoft {
	return SClusterResetResetTypeSoft{cs: append(c.cs, "SOFT"), cf: c.cf, ks: c.ks}
}

func (c SClusterReset) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterReset() (c SClusterReset) {
	c.cs = append(b.get(), "CLUSTER", "RESET")
	c.ks = InitSlot
	return
}

type SClusterResetResetTypeHard SCompleted

func (c SClusterResetResetTypeHard) Build() SCompleted {
	return SCompleted(c)
}

type SClusterResetResetTypeSoft SCompleted

func (c SClusterResetResetTypeSoft) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSaveconfig SCompleted

func (c SClusterSaveconfig) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterSaveconfig() (c SClusterSaveconfig) {
	c.cs = append(b.get(), "CLUSTER", "SAVECONFIG")
	c.ks = InitSlot
	return
}

type SClusterSetConfigEpoch SCompleted

func (c SClusterSetConfigEpoch) ConfigEpoch(ConfigEpoch int64) SClusterSetConfigEpochConfigEpoch {
	return SClusterSetConfigEpochConfigEpoch{cs: append(c.cs, strconv.FormatInt(ConfigEpoch, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterSetConfigEpoch() (c SClusterSetConfigEpoch) {
	c.cs = append(b.get(), "CLUSTER", "SET-CONFIG-EPOCH")
	c.ks = InitSlot
	return
}

type SClusterSetConfigEpochConfigEpoch SCompleted

func (c SClusterSetConfigEpochConfigEpoch) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslot SCompleted

func (c SClusterSetslot) Slot(Slot int64) SClusterSetslotSlot {
	return SClusterSetslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterSetslot() (c SClusterSetslot) {
	c.cs = append(b.get(), "CLUSTER", "SETSLOT")
	c.ks = InitSlot
	return
}

type SClusterSetslotNodeId SCompleted

func (c SClusterSetslotNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSlot SCompleted

func (c SClusterSetslotSlot) Importing() SClusterSetslotSubcommandImporting {
	return SClusterSetslotSubcommandImporting{cs: append(c.cs, "IMPORTING"), cf: c.cf, ks: c.ks}
}

func (c SClusterSetslotSlot) Migrating() SClusterSetslotSubcommandMigrating {
	return SClusterSetslotSubcommandMigrating{cs: append(c.cs, "MIGRATING"), cf: c.cf, ks: c.ks}
}

func (c SClusterSetslotSlot) Stable() SClusterSetslotSubcommandStable {
	return SClusterSetslotSubcommandStable{cs: append(c.cs, "STABLE"), cf: c.cf, ks: c.ks}
}

func (c SClusterSetslotSlot) Node() SClusterSetslotSubcommandNode {
	return SClusterSetslotSubcommandNode{cs: append(c.cs, "NODE"), cf: c.cf, ks: c.ks}
}

type SClusterSetslotSubcommandImporting SCompleted

func (c SClusterSetslotSubcommandImporting) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c SClusterSetslotSubcommandImporting) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandMigrating SCompleted

func (c SClusterSetslotSubcommandMigrating) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c SClusterSetslotSubcommandMigrating) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandNode SCompleted

func (c SClusterSetslotSubcommandNode) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c SClusterSetslotSubcommandNode) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandStable SCompleted

func (c SClusterSetslotSubcommandStable) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (c SClusterSetslotSubcommandStable) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSlaves SCompleted

func (c SClusterSlaves) NodeId(NodeId string) SClusterSlavesNodeId {
	return SClusterSlavesNodeId{cs: append(c.cs, NodeId), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ClusterSlaves() (c SClusterSlaves) {
	c.cs = append(b.get(), "CLUSTER", "SLAVES")
	c.ks = InitSlot
	return
}

type SClusterSlavesNodeId SCompleted

func (c SClusterSlavesNodeId) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSlots SCompleted

func (c SClusterSlots) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ClusterSlots() (c SClusterSlots) {
	c.cs = append(b.get(), "CLUSTER", "SLOTS")
	c.ks = InitSlot
	return
}

type SCommand SCompleted

func (c SCommand) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Command() (c SCommand) {
	c.cs = append(b.get(), "COMMAND")
	c.ks = InitSlot
	return
}

type SCommandCount SCompleted

func (c SCommandCount) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) CommandCount() (c SCommandCount) {
	c.cs = append(b.get(), "COMMAND", "COUNT")
	c.ks = InitSlot
	return
}

type SCommandGetkeys SCompleted

func (c SCommandGetkeys) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) CommandGetkeys() (c SCommandGetkeys) {
	c.cs = append(b.get(), "COMMAND", "GETKEYS")
	c.ks = InitSlot
	return
}

type SCommandInfo SCompleted

func (c SCommandInfo) CommandName(CommandName ...string) SCommandInfoCommandName {
	return SCommandInfoCommandName{cs: append(c.cs, CommandName...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) CommandInfo() (c SCommandInfo) {
	c.cs = append(b.get(), "COMMAND", "INFO")
	c.ks = InitSlot
	return
}

type SCommandInfoCommandName SCompleted

func (c SCommandInfoCommandName) CommandName(CommandName ...string) SCommandInfoCommandName {
	return SCommandInfoCommandName{cs: append(c.cs, CommandName...), cf: c.cf, ks: c.ks}
}

func (c SCommandInfoCommandName) Build() SCompleted {
	return SCompleted(c)
}

type SConfigGet SCompleted

func (c SConfigGet) Parameter(Parameter string) SConfigGetParameter {
	return SConfigGetParameter{cs: append(c.cs, Parameter), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ConfigGet() (c SConfigGet) {
	c.cs = append(b.get(), "CONFIG", "GET")
	c.ks = InitSlot
	return
}

type SConfigGetParameter SCompleted

func (c SConfigGetParameter) Build() SCompleted {
	return SCompleted(c)
}

type SConfigResetstat SCompleted

func (c SConfigResetstat) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ConfigResetstat() (c SConfigResetstat) {
	c.cs = append(b.get(), "CONFIG", "RESETSTAT")
	c.ks = InitSlot
	return
}

type SConfigRewrite SCompleted

func (c SConfigRewrite) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ConfigRewrite() (c SConfigRewrite) {
	c.cs = append(b.get(), "CONFIG", "REWRITE")
	c.ks = InitSlot
	return
}

type SConfigSet SCompleted

func (c SConfigSet) Parameter(Parameter string) SConfigSetParameter {
	return SConfigSetParameter{cs: append(c.cs, Parameter), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ConfigSet() (c SConfigSet) {
	c.cs = append(b.get(), "CONFIG", "SET")
	c.ks = InitSlot
	return
}

type SConfigSetParameter SCompleted

func (c SConfigSetParameter) Value(Value string) SConfigSetValue {
	return SConfigSetValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SConfigSetValue SCompleted

func (c SConfigSetValue) Build() SCompleted {
	return SCompleted(c)
}

type SCopy SCompleted

func (c SCopy) Source(Source string) SCopySource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SCopySource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Copy() (c SCopy) {
	c.cs = append(b.get(), "COPY")
	c.ks = InitSlot
	return
}

type SCopyDb SCompleted

func (c SCopyDb) Replace() SCopyReplaceReplace {
	return SCopyReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c SCopyDb) Build() SCompleted {
	return SCompleted(c)
}

type SCopyDestination SCompleted

func (c SCopyDestination) Db(DestinationDb int64) SCopyDb {
	return SCopyDb{cs: append(c.cs, "DB", strconv.FormatInt(DestinationDb, 10)), cf: c.cf, ks: c.ks}
}

func (c SCopyDestination) Replace() SCopyReplaceReplace {
	return SCopyReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c SCopyDestination) Build() SCompleted {
	return SCompleted(c)
}

type SCopyReplaceReplace SCompleted

func (c SCopyReplaceReplace) Build() SCompleted {
	return SCompleted(c)
}

type SCopySource SCompleted

func (c SCopySource) Destination(Destination string) SCopyDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SCopyDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type SDbsize SCompleted

func (c SDbsize) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Dbsize() (c SDbsize) {
	c.cs = append(b.get(), "DBSIZE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SDebugObject SCompleted

func (c SDebugObject) Key(Key string) SDebugObjectKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SDebugObjectKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) DebugObject() (c SDebugObject) {
	c.cs = append(b.get(), "DEBUG", "OBJECT")
	c.ks = InitSlot
	return
}

type SDebugObjectKey SCompleted

func (c SDebugObjectKey) Build() SCompleted {
	return SCompleted(c)
}

type SDebugSegfault SCompleted

func (c SDebugSegfault) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) DebugSegfault() (c SDebugSegfault) {
	c.cs = append(b.get(), "DEBUG", "SEGFAULT")
	c.ks = InitSlot
	return
}

type SDecr SCompleted

func (c SDecr) Key(Key string) SDecrKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SDecrKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Decr() (c SDecr) {
	c.cs = append(b.get(), "DECR")
	c.ks = InitSlot
	return
}

type SDecrKey SCompleted

func (c SDecrKey) Build() SCompleted {
	return SCompleted(c)
}

type SDecrby SCompleted

func (c SDecrby) Key(Key string) SDecrbyKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SDecrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Decrby() (c SDecrby) {
	c.cs = append(b.get(), "DECRBY")
	c.ks = InitSlot
	return
}

type SDecrbyDecrement SCompleted

func (c SDecrbyDecrement) Build() SCompleted {
	return SCompleted(c)
}

type SDecrbyKey SCompleted

func (c SDecrbyKey) Decrement(Decrement int64) SDecrbyDecrement {
	return SDecrbyDecrement{cs: append(c.cs, strconv.FormatInt(Decrement, 10)), cf: c.cf, ks: c.ks}
}

type SDel SCompleted

func (c SDel) Key(Key ...string) SDelKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SDelKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Del() (c SDel) {
	c.cs = append(b.get(), "DEL")
	c.ks = InitSlot
	return
}

type SDelKey SCompleted

func (c SDelKey) Key(Key ...string) SDelKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SDelKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SDelKey) Build() SCompleted {
	return SCompleted(c)
}

type SDiscard SCompleted

func (c SDiscard) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Discard() (c SDiscard) {
	c.cs = append(b.get(), "DISCARD")
	c.ks = InitSlot
	return
}

type SDump SCompleted

func (c SDump) Key(Key string) SDumpKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SDumpKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Dump() (c SDump) {
	c.cs = append(b.get(), "DUMP")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SDumpKey SCompleted

func (c SDumpKey) Build() SCompleted {
	return SCompleted(c)
}

type SEcho SCompleted

func (c SEcho) Message(Message string) SEchoMessage {
	return SEchoMessage{cs: append(c.cs, Message), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Echo() (c SEcho) {
	c.cs = append(b.get(), "ECHO")
	c.ks = InitSlot
	return
}

type SEchoMessage SCompleted

func (c SEchoMessage) Build() SCompleted {
	return SCompleted(c)
}

type SEval SCompleted

func (c SEval) Script(Script string) SEvalScript {
	return SEvalScript{cs: append(c.cs, Script), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Eval() (c SEval) {
	c.cs = append(b.get(), "EVAL")
	c.ks = InitSlot
	return
}

type SEvalArg SCompleted

func (c SEvalArg) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalKey SCompleted

func (c SEvalKey) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalKey) Key(Key ...string) SEvalKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SEvalKey) Build() SCompleted {
	return SCompleted(c)
}

type SEvalNumkeys SCompleted

func (c SEvalNumkeys) Key(Key ...string) SEvalKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SEvalNumkeys) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalNumkeys) Build() SCompleted {
	return SCompleted(c)
}

type SEvalRo SCompleted

func (c SEvalRo) Script(Script string) SEvalRoScript {
	return SEvalRoScript{cs: append(c.cs, Script), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) EvalRo() (c SEvalRo) {
	c.cs = append(b.get(), "EVAL_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SEvalRoArg SCompleted

func (c SEvalRoArg) Arg(Arg ...string) SEvalRoArg {
	return SEvalRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalRoArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalRoKey SCompleted

func (c SEvalRoKey) Arg(Arg ...string) SEvalRoArg {
	return SEvalRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalRoKey) Key(Key ...string) SEvalRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SEvalRoNumkeys SCompleted

func (c SEvalRoNumkeys) Key(Key ...string) SEvalRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SEvalRoScript SCompleted

func (c SEvalRoScript) Numkeys(Numkeys int64) SEvalRoNumkeys {
	return SEvalRoNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SEvalScript SCompleted

func (c SEvalScript) Numkeys(Numkeys int64) SEvalNumkeys {
	return SEvalNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SEvalsha SCompleted

func (c SEvalsha) Sha1(Sha1 string) SEvalshaSha1 {
	return SEvalshaSha1{cs: append(c.cs, Sha1), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Evalsha() (c SEvalsha) {
	c.cs = append(b.get(), "EVALSHA")
	c.ks = InitSlot
	return
}

type SEvalshaArg SCompleted

func (c SEvalshaArg) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalshaArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaKey SCompleted

func (c SEvalshaKey) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalshaKey) Key(Key ...string) SEvalshaKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SEvalshaKey) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaNumkeys SCompleted

func (c SEvalshaNumkeys) Key(Key ...string) SEvalshaKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SEvalshaNumkeys) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalshaNumkeys) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaRo SCompleted

func (c SEvalshaRo) Sha1(Sha1 string) SEvalshaRoSha1 {
	return SEvalshaRoSha1{cs: append(c.cs, Sha1), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) EvalshaRo() (c SEvalshaRo) {
	c.cs = append(b.get(), "EVALSHA_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SEvalshaRoArg SCompleted

func (c SEvalshaRoArg) Arg(Arg ...string) SEvalshaRoArg {
	return SEvalshaRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalshaRoArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaRoKey SCompleted

func (c SEvalshaRoKey) Arg(Arg ...string) SEvalshaRoArg {
	return SEvalshaRoArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SEvalshaRoKey) Key(Key ...string) SEvalshaRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SEvalshaRoNumkeys SCompleted

func (c SEvalshaRoNumkeys) Key(Key ...string) SEvalshaRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaRoKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SEvalshaRoSha1 SCompleted

func (c SEvalshaRoSha1) Numkeys(Numkeys int64) SEvalshaRoNumkeys {
	return SEvalshaRoNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SEvalshaSha1 SCompleted

func (c SEvalshaSha1) Numkeys(Numkeys int64) SEvalshaNumkeys {
	return SEvalshaNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SExec SCompleted

func (c SExec) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Exec() (c SExec) {
	c.cs = append(b.get(), "EXEC")
	c.ks = InitSlot
	return
}

type SExists SCompleted

func (c SExists) Key(Key ...string) SExistsKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SExistsKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Exists() (c SExists) {
	c.cs = append(b.get(), "EXISTS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SExistsKey SCompleted

func (c SExistsKey) Key(Key ...string) SExistsKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SExistsKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SExistsKey) Build() SCompleted {
	return SCompleted(c)
}

type SExpire SCompleted

func (c SExpire) Key(Key string) SExpireKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SExpireKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Expire() (c SExpire) {
	c.cs = append(b.get(), "EXPIRE")
	c.ks = InitSlot
	return
}

type SExpireConditionGt SCompleted

func (c SExpireConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireConditionLt SCompleted

func (c SExpireConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireConditionNx SCompleted

func (c SExpireConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireConditionXx SCompleted

func (c SExpireConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireKey SCompleted

func (c SExpireKey) Seconds(Seconds int64) SExpireSeconds {
	return SExpireSeconds{cs: append(c.cs, strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

type SExpireSeconds SCompleted

func (c SExpireSeconds) Nx() SExpireConditionNx {
	return SExpireConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SExpireSeconds) Xx() SExpireConditionXx {
	return SExpireConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SExpireSeconds) Gt() SExpireConditionGt {
	return SExpireConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c SExpireSeconds) Lt() SExpireConditionLt {
	return SExpireConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c SExpireSeconds) Build() SCompleted {
	return SCompleted(c)
}

type SExpireat SCompleted

func (c SExpireat) Key(Key string) SExpireatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SExpireatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Expireat() (c SExpireat) {
	c.cs = append(b.get(), "EXPIREAT")
	c.ks = InitSlot
	return
}

type SExpireatConditionGt SCompleted

func (c SExpireatConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatConditionLt SCompleted

func (c SExpireatConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatConditionNx SCompleted

func (c SExpireatConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatConditionXx SCompleted

func (c SExpireatConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SExpireatKey SCompleted

func (c SExpireatKey) Timestamp(Timestamp int64) SExpireatTimestamp {
	return SExpireatTimestamp{cs: append(c.cs, strconv.FormatInt(Timestamp, 10)), cf: c.cf, ks: c.ks}
}

type SExpireatTimestamp SCompleted

func (c SExpireatTimestamp) Nx() SExpireatConditionNx {
	return SExpireatConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SExpireatTimestamp) Xx() SExpireatConditionXx {
	return SExpireatConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SExpireatTimestamp) Gt() SExpireatConditionGt {
	return SExpireatConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c SExpireatTimestamp) Lt() SExpireatConditionLt {
	return SExpireatConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c SExpireatTimestamp) Build() SCompleted {
	return SCompleted(c)
}

type SExpiretime SCompleted

func (c SExpiretime) Key(Key string) SExpiretimeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SExpiretimeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Expiretime() (c SExpiretime) {
	c.cs = append(b.get(), "EXPIRETIME")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SExpiretimeKey SCompleted

func (c SExpiretimeKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SExpiretimeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SFailover SCompleted

func (c SFailover) To() SFailoverTargetTo {
	return SFailoverTargetTo{cs: append(c.cs, "TO"), cf: c.cf, ks: c.ks}
}

func (c SFailover) Abort() SFailoverAbort {
	return SFailoverAbort{cs: append(c.cs, "ABORT"), cf: c.cf, ks: c.ks}
}

func (c SFailover) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SFailover) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Failover() (c SFailover) {
	c.cs = append(b.get(), "FAILOVER")
	c.ks = InitSlot
	return
}

type SFailoverAbort SCompleted

func (c SFailoverAbort) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SFailoverAbort) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetForce SCompleted

func (c SFailoverTargetForce) Abort() SFailoverAbort {
	return SFailoverAbort{cs: append(c.cs, "ABORT"), cf: c.cf, ks: c.ks}
}

func (c SFailoverTargetForce) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SFailoverTargetForce) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetHost SCompleted

func (c SFailoverTargetHost) Port(Port int64) SFailoverTargetPort {
	return SFailoverTargetPort{cs: append(c.cs, strconv.FormatInt(Port, 10)), cf: c.cf, ks: c.ks}
}

type SFailoverTargetPort SCompleted

func (c SFailoverTargetPort) Force() SFailoverTargetForce {
	return SFailoverTargetForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c SFailoverTargetPort) Abort() SFailoverAbort {
	return SFailoverAbort{cs: append(c.cs, "ABORT"), cf: c.cf, ks: c.ks}
}

func (c SFailoverTargetPort) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SFailoverTargetPort) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetTo SCompleted

func (c SFailoverTargetTo) Host(Host string) SFailoverTargetHost {
	return SFailoverTargetHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

type SFailoverTimeout SCompleted

func (c SFailoverTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SFlushall SCompleted

func (c SFlushall) Async() SFlushallAsyncAsync {
	return SFlushallAsyncAsync{cs: append(c.cs, "ASYNC"), cf: c.cf, ks: c.ks}
}

func (c SFlushall) Sync() SFlushallAsyncSync {
	return SFlushallAsyncSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c SFlushall) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Flushall() (c SFlushall) {
	c.cs = append(b.get(), "FLUSHALL")
	c.ks = InitSlot
	return
}

type SFlushallAsyncAsync SCompleted

func (c SFlushallAsyncAsync) Build() SCompleted {
	return SCompleted(c)
}

type SFlushallAsyncSync SCompleted

func (c SFlushallAsyncSync) Build() SCompleted {
	return SCompleted(c)
}

type SFlushdb SCompleted

func (c SFlushdb) Async() SFlushdbAsyncAsync {
	return SFlushdbAsyncAsync{cs: append(c.cs, "ASYNC"), cf: c.cf, ks: c.ks}
}

func (c SFlushdb) Sync() SFlushdbAsyncSync {
	return SFlushdbAsyncSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c SFlushdb) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Flushdb() (c SFlushdb) {
	c.cs = append(b.get(), "FLUSHDB")
	c.ks = InitSlot
	return
}

type SFlushdbAsyncAsync SCompleted

func (c SFlushdbAsyncAsync) Build() SCompleted {
	return SCompleted(c)
}

type SFlushdbAsyncSync SCompleted

func (c SFlushdbAsyncSync) Build() SCompleted {
	return SCompleted(c)
}

type SGeoadd SCompleted

func (c SGeoadd) Key(Key string) SGeoaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Geoadd() (c SGeoadd) {
	c.cs = append(b.get(), "GEOADD")
	c.ks = InitSlot
	return
}

type SGeoaddChangeCh SCompleted

func (c SGeoaddChangeCh) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SGeoaddConditionNx SCompleted

func (c SGeoaddConditionNx) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SGeoaddConditionNx) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SGeoaddConditionXx SCompleted

func (c SGeoaddConditionXx) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SGeoaddConditionXx) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SGeoaddKey SCompleted

func (c SGeoaddKey) Nx() SGeoaddConditionNx {
	return SGeoaddConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SGeoaddKey) Xx() SGeoaddConditionXx {
	return SGeoaddConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SGeoaddKey) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SGeoaddKey) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SGeoaddLongitudeLatitudeMember SCompleted

func (c SGeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member), cf: c.cf, ks: c.ks}
}

func (c SGeoaddLongitudeLatitudeMember) Build() SCompleted {
	return SCompleted(c)
}

type SGeodist SCompleted

func (c SGeodist) Key(Key string) SGeodistKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeodistKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Geodist() (c SGeodist) {
	c.cs = append(b.get(), "GEODIST")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeodistKey SCompleted

func (c SGeodistKey) Member1(Member1 string) SGeodistMember1 {
	return SGeodistMember1{cs: append(c.cs, Member1), cf: c.cf, ks: c.ks}
}

type SGeodistMember1 SCompleted

func (c SGeodistMember1) Member2(Member2 string) SGeodistMember2 {
	return SGeodistMember2{cs: append(c.cs, Member2), cf: c.cf, ks: c.ks}
}

type SGeodistMember2 SCompleted

func (c SGeodistMember2) M() SGeodistUnitM {
	return SGeodistUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeodistMember2) Km() SGeodistUnitKm {
	return SGeodistUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeodistMember2) Ft() SGeodistUnitFt {
	return SGeodistUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeodistMember2) Mi() SGeodistUnitMi {
	return SGeodistUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

func (c SGeodistMember2) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistMember2) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitFt SCompleted

func (c SGeodistUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitKm SCompleted

func (c SGeodistUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitM SCompleted

func (c SGeodistUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeodistUnitMi SCompleted

func (c SGeodistUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeodistUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeohash SCompleted

func (c SGeohash) Key(Key string) SGeohashKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeohashKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Geohash() (c SGeohash) {
	c.cs = append(b.get(), "GEOHASH")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeohashKey SCompleted

func (c SGeohashKey) Member(Member ...string) SGeohashMember {
	return SGeohashMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SGeohashMember SCompleted

func (c SGeohashMember) Member(Member ...string) SGeohashMember {
	return SGeohashMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SGeohashMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeohashMember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeopos SCompleted

func (c SGeopos) Key(Key string) SGeoposKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoposKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Geopos() (c SGeopos) {
	c.cs = append(b.get(), "GEOPOS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeoposKey SCompleted

func (c SGeoposKey) Member(Member ...string) SGeoposMember {
	return SGeoposMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SGeoposMember SCompleted

func (c SGeoposMember) Member(Member ...string) SGeoposMember {
	return SGeoposMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SGeoposMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoposMember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradius SCompleted

func (c SGeoradius) Key(Key string) SGeoradiusKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Georadius() (c SGeoradius) {
	c.cs = append(b.get(), "GEORADIUS")
	c.ks = InitSlot
	return
}

type SGeoradiusCountAnyAny SCompleted

func (c SGeoradiusCountAnyAny) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountAnyAny) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountAnyAny) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountAnyAny) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusCountCount SCompleted

func (c SGeoradiusCountCount) Any() SGeoradiusCountAnyAny {
	return SGeoradiusCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountCount) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountCount) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountCount) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountCount) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusKey SCompleted

func (c SGeoradiusKey) Longitude(Longitude float64) SGeoradiusLongitude {
	return SGeoradiusLongitude{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusLatitude SCompleted

func (c SGeoradiusLatitude) Radius(Radius float64) SGeoradiusRadius {
	return SGeoradiusRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusLongitude SCompleted

func (c SGeoradiusLongitude) Latitude(Latitude float64) SGeoradiusLatitude {
	return SGeoradiusLatitude{cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusOrderAsc SCompleted

func (c SGeoradiusOrderAsc) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusOrderAsc) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusOrderDesc SCompleted

func (c SGeoradiusOrderDesc) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusOrderDesc) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusRadius SCompleted

func (c SGeoradiusRadius) M() SGeoradiusUnitM {
	return SGeoradiusUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRadius) Km() SGeoradiusUnitKm {
	return SGeoradiusUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRadius) Ft() SGeoradiusUnitFt {
	return SGeoradiusUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRadius) Mi() SGeoradiusUnitMi {
	return SGeoradiusUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeoradiusRo SCompleted

func (c SGeoradiusRo) Key(Key string) SGeoradiusRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) GeoradiusRo() (c SGeoradiusRo) {
	c.cs = append(b.get(), "GEORADIUS_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeoradiusRoCountAnyAny SCompleted

func (c SGeoradiusRoCountAnyAny) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoCountAnyAny) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoCountAnyAny) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoCountCount SCompleted

func (c SGeoradiusRoCountCount) Any() SGeoradiusRoCountAnyAny {
	return SGeoradiusRoCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoCountCount) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoCountCount) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoCountCount) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoKey SCompleted

func (c SGeoradiusRoKey) Longitude(Longitude float64) SGeoradiusRoLongitude {
	return SGeoradiusRoLongitude{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusRoLatitude SCompleted

func (c SGeoradiusRoLatitude) Radius(Radius float64) SGeoradiusRoRadius {
	return SGeoradiusRoRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusRoLongitude SCompleted

func (c SGeoradiusRoLongitude) Latitude(Latitude float64) SGeoradiusRoLatitude {
	return SGeoradiusRoLatitude{cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusRoOrderAsc SCompleted

func (c SGeoradiusRoOrderAsc) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoOrderDesc SCompleted

func (c SGeoradiusRoOrderDesc) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoRadius SCompleted

func (c SGeoradiusRoRadius) M() SGeoradiusRoUnitM {
	return SGeoradiusRoUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoRadius) Km() SGeoradiusRoUnitKm {
	return SGeoradiusRoUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoRadius) Ft() SGeoradiusRoUnitFt {
	return SGeoradiusRoUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoRadius) Mi() SGeoradiusRoUnitMi {
	return SGeoradiusRoUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeoradiusRoStoredist SCompleted

func (c SGeoradiusRoStoredist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoStoredist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitFt SCompleted

func (c SGeoradiusRoUnitFt) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitFt) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitFt) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitFt) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitFt) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitFt) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitFt) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitKm SCompleted

func (c SGeoradiusRoUnitKm) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitKm) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitKm) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitKm) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitKm) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitKm) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitKm) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitM SCompleted

func (c SGeoradiusRoUnitM) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitM) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitM) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitM) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitM) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitM) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitM) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitMi SCompleted

func (c SGeoradiusRoUnitMi) Withcoord() SGeoradiusRoWithcoordWithcoord {
	return SGeoradiusRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitMi) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitMi) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitMi) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitMi) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitMi) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitMi) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithcoordWithcoord SCompleted

func (c SGeoradiusRoWithcoordWithcoord) Withdist() SGeoradiusRoWithdistWithdist {
	return SGeoradiusRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithcoordWithcoord) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithcoordWithcoord) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithcoordWithcoord) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithcoordWithcoord) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithcoordWithcoord) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithdistWithdist SCompleted

func (c SGeoradiusRoWithdistWithdist) Withhash() SGeoradiusRoWithhashWithhash {
	return SGeoradiusRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithdistWithdist) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithdistWithdist) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithdistWithdist) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithdistWithdist) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithhashWithhash SCompleted

func (c SGeoradiusRoWithhashWithhash) Count(Count int64) SGeoradiusRoCountCount {
	return SGeoradiusRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithhashWithhash) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithhashWithhash) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithhashWithhash) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusRoWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoWithhashWithhash) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusStore SCompleted

func (c SGeoradiusStore) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusStore) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusStoredist SCompleted

func (c SGeoradiusStoredist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitFt SCompleted

func (c SGeoradiusUnitFt) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitKm SCompleted

func (c SGeoradiusUnitKm) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitM SCompleted

func (c SGeoradiusUnitM) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitMi SCompleted

func (c SGeoradiusUnitMi) Withcoord() SGeoradiusWithcoordWithcoord {
	return SGeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusWithcoordWithcoord SCompleted

func (c SGeoradiusWithcoordWithcoord) Withdist() SGeoradiusWithdistWithdist {
	return SGeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithcoordWithcoord) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithcoordWithcoord) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithcoordWithcoord) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithcoordWithcoord) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithcoordWithcoord) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithcoordWithcoord) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusWithdistWithdist SCompleted

func (c SGeoradiusWithdistWithdist) Withhash() SGeoradiusWithhashWithhash {
	return SGeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithdistWithdist) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithdistWithdist) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithdistWithdist) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithdistWithdist) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithdistWithdist) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusWithhashWithhash SCompleted

func (c SGeoradiusWithhashWithhash) Count(Count int64) SGeoradiusCountCount {
	return SGeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithhashWithhash) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithhashWithhash) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithhashWithhash) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithhashWithhash) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymember SCompleted

func (c SGeoradiusbymember) Key(Key string) SGeoradiusbymemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Georadiusbymember() (c SGeoradiusbymember) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER")
	c.ks = InitSlot
	return
}

type SGeoradiusbymemberCountAnyAny SCompleted

func (c SGeoradiusbymemberCountAnyAny) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountAnyAny) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountAnyAny) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountAnyAny) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberCountCount SCompleted

func (c SGeoradiusbymemberCountCount) Any() SGeoradiusbymemberCountAnyAny {
	return SGeoradiusbymemberCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountCount) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountCount) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountCount) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountCount) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberKey SCompleted

func (c SGeoradiusbymemberKey) Member(Member string) SGeoradiusbymemberMember {
	return SGeoradiusbymemberMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SGeoradiusbymemberMember SCompleted

func (c SGeoradiusbymemberMember) Radius(Radius float64) SGeoradiusbymemberRadius {
	return SGeoradiusbymemberRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusbymemberOrderAsc SCompleted

func (c SGeoradiusbymemberOrderAsc) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberOrderAsc) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberOrderDesc SCompleted

func (c SGeoradiusbymemberOrderDesc) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberOrderDesc) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberRadius SCompleted

func (c SGeoradiusbymemberRadius) M() SGeoradiusbymemberUnitM {
	return SGeoradiusbymemberUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRadius) Km() SGeoradiusbymemberUnitKm {
	return SGeoradiusbymemberUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRadius) Ft() SGeoradiusbymemberUnitFt {
	return SGeoradiusbymemberUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRadius) Mi() SGeoradiusbymemberUnitMi {
	return SGeoradiusbymemberUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeoradiusbymemberRo SCompleted

func (c SGeoradiusbymemberRo) Key(Key string) SGeoradiusbymemberRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) GeoradiusbymemberRo() (c SGeoradiusbymemberRo) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeoradiusbymemberRoCountAnyAny SCompleted

func (c SGeoradiusbymemberRoCountAnyAny) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoCountAnyAny) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoCountAnyAny) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoCountCount SCompleted

func (c SGeoradiusbymemberRoCountCount) Any() SGeoradiusbymemberRoCountAnyAny {
	return SGeoradiusbymemberRoCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoCountCount) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoCountCount) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoCountCount) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoKey SCompleted

func (c SGeoradiusbymemberRoKey) Member(Member string) SGeoradiusbymemberRoMember {
	return SGeoradiusbymemberRoMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SGeoradiusbymemberRoMember SCompleted

func (c SGeoradiusbymemberRoMember) Radius(Radius float64) SGeoradiusbymemberRoRadius {
	return SGeoradiusbymemberRoRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeoradiusbymemberRoOrderAsc SCompleted

func (c SGeoradiusbymemberRoOrderAsc) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoOrderDesc SCompleted

func (c SGeoradiusbymemberRoOrderDesc) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoRadius SCompleted

func (c SGeoradiusbymemberRoRadius) M() SGeoradiusbymemberRoUnitM {
	return SGeoradiusbymemberRoUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoRadius) Km() SGeoradiusbymemberRoUnitKm {
	return SGeoradiusbymemberRoUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoRadius) Ft() SGeoradiusbymemberRoUnitFt {
	return SGeoradiusbymemberRoUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoRadius) Mi() SGeoradiusbymemberRoUnitMi {
	return SGeoradiusbymemberRoUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeoradiusbymemberRoStoredist SCompleted

func (c SGeoradiusbymemberRoStoredist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoStoredist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitFt SCompleted

func (c SGeoradiusbymemberRoUnitFt) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitFt) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitFt) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitFt) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitFt) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitFt) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitFt) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitKm SCompleted

func (c SGeoradiusbymemberRoUnitKm) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitKm) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitKm) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitKm) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitKm) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitKm) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitKm) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitM SCompleted

func (c SGeoradiusbymemberRoUnitM) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitM) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitM) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitM) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitM) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitM) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitM) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitMi SCompleted

func (c SGeoradiusbymemberRoUnitMi) Withcoord() SGeoradiusbymemberRoWithcoordWithcoord {
	return SGeoradiusbymemberRoWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitMi) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitMi) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitMi) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitMi) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitMi) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitMi) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithcoordWithcoord SCompleted

func (c SGeoradiusbymemberRoWithcoordWithcoord) Withdist() SGeoradiusbymemberRoWithdistWithdist {
	return SGeoradiusbymemberRoWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithdistWithdist SCompleted

func (c SGeoradiusbymemberRoWithdistWithdist) Withhash() SGeoradiusbymemberRoWithhashWithhash {
	return SGeoradiusbymemberRoWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithhashWithhash SCompleted

func (c SGeoradiusbymemberRoWithhashWithhash) Count(Count int64) SGeoradiusbymemberRoCountCount {
	return SGeoradiusbymemberRoCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberRoWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoWithhashWithhash) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberStore SCompleted

func (c SGeoradiusbymemberStore) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberStore) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberStoredist SCompleted

func (c SGeoradiusbymemberStoredist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitFt SCompleted

func (c SGeoradiusbymemberUnitFt) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitKm SCompleted

func (c SGeoradiusbymemberUnitKm) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitM SCompleted

func (c SGeoradiusbymemberUnitM) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitMi SCompleted

func (c SGeoradiusbymemberUnitMi) Withcoord() SGeoradiusbymemberWithcoordWithcoord {
	return SGeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberWithcoordWithcoord SCompleted

func (c SGeoradiusbymemberWithcoordWithcoord) Withdist() SGeoradiusbymemberWithdistWithdist {
	return SGeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberWithdistWithdist SCompleted

func (c SGeoradiusbymemberWithdistWithdist) Withhash() SGeoradiusbymemberWithhashWithhash {
	return SGeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithdistWithdist) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithdistWithdist) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithdistWithdist) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithdistWithdist) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithdistWithdist) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberWithhashWithhash SCompleted

func (c SGeoradiusbymemberWithhashWithhash) Count(Count int64) SGeoradiusbymemberCountCount {
	return SGeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithhashWithhash) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithhashWithhash) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithhashWithhash) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cs: append(c.cs, "STORE", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithhashWithhash) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key), cf: c.cf, ks: c.ks}
}

func (c SGeoradiusbymemberWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearch SCompleted

func (c SGeosearch) Key(Key string) SGeosearchKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeosearchKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Geosearch() (c SGeosearch) {
	c.cs = append(b.get(), "GEOSEARCH")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeosearchBoxBybox SCompleted

func (c SGeosearchBoxBybox) Height(Height float64) SGeosearchBoxHeight {
	return SGeosearchBoxHeight{cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeosearchBoxHeight SCompleted

func (c SGeosearchBoxHeight) M() SGeosearchBoxUnitM {
	return SGeosearchBoxUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxHeight) Km() SGeosearchBoxUnitKm {
	return SGeosearchBoxUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxHeight) Ft() SGeosearchBoxUnitFt {
	return SGeosearchBoxUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxHeight) Mi() SGeosearchBoxUnitMi {
	return SGeosearchBoxUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeosearchBoxUnitFt SCompleted

func (c SGeosearchBoxUnitFt) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitFt) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitFt) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitFt) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitFt) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitFt) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitKm SCompleted

func (c SGeosearchBoxUnitKm) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitKm) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitKm) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitKm) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitKm) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitKm) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitM SCompleted

func (c SGeosearchBoxUnitM) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitM) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitM) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitM) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitM) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitM) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitMi SCompleted

func (c SGeosearchBoxUnitMi) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitMi) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitMi) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitMi) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitMi) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitMi) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchBoxUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleByradius SCompleted

func (c SGeosearchCircleByradius) M() SGeosearchCircleUnitM {
	return SGeosearchCircleUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleByradius) Km() SGeosearchCircleUnitKm {
	return SGeosearchCircleUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleByradius) Ft() SGeosearchCircleUnitFt {
	return SGeosearchCircleUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleByradius) Mi() SGeosearchCircleUnitMi {
	return SGeosearchCircleUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeosearchCircleUnitFt SCompleted

func (c SGeosearchCircleUnitFt) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitFt) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitFt) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitFt) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitFt) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitFt) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitFt) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitKm SCompleted

func (c SGeosearchCircleUnitKm) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitKm) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitKm) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitKm) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitKm) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitKm) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitKm) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitM SCompleted

func (c SGeosearchCircleUnitM) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitM) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitM) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitM) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitM) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitM) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitM) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitMi SCompleted

func (c SGeosearchCircleUnitMi) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitMi) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitMi) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitMi) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitMi) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitMi) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitMi) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCircleUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCountAnyAny SCompleted

func (c SGeosearchCountAnyAny) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCountAnyAny) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCountAnyAny) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCountCount SCompleted

func (c SGeosearchCountCount) Any() SGeosearchCountAnyAny {
	return SGeosearchCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCountCount) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCountCount) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCountCount) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchFromlonlat SCompleted

func (c SGeosearchFromlonlat) Byradius(Radius float64) SGeosearchCircleByradius {
	return SGeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFromlonlat) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchFromlonlat) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchFrommember SCompleted

func (c SGeosearchFrommember) Fromlonlat(Longitude float64, Latitude float64) SGeosearchFromlonlat {
	return SGeosearchFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Byradius(Radius float64) SGeosearchCircleByradius {
	return SGeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchFrommember) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchFrommember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchKey SCompleted

func (c SGeosearchKey) Frommember(Member string) SGeosearchFrommember {
	return SGeosearchFrommember{cs: append(c.cs, "FROMMEMBER", Member), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Fromlonlat(Longitude float64, Latitude float64) SGeosearchFromlonlat {
	return SGeosearchFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Byradius(Radius float64) SGeosearchCircleByradius {
	return SGeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Bybox(Width float64) SGeosearchBoxBybox {
	return SGeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Asc() SGeosearchOrderAsc {
	return SGeosearchOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Desc() SGeosearchOrderDesc {
	return SGeosearchOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchOrderAsc SCompleted

func (c SGeosearchOrderAsc) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderAsc) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderAsc) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderAsc) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchOrderDesc SCompleted

func (c SGeosearchOrderDesc) Count(Count int64) SGeosearchCountCount {
	return SGeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderDesc) Withcoord() SGeosearchWithcoordWithcoord {
	return SGeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderDesc) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderDesc) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchWithcoordWithcoord SCompleted

func (c SGeosearchWithcoordWithcoord) Withdist() SGeosearchWithdistWithdist {
	return SGeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchWithcoordWithcoord) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchWithdistWithdist SCompleted

func (c SGeosearchWithdistWithdist) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchWithhashWithhash SCompleted

func (c SGeosearchWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchWithhashWithhash) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchstore SCompleted

func (c SGeosearchstore) Destination(Destination string) SGeosearchstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SGeosearchstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Geosearchstore() (c SGeosearchstore) {
	c.cs = append(b.get(), "GEOSEARCHSTORE")
	c.ks = InitSlot
	return
}

type SGeosearchstoreBoxBybox SCompleted

func (c SGeosearchstoreBoxBybox) Height(Height float64) SGeosearchstoreBoxHeight {
	return SGeosearchstoreBoxHeight{cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SGeosearchstoreBoxHeight SCompleted

func (c SGeosearchstoreBoxHeight) M() SGeosearchstoreBoxUnitM {
	return SGeosearchstoreBoxUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxHeight) Km() SGeosearchstoreBoxUnitKm {
	return SGeosearchstoreBoxUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxHeight) Ft() SGeosearchstoreBoxUnitFt {
	return SGeosearchstoreBoxUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxHeight) Mi() SGeosearchstoreBoxUnitMi {
	return SGeosearchstoreBoxUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeosearchstoreBoxUnitFt SCompleted

func (c SGeosearchstoreBoxUnitFt) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitFt) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitFt) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitFt) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreBoxUnitKm SCompleted

func (c SGeosearchstoreBoxUnitKm) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitKm) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitKm) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitKm) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreBoxUnitM SCompleted

func (c SGeosearchstoreBoxUnitM) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitM) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitM) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitM) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreBoxUnitMi SCompleted

func (c SGeosearchstoreBoxUnitMi) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitMi) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitMi) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitMi) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreBoxUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleByradius SCompleted

func (c SGeosearchstoreCircleByradius) M() SGeosearchstoreCircleUnitM {
	return SGeosearchstoreCircleUnitM{cs: append(c.cs, "m"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleByradius) Km() SGeosearchstoreCircleUnitKm {
	return SGeosearchstoreCircleUnitKm{cs: append(c.cs, "km"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleByradius) Ft() SGeosearchstoreCircleUnitFt {
	return SGeosearchstoreCircleUnitFt{cs: append(c.cs, "ft"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleByradius) Mi() SGeosearchstoreCircleUnitMi {
	return SGeosearchstoreCircleUnitMi{cs: append(c.cs, "mi"), cf: c.cf, ks: c.ks}
}

type SGeosearchstoreCircleUnitFt SCompleted

func (c SGeosearchstoreCircleUnitFt) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitFt) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitFt) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitFt) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitFt) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleUnitKm SCompleted

func (c SGeosearchstoreCircleUnitKm) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitKm) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitKm) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitKm) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitKm) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleUnitM SCompleted

func (c SGeosearchstoreCircleUnitM) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitM) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitM) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitM) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitM) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleUnitMi SCompleted

func (c SGeosearchstoreCircleUnitMi) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitMi) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitMi) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitMi) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitMi) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCircleUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCountAnyAny SCompleted

func (c SGeosearchstoreCountAnyAny) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCountCount SCompleted

func (c SGeosearchstoreCountCount) Any() SGeosearchstoreCountAnyAny {
	return SGeosearchstoreCountAnyAny{cs: append(c.cs, "ANY"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCountCount) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreDestination SCompleted

func (c SGeosearchstoreDestination) Source(Source string) SGeosearchstoreSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SGeosearchstoreSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

type SGeosearchstoreFromlonlat SCompleted

func (c SGeosearchstoreFromlonlat) Byradius(Radius float64) SGeosearchstoreCircleByradius {
	return SGeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFromlonlat) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFromlonlat) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFromlonlat) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFromlonlat) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFromlonlat) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFromlonlat) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreFrommember SCompleted

func (c SGeosearchstoreFrommember) Fromlonlat(Longitude float64, Latitude float64) SGeosearchstoreFromlonlat {
	return SGeosearchstoreFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFrommember) Byradius(Radius float64) SGeosearchstoreCircleByradius {
	return SGeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFrommember) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFrommember) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFrommember) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFrommember) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFrommember) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreFrommember) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreOrderAsc SCompleted

func (c SGeosearchstoreOrderAsc) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreOrderAsc) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreOrderDesc SCompleted

func (c SGeosearchstoreOrderDesc) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreOrderDesc) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreSource SCompleted

func (c SGeosearchstoreSource) Frommember(Member string) SGeosearchstoreFrommember {
	return SGeosearchstoreFrommember{cs: append(c.cs, "FROMMEMBER", Member), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Fromlonlat(Longitude float64, Latitude float64) SGeosearchstoreFromlonlat {
	return SGeosearchstoreFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Byradius(Radius float64) SGeosearchstoreCircleByradius {
	return SGeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Bybox(Width float64) SGeosearchstoreBoxBybox {
	return SGeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Asc() SGeosearchstoreOrderAsc {
	return SGeosearchstoreOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Desc() SGeosearchstoreOrderDesc {
	return SGeosearchstoreOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST"), cf: c.cf, ks: c.ks}
}

func (c SGeosearchstoreSource) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreStoredistStoredist SCompleted

func (c SGeosearchstoreStoredistStoredist) Build() SCompleted {
	return SCompleted(c)
}

type SGet SCompleted

func (c SGet) Key(Key string) SGetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Get() (c SGet) {
	c.cs = append(b.get(), "GET")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGetKey SCompleted

func (c SGetKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SGetKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGetbit SCompleted

func (c SGetbit) Key(Key string) SGetbitKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetbitKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Getbit() (c SGetbit) {
	c.cs = append(b.get(), "GETBIT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGetbitKey SCompleted

func (c SGetbitKey) Offset(Offset int64) SGetbitOffset {
	return SGetbitOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type SGetbitOffset SCompleted

func (c SGetbitOffset) Build() SCompleted {
	return SCompleted(c)
}

func (c SGetbitOffset) Cache() SCacheable {
	return SCacheable(c)
}

type SGetdel SCompleted

func (c SGetdel) Key(Key string) SGetdelKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetdelKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Getdel() (c SGetdel) {
	c.cs = append(b.get(), "GETDEL")
	c.ks = InitSlot
	return
}

type SGetdelKey SCompleted

func (c SGetdelKey) Build() SCompleted {
	return SCompleted(c)
}

type SGetex SCompleted

func (c SGetex) Key(Key string) SGetexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Getex() (c SGetex) {
	c.cs = append(b.get(), "GETEX")
	c.ks = InitSlot
	return
}

type SGetexExpirationEx SCompleted

func (c SGetexExpirationEx) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationExat SCompleted

func (c SGetexExpirationExat) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationPersist SCompleted

func (c SGetexExpirationPersist) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationPx SCompleted

func (c SGetexExpirationPx) Build() SCompleted {
	return SCompleted(c)
}

type SGetexExpirationPxat SCompleted

func (c SGetexExpirationPxat) Build() SCompleted {
	return SCompleted(c)
}

type SGetexKey SCompleted

func (c SGetexKey) Ex(Seconds int64) SGetexExpirationEx {
	return SGetexExpirationEx{cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SGetexKey) Px(Milliseconds int64) SGetexExpirationPx {
	return SGetexExpirationPx{cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SGetexKey) Exat(Timestamp int64) SGetexExpirationExat {
	return SGetexExpirationExat{cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c SGetexKey) Pxat(Millisecondstimestamp int64) SGetexExpirationPxat {
	return SGetexExpirationPxat{cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c SGetexKey) Persist() SGetexExpirationPersist {
	return SGetexExpirationPersist{cs: append(c.cs, "PERSIST"), cf: c.cf, ks: c.ks}
}

func (c SGetexKey) Build() SCompleted {
	return SCompleted(c)
}

type SGetrange SCompleted

func (c SGetrange) Key(Key string) SGetrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Getrange() (c SGetrange) {
	c.cs = append(b.get(), "GETRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGetrangeEnd SCompleted

func (c SGetrangeEnd) Build() SCompleted {
	return SCompleted(c)
}

func (c SGetrangeEnd) Cache() SCacheable {
	return SCacheable(c)
}

type SGetrangeKey SCompleted

func (c SGetrangeKey) Start(Start int64) SGetrangeStart {
	return SGetrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type SGetrangeStart SCompleted

func (c SGetrangeStart) End(End int64) SGetrangeEnd {
	return SGetrangeEnd{cs: append(c.cs, strconv.FormatInt(End, 10)), cf: c.cf, ks: c.ks}
}

type SGetset SCompleted

func (c SGetset) Key(Key string) SGetsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Getset() (c SGetset) {
	c.cs = append(b.get(), "GETSET")
	c.ks = InitSlot
	return
}

type SGetsetKey SCompleted

func (c SGetsetKey) Value(Value string) SGetsetValue {
	return SGetsetValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SGetsetValue SCompleted

func (c SGetsetValue) Build() SCompleted {
	return SCompleted(c)
}

type SHdel SCompleted

func (c SHdel) Key(Key string) SHdelKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHdelKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hdel() (c SHdel) {
	c.cs = append(b.get(), "HDEL")
	c.ks = InitSlot
	return
}

type SHdelField SCompleted

func (c SHdelField) Field(Field ...string) SHdelField {
	return SHdelField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

func (c SHdelField) Build() SCompleted {
	return SCompleted(c)
}

type SHdelKey SCompleted

func (c SHdelKey) Field(Field ...string) SHdelField {
	return SHdelField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

type SHello SCompleted

func (c SHello) Protover(Protover int64) SHelloArgumentsProtover {
	return SHelloArgumentsProtover{cs: append(c.cs, strconv.FormatInt(Protover, 10)), cf: c.cf, ks: c.ks}
}

func (c SHello) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Hello() (c SHello) {
	c.cs = append(b.get(), "HELLO")
	c.ks = InitSlot
	return
}

type SHelloArgumentsAuth SCompleted

func (c SHelloArgumentsAuth) Setname(Clientname string) SHelloArgumentsSetname {
	return SHelloArgumentsSetname{cs: append(c.cs, "SETNAME", Clientname), cf: c.cf, ks: c.ks}
}

func (c SHelloArgumentsAuth) Build() SCompleted {
	return SCompleted(c)
}

type SHelloArgumentsProtover SCompleted

func (c SHelloArgumentsProtover) Auth(Username string, Password string) SHelloArgumentsAuth {
	return SHelloArgumentsAuth{cs: append(c.cs, "AUTH", Username, Password), cf: c.cf, ks: c.ks}
}

func (c SHelloArgumentsProtover) Setname(Clientname string) SHelloArgumentsSetname {
	return SHelloArgumentsSetname{cs: append(c.cs, "SETNAME", Clientname), cf: c.cf, ks: c.ks}
}

func (c SHelloArgumentsProtover) Build() SCompleted {
	return SCompleted(c)
}

type SHelloArgumentsSetname SCompleted

func (c SHelloArgumentsSetname) Build() SCompleted {
	return SCompleted(c)
}

type SHexists SCompleted

func (c SHexists) Key(Key string) SHexistsKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHexistsKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hexists() (c SHexists) {
	c.cs = append(b.get(), "HEXISTS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHexistsField SCompleted

func (c SHexistsField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHexistsField) Cache() SCacheable {
	return SCacheable(c)
}

type SHexistsKey SCompleted

func (c SHexistsKey) Field(Field string) SHexistsField {
	return SHexistsField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type SHget SCompleted

func (c SHget) Key(Key string) SHgetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHgetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hget() (c SHget) {
	c.cs = append(b.get(), "HGET")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHgetField SCompleted

func (c SHgetField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHgetField) Cache() SCacheable {
	return SCacheable(c)
}

type SHgetKey SCompleted

func (c SHgetKey) Field(Field string) SHgetField {
	return SHgetField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type SHgetall SCompleted

func (c SHgetall) Key(Key string) SHgetallKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHgetallKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hgetall() (c SHgetall) {
	c.cs = append(b.get(), "HGETALL")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHgetallKey SCompleted

func (c SHgetallKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHgetallKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHincrby SCompleted

func (c SHincrby) Key(Key string) SHincrbyKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHincrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hincrby() (c SHincrby) {
	c.cs = append(b.get(), "HINCRBY")
	c.ks = InitSlot
	return
}

type SHincrbyField SCompleted

func (c SHincrbyField) Increment(Increment int64) SHincrbyIncrement {
	return SHincrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

type SHincrbyIncrement SCompleted

func (c SHincrbyIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SHincrbyKey SCompleted

func (c SHincrbyKey) Field(Field string) SHincrbyField {
	return SHincrbyField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type SHincrbyfloat SCompleted

func (c SHincrbyfloat) Key(Key string) SHincrbyfloatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHincrbyfloatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hincrbyfloat() (c SHincrbyfloat) {
	c.cs = append(b.get(), "HINCRBYFLOAT")
	c.ks = InitSlot
	return
}

type SHincrbyfloatField SCompleted

func (c SHincrbyfloatField) Increment(Increment float64) SHincrbyfloatIncrement {
	return SHincrbyfloatIncrement{cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SHincrbyfloatIncrement SCompleted

func (c SHincrbyfloatIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SHincrbyfloatKey SCompleted

func (c SHincrbyfloatKey) Field(Field string) SHincrbyfloatField {
	return SHincrbyfloatField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type SHkeys SCompleted

func (c SHkeys) Key(Key string) SHkeysKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHkeysKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hkeys() (c SHkeys) {
	c.cs = append(b.get(), "HKEYS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHkeysKey SCompleted

func (c SHkeysKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHkeysKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHlen SCompleted

func (c SHlen) Key(Key string) SHlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hlen() (c SHlen) {
	c.cs = append(b.get(), "HLEN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHlenKey SCompleted

func (c SHlenKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHlenKey) Cache() SCacheable {
	return SCacheable(c)
}

type SHmget SCompleted

func (c SHmget) Key(Key string) SHmgetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHmgetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hmget() (c SHmget) {
	c.cs = append(b.get(), "HMGET")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHmgetField SCompleted

func (c SHmgetField) Field(Field ...string) SHmgetField {
	return SHmgetField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

func (c SHmgetField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHmgetField) Cache() SCacheable {
	return SCacheable(c)
}

type SHmgetKey SCompleted

func (c SHmgetKey) Field(Field ...string) SHmgetField {
	return SHmgetField{cs: append(c.cs, Field...), cf: c.cf, ks: c.ks}
}

type SHmset SCompleted

func (c SHmset) Key(Key string) SHmsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHmsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hmset() (c SHmset) {
	c.cs = append(b.get(), "HMSET")
	c.ks = InitSlot
	return
}

type SHmsetFieldValue SCompleted

func (c SHmsetFieldValue) FieldValue(Field string, Value string) SHmsetFieldValue {
	return SHmsetFieldValue{cs: append(c.cs, Field, Value), cf: c.cf, ks: c.ks}
}

func (c SHmsetFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SHmsetKey SCompleted

func (c SHmsetKey) FieldValue() SHmsetFieldValue {
	return SHmsetFieldValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SHrandfield SCompleted

func (c SHrandfield) Key(Key string) SHrandfieldKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHrandfieldKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hrandfield() (c SHrandfield) {
	c.cs = append(b.get(), "HRANDFIELD")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHrandfieldKey SCompleted

func (c SHrandfieldKey) Count(Count int64) SHrandfieldOptionsCount {
	return SHrandfieldOptionsCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SHrandfieldKey) Build() SCompleted {
	return SCompleted(c)
}

type SHrandfieldOptionsCount SCompleted

func (c SHrandfieldOptionsCount) Withvalues() SHrandfieldOptionsWithvaluesWithvalues {
	return SHrandfieldOptionsWithvaluesWithvalues{cs: append(c.cs, "WITHVALUES"), cf: c.cf, ks: c.ks}
}

func (c SHrandfieldOptionsCount) Build() SCompleted {
	return SCompleted(c)
}

type SHrandfieldOptionsWithvaluesWithvalues SCompleted

func (c SHrandfieldOptionsWithvaluesWithvalues) Build() SCompleted {
	return SCompleted(c)
}

type SHscan SCompleted

func (c SHscan) Key(Key string) SHscanKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHscanKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hscan() (c SHscan) {
	c.cs = append(b.get(), "HSCAN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHscanCount SCompleted

func (c SHscanCount) Build() SCompleted {
	return SCompleted(c)
}

type SHscanCursor SCompleted

func (c SHscanCursor) Match(Pattern string) SHscanMatch {
	return SHscanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c SHscanCursor) Count(Count int64) SHscanCount {
	return SHscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SHscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SHscanKey SCompleted

func (c SHscanKey) Cursor(Cursor int64) SHscanCursor {
	return SHscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

type SHscanMatch SCompleted

func (c SHscanMatch) Count(Count int64) SHscanCount {
	return SHscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SHscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SHset SCompleted

func (c SHset) Key(Key string) SHsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hset() (c SHset) {
	c.cs = append(b.get(), "HSET")
	c.ks = InitSlot
	return
}

type SHsetFieldValue SCompleted

func (c SHsetFieldValue) FieldValue(Field string, Value string) SHsetFieldValue {
	return SHsetFieldValue{cs: append(c.cs, Field, Value), cf: c.cf, ks: c.ks}
}

func (c SHsetFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SHsetKey SCompleted

func (c SHsetKey) FieldValue() SHsetFieldValue {
	return SHsetFieldValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SHsetnx SCompleted

func (c SHsetnx) Key(Key string) SHsetnxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHsetnxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hsetnx() (c SHsetnx) {
	c.cs = append(b.get(), "HSETNX")
	c.ks = InitSlot
	return
}

type SHsetnxField SCompleted

func (c SHsetnxField) Value(Value string) SHsetnxValue {
	return SHsetnxValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SHsetnxKey SCompleted

func (c SHsetnxKey) Field(Field string) SHsetnxField {
	return SHsetnxField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type SHsetnxValue SCompleted

func (c SHsetnxValue) Build() SCompleted {
	return SCompleted(c)
}

type SHstrlen SCompleted

func (c SHstrlen) Key(Key string) SHstrlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHstrlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hstrlen() (c SHstrlen) {
	c.cs = append(b.get(), "HSTRLEN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHstrlenField SCompleted

func (c SHstrlenField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHstrlenField) Cache() SCacheable {
	return SCacheable(c)
}

type SHstrlenKey SCompleted

func (c SHstrlenKey) Field(Field string) SHstrlenField {
	return SHstrlenField{cs: append(c.cs, Field), cf: c.cf, ks: c.ks}
}

type SHvals SCompleted

func (c SHvals) Key(Key string) SHvalsKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHvalsKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Hvals() (c SHvals) {
	c.cs = append(b.get(), "HVALS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHvalsKey SCompleted

func (c SHvalsKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SHvalsKey) Cache() SCacheable {
	return SCacheable(c)
}

type SIncr SCompleted

func (c SIncr) Key(Key string) SIncrKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SIncrKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Incr() (c SIncr) {
	c.cs = append(b.get(), "INCR")
	c.ks = InitSlot
	return
}

type SIncrKey SCompleted

func (c SIncrKey) Build() SCompleted {
	return SCompleted(c)
}

type SIncrby SCompleted

func (c SIncrby) Key(Key string) SIncrbyKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SIncrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Incrby() (c SIncrby) {
	c.cs = append(b.get(), "INCRBY")
	c.ks = InitSlot
	return
}

type SIncrbyIncrement SCompleted

func (c SIncrbyIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SIncrbyKey SCompleted

func (c SIncrbyKey) Increment(Increment int64) SIncrbyIncrement {
	return SIncrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

type SIncrbyfloat SCompleted

func (c SIncrbyfloat) Key(Key string) SIncrbyfloatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SIncrbyfloatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Incrbyfloat() (c SIncrbyfloat) {
	c.cs = append(b.get(), "INCRBYFLOAT")
	c.ks = InitSlot
	return
}

type SIncrbyfloatIncrement SCompleted

func (c SIncrbyfloatIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SIncrbyfloatKey SCompleted

func (c SIncrbyfloatKey) Increment(Increment float64) SIncrbyfloatIncrement {
	return SIncrbyfloatIncrement{cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SInfo SCompleted

func (c SInfo) Section(Section string) SInfoSection {
	return SInfoSection{cs: append(c.cs, Section), cf: c.cf, ks: c.ks}
}

func (c SInfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Info() (c SInfo) {
	c.cs = append(b.get(), "INFO")
	c.ks = InitSlot
	return
}

type SInfoSection SCompleted

func (c SInfoSection) Build() SCompleted {
	return SCompleted(c)
}

type SKeys SCompleted

func (c SKeys) Pattern(Pattern string) SKeysPattern {
	return SKeysPattern{cs: append(c.cs, Pattern), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Keys() (c SKeys) {
	c.cs = append(b.get(), "KEYS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SKeysPattern SCompleted

func (c SKeysPattern) Build() SCompleted {
	return SCompleted(c)
}

type SLastsave SCompleted

func (c SLastsave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Lastsave() (c SLastsave) {
	c.cs = append(b.get(), "LASTSAVE")
	c.ks = InitSlot
	return
}

type SLatencyDoctor SCompleted

func (c SLatencyDoctor) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyDoctor() (c SLatencyDoctor) {
	c.cs = append(b.get(), "LATENCY", "DOCTOR")
	c.ks = InitSlot
	return
}

type SLatencyGraph SCompleted

func (c SLatencyGraph) Event(Event string) SLatencyGraphEvent {
	return SLatencyGraphEvent{cs: append(c.cs, Event), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) LatencyGraph() (c SLatencyGraph) {
	c.cs = append(b.get(), "LATENCY", "GRAPH")
	c.ks = InitSlot
	return
}

type SLatencyGraphEvent SCompleted

func (c SLatencyGraphEvent) Build() SCompleted {
	return SCompleted(c)
}

type SLatencyHelp SCompleted

func (c SLatencyHelp) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyHelp() (c SLatencyHelp) {
	c.cs = append(b.get(), "LATENCY", "HELP")
	c.ks = InitSlot
	return
}

type SLatencyHistory SCompleted

func (c SLatencyHistory) Event(Event string) SLatencyHistoryEvent {
	return SLatencyHistoryEvent{cs: append(c.cs, Event), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) LatencyHistory() (c SLatencyHistory) {
	c.cs = append(b.get(), "LATENCY", "HISTORY")
	c.ks = InitSlot
	return
}

type SLatencyHistoryEvent SCompleted

func (c SLatencyHistoryEvent) Build() SCompleted {
	return SCompleted(c)
}

type SLatencyLatest SCompleted

func (c SLatencyLatest) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyLatest() (c SLatencyLatest) {
	c.cs = append(b.get(), "LATENCY", "LATEST")
	c.ks = InitSlot
	return
}

type SLatencyReset SCompleted

func (c SLatencyReset) Event(Event ...string) SLatencyResetEvent {
	return SLatencyResetEvent{cs: append(c.cs, Event...), cf: c.cf, ks: c.ks}
}

func (c SLatencyReset) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) LatencyReset() (c SLatencyReset) {
	c.cs = append(b.get(), "LATENCY", "RESET")
	c.ks = InitSlot
	return
}

type SLatencyResetEvent SCompleted

func (c SLatencyResetEvent) Event(Event ...string) SLatencyResetEvent {
	return SLatencyResetEvent{cs: append(c.cs, Event...), cf: c.cf, ks: c.ks}
}

func (c SLatencyResetEvent) Build() SCompleted {
	return SCompleted(c)
}

type SLindex SCompleted

func (c SLindex) Key(Key string) SLindexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLindexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lindex() (c SLindex) {
	c.cs = append(b.get(), "LINDEX")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SLindexIndex SCompleted

func (c SLindexIndex) Build() SCompleted {
	return SCompleted(c)
}

func (c SLindexIndex) Cache() SCacheable {
	return SCacheable(c)
}

type SLindexKey SCompleted

func (c SLindexKey) Index(Index int64) SLindexIndex {
	return SLindexIndex{cs: append(c.cs, strconv.FormatInt(Index, 10)), cf: c.cf, ks: c.ks}
}

type SLinsert SCompleted

func (c SLinsert) Key(Key string) SLinsertKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLinsertKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Linsert() (c SLinsert) {
	c.cs = append(b.get(), "LINSERT")
	c.ks = InitSlot
	return
}

type SLinsertElement SCompleted

func (c SLinsertElement) Build() SCompleted {
	return SCompleted(c)
}

type SLinsertKey SCompleted

func (c SLinsertKey) Before() SLinsertWhereBefore {
	return SLinsertWhereBefore{cs: append(c.cs, "BEFORE"), cf: c.cf, ks: c.ks}
}

func (c SLinsertKey) After() SLinsertWhereAfter {
	return SLinsertWhereAfter{cs: append(c.cs, "AFTER"), cf: c.cf, ks: c.ks}
}

type SLinsertPivot SCompleted

func (c SLinsertPivot) Element(Element string) SLinsertElement {
	return SLinsertElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type SLinsertWhereAfter SCompleted

func (c SLinsertWhereAfter) Pivot(Pivot string) SLinsertPivot {
	return SLinsertPivot{cs: append(c.cs, Pivot), cf: c.cf, ks: c.ks}
}

type SLinsertWhereBefore SCompleted

func (c SLinsertWhereBefore) Pivot(Pivot string) SLinsertPivot {
	return SLinsertPivot{cs: append(c.cs, Pivot), cf: c.cf, ks: c.ks}
}

type SLlen SCompleted

func (c SLlen) Key(Key string) SLlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Llen() (c SLlen) {
	c.cs = append(b.get(), "LLEN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SLlenKey SCompleted

func (c SLlenKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SLlenKey) Cache() SCacheable {
	return SCacheable(c)
}

type SLmove SCompleted

func (c SLmove) Source(Source string) SLmoveSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SLmoveSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lmove() (c SLmove) {
	c.cs = append(b.get(), "LMOVE")
	c.ks = InitSlot
	return
}

type SLmoveDestination SCompleted

func (c SLmoveDestination) Left() SLmoveWherefromLeft {
	return SLmoveWherefromLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SLmoveDestination) Right() SLmoveWherefromRight {
	return SLmoveWherefromRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SLmoveSource SCompleted

func (c SLmoveSource) Destination(Destination string) SLmoveDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SLmoveDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type SLmoveWherefromLeft SCompleted

func (c SLmoveWherefromLeft) Left() SLmoveWheretoLeft {
	return SLmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SLmoveWherefromLeft) Right() SLmoveWheretoRight {
	return SLmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SLmoveWherefromRight SCompleted

func (c SLmoveWherefromRight) Left() SLmoveWheretoLeft {
	return SLmoveWheretoLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SLmoveWherefromRight) Right() SLmoveWheretoRight {
	return SLmoveWheretoRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SLmoveWheretoLeft SCompleted

func (c SLmoveWheretoLeft) Build() SCompleted {
	return SCompleted(c)
}

type SLmoveWheretoRight SCompleted

func (c SLmoveWheretoRight) Build() SCompleted {
	return SCompleted(c)
}

type SLmpop SCompleted

func (c SLmpop) Numkeys(Numkeys int64) SLmpopNumkeys {
	return SLmpopNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lmpop() (c SLmpop) {
	c.cs = append(b.get(), "LMPOP")
	c.ks = InitSlot
	return
}

type SLmpopCount SCompleted

func (c SLmpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SLmpopKey SCompleted

func (c SLmpopKey) Left() SLmpopWhereLeft {
	return SLmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SLmpopKey) Right() SLmpopWhereRight {
	return SLmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

func (c SLmpopKey) Key(Key ...string) SLmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SLmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SLmpopNumkeys SCompleted

func (c SLmpopNumkeys) Key(Key ...string) SLmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SLmpopKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SLmpopNumkeys) Left() SLmpopWhereLeft {
	return SLmpopWhereLeft{cs: append(c.cs, "LEFT"), cf: c.cf, ks: c.ks}
}

func (c SLmpopNumkeys) Right() SLmpopWhereRight {
	return SLmpopWhereRight{cs: append(c.cs, "RIGHT"), cf: c.cf, ks: c.ks}
}

type SLmpopWhereLeft SCompleted

func (c SLmpopWhereLeft) Count(Count int64) SLmpopCount {
	return SLmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SLmpopWhereLeft) Build() SCompleted {
	return SCompleted(c)
}

type SLmpopWhereRight SCompleted

func (c SLmpopWhereRight) Count(Count int64) SLmpopCount {
	return SLmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SLmpopWhereRight) Build() SCompleted {
	return SCompleted(c)
}

type SLolwut SCompleted

func (c SLolwut) Version(Version int64) SLolwutVersion {
	return SLolwutVersion{cs: append(c.cs, "VERSION", strconv.FormatInt(Version, 10)), cf: c.cf, ks: c.ks}
}

func (c SLolwut) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Lolwut() (c SLolwut) {
	c.cs = append(b.get(), "LOLWUT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SLolwutVersion SCompleted

func (c SLolwutVersion) Build() SCompleted {
	return SCompleted(c)
}

type SLpop SCompleted

func (c SLpop) Key(Key string) SLpopKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLpopKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lpop() (c SLpop) {
	c.cs = append(b.get(), "LPOP")
	c.ks = InitSlot
	return
}

type SLpopCount SCompleted

func (c SLpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SLpopKey SCompleted

func (c SLpopKey) Count(Count int64) SLpopCount {
	return SLpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SLpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SLpos SCompleted

func (c SLpos) Key(Key string) SLposKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLposKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lpos() (c SLpos) {
	c.cs = append(b.get(), "LPOS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SLposCount SCompleted

func (c SLposCount) Maxlen(Len int64) SLposMaxlen {
	return SLposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10)), cf: c.cf, ks: c.ks}
}

func (c SLposCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposCount) Cache() SCacheable {
	return SCacheable(c)
}

type SLposElement SCompleted

func (c SLposElement) Rank(Rank int64) SLposRank {
	return SLposRank{cs: append(c.cs, "RANK", strconv.FormatInt(Rank, 10)), cf: c.cf, ks: c.ks}
}

func (c SLposElement) Count(NumMatches int64) SLposCount {
	return SLposCount{cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10)), cf: c.cf, ks: c.ks}
}

func (c SLposElement) Maxlen(Len int64) SLposMaxlen {
	return SLposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10)), cf: c.cf, ks: c.ks}
}

func (c SLposElement) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposElement) Cache() SCacheable {
	return SCacheable(c)
}

type SLposKey SCompleted

func (c SLposKey) Element(Element string) SLposElement {
	return SLposElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type SLposMaxlen SCompleted

func (c SLposMaxlen) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposMaxlen) Cache() SCacheable {
	return SCacheable(c)
}

type SLposRank SCompleted

func (c SLposRank) Count(NumMatches int64) SLposCount {
	return SLposCount{cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10)), cf: c.cf, ks: c.ks}
}

func (c SLposRank) Maxlen(Len int64) SLposMaxlen {
	return SLposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10)), cf: c.cf, ks: c.ks}
}

func (c SLposRank) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposRank) Cache() SCacheable {
	return SCacheable(c)
}

type SLpush SCompleted

func (c SLpush) Key(Key string) SLpushKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLpushKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lpush() (c SLpush) {
	c.cs = append(b.get(), "LPUSH")
	c.ks = InitSlot
	return
}

type SLpushElement SCompleted

func (c SLpushElement) Element(Element ...string) SLpushElement {
	return SLpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c SLpushElement) Build() SCompleted {
	return SCompleted(c)
}

type SLpushKey SCompleted

func (c SLpushKey) Element(Element ...string) SLpushElement {
	return SLpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type SLpushx SCompleted

func (c SLpushx) Key(Key string) SLpushxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLpushxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lpushx() (c SLpushx) {
	c.cs = append(b.get(), "LPUSHX")
	c.ks = InitSlot
	return
}

type SLpushxElement SCompleted

func (c SLpushxElement) Element(Element ...string) SLpushxElement {
	return SLpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c SLpushxElement) Build() SCompleted {
	return SCompleted(c)
}

type SLpushxKey SCompleted

func (c SLpushxKey) Element(Element ...string) SLpushxElement {
	return SLpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type SLrange SCompleted

func (c SLrange) Key(Key string) SLrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lrange() (c SLrange) {
	c.cs = append(b.get(), "LRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SLrangeKey SCompleted

func (c SLrangeKey) Start(Start int64) SLrangeStart {
	return SLrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type SLrangeStart SCompleted

func (c SLrangeStart) Stop(Stop int64) SLrangeStop {
	return SLrangeStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type SLrangeStop SCompleted

func (c SLrangeStop) Build() SCompleted {
	return SCompleted(c)
}

func (c SLrangeStop) Cache() SCacheable {
	return SCacheable(c)
}

type SLrem SCompleted

func (c SLrem) Key(Key string) SLremKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLremKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lrem() (c SLrem) {
	c.cs = append(b.get(), "LREM")
	c.ks = InitSlot
	return
}

type SLremCount SCompleted

func (c SLremCount) Element(Element string) SLremElement {
	return SLremElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type SLremElement SCompleted

func (c SLremElement) Build() SCompleted {
	return SCompleted(c)
}

type SLremKey SCompleted

func (c SLremKey) Count(Count int64) SLremCount {
	return SLremCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

type SLset SCompleted

func (c SLset) Key(Key string) SLsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLsetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Lset() (c SLset) {
	c.cs = append(b.get(), "LSET")
	c.ks = InitSlot
	return
}

type SLsetElement SCompleted

func (c SLsetElement) Build() SCompleted {
	return SCompleted(c)
}

type SLsetIndex SCompleted

func (c SLsetIndex) Element(Element string) SLsetElement {
	return SLsetElement{cs: append(c.cs, Element), cf: c.cf, ks: c.ks}
}

type SLsetKey SCompleted

func (c SLsetKey) Index(Index int64) SLsetIndex {
	return SLsetIndex{cs: append(c.cs, strconv.FormatInt(Index, 10)), cf: c.cf, ks: c.ks}
}

type SLtrim SCompleted

func (c SLtrim) Key(Key string) SLtrimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLtrimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Ltrim() (c SLtrim) {
	c.cs = append(b.get(), "LTRIM")
	c.ks = InitSlot
	return
}

type SLtrimKey SCompleted

func (c SLtrimKey) Start(Start int64) SLtrimStart {
	return SLtrimStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type SLtrimStart SCompleted

func (c SLtrimStart) Stop(Stop int64) SLtrimStop {
	return SLtrimStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type SLtrimStop SCompleted

func (c SLtrimStop) Build() SCompleted {
	return SCompleted(c)
}

type SMemoryDoctor SCompleted

func (c SMemoryDoctor) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryDoctor() (c SMemoryDoctor) {
	c.cs = append(b.get(), "MEMORY", "DOCTOR")
	c.ks = InitSlot
	return
}

type SMemoryHelp SCompleted

func (c SMemoryHelp) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryHelp() (c SMemoryHelp) {
	c.cs = append(b.get(), "MEMORY", "HELP")
	c.ks = InitSlot
	return
}

type SMemoryMallocStats SCompleted

func (c SMemoryMallocStats) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryMallocStats() (c SMemoryMallocStats) {
	c.cs = append(b.get(), "MEMORY", "MALLOC-STATS")
	c.ks = InitSlot
	return
}

type SMemoryPurge SCompleted

func (c SMemoryPurge) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryPurge() (c SMemoryPurge) {
	c.cs = append(b.get(), "MEMORY", "PURGE")
	c.ks = InitSlot
	return
}

type SMemoryStats SCompleted

func (c SMemoryStats) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) MemoryStats() (c SMemoryStats) {
	c.cs = append(b.get(), "MEMORY", "STATS")
	c.ks = InitSlot
	return
}

type SMemoryUsage SCompleted

func (c SMemoryUsage) Key(Key string) SMemoryUsageKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SMemoryUsageKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) MemoryUsage() (c SMemoryUsage) {
	c.cs = append(b.get(), "MEMORY", "USAGE")
	c.ks = InitSlot
	return
}

type SMemoryUsageKey SCompleted

func (c SMemoryUsageKey) Samples(Count int64) SMemoryUsageSamples {
	return SMemoryUsageSamples{cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SMemoryUsageKey) Build() SCompleted {
	return SCompleted(c)
}

type SMemoryUsageSamples SCompleted

func (c SMemoryUsageSamples) Build() SCompleted {
	return SCompleted(c)
}

type SMget SCompleted

func (c SMget) Key(Key ...string) SMgetKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SMgetKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Mget() (c SMget) {
	c.cs = append(b.get(), "MGET")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SMgetKey SCompleted

func (c SMgetKey) Key(Key ...string) SMgetKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SMgetKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SMgetKey) Build() SCompleted {
	return SCompleted(c)
}

type SMigrate SCompleted

func (c SMigrate) Host(Host string) SMigrateHost {
	return SMigrateHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Migrate() (c SMigrate) {
	c.cs = append(b.get(), "MIGRATE")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SMigrateAuth SCompleted

func (c SMigrateAuth) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c SMigrateAuth) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SMigrateAuth) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateAuth2 SCompleted

func (c SMigrateAuth2) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SMigrateAuth2) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateCopyCopy SCompleted

func (c SMigrateCopyCopy) Replace() SMigrateReplaceReplace {
	return SMigrateReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c SMigrateCopyCopy) Auth(Password string) SMigrateAuth {
	return SMigrateAuth{cs: append(c.cs, "AUTH", Password), cf: c.cf, ks: c.ks}
}

func (c SMigrateCopyCopy) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c SMigrateCopyCopy) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SMigrateCopyCopy) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateDestinationDb SCompleted

func (c SMigrateDestinationDb) Timeout(Timeout int64) SMigrateTimeout {
	return SMigrateTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10)), cf: c.cf, ks: c.ks}
}

type SMigrateHost SCompleted

func (c SMigrateHost) Port(Port string) SMigratePort {
	return SMigratePort{cs: append(c.cs, Port), cf: c.cf, ks: c.ks}
}

type SMigrateKeyEmpty SCompleted

func (c SMigrateKeyEmpty) DestinationDb(DestinationDb int64) SMigrateDestinationDb {
	return SMigrateDestinationDb{cs: append(c.cs, strconv.FormatInt(DestinationDb, 10)), cf: c.cf, ks: c.ks}
}

type SMigrateKeyKey SCompleted

func (c SMigrateKeyKey) DestinationDb(DestinationDb int64) SMigrateDestinationDb {
	return SMigrateDestinationDb{cs: append(c.cs, strconv.FormatInt(DestinationDb, 10)), cf: c.cf, ks: c.ks}
}

type SMigrateKeys SCompleted

func (c SMigrateKeys) Keys(Keys ...string) SMigrateKeys {
	for _, k := range Keys {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SMigrateKeys{cs: append(c.cs, Keys...), cf: c.cf, ks: c.ks}
}

func (c SMigrateKeys) Build() SCompleted {
	return SCompleted(c)
}

type SMigratePort SCompleted

func (c SMigratePort) Key() SMigrateKeyKey {
	return SMigrateKeyKey{cs: append(c.cs, "key"), cf: c.cf, ks: c.ks}
}

func (c SMigratePort) Empty() SMigrateKeyEmpty {
	return SMigrateKeyEmpty{cs: append(c.cs, "\"\""), cf: c.cf, ks: c.ks}
}

type SMigrateReplaceReplace SCompleted

func (c SMigrateReplaceReplace) Auth(Password string) SMigrateAuth {
	return SMigrateAuth{cs: append(c.cs, "AUTH", Password), cf: c.cf, ks: c.ks}
}

func (c SMigrateReplaceReplace) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c SMigrateReplaceReplace) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SMigrateReplaceReplace) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateTimeout SCompleted

func (c SMigrateTimeout) Copy() SMigrateCopyCopy {
	return SMigrateCopyCopy{cs: append(c.cs, "COPY"), cf: c.cf, ks: c.ks}
}

func (c SMigrateTimeout) Replace() SMigrateReplaceReplace {
	return SMigrateReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c SMigrateTimeout) Auth(Password string) SMigrateAuth {
	return SMigrateAuth{cs: append(c.cs, "AUTH", Password), cf: c.cf, ks: c.ks}
}

func (c SMigrateTimeout) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword), cf: c.cf, ks: c.ks}
}

func (c SMigrateTimeout) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SMigrateTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SModuleList SCompleted

func (c SModuleList) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ModuleList() (c SModuleList) {
	c.cs = append(b.get(), "MODULE", "LIST")
	c.ks = InitSlot
	return
}

type SModuleLoad SCompleted

func (c SModuleLoad) Path(Path string) SModuleLoadPath {
	return SModuleLoadPath{cs: append(c.cs, Path), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ModuleLoad() (c SModuleLoad) {
	c.cs = append(b.get(), "MODULE", "LOAD")
	c.ks = InitSlot
	return
}

type SModuleLoadArg SCompleted

func (c SModuleLoadArg) Arg(Arg ...string) SModuleLoadArg {
	return SModuleLoadArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SModuleLoadArg) Build() SCompleted {
	return SCompleted(c)
}

type SModuleLoadPath SCompleted

func (c SModuleLoadPath) Arg(Arg ...string) SModuleLoadArg {
	return SModuleLoadArg{cs: append(c.cs, Arg...), cf: c.cf, ks: c.ks}
}

func (c SModuleLoadPath) Build() SCompleted {
	return SCompleted(c)
}

type SModuleUnload SCompleted

func (c SModuleUnload) Name(Name string) SModuleUnloadName {
	return SModuleUnloadName{cs: append(c.cs, Name), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ModuleUnload() (c SModuleUnload) {
	c.cs = append(b.get(), "MODULE", "UNLOAD")
	c.ks = InitSlot
	return
}

type SModuleUnloadName SCompleted

func (c SModuleUnloadName) Build() SCompleted {
	return SCompleted(c)
}

type SMonitor SCompleted

func (c SMonitor) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Monitor() (c SMonitor) {
	c.cs = append(b.get(), "MONITOR")
	c.ks = InitSlot
	return
}

type SMove SCompleted

func (c SMove) Key(Key string) SMoveKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SMoveKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Move() (c SMove) {
	c.cs = append(b.get(), "MOVE")
	c.ks = InitSlot
	return
}

type SMoveDb SCompleted

func (c SMoveDb) Build() SCompleted {
	return SCompleted(c)
}

type SMoveKey SCompleted

func (c SMoveKey) Db(Db int64) SMoveDb {
	return SMoveDb{cs: append(c.cs, strconv.FormatInt(Db, 10)), cf: c.cf, ks: c.ks}
}

type SMset SCompleted

func (c SMset) KeyValue() SMsetKeyValue {
	return SMsetKeyValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Mset() (c SMset) {
	c.cs = append(b.get(), "MSET")
	c.ks = InitSlot
	return
}

type SMsetKeyValue SCompleted

func (c SMsetKeyValue) KeyValue(Key string, Value string) SMsetKeyValue {
	c.ks = checkSlot(c.ks, slot(Key))
	return SMsetKeyValue{cs: append(c.cs, Key, Value), cf: c.cf, ks: c.ks}
}

func (c SMsetKeyValue) Build() SCompleted {
	return SCompleted(c)
}

type SMsetnx SCompleted

func (c SMsetnx) KeyValue() SMsetnxKeyValue {
	return SMsetnxKeyValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Msetnx() (c SMsetnx) {
	c.cs = append(b.get(), "MSETNX")
	c.ks = InitSlot
	return
}

type SMsetnxKeyValue SCompleted

func (c SMsetnxKeyValue) KeyValue(Key string, Value string) SMsetnxKeyValue {
	c.ks = checkSlot(c.ks, slot(Key))
	return SMsetnxKeyValue{cs: append(c.cs, Key, Value), cf: c.cf, ks: c.ks}
}

func (c SMsetnxKeyValue) Build() SCompleted {
	return SCompleted(c)
}

type SMulti SCompleted

func (c SMulti) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Multi() (c SMulti) {
	c.cs = append(b.get(), "MULTI")
	c.ks = InitSlot
	return
}

type SObject SCompleted

func (c SObject) Subcommand(Subcommand string) SObjectSubcommand {
	return SObjectSubcommand{cs: append(c.cs, Subcommand), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Object() (c SObject) {
	c.cs = append(b.get(), "OBJECT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SObjectArguments SCompleted

func (c SObjectArguments) Arguments(Arguments ...string) SObjectArguments {
	return SObjectArguments{cs: append(c.cs, Arguments...), cf: c.cf, ks: c.ks}
}

func (c SObjectArguments) Build() SCompleted {
	return SCompleted(c)
}

type SObjectSubcommand SCompleted

func (c SObjectSubcommand) Arguments(Arguments ...string) SObjectArguments {
	return SObjectArguments{cs: append(c.cs, Arguments...), cf: c.cf, ks: c.ks}
}

func (c SObjectSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SPersist SCompleted

func (c SPersist) Key(Key string) SPersistKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPersistKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Persist() (c SPersist) {
	c.cs = append(b.get(), "PERSIST")
	c.ks = InitSlot
	return
}

type SPersistKey SCompleted

func (c SPersistKey) Build() SCompleted {
	return SCompleted(c)
}

type SPexpire SCompleted

func (c SPexpire) Key(Key string) SPexpireKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPexpireKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pexpire() (c SPexpire) {
	c.cs = append(b.get(), "PEXPIRE")
	c.ks = InitSlot
	return
}

type SPexpireConditionGt SCompleted

func (c SPexpireConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireConditionLt SCompleted

func (c SPexpireConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireConditionNx SCompleted

func (c SPexpireConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireConditionXx SCompleted

func (c SPexpireConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireKey SCompleted

func (c SPexpireKey) Milliseconds(Milliseconds int64) SPexpireMilliseconds {
	return SPexpireMilliseconds{cs: append(c.cs, strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

type SPexpireMilliseconds SCompleted

func (c SPexpireMilliseconds) Nx() SPexpireConditionNx {
	return SPexpireConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SPexpireMilliseconds) Xx() SPexpireConditionXx {
	return SPexpireConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SPexpireMilliseconds) Gt() SPexpireConditionGt {
	return SPexpireConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c SPexpireMilliseconds) Lt() SPexpireConditionLt {
	return SPexpireConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c SPexpireMilliseconds) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireat SCompleted

func (c SPexpireat) Key(Key string) SPexpireatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPexpireatKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pexpireat() (c SPexpireat) {
	c.cs = append(b.get(), "PEXPIREAT")
	c.ks = InitSlot
	return
}

type SPexpireatConditionGt SCompleted

func (c SPexpireatConditionGt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatConditionLt SCompleted

func (c SPexpireatConditionLt) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatConditionNx SCompleted

func (c SPexpireatConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatConditionXx SCompleted

func (c SPexpireatConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SPexpireatKey SCompleted

func (c SPexpireatKey) MillisecondsTimestamp(MillisecondsTimestamp int64) SPexpireatMillisecondsTimestamp {
	return SPexpireatMillisecondsTimestamp{cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10)), cf: c.cf, ks: c.ks}
}

type SPexpireatMillisecondsTimestamp SCompleted

func (c SPexpireatMillisecondsTimestamp) Nx() SPexpireatConditionNx {
	return SPexpireatConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SPexpireatMillisecondsTimestamp) Xx() SPexpireatConditionXx {
	return SPexpireatConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SPexpireatMillisecondsTimestamp) Gt() SPexpireatConditionGt {
	return SPexpireatConditionGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c SPexpireatMillisecondsTimestamp) Lt() SPexpireatConditionLt {
	return SPexpireatConditionLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c SPexpireatMillisecondsTimestamp) Build() SCompleted {
	return SCompleted(c)
}

type SPexpiretime SCompleted

func (c SPexpiretime) Key(Key string) SPexpiretimeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPexpiretimeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pexpiretime() (c SPexpiretime) {
	c.cs = append(b.get(), "PEXPIRETIME")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SPexpiretimeKey SCompleted

func (c SPexpiretimeKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SPexpiretimeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SPfadd SCompleted

func (c SPfadd) Key(Key string) SPfaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPfaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pfadd() (c SPfadd) {
	c.cs = append(b.get(), "PFADD")
	c.ks = InitSlot
	return
}

type SPfaddElement SCompleted

func (c SPfaddElement) Element(Element ...string) SPfaddElement {
	return SPfaddElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c SPfaddElement) Build() SCompleted {
	return SCompleted(c)
}

type SPfaddKey SCompleted

func (c SPfaddKey) Element(Element ...string) SPfaddElement {
	return SPfaddElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c SPfaddKey) Build() SCompleted {
	return SCompleted(c)
}

type SPfcount SCompleted

func (c SPfcount) Key(Key ...string) SPfcountKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SPfcountKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pfcount() (c SPfcount) {
	c.cs = append(b.get(), "PFCOUNT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SPfcountKey SCompleted

func (c SPfcountKey) Key(Key ...string) SPfcountKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SPfcountKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SPfcountKey) Build() SCompleted {
	return SCompleted(c)
}

type SPfmerge SCompleted

func (c SPfmerge) Destkey(Destkey string) SPfmergeDestkey {
	c.ks = checkSlot(c.ks, slot(Destkey))
	return SPfmergeDestkey{cs: append(c.cs, Destkey), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pfmerge() (c SPfmerge) {
	c.cs = append(b.get(), "PFMERGE")
	c.ks = InitSlot
	return
}

type SPfmergeDestkey SCompleted

func (c SPfmergeDestkey) Sourcekey(Sourcekey ...string) SPfmergeSourcekey {
	for _, k := range Sourcekey {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SPfmergeSourcekey{cs: append(c.cs, Sourcekey...), cf: c.cf, ks: c.ks}
}

type SPfmergeSourcekey SCompleted

func (c SPfmergeSourcekey) Sourcekey(Sourcekey ...string) SPfmergeSourcekey {
	for _, k := range Sourcekey {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SPfmergeSourcekey{cs: append(c.cs, Sourcekey...), cf: c.cf, ks: c.ks}
}

func (c SPfmergeSourcekey) Build() SCompleted {
	return SCompleted(c)
}

type SPing SCompleted

func (c SPing) Message(Message string) SPingMessage {
	return SPingMessage{cs: append(c.cs, Message), cf: c.cf, ks: c.ks}
}

func (c SPing) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Ping() (c SPing) {
	c.cs = append(b.get(), "PING")
	c.ks = InitSlot
	return
}

type SPingMessage SCompleted

func (c SPingMessage) Build() SCompleted {
	return SCompleted(c)
}

type SPsetex SCompleted

func (c SPsetex) Key(Key string) SPsetexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPsetexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Psetex() (c SPsetex) {
	c.cs = append(b.get(), "PSETEX")
	c.ks = InitSlot
	return
}

type SPsetexKey SCompleted

func (c SPsetexKey) Milliseconds(Milliseconds int64) SPsetexMilliseconds {
	return SPsetexMilliseconds{cs: append(c.cs, strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

type SPsetexMilliseconds SCompleted

func (c SPsetexMilliseconds) Value(Value string) SPsetexValue {
	return SPsetexValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SPsetexValue SCompleted

func (c SPsetexValue) Build() SCompleted {
	return SCompleted(c)
}

type SPsubscribe SCompleted

func (c SPsubscribe) Pattern(Pattern ...string) SPsubscribePattern {
	return SPsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Psubscribe() (c SPsubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	c.cf = noRetTag
	c.ks = InitSlot
	return
}

type SPsubscribePattern SCompleted

func (c SPsubscribePattern) Pattern(Pattern ...string) SPsubscribePattern {
	return SPsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SPsubscribePattern) Build() SCompleted {
	return SCompleted(c)
}

type SPsync SCompleted

func (c SPsync) Replicationid(Replicationid int64) SPsyncReplicationid {
	return SPsyncReplicationid{cs: append(c.cs, strconv.FormatInt(Replicationid, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Psync() (c SPsync) {
	c.cs = append(b.get(), "PSYNC")
	c.ks = InitSlot
	return
}

type SPsyncOffset SCompleted

func (c SPsyncOffset) Build() SCompleted {
	return SCompleted(c)
}

type SPsyncReplicationid SCompleted

func (c SPsyncReplicationid) Offset(Offset int64) SPsyncOffset {
	return SPsyncOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type SPttl SCompleted

func (c SPttl) Key(Key string) SPttlKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPttlKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pttl() (c SPttl) {
	c.cs = append(b.get(), "PTTL")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SPttlKey SCompleted

func (c SPttlKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SPttlKey) Cache() SCacheable {
	return SCacheable(c)
}

type SPublish SCompleted

func (c SPublish) Channel(Channel string) SPublishChannel {
	return SPublishChannel{cs: append(c.cs, Channel), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Publish() (c SPublish) {
	c.cs = append(b.get(), "PUBLISH")
	c.ks = InitSlot
	return
}

type SPublishChannel SCompleted

func (c SPublishChannel) Message(Message string) SPublishMessage {
	return SPublishMessage{cs: append(c.cs, Message), cf: c.cf, ks: c.ks}
}

type SPublishMessage SCompleted

func (c SPublishMessage) Build() SCompleted {
	return SCompleted(c)
}

type SPubsub SCompleted

func (c SPubsub) Subcommand(Subcommand string) SPubsubSubcommand {
	return SPubsubSubcommand{cs: append(c.cs, Subcommand), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Pubsub() (c SPubsub) {
	c.cs = append(b.get(), "PUBSUB")
	c.ks = InitSlot
	return
}

type SPubsubArgument SCompleted

func (c SPubsubArgument) Argument(Argument ...string) SPubsubArgument {
	return SPubsubArgument{cs: append(c.cs, Argument...), cf: c.cf, ks: c.ks}
}

func (c SPubsubArgument) Build() SCompleted {
	return SCompleted(c)
}

type SPubsubSubcommand SCompleted

func (c SPubsubSubcommand) Argument(Argument ...string) SPubsubArgument {
	return SPubsubArgument{cs: append(c.cs, Argument...), cf: c.cf, ks: c.ks}
}

func (c SPubsubSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SPunsubscribe SCompleted

func (c SPunsubscribe) Pattern(Pattern ...string) SPunsubscribePattern {
	return SPunsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SPunsubscribe) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Punsubscribe() (c SPunsubscribe) {
	c.cs = append(b.get(), "PUNSUBSCRIBE")
	c.cf = noRetTag
	c.ks = InitSlot
	return
}

type SPunsubscribePattern SCompleted

func (c SPunsubscribePattern) Pattern(Pattern ...string) SPunsubscribePattern {
	return SPunsubscribePattern{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SPunsubscribePattern) Build() SCompleted {
	return SCompleted(c)
}

type SQuit SCompleted

func (c SQuit) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Quit() (c SQuit) {
	c.cs = append(b.get(), "QUIT")
	c.ks = InitSlot
	return
}

type SRandomkey SCompleted

func (c SRandomkey) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Randomkey() (c SRandomkey) {
	c.cs = append(b.get(), "RANDOMKEY")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SReadonly SCompleted

func (c SReadonly) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Readonly() (c SReadonly) {
	c.cs = append(b.get(), "READONLY")
	c.ks = InitSlot
	return
}

type SReadwrite SCompleted

func (c SReadwrite) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Readwrite() (c SReadwrite) {
	c.cs = append(b.get(), "READWRITE")
	c.ks = InitSlot
	return
}

type SRename SCompleted

func (c SRename) Key(Key string) SRenameKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRenameKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Rename() (c SRename) {
	c.cs = append(b.get(), "RENAME")
	c.ks = InitSlot
	return
}

type SRenameKey SCompleted

func (c SRenameKey) Newkey(Newkey string) SRenameNewkey {
	c.ks = checkSlot(c.ks, slot(Newkey))
	return SRenameNewkey{cs: append(c.cs, Newkey), cf: c.cf, ks: c.ks}
}

type SRenameNewkey SCompleted

func (c SRenameNewkey) Build() SCompleted {
	return SCompleted(c)
}

type SRenamenx SCompleted

func (c SRenamenx) Key(Key string) SRenamenxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRenamenxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Renamenx() (c SRenamenx) {
	c.cs = append(b.get(), "RENAMENX")
	c.ks = InitSlot
	return
}

type SRenamenxKey SCompleted

func (c SRenamenxKey) Newkey(Newkey string) SRenamenxNewkey {
	c.ks = checkSlot(c.ks, slot(Newkey))
	return SRenamenxNewkey{cs: append(c.cs, Newkey), cf: c.cf, ks: c.ks}
}

type SRenamenxNewkey SCompleted

func (c SRenamenxNewkey) Build() SCompleted {
	return SCompleted(c)
}

type SReplicaof SCompleted

func (c SReplicaof) Host(Host string) SReplicaofHost {
	return SReplicaofHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Replicaof() (c SReplicaof) {
	c.cs = append(b.get(), "REPLICAOF")
	c.ks = InitSlot
	return
}

type SReplicaofHost SCompleted

func (c SReplicaofHost) Port(Port string) SReplicaofPort {
	return SReplicaofPort{cs: append(c.cs, Port), cf: c.cf, ks: c.ks}
}

type SReplicaofPort SCompleted

func (c SReplicaofPort) Build() SCompleted {
	return SCompleted(c)
}

type SReset SCompleted

func (c SReset) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Reset() (c SReset) {
	c.cs = append(b.get(), "RESET")
	c.ks = InitSlot
	return
}

type SRestore SCompleted

func (c SRestore) Key(Key string) SRestoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRestoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Restore() (c SRestore) {
	c.cs = append(b.get(), "RESTORE")
	c.ks = InitSlot
	return
}

type SRestoreAbsttlAbsttl SCompleted

func (c SRestoreAbsttlAbsttl) Idletime(Seconds int64) SRestoreIdletime {
	return SRestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SRestoreAbsttlAbsttl) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c SRestoreAbsttlAbsttl) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreFreq SCompleted

func (c SRestoreFreq) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreIdletime SCompleted

func (c SRestoreIdletime) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c SRestoreIdletime) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreKey SCompleted

func (c SRestoreKey) Ttl(Ttl int64) SRestoreTtl {
	return SRestoreTtl{cs: append(c.cs, strconv.FormatInt(Ttl, 10)), cf: c.cf, ks: c.ks}
}

type SRestoreReplaceReplace SCompleted

func (c SRestoreReplaceReplace) Absttl() SRestoreAbsttlAbsttl {
	return SRestoreAbsttlAbsttl{cs: append(c.cs, "ABSTTL"), cf: c.cf, ks: c.ks}
}

func (c SRestoreReplaceReplace) Idletime(Seconds int64) SRestoreIdletime {
	return SRestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SRestoreReplaceReplace) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c SRestoreReplaceReplace) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreSerializedValue SCompleted

func (c SRestoreSerializedValue) Replace() SRestoreReplaceReplace {
	return SRestoreReplaceReplace{cs: append(c.cs, "REPLACE"), cf: c.cf, ks: c.ks}
}

func (c SRestoreSerializedValue) Absttl() SRestoreAbsttlAbsttl {
	return SRestoreAbsttlAbsttl{cs: append(c.cs, "ABSTTL"), cf: c.cf, ks: c.ks}
}

func (c SRestoreSerializedValue) Idletime(Seconds int64) SRestoreIdletime {
	return SRestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SRestoreSerializedValue) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10)), cf: c.cf, ks: c.ks}
}

func (c SRestoreSerializedValue) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreTtl SCompleted

func (c SRestoreTtl) SerializedValue(SerializedValue string) SRestoreSerializedValue {
	return SRestoreSerializedValue{cs: append(c.cs, SerializedValue), cf: c.cf, ks: c.ks}
}

type SRole SCompleted

func (c SRole) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Role() (c SRole) {
	c.cs = append(b.get(), "ROLE")
	c.ks = InitSlot
	return
}

type SRpop SCompleted

func (c SRpop) Key(Key string) SRpopKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRpopKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Rpop() (c SRpop) {
	c.cs = append(b.get(), "RPOP")
	c.ks = InitSlot
	return
}

type SRpopCount SCompleted

func (c SRpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SRpopKey SCompleted

func (c SRpopKey) Count(Count int64) SRpopCount {
	return SRpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SRpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SRpoplpush SCompleted

func (c SRpoplpush) Source(Source string) SRpoplpushSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SRpoplpushSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Rpoplpush() (c SRpoplpush) {
	c.cs = append(b.get(), "RPOPLPUSH")
	c.ks = InitSlot
	return
}

type SRpoplpushDestination SCompleted

func (c SRpoplpushDestination) Build() SCompleted {
	return SCompleted(c)
}

type SRpoplpushSource SCompleted

func (c SRpoplpushSource) Destination(Destination string) SRpoplpushDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SRpoplpushDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type SRpush SCompleted

func (c SRpush) Key(Key string) SRpushKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRpushKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Rpush() (c SRpush) {
	c.cs = append(b.get(), "RPUSH")
	c.ks = InitSlot
	return
}

type SRpushElement SCompleted

func (c SRpushElement) Element(Element ...string) SRpushElement {
	return SRpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c SRpushElement) Build() SCompleted {
	return SCompleted(c)
}

type SRpushKey SCompleted

func (c SRpushKey) Element(Element ...string) SRpushElement {
	return SRpushElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type SRpushx SCompleted

func (c SRpushx) Key(Key string) SRpushxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRpushxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Rpushx() (c SRpushx) {
	c.cs = append(b.get(), "RPUSHX")
	c.ks = InitSlot
	return
}

type SRpushxElement SCompleted

func (c SRpushxElement) Element(Element ...string) SRpushxElement {
	return SRpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

func (c SRpushxElement) Build() SCompleted {
	return SCompleted(c)
}

type SRpushxKey SCompleted

func (c SRpushxKey) Element(Element ...string) SRpushxElement {
	return SRpushxElement{cs: append(c.cs, Element...), cf: c.cf, ks: c.ks}
}

type SSadd SCompleted

func (c SSadd) Key(Key string) SSaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sadd() (c SSadd) {
	c.cs = append(b.get(), "SADD")
	c.ks = InitSlot
	return
}

type SSaddKey SCompleted

func (c SSaddKey) Member(Member ...string) SSaddMember {
	return SSaddMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SSaddMember SCompleted

func (c SSaddMember) Member(Member ...string) SSaddMember {
	return SSaddMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SSaddMember) Build() SCompleted {
	return SCompleted(c)
}

type SSave SCompleted

func (c SSave) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Save() (c SSave) {
	c.cs = append(b.get(), "SAVE")
	c.ks = InitSlot
	return
}

type SScan SCompleted

func (c SScan) Cursor(Cursor int64) SScanCursor {
	return SScanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Scan() (c SScan) {
	c.cs = append(b.get(), "SCAN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SScanCount SCompleted

func (c SScanCount) Type(Type string) SScanType {
	return SScanType{cs: append(c.cs, "TYPE", Type), cf: c.cf, ks: c.ks}
}

func (c SScanCount) Build() SCompleted {
	return SCompleted(c)
}

type SScanCursor SCompleted

func (c SScanCursor) Match(Pattern string) SScanMatch {
	return SScanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c SScanCursor) Count(Count int64) SScanCount {
	return SScanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SScanCursor) Type(Type string) SScanType {
	return SScanType{cs: append(c.cs, "TYPE", Type), cf: c.cf, ks: c.ks}
}

func (c SScanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SScanMatch SCompleted

func (c SScanMatch) Count(Count int64) SScanCount {
	return SScanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SScanMatch) Type(Type string) SScanType {
	return SScanType{cs: append(c.cs, "TYPE", Type), cf: c.cf, ks: c.ks}
}

func (c SScanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SScanType SCompleted

func (c SScanType) Build() SCompleted {
	return SCompleted(c)
}

type SScard SCompleted

func (c SScard) Key(Key string) SScardKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SScardKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Scard() (c SScard) {
	c.cs = append(b.get(), "SCARD")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SScardKey SCompleted

func (c SScardKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SScardKey) Cache() SCacheable {
	return SCacheable(c)
}

type SScriptDebug SCompleted

func (c SScriptDebug) Yes() SScriptDebugModeYes {
	return SScriptDebugModeYes{cs: append(c.cs, "YES"), cf: c.cf, ks: c.ks}
}

func (c SScriptDebug) Sync() SScriptDebugModeSync {
	return SScriptDebugModeSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c SScriptDebug) No() SScriptDebugModeNo {
	return SScriptDebugModeNo{cs: append(c.cs, "NO"), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ScriptDebug() (c SScriptDebug) {
	c.cs = append(b.get(), "SCRIPT", "DEBUG")
	c.ks = InitSlot
	return
}

type SScriptDebugModeNo SCompleted

func (c SScriptDebugModeNo) Build() SCompleted {
	return SCompleted(c)
}

type SScriptDebugModeSync SCompleted

func (c SScriptDebugModeSync) Build() SCompleted {
	return SCompleted(c)
}

type SScriptDebugModeYes SCompleted

func (c SScriptDebugModeYes) Build() SCompleted {
	return SCompleted(c)
}

type SScriptExists SCompleted

func (c SScriptExists) Sha1(Sha1 ...string) SScriptExistsSha1 {
	return SScriptExistsSha1{cs: append(c.cs, Sha1...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ScriptExists() (c SScriptExists) {
	c.cs = append(b.get(), "SCRIPT", "EXISTS")
	c.ks = InitSlot
	return
}

type SScriptExistsSha1 SCompleted

func (c SScriptExistsSha1) Sha1(Sha1 ...string) SScriptExistsSha1 {
	return SScriptExistsSha1{cs: append(c.cs, Sha1...), cf: c.cf, ks: c.ks}
}

func (c SScriptExistsSha1) Build() SCompleted {
	return SCompleted(c)
}

type SScriptFlush SCompleted

func (c SScriptFlush) Async() SScriptFlushAsyncAsync {
	return SScriptFlushAsyncAsync{cs: append(c.cs, "ASYNC"), cf: c.cf, ks: c.ks}
}

func (c SScriptFlush) Sync() SScriptFlushAsyncSync {
	return SScriptFlushAsyncSync{cs: append(c.cs, "SYNC"), cf: c.cf, ks: c.ks}
}

func (c SScriptFlush) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ScriptFlush() (c SScriptFlush) {
	c.cs = append(b.get(), "SCRIPT", "FLUSH")
	c.ks = InitSlot
	return
}

type SScriptFlushAsyncAsync SCompleted

func (c SScriptFlushAsyncAsync) Build() SCompleted {
	return SCompleted(c)
}

type SScriptFlushAsyncSync SCompleted

func (c SScriptFlushAsyncSync) Build() SCompleted {
	return SCompleted(c)
}

type SScriptKill SCompleted

func (c SScriptKill) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) ScriptKill() (c SScriptKill) {
	c.cs = append(b.get(), "SCRIPT", "KILL")
	c.ks = InitSlot
	return
}

type SScriptLoad SCompleted

func (c SScriptLoad) Script(Script string) SScriptLoadScript {
	return SScriptLoadScript{cs: append(c.cs, Script), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) ScriptLoad() (c SScriptLoad) {
	c.cs = append(b.get(), "SCRIPT", "LOAD")
	c.ks = InitSlot
	return
}

type SScriptLoadScript SCompleted

func (c SScriptLoadScript) Build() SCompleted {
	return SCompleted(c)
}

type SSdiff SCompleted

func (c SSdiff) Key(Key ...string) SSdiffKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sdiff() (c SSdiff) {
	c.cs = append(b.get(), "SDIFF")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSdiffKey SCompleted

func (c SSdiffKey) Key(Key ...string) SSdiffKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SSdiffKey) Build() SCompleted {
	return SCompleted(c)
}

type SSdiffstore SCompleted

func (c SSdiffstore) Destination(Destination string) SSdiffstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSdiffstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sdiffstore() (c SSdiffstore) {
	c.cs = append(b.get(), "SDIFFSTORE")
	c.ks = InitSlot
	return
}

type SSdiffstoreDestination SCompleted

func (c SSdiffstoreDestination) Key(Key ...string) SSdiffstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SSdiffstoreKey SCompleted

func (c SSdiffstoreKey) Key(Key ...string) SSdiffstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SSdiffstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSelect SCompleted

func (c SSelect) Index(Index int64) SSelectIndex {
	return SSelectIndex{cs: append(c.cs, strconv.FormatInt(Index, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Select() (c SSelect) {
	c.cs = append(b.get(), "SELECT")
	c.ks = InitSlot
	return
}

type SSelectIndex SCompleted

func (c SSelectIndex) Build() SCompleted {
	return SCompleted(c)
}

type SSet SCompleted

func (c SSet) Key(Key string) SSetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Set() (c SSet) {
	c.cs = append(b.get(), "SET")
	c.ks = InitSlot
	return
}

type SSetConditionNx SCompleted

func (c SSetConditionNx) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SSetConditionXx SCompleted

func (c SSetConditionXx) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationEx SCompleted

func (c SSetExpirationEx) Nx() SSetConditionNx {
	return SSetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationEx) Xx() SSetConditionXx {
	return SSetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationEx) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationEx) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationExat SCompleted

func (c SSetExpirationExat) Nx() SSetConditionNx {
	return SSetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationExat) Xx() SSetConditionXx {
	return SSetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationExat) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationExat) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationKeepttl SCompleted

func (c SSetExpirationKeepttl) Nx() SSetConditionNx {
	return SSetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationKeepttl) Xx() SSetConditionXx {
	return SSetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationKeepttl) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationKeepttl) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationPx SCompleted

func (c SSetExpirationPx) Nx() SSetConditionNx {
	return SSetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationPx) Xx() SSetConditionXx {
	return SSetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationPx) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationPx) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationPxat SCompleted

func (c SSetExpirationPxat) Nx() SSetConditionNx {
	return SSetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationPxat) Xx() SSetConditionXx {
	return SSetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationPxat) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetExpirationPxat) Build() SCompleted {
	return SCompleted(c)
}

type SSetGetGet SCompleted

func (c SSetGetGet) Build() SCompleted {
	return SCompleted(c)
}

type SSetKey SCompleted

func (c SSetKey) Value(Value string) SSetValue {
	return SSetValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SSetValue SCompleted

func (c SSetValue) Ex(Seconds int64) SSetExpirationEx {
	return SSetExpirationEx{cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Px(Milliseconds int64) SSetExpirationPx {
	return SSetExpirationPx{cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Exat(Timestamp int64) SSetExpirationExat {
	return SSetExpirationExat{cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Pxat(Millisecondstimestamp int64) SSetExpirationPxat {
	return SSetExpirationPxat{cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10)), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Keepttl() SSetExpirationKeepttl {
	return SSetExpirationKeepttl{cs: append(c.cs, "KEEPTTL"), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Nx() SSetConditionNx {
	return SSetConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Xx() SSetConditionXx {
	return SSetConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Get() SSetGetGet {
	return SSetGetGet{cs: append(c.cs, "GET"), cf: c.cf, ks: c.ks}
}

func (c SSetValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetbit SCompleted

func (c SSetbit) Key(Key string) SSetbitKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetbitKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Setbit() (c SSetbit) {
	c.cs = append(b.get(), "SETBIT")
	c.ks = InitSlot
	return
}

type SSetbitKey SCompleted

func (c SSetbitKey) Offset(Offset int64) SSetbitOffset {
	return SSetbitOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type SSetbitOffset SCompleted

func (c SSetbitOffset) Value(Value int64) SSetbitValue {
	return SSetbitValue{cs: append(c.cs, strconv.FormatInt(Value, 10)), cf: c.cf, ks: c.ks}
}

type SSetbitValue SCompleted

func (c SSetbitValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetex SCompleted

func (c SSetex) Key(Key string) SSetexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Setex() (c SSetex) {
	c.cs = append(b.get(), "SETEX")
	c.ks = InitSlot
	return
}

type SSetexKey SCompleted

func (c SSetexKey) Seconds(Seconds int64) SSetexSeconds {
	return SSetexSeconds{cs: append(c.cs, strconv.FormatInt(Seconds, 10)), cf: c.cf, ks: c.ks}
}

type SSetexSeconds SCompleted

func (c SSetexSeconds) Value(Value string) SSetexValue {
	return SSetexValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SSetexValue SCompleted

func (c SSetexValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetnx SCompleted

func (c SSetnx) Key(Key string) SSetnxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetnxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Setnx() (c SSetnx) {
	c.cs = append(b.get(), "SETNX")
	c.ks = InitSlot
	return
}

type SSetnxKey SCompleted

func (c SSetnxKey) Value(Value string) SSetnxValue {
	return SSetnxValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SSetnxValue SCompleted

func (c SSetnxValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetrange SCompleted

func (c SSetrange) Key(Key string) SSetrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Setrange() (c SSetrange) {
	c.cs = append(b.get(), "SETRANGE")
	c.ks = InitSlot
	return
}

type SSetrangeKey SCompleted

func (c SSetrangeKey) Offset(Offset int64) SSetrangeOffset {
	return SSetrangeOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10)), cf: c.cf, ks: c.ks}
}

type SSetrangeOffset SCompleted

func (c SSetrangeOffset) Value(Value string) SSetrangeValue {
	return SSetrangeValue{cs: append(c.cs, Value), cf: c.cf, ks: c.ks}
}

type SSetrangeValue SCompleted

func (c SSetrangeValue) Build() SCompleted {
	return SCompleted(c)
}

type SShutdown SCompleted

func (c SShutdown) Nosave() SShutdownSaveModeNosave {
	return SShutdownSaveModeNosave{cs: append(c.cs, "NOSAVE"), cf: c.cf, ks: c.ks}
}

func (c SShutdown) Save() SShutdownSaveModeSave {
	return SShutdownSaveModeSave{cs: append(c.cs, "SAVE"), cf: c.cf, ks: c.ks}
}

func (c SShutdown) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Shutdown() (c SShutdown) {
	c.cs = append(b.get(), "SHUTDOWN")
	c.ks = InitSlot
	return
}

type SShutdownSaveModeNosave SCompleted

func (c SShutdownSaveModeNosave) Build() SCompleted {
	return SCompleted(c)
}

type SShutdownSaveModeSave SCompleted

func (c SShutdownSaveModeSave) Build() SCompleted {
	return SCompleted(c)
}

type SSinter SCompleted

func (c SSinter) Key(Key ...string) SSinterKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sinter() (c SSinter) {
	c.cs = append(b.get(), "SINTER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSinterKey SCompleted

func (c SSinterKey) Key(Key ...string) SSinterKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SSinterKey) Build() SCompleted {
	return SCompleted(c)
}

type SSintercard SCompleted

func (c SSintercard) Key(Key ...string) SSintercardKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sintercard() (c SSintercard) {
	c.cs = append(b.get(), "SINTERCARD")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSintercardKey SCompleted

func (c SSintercardKey) Key(Key ...string) SSintercardKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SSintercardKey) Build() SCompleted {
	return SCompleted(c)
}

type SSinterstore SCompleted

func (c SSinterstore) Destination(Destination string) SSinterstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSinterstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sinterstore() (c SSinterstore) {
	c.cs = append(b.get(), "SINTERSTORE")
	c.ks = InitSlot
	return
}

type SSinterstoreDestination SCompleted

func (c SSinterstoreDestination) Key(Key ...string) SSinterstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SSinterstoreKey SCompleted

func (c SSinterstoreKey) Key(Key ...string) SSinterstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SSinterstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSismember SCompleted

func (c SSismember) Key(Key string) SSismemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSismemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sismember() (c SSismember) {
	c.cs = append(b.get(), "SISMEMBER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSismemberKey SCompleted

func (c SSismemberKey) Member(Member string) SSismemberMember {
	return SSismemberMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SSismemberMember SCompleted

func (c SSismemberMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SSismemberMember) Cache() SCacheable {
	return SCacheable(c)
}

type SSlaveof SCompleted

func (c SSlaveof) Host(Host string) SSlaveofHost {
	return SSlaveofHost{cs: append(c.cs, Host), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Slaveof() (c SSlaveof) {
	c.cs = append(b.get(), "SLAVEOF")
	c.ks = InitSlot
	return
}

type SSlaveofHost SCompleted

func (c SSlaveofHost) Port(Port string) SSlaveofPort {
	return SSlaveofPort{cs: append(c.cs, Port), cf: c.cf, ks: c.ks}
}

type SSlaveofPort SCompleted

func (c SSlaveofPort) Build() SCompleted {
	return SCompleted(c)
}

type SSlowlog SCompleted

func (c SSlowlog) Subcommand(Subcommand string) SSlowlogSubcommand {
	return SSlowlogSubcommand{cs: append(c.cs, Subcommand), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Slowlog() (c SSlowlog) {
	c.cs = append(b.get(), "SLOWLOG")
	c.ks = InitSlot
	return
}

type SSlowlogArgument SCompleted

func (c SSlowlogArgument) Build() SCompleted {
	return SCompleted(c)
}

type SSlowlogSubcommand SCompleted

func (c SSlowlogSubcommand) Argument(Argument string) SSlowlogArgument {
	return SSlowlogArgument{cs: append(c.cs, Argument), cf: c.cf, ks: c.ks}
}

func (c SSlowlogSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SSmembers SCompleted

func (c SSmembers) Key(Key string) SSmembersKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSmembersKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Smembers() (c SSmembers) {
	c.cs = append(b.get(), "SMEMBERS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSmembersKey SCompleted

func (c SSmembersKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SSmembersKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSmismember SCompleted

func (c SSmismember) Key(Key string) SSmismemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSmismemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Smismember() (c SSmismember) {
	c.cs = append(b.get(), "SMISMEMBER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSmismemberKey SCompleted

func (c SSmismemberKey) Member(Member ...string) SSmismemberMember {
	return SSmismemberMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SSmismemberMember SCompleted

func (c SSmismemberMember) Member(Member ...string) SSmismemberMember {
	return SSmismemberMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SSmismemberMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SSmismemberMember) Cache() SCacheable {
	return SCacheable(c)
}

type SSmove SCompleted

func (c SSmove) Source(Source string) SSmoveSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SSmoveSource{cs: append(c.cs, Source), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Smove() (c SSmove) {
	c.cs = append(b.get(), "SMOVE")
	c.ks = InitSlot
	return
}

type SSmoveDestination SCompleted

func (c SSmoveDestination) Member(Member string) SSmoveMember {
	return SSmoveMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SSmoveMember SCompleted

func (c SSmoveMember) Build() SCompleted {
	return SCompleted(c)
}

type SSmoveSource SCompleted

func (c SSmoveSource) Destination(Destination string) SSmoveDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSmoveDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

type SSort SCompleted

func (c SSort) Key(Key string) SSortKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSortKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sort() (c SSort) {
	c.cs = append(b.get(), "SORT")
	c.ks = InitSlot
	return
}

type SSortBy SCompleted

func (c SSortBy) Limit(Offset int64, Count int64) SSortLimit {
	return SSortLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSortBy) Get(Pattern ...string) SSortGet {
	c.cs = append(c.cs, "GET")
	return SSortGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SSortBy) Asc() SSortOrderAsc {
	return SSortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortBy) Desc() SSortOrderDesc {
	return SSortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortBy) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortBy) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SSortBy) Build() SCompleted {
	return SCompleted(c)
}

type SSortGet SCompleted

func (c SSortGet) Asc() SSortOrderAsc {
	return SSortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortGet) Desc() SSortOrderDesc {
	return SSortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortGet) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortGet) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SSortGet) Get(Get ...string) SSortGet {
	return SSortGet{cs: append(c.cs, Get...), cf: c.cf, ks: c.ks}
}

func (c SSortGet) Build() SCompleted {
	return SCompleted(c)
}

type SSortKey SCompleted

func (c SSortKey) By(Pattern string) SSortBy {
	return SSortBy{cs: append(c.cs, "BY", Pattern), cf: c.cf, ks: c.ks}
}

func (c SSortKey) Limit(Offset int64, Count int64) SSortLimit {
	return SSortLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSortKey) Get(Pattern ...string) SSortGet {
	c.cs = append(c.cs, "GET")
	return SSortGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SSortKey) Asc() SSortOrderAsc {
	return SSortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortKey) Desc() SSortOrderDesc {
	return SSortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortKey) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortKey) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SSortKey) Build() SCompleted {
	return SCompleted(c)
}

type SSortLimit SCompleted

func (c SSortLimit) Get(Pattern ...string) SSortGet {
	c.cs = append(c.cs, "GET")
	return SSortGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SSortLimit) Asc() SSortOrderAsc {
	return SSortOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortLimit) Desc() SSortOrderDesc {
	return SSortOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortLimit) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortLimit) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SSortLimit) Build() SCompleted {
	return SCompleted(c)
}

type SSortOrderAsc SCompleted

func (c SSortOrderAsc) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortOrderAsc) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SSortOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SSortOrderDesc SCompleted

func (c SSortOrderDesc) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortOrderDesc) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SSortOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SSortRo SCompleted

func (c SSortRo) Key(Key string) SSortRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSortRoKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) SortRo() (c SSortRo) {
	c.cs = append(b.get(), "SORT_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSortRoBy SCompleted

func (c SSortRoBy) Limit(Offset int64, Count int64) SSortRoLimit {
	return SSortRoLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSortRoBy) Get(Pattern ...string) SSortRoGet {
	c.cs = append(c.cs, "GET")
	return SSortRoGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SSortRoBy) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoBy) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoBy) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortRoBy) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoBy) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoGet SCompleted

func (c SSortRoGet) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoGet) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoGet) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortRoGet) Get(Get ...string) SSortRoGet {
	return SSortRoGet{cs: append(c.cs, Get...), cf: c.cf, ks: c.ks}
}

func (c SSortRoGet) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoGet) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoKey SCompleted

func (c SSortRoKey) By(Pattern string) SSortRoBy {
	return SSortRoBy{cs: append(c.cs, "BY", Pattern), cf: c.cf, ks: c.ks}
}

func (c SSortRoKey) Limit(Offset int64, Count int64) SSortRoLimit {
	return SSortRoLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSortRoKey) Get(Pattern ...string) SSortRoGet {
	c.cs = append(c.cs, "GET")
	return SSortRoGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SSortRoKey) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoKey) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoKey) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortRoKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoLimit SCompleted

func (c SSortRoLimit) Get(Pattern ...string) SSortRoGet {
	c.cs = append(c.cs, "GET")
	return SSortRoGet{cs: append(c.cs, Pattern...), cf: c.cf, ks: c.ks}
}

func (c SSortRoLimit) Asc() SSortRoOrderAsc {
	return SSortRoOrderAsc{cs: append(c.cs, "ASC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoLimit) Desc() SSortRoOrderDesc {
	return SSortRoOrderDesc{cs: append(c.cs, "DESC"), cf: c.cf, ks: c.ks}
}

func (c SSortRoLimit) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortRoLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoOrderAsc SCompleted

func (c SSortRoOrderAsc) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortRoOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoOrderDesc SCompleted

func (c SSortRoOrderDesc) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cs: append(c.cs, "ALPHA"), cf: c.cf, ks: c.ks}
}

func (c SSortRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoSortingAlpha SCompleted

func (c SSortRoSortingAlpha) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoSortingAlpha) Cache() SCacheable {
	return SCacheable(c)
}

type SSortSortingAlpha SCompleted

func (c SSortSortingAlpha) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cs: append(c.cs, "STORE", Destination), cf: c.cf, ks: c.ks}
}

func (c SSortSortingAlpha) Build() SCompleted {
	return SCompleted(c)
}

type SSortStore SCompleted

func (c SSortStore) Build() SCompleted {
	return SCompleted(c)
}

type SSpop SCompleted

func (c SSpop) Key(Key string) SSpopKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSpopKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Spop() (c SSpop) {
	c.cs = append(b.get(), "SPOP")
	c.ks = InitSlot
	return
}

type SSpopCount SCompleted

func (c SSpopCount) Build() SCompleted {
	return SCompleted(c)
}

type SSpopKey SCompleted

func (c SSpopKey) Count(Count int64) SSpopCount {
	return SSpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SSrandmember SCompleted

func (c SSrandmember) Key(Key string) SSrandmemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSrandmemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Srandmember() (c SSrandmember) {
	c.cs = append(b.get(), "SRANDMEMBER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSrandmemberCount SCompleted

func (c SSrandmemberCount) Build() SCompleted {
	return SCompleted(c)
}

type SSrandmemberKey SCompleted

func (c SSrandmemberKey) Count(Count int64) SSrandmemberCount {
	return SSrandmemberCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSrandmemberKey) Build() SCompleted {
	return SCompleted(c)
}

type SSrem SCompleted

func (c SSrem) Key(Key string) SSremKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSremKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Srem() (c SSrem) {
	c.cs = append(b.get(), "SREM")
	c.ks = InitSlot
	return
}

type SSremKey SCompleted

func (c SSremKey) Member(Member ...string) SSremMember {
	return SSremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SSremMember SCompleted

func (c SSremMember) Member(Member ...string) SSremMember {
	return SSremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SSremMember) Build() SCompleted {
	return SCompleted(c)
}

type SSscan SCompleted

func (c SSscan) Key(Key string) SSscanKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSscanKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sscan() (c SSscan) {
	c.cs = append(b.get(), "SSCAN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSscanCount SCompleted

func (c SSscanCount) Build() SCompleted {
	return SCompleted(c)
}

type SSscanCursor SCompleted

func (c SSscanCursor) Match(Pattern string) SSscanMatch {
	return SSscanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c SSscanCursor) Count(Count int64) SSscanCount {
	return SSscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SSscanKey SCompleted

func (c SSscanKey) Cursor(Cursor int64) SSscanCursor {
	return SSscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

type SSscanMatch SCompleted

func (c SSscanMatch) Count(Count int64) SSscanCount {
	return SSscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SSscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SStralgo SCompleted

func (c SStralgo) Lcs() SStralgoAlgorithmLcs {
	return SStralgoAlgorithmLcs{cs: append(c.cs, "LCS"), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Stralgo() (c SStralgo) {
	c.cs = append(b.get(), "STRALGO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SStralgoAlgoSpecificArgument SCompleted

func (c SStralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) SStralgoAlgoSpecificArgument {
	return SStralgoAlgoSpecificArgument{cs: append(c.cs, AlgoSpecificArgument...), cf: c.cf, ks: c.ks}
}

func (c SStralgoAlgoSpecificArgument) Build() SCompleted {
	return SCompleted(c)
}

type SStralgoAlgorithmLcs SCompleted

func (c SStralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) SStralgoAlgoSpecificArgument {
	return SStralgoAlgoSpecificArgument{cs: append(c.cs, AlgoSpecificArgument...), cf: c.cf, ks: c.ks}
}

type SStrlen SCompleted

func (c SStrlen) Key(Key string) SStrlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SStrlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Strlen() (c SStrlen) {
	c.cs = append(b.get(), "STRLEN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SStrlenKey SCompleted

func (c SStrlenKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SStrlenKey) Cache() SCacheable {
	return SCacheable(c)
}

type SSubscribe SCompleted

func (c SSubscribe) Channel(Channel ...string) SSubscribeChannel {
	return SSubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Subscribe() (c SSubscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	c.cf = noRetTag
	c.ks = InitSlot
	return
}

type SSubscribeChannel SCompleted

func (c SSubscribeChannel) Channel(Channel ...string) SSubscribeChannel {
	return SSubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (c SSubscribeChannel) Build() SCompleted {
	return SCompleted(c)
}

type SSunion SCompleted

func (c SSunion) Key(Key ...string) SSunionKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sunion() (c SSunion) {
	c.cs = append(b.get(), "SUNION")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSunionKey SCompleted

func (c SSunionKey) Key(Key ...string) SSunionKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SSunionKey) Build() SCompleted {
	return SCompleted(c)
}

type SSunionstore SCompleted

func (c SSunionstore) Destination(Destination string) SSunionstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSunionstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Sunionstore() (c SSunionstore) {
	c.cs = append(b.get(), "SUNIONSTORE")
	c.ks = InitSlot
	return
}

type SSunionstoreDestination SCompleted

func (c SSunionstoreDestination) Key(Key ...string) SSunionstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SSunionstoreKey SCompleted

func (c SSunionstoreKey) Key(Key ...string) SSunionstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SSunionstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSwapdb SCompleted

func (c SSwapdb) Index1(Index1 int64) SSwapdbIndex1 {
	return SSwapdbIndex1{cs: append(c.cs, strconv.FormatInt(Index1, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Swapdb() (c SSwapdb) {
	c.cs = append(b.get(), "SWAPDB")
	c.ks = InitSlot
	return
}

type SSwapdbIndex1 SCompleted

func (c SSwapdbIndex1) Index2(Index2 int64) SSwapdbIndex2 {
	return SSwapdbIndex2{cs: append(c.cs, strconv.FormatInt(Index2, 10)), cf: c.cf, ks: c.ks}
}

type SSwapdbIndex2 SCompleted

func (c SSwapdbIndex2) Build() SCompleted {
	return SCompleted(c)
}

type SSync SCompleted

func (c SSync) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Sync() (c SSync) {
	c.cs = append(b.get(), "SYNC")
	c.ks = InitSlot
	return
}

type STime SCompleted

func (c STime) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Time() (c STime) {
	c.cs = append(b.get(), "TIME")
	c.ks = InitSlot
	return
}

type STouch SCompleted

func (c STouch) Key(Key ...string) STouchKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return STouchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Touch() (c STouch) {
	c.cs = append(b.get(), "TOUCH")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type STouchKey SCompleted

func (c STouchKey) Key(Key ...string) STouchKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return STouchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c STouchKey) Build() SCompleted {
	return SCompleted(c)
}

type STtl SCompleted

func (c STtl) Key(Key string) STtlKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return STtlKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Ttl() (c STtl) {
	c.cs = append(b.get(), "TTL")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type STtlKey SCompleted

func (c STtlKey) Build() SCompleted {
	return SCompleted(c)
}

func (c STtlKey) Cache() SCacheable {
	return SCacheable(c)
}

type SType SCompleted

func (c SType) Key(Key string) STypeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return STypeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Type() (c SType) {
	c.cs = append(b.get(), "TYPE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type STypeKey SCompleted

func (c STypeKey) Build() SCompleted {
	return SCompleted(c)
}

func (c STypeKey) Cache() SCacheable {
	return SCacheable(c)
}

type SUnlink SCompleted

func (c SUnlink) Key(Key ...string) SUnlinkKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SUnlinkKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Unlink() (c SUnlink) {
	c.cs = append(b.get(), "UNLINK")
	c.ks = InitSlot
	return
}

type SUnlinkKey SCompleted

func (c SUnlinkKey) Key(Key ...string) SUnlinkKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SUnlinkKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SUnlinkKey) Build() SCompleted {
	return SCompleted(c)
}

type SUnsubscribe SCompleted

func (c SUnsubscribe) Channel(Channel ...string) SUnsubscribeChannel {
	return SUnsubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (c SUnsubscribe) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Unsubscribe() (c SUnsubscribe) {
	c.cs = append(b.get(), "UNSUBSCRIBE")
	c.cf = noRetTag
	c.ks = InitSlot
	return
}

type SUnsubscribeChannel SCompleted

func (c SUnsubscribeChannel) Channel(Channel ...string) SUnsubscribeChannel {
	return SUnsubscribeChannel{cs: append(c.cs, Channel...), cf: c.cf, ks: c.ks}
}

func (c SUnsubscribeChannel) Build() SCompleted {
	return SCompleted(c)
}

type SUnwatch SCompleted

func (c SUnwatch) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Unwatch() (c SUnwatch) {
	c.cs = append(b.get(), "UNWATCH")
	c.ks = InitSlot
	return
}

type SWait SCompleted

func (c SWait) Numreplicas(Numreplicas int64) SWaitNumreplicas {
	return SWaitNumreplicas{cs: append(c.cs, strconv.FormatInt(Numreplicas, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Wait() (c SWait) {
	c.cs = append(b.get(), "WAIT")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SWaitNumreplicas SCompleted

func (c SWaitNumreplicas) Timeout(Timeout int64) SWaitTimeout {
	return SWaitTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10)), cf: c.cf, ks: c.ks}
}

type SWaitTimeout SCompleted

func (c SWaitTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SWatch SCompleted

func (c SWatch) Key(Key ...string) SWatchKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SWatchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Watch() (c SWatch) {
	c.cs = append(b.get(), "WATCH")
	c.ks = InitSlot
	return
}

type SWatchKey SCompleted

func (c SWatchKey) Key(Key ...string) SWatchKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SWatchKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SWatchKey) Build() SCompleted {
	return SCompleted(c)
}

type SXack SCompleted

func (c SXack) Key(Key string) SXackKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXackKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xack() (c SXack) {
	c.cs = append(b.get(), "XACK")
	c.ks = InitSlot
	return
}

type SXackGroup SCompleted

func (c SXackGroup) Id(Id ...string) SXackId {
	return SXackId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

type SXackId SCompleted

func (c SXackId) Id(Id ...string) SXackId {
	return SXackId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c SXackId) Build() SCompleted {
	return SCompleted(c)
}

type SXackKey SCompleted

func (c SXackKey) Group(Group string) SXackGroup {
	return SXackGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type SXadd SCompleted

func (c SXadd) Key(Key string) SXaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xadd() (c SXadd) {
	c.cs = append(b.get(), "XADD")
	c.ks = InitSlot
	return
}

type SXaddFieldValue SCompleted

func (c SXaddFieldValue) FieldValue(Field string, Value string) SXaddFieldValue {
	return SXaddFieldValue{cs: append(c.cs, Field, Value), cf: c.cf, ks: c.ks}
}

func (c SXaddFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SXaddId SCompleted

func (c SXaddId) FieldValue() SXaddFieldValue {
	return SXaddFieldValue{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SXaddKey SCompleted

func (c SXaddKey) Nomkstream() SXaddNomkstream {
	return SXaddNomkstream{cs: append(c.cs, "NOMKSTREAM"), cf: c.cf, ks: c.ks}
}

func (c SXaddKey) Maxlen() SXaddTrimStrategyMaxlen {
	return SXaddTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN"), cf: c.cf, ks: c.ks}
}

func (c SXaddKey) Minid() SXaddTrimStrategyMinid {
	return SXaddTrimStrategyMinid{cs: append(c.cs, "MINID"), cf: c.cf, ks: c.ks}
}

func (c SXaddKey) Id(Id string) SXaddId {
	return SXaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type SXaddNomkstream SCompleted

func (c SXaddNomkstream) Maxlen() SXaddTrimStrategyMaxlen {
	return SXaddTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN"), cf: c.cf, ks: c.ks}
}

func (c SXaddNomkstream) Minid() SXaddTrimStrategyMinid {
	return SXaddTrimStrategyMinid{cs: append(c.cs, "MINID"), cf: c.cf, ks: c.ks}
}

func (c SXaddNomkstream) Id(Id string) SXaddId {
	return SXaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type SXaddTrimLimit SCompleted

func (c SXaddTrimLimit) Id(Id string) SXaddId {
	return SXaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type SXaddTrimOperatorAlmost SCompleted

func (c SXaddTrimOperatorAlmost) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXaddTrimOperatorExact SCompleted

func (c SXaddTrimOperatorExact) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXaddTrimStrategyMaxlen SCompleted

func (c SXaddTrimStrategyMaxlen) Exact() SXaddTrimOperatorExact {
	return SXaddTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c SXaddTrimStrategyMaxlen) Almost() SXaddTrimOperatorAlmost {
	return SXaddTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c SXaddTrimStrategyMaxlen) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXaddTrimStrategyMinid SCompleted

func (c SXaddTrimStrategyMinid) Exact() SXaddTrimOperatorExact {
	return SXaddTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c SXaddTrimStrategyMinid) Almost() SXaddTrimOperatorAlmost {
	return SXaddTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c SXaddTrimStrategyMinid) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXaddTrimThreshold SCompleted

func (c SXaddTrimThreshold) Limit(Count int64) SXaddTrimLimit {
	return SXaddTrimLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXaddTrimThreshold) Id(Id string) SXaddId {
	return SXaddId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type SXautoclaim SCompleted

func (c SXautoclaim) Key(Key string) SXautoclaimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXautoclaimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xautoclaim() (c SXautoclaim) {
	c.cs = append(b.get(), "XAUTOCLAIM")
	c.ks = InitSlot
	return
}

type SXautoclaimConsumer SCompleted

func (c SXautoclaimConsumer) MinIdleTime(MinIdleTime string) SXautoclaimMinIdleTime {
	return SXautoclaimMinIdleTime{cs: append(c.cs, MinIdleTime), cf: c.cf, ks: c.ks}
}

type SXautoclaimCount SCompleted

func (c SXautoclaimCount) Justid() SXautoclaimJustidJustid {
	return SXautoclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c SXautoclaimCount) Build() SCompleted {
	return SCompleted(c)
}

type SXautoclaimGroup SCompleted

func (c SXautoclaimGroup) Consumer(Consumer string) SXautoclaimConsumer {
	return SXautoclaimConsumer{cs: append(c.cs, Consumer), cf: c.cf, ks: c.ks}
}

type SXautoclaimJustidJustid SCompleted

func (c SXautoclaimJustidJustid) Build() SCompleted {
	return SCompleted(c)
}

type SXautoclaimKey SCompleted

func (c SXautoclaimKey) Group(Group string) SXautoclaimGroup {
	return SXautoclaimGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type SXautoclaimMinIdleTime SCompleted

func (c SXautoclaimMinIdleTime) Start(Start string) SXautoclaimStart {
	return SXautoclaimStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type SXautoclaimStart SCompleted

func (c SXautoclaimStart) Count(Count int64) SXautoclaimCount {
	return SXautoclaimCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXautoclaimStart) Justid() SXautoclaimJustidJustid {
	return SXautoclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c SXautoclaimStart) Build() SCompleted {
	return SCompleted(c)
}

type SXclaim SCompleted

func (c SXclaim) Key(Key string) SXclaimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXclaimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xclaim() (c SXclaim) {
	c.cs = append(b.get(), "XCLAIM")
	c.ks = InitSlot
	return
}

type SXclaimConsumer SCompleted

func (c SXclaimConsumer) MinIdleTime(MinIdleTime string) SXclaimMinIdleTime {
	return SXclaimMinIdleTime{cs: append(c.cs, MinIdleTime), cf: c.cf, ks: c.ks}
}

type SXclaimForceForce SCompleted

func (c SXclaimForceForce) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c SXclaimForceForce) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimGroup SCompleted

func (c SXclaimGroup) Consumer(Consumer string) SXclaimConsumer {
	return SXclaimConsumer{cs: append(c.cs, Consumer), cf: c.cf, ks: c.ks}
}

type SXclaimId SCompleted

func (c SXclaimId) Idle(Ms int64) SXclaimIdle {
	return SXclaimIdle{cs: append(c.cs, "IDLE", strconv.FormatInt(Ms, 10)), cf: c.cf, ks: c.ks}
}

func (c SXclaimId) Time(MsUnixTime int64) SXclaimTime {
	return SXclaimTime{cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10)), cf: c.cf, ks: c.ks}
}

func (c SXclaimId) Retrycount(Count int64) SXclaimRetrycount {
	return SXclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXclaimId) Force() SXclaimForceForce {
	return SXclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c SXclaimId) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c SXclaimId) Id(Id ...string) SXclaimId {
	return SXclaimId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c SXclaimId) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimIdle SCompleted

func (c SXclaimIdle) Time(MsUnixTime int64) SXclaimTime {
	return SXclaimTime{cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10)), cf: c.cf, ks: c.ks}
}

func (c SXclaimIdle) Retrycount(Count int64) SXclaimRetrycount {
	return SXclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXclaimIdle) Force() SXclaimForceForce {
	return SXclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c SXclaimIdle) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c SXclaimIdle) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimJustidJustid SCompleted

func (c SXclaimJustidJustid) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimKey SCompleted

func (c SXclaimKey) Group(Group string) SXclaimGroup {
	return SXclaimGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type SXclaimMinIdleTime SCompleted

func (c SXclaimMinIdleTime) Id(Id ...string) SXclaimId {
	return SXclaimId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

type SXclaimRetrycount SCompleted

func (c SXclaimRetrycount) Force() SXclaimForceForce {
	return SXclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c SXclaimRetrycount) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c SXclaimRetrycount) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimTime SCompleted

func (c SXclaimTime) Retrycount(Count int64) SXclaimRetrycount {
	return SXclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXclaimTime) Force() SXclaimForceForce {
	return SXclaimForceForce{cs: append(c.cs, "FORCE"), cf: c.cf, ks: c.ks}
}

func (c SXclaimTime) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cs: append(c.cs, "JUSTID"), cf: c.cf, ks: c.ks}
}

func (c SXclaimTime) Build() SCompleted {
	return SCompleted(c)
}

type SXdel SCompleted

func (c SXdel) Key(Key string) SXdelKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXdelKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xdel() (c SXdel) {
	c.cs = append(b.get(), "XDEL")
	c.ks = InitSlot
	return
}

type SXdelId SCompleted

func (c SXdelId) Id(Id ...string) SXdelId {
	return SXdelId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c SXdelId) Build() SCompleted {
	return SCompleted(c)
}

type SXdelKey SCompleted

func (c SXdelKey) Id(Id ...string) SXdelId {
	return SXdelId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

type SXgroup SCompleted

func (c SXgroup) Create(Key string, Groupname string) SXgroupCreateCreate {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateCreate{cs: append(c.cs, "CREATE", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroup) Setid(Key string, Groupname string) SXgroupSetidSetid {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroup) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroup) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroup) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroup) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Xgroup() (c SXgroup) {
	c.cs = append(b.get(), "XGROUP")
	c.ks = InitSlot
	return
}

type SXgroupCreateCreate SCompleted

func (c SXgroupCreateCreate) Id(Id string) SXgroupCreateId {
	return SXgroupCreateId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type SXgroupCreateId SCompleted

func (c SXgroupCreateId) Mkstream() SXgroupCreateMkstream {
	return SXgroupCreateMkstream{cs: append(c.cs, "MKSTREAM"), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateId) Setid(Key string, Groupname string) SXgroupSetidSetid {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateId) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateId) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateId) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateId) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupCreateMkstream SCompleted

func (c SXgroupCreateMkstream) Setid(Key string, Groupname string) SXgroupSetidSetid {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateMkstream) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateMkstream) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateMkstream) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateMkstream) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupCreateconsumer SCompleted

func (c SXgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupCreateconsumer) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupDelconsumer SCompleted

func (c SXgroupDelconsumer) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupDestroy SCompleted

func (c SXgroupDestroy) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupDestroy) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupDestroy) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupSetidId SCompleted

func (c SXgroupSetidId) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXgroupSetidId) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupSetidId) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername), cf: c.cf, ks: c.ks}
}

func (c SXgroupSetidId) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupSetidSetid SCompleted

func (c SXgroupSetidSetid) Id(Id string) SXgroupSetidId {
	return SXgroupSetidId{cs: append(c.cs, Id), cf: c.cf, ks: c.ks}
}

type SXinfo SCompleted

func (c SXinfo) Consumers(Key string, Groupname string) SXinfoConsumers {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoConsumers{cs: append(c.cs, "CONSUMERS", Key, Groupname), cf: c.cf, ks: c.ks}
}

func (c SXinfo) Groups(Key string) SXinfoGroups {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoGroups{cs: append(c.cs, "GROUPS", Key), cf: c.cf, ks: c.ks}
}

func (c SXinfo) Stream(Key string) SXinfoStream {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoStream{cs: append(c.cs, "STREAM", Key), cf: c.cf, ks: c.ks}
}

func (c SXinfo) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c SXinfo) Build() SCompleted {
	return SCompleted(c)
}

func (b *SBuilder) Xinfo() (c SXinfo) {
	c.cs = append(b.get(), "XINFO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SXinfoConsumers SCompleted

func (c SXinfoConsumers) Groups(Key string) SXinfoGroups {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoGroups{cs: append(c.cs, "GROUPS", Key), cf: c.cf, ks: c.ks}
}

func (c SXinfoConsumers) Stream(Key string) SXinfoStream {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoStream{cs: append(c.cs, "STREAM", Key), cf: c.cf, ks: c.ks}
}

func (c SXinfoConsumers) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c SXinfoConsumers) Build() SCompleted {
	return SCompleted(c)
}

type SXinfoGroups SCompleted

func (c SXinfoGroups) Stream(Key string) SXinfoStream {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoStream{cs: append(c.cs, "STREAM", Key), cf: c.cf, ks: c.ks}
}

func (c SXinfoGroups) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c SXinfoGroups) Build() SCompleted {
	return SCompleted(c)
}

type SXinfoHelpHelp SCompleted

func (c SXinfoHelpHelp) Build() SCompleted {
	return SCompleted(c)
}

type SXinfoStream SCompleted

func (c SXinfoStream) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cs: append(c.cs, "HELP"), cf: c.cf, ks: c.ks}
}

func (c SXinfoStream) Build() SCompleted {
	return SCompleted(c)
}

type SXlen SCompleted

func (c SXlen) Key(Key string) SXlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXlenKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xlen() (c SXlen) {
	c.cs = append(b.get(), "XLEN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SXlenKey SCompleted

func (c SXlenKey) Build() SCompleted {
	return SCompleted(c)
}

type SXpending SCompleted

func (c SXpending) Key(Key string) SXpendingKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXpendingKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xpending() (c SXpending) {
	c.cs = append(b.get(), "XPENDING")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SXpendingFiltersConsumer SCompleted

func (c SXpendingFiltersConsumer) Build() SCompleted {
	return SCompleted(c)
}

type SXpendingFiltersCount SCompleted

func (c SXpendingFiltersCount) Consumer(Consumer string) SXpendingFiltersConsumer {
	return SXpendingFiltersConsumer{cs: append(c.cs, Consumer), cf: c.cf, ks: c.ks}
}

func (c SXpendingFiltersCount) Build() SCompleted {
	return SCompleted(c)
}

type SXpendingFiltersEnd SCompleted

func (c SXpendingFiltersEnd) Count(Count int64) SXpendingFiltersCount {
	return SXpendingFiltersCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

type SXpendingFiltersIdle SCompleted

func (c SXpendingFiltersIdle) Start(Start string) SXpendingFiltersStart {
	return SXpendingFiltersStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type SXpendingFiltersStart SCompleted

func (c SXpendingFiltersStart) End(End string) SXpendingFiltersEnd {
	return SXpendingFiltersEnd{cs: append(c.cs, End), cf: c.cf, ks: c.ks}
}

type SXpendingGroup SCompleted

func (c SXpendingGroup) Idle(MinIdleTime int64) SXpendingFiltersIdle {
	return SXpendingFiltersIdle{cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10)), cf: c.cf, ks: c.ks}
}

func (c SXpendingGroup) Start(Start string) SXpendingFiltersStart {
	return SXpendingFiltersStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type SXpendingKey SCompleted

func (c SXpendingKey) Group(Group string) SXpendingGroup {
	return SXpendingGroup{cs: append(c.cs, Group), cf: c.cf, ks: c.ks}
}

type SXrange SCompleted

func (c SXrange) Key(Key string) SXrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xrange() (c SXrange) {
	c.cs = append(b.get(), "XRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SXrangeCount SCompleted

func (c SXrangeCount) Build() SCompleted {
	return SCompleted(c)
}

type SXrangeEnd SCompleted

func (c SXrangeEnd) Count(Count int64) SXrangeCount {
	return SXrangeCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXrangeEnd) Build() SCompleted {
	return SCompleted(c)
}

type SXrangeKey SCompleted

func (c SXrangeKey) Start(Start string) SXrangeStart {
	return SXrangeStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type SXrangeStart SCompleted

func (c SXrangeStart) End(End string) SXrangeEnd {
	return SXrangeEnd{cs: append(c.cs, End), cf: c.cf, ks: c.ks}
}

type SXread SCompleted

func (c SXread) Count(Count int64) SXreadCount {
	return SXreadCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXread) Block(Milliseconds int64) SXreadBlock {
	c.cf = blockTag
	return SXreadBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SXread) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xread() (c SXread) {
	c.cs = append(b.get(), "XREAD")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SXreadBlock SCompleted

func (c SXreadBlock) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type SXreadCount SCompleted

func (c SXreadCount) Block(Milliseconds int64) SXreadBlock {
	c.cf = blockTag
	return SXreadBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SXreadCount) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type SXreadId SCompleted

func (c SXreadId) Id(Id ...string) SXreadId {
	return SXreadId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c SXreadId) Build() SCompleted {
	return SCompleted(c)
}

type SXreadKey SCompleted

func (c SXreadKey) Id(Id ...string) SXreadId {
	return SXreadId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c SXreadKey) Key(Key ...string) SXreadKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SXreadStreamsStreams SCompleted

func (c SXreadStreamsStreams) Key(Key ...string) SXreadKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SXreadgroup SCompleted

func (c SXreadgroup) Group(Group string, Consumer string) SXreadgroupGroup {
	return SXreadgroupGroup{cs: append(c.cs, "GROUP", Group, Consumer), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xreadgroup() (c SXreadgroup) {
	c.cs = append(b.get(), "XREADGROUP")
	c.ks = InitSlot
	return
}

type SXreadgroupBlock SCompleted

func (c SXreadgroupBlock) Noack() SXreadgroupNoackNoack {
	return SXreadgroupNoackNoack{cs: append(c.cs, "NOACK"), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupBlock) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type SXreadgroupCount SCompleted

func (c SXreadgroupCount) Block(Milliseconds int64) SXreadgroupBlock {
	c.cf = blockTag
	return SXreadgroupBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupCount) Noack() SXreadgroupNoackNoack {
	return SXreadgroupNoackNoack{cs: append(c.cs, "NOACK"), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupCount) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type SXreadgroupGroup SCompleted

func (c SXreadgroupGroup) Count(Count int64) SXreadgroupCount {
	return SXreadgroupCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupGroup) Block(Milliseconds int64) SXreadgroupBlock {
	c.cf = blockTag
	return SXreadgroupBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10)), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupGroup) Noack() SXreadgroupNoackNoack {
	return SXreadgroupNoackNoack{cs: append(c.cs, "NOACK"), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupGroup) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type SXreadgroupId SCompleted

func (c SXreadgroupId) Id(Id ...string) SXreadgroupId {
	return SXreadgroupId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupId) Build() SCompleted {
	return SCompleted(c)
}

type SXreadgroupKey SCompleted

func (c SXreadgroupKey) Id(Id ...string) SXreadgroupId {
	return SXreadgroupId{cs: append(c.cs, Id...), cf: c.cf, ks: c.ks}
}

func (c SXreadgroupKey) Key(Key ...string) SXreadgroupKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadgroupKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SXreadgroupNoackNoack SCompleted

func (c SXreadgroupNoackNoack) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cs: append(c.cs, "STREAMS"), cf: c.cf, ks: c.ks}
}

type SXreadgroupStreamsStreams SCompleted

func (c SXreadgroupStreamsStreams) Key(Key ...string) SXreadgroupKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadgroupKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SXrevrange SCompleted

func (c SXrevrange) Key(Key string) SXrevrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXrevrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xrevrange() (c SXrevrange) {
	c.cs = append(b.get(), "XREVRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SXrevrangeCount SCompleted

func (c SXrevrangeCount) Build() SCompleted {
	return SCompleted(c)
}

type SXrevrangeEnd SCompleted

func (c SXrevrangeEnd) Start(Start string) SXrevrangeStart {
	return SXrevrangeStart{cs: append(c.cs, Start), cf: c.cf, ks: c.ks}
}

type SXrevrangeKey SCompleted

func (c SXrevrangeKey) End(End string) SXrevrangeEnd {
	return SXrevrangeEnd{cs: append(c.cs, End), cf: c.cf, ks: c.ks}
}

type SXrevrangeStart SCompleted

func (c SXrevrangeStart) Count(Count int64) SXrevrangeCount {
	return SXrevrangeCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXrevrangeStart) Build() SCompleted {
	return SCompleted(c)
}

type SXtrim SCompleted

func (c SXtrim) Key(Key string) SXtrimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXtrimKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Xtrim() (c SXtrim) {
	c.cs = append(b.get(), "XTRIM")
	c.ks = InitSlot
	return
}

type SXtrimKey SCompleted

func (c SXtrimKey) Maxlen() SXtrimTrimStrategyMaxlen {
	return SXtrimTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN"), cf: c.cf, ks: c.ks}
}

func (c SXtrimKey) Minid() SXtrimTrimStrategyMinid {
	return SXtrimTrimStrategyMinid{cs: append(c.cs, "MINID"), cf: c.cf, ks: c.ks}
}

type SXtrimTrimLimit SCompleted

func (c SXtrimTrimLimit) Build() SCompleted {
	return SCompleted(c)
}

type SXtrimTrimOperatorAlmost SCompleted

func (c SXtrimTrimOperatorAlmost) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXtrimTrimOperatorExact SCompleted

func (c SXtrimTrimOperatorExact) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXtrimTrimStrategyMaxlen SCompleted

func (c SXtrimTrimStrategyMaxlen) Exact() SXtrimTrimOperatorExact {
	return SXtrimTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c SXtrimTrimStrategyMaxlen) Almost() SXtrimTrimOperatorAlmost {
	return SXtrimTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c SXtrimTrimStrategyMaxlen) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXtrimTrimStrategyMinid SCompleted

func (c SXtrimTrimStrategyMinid) Exact() SXtrimTrimOperatorExact {
	return SXtrimTrimOperatorExact{cs: append(c.cs, "="), cf: c.cf, ks: c.ks}
}

func (c SXtrimTrimStrategyMinid) Almost() SXtrimTrimOperatorAlmost {
	return SXtrimTrimOperatorAlmost{cs: append(c.cs, "~"), cf: c.cf, ks: c.ks}
}

func (c SXtrimTrimStrategyMinid) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cs: append(c.cs, Threshold), cf: c.cf, ks: c.ks}
}

type SXtrimTrimThreshold SCompleted

func (c SXtrimTrimThreshold) Limit(Count int64) SXtrimTrimLimit {
	return SXtrimTrimLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SXtrimTrimThreshold) Build() SCompleted {
	return SCompleted(c)
}

type SZadd SCompleted

func (c SZadd) Key(Key string) SZaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZaddKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zadd() (c SZadd) {
	c.cs = append(b.get(), "ZADD")
	c.ks = InitSlot
	return
}

type SZaddChangeCh SCompleted

func (c SZaddChangeCh) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c SZaddChangeCh) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SZaddComparisonGt SCompleted

func (c SZaddComparisonGt) Ch() SZaddChangeCh {
	return SZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SZaddComparisonGt) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c SZaddComparisonGt) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SZaddComparisonLt SCompleted

func (c SZaddComparisonLt) Ch() SZaddChangeCh {
	return SZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SZaddComparisonLt) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c SZaddComparisonLt) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SZaddConditionNx SCompleted

func (c SZaddConditionNx) Gt() SZaddComparisonGt {
	return SZaddComparisonGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionNx) Lt() SZaddComparisonLt {
	return SZaddComparisonLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionNx) Ch() SZaddChangeCh {
	return SZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionNx) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionNx) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SZaddConditionXx SCompleted

func (c SZaddConditionXx) Gt() SZaddComparisonGt {
	return SZaddComparisonGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionXx) Lt() SZaddComparisonLt {
	return SZaddComparisonLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionXx) Ch() SZaddChangeCh {
	return SZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionXx) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c SZaddConditionXx) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SZaddIncrementIncr SCompleted

func (c SZaddIncrementIncr) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SZaddKey SCompleted

func (c SZaddKey) Nx() SZaddConditionNx {
	return SZaddConditionNx{cs: append(c.cs, "NX"), cf: c.cf, ks: c.ks}
}

func (c SZaddKey) Xx() SZaddConditionXx {
	return SZaddConditionXx{cs: append(c.cs, "XX"), cf: c.cf, ks: c.ks}
}

func (c SZaddKey) Gt() SZaddComparisonGt {
	return SZaddComparisonGt{cs: append(c.cs, "GT"), cf: c.cf, ks: c.ks}
}

func (c SZaddKey) Lt() SZaddComparisonLt {
	return SZaddComparisonLt{cs: append(c.cs, "LT"), cf: c.cf, ks: c.ks}
}

func (c SZaddKey) Ch() SZaddChangeCh {
	return SZaddChangeCh{cs: append(c.cs, "CH"), cf: c.cf, ks: c.ks}
}

func (c SZaddKey) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cs: append(c.cs, "INCR"), cf: c.cf, ks: c.ks}
}

func (c SZaddKey) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cs: c.cs, cf: c.cf, ks: c.ks}
}

type SZaddScoreMember SCompleted

func (c SZaddScoreMember) ScoreMember(Score float64, Member string) SZaddScoreMember {
	return SZaddScoreMember{cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member), cf: c.cf, ks: c.ks}
}

func (c SZaddScoreMember) Build() SCompleted {
	return SCompleted(c)
}

type SZcard SCompleted

func (c SZcard) Key(Key string) SZcardKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZcardKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zcard() (c SZcard) {
	c.cs = append(b.get(), "ZCARD")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZcardKey SCompleted

func (c SZcardKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SZcardKey) Cache() SCacheable {
	return SCacheable(c)
}

type SZcount SCompleted

func (c SZcount) Key(Key string) SZcountKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZcountKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zcount() (c SZcount) {
	c.cs = append(b.get(), "ZCOUNT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZcountKey SCompleted

func (c SZcountKey) Min(Min float64) SZcountMin {
	return SZcountMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZcountMax SCompleted

func (c SZcountMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZcountMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZcountMin SCompleted

func (c SZcountMin) Max(Max float64) SZcountMax {
	return SZcountMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZdiff SCompleted

func (c SZdiff) Numkeys(Numkeys int64) SZdiffNumkeys {
	return SZdiffNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zdiff() (c SZdiff) {
	c.cs = append(b.get(), "ZDIFF")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZdiffKey SCompleted

func (c SZdiffKey) Withscores() SZdiffWithscoresWithscores {
	return SZdiffWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZdiffKey) Key(Key ...string) SZdiffKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SZdiffKey) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffNumkeys SCompleted

func (c SZdiffNumkeys) Key(Key ...string) SZdiffKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SZdiffWithscoresWithscores SCompleted

func (c SZdiffWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffstore SCompleted

func (c SZdiffstore) Destination(Destination string) SZdiffstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SZdiffstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zdiffstore() (c SZdiffstore) {
	c.cs = append(b.get(), "ZDIFFSTORE")
	c.ks = InitSlot
	return
}

type SZdiffstoreDestination SCompleted

func (c SZdiffstoreDestination) Numkeys(Numkeys int64) SZdiffstoreNumkeys {
	return SZdiffstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SZdiffstoreKey SCompleted

func (c SZdiffstoreKey) Key(Key ...string) SZdiffstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SZdiffstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffstoreNumkeys SCompleted

func (c SZdiffstoreNumkeys) Key(Key ...string) SZdiffstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SZincrby SCompleted

func (c SZincrby) Key(Key string) SZincrbyKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZincrbyKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zincrby() (c SZincrby) {
	c.cs = append(b.get(), "ZINCRBY")
	c.ks = InitSlot
	return
}

type SZincrbyIncrement SCompleted

func (c SZincrbyIncrement) Member(Member string) SZincrbyMember {
	return SZincrbyMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SZincrbyKey SCompleted

func (c SZincrbyKey) Increment(Increment int64) SZincrbyIncrement {
	return SZincrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10)), cf: c.cf, ks: c.ks}
}

type SZincrbyMember SCompleted

func (c SZincrbyMember) Build() SCompleted {
	return SCompleted(c)
}

type SZinter SCompleted

func (c SZinter) Numkeys(Numkeys int64) SZinterNumkeys {
	return SZinterNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zinter() (c SZinter) {
	c.cs = append(b.get(), "ZINTER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZinterAggregateMax SCompleted

func (c SZinterAggregateMax) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZinterAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZinterAggregateMin SCompleted

func (c SZinterAggregateMin) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZinterAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZinterAggregateSum SCompleted

func (c SZinterAggregateSum) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZinterAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZinterKey SCompleted

func (c SZinterKey) Weights(Weight ...int64) SZinterWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZinterKey) Sum() SZinterAggregateSum {
	return SZinterAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZinterKey) Min() SZinterAggregateMin {
	return SZinterAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZinterKey) Max() SZinterAggregateMax {
	return SZinterAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZinterKey) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZinterKey) Key(Key ...string) SZinterKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SZinterKey) Build() SCompleted {
	return SCompleted(c)
}

type SZinterNumkeys SCompleted

func (c SZinterNumkeys) Key(Key ...string) SZinterKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SZinterWeights SCompleted

func (c SZinterWeights) Sum() SZinterAggregateSum {
	return SZinterAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZinterWeights) Min() SZinterAggregateMin {
	return SZinterAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZinterWeights) Max() SZinterAggregateMax {
	return SZinterAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZinterWeights) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZinterWeights) Weights(Weights ...int64) SZinterWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZinterWeights) Build() SCompleted {
	return SCompleted(c)
}

type SZinterWithscoresWithscores SCompleted

func (c SZinterWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZintercard SCompleted

func (c SZintercard) Numkeys(Numkeys int64) SZintercardNumkeys {
	return SZintercardNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zintercard() (c SZintercard) {
	c.cs = append(b.get(), "ZINTERCARD")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZintercardKey SCompleted

func (c SZintercardKey) Key(Key ...string) SZintercardKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SZintercardKey) Build() SCompleted {
	return SCompleted(c)
}

type SZintercardNumkeys SCompleted

func (c SZintercardNumkeys) Key(Key ...string) SZintercardKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZintercardKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SZinterstore SCompleted

func (c SZinterstore) Destination(Destination string) SZinterstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SZinterstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zinterstore() (c SZinterstore) {
	c.cs = append(b.get(), "ZINTERSTORE")
	c.ks = InitSlot
	return
}

type SZinterstoreAggregateMax SCompleted

func (c SZinterstoreAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreAggregateMin SCompleted

func (c SZinterstoreAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreAggregateSum SCompleted

func (c SZinterstoreAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreDestination SCompleted

func (c SZinterstoreDestination) Numkeys(Numkeys int64) SZinterstoreNumkeys {
	return SZinterstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SZinterstoreKey SCompleted

func (c SZinterstoreKey) Weights(Weight ...int64) SZinterstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZinterstoreKey) Sum() SZinterstoreAggregateSum {
	return SZinterstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZinterstoreKey) Min() SZinterstoreAggregateMin {
	return SZinterstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZinterstoreKey) Max() SZinterstoreAggregateMax {
	return SZinterstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZinterstoreKey) Key(Key ...string) SZinterstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SZinterstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreNumkeys SCompleted

func (c SZinterstoreNumkeys) Key(Key ...string) SZinterstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SZinterstoreWeights SCompleted

func (c SZinterstoreWeights) Sum() SZinterstoreAggregateSum {
	return SZinterstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZinterstoreWeights) Min() SZinterstoreAggregateMin {
	return SZinterstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZinterstoreWeights) Max() SZinterstoreAggregateMax {
	return SZinterstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZinterstoreWeights) Weights(Weights ...int64) SZinterstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZinterstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZinterstoreWeights) Build() SCompleted {
	return SCompleted(c)
}

type SZlexcount SCompleted

func (c SZlexcount) Key(Key string) SZlexcountKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZlexcountKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zlexcount() (c SZlexcount) {
	c.cs = append(b.get(), "ZLEXCOUNT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZlexcountKey SCompleted

func (c SZlexcountKey) Min(Min string) SZlexcountMin {
	return SZlexcountMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type SZlexcountMax SCompleted

func (c SZlexcountMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZlexcountMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZlexcountMin SCompleted

func (c SZlexcountMin) Max(Max string) SZlexcountMax {
	return SZlexcountMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type SZmscore SCompleted

func (c SZmscore) Key(Key string) SZmscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZmscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zmscore() (c SZmscore) {
	c.cs = append(b.get(), "ZMSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZmscoreKey SCompleted

func (c SZmscoreKey) Member(Member ...string) SZmscoreMember {
	return SZmscoreMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SZmscoreMember SCompleted

func (c SZmscoreMember) Member(Member ...string) SZmscoreMember {
	return SZmscoreMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SZmscoreMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZmscoreMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZpopmax SCompleted

func (c SZpopmax) Key(Key string) SZpopmaxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZpopmaxKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zpopmax() (c SZpopmax) {
	c.cs = append(b.get(), "ZPOPMAX")
	c.ks = InitSlot
	return
}

type SZpopmaxCount SCompleted

func (c SZpopmaxCount) Build() SCompleted {
	return SCompleted(c)
}

type SZpopmaxKey SCompleted

func (c SZpopmaxKey) Count(Count int64) SZpopmaxCount {
	return SZpopmaxCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZpopmaxKey) Build() SCompleted {
	return SCompleted(c)
}

type SZpopmin SCompleted

func (c SZpopmin) Key(Key string) SZpopminKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZpopminKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zpopmin() (c SZpopmin) {
	c.cs = append(b.get(), "ZPOPMIN")
	c.ks = InitSlot
	return
}

type SZpopminCount SCompleted

func (c SZpopminCount) Build() SCompleted {
	return SCompleted(c)
}

type SZpopminKey SCompleted

func (c SZpopminKey) Count(Count int64) SZpopminCount {
	return SZpopminCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZpopminKey) Build() SCompleted {
	return SCompleted(c)
}

type SZrandmember SCompleted

func (c SZrandmember) Key(Key string) SZrandmemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrandmemberKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrandmember() (c SZrandmember) {
	c.cs = append(b.get(), "ZRANDMEMBER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrandmemberKey SCompleted

func (c SZrandmemberKey) Count(Count int64) SZrandmemberOptionsCount {
	return SZrandmemberOptionsCount{cs: append(c.cs, strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrandmemberKey) Build() SCompleted {
	return SCompleted(c)
}

type SZrandmemberOptionsCount SCompleted

func (c SZrandmemberOptionsCount) Withscores() SZrandmemberOptionsWithscoresWithscores {
	return SZrandmemberOptionsWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrandmemberOptionsCount) Build() SCompleted {
	return SCompleted(c)
}

type SZrandmemberOptionsWithscoresWithscores SCompleted

func (c SZrandmemberOptionsWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZrange SCompleted

func (c SZrange) Key(Key string) SZrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrange() (c SZrange) {
	c.cs = append(b.get(), "ZRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrangeKey SCompleted

func (c SZrangeKey) Min(Min string) SZrangeMin {
	return SZrangeMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type SZrangeLimit SCompleted

func (c SZrangeLimit) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrangeLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeMax SCompleted

func (c SZrangeMax) Byscore() SZrangeSortbyByscore {
	return SZrangeSortbyByscore{cs: append(c.cs, "BYSCORE"), cf: c.cf, ks: c.ks}
}

func (c SZrangeMax) Bylex() SZrangeSortbyBylex {
	return SZrangeSortbyBylex{cs: append(c.cs, "BYLEX"), cf: c.cf, ks: c.ks}
}

func (c SZrangeMax) Rev() SZrangeRevRev {
	return SZrangeRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c SZrangeMax) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangeMax) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrangeMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeMin SCompleted

func (c SZrangeMin) Max(Max string) SZrangeMax {
	return SZrangeMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type SZrangeRevRev SCompleted

func (c SZrangeRevRev) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangeRevRev) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrangeRevRev) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeRevRev) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeSortbyBylex SCompleted

func (c SZrangeSortbyBylex) Rev() SZrangeRevRev {
	return SZrangeRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c SZrangeSortbyBylex) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangeSortbyBylex) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrangeSortbyBylex) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeSortbyBylex) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeSortbyByscore SCompleted

func (c SZrangeSortbyByscore) Rev() SZrangeRevRev {
	return SZrangeRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c SZrangeSortbyByscore) Limit(Offset int64, Count int64) SZrangeLimit {
	return SZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangeSortbyByscore) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrangeSortbyByscore) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeSortbyByscore) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeWithscoresWithscores SCompleted

func (c SZrangeWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylex SCompleted

func (c SZrangebylex) Key(Key string) SZrangebylexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrangebylexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrangebylex() (c SZrangebylex) {
	c.cs = append(b.get(), "ZRANGEBYLEX")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrangebylexKey SCompleted

func (c SZrangebylexKey) Min(Min string) SZrangebylexMin {
	return SZrangebylexMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type SZrangebylexLimit SCompleted

func (c SZrangebylexLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebylexLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylexMax SCompleted

func (c SZrangebylexMax) Limit(Offset int64, Count int64) SZrangebylexLimit {
	return SZrangebylexLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangebylexMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebylexMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylexMin SCompleted

func (c SZrangebylexMin) Max(Max string) SZrangebylexMax {
	return SZrangebylexMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type SZrangebyscore SCompleted

func (c SZrangebyscore) Key(Key string) SZrangebyscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrangebyscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrangebyscore() (c SZrangebyscore) {
	c.cs = append(b.get(), "ZRANGEBYSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrangebyscoreKey SCompleted

func (c SZrangebyscoreKey) Min(Min float64) SZrangebyscoreMin {
	return SZrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZrangebyscoreLimit SCompleted

func (c SZrangebyscoreLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebyscoreLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebyscoreMax SCompleted

func (c SZrangebyscoreMax) Withscores() SZrangebyscoreWithscoresWithscores {
	return SZrangebyscoreWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrangebyscoreMax) Limit(Offset int64, Count int64) SZrangebyscoreLimit {
	return SZrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangebyscoreMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebyscoreMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebyscoreMin SCompleted

func (c SZrangebyscoreMin) Max(Max float64) SZrangebyscoreMax {
	return SZrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZrangebyscoreWithscoresWithscores SCompleted

func (c SZrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) SZrangebyscoreLimit {
	return SZrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangebyscoreWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebyscoreWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangestore SCompleted

func (c SZrangestore) Dst(Dst string) SZrangestoreDst {
	c.ks = checkSlot(c.ks, slot(Dst))
	return SZrangestoreDst{cs: append(c.cs, Dst), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrangestore() (c SZrangestore) {
	c.cs = append(b.get(), "ZRANGESTORE")
	c.ks = InitSlot
	return
}

type SZrangestoreDst SCompleted

func (c SZrangestoreDst) Src(Src string) SZrangestoreSrc {
	c.ks = checkSlot(c.ks, slot(Src))
	return SZrangestoreSrc{cs: append(c.cs, Src), cf: c.cf, ks: c.ks}
}

type SZrangestoreLimit SCompleted

func (c SZrangestoreLimit) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreMax SCompleted

func (c SZrangestoreMax) Byscore() SZrangestoreSortbyByscore {
	return SZrangestoreSortbyByscore{cs: append(c.cs, "BYSCORE"), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreMax) Bylex() SZrangestoreSortbyBylex {
	return SZrangestoreSortbyBylex{cs: append(c.cs, "BYLEX"), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreMax) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreMax) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreMax) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreMin SCompleted

func (c SZrangestoreMin) Max(Max string) SZrangestoreMax {
	return SZrangestoreMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type SZrangestoreRevRev SCompleted

func (c SZrangestoreRevRev) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreRevRev) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSortbyBylex SCompleted

func (c SZrangestoreSortbyBylex) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreSortbyBylex) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreSortbyBylex) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSortbyByscore SCompleted

func (c SZrangestoreSortbyByscore) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cs: append(c.cs, "REV"), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreSortbyByscore) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrangestoreSortbyByscore) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSrc SCompleted

func (c SZrangestoreSrc) Min(Min string) SZrangestoreMin {
	return SZrangestoreMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type SZrank SCompleted

func (c SZrank) Key(Key string) SZrankKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrankKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrank() (c SZrank) {
	c.cs = append(b.get(), "ZRANK")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrankKey SCompleted

func (c SZrankKey) Member(Member string) SZrankMember {
	return SZrankMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SZrankMember SCompleted

func (c SZrankMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrankMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZrem SCompleted

func (c SZrem) Key(Key string) SZremKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZremKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrem() (c SZrem) {
	c.cs = append(b.get(), "ZREM")
	c.ks = InitSlot
	return
}

type SZremKey SCompleted

func (c SZremKey) Member(Member ...string) SZremMember {
	return SZremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

type SZremMember SCompleted

func (c SZremMember) Member(Member ...string) SZremMember {
	return SZremMember{cs: append(c.cs, Member...), cf: c.cf, ks: c.ks}
}

func (c SZremMember) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebylex SCompleted

func (c SZremrangebylex) Key(Key string) SZremrangebylexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZremrangebylexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zremrangebylex() (c SZremrangebylex) {
	c.cs = append(b.get(), "ZREMRANGEBYLEX")
	c.ks = InitSlot
	return
}

type SZremrangebylexKey SCompleted

func (c SZremrangebylexKey) Min(Min string) SZremrangebylexMin {
	return SZremrangebylexMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type SZremrangebylexMax SCompleted

func (c SZremrangebylexMax) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebylexMin SCompleted

func (c SZremrangebylexMin) Max(Max string) SZremrangebylexMax {
	return SZremrangebylexMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type SZremrangebyrank SCompleted

func (c SZremrangebyrank) Key(Key string) SZremrangebyrankKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZremrangebyrankKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zremrangebyrank() (c SZremrangebyrank) {
	c.cs = append(b.get(), "ZREMRANGEBYRANK")
	c.ks = InitSlot
	return
}

type SZremrangebyrankKey SCompleted

func (c SZremrangebyrankKey) Start(Start int64) SZremrangebyrankStart {
	return SZremrangebyrankStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type SZremrangebyrankStart SCompleted

func (c SZremrangebyrankStart) Stop(Stop int64) SZremrangebyrankStop {
	return SZremrangebyrankStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type SZremrangebyrankStop SCompleted

func (c SZremrangebyrankStop) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebyscore SCompleted

func (c SZremrangebyscore) Key(Key string) SZremrangebyscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZremrangebyscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zremrangebyscore() (c SZremrangebyscore) {
	c.cs = append(b.get(), "ZREMRANGEBYSCORE")
	c.ks = InitSlot
	return
}

type SZremrangebyscoreKey SCompleted

func (c SZremrangebyscoreKey) Min(Min float64) SZremrangebyscoreMin {
	return SZremrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZremrangebyscoreMax SCompleted

func (c SZremrangebyscoreMax) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebyscoreMin SCompleted

func (c SZremrangebyscoreMin) Max(Max float64) SZremrangebyscoreMax {
	return SZremrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZrevrange SCompleted

func (c SZrevrange) Key(Key string) SZrevrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrevrangeKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrevrange() (c SZrevrange) {
	c.cs = append(b.get(), "ZREVRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrangeKey SCompleted

func (c SZrevrangeKey) Start(Start int64) SZrevrangeStart {
	return SZrevrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10)), cf: c.cf, ks: c.ks}
}

type SZrevrangeStart SCompleted

func (c SZrevrangeStart) Stop(Stop int64) SZrevrangeStop {
	return SZrevrangeStop{cs: append(c.cs, strconv.FormatInt(Stop, 10)), cf: c.cf, ks: c.ks}
}

type SZrevrangeStop SCompleted

func (c SZrevrangeStop) Withscores() SZrevrangeWithscoresWithscores {
	return SZrevrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrevrangeStop) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangeStop) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangeWithscoresWithscores SCompleted

func (c SZrevrangeWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangeWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebylex SCompleted

func (c SZrevrangebylex) Key(Key string) SZrevrangebylexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrevrangebylexKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrevrangebylex() (c SZrevrangebylex) {
	c.cs = append(b.get(), "ZREVRANGEBYLEX")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrangebylexKey SCompleted

func (c SZrevrangebylexKey) Max(Max string) SZrevrangebylexMax {
	return SZrevrangebylexMax{cs: append(c.cs, Max), cf: c.cf, ks: c.ks}
}

type SZrevrangebylexLimit SCompleted

func (c SZrevrangebylexLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebylexLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebylexMax SCompleted

func (c SZrevrangebylexMax) Min(Min string) SZrevrangebylexMin {
	return SZrevrangebylexMin{cs: append(c.cs, Min), cf: c.cf, ks: c.ks}
}

type SZrevrangebylexMin SCompleted

func (c SZrevrangebylexMin) Limit(Offset int64, Count int64) SZrevrangebylexLimit {
	return SZrevrangebylexLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrevrangebylexMin) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebylexMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscore SCompleted

func (c SZrevrangebyscore) Key(Key string) SZrevrangebyscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrevrangebyscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrevrangebyscore() (c SZrevrangebyscore) {
	c.cs = append(b.get(), "ZREVRANGEBYSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrangebyscoreKey SCompleted

func (c SZrevrangebyscoreKey) Max(Max float64) SZrevrangebyscoreMax {
	return SZrevrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZrevrangebyscoreLimit SCompleted

func (c SZrevrangebyscoreLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebyscoreLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscoreMax SCompleted

func (c SZrevrangebyscoreMax) Min(Min float64) SZrevrangebyscoreMin {
	return SZrevrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64)), cf: c.cf, ks: c.ks}
}

type SZrevrangebyscoreMin SCompleted

func (c SZrevrangebyscoreMin) Withscores() SZrevrangebyscoreWithscoresWithscores {
	return SZrevrangebyscoreWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZrevrangebyscoreMin) Limit(Offset int64, Count int64) SZrevrangebyscoreLimit {
	return SZrevrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrevrangebyscoreMin) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebyscoreMin) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrangebyscoreWithscoresWithscores SCompleted

func (c SZrevrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) SZrevrangebyscoreLimit {
	return SZrevrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZrevrangebyscoreWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrangebyscoreWithscoresWithscores) Cache() SCacheable {
	return SCacheable(c)
}

type SZrevrank SCompleted

func (c SZrevrank) Key(Key string) SZrevrankKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrevrankKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zrevrank() (c SZrevrank) {
	c.cs = append(b.get(), "ZREVRANK")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrankKey SCompleted

func (c SZrevrankKey) Member(Member string) SZrevrankMember {
	return SZrevrankMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SZrevrankMember SCompleted

func (c SZrevrankMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrevrankMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZscan SCompleted

func (c SZscan) Key(Key string) SZscanKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZscanKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zscan() (c SZscan) {
	c.cs = append(b.get(), "ZSCAN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZscanCount SCompleted

func (c SZscanCount) Build() SCompleted {
	return SCompleted(c)
}

type SZscanCursor SCompleted

func (c SZscanCursor) Match(Pattern string) SZscanMatch {
	return SZscanMatch{cs: append(c.cs, "MATCH", Pattern), cf: c.cf, ks: c.ks}
}

func (c SZscanCursor) Count(Count int64) SZscanCount {
	return SZscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SZscanKey SCompleted

func (c SZscanKey) Cursor(Cursor int64) SZscanCursor {
	return SZscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10)), cf: c.cf, ks: c.ks}
}

type SZscanMatch SCompleted

func (c SZscanMatch) Count(Count int64) SZscanCount {
	return SZscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10)), cf: c.cf, ks: c.ks}
}

func (c SZscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SZscore SCompleted

func (c SZscore) Key(Key string) SZscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZscoreKey{cs: append(c.cs, Key), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zscore() (c SZscore) {
	c.cs = append(b.get(), "ZSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZscoreKey SCompleted

func (c SZscoreKey) Member(Member string) SZscoreMember {
	return SZscoreMember{cs: append(c.cs, Member), cf: c.cf, ks: c.ks}
}

type SZscoreMember SCompleted

func (c SZscoreMember) Build() SCompleted {
	return SCompleted(c)
}

func (c SZscoreMember) Cache() SCacheable {
	return SCacheable(c)
}

type SZunion SCompleted

func (c SZunion) Numkeys(Numkeys int64) SZunionNumkeys {
	return SZunionNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zunion() (c SZunion) {
	c.cs = append(b.get(), "ZUNION")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZunionAggregateMax SCompleted

func (c SZunionAggregateMax) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZunionAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZunionAggregateMin SCompleted

func (c SZunionAggregateMin) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZunionAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZunionAggregateSum SCompleted

func (c SZunionAggregateSum) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZunionAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZunionKey SCompleted

func (c SZunionKey) Weights(Weight ...int64) SZunionWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZunionKey) Sum() SZunionAggregateSum {
	return SZunionAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZunionKey) Min() SZunionAggregateMin {
	return SZunionAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZunionKey) Max() SZunionAggregateMax {
	return SZunionAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZunionKey) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZunionKey) Key(Key ...string) SZunionKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SZunionKey) Build() SCompleted {
	return SCompleted(c)
}

type SZunionNumkeys SCompleted

func (c SZunionNumkeys) Key(Key ...string) SZunionKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SZunionWeights SCompleted

func (c SZunionWeights) Sum() SZunionAggregateSum {
	return SZunionAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZunionWeights) Min() SZunionAggregateMin {
	return SZunionAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZunionWeights) Max() SZunionAggregateMax {
	return SZunionAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZunionWeights) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES"), cf: c.cf, ks: c.ks}
}

func (c SZunionWeights) Weights(Weights ...int64) SZunionWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZunionWeights) Build() SCompleted {
	return SCompleted(c)
}

type SZunionWithscoresWithscores SCompleted

func (c SZunionWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstore SCompleted

func (c SZunionstore) Destination(Destination string) SZunionstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SZunionstoreDestination{cs: append(c.cs, Destination), cf: c.cf, ks: c.ks}
}

func (b *SBuilder) Zunionstore() (c SZunionstore) {
	c.cs = append(b.get(), "ZUNIONSTORE")
	c.ks = InitSlot
	return
}

type SZunionstoreAggregateMax SCompleted

func (c SZunionstoreAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreAggregateMin SCompleted

func (c SZunionstoreAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreAggregateSum SCompleted

func (c SZunionstoreAggregateSum) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreDestination SCompleted

func (c SZunionstoreDestination) Numkeys(Numkeys int64) SZunionstoreNumkeys {
	return SZunionstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10)), cf: c.cf, ks: c.ks}
}

type SZunionstoreKey SCompleted

func (c SZunionstoreKey) Weights(Weight ...int64) SZunionstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZunionstoreKey) Sum() SZunionstoreAggregateSum {
	return SZunionstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZunionstoreKey) Min() SZunionstoreAggregateMin {
	return SZunionstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZunionstoreKey) Max() SZunionstoreAggregateMax {
	return SZunionstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZunionstoreKey) Key(Key ...string) SZunionstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

func (c SZunionstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreNumkeys SCompleted

func (c SZunionstoreNumkeys) Key(Key ...string) SZunionstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionstoreKey{cs: append(c.cs, Key...), cf: c.cf, ks: c.ks}
}

type SZunionstoreWeights SCompleted

func (c SZunionstoreWeights) Sum() SZunionstoreAggregateSum {
	return SZunionstoreAggregateSum{cs: append(c.cs, "SUM"), cf: c.cf, ks: c.ks}
}

func (c SZunionstoreWeights) Min() SZunionstoreAggregateMin {
	return SZunionstoreAggregateMin{cs: append(c.cs, "MIN"), cf: c.cf, ks: c.ks}
}

func (c SZunionstoreWeights) Max() SZunionstoreAggregateMax {
	return SZunionstoreAggregateMax{cs: append(c.cs, "MAX"), cf: c.cf, ks: c.ks}
}

func (c SZunionstoreWeights) Weights(Weights ...int64) SZunionstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SZunionstoreWeights{cs: c.cs, cf: c.cf, ks: c.ks}
}

func (c SZunionstoreWeights) Build() SCompleted {
	return SCompleted(c)
}

