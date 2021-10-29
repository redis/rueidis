package cmds

import "strconv"

type AclCat struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c AclCatCategoryname) Build() Completed {
	return Completed(c)
}

type AclDeluser struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c AclDeluserUsername) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cf: c.cf, cs: append(c.cs, Username...)}
}

func (c AclDeluserUsername) Build() Completed {
	return Completed(c)
}

type AclGenpass struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c AclGenpassBits) Build() Completed {
	return Completed(c)
}

type AclGetuser struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c AclGetuserUsername) Build() Completed {
	return Completed(c)
}

type AclHelp struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c AclLogCountOrReset) Build() Completed {
	return Completed(c)
}

type AclSave struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c AclSetuserRule) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c AclSetuserRule) Build() Completed {
	return Completed(c)
}

type AclSetuserUsername struct {
	cs []string
	cf uint32
}

func (c AclSetuserUsername) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cf: c.cf, cs: append(c.cs, Rule...)}
}

func (c AclSetuserUsername) Build() Completed {
	return Completed(c)
}

type AclUsers struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c AppendKey) Value(Value string) AppendValue {
	return AppendValue{cf: c.cf, cs: append(c.cs, Value)}
}

type AppendValue struct {
	cs []string
	cf uint32
}

func (c AppendValue) Build() Completed {
	return Completed(c)
}

type Asking struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c AuthPassword) Build() Completed {
	return Completed(c)
}

type AuthUsername struct {
	cs []string
	cf uint32
}

func (c AuthUsername) Password(Password string) AuthPassword {
	return AuthPassword{cf: c.cf, cs: append(c.cs, Password)}
}

type Bgrewriteaof struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c BgsaveScheduleSchedule) Build() Completed {
	return Completed(c)
}

type Bitcount struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c BitcountStartEnd) Build() Completed {
	return Completed(c)
}

func (c BitcountStartEnd) Cache() Cacheable {
	return Cacheable(c)
}

type Bitfield struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c BitfieldFail) Build() Completed {
	return Completed(c)
}

type BitfieldGet struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c BitfieldRoGet) Build() Completed {
	return Completed(c)
}

func (c BitfieldRoGet) Cache() Cacheable {
	return Cacheable(c)
}

type BitfieldRoKey struct {
	cs []string
	cf uint32
}

func (c BitfieldRoKey) Get(Type string, Offset int64) BitfieldRoGet {
	return BitfieldRoGet{cf: c.cf, cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

func (c BitfieldRoKey) Cache() Cacheable {
	return Cacheable(c)
}

type BitfieldSat struct {
	cs []string
	cf uint32
}

func (c BitfieldSat) Build() Completed {
	return Completed(c)
}

type BitfieldSet struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c BitfieldWrap) Build() Completed {
	return Completed(c)
}

type Bitop struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c BitopDestkey) Key(Key ...string) BitopKey {
	return BitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BitopKey struct {
	cs []string
	cf uint32
}

func (c BitopKey) Key(Key ...string) BitopKey {
	return BitopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c BitopKey) Build() Completed {
	return Completed(c)
}

type BitopOperation struct {
	cs []string
	cf uint32
}

func (c BitopOperation) Destkey(Destkey string) BitopDestkey {
	return BitopDestkey{cf: c.cf, cs: append(c.cs, Destkey)}
}

type Bitpos struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c BitposBit) Start(Start int64) BitposIndexStart {
	return BitposIndexStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c BitposBit) Cache() Cacheable {
	return Cacheable(c)
}

type BitposIndexEnd struct {
	cs []string
	cf uint32
}

func (c BitposIndexEnd) Build() Completed {
	return Completed(c)
}

func (c BitposIndexEnd) Cache() Cacheable {
	return Cacheable(c)
}

type BitposIndexStart struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c BitposKey) Bit(Bit int64) BitposBit {
	return BitposBit{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Bit, 10))}
}

func (c BitposKey) Cache() Cacheable {
	return Cacheable(c)
}

type Blmove struct {
	cs []string
	cf uint32
}

func (c Blmove) Source(Source string) BlmoveSource {
	return BlmoveSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Blmove() (c Blmove) {
	c.cf = blockTag
	c.cs = append(b.get(), "BLMOVE")
	return
}

type BlmoveDestination struct {
	cs []string
	cf uint32
}

func (c BlmoveDestination) Left() BlmoveWherefromLeft {
	return BlmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveDestination) Right() BlmoveWherefromRight {
	return BlmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveSource struct {
	cs []string
	cf uint32
}

func (c BlmoveSource) Destination(Destination string) BlmoveDestination {
	return BlmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type BlmoveTimeout struct {
	cs []string
	cf uint32
}

func (c BlmoveTimeout) Build() Completed {
	return Completed(c)
}

type BlmoveWherefromLeft struct {
	cs []string
	cf uint32
}

func (c BlmoveWherefromLeft) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromLeft) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveWherefromRight struct {
	cs []string
	cf uint32
}

func (c BlmoveWherefromRight) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromRight) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type BlmoveWheretoLeft struct {
	cs []string
	cf uint32
}

func (c BlmoveWheretoLeft) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BlmoveWheretoRight struct {
	cs []string
	cf uint32
}

func (c BlmoveWheretoRight) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type Blmpop struct {
	cs []string
	cf uint32
}

func (c Blmpop) Timeout(Timeout float64) BlmpopTimeout {
	return BlmpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (b *Builder) Blmpop() (c Blmpop) {
	c.cf = blockTag
	c.cs = append(b.get(), "BLMPOP")
	return
}

type BlmpopCount struct {
	cs []string
	cf uint32
}

func (c BlmpopCount) Build() Completed {
	return Completed(c)
}

type BlmpopKey struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c BlmpopTimeout) Numkeys(Numkeys int64) BlmpopNumkeys {
	return BlmpopNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type BlmpopWhereLeft struct {
	cs []string
	cf uint32
}

func (c BlmpopWhereLeft) Count(Count int64) BlmpopCount {
	return BlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type BlmpopWhereRight struct {
	cs []string
	cf uint32
}

func (c BlmpopWhereRight) Count(Count int64) BlmpopCount {
	return BlmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Blpop struct {
	cs []string
	cf uint32
}

func (c Blpop) Key(Key ...string) BlpopKey {
	return BlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Blpop() (c Blpop) {
	c.cf = blockTag
	c.cs = append(b.get(), "BLPOP")
	return
}

type BlpopKey struct {
	cs []string
	cf uint32
}

func (c BlpopKey) Timeout(Timeout float64) BlpopTimeout {
	return BlpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BlpopKey) Key(Key ...string) BlpopKey {
	return BlpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BlpopTimeout struct {
	cs []string
	cf uint32
}

func (c BlpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpop struct {
	cs []string
	cf uint32
}

func (c Brpop) Key(Key ...string) BrpopKey {
	return BrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Brpop() (c Brpop) {
	c.cf = blockTag
	c.cs = append(b.get(), "BRPOP")
	return
}

type BrpopKey struct {
	cs []string
	cf uint32
}

func (c BrpopKey) Timeout(Timeout float64) BrpopTimeout {
	return BrpopTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BrpopKey) Key(Key ...string) BrpopKey {
	return BrpopKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BrpopTimeout struct {
	cs []string
	cf uint32
}

func (c BrpopTimeout) Build() Completed {
	return Completed(c)
}

type Brpoplpush struct {
	cs []string
	cf uint32
}

func (c Brpoplpush) Source(Source string) BrpoplpushSource {
	return BrpoplpushSource{cf: c.cf, cs: append(c.cs, Source)}
}

func (b *Builder) Brpoplpush() (c Brpoplpush) {
	c.cf = blockTag
	c.cs = append(b.get(), "BRPOPLPUSH")
	return
}

type BrpoplpushDestination struct {
	cs []string
	cf uint32
}

func (c BrpoplpushDestination) Timeout(Timeout float64) BrpoplpushTimeout {
	return BrpoplpushTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BrpoplpushSource struct {
	cs []string
	cf uint32
}

func (c BrpoplpushSource) Destination(Destination string) BrpoplpushDestination {
	return BrpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type BrpoplpushTimeout struct {
	cs []string
	cf uint32
}

func (c BrpoplpushTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmax struct {
	cs []string
	cf uint32
}

func (c Bzpopmax) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmax() (c Bzpopmax) {
	c.cf = blockTag
	c.cs = append(b.get(), "BZPOPMAX")
	return
}

type BzpopmaxKey struct {
	cs []string
	cf uint32
}

func (c BzpopmaxKey) Timeout(Timeout float64) BzpopmaxTimeout {
	return BzpopmaxTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopmaxKey) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BzpopmaxTimeout struct {
	cs []string
	cf uint32
}

func (c BzpopmaxTimeout) Build() Completed {
	return Completed(c)
}

type Bzpopmin struct {
	cs []string
	cf uint32
}

func (c Bzpopmin) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmin() (c Bzpopmin) {
	c.cf = blockTag
	c.cs = append(b.get(), "BZPOPMIN")
	return
}

type BzpopminKey struct {
	cs []string
	cf uint32
}

func (c BzpopminKey) Timeout(Timeout float64) BzpopminTimeout {
	return BzpopminTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopminKey) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type BzpopminTimeout struct {
	cs []string
	cf uint32
}

func (c BzpopminTimeout) Build() Completed {
	return Completed(c)
}

type ClientCaching struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClientCachingModeNo) Build() Completed {
	return Completed(c)
}

type ClientCachingModeYes struct {
	cs []string
	cf uint32
}

func (c ClientCachingModeYes) Build() Completed {
	return Completed(c)
}

type ClientGetname struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClientKillLaddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cf: c.cf, cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillLaddr) Build() Completed {
	return Completed(c)
}

type ClientKillMaster struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClientKillSkipme) Build() Completed {
	return Completed(c)
}

type ClientKillSlave struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClientListIdId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cf: c.cf, cs: c.cs}
}

type ClientListMaster struct {
	cs []string
	cf uint32
}

func (c ClientListMaster) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientListNormal struct {
	cs []string
	cf uint32
}

func (c ClientListNormal) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientListPubsub struct {
	cs []string
	cf uint32
}

func (c ClientListPubsub) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientListReplica struct {
	cs []string
	cf uint32
}

func (c ClientListReplica) Id() ClientListIdId {
	return ClientListIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

type ClientNoEvict struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClientNoEvictEnabledOff) Build() Completed {
	return Completed(c)
}

type ClientNoEvictEnabledOn struct {
	cs []string
	cf uint32
}

func (c ClientNoEvictEnabledOn) Build() Completed {
	return Completed(c)
}

type ClientPause struct {
	cs []string
	cf uint32
}

func (c ClientPause) Timeout(Timeout int64) ClientPauseTimeout {
	return ClientPauseTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

func (b *Builder) ClientPause() (c ClientPause) {
	c.cf = blockTag
	c.cs = append(b.get(), "CLIENT", "PAUSE")
	return
}

type ClientPauseModeAll struct {
	cs []string
	cf uint32
}

func (c ClientPauseModeAll) Build() Completed {
	return Completed(c)
}

type ClientPauseModeWrite struct {
	cs []string
	cf uint32
}

func (c ClientPauseModeWrite) Build() Completed {
	return Completed(c)
}

type ClientPauseTimeout struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClientReplyReplyModeOff) Build() Completed {
	return Completed(c)
}

type ClientReplyReplyModeOn struct {
	cs []string
	cf uint32
}

func (c ClientReplyReplyModeOn) Build() Completed {
	return Completed(c)
}

type ClientReplyReplyModeSkip struct {
	cs []string
	cf uint32
}

func (c ClientReplyReplyModeSkip) Build() Completed {
	return Completed(c)
}

type ClientSetname struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClientSetnameConnectionName) Build() Completed {
	return Completed(c)
}

type ClientTracking struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClientTrackingNoloopNoloop) Build() Completed {
	return Completed(c)
}

type ClientTrackingOptinOptin struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClientTrackingOptoutOptout) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cf: c.cf, cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingOptoutOptout) Build() Completed {
	return Completed(c)
}

type ClientTrackingPrefix struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClientUnblockUnblockTypeError) Build() Completed {
	return Completed(c)
}

type ClientUnblockUnblockTypeTimeout struct {
	cs []string
	cf uint32
}

func (c ClientUnblockUnblockTypeTimeout) Build() Completed {
	return Completed(c)
}

type ClientUnpause struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClusterCountFailureReportsNodeId) Build() Completed {
	return Completed(c)
}

type ClusterCountkeysinslot struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterCountkeysinslotSlot) Build() Completed {
	return Completed(c)
}

type ClusterDelslots struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClusterFailoverOptionsForce) Build() Completed {
	return Completed(c)
}

type ClusterFailoverOptionsTakeover struct {
	cs []string
	cf uint32
}

func (c ClusterFailoverOptionsTakeover) Build() Completed {
	return Completed(c)
}

type ClusterFlushslots struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClusterForgetNodeId) Build() Completed {
	return Completed(c)
}

type ClusterGetkeysinslot struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterGetkeysinslotCount) Build() Completed {
	return Completed(c)
}

type ClusterGetkeysinslotSlot struct {
	cs []string
	cf uint32
}

func (c ClusterGetkeysinslotSlot) Count(Count int64) ClusterGetkeysinslotCount {
	return ClusterGetkeysinslotCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type ClusterInfo struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClusterKeyslotKey) Build() Completed {
	return Completed(c)
}

type ClusterMeet struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterMeetIp) Port(Port int64) ClusterMeetPort {
	return ClusterMeetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type ClusterMeetPort struct {
	cs []string
	cf uint32
}

func (c ClusterMeetPort) Build() Completed {
	return Completed(c)
}

type ClusterMyid struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClusterReplicasNodeId) Build() Completed {
	return Completed(c)
}

type ClusterReplicate struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterReplicateNodeId) Build() Completed {
	return Completed(c)
}

type ClusterReset struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterResetResetTypeHard) Build() Completed {
	return Completed(c)
}

type ClusterResetResetTypeSoft struct {
	cs []string
	cf uint32
}

func (c ClusterResetResetTypeSoft) Build() Completed {
	return Completed(c)
}

type ClusterSaveconfig struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ClusterSetConfigEpochConfigEpoch) Build() Completed {
	return Completed(c)
}

type ClusterSetslot struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterSetslotNodeId) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSlot struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterSetslotSubcommandImporting) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandImporting) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandMigrating struct {
	cs []string
	cf uint32
}

func (c ClusterSetslotSubcommandMigrating) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandMigrating) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandNode struct {
	cs []string
	cf uint32
}

func (c ClusterSetslotSubcommandNode) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandNode) Build() Completed {
	return Completed(c)
}

type ClusterSetslotSubcommandStable struct {
	cs []string
	cf uint32
}

func (c ClusterSetslotSubcommandStable) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cf: c.cf, cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandStable) Build() Completed {
	return Completed(c)
}

type ClusterSlaves struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ClusterSlavesNodeId) Build() Completed {
	return Completed(c)
}

type ClusterSlots struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c CommandInfoCommandName) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cf: c.cf, cs: append(c.cs, CommandName...)}
}

func (c CommandInfoCommandName) Build() Completed {
	return Completed(c)
}

type ConfigGet struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ConfigGetParameter) Build() Completed {
	return Completed(c)
}

type ConfigResetstat struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ConfigSetParameter) Value(Value string) ConfigSetValue {
	return ConfigSetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type ConfigSetValue struct {
	cs []string
	cf uint32
}

func (c ConfigSetValue) Build() Completed {
	return Completed(c)
}

type Copy struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c CopyDb) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cf: c.cf, cs: append(c.cs, "REPLACE")}
}

func (c CopyDb) Build() Completed {
	return Completed(c)
}

type CopyDestination struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c CopyReplaceReplace) Build() Completed {
	return Completed(c)
}

type CopySource struct {
	cs []string
	cf uint32
}

func (c CopySource) Destination(Destination string) CopyDestination {
	return CopyDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Dbsize struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c DebugObjectKey) Build() Completed {
	return Completed(c)
}

type DebugSegfault struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c DecrKey) Build() Completed {
	return Completed(c)
}

type Decrby struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c DecrbyDecrement) Build() Completed {
	return Completed(c)
}

type DecrbyKey struct {
	cs []string
	cf uint32
}

func (c DecrbyKey) Decrement(Decrement int64) DecrbyDecrement {
	return DecrbyDecrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Decrement, 10))}
}

type Del struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c DelKey) Key(Key ...string) DelKey {
	return DelKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c DelKey) Build() Completed {
	return Completed(c)
}

type Discard struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c DumpKey) Build() Completed {
	return Completed(c)
}

type Echo struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c EchoMessage) Build() Completed {
	return Completed(c)
}

type Eval struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c EvalArg) Arg(Arg ...string) EvalArg {
	return EvalArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalArg) Build() Completed {
	return Completed(c)
}

type EvalKey struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c EvalRoArg) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalRoArg) Build() Completed {
	return Completed(c)
}

type EvalRoKey struct {
	cs []string
	cf uint32
}

func (c EvalRoKey) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalRoKey) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalRoNumkeys struct {
	cs []string
	cf uint32
}

func (c EvalRoNumkeys) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalRoScript struct {
	cs []string
	cf uint32
}

func (c EvalRoScript) Numkeys(Numkeys int64) EvalRoNumkeys {
	return EvalRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalScript struct {
	cs []string
	cf uint32
}

func (c EvalScript) Numkeys(Numkeys int64) EvalNumkeys {
	return EvalNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type Evalsha struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c EvalshaArg) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaArg) Build() Completed {
	return Completed(c)
}

type EvalshaKey struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c EvalshaRoArg) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaRoArg) Build() Completed {
	return Completed(c)
}

type EvalshaRoKey struct {
	cs []string
	cf uint32
}

func (c EvalshaRoKey) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c EvalshaRoKey) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalshaRoNumkeys struct {
	cs []string
	cf uint32
}

func (c EvalshaRoNumkeys) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type EvalshaRoSha1 struct {
	cs []string
	cf uint32
}

func (c EvalshaRoSha1) Numkeys(Numkeys int64) EvalshaRoNumkeys {
	return EvalshaRoNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalshaSha1 struct {
	cs []string
	cf uint32
}

func (c EvalshaSha1) Numkeys(Numkeys int64) EvalshaNumkeys {
	return EvalshaNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type Exec struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ExistsKey) Key(Key ...string) ExistsKey {
	return ExistsKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ExistsKey) Build() Completed {
	return Completed(c)
}

type Expire struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ExpireConditionGt) Build() Completed {
	return Completed(c)
}

type ExpireConditionLt struct {
	cs []string
	cf uint32
}

func (c ExpireConditionLt) Build() Completed {
	return Completed(c)
}

type ExpireConditionNx struct {
	cs []string
	cf uint32
}

func (c ExpireConditionNx) Build() Completed {
	return Completed(c)
}

type ExpireConditionXx struct {
	cs []string
	cf uint32
}

func (c ExpireConditionXx) Build() Completed {
	return Completed(c)
}

type ExpireKey struct {
	cs []string
	cf uint32
}

func (c ExpireKey) Seconds(Seconds int64) ExpireSeconds {
	return ExpireSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type ExpireSeconds struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ExpireatConditionGt) Build() Completed {
	return Completed(c)
}

type ExpireatConditionLt struct {
	cs []string
	cf uint32
}

func (c ExpireatConditionLt) Build() Completed {
	return Completed(c)
}

type ExpireatConditionNx struct {
	cs []string
	cf uint32
}

func (c ExpireatConditionNx) Build() Completed {
	return Completed(c)
}

type ExpireatConditionXx struct {
	cs []string
	cf uint32
}

func (c ExpireatConditionXx) Build() Completed {
	return Completed(c)
}

type ExpireatKey struct {
	cs []string
	cf uint32
}

func (c ExpireatKey) Timestamp(Timestamp int64) ExpireatTimestamp {
	return ExpireatTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timestamp, 10))}
}

type ExpireatTimestamp struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ExpiretimeKey) Build() Completed {
	return Completed(c)
}

type Failover struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c FailoverAbort) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cf: c.cf, cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverAbort) Build() Completed {
	return Completed(c)
}

type FailoverTargetForce struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c FailoverTargetHost) Port(Port int64) FailoverTargetPort {
	return FailoverTargetPort{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type FailoverTargetPort struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c FailoverTargetTo) Host(Host string) FailoverTargetHost {
	return FailoverTargetHost{cf: c.cf, cs: append(c.cs, Host)}
}

type FailoverTimeout struct {
	cs []string
	cf uint32
}

func (c FailoverTimeout) Build() Completed {
	return Completed(c)
}

type Flushall struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c FlushallAsyncAsync) Build() Completed {
	return Completed(c)
}

type FlushallAsyncSync struct {
	cs []string
	cf uint32
}

func (c FlushallAsyncSync) Build() Completed {
	return Completed(c)
}

type Flushdb struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c FlushdbAsyncAsync) Build() Completed {
	return Completed(c)
}

type FlushdbAsyncSync struct {
	cs []string
	cf uint32
}

func (c FlushdbAsyncSync) Build() Completed {
	return Completed(c)
}

type Geoadd struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeoaddChangeCh) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type GeoaddConditionNx struct {
	cs []string
	cf uint32
}

func (c GeoaddConditionNx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c GeoaddConditionNx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type GeoaddConditionXx struct {
	cs []string
	cf uint32
}

func (c GeoaddConditionXx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cf: c.cf, cs: append(c.cs, "CH")}
}

func (c GeoaddConditionXx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, )}
}

type GeoaddKey struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member)}
}

func (c GeoaddLongitudeLatitudeMember) Build() Completed {
	return Completed(c)
}

type Geodist struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeodistKey) Member1(Member1 string) GeodistMember1 {
	return GeodistMember1{cf: c.cf, cs: append(c.cs, Member1)}
}

func (c GeodistKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistMember1 struct {
	cs []string
	cf uint32
}

func (c GeodistMember1) Member2(Member2 string) GeodistMember2 {
	return GeodistMember2{cf: c.cf, cs: append(c.cs, Member2)}
}

func (c GeodistMember1) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistMember2 struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeodistUnitFt) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitFt) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitKm struct {
	cs []string
	cf uint32
}

func (c GeodistUnitKm) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitKm) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitM struct {
	cs []string
	cf uint32
}

func (c GeodistUnitM) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitM) Cache() Cacheable {
	return Cacheable(c)
}

type GeodistUnitMi struct {
	cs []string
	cf uint32
}

func (c GeodistUnitMi) Build() Completed {
	return Completed(c)
}

func (c GeodistUnitMi) Cache() Cacheable {
	return Cacheable(c)
}

type Geohash struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeohashKey) Member(Member ...string) GeohashMember {
	return GeohashMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeohashKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeohashMember struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GeoposKey) Member(Member ...string) GeoposMember {
	return GeoposMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c GeoposKey) Cache() Cacheable {
	return Cacheable(c)
}

type GeoposMember struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GeoradiusKey) Longitude(Longitude float64) GeoradiusLongitude {
	return GeoradiusLongitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type GeoradiusLatitude struct {
	cs []string
	cf uint32
}

func (c GeoradiusLatitude) Radius(Radius float64) GeoradiusRadius {
	return GeoradiusRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusLongitude struct {
	cs []string
	cf uint32
}

func (c GeoradiusLongitude) Latitude(Latitude float64) GeoradiusLatitude {
	return GeoradiusLatitude{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type GeoradiusOrderAsc struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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

type GeoradiusStore struct {
	cs []string
	cf uint32
}

func (c GeoradiusStore) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusStore) Build() Completed {
	return Completed(c)
}

type GeoradiusStoredist struct {
	cs []string
	cf uint32
}

func (c GeoradiusStoredist) Build() Completed {
	return Completed(c)
}

type GeoradiusUnitFt struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GeoradiusbymemberKey) Member(Member string) GeoradiusbymemberMember {
	return GeoradiusbymemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

type GeoradiusbymemberMember struct {
	cs []string
	cf uint32
}

func (c GeoradiusbymemberMember) Radius(Radius float64) GeoradiusbymemberRadius {
	return GeoradiusbymemberRadius{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusbymemberOrderAsc struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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

type GeoradiusbymemberStore struct {
	cs []string
	cf uint32
}

func (c GeoradiusbymemberStore) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberStore) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberStoredist struct {
	cs []string
	cf uint32
}

func (c GeoradiusbymemberStoredist) Build() Completed {
	return Completed(c)
}

type GeoradiusbymemberUnitFt struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GeosearchBoxBybox) Height(Height float64) GeosearchBoxHeight {
	return GeosearchBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

func (c GeosearchBoxBybox) Cache() Cacheable {
	return Cacheable(c)
}

type GeosearchBoxHeight struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GeosearchWithhashWithhash) Build() Completed {
	return Completed(c)
}

func (c GeosearchWithhashWithhash) Cache() Cacheable {
	return Cacheable(c)
}

type Geosearchstore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeosearchstoreBoxBybox) Height(Height float64) GeosearchstoreBoxHeight {
	return GeosearchstoreBoxHeight{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type GeosearchstoreBoxHeight struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GeosearchstoreCountAnyAny) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreCountAnyAny) Build() Completed {
	return Completed(c)
}

type GeosearchstoreCountCount struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeosearchstoreDestination) Source(Source string) GeosearchstoreSource {
	return GeosearchstoreSource{cf: c.cf, cs: append(c.cs, Source)}
}

type GeosearchstoreFromlonlat struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GeosearchstoreOrderAsc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderAsc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreOrderDesc struct {
	cs []string
	cf uint32
}

func (c GeosearchstoreOrderDesc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderDesc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cf: c.cf, cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreSource struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GeosearchstoreStoredistStoredist) Build() Completed {
	return Completed(c)
}

type Get struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GetKey) Build() Completed {
	return Completed(c)
}

func (c GetKey) Cache() Cacheable {
	return Cacheable(c)
}

type Getbit struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GetbitKey) Offset(Offset int64) GetbitOffset {
	return GetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

func (c GetbitKey) Cache() Cacheable {
	return Cacheable(c)
}

type GetbitOffset struct {
	cs []string
	cf uint32
}

func (c GetbitOffset) Build() Completed {
	return Completed(c)
}

func (c GetbitOffset) Cache() Cacheable {
	return Cacheable(c)
}

type Getdel struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GetdelKey) Build() Completed {
	return Completed(c)
}

type Getex struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GetexExpirationEx) Build() Completed {
	return Completed(c)
}

type GetexExpirationExat struct {
	cs []string
	cf uint32
}

func (c GetexExpirationExat) Build() Completed {
	return Completed(c)
}

type GetexExpirationPersist struct {
	cs []string
	cf uint32
}

func (c GetexExpirationPersist) Build() Completed {
	return Completed(c)
}

type GetexExpirationPx struct {
	cs []string
	cf uint32
}

func (c GetexExpirationPx) Build() Completed {
	return Completed(c)
}

type GetexExpirationPxat struct {
	cs []string
	cf uint32
}

func (c GetexExpirationPxat) Build() Completed {
	return Completed(c)
}

type GetexKey struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c GetrangeEnd) Build() Completed {
	return Completed(c)
}

func (c GetrangeEnd) Cache() Cacheable {
	return Cacheable(c)
}

type GetrangeKey struct {
	cs []string
	cf uint32
}

func (c GetrangeKey) Start(Start int64) GetrangeStart {
	return GetrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c GetrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type GetrangeStart struct {
	cs []string
	cf uint32
}

func (c GetrangeStart) End(End int64) GetrangeEnd {
	return GetrangeEnd{cf: c.cf, cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c GetrangeStart) Cache() Cacheable {
	return Cacheable(c)
}

type Getset struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c GetsetKey) Value(Value string) GetsetValue {
	return GetsetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type GetsetValue struct {
	cs []string
	cf uint32
}

func (c GetsetValue) Build() Completed {
	return Completed(c)
}

type Hdel struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HdelField) Field(Field ...string) HdelField {
	return HdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c HdelField) Build() Completed {
	return Completed(c)
}

type HdelKey struct {
	cs []string
	cf uint32
}

func (c HdelKey) Field(Field ...string) HdelField {
	return HdelField{cf: c.cf, cs: append(c.cs, Field...)}
}

type Hello struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HelloArgumentsAuth) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cf: c.cf, cs: append(c.cs, "SETNAME", Clientname)}
}

func (c HelloArgumentsAuth) Build() Completed {
	return Completed(c)
}

type HelloArgumentsProtover struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HelloArgumentsSetname) Build() Completed {
	return Completed(c)
}

type Hexists struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HexistsField) Build() Completed {
	return Completed(c)
}

func (c HexistsField) Cache() Cacheable {
	return Cacheable(c)
}

type HexistsKey struct {
	cs []string
	cf uint32
}

func (c HexistsKey) Field(Field string) HexistsField {
	return HexistsField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c HexistsKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hget struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HgetField) Build() Completed {
	return Completed(c)
}

func (c HgetField) Cache() Cacheable {
	return Cacheable(c)
}

type HgetKey struct {
	cs []string
	cf uint32
}

func (c HgetKey) Field(Field string) HgetField {
	return HgetField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c HgetKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hgetall struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HgetallKey) Build() Completed {
	return Completed(c)
}

func (c HgetallKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hincrby struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HincrbyField) Increment(Increment int64) HincrbyIncrement {
	return HincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type HincrbyIncrement struct {
	cs []string
	cf uint32
}

func (c HincrbyIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyKey struct {
	cs []string
	cf uint32
}

func (c HincrbyKey) Field(Field string) HincrbyField {
	return HincrbyField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hincrbyfloat struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HincrbyfloatField) Increment(Increment float64) HincrbyfloatIncrement {
	return HincrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type HincrbyfloatIncrement struct {
	cs []string
	cf uint32
}

func (c HincrbyfloatIncrement) Build() Completed {
	return Completed(c)
}

type HincrbyfloatKey struct {
	cs []string
	cf uint32
}

func (c HincrbyfloatKey) Field(Field string) HincrbyfloatField {
	return HincrbyfloatField{cf: c.cf, cs: append(c.cs, Field)}
}

type Hkeys struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HkeysKey) Build() Completed {
	return Completed(c)
}

func (c HkeysKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hlen struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HlenKey) Build() Completed {
	return Completed(c)
}

func (c HlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hmget struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c HmgetKey) Field(Field ...string) HmgetField {
	return HmgetField{cf: c.cf, cs: append(c.cs, Field...)}
}

func (c HmgetKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hmset struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HmsetFieldValue) FieldValue(Field string, Value string) HmsetFieldValue {
	return HmsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c HmsetFieldValue) Build() Completed {
	return Completed(c)
}

type HmsetKey struct {
	cs []string
	cf uint32
}

func (c HmsetKey) FieldValue() HmsetFieldValue {
	return HmsetFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type Hrandfield struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HrandfieldKey) Count(Count int64) HrandfieldOptionsCount {
	return HrandfieldOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type HrandfieldOptionsCount struct {
	cs []string
	cf uint32
}

func (c HrandfieldOptionsCount) Withvalues() HrandfieldOptionsWithvaluesWithvalues {
	return HrandfieldOptionsWithvaluesWithvalues{cf: c.cf, cs: append(c.cs, "WITHVALUES")}
}

func (c HrandfieldOptionsCount) Build() Completed {
	return Completed(c)
}

type HrandfieldOptionsWithvaluesWithvalues struct {
	cs []string
	cf uint32
}

func (c HrandfieldOptionsWithvaluesWithvalues) Build() Completed {
	return Completed(c)
}

type Hscan struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HscanCount) Build() Completed {
	return Completed(c)
}

type HscanCursor struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HscanKey) Cursor(Cursor int64) HscanCursor {
	return HscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type HscanMatch struct {
	cs []string
	cf uint32
}

func (c HscanMatch) Count(Count int64) HscanCount {
	return HscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c HscanMatch) Build() Completed {
	return Completed(c)
}

type Hset struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HsetFieldValue) FieldValue(Field string, Value string) HsetFieldValue {
	return HsetFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c HsetFieldValue) Build() Completed {
	return Completed(c)
}

type HsetKey struct {
	cs []string
	cf uint32
}

func (c HsetKey) FieldValue() HsetFieldValue {
	return HsetFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type Hsetnx struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HsetnxField) Value(Value string) HsetnxValue {
	return HsetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type HsetnxKey struct {
	cs []string
	cf uint32
}

func (c HsetnxKey) Field(Field string) HsetnxField {
	return HsetnxField{cf: c.cf, cs: append(c.cs, Field)}
}

type HsetnxValue struct {
	cs []string
	cf uint32
}

func (c HsetnxValue) Build() Completed {
	return Completed(c)
}

type Hstrlen struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HstrlenField) Build() Completed {
	return Completed(c)
}

func (c HstrlenField) Cache() Cacheable {
	return Cacheable(c)
}

type HstrlenKey struct {
	cs []string
	cf uint32
}

func (c HstrlenKey) Field(Field string) HstrlenField {
	return HstrlenField{cf: c.cf, cs: append(c.cs, Field)}
}

func (c HstrlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Hvals struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c HvalsKey) Build() Completed {
	return Completed(c)
}

func (c HvalsKey) Cache() Cacheable {
	return Cacheable(c)
}

type Incr struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c IncrKey) Build() Completed {
	return Completed(c)
}

type Incrby struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c IncrbyIncrement) Build() Completed {
	return Completed(c)
}

type IncrbyKey struct {
	cs []string
	cf uint32
}

func (c IncrbyKey) Increment(Increment int64) IncrbyIncrement {
	return IncrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type Incrbyfloat struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c IncrbyfloatIncrement) Build() Completed {
	return Completed(c)
}

type IncrbyfloatKey struct {
	cs []string
	cf uint32
}

func (c IncrbyfloatKey) Increment(Increment float64) IncrbyfloatIncrement {
	return IncrbyfloatIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type Info struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c InfoSection) Build() Completed {
	return Completed(c)
}

type Keys struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c KeysPattern) Build() Completed {
	return Completed(c)
}

type Lastsave struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c LatencyGraphEvent) Build() Completed {
	return Completed(c)
}

type LatencyHelp struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c LatencyHistoryEvent) Build() Completed {
	return Completed(c)
}

type LatencyLatest struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c LatencyResetEvent) Event(Event ...string) LatencyResetEvent {
	return LatencyResetEvent{cf: c.cf, cs: append(c.cs, Event...)}
}

func (c LatencyResetEvent) Build() Completed {
	return Completed(c)
}

type Lindex struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LindexIndex) Build() Completed {
	return Completed(c)
}

func (c LindexIndex) Cache() Cacheable {
	return Cacheable(c)
}

type LindexKey struct {
	cs []string
	cf uint32
}

func (c LindexKey) Index(Index int64) LindexIndex {
	return LindexIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

func (c LindexKey) Cache() Cacheable {
	return Cacheable(c)
}

type Linsert struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LinsertElement) Build() Completed {
	return Completed(c)
}

type LinsertKey struct {
	cs []string
	cf uint32
}

func (c LinsertKey) Before() LinsertWhereBefore {
	return LinsertWhereBefore{cf: c.cf, cs: append(c.cs, "BEFORE")}
}

func (c LinsertKey) After() LinsertWhereAfter {
	return LinsertWhereAfter{cf: c.cf, cs: append(c.cs, "AFTER")}
}

type LinsertPivot struct {
	cs []string
	cf uint32
}

func (c LinsertPivot) Element(Element string) LinsertElement {
	return LinsertElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LinsertWhereAfter struct {
	cs []string
	cf uint32
}

func (c LinsertWhereAfter) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type LinsertWhereBefore struct {
	cs []string
	cf uint32
}

func (c LinsertWhereBefore) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cf: c.cf, cs: append(c.cs, Pivot)}
}

type Llen struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LlenKey) Build() Completed {
	return Completed(c)
}

func (c LlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Lmove struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LmoveDestination) Left() LmoveWherefromLeft {
	return LmoveWherefromLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveDestination) Right() LmoveWherefromRight {
	return LmoveWherefromRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveSource struct {
	cs []string
	cf uint32
}

func (c LmoveSource) Destination(Destination string) LmoveDestination {
	return LmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type LmoveWherefromLeft struct {
	cs []string
	cf uint32
}

func (c LmoveWherefromLeft) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromLeft) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveWherefromRight struct {
	cs []string
	cf uint32
}

func (c LmoveWherefromRight) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cf: c.cf, cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromRight) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cf: c.cf, cs: append(c.cs, "RIGHT")}
}

type LmoveWheretoLeft struct {
	cs []string
	cf uint32
}

func (c LmoveWheretoLeft) Build() Completed {
	return Completed(c)
}

type LmoveWheretoRight struct {
	cs []string
	cf uint32
}

func (c LmoveWheretoRight) Build() Completed {
	return Completed(c)
}

type Lmpop struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LmpopCount) Build() Completed {
	return Completed(c)
}

type LmpopKey struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c LmpopWhereLeft) Count(Count int64) LmpopCount {
	return LmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereLeft) Build() Completed {
	return Completed(c)
}

type LmpopWhereRight struct {
	cs []string
	cf uint32
}

func (c LmpopWhereRight) Count(Count int64) LmpopCount {
	return LmpopCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereRight) Build() Completed {
	return Completed(c)
}

type Lolwut struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LolwutVersion) Build() Completed {
	return Completed(c)
}

type Lpop struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LpopCount) Build() Completed {
	return Completed(c)
}

type LpopKey struct {
	cs []string
	cf uint32
}

func (c LpopKey) Count(Count int64) LpopCount {
	return LpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c LpopKey) Build() Completed {
	return Completed(c)
}

type Lpos struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c LposKey) Element(Element string) LposElement {
	return LposElement{cf: c.cf, cs: append(c.cs, Element)}
}

func (c LposKey) Cache() Cacheable {
	return Cacheable(c)
}

type LposMaxlen struct {
	cs []string
	cf uint32
}

func (c LposMaxlen) Build() Completed {
	return Completed(c)
}

func (c LposMaxlen) Cache() Cacheable {
	return Cacheable(c)
}

type LposRank struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c LpushElement) Element(Element ...string) LpushElement {
	return LpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c LpushElement) Build() Completed {
	return Completed(c)
}

type LpushKey struct {
	cs []string
	cf uint32
}

func (c LpushKey) Element(Element ...string) LpushElement {
	return LpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Lpushx struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LpushxElement) Element(Element ...string) LpushxElement {
	return LpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c LpushxElement) Build() Completed {
	return Completed(c)
}

type LpushxKey struct {
	cs []string
	cf uint32
}

func (c LpushxKey) Element(Element ...string) LpushxElement {
	return LpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Lrange struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LrangeKey) Start(Start int64) LrangeStart {
	return LrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c LrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type LrangeStart struct {
	cs []string
	cf uint32
}

func (c LrangeStart) Stop(Stop int64) LrangeStop {
	return LrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

func (c LrangeStart) Cache() Cacheable {
	return Cacheable(c)
}

type LrangeStop struct {
	cs []string
	cf uint32
}

func (c LrangeStop) Build() Completed {
	return Completed(c)
}

func (c LrangeStop) Cache() Cacheable {
	return Cacheable(c)
}

type Lrem struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LremCount) Element(Element string) LremElement {
	return LremElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LremElement struct {
	cs []string
	cf uint32
}

func (c LremElement) Build() Completed {
	return Completed(c)
}

type LremKey struct {
	cs []string
	cf uint32
}

func (c LremKey) Count(Count int64) LremCount {
	return LremCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type Lset struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LsetElement) Build() Completed {
	return Completed(c)
}

type LsetIndex struct {
	cs []string
	cf uint32
}

func (c LsetIndex) Element(Element string) LsetElement {
	return LsetElement{cf: c.cf, cs: append(c.cs, Element)}
}

type LsetKey struct {
	cs []string
	cf uint32
}

func (c LsetKey) Index(Index int64) LsetIndex {
	return LsetIndex{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type Ltrim struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c LtrimKey) Start(Start int64) LtrimStart {
	return LtrimStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type LtrimStart struct {
	cs []string
	cf uint32
}

func (c LtrimStart) Stop(Stop int64) LtrimStop {
	return LtrimStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type LtrimStop struct {
	cs []string
	cf uint32
}

func (c LtrimStop) Build() Completed {
	return Completed(c)
}

type MemoryDoctor struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c MemoryUsageKey) Samples(Count int64) MemoryUsageSamples {
	return MemoryUsageSamples{cf: c.cf, cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10))}
}

func (c MemoryUsageKey) Build() Completed {
	return Completed(c)
}

type MemoryUsageSamples struct {
	cs []string
	cf uint32
}

func (c MemoryUsageSamples) Build() Completed {
	return Completed(c)
}

type Mget struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c MgetKey) Key(Key ...string) MgetKey {
	return MgetKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c MgetKey) Build() Completed {
	return Completed(c)
}

type Migrate struct {
	cs []string
	cf uint32
}

func (c Migrate) Host(Host string) MigrateHost {
	return MigrateHost{cf: c.cf, cs: append(c.cs, Host)}
}

func (b *Builder) Migrate() (c Migrate) {
	c.cf = blockTag
	c.cs = append(b.get(), "MIGRATE")
	return
}

type MigrateAuth struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c MigrateDestinationDb) Timeout(Timeout int64) MigrateTimeout {
	return MigrateTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type MigrateHost struct {
	cs []string
	cf uint32
}

func (c MigrateHost) Port(Port string) MigratePort {
	return MigratePort{cf: c.cf, cs: append(c.cs, Port)}
}

type MigrateKeyEmpty struct {
	cs []string
	cf uint32
}

func (c MigrateKeyEmpty) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeyKey struct {
	cs []string
	cf uint32
}

func (c MigrateKeyKey) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeys struct {
	cs []string
	cf uint32
}

func (c MigrateKeys) Keys(Keys ...string) MigrateKeys {
	return MigrateKeys{cf: c.cf, cs: append(c.cs, Keys...)}
}

func (c MigrateKeys) Build() Completed {
	return Completed(c)
}

type MigratePort struct {
	cs []string
	cf uint32
}

func (c MigratePort) Key() MigrateKeyKey {
	return MigrateKeyKey{cf: c.cf, cs: append(c.cs, "key")}
}

func (c MigratePort) Empty() MigrateKeyEmpty {
	return MigrateKeyEmpty{cf: c.cf, cs: append(c.cs, "\"\"")}
}

type MigrateReplaceReplace struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ModuleLoadArg) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c ModuleLoadArg) Build() Completed {
	return Completed(c)
}

type ModuleLoadPath struct {
	cs []string
	cf uint32
}

func (c ModuleLoadPath) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cf: c.cf, cs: append(c.cs, Arg...)}
}

func (c ModuleLoadPath) Build() Completed {
	return Completed(c)
}

type ModuleUnload struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ModuleUnloadName) Build() Completed {
	return Completed(c)
}

type Monitor struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c MoveDb) Build() Completed {
	return Completed(c)
}

type MoveKey struct {
	cs []string
	cf uint32
}

func (c MoveKey) Db(Db int64) MoveDb {
	return MoveDb{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Db, 10))}
}

type Mset struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c MsetKeyValue) KeyValue(Key string, Value string) MsetKeyValue {
	return MsetKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c MsetKeyValue) Build() Completed {
	return Completed(c)
}

type Msetnx struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c MsetnxKeyValue) KeyValue(Key string, Value string) MsetnxKeyValue {
	return MsetnxKeyValue{cf: c.cf, cs: append(c.cs, Key, Value)}
}

func (c MsetnxKeyValue) Build() Completed {
	return Completed(c)
}

type Multi struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ObjectArguments) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c ObjectArguments) Build() Completed {
	return Completed(c)
}

type ObjectSubcommand struct {
	cs []string
	cf uint32
}

func (c ObjectSubcommand) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cf: c.cf, cs: append(c.cs, Arguments...)}
}

func (c ObjectSubcommand) Build() Completed {
	return Completed(c)
}

type Persist struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PersistKey) Build() Completed {
	return Completed(c)
}

type Pexpire struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PexpireConditionGt) Build() Completed {
	return Completed(c)
}

type PexpireConditionLt struct {
	cs []string
	cf uint32
}

func (c PexpireConditionLt) Build() Completed {
	return Completed(c)
}

type PexpireConditionNx struct {
	cs []string
	cf uint32
}

func (c PexpireConditionNx) Build() Completed {
	return Completed(c)
}

type PexpireConditionXx struct {
	cs []string
	cf uint32
}

func (c PexpireConditionXx) Build() Completed {
	return Completed(c)
}

type PexpireKey struct {
	cs []string
	cf uint32
}

func (c PexpireKey) Milliseconds(Milliseconds int64) PexpireMilliseconds {
	return PexpireMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PexpireMilliseconds struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c PexpireatConditionGt) Build() Completed {
	return Completed(c)
}

type PexpireatConditionLt struct {
	cs []string
	cf uint32
}

func (c PexpireatConditionLt) Build() Completed {
	return Completed(c)
}

type PexpireatConditionNx struct {
	cs []string
	cf uint32
}

func (c PexpireatConditionNx) Build() Completed {
	return Completed(c)
}

type PexpireatConditionXx struct {
	cs []string
	cf uint32
}

func (c PexpireatConditionXx) Build() Completed {
	return Completed(c)
}

type PexpireatKey struct {
	cs []string
	cf uint32
}

func (c PexpireatKey) MillisecondsTimestamp(MillisecondsTimestamp int64) PexpireatMillisecondsTimestamp {
	return PexpireatMillisecondsTimestamp{cf: c.cf, cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10))}
}

type PexpireatMillisecondsTimestamp struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c PexpiretimeKey) Build() Completed {
	return Completed(c)
}

type Pfadd struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PfaddElement) Element(Element ...string) PfaddElement {
	return PfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c PfaddElement) Build() Completed {
	return Completed(c)
}

type PfaddKey struct {
	cs []string
	cf uint32
}

func (c PfaddKey) Element(Element ...string) PfaddElement {
	return PfaddElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c PfaddKey) Build() Completed {
	return Completed(c)
}

type Pfcount struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PfcountKey) Key(Key ...string) PfcountKey {
	return PfcountKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c PfcountKey) Build() Completed {
	return Completed(c)
}

type Pfmerge struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PfmergeDestkey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

type PfmergeSourcekey struct {
	cs []string
	cf uint32
}

func (c PfmergeSourcekey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cf: c.cf, cs: append(c.cs, Sourcekey...)}
}

func (c PfmergeSourcekey) Build() Completed {
	return Completed(c)
}

type Ping struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PingMessage) Build() Completed {
	return Completed(c)
}

type Psetex struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PsetexKey) Milliseconds(Milliseconds int64) PsetexMilliseconds {
	return PsetexMilliseconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PsetexMilliseconds struct {
	cs []string
	cf uint32
}

func (c PsetexMilliseconds) Value(Value string) PsetexValue {
	return PsetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type PsetexValue struct {
	cs []string
	cf uint32
}

func (c PsetexValue) Build() Completed {
	return Completed(c)
}

type Psubscribe struct {
	cs []string
	cf uint32
}

func (c Psubscribe) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (b *Builder) Psubscribe() (c Psubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	return
}

type PsubscribePattern struct {
	cs []string
	cf uint32
}

func (c PsubscribePattern) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c PsubscribePattern) Build() Completed {
	return Completed(c)
}

type Psync struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PsyncOffset) Build() Completed {
	return Completed(c)
}

type PsyncReplicationid struct {
	cs []string
	cf uint32
}

func (c PsyncReplicationid) Offset(Offset int64) PsyncOffset {
	return PsyncOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type Pttl struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PttlKey) Build() Completed {
	return Completed(c)
}

func (c PttlKey) Cache() Cacheable {
	return Cacheable(c)
}

type Publish struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PublishChannel) Message(Message string) PublishMessage {
	return PublishMessage{cf: c.cf, cs: append(c.cs, Message)}
}

type PublishMessage struct {
	cs []string
	cf uint32
}

func (c PublishMessage) Build() Completed {
	return Completed(c)
}

type Pubsub struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c PubsubArgument) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c PubsubArgument) Build() Completed {
	return Completed(c)
}

type PubsubSubcommand struct {
	cs []string
	cf uint32
}

func (c PubsubSubcommand) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cf: c.cf, cs: append(c.cs, Argument...)}
}

func (c PubsubSubcommand) Build() Completed {
	return Completed(c)
}

type Punsubscribe struct {
	cs []string
	cf uint32
}

func (c Punsubscribe) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c Punsubscribe) Build() Completed {
	return Completed(c)
}

func (b *Builder) Punsubscribe() (c Punsubscribe) {
	c.cs = append(b.get(), "PUNSUBSCRIBE")
	return
}

type PunsubscribePattern struct {
	cs []string
	cf uint32
}

func (c PunsubscribePattern) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cf: c.cf, cs: append(c.cs, Pattern...)}
}

func (c PunsubscribePattern) Build() Completed {
	return Completed(c)
}

type Quit struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c RenameKey) Newkey(Newkey string) RenameNewkey {
	return RenameNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type RenameNewkey struct {
	cs []string
	cf uint32
}

func (c RenameNewkey) Build() Completed {
	return Completed(c)
}

type Renamenx struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c RenamenxKey) Newkey(Newkey string) RenamenxNewkey {
	return RenamenxNewkey{cf: c.cf, cs: append(c.cs, Newkey)}
}

type RenamenxNewkey struct {
	cs []string
	cf uint32
}

func (c RenamenxNewkey) Build() Completed {
	return Completed(c)
}

type Replicaof struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ReplicaofHost) Port(Port string) ReplicaofPort {
	return ReplicaofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type ReplicaofPort struct {
	cs []string
	cf uint32
}

func (c ReplicaofPort) Build() Completed {
	return Completed(c)
}

type Reset struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c RestoreFreq) Build() Completed {
	return Completed(c)
}

type RestoreIdletime struct {
	cs []string
	cf uint32
}

func (c RestoreIdletime) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cf: c.cf, cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreIdletime) Build() Completed {
	return Completed(c)
}

type RestoreKey struct {
	cs []string
	cf uint32
}

func (c RestoreKey) Ttl(Ttl int64) RestoreTtl {
	return RestoreTtl{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Ttl, 10))}
}

type RestoreReplaceReplace struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c RestoreTtl) SerializedValue(SerializedValue string) RestoreSerializedValue {
	return RestoreSerializedValue{cf: c.cf, cs: append(c.cs, SerializedValue)}
}

type Role struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c RpopCount) Build() Completed {
	return Completed(c)
}

type RpopKey struct {
	cs []string
	cf uint32
}

func (c RpopKey) Count(Count int64) RpopCount {
	return RpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c RpopKey) Build() Completed {
	return Completed(c)
}

type Rpoplpush struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c RpoplpushDestination) Build() Completed {
	return Completed(c)
}

type RpoplpushSource struct {
	cs []string
	cf uint32
}

func (c RpoplpushSource) Destination(Destination string) RpoplpushDestination {
	return RpoplpushDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Rpush struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c RpushElement) Element(Element ...string) RpushElement {
	return RpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c RpushElement) Build() Completed {
	return Completed(c)
}

type RpushKey struct {
	cs []string
	cf uint32
}

func (c RpushKey) Element(Element ...string) RpushElement {
	return RpushElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Rpushx struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c RpushxElement) Element(Element ...string) RpushxElement {
	return RpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

func (c RpushxElement) Build() Completed {
	return Completed(c)
}

type RpushxKey struct {
	cs []string
	cf uint32
}

func (c RpushxKey) Element(Element ...string) RpushxElement {
	return RpushxElement{cf: c.cf, cs: append(c.cs, Element...)}
}

type Sadd struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SaddKey) Member(Member ...string) SaddMember {
	return SaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SaddMember struct {
	cs []string
	cf uint32
}

func (c SaddMember) Member(Member ...string) SaddMember {
	return SaddMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SaddMember) Build() Completed {
	return Completed(c)
}

type Save struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ScanCount) Type(Type string) ScanType {
	return ScanType{cf: c.cf, cs: append(c.cs, "TYPE", Type)}
}

func (c ScanCount) Build() Completed {
	return Completed(c)
}

type ScanCursor struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ScanType) Build() Completed {
	return Completed(c)
}

type Scard struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ScardKey) Build() Completed {
	return Completed(c)
}

func (c ScardKey) Cache() Cacheable {
	return Cacheable(c)
}

type ScriptDebug struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ScriptDebugModeNo) Build() Completed {
	return Completed(c)
}

type ScriptDebugModeSync struct {
	cs []string
	cf uint32
}

func (c ScriptDebugModeSync) Build() Completed {
	return Completed(c)
}

type ScriptDebugModeYes struct {
	cs []string
	cf uint32
}

func (c ScriptDebugModeYes) Build() Completed {
	return Completed(c)
}

type ScriptExists struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ScriptExistsSha1) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cf: c.cf, cs: append(c.cs, Sha1...)}
}

func (c ScriptExistsSha1) Build() Completed {
	return Completed(c)
}

type ScriptFlush struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ScriptFlushAsyncAsync) Build() Completed {
	return Completed(c)
}

type ScriptFlushAsyncSync struct {
	cs []string
	cf uint32
}

func (c ScriptFlushAsyncSync) Build() Completed {
	return Completed(c)
}

type ScriptKill struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ScriptLoadScript) Build() Completed {
	return Completed(c)
}

type Sdiff struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SdiffKey) Key(Key ...string) SdiffKey {
	return SdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SdiffKey) Build() Completed {
	return Completed(c)
}

type Sdiffstore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SdiffstoreDestination) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SdiffstoreKey struct {
	cs []string
	cf uint32
}

func (c SdiffstoreKey) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SdiffstoreKey) Build() Completed {
	return Completed(c)
}

type Select struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SelectIndex) Build() Completed {
	return Completed(c)
}

type Set struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SetConditionNx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetConditionNx) Build() Completed {
	return Completed(c)
}

type SetConditionXx struct {
	cs []string
	cf uint32
}

func (c SetConditionXx) Get() SetGetGet {
	return SetGetGet{cf: c.cf, cs: append(c.cs, "GET")}
}

func (c SetConditionXx) Build() Completed {
	return Completed(c)
}

type SetExpirationEx struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c SetGetGet) Build() Completed {
	return Completed(c)
}

type SetKey struct {
	cs []string
	cf uint32
}

func (c SetKey) Value(Value string) SetValue {
	return SetValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetValue struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c SetbitKey) Offset(Offset int64) SetbitOffset {
	return SetbitOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetbitOffset struct {
	cs []string
	cf uint32
}

func (c SetbitOffset) Value(Value int64) SetbitValue {
	return SetbitValue{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Value, 10))}
}

type SetbitValue struct {
	cs []string
	cf uint32
}

func (c SetbitValue) Build() Completed {
	return Completed(c)
}

type Setex struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SetexKey) Seconds(Seconds int64) SetexSeconds {
	return SetexSeconds{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SetexSeconds struct {
	cs []string
	cf uint32
}

func (c SetexSeconds) Value(Value string) SetexValue {
	return SetexValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetexValue struct {
	cs []string
	cf uint32
}

func (c SetexValue) Build() Completed {
	return Completed(c)
}

type Setnx struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SetnxKey) Value(Value string) SetnxValue {
	return SetnxValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetnxValue struct {
	cs []string
	cf uint32
}

func (c SetnxValue) Build() Completed {
	return Completed(c)
}

type Setrange struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SetrangeKey) Offset(Offset int64) SetrangeOffset {
	return SetrangeOffset{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetrangeOffset struct {
	cs []string
	cf uint32
}

func (c SetrangeOffset) Value(Value string) SetrangeValue {
	return SetrangeValue{cf: c.cf, cs: append(c.cs, Value)}
}

type SetrangeValue struct {
	cs []string
	cf uint32
}

func (c SetrangeValue) Build() Completed {
	return Completed(c)
}

type Shutdown struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ShutdownSaveModeNosave) Build() Completed {
	return Completed(c)
}

type ShutdownSaveModeSave struct {
	cs []string
	cf uint32
}

func (c ShutdownSaveModeSave) Build() Completed {
	return Completed(c)
}

type Sinter struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SinterKey) Key(Key ...string) SinterKey {
	return SinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SinterKey) Build() Completed {
	return Completed(c)
}

type Sintercard struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SintercardKey) Key(Key ...string) SintercardKey {
	return SintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SintercardKey) Build() Completed {
	return Completed(c)
}

type Sinterstore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SinterstoreDestination) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SinterstoreKey struct {
	cs []string
	cf uint32
}

func (c SinterstoreKey) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SinterstoreKey) Build() Completed {
	return Completed(c)
}

type Sismember struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SismemberKey) Member(Member string) SismemberMember {
	return SismemberMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c SismemberKey) Cache() Cacheable {
	return Cacheable(c)
}

type SismemberMember struct {
	cs []string
	cf uint32
}

func (c SismemberMember) Build() Completed {
	return Completed(c)
}

func (c SismemberMember) Cache() Cacheable {
	return Cacheable(c)
}

type Slaveof struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SlaveofHost) Port(Port string) SlaveofPort {
	return SlaveofPort{cf: c.cf, cs: append(c.cs, Port)}
}

type SlaveofPort struct {
	cs []string
	cf uint32
}

func (c SlaveofPort) Build() Completed {
	return Completed(c)
}

type Slowlog struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SlowlogArgument) Build() Completed {
	return Completed(c)
}

type SlowlogSubcommand struct {
	cs []string
	cf uint32
}

func (c SlowlogSubcommand) Argument(Argument string) SlowlogArgument {
	return SlowlogArgument{cf: c.cf, cs: append(c.cs, Argument)}
}

func (c SlowlogSubcommand) Build() Completed {
	return Completed(c)
}

type Smembers struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SmembersKey) Build() Completed {
	return Completed(c)
}

func (c SmembersKey) Cache() Cacheable {
	return Cacheable(c)
}

type Smismember struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SmismemberKey) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SmismemberKey) Cache() Cacheable {
	return Cacheable(c)
}

type SmismemberMember struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c SmoveDestination) Member(Member string) SmoveMember {
	return SmoveMember{cf: c.cf, cs: append(c.cs, Member)}
}

type SmoveMember struct {
	cs []string
	cf uint32
}

func (c SmoveMember) Build() Completed {
	return Completed(c)
}

type SmoveSource struct {
	cs []string
	cf uint32
}

func (c SmoveSource) Destination(Destination string) SmoveDestination {
	return SmoveDestination{cf: c.cf, cs: append(c.cs, Destination)}
}

type Sort struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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

type SortRoGet struct {
	cs []string
	cf uint32
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

type SortRoKey struct {
	cs []string
	cf uint32
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

type SortRoLimit struct {
	cs []string
	cf uint32
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

type SortRoOrderAsc struct {
	cs []string
	cf uint32
}

func (c SortRoOrderAsc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoOrderAsc) Build() Completed {
	return Completed(c)
}

type SortRoOrderDesc struct {
	cs []string
	cf uint32
}

func (c SortRoOrderDesc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cf: c.cf, cs: append(c.cs, "ALPHA")}
}

func (c SortRoOrderDesc) Build() Completed {
	return Completed(c)
}

type SortRoSortingAlpha struct {
	cs []string
	cf uint32
}

func (c SortRoSortingAlpha) Build() Completed {
	return Completed(c)
}

type SortSortingAlpha struct {
	cs []string
	cf uint32
}

func (c SortSortingAlpha) Store(Destination string) SortStore {
	return SortStore{cf: c.cf, cs: append(c.cs, "STORE", Destination)}
}

func (c SortSortingAlpha) Build() Completed {
	return Completed(c)
}

type SortStore struct {
	cs []string
	cf uint32
}

func (c SortStore) Build() Completed {
	return Completed(c)
}

type Spop struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SpopCount) Build() Completed {
	return Completed(c)
}

type SpopKey struct {
	cs []string
	cf uint32
}

func (c SpopKey) Count(Count int64) SpopCount {
	return SpopCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SpopKey) Build() Completed {
	return Completed(c)
}

type Srandmember struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SrandmemberCount) Build() Completed {
	return Completed(c)
}

type SrandmemberKey struct {
	cs []string
	cf uint32
}

func (c SrandmemberKey) Count(Count int64) SrandmemberCount {
	return SrandmemberCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SrandmemberKey) Build() Completed {
	return Completed(c)
}

type Srem struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SremKey) Member(Member ...string) SremMember {
	return SremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type SremMember struct {
	cs []string
	cf uint32
}

func (c SremMember) Member(Member ...string) SremMember {
	return SremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c SremMember) Build() Completed {
	return Completed(c)
}

type Sscan struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SscanCount) Build() Completed {
	return Completed(c)
}

type SscanCursor struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SscanKey) Cursor(Cursor int64) SscanCursor {
	return SscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SscanMatch struct {
	cs []string
	cf uint32
}

func (c SscanMatch) Count(Count int64) SscanCount {
	return SscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SscanMatch) Build() Completed {
	return Completed(c)
}

type Stralgo struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c StralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

func (c StralgoAlgoSpecificArgument) Build() Completed {
	return Completed(c)
}

type StralgoAlgorithmLcs struct {
	cs []string
	cf uint32
}

func (c StralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cf: c.cf, cs: append(c.cs, AlgoSpecificArgument...)}
}

type Strlen struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c StrlenKey) Build() Completed {
	return Completed(c)
}

func (c StrlenKey) Cache() Cacheable {
	return Cacheable(c)
}

type Subscribe struct {
	cs []string
	cf uint32
}

func (c Subscribe) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (b *Builder) Subscribe() (c Subscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	return
}

type SubscribeChannel struct {
	cs []string
	cf uint32
}

func (c SubscribeChannel) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c SubscribeChannel) Build() Completed {
	return Completed(c)
}

type Sunion struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SunionKey) Key(Key ...string) SunionKey {
	return SunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SunionKey) Build() Completed {
	return Completed(c)
}

type Sunionstore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SunionstoreDestination) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type SunionstoreKey struct {
	cs []string
	cf uint32
}

func (c SunionstoreKey) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c SunionstoreKey) Build() Completed {
	return Completed(c)
}

type Swapdb struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c SwapdbIndex1) Index2(Index2 int64) SwapdbIndex2 {
	return SwapdbIndex2{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Index2, 10))}
}

type SwapdbIndex2 struct {
	cs []string
	cf uint32
}

func (c SwapdbIndex2) Build() Completed {
	return Completed(c)
}

type Sync struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c TouchKey) Key(Key ...string) TouchKey {
	return TouchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c TouchKey) Build() Completed {
	return Completed(c)
}

type Ttl struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c TtlKey) Build() Completed {
	return Completed(c)
}

func (c TtlKey) Cache() Cacheable {
	return Cacheable(c)
}

type Type struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c TypeKey) Build() Completed {
	return Completed(c)
}

func (c TypeKey) Cache() Cacheable {
	return Cacheable(c)
}

type Unlink struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c UnlinkKey) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c UnlinkKey) Build() Completed {
	return Completed(c)
}

type Unsubscribe struct {
	cs []string
	cf uint32
}

func (c Unsubscribe) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c Unsubscribe) Build() Completed {
	return Completed(c)
}

func (b *Builder) Unsubscribe() (c Unsubscribe) {
	c.cs = append(b.get(), "UNSUBSCRIBE")
	return
}

type UnsubscribeChannel struct {
	cs []string
	cf uint32
}

func (c UnsubscribeChannel) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cf: c.cf, cs: append(c.cs, Channel...)}
}

func (c UnsubscribeChannel) Build() Completed {
	return Completed(c)
}

type Unwatch struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c Wait) Numreplicas(Numreplicas int64) WaitNumreplicas {
	return WaitNumreplicas{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numreplicas, 10))}
}

func (b *Builder) Wait() (c Wait) {
	c.cf = blockTag
	c.cs = append(b.get(), "WAIT")
	return
}

type WaitNumreplicas struct {
	cs []string
	cf uint32
}

func (c WaitNumreplicas) Timeout(Timeout int64) WaitTimeout {
	return WaitTimeout{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type WaitTimeout struct {
	cs []string
	cf uint32
}

func (c WaitTimeout) Build() Completed {
	return Completed(c)
}

type Watch struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c WatchKey) Key(Key ...string) WatchKey {
	return WatchKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c WatchKey) Build() Completed {
	return Completed(c)
}

type Xack struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XackGroup) Id(Id ...string) XackId {
	return XackId{cf: c.cf, cs: append(c.cs, Id...)}
}

type XackId struct {
	cs []string
	cf uint32
}

func (c XackId) Id(Id ...string) XackId {
	return XackId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XackId) Build() Completed {
	return Completed(c)
}

type XackKey struct {
	cs []string
	cf uint32
}

func (c XackKey) Group(Group string) XackGroup {
	return XackGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type Xadd struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XaddFieldValue) FieldValue(Field string, Value string) XaddFieldValue {
	return XaddFieldValue{cf: c.cf, cs: append(c.cs, Field, Value)}
}

func (c XaddFieldValue) Build() Completed {
	return Completed(c)
}

type XaddId struct {
	cs []string
	cf uint32
}

func (c XaddId) FieldValue() XaddFieldValue {
	return XaddFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type XaddKey struct {
	cs []string
	cf uint32
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

func (c XaddKey) Wildcard() XaddWildcard {
	return XaddWildcard{cf: c.cf, cs: append(c.cs, "*")}
}

func (c XaddKey) Id() XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, "ID")}
}

type XaddNomkstream struct {
	cs []string
	cf uint32
}

func (c XaddNomkstream) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c XaddNomkstream) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

func (c XaddNomkstream) Wildcard() XaddWildcard {
	return XaddWildcard{cf: c.cf, cs: append(c.cs, "*")}
}

func (c XaddNomkstream) Id() XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, "ID")}
}

type XaddTrimLimit struct {
	cs []string
	cf uint32
}

func (c XaddTrimLimit) Wildcard() XaddWildcard {
	return XaddWildcard{cf: c.cf, cs: append(c.cs, "*")}
}

func (c XaddTrimLimit) Id() XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, "ID")}
}

type XaddTrimOperatorAlmost struct {
	cs []string
	cf uint32
}

func (c XaddTrimOperatorAlmost) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimOperatorExact struct {
	cs []string
	cf uint32
}

func (c XaddTrimOperatorExact) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XaddTrimStrategyMaxlen struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c XaddTrimThreshold) Limit(Count int64) XaddTrimLimit {
	return XaddTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XaddTrimThreshold) Wildcard() XaddWildcard {
	return XaddWildcard{cf: c.cf, cs: append(c.cs, "*")}
}

func (c XaddTrimThreshold) Id() XaddId {
	return XaddId{cf: c.cf, cs: append(c.cs, "ID")}
}

type XaddWildcard struct {
	cs []string
	cf uint32
}

func (c XaddWildcard) FieldValue() XaddFieldValue {
	return XaddFieldValue{cf: c.cf, cs: append(c.cs, )}
}

type Xautoclaim struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XautoclaimConsumer) MinIdleTime(MinIdleTime string) XautoclaimMinIdleTime {
	return XautoclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type XautoclaimCount struct {
	cs []string
	cf uint32
}

func (c XautoclaimCount) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XautoclaimCount) Build() Completed {
	return Completed(c)
}

type XautoclaimGroup struct {
	cs []string
	cf uint32
}

func (c XautoclaimGroup) Consumer(Consumer string) XautoclaimConsumer {
	return XautoclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type XautoclaimJustidJustid struct {
	cs []string
	cf uint32
}

func (c XautoclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XautoclaimKey struct {
	cs []string
	cf uint32
}

func (c XautoclaimKey) Group(Group string) XautoclaimGroup {
	return XautoclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type XautoclaimMinIdleTime struct {
	cs []string
	cf uint32
}

func (c XautoclaimMinIdleTime) Start(Start string) XautoclaimStart {
	return XautoclaimStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XautoclaimStart struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c XclaimConsumer) MinIdleTime(MinIdleTime string) XclaimMinIdleTime {
	return XclaimMinIdleTime{cf: c.cf, cs: append(c.cs, MinIdleTime)}
}

type XclaimForceForce struct {
	cs []string
	cf uint32
}

func (c XclaimForceForce) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cf: c.cf, cs: append(c.cs, "JUSTID")}
}

func (c XclaimForceForce) Build() Completed {
	return Completed(c)
}

type XclaimGroup struct {
	cs []string
	cf uint32
}

func (c XclaimGroup) Consumer(Consumer string) XclaimConsumer {
	return XclaimConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

type XclaimId struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c XclaimJustidJustid) Build() Completed {
	return Completed(c)
}

type XclaimKey struct {
	cs []string
	cf uint32
}

func (c XclaimKey) Group(Group string) XclaimGroup {
	return XclaimGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type XclaimMinIdleTime struct {
	cs []string
	cf uint32
}

func (c XclaimMinIdleTime) Id(Id ...string) XclaimId {
	return XclaimId{cf: c.cf, cs: append(c.cs, Id...)}
}

type XclaimRetrycount struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c XdelId) Id(Id ...string) XdelId {
	return XdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XdelId) Build() Completed {
	return Completed(c)
}

type XdelKey struct {
	cs []string
	cf uint32
}

func (c XdelKey) Id(Id ...string) XdelId {
	return XdelId{cf: c.cf, cs: append(c.cs, Id...)}
}

type Xgroup struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XgroupCreateCreate) Id() XgroupCreateIdId {
	return XgroupCreateIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c XgroupCreateCreate) Lastid() XgroupCreateIdLastID {
	return XgroupCreateIdLastID{cf: c.cf, cs: append(c.cs, "$")}
}

type XgroupCreateIdId struct {
	cs []string
	cf uint32
}

func (c XgroupCreateIdId) Mkstream() XgroupCreateMkstream {
	return XgroupCreateMkstream{cf: c.cf, cs: append(c.cs, "MKSTREAM")}
}

func (c XgroupCreateIdId) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c XgroupCreateIdId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupCreateIdId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateIdId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type XgroupCreateIdLastID struct {
	cs []string
	cf uint32
}

func (c XgroupCreateIdLastID) Mkstream() XgroupCreateMkstream {
	return XgroupCreateMkstream{cf: c.cf, cs: append(c.cs, "MKSTREAM")}
}

func (c XgroupCreateIdLastID) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cf: c.cf, cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c XgroupCreateIdLastID) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupCreateIdLastID) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateIdLastID) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type XgroupCreateMkstream struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateconsumer) Build() Completed {
	return Completed(c)
}

type XgroupDelconsumer struct {
	cs []string
	cf uint32
}

func (c XgroupDelconsumer) Build() Completed {
	return Completed(c)
}

type XgroupDestroy struct {
	cs []string
	cf uint32
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

type XgroupSetidIdId struct {
	cs []string
	cf uint32
}

func (c XgroupSetidIdId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupSetidIdId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdId) Build() Completed {
	return Completed(c)
}

type XgroupSetidIdLastID struct {
	cs []string
	cf uint32
}

func (c XgroupSetidIdLastID) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cf: c.cf, cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupSetidIdLastID) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cf: c.cf, cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdLastID) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cf: c.cf, cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdLastID) Build() Completed {
	return Completed(c)
}

type XgroupSetidSetid struct {
	cs []string
	cf uint32
}

func (c XgroupSetidSetid) Id() XgroupSetidIdId {
	return XgroupSetidIdId{cf: c.cf, cs: append(c.cs, "ID")}
}

func (c XgroupSetidSetid) Lastid() XgroupSetidIdLastID {
	return XgroupSetidIdLastID{cf: c.cf, cs: append(c.cs, "$")}
}

type Xinfo struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c XinfoHelpHelp) Build() Completed {
	return Completed(c)
}

type XinfoStream struct {
	cs []string
	cf uint32
}

func (c XinfoStream) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cf: c.cf, cs: append(c.cs, "HELP")}
}

func (c XinfoStream) Build() Completed {
	return Completed(c)
}

type Xlen struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XlenKey) Build() Completed {
	return Completed(c)
}

type Xpending struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XpendingFiltersConsumer) Build() Completed {
	return Completed(c)
}

type XpendingFiltersCount struct {
	cs []string
	cf uint32
}

func (c XpendingFiltersCount) Consumer(Consumer string) XpendingFiltersConsumer {
	return XpendingFiltersConsumer{cf: c.cf, cs: append(c.cs, Consumer)}
}

func (c XpendingFiltersCount) Build() Completed {
	return Completed(c)
}

type XpendingFiltersEnd struct {
	cs []string
	cf uint32
}

func (c XpendingFiltersEnd) Count(Count int64) XpendingFiltersCount {
	return XpendingFiltersCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type XpendingFiltersIdle struct {
	cs []string
	cf uint32
}

func (c XpendingFiltersIdle) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XpendingFiltersStart struct {
	cs []string
	cf uint32
}

func (c XpendingFiltersStart) End(End string) XpendingFiltersEnd {
	return XpendingFiltersEnd{cf: c.cf, cs: append(c.cs, End)}
}

type XpendingGroup struct {
	cs []string
	cf uint32
}

func (c XpendingGroup) Idle(MinIdleTime int64) XpendingFiltersIdle {
	return XpendingFiltersIdle{cf: c.cf, cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10))}
}

func (c XpendingGroup) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XpendingKey struct {
	cs []string
	cf uint32
}

func (c XpendingKey) Group(Group string) XpendingGroup {
	return XpendingGroup{cf: c.cf, cs: append(c.cs, Group)}
}

type Xrange struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XrangeCount) Build() Completed {
	return Completed(c)
}

type XrangeEnd struct {
	cs []string
	cf uint32
}

func (c XrangeEnd) Count(Count int64) XrangeCount {
	return XrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrangeEnd) Build() Completed {
	return Completed(c)
}

type XrangeKey struct {
	cs []string
	cf uint32
}

func (c XrangeKey) Start(Start string) XrangeStart {
	return XrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XrangeStart struct {
	cs []string
	cf uint32
}

func (c XrangeStart) End(End string) XrangeEnd {
	return XrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type Xread struct {
	cs []string
	cf uint32
}

func (c Xread) Count(Count int64) XreadCount {
	return XreadCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c Xread) Block(Milliseconds int64) XreadBlock {
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
	cf uint32
}

func (c XreadBlock) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadCount struct {
	cs []string
	cf uint32
}

func (c XreadCount) Block(Milliseconds int64) XreadBlock {
	return XreadBlock{cf: c.cf, cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadCount) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadId struct {
	cs []string
	cf uint32
}

func (c XreadId) Id(Id ...string) XreadId {
	return XreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadId) Build() Completed {
	return Completed(c)
}

type XreadKey struct {
	cs []string
	cf uint32
}

func (c XreadKey) Id(Id ...string) XreadId {
	return XreadId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadKey) Key(Key ...string) XreadKey {
	return XreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type XreadStreamsStreams struct {
	cs []string
	cf uint32
}

func (c XreadStreamsStreams) Key(Key ...string) XreadKey {
	return XreadKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Xreadgroup struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XreadgroupBlock) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cf: c.cf, cs: append(c.cs, "NOACK")}
}

func (c XreadgroupBlock) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupCount struct {
	cs []string
	cf uint32
}

func (c XreadgroupCount) Block(Milliseconds int64) XreadgroupBlock {
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
	cf uint32
}

func (c XreadgroupGroup) Count(Count int64) XreadgroupCount {
	return XreadgroupCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XreadgroupGroup) Block(Milliseconds int64) XreadgroupBlock {
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
	cf uint32
}

func (c XreadgroupId) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadgroupId) Build() Completed {
	return Completed(c)
}

type XreadgroupKey struct {
	cs []string
	cf uint32
}

func (c XreadgroupKey) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cf: c.cf, cs: append(c.cs, Id...)}
}

func (c XreadgroupKey) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type XreadgroupNoackNoack struct {
	cs []string
	cf uint32
}

func (c XreadgroupNoackNoack) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cf: c.cf, cs: append(c.cs, "STREAMS")}
}

type XreadgroupStreamsStreams struct {
	cs []string
	cf uint32
}

func (c XreadgroupStreamsStreams) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Xrevrange struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XrevrangeCount) Build() Completed {
	return Completed(c)
}

type XrevrangeEnd struct {
	cs []string
	cf uint32
}

func (c XrevrangeEnd) Start(Start string) XrevrangeStart {
	return XrevrangeStart{cf: c.cf, cs: append(c.cs, Start)}
}

type XrevrangeKey struct {
	cs []string
	cf uint32
}

func (c XrevrangeKey) End(End string) XrevrangeEnd {
	return XrevrangeEnd{cf: c.cf, cs: append(c.cs, End)}
}

type XrevrangeStart struct {
	cs []string
	cf uint32
}

func (c XrevrangeStart) Count(Count int64) XrevrangeCount {
	return XrevrangeCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrevrangeStart) Build() Completed {
	return Completed(c)
}

type Xtrim struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c XtrimKey) Maxlen() XtrimTrimStrategyMaxlen {
	return XtrimTrimStrategyMaxlen{cf: c.cf, cs: append(c.cs, "MAXLEN")}
}

func (c XtrimKey) Minid() XtrimTrimStrategyMinid {
	return XtrimTrimStrategyMinid{cf: c.cf, cs: append(c.cs, "MINID")}
}

type XtrimTrimLimit struct {
	cs []string
	cf uint32
}

func (c XtrimTrimLimit) Build() Completed {
	return Completed(c)
}

type XtrimTrimOperatorAlmost struct {
	cs []string
	cf uint32
}

func (c XtrimTrimOperatorAlmost) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimOperatorExact struct {
	cs []string
	cf uint32
}

func (c XtrimTrimOperatorExact) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cf: c.cf, cs: append(c.cs, Threshold)}
}

type XtrimTrimStrategyMaxlen struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c XtrimTrimThreshold) Limit(Count int64) XtrimTrimLimit {
	return XtrimTrimLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XtrimTrimThreshold) Build() Completed {
	return Completed(c)
}

type Zadd struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZaddChangeCh) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cf: c.cf, cs: append(c.cs, "INCR")}
}

func (c ZaddChangeCh) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddComparisonGt struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZaddIncrementIncr) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, )}
}

type ZaddKey struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZaddScoreMember) ScoreMember(Score float64, Member string) ZaddScoreMember {
	return ZaddScoreMember{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member)}
}

func (c ZaddScoreMember) Build() Completed {
	return Completed(c)
}

type Zcard struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZcardKey) Build() Completed {
	return Completed(c)
}

func (c ZcardKey) Cache() Cacheable {
	return Cacheable(c)
}

type Zcount struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZcountKey) Min(Min float64) ZcountMin {
	return ZcountMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c ZcountKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZcountMax struct {
	cs []string
	cf uint32
}

func (c ZcountMax) Build() Completed {
	return Completed(c)
}

func (c ZcountMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZcountMin struct {
	cs []string
	cf uint32
}

func (c ZcountMin) Max(Max float64) ZcountMax {
	return ZcountMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c ZcountMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zdiff struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZdiffNumkeys) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZdiffWithscoresWithscores struct {
	cs []string
	cf uint32
}

func (c ZdiffWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zdiffstore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZdiffstoreDestination) Numkeys(Numkeys int64) ZdiffstoreNumkeys {
	return ZdiffstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZdiffstoreKey struct {
	cs []string
	cf uint32
}

func (c ZdiffstoreKey) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZdiffstoreKey) Build() Completed {
	return Completed(c)
}

type ZdiffstoreNumkeys struct {
	cs []string
	cf uint32
}

func (c ZdiffstoreNumkeys) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Zincrby struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZincrbyIncrement) Member(Member string) ZincrbyMember {
	return ZincrbyMember{cf: c.cf, cs: append(c.cs, Member)}
}

type ZincrbyKey struct {
	cs []string
	cf uint32
}

func (c ZincrbyKey) Increment(Increment int64) ZincrbyIncrement {
	return ZincrbyIncrement{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type ZincrbyMember struct {
	cs []string
	cf uint32
}

func (c ZincrbyMember) Build() Completed {
	return Completed(c)
}

type Zinter struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZinterAggregateMax) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMax) Build() Completed {
	return Completed(c)
}

type ZinterAggregateMin struct {
	cs []string
	cf uint32
}

func (c ZinterAggregateMin) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMin) Build() Completed {
	return Completed(c)
}

type ZinterAggregateSum struct {
	cs []string
	cf uint32
}

func (c ZinterAggregateSum) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateSum) Build() Completed {
	return Completed(c)
}

type ZinterKey struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZinterNumkeys) Key(Key ...string) ZinterKey {
	return ZinterKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZinterWeights struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZinterWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zintercard struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZintercardKey) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

func (c ZintercardKey) Build() Completed {
	return Completed(c)
}

type ZintercardNumkeys struct {
	cs []string
	cf uint32
}

func (c ZintercardNumkeys) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type Zinterstore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZinterstoreAggregateMax) Build() Completed {
	return Completed(c)
}

type ZinterstoreAggregateMin struct {
	cs []string
	cf uint32
}

func (c ZinterstoreAggregateMin) Build() Completed {
	return Completed(c)
}

type ZinterstoreAggregateSum struct {
	cs []string
	cf uint32
}

func (c ZinterstoreAggregateSum) Build() Completed {
	return Completed(c)
}

type ZinterstoreDestination struct {
	cs []string
	cf uint32
}

func (c ZinterstoreDestination) Numkeys(Numkeys int64) ZinterstoreNumkeys {
	return ZinterstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZinterstoreKey struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZinterstoreNumkeys) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZinterstoreWeights struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZlexcountKey) Min(Min string) ZlexcountMin {
	return ZlexcountMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZlexcountKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZlexcountMax struct {
	cs []string
	cf uint32
}

func (c ZlexcountMax) Build() Completed {
	return Completed(c)
}

func (c ZlexcountMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZlexcountMin struct {
	cs []string
	cf uint32
}

func (c ZlexcountMin) Max(Max string) ZlexcountMax {
	return ZlexcountMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZlexcountMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zmscore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZmscoreKey) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c ZmscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZmscoreMember struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZpopmaxCount) Build() Completed {
	return Completed(c)
}

type ZpopmaxKey struct {
	cs []string
	cf uint32
}

func (c ZpopmaxKey) Count(Count int64) ZpopmaxCount {
	return ZpopmaxCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopmaxKey) Build() Completed {
	return Completed(c)
}

type Zpopmin struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZpopminCount) Build() Completed {
	return Completed(c)
}

type ZpopminKey struct {
	cs []string
	cf uint32
}

func (c ZpopminKey) Count(Count int64) ZpopminCount {
	return ZpopminCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopminKey) Build() Completed {
	return Completed(c)
}

type Zrandmember struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrandmemberKey) Count(Count int64) ZrandmemberOptionsCount {
	return ZrandmemberOptionsCount{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type ZrandmemberOptionsCount struct {
	cs []string
	cf uint32
}

func (c ZrandmemberOptionsCount) Withscores() ZrandmemberOptionsWithscoresWithscores {
	return ZrandmemberOptionsWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZrandmemberOptionsCount) Build() Completed {
	return Completed(c)
}

type ZrandmemberOptionsWithscoresWithscores struct {
	cs []string
	cf uint32
}

func (c ZrandmemberOptionsWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zrange struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrangeKey) Min(Min string) ZrangeMin {
	return ZrangeMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeLimit struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZrangeMin) Max(Max string) ZrangeMax {
	return ZrangeMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZrangeMin) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangeRevRev struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZrangeWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrangeWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangebylex struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrangebylexKey) Min(Min string) ZrangebylexMin {
	return ZrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZrangebylexKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexLimit struct {
	cs []string
	cf uint32
}

func (c ZrangebylexLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangebylexLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebylexMax struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrangebylexMin) Max(Max string) ZrangebylexMax {
	return ZrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZrangebylexMin) Cache() Cacheable {
	return Cacheable(c)
}

type Zrangebyscore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrangebyscoreKey) Min(Min float64) ZrangebyscoreMin {
	return ZrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c ZrangebyscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreLimit struct {
	cs []string
	cf uint32
}

func (c ZrangebyscoreLimit) Build() Completed {
	return Completed(c)
}

func (c ZrangebyscoreLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreMax struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrangebyscoreMin) Max(Max float64) ZrangebyscoreMax {
	return ZrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c ZrangebyscoreMin) Cache() Cacheable {
	return Cacheable(c)
}

type ZrangebyscoreWithscoresWithscores struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZrangestoreDst) Src(Src string) ZrangestoreSrc {
	return ZrangestoreSrc{cf: c.cf, cs: append(c.cs, Src)}
}

type ZrangestoreLimit struct {
	cs []string
	cf uint32
}

func (c ZrangestoreLimit) Build() Completed {
	return Completed(c)
}

type ZrangestoreMax struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrangestoreMin) Max(Max string) ZrangestoreMax {
	return ZrangestoreMax{cf: c.cf, cs: append(c.cs, Max)}
}

type ZrangestoreRevRev struct {
	cs []string
	cf uint32
}

func (c ZrangestoreRevRev) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cf: c.cf, cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreRevRev) Build() Completed {
	return Completed(c)
}

type ZrangestoreSortbyBylex struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZrangestoreSrc) Min(Min string) ZrangestoreMin {
	return ZrangestoreMin{cf: c.cf, cs: append(c.cs, Min)}
}

type Zrank struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrankKey) Member(Member string) ZrankMember {
	return ZrankMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c ZrankKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrankMember struct {
	cs []string
	cf uint32
}

func (c ZrankMember) Build() Completed {
	return Completed(c)
}

func (c ZrankMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zrem struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZremKey) Member(Member ...string) ZremMember {
	return ZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

type ZremMember struct {
	cs []string
	cf uint32
}

func (c ZremMember) Member(Member ...string) ZremMember {
	return ZremMember{cf: c.cf, cs: append(c.cs, Member...)}
}

func (c ZremMember) Build() Completed {
	return Completed(c)
}

type Zremrangebylex struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZremrangebylexKey) Min(Min string) ZremrangebylexMin {
	return ZremrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

type ZremrangebylexMax struct {
	cs []string
	cf uint32
}

func (c ZremrangebylexMax) Build() Completed {
	return Completed(c)
}

type ZremrangebylexMin struct {
	cs []string
	cf uint32
}

func (c ZremrangebylexMin) Max(Max string) ZremrangebylexMax {
	return ZremrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

type Zremrangebyrank struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZremrangebyrankKey) Start(Start int64) ZremrangebyrankStart {
	return ZremrangebyrankStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type ZremrangebyrankStart struct {
	cs []string
	cf uint32
}

func (c ZremrangebyrankStart) Stop(Stop int64) ZremrangebyrankStop {
	return ZremrangebyrankStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type ZremrangebyrankStop struct {
	cs []string
	cf uint32
}

func (c ZremrangebyrankStop) Build() Completed {
	return Completed(c)
}

type Zremrangebyscore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZremrangebyscoreKey) Min(Min float64) ZremrangebyscoreMin {
	return ZremrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZremrangebyscoreMax struct {
	cs []string
	cf uint32
}

func (c ZremrangebyscoreMax) Build() Completed {
	return Completed(c)
}

type ZremrangebyscoreMin struct {
	cs []string
	cf uint32
}

func (c ZremrangebyscoreMin) Max(Max float64) ZremrangebyscoreMax {
	return ZremrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type Zrevrange struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrevrangeKey) Start(Start int64) ZrevrangeStart {
	return ZrevrangeStart{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

func (c ZrevrangeKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangeStart struct {
	cs []string
	cf uint32
}

func (c ZrevrangeStart) Stop(Stop int64) ZrevrangeStop {
	return ZrevrangeStop{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

func (c ZrevrangeStart) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangeStop struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrevrangeWithscoresWithscores) Build() Completed {
	return Completed(c)
}

func (c ZrevrangeWithscoresWithscores) Cache() Cacheable {
	return Cacheable(c)
}

type Zrevrangebylex struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZrevrangebylexKey) Max(Max string) ZrevrangebylexMax {
	return ZrevrangebylexMax{cf: c.cf, cs: append(c.cs, Max)}
}

func (c ZrevrangebylexKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebylexLimit struct {
	cs []string
	cf uint32
}

func (c ZrevrangebylexLimit) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebylexLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebylexMax struct {
	cs []string
	cf uint32
}

func (c ZrevrangebylexMax) Min(Min string) ZrevrangebylexMin {
	return ZrevrangebylexMin{cf: c.cf, cs: append(c.cs, Min)}
}

func (c ZrevrangebylexMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebylexMin struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZrevrangebyscoreKey) Max(Max float64) ZrevrangebyscoreMax {
	return ZrevrangebyscoreMax{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

func (c ZrevrangebyscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreLimit struct {
	cs []string
	cf uint32
}

func (c ZrevrangebyscoreLimit) Build() Completed {
	return Completed(c)
}

func (c ZrevrangebyscoreLimit) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreMax struct {
	cs []string
	cf uint32
}

func (c ZrevrangebyscoreMax) Min(Min float64) ZrevrangebyscoreMin {
	return ZrevrangebyscoreMin{cf: c.cf, cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

func (c ZrevrangebyscoreMax) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrangebyscoreMin struct {
	cs []string
	cf uint32
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
	cf uint32
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
	cf uint32
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
	cf uint32
}

func (c ZrevrankKey) Member(Member string) ZrevrankMember {
	return ZrevrankMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c ZrevrankKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZrevrankMember struct {
	cs []string
	cf uint32
}

func (c ZrevrankMember) Build() Completed {
	return Completed(c)
}

func (c ZrevrankMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zscan struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZscanCount) Build() Completed {
	return Completed(c)
}

type ZscanCursor struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZscanKey) Cursor(Cursor int64) ZscanCursor {
	return ZscanCursor{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type ZscanMatch struct {
	cs []string
	cf uint32
}

func (c ZscanMatch) Count(Count int64) ZscanCount {
	return ZscanCount{cf: c.cf, cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ZscanMatch) Build() Completed {
	return Completed(c)
}

type Zscore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZscoreKey) Member(Member string) ZscoreMember {
	return ZscoreMember{cf: c.cf, cs: append(c.cs, Member)}
}

func (c ZscoreKey) Cache() Cacheable {
	return Cacheable(c)
}

type ZscoreMember struct {
	cs []string
	cf uint32
}

func (c ZscoreMember) Build() Completed {
	return Completed(c)
}

func (c ZscoreMember) Cache() Cacheable {
	return Cacheable(c)
}

type Zunion struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZunionAggregateMax) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMax) Build() Completed {
	return Completed(c)
}

type ZunionAggregateMin struct {
	cs []string
	cf uint32
}

func (c ZunionAggregateMin) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMin) Build() Completed {
	return Completed(c)
}

type ZunionAggregateSum struct {
	cs []string
	cf uint32
}

func (c ZunionAggregateSum) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cf: c.cf, cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateSum) Build() Completed {
	return Completed(c)
}

type ZunionKey struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZunionNumkeys) Key(Key ...string) ZunionKey {
	return ZunionKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZunionWeights struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZunionWithscoresWithscores) Build() Completed {
	return Completed(c)
}

type Zunionstore struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZunionstoreAggregateMax) Build() Completed {
	return Completed(c)
}

type ZunionstoreAggregateMin struct {
	cs []string
	cf uint32
}

func (c ZunionstoreAggregateMin) Build() Completed {
	return Completed(c)
}

type ZunionstoreAggregateSum struct {
	cs []string
	cf uint32
}

func (c ZunionstoreAggregateSum) Build() Completed {
	return Completed(c)
}

type ZunionstoreDestination struct {
	cs []string
	cf uint32
}

func (c ZunionstoreDestination) Numkeys(Numkeys int64) ZunionstoreNumkeys {
	return ZunionstoreNumkeys{cf: c.cf, cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZunionstoreKey struct {
	cs []string
	cf uint32
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
	cf uint32
}

func (c ZunionstoreNumkeys) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cf: c.cf, cs: append(c.cs, Key...)}
}

type ZunionstoreWeights struct {
	cs []string
	cf uint32
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

