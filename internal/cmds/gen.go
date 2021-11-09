// Code generated DO NOT EDIT

package cmds

import "strconv"

type AclCat Completed

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

type AclCatCategoryname Completed

func (c AclCatCategoryname) Build() Completed {
	return Completed(c)
}

type AclDeluser Completed

func (c AclDeluser) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (b *Builder) AclDeluser() (c AclDeluser) {
	c.cs = append(b.get(), "ACL", "DELUSER")
	return
}

type AclDeluserUsername Completed

func (c AclDeluserUsername) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (c AclDeluserUsername) Build() Completed {
	return Completed(c)
}

type AclGenpass Completed

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

type AclGenpassBits Completed

func (c AclGenpassBits) Build() Completed {
	return Completed(c)
}

type AclGetuser Completed

func (c AclGetuser) Username(Username string) AclGetuserUsername {
	return AclGetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
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
	return AclLogCountOrReset{cf: c.cf, cs: append(c.cs, CountOrReset)}
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
	return AclSetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (b *Builder) AclSetuser() (c AclSetuser) {
	c.cs = append(b.get(), "ACL", "SETUSER")
	return
}

type AclSetuserRule Completed

func (c AclSetuserRule) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c AclSetuserRule) Build() Completed {
	return Completed(c)
}

type AclSetuserUsername Completed

func (c AclSetuserUsername) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
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
	return AppendKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Append() (c Append) {
	c.cs = append(b.get(), "APPEND")
	return
}

type AppendKey Completed

func (c AppendKey) Value(Value string) AppendValue {
	return AppendValue{cf: c.cf, cs: append(c.cs, Value)}
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
	return AuthUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (c Auth) Password(Password string) AuthPassword {
	return AuthPassword{cf: c.cf, cs: append(c.cs, Password)}
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
	return AuthPassword{cf: c.cf, cs: append(c.cs, Password)}
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
	return BgsaveScheduleSchedule{cf: c.cf, cs: append(c.cs, "SCHEDULE")}
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
	return BitcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Bitcount() (c Bitcount) {
	c.cs = append(b.get(), "BITCOUNT")
	return
}

type BitcountKey Completed

func (c BitcountKey) StartEnd(Start int64, End int64) BitcountStartEnd {
	return BitcountStartEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10), strconv.FormatInt(End, 10))}
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
	return BitfieldKey{cf: c.cf, cs: append(c.cs, Key)}
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

type BitfieldIncrby Completed

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

type BitfieldKey Completed

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

type BitfieldRo Completed

func (c BitfieldRo) Key(Key string) BitfieldRoKey {
	return BitfieldRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) BitfieldRo() (c BitfieldRo) {
	c.cs = append(b.get(), "BITFIELD_RO")
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
	return BitfieldRoGet{cf: c.cf, cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

type BitfieldSat Completed

func (c BitfieldSat) Build() Completed {
	return Completed(c)
}

type BitfieldSet Completed

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

type BitfieldWrap Completed

func (c BitfieldWrap) Build() Completed {
	return Completed(c)
}

type Bitop Completed

func (c Bitop) Operation(Operation string) BitopOperation {
	return BitopOperation{cf: c.cf, cs: append(c.cs, Operation)}
}

func (b *Builder) Bitop() (c Bitop) {
	c.cs = append(b.get(), "BITOP")
	return
}

type BitopDestkey Completed

func (c BitopDestkey) Key(Key ...string) BitopKey {
	return BitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BitopKey Completed

func (c BitopKey) Key(Key ...string) BitopKey {
	return BitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c BitopKey) Build() Completed {
	return Completed(c)
}

type BitopOperation Completed

func (c BitopOperation) Destkey(Destkey string) BitopDestkey {
	return BitopDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

type Bitpos Completed

func (c Bitpos) Key(Key string) BitposKey {
	return BitposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Bitpos() (c Bitpos) {
	c.cs = append(b.get(), "BITPOS")
	return
}

type BitposBit Completed

func (c BitposBit) Start(Start int64) BitposIndexStart {
	return BitposIndexStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
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
	return BitposIndexEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c BitposIndexStart) Build() Completed {
	return Completed(c)
}

func (c BitposIndexStart) Cache() Cacheable {
	return Cacheable(c)
}

type BitposKey Completed

func (c BitposKey) Bit(Bit int64) BitposBit {
	return BitposBit{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bit, 10))}
}

type Blmove Completed

func (c Blmove) Source(Source string) BlmoveSource {
	return BlmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Blmove() (c Blmove) {
	c.cs = append(b.get(), "BLMOVE")
	c.cf = blockTag
	return
}

type BlmoveDestination Completed

func (c BlmoveDestination) Left() BlmoveWherefromLeft {
	return BlmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveDestination) Right() BlmoveWherefromRight {
	return BlmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveSource Completed

func (c BlmoveSource) Destination(Destination string) BlmoveDestination {
	return BlmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type BlmoveTimeout Completed

func (c BlmoveTimeout) Build() Completed {
	return Completed(c)
}

type BlmoveWherefromLeft Completed

func (c BlmoveWherefromLeft) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromLeft) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveWherefromRight Completed

func (c BlmoveWherefromRight) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromRight) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveWheretoLeft Completed

func (c BlmoveWheretoLeft) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BlmoveWheretoRight Completed

func (c BlmoveWheretoRight) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type Blmpop Completed

func (c Blmpop) Timeout(Timeout float64) BlmpopTimeout {
	return BlmpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
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
	return BlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmpopKey) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c BlmpopKey) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BlmpopNumkeys Completed

func (c BlmpopNumkeys) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c BlmpopNumkeys) Left() BlmpopWhereLeft {
	return BlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmpopNumkeys) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmpopTimeout Completed

func (c BlmpopTimeout) Numkeys(Numkeys int64) BlmpopNumkeys {
	return BlmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type BlmpopWhereLeft Completed

func (c BlmpopWhereLeft) Count(Count int64) BlmpopCount {
	return BlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type BlmpopWhereRight Completed

func (c BlmpopWhereRight) Count(Count int64) BlmpopCount {
	return BlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Blpop Completed

func (c Blpop) Key(Key ...string) BlpopKey {
	return BlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Blpop() (c Blpop) {
	c.cs = append(b.get(), "BLPOP")
	c.cf = blockTag
	return
}

type BlpopKey Completed

func (c BlpopKey) Timeout(Timeout float64) BlpopTimeout {
	return BlpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BlpopKey) Key(Key ...string) BlpopKey {
	return BlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BlpopTimeout Completed

func (c BlpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpop Completed

func (c Brpop) Key(Key ...string) BrpopKey {
	return BrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Brpop() (c Brpop) {
	c.cs = append(b.get(), "BRPOP")
	c.cf = blockTag
	return
}

type BrpopKey Completed

func (c BrpopKey) Timeout(Timeout float64) BrpopTimeout {
	return BrpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BrpopKey) Key(Key ...string) BrpopKey {
	return BrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BrpopTimeout Completed

func (c BrpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpoplpush Completed

func (c Brpoplpush) Source(Source string) BrpoplpushSource {
	return BrpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Brpoplpush() (c Brpoplpush) {
	c.cs = append(b.get(), "BRPOPLPUSH")
	c.cf = blockTag
	return
}

type BrpoplpushDestination Completed

func (c BrpoplpushDestination) Timeout(Timeout float64) BrpoplpushTimeout {
	return BrpoplpushTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BrpoplpushSource Completed

func (c BrpoplpushSource) Destination(Destination string) BrpoplpushDestination {
	return BrpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type BrpoplpushTimeout Completed

func (c BrpoplpushTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmax Completed

func (c Bzpopmax) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmax() (c Bzpopmax) {
	c.cs = append(b.get(), "BZPOPMAX")
	c.cf = blockTag
	return
}

type BzpopmaxKey Completed

func (c BzpopmaxKey) Timeout(Timeout float64) BzpopmaxTimeout {
	return BzpopmaxTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopmaxKey) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BzpopmaxTimeout Completed

func (c BzpopmaxTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmin Completed

func (c Bzpopmin) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmin() (c Bzpopmin) {
	c.cs = append(b.get(), "BZPOPMIN")
	c.cf = blockTag
	return
}

type BzpopminKey Completed

func (c BzpopminKey) Timeout(Timeout float64) BzpopminTimeout {
	return BzpopminTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopminKey) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BzpopminTimeout Completed

func (c BzpopminTimeout) Build() Completed {
	return Completed(c)
}

type ClientCaching Completed

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

type ClientKillAddr Completed

func (c ClientKillAddr) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillAddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillAddr) Build() Completed {
	return Completed(c)
}

type ClientKillId Completed

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

type ClientKillIpPort Completed

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

type ClientKillLaddr Completed

func (c ClientKillLaddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillLaddr) Build() Completed {
	return Completed(c)
}

type ClientKillMaster Completed

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

type ClientKillNormal Completed

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

type ClientKillPubsub Completed

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

type ClientKillSkipme Completed

func (c ClientKillSkipme) Build() Completed {
	return Completed(c)
}

type ClientKillSlave Completed

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

type ClientKillUser Completed

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

type ClientList Completed

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
	return ClientListIdClientId{cf: c.cf, cs: c.cs}
}

func (c ClientListIdClientId) Build() Completed {
	return Completed(c)
}

type ClientListIdId Completed

func (c ClientListIdId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cf: c.cf, cs: c.cs}
}

type ClientListMaster Completed

func (c ClientListMaster) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c ClientListMaster) Build() Completed {
	return Completed(c)
}

type ClientListNormal Completed

func (c ClientListNormal) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c ClientListNormal) Build() Completed {
	return Completed(c)
}

type ClientListPubsub Completed

func (c ClientListPubsub) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c ClientListPubsub) Build() Completed {
	return Completed(c)
}

type ClientListReplica Completed

func (c ClientListReplica) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c ClientListReplica) Build() Completed {
	return Completed(c)
}

type ClientNoEvict Completed

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
	return ClientPauseTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
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
	return ClientPauseModeWrite{cf: c.cf, cs: append(c.cs, "WRITE")}
}

func (c ClientPauseTimeout) All() ClientPauseModeAll {
	return ClientPauseModeAll{cf: c.cf, cs: append(c.cs, "ALL")}
}

func (c ClientPauseTimeout) Build() Completed {
	return Completed(c)
}

type ClientReply Completed

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
	return ClientSetnameConnectionName{cf: c.cf, cs: append(c.cs, ConnectionName)}
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
	return ClientTrackingStatusOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c ClientTracking) Off() ClientTrackingStatusOff {
	return ClientTrackingStatusOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (b *Builder) ClientTracking() (c ClientTracking) {
	c.cs = append(b.get(), "CLIENT", "TRACKING")
	return
}

type ClientTrackingBcastBcast Completed

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

type ClientTrackingNoloopNoloop Completed

func (c ClientTrackingNoloopNoloop) Build() Completed {
	return Completed(c)
}

type ClientTrackingOptinOptin Completed

func (c ClientTrackingOptinOptin) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingOptinOptin) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingOptinOptin) Build() Completed {
	return Completed(c)
}

type ClientTrackingOptoutOptout Completed

func (c ClientTrackingOptoutOptout) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingOptoutOptout) Build() Completed {
	return Completed(c)
}

type ClientTrackingPrefix Completed

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

type ClientTrackingRedirect Completed

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

type ClientTrackingStatusOff Completed

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

type ClientTrackingStatusOn Completed

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
	return ClientUnblockClientId{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ClientId, 10))}
}

func (b *Builder) ClientUnblock() (c ClientUnblock) {
	c.cs = append(b.get(), "CLIENT", "UNBLOCK")
	return
}

type ClientUnblockClientId Completed

func (c ClientUnblockClientId) Timeout() ClientUnblockUnblockTypeTimeout {
	return ClientUnblockUnblockTypeTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT")}
}

func (c ClientUnblockClientId) Error() ClientUnblockUnblockTypeError {
	return ClientUnblockUnblockTypeError{cf: c.cf, cs: append(c.cs, "ERROR")}
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
	return ClusterAddslotsSlot{cf: c.cf, cs: c.cs}
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
	return ClusterAddslotsSlot{cf: c.cf, cs: c.cs}
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
	return ClusterCountFailureReportsNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return ClusterCountkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
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
	return ClusterDelslotsSlot{cf: c.cf, cs: c.cs}
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
	return ClusterDelslotsSlot{cf: c.cf, cs: c.cs}
}

func (c ClusterDelslotsSlot) Build() Completed {
	return Completed(c)
}

type ClusterFailover Completed

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
	return ClusterForgetNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return ClusterGetkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
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
	return ClusterGetkeysinslotCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
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
	return ClusterKeyslotKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return ClusterMeetIp{cf: c.cf, cs: append(c.cs, Ip)}
}

func (b *Builder) ClusterMeet() (c ClusterMeet) {
	c.cs = append(b.get(), "CLUSTER", "MEET")
	return
}

type ClusterMeetIp Completed

func (c ClusterMeetIp) Port(Port int64) ClusterMeetPort {
	return ClusterMeetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
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
	return ClusterReplicasNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return ClusterReplicateNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return ClusterSetConfigEpochConfigEpoch{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ConfigEpoch, 10))}
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
	return ClusterSetslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
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

type ClusterSetslotSubcommandImporting Completed

func (c ClusterSetslotSubcommandImporting) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandImporting) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandMigrating Completed

func (c ClusterSetslotSubcommandMigrating) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandMigrating) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandNode Completed

func (c ClusterSetslotSubcommandNode) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandNode) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandStable Completed

func (c ClusterSetslotSubcommandStable) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandStable) Build() Completed {
	return Completed(c)
}

type ClusterSlaves Completed

func (c ClusterSlaves) NodeId(NodeId string) ClusterSlavesNodeId {
	return ClusterSlavesNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return CommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (b *Builder) CommandInfo() (c CommandInfo) {
	c.cs = append(b.get(), "COMMAND", "INFO")
	return
}

type CommandInfoCommandName Completed

func (c CommandInfoCommandName) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (c CommandInfoCommandName) Build() Completed {
	return Completed(c)
}

type ConfigGet Completed

func (c ConfigGet) Parameter(Parameter string) ConfigGetParameter {
	return ConfigGetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
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
	return ConfigSetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
}

func (b *Builder) ConfigSet() (c ConfigSet) {
	c.cs = append(b.get(), "CONFIG", "SET")
	return
}

type ConfigSetParameter Completed

func (c ConfigSetParameter) Value(Value string) ConfigSetValue {
	return ConfigSetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type ConfigSetValue Completed

func (c ConfigSetValue) Build() Completed {
	return Completed(c)
}

type Copy Completed

func (c Copy) Source(Source string) CopySource {
	return CopySource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Copy() (c Copy) {
	c.cs = append(b.get(), "COPY")
	return
}

type CopyDb Completed

func (c CopyDb) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c CopyDb) Build() Completed {
	return Completed(c)
}

type CopyDestination Completed

func (c CopyDestination) Db(DestinationDb int64) CopyDb {
	return CopyDb{cf: c.cf, cs: append(c.cs, "DB", strconv.FormatInt(DestinationDb, 10))}
}

func (c CopyDestination) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
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
	return CopyDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Dbsize Completed

func (c Dbsize) Build() Completed {
	return Completed(c)
}

func (b *Builder) Dbsize() (c Dbsize) {
	c.cs = append(b.get(), "DBSIZE")
	return
}

type DebugObject Completed

func (c DebugObject) Key(Key string) DebugObjectKey {
	return DebugObjectKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return DecrKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return DecrbyKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return DecrbyDecrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Decrement, 10))}
}

type Del Completed

func (c Del) Key(Key ...string) DelKey {
	return DelKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Del() (c Del) {
	c.cs = append(b.get(), "DEL")
	return
}

type DelKey Completed

func (c DelKey) Key(Key ...string) DelKey {
	return DelKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return DumpKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Dump() (c Dump) {
	c.cs = append(b.get(), "DUMP")
	return
}

type DumpKey Completed

func (c DumpKey) Build() Completed {
	return Completed(c)
}

type Echo Completed

func (c Echo) Message(Message string) EchoMessage {
	return EchoMessage{cf: c.cf, cs: append(c.cs, Message)}
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
	return EvalScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *Builder) Eval() (c Eval) {
	c.cs = append(b.get(), "EVAL")
	return
}

type EvalArg Completed

func (c EvalArg) Arg(Arg ...string) EvalArg {
	return EvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalArg) Build() Completed {
	return Completed(c)
}

type EvalKey Completed

func (c EvalKey) Arg(Arg ...string) EvalArg {
	return EvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalKey) Key(Key ...string) EvalKey {
	return EvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalKey) Build() Completed {
	return Completed(c)
}

type EvalNumkeys Completed

func (c EvalNumkeys) Key(Key ...string) EvalKey {
	return EvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalNumkeys) Arg(Arg ...string) EvalArg {
	return EvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalNumkeys) Build() Completed {
	return Completed(c)
}

type EvalRo Completed

func (c EvalRo) Script(Script string) EvalRoScript {
	return EvalRoScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *Builder) EvalRo() (c EvalRo) {
	c.cs = append(b.get(), "EVAL_RO")
	return
}

type EvalRoArg Completed

func (c EvalRoArg) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalRoArg) Build() Completed {
	return Completed(c)
}

type EvalRoKey Completed

func (c EvalRoKey) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalRoKey) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalRoNumkeys Completed

func (c EvalRoNumkeys) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalRoScript Completed

func (c EvalRoScript) Numkeys(Numkeys int64) EvalRoNumkeys {
	return EvalRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalScript Completed

func (c EvalScript) Numkeys(Numkeys int64) EvalNumkeys {
	return EvalNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type Evalsha Completed

func (c Evalsha) Sha1(Sha1 string) EvalshaSha1 {
	return EvalshaSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *Builder) Evalsha() (c Evalsha) {
	c.cs = append(b.get(), "EVALSHA")
	return
}

type EvalshaArg Completed

func (c EvalshaArg) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaArg) Build() Completed {
	return Completed(c)
}

type EvalshaKey Completed

func (c EvalshaKey) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaKey) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalshaKey) Build() Completed {
	return Completed(c)
}

type EvalshaNumkeys Completed

func (c EvalshaNumkeys) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c EvalshaNumkeys) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaNumkeys) Build() Completed {
	return Completed(c)
}

type EvalshaRo Completed

func (c EvalshaRo) Sha1(Sha1 string) EvalshaRoSha1 {
	return EvalshaRoSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *Builder) EvalshaRo() (c EvalshaRo) {
	c.cs = append(b.get(), "EVALSHA_RO")
	return
}

type EvalshaRoArg Completed

func (c EvalshaRoArg) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaRoArg) Build() Completed {
	return Completed(c)
}

type EvalshaRoKey Completed

func (c EvalshaRoKey) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaRoKey) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalshaRoNumkeys Completed

func (c EvalshaRoNumkeys) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalshaRoSha1 Completed

func (c EvalshaRoSha1) Numkeys(Numkeys int64) EvalshaRoNumkeys {
	return EvalshaRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalshaSha1 Completed

func (c EvalshaSha1) Numkeys(Numkeys int64) EvalshaNumkeys {
	return EvalshaNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
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
	return ExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Exists() (c Exists) {
	c.cs = append(b.get(), "EXISTS")
	return
}

type ExistsKey Completed

func (c ExistsKey) Key(Key ...string) ExistsKey {
	return ExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ExistsKey) Build() Completed {
	return Completed(c)
}

type Expire Completed

func (c Expire) Key(Key string) ExpireKey {
	return ExpireKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return ExpireSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type ExpireSeconds Completed

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

type Expireat Completed

func (c Expireat) Key(Key string) ExpireatKey {
	return ExpireatKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return ExpireatTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timestamp, 10))}
}

type ExpireatTimestamp Completed

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

type Expiretime Completed

func (c Expiretime) Key(Key string) ExpiretimeKey {
	return ExpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Expiretime() (c Expiretime) {
	c.cs = append(b.get(), "EXPIRETIME")
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
	return FailoverTargetTo{cf: c.cf, cs: append(c.cs, "TO")}
}

func (c Failover) Abort() FailoverAbort {
	return FailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c Failover) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
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
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverAbort) Build() Completed {
	return Completed(c)
}

type FailoverTargetForce Completed

func (c FailoverTargetForce) Abort() FailoverAbort {
	return FailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c FailoverTargetForce) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverTargetForce) Build() Completed {
	return Completed(c)
}

type FailoverTargetHost Completed

func (c FailoverTargetHost) Port(Port int64) FailoverTargetPort {
	return FailoverTargetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type FailoverTargetPort Completed

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

type FailoverTargetTo Completed

func (c FailoverTargetTo) Host(Host string) FailoverTargetHost {
	return FailoverTargetHost{cf: c.cf, cs: append(c.cs, Host)}
}

type FailoverTimeout Completed

func (c FailoverTimeout) Build() Completed {
	return Completed(c)
}

type Flushall Completed

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
	return GeoaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geoadd() (c Geoadd) {
	c.cs = append(b.get(), "GEOADD")
	return
}

type GeoaddChangeCh Completed

func (c GeoaddChangeCh) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type GeoaddConditionNx Completed

func (c GeoaddConditionNx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c GeoaddConditionNx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type GeoaddConditionXx Completed

func (c GeoaddConditionXx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c GeoaddConditionXx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type GeoaddKey Completed

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
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type GeoaddLongitudeLatitudeMember Completed

func (c GeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member)}
}

func (c GeoaddLongitudeLatitudeMember) Build() Completed {
	return Completed(c)
}

type Geodist Completed

func (c Geodist) Key(Key string) GeodistKey {
	return GeodistKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geodist() (c Geodist) {
	c.cs = append(b.get(), "GEODIST")
	return
}

type GeodistKey Completed

func (c GeodistKey) Member1(Member1 string) GeodistMember1 {
	return GeodistMember1{cf: c.cf, cs: append(c.cs, Member1)}
}

type GeodistMember1 Completed

func (c GeodistMember1) Member2(Member2 string) GeodistMember2 {
	return GeodistMember2{cf: c.cf, cs: append(c.cs, Member2)}
}

type GeodistMember2 Completed

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
	return GeohashKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geohash() (c Geohash) {
	c.cs = append(b.get(), "GEOHASH")
	return
}

type GeohashKey Completed

func (c GeohashKey) Member(Member ...string) GeohashMember {
	return GeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type GeohashMember Completed

func (c GeohashMember) Member(Member ...string) GeohashMember {
	return GeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeohashMember) Build() Completed {
	return Completed(c)
}

func (c GeohashMember) Cache() Cacheable {
	return Cacheable(c)
}

type Geopos Completed

func (c Geopos) Key(Key string) GeoposKey {
	return GeoposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geopos() (c Geopos) {
	c.cs = append(b.get(), "GEOPOS")
	return
}

type GeoposKey Completed

func (c GeoposKey) Member(Member ...string) GeoposMember {
	return GeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type GeoposMember Completed

func (c GeoposMember) Member(Member ...string) GeoposMember {
	return GeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeoposMember) Build() Completed {
	return Completed(c)
}

func (c GeoposMember) Cache() Cacheable {
	return Cacheable(c)
}

type Georadius Completed

func (c Georadius) Key(Key string) GeoradiusKey {
	return GeoradiusKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Georadius() (c Georadius) {
	c.cs = append(b.get(), "GEORADIUS")
	return
}

type GeoradiusCountAnyAny Completed

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

type GeoradiusCountCount Completed

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

type GeoradiusKey Completed

func (c GeoradiusKey) Longitude(Longitude float64) GeoradiusLongitude {
	return GeoradiusLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type GeoradiusLatitude Completed

func (c GeoradiusLatitude) Radius(Radius float64) GeoradiusRadius {
	return GeoradiusRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusLongitude Completed

func (c GeoradiusLongitude) Latitude(Latitude float64) GeoradiusLatitude {
	return GeoradiusLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type GeoradiusOrderAsc Completed

func (c GeoradiusOrderAsc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusOrderAsc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusOrderAsc) Build() Completed {
	return Completed(c)
}

type GeoradiusOrderDesc Completed

func (c GeoradiusOrderDesc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusOrderDesc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusOrderDesc) Build() Completed {
	return Completed(c)
}

type GeoradiusRadius Completed

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

type GeoradiusRo Completed

func (c GeoradiusRo) Key(Key string) GeoradiusRoKey {
	return GeoradiusRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) GeoradiusRo() (c GeoradiusRo) {
	c.cs = append(b.get(), "GEORADIUS_RO")
	return
}

type GeoradiusRoCountAnyAny Completed

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

type GeoradiusRoCountCount Completed

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

type GeoradiusRoKey Completed

func (c GeoradiusRoKey) Longitude(Longitude float64) GeoradiusRoLongitude {
	return GeoradiusRoLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type GeoradiusRoLatitude Completed

func (c GeoradiusRoLatitude) Radius(Radius float64) GeoradiusRoRadius {
	return GeoradiusRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusRoLongitude Completed

func (c GeoradiusRoLongitude) Latitude(Latitude float64) GeoradiusRoLatitude {
	return GeoradiusRoLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type GeoradiusRoOrderAsc Completed

func (c GeoradiusRoOrderAsc) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoOrderDesc Completed

func (c GeoradiusRoOrderDesc) Storedist(Key string) GeoradiusRoStoredist {
	return GeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoRadius Completed

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

type GeoradiusRoStoredist Completed

func (c GeoradiusRoStoredist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoStoredist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitFt Completed

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

func (c GeoradiusRoUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitKm Completed

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

func (c GeoradiusRoUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitM Completed

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

func (c GeoradiusRoUnitM) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoUnitMi Completed

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

func (c GeoradiusRoUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithcoordWithcoord Completed

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

func (c GeoradiusRoWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithdistWithdist Completed

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

func (c GeoradiusRoWithdistWithdist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusRoWithhashWithhash Completed

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

func (c GeoradiusRoWithhashWithhash) Build() Completed {
	return Completed(c)
}

func (c GeoradiusRoWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusStore Completed

func (c GeoradiusStore) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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

func (c GeoradiusUnitFt) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitKm Completed

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

func (c GeoradiusUnitKm) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitM Completed

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

func (c GeoradiusUnitM) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitMi Completed

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

func (c GeoradiusUnitMi) Build() Completed {
	return Completed(c)
}

type GeoradiusWithcoordWithcoord Completed

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

func (c GeoradiusWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

type GeoradiusWithdistWithdist Completed

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

func (c GeoradiusWithdistWithdist) Build() Completed {
	return Completed(c)
}

type GeoradiusWithhashWithhash Completed

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

func (c GeoradiusWithhashWithhash) Build() Completed {
	return Completed(c)
}

type Georadiusbymember Completed

func (c Georadiusbymember) Key(Key string) GeoradiusbymemberKey {
	return GeoradiusbymemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Georadiusbymember() (c Georadiusbymember) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER")
	return
}

type GeoradiusbymemberCountAnyAny Completed

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

type GeoradiusbymemberCountCount Completed

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

type GeoradiusbymemberKey Completed

func (c GeoradiusbymemberKey) Member(Member string) GeoradiusbymemberMember {
	return GeoradiusbymemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

type GeoradiusbymemberMember Completed

func (c GeoradiusbymemberMember) Radius(Radius float64) GeoradiusbymemberRadius {
	return GeoradiusbymemberRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusbymemberOrderAsc Completed

func (c GeoradiusbymemberOrderAsc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberOrderAsc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberOrderAsc) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberOrderDesc Completed

func (c GeoradiusbymemberOrderDesc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberOrderDesc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberOrderDesc) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberRadius Completed

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

type GeoradiusbymemberRo Completed

func (c GeoradiusbymemberRo) Key(Key string) GeoradiusbymemberRoKey {
	return GeoradiusbymemberRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) GeoradiusbymemberRo() (c GeoradiusbymemberRo) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER_RO")
	return
}

type GeoradiusbymemberRoCountAnyAny Completed

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

type GeoradiusbymemberRoCountCount Completed

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

type GeoradiusbymemberRoKey Completed

func (c GeoradiusbymemberRoKey) Member(Member string) GeoradiusbymemberRoMember {
	return GeoradiusbymemberRoMember{cf: c.cf, cs: append(c.cs, Member)}
}

type GeoradiusbymemberRoMember Completed

func (c GeoradiusbymemberRoMember) Radius(Radius float64) GeoradiusbymemberRoRadius {
	return GeoradiusbymemberRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusbymemberRoOrderAsc Completed

func (c GeoradiusbymemberRoOrderAsc) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoOrderDesc Completed

func (c GeoradiusbymemberRoOrderDesc) Storedist(Key string) GeoradiusbymemberRoStoredist {
	return GeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberRoOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoRadius Completed

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

type GeoradiusbymemberRoStoredist Completed

func (c GeoradiusbymemberRoStoredist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoStoredist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitFt Completed

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

func (c GeoradiusbymemberRoUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitKm Completed

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

func (c GeoradiusbymemberRoUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitM Completed

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

func (c GeoradiusbymemberRoUnitM) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoUnitMi Completed

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

func (c GeoradiusbymemberRoUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithcoordWithcoord Completed

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

func (c GeoradiusbymemberRoWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoWithcoordWithcoord) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithdistWithdist Completed

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

func (c GeoradiusbymemberRoWithdistWithdist) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoWithdistWithdist) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberRoWithhashWithhash Completed

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

func (c GeoradiusbymemberRoWithhashWithhash) Build() Completed {
	return Completed(c)
}

func (c GeoradiusbymemberRoWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type GeoradiusbymemberStore Completed

func (c GeoradiusbymemberStore) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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

func (c GeoradiusbymemberUnitFt) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitKm Completed

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

func (c GeoradiusbymemberUnitKm) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitM Completed

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

func (c GeoradiusbymemberUnitM) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitMi Completed

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

func (c GeoradiusbymemberUnitMi) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberWithcoordWithcoord Completed

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

func (c GeoradiusbymemberWithcoordWithcoord) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberWithdistWithdist Completed

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

func (c GeoradiusbymemberWithdistWithdist) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberWithhashWithhash Completed

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

func (c GeoradiusbymemberWithhashWithhash) Build() Completed {
	return Completed(c)
}

type Geosearch Completed

func (c Geosearch) Key(Key string) GeosearchKey {
	return GeosearchKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Geosearch() (c Geosearch) {
	c.cs = append(b.get(), "GEOSEARCH")
	return
}

type GeosearchBoxBybox Completed

func (c GeosearchBoxBybox) Height(Height float64) GeosearchBoxHeight {
	return GeosearchBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type GeosearchBoxHeight Completed

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

type GeosearchBoxUnitFt Completed

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

func (c GeosearchBoxUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitKm Completed

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

func (c GeosearchBoxUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitM Completed

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

func (c GeosearchBoxUnitM) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxUnitMi Completed

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

func (c GeosearchBoxUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeosearchBoxUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleByradius Completed

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

type GeosearchCircleUnitFt Completed

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

func (c GeosearchCircleUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitKm Completed

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

func (c GeosearchCircleUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitM Completed

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

func (c GeosearchCircleUnitM) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCircleUnitMi Completed

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

func (c GeosearchCircleUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeosearchCircleUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchCountAnyAny Completed

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

type GeosearchCountCount Completed

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

type GeosearchFromlonlat Completed

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

func (c GeosearchFromlonlat) Build() Completed {
	return Completed(c)
}

func (c GeosearchFromlonlat) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchFrommember Completed

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

func (c GeosearchFrommember) Build() Completed {
	return Completed(c)
}

func (c GeosearchFrommember) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchKey Completed

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

func (c GeosearchKey) Build() Completed {
	return Completed(c)
}

func (c GeosearchKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchOrderAsc Completed

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

func (c GeosearchOrderAsc) Build() Completed {
	return Completed(c)
}

func (c GeosearchOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchOrderDesc Completed

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

func (c GeosearchOrderDesc) Build() Completed {
	return Completed(c)
}

func (c GeosearchOrderDesc) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchWithcoordWithcoord Completed

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

type GeosearchWithdistWithdist Completed

func (c GeosearchWithdistWithdist) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
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
	return GeosearchstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Geosearchstore() (c Geosearchstore) {
	c.cs = append(b.get(), "GEOSEARCHSTORE")
	return
}

type GeosearchstoreBoxBybox Completed

func (c GeosearchstoreBoxBybox) Height(Height float64) GeosearchstoreBoxHeight {
	return GeosearchstoreBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type GeosearchstoreBoxHeight Completed

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

type GeosearchstoreBoxUnitFt Completed

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

func (c GeosearchstoreBoxUnitFt) Build() Completed {
	return Completed(c)
}

type GeosearchstoreBoxUnitKm Completed

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

func (c GeosearchstoreBoxUnitKm) Build() Completed {
	return Completed(c)
}

type GeosearchstoreBoxUnitM Completed

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

func (c GeosearchstoreBoxUnitM) Build() Completed {
	return Completed(c)
}

type GeosearchstoreBoxUnitMi Completed

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

func (c GeosearchstoreBoxUnitMi) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleByradius Completed

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

type GeosearchstoreCircleUnitFt Completed

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

func (c GeosearchstoreCircleUnitFt) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleUnitKm Completed

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

func (c GeosearchstoreCircleUnitKm) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleUnitM Completed

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

func (c GeosearchstoreCircleUnitM) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCircleUnitMi Completed

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

func (c GeosearchstoreCircleUnitMi) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCountAnyAny Completed

func (c GeosearchstoreCountAnyAny) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCountCount Completed

func (c GeosearchstoreCountCount) Any() GeosearchstoreCountAnyAny {
	return GeosearchstoreCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c GeosearchstoreCountCount) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreCountCount) Build() Completed {
	return Completed(c)
}

type GeosearchstoreDestination Completed

func (c GeosearchstoreDestination) Source(Source string) GeosearchstoreSource {
	return GeosearchstoreSource{cf: c.cf, cs: append(c.cs, Source)}
}

type GeosearchstoreFromlonlat Completed

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

func (c GeosearchstoreFromlonlat) Build() Completed {
	return Completed(c)
}

type GeosearchstoreFrommember Completed

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

func (c GeosearchstoreFrommember) Build() Completed {
	return Completed(c)
}

type GeosearchstoreOrderAsc Completed

func (c GeosearchstoreOrderAsc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderAsc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreOrderAsc) Build() Completed {
	return Completed(c)
}

type GeosearchstoreOrderDesc Completed

func (c GeosearchstoreOrderDesc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderDesc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreOrderDesc) Build() Completed {
	return Completed(c)
}

type GeosearchstoreSource Completed

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

func (c GeosearchstoreSource) Build() Completed {
	return Completed(c)
}

type GeosearchstoreStoredistStoredist Completed

func (c GeosearchstoreStoredistStoredist) Build() Completed {
	return Completed(c)
}

type Get Completed

func (c Get) Key(Key string) GetKey {
	return GetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Get() (c Get) {
	c.cs = append(b.get(), "GET")
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
	return GetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getbit() (c Getbit) {
	c.cs = append(b.get(), "GETBIT")
	return
}

type GetbitKey Completed

func (c GetbitKey) Offset(Offset int64) GetbitOffset {
	return GetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
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
	return GetdelKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return GetexKey{cf: c.cf, cs: append(c.cs, Key)}
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

type Getrange Completed

func (c Getrange) Key(Key string) GetrangeKey {
	return GetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getrange() (c Getrange) {
	c.cs = append(b.get(), "GETRANGE")
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
	return GetrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type GetrangeStart Completed

func (c GetrangeStart) End(End int64) GetrangeEnd {
	return GetrangeEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

type Getset Completed

func (c Getset) Key(Key string) GetsetKey {
	return GetsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Getset() (c Getset) {
	c.cs = append(b.get(), "GETSET")
	return
}

type GetsetKey Completed

func (c GetsetKey) Value(Value string) GetsetValue {
	return GetsetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type GetsetValue Completed

func (c GetsetValue) Build() Completed {
	return Completed(c)
}

type Hdel Completed

func (c Hdel) Key(Key string) HdelKey {
	return HdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hdel() (c Hdel) {
	c.cs = append(b.get(), "HDEL")
	return
}

type HdelField Completed

func (c HdelField) Field(Field ...string) HdelField {
	return HdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c HdelField) Build() Completed {
	return Completed(c)
}

type HdelKey Completed

func (c HdelKey) Field(Field ...string) HdelField {
	return HdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

type Hello Completed

func (c Hello) Protover(Protover int64) HelloArgumentsProtover {
	return HelloArgumentsProtover{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Protover, 10))}
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
	return HelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
}

func (c HelloArgumentsAuth) Build() Completed {
	return Completed(c)
}

type HelloArgumentsProtover Completed

func (c HelloArgumentsProtover) Auth(Username string, Password string) HelloArgumentsAuth {
	return HelloArgumentsAuth{cf: c.cf, cs: append(c.cs, "AUTH", Username, Password)}
}

func (c HelloArgumentsProtover) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
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
	return HexistsKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hexists() (c Hexists) {
	c.cs = append(b.get(), "HEXISTS")
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
	return HexistsField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hget Completed

func (c Hget) Key(Key string) HgetKey {
	return HgetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hget() (c Hget) {
	c.cs = append(b.get(), "HGET")
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
	return HgetField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hgetall Completed

func (c Hgetall) Key(Key string) HgetallKey {
	return HgetallKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hgetall() (c Hgetall) {
	c.cs = append(b.get(), "HGETALL")
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
	return HincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hincrby() (c Hincrby) {
	c.cs = append(b.get(), "HINCRBY")
	return
}

type HincrbyField Completed

func (c HincrbyField) Increment(Increment int64) HincrbyIncrement {
	return HincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type HincrbyIncrement Completed

func (c HincrbyIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyKey Completed

func (c HincrbyKey) Field(Field string) HincrbyField {
	return HincrbyField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hincrbyfloat Completed

func (c Hincrbyfloat) Key(Key string) HincrbyfloatKey {
	return HincrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hincrbyfloat() (c Hincrbyfloat) {
	c.cs = append(b.get(), "HINCRBYFLOAT")
	return
}

type HincrbyfloatField Completed

func (c HincrbyfloatField) Increment(Increment float64) HincrbyfloatIncrement {
	return HincrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type HincrbyfloatIncrement Completed

func (c HincrbyfloatIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyfloatKey Completed

func (c HincrbyfloatKey) Field(Field string) HincrbyfloatField {
	return HincrbyfloatField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hkeys Completed

func (c Hkeys) Key(Key string) HkeysKey {
	return HkeysKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hkeys() (c Hkeys) {
	c.cs = append(b.get(), "HKEYS")
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
	return HlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hlen() (c Hlen) {
	c.cs = append(b.get(), "HLEN")
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
	return HmgetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hmget() (c Hmget) {
	c.cs = append(b.get(), "HMGET")
	return
}

type HmgetField Completed

func (c HmgetField) Field(Field ...string) HmgetField {
	return HmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c HmgetField) Build() Completed {
	return Completed(c)
}

func (c HmgetField) Cache() Cacheable {
	return Cacheable(c)
}

type HmgetKey Completed

func (c HmgetKey) Field(Field ...string) HmgetField {
	return HmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

type Hmset Completed

func (c Hmset) Key(Key string) HmsetKey {
	return HmsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hmset() (c Hmset) {
	c.cs = append(b.get(), "HMSET")
	return
}

type HmsetFieldValue Completed

func (c HmsetFieldValue) FieldValue(Field string, Value string) HmsetFieldValue {
	return HmsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c HmsetFieldValue) Build() Completed {
	return Completed(c)
}

type HmsetKey Completed

func (c HmsetKey) FieldValue() HmsetFieldValue {
	return HmsetFieldValue{cf: c.cf, cs: c.cs}
}

type Hrandfield Completed

func (c Hrandfield) Key(Key string) HrandfieldKey {
	return HrandfieldKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hrandfield() (c Hrandfield) {
	c.cs = append(b.get(), "HRANDFIELD")
	return
}

type HrandfieldKey Completed

func (c HrandfieldKey) Count(Count int64) HrandfieldOptionsCount {
	return HrandfieldOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c HrandfieldKey) Build() Completed {
	return Completed(c)
}

type HrandfieldOptionsCount Completed

func (c HrandfieldOptionsCount) Withvalues() HrandfieldOptionsWithvaluesWithvalues {
	return HrandfieldOptionsWithvaluesWithvalues{cf: c.cf, cs: append(c.cs, "WITHVALUES")}
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
	return HscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hscan() (c Hscan) {
	c.cs = append(b.get(), "HSCAN")
	return
}

type HscanCount Completed

func (c HscanCount) Build() Completed {
	return Completed(c)
}

type HscanCursor Completed

func (c HscanCursor) Match(Pattern string) HscanMatch {
	return HscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c HscanCursor) Count(Count int64) HscanCount {
	return HscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c HscanCursor) Build() Completed {
	return Completed(c)
}

type HscanKey Completed

func (c HscanKey) Cursor(Cursor int64) HscanCursor {
	return HscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type HscanMatch Completed

func (c HscanMatch) Count(Count int64) HscanCount {
	return HscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c HscanMatch) Build() Completed {
	return Completed(c)
}

type Hset Completed

func (c Hset) Key(Key string) HsetKey {
	return HsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hset() (c Hset) {
	c.cs = append(b.get(), "HSET")
	return
}

type HsetFieldValue Completed

func (c HsetFieldValue) FieldValue(Field string, Value string) HsetFieldValue {
	return HsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c HsetFieldValue) Build() Completed {
	return Completed(c)
}

type HsetKey Completed

func (c HsetKey) FieldValue() HsetFieldValue {
	return HsetFieldValue{cf: c.cf, cs: c.cs}
}

type Hsetnx Completed

func (c Hsetnx) Key(Key string) HsetnxKey {
	return HsetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hsetnx() (c Hsetnx) {
	c.cs = append(b.get(), "HSETNX")
	return
}

type HsetnxField Completed

func (c HsetnxField) Value(Value string) HsetnxValue {
	return HsetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type HsetnxKey Completed

func (c HsetnxKey) Field(Field string) HsetnxField {
	return HsetnxField{cf: c.cf, cs: append(c.cs, Field)}
}

type HsetnxValue Completed

func (c HsetnxValue) Build() Completed {
	return Completed(c)
}

type Hstrlen Completed

func (c Hstrlen) Key(Key string) HstrlenKey {
	return HstrlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hstrlen() (c Hstrlen) {
	c.cs = append(b.get(), "HSTRLEN")
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
	return HstrlenField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hvals Completed

func (c Hvals) Key(Key string) HvalsKey {
	return HvalsKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Hvals() (c Hvals) {
	c.cs = append(b.get(), "HVALS")
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
	return IncrKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return IncrbyKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return IncrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type Incrbyfloat Completed

func (c Incrbyfloat) Key(Key string) IncrbyfloatKey {
	return IncrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return IncrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type Info Completed

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

type InfoSection Completed

func (c InfoSection) Build() Completed {
	return Completed(c)
}

type Keys Completed

func (c Keys) Pattern(Pattern string) KeysPattern {
	return KeysPattern{cf: c.cf, cs: append(c.cs, Pattern)}
}

func (b *Builder) Keys() (c Keys) {
	c.cs = append(b.get(), "KEYS")
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
	return LatencyGraphEvent{cf: c.cf, cs: append(c.cs, Event)}
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
	return LatencyHistoryEvent{cf: c.cf, cs: append(c.cs, Event)}
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
	return LatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
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
	return LatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
}

func (c LatencyResetEvent) Build() Completed {
	return Completed(c)
}

type Lindex Completed

func (c Lindex) Key(Key string) LindexKey {
	return LindexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lindex() (c Lindex) {
	c.cs = append(b.get(), "LINDEX")
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
	return LindexIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type Linsert Completed

func (c Linsert) Key(Key string) LinsertKey {
	return LinsertKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return LinsertWhereBefore{cf: c.cf, cs: append(c.cs, "BEFORE")}
}

func (c LinsertKey) After() LinsertWhereAfter {
	return LinsertWhereAfter{cf: c.cf, cs: append(c.cs, "AFTER")}
}

type LinsertPivot Completed

func (c LinsertPivot) Element(Element string) LinsertElement {
	return LinsertElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LinsertWhereAfter Completed

func (c LinsertWhereAfter) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type LinsertWhereBefore Completed

func (c LinsertWhereBefore) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type Llen Completed

func (c Llen) Key(Key string) LlenKey {
	return LlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Llen() (c Llen) {
	c.cs = append(b.get(), "LLEN")
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
	return LmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Lmove() (c Lmove) {
	c.cs = append(b.get(), "LMOVE")
	return
}

type LmoveDestination Completed

func (c LmoveDestination) Left() LmoveWherefromLeft {
	return LmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveDestination) Right() LmoveWherefromRight {
	return LmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveSource Completed

func (c LmoveSource) Destination(Destination string) LmoveDestination {
	return LmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type LmoveWherefromLeft Completed

func (c LmoveWherefromLeft) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromLeft) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveWherefromRight Completed

func (c LmoveWherefromRight) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromRight) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
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
	return LmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
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
	return LmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmpopKey) Right() LmpopWhereRight {
	return LmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c LmpopKey) Key(Key ...string) LmpopKey {
	return LmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type LmpopNumkeys Completed

func (c LmpopNumkeys) Key(Key ...string) LmpopKey {
	return LmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c LmpopNumkeys) Left() LmpopWhereLeft {
	return LmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmpopNumkeys) Right() LmpopWhereRight {
	return LmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmpopWhereLeft Completed

func (c LmpopWhereLeft) Count(Count int64) LmpopCount {
	return LmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type LmpopWhereRight Completed

func (c LmpopWhereRight) Count(Count int64) LmpopCount {
	return LmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Lolwut Completed

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

type LolwutVersion Completed

func (c LolwutVersion) Build() Completed {
	return Completed(c)
}

type Lpop Completed

func (c Lpop) Key(Key string) LpopKey {
	return LpopKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return LpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c LpopKey) Build() Completed {
	return Completed(c)
}

type Lpos Completed

func (c Lpos) Key(Key string) LposKey {
	return LposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lpos() (c Lpos) {
	c.cs = append(b.get(), "LPOS")
	return
}

type LposCount Completed

func (c LposCount) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c LposCount) Build() Completed {
	return Completed(c)
}

func (c LposCount) Cache() Cacheable {
	return Cacheable(c)
}

type LposElement Completed

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

type LposKey Completed

func (c LposKey) Element(Element string) LposElement {
	return LposElement{cf: c.cf, cs: append(c.cs, Element)}
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

type Lpush Completed

func (c Lpush) Key(Key string) LpushKey {
	return LpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lpush() (c Lpush) {
	c.cs = append(b.get(), "LPUSH")
	return
}

type LpushElement Completed

func (c LpushElement) Element(Element ...string) LpushElement {
	return LpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c LpushElement) Build() Completed {
	return Completed(c)
}

type LpushKey Completed

func (c LpushKey) Element(Element ...string) LpushElement {
	return LpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Lpushx Completed

func (c Lpushx) Key(Key string) LpushxKey {
	return LpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lpushx() (c Lpushx) {
	c.cs = append(b.get(), "LPUSHX")
	return
}

type LpushxElement Completed

func (c LpushxElement) Element(Element ...string) LpushxElement {
	return LpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c LpushxElement) Build() Completed {
	return Completed(c)
}

type LpushxKey Completed

func (c LpushxKey) Element(Element ...string) LpushxElement {
	return LpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Lrange Completed

func (c Lrange) Key(Key string) LrangeKey {
	return LrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lrange() (c Lrange) {
	c.cs = append(b.get(), "LRANGE")
	return
}

type LrangeKey Completed

func (c LrangeKey) Start(Start int64) LrangeStart {
	return LrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type LrangeStart Completed

func (c LrangeStart) Stop(Stop int64) LrangeStop {
	return LrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
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
	return LremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Lrem() (c Lrem) {
	c.cs = append(b.get(), "LREM")
	return
}

type LremCount Completed

func (c LremCount) Element(Element string) LremElement {
	return LremElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LremElement Completed

func (c LremElement) Build() Completed {
	return Completed(c)
}

type LremKey Completed

func (c LremKey) Count(Count int64) LremCount {
	return LremCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type Lset Completed

func (c Lset) Key(Key string) LsetKey {
	return LsetKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return LsetElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LsetKey Completed

func (c LsetKey) Index(Index int64) LsetIndex {
	return LsetIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type Ltrim Completed

func (c Ltrim) Key(Key string) LtrimKey {
	return LtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Ltrim() (c Ltrim) {
	c.cs = append(b.get(), "LTRIM")
	return
}

type LtrimKey Completed

func (c LtrimKey) Start(Start int64) LtrimStart {
	return LtrimStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type LtrimStart Completed

func (c LtrimStart) Stop(Stop int64) LtrimStop {
	return LtrimStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
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
	return MemoryUsageKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) MemoryUsage() (c MemoryUsage) {
	c.cs = append(b.get(), "MEMORY", "USAGE")
	return
}

type MemoryUsageKey Completed

func (c MemoryUsageKey) Samples(Count int64) MemoryUsageSamples {
	return MemoryUsageSamples{cf: c.cf, cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10))}
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
	return MgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Mget() (c Mget) {
	c.cs = append(b.get(), "MGET")
	return
}

type MgetKey Completed

func (c MgetKey) Key(Key ...string) MgetKey {
	return MgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MgetKey) Build() Completed {
	return Completed(c)
}

type Migrate Completed

func (c Migrate) Host(Host string) MigrateHost {
	return MigrateHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *Builder) Migrate() (c Migrate) {
	c.cs = append(b.get(), "MIGRATE")
	c.cf = blockTag
	return
}

type MigrateAuth Completed

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

type MigrateAuth2 Completed

func (c MigrateAuth2) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MigrateAuth2) Build() Completed {
	return Completed(c)
}

type MigrateCopyCopy Completed

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

type MigrateDestinationDb Completed

func (c MigrateDestinationDb) Timeout(Timeout int64) MigrateTimeout {
	return MigrateTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type MigrateHost Completed

func (c MigrateHost) Port(Port string) MigratePort {
	return MigratePort{cf: c.cf, cs: append(c.cs, Port)}
}

type MigrateKeyEmpty Completed

func (c MigrateKeyEmpty) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeyKey Completed

func (c MigrateKeyKey) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeys Completed

func (c MigrateKeys) Keys(Keys ...string) MigrateKeys {
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Keys...)}
}

func (c MigrateKeys) Build() Completed {
	return Completed(c)
}

type MigratePort Completed

func (c MigratePort) Key() MigrateKeyKey {
	return MigrateKeyKey{cf: c.cf, cs: append(c.cs, "key")}
}

func (c MigratePort) Empty() MigrateKeyEmpty {
	return MigrateKeyEmpty{cf: c.cf, cs: append(c.cs, "\"\"")}
}

type MigrateReplaceReplace Completed

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

type MigrateTimeout Completed

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
	return ModuleLoadPath{cf: c.cf, cs: append(c.cs, Path)}
}

func (b *Builder) ModuleLoad() (c ModuleLoad) {
	c.cs = append(b.get(), "MODULE", "LOAD")
	return
}

type ModuleLoadArg Completed

func (c ModuleLoadArg) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c ModuleLoadArg) Build() Completed {
	return Completed(c)
}

type ModuleLoadPath Completed

func (c ModuleLoadPath) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c ModuleLoadPath) Build() Completed {
	return Completed(c)
}

type ModuleUnload Completed

func (c ModuleUnload) Name(Name string) ModuleUnloadName {
	return ModuleUnloadName{cf: c.cf, cs: append(c.cs, Name)}
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
	return MoveKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return MoveDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Db, 10))}
}

type Mset Completed

func (c Mset) KeyValue() MsetKeyValue {
	return MsetKeyValue{cf: c.cf, cs: c.cs}
}

func (b *Builder) Mset() (c Mset) {
	c.cs = append(b.get(), "MSET")
	return
}

type MsetKeyValue Completed

func (c MsetKeyValue) KeyValue(Key string, Value string) MsetKeyValue {
	return MsetKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c MsetKeyValue) Build() Completed {
	return Completed(c)
}

type Msetnx Completed

func (c Msetnx) KeyValue() MsetnxKeyValue {
	return MsetnxKeyValue{cf: c.cf, cs: c.cs}
}

func (b *Builder) Msetnx() (c Msetnx) {
	c.cs = append(b.get(), "MSETNX")
	return
}

type MsetnxKeyValue Completed

func (c MsetnxKeyValue) KeyValue(Key string, Value string) MsetnxKeyValue {
	return MsetnxKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
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
	return ObjectSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *Builder) Object() (c Object) {
	c.cs = append(b.get(), "OBJECT")
	return
}

type ObjectArguments Completed

func (c ObjectArguments) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c ObjectArguments) Build() Completed {
	return Completed(c)
}

type ObjectSubcommand Completed

func (c ObjectSubcommand) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c ObjectSubcommand) Build() Completed {
	return Completed(c)
}

type Persist Completed

func (c Persist) Key(Key string) PersistKey {
	return PersistKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return PexpireKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return PexpireMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PexpireMilliseconds Completed

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

type Pexpireat Completed

func (c Pexpireat) Key(Key string) PexpireatKey {
	return PexpireatKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return PexpireatMillisecondsTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10))}
}

type PexpireatMillisecondsTimestamp Completed

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

type Pexpiretime Completed

func (c Pexpiretime) Key(Key string) PexpiretimeKey {
	return PexpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pexpiretime() (c Pexpiretime) {
	c.cs = append(b.get(), "PEXPIRETIME")
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
	return PfaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pfadd() (c Pfadd) {
	c.cs = append(b.get(), "PFADD")
	return
}

type PfaddElement Completed

func (c PfaddElement) Element(Element ...string) PfaddElement {
	return PfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c PfaddElement) Build() Completed {
	return Completed(c)
}

type PfaddKey Completed

func (c PfaddKey) Element(Element ...string) PfaddElement {
	return PfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c PfaddKey) Build() Completed {
	return Completed(c)
}

type Pfcount Completed

func (c Pfcount) Key(Key ...string) PfcountKey {
	return PfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Pfcount() (c Pfcount) {
	c.cs = append(b.get(), "PFCOUNT")
	return
}

type PfcountKey Completed

func (c PfcountKey) Key(Key ...string) PfcountKey {
	return PfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c PfcountKey) Build() Completed {
	return Completed(c)
}

type Pfmerge Completed

func (c Pfmerge) Destkey(Destkey string) PfmergeDestkey {
	return PfmergeDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

func (b *Builder) Pfmerge() (c Pfmerge) {
	c.cs = append(b.get(), "PFMERGE")
	return
}

type PfmergeDestkey Completed

func (c PfmergeDestkey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

type PfmergeSourcekey Completed

func (c PfmergeSourcekey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

func (c PfmergeSourcekey) Build() Completed {
	return Completed(c)
}

type Ping Completed

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

type PingMessage Completed

func (c PingMessage) Build() Completed {
	return Completed(c)
}

type Psetex Completed

func (c Psetex) Key(Key string) PsetexKey {
	return PsetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Psetex() (c Psetex) {
	c.cs = append(b.get(), "PSETEX")
	return
}

type PsetexKey Completed

func (c PsetexKey) Milliseconds(Milliseconds int64) PsetexMilliseconds {
	return PsetexMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PsetexMilliseconds Completed

func (c PsetexMilliseconds) Value(Value string) PsetexValue {
	return PsetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type PsetexValue Completed

func (c PsetexValue) Build() Completed {
	return Completed(c)
}

type Psubscribe Completed

func (c Psubscribe) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (b *Builder) Psubscribe() (c Psubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	c.cf = noRetTag
	return
}

type PsubscribePattern Completed

func (c PsubscribePattern) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c PsubscribePattern) Build() Completed {
	return Completed(c)
}

type Psync Completed

func (c Psync) Replicationid(Replicationid int64) PsyncReplicationid {
	return PsyncReplicationid{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Replicationid, 10))}
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
	return PsyncOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type Pttl Completed

func (c Pttl) Key(Key string) PttlKey {
	return PttlKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Pttl() (c Pttl) {
	c.cs = append(b.get(), "PTTL")
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
	return PublishChannel{cf: c.cf, cs: append(c.cs, Channel)}
}

func (b *Builder) Publish() (c Publish) {
	c.cs = append(b.get(), "PUBLISH")
	return
}

type PublishChannel Completed

func (c PublishChannel) Message(Message string) PublishMessage {
	return PublishMessage{cf: c.cf, cs: append(c.cs, Message)}
}

type PublishMessage Completed

func (c PublishMessage) Build() Completed {
	return Completed(c)
}

type Pubsub Completed

func (c Pubsub) Subcommand(Subcommand string) PubsubSubcommand {
	return PubsubSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *Builder) Pubsub() (c Pubsub) {
	c.cs = append(b.get(), "PUBSUB")
	return
}

type PubsubArgument Completed

func (c PubsubArgument) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c PubsubArgument) Build() Completed {
	return Completed(c)
}

type PubsubSubcommand Completed

func (c PubsubSubcommand) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c PubsubSubcommand) Build() Completed {
	return Completed(c)
}

type Punsubscribe Completed

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

type PunsubscribePattern Completed

func (c PunsubscribePattern) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
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
	return RenameKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Rename() (c Rename) {
	c.cs = append(b.get(), "RENAME")
	return
}

type RenameKey Completed

func (c RenameKey) Newkey(Newkey string) RenameNewkey {
	return RenameNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type RenameNewkey Completed

func (c RenameNewkey) Build() Completed {
	return Completed(c)
}

type Renamenx Completed

func (c Renamenx) Key(Key string) RenamenxKey {
	return RenamenxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Renamenx() (c Renamenx) {
	c.cs = append(b.get(), "RENAMENX")
	return
}

type RenamenxKey Completed

func (c RenamenxKey) Newkey(Newkey string) RenamenxNewkey {
	return RenamenxNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type RenamenxNewkey Completed

func (c RenamenxNewkey) Build() Completed {
	return Completed(c)
}

type Replicaof Completed

func (c Replicaof) Host(Host string) ReplicaofHost {
	return ReplicaofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *Builder) Replicaof() (c Replicaof) {
	c.cs = append(b.get(), "REPLICAOF")
	return
}

type ReplicaofHost Completed

func (c ReplicaofHost) Port(Port string) ReplicaofPort {
	return ReplicaofPort{cf: c.cf, cs: append(c.cs, Port)}
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
	return RestoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Restore() (c Restore) {
	c.cs = append(b.get(), "RESTORE")
	return
}

type RestoreAbsttlAbsttl Completed

func (c RestoreAbsttlAbsttl) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c RestoreAbsttlAbsttl) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
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
	return RestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreIdletime) Build() Completed {
	return Completed(c)
}

type RestoreKey Completed

func (c RestoreKey) Ttl(Ttl int64) RestoreTtl {
	return RestoreTtl{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Ttl, 10))}
}

type RestoreReplaceReplace Completed

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

type RestoreSerializedValue Completed

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

type RestoreTtl Completed

func (c RestoreTtl) SerializedValue(SerializedValue string) RestoreSerializedValue {
	return RestoreSerializedValue{cf: c.cf, cs: append(c.cs, SerializedValue)}
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
	return RpopKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return RpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c RpopKey) Build() Completed {
	return Completed(c)
}

type Rpoplpush Completed

func (c Rpoplpush) Source(Source string) RpoplpushSource {
	return RpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
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
	return RpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Rpush Completed

func (c Rpush) Key(Key string) RpushKey {
	return RpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Rpush() (c Rpush) {
	c.cs = append(b.get(), "RPUSH")
	return
}

type RpushElement Completed

func (c RpushElement) Element(Element ...string) RpushElement {
	return RpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c RpushElement) Build() Completed {
	return Completed(c)
}

type RpushKey Completed

func (c RpushKey) Element(Element ...string) RpushElement {
	return RpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Rpushx Completed

func (c Rpushx) Key(Key string) RpushxKey {
	return RpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Rpushx() (c Rpushx) {
	c.cs = append(b.get(), "RPUSHX")
	return
}

type RpushxElement Completed

func (c RpushxElement) Element(Element ...string) RpushxElement {
	return RpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c RpushxElement) Build() Completed {
	return Completed(c)
}

type RpushxKey Completed

func (c RpushxKey) Element(Element ...string) RpushxElement {
	return RpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Sadd Completed

func (c Sadd) Key(Key string) SaddKey {
	return SaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sadd() (c Sadd) {
	c.cs = append(b.get(), "SADD")
	return
}

type SaddKey Completed

func (c SaddKey) Member(Member ...string) SaddMember {
	return SaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SaddMember Completed

func (c SaddMember) Member(Member ...string) SaddMember {
	return SaddMember{cf: c.cf, cs: append(c.cs, Member...)}
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
	return ScanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

func (b *Builder) Scan() (c Scan) {
	c.cs = append(b.get(), "SCAN")
	return
}

type ScanCount Completed

func (c ScanCount) Type(Type string) ScanType {
	return ScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c ScanCount) Build() Completed {
	return Completed(c)
}

type ScanCursor Completed

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

type ScanMatch Completed

func (c ScanMatch) Count(Count int64) ScanCount {
	return ScanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ScanMatch) Type(Type string) ScanType {
	return ScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
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
	return ScardKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Scard() (c Scard) {
	c.cs = append(b.get(), "SCARD")
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
	return ScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (b *Builder) ScriptExists() (c ScriptExists) {
	c.cs = append(b.get(), "SCRIPT", "EXISTS")
	return
}

type ScriptExistsSha1 Completed

func (c ScriptExistsSha1) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (c ScriptExistsSha1) Build() Completed {
	return Completed(c)
}

type ScriptFlush Completed

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
	return ScriptLoadScript{cf: c.cf, cs: append(c.cs, Script)}
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
	return SdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sdiff() (c Sdiff) {
	c.cs = append(b.get(), "SDIFF")
	return
}

type SdiffKey Completed

func (c SdiffKey) Key(Key ...string) SdiffKey {
	return SdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SdiffKey) Build() Completed {
	return Completed(c)
}

type Sdiffstore Completed

func (c Sdiffstore) Destination(Destination string) SdiffstoreDestination {
	return SdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Sdiffstore() (c Sdiffstore) {
	c.cs = append(b.get(), "SDIFFSTORE")
	return
}

type SdiffstoreDestination Completed

func (c SdiffstoreDestination) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SdiffstoreKey Completed

func (c SdiffstoreKey) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SdiffstoreKey) Build() Completed {
	return Completed(c)
}

type Select Completed

func (c Select) Index(Index int64) SelectIndex {
	return SelectIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
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
	return SetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Set() (c Set) {
	c.cs = append(b.get(), "SET")
	return
}

type SetConditionNx Completed

func (c SetConditionNx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetConditionNx) Build() Completed {
	return Completed(c)
}

type SetConditionXx Completed

func (c SetConditionXx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetConditionXx) Build() Completed {
	return Completed(c)
}

type SetExpirationEx Completed

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

type SetExpirationExat Completed

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

type SetExpirationKeepttl Completed

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

type SetExpirationPx Completed

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

type SetExpirationPxat Completed

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

type SetGetGet Completed

func (c SetGetGet) Build() Completed {
	return Completed(c)
}

type SetKey Completed

func (c SetKey) Value(Value string) SetValue {
	return SetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetValue Completed

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

type Setbit Completed

func (c Setbit) Key(Key string) SetbitKey {
	return SetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setbit() (c Setbit) {
	c.cs = append(b.get(), "SETBIT")
	return
}

type SetbitKey Completed

func (c SetbitKey) Offset(Offset int64) SetbitOffset {
	return SetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetbitOffset Completed

func (c SetbitOffset) Value(Value int64) SetbitValue {
	return SetbitValue{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Value, 10))}
}

type SetbitValue Completed

func (c SetbitValue) Build() Completed {
	return Completed(c)
}

type Setex Completed

func (c Setex) Key(Key string) SetexKey {
	return SetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setex() (c Setex) {
	c.cs = append(b.get(), "SETEX")
	return
}

type SetexKey Completed

func (c SetexKey) Seconds(Seconds int64) SetexSeconds {
	return SetexSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SetexSeconds Completed

func (c SetexSeconds) Value(Value string) SetexValue {
	return SetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetexValue Completed

func (c SetexValue) Build() Completed {
	return Completed(c)
}

type Setnx Completed

func (c Setnx) Key(Key string) SetnxKey {
	return SetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setnx() (c Setnx) {
	c.cs = append(b.get(), "SETNX")
	return
}

type SetnxKey Completed

func (c SetnxKey) Value(Value string) SetnxValue {
	return SetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetnxValue Completed

func (c SetnxValue) Build() Completed {
	return Completed(c)
}

type Setrange Completed

func (c Setrange) Key(Key string) SetrangeKey {
	return SetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Setrange() (c Setrange) {
	c.cs = append(b.get(), "SETRANGE")
	return
}

type SetrangeKey Completed

func (c SetrangeKey) Offset(Offset int64) SetrangeOffset {
	return SetrangeOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetrangeOffset Completed

func (c SetrangeOffset) Value(Value string) SetrangeValue {
	return SetrangeValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetrangeValue Completed

func (c SetrangeValue) Build() Completed {
	return Completed(c)
}

type Shutdown Completed

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
	return SinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sinter() (c Sinter) {
	c.cs = append(b.get(), "SINTER")
	return
}

type SinterKey Completed

func (c SinterKey) Key(Key ...string) SinterKey {
	return SinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SinterKey) Build() Completed {
	return Completed(c)
}

type Sintercard Completed

func (c Sintercard) Key(Key ...string) SintercardKey {
	return SintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sintercard() (c Sintercard) {
	c.cs = append(b.get(), "SINTERCARD")
	return
}

type SintercardKey Completed

func (c SintercardKey) Key(Key ...string) SintercardKey {
	return SintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SintercardKey) Build() Completed {
	return Completed(c)
}

type Sinterstore Completed

func (c Sinterstore) Destination(Destination string) SinterstoreDestination {
	return SinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Sinterstore() (c Sinterstore) {
	c.cs = append(b.get(), "SINTERSTORE")
	return
}

type SinterstoreDestination Completed

func (c SinterstoreDestination) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SinterstoreKey Completed

func (c SinterstoreKey) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SinterstoreKey) Build() Completed {
	return Completed(c)
}

type Sismember Completed

func (c Sismember) Key(Key string) SismemberKey {
	return SismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sismember() (c Sismember) {
	c.cs = append(b.get(), "SISMEMBER")
	return
}

type SismemberKey Completed

func (c SismemberKey) Member(Member string) SismemberMember {
	return SismemberMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return SlaveofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *Builder) Slaveof() (c Slaveof) {
	c.cs = append(b.get(), "SLAVEOF")
	return
}

type SlaveofHost Completed

func (c SlaveofHost) Port(Port string) SlaveofPort {
	return SlaveofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type SlaveofPort Completed

func (c SlaveofPort) Build() Completed {
	return Completed(c)
}

type Slowlog Completed

func (c Slowlog) Subcommand(Subcommand string) SlowlogSubcommand {
	return SlowlogSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
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
	return SlowlogArgument{cf: c.cf, cs: append(c.cs, Argument)}
}

func (c SlowlogSubcommand) Build() Completed {
	return Completed(c)
}

type Smembers Completed

func (c Smembers) Key(Key string) SmembersKey {
	return SmembersKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Smembers() (c Smembers) {
	c.cs = append(b.get(), "SMEMBERS")
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
	return SmismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Smismember() (c Smismember) {
	c.cs = append(b.get(), "SMISMEMBER")
	return
}

type SmismemberKey Completed

func (c SmismemberKey) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SmismemberMember Completed

func (c SmismemberMember) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SmismemberMember) Build() Completed {
	return Completed(c)
}

func (c SmismemberMember) Cache() Cacheable {
	return Cacheable(c)
}

type Smove Completed

func (c Smove) Source(Source string) SmoveSource {
	return SmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Smove() (c Smove) {
	c.cs = append(b.get(), "SMOVE")
	return
}

type SmoveDestination Completed

func (c SmoveDestination) Member(Member string) SmoveMember {
	return SmoveMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SmoveMember Completed

func (c SmoveMember) Build() Completed {
	return Completed(c)
}

type SmoveSource Completed

func (c SmoveSource) Destination(Destination string) SmoveDestination {
	return SmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Sort Completed

func (c Sort) Key(Key string) SortKey {
	return SortKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sort() (c Sort) {
	c.cs = append(b.get(), "SORT")
	return
}

type SortBy Completed

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

type SortGet Completed

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

type SortKey Completed

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

type SortLimit Completed

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

type SortOrderAsc Completed

func (c SortOrderAsc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortOrderAsc) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortOrderAsc) Build() Completed {
	return Completed(c)
}

type SortOrderDesc Completed

func (c SortOrderDesc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortOrderDesc) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortOrderDesc) Build() Completed {
	return Completed(c)
}

type SortRo Completed

func (c SortRo) Key(Key string) SortRoKey {
	return SortRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) SortRo() (c SortRo) {
	c.cs = append(b.get(), "SORT_RO")
	return
}

type SortRoBy Completed

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

type SortRoGet Completed

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

type SortRoKey Completed

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

type SortRoLimit Completed

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

type SortRoOrderAsc Completed

func (c SortRoOrderAsc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoOrderAsc) Build() Completed {
	return Completed(c)
}

func (c SortRoOrderAsc) Cache() Cacheable {
	return Cacheable(c)
}

type SortRoOrderDesc Completed

func (c SortRoOrderDesc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
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
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
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
	return SpopKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SpopKey) Build() Completed {
	return Completed(c)
}

type Srandmember Completed

func (c Srandmember) Key(Key string) SrandmemberKey {
	return SrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Srandmember() (c Srandmember) {
	c.cs = append(b.get(), "SRANDMEMBER")
	return
}

type SrandmemberCount Completed

func (c SrandmemberCount) Build() Completed {
	return Completed(c)
}

type SrandmemberKey Completed

func (c SrandmemberKey) Count(Count int64) SrandmemberCount {
	return SrandmemberCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SrandmemberKey) Build() Completed {
	return Completed(c)
}

type Srem Completed

func (c Srem) Key(Key string) SremKey {
	return SremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Srem() (c Srem) {
	c.cs = append(b.get(), "SREM")
	return
}

type SremKey Completed

func (c SremKey) Member(Member ...string) SremMember {
	return SremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SremMember Completed

func (c SremMember) Member(Member ...string) SremMember {
	return SremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SremMember) Build() Completed {
	return Completed(c)
}

type Sscan Completed

func (c Sscan) Key(Key string) SscanKey {
	return SscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Sscan() (c Sscan) {
	c.cs = append(b.get(), "SSCAN")
	return
}

type SscanCount Completed

func (c SscanCount) Build() Completed {
	return Completed(c)
}

type SscanCursor Completed

func (c SscanCursor) Match(Pattern string) SscanMatch {
	return SscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SscanCursor) Count(Count int64) SscanCount {
	return SscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SscanCursor) Build() Completed {
	return Completed(c)
}

type SscanKey Completed

func (c SscanKey) Cursor(Cursor int64) SscanCursor {
	return SscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SscanMatch Completed

func (c SscanMatch) Count(Count int64) SscanCount {
	return SscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SscanMatch) Build() Completed {
	return Completed(c)
}

type Stralgo Completed

func (c Stralgo) Lcs() StralgoAlgorithmLcs {
	return StralgoAlgorithmLcs{cf: c.cf, cs: append(c.cs, "LCS")}
}

func (b *Builder) Stralgo() (c Stralgo) {
	c.cs = append(b.get(), "STRALGO")
	return
}

type StralgoAlgoSpecificArgument Completed

func (c StralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

func (c StralgoAlgoSpecificArgument) Build() Completed {
	return Completed(c)
}

type StralgoAlgorithmLcs Completed

func (c StralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

type Strlen Completed

func (c Strlen) Key(Key string) StrlenKey {
	return StrlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Strlen() (c Strlen) {
	c.cs = append(b.get(), "STRLEN")
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
	return SubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (b *Builder) Subscribe() (c Subscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	c.cf = noRetTag
	return
}

type SubscribeChannel Completed

func (c SubscribeChannel) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c SubscribeChannel) Build() Completed {
	return Completed(c)
}

type Sunion Completed

func (c Sunion) Key(Key ...string) SunionKey {
	return SunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Sunion() (c Sunion) {
	c.cs = append(b.get(), "SUNION")
	return
}

type SunionKey Completed

func (c SunionKey) Key(Key ...string) SunionKey {
	return SunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SunionKey) Build() Completed {
	return Completed(c)
}

type Sunionstore Completed

func (c Sunionstore) Destination(Destination string) SunionstoreDestination {
	return SunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Sunionstore() (c Sunionstore) {
	c.cs = append(b.get(), "SUNIONSTORE")
	return
}

type SunionstoreDestination Completed

func (c SunionstoreDestination) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SunionstoreKey Completed

func (c SunionstoreKey) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SunionstoreKey) Build() Completed {
	return Completed(c)
}

type Swapdb Completed

func (c Swapdb) Index1(Index1 int64) SwapdbIndex1 {
	return SwapdbIndex1{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index1, 10))}
}

func (b *Builder) Swapdb() (c Swapdb) {
	c.cs = append(b.get(), "SWAPDB")
	return
}

type SwapdbIndex1 Completed

func (c SwapdbIndex1) Index2(Index2 int64) SwapdbIndex2 {
	return SwapdbIndex2{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index2, 10))}
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
	return TouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Touch() (c Touch) {
	c.cs = append(b.get(), "TOUCH")
	return
}

type TouchKey Completed

func (c TouchKey) Key(Key ...string) TouchKey {
	return TouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c TouchKey) Build() Completed {
	return Completed(c)
}

type Ttl Completed

func (c Ttl) Key(Key string) TtlKey {
	return TtlKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Ttl() (c Ttl) {
	c.cs = append(b.get(), "TTL")
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
	return TypeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Type() (c Type) {
	c.cs = append(b.get(), "TYPE")
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
	return UnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Unlink() (c Unlink) {
	c.cs = append(b.get(), "UNLINK")
	return
}

type UnlinkKey Completed

func (c UnlinkKey) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c UnlinkKey) Build() Completed {
	return Completed(c)
}

type Unsubscribe Completed

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

type UnsubscribeChannel Completed

func (c UnsubscribeChannel) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
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
	return WaitNumreplicas{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numreplicas, 10))}
}

func (b *Builder) Wait() (c Wait) {
	c.cs = append(b.get(), "WAIT")
	c.cf = blockTag
	return
}

type WaitNumreplicas Completed

func (c WaitNumreplicas) Timeout(Timeout int64) WaitTimeout {
	return WaitTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type WaitTimeout Completed

func (c WaitTimeout) Build() Completed {
	return Completed(c)
}

type Watch Completed

func (c Watch) Key(Key ...string) WatchKey {
	return WatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Watch() (c Watch) {
	c.cs = append(b.get(), "WATCH")
	return
}

type WatchKey Completed

func (c WatchKey) Key(Key ...string) WatchKey {
	return WatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c WatchKey) Build() Completed {
	return Completed(c)
}

type Xack Completed

func (c Xack) Key(Key string) XackKey {
	return XackKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xack() (c Xack) {
	c.cs = append(b.get(), "XACK")
	return
}

type XackGroup Completed

func (c XackGroup) Id(Id ...string) XackId {
	return XackId{cf: c.cf, cs: append(c.cs, Id...)}
}

type XackId Completed

func (c XackId) Id(Id ...string) XackId {
	return XackId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XackId) Build() Completed {
	return Completed(c)
}

type XackKey Completed

func (c XackKey) Group(Group string) XackGroup {
	return XackGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type Xadd Completed

func (c Xadd) Key(Key string) XaddKey {
	return XaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xadd() (c Xadd) {
	c.cs = append(b.get(), "XADD")
	return
}

type XaddFieldValue Completed

func (c XaddFieldValue) FieldValue(Field string, Value string) XaddFieldValue {
	return XaddFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c XaddFieldValue) Build() Completed {
	return Completed(c)
}

type XaddId Completed

func (c XaddId) FieldValue() XaddFieldValue {
	return XaddFieldValue{cf: c.cf, cs: c.cs}
}

type XaddKey Completed

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

type XaddNomkstream Completed

func (c XaddNomkstream) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c XaddNomkstream) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

func (c XaddNomkstream) Id(Id string) XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type XaddTrimLimit Completed

func (c XaddTrimLimit) Id(Id string) XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type XaddTrimOperatorAlmost Completed

func (c XaddTrimOperatorAlmost) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimOperatorExact Completed

func (c XaddTrimOperatorExact) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimStrategyMaxlen Completed

func (c XaddTrimStrategyMaxlen) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XaddTrimStrategyMaxlen) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XaddTrimStrategyMaxlen) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimStrategyMinid Completed

func (c XaddTrimStrategyMinid) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XaddTrimStrategyMinid) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XaddTrimStrategyMinid) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimThreshold Completed

func (c XaddTrimThreshold) Limit(Count int64) XaddTrimLimit {
	return XaddTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XaddTrimThreshold) Id(Id string) XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type Xautoclaim Completed

func (c Xautoclaim) Key(Key string) XautoclaimKey {
	return XautoclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xautoclaim() (c Xautoclaim) {
	c.cs = append(b.get(), "XAUTOCLAIM")
	return
}

type XautoclaimConsumer Completed

func (c XautoclaimConsumer) MinIdleTime(MinIdleTime string) XautoclaimMinIdleTime {
	return XautoclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type XautoclaimCount Completed

func (c XautoclaimCount) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XautoclaimCount) Build() Completed {
	return Completed(c)
}

type XautoclaimGroup Completed

func (c XautoclaimGroup) Consumer(Consumer string) XautoclaimConsumer {
	return XautoclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type XautoclaimJustidJustid Completed

func (c XautoclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XautoclaimKey Completed

func (c XautoclaimKey) Group(Group string) XautoclaimGroup {
	return XautoclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type XautoclaimMinIdleTime Completed

func (c XautoclaimMinIdleTime) Start(Start string) XautoclaimStart {
	return XautoclaimStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XautoclaimStart Completed

func (c XautoclaimStart) Count(Count int64) XautoclaimCount {
	return XautoclaimCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XautoclaimStart) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XautoclaimStart) Build() Completed {
	return Completed(c)
}

type Xclaim Completed

func (c Xclaim) Key(Key string) XclaimKey {
	return XclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xclaim() (c Xclaim) {
	c.cs = append(b.get(), "XCLAIM")
	return
}

type XclaimConsumer Completed

func (c XclaimConsumer) MinIdleTime(MinIdleTime string) XclaimMinIdleTime {
	return XclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type XclaimForceForce Completed

func (c XclaimForceForce) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimForceForce) Build() Completed {
	return Completed(c)
}

type XclaimGroup Completed

func (c XclaimGroup) Consumer(Consumer string) XclaimConsumer {
	return XclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type XclaimId Completed

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

type XclaimIdle Completed

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

type XclaimJustidJustid Completed

func (c XclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XclaimKey Completed

func (c XclaimKey) Group(Group string) XclaimGroup {
	return XclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type XclaimMinIdleTime Completed

func (c XclaimMinIdleTime) Id(Id ...string) XclaimId {
	return XclaimId{cf: c.cf, cs: append(c.cs, Id...)}
}

type XclaimRetrycount Completed

func (c XclaimRetrycount) Force() XclaimForceForce {
	return XclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c XclaimRetrycount) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimRetrycount) Build() Completed {
	return Completed(c)
}

type XclaimTime Completed

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

type Xdel Completed

func (c Xdel) Key(Key string) XdelKey {
	return XdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xdel() (c Xdel) {
	c.cs = append(b.get(), "XDEL")
	return
}

type XdelId Completed

func (c XdelId) Id(Id ...string) XdelId {
	return XdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XdelId) Build() Completed {
	return Completed(c)
}

type XdelKey Completed

func (c XdelKey) Id(Id ...string) XdelId {
	return XdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

type Xgroup Completed

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

func (c Xgroup) Build() Completed {
	return Completed(c)
}

func (b *Builder) Xgroup() (c Xgroup) {
	c.cs = append(b.get(), "XGROUP")
	return
}

type XgroupCreateCreate Completed

func (c XgroupCreateCreate) Id(Id string) XgroupCreateId {
	return XgroupCreateId{cf: c.cf, cs: append(c.cs, Id)}
}

type XgroupCreateId Completed

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

func (c XgroupCreateId) Build() Completed {
	return Completed(c)
}

type XgroupCreateMkstream Completed

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

func (c XgroupCreateMkstream) Build() Completed {
	return Completed(c)
}

type XgroupCreateconsumer Completed

func (c XgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
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
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupDestroy) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupDestroy) Build() Completed {
	return Completed(c)
}

type XgroupSetidId Completed

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

type XgroupSetidSetid Completed

func (c XgroupSetidSetid) Id(Id string) XgroupSetidId {
	return XgroupSetidId{cf: c.cf, cs: append(c.cs, Id)}
}

type Xinfo Completed

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

type XinfoConsumers Completed

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

type XinfoGroups Completed

func (c XinfoGroups) Stream(Key string) XinfoStream {
	return XinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c XinfoGroups) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
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
	return XinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c XinfoStream) Build() Completed {
	return Completed(c)
}

type Xlen Completed

func (c Xlen) Key(Key string) XlenKey {
	return XlenKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xlen() (c Xlen) {
	c.cs = append(b.get(), "XLEN")
	return
}

type XlenKey Completed

func (c XlenKey) Build() Completed {
	return Completed(c)
}

type Xpending Completed

func (c Xpending) Key(Key string) XpendingKey {
	return XpendingKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xpending() (c Xpending) {
	c.cs = append(b.get(), "XPENDING")
	return
}

type XpendingFiltersConsumer Completed

func (c XpendingFiltersConsumer) Build() Completed {
	return Completed(c)
}

type XpendingFiltersCount Completed

func (c XpendingFiltersCount) Consumer(Consumer string) XpendingFiltersConsumer {
	return XpendingFiltersConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

func (c XpendingFiltersCount) Build() Completed {
	return Completed(c)
}

type XpendingFiltersEnd Completed

func (c XpendingFiltersEnd) Count(Count int64) XpendingFiltersCount {
	return XpendingFiltersCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type XpendingFiltersIdle Completed

func (c XpendingFiltersIdle) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XpendingFiltersStart Completed

func (c XpendingFiltersStart) End(End string) XpendingFiltersEnd {
	return XpendingFiltersEnd{cf: c.cf, cs: append(c.cs, End)}
}

type XpendingGroup Completed

func (c XpendingGroup) Idle(MinIdleTime int64) XpendingFiltersIdle {
	return XpendingFiltersIdle{cf: c.cf, cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10))}
}

func (c XpendingGroup) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XpendingKey Completed

func (c XpendingKey) Group(Group string) XpendingGroup {
	return XpendingGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type Xrange Completed

func (c Xrange) Key(Key string) XrangeKey {
	return XrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xrange() (c Xrange) {
	c.cs = append(b.get(), "XRANGE")
	return
}

type XrangeCount Completed

func (c XrangeCount) Build() Completed {
	return Completed(c)
}

type XrangeEnd Completed

func (c XrangeEnd) Count(Count int64) XrangeCount {
	return XrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrangeEnd) Build() Completed {
	return Completed(c)
}

type XrangeKey Completed

func (c XrangeKey) Start(Start string) XrangeStart {
	return XrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XrangeStart Completed

func (c XrangeStart) End(End string) XrangeEnd {
	return XrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type Xread Completed

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

type XreadBlock Completed

func (c XreadBlock) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadCount Completed

func (c XreadCount) Block(Milliseconds int64) XreadBlock {
	c.cf = blockTag
	return XreadBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadCount) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadId Completed

func (c XreadId) Id(Id ...string) XreadId {
	return XreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadId) Build() Completed {
	return Completed(c)
}

type XreadKey Completed

func (c XreadKey) Id(Id ...string) XreadId {
	return XreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadKey) Key(Key ...string) XreadKey {
	return XreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type XreadStreamsStreams Completed

func (c XreadStreamsStreams) Key(Key ...string) XreadKey {
	return XreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Xreadgroup Completed

func (c Xreadgroup) Group(Group string, Consumer string) XreadgroupGroup {
	return XreadgroupGroup{cf: c.cf, cs: append(c.cs, "GROUP", Group, Consumer)}
}

func (b *Builder) Xreadgroup() (c Xreadgroup) {
	c.cs = append(b.get(), "XREADGROUP")
	return
}

type XreadgroupBlock Completed

func (c XreadgroupBlock) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c XreadgroupBlock) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupCount Completed

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

type XreadgroupGroup Completed

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

type XreadgroupId Completed

func (c XreadgroupId) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadgroupId) Build() Completed {
	return Completed(c)
}

type XreadgroupKey Completed

func (c XreadgroupKey) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadgroupKey) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type XreadgroupNoackNoack Completed

func (c XreadgroupNoackNoack) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupStreamsStreams Completed

func (c XreadgroupStreamsStreams) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Xrevrange Completed

func (c Xrevrange) Key(Key string) XrevrangeKey {
	return XrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xrevrange() (c Xrevrange) {
	c.cs = append(b.get(), "XREVRANGE")
	return
}

type XrevrangeCount Completed

func (c XrevrangeCount) Build() Completed {
	return Completed(c)
}

type XrevrangeEnd Completed

func (c XrevrangeEnd) Start(Start string) XrevrangeStart {
	return XrevrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XrevrangeKey Completed

func (c XrevrangeKey) End(End string) XrevrangeEnd {
	return XrevrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type XrevrangeStart Completed

func (c XrevrangeStart) Count(Count int64) XrevrangeCount {
	return XrevrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrevrangeStart) Build() Completed {
	return Completed(c)
}

type Xtrim Completed

func (c Xtrim) Key(Key string) XtrimKey {
	return XtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Xtrim() (c Xtrim) {
	c.cs = append(b.get(), "XTRIM")
	return
}

type XtrimKey Completed

func (c XtrimKey) Maxlen() XtrimTrimStrategyMaxlen {
	return XtrimTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c XtrimKey) Minid() XtrimTrimStrategyMinid {
	return XtrimTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

type XtrimTrimLimit Completed

func (c XtrimTrimLimit) Build() Completed {
	return Completed(c)
}

type XtrimTrimOperatorAlmost Completed

func (c XtrimTrimOperatorAlmost) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimOperatorExact Completed

func (c XtrimTrimOperatorExact) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimStrategyMaxlen Completed

func (c XtrimTrimStrategyMaxlen) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XtrimTrimStrategyMaxlen) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XtrimTrimStrategyMaxlen) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimStrategyMinid Completed

func (c XtrimTrimStrategyMinid) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c XtrimTrimStrategyMinid) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c XtrimTrimStrategyMinid) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimThreshold Completed

func (c XtrimTrimThreshold) Limit(Count int64) XtrimTrimLimit {
	return XtrimTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XtrimTrimThreshold) Build() Completed {
	return Completed(c)
}

type Zadd Completed

func (c Zadd) Key(Key string) ZaddKey {
	return ZaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zadd() (c Zadd) {
	c.cs = append(b.get(), "ZADD")
	return
}

type ZaddChangeCh Completed

func (c ZaddChangeCh) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddChangeCh) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: c.cs}
}

type ZaddComparisonGt Completed

func (c ZaddComparisonGt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c ZaddComparisonGt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddComparisonGt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: c.cs}
}

type ZaddComparisonLt Completed

func (c ZaddComparisonLt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c ZaddComparisonLt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddComparisonLt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: c.cs}
}

type ZaddConditionNx Completed

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
	return ZaddScoreMember{cf: c.cf, cs: c.cs}
}

type ZaddConditionXx Completed

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
	return ZaddScoreMember{cf: c.cf, cs: c.cs}
}

type ZaddIncrementIncr Completed

func (c ZaddIncrementIncr) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: c.cs}
}

type ZaddKey Completed

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
	return ZaddScoreMember{cf: c.cf, cs: c.cs}
}

type ZaddScoreMember Completed

func (c ZaddScoreMember) ScoreMember(Score float64, Member string) ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member)}
}

func (c ZaddScoreMember) Build() Completed {
	return Completed(c)
}

type Zcard Completed

func (c Zcard) Key(Key string) ZcardKey {
	return ZcardKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zcard() (c Zcard) {
	c.cs = append(b.get(), "ZCARD")
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
	return ZcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zcount() (c Zcount) {
	c.cs = append(b.get(), "ZCOUNT")
	return
}

type ZcountKey Completed

func (c ZcountKey) Min(Min float64) ZcountMin {
	return ZcountMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
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
	return ZcountMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type Zdiff Completed

func (c Zdiff) Numkeys(Numkeys int64) ZdiffNumkeys {
	return ZdiffNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zdiff() (c Zdiff) {
	c.cs = append(b.get(), "ZDIFF")
	return
}

type ZdiffKey Completed

func (c ZdiffKey) Withscores() ZdiffWithscoresWithscores {
	return ZdiffWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZdiffKey) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZdiffKey) Build() Completed {
	return Completed(c)
}

type ZdiffNumkeys Completed

func (c ZdiffNumkeys) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZdiffWithscoresWithscores Completed

func (c ZdiffWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zdiffstore Completed

func (c Zdiffstore) Destination(Destination string) ZdiffstoreDestination {
	return ZdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *Builder) Zdiffstore() (c Zdiffstore) {
	c.cs = append(b.get(), "ZDIFFSTORE")
	return
}

type ZdiffstoreDestination Completed

func (c ZdiffstoreDestination) Numkeys(Numkeys int64) ZdiffstoreNumkeys {
	return ZdiffstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZdiffstoreKey Completed

func (c ZdiffstoreKey) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZdiffstoreKey) Build() Completed {
	return Completed(c)
}

type ZdiffstoreNumkeys Completed

func (c ZdiffstoreNumkeys) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Zincrby Completed

func (c Zincrby) Key(Key string) ZincrbyKey {
	return ZincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zincrby() (c Zincrby) {
	c.cs = append(b.get(), "ZINCRBY")
	return
}

type ZincrbyIncrement Completed

func (c ZincrbyIncrement) Member(Member string) ZincrbyMember {
	return ZincrbyMember{cf: c.cf, cs: append(c.cs, Member)}
}

type ZincrbyKey Completed

func (c ZincrbyKey) Increment(Increment int64) ZincrbyIncrement {
	return ZincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type ZincrbyMember Completed

func (c ZincrbyMember) Build() Completed {
	return Completed(c)
}

type Zinter Completed

func (c Zinter) Numkeys(Numkeys int64) ZinterNumkeys {
	return ZinterNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zinter() (c Zinter) {
	c.cs = append(b.get(), "ZINTER")
	return
}

type ZinterAggregateMax Completed

func (c ZinterAggregateMax) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMax) Build() Completed {
	return Completed(c)
}

type ZinterAggregateMin Completed

func (c ZinterAggregateMin) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMin) Build() Completed {
	return Completed(c)
}

type ZinterAggregateSum Completed

func (c ZinterAggregateSum) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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

type ZinterNumkeys Completed

func (c ZinterNumkeys) Key(Key ...string) ZinterKey {
	return ZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZinterWeights Completed

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

type ZinterWithscoresWithscores Completed

func (c ZinterWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zintercard Completed

func (c Zintercard) Numkeys(Numkeys int64) ZintercardNumkeys {
	return ZintercardNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zintercard() (c Zintercard) {
	c.cs = append(b.get(), "ZINTERCARD")
	return
}

type ZintercardKey Completed

func (c ZintercardKey) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZintercardKey) Build() Completed {
	return Completed(c)
}

type ZintercardNumkeys Completed

func (c ZintercardNumkeys) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Zinterstore Completed

func (c Zinterstore) Destination(Destination string) ZinterstoreDestination {
	return ZinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return ZinterstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZinterstoreKey Completed

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

type ZinterstoreNumkeys Completed

func (c ZinterstoreNumkeys) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZinterstoreWeights Completed

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

type Zlexcount Completed

func (c Zlexcount) Key(Key string) ZlexcountKey {
	return ZlexcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zlexcount() (c Zlexcount) {
	c.cs = append(b.get(), "ZLEXCOUNT")
	return
}

type ZlexcountKey Completed

func (c ZlexcountKey) Min(Min string) ZlexcountMin {
	return ZlexcountMin{cf: c.cf, cs: append(c.cs, Min)}
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
	return ZlexcountMax{cf: c.cf, cs: append(c.cs, Max)}
}

type Zmscore Completed

func (c Zmscore) Key(Key string) ZmscoreKey {
	return ZmscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zmscore() (c Zmscore) {
	c.cs = append(b.get(), "ZMSCORE")
	return
}

type ZmscoreKey Completed

func (c ZmscoreKey) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type ZmscoreMember Completed

func (c ZmscoreMember) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c ZmscoreMember) Build() Completed {
	return Completed(c)
}

func (c ZmscoreMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zpopmax Completed

func (c Zpopmax) Key(Key string) ZpopmaxKey {
	return ZpopmaxKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return ZpopmaxCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopmaxKey) Build() Completed {
	return Completed(c)
}

type Zpopmin Completed

func (c Zpopmin) Key(Key string) ZpopminKey {
	return ZpopminKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return ZpopminCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopminKey) Build() Completed {
	return Completed(c)
}

type Zrandmember Completed

func (c Zrandmember) Key(Key string) ZrandmemberKey {
	return ZrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrandmember() (c Zrandmember) {
	c.cs = append(b.get(), "ZRANDMEMBER")
	return
}

type ZrandmemberKey Completed

func (c ZrandmemberKey) Count(Count int64) ZrandmemberOptionsCount {
	return ZrandmemberOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZrandmemberKey) Build() Completed {
	return Completed(c)
}

type ZrandmemberOptionsCount Completed

func (c ZrandmemberOptionsCount) Withscores() ZrandmemberOptionsWithscoresWithscores {
	return ZrandmemberOptionsWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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
	return ZrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrange() (c Zrange) {
	c.cs = append(b.get(), "ZRANGE")
	return
}

type ZrangeKey Completed

func (c ZrangeKey) Min(Min string) ZrangeMin {
	return ZrangeMin{cf: c.cf, cs: append(c.cs, Min)}
}

type ZrangeLimit Completed

func (c ZrangeLimit) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangeLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeMax Completed

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

type ZrangeMin Completed

func (c ZrangeMin) Max(Max string) ZrangeMax {
	return ZrangeMax{cf: c.cf, cs: append(c.cs, Max)}
}

type ZrangeRevRev Completed

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

type ZrangeSortbyBylex Completed

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

type ZrangeSortbyByscore Completed

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

type ZrangeWithscoresWithscores Completed

func (c ZrangeWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrangeWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangebylex Completed

func (c Zrangebylex) Key(Key string) ZrangebylexKey {
	return ZrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrangebylex() (c Zrangebylex) {
	c.cs = append(b.get(), "ZRANGEBYLEX")
	return
}

type ZrangebylexKey Completed

func (c ZrangebylexKey) Min(Min string) ZrangebylexMin {
	return ZrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
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
	return ZrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebylexMax) Build() Completed {
	return Completed(c)
}

func (c ZrangebylexMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexMin Completed

func (c ZrangebylexMin) Max(Max string) ZrangebylexMax {
	return ZrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

type Zrangebyscore Completed

func (c Zrangebyscore) Key(Key string) ZrangebyscoreKey {
	return ZrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrangebyscore() (c Zrangebyscore) {
	c.cs = append(b.get(), "ZRANGEBYSCORE")
	return
}

type ZrangebyscoreKey Completed

func (c ZrangebyscoreKey) Min(Min float64) ZrangebyscoreMin {
	return ZrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
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

type ZrangebyscoreMin Completed

func (c ZrangebyscoreMin) Max(Max float64) ZrangebyscoreMax {
	return ZrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type ZrangebyscoreWithscoresWithscores Completed

func (c ZrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrangebyscoreLimit {
	return ZrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebyscoreWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangestore Completed

func (c Zrangestore) Dst(Dst string) ZrangestoreDst {
	return ZrangestoreDst{cf: c.cf, cs: append(c.cs, Dst)}
}

func (b *Builder) Zrangestore() (c Zrangestore) {
	c.cs = append(b.get(), "ZRANGESTORE")
	return
}

type ZrangestoreDst Completed

func (c ZrangestoreDst) Src(Src string) ZrangestoreSrc {
	return ZrangestoreSrc{cf: c.cf, cs: append(c.cs, Src)}
}

type ZrangestoreLimit Completed

func (c ZrangestoreLimit) Build() Completed {
	return Completed(c)
}

type ZrangestoreMax Completed

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

type ZrangestoreMin Completed

func (c ZrangestoreMin) Max(Max string) ZrangestoreMax {
	return ZrangestoreMax{cf: c.cf, cs: append(c.cs, Max)}
}

type ZrangestoreRevRev Completed

func (c ZrangestoreRevRev) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreRevRev) Build() Completed {
	return Completed(c)
}

type ZrangestoreSortbyBylex Completed

func (c ZrangestoreSortbyBylex) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangestoreSortbyBylex) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreSortbyBylex) Build() Completed {
	return Completed(c)
}

type ZrangestoreSortbyByscore Completed

func (c ZrangestoreSortbyByscore) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c ZrangestoreSortbyByscore) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreSortbyByscore) Build() Completed {
	return Completed(c)
}

type ZrangestoreSrc Completed

func (c ZrangestoreSrc) Min(Min string) ZrangestoreMin {
	return ZrangestoreMin{cf: c.cf, cs: append(c.cs, Min)}
}

type Zrank Completed

func (c Zrank) Key(Key string) ZrankKey {
	return ZrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrank() (c Zrank) {
	c.cs = append(b.get(), "ZRANK")
	return
}

type ZrankKey Completed

func (c ZrankKey) Member(Member string) ZrankMember {
	return ZrankMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return ZremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrem() (c Zrem) {
	c.cs = append(b.get(), "ZREM")
	return
}

type ZremKey Completed

func (c ZremKey) Member(Member ...string) ZremMember {
	return ZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type ZremMember Completed

func (c ZremMember) Member(Member ...string) ZremMember {
	return ZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c ZremMember) Build() Completed {
	return Completed(c)
}

type Zremrangebylex Completed

func (c Zremrangebylex) Key(Key string) ZremrangebylexKey {
	return ZremrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebylex() (c Zremrangebylex) {
	c.cs = append(b.get(), "ZREMRANGEBYLEX")
	return
}

type ZremrangebylexKey Completed

func (c ZremrangebylexKey) Min(Min string) ZremrangebylexMin {
	return ZremrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

type ZremrangebylexMax Completed

func (c ZremrangebylexMax) Build() Completed {
	return Completed(c)
}

type ZremrangebylexMin Completed

func (c ZremrangebylexMin) Max(Max string) ZremrangebylexMax {
	return ZremrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

type Zremrangebyrank Completed

func (c Zremrangebyrank) Key(Key string) ZremrangebyrankKey {
	return ZremrangebyrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebyrank() (c Zremrangebyrank) {
	c.cs = append(b.get(), "ZREMRANGEBYRANK")
	return
}

type ZremrangebyrankKey Completed

func (c ZremrangebyrankKey) Start(Start int64) ZremrangebyrankStart {
	return ZremrangebyrankStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type ZremrangebyrankStart Completed

func (c ZremrangebyrankStart) Stop(Stop int64) ZremrangebyrankStop {
	return ZremrangebyrankStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type ZremrangebyrankStop Completed

func (c ZremrangebyrankStop) Build() Completed {
	return Completed(c)
}

type Zremrangebyscore Completed

func (c Zremrangebyscore) Key(Key string) ZremrangebyscoreKey {
	return ZremrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebyscore() (c Zremrangebyscore) {
	c.cs = append(b.get(), "ZREMRANGEBYSCORE")
	return
}

type ZremrangebyscoreKey Completed

func (c ZremrangebyscoreKey) Min(Min float64) ZremrangebyscoreMin {
	return ZremrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZremrangebyscoreMax Completed

func (c ZremrangebyscoreMax) Build() Completed {
	return Completed(c)
}

type ZremrangebyscoreMin Completed

func (c ZremrangebyscoreMin) Max(Max float64) ZremrangebyscoreMax {
	return ZremrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type Zrevrange Completed

func (c Zrevrange) Key(Key string) ZrevrangeKey {
	return ZrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrange() (c Zrevrange) {
	c.cs = append(b.get(), "ZREVRANGE")
	return
}

type ZrevrangeKey Completed

func (c ZrevrangeKey) Start(Start int64) ZrevrangeStart {
	return ZrevrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type ZrevrangeStart Completed

func (c ZrevrangeStart) Stop(Stop int64) ZrevrangeStop {
	return ZrevrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type ZrevrangeStop Completed

func (c ZrevrangeStop) Withscores() ZrevrangeWithscoresWithscores {
	return ZrevrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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
	return ZrevrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrangebylex() (c Zrevrangebylex) {
	c.cs = append(b.get(), "ZREVRANGEBYLEX")
	return
}

type ZrevrangebylexKey Completed

func (c ZrevrangebylexKey) Max(Max string) ZrevrangebylexMax {
	return ZrevrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
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
	return ZrevrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

type ZrevrangebylexMin Completed

func (c ZrevrangebylexMin) Limit(Offset int64, Count int64) ZrevrangebylexLimit {
	return ZrevrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebylexMin) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebylexMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrangebyscore Completed

func (c Zrevrangebyscore) Key(Key string) ZrevrangebyscoreKey {
	return ZrevrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrangebyscore() (c Zrevrangebyscore) {
	c.cs = append(b.get(), "ZREVRANGEBYSCORE")
	return
}

type ZrevrangebyscoreKey Completed

func (c ZrevrangebyscoreKey) Max(Max float64) ZrevrangebyscoreMax {
	return ZrevrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
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
	return ZrevrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZrevrangebyscoreMin Completed

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

type ZrevrangebyscoreWithscoresWithscores Completed

func (c ZrevrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrevrangebyscoreLimit {
	return ZrevrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebyscoreWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrank Completed

func (c Zrevrank) Key(Key string) ZrevrankKey {
	return ZrevrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrank() (c Zrevrank) {
	c.cs = append(b.get(), "ZREVRANK")
	return
}

type ZrevrankKey Completed

func (c ZrevrankKey) Member(Member string) ZrevrankMember {
	return ZrevrankMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return ZscanKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zscan() (c Zscan) {
	c.cs = append(b.get(), "ZSCAN")
	return
}

type ZscanCount Completed

func (c ZscanCount) Build() Completed {
	return Completed(c)
}

type ZscanCursor Completed

func (c ZscanCursor) Match(Pattern string) ZscanMatch {
	return ZscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c ZscanCursor) Count(Count int64) ZscanCount {
	return ZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ZscanCursor) Build() Completed {
	return Completed(c)
}

type ZscanKey Completed

func (c ZscanKey) Cursor(Cursor int64) ZscanCursor {
	return ZscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type ZscanMatch Completed

func (c ZscanMatch) Count(Count int64) ZscanCount {
	return ZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ZscanMatch) Build() Completed {
	return Completed(c)
}

type Zscore Completed

func (c Zscore) Key(Key string) ZscoreKey {
	return ZscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *Builder) Zscore() (c Zscore) {
	c.cs = append(b.get(), "ZSCORE")
	return
}

type ZscoreKey Completed

func (c ZscoreKey) Member(Member string) ZscoreMember {
	return ZscoreMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return ZunionNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zunion() (c Zunion) {
	c.cs = append(b.get(), "ZUNION")
	return
}

type ZunionAggregateMax Completed

func (c ZunionAggregateMax) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMax) Build() Completed {
	return Completed(c)
}

type ZunionAggregateMin Completed

func (c ZunionAggregateMin) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMin) Build() Completed {
	return Completed(c)
}

type ZunionAggregateSum Completed

func (c ZunionAggregateSum) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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

type ZunionNumkeys Completed

func (c ZunionNumkeys) Key(Key ...string) ZunionKey {
	return ZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZunionWeights Completed

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

type ZunionWithscoresWithscores Completed

func (c ZunionWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zunionstore Completed

func (c Zunionstore) Destination(Destination string) ZunionstoreDestination {
	return ZunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return ZunionstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZunionstoreKey Completed

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

type ZunionstoreNumkeys Completed

func (c ZunionstoreNumkeys) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZunionstoreWeights Completed

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

type SAclCat SCompleted

func (c SAclCat) Categoryname(Categoryname string) SAclCatCategoryname {
	return SAclCatCategoryname{cf: c.cf, cs: append(c.cs, Categoryname)}
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
	return SAclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (b *SBuilder) AclDeluser() (c SAclDeluser) {
	c.cs = append(b.get(), "ACL", "DELUSER")
	c.ks = InitSlot
	return
}

type SAclDeluserUsername SCompleted

func (c SAclDeluserUsername) Username(Username ...string) SAclDeluserUsername {
	return SAclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (c SAclDeluserUsername) Build() SCompleted {
	return SCompleted(c)
}

type SAclGenpass SCompleted

func (c SAclGenpass) Bits(Bits int64) SAclGenpassBits {
	return SAclGenpassBits{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bits, 10))}
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
	return SAclGetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
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
	return SAclLogCountOrReset{cf: c.cf, cs: append(c.cs, CountOrReset)}
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
	return SAclSetuserUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (b *SBuilder) AclSetuser() (c SAclSetuser) {
	c.cs = append(b.get(), "ACL", "SETUSER")
	c.ks = InitSlot
	return
}

type SAclSetuserRule SCompleted

func (c SAclSetuserRule) Rule(Rule ...string) SAclSetuserRule {
	return SAclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c SAclSetuserRule) Build() SCompleted {
	return SCompleted(c)
}

type SAclSetuserUsername SCompleted

func (c SAclSetuserUsername) Rule(Rule ...string) SAclSetuserRule {
	return SAclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
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
	return SAppendKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Append() (c SAppend) {
	c.cs = append(b.get(), "APPEND")
	c.ks = InitSlot
	return
}

type SAppendKey SCompleted

func (c SAppendKey) Value(Value string) SAppendValue {
	return SAppendValue{cf: c.cf, cs: append(c.cs, Value)}
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
	return SAuthUsername{cf: c.cf, cs: append(c.cs, Username)}
}

func (c SAuth) Password(Password string) SAuthPassword {
	return SAuthPassword{cf: c.cf, cs: append(c.cs, Password)}
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
	return SAuthPassword{cf: c.cf, cs: append(c.cs, Password)}
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
	return SBgsaveScheduleSchedule{cf: c.cf, cs: append(c.cs, "SCHEDULE")}
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
	return SBitcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Bitcount() (c SBitcount) {
	c.cs = append(b.get(), "BITCOUNT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SBitcountKey SCompleted

func (c SBitcountKey) StartEnd(Start int64, End int64) SBitcountStartEnd {
	return SBitcountStartEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10), strconv.FormatInt(End, 10))}
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
	return SBitfieldKey{cf: c.cf, cs: append(c.cs, Key)}
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

type SBitfieldIncrby SCompleted

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

type SBitfieldKey SCompleted

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

type SBitfieldRo SCompleted

func (c SBitfieldRo) Key(Key string) SBitfieldRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SBitfieldRoKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SBitfieldRoGet{cf: c.cf, cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

type SBitfieldSat SCompleted

func (c SBitfieldSat) Build() SCompleted {
	return SCompleted(c)
}

type SBitfieldSet SCompleted

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

type SBitfieldWrap SCompleted

func (c SBitfieldWrap) Build() SCompleted {
	return SCompleted(c)
}

type SBitop SCompleted

func (c SBitop) Operation(Operation string) SBitopOperation {
	return SBitopOperation{cf: c.cf, cs: append(c.cs, Operation)}
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
	return SBitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBitopKey SCompleted

func (c SBitopKey) Key(Key ...string) SBitopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SBitopKey) Build() SCompleted {
	return SCompleted(c)
}

type SBitopOperation SCompleted

func (c SBitopOperation) Destkey(Destkey string) SBitopDestkey {
	c.ks = checkSlot(c.ks, slot(Destkey))
	return SBitopDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

type SBitpos SCompleted

func (c SBitpos) Key(Key string) SBitposKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SBitposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Bitpos() (c SBitpos) {
	c.cs = append(b.get(), "BITPOS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SBitposBit SCompleted

func (c SBitposBit) Start(Start int64) SBitposIndexStart {
	return SBitposIndexStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
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
	return SBitposIndexEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c SBitposIndexStart) Build() SCompleted {
	return SCompleted(c)
}

func (c SBitposIndexStart) Cache() SCacheable {
	return SCacheable(c)
}

type SBitposKey SCompleted

func (c SBitposKey) Bit(Bit int64) SBitposBit {
	return SBitposBit{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bit, 10))}
}

type SBlmove SCompleted

func (c SBlmove) Source(Source string) SBlmoveSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SBlmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Blmove() (c SBlmove) {
	c.cs = append(b.get(), "BLMOVE")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBlmoveDestination SCompleted

func (c SBlmoveDestination) Left() SBlmoveWherefromLeft {
	return SBlmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmoveDestination) Right() SBlmoveWherefromRight {
	return SBlmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmoveSource SCompleted

func (c SBlmoveSource) Destination(Destination string) SBlmoveDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SBlmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SBlmoveTimeout SCompleted

func (c SBlmoveTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBlmoveWherefromLeft SCompleted

func (c SBlmoveWherefromLeft) Left() SBlmoveWheretoLeft {
	return SBlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmoveWherefromLeft) Right() SBlmoveWheretoRight {
	return SBlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmoveWherefromRight SCompleted

func (c SBlmoveWherefromRight) Left() SBlmoveWheretoLeft {
	return SBlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmoveWherefromRight) Right() SBlmoveWheretoRight {
	return SBlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmoveWheretoLeft SCompleted

func (c SBlmoveWheretoLeft) Timeout(Timeout float64) SBlmoveTimeout {
	return SBlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type SBlmoveWheretoRight SCompleted

func (c SBlmoveWheretoRight) Timeout(Timeout float64) SBlmoveTimeout {
	return SBlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type SBlmpop SCompleted

func (c SBlmpop) Timeout(Timeout float64) SBlmpopTimeout {
	return SBlmpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
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
	return SBlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmpopKey) Right() SBlmpopWhereRight {
	return SBlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c SBlmpopKey) Key(Key ...string) SBlmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBlmpopNumkeys SCompleted

func (c SBlmpopNumkeys) Key(Key ...string) SBlmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SBlmpopNumkeys) Left() SBlmpopWhereLeft {
	return SBlmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SBlmpopNumkeys) Right() SBlmpopWhereRight {
	return SBlmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SBlmpopTimeout SCompleted

func (c SBlmpopTimeout) Numkeys(Numkeys int64) SBlmpopNumkeys {
	return SBlmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SBlmpopWhereLeft SCompleted

func (c SBlmpopWhereLeft) Count(Count int64) SBlmpopCount {
	return SBlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SBlmpopWhereLeft) Build() SCompleted {
	return SCompleted(c)
}

type SBlmpopWhereRight SCompleted

func (c SBlmpopWhereRight) Count(Count int64) SBlmpopCount {
	return SBlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SBlmpopWhereRight) Build() SCompleted {
	return SCompleted(c)
}

type SBlpop SCompleted

func (c SBlpop) Key(Key ...string) SBlpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Blpop() (c SBlpop) {
	c.cs = append(b.get(), "BLPOP")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBlpopKey SCompleted

func (c SBlpopKey) Timeout(Timeout float64) SBlpopTimeout {
	return SBlpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBlpopKey) Key(Key ...string) SBlpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SBrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Brpop() (c SBrpop) {
	c.cs = append(b.get(), "BRPOP")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBrpopKey SCompleted

func (c SBrpopKey) Timeout(Timeout float64) SBrpopTimeout {
	return SBrpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBrpopKey) Key(Key ...string) SBrpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBrpopTimeout SCompleted

func (c SBrpopTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SBrpoplpush SCompleted

func (c SBrpoplpush) Source(Source string) SBrpoplpushSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SBrpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Brpoplpush() (c SBrpoplpush) {
	c.cs = append(b.get(), "BRPOPLPUSH")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBrpoplpushDestination SCompleted

func (c SBrpoplpushDestination) Timeout(Timeout float64) SBrpoplpushTimeout {
	return SBrpoplpushTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type SBrpoplpushSource SCompleted

func (c SBrpoplpushSource) Destination(Destination string) SBrpoplpushDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SBrpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return SBzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Bzpopmax() (c SBzpopmax) {
	c.cs = append(b.get(), "BZPOPMAX")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBzpopmaxKey SCompleted

func (c SBzpopmaxKey) Timeout(Timeout float64) SBzpopmaxTimeout {
	return SBzpopmaxTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBzpopmaxKey) Key(Key ...string) SBzpopmaxKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SBzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *SBuilder) Bzpopmin() (c SBzpopmin) {
	c.cs = append(b.get(), "BZPOPMIN")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SBzpopminKey SCompleted

func (c SBzpopminKey) Timeout(Timeout float64) SBzpopminTimeout {
	return SBzpopminTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c SBzpopminKey) Key(Key ...string) SBzpopminKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SBzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SBzpopminTimeout SCompleted

func (c SBzpopminTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientCaching SCompleted

func (c SClientCaching) Yes() SClientCachingModeYes {
	return SClientCachingModeYes{cf: c.cf, cs: append(c.cs, "YES")}
}

func (c SClientCaching) No() SClientCachingModeNo {
	return SClientCachingModeNo{cf: c.cf, cs: append(c.cs, "NO")}
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
	c.ks = InitSlot
	return
}

type SClientKillAddr SCompleted

func (c SClientKillAddr) Laddr(IpPort string) SClientKillLaddr {
	return SClientKillLaddr{cf: c.cf, cs: append(c.cs, "LADDR", IpPort)}
}

func (c SClientKillAddr) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillAddr) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillId SCompleted

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

type SClientKillIpPort SCompleted

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

type SClientKillLaddr SCompleted

func (c SClientKillLaddr) Skipme(YesNo string) SClientKillSkipme {
	return SClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c SClientKillLaddr) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillMaster SCompleted

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

type SClientKillNormal SCompleted

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

type SClientKillPubsub SCompleted

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

type SClientKillSkipme SCompleted

func (c SClientKillSkipme) Build() SCompleted {
	return SCompleted(c)
}

type SClientKillSlave SCompleted

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

type SClientKillUser SCompleted

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

type SClientList SCompleted

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
	return SClientListIdClientId{cf: c.cf, cs: c.cs}
}

func (c SClientListIdClientId) Build() SCompleted {
	return SCompleted(c)
}

type SClientListIdId SCompleted

func (c SClientListIdId) ClientId(ClientId ...int64) SClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return SClientListIdClientId{cf: c.cf, cs: c.cs}
}

type SClientListMaster SCompleted

func (c SClientListMaster) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c SClientListMaster) Build() SCompleted {
	return SCompleted(c)
}

type SClientListNormal SCompleted

func (c SClientListNormal) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c SClientListNormal) Build() SCompleted {
	return SCompleted(c)
}

type SClientListPubsub SCompleted

func (c SClientListPubsub) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c SClientListPubsub) Build() SCompleted {
	return SCompleted(c)
}

type SClientListReplica SCompleted

func (c SClientListReplica) Id() SClientListIdId {
	return SClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c SClientListReplica) Build() SCompleted {
	return SCompleted(c)
}

type SClientNoEvict SCompleted

func (c SClientNoEvict) On() SClientNoEvictEnabledOn {
	return SClientNoEvictEnabledOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c SClientNoEvict) Off() SClientNoEvictEnabledOff {
	return SClientNoEvictEnabledOff{cf: c.cf, cs: append(c.cs, "OFF")}
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
	return SClientPauseTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
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
	return SClientPauseModeWrite{cf: c.cf, cs: append(c.cs, "WRITE")}
}

func (c SClientPauseTimeout) All() SClientPauseModeAll {
	return SClientPauseModeAll{cf: c.cf, cs: append(c.cs, "ALL")}
}

func (c SClientPauseTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SClientReply SCompleted

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
	return SClientSetnameConnectionName{cf: c.cf, cs: append(c.cs, ConnectionName)}
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
	return SClientTrackingStatusOn{cf: c.cf, cs: append(c.cs, "ON")}
}

func (c SClientTracking) Off() SClientTrackingStatusOff {
	return SClientTrackingStatusOff{cf: c.cf, cs: append(c.cs, "OFF")}
}

func (b *SBuilder) ClientTracking() (c SClientTracking) {
	c.cs = append(b.get(), "CLIENT", "TRACKING")
	c.ks = InitSlot
	return
}

type SClientTrackingBcastBcast SCompleted

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

type SClientTrackingNoloopNoloop SCompleted

func (c SClientTrackingNoloopNoloop) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingOptinOptin SCompleted

func (c SClientTrackingOptinOptin) Optout() SClientTrackingOptoutOptout {
	return SClientTrackingOptoutOptout{cf: c.cf, cs: append(c.cs, "OPTOUT")}
}

func (c SClientTrackingOptinOptin) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingOptinOptin) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingOptoutOptout SCompleted

func (c SClientTrackingOptoutOptout) Noloop() SClientTrackingNoloopNoloop {
	return SClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c SClientTrackingOptoutOptout) Build() SCompleted {
	return SCompleted(c)
}

type SClientTrackingPrefix SCompleted

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

type SClientTrackingRedirect SCompleted

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

type SClientTrackingStatusOff SCompleted

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

type SClientTrackingStatusOn SCompleted

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
	return SClientUnblockClientId{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ClientId, 10))}
}

func (b *SBuilder) ClientUnblock() (c SClientUnblock) {
	c.cs = append(b.get(), "CLIENT", "UNBLOCK")
	c.ks = InitSlot
	return
}

type SClientUnblockClientId SCompleted

func (c SClientUnblockClientId) Timeout() SClientUnblockUnblockTypeTimeout {
	return SClientUnblockUnblockTypeTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT")}
}

func (c SClientUnblockClientId) Error() SClientUnblockUnblockTypeError {
	return SClientUnblockUnblockTypeError{cf: c.cf, cs: append(c.cs, "ERROR")}
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
	return SClusterAddslotsSlot{cf: c.cf, cs: c.cs}
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
	return SClusterAddslotsSlot{cf: c.cf, cs: c.cs}
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
	return SClusterCountFailureReportsNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return SClusterCountkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
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
	return SClusterDelslotsSlot{cf: c.cf, cs: c.cs}
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
	return SClusterDelslotsSlot{cf: c.cf, cs: c.cs}
}

func (c SClusterDelslotsSlot) Build() SCompleted {
	return SCompleted(c)
}

type SClusterFailover SCompleted

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
	return SClusterForgetNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return SClusterGetkeysinslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
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
	return SClusterGetkeysinslotCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
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
	return SClusterKeyslotKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SClusterMeetIp{cf: c.cf, cs: append(c.cs, Ip)}
}

func (b *SBuilder) ClusterMeet() (c SClusterMeet) {
	c.cs = append(b.get(), "CLUSTER", "MEET")
	c.ks = InitSlot
	return
}

type SClusterMeetIp SCompleted

func (c SClusterMeetIp) Port(Port int64) SClusterMeetPort {
	return SClusterMeetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
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
	return SClusterReplicasNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return SClusterReplicateNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return SClusterSetConfigEpochConfigEpoch{cf: c.cf, cs: append(c.cs, strconv.FormatInt(ConfigEpoch, 10))}
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
	return SClusterSetslotSlot{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Slot, 10))}
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

type SClusterSetslotSubcommandImporting SCompleted

func (c SClusterSetslotSubcommandImporting) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandImporting) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandMigrating SCompleted

func (c SClusterSetslotSubcommandMigrating) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandMigrating) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandNode SCompleted

func (c SClusterSetslotSubcommandNode) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandNode) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSetslotSubcommandStable SCompleted

func (c SClusterSetslotSubcommandStable) NodeId(NodeId string) SClusterSetslotNodeId {
	return SClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c SClusterSetslotSubcommandStable) Build() SCompleted {
	return SCompleted(c)
}

type SClusterSlaves SCompleted

func (c SClusterSlaves) NodeId(NodeId string) SClusterSlavesNodeId {
	return SClusterSlavesNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
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
	return SCommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (b *SBuilder) CommandInfo() (c SCommandInfo) {
	c.cs = append(b.get(), "COMMAND", "INFO")
	c.ks = InitSlot
	return
}

type SCommandInfoCommandName SCompleted

func (c SCommandInfoCommandName) CommandName(CommandName ...string) SCommandInfoCommandName {
	return SCommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (c SCommandInfoCommandName) Build() SCompleted {
	return SCompleted(c)
}

type SConfigGet SCompleted

func (c SConfigGet) Parameter(Parameter string) SConfigGetParameter {
	return SConfigGetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
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
	return SConfigSetParameter{cf: c.cf, cs: append(c.cs, Parameter)}
}

func (b *SBuilder) ConfigSet() (c SConfigSet) {
	c.cs = append(b.get(), "CONFIG", "SET")
	c.ks = InitSlot
	return
}

type SConfigSetParameter SCompleted

func (c SConfigSetParameter) Value(Value string) SConfigSetValue {
	return SConfigSetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SConfigSetValue SCompleted

func (c SConfigSetValue) Build() SCompleted {
	return SCompleted(c)
}

type SCopy SCompleted

func (c SCopy) Source(Source string) SCopySource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SCopySource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Copy() (c SCopy) {
	c.cs = append(b.get(), "COPY")
	c.ks = InitSlot
	return
}

type SCopyDb SCompleted

func (c SCopyDb) Replace() SCopyReplaceReplace {
	return SCopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c SCopyDb) Build() SCompleted {
	return SCompleted(c)
}

type SCopyDestination SCompleted

func (c SCopyDestination) Db(DestinationDb int64) SCopyDb {
	return SCopyDb{cf: c.cf, cs: append(c.cs, "DB", strconv.FormatInt(DestinationDb, 10))}
}

func (c SCopyDestination) Replace() SCopyReplaceReplace {
	return SCopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
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
	return SCopyDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return SDebugObjectKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SDecrKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SDecrbyKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SDecrbyDecrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Decrement, 10))}
}

type SDel SCompleted

func (c SDel) Key(Key ...string) SDelKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SDelKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SDelKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SDumpKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SEchoMessage{cf: c.cf, cs: append(c.cs, Message)}
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
	return SEvalScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *SBuilder) Eval() (c SEval) {
	c.cs = append(b.get(), "EVAL")
	c.ks = InitSlot
	return
}

type SEvalArg SCompleted

func (c SEvalArg) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalKey SCompleted

func (c SEvalKey) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalKey) Key(Key ...string) SEvalKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalKey) Build() SCompleted {
	return SCompleted(c)
}

type SEvalNumkeys SCompleted

func (c SEvalNumkeys) Key(Key ...string) SEvalKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalNumkeys) Arg(Arg ...string) SEvalArg {
	return SEvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalNumkeys) Build() SCompleted {
	return SCompleted(c)
}

type SEvalRo SCompleted

func (c SEvalRo) Script(Script string) SEvalRoScript {
	return SEvalRoScript{cf: c.cf, cs: append(c.cs, Script)}
}

func (b *SBuilder) EvalRo() (c SEvalRo) {
	c.cs = append(b.get(), "EVAL_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SEvalRoArg SCompleted

func (c SEvalRoArg) Arg(Arg ...string) SEvalRoArg {
	return SEvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalRoArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalRoKey SCompleted

func (c SEvalRoKey) Arg(Arg ...string) SEvalRoArg {
	return SEvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalRoKey) Key(Key ...string) SEvalRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalRoNumkeys SCompleted

func (c SEvalRoNumkeys) Key(Key ...string) SEvalRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalRoScript SCompleted

func (c SEvalRoScript) Numkeys(Numkeys int64) SEvalRoNumkeys {
	return SEvalRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SEvalScript SCompleted

func (c SEvalScript) Numkeys(Numkeys int64) SEvalNumkeys {
	return SEvalNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SEvalsha SCompleted

func (c SEvalsha) Sha1(Sha1 string) SEvalshaSha1 {
	return SEvalshaSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *SBuilder) Evalsha() (c SEvalsha) {
	c.cs = append(b.get(), "EVALSHA")
	c.ks = InitSlot
	return
}

type SEvalshaArg SCompleted

func (c SEvalshaArg) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaKey SCompleted

func (c SEvalshaKey) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaKey) Key(Key ...string) SEvalshaKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalshaKey) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaNumkeys SCompleted

func (c SEvalshaNumkeys) Key(Key ...string) SEvalshaKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SEvalshaNumkeys) Arg(Arg ...string) SEvalshaArg {
	return SEvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaNumkeys) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaRo SCompleted

func (c SEvalshaRo) Sha1(Sha1 string) SEvalshaRoSha1 {
	return SEvalshaRoSha1{cf: c.cf, cs: append(c.cs, Sha1)}
}

func (b *SBuilder) EvalshaRo() (c SEvalshaRo) {
	c.cs = append(b.get(), "EVALSHA_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SEvalshaRoArg SCompleted

func (c SEvalshaRoArg) Arg(Arg ...string) SEvalshaRoArg {
	return SEvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaRoArg) Build() SCompleted {
	return SCompleted(c)
}

type SEvalshaRoKey SCompleted

func (c SEvalshaRoKey) Arg(Arg ...string) SEvalshaRoArg {
	return SEvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SEvalshaRoKey) Key(Key ...string) SEvalshaRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalshaRoNumkeys SCompleted

func (c SEvalshaRoNumkeys) Key(Key ...string) SEvalshaRoKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SEvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SEvalshaRoSha1 SCompleted

func (c SEvalshaRoSha1) Numkeys(Numkeys int64) SEvalshaRoNumkeys {
	return SEvalshaRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SEvalshaSha1 SCompleted

func (c SEvalshaSha1) Numkeys(Numkeys int64) SEvalshaNumkeys {
	return SEvalshaNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
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
	return SExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SExistsKey) Build() SCompleted {
	return SCompleted(c)
}

type SExpire SCompleted

func (c SExpire) Key(Key string) SExpireKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SExpireKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SExpireSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SExpireSeconds SCompleted

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

type SExpireat SCompleted

func (c SExpireat) Key(Key string) SExpireatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SExpireatKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SExpireatTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timestamp, 10))}
}

type SExpireatTimestamp SCompleted

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

type SExpiretime SCompleted

func (c SExpiretime) Key(Key string) SExpiretimeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SExpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SFailoverTargetTo{cf: c.cf, cs: append(c.cs, "TO")}
}

func (c SFailover) Abort() SFailoverAbort {
	return SFailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c SFailover) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
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
	return SFailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c SFailoverAbort) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetForce SCompleted

func (c SFailoverTargetForce) Abort() SFailoverAbort {
	return SFailoverAbort{cf: c.cf, cs: append(c.cs, "ABORT")}
}

func (c SFailoverTargetForce) Timeout(Milliseconds int64) SFailoverTimeout {
	return SFailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c SFailoverTargetForce) Build() SCompleted {
	return SCompleted(c)
}

type SFailoverTargetHost SCompleted

func (c SFailoverTargetHost) Port(Port int64) SFailoverTargetPort {
	return SFailoverTargetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type SFailoverTargetPort SCompleted

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

type SFailoverTargetTo SCompleted

func (c SFailoverTargetTo) Host(Host string) SFailoverTargetHost {
	return SFailoverTargetHost{cf: c.cf, cs: append(c.cs, Host)}
}

type SFailoverTimeout SCompleted

func (c SFailoverTimeout) Build() SCompleted {
	return SCompleted(c)
}

type SFlushall SCompleted

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
	return SGeoaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geoadd() (c SGeoadd) {
	c.cs = append(b.get(), "GEOADD")
	c.ks = InitSlot
	return
}

type SGeoaddChangeCh SCompleted

func (c SGeoaddChangeCh) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type SGeoaddConditionNx SCompleted

func (c SGeoaddConditionNx) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SGeoaddConditionNx) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type SGeoaddConditionXx SCompleted

func (c SGeoaddConditionXx) Ch() SGeoaddChangeCh {
	return SGeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SGeoaddConditionXx) LongitudeLatitudeMember() SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type SGeoaddKey SCompleted

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
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: c.cs}
}

type SGeoaddLongitudeLatitudeMember SCompleted

func (c SGeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) SGeoaddLongitudeLatitudeMember {
	return SGeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member)}
}

func (c SGeoaddLongitudeLatitudeMember) Build() SCompleted {
	return SCompleted(c)
}

type SGeodist SCompleted

func (c SGeodist) Key(Key string) SGeodistKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeodistKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geodist() (c SGeodist) {
	c.cs = append(b.get(), "GEODIST")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeodistKey SCompleted

func (c SGeodistKey) Member1(Member1 string) SGeodistMember1 {
	return SGeodistMember1{cf: c.cf, cs: append(c.cs, Member1)}
}

type SGeodistMember1 SCompleted

func (c SGeodistMember1) Member2(Member2 string) SGeodistMember2 {
	return SGeodistMember2{cf: c.cf, cs: append(c.cs, Member2)}
}

type SGeodistMember2 SCompleted

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
	return SGeohashKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geohash() (c SGeohash) {
	c.cs = append(b.get(), "GEOHASH")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeohashKey SCompleted

func (c SGeohashKey) Member(Member ...string) SGeohashMember {
	return SGeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SGeohashMember SCompleted

func (c SGeohashMember) Member(Member ...string) SGeohashMember {
	return SGeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
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
	return SGeoposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geopos() (c SGeopos) {
	c.cs = append(b.get(), "GEOPOS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeoposKey SCompleted

func (c SGeoposKey) Member(Member ...string) SGeoposMember {
	return SGeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SGeoposMember SCompleted

func (c SGeoposMember) Member(Member ...string) SGeoposMember {
	return SGeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
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
	return SGeoradiusKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Georadius() (c SGeoradius) {
	c.cs = append(b.get(), "GEORADIUS")
	c.ks = InitSlot
	return
}

type SGeoradiusCountAnyAny SCompleted

func (c SGeoradiusCountAnyAny) Asc() SGeoradiusOrderAsc {
	return SGeoradiusOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusCountAnyAny) Desc() SGeoradiusOrderDesc {
	return SGeoradiusOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusCountAnyAny) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusCountAnyAny) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusCountCount SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusCountCount) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusKey SCompleted

func (c SGeoradiusKey) Longitude(Longitude float64) SGeoradiusLongitude {
	return SGeoradiusLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type SGeoradiusLatitude SCompleted

func (c SGeoradiusLatitude) Radius(Radius float64) SGeoradiusRadius {
	return SGeoradiusRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type SGeoradiusLongitude SCompleted

func (c SGeoradiusLongitude) Latitude(Latitude float64) SGeoradiusLatitude {
	return SGeoradiusLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type SGeoradiusOrderAsc SCompleted

func (c SGeoradiusOrderAsc) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusOrderAsc) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusOrderDesc SCompleted

func (c SGeoradiusOrderDesc) Store(Key string) SGeoradiusStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusOrderDesc) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusRadius SCompleted

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

type SGeoradiusRo SCompleted

func (c SGeoradiusRo) Key(Key string) SGeoradiusRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) GeoradiusRo() (c SGeoradiusRo) {
	c.cs = append(b.get(), "GEORADIUS_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeoradiusRoCountAnyAny SCompleted

func (c SGeoradiusRoCountAnyAny) Asc() SGeoradiusRoOrderAsc {
	return SGeoradiusRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusRoCountAnyAny) Desc() SGeoradiusRoOrderDesc {
	return SGeoradiusRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusRoCountAnyAny) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoCountCount SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoKey SCompleted

func (c SGeoradiusRoKey) Longitude(Longitude float64) SGeoradiusRoLongitude {
	return SGeoradiusRoLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type SGeoradiusRoLatitude SCompleted

func (c SGeoradiusRoLatitude) Radius(Radius float64) SGeoradiusRoRadius {
	return SGeoradiusRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type SGeoradiusRoLongitude SCompleted

func (c SGeoradiusRoLongitude) Latitude(Latitude float64) SGeoradiusRoLatitude {
	return SGeoradiusRoLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type SGeoradiusRoOrderAsc SCompleted

func (c SGeoradiusRoOrderAsc) Storedist(Key string) SGeoradiusRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoRadius SCompleted

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

type SGeoradiusRoStoredist SCompleted

func (c SGeoradiusRoStoredist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoStoredist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitFt SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitKm SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitM SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoUnitMi SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithcoordWithcoord SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithdistWithdist SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusRoWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusRoWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusRoWithhashWithhash SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitFt) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitKm SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitKm) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitM SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitM) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusUnitMi SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusUnitMi) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusWithcoordWithcoord SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusWithcoordWithcoord) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusWithdistWithdist SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusWithdistWithdist) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusWithhashWithhash SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusWithhashWithhash) Storedist(Key string) SGeoradiusStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymember SCompleted

func (c SGeoradiusbymember) Key(Key string) SGeoradiusbymemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Georadiusbymember() (c SGeoradiusbymember) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER")
	c.ks = InitSlot
	return
}

type SGeoradiusbymemberCountAnyAny SCompleted

func (c SGeoradiusbymemberCountAnyAny) Asc() SGeoradiusbymemberOrderAsc {
	return SGeoradiusbymemberOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberCountAnyAny) Desc() SGeoradiusbymemberOrderDesc {
	return SGeoradiusbymemberOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberCountAnyAny) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberCountAnyAny) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberCountCount SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberCountCount) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberKey SCompleted

func (c SGeoradiusbymemberKey) Member(Member string) SGeoradiusbymemberMember {
	return SGeoradiusbymemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SGeoradiusbymemberMember SCompleted

func (c SGeoradiusbymemberMember) Radius(Radius float64) SGeoradiusbymemberRadius {
	return SGeoradiusbymemberRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type SGeoradiusbymemberOrderAsc SCompleted

func (c SGeoradiusbymemberOrderAsc) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberOrderAsc) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberOrderDesc SCompleted

func (c SGeoradiusbymemberOrderDesc) Store(Key string) SGeoradiusbymemberStore {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberOrderDesc) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberRadius SCompleted

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

type SGeoradiusbymemberRo SCompleted

func (c SGeoradiusbymemberRo) Key(Key string) SGeoradiusbymemberRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) GeoradiusbymemberRo() (c SGeoradiusbymemberRo) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeoradiusbymemberRoCountAnyAny SCompleted

func (c SGeoradiusbymemberRoCountAnyAny) Asc() SGeoradiusbymemberRoOrderAsc {
	return SGeoradiusbymemberRoOrderAsc{cf: c.cf, cs: append(c.cs, "ASC")}
}

func (c SGeoradiusbymemberRoCountAnyAny) Desc() SGeoradiusbymemberRoOrderDesc {
	return SGeoradiusbymemberRoOrderDesc{cf: c.cf, cs: append(c.cs, "DESC")}
}

func (c SGeoradiusbymemberRoCountAnyAny) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoCountAnyAny) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoCountCount SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoCountCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoCountCount) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoKey SCompleted

func (c SGeoradiusbymemberRoKey) Member(Member string) SGeoradiusbymemberRoMember {
	return SGeoradiusbymemberRoMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SGeoradiusbymemberRoMember SCompleted

func (c SGeoradiusbymemberRoMember) Radius(Radius float64) SGeoradiusbymemberRoRadius {
	return SGeoradiusbymemberRoRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type SGeoradiusbymemberRoOrderAsc SCompleted

func (c SGeoradiusbymemberRoOrderAsc) Storedist(Key string) SGeoradiusbymemberRoStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoRadius SCompleted

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

type SGeoradiusbymemberRoStoredist SCompleted

func (c SGeoradiusbymemberRoStoredist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoStoredist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitFt SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitKm SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitM SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoUnitMi SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithcoordWithcoord SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoWithcoordWithcoord) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithdistWithdist SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberRoWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeoradiusbymemberRoWithdistWithdist) Cache() SCacheable {
	return SCacheable(c)
}

type SGeoradiusbymemberRoWithhashWithhash SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberRoStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitFt) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitKm SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitKm) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitM SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitM) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberUnitMi SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberUnitMi) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberWithcoordWithcoord SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberWithcoordWithcoord) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberWithdistWithdist SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberWithdistWithdist) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberWithdistWithdist) Build() SCompleted {
	return SCompleted(c)
}

type SGeoradiusbymemberWithhashWithhash SCompleted

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
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStore{cf: c.cf, cs: append(c.cs, "STORE", Key)}
}

func (c SGeoradiusbymemberWithhashWithhash) Storedist(Key string) SGeoradiusbymemberStoredist {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c SGeoradiusbymemberWithhashWithhash) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearch SCompleted

func (c SGeosearch) Key(Key string) SGeosearchKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGeosearchKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Geosearch() (c SGeosearch) {
	c.cs = append(b.get(), "GEOSEARCH")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGeosearchBoxBybox SCompleted

func (c SGeosearchBoxBybox) Height(Height float64) SGeosearchBoxHeight {
	return SGeosearchBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type SGeosearchBoxHeight SCompleted

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

type SGeosearchBoxUnitFt SCompleted

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

func (c SGeosearchBoxUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitKm SCompleted

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

func (c SGeosearchBoxUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitM SCompleted

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

func (c SGeosearchBoxUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchBoxUnitMi SCompleted

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

func (c SGeosearchBoxUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchBoxUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleByradius SCompleted

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

type SGeosearchCircleUnitFt SCompleted

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

func (c SGeosearchCircleUnitFt) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitFt) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitKm SCompleted

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

func (c SGeosearchCircleUnitKm) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitKm) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitM SCompleted

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

func (c SGeosearchCircleUnitM) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitM) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCircleUnitMi SCompleted

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

func (c SGeosearchCircleUnitMi) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchCircleUnitMi) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchCountAnyAny SCompleted

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

type SGeosearchCountCount SCompleted

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

type SGeosearchFromlonlat SCompleted

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

func (c SGeosearchFromlonlat) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchFromlonlat) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchFrommember SCompleted

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

func (c SGeosearchFrommember) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchFrommember) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchKey SCompleted

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

func (c SGeosearchKey) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchKey) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchOrderAsc SCompleted

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

func (c SGeosearchOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchOrderDesc SCompleted

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

func (c SGeosearchOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

func (c SGeosearchOrderDesc) Cache() SCacheable {
	return SCacheable(c)
}

type SGeosearchWithcoordWithcoord SCompleted

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

type SGeosearchWithdistWithdist SCompleted

func (c SGeosearchWithdistWithdist) Withhash() SGeosearchWithhashWithhash {
	return SGeosearchWithhashWithhash{cf: c.cf, cs: append(c.cs, "WITHHASH")}
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
	return SGeosearchstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Geosearchstore() (c SGeosearchstore) {
	c.cs = append(b.get(), "GEOSEARCHSTORE")
	c.ks = InitSlot
	return
}

type SGeosearchstoreBoxBybox SCompleted

func (c SGeosearchstoreBoxBybox) Height(Height float64) SGeosearchstoreBoxHeight {
	return SGeosearchstoreBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type SGeosearchstoreBoxHeight SCompleted

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

type SGeosearchstoreBoxUnitFt SCompleted

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

func (c SGeosearchstoreBoxUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreBoxUnitKm SCompleted

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

func (c SGeosearchstoreBoxUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreBoxUnitM SCompleted

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

func (c SGeosearchstoreBoxUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreBoxUnitMi SCompleted

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

func (c SGeosearchstoreBoxUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleByradius SCompleted

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

type SGeosearchstoreCircleUnitFt SCompleted

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

func (c SGeosearchstoreCircleUnitFt) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleUnitKm SCompleted

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

func (c SGeosearchstoreCircleUnitKm) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleUnitM SCompleted

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

func (c SGeosearchstoreCircleUnitM) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCircleUnitMi SCompleted

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

func (c SGeosearchstoreCircleUnitMi) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCountAnyAny SCompleted

func (c SGeosearchstoreCountAnyAny) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c SGeosearchstoreCountAnyAny) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreCountCount SCompleted

func (c SGeosearchstoreCountCount) Any() SGeosearchstoreCountAnyAny {
	return SGeosearchstoreCountAnyAny{cf: c.cf, cs: append(c.cs, "ANY")}
}

func (c SGeosearchstoreCountCount) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c SGeosearchstoreCountCount) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreDestination SCompleted

func (c SGeosearchstoreDestination) Source(Source string) SGeosearchstoreSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SGeosearchstoreSource{cf: c.cf, cs: append(c.cs, Source)}
}

type SGeosearchstoreFromlonlat SCompleted

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

func (c SGeosearchstoreFromlonlat) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreFrommember SCompleted

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

func (c SGeosearchstoreFrommember) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreOrderAsc SCompleted

func (c SGeosearchstoreOrderAsc) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreOrderAsc) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c SGeosearchstoreOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreOrderDesc SCompleted

func (c SGeosearchstoreOrderDesc) Count(Count int64) SGeosearchstoreCountCount {
	return SGeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SGeosearchstoreOrderDesc) Storedist() SGeosearchstoreStoredistStoredist {
	return SGeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c SGeosearchstoreOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SGeosearchstoreSource SCompleted

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
	return SGetKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SGetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Getbit() (c SGetbit) {
	c.cs = append(b.get(), "GETBIT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SGetbitKey SCompleted

func (c SGetbitKey) Offset(Offset int64) SGetbitOffset {
	return SGetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
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
	return SGetdelKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SGetexKey{cf: c.cf, cs: append(c.cs, Key)}
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

type SGetrange SCompleted

func (c SGetrange) Key(Key string) SGetrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SGetrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type SGetrangeStart SCompleted

func (c SGetrangeStart) End(End int64) SGetrangeEnd {
	return SGetrangeEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

type SGetset SCompleted

func (c SGetset) Key(Key string) SGetsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SGetsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Getset() (c SGetset) {
	c.cs = append(b.get(), "GETSET")
	c.ks = InitSlot
	return
}

type SGetsetKey SCompleted

func (c SGetsetKey) Value(Value string) SGetsetValue {
	return SGetsetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SGetsetValue SCompleted

func (c SGetsetValue) Build() SCompleted {
	return SCompleted(c)
}

type SHdel SCompleted

func (c SHdel) Key(Key string) SHdelKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hdel() (c SHdel) {
	c.cs = append(b.get(), "HDEL")
	c.ks = InitSlot
	return
}

type SHdelField SCompleted

func (c SHdelField) Field(Field ...string) SHdelField {
	return SHdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c SHdelField) Build() SCompleted {
	return SCompleted(c)
}

type SHdelKey SCompleted

func (c SHdelKey) Field(Field ...string) SHdelField {
	return SHdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

type SHello SCompleted

func (c SHello) Protover(Protover int64) SHelloArgumentsProtover {
	return SHelloArgumentsProtover{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Protover, 10))}
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
	return SHelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
}

func (c SHelloArgumentsAuth) Build() SCompleted {
	return SCompleted(c)
}

type SHelloArgumentsProtover SCompleted

func (c SHelloArgumentsProtover) Auth(Username string, Password string) SHelloArgumentsAuth {
	return SHelloArgumentsAuth{cf: c.cf, cs: append(c.cs, "AUTH", Username, Password)}
}

func (c SHelloArgumentsProtover) Setname(Clientname string) SHelloArgumentsSetname {
	return SHelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
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
	return SHexistsKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SHexistsField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHget SCompleted

func (c SHget) Key(Key string) SHgetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHgetKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SHgetField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHgetall SCompleted

func (c SHgetall) Key(Key string) SHgetallKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHgetallKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SHincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hincrby() (c SHincrby) {
	c.cs = append(b.get(), "HINCRBY")
	c.ks = InitSlot
	return
}

type SHincrbyField SCompleted

func (c SHincrbyField) Increment(Increment int64) SHincrbyIncrement {
	return SHincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type SHincrbyIncrement SCompleted

func (c SHincrbyIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SHincrbyKey SCompleted

func (c SHincrbyKey) Field(Field string) SHincrbyField {
	return SHincrbyField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHincrbyfloat SCompleted

func (c SHincrbyfloat) Key(Key string) SHincrbyfloatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHincrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hincrbyfloat() (c SHincrbyfloat) {
	c.cs = append(b.get(), "HINCRBYFLOAT")
	c.ks = InitSlot
	return
}

type SHincrbyfloatField SCompleted

func (c SHincrbyfloatField) Increment(Increment float64) SHincrbyfloatIncrement {
	return SHincrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type SHincrbyfloatIncrement SCompleted

func (c SHincrbyfloatIncrement) Build() SCompleted {
	return SCompleted(c)
}

type SHincrbyfloatKey SCompleted

func (c SHincrbyfloatKey) Field(Field string) SHincrbyfloatField {
	return SHincrbyfloatField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHkeys SCompleted

func (c SHkeys) Key(Key string) SHkeysKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHkeysKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SHlenKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SHmgetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hmget() (c SHmget) {
	c.cs = append(b.get(), "HMGET")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHmgetField SCompleted

func (c SHmgetField) Field(Field ...string) SHmgetField {
	return SHmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c SHmgetField) Build() SCompleted {
	return SCompleted(c)
}

func (c SHmgetField) Cache() SCacheable {
	return SCacheable(c)
}

type SHmgetKey SCompleted

func (c SHmgetKey) Field(Field ...string) SHmgetField {
	return SHmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

type SHmset SCompleted

func (c SHmset) Key(Key string) SHmsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHmsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hmset() (c SHmset) {
	c.cs = append(b.get(), "HMSET")
	c.ks = InitSlot
	return
}

type SHmsetFieldValue SCompleted

func (c SHmsetFieldValue) FieldValue(Field string, Value string) SHmsetFieldValue {
	return SHmsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c SHmsetFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SHmsetKey SCompleted

func (c SHmsetKey) FieldValue() SHmsetFieldValue {
	return SHmsetFieldValue{cf: c.cf, cs: c.cs}
}

type SHrandfield SCompleted

func (c SHrandfield) Key(Key string) SHrandfieldKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHrandfieldKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hrandfield() (c SHrandfield) {
	c.cs = append(b.get(), "HRANDFIELD")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SHrandfieldKey SCompleted

func (c SHrandfieldKey) Count(Count int64) SHrandfieldOptionsCount {
	return SHrandfieldOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SHrandfieldKey) Build() SCompleted {
	return SCompleted(c)
}

type SHrandfieldOptionsCount SCompleted

func (c SHrandfieldOptionsCount) Withvalues() SHrandfieldOptionsWithvaluesWithvalues {
	return SHrandfieldOptionsWithvaluesWithvalues{cf: c.cf, cs: append(c.cs, "WITHVALUES")}
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
	return SHscanKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SHscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SHscanCursor) Count(Count int64) SHscanCount {
	return SHscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SHscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SHscanKey SCompleted

func (c SHscanKey) Cursor(Cursor int64) SHscanCursor {
	return SHscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SHscanMatch SCompleted

func (c SHscanMatch) Count(Count int64) SHscanCount {
	return SHscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SHscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SHset SCompleted

func (c SHset) Key(Key string) SHsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHsetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hset() (c SHset) {
	c.cs = append(b.get(), "HSET")
	c.ks = InitSlot
	return
}

type SHsetFieldValue SCompleted

func (c SHsetFieldValue) FieldValue(Field string, Value string) SHsetFieldValue {
	return SHsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c SHsetFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SHsetKey SCompleted

func (c SHsetKey) FieldValue() SHsetFieldValue {
	return SHsetFieldValue{cf: c.cf, cs: c.cs}
}

type SHsetnx SCompleted

func (c SHsetnx) Key(Key string) SHsetnxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHsetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Hsetnx() (c SHsetnx) {
	c.cs = append(b.get(), "HSETNX")
	c.ks = InitSlot
	return
}

type SHsetnxField SCompleted

func (c SHsetnxField) Value(Value string) SHsetnxValue {
	return SHsetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SHsetnxKey SCompleted

func (c SHsetnxKey) Field(Field string) SHsetnxField {
	return SHsetnxField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHsetnxValue SCompleted

func (c SHsetnxValue) Build() SCompleted {
	return SCompleted(c)
}

type SHstrlen SCompleted

func (c SHstrlen) Key(Key string) SHstrlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHstrlenKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SHstrlenField{cf: c.cf, cs: append(c.cs, Field)}
}

type SHvals SCompleted

func (c SHvals) Key(Key string) SHvalsKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SHvalsKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SIncrKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SIncrbyKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SIncrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type SIncrbyfloat SCompleted

func (c SIncrbyfloat) Key(Key string) SIncrbyfloatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SIncrbyfloatKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SIncrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type SInfo SCompleted

func (c SInfo) Section(Section string) SInfoSection {
	return SInfoSection{cf: c.cf, cs: append(c.cs, Section)}
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
	return SKeysPattern{cf: c.cf, cs: append(c.cs, Pattern)}
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
	return SLatencyGraphEvent{cf: c.cf, cs: append(c.cs, Event)}
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
	return SLatencyHistoryEvent{cf: c.cf, cs: append(c.cs, Event)}
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
	return SLatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
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
	return SLatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
}

func (c SLatencyResetEvent) Build() SCompleted {
	return SCompleted(c)
}

type SLindex SCompleted

func (c SLindex) Key(Key string) SLindexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLindexKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SLindexIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type SLinsert SCompleted

func (c SLinsert) Key(Key string) SLinsertKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLinsertKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SLinsertWhereBefore{cf: c.cf, cs: append(c.cs, "BEFORE")}
}

func (c SLinsertKey) After() SLinsertWhereAfter {
	return SLinsertWhereAfter{cf: c.cf, cs: append(c.cs, "AFTER")}
}

type SLinsertPivot SCompleted

func (c SLinsertPivot) Element(Element string) SLinsertElement {
	return SLinsertElement{cf: c.cf, cs: append(c.cs, Element)}
}

type SLinsertWhereAfter SCompleted

func (c SLinsertWhereAfter) Pivot(Pivot string) SLinsertPivot {
	return SLinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type SLinsertWhereBefore SCompleted

func (c SLinsertWhereBefore) Pivot(Pivot string) SLinsertPivot {
	return SLinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type SLlen SCompleted

func (c SLlen) Key(Key string) SLlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLlenKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SLmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Lmove() (c SLmove) {
	c.cs = append(b.get(), "LMOVE")
	c.ks = InitSlot
	return
}

type SLmoveDestination SCompleted

func (c SLmoveDestination) Left() SLmoveWherefromLeft {
	return SLmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmoveDestination) Right() SLmoveWherefromRight {
	return SLmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SLmoveSource SCompleted

func (c SLmoveSource) Destination(Destination string) SLmoveDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SLmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SLmoveWherefromLeft SCompleted

func (c SLmoveWherefromLeft) Left() SLmoveWheretoLeft {
	return SLmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmoveWherefromLeft) Right() SLmoveWheretoRight {
	return SLmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SLmoveWherefromRight SCompleted

func (c SLmoveWherefromRight) Left() SLmoveWheretoLeft {
	return SLmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmoveWherefromRight) Right() SLmoveWheretoRight {
	return SLmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
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
	return SLmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
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
	return SLmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmpopKey) Right() SLmpopWhereRight {
	return SLmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

func (c SLmpopKey) Key(Key ...string) SLmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SLmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SLmpopNumkeys SCompleted

func (c SLmpopNumkeys) Key(Key ...string) SLmpopKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SLmpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SLmpopNumkeys) Left() SLmpopWhereLeft {
	return SLmpopWhereLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c SLmpopNumkeys) Right() SLmpopWhereRight {
	return SLmpopWhereRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type SLmpopWhereLeft SCompleted

func (c SLmpopWhereLeft) Count(Count int64) SLmpopCount {
	return SLmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SLmpopWhereLeft) Build() SCompleted {
	return SCompleted(c)
}

type SLmpopWhereRight SCompleted

func (c SLmpopWhereRight) Count(Count int64) SLmpopCount {
	return SLmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SLmpopWhereRight) Build() SCompleted {
	return SCompleted(c)
}

type SLolwut SCompleted

func (c SLolwut) Version(Version int64) SLolwutVersion {
	return SLolwutVersion{cf: c.cf, cs: append(c.cs, "VERSION", strconv.FormatInt(Version, 10))}
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
	return SLpopKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SLpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SLpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SLpos SCompleted

func (c SLpos) Key(Key string) SLposKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLposKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lpos() (c SLpos) {
	c.cs = append(b.get(), "LPOS")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SLposCount SCompleted

func (c SLposCount) Maxlen(Len int64) SLposMaxlen {
	return SLposMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c SLposCount) Build() SCompleted {
	return SCompleted(c)
}

func (c SLposCount) Cache() SCacheable {
	return SCacheable(c)
}

type SLposElement SCompleted

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

type SLposKey SCompleted

func (c SLposKey) Element(Element string) SLposElement {
	return SLposElement{cf: c.cf, cs: append(c.cs, Element)}
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

type SLpush SCompleted

func (c SLpush) Key(Key string) SLpushKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lpush() (c SLpush) {
	c.cs = append(b.get(), "LPUSH")
	c.ks = InitSlot
	return
}

type SLpushElement SCompleted

func (c SLpushElement) Element(Element ...string) SLpushElement {
	return SLpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SLpushElement) Build() SCompleted {
	return SCompleted(c)
}

type SLpushKey SCompleted

func (c SLpushKey) Element(Element ...string) SLpushElement {
	return SLpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SLpushx SCompleted

func (c SLpushx) Key(Key string) SLpushxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lpushx() (c SLpushx) {
	c.cs = append(b.get(), "LPUSHX")
	c.ks = InitSlot
	return
}

type SLpushxElement SCompleted

func (c SLpushxElement) Element(Element ...string) SLpushxElement {
	return SLpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SLpushxElement) Build() SCompleted {
	return SCompleted(c)
}

type SLpushxKey SCompleted

func (c SLpushxKey) Element(Element ...string) SLpushxElement {
	return SLpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SLrange SCompleted

func (c SLrange) Key(Key string) SLrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lrange() (c SLrange) {
	c.cs = append(b.get(), "LRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SLrangeKey SCompleted

func (c SLrangeKey) Start(Start int64) SLrangeStart {
	return SLrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type SLrangeStart SCompleted

func (c SLrangeStart) Stop(Stop int64) SLrangeStop {
	return SLrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
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
	return SLremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Lrem() (c SLrem) {
	c.cs = append(b.get(), "LREM")
	c.ks = InitSlot
	return
}

type SLremCount SCompleted

func (c SLremCount) Element(Element string) SLremElement {
	return SLremElement{cf: c.cf, cs: append(c.cs, Element)}
}

type SLremElement SCompleted

func (c SLremElement) Build() SCompleted {
	return SCompleted(c)
}

type SLremKey SCompleted

func (c SLremKey) Count(Count int64) SLremCount {
	return SLremCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type SLset SCompleted

func (c SLset) Key(Key string) SLsetKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLsetKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SLsetElement{cf: c.cf, cs: append(c.cs, Element)}
}

type SLsetKey SCompleted

func (c SLsetKey) Index(Index int64) SLsetIndex {
	return SLsetIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type SLtrim SCompleted

func (c SLtrim) Key(Key string) SLtrimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SLtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Ltrim() (c SLtrim) {
	c.cs = append(b.get(), "LTRIM")
	c.ks = InitSlot
	return
}

type SLtrimKey SCompleted

func (c SLtrimKey) Start(Start int64) SLtrimStart {
	return SLtrimStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type SLtrimStart SCompleted

func (c SLtrimStart) Stop(Stop int64) SLtrimStop {
	return SLtrimStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
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
	return SMemoryUsageKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) MemoryUsage() (c SMemoryUsage) {
	c.cs = append(b.get(), "MEMORY", "USAGE")
	c.ks = InitSlot
	return
}

type SMemoryUsageKey SCompleted

func (c SMemoryUsageKey) Samples(Count int64) SMemoryUsageSamples {
	return SMemoryUsageSamples{cf: c.cf, cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10))}
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
	return SMgetKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SMgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMgetKey) Build() SCompleted {
	return SCompleted(c)
}

type SMigrate SCompleted

func (c SMigrate) Host(Host string) SMigrateHost {
	return SMigrateHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *SBuilder) Migrate() (c SMigrate) {
	c.cs = append(b.get(), "MIGRATE")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SMigrateAuth SCompleted

func (c SMigrateAuth) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c SMigrateAuth) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateAuth2) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateCopyCopy SCompleted

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
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateCopyCopy) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateDestinationDb SCompleted

func (c SMigrateDestinationDb) Timeout(Timeout int64) SMigrateTimeout {
	return SMigrateTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type SMigrateHost SCompleted

func (c SMigrateHost) Port(Port string) SMigratePort {
	return SMigratePort{cf: c.cf, cs: append(c.cs, Port)}
}

type SMigrateKeyEmpty SCompleted

func (c SMigrateKeyEmpty) DestinationDb(DestinationDb int64) SMigrateDestinationDb {
	return SMigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type SMigrateKeyKey SCompleted

func (c SMigrateKeyKey) DestinationDb(DestinationDb int64) SMigrateDestinationDb {
	return SMigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type SMigrateKeys SCompleted

func (c SMigrateKeys) Keys(Keys ...string) SMigrateKeys {
	for _, k := range Keys {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Keys...)}
}

func (c SMigrateKeys) Build() SCompleted {
	return SCompleted(c)
}

type SMigratePort SCompleted

func (c SMigratePort) Key() SMigrateKeyKey {
	return SMigrateKeyKey{cf: c.cf, cs: append(c.cs, "key")}
}

func (c SMigratePort) Empty() SMigrateKeyEmpty {
	return SMigrateKeyEmpty{cf: c.cf, cs: append(c.cs, "\"\"")}
}

type SMigrateReplaceReplace SCompleted

func (c SMigrateReplaceReplace) Auth(Password string) SMigrateAuth {
	return SMigrateAuth{cf: c.cf, cs: append(c.cs, "AUTH", Password)}
}

func (c SMigrateReplaceReplace) Auth2(UsernamePassword string) SMigrateAuth2 {
	return SMigrateAuth2{cf: c.cf, cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c SMigrateReplaceReplace) Keys(Key ...string) SMigrateKeys {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SMigrateReplaceReplace) Build() SCompleted {
	return SCompleted(c)
}

type SMigrateTimeout SCompleted

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
		c.ks = checkSlot(c.ks, slot(k))
	}
	c.cs = append(c.cs, "KEYS")
	return SMigrateKeys{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SModuleLoadPath{cf: c.cf, cs: append(c.cs, Path)}
}

func (b *SBuilder) ModuleLoad() (c SModuleLoad) {
	c.cs = append(b.get(), "MODULE", "LOAD")
	c.ks = InitSlot
	return
}

type SModuleLoadArg SCompleted

func (c SModuleLoadArg) Arg(Arg ...string) SModuleLoadArg {
	return SModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SModuleLoadArg) Build() SCompleted {
	return SCompleted(c)
}

type SModuleLoadPath SCompleted

func (c SModuleLoadPath) Arg(Arg ...string) SModuleLoadArg {
	return SModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c SModuleLoadPath) Build() SCompleted {
	return SCompleted(c)
}

type SModuleUnload SCompleted

func (c SModuleUnload) Name(Name string) SModuleUnloadName {
	return SModuleUnloadName{cf: c.cf, cs: append(c.cs, Name)}
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
	return SMoveKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SMoveDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Db, 10))}
}

type SMset SCompleted

func (c SMset) KeyValue() SMsetKeyValue {
	return SMsetKeyValue{cf: c.cf, cs: c.cs}
}

func (b *SBuilder) Mset() (c SMset) {
	c.cs = append(b.get(), "MSET")
	c.ks = InitSlot
	return
}

type SMsetKeyValue SCompleted

func (c SMsetKeyValue) KeyValue(Key string, Value string) SMsetKeyValue {
	c.ks = checkSlot(c.ks, slot(Key))
	return SMsetKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c SMsetKeyValue) Build() SCompleted {
	return SCompleted(c)
}

type SMsetnx SCompleted

func (c SMsetnx) KeyValue() SMsetnxKeyValue {
	return SMsetnxKeyValue{cf: c.cf, cs: c.cs}
}

func (b *SBuilder) Msetnx() (c SMsetnx) {
	c.cs = append(b.get(), "MSETNX")
	c.ks = InitSlot
	return
}

type SMsetnxKeyValue SCompleted

func (c SMsetnxKeyValue) KeyValue(Key string, Value string) SMsetnxKeyValue {
	c.ks = checkSlot(c.ks, slot(Key))
	return SMsetnxKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
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
	return SObjectSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *SBuilder) Object() (c SObject) {
	c.cs = append(b.get(), "OBJECT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SObjectArguments SCompleted

func (c SObjectArguments) Arguments(Arguments ...string) SObjectArguments {
	return SObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c SObjectArguments) Build() SCompleted {
	return SCompleted(c)
}

type SObjectSubcommand SCompleted

func (c SObjectSubcommand) Arguments(Arguments ...string) SObjectArguments {
	return SObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c SObjectSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SPersist SCompleted

func (c SPersist) Key(Key string) SPersistKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPersistKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SPexpireKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SPexpireMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type SPexpireMilliseconds SCompleted

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

type SPexpireat SCompleted

func (c SPexpireat) Key(Key string) SPexpireatKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPexpireatKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SPexpireatMillisecondsTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10))}
}

type SPexpireatMillisecondsTimestamp SCompleted

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

type SPexpiretime SCompleted

func (c SPexpiretime) Key(Key string) SPexpiretimeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPexpiretimeKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SPfaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Pfadd() (c SPfadd) {
	c.cs = append(b.get(), "PFADD")
	c.ks = InitSlot
	return
}

type SPfaddElement SCompleted

func (c SPfaddElement) Element(Element ...string) SPfaddElement {
	return SPfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SPfaddElement) Build() SCompleted {
	return SCompleted(c)
}

type SPfaddKey SCompleted

func (c SPfaddKey) Element(Element ...string) SPfaddElement {
	return SPfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SPfaddKey) Build() SCompleted {
	return SCompleted(c)
}

type SPfcount SCompleted

func (c SPfcount) Key(Key ...string) SPfcountKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SPfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SPfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SPfcountKey) Build() SCompleted {
	return SCompleted(c)
}

type SPfmerge SCompleted

func (c SPfmerge) Destkey(Destkey string) SPfmergeDestkey {
	c.ks = checkSlot(c.ks, slot(Destkey))
	return SPfmergeDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
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
	return SPfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

type SPfmergeSourcekey SCompleted

func (c SPfmergeSourcekey) Sourcekey(Sourcekey ...string) SPfmergeSourcekey {
	for _, k := range Sourcekey {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SPfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

func (c SPfmergeSourcekey) Build() SCompleted {
	return SCompleted(c)
}

type SPing SCompleted

func (c SPing) Message(Message string) SPingMessage {
	return SPingMessage{cf: c.cf, cs: append(c.cs, Message)}
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
	return SPsetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Psetex() (c SPsetex) {
	c.cs = append(b.get(), "PSETEX")
	c.ks = InitSlot
	return
}

type SPsetexKey SCompleted

func (c SPsetexKey) Milliseconds(Milliseconds int64) SPsetexMilliseconds {
	return SPsetexMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type SPsetexMilliseconds SCompleted

func (c SPsetexMilliseconds) Value(Value string) SPsetexValue {
	return SPsetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SPsetexValue SCompleted

func (c SPsetexValue) Build() SCompleted {
	return SCompleted(c)
}

type SPsubscribe SCompleted

func (c SPsubscribe) Pattern(Pattern ...string) SPsubscribePattern {
	return SPsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (b *SBuilder) Psubscribe() (c SPsubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	c.cf = noRetTag
	c.ks = InitSlot
	return
}

type SPsubscribePattern SCompleted

func (c SPsubscribePattern) Pattern(Pattern ...string) SPsubscribePattern {
	return SPsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c SPsubscribePattern) Build() SCompleted {
	return SCompleted(c)
}

type SPsync SCompleted

func (c SPsync) Replicationid(Replicationid int64) SPsyncReplicationid {
	return SPsyncReplicationid{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Replicationid, 10))}
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
	return SPsyncOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SPttl SCompleted

func (c SPttl) Key(Key string) SPttlKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SPttlKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SPublishChannel{cf: c.cf, cs: append(c.cs, Channel)}
}

func (b *SBuilder) Publish() (c SPublish) {
	c.cs = append(b.get(), "PUBLISH")
	c.ks = InitSlot
	return
}

type SPublishChannel SCompleted

func (c SPublishChannel) Message(Message string) SPublishMessage {
	return SPublishMessage{cf: c.cf, cs: append(c.cs, Message)}
}

type SPublishMessage SCompleted

func (c SPublishMessage) Build() SCompleted {
	return SCompleted(c)
}

type SPubsub SCompleted

func (c SPubsub) Subcommand(Subcommand string) SPubsubSubcommand {
	return SPubsubSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
}

func (b *SBuilder) Pubsub() (c SPubsub) {
	c.cs = append(b.get(), "PUBSUB")
	c.ks = InitSlot
	return
}

type SPubsubArgument SCompleted

func (c SPubsubArgument) Argument(Argument ...string) SPubsubArgument {
	return SPubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c SPubsubArgument) Build() SCompleted {
	return SCompleted(c)
}

type SPubsubSubcommand SCompleted

func (c SPubsubSubcommand) Argument(Argument ...string) SPubsubArgument {
	return SPubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c SPubsubSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SPunsubscribe SCompleted

func (c SPunsubscribe) Pattern(Pattern ...string) SPunsubscribePattern {
	return SPunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
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
	return SPunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
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
	return SRenameKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Rename() (c SRename) {
	c.cs = append(b.get(), "RENAME")
	c.ks = InitSlot
	return
}

type SRenameKey SCompleted

func (c SRenameKey) Newkey(Newkey string) SRenameNewkey {
	c.ks = checkSlot(c.ks, slot(Newkey))
	return SRenameNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type SRenameNewkey SCompleted

func (c SRenameNewkey) Build() SCompleted {
	return SCompleted(c)
}

type SRenamenx SCompleted

func (c SRenamenx) Key(Key string) SRenamenxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRenamenxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Renamenx() (c SRenamenx) {
	c.cs = append(b.get(), "RENAMENX")
	c.ks = InitSlot
	return
}

type SRenamenxKey SCompleted

func (c SRenamenxKey) Newkey(Newkey string) SRenamenxNewkey {
	c.ks = checkSlot(c.ks, slot(Newkey))
	return SRenamenxNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type SRenamenxNewkey SCompleted

func (c SRenamenxNewkey) Build() SCompleted {
	return SCompleted(c)
}

type SReplicaof SCompleted

func (c SReplicaof) Host(Host string) SReplicaofHost {
	return SReplicaofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *SBuilder) Replicaof() (c SReplicaof) {
	c.cs = append(b.get(), "REPLICAOF")
	c.ks = InitSlot
	return
}

type SReplicaofHost SCompleted

func (c SReplicaofHost) Port(Port string) SReplicaofPort {
	return SReplicaofPort{cf: c.cf, cs: append(c.cs, Port)}
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
	return SRestoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Restore() (c SRestore) {
	c.cs = append(b.get(), "RESTORE")
	c.ks = InitSlot
	return
}

type SRestoreAbsttlAbsttl SCompleted

func (c SRestoreAbsttlAbsttl) Idletime(Seconds int64) SRestoreIdletime {
	return SRestoreIdletime{cf: c.cf, cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c SRestoreAbsttlAbsttl) Freq(Frequency int64) SRestoreFreq {
	return SRestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
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
	return SRestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c SRestoreIdletime) Build() SCompleted {
	return SCompleted(c)
}

type SRestoreKey SCompleted

func (c SRestoreKey) Ttl(Ttl int64) SRestoreTtl {
	return SRestoreTtl{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Ttl, 10))}
}

type SRestoreReplaceReplace SCompleted

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

type SRestoreSerializedValue SCompleted

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

type SRestoreTtl SCompleted

func (c SRestoreTtl) SerializedValue(SerializedValue string) SRestoreSerializedValue {
	return SRestoreSerializedValue{cf: c.cf, cs: append(c.cs, SerializedValue)}
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
	return SRpopKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SRpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SRpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SRpoplpush SCompleted

func (c SRpoplpush) Source(Source string) SRpoplpushSource {
	c.ks = checkSlot(c.ks, slot(Source))
	return SRpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
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
	return SRpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SRpush SCompleted

func (c SRpush) Key(Key string) SRpushKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRpushKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Rpush() (c SRpush) {
	c.cs = append(b.get(), "RPUSH")
	c.ks = InitSlot
	return
}

type SRpushElement SCompleted

func (c SRpushElement) Element(Element ...string) SRpushElement {
	return SRpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SRpushElement) Build() SCompleted {
	return SCompleted(c)
}

type SRpushKey SCompleted

func (c SRpushKey) Element(Element ...string) SRpushElement {
	return SRpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SRpushx SCompleted

func (c SRpushx) Key(Key string) SRpushxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SRpushxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Rpushx() (c SRpushx) {
	c.cs = append(b.get(), "RPUSHX")
	c.ks = InitSlot
	return
}

type SRpushxElement SCompleted

func (c SRpushxElement) Element(Element ...string) SRpushxElement {
	return SRpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c SRpushxElement) Build() SCompleted {
	return SCompleted(c)
}

type SRpushxKey SCompleted

func (c SRpushxKey) Element(Element ...string) SRpushxElement {
	return SRpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type SSadd SCompleted

func (c SSadd) Key(Key string) SSaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Sadd() (c SSadd) {
	c.cs = append(b.get(), "SADD")
	c.ks = InitSlot
	return
}

type SSaddKey SCompleted

func (c SSaddKey) Member(Member ...string) SSaddMember {
	return SSaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SSaddMember SCompleted

func (c SSaddMember) Member(Member ...string) SSaddMember {
	return SSaddMember{cf: c.cf, cs: append(c.cs, Member...)}
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
	return SScanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

func (b *SBuilder) Scan() (c SScan) {
	c.cs = append(b.get(), "SCAN")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SScanCount SCompleted

func (c SScanCount) Type(Type string) SScanType {
	return SScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c SScanCount) Build() SCompleted {
	return SCompleted(c)
}

type SScanCursor SCompleted

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

type SScanMatch SCompleted

func (c SScanMatch) Count(Count int64) SScanCount {
	return SScanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SScanMatch) Type(Type string) SScanType {
	return SScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
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
	return SScardKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (b *SBuilder) ScriptExists() (c SScriptExists) {
	c.cs = append(b.get(), "SCRIPT", "EXISTS")
	c.ks = InitSlot
	return
}

type SScriptExistsSha1 SCompleted

func (c SScriptExistsSha1) Sha1(Sha1 ...string) SScriptExistsSha1 {
	return SScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (c SScriptExistsSha1) Build() SCompleted {
	return SCompleted(c)
}

type SScriptFlush SCompleted

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
	return SScriptLoadScript{cf: c.cf, cs: append(c.cs, Script)}
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
	return SSdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SSdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSdiffKey) Build() SCompleted {
	return SCompleted(c)
}

type SSdiffstore SCompleted

func (c SSdiffstore) Destination(Destination string) SSdiffstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return SSdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SSdiffstoreKey SCompleted

func (c SSdiffstoreKey) Key(Key ...string) SSdiffstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSdiffstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSelect SCompleted

func (c SSelect) Index(Index int64) SSelectIndex {
	return SSelectIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
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
	return SSetKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Set() (c SSet) {
	c.cs = append(b.get(), "SET")
	c.ks = InitSlot
	return
}

type SSetConditionNx SCompleted

func (c SSetConditionNx) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetConditionNx) Build() SCompleted {
	return SCompleted(c)
}

type SSetConditionXx SCompleted

func (c SSetConditionXx) Get() SSetGetGet {
	return SSetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SSetConditionXx) Build() SCompleted {
	return SCompleted(c)
}

type SSetExpirationEx SCompleted

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

type SSetExpirationExat SCompleted

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

type SSetExpirationKeepttl SCompleted

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

type SSetExpirationPx SCompleted

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

type SSetExpirationPxat SCompleted

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

type SSetGetGet SCompleted

func (c SSetGetGet) Build() SCompleted {
	return SCompleted(c)
}

type SSetKey SCompleted

func (c SSetKey) Value(Value string) SSetValue {
	return SSetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetValue SCompleted

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

type SSetbit SCompleted

func (c SSetbit) Key(Key string) SSetbitKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetbitKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setbit() (c SSetbit) {
	c.cs = append(b.get(), "SETBIT")
	c.ks = InitSlot
	return
}

type SSetbitKey SCompleted

func (c SSetbitKey) Offset(Offset int64) SSetbitOffset {
	return SSetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SSetbitOffset SCompleted

func (c SSetbitOffset) Value(Value int64) SSetbitValue {
	return SSetbitValue{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Value, 10))}
}

type SSetbitValue SCompleted

func (c SSetbitValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetex SCompleted

func (c SSetex) Key(Key string) SSetexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setex() (c SSetex) {
	c.cs = append(b.get(), "SETEX")
	c.ks = InitSlot
	return
}

type SSetexKey SCompleted

func (c SSetexKey) Seconds(Seconds int64) SSetexSeconds {
	return SSetexSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SSetexSeconds SCompleted

func (c SSetexSeconds) Value(Value string) SSetexValue {
	return SSetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetexValue SCompleted

func (c SSetexValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetnx SCompleted

func (c SSetnx) Key(Key string) SSetnxKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetnxKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setnx() (c SSetnx) {
	c.cs = append(b.get(), "SETNX")
	c.ks = InitSlot
	return
}

type SSetnxKey SCompleted

func (c SSetnxKey) Value(Value string) SSetnxValue {
	return SSetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetnxValue SCompleted

func (c SSetnxValue) Build() SCompleted {
	return SCompleted(c)
}

type SSetrange SCompleted

func (c SSetrange) Key(Key string) SSetrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSetrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Setrange() (c SSetrange) {
	c.cs = append(b.get(), "SETRANGE")
	c.ks = InitSlot
	return
}

type SSetrangeKey SCompleted

func (c SSetrangeKey) Offset(Offset int64) SSetrangeOffset {
	return SSetrangeOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SSetrangeOffset SCompleted

func (c SSetrangeOffset) Value(Value string) SSetrangeValue {
	return SSetrangeValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SSetrangeValue SCompleted

func (c SSetrangeValue) Build() SCompleted {
	return SCompleted(c)
}

type SShutdown SCompleted

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
	return SSinterKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SSinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSinterKey) Build() SCompleted {
	return SCompleted(c)
}

type SSintercard SCompleted

func (c SSintercard) Key(Key ...string) SSintercardKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SSintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSintercardKey) Build() SCompleted {
	return SCompleted(c)
}

type SSinterstore SCompleted

func (c SSinterstore) Destination(Destination string) SSinterstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return SSinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SSinterstoreKey SCompleted

func (c SSinterstoreKey) Key(Key ...string) SSinterstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSinterstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSismember SCompleted

func (c SSismember) Key(Key string) SSismemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Sismember() (c SSismember) {
	c.cs = append(b.get(), "SISMEMBER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSismemberKey SCompleted

func (c SSismemberKey) Member(Member string) SSismemberMember {
	return SSismemberMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return SSlaveofHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *SBuilder) Slaveof() (c SSlaveof) {
	c.cs = append(b.get(), "SLAVEOF")
	c.ks = InitSlot
	return
}

type SSlaveofHost SCompleted

func (c SSlaveofHost) Port(Port string) SSlaveofPort {
	return SSlaveofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type SSlaveofPort SCompleted

func (c SSlaveofPort) Build() SCompleted {
	return SCompleted(c)
}

type SSlowlog SCompleted

func (c SSlowlog) Subcommand(Subcommand string) SSlowlogSubcommand {
	return SSlowlogSubcommand{cf: c.cf, cs: append(c.cs, Subcommand)}
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
	return SSlowlogArgument{cf: c.cf, cs: append(c.cs, Argument)}
}

func (c SSlowlogSubcommand) Build() SCompleted {
	return SCompleted(c)
}

type SSmembers SCompleted

func (c SSmembers) Key(Key string) SSmembersKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSmembersKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SSmismemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Smismember() (c SSmismember) {
	c.cs = append(b.get(), "SMISMEMBER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSmismemberKey SCompleted

func (c SSmismemberKey) Member(Member ...string) SSmismemberMember {
	return SSmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SSmismemberMember SCompleted

func (c SSmismemberMember) Member(Member ...string) SSmismemberMember {
	return SSmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
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
	return SSmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *SBuilder) Smove() (c SSmove) {
	c.cs = append(b.get(), "SMOVE")
	c.ks = InitSlot
	return
}

type SSmoveDestination SCompleted

func (c SSmoveDestination) Member(Member string) SSmoveMember {
	return SSmoveMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SSmoveMember SCompleted

func (c SSmoveMember) Build() SCompleted {
	return SCompleted(c)
}

type SSmoveSource SCompleted

func (c SSmoveSource) Destination(Destination string) SSmoveDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type SSort SCompleted

func (c SSort) Key(Key string) SSortKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSortKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Sort() (c SSort) {
	c.cs = append(b.get(), "SORT")
	c.ks = InitSlot
	return
}

type SSortBy SCompleted

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
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortBy) Build() SCompleted {
	return SCompleted(c)
}

type SSortGet SCompleted

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
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortGet) Get(Get ...string) SSortGet {
	return SSortGet{cf: c.cf, cs: append(c.cs, Get...)}
}

func (c SSortGet) Build() SCompleted {
	return SCompleted(c)
}

type SSortKey SCompleted

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
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortKey) Build() SCompleted {
	return SCompleted(c)
}

type SSortLimit SCompleted

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
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortLimit) Build() SCompleted {
	return SCompleted(c)
}

type SSortOrderAsc SCompleted

func (c SSortOrderAsc) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortOrderAsc) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

type SSortOrderDesc SCompleted

func (c SSortOrderDesc) Alpha() SSortSortingAlpha {
	return SSortSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortOrderDesc) Store(Destination string) SSortStore {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SSortOrderDesc) Build() SCompleted {
	return SCompleted(c)
}

type SSortRo SCompleted

func (c SSortRo) Key(Key string) SSortRoKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSortRoKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) SortRo() (c SSortRo) {
	c.cs = append(b.get(), "SORT_RO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SSortRoBy SCompleted

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

type SSortRoGet SCompleted

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

type SSortRoKey SCompleted

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

type SSortRoLimit SCompleted

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

type SSortRoOrderAsc SCompleted

func (c SSortRoOrderAsc) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SSortRoOrderAsc) Build() SCompleted {
	return SCompleted(c)
}

func (c SSortRoOrderAsc) Cache() SCacheable {
	return SCacheable(c)
}

type SSortRoOrderDesc SCompleted

func (c SSortRoOrderDesc) Alpha() SSortRoSortingAlpha {
	return SSortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
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
	return SSortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
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
	return SSpopKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SSpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SSpopKey) Build() SCompleted {
	return SCompleted(c)
}

type SSrandmember SCompleted

func (c SSrandmember) Key(Key string) SSrandmemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SSrandmemberCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SSrandmemberKey) Build() SCompleted {
	return SCompleted(c)
}

type SSrem SCompleted

func (c SSrem) Key(Key string) SSremKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Srem() (c SSrem) {
	c.cs = append(b.get(), "SREM")
	c.ks = InitSlot
	return
}

type SSremKey SCompleted

func (c SSremKey) Member(Member ...string) SSremMember {
	return SSremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SSremMember SCompleted

func (c SSremMember) Member(Member ...string) SSremMember {
	return SSremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SSremMember) Build() SCompleted {
	return SCompleted(c)
}

type SSscan SCompleted

func (c SSscan) Key(Key string) SSscanKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SSscanKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SSscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SSscanCursor) Count(Count int64) SSscanCount {
	return SSscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SSscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SSscanKey SCompleted

func (c SSscanKey) Cursor(Cursor int64) SSscanCursor {
	return SSscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SSscanMatch SCompleted

func (c SSscanMatch) Count(Count int64) SSscanCount {
	return SSscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SSscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SStralgo SCompleted

func (c SStralgo) Lcs() SStralgoAlgorithmLcs {
	return SStralgoAlgorithmLcs{cf: c.cf, cs: append(c.cs, "LCS")}
}

func (b *SBuilder) Stralgo() (c SStralgo) {
	c.cs = append(b.get(), "STRALGO")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SStralgoAlgoSpecificArgument SCompleted

func (c SStralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) SStralgoAlgoSpecificArgument {
	return SStralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

func (c SStralgoAlgoSpecificArgument) Build() SCompleted {
	return SCompleted(c)
}

type SStralgoAlgorithmLcs SCompleted

func (c SStralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) SStralgoAlgoSpecificArgument {
	return SStralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

type SStrlen SCompleted

func (c SStrlen) Key(Key string) SStrlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SStrlenKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SSubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (b *SBuilder) Subscribe() (c SSubscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	c.cf = noRetTag
	c.ks = InitSlot
	return
}

type SSubscribeChannel SCompleted

func (c SSubscribeChannel) Channel(Channel ...string) SSubscribeChannel {
	return SSubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c SSubscribeChannel) Build() SCompleted {
	return SCompleted(c)
}

type SSunion SCompleted

func (c SSunion) Key(Key ...string) SSunionKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSunionKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SSunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSunionKey) Build() SCompleted {
	return SCompleted(c)
}

type SSunionstore SCompleted

func (c SSunionstore) Destination(Destination string) SSunionstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SSunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return SSunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SSunionstoreKey SCompleted

func (c SSunionstoreKey) Key(Key ...string) SSunionstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SSunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SSunionstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SSwapdb SCompleted

func (c SSwapdb) Index1(Index1 int64) SSwapdbIndex1 {
	return SSwapdbIndex1{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index1, 10))}
}

func (b *SBuilder) Swapdb() (c SSwapdb) {
	c.cs = append(b.get(), "SWAPDB")
	c.ks = InitSlot
	return
}

type SSwapdbIndex1 SCompleted

func (c SSwapdbIndex1) Index2(Index2 int64) SSwapdbIndex2 {
	return SSwapdbIndex2{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index2, 10))}
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
	return STouchKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return STouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c STouchKey) Build() SCompleted {
	return SCompleted(c)
}

type STtl SCompleted

func (c STtl) Key(Key string) STtlKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return STtlKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return STypeKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SUnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SUnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SUnlinkKey) Build() SCompleted {
	return SCompleted(c)
}

type SUnsubscribe SCompleted

func (c SUnsubscribe) Channel(Channel ...string) SUnsubscribeChannel {
	return SUnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
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
	return SUnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
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
	return SWaitNumreplicas{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numreplicas, 10))}
}

func (b *SBuilder) Wait() (c SWait) {
	c.cs = append(b.get(), "WAIT")
	c.cf = blockTag
	c.ks = InitSlot
	return
}

type SWaitNumreplicas SCompleted

func (c SWaitNumreplicas) Timeout(Timeout int64) SWaitTimeout {
	return SWaitTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
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
	return SWatchKey{cf: c.cf, cs: append(c.cs, Key...)}
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
	return SWatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SWatchKey) Build() SCompleted {
	return SCompleted(c)
}

type SXack SCompleted

func (c SXack) Key(Key string) SXackKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXackKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xack() (c SXack) {
	c.cs = append(b.get(), "XACK")
	c.ks = InitSlot
	return
}

type SXackGroup SCompleted

func (c SXackGroup) Id(Id ...string) SXackId {
	return SXackId{cf: c.cf, cs: append(c.cs, Id...)}
}

type SXackId SCompleted

func (c SXackId) Id(Id ...string) SXackId {
	return SXackId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXackId) Build() SCompleted {
	return SCompleted(c)
}

type SXackKey SCompleted

func (c SXackKey) Group(Group string) SXackGroup {
	return SXackGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXadd SCompleted

func (c SXadd) Key(Key string) SXaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xadd() (c SXadd) {
	c.cs = append(b.get(), "XADD")
	c.ks = InitSlot
	return
}

type SXaddFieldValue SCompleted

func (c SXaddFieldValue) FieldValue(Field string, Value string) SXaddFieldValue {
	return SXaddFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c SXaddFieldValue) Build() SCompleted {
	return SCompleted(c)
}

type SXaddId SCompleted

func (c SXaddId) FieldValue() SXaddFieldValue {
	return SXaddFieldValue{cf: c.cf, cs: c.cs}
}

type SXaddKey SCompleted

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

type SXaddNomkstream SCompleted

func (c SXaddNomkstream) Maxlen() SXaddTrimStrategyMaxlen {
	return SXaddTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c SXaddNomkstream) Minid() SXaddTrimStrategyMinid {
	return SXaddTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

func (c SXaddNomkstream) Id(Id string) SXaddId {
	return SXaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXaddTrimLimit SCompleted

func (c SXaddTrimLimit) Id(Id string) SXaddId {
	return SXaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXaddTrimOperatorAlmost SCompleted

func (c SXaddTrimOperatorAlmost) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimOperatorExact SCompleted

func (c SXaddTrimOperatorExact) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimStrategyMaxlen SCompleted

func (c SXaddTrimStrategyMaxlen) Exact() SXaddTrimOperatorExact {
	return SXaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXaddTrimStrategyMaxlen) Almost() SXaddTrimOperatorAlmost {
	return SXaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXaddTrimStrategyMaxlen) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimStrategyMinid SCompleted

func (c SXaddTrimStrategyMinid) Exact() SXaddTrimOperatorExact {
	return SXaddTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXaddTrimStrategyMinid) Almost() SXaddTrimOperatorAlmost {
	return SXaddTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXaddTrimStrategyMinid) Threshold(Threshold string) SXaddTrimThreshold {
	return SXaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXaddTrimThreshold SCompleted

func (c SXaddTrimThreshold) Limit(Count int64) SXaddTrimLimit {
	return SXaddTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c SXaddTrimThreshold) Id(Id string) SXaddId {
	return SXaddId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXautoclaim SCompleted

func (c SXautoclaim) Key(Key string) SXautoclaimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXautoclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xautoclaim() (c SXautoclaim) {
	c.cs = append(b.get(), "XAUTOCLAIM")
	c.ks = InitSlot
	return
}

type SXautoclaimConsumer SCompleted

func (c SXautoclaimConsumer) MinIdleTime(MinIdleTime string) SXautoclaimMinIdleTime {
	return SXautoclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type SXautoclaimCount SCompleted

func (c SXautoclaimCount) Justid() SXautoclaimJustidJustid {
	return SXautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXautoclaimCount) Build() SCompleted {
	return SCompleted(c)
}

type SXautoclaimGroup SCompleted

func (c SXautoclaimGroup) Consumer(Consumer string) SXautoclaimConsumer {
	return SXautoclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type SXautoclaimJustidJustid SCompleted

func (c SXautoclaimJustidJustid) Build() SCompleted {
	return SCompleted(c)
}

type SXautoclaimKey SCompleted

func (c SXautoclaimKey) Group(Group string) SXautoclaimGroup {
	return SXautoclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXautoclaimMinIdleTime SCompleted

func (c SXautoclaimMinIdleTime) Start(Start string) SXautoclaimStart {
	return SXautoclaimStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXautoclaimStart SCompleted

func (c SXautoclaimStart) Count(Count int64) SXautoclaimCount {
	return SXautoclaimCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXautoclaimStart) Justid() SXautoclaimJustidJustid {
	return SXautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXautoclaimStart) Build() SCompleted {
	return SCompleted(c)
}

type SXclaim SCompleted

func (c SXclaim) Key(Key string) SXclaimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXclaimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xclaim() (c SXclaim) {
	c.cs = append(b.get(), "XCLAIM")
	c.ks = InitSlot
	return
}

type SXclaimConsumer SCompleted

func (c SXclaimConsumer) MinIdleTime(MinIdleTime string) SXclaimMinIdleTime {
	return SXclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type SXclaimForceForce SCompleted

func (c SXclaimForceForce) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXclaimForceForce) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimGroup SCompleted

func (c SXclaimGroup) Consumer(Consumer string) SXclaimConsumer {
	return SXclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type SXclaimId SCompleted

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

type SXclaimIdle SCompleted

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

type SXclaimJustidJustid SCompleted

func (c SXclaimJustidJustid) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimKey SCompleted

func (c SXclaimKey) Group(Group string) SXclaimGroup {
	return SXclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXclaimMinIdleTime SCompleted

func (c SXclaimMinIdleTime) Id(Id ...string) SXclaimId {
	return SXclaimId{cf: c.cf, cs: append(c.cs, Id...)}
}

type SXclaimRetrycount SCompleted

func (c SXclaimRetrycount) Force() SXclaimForceForce {
	return SXclaimForceForce{cf: c.cf, cs: append(c.cs, "FORCE")}
}

func (c SXclaimRetrycount) Justid() SXclaimJustidJustid {
	return SXclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c SXclaimRetrycount) Build() SCompleted {
	return SCompleted(c)
}

type SXclaimTime SCompleted

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

type SXdel SCompleted

func (c SXdel) Key(Key string) SXdelKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXdelKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xdel() (c SXdel) {
	c.cs = append(b.get(), "XDEL")
	c.ks = InitSlot
	return
}

type SXdelId SCompleted

func (c SXdelId) Id(Id ...string) SXdelId {
	return SXdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXdelId) Build() SCompleted {
	return SCompleted(c)
}

type SXdelKey SCompleted

func (c SXdelKey) Id(Id ...string) SXdelId {
	return SXdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

type SXgroup SCompleted

func (c SXgroup) Create(Key string, Groupname string) SXgroupCreateCreate {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateCreate{cf: c.cf, cs: append(c.cs, "CREATE", Key, Groupname)}
}

func (c SXgroup) Setid(Key string, Groupname string) SXgroupSetidSetid {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c SXgroup) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroup) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroup) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
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
	return SXgroupCreateId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXgroupCreateId SCompleted

func (c SXgroupCreateId) Mkstream() SXgroupCreateMkstream {
	return SXgroupCreateMkstream{cf: c.cf, cs: append(c.cs, "MKSTREAM")}
}

func (c SXgroupCreateId) Setid(Key string, Groupname string) SXgroupSetidSetid {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c SXgroupCreateId) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroupCreateId) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupCreateId) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupCreateId) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupCreateMkstream SCompleted

func (c SXgroupCreateMkstream) Setid(Key string, Groupname string) SXgroupSetidSetid {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c SXgroupCreateMkstream) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroupCreateMkstream) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupCreateMkstream) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupCreateMkstream) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupCreateconsumer SCompleted

func (c SXgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
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
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupDestroy) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupDestroy) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupSetidId SCompleted

func (c SXgroupSetidId) Destroy(Key string, Groupname string) SXgroupDestroy {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c SXgroupSetidId) Createconsumer(Key string, Groupname string, Consumername string) SXgroupCreateconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupSetidId) Delconsumer(Key string, Groupname string, Consumername string) SXgroupDelconsumer {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c SXgroupSetidId) Build() SCompleted {
	return SCompleted(c)
}

type SXgroupSetidSetid SCompleted

func (c SXgroupSetidSetid) Id(Id string) SXgroupSetidId {
	return SXgroupSetidId{cf: c.cf, cs: append(c.cs, Id)}
}

type SXinfo SCompleted

func (c SXinfo) Consumers(Key string, Groupname string) SXinfoConsumers {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoConsumers{cf: c.cf, cs: append(c.cs, "CONSUMERS", Key, Groupname)}
}

func (c SXinfo) Groups(Key string) SXinfoGroups {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoGroups{cf: c.cf, cs: append(c.cs, "GROUPS", Key)}
}

func (c SXinfo) Stream(Key string) SXinfoStream {
	c.ks = checkSlot(c.ks, slot(Key))
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
	c.ks = InitSlot
	return
}

type SXinfoConsumers SCompleted

func (c SXinfoConsumers) Groups(Key string) SXinfoGroups {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoGroups{cf: c.cf, cs: append(c.cs, "GROUPS", Key)}
}

func (c SXinfoConsumers) Stream(Key string) SXinfoStream {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c SXinfoConsumers) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c SXinfoConsumers) Build() SCompleted {
	return SCompleted(c)
}

type SXinfoGroups SCompleted

func (c SXinfoGroups) Stream(Key string) SXinfoStream {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXinfoStream{cf: c.cf, cs: append(c.cs, "STREAM", Key)}
}

func (c SXinfoGroups) Help() SXinfoHelpHelp {
	return SXinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
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
	return SXinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c SXinfoStream) Build() SCompleted {
	return SCompleted(c)
}

type SXlen SCompleted

func (c SXlen) Key(Key string) SXlenKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXlenKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SXpendingKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SXpendingFiltersConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

func (c SXpendingFiltersCount) Build() SCompleted {
	return SCompleted(c)
}

type SXpendingFiltersEnd SCompleted

func (c SXpendingFiltersEnd) Count(Count int64) SXpendingFiltersCount {
	return SXpendingFiltersCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type SXpendingFiltersIdle SCompleted

func (c SXpendingFiltersIdle) Start(Start string) SXpendingFiltersStart {
	return SXpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXpendingFiltersStart SCompleted

func (c SXpendingFiltersStart) End(End string) SXpendingFiltersEnd {
	return SXpendingFiltersEnd{cf: c.cf, cs: append(c.cs, End)}
}

type SXpendingGroup SCompleted

func (c SXpendingGroup) Idle(MinIdleTime int64) SXpendingFiltersIdle {
	return SXpendingFiltersIdle{cf: c.cf, cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10))}
}

func (c SXpendingGroup) Start(Start string) SXpendingFiltersStart {
	return SXpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXpendingKey SCompleted

func (c SXpendingKey) Group(Group string) SXpendingGroup {
	return SXpendingGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type SXrange SCompleted

func (c SXrange) Key(Key string) SXrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXrangeKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SXrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXrangeEnd) Build() SCompleted {
	return SCompleted(c)
}

type SXrangeKey SCompleted

func (c SXrangeKey) Start(Start string) SXrangeStart {
	return SXrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXrangeStart SCompleted

func (c SXrangeStart) End(End string) SXrangeEnd {
	return SXrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type SXread SCompleted

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
	c.ks = InitSlot
	return
}

type SXreadBlock SCompleted

func (c SXreadBlock) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadCount SCompleted

func (c SXreadCount) Block(Milliseconds int64) SXreadBlock {
	c.cf = blockTag
	return SXreadBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c SXreadCount) Streams() SXreadStreamsStreams {
	return SXreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadId SCompleted

func (c SXreadId) Id(Id ...string) SXreadId {
	return SXreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadId) Build() SCompleted {
	return SCompleted(c)
}

type SXreadKey SCompleted

func (c SXreadKey) Id(Id ...string) SXreadId {
	return SXreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadKey) Key(Key ...string) SXreadKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXreadStreamsStreams SCompleted

func (c SXreadStreamsStreams) Key(Key ...string) SXreadKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXreadgroup SCompleted

func (c SXreadgroup) Group(Group string, Consumer string) SXreadgroupGroup {
	return SXreadgroupGroup{cf: c.cf, cs: append(c.cs, "GROUP", Group, Consumer)}
}

func (b *SBuilder) Xreadgroup() (c SXreadgroup) {
	c.cs = append(b.get(), "XREADGROUP")
	c.ks = InitSlot
	return
}

type SXreadgroupBlock SCompleted

func (c SXreadgroupBlock) Noack() SXreadgroupNoackNoack {
	return SXreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c SXreadgroupBlock) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadgroupCount SCompleted

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

type SXreadgroupGroup SCompleted

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

type SXreadgroupId SCompleted

func (c SXreadgroupId) Id(Id ...string) SXreadgroupId {
	return SXreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadgroupId) Build() SCompleted {
	return SCompleted(c)
}

type SXreadgroupKey SCompleted

func (c SXreadgroupKey) Id(Id ...string) SXreadgroupId {
	return SXreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c SXreadgroupKey) Key(Key ...string) SXreadgroupKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXreadgroupNoackNoack SCompleted

func (c SXreadgroupNoackNoack) Streams() SXreadgroupStreamsStreams {
	return SXreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type SXreadgroupStreamsStreams SCompleted

func (c SXreadgroupStreamsStreams) Key(Key ...string) SXreadgroupKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SXreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SXrevrange SCompleted

func (c SXrevrange) Key(Key string) SXrevrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SXrevrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type SXrevrangeKey SCompleted

func (c SXrevrangeKey) End(End string) SXrevrangeEnd {
	return SXrevrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type SXrevrangeStart SCompleted

func (c SXrevrangeStart) Count(Count int64) SXrevrangeCount {
	return SXrevrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SXrevrangeStart) Build() SCompleted {
	return SCompleted(c)
}

type SXtrim SCompleted

func (c SXtrim) Key(Key string) SXtrimKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SXtrimKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Xtrim() (c SXtrim) {
	c.cs = append(b.get(), "XTRIM")
	c.ks = InitSlot
	return
}

type SXtrimKey SCompleted

func (c SXtrimKey) Maxlen() SXtrimTrimStrategyMaxlen {
	return SXtrimTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c SXtrimKey) Minid() SXtrimTrimStrategyMinid {
	return SXtrimTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

type SXtrimTrimLimit SCompleted

func (c SXtrimTrimLimit) Build() SCompleted {
	return SCompleted(c)
}

type SXtrimTrimOperatorAlmost SCompleted

func (c SXtrimTrimOperatorAlmost) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimOperatorExact SCompleted

func (c SXtrimTrimOperatorExact) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimStrategyMaxlen SCompleted

func (c SXtrimTrimStrategyMaxlen) Exact() SXtrimTrimOperatorExact {
	return SXtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXtrimTrimStrategyMaxlen) Almost() SXtrimTrimOperatorAlmost {
	return SXtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXtrimTrimStrategyMaxlen) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimStrategyMinid SCompleted

func (c SXtrimTrimStrategyMinid) Exact() SXtrimTrimOperatorExact {
	return SXtrimTrimOperatorExact{cf: c.cf, cs: append(c.cs, "=")}
}

func (c SXtrimTrimStrategyMinid) Almost() SXtrimTrimOperatorAlmost {
	return SXtrimTrimOperatorAlmost{cf: c.cf, cs: append(c.cs, "~")}
}

func (c SXtrimTrimStrategyMinid) Threshold(Threshold string) SXtrimTrimThreshold {
	return SXtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type SXtrimTrimThreshold SCompleted

func (c SXtrimTrimThreshold) Limit(Count int64) SXtrimTrimLimit {
	return SXtrimTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c SXtrimTrimThreshold) Build() SCompleted {
	return SCompleted(c)
}

type SZadd SCompleted

func (c SZadd) Key(Key string) SZaddKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZaddKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zadd() (c SZadd) {
	c.cs = append(b.get(), "ZADD")
	c.ks = InitSlot
	return
}

type SZaddChangeCh SCompleted

func (c SZaddChangeCh) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddChangeCh) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: c.cs}
}

type SZaddComparisonGt SCompleted

func (c SZaddComparisonGt) Ch() SZaddChangeCh {
	return SZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SZaddComparisonGt) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddComparisonGt) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: c.cs}
}

type SZaddComparisonLt SCompleted

func (c SZaddComparisonLt) Ch() SZaddChangeCh {
	return SZaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c SZaddComparisonLt) Incr() SZaddIncrementIncr {
	return SZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c SZaddComparisonLt) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: c.cs}
}

type SZaddConditionNx SCompleted

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
	return SZaddScoreMember{cf: c.cf, cs: c.cs}
}

type SZaddConditionXx SCompleted

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
	return SZaddScoreMember{cf: c.cf, cs: c.cs}
}

type SZaddIncrementIncr SCompleted

func (c SZaddIncrementIncr) ScoreMember() SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: c.cs}
}

type SZaddKey SCompleted

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
	return SZaddScoreMember{cf: c.cf, cs: c.cs}
}

type SZaddScoreMember SCompleted

func (c SZaddScoreMember) ScoreMember(Score float64, Member string) SZaddScoreMember {
	return SZaddScoreMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member)}
}

func (c SZaddScoreMember) Build() SCompleted {
	return SCompleted(c)
}

type SZcard SCompleted

func (c SZcard) Key(Key string) SZcardKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZcardKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SZcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zcount() (c SZcount) {
	c.cs = append(b.get(), "ZCOUNT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZcountKey SCompleted

func (c SZcountKey) Min(Min float64) SZcountMin {
	return SZcountMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
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
	return SZcountMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type SZdiff SCompleted

func (c SZdiff) Numkeys(Numkeys int64) SZdiffNumkeys {
	return SZdiffNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Zdiff() (c SZdiff) {
	c.cs = append(b.get(), "ZDIFF")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZdiffKey SCompleted

func (c SZdiffKey) Withscores() SZdiffWithscoresWithscores {
	return SZdiffWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZdiffKey) Key(Key ...string) SZdiffKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZdiffKey) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffNumkeys SCompleted

func (c SZdiffNumkeys) Key(Key ...string) SZdiffKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZdiffWithscoresWithscores SCompleted

func (c SZdiffWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffstore SCompleted

func (c SZdiffstore) Destination(Destination string) SZdiffstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SZdiffstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

func (b *SBuilder) Zdiffstore() (c SZdiffstore) {
	c.cs = append(b.get(), "ZDIFFSTORE")
	c.ks = InitSlot
	return
}

type SZdiffstoreDestination SCompleted

func (c SZdiffstoreDestination) Numkeys(Numkeys int64) SZdiffstoreNumkeys {
	return SZdiffstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SZdiffstoreKey SCompleted

func (c SZdiffstoreKey) Key(Key ...string) SZdiffstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZdiffstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZdiffstoreNumkeys SCompleted

func (c SZdiffstoreNumkeys) Key(Key ...string) SZdiffstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZincrby SCompleted

func (c SZincrby) Key(Key string) SZincrbyKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZincrbyKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zincrby() (c SZincrby) {
	c.cs = append(b.get(), "ZINCRBY")
	c.ks = InitSlot
	return
}

type SZincrbyIncrement SCompleted

func (c SZincrbyIncrement) Member(Member string) SZincrbyMember {
	return SZincrbyMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SZincrbyKey SCompleted

func (c SZincrbyKey) Increment(Increment int64) SZincrbyIncrement {
	return SZincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type SZincrbyMember SCompleted

func (c SZincrbyMember) Build() SCompleted {
	return SCompleted(c)
}

type SZinter SCompleted

func (c SZinter) Numkeys(Numkeys int64) SZinterNumkeys {
	return SZinterNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Zinter() (c SZinter) {
	c.cs = append(b.get(), "ZINTER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZinterAggregateMax SCompleted

func (c SZinterAggregateMax) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZinterAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZinterAggregateMin SCompleted

func (c SZinterAggregateMin) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZinterAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZinterAggregateSum SCompleted

func (c SZinterAggregateSum) Withscores() SZinterWithscoresWithscores {
	return SZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZinterKey) Build() SCompleted {
	return SCompleted(c)
}

type SZinterNumkeys SCompleted

func (c SZinterNumkeys) Key(Key ...string) SZinterKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZinterWeights SCompleted

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

type SZinterWithscoresWithscores SCompleted

func (c SZinterWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZintercard SCompleted

func (c SZintercard) Numkeys(Numkeys int64) SZintercardNumkeys {
	return SZintercardNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
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
	return SZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZintercardKey) Build() SCompleted {
	return SCompleted(c)
}

type SZintercardNumkeys SCompleted

func (c SZintercardNumkeys) Key(Key ...string) SZintercardKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZinterstore SCompleted

func (c SZinterstore) Destination(Destination string) SZinterstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SZinterstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return SZinterstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SZinterstoreKey SCompleted

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
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZinterstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZinterstoreNumkeys SCompleted

func (c SZinterstoreNumkeys) Key(Key ...string) SZinterstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZinterstoreWeights SCompleted

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

type SZlexcount SCompleted

func (c SZlexcount) Key(Key string) SZlexcountKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZlexcountKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zlexcount() (c SZlexcount) {
	c.cs = append(b.get(), "ZLEXCOUNT")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZlexcountKey SCompleted

func (c SZlexcountKey) Min(Min string) SZlexcountMin {
	return SZlexcountMin{cf: c.cf, cs: append(c.cs, Min)}
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
	return SZlexcountMax{cf: c.cf, cs: append(c.cs, Max)}
}

type SZmscore SCompleted

func (c SZmscore) Key(Key string) SZmscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZmscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zmscore() (c SZmscore) {
	c.cs = append(b.get(), "ZMSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZmscoreKey SCompleted

func (c SZmscoreKey) Member(Member ...string) SZmscoreMember {
	return SZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SZmscoreMember SCompleted

func (c SZmscoreMember) Member(Member ...string) SZmscoreMember {
	return SZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
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
	return SZpopmaxKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SZpopmaxCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SZpopmaxKey) Build() SCompleted {
	return SCompleted(c)
}

type SZpopmin SCompleted

func (c SZpopmin) Key(Key string) SZpopminKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZpopminKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SZpopminCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SZpopminKey) Build() SCompleted {
	return SCompleted(c)
}

type SZrandmember SCompleted

func (c SZrandmember) Key(Key string) SZrandmemberKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrandmemberKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrandmember() (c SZrandmember) {
	c.cs = append(b.get(), "ZRANDMEMBER")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrandmemberKey SCompleted

func (c SZrandmemberKey) Count(Count int64) SZrandmemberOptionsCount {
	return SZrandmemberOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SZrandmemberKey) Build() SCompleted {
	return SCompleted(c)
}

type SZrandmemberOptionsCount SCompleted

func (c SZrandmemberOptionsCount) Withscores() SZrandmemberOptionsWithscoresWithscores {
	return SZrandmemberOptionsWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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
	return SZrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrange() (c SZrange) {
	c.cs = append(b.get(), "ZRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrangeKey SCompleted

func (c SZrangeKey) Min(Min string) SZrangeMin {
	return SZrangeMin{cf: c.cf, cs: append(c.cs, Min)}
}

type SZrangeLimit SCompleted

func (c SZrangeLimit) Withscores() SZrangeWithscoresWithscores {
	return SZrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZrangeLimit) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangeLimit) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangeMax SCompleted

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

type SZrangeMin SCompleted

func (c SZrangeMin) Max(Max string) SZrangeMax {
	return SZrangeMax{cf: c.cf, cs: append(c.cs, Max)}
}

type SZrangeRevRev SCompleted

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

type SZrangeSortbyBylex SCompleted

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

type SZrangeSortbyByscore SCompleted

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
	return SZrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrangebylex() (c SZrangebylex) {
	c.cs = append(b.get(), "ZRANGEBYLEX")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrangebylexKey SCompleted

func (c SZrangebylexKey) Min(Min string) SZrangebylexMin {
	return SZrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
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
	return SZrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangebylexMax) Build() SCompleted {
	return SCompleted(c)
}

func (c SZrangebylexMax) Cache() SCacheable {
	return SCacheable(c)
}

type SZrangebylexMin SCompleted

func (c SZrangebylexMin) Max(Max string) SZrangebylexMax {
	return SZrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

type SZrangebyscore SCompleted

func (c SZrangebyscore) Key(Key string) SZrangebyscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrangebyscore() (c SZrangebyscore) {
	c.cs = append(b.get(), "ZRANGEBYSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrangebyscoreKey SCompleted

func (c SZrangebyscoreKey) Min(Min float64) SZrangebyscoreMin {
	return SZrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
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

type SZrangebyscoreMin SCompleted

func (c SZrangebyscoreMin) Max(Max float64) SZrangebyscoreMax {
	return SZrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type SZrangebyscoreWithscoresWithscores SCompleted

func (c SZrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) SZrangebyscoreLimit {
	return SZrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
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
	return SZrangestoreDst{cf: c.cf, cs: append(c.cs, Dst)}
}

func (b *SBuilder) Zrangestore() (c SZrangestore) {
	c.cs = append(b.get(), "ZRANGESTORE")
	c.ks = InitSlot
	return
}

type SZrangestoreDst SCompleted

func (c SZrangestoreDst) Src(Src string) SZrangestoreSrc {
	c.ks = checkSlot(c.ks, slot(Src))
	return SZrangestoreSrc{cf: c.cf, cs: append(c.cs, Src)}
}

type SZrangestoreLimit SCompleted

func (c SZrangestoreLimit) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreMax SCompleted

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

type SZrangestoreMin SCompleted

func (c SZrangestoreMin) Max(Max string) SZrangestoreMax {
	return SZrangestoreMax{cf: c.cf, cs: append(c.cs, Max)}
}

type SZrangestoreRevRev SCompleted

func (c SZrangestoreRevRev) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangestoreRevRev) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSortbyBylex SCompleted

func (c SZrangestoreSortbyBylex) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangestoreSortbyBylex) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangestoreSortbyBylex) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSortbyByscore SCompleted

func (c SZrangestoreSortbyByscore) Rev() SZrangestoreRevRev {
	return SZrangestoreRevRev{cf: c.cf, cs: append(c.cs, "REV")}
}

func (c SZrangestoreSortbyByscore) Limit(Offset int64, Count int64) SZrangestoreLimit {
	return SZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SZrangestoreSortbyByscore) Build() SCompleted {
	return SCompleted(c)
}

type SZrangestoreSrc SCompleted

func (c SZrangestoreSrc) Min(Min string) SZrangestoreMin {
	return SZrangestoreMin{cf: c.cf, cs: append(c.cs, Min)}
}

type SZrank SCompleted

func (c SZrank) Key(Key string) SZrankKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrank() (c SZrank) {
	c.cs = append(b.get(), "ZRANK")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrankKey SCompleted

func (c SZrankKey) Member(Member string) SZrankMember {
	return SZrankMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return SZremKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrem() (c SZrem) {
	c.cs = append(b.get(), "ZREM")
	c.ks = InitSlot
	return
}

type SZremKey SCompleted

func (c SZremKey) Member(Member ...string) SZremMember {
	return SZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SZremMember SCompleted

func (c SZremMember) Member(Member ...string) SZremMember {
	return SZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SZremMember) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebylex SCompleted

func (c SZremrangebylex) Key(Key string) SZremrangebylexKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZremrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zremrangebylex() (c SZremrangebylex) {
	c.cs = append(b.get(), "ZREMRANGEBYLEX")
	c.ks = InitSlot
	return
}

type SZremrangebylexKey SCompleted

func (c SZremrangebylexKey) Min(Min string) SZremrangebylexMin {
	return SZremrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

type SZremrangebylexMax SCompleted

func (c SZremrangebylexMax) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebylexMin SCompleted

func (c SZremrangebylexMin) Max(Max string) SZremrangebylexMax {
	return SZremrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

type SZremrangebyrank SCompleted

func (c SZremrangebyrank) Key(Key string) SZremrangebyrankKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZremrangebyrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zremrangebyrank() (c SZremrangebyrank) {
	c.cs = append(b.get(), "ZREMRANGEBYRANK")
	c.ks = InitSlot
	return
}

type SZremrangebyrankKey SCompleted

func (c SZremrangebyrankKey) Start(Start int64) SZremrangebyrankStart {
	return SZremrangebyrankStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type SZremrangebyrankStart SCompleted

func (c SZremrangebyrankStart) Stop(Stop int64) SZremrangebyrankStop {
	return SZremrangebyrankStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type SZremrangebyrankStop SCompleted

func (c SZremrangebyrankStop) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebyscore SCompleted

func (c SZremrangebyscore) Key(Key string) SZremrangebyscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZremrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zremrangebyscore() (c SZremrangebyscore) {
	c.cs = append(b.get(), "ZREMRANGEBYSCORE")
	c.ks = InitSlot
	return
}

type SZremrangebyscoreKey SCompleted

func (c SZremrangebyscoreKey) Min(Min float64) SZremrangebyscoreMin {
	return SZremrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type SZremrangebyscoreMax SCompleted

func (c SZremrangebyscoreMax) Build() SCompleted {
	return SCompleted(c)
}

type SZremrangebyscoreMin SCompleted

func (c SZremrangebyscoreMin) Max(Max float64) SZremrangebyscoreMax {
	return SZremrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type SZrevrange SCompleted

func (c SZrevrange) Key(Key string) SZrevrangeKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZrevrangeKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrange() (c SZrevrange) {
	c.cs = append(b.get(), "ZREVRANGE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrangeKey SCompleted

func (c SZrevrangeKey) Start(Start int64) SZrevrangeStart {
	return SZrevrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type SZrevrangeStart SCompleted

func (c SZrevrangeStart) Stop(Stop int64) SZrevrangeStop {
	return SZrevrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type SZrevrangeStop SCompleted

func (c SZrevrangeStop) Withscores() SZrevrangeWithscoresWithscores {
	return SZrevrangeWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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
	return SZrevrangebylexKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrangebylex() (c SZrevrangebylex) {
	c.cs = append(b.get(), "ZREVRANGEBYLEX")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrangebylexKey SCompleted

func (c SZrevrangebylexKey) Max(Max string) SZrevrangebylexMax {
	return SZrevrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
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
	return SZrevrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

type SZrevrangebylexMin SCompleted

func (c SZrevrangebylexMin) Limit(Offset int64, Count int64) SZrevrangebylexLimit {
	return SZrevrangebylexLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
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
	return SZrevrangebyscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrangebyscore() (c SZrevrangebyscore) {
	c.cs = append(b.get(), "ZREVRANGEBYSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrangebyscoreKey SCompleted

func (c SZrevrangebyscoreKey) Max(Max float64) SZrevrangebyscoreMax {
	return SZrevrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
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
	return SZrevrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type SZrevrangebyscoreMin SCompleted

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

type SZrevrangebyscoreWithscoresWithscores SCompleted

func (c SZrevrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) SZrevrangebyscoreLimit {
	return SZrevrangebyscoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
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
	return SZrevrankKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zrevrank() (c SZrevrank) {
	c.cs = append(b.get(), "ZREVRANK")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZrevrankKey SCompleted

func (c SZrevrankKey) Member(Member string) SZrevrankMember {
	return SZrevrankMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return SZscanKey{cf: c.cf, cs: append(c.cs, Key)}
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
	return SZscanMatch{cf: c.cf, cs: append(c.cs, "MATCH", Pattern)}
}

func (c SZscanCursor) Count(Count int64) SZscanCount {
	return SZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SZscanCursor) Build() SCompleted {
	return SCompleted(c)
}

type SZscanKey SCompleted

func (c SZscanKey) Cursor(Cursor int64) SZscanCursor {
	return SZscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SZscanMatch SCompleted

func (c SZscanMatch) Count(Count int64) SZscanCount {
	return SZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SZscanMatch) Build() SCompleted {
	return SCompleted(c)
}

type SZscore SCompleted

func (c SZscore) Key(Key string) SZscoreKey {
	c.ks = checkSlot(c.ks, slot(Key))
	return SZscoreKey{cf: c.cf, cs: append(c.cs, Key)}
}

func (b *SBuilder) Zscore() (c SZscore) {
	c.cs = append(b.get(), "ZSCORE")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZscoreKey SCompleted

func (c SZscoreKey) Member(Member string) SZscoreMember {
	return SZscoreMember{cf: c.cf, cs: append(c.cs, Member)}
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
	return SZunionNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *SBuilder) Zunion() (c SZunion) {
	c.cs = append(b.get(), "ZUNION")
	c.cf = readonly
	c.ks = InitSlot
	return
}

type SZunionAggregateMax SCompleted

func (c SZunionAggregateMax) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZunionAggregateMax) Build() SCompleted {
	return SCompleted(c)
}

type SZunionAggregateMin SCompleted

func (c SZunionAggregateMin) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c SZunionAggregateMin) Build() SCompleted {
	return SCompleted(c)
}

type SZunionAggregateSum SCompleted

func (c SZunionAggregateSum) Withscores() SZunionWithscoresWithscores {
	return SZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
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
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZunionKey) Build() SCompleted {
	return SCompleted(c)
}

type SZunionNumkeys SCompleted

func (c SZunionNumkeys) Key(Key ...string) SZunionKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZunionWeights SCompleted

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

type SZunionWithscoresWithscores SCompleted

func (c SZunionWithscoresWithscores) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstore SCompleted

func (c SZunionstore) Destination(Destination string) SZunionstoreDestination {
	c.ks = checkSlot(c.ks, slot(Destination))
	return SZunionstoreDestination{cf: c.cf, cs: append(c.cs, Destination)}
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
	return SZunionstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type SZunionstoreKey SCompleted

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
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SZunionstoreKey) Build() SCompleted {
	return SCompleted(c)
}

type SZunionstoreNumkeys SCompleted

func (c SZunionstoreNumkeys) Key(Key ...string) SZunionstoreKey {
	for _, k := range Key {
		c.ks = checkSlot(c.ks, slot(k))
	}
	return SZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SZunionstoreWeights SCompleted

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

