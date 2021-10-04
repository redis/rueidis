package cmds

import "strconv"

type AclCat struct {
	cs []string
}

func (c AclCat) Categoryname(Categoryname string) AclCatCategoryname {
	return AclCatCategoryname{cs: append(c.cs, Categoryname)}
}

func (c AclCat) Build() []string {
	return c.cs
}

func (b *Builder) AclCat() (c AclCat) {
	c.cs = append(b.get(), "ACL", "CAT")
	return
}

type AclCatCategoryname struct {
	cs []string
}

func (c AclCatCategoryname) Build() []string {
	return c.cs
}

type AclDeluser struct {
	cs []string
}

func (c AclDeluser) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cs: append(c.cs, Username...)}
}

func (b *Builder) AclDeluser() (c AclDeluser) {
	c.cs = append(b.get(), "ACL", "DELUSER")
	return
}

type AclDeluserUsername struct {
	cs []string
}

func (c AclDeluserUsername) Username(Username ...string) AclDeluserUsername {
	return AclDeluserUsername{cs: append(c.cs, Username...)}
}

func (c AclDeluserUsername) Build() []string {
	return c.cs
}

type AclGenpass struct {
	cs []string
}

func (c AclGenpass) Bits(Bits int64) AclGenpassBits {
	return AclGenpassBits{cs: append(c.cs, strconv.FormatInt(Bits, 10))}
}

func (c AclGenpass) Build() []string {
	return c.cs
}

func (b *Builder) AclGenpass() (c AclGenpass) {
	c.cs = append(b.get(), "ACL", "GENPASS")
	return
}

type AclGenpassBits struct {
	cs []string
}

func (c AclGenpassBits) Build() []string {
	return c.cs
}

type AclGetuser struct {
	cs []string
}

func (c AclGetuser) Username(Username string) AclGetuserUsername {
	return AclGetuserUsername{cs: append(c.cs, Username)}
}

func (b *Builder) AclGetuser() (c AclGetuser) {
	c.cs = append(b.get(), "ACL", "GETUSER")
	return
}

type AclGetuserUsername struct {
	cs []string
}

func (c AclGetuserUsername) Build() []string {
	return c.cs
}

type AclHelp struct {
	cs []string
}

func (c AclHelp) Build() []string {
	return c.cs
}

func (b *Builder) AclHelp() (c AclHelp) {
	c.cs = append(b.get(), "ACL", "HELP")
	return
}

type AclList struct {
	cs []string
}

func (c AclList) Build() []string {
	return c.cs
}

func (b *Builder) AclList() (c AclList) {
	c.cs = append(b.get(), "ACL", "LIST")
	return
}

type AclLoad struct {
	cs []string
}

func (c AclLoad) Build() []string {
	return c.cs
}

func (b *Builder) AclLoad() (c AclLoad) {
	c.cs = append(b.get(), "ACL", "LOAD")
	return
}

type AclLog struct {
	cs []string
}

func (c AclLog) CountOrReset(CountOrReset string) AclLogCountOrReset {
	return AclLogCountOrReset{cs: append(c.cs, CountOrReset)}
}

func (c AclLog) Build() []string {
	return c.cs
}

func (b *Builder) AclLog() (c AclLog) {
	c.cs = append(b.get(), "ACL", "LOG")
	return
}

type AclLogCountOrReset struct {
	cs []string
}

func (c AclLogCountOrReset) Build() []string {
	return c.cs
}

type AclSave struct {
	cs []string
}

func (c AclSave) Build() []string {
	return c.cs
}

func (b *Builder) AclSave() (c AclSave) {
	c.cs = append(b.get(), "ACL", "SAVE")
	return
}

type AclSetuser struct {
	cs []string
}

func (c AclSetuser) Username(Username string) AclSetuserUsername {
	return AclSetuserUsername{cs: append(c.cs, Username)}
}

func (b *Builder) AclSetuser() (c AclSetuser) {
	c.cs = append(b.get(), "ACL", "SETUSER")
	return
}

type AclSetuserRule struct {
	cs []string
}

func (c AclSetuserRule) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cs: append(c.cs, Rule...)}
}

func (c AclSetuserRule) Build() []string {
	return c.cs
}

type AclSetuserUsername struct {
	cs []string
}

func (c AclSetuserUsername) Rule(Rule ...string) AclSetuserRule {
	return AclSetuserRule{cs: append(c.cs, Rule...)}
}

func (c AclSetuserUsername) Build() []string {
	return c.cs
}

type AclUsers struct {
	cs []string
}

func (c AclUsers) Build() []string {
	return c.cs
}

func (b *Builder) AclUsers() (c AclUsers) {
	c.cs = append(b.get(), "ACL", "USERS")
	return
}

type AclWhoami struct {
	cs []string
}

func (c AclWhoami) Build() []string {
	return c.cs
}

func (b *Builder) AclWhoami() (c AclWhoami) {
	c.cs = append(b.get(), "ACL", "WHOAMI")
	return
}

type Append struct {
	cs []string
}

func (c Append) Key(Key string) AppendKey {
	return AppendKey{cs: append(c.cs, Key)}
}

func (b *Builder) Append() (c Append) {
	c.cs = append(b.get(), "APPEND")
	return
}

type AppendKey struct {
	cs []string
}

func (c AppendKey) Value(Value string) AppendValue {
	return AppendValue{cs: append(c.cs, Value)}
}

type AppendValue struct {
	cs []string
}

func (c AppendValue) Build() []string {
	return c.cs
}

type Asking struct {
	cs []string
}

func (c Asking) Build() []string {
	return c.cs
}

func (b *Builder) Asking() (c Asking) {
	c.cs = append(b.get(), "ASKING")
	return
}

type Auth struct {
	cs []string
}

func (c Auth) Username(Username string) AuthUsername {
	return AuthUsername{cs: append(c.cs, Username)}
}

func (c Auth) Password(Password string) AuthPassword {
	return AuthPassword{cs: append(c.cs, Password)}
}

func (b *Builder) Auth() (c Auth) {
	c.cs = append(b.get(), "AUTH")
	return
}

type AuthPassword struct {
	cs []string
}

func (c AuthPassword) Build() []string {
	return c.cs
}

type AuthUsername struct {
	cs []string
}

func (c AuthUsername) Password(Password string) AuthPassword {
	return AuthPassword{cs: append(c.cs, Password)}
}

type Bgrewriteaof struct {
	cs []string
}

func (c Bgrewriteaof) Build() []string {
	return c.cs
}

func (b *Builder) Bgrewriteaof() (c Bgrewriteaof) {
	c.cs = append(b.get(), "BGREWRITEAOF")
	return
}

type Bgsave struct {
	cs []string
}

func (c Bgsave) Schedule() BgsaveScheduleSchedule {
	return BgsaveScheduleSchedule{cs: append(c.cs, "SCHEDULE")}
}

func (c Bgsave) Build() []string {
	return c.cs
}

func (b *Builder) Bgsave() (c Bgsave) {
	c.cs = append(b.get(), "BGSAVE")
	return
}

type BgsaveScheduleSchedule struct {
	cs []string
}

func (c BgsaveScheduleSchedule) Build() []string {
	return c.cs
}

type Bitcount struct {
	cs []string
}

func (c Bitcount) Key(Key string) BitcountKey {
	return BitcountKey{cs: append(c.cs, Key)}
}

func (b *Builder) Bitcount() (c Bitcount) {
	c.cs = append(b.get(), "BITCOUNT")
	return
}

type BitcountKey struct {
	cs []string
}

func (c BitcountKey) StartEnd(Start int64, End int64) BitcountStartEnd {
	return BitcountStartEnd{cs: append(c.cs, strconv.FormatInt(Start, 10), strconv.FormatInt(End, 10))}
}

func (c BitcountKey) Build() []string {
	return c.cs
}

type BitcountStartEnd struct {
	cs []string
}

func (c BitcountStartEnd) Build() []string {
	return c.cs
}

type Bitfield struct {
	cs []string
}

func (c Bitfield) Key(Key string) BitfieldKey {
	return BitfieldKey{cs: append(c.cs, Key)}
}

func (b *Builder) Bitfield() (c Bitfield) {
	c.cs = append(b.get(), "BITFIELD")
	return
}

type BitfieldFail struct {
	cs []string
}

func (c BitfieldFail) Build() []string {
	return c.cs
}

type BitfieldGet struct {
	cs []string
}

func (c BitfieldGet) Set(Type string, Offset int64, Value int64) BitfieldSet {
	return BitfieldSet{cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10))}
}

func (c BitfieldGet) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c BitfieldGet) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP")}
}

func (c BitfieldGet) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT")}
}

func (c BitfieldGet) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL")}
}

func (c BitfieldGet) Build() []string {
	return c.cs
}

type BitfieldIncrby struct {
	cs []string
}

func (c BitfieldIncrby) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP")}
}

func (c BitfieldIncrby) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT")}
}

func (c BitfieldIncrby) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL")}
}

func (c BitfieldIncrby) Build() []string {
	return c.cs
}

type BitfieldKey struct {
	cs []string
}

func (c BitfieldKey) Get(Type string, Offset int64) BitfieldGet {
	return BitfieldGet{cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

func (c BitfieldKey) Set(Type string, Offset int64, Value int64) BitfieldSet {
	return BitfieldSet{cs: append(c.cs, "SET", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Value, 10))}
}

func (c BitfieldKey) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c BitfieldKey) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP")}
}

func (c BitfieldKey) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT")}
}

func (c BitfieldKey) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL")}
}

func (c BitfieldKey) Build() []string {
	return c.cs
}

type BitfieldRo struct {
	cs []string
}

func (c BitfieldRo) Key(Key string) BitfieldRoKey {
	return BitfieldRoKey{cs: append(c.cs, Key)}
}

func (b *Builder) BitfieldRo() (c BitfieldRo) {
	c.cs = append(b.get(), "BITFIELD_RO")
	return
}

type BitfieldRoGet struct {
	cs []string
}

func (c BitfieldRoGet) Build() []string {
	return c.cs
}

type BitfieldRoKey struct {
	cs []string
}

func (c BitfieldRoKey) Get(Type string, Offset int64) BitfieldRoGet {
	return BitfieldRoGet{cs: append(c.cs, "GET", Type, strconv.FormatInt(Offset, 10))}
}

type BitfieldSat struct {
	cs []string
}

func (c BitfieldSat) Build() []string {
	return c.cs
}

type BitfieldSet struct {
	cs []string
}

func (c BitfieldSet) Incrby(Type string, Offset int64, Increment int64) BitfieldIncrby {
	return BitfieldIncrby{cs: append(c.cs, "INCRBY", Type, strconv.FormatInt(Offset, 10), strconv.FormatInt(Increment, 10))}
}

func (c BitfieldSet) Wrap() BitfieldWrap {
	return BitfieldWrap{cs: append(c.cs, "WRAP")}
}

func (c BitfieldSet) Sat() BitfieldSat {
	return BitfieldSat{cs: append(c.cs, "SAT")}
}

func (c BitfieldSet) Fail() BitfieldFail {
	return BitfieldFail{cs: append(c.cs, "FAIL")}
}

func (c BitfieldSet) Build() []string {
	return c.cs
}

type BitfieldWrap struct {
	cs []string
}

func (c BitfieldWrap) Build() []string {
	return c.cs
}

type Bitop struct {
	cs []string
}

func (c Bitop) Operation(Operation string) BitopOperation {
	return BitopOperation{cs: append(c.cs, Operation)}
}

func (b *Builder) Bitop() (c Bitop) {
	c.cs = append(b.get(), "BITOP")
	return
}

type BitopDestkey struct {
	cs []string
}

func (c BitopDestkey) Key(Key ...string) BitopKey {
	return BitopKey{cs: append(c.cs, Key...)}
}

type BitopKey struct {
	cs []string
}

func (c BitopKey) Key(Key ...string) BitopKey {
	return BitopKey{cs: append(c.cs, Key...)}
}

func (c BitopKey) Build() []string {
	return c.cs
}

type BitopOperation struct {
	cs []string
}

func (c BitopOperation) Destkey(Destkey string) BitopDestkey {
	return BitopDestkey{cs: append(c.cs, Destkey)}
}

type Bitpos struct {
	cs []string
}

func (c Bitpos) Key(Key string) BitposKey {
	return BitposKey{cs: append(c.cs, Key)}
}

func (b *Builder) Bitpos() (c Bitpos) {
	c.cs = append(b.get(), "BITPOS")
	return
}

type BitposBit struct {
	cs []string
}

func (c BitposBit) Start(Start int64) BitposIndexStart {
	return BitposIndexStart{cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type BitposIndexEnd struct {
	cs []string
}

func (c BitposIndexEnd) Build() []string {
	return c.cs
}

type BitposIndexStart struct {
	cs []string
}

func (c BitposIndexStart) End(End int64) BitposIndexEnd {
	return BitposIndexEnd{cs: append(c.cs, strconv.FormatInt(End, 10))}
}

func (c BitposIndexStart) Build() []string {
	return c.cs
}

type BitposKey struct {
	cs []string
}

func (c BitposKey) Bit(Bit int64) BitposBit {
	return BitposBit{cs: append(c.cs, strconv.FormatInt(Bit, 10))}
}

type Blmove struct {
	cs []string
}

func (c Blmove) Source(Source string) BlmoveSource {
	return BlmoveSource{cs: append(c.cs, Source)}
}

func (b *Builder) Blmove() (c Blmove) {
	c.cs = append(b.get(), "BLMOVE")
	return
}

type BlmoveDestination struct {
	cs []string
}

func (c BlmoveDestination) Left() BlmoveWherefromLeft {
	return BlmoveWherefromLeft{cs: append(c.cs, "LEFT")}
}

func (c BlmoveDestination) Right() BlmoveWherefromRight {
	return BlmoveWherefromRight{cs: append(c.cs, "RIGHT")}
}

type BlmoveSource struct {
	cs []string
}

func (c BlmoveSource) Destination(Destination string) BlmoveDestination {
	return BlmoveDestination{cs: append(c.cs, Destination)}
}

type BlmoveTimeout struct {
	cs []string
}

func (c BlmoveTimeout) Build() []string {
	return c.cs
}

type BlmoveWherefromLeft struct {
	cs []string
}

func (c BlmoveWherefromLeft) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromLeft) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cs: append(c.cs, "RIGHT")}
}

type BlmoveWherefromRight struct {
	cs []string
}

func (c BlmoveWherefromRight) Left() BlmoveWheretoLeft {
	return BlmoveWheretoLeft{cs: append(c.cs, "LEFT")}
}

func (c BlmoveWherefromRight) Right() BlmoveWheretoRight {
	return BlmoveWheretoRight{cs: append(c.cs, "RIGHT")}
}

type BlmoveWheretoLeft struct {
	cs []string
}

func (c BlmoveWheretoLeft) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BlmoveWheretoRight struct {
	cs []string
}

func (c BlmoveWheretoRight) Timeout(Timeout float64) BlmoveTimeout {
	return BlmoveTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type Blmpop struct {
	cs []string
}

func (c Blmpop) Timeout(Timeout float64) BlmpopTimeout {
	return BlmpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (b *Builder) Blmpop() (c Blmpop) {
	c.cs = append(b.get(), "BLMPOP")
	return
}

type BlmpopCount struct {
	cs []string
}

func (c BlmpopCount) Build() []string {
	return c.cs
}

type BlmpopKey struct {
	cs []string
}

func (c BlmpopKey) Left() BlmpopWhereLeft {
	return BlmpopWhereLeft{cs: append(c.cs, "LEFT")}
}

func (c BlmpopKey) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cs: append(c.cs, "RIGHT")}
}

func (c BlmpopKey) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cs: append(c.cs, Key...)}
}

type BlmpopNumkeys struct {
	cs []string
}

func (c BlmpopNumkeys) Key(Key ...string) BlmpopKey {
	return BlmpopKey{cs: append(c.cs, Key...)}
}

func (c BlmpopNumkeys) Left() BlmpopWhereLeft {
	return BlmpopWhereLeft{cs: append(c.cs, "LEFT")}
}

func (c BlmpopNumkeys) Right() BlmpopWhereRight {
	return BlmpopWhereRight{cs: append(c.cs, "RIGHT")}
}

type BlmpopTimeout struct {
	cs []string
}

func (c BlmpopTimeout) Numkeys(Numkeys int64) BlmpopNumkeys {
	return BlmpopNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type BlmpopWhereLeft struct {
	cs []string
}

func (c BlmpopWhereLeft) Count(Count int64) BlmpopCount {
	return BlmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereLeft) Build() []string {
	return c.cs
}

type BlmpopWhereRight struct {
	cs []string
}

func (c BlmpopWhereRight) Count(Count int64) BlmpopCount {
	return BlmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c BlmpopWhereRight) Build() []string {
	return c.cs
}

type Blpop struct {
	cs []string
}

func (c Blpop) Key(Key ...string) BlpopKey {
	return BlpopKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Blpop() (c Blpop) {
	c.cs = append(b.get(), "BLPOP")
	return
}

type BlpopKey struct {
	cs []string
}

func (c BlpopKey) Timeout(Timeout float64) BlpopTimeout {
	return BlpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BlpopKey) Key(Key ...string) BlpopKey {
	return BlpopKey{cs: append(c.cs, Key...)}
}

type BlpopTimeout struct {
	cs []string
}

func (c BlpopTimeout) Build() []string {
	return c.cs
}

type Brpop struct {
	cs []string
}

func (c Brpop) Key(Key ...string) BrpopKey {
	return BrpopKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Brpop() (c Brpop) {
	c.cs = append(b.get(), "BRPOP")
	return
}

type BrpopKey struct {
	cs []string
}

func (c BrpopKey) Timeout(Timeout float64) BrpopTimeout {
	return BrpopTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BrpopKey) Key(Key ...string) BrpopKey {
	return BrpopKey{cs: append(c.cs, Key...)}
}

type BrpopTimeout struct {
	cs []string
}

func (c BrpopTimeout) Build() []string {
	return c.cs
}

type Brpoplpush struct {
	cs []string
}

func (c Brpoplpush) Source(Source string) BrpoplpushSource {
	return BrpoplpushSource{cs: append(c.cs, Source)}
}

func (b *Builder) Brpoplpush() (c Brpoplpush) {
	c.cs = append(b.get(), "BRPOPLPUSH")
	return
}

type BrpoplpushDestination struct {
	cs []string
}

func (c BrpoplpushDestination) Timeout(Timeout float64) BrpoplpushTimeout {
	return BrpoplpushTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

type BrpoplpushSource struct {
	cs []string
}

func (c BrpoplpushSource) Destination(Destination string) BrpoplpushDestination {
	return BrpoplpushDestination{cs: append(c.cs, Destination)}
}

type BrpoplpushTimeout struct {
	cs []string
}

func (c BrpoplpushTimeout) Build() []string {
	return c.cs
}

type Bzpopmax struct {
	cs []string
}

func (c Bzpopmax) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmax() (c Bzpopmax) {
	c.cs = append(b.get(), "BZPOPMAX")
	return
}

type BzpopmaxKey struct {
	cs []string
}

func (c BzpopmaxKey) Timeout(Timeout float64) BzpopmaxTimeout {
	return BzpopmaxTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopmaxKey) Key(Key ...string) BzpopmaxKey {
	return BzpopmaxKey{cs: append(c.cs, Key...)}
}

type BzpopmaxTimeout struct {
	cs []string
}

func (c BzpopmaxTimeout) Build() []string {
	return c.cs
}

type Bzpopmin struct {
	cs []string
}

func (c Bzpopmin) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Bzpopmin() (c Bzpopmin) {
	c.cs = append(b.get(), "BZPOPMIN")
	return
}

type BzpopminKey struct {
	cs []string
}

func (c BzpopminKey) Timeout(Timeout float64) BzpopminTimeout {
	return BzpopminTimeout{cs: append(c.cs, strconv.FormatFloat(Timeout, 'f', -1, 64))}
}

func (c BzpopminKey) Key(Key ...string) BzpopminKey {
	return BzpopminKey{cs: append(c.cs, Key...)}
}

type BzpopminTimeout struct {
	cs []string
}

func (c BzpopminTimeout) Build() []string {
	return c.cs
}

type ClientCaching struct {
	cs []string
}

func (c ClientCaching) Yes() ClientCachingModeYes {
	return ClientCachingModeYes{cs: append(c.cs, "YES")}
}

func (c ClientCaching) No() ClientCachingModeNo {
	return ClientCachingModeNo{cs: append(c.cs, "NO")}
}

func (b *Builder) ClientCaching() (c ClientCaching) {
	c.cs = append(b.get(), "CLIENT", "CACHING")
	return
}

type ClientCachingModeNo struct {
	cs []string
}

func (c ClientCachingModeNo) Build() []string {
	return c.cs
}

type ClientCachingModeYes struct {
	cs []string
}

func (c ClientCachingModeYes) Build() []string {
	return c.cs
}

type ClientGetname struct {
	cs []string
}

func (c ClientGetname) Build() []string {
	return c.cs
}

func (b *Builder) ClientGetname() (c ClientGetname) {
	c.cs = append(b.get(), "CLIENT", "GETNAME")
	return
}

type ClientGetredir struct {
	cs []string
}

func (c ClientGetredir) Build() []string {
	return c.cs
}

func (b *Builder) ClientGetredir() (c ClientGetredir) {
	c.cs = append(b.get(), "CLIENT", "GETREDIR")
	return
}

type ClientId struct {
	cs []string
}

func (c ClientId) Build() []string {
	return c.cs
}

func (b *Builder) ClientId() (c ClientId) {
	c.cs = append(b.get(), "CLIENT", "ID")
	return
}

type ClientInfo struct {
	cs []string
}

func (c ClientInfo) Build() []string {
	return c.cs
}

func (b *Builder) ClientInfo() (c ClientInfo) {
	c.cs = append(b.get(), "CLIENT", "INFO")
	return
}

type ClientKill struct {
	cs []string
}

func (c ClientKill) IpPort(IpPort string) ClientKillIpPort {
	return ClientKillIpPort{cs: append(c.cs, IpPort)}
}

func (c ClientKill) Id(ClientId int64) ClientKillId {
	return ClientKillId{cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10))}
}

func (c ClientKill) Normal() ClientKillNormal {
	return ClientKillNormal{cs: append(c.cs, "normal")}
}

func (c ClientKill) Master() ClientKillMaster {
	return ClientKillMaster{cs: append(c.cs, "master")}
}

func (c ClientKill) Slave() ClientKillSlave {
	return ClientKillSlave{cs: append(c.cs, "slave")}
}

func (c ClientKill) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cs: append(c.cs, "pubsub")}
}

func (c ClientKill) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username)}
}

func (c ClientKill) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKill) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKill) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKill) Build() []string {
	return c.cs
}

func (b *Builder) ClientKill() (c ClientKill) {
	c.cs = append(b.get(), "CLIENT", "KILL")
	return
}

type ClientKillAddr struct {
	cs []string
}

func (c ClientKillAddr) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillAddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillAddr) Build() []string {
	return c.cs
}

type ClientKillId struct {
	cs []string
}

func (c ClientKillId) Normal() ClientKillNormal {
	return ClientKillNormal{cs: append(c.cs, "normal")}
}

func (c ClientKillId) Master() ClientKillMaster {
	return ClientKillMaster{cs: append(c.cs, "master")}
}

func (c ClientKillId) Slave() ClientKillSlave {
	return ClientKillSlave{cs: append(c.cs, "slave")}
}

func (c ClientKillId) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cs: append(c.cs, "pubsub")}
}

func (c ClientKillId) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username)}
}

func (c ClientKillId) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillId) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillId) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillId) Build() []string {
	return c.cs
}

type ClientKillIpPort struct {
	cs []string
}

func (c ClientKillIpPort) Id(ClientId int64) ClientKillId {
	return ClientKillId{cs: append(c.cs, "ID", strconv.FormatInt(ClientId, 10))}
}

func (c ClientKillIpPort) Normal() ClientKillNormal {
	return ClientKillNormal{cs: append(c.cs, "normal")}
}

func (c ClientKillIpPort) Master() ClientKillMaster {
	return ClientKillMaster{cs: append(c.cs, "master")}
}

func (c ClientKillIpPort) Slave() ClientKillSlave {
	return ClientKillSlave{cs: append(c.cs, "slave")}
}

func (c ClientKillIpPort) Pubsub() ClientKillPubsub {
	return ClientKillPubsub{cs: append(c.cs, "pubsub")}
}

func (c ClientKillIpPort) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username)}
}

func (c ClientKillIpPort) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillIpPort) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillIpPort) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillIpPort) Build() []string {
	return c.cs
}

type ClientKillLaddr struct {
	cs []string
}

func (c ClientKillLaddr) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillLaddr) Build() []string {
	return c.cs
}

type ClientKillMaster struct {
	cs []string
}

func (c ClientKillMaster) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username)}
}

func (c ClientKillMaster) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillMaster) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillMaster) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillMaster) Build() []string {
	return c.cs
}

type ClientKillNormal struct {
	cs []string
}

func (c ClientKillNormal) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username)}
}

func (c ClientKillNormal) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillNormal) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillNormal) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillNormal) Build() []string {
	return c.cs
}

type ClientKillPubsub struct {
	cs []string
}

func (c ClientKillPubsub) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username)}
}

func (c ClientKillPubsub) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillPubsub) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillPubsub) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillPubsub) Build() []string {
	return c.cs
}

type ClientKillSkipme struct {
	cs []string
}

func (c ClientKillSkipme) Build() []string {
	return c.cs
}

type ClientKillSlave struct {
	cs []string
}

func (c ClientKillSlave) User(Username string) ClientKillUser {
	return ClientKillUser{cs: append(c.cs, "USER", Username)}
}

func (c ClientKillSlave) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillSlave) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillSlave) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillSlave) Build() []string {
	return c.cs
}

type ClientKillUser struct {
	cs []string
}

func (c ClientKillUser) Addr(IpPort string) ClientKillAddr {
	return ClientKillAddr{cs: append(c.cs, "ADDR", IpPort)}
}

func (c ClientKillUser) Laddr(IpPort string) ClientKillLaddr {
	return ClientKillLaddr{cs: append(c.cs, "LADDR", IpPort)}
}

func (c ClientKillUser) Skipme(YesNo string) ClientKillSkipme {
	return ClientKillSkipme{cs: append(c.cs, "SKIPME", YesNo)}
}

func (c ClientKillUser) Build() []string {
	return c.cs
}

type ClientList struct {
	cs []string
}

func (c ClientList) Normal() ClientListNormal {
	return ClientListNormal{cs: append(c.cs, "normal")}
}

func (c ClientList) Master() ClientListMaster {
	return ClientListMaster{cs: append(c.cs, "master")}
}

func (c ClientList) Replica() ClientListReplica {
	return ClientListReplica{cs: append(c.cs, "replica")}
}

func (c ClientList) Pubsub() ClientListPubsub {
	return ClientListPubsub{cs: append(c.cs, "pubsub")}
}

func (c ClientList) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID")}
}

func (b *Builder) ClientList() (c ClientList) {
	c.cs = append(b.get(), "CLIENT", "LIST")
	return
}

type ClientListIdClientId struct {
	cs []string
}

func (c ClientListIdClientId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cs: c.cs}
}

func (c ClientListIdClientId) Build() []string {
	return c.cs
}

type ClientListIdId struct {
	cs []string
}

func (c ClientListIdId) ClientId(ClientId ...int64) ClientListIdClientId {
	for _, n := range ClientId {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClientListIdClientId{cs: c.cs}
}

type ClientListMaster struct {
	cs []string
}

func (c ClientListMaster) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID")}
}

type ClientListNormal struct {
	cs []string
}

func (c ClientListNormal) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID")}
}

type ClientListPubsub struct {
	cs []string
}

func (c ClientListPubsub) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID")}
}

type ClientListReplica struct {
	cs []string
}

func (c ClientListReplica) Id() ClientListIdId {
	return ClientListIdId{cs: append(c.cs, "ID")}
}

type ClientNoEvict struct {
	cs []string
}

func (c ClientNoEvict) On() ClientNoEvictEnabledOn {
	return ClientNoEvictEnabledOn{cs: append(c.cs, "ON")}
}

func (c ClientNoEvict) Off() ClientNoEvictEnabledOff {
	return ClientNoEvictEnabledOff{cs: append(c.cs, "OFF")}
}

func (b *Builder) ClientNoEvict() (c ClientNoEvict) {
	c.cs = append(b.get(), "CLIENT", "NO-EVICT")
	return
}

type ClientNoEvictEnabledOff struct {
	cs []string
}

func (c ClientNoEvictEnabledOff) Build() []string {
	return c.cs
}

type ClientNoEvictEnabledOn struct {
	cs []string
}

func (c ClientNoEvictEnabledOn) Build() []string {
	return c.cs
}

type ClientPause struct {
	cs []string
}

func (c ClientPause) Timeout(Timeout int64) ClientPauseTimeout {
	return ClientPauseTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

func (b *Builder) ClientPause() (c ClientPause) {
	c.cs = append(b.get(), "CLIENT", "PAUSE")
	return
}

type ClientPauseModeAll struct {
	cs []string
}

func (c ClientPauseModeAll) Build() []string {
	return c.cs
}

type ClientPauseModeWrite struct {
	cs []string
}

func (c ClientPauseModeWrite) Build() []string {
	return c.cs
}

type ClientPauseTimeout struct {
	cs []string
}

func (c ClientPauseTimeout) Write() ClientPauseModeWrite {
	return ClientPauseModeWrite{cs: append(c.cs, "WRITE")}
}

func (c ClientPauseTimeout) All() ClientPauseModeAll {
	return ClientPauseModeAll{cs: append(c.cs, "ALL")}
}

func (c ClientPauseTimeout) Build() []string {
	return c.cs
}

type ClientReply struct {
	cs []string
}

func (c ClientReply) On() ClientReplyReplyModeOn {
	return ClientReplyReplyModeOn{cs: append(c.cs, "ON")}
}

func (c ClientReply) Off() ClientReplyReplyModeOff {
	return ClientReplyReplyModeOff{cs: append(c.cs, "OFF")}
}

func (c ClientReply) Skip() ClientReplyReplyModeSkip {
	return ClientReplyReplyModeSkip{cs: append(c.cs, "SKIP")}
}

func (b *Builder) ClientReply() (c ClientReply) {
	c.cs = append(b.get(), "CLIENT", "REPLY")
	return
}

type ClientReplyReplyModeOff struct {
	cs []string
}

func (c ClientReplyReplyModeOff) Build() []string {
	return c.cs
}

type ClientReplyReplyModeOn struct {
	cs []string
}

func (c ClientReplyReplyModeOn) Build() []string {
	return c.cs
}

type ClientReplyReplyModeSkip struct {
	cs []string
}

func (c ClientReplyReplyModeSkip) Build() []string {
	return c.cs
}

type ClientSetname struct {
	cs []string
}

func (c ClientSetname) ConnectionName(ConnectionName string) ClientSetnameConnectionName {
	return ClientSetnameConnectionName{cs: append(c.cs, ConnectionName)}
}

func (b *Builder) ClientSetname() (c ClientSetname) {
	c.cs = append(b.get(), "CLIENT", "SETNAME")
	return
}

type ClientSetnameConnectionName struct {
	cs []string
}

func (c ClientSetnameConnectionName) Build() []string {
	return c.cs
}

type ClientTracking struct {
	cs []string
}

func (c ClientTracking) On() ClientTrackingStatusOn {
	return ClientTrackingStatusOn{cs: append(c.cs, "ON")}
}

func (c ClientTracking) Off() ClientTrackingStatusOff {
	return ClientTrackingStatusOff{cs: append(c.cs, "OFF")}
}

func (b *Builder) ClientTracking() (c ClientTracking) {
	c.cs = append(b.get(), "CLIENT", "TRACKING")
	return
}

type ClientTrackingBcastBcast struct {
	cs []string
}

func (c ClientTrackingBcastBcast) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingBcastBcast) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingBcastBcast) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingBcastBcast) Build() []string {
	return c.cs
}

type ClientTrackingNoloopNoloop struct {
	cs []string
}

func (c ClientTrackingNoloopNoloop) Build() []string {
	return c.cs
}

type ClientTrackingOptinOptin struct {
	cs []string
}

func (c ClientTrackingOptinOptin) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingOptinOptin) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingOptinOptin) Build() []string {
	return c.cs
}

type ClientTrackingOptoutOptout struct {
	cs []string
}

func (c ClientTrackingOptoutOptout) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingOptoutOptout) Build() []string {
	return c.cs
}

type ClientTrackingPrefix struct {
	cs []string
}

func (c ClientTrackingPrefix) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingPrefix) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingPrefix) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingPrefix) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingPrefix) Prefix(Prefix ...string) ClientTrackingPrefix {
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingPrefix) Build() []string {
	return c.cs
}

type ClientTrackingRedirect struct {
	cs []string
}

func (c ClientTrackingRedirect) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingRedirect) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingRedirect) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingRedirect) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingRedirect) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingRedirect) Build() []string {
	return c.cs
}

type ClientTrackingStatusOff struct {
	cs []string
}

func (c ClientTrackingStatusOff) Redirect(ClientId int64) ClientTrackingRedirect {
	return ClientTrackingRedirect{cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10))}
}

func (c ClientTrackingStatusOff) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingStatusOff) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingStatusOff) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingStatusOff) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingStatusOff) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingStatusOff) Build() []string {
	return c.cs
}

type ClientTrackingStatusOn struct {
	cs []string
}

func (c ClientTrackingStatusOn) Redirect(ClientId int64) ClientTrackingRedirect {
	return ClientTrackingRedirect{cs: append(c.cs, "REDIRECT", strconv.FormatInt(ClientId, 10))}
}

func (c ClientTrackingStatusOn) Prefix(Prefix ...string) ClientTrackingPrefix {
	c.cs = append(c.cs, "PREFIX")
	return ClientTrackingPrefix{cs: append(c.cs, Prefix...)}
}

func (c ClientTrackingStatusOn) Bcast() ClientTrackingBcastBcast {
	return ClientTrackingBcastBcast{cs: append(c.cs, "BCAST")}
}

func (c ClientTrackingStatusOn) Optin() ClientTrackingOptinOptin {
	return ClientTrackingOptinOptin{cs: append(c.cs, "OPTIN")}
}

func (c ClientTrackingStatusOn) Optout() ClientTrackingOptoutOptout {
	return ClientTrackingOptoutOptout{cs: append(c.cs, "OPTOUT")}
}

func (c ClientTrackingStatusOn) Noloop() ClientTrackingNoloopNoloop {
	return ClientTrackingNoloopNoloop{cs: append(c.cs, "NOLOOP")}
}

func (c ClientTrackingStatusOn) Build() []string {
	return c.cs
}

type ClientTrackinginfo struct {
	cs []string
}

func (c ClientTrackinginfo) Build() []string {
	return c.cs
}

func (b *Builder) ClientTrackinginfo() (c ClientTrackinginfo) {
	c.cs = append(b.get(), "CLIENT", "TRACKINGINFO")
	return
}

type ClientUnblock struct {
	cs []string
}

func (c ClientUnblock) ClientId(ClientId int64) ClientUnblockClientId {
	return ClientUnblockClientId{cs: append(c.cs, strconv.FormatInt(ClientId, 10))}
}

func (b *Builder) ClientUnblock() (c ClientUnblock) {
	c.cs = append(b.get(), "CLIENT", "UNBLOCK")
	return
}

type ClientUnblockClientId struct {
	cs []string
}

func (c ClientUnblockClientId) Timeout() ClientUnblockUnblockTypeTimeout {
	return ClientUnblockUnblockTypeTimeout{cs: append(c.cs, "TIMEOUT")}
}

func (c ClientUnblockClientId) Error() ClientUnblockUnblockTypeError {
	return ClientUnblockUnblockTypeError{cs: append(c.cs, "ERROR")}
}

func (c ClientUnblockClientId) Build() []string {
	return c.cs
}

type ClientUnblockUnblockTypeError struct {
	cs []string
}

func (c ClientUnblockUnblockTypeError) Build() []string {
	return c.cs
}

type ClientUnblockUnblockTypeTimeout struct {
	cs []string
}

func (c ClientUnblockUnblockTypeTimeout) Build() []string {
	return c.cs
}

type ClientUnpause struct {
	cs []string
}

func (c ClientUnpause) Build() []string {
	return c.cs
}

func (b *Builder) ClientUnpause() (c ClientUnpause) {
	c.cs = append(b.get(), "CLIENT", "UNPAUSE")
	return
}

type ClusterAddslots struct {
	cs []string
}

func (c ClusterAddslots) Slot(Slot ...int64) ClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterAddslotsSlot{cs: c.cs}
}

func (b *Builder) ClusterAddslots() (c ClusterAddslots) {
	c.cs = append(b.get(), "CLUSTER", "ADDSLOTS")
	return
}

type ClusterAddslotsSlot struct {
	cs []string
}

func (c ClusterAddslotsSlot) Slot(Slot ...int64) ClusterAddslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterAddslotsSlot{cs: c.cs}
}

func (c ClusterAddslotsSlot) Build() []string {
	return c.cs
}

type ClusterBumpepoch struct {
	cs []string
}

func (c ClusterBumpepoch) Build() []string {
	return c.cs
}

func (b *Builder) ClusterBumpepoch() (c ClusterBumpepoch) {
	c.cs = append(b.get(), "CLUSTER", "BUMPEPOCH")
	return
}

type ClusterCountFailureReports struct {
	cs []string
}

func (c ClusterCountFailureReports) NodeId(NodeId string) ClusterCountFailureReportsNodeId {
	return ClusterCountFailureReportsNodeId{cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterCountFailureReports() (c ClusterCountFailureReports) {
	c.cs = append(b.get(), "CLUSTER", "COUNT-FAILURE-REPORTS")
	return
}

type ClusterCountFailureReportsNodeId struct {
	cs []string
}

func (c ClusterCountFailureReportsNodeId) Build() []string {
	return c.cs
}

type ClusterCountkeysinslot struct {
	cs []string
}

func (c ClusterCountkeysinslot) Slot(Slot int64) ClusterCountkeysinslotSlot {
	return ClusterCountkeysinslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *Builder) ClusterCountkeysinslot() (c ClusterCountkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "COUNTKEYSINSLOT")
	return
}

type ClusterCountkeysinslotSlot struct {
	cs []string
}

func (c ClusterCountkeysinslotSlot) Build() []string {
	return c.cs
}

type ClusterDelslots struct {
	cs []string
}

func (c ClusterDelslots) Slot(Slot ...int64) ClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterDelslotsSlot{cs: c.cs}
}

func (b *Builder) ClusterDelslots() (c ClusterDelslots) {
	c.cs = append(b.get(), "CLUSTER", "DELSLOTS")
	return
}

type ClusterDelslotsSlot struct {
	cs []string
}

func (c ClusterDelslotsSlot) Slot(Slot ...int64) ClusterDelslotsSlot {
	for _, n := range Slot {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ClusterDelslotsSlot{cs: c.cs}
}

func (c ClusterDelslotsSlot) Build() []string {
	return c.cs
}

type ClusterFailover struct {
	cs []string
}

func (c ClusterFailover) Force() ClusterFailoverOptionsForce {
	return ClusterFailoverOptionsForce{cs: append(c.cs, "FORCE")}
}

func (c ClusterFailover) Takeover() ClusterFailoverOptionsTakeover {
	return ClusterFailoverOptionsTakeover{cs: append(c.cs, "TAKEOVER")}
}

func (c ClusterFailover) Build() []string {
	return c.cs
}

func (b *Builder) ClusterFailover() (c ClusterFailover) {
	c.cs = append(b.get(), "CLUSTER", "FAILOVER")
	return
}

type ClusterFailoverOptionsForce struct {
	cs []string
}

func (c ClusterFailoverOptionsForce) Build() []string {
	return c.cs
}

type ClusterFailoverOptionsTakeover struct {
	cs []string
}

func (c ClusterFailoverOptionsTakeover) Build() []string {
	return c.cs
}

type ClusterFlushslots struct {
	cs []string
}

func (c ClusterFlushslots) Build() []string {
	return c.cs
}

func (b *Builder) ClusterFlushslots() (c ClusterFlushslots) {
	c.cs = append(b.get(), "CLUSTER", "FLUSHSLOTS")
	return
}

type ClusterForget struct {
	cs []string
}

func (c ClusterForget) NodeId(NodeId string) ClusterForgetNodeId {
	return ClusterForgetNodeId{cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterForget() (c ClusterForget) {
	c.cs = append(b.get(), "CLUSTER", "FORGET")
	return
}

type ClusterForgetNodeId struct {
	cs []string
}

func (c ClusterForgetNodeId) Build() []string {
	return c.cs
}

type ClusterGetkeysinslot struct {
	cs []string
}

func (c ClusterGetkeysinslot) Slot(Slot int64) ClusterGetkeysinslotSlot {
	return ClusterGetkeysinslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *Builder) ClusterGetkeysinslot() (c ClusterGetkeysinslot) {
	c.cs = append(b.get(), "CLUSTER", "GETKEYSINSLOT")
	return
}

type ClusterGetkeysinslotCount struct {
	cs []string
}

func (c ClusterGetkeysinslotCount) Build() []string {
	return c.cs
}

type ClusterGetkeysinslotSlot struct {
	cs []string
}

func (c ClusterGetkeysinslotSlot) Count(Count int64) ClusterGetkeysinslotCount {
	return ClusterGetkeysinslotCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type ClusterInfo struct {
	cs []string
}

func (c ClusterInfo) Build() []string {
	return c.cs
}

func (b *Builder) ClusterInfo() (c ClusterInfo) {
	c.cs = append(b.get(), "CLUSTER", "INFO")
	return
}

type ClusterKeyslot struct {
	cs []string
}

func (c ClusterKeyslot) Key(Key string) ClusterKeyslotKey {
	return ClusterKeyslotKey{cs: append(c.cs, Key)}
}

func (b *Builder) ClusterKeyslot() (c ClusterKeyslot) {
	c.cs = append(b.get(), "CLUSTER", "KEYSLOT")
	return
}

type ClusterKeyslotKey struct {
	cs []string
}

func (c ClusterKeyslotKey) Build() []string {
	return c.cs
}

type ClusterMeet struct {
	cs []string
}

func (c ClusterMeet) Ip(Ip string) ClusterMeetIp {
	return ClusterMeetIp{cs: append(c.cs, Ip)}
}

func (b *Builder) ClusterMeet() (c ClusterMeet) {
	c.cs = append(b.get(), "CLUSTER", "MEET")
	return
}

type ClusterMeetIp struct {
	cs []string
}

func (c ClusterMeetIp) Port(Port int64) ClusterMeetPort {
	return ClusterMeetPort{cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type ClusterMeetPort struct {
	cs []string
}

func (c ClusterMeetPort) Build() []string {
	return c.cs
}

type ClusterMyid struct {
	cs []string
}

func (c ClusterMyid) Build() []string {
	return c.cs
}

func (b *Builder) ClusterMyid() (c ClusterMyid) {
	c.cs = append(b.get(), "CLUSTER", "MYID")
	return
}

type ClusterNodes struct {
	cs []string
}

func (c ClusterNodes) Build() []string {
	return c.cs
}

func (b *Builder) ClusterNodes() (c ClusterNodes) {
	c.cs = append(b.get(), "CLUSTER", "NODES")
	return
}

type ClusterReplicas struct {
	cs []string
}

func (c ClusterReplicas) NodeId(NodeId string) ClusterReplicasNodeId {
	return ClusterReplicasNodeId{cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterReplicas() (c ClusterReplicas) {
	c.cs = append(b.get(), "CLUSTER", "REPLICAS")
	return
}

type ClusterReplicasNodeId struct {
	cs []string
}

func (c ClusterReplicasNodeId) Build() []string {
	return c.cs
}

type ClusterReplicate struct {
	cs []string
}

func (c ClusterReplicate) NodeId(NodeId string) ClusterReplicateNodeId {
	return ClusterReplicateNodeId{cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterReplicate() (c ClusterReplicate) {
	c.cs = append(b.get(), "CLUSTER", "REPLICATE")
	return
}

type ClusterReplicateNodeId struct {
	cs []string
}

func (c ClusterReplicateNodeId) Build() []string {
	return c.cs
}

type ClusterReset struct {
	cs []string
}

func (c ClusterReset) Hard() ClusterResetResetTypeHard {
	return ClusterResetResetTypeHard{cs: append(c.cs, "HARD")}
}

func (c ClusterReset) Soft() ClusterResetResetTypeSoft {
	return ClusterResetResetTypeSoft{cs: append(c.cs, "SOFT")}
}

func (c ClusterReset) Build() []string {
	return c.cs
}

func (b *Builder) ClusterReset() (c ClusterReset) {
	c.cs = append(b.get(), "CLUSTER", "RESET")
	return
}

type ClusterResetResetTypeHard struct {
	cs []string
}

func (c ClusterResetResetTypeHard) Build() []string {
	return c.cs
}

type ClusterResetResetTypeSoft struct {
	cs []string
}

func (c ClusterResetResetTypeSoft) Build() []string {
	return c.cs
}

type ClusterSaveconfig struct {
	cs []string
}

func (c ClusterSaveconfig) Build() []string {
	return c.cs
}

func (b *Builder) ClusterSaveconfig() (c ClusterSaveconfig) {
	c.cs = append(b.get(), "CLUSTER", "SAVECONFIG")
	return
}

type ClusterSetConfigEpoch struct {
	cs []string
}

func (c ClusterSetConfigEpoch) ConfigEpoch(ConfigEpoch int64) ClusterSetConfigEpochConfigEpoch {
	return ClusterSetConfigEpochConfigEpoch{cs: append(c.cs, strconv.FormatInt(ConfigEpoch, 10))}
}

func (b *Builder) ClusterSetConfigEpoch() (c ClusterSetConfigEpoch) {
	c.cs = append(b.get(), "CLUSTER", "SET-CONFIG-EPOCH")
	return
}

type ClusterSetConfigEpochConfigEpoch struct {
	cs []string
}

func (c ClusterSetConfigEpochConfigEpoch) Build() []string {
	return c.cs
}

type ClusterSetslot struct {
	cs []string
}

func (c ClusterSetslot) Slot(Slot int64) ClusterSetslotSlot {
	return ClusterSetslotSlot{cs: append(c.cs, strconv.FormatInt(Slot, 10))}
}

func (b *Builder) ClusterSetslot() (c ClusterSetslot) {
	c.cs = append(b.get(), "CLUSTER", "SETSLOT")
	return
}

type ClusterSetslotNodeId struct {
	cs []string
}

func (c ClusterSetslotNodeId) Build() []string {
	return c.cs
}

type ClusterSetslotSlot struct {
	cs []string
}

func (c ClusterSetslotSlot) Importing() ClusterSetslotSubcommandImporting {
	return ClusterSetslotSubcommandImporting{cs: append(c.cs, "IMPORTING")}
}

func (c ClusterSetslotSlot) Migrating() ClusterSetslotSubcommandMigrating {
	return ClusterSetslotSubcommandMigrating{cs: append(c.cs, "MIGRATING")}
}

func (c ClusterSetslotSlot) Stable() ClusterSetslotSubcommandStable {
	return ClusterSetslotSubcommandStable{cs: append(c.cs, "STABLE")}
}

func (c ClusterSetslotSlot) Node() ClusterSetslotSubcommandNode {
	return ClusterSetslotSubcommandNode{cs: append(c.cs, "NODE")}
}

type ClusterSetslotSubcommandImporting struct {
	cs []string
}

func (c ClusterSetslotSubcommandImporting) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandImporting) Build() []string {
	return c.cs
}

type ClusterSetslotSubcommandMigrating struct {
	cs []string
}

func (c ClusterSetslotSubcommandMigrating) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandMigrating) Build() []string {
	return c.cs
}

type ClusterSetslotSubcommandNode struct {
	cs []string
}

func (c ClusterSetslotSubcommandNode) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandNode) Build() []string {
	return c.cs
}

type ClusterSetslotSubcommandStable struct {
	cs []string
}

func (c ClusterSetslotSubcommandStable) NodeId(NodeId string) ClusterSetslotNodeId {
	return ClusterSetslotNodeId{cs: append(c.cs, NodeId)}
}

func (c ClusterSetslotSubcommandStable) Build() []string {
	return c.cs
}

type ClusterSlaves struct {
	cs []string
}

func (c ClusterSlaves) NodeId(NodeId string) ClusterSlavesNodeId {
	return ClusterSlavesNodeId{cs: append(c.cs, NodeId)}
}

func (b *Builder) ClusterSlaves() (c ClusterSlaves) {
	c.cs = append(b.get(), "CLUSTER", "SLAVES")
	return
}

type ClusterSlavesNodeId struct {
	cs []string
}

func (c ClusterSlavesNodeId) Build() []string {
	return c.cs
}

type ClusterSlots struct {
	cs []string
}

func (c ClusterSlots) Build() []string {
	return c.cs
}

func (b *Builder) ClusterSlots() (c ClusterSlots) {
	c.cs = append(b.get(), "CLUSTER", "SLOTS")
	return
}

type Command struct {
	cs []string
}

func (c Command) Build() []string {
	return c.cs
}

func (b *Builder) Command() (c Command) {
	c.cs = append(b.get(), "COMMAND")
	return
}

type CommandCount struct {
	cs []string
}

func (c CommandCount) Build() []string {
	return c.cs
}

func (b *Builder) CommandCount() (c CommandCount) {
	c.cs = append(b.get(), "COMMAND", "COUNT")
	return
}

type CommandGetkeys struct {
	cs []string
}

func (c CommandGetkeys) Build() []string {
	return c.cs
}

func (b *Builder) CommandGetkeys() (c CommandGetkeys) {
	c.cs = append(b.get(), "COMMAND", "GETKEYS")
	return
}

type CommandInfo struct {
	cs []string
}

func (c CommandInfo) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cs: append(c.cs, CommandName...)}
}

func (b *Builder) CommandInfo() (c CommandInfo) {
	c.cs = append(b.get(), "COMMAND", "INFO")
	return
}

type CommandInfoCommandName struct {
	cs []string
}

func (c CommandInfoCommandName) CommandName(CommandName ...string) CommandInfoCommandName {
	return CommandInfoCommandName{cs: append(c.cs, CommandName...)}
}

func (c CommandInfoCommandName) Build() []string {
	return c.cs
}

type ConfigGet struct {
	cs []string
}

func (c ConfigGet) Parameter(Parameter string) ConfigGetParameter {
	return ConfigGetParameter{cs: append(c.cs, Parameter)}
}

func (b *Builder) ConfigGet() (c ConfigGet) {
	c.cs = append(b.get(), "CONFIG", "GET")
	return
}

type ConfigGetParameter struct {
	cs []string
}

func (c ConfigGetParameter) Build() []string {
	return c.cs
}

type ConfigResetstat struct {
	cs []string
}

func (c ConfigResetstat) Build() []string {
	return c.cs
}

func (b *Builder) ConfigResetstat() (c ConfigResetstat) {
	c.cs = append(b.get(), "CONFIG", "RESETSTAT")
	return
}

type ConfigRewrite struct {
	cs []string
}

func (c ConfigRewrite) Build() []string {
	return c.cs
}

func (b *Builder) ConfigRewrite() (c ConfigRewrite) {
	c.cs = append(b.get(), "CONFIG", "REWRITE")
	return
}

type ConfigSet struct {
	cs []string
}

func (c ConfigSet) Parameter(Parameter string) ConfigSetParameter {
	return ConfigSetParameter{cs: append(c.cs, Parameter)}
}

func (b *Builder) ConfigSet() (c ConfigSet) {
	c.cs = append(b.get(), "CONFIG", "SET")
	return
}

type ConfigSetParameter struct {
	cs []string
}

func (c ConfigSetParameter) Value(Value string) ConfigSetValue {
	return ConfigSetValue{cs: append(c.cs, Value)}
}

type ConfigSetValue struct {
	cs []string
}

func (c ConfigSetValue) Build() []string {
	return c.cs
}

type Copy struct {
	cs []string
}

func (c Copy) Source(Source string) CopySource {
	return CopySource{cs: append(c.cs, Source)}
}

func (b *Builder) Copy() (c Copy) {
	c.cs = append(b.get(), "COPY")
	return
}

type CopyDb struct {
	cs []string
}

func (c CopyDb) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cs: append(c.cs, "REPLACE")}
}

func (c CopyDb) Build() []string {
	return c.cs
}

type CopyDestination struct {
	cs []string
}

func (c CopyDestination) Db(DestinationDb int64) CopyDb {
	return CopyDb{cs: append(c.cs, "DB", strconv.FormatInt(DestinationDb, 10))}
}

func (c CopyDestination) Replace() CopyReplaceReplace {
	return CopyReplaceReplace{cs: append(c.cs, "REPLACE")}
}

func (c CopyDestination) Build() []string {
	return c.cs
}

type CopyReplaceReplace struct {
	cs []string
}

func (c CopyReplaceReplace) Build() []string {
	return c.cs
}

type CopySource struct {
	cs []string
}

func (c CopySource) Destination(Destination string) CopyDestination {
	return CopyDestination{cs: append(c.cs, Destination)}
}

type Dbsize struct {
	cs []string
}

func (c Dbsize) Build() []string {
	return c.cs
}

func (b *Builder) Dbsize() (c Dbsize) {
	c.cs = append(b.get(), "DBSIZE")
	return
}

type DebugObject struct {
	cs []string
}

func (c DebugObject) Key(Key string) DebugObjectKey {
	return DebugObjectKey{cs: append(c.cs, Key)}
}

func (b *Builder) DebugObject() (c DebugObject) {
	c.cs = append(b.get(), "DEBUG", "OBJECT")
	return
}

type DebugObjectKey struct {
	cs []string
}

func (c DebugObjectKey) Build() []string {
	return c.cs
}

type DebugSegfault struct {
	cs []string
}

func (c DebugSegfault) Build() []string {
	return c.cs
}

func (b *Builder) DebugSegfault() (c DebugSegfault) {
	c.cs = append(b.get(), "DEBUG", "SEGFAULT")
	return
}

type Decr struct {
	cs []string
}

func (c Decr) Key(Key string) DecrKey {
	return DecrKey{cs: append(c.cs, Key)}
}

func (b *Builder) Decr() (c Decr) {
	c.cs = append(b.get(), "DECR")
	return
}

type DecrKey struct {
	cs []string
}

func (c DecrKey) Build() []string {
	return c.cs
}

type Decrby struct {
	cs []string
}

func (c Decrby) Key(Key string) DecrbyKey {
	return DecrbyKey{cs: append(c.cs, Key)}
}

func (b *Builder) Decrby() (c Decrby) {
	c.cs = append(b.get(), "DECRBY")
	return
}

type DecrbyDecrement struct {
	cs []string
}

func (c DecrbyDecrement) Build() []string {
	return c.cs
}

type DecrbyKey struct {
	cs []string
}

func (c DecrbyKey) Decrement(Decrement int64) DecrbyDecrement {
	return DecrbyDecrement{cs: append(c.cs, strconv.FormatInt(Decrement, 10))}
}

type Del struct {
	cs []string
}

func (c Del) Key(Key ...string) DelKey {
	return DelKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Del() (c Del) {
	c.cs = append(b.get(), "DEL")
	return
}

type DelKey struct {
	cs []string
}

func (c DelKey) Key(Key ...string) DelKey {
	return DelKey{cs: append(c.cs, Key...)}
}

func (c DelKey) Build() []string {
	return c.cs
}

type Discard struct {
	cs []string
}

func (c Discard) Build() []string {
	return c.cs
}

func (b *Builder) Discard() (c Discard) {
	c.cs = append(b.get(), "DISCARD")
	return
}

type Dump struct {
	cs []string
}

func (c Dump) Key(Key string) DumpKey {
	return DumpKey{cs: append(c.cs, Key)}
}

func (b *Builder) Dump() (c Dump) {
	c.cs = append(b.get(), "DUMP")
	return
}

type DumpKey struct {
	cs []string
}

func (c DumpKey) Build() []string {
	return c.cs
}

type Echo struct {
	cs []string
}

func (c Echo) Message(Message string) EchoMessage {
	return EchoMessage{cs: append(c.cs, Message)}
}

func (b *Builder) Echo() (c Echo) {
	c.cs = append(b.get(), "ECHO")
	return
}

type EchoMessage struct {
	cs []string
}

func (c EchoMessage) Build() []string {
	return c.cs
}

type Eval struct {
	cs []string
}

func (c Eval) Script(Script string) EvalScript {
	return EvalScript{cs: append(c.cs, Script)}
}

func (b *Builder) Eval() (c Eval) {
	c.cs = append(b.get(), "EVAL")
	return
}

type EvalArg struct {
	cs []string
}

func (c EvalArg) Arg(Arg ...string) EvalArg {
	return EvalArg{cs: append(c.cs, Arg...)}
}

func (c EvalArg) Build() []string {
	return c.cs
}

type EvalKey struct {
	cs []string
}

func (c EvalKey) Arg(Arg ...string) EvalArg {
	return EvalArg{cs: append(c.cs, Arg...)}
}

func (c EvalKey) Key(Key ...string) EvalKey {
	return EvalKey{cs: append(c.cs, Key...)}
}

func (c EvalKey) Build() []string {
	return c.cs
}

type EvalNumkeys struct {
	cs []string
}

func (c EvalNumkeys) Key(Key ...string) EvalKey {
	return EvalKey{cs: append(c.cs, Key...)}
}

func (c EvalNumkeys) Arg(Arg ...string) EvalArg {
	return EvalArg{cs: append(c.cs, Arg...)}
}

func (c EvalNumkeys) Build() []string {
	return c.cs
}

type EvalRo struct {
	cs []string
}

func (c EvalRo) Script(Script string) EvalRoScript {
	return EvalRoScript{cs: append(c.cs, Script)}
}

func (b *Builder) EvalRo() (c EvalRo) {
	c.cs = append(b.get(), "EVAL_RO")
	return
}

type EvalRoArg struct {
	cs []string
}

func (c EvalRoArg) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cs: append(c.cs, Arg...)}
}

func (c EvalRoArg) Build() []string {
	return c.cs
}

type EvalRoKey struct {
	cs []string
}

func (c EvalRoKey) Arg(Arg ...string) EvalRoArg {
	return EvalRoArg{cs: append(c.cs, Arg...)}
}

func (c EvalRoKey) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cs: append(c.cs, Key...)}
}

type EvalRoNumkeys struct {
	cs []string
}

func (c EvalRoNumkeys) Key(Key ...string) EvalRoKey {
	return EvalRoKey{cs: append(c.cs, Key...)}
}

type EvalRoScript struct {
	cs []string
}

func (c EvalRoScript) Numkeys(Numkeys int64) EvalRoNumkeys {
	return EvalRoNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalScript struct {
	cs []string
}

func (c EvalScript) Numkeys(Numkeys int64) EvalNumkeys {
	return EvalNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type Evalsha struct {
	cs []string
}

func (c Evalsha) Sha1(Sha1 string) EvalshaSha1 {
	return EvalshaSha1{cs: append(c.cs, Sha1)}
}

func (b *Builder) Evalsha() (c Evalsha) {
	c.cs = append(b.get(), "EVALSHA")
	return
}

type EvalshaArg struct {
	cs []string
}

func (c EvalshaArg) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cs: append(c.cs, Arg...)}
}

func (c EvalshaArg) Build() []string {
	return c.cs
}

type EvalshaKey struct {
	cs []string
}

func (c EvalshaKey) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cs: append(c.cs, Arg...)}
}

func (c EvalshaKey) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cs: append(c.cs, Key...)}
}

func (c EvalshaKey) Build() []string {
	return c.cs
}

type EvalshaNumkeys struct {
	cs []string
}

func (c EvalshaNumkeys) Key(Key ...string) EvalshaKey {
	return EvalshaKey{cs: append(c.cs, Key...)}
}

func (c EvalshaNumkeys) Arg(Arg ...string) EvalshaArg {
	return EvalshaArg{cs: append(c.cs, Arg...)}
}

func (c EvalshaNumkeys) Build() []string {
	return c.cs
}

type EvalshaRo struct {
	cs []string
}

func (c EvalshaRo) Sha1(Sha1 string) EvalshaRoSha1 {
	return EvalshaRoSha1{cs: append(c.cs, Sha1)}
}

func (b *Builder) EvalshaRo() (c EvalshaRo) {
	c.cs = append(b.get(), "EVALSHA_RO")
	return
}

type EvalshaRoArg struct {
	cs []string
}

func (c EvalshaRoArg) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cs: append(c.cs, Arg...)}
}

func (c EvalshaRoArg) Build() []string {
	return c.cs
}

type EvalshaRoKey struct {
	cs []string
}

func (c EvalshaRoKey) Arg(Arg ...string) EvalshaRoArg {
	return EvalshaRoArg{cs: append(c.cs, Arg...)}
}

func (c EvalshaRoKey) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cs: append(c.cs, Key...)}
}

type EvalshaRoNumkeys struct {
	cs []string
}

func (c EvalshaRoNumkeys) Key(Key ...string) EvalshaRoKey {
	return EvalshaRoKey{cs: append(c.cs, Key...)}
}

type EvalshaRoSha1 struct {
	cs []string
}

func (c EvalshaRoSha1) Numkeys(Numkeys int64) EvalshaRoNumkeys {
	return EvalshaRoNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type EvalshaSha1 struct {
	cs []string
}

func (c EvalshaSha1) Numkeys(Numkeys int64) EvalshaNumkeys {
	return EvalshaNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type Exec struct {
	cs []string
}

func (c Exec) Build() []string {
	return c.cs
}

func (b *Builder) Exec() (c Exec) {
	c.cs = append(b.get(), "EXEC")
	return
}

type Exists struct {
	cs []string
}

func (c Exists) Key(Key ...string) ExistsKey {
	return ExistsKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Exists() (c Exists) {
	c.cs = append(b.get(), "EXISTS")
	return
}

type ExistsKey struct {
	cs []string
}

func (c ExistsKey) Key(Key ...string) ExistsKey {
	return ExistsKey{cs: append(c.cs, Key...)}
}

func (c ExistsKey) Build() []string {
	return c.cs
}

type Expire struct {
	cs []string
}

func (c Expire) Key(Key string) ExpireKey {
	return ExpireKey{cs: append(c.cs, Key)}
}

func (b *Builder) Expire() (c Expire) {
	c.cs = append(b.get(), "EXPIRE")
	return
}

type ExpireConditionGt struct {
	cs []string
}

func (c ExpireConditionGt) Build() []string {
	return c.cs
}

type ExpireConditionLt struct {
	cs []string
}

func (c ExpireConditionLt) Build() []string {
	return c.cs
}

type ExpireConditionNx struct {
	cs []string
}

func (c ExpireConditionNx) Build() []string {
	return c.cs
}

type ExpireConditionXx struct {
	cs []string
}

func (c ExpireConditionXx) Build() []string {
	return c.cs
}

type ExpireKey struct {
	cs []string
}

func (c ExpireKey) Seconds(Seconds int64) ExpireSeconds {
	return ExpireSeconds{cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type ExpireSeconds struct {
	cs []string
}

func (c ExpireSeconds) Nx() ExpireConditionNx {
	return ExpireConditionNx{cs: append(c.cs, "NX")}
}

func (c ExpireSeconds) Xx() ExpireConditionXx {
	return ExpireConditionXx{cs: append(c.cs, "XX")}
}

func (c ExpireSeconds) Gt() ExpireConditionGt {
	return ExpireConditionGt{cs: append(c.cs, "GT")}
}

func (c ExpireSeconds) Lt() ExpireConditionLt {
	return ExpireConditionLt{cs: append(c.cs, "LT")}
}

func (c ExpireSeconds) Build() []string {
	return c.cs
}

type Expireat struct {
	cs []string
}

func (c Expireat) Key(Key string) ExpireatKey {
	return ExpireatKey{cs: append(c.cs, Key)}
}

func (b *Builder) Expireat() (c Expireat) {
	c.cs = append(b.get(), "EXPIREAT")
	return
}

type ExpireatConditionGt struct {
	cs []string
}

func (c ExpireatConditionGt) Build() []string {
	return c.cs
}

type ExpireatConditionLt struct {
	cs []string
}

func (c ExpireatConditionLt) Build() []string {
	return c.cs
}

type ExpireatConditionNx struct {
	cs []string
}

func (c ExpireatConditionNx) Build() []string {
	return c.cs
}

type ExpireatConditionXx struct {
	cs []string
}

func (c ExpireatConditionXx) Build() []string {
	return c.cs
}

type ExpireatKey struct {
	cs []string
}

func (c ExpireatKey) Timestamp(Timestamp int64) ExpireatTimestamp {
	return ExpireatTimestamp{cs: append(c.cs, strconv.FormatInt(Timestamp, 10))}
}

type ExpireatTimestamp struct {
	cs []string
}

func (c ExpireatTimestamp) Nx() ExpireatConditionNx {
	return ExpireatConditionNx{cs: append(c.cs, "NX")}
}

func (c ExpireatTimestamp) Xx() ExpireatConditionXx {
	return ExpireatConditionXx{cs: append(c.cs, "XX")}
}

func (c ExpireatTimestamp) Gt() ExpireatConditionGt {
	return ExpireatConditionGt{cs: append(c.cs, "GT")}
}

func (c ExpireatTimestamp) Lt() ExpireatConditionLt {
	return ExpireatConditionLt{cs: append(c.cs, "LT")}
}

func (c ExpireatTimestamp) Build() []string {
	return c.cs
}

type Expiretime struct {
	cs []string
}

func (c Expiretime) Key(Key string) ExpiretimeKey {
	return ExpiretimeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Expiretime() (c Expiretime) {
	c.cs = append(b.get(), "EXPIRETIME")
	return
}

type ExpiretimeKey struct {
	cs []string
}

func (c ExpiretimeKey) Build() []string {
	return c.cs
}

type Failover struct {
	cs []string
}

func (c Failover) To() FailoverTargetTo {
	return FailoverTargetTo{cs: append(c.cs, "TO")}
}

func (c Failover) Abort() FailoverAbort {
	return FailoverAbort{cs: append(c.cs, "ABORT")}
}

func (c Failover) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (b *Builder) Failover() (c Failover) {
	c.cs = append(b.get(), "FAILOVER")
	return
}

type FailoverAbort struct {
	cs []string
}

func (c FailoverAbort) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverAbort) Build() []string {
	return c.cs
}

type FailoverTargetForce struct {
	cs []string
}

func (c FailoverTargetForce) Abort() FailoverAbort {
	return FailoverAbort{cs: append(c.cs, "ABORT")}
}

func (c FailoverTargetForce) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverTargetForce) Build() []string {
	return c.cs
}

type FailoverTargetHost struct {
	cs []string
}

func (c FailoverTargetHost) Port(Port int64) FailoverTargetPort {
	return FailoverTargetPort{cs: append(c.cs, strconv.FormatInt(Port, 10))}
}

type FailoverTargetPort struct {
	cs []string
}

func (c FailoverTargetPort) Force() FailoverTargetForce {
	return FailoverTargetForce{cs: append(c.cs, "FORCE")}
}

func (c FailoverTargetPort) Abort() FailoverAbort {
	return FailoverAbort{cs: append(c.cs, "ABORT")}
}

func (c FailoverTargetPort) Timeout(Milliseconds int64) FailoverTimeout {
	return FailoverTimeout{cs: append(c.cs, "TIMEOUT", strconv.FormatInt(Milliseconds, 10))}
}

func (c FailoverTargetPort) Build() []string {
	return c.cs
}

type FailoverTargetTo struct {
	cs []string
}

func (c FailoverTargetTo) Host(Host string) FailoverTargetHost {
	return FailoverTargetHost{cs: append(c.cs, Host)}
}

type FailoverTimeout struct {
	cs []string
}

func (c FailoverTimeout) Build() []string {
	return c.cs
}

type Flushall struct {
	cs []string
}

func (c Flushall) Async() FlushallAsyncAsync {
	return FlushallAsyncAsync{cs: append(c.cs, "ASYNC")}
}

func (c Flushall) Sync() FlushallAsyncSync {
	return FlushallAsyncSync{cs: append(c.cs, "SYNC")}
}

func (c Flushall) Build() []string {
	return c.cs
}

func (b *Builder) Flushall() (c Flushall) {
	c.cs = append(b.get(), "FLUSHALL")
	return
}

type FlushallAsyncAsync struct {
	cs []string
}

func (c FlushallAsyncAsync) Build() []string {
	return c.cs
}

type FlushallAsyncSync struct {
	cs []string
}

func (c FlushallAsyncSync) Build() []string {
	return c.cs
}

type Flushdb struct {
	cs []string
}

func (c Flushdb) Async() FlushdbAsyncAsync {
	return FlushdbAsyncAsync{cs: append(c.cs, "ASYNC")}
}

func (c Flushdb) Sync() FlushdbAsyncSync {
	return FlushdbAsyncSync{cs: append(c.cs, "SYNC")}
}

func (c Flushdb) Build() []string {
	return c.cs
}

func (b *Builder) Flushdb() (c Flushdb) {
	c.cs = append(b.get(), "FLUSHDB")
	return
}

type FlushdbAsyncAsync struct {
	cs []string
}

func (c FlushdbAsyncAsync) Build() []string {
	return c.cs
}

type FlushdbAsyncSync struct {
	cs []string
}

func (c FlushdbAsyncSync) Build() []string {
	return c.cs
}

type Geoadd struct {
	cs []string
}

func (c Geoadd) Key(Key string) GeoaddKey {
	return GeoaddKey{cs: append(c.cs, Key)}
}

func (b *Builder) Geoadd() (c Geoadd) {
	c.cs = append(b.get(), "GEOADD")
	return
}

type GeoaddChangeCh struct {
	cs []string
}

func (c GeoaddChangeCh) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: append(c.cs)}
}

type GeoaddConditionNx struct {
	cs []string
}

func (c GeoaddConditionNx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cs: append(c.cs, "CH")}
}

func (c GeoaddConditionNx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: append(c.cs)}
}

type GeoaddConditionXx struct {
	cs []string
}

func (c GeoaddConditionXx) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cs: append(c.cs, "CH")}
}

func (c GeoaddConditionXx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: append(c.cs)}
}

type GeoaddKey struct {
	cs []string
}

func (c GeoaddKey) Nx() GeoaddConditionNx {
	return GeoaddConditionNx{cs: append(c.cs, "NX")}
}

func (c GeoaddKey) Xx() GeoaddConditionXx {
	return GeoaddConditionXx{cs: append(c.cs, "XX")}
}

func (c GeoaddKey) Ch() GeoaddChangeCh {
	return GeoaddChangeCh{cs: append(c.cs, "CH")}
}

func (c GeoaddKey) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: append(c.cs)}
}

type GeoaddLongitudeLatitudeMember struct {
	cs []string
}

func (c GeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(Longitude float64, Latitude float64, Member string) GeoaddLongitudeLatitudeMember {
	return GeoaddLongitudeLatitudeMember{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64), Member)}
}

func (c GeoaddLongitudeLatitudeMember) Build() []string {
	return c.cs
}

type Geodist struct {
	cs []string
}

func (c Geodist) Key(Key string) GeodistKey {
	return GeodistKey{cs: append(c.cs, Key)}
}

func (b *Builder) Geodist() (c Geodist) {
	c.cs = append(b.get(), "GEODIST")
	return
}

type GeodistKey struct {
	cs []string
}

func (c GeodistKey) Member1(Member1 string) GeodistMember1 {
	return GeodistMember1{cs: append(c.cs, Member1)}
}

type GeodistMember1 struct {
	cs []string
}

func (c GeodistMember1) Member2(Member2 string) GeodistMember2 {
	return GeodistMember2{cs: append(c.cs, Member2)}
}

type GeodistMember2 struct {
	cs []string
}

func (c GeodistMember2) M() GeodistUnitM {
	return GeodistUnitM{cs: append(c.cs, "m")}
}

func (c GeodistMember2) Km() GeodistUnitKm {
	return GeodistUnitKm{cs: append(c.cs, "km")}
}

func (c GeodistMember2) Ft() GeodistUnitFt {
	return GeodistUnitFt{cs: append(c.cs, "ft")}
}

func (c GeodistMember2) Mi() GeodistUnitMi {
	return GeodistUnitMi{cs: append(c.cs, "mi")}
}

func (c GeodistMember2) Build() []string {
	return c.cs
}

type GeodistUnitFt struct {
	cs []string
}

func (c GeodistUnitFt) Build() []string {
	return c.cs
}

type GeodistUnitKm struct {
	cs []string
}

func (c GeodistUnitKm) Build() []string {
	return c.cs
}

type GeodistUnitM struct {
	cs []string
}

func (c GeodistUnitM) Build() []string {
	return c.cs
}

type GeodistUnitMi struct {
	cs []string
}

func (c GeodistUnitMi) Build() []string {
	return c.cs
}

type Geohash struct {
	cs []string
}

func (c Geohash) Key(Key string) GeohashKey {
	return GeohashKey{cs: append(c.cs, Key)}
}

func (b *Builder) Geohash() (c Geohash) {
	c.cs = append(b.get(), "GEOHASH")
	return
}

type GeohashKey struct {
	cs []string
}

func (c GeohashKey) Member(Member ...string) GeohashMember {
	return GeohashMember{cs: append(c.cs, Member...)}
}

type GeohashMember struct {
	cs []string
}

func (c GeohashMember) Member(Member ...string) GeohashMember {
	return GeohashMember{cs: append(c.cs, Member...)}
}

func (c GeohashMember) Build() []string {
	return c.cs
}

type Geopos struct {
	cs []string
}

func (c Geopos) Key(Key string) GeoposKey {
	return GeoposKey{cs: append(c.cs, Key)}
}

func (b *Builder) Geopos() (c Geopos) {
	c.cs = append(b.get(), "GEOPOS")
	return
}

type GeoposKey struct {
	cs []string
}

func (c GeoposKey) Member(Member ...string) GeoposMember {
	return GeoposMember{cs: append(c.cs, Member...)}
}

type GeoposMember struct {
	cs []string
}

func (c GeoposMember) Member(Member ...string) GeoposMember {
	return GeoposMember{cs: append(c.cs, Member...)}
}

func (c GeoposMember) Build() []string {
	return c.cs
}

type Georadius struct {
	cs []string
}

func (c Georadius) Key(Key string) GeoradiusKey {
	return GeoradiusKey{cs: append(c.cs, Key)}
}

func (b *Builder) Georadius() (c Georadius) {
	c.cs = append(b.get(), "GEORADIUS")
	return
}

type GeoradiusCountAnyAny struct {
	cs []string
}

func (c GeoradiusCountAnyAny) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusCountAnyAny) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusCountAnyAny) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusCountAnyAny) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusCountAnyAny) Build() []string {
	return c.cs
}

type GeoradiusCountCount struct {
	cs []string
}

func (c GeoradiusCountCount) Any() GeoradiusCountAnyAny {
	return GeoradiusCountAnyAny{cs: append(c.cs, "ANY")}
}

func (c GeoradiusCountCount) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusCountCount) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusCountCount) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusCountCount) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusCountCount) Build() []string {
	return c.cs
}

type GeoradiusKey struct {
	cs []string
}

func (c GeoradiusKey) Longitude(Longitude float64) GeoradiusLongitude {
	return GeoradiusLongitude{cs: append(c.cs, strconv.FormatFloat(Longitude, 'f', -1, 64))}
}

type GeoradiusLatitude struct {
	cs []string
}

func (c GeoradiusLatitude) Radius(Radius float64) GeoradiusRadius {
	return GeoradiusRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusLongitude struct {
	cs []string
}

func (c GeoradiusLongitude) Latitude(Latitude float64) GeoradiusLatitude {
	return GeoradiusLatitude{cs: append(c.cs, strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

type GeoradiusOrderAsc struct {
	cs []string
}

func (c GeoradiusOrderAsc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusOrderAsc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusOrderAsc) Build() []string {
	return c.cs
}

type GeoradiusOrderDesc struct {
	cs []string
}

func (c GeoradiusOrderDesc) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusOrderDesc) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusOrderDesc) Build() []string {
	return c.cs
}

type GeoradiusRadius struct {
	cs []string
}

func (c GeoradiusRadius) M() GeoradiusUnitM {
	return GeoradiusUnitM{cs: append(c.cs, "m")}
}

func (c GeoradiusRadius) Km() GeoradiusUnitKm {
	return GeoradiusUnitKm{cs: append(c.cs, "km")}
}

func (c GeoradiusRadius) Ft() GeoradiusUnitFt {
	return GeoradiusUnitFt{cs: append(c.cs, "ft")}
}

func (c GeoradiusRadius) Mi() GeoradiusUnitMi {
	return GeoradiusUnitMi{cs: append(c.cs, "mi")}
}

type GeoradiusStore struct {
	cs []string
}

func (c GeoradiusStore) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusStore) Build() []string {
	return c.cs
}

type GeoradiusStoredist struct {
	cs []string
}

func (c GeoradiusStoredist) Build() []string {
	return c.cs
}

type GeoradiusUnitFt struct {
	cs []string
}

func (c GeoradiusUnitFt) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitFt) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitFt) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitFt) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitFt) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitFt) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitFt) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitFt) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusUnitKm struct {
	cs []string
}

func (c GeoradiusUnitKm) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitKm) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitKm) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitKm) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitKm) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitKm) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitKm) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitKm) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusUnitM struct {
	cs []string
}

func (c GeoradiusUnitM) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitM) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitM) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitM) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitM) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitM) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitM) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitM) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusUnitMi struct {
	cs []string
}

func (c GeoradiusUnitMi) Withcoord() GeoradiusWithcoordWithcoord {
	return GeoradiusWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusUnitMi) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusUnitMi) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusUnitMi) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusUnitMi) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusUnitMi) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusUnitMi) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusUnitMi) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusWithcoordWithcoord struct {
	cs []string
}

func (c GeoradiusWithcoordWithcoord) Withdist() GeoradiusWithdistWithdist {
	return GeoradiusWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusWithcoordWithcoord) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusWithcoordWithcoord) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusWithcoordWithcoord) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusWithcoordWithcoord) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusWithcoordWithcoord) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusWithcoordWithcoord) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusWithdistWithdist struct {
	cs []string
}

func (c GeoradiusWithdistWithdist) Withhash() GeoradiusWithhashWithhash {
	return GeoradiusWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusWithdistWithdist) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusWithdistWithdist) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusWithdistWithdist) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusWithdistWithdist) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusWithdistWithdist) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusWithhashWithhash struct {
	cs []string
}

func (c GeoradiusWithhashWithhash) Count(Count int64) GeoradiusCountCount {
	return GeoradiusCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusWithhashWithhash) Asc() GeoradiusOrderAsc {
	return GeoradiusOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusWithhashWithhash) Desc() GeoradiusOrderDesc {
	return GeoradiusOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusWithhashWithhash) Store(Key string) GeoradiusStore {
	return GeoradiusStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusWithhashWithhash) Storedist(Key string) GeoradiusStoredist {
	return GeoradiusStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type Georadiusbymember struct {
	cs []string
}

func (c Georadiusbymember) Key(Key string) GeoradiusbymemberKey {
	return GeoradiusbymemberKey{cs: append(c.cs, Key)}
}

func (b *Builder) Georadiusbymember() (c Georadiusbymember) {
	c.cs = append(b.get(), "GEORADIUSBYMEMBER")
	return
}

type GeoradiusbymemberCountAnyAny struct {
	cs []string
}

func (c GeoradiusbymemberCountAnyAny) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberCountAnyAny) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberCountAnyAny) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberCountAnyAny) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberCountAnyAny) Build() []string {
	return c.cs
}

type GeoradiusbymemberCountCount struct {
	cs []string
}

func (c GeoradiusbymemberCountCount) Any() GeoradiusbymemberCountAnyAny {
	return GeoradiusbymemberCountAnyAny{cs: append(c.cs, "ANY")}
}

func (c GeoradiusbymemberCountCount) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberCountCount) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberCountCount) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberCountCount) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberCountCount) Build() []string {
	return c.cs
}

type GeoradiusbymemberKey struct {
	cs []string
}

func (c GeoradiusbymemberKey) Member(Member string) GeoradiusbymemberMember {
	return GeoradiusbymemberMember{cs: append(c.cs, Member)}
}

type GeoradiusbymemberMember struct {
	cs []string
}

func (c GeoradiusbymemberMember) Radius(Radius float64) GeoradiusbymemberRadius {
	return GeoradiusbymemberRadius{cs: append(c.cs, strconv.FormatFloat(Radius, 'f', -1, 64))}
}

type GeoradiusbymemberOrderAsc struct {
	cs []string
}

func (c GeoradiusbymemberOrderAsc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberOrderAsc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberOrderAsc) Build() []string {
	return c.cs
}

type GeoradiusbymemberOrderDesc struct {
	cs []string
}

func (c GeoradiusbymemberOrderDesc) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberOrderDesc) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberOrderDesc) Build() []string {
	return c.cs
}

type GeoradiusbymemberRadius struct {
	cs []string
}

func (c GeoradiusbymemberRadius) M() GeoradiusbymemberUnitM {
	return GeoradiusbymemberUnitM{cs: append(c.cs, "m")}
}

func (c GeoradiusbymemberRadius) Km() GeoradiusbymemberUnitKm {
	return GeoradiusbymemberUnitKm{cs: append(c.cs, "km")}
}

func (c GeoradiusbymemberRadius) Ft() GeoradiusbymemberUnitFt {
	return GeoradiusbymemberUnitFt{cs: append(c.cs, "ft")}
}

func (c GeoradiusbymemberRadius) Mi() GeoradiusbymemberUnitMi {
	return GeoradiusbymemberUnitMi{cs: append(c.cs, "mi")}
}

type GeoradiusbymemberStore struct {
	cs []string
}

func (c GeoradiusbymemberStore) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

func (c GeoradiusbymemberStore) Build() []string {
	return c.cs
}

type GeoradiusbymemberStoredist struct {
	cs []string
}

func (c GeoradiusbymemberStoredist) Build() []string {
	return c.cs
}

type GeoradiusbymemberUnitFt struct {
	cs []string
}

func (c GeoradiusbymemberUnitFt) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitFt) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitFt) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitFt) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitFt) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitFt) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitFt) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitFt) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberUnitKm struct {
	cs []string
}

func (c GeoradiusbymemberUnitKm) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitKm) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitKm) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitKm) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitKm) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitKm) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitKm) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitKm) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberUnitM struct {
	cs []string
}

func (c GeoradiusbymemberUnitM) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitM) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitM) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitM) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitM) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitM) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitM) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitM) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberUnitMi struct {
	cs []string
}

func (c GeoradiusbymemberUnitMi) Withcoord() GeoradiusbymemberWithcoordWithcoord {
	return GeoradiusbymemberWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeoradiusbymemberUnitMi) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberUnitMi) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberUnitMi) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberUnitMi) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberUnitMi) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberUnitMi) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberUnitMi) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberWithcoordWithcoord struct {
	cs []string
}

func (c GeoradiusbymemberWithcoordWithcoord) Withdist() GeoradiusbymemberWithdistWithdist {
	return GeoradiusbymemberWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberWithcoordWithcoord) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberWithcoordWithcoord) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberWithcoordWithcoord) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberWithdistWithdist struct {
	cs []string
}

func (c GeoradiusbymemberWithdistWithdist) Withhash() GeoradiusbymemberWithhashWithhash {
	return GeoradiusbymemberWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeoradiusbymemberWithdistWithdist) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberWithdistWithdist) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberWithdistWithdist) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberWithdistWithdist) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberWithdistWithdist) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type GeoradiusbymemberWithhashWithhash struct {
	cs []string
}

func (c GeoradiusbymemberWithhashWithhash) Count(Count int64) GeoradiusbymemberCountCount {
	return GeoradiusbymemberCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeoradiusbymemberWithhashWithhash) Asc() GeoradiusbymemberOrderAsc {
	return GeoradiusbymemberOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeoradiusbymemberWithhashWithhash) Desc() GeoradiusbymemberOrderDesc {
	return GeoradiusbymemberOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeoradiusbymemberWithhashWithhash) Store(Key string) GeoradiusbymemberStore {
	return GeoradiusbymemberStore{cs: append(c.cs, "STORE", Key)}
}

func (c GeoradiusbymemberWithhashWithhash) Storedist(Key string) GeoradiusbymemberStoredist {
	return GeoradiusbymemberStoredist{cs: append(c.cs, "STOREDIST", Key)}
}

type Geosearch struct {
	cs []string
}

func (c Geosearch) Key(Key string) GeosearchKey {
	return GeosearchKey{cs: append(c.cs, Key)}
}

func (b *Builder) Geosearch() (c Geosearch) {
	c.cs = append(b.get(), "GEOSEARCH")
	return
}

type GeosearchBoxBybox struct {
	cs []string
}

func (c GeosearchBoxBybox) Height(Height float64) GeosearchBoxHeight {
	return GeosearchBoxHeight{cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type GeosearchBoxHeight struct {
	cs []string
}

func (c GeosearchBoxHeight) M() GeosearchBoxUnitM {
	return GeosearchBoxUnitM{cs: append(c.cs, "m")}
}

func (c GeosearchBoxHeight) Km() GeosearchBoxUnitKm {
	return GeosearchBoxUnitKm{cs: append(c.cs, "km")}
}

func (c GeosearchBoxHeight) Ft() GeosearchBoxUnitFt {
	return GeosearchBoxUnitFt{cs: append(c.cs, "ft")}
}

func (c GeosearchBoxHeight) Mi() GeosearchBoxUnitMi {
	return GeosearchBoxUnitMi{cs: append(c.cs, "mi")}
}

type GeosearchBoxUnitFt struct {
	cs []string
}

func (c GeosearchBoxUnitFt) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitFt) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitFt) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitFt) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitFt) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitFt) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchBoxUnitKm struct {
	cs []string
}

func (c GeosearchBoxUnitKm) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitKm) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitKm) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitKm) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitKm) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitKm) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchBoxUnitM struct {
	cs []string
}

func (c GeosearchBoxUnitM) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitM) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitM) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitM) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitM) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitM) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchBoxUnitMi struct {
	cs []string
}

func (c GeosearchBoxUnitMi) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchBoxUnitMi) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchBoxUnitMi) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchBoxUnitMi) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchBoxUnitMi) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchBoxUnitMi) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchCircleByradius struct {
	cs []string
}

func (c GeosearchCircleByradius) M() GeosearchCircleUnitM {
	return GeosearchCircleUnitM{cs: append(c.cs, "m")}
}

func (c GeosearchCircleByradius) Km() GeosearchCircleUnitKm {
	return GeosearchCircleUnitKm{cs: append(c.cs, "km")}
}

func (c GeosearchCircleByradius) Ft() GeosearchCircleUnitFt {
	return GeosearchCircleUnitFt{cs: append(c.cs, "ft")}
}

func (c GeosearchCircleByradius) Mi() GeosearchCircleUnitMi {
	return GeosearchCircleUnitMi{cs: append(c.cs, "mi")}
}

type GeosearchCircleUnitFt struct {
	cs []string
}

func (c GeosearchCircleUnitFt) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitFt) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitFt) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitFt) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitFt) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitFt) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitFt) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchCircleUnitKm struct {
	cs []string
}

func (c GeosearchCircleUnitKm) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitKm) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitKm) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitKm) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitKm) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitKm) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitKm) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchCircleUnitM struct {
	cs []string
}

func (c GeosearchCircleUnitM) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitM) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitM) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitM) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitM) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitM) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitM) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchCircleUnitMi struct {
	cs []string
}

func (c GeosearchCircleUnitMi) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchCircleUnitMi) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchCircleUnitMi) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchCircleUnitMi) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchCircleUnitMi) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCircleUnitMi) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCircleUnitMi) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchCountAnyAny struct {
	cs []string
}

func (c GeosearchCountAnyAny) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCountAnyAny) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCountAnyAny) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCountAnyAny) Build() []string {
	return c.cs
}

type GeosearchCountCount struct {
	cs []string
}

func (c GeosearchCountCount) Any() GeosearchCountAnyAny {
	return GeosearchCountAnyAny{cs: append(c.cs, "ANY")}
}

func (c GeosearchCountCount) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchCountCount) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchCountCount) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchCountCount) Build() []string {
	return c.cs
}

type GeosearchFromlonlat struct {
	cs []string
}

func (c GeosearchFromlonlat) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchFromlonlat) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchFromlonlat) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchFromlonlat) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchFromlonlat) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchFromlonlat) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchFromlonlat) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchFromlonlat) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchFrommember struct {
	cs []string
}

func (c GeosearchFrommember) Fromlonlat(Longitude float64, Latitude float64) GeosearchFromlonlat {
	return GeosearchFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchFrommember) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchFrommember) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchFrommember) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchFrommember) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchFrommember) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchFrommember) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchFrommember) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchFrommember) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchKey struct {
	cs []string
}

func (c GeosearchKey) Frommember(Member string) GeosearchFrommember {
	return GeosearchFrommember{cs: append(c.cs, "FROMMEMBER", Member)}
}

func (c GeosearchKey) Fromlonlat(Longitude float64, Latitude float64) GeosearchFromlonlat {
	return GeosearchFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchKey) Byradius(Radius float64) GeosearchCircleByradius {
	return GeosearchCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchKey) Bybox(Width float64) GeosearchBoxBybox {
	return GeosearchBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchKey) Asc() GeosearchOrderAsc {
	return GeosearchOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchKey) Desc() GeosearchOrderDesc {
	return GeosearchOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchKey) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchKey) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchKey) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchKey) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchOrderAsc struct {
	cs []string
}

func (c GeosearchOrderAsc) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchOrderAsc) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchOrderAsc) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchOrderAsc) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchOrderDesc struct {
	cs []string
}

func (c GeosearchOrderDesc) Count(Count int64) GeosearchCountCount {
	return GeosearchCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchOrderDesc) Withcoord() GeosearchWithcoordWithcoord {
	return GeosearchWithcoordWithcoord{cs: append(c.cs, "WITHCOORD")}
}

func (c GeosearchOrderDesc) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchOrderDesc) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

type GeosearchWithcoordWithcoord struct {
	cs []string
}

func (c GeosearchWithcoordWithcoord) Withdist() GeosearchWithdistWithdist {
	return GeosearchWithdistWithdist{cs: append(c.cs, "WITHDIST")}
}

func (c GeosearchWithcoordWithcoord) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchWithcoordWithcoord) Build() []string {
	return c.cs
}

type GeosearchWithdistWithdist struct {
	cs []string
}

func (c GeosearchWithdistWithdist) Withhash() GeosearchWithhashWithhash {
	return GeosearchWithhashWithhash{cs: append(c.cs, "WITHHASH")}
}

func (c GeosearchWithdistWithdist) Build() []string {
	return c.cs
}

type GeosearchWithhashWithhash struct {
	cs []string
}

func (c GeosearchWithhashWithhash) Build() []string {
	return c.cs
}

type Geosearchstore struct {
	cs []string
}

func (c Geosearchstore) Destination(Destination string) GeosearchstoreDestination {
	return GeosearchstoreDestination{cs: append(c.cs, Destination)}
}

func (b *Builder) Geosearchstore() (c Geosearchstore) {
	c.cs = append(b.get(), "GEOSEARCHSTORE")
	return
}

type GeosearchstoreBoxBybox struct {
	cs []string
}

func (c GeosearchstoreBoxBybox) Height(Height float64) GeosearchstoreBoxHeight {
	return GeosearchstoreBoxHeight{cs: append(c.cs, strconv.FormatFloat(Height, 'f', -1, 64))}
}

type GeosearchstoreBoxHeight struct {
	cs []string
}

func (c GeosearchstoreBoxHeight) M() GeosearchstoreBoxUnitM {
	return GeosearchstoreBoxUnitM{cs: append(c.cs, "m")}
}

func (c GeosearchstoreBoxHeight) Km() GeosearchstoreBoxUnitKm {
	return GeosearchstoreBoxUnitKm{cs: append(c.cs, "km")}
}

func (c GeosearchstoreBoxHeight) Ft() GeosearchstoreBoxUnitFt {
	return GeosearchstoreBoxUnitFt{cs: append(c.cs, "ft")}
}

func (c GeosearchstoreBoxHeight) Mi() GeosearchstoreBoxUnitMi {
	return GeosearchstoreBoxUnitMi{cs: append(c.cs, "mi")}
}

type GeosearchstoreBoxUnitFt struct {
	cs []string
}

func (c GeosearchstoreBoxUnitFt) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitFt) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitFt) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitFt) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreBoxUnitKm struct {
	cs []string
}

func (c GeosearchstoreBoxUnitKm) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitKm) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitKm) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitKm) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreBoxUnitM struct {
	cs []string
}

func (c GeosearchstoreBoxUnitM) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitM) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitM) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitM) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreBoxUnitMi struct {
	cs []string
}

func (c GeosearchstoreBoxUnitMi) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreBoxUnitMi) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreBoxUnitMi) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreBoxUnitMi) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleByradius struct {
	cs []string
}

func (c GeosearchstoreCircleByradius) M() GeosearchstoreCircleUnitM {
	return GeosearchstoreCircleUnitM{cs: append(c.cs, "m")}
}

func (c GeosearchstoreCircleByradius) Km() GeosearchstoreCircleUnitKm {
	return GeosearchstoreCircleUnitKm{cs: append(c.cs, "km")}
}

func (c GeosearchstoreCircleByradius) Ft() GeosearchstoreCircleUnitFt {
	return GeosearchstoreCircleUnitFt{cs: append(c.cs, "ft")}
}

func (c GeosearchstoreCircleByradius) Mi() GeosearchstoreCircleUnitMi {
	return GeosearchstoreCircleUnitMi{cs: append(c.cs, "mi")}
}

type GeosearchstoreCircleUnitFt struct {
	cs []string
}

func (c GeosearchstoreCircleUnitFt) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitFt) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitFt) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitFt) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitFt) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleUnitKm struct {
	cs []string
}

func (c GeosearchstoreCircleUnitKm) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitKm) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitKm) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitKm) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitKm) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleUnitM struct {
	cs []string
}

func (c GeosearchstoreCircleUnitM) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitM) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitM) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitM) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitM) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCircleUnitMi struct {
	cs []string
}

func (c GeosearchstoreCircleUnitMi) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreCircleUnitMi) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreCircleUnitMi) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreCircleUnitMi) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreCircleUnitMi) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreCountAnyAny struct {
	cs []string
}

func (c GeosearchstoreCountAnyAny) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreCountAnyAny) Build() []string {
	return c.cs
}

type GeosearchstoreCountCount struct {
	cs []string
}

func (c GeosearchstoreCountCount) Any() GeosearchstoreCountAnyAny {
	return GeosearchstoreCountAnyAny{cs: append(c.cs, "ANY")}
}

func (c GeosearchstoreCountCount) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

func (c GeosearchstoreCountCount) Build() []string {
	return c.cs
}

type GeosearchstoreDestination struct {
	cs []string
}

func (c GeosearchstoreDestination) Source(Source string) GeosearchstoreSource {
	return GeosearchstoreSource{cs: append(c.cs, Source)}
}

type GeosearchstoreFromlonlat struct {
	cs []string
}

func (c GeosearchstoreFromlonlat) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchstoreFromlonlat) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreFromlonlat) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreFromlonlat) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreFromlonlat) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreFromlonlat) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreFrommember struct {
	cs []string
}

func (c GeosearchstoreFrommember) Fromlonlat(Longitude float64, Latitude float64) GeosearchstoreFromlonlat {
	return GeosearchstoreFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchstoreFrommember) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchstoreFrommember) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreFrommember) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreFrommember) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreFrommember) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreFrommember) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreOrderAsc struct {
	cs []string
}

func (c GeosearchstoreOrderAsc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderAsc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreOrderDesc struct {
	cs []string
}

func (c GeosearchstoreOrderDesc) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreOrderDesc) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreSource struct {
	cs []string
}

func (c GeosearchstoreSource) Frommember(Member string) GeosearchstoreFrommember {
	return GeosearchstoreFrommember{cs: append(c.cs, "FROMMEMBER", Member)}
}

func (c GeosearchstoreSource) Fromlonlat(Longitude float64, Latitude float64) GeosearchstoreFromlonlat {
	return GeosearchstoreFromlonlat{cs: append(c.cs, "FROMLONLAT", strconv.FormatFloat(Longitude, 'f', -1, 64), strconv.FormatFloat(Latitude, 'f', -1, 64))}
}

func (c GeosearchstoreSource) Byradius(Radius float64) GeosearchstoreCircleByradius {
	return GeosearchstoreCircleByradius{cs: append(c.cs, "BYRADIUS", strconv.FormatFloat(Radius, 'f', -1, 64))}
}

func (c GeosearchstoreSource) Bybox(Width float64) GeosearchstoreBoxBybox {
	return GeosearchstoreBoxBybox{cs: append(c.cs, "BYBOX", strconv.FormatFloat(Width, 'f', -1, 64))}
}

func (c GeosearchstoreSource) Asc() GeosearchstoreOrderAsc {
	return GeosearchstoreOrderAsc{cs: append(c.cs, "ASC")}
}

func (c GeosearchstoreSource) Desc() GeosearchstoreOrderDesc {
	return GeosearchstoreOrderDesc{cs: append(c.cs, "DESC")}
}

func (c GeosearchstoreSource) Count(Count int64) GeosearchstoreCountCount {
	return GeosearchstoreCountCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c GeosearchstoreSource) Storedist() GeosearchstoreStoredistStoredist {
	return GeosearchstoreStoredistStoredist{cs: append(c.cs, "STOREDIST")}
}

type GeosearchstoreStoredistStoredist struct {
	cs []string
}

func (c GeosearchstoreStoredistStoredist) Build() []string {
	return c.cs
}

type Get struct {
	cs []string
}

func (c Get) Key(Key string) GetKey {
	return GetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Get() (c Get) {
	c.cs = append(b.get(), "GET")
	return
}

type GetKey struct {
	cs []string
}

func (c GetKey) Build() []string {
	return c.cs
}

type Getbit struct {
	cs []string
}

func (c Getbit) Key(Key string) GetbitKey {
	return GetbitKey{cs: append(c.cs, Key)}
}

func (b *Builder) Getbit() (c Getbit) {
	c.cs = append(b.get(), "GETBIT")
	return
}

type GetbitKey struct {
	cs []string
}

func (c GetbitKey) Offset(Offset int64) GetbitOffset {
	return GetbitOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type GetbitOffset struct {
	cs []string
}

func (c GetbitOffset) Build() []string {
	return c.cs
}

type Getdel struct {
	cs []string
}

func (c Getdel) Key(Key string) GetdelKey {
	return GetdelKey{cs: append(c.cs, Key)}
}

func (b *Builder) Getdel() (c Getdel) {
	c.cs = append(b.get(), "GETDEL")
	return
}

type GetdelKey struct {
	cs []string
}

func (c GetdelKey) Build() []string {
	return c.cs
}

type Getex struct {
	cs []string
}

func (c Getex) Key(Key string) GetexKey {
	return GetexKey{cs: append(c.cs, Key)}
}

func (b *Builder) Getex() (c Getex) {
	c.cs = append(b.get(), "GETEX")
	return
}

type GetexExpirationEx struct {
	cs []string
}

func (c GetexExpirationEx) Build() []string {
	return c.cs
}

type GetexExpirationExat struct {
	cs []string
}

func (c GetexExpirationExat) Build() []string {
	return c.cs
}

type GetexExpirationPersist struct {
	cs []string
}

func (c GetexExpirationPersist) Build() []string {
	return c.cs
}

type GetexExpirationPx struct {
	cs []string
}

func (c GetexExpirationPx) Build() []string {
	return c.cs
}

type GetexExpirationPxat struct {
	cs []string
}

func (c GetexExpirationPxat) Build() []string {
	return c.cs
}

type GetexKey struct {
	cs []string
}

func (c GetexKey) Ex(Seconds int64) GetexExpirationEx {
	return GetexExpirationEx{cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10))}
}

func (c GetexKey) Px(Milliseconds int64) GetexExpirationPx {
	return GetexExpirationPx{cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10))}
}

func (c GetexKey) Exat(Timestamp int64) GetexExpirationExat {
	return GetexExpirationExat{cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10))}
}

func (c GetexKey) Pxat(Millisecondstimestamp int64) GetexExpirationPxat {
	return GetexExpirationPxat{cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10))}
}

func (c GetexKey) Persist() GetexExpirationPersist {
	return GetexExpirationPersist{cs: append(c.cs, "PERSIST")}
}

func (c GetexKey) Build() []string {
	return c.cs
}

type Getrange struct {
	cs []string
}

func (c Getrange) Key(Key string) GetrangeKey {
	return GetrangeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Getrange() (c Getrange) {
	c.cs = append(b.get(), "GETRANGE")
	return
}

type GetrangeEnd struct {
	cs []string
}

func (c GetrangeEnd) Build() []string {
	return c.cs
}

type GetrangeKey struct {
	cs []string
}

func (c GetrangeKey) Start(Start int64) GetrangeStart {
	return GetrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type GetrangeStart struct {
	cs []string
}

func (c GetrangeStart) End(End int64) GetrangeEnd {
	return GetrangeEnd{cs: append(c.cs, strconv.FormatInt(End, 10))}
}

type Getset struct {
	cs []string
}

func (c Getset) Key(Key string) GetsetKey {
	return GetsetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Getset() (c Getset) {
	c.cs = append(b.get(), "GETSET")
	return
}

type GetsetKey struct {
	cs []string
}

func (c GetsetKey) Value(Value string) GetsetValue {
	return GetsetValue{cs: append(c.cs, Value)}
}

type GetsetValue struct {
	cs []string
}

func (c GetsetValue) Build() []string {
	return c.cs
}

type Hdel struct {
	cs []string
}

func (c Hdel) Key(Key string) HdelKey {
	return HdelKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hdel() (c Hdel) {
	c.cs = append(b.get(), "HDEL")
	return
}

type HdelField struct {
	cs []string
}

func (c HdelField) Field(Field ...string) HdelField {
	return HdelField{cs: append(c.cs, Field...)}
}

func (c HdelField) Build() []string {
	return c.cs
}

type HdelKey struct {
	cs []string
}

func (c HdelKey) Field(Field ...string) HdelField {
	return HdelField{cs: append(c.cs, Field...)}
}

type Hello struct {
	cs []string
}

func (c Hello) Protover(Protover int64) HelloArgumentsProtover {
	return HelloArgumentsProtover{cs: append(c.cs, strconv.FormatInt(Protover, 10))}
}

func (b *Builder) Hello() (c Hello) {
	c.cs = append(b.get(), "HELLO")
	return
}

type HelloArgumentsAuth struct {
	cs []string
}

func (c HelloArgumentsAuth) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cs: append(c.cs, "SETNAME", Clientname)}
}

func (c HelloArgumentsAuth) Build() []string {
	return c.cs
}

type HelloArgumentsProtover struct {
	cs []string
}

func (c HelloArgumentsProtover) Auth(Username string, Password string) HelloArgumentsAuth {
	return HelloArgumentsAuth{cs: append(c.cs, "AUTH", Username, Password)}
}

func (c HelloArgumentsProtover) Setname(Clientname string) HelloArgumentsSetname {
	return HelloArgumentsSetname{cs: append(c.cs, "SETNAME", Clientname)}
}

func (c HelloArgumentsProtover) Build() []string {
	return c.cs
}

type HelloArgumentsSetname struct {
	cs []string
}

func (c HelloArgumentsSetname) Build() []string {
	return c.cs
}

type Hexists struct {
	cs []string
}

func (c Hexists) Key(Key string) HexistsKey {
	return HexistsKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hexists() (c Hexists) {
	c.cs = append(b.get(), "HEXISTS")
	return
}

type HexistsField struct {
	cs []string
}

func (c HexistsField) Build() []string {
	return c.cs
}

type HexistsKey struct {
	cs []string
}

func (c HexistsKey) Field(Field string) HexistsField {
	return HexistsField{cs: append(c.cs, Field)}
}

type Hget struct {
	cs []string
}

func (c Hget) Key(Key string) HgetKey {
	return HgetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hget() (c Hget) {
	c.cs = append(b.get(), "HGET")
	return
}

type HgetField struct {
	cs []string
}

func (c HgetField) Build() []string {
	return c.cs
}

type HgetKey struct {
	cs []string
}

func (c HgetKey) Field(Field string) HgetField {
	return HgetField{cs: append(c.cs, Field)}
}

type Hgetall struct {
	cs []string
}

func (c Hgetall) Key(Key string) HgetallKey {
	return HgetallKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hgetall() (c Hgetall) {
	c.cs = append(b.get(), "HGETALL")
	return
}

type HgetallKey struct {
	cs []string
}

func (c HgetallKey) Build() []string {
	return c.cs
}

type Hincrby struct {
	cs []string
}

func (c Hincrby) Key(Key string) HincrbyKey {
	return HincrbyKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hincrby() (c Hincrby) {
	c.cs = append(b.get(), "HINCRBY")
	return
}

type HincrbyField struct {
	cs []string
}

func (c HincrbyField) Increment(Increment int64) HincrbyIncrement {
	return HincrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type HincrbyIncrement struct {
	cs []string
}

func (c HincrbyIncrement) Build() []string {
	return c.cs
}

type HincrbyKey struct {
	cs []string
}

func (c HincrbyKey) Field(Field string) HincrbyField {
	return HincrbyField{cs: append(c.cs, Field)}
}

type Hincrbyfloat struct {
	cs []string
}

func (c Hincrbyfloat) Key(Key string) HincrbyfloatKey {
	return HincrbyfloatKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hincrbyfloat() (c Hincrbyfloat) {
	c.cs = append(b.get(), "HINCRBYFLOAT")
	return
}

type HincrbyfloatField struct {
	cs []string
}

func (c HincrbyfloatField) Increment(Increment float64) HincrbyfloatIncrement {
	return HincrbyfloatIncrement{cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type HincrbyfloatIncrement struct {
	cs []string
}

func (c HincrbyfloatIncrement) Build() []string {
	return c.cs
}

type HincrbyfloatKey struct {
	cs []string
}

func (c HincrbyfloatKey) Field(Field string) HincrbyfloatField {
	return HincrbyfloatField{cs: append(c.cs, Field)}
}

type Hkeys struct {
	cs []string
}

func (c Hkeys) Key(Key string) HkeysKey {
	return HkeysKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hkeys() (c Hkeys) {
	c.cs = append(b.get(), "HKEYS")
	return
}

type HkeysKey struct {
	cs []string
}

func (c HkeysKey) Build() []string {
	return c.cs
}

type Hlen struct {
	cs []string
}

func (c Hlen) Key(Key string) HlenKey {
	return HlenKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hlen() (c Hlen) {
	c.cs = append(b.get(), "HLEN")
	return
}

type HlenKey struct {
	cs []string
}

func (c HlenKey) Build() []string {
	return c.cs
}

type Hmget struct {
	cs []string
}

func (c Hmget) Key(Key string) HmgetKey {
	return HmgetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hmget() (c Hmget) {
	c.cs = append(b.get(), "HMGET")
	return
}

type HmgetField struct {
	cs []string
}

func (c HmgetField) Field(Field ...string) HmgetField {
	return HmgetField{cs: append(c.cs, Field...)}
}

func (c HmgetField) Build() []string {
	return c.cs
}

type HmgetKey struct {
	cs []string
}

func (c HmgetKey) Field(Field ...string) HmgetField {
	return HmgetField{cs: append(c.cs, Field...)}
}

type Hmset struct {
	cs []string
}

func (c Hmset) Key(Key string) HmsetKey {
	return HmsetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hmset() (c Hmset) {
	c.cs = append(b.get(), "HMSET")
	return
}

type HmsetFieldValue struct {
	cs []string
}

func (c HmsetFieldValue) FieldValue(Field string, Value string) HmsetFieldValue {
	return HmsetFieldValue{cs: append(c.cs, Field, Value)}
}

func (c HmsetFieldValue) Build() []string {
	return c.cs
}

type HmsetKey struct {
	cs []string
}

func (c HmsetKey) FieldValue() HmsetFieldValue {
	return HmsetFieldValue{cs: append(c.cs)}
}

type Hrandfield struct {
	cs []string
}

func (c Hrandfield) Key(Key string) HrandfieldKey {
	return HrandfieldKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hrandfield() (c Hrandfield) {
	c.cs = append(b.get(), "HRANDFIELD")
	return
}

type HrandfieldKey struct {
	cs []string
}

func (c HrandfieldKey) Count(Count int64) HrandfieldOptionsCount {
	return HrandfieldOptionsCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type HrandfieldOptionsCount struct {
	cs []string
}

func (c HrandfieldOptionsCount) Withvalues() HrandfieldOptionsWithvaluesWithvalues {
	return HrandfieldOptionsWithvaluesWithvalues{cs: append(c.cs, "WITHVALUES")}
}

func (c HrandfieldOptionsCount) Build() []string {
	return c.cs
}

type HrandfieldOptionsWithvaluesWithvalues struct {
	cs []string
}

func (c HrandfieldOptionsWithvaluesWithvalues) Build() []string {
	return c.cs
}

type Hscan struct {
	cs []string
}

func (c Hscan) Key(Key string) HscanKey {
	return HscanKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hscan() (c Hscan) {
	c.cs = append(b.get(), "HSCAN")
	return
}

type HscanCount struct {
	cs []string
}

func (c HscanCount) Build() []string {
	return c.cs
}

type HscanCursor struct {
	cs []string
}

func (c HscanCursor) Match(Pattern string) HscanMatch {
	return HscanMatch{cs: append(c.cs, "MATCH", Pattern)}
}

func (c HscanCursor) Count(Count int64) HscanCount {
	return HscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c HscanCursor) Build() []string {
	return c.cs
}

type HscanKey struct {
	cs []string
}

func (c HscanKey) Cursor(Cursor int64) HscanCursor {
	return HscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type HscanMatch struct {
	cs []string
}

func (c HscanMatch) Count(Count int64) HscanCount {
	return HscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c HscanMatch) Build() []string {
	return c.cs
}

type Hset struct {
	cs []string
}

func (c Hset) Key(Key string) HsetKey {
	return HsetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hset() (c Hset) {
	c.cs = append(b.get(), "HSET")
	return
}

type HsetFieldValue struct {
	cs []string
}

func (c HsetFieldValue) FieldValue(Field string, Value string) HsetFieldValue {
	return HsetFieldValue{cs: append(c.cs, Field, Value)}
}

func (c HsetFieldValue) Build() []string {
	return c.cs
}

type HsetKey struct {
	cs []string
}

func (c HsetKey) FieldValue() HsetFieldValue {
	return HsetFieldValue{cs: append(c.cs)}
}

type Hsetnx struct {
	cs []string
}

func (c Hsetnx) Key(Key string) HsetnxKey {
	return HsetnxKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hsetnx() (c Hsetnx) {
	c.cs = append(b.get(), "HSETNX")
	return
}

type HsetnxField struct {
	cs []string
}

func (c HsetnxField) Value(Value string) HsetnxValue {
	return HsetnxValue{cs: append(c.cs, Value)}
}

type HsetnxKey struct {
	cs []string
}

func (c HsetnxKey) Field(Field string) HsetnxField {
	return HsetnxField{cs: append(c.cs, Field)}
}

type HsetnxValue struct {
	cs []string
}

func (c HsetnxValue) Build() []string {
	return c.cs
}

type Hstrlen struct {
	cs []string
}

func (c Hstrlen) Key(Key string) HstrlenKey {
	return HstrlenKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hstrlen() (c Hstrlen) {
	c.cs = append(b.get(), "HSTRLEN")
	return
}

type HstrlenField struct {
	cs []string
}

func (c HstrlenField) Build() []string {
	return c.cs
}

type HstrlenKey struct {
	cs []string
}

func (c HstrlenKey) Field(Field string) HstrlenField {
	return HstrlenField{cs: append(c.cs, Field)}
}

type Hvals struct {
	cs []string
}

func (c Hvals) Key(Key string) HvalsKey {
	return HvalsKey{cs: append(c.cs, Key)}
}

func (b *Builder) Hvals() (c Hvals) {
	c.cs = append(b.get(), "HVALS")
	return
}

type HvalsKey struct {
	cs []string
}

func (c HvalsKey) Build() []string {
	return c.cs
}

type Incr struct {
	cs []string
}

func (c Incr) Key(Key string) IncrKey {
	return IncrKey{cs: append(c.cs, Key)}
}

func (b *Builder) Incr() (c Incr) {
	c.cs = append(b.get(), "INCR")
	return
}

type IncrKey struct {
	cs []string
}

func (c IncrKey) Build() []string {
	return c.cs
}

type Incrby struct {
	cs []string
}

func (c Incrby) Key(Key string) IncrbyKey {
	return IncrbyKey{cs: append(c.cs, Key)}
}

func (b *Builder) Incrby() (c Incrby) {
	c.cs = append(b.get(), "INCRBY")
	return
}

type IncrbyIncrement struct {
	cs []string
}

func (c IncrbyIncrement) Build() []string {
	return c.cs
}

type IncrbyKey struct {
	cs []string
}

func (c IncrbyKey) Increment(Increment int64) IncrbyIncrement {
	return IncrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type Incrbyfloat struct {
	cs []string
}

func (c Incrbyfloat) Key(Key string) IncrbyfloatKey {
	return IncrbyfloatKey{cs: append(c.cs, Key)}
}

func (b *Builder) Incrbyfloat() (c Incrbyfloat) {
	c.cs = append(b.get(), "INCRBYFLOAT")
	return
}

type IncrbyfloatIncrement struct {
	cs []string
}

func (c IncrbyfloatIncrement) Build() []string {
	return c.cs
}

type IncrbyfloatKey struct {
	cs []string
}

func (c IncrbyfloatKey) Increment(Increment float64) IncrbyfloatIncrement {
	return IncrbyfloatIncrement{cs: append(c.cs, strconv.FormatFloat(Increment, 'f', -1, 64))}
}

type Info struct {
	cs []string
}

func (c Info) Section(Section string) InfoSection {
	return InfoSection{cs: append(c.cs, Section)}
}

func (c Info) Build() []string {
	return c.cs
}

func (b *Builder) Info() (c Info) {
	c.cs = append(b.get(), "INFO")
	return
}

type InfoSection struct {
	cs []string
}

func (c InfoSection) Build() []string {
	return c.cs
}

type Keys struct {
	cs []string
}

func (c Keys) Pattern(Pattern string) KeysPattern {
	return KeysPattern{cs: append(c.cs, Pattern)}
}

func (b *Builder) Keys() (c Keys) {
	c.cs = append(b.get(), "KEYS")
	return
}

type KeysPattern struct {
	cs []string
}

func (c KeysPattern) Build() []string {
	return c.cs
}

type Lastsave struct {
	cs []string
}

func (c Lastsave) Build() []string {
	return c.cs
}

func (b *Builder) Lastsave() (c Lastsave) {
	c.cs = append(b.get(), "LASTSAVE")
	return
}

type LatencyDoctor struct {
	cs []string
}

func (c LatencyDoctor) Build() []string {
	return c.cs
}

func (b *Builder) LatencyDoctor() (c LatencyDoctor) {
	c.cs = append(b.get(), "LATENCY", "DOCTOR")
	return
}

type LatencyGraph struct {
	cs []string
}

func (c LatencyGraph) Event(Event string) LatencyGraphEvent {
	return LatencyGraphEvent{cs: append(c.cs, Event)}
}

func (b *Builder) LatencyGraph() (c LatencyGraph) {
	c.cs = append(b.get(), "LATENCY", "GRAPH")
	return
}

type LatencyGraphEvent struct {
	cs []string
}

func (c LatencyGraphEvent) Build() []string {
	return c.cs
}

type LatencyHelp struct {
	cs []string
}

func (c LatencyHelp) Build() []string {
	return c.cs
}

func (b *Builder) LatencyHelp() (c LatencyHelp) {
	c.cs = append(b.get(), "LATENCY", "HELP")
	return
}

type LatencyHistory struct {
	cs []string
}

func (c LatencyHistory) Event(Event string) LatencyHistoryEvent {
	return LatencyHistoryEvent{cs: append(c.cs, Event)}
}

func (b *Builder) LatencyHistory() (c LatencyHistory) {
	c.cs = append(b.get(), "LATENCY", "HISTORY")
	return
}

type LatencyHistoryEvent struct {
	cs []string
}

func (c LatencyHistoryEvent) Build() []string {
	return c.cs
}

type LatencyLatest struct {
	cs []string
}

func (c LatencyLatest) Build() []string {
	return c.cs
}

func (b *Builder) LatencyLatest() (c LatencyLatest) {
	c.cs = append(b.get(), "LATENCY", "LATEST")
	return
}

type LatencyReset struct {
	cs []string
}

func (c LatencyReset) Event(Event ...string) LatencyResetEvent {
	return LatencyResetEvent{cs: append(c.cs, Event...)}
}

func (c LatencyReset) Build() []string {
	return c.cs
}

func (b *Builder) LatencyReset() (c LatencyReset) {
	c.cs = append(b.get(), "LATENCY", "RESET")
	return
}

type LatencyResetEvent struct {
	cs []string
}

func (c LatencyResetEvent) Event(Event ...string) LatencyResetEvent {
	return LatencyResetEvent{cs: append(c.cs, Event...)}
}

func (c LatencyResetEvent) Build() []string {
	return c.cs
}

type Lindex struct {
	cs []string
}

func (c Lindex) Key(Key string) LindexKey {
	return LindexKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lindex() (c Lindex) {
	c.cs = append(b.get(), "LINDEX")
	return
}

type LindexIndex struct {
	cs []string
}

func (c LindexIndex) Build() []string {
	return c.cs
}

type LindexKey struct {
	cs []string
}

func (c LindexKey) Index(Index int64) LindexIndex {
	return LindexIndex{cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type Linsert struct {
	cs []string
}

func (c Linsert) Key(Key string) LinsertKey {
	return LinsertKey{cs: append(c.cs, Key)}
}

func (b *Builder) Linsert() (c Linsert) {
	c.cs = append(b.get(), "LINSERT")
	return
}

type LinsertElement struct {
	cs []string
}

func (c LinsertElement) Build() []string {
	return c.cs
}

type LinsertKey struct {
	cs []string
}

func (c LinsertKey) Before() LinsertWhereBefore {
	return LinsertWhereBefore{cs: append(c.cs, "BEFORE")}
}

func (c LinsertKey) After() LinsertWhereAfter {
	return LinsertWhereAfter{cs: append(c.cs, "AFTER")}
}

type LinsertPivot struct {
	cs []string
}

func (c LinsertPivot) Element(Element string) LinsertElement {
	return LinsertElement{cs: append(c.cs, Element)}
}

type LinsertWhereAfter struct {
	cs []string
}

func (c LinsertWhereAfter) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cs: append(c.cs, Pivot)}
}

type LinsertWhereBefore struct {
	cs []string
}

func (c LinsertWhereBefore) Pivot(Pivot string) LinsertPivot {
	return LinsertPivot{cs: append(c.cs, Pivot)}
}

type Llen struct {
	cs []string
}

func (c Llen) Key(Key string) LlenKey {
	return LlenKey{cs: append(c.cs, Key)}
}

func (b *Builder) Llen() (c Llen) {
	c.cs = append(b.get(), "LLEN")
	return
}

type LlenKey struct {
	cs []string
}

func (c LlenKey) Build() []string {
	return c.cs
}

type Lmove struct {
	cs []string
}

func (c Lmove) Source(Source string) LmoveSource {
	return LmoveSource{cs: append(c.cs, Source)}
}

func (b *Builder) Lmove() (c Lmove) {
	c.cs = append(b.get(), "LMOVE")
	return
}

type LmoveDestination struct {
	cs []string
}

func (c LmoveDestination) Left() LmoveWherefromLeft {
	return LmoveWherefromLeft{cs: append(c.cs, "LEFT")}
}

func (c LmoveDestination) Right() LmoveWherefromRight {
	return LmoveWherefromRight{cs: append(c.cs, "RIGHT")}
}

type LmoveSource struct {
	cs []string
}

func (c LmoveSource) Destination(Destination string) LmoveDestination {
	return LmoveDestination{cs: append(c.cs, Destination)}
}

type LmoveWherefromLeft struct {
	cs []string
}

func (c LmoveWherefromLeft) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromLeft) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cs: append(c.cs, "RIGHT")}
}

type LmoveWherefromRight struct {
	cs []string
}

func (c LmoveWherefromRight) Left() LmoveWheretoLeft {
	return LmoveWheretoLeft{cs: append(c.cs, "LEFT")}
}

func (c LmoveWherefromRight) Right() LmoveWheretoRight {
	return LmoveWheretoRight{cs: append(c.cs, "RIGHT")}
}

type LmoveWheretoLeft struct {
	cs []string
}

func (c LmoveWheretoLeft) Build() []string {
	return c.cs
}

type LmoveWheretoRight struct {
	cs []string
}

func (c LmoveWheretoRight) Build() []string {
	return c.cs
}

type Lmpop struct {
	cs []string
}

func (c Lmpop) Numkeys(Numkeys int64) LmpopNumkeys {
	return LmpopNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Lmpop() (c Lmpop) {
	c.cs = append(b.get(), "LMPOP")
	return
}

type LmpopCount struct {
	cs []string
}

func (c LmpopCount) Build() []string {
	return c.cs
}

type LmpopKey struct {
	cs []string
}

func (c LmpopKey) Left() LmpopWhereLeft {
	return LmpopWhereLeft{cs: append(c.cs, "LEFT")}
}

func (c LmpopKey) Right() LmpopWhereRight {
	return LmpopWhereRight{cs: append(c.cs, "RIGHT")}
}

func (c LmpopKey) Key(Key ...string) LmpopKey {
	return LmpopKey{cs: append(c.cs, Key...)}
}

type LmpopNumkeys struct {
	cs []string
}

func (c LmpopNumkeys) Key(Key ...string) LmpopKey {
	return LmpopKey{cs: append(c.cs, Key...)}
}

func (c LmpopNumkeys) Left() LmpopWhereLeft {
	return LmpopWhereLeft{cs: append(c.cs, "LEFT")}
}

func (c LmpopNumkeys) Right() LmpopWhereRight {
	return LmpopWhereRight{cs: append(c.cs, "RIGHT")}
}

type LmpopWhereLeft struct {
	cs []string
}

func (c LmpopWhereLeft) Count(Count int64) LmpopCount {
	return LmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereLeft) Build() []string {
	return c.cs
}

type LmpopWhereRight struct {
	cs []string
}

func (c LmpopWhereRight) Count(Count int64) LmpopCount {
	return LmpopCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c LmpopWhereRight) Build() []string {
	return c.cs
}

type Lolwut struct {
	cs []string
}

func (c Lolwut) Version(Version int64) LolwutVersion {
	return LolwutVersion{cs: append(c.cs, "VERSION", strconv.FormatInt(Version, 10))}
}

func (c Lolwut) Build() []string {
	return c.cs
}

func (b *Builder) Lolwut() (c Lolwut) {
	c.cs = append(b.get(), "LOLWUT")
	return
}

type LolwutVersion struct {
	cs []string
}

func (c LolwutVersion) Build() []string {
	return c.cs
}

type Lpop struct {
	cs []string
}

func (c Lpop) Key(Key string) LpopKey {
	return LpopKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lpop() (c Lpop) {
	c.cs = append(b.get(), "LPOP")
	return
}

type LpopCount struct {
	cs []string
}

func (c LpopCount) Build() []string {
	return c.cs
}

type LpopKey struct {
	cs []string
}

func (c LpopKey) Count(Count int64) LpopCount {
	return LpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c LpopKey) Build() []string {
	return c.cs
}

type Lpos struct {
	cs []string
}

func (c Lpos) Key(Key string) LposKey {
	return LposKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lpos() (c Lpos) {
	c.cs = append(b.get(), "LPOS")
	return
}

type LposCount struct {
	cs []string
}

func (c LposCount) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c LposCount) Build() []string {
	return c.cs
}

type LposElement struct {
	cs []string
}

func (c LposElement) Rank(Rank int64) LposRank {
	return LposRank{cs: append(c.cs, "RANK", strconv.FormatInt(Rank, 10))}
}

func (c LposElement) Count(NumMatches int64) LposCount {
	return LposCount{cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10))}
}

func (c LposElement) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c LposElement) Build() []string {
	return c.cs
}

type LposKey struct {
	cs []string
}

func (c LposKey) Element(Element string) LposElement {
	return LposElement{cs: append(c.cs, Element)}
}

type LposMaxlen struct {
	cs []string
}

func (c LposMaxlen) Build() []string {
	return c.cs
}

type LposRank struct {
	cs []string
}

func (c LposRank) Count(NumMatches int64) LposCount {
	return LposCount{cs: append(c.cs, "COUNT", strconv.FormatInt(NumMatches, 10))}
}

func (c LposRank) Maxlen(Len int64) LposMaxlen {
	return LposMaxlen{cs: append(c.cs, "MAXLEN", strconv.FormatInt(Len, 10))}
}

func (c LposRank) Build() []string {
	return c.cs
}

type Lpush struct {
	cs []string
}

func (c Lpush) Key(Key string) LpushKey {
	return LpushKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lpush() (c Lpush) {
	c.cs = append(b.get(), "LPUSH")
	return
}

type LpushElement struct {
	cs []string
}

func (c LpushElement) Element(Element ...string) LpushElement {
	return LpushElement{cs: append(c.cs, Element...)}
}

func (c LpushElement) Build() []string {
	return c.cs
}

type LpushKey struct {
	cs []string
}

func (c LpushKey) Element(Element ...string) LpushElement {
	return LpushElement{cs: append(c.cs, Element...)}
}

type Lpushx struct {
	cs []string
}

func (c Lpushx) Key(Key string) LpushxKey {
	return LpushxKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lpushx() (c Lpushx) {
	c.cs = append(b.get(), "LPUSHX")
	return
}

type LpushxElement struct {
	cs []string
}

func (c LpushxElement) Element(Element ...string) LpushxElement {
	return LpushxElement{cs: append(c.cs, Element...)}
}

func (c LpushxElement) Build() []string {
	return c.cs
}

type LpushxKey struct {
	cs []string
}

func (c LpushxKey) Element(Element ...string) LpushxElement {
	return LpushxElement{cs: append(c.cs, Element...)}
}

type Lrange struct {
	cs []string
}

func (c Lrange) Key(Key string) LrangeKey {
	return LrangeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lrange() (c Lrange) {
	c.cs = append(b.get(), "LRANGE")
	return
}

type LrangeKey struct {
	cs []string
}

func (c LrangeKey) Start(Start int64) LrangeStart {
	return LrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type LrangeStart struct {
	cs []string
}

func (c LrangeStart) Stop(Stop int64) LrangeStop {
	return LrangeStop{cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type LrangeStop struct {
	cs []string
}

func (c LrangeStop) Build() []string {
	return c.cs
}

type Lrem struct {
	cs []string
}

func (c Lrem) Key(Key string) LremKey {
	return LremKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lrem() (c Lrem) {
	c.cs = append(b.get(), "LREM")
	return
}

type LremCount struct {
	cs []string
}

func (c LremCount) Element(Element string) LremElement {
	return LremElement{cs: append(c.cs, Element)}
}

type LremElement struct {
	cs []string
}

func (c LremElement) Build() []string {
	return c.cs
}

type LremKey struct {
	cs []string
}

func (c LremKey) Count(Count int64) LremCount {
	return LremCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type Lset struct {
	cs []string
}

func (c Lset) Key(Key string) LsetKey {
	return LsetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Lset() (c Lset) {
	c.cs = append(b.get(), "LSET")
	return
}

type LsetElement struct {
	cs []string
}

func (c LsetElement) Build() []string {
	return c.cs
}

type LsetIndex struct {
	cs []string
}

func (c LsetIndex) Element(Element string) LsetElement {
	return LsetElement{cs: append(c.cs, Element)}
}

type LsetKey struct {
	cs []string
}

func (c LsetKey) Index(Index int64) LsetIndex {
	return LsetIndex{cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

type Ltrim struct {
	cs []string
}

func (c Ltrim) Key(Key string) LtrimKey {
	return LtrimKey{cs: append(c.cs, Key)}
}

func (b *Builder) Ltrim() (c Ltrim) {
	c.cs = append(b.get(), "LTRIM")
	return
}

type LtrimKey struct {
	cs []string
}

func (c LtrimKey) Start(Start int64) LtrimStart {
	return LtrimStart{cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type LtrimStart struct {
	cs []string
}

func (c LtrimStart) Stop(Stop int64) LtrimStop {
	return LtrimStop{cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type LtrimStop struct {
	cs []string
}

func (c LtrimStop) Build() []string {
	return c.cs
}

type MemoryDoctor struct {
	cs []string
}

func (c MemoryDoctor) Build() []string {
	return c.cs
}

func (b *Builder) MemoryDoctor() (c MemoryDoctor) {
	c.cs = append(b.get(), "MEMORY", "DOCTOR")
	return
}

type MemoryHelp struct {
	cs []string
}

func (c MemoryHelp) Build() []string {
	return c.cs
}

func (b *Builder) MemoryHelp() (c MemoryHelp) {
	c.cs = append(b.get(), "MEMORY", "HELP")
	return
}

type MemoryMallocStats struct {
	cs []string
}

func (c MemoryMallocStats) Build() []string {
	return c.cs
}

func (b *Builder) MemoryMallocStats() (c MemoryMallocStats) {
	c.cs = append(b.get(), "MEMORY", "MALLOC-STATS")
	return
}

type MemoryPurge struct {
	cs []string
}

func (c MemoryPurge) Build() []string {
	return c.cs
}

func (b *Builder) MemoryPurge() (c MemoryPurge) {
	c.cs = append(b.get(), "MEMORY", "PURGE")
	return
}

type MemoryStats struct {
	cs []string
}

func (c MemoryStats) Build() []string {
	return c.cs
}

func (b *Builder) MemoryStats() (c MemoryStats) {
	c.cs = append(b.get(), "MEMORY", "STATS")
	return
}

type MemoryUsage struct {
	cs []string
}

func (c MemoryUsage) Key(Key string) MemoryUsageKey {
	return MemoryUsageKey{cs: append(c.cs, Key)}
}

func (b *Builder) MemoryUsage() (c MemoryUsage) {
	c.cs = append(b.get(), "MEMORY", "USAGE")
	return
}

type MemoryUsageKey struct {
	cs []string
}

func (c MemoryUsageKey) Samples(Count int64) MemoryUsageSamples {
	return MemoryUsageSamples{cs: append(c.cs, "SAMPLES", strconv.FormatInt(Count, 10))}
}

func (c MemoryUsageKey) Build() []string {
	return c.cs
}

type MemoryUsageSamples struct {
	cs []string
}

func (c MemoryUsageSamples) Build() []string {
	return c.cs
}

type Mget struct {
	cs []string
}

func (c Mget) Key(Key ...string) MgetKey {
	return MgetKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Mget() (c Mget) {
	c.cs = append(b.get(), "MGET")
	return
}

type MgetKey struct {
	cs []string
}

func (c MgetKey) Key(Key ...string) MgetKey {
	return MgetKey{cs: append(c.cs, Key...)}
}

func (c MgetKey) Build() []string {
	return c.cs
}

type Migrate struct {
	cs []string
}

func (c Migrate) Host(Host string) MigrateHost {
	return MigrateHost{cs: append(c.cs, Host)}
}

func (b *Builder) Migrate() (c Migrate) {
	c.cs = append(b.get(), "MIGRATE")
	return
}

type MigrateAuth struct {
	cs []string
}

func (c MigrateAuth) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateAuth) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...)}
}

func (c MigrateAuth) Build() []string {
	return c.cs
}

type MigrateAuth2 struct {
	cs []string
}

func (c MigrateAuth2) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...)}
}

func (c MigrateAuth2) Build() []string {
	return c.cs
}

type MigrateCopyCopy struct {
	cs []string
}

func (c MigrateCopyCopy) Replace() MigrateReplaceReplace {
	return MigrateReplaceReplace{cs: append(c.cs, "REPLACE")}
}

func (c MigrateCopyCopy) Auth(Password string) MigrateAuth {
	return MigrateAuth{cs: append(c.cs, "AUTH", Password)}
}

func (c MigrateCopyCopy) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateCopyCopy) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...)}
}

func (c MigrateCopyCopy) Build() []string {
	return c.cs
}

type MigrateDestinationDb struct {
	cs []string
}

func (c MigrateDestinationDb) Timeout(Timeout int64) MigrateTimeout {
	return MigrateTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type MigrateHost struct {
	cs []string
}

func (c MigrateHost) Port(Port string) MigratePort {
	return MigratePort{cs: append(c.cs, Port)}
}

type MigrateKeyEmpty struct {
	cs []string
}

func (c MigrateKeyEmpty) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeyKey struct {
	cs []string
}

func (c MigrateKeyKey) DestinationDb(DestinationDb int64) MigrateDestinationDb {
	return MigrateDestinationDb{cs: append(c.cs, strconv.FormatInt(DestinationDb, 10))}
}

type MigrateKeys struct {
	cs []string
}

func (c MigrateKeys) Keys(Keys ...string) MigrateKeys {
	return MigrateKeys{cs: append(c.cs, Keys...)}
}

func (c MigrateKeys) Build() []string {
	return c.cs
}

type MigratePort struct {
	cs []string
}

func (c MigratePort) Key() MigrateKeyKey {
	return MigrateKeyKey{cs: append(c.cs, "key")}
}

func (c MigratePort) Empty() MigrateKeyEmpty {
	return MigrateKeyEmpty{cs: append(c.cs, "\"\"")}
}

type MigrateReplaceReplace struct {
	cs []string
}

func (c MigrateReplaceReplace) Auth(Password string) MigrateAuth {
	return MigrateAuth{cs: append(c.cs, "AUTH", Password)}
}

func (c MigrateReplaceReplace) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateReplaceReplace) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...)}
}

func (c MigrateReplaceReplace) Build() []string {
	return c.cs
}

type MigrateTimeout struct {
	cs []string
}

func (c MigrateTimeout) Copy() MigrateCopyCopy {
	return MigrateCopyCopy{cs: append(c.cs, "COPY")}
}

func (c MigrateTimeout) Replace() MigrateReplaceReplace {
	return MigrateReplaceReplace{cs: append(c.cs, "REPLACE")}
}

func (c MigrateTimeout) Auth(Password string) MigrateAuth {
	return MigrateAuth{cs: append(c.cs, "AUTH", Password)}
}

func (c MigrateTimeout) Auth2(UsernamePassword string) MigrateAuth2 {
	return MigrateAuth2{cs: append(c.cs, "AUTH2", UsernamePassword)}
}

func (c MigrateTimeout) Keys(Key ...string) MigrateKeys {
	c.cs = append(c.cs, "KEYS")
	return MigrateKeys{cs: append(c.cs, Key...)}
}

func (c MigrateTimeout) Build() []string {
	return c.cs
}

type ModuleList struct {
	cs []string
}

func (c ModuleList) Build() []string {
	return c.cs
}

func (b *Builder) ModuleList() (c ModuleList) {
	c.cs = append(b.get(), "MODULE", "LIST")
	return
}

type ModuleLoad struct {
	cs []string
}

func (c ModuleLoad) Path(Path string) ModuleLoadPath {
	return ModuleLoadPath{cs: append(c.cs, Path)}
}

func (b *Builder) ModuleLoad() (c ModuleLoad) {
	c.cs = append(b.get(), "MODULE", "LOAD")
	return
}

type ModuleLoadArg struct {
	cs []string
}

func (c ModuleLoadArg) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cs: append(c.cs, Arg...)}
}

func (c ModuleLoadArg) Build() []string {
	return c.cs
}

type ModuleLoadPath struct {
	cs []string
}

func (c ModuleLoadPath) Arg(Arg ...string) ModuleLoadArg {
	return ModuleLoadArg{cs: append(c.cs, Arg...)}
}

func (c ModuleLoadPath) Build() []string {
	return c.cs
}

type ModuleUnload struct {
	cs []string
}

func (c ModuleUnload) Name(Name string) ModuleUnloadName {
	return ModuleUnloadName{cs: append(c.cs, Name)}
}

func (b *Builder) ModuleUnload() (c ModuleUnload) {
	c.cs = append(b.get(), "MODULE", "UNLOAD")
	return
}

type ModuleUnloadName struct {
	cs []string
}

func (c ModuleUnloadName) Build() []string {
	return c.cs
}

type Monitor struct {
	cs []string
}

func (c Monitor) Build() []string {
	return c.cs
}

func (b *Builder) Monitor() (c Monitor) {
	c.cs = append(b.get(), "MONITOR")
	return
}

type Move struct {
	cs []string
}

func (c Move) Key(Key string) MoveKey {
	return MoveKey{cs: append(c.cs, Key)}
}

func (b *Builder) Move() (c Move) {
	c.cs = append(b.get(), "MOVE")
	return
}

type MoveDb struct {
	cs []string
}

func (c MoveDb) Build() []string {
	return c.cs
}

type MoveKey struct {
	cs []string
}

func (c MoveKey) Db(Db int64) MoveDb {
	return MoveDb{cs: append(c.cs, strconv.FormatInt(Db, 10))}
}

type Mset struct {
	cs []string
}

func (c Mset) KeyValue() MsetKeyValue {
	return MsetKeyValue{cs: append(c.cs)}
}

func (b *Builder) Mset() (c Mset) {
	c.cs = append(b.get(), "MSET")
	return
}

type MsetKeyValue struct {
	cs []string
}

func (c MsetKeyValue) KeyValue(Key string, Value string) MsetKeyValue {
	return MsetKeyValue{cs: append(c.cs, Key, Value)}
}

func (c MsetKeyValue) Build() []string {
	return c.cs
}

type Msetnx struct {
	cs []string
}

func (c Msetnx) KeyValue() MsetnxKeyValue {
	return MsetnxKeyValue{cs: append(c.cs)}
}

func (b *Builder) Msetnx() (c Msetnx) {
	c.cs = append(b.get(), "MSETNX")
	return
}

type MsetnxKeyValue struct {
	cs []string
}

func (c MsetnxKeyValue) KeyValue(Key string, Value string) MsetnxKeyValue {
	return MsetnxKeyValue{cs: append(c.cs, Key, Value)}
}

func (c MsetnxKeyValue) Build() []string {
	return c.cs
}

type Multi struct {
	cs []string
}

func (c Multi) Build() []string {
	return c.cs
}

func (b *Builder) Multi() (c Multi) {
	c.cs = append(b.get(), "MULTI")
	return
}

type Object struct {
	cs []string
}

func (c Object) Subcommand(Subcommand string) ObjectSubcommand {
	return ObjectSubcommand{cs: append(c.cs, Subcommand)}
}

func (b *Builder) Object() (c Object) {
	c.cs = append(b.get(), "OBJECT")
	return
}

type ObjectArguments struct {
	cs []string
}

func (c ObjectArguments) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cs: append(c.cs, Arguments...)}
}

func (c ObjectArguments) Build() []string {
	return c.cs
}

type ObjectSubcommand struct {
	cs []string
}

func (c ObjectSubcommand) Arguments(Arguments ...string) ObjectArguments {
	return ObjectArguments{cs: append(c.cs, Arguments...)}
}

func (c ObjectSubcommand) Build() []string {
	return c.cs
}

type Persist struct {
	cs []string
}

func (c Persist) Key(Key string) PersistKey {
	return PersistKey{cs: append(c.cs, Key)}
}

func (b *Builder) Persist() (c Persist) {
	c.cs = append(b.get(), "PERSIST")
	return
}

type PersistKey struct {
	cs []string
}

func (c PersistKey) Build() []string {
	return c.cs
}

type Pexpire struct {
	cs []string
}

func (c Pexpire) Key(Key string) PexpireKey {
	return PexpireKey{cs: append(c.cs, Key)}
}

func (b *Builder) Pexpire() (c Pexpire) {
	c.cs = append(b.get(), "PEXPIRE")
	return
}

type PexpireConditionGt struct {
	cs []string
}

func (c PexpireConditionGt) Build() []string {
	return c.cs
}

type PexpireConditionLt struct {
	cs []string
}

func (c PexpireConditionLt) Build() []string {
	return c.cs
}

type PexpireConditionNx struct {
	cs []string
}

func (c PexpireConditionNx) Build() []string {
	return c.cs
}

type PexpireConditionXx struct {
	cs []string
}

func (c PexpireConditionXx) Build() []string {
	return c.cs
}

type PexpireKey struct {
	cs []string
}

func (c PexpireKey) Milliseconds(Milliseconds int64) PexpireMilliseconds {
	return PexpireMilliseconds{cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PexpireMilliseconds struct {
	cs []string
}

func (c PexpireMilliseconds) Nx() PexpireConditionNx {
	return PexpireConditionNx{cs: append(c.cs, "NX")}
}

func (c PexpireMilliseconds) Xx() PexpireConditionXx {
	return PexpireConditionXx{cs: append(c.cs, "XX")}
}

func (c PexpireMilliseconds) Gt() PexpireConditionGt {
	return PexpireConditionGt{cs: append(c.cs, "GT")}
}

func (c PexpireMilliseconds) Lt() PexpireConditionLt {
	return PexpireConditionLt{cs: append(c.cs, "LT")}
}

func (c PexpireMilliseconds) Build() []string {
	return c.cs
}

type Pexpireat struct {
	cs []string
}

func (c Pexpireat) Key(Key string) PexpireatKey {
	return PexpireatKey{cs: append(c.cs, Key)}
}

func (b *Builder) Pexpireat() (c Pexpireat) {
	c.cs = append(b.get(), "PEXPIREAT")
	return
}

type PexpireatConditionGt struct {
	cs []string
}

func (c PexpireatConditionGt) Build() []string {
	return c.cs
}

type PexpireatConditionLt struct {
	cs []string
}

func (c PexpireatConditionLt) Build() []string {
	return c.cs
}

type PexpireatConditionNx struct {
	cs []string
}

func (c PexpireatConditionNx) Build() []string {
	return c.cs
}

type PexpireatConditionXx struct {
	cs []string
}

func (c PexpireatConditionXx) Build() []string {
	return c.cs
}

type PexpireatKey struct {
	cs []string
}

func (c PexpireatKey) MillisecondsTimestamp(MillisecondsTimestamp int64) PexpireatMillisecondsTimestamp {
	return PexpireatMillisecondsTimestamp{cs: append(c.cs, strconv.FormatInt(MillisecondsTimestamp, 10))}
}

type PexpireatMillisecondsTimestamp struct {
	cs []string
}

func (c PexpireatMillisecondsTimestamp) Nx() PexpireatConditionNx {
	return PexpireatConditionNx{cs: append(c.cs, "NX")}
}

func (c PexpireatMillisecondsTimestamp) Xx() PexpireatConditionXx {
	return PexpireatConditionXx{cs: append(c.cs, "XX")}
}

func (c PexpireatMillisecondsTimestamp) Gt() PexpireatConditionGt {
	return PexpireatConditionGt{cs: append(c.cs, "GT")}
}

func (c PexpireatMillisecondsTimestamp) Lt() PexpireatConditionLt {
	return PexpireatConditionLt{cs: append(c.cs, "LT")}
}

func (c PexpireatMillisecondsTimestamp) Build() []string {
	return c.cs
}

type Pexpiretime struct {
	cs []string
}

func (c Pexpiretime) Key(Key string) PexpiretimeKey {
	return PexpiretimeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Pexpiretime() (c Pexpiretime) {
	c.cs = append(b.get(), "PEXPIRETIME")
	return
}

type PexpiretimeKey struct {
	cs []string
}

func (c PexpiretimeKey) Build() []string {
	return c.cs
}

type Pfadd struct {
	cs []string
}

func (c Pfadd) Key(Key string) PfaddKey {
	return PfaddKey{cs: append(c.cs, Key)}
}

func (b *Builder) Pfadd() (c Pfadd) {
	c.cs = append(b.get(), "PFADD")
	return
}

type PfaddElement struct {
	cs []string
}

func (c PfaddElement) Element(Element ...string) PfaddElement {
	return PfaddElement{cs: append(c.cs, Element...)}
}

func (c PfaddElement) Build() []string {
	return c.cs
}

type PfaddKey struct {
	cs []string
}

func (c PfaddKey) Element(Element ...string) PfaddElement {
	return PfaddElement{cs: append(c.cs, Element...)}
}

func (c PfaddKey) Build() []string {
	return c.cs
}

type Pfcount struct {
	cs []string
}

func (c Pfcount) Key(Key ...string) PfcountKey {
	return PfcountKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Pfcount() (c Pfcount) {
	c.cs = append(b.get(), "PFCOUNT")
	return
}

type PfcountKey struct {
	cs []string
}

func (c PfcountKey) Key(Key ...string) PfcountKey {
	return PfcountKey{cs: append(c.cs, Key...)}
}

func (c PfcountKey) Build() []string {
	return c.cs
}

type Pfmerge struct {
	cs []string
}

func (c Pfmerge) Destkey(Destkey string) PfmergeDestkey {
	return PfmergeDestkey{cs: append(c.cs, Destkey)}
}

func (b *Builder) Pfmerge() (c Pfmerge) {
	c.cs = append(b.get(), "PFMERGE")
	return
}

type PfmergeDestkey struct {
	cs []string
}

func (c PfmergeDestkey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cs: append(c.cs, Sourcekey...)}
}

type PfmergeSourcekey struct {
	cs []string
}

func (c PfmergeSourcekey) Sourcekey(Sourcekey ...string) PfmergeSourcekey {
	return PfmergeSourcekey{cs: append(c.cs, Sourcekey...)}
}

func (c PfmergeSourcekey) Build() []string {
	return c.cs
}

type Ping struct {
	cs []string
}

func (c Ping) Message(Message string) PingMessage {
	return PingMessage{cs: append(c.cs, Message)}
}

func (c Ping) Build() []string {
	return c.cs
}

func (b *Builder) Ping() (c Ping) {
	c.cs = append(b.get(), "PING")
	return
}

type PingMessage struct {
	cs []string
}

func (c PingMessage) Build() []string {
	return c.cs
}

type Psetex struct {
	cs []string
}

func (c Psetex) Key(Key string) PsetexKey {
	return PsetexKey{cs: append(c.cs, Key)}
}

func (b *Builder) Psetex() (c Psetex) {
	c.cs = append(b.get(), "PSETEX")
	return
}

type PsetexKey struct {
	cs []string
}

func (c PsetexKey) Milliseconds(Milliseconds int64) PsetexMilliseconds {
	return PsetexMilliseconds{cs: append(c.cs, strconv.FormatInt(Milliseconds, 10))}
}

type PsetexMilliseconds struct {
	cs []string
}

func (c PsetexMilliseconds) Value(Value string) PsetexValue {
	return PsetexValue{cs: append(c.cs, Value)}
}

type PsetexValue struct {
	cs []string
}

func (c PsetexValue) Build() []string {
	return c.cs
}

type Psubscribe struct {
	cs []string
}

func (c Psubscribe) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cs: append(c.cs, Pattern...)}
}

func (b *Builder) Psubscribe() (c Psubscribe) {
	c.cs = append(b.get(), "PSUBSCRIBE")
	return
}

type PsubscribePattern struct {
	cs []string
}

func (c PsubscribePattern) Pattern(Pattern ...string) PsubscribePattern {
	return PsubscribePattern{cs: append(c.cs, Pattern...)}
}

func (c PsubscribePattern) Build() []string {
	return c.cs
}

type Psync struct {
	cs []string
}

func (c Psync) Replicationid(Replicationid int64) PsyncReplicationid {
	return PsyncReplicationid{cs: append(c.cs, strconv.FormatInt(Replicationid, 10))}
}

func (b *Builder) Psync() (c Psync) {
	c.cs = append(b.get(), "PSYNC")
	return
}

type PsyncOffset struct {
	cs []string
}

func (c PsyncOffset) Build() []string {
	return c.cs
}

type PsyncReplicationid struct {
	cs []string
}

func (c PsyncReplicationid) Offset(Offset int64) PsyncOffset {
	return PsyncOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type Pttl struct {
	cs []string
}

func (c Pttl) Key(Key string) PttlKey {
	return PttlKey{cs: append(c.cs, Key)}
}

func (b *Builder) Pttl() (c Pttl) {
	c.cs = append(b.get(), "PTTL")
	return
}

type PttlKey struct {
	cs []string
}

func (c PttlKey) Build() []string {
	return c.cs
}

type Publish struct {
	cs []string
}

func (c Publish) Channel(Channel string) PublishChannel {
	return PublishChannel{cs: append(c.cs, Channel)}
}

func (b *Builder) Publish() (c Publish) {
	c.cs = append(b.get(), "PUBLISH")
	return
}

type PublishChannel struct {
	cs []string
}

func (c PublishChannel) Message(Message string) PublishMessage {
	return PublishMessage{cs: append(c.cs, Message)}
}

type PublishMessage struct {
	cs []string
}

func (c PublishMessage) Build() []string {
	return c.cs
}

type Pubsub struct {
	cs []string
}

func (c Pubsub) Subcommand(Subcommand string) PubsubSubcommand {
	return PubsubSubcommand{cs: append(c.cs, Subcommand)}
}

func (b *Builder) Pubsub() (c Pubsub) {
	c.cs = append(b.get(), "PUBSUB")
	return
}

type PubsubArgument struct {
	cs []string
}

func (c PubsubArgument) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cs: append(c.cs, Argument...)}
}

func (c PubsubArgument) Build() []string {
	return c.cs
}

type PubsubSubcommand struct {
	cs []string
}

func (c PubsubSubcommand) Argument(Argument ...string) PubsubArgument {
	return PubsubArgument{cs: append(c.cs, Argument...)}
}

func (c PubsubSubcommand) Build() []string {
	return c.cs
}

type Punsubscribe struct {
	cs []string
}

func (c Punsubscribe) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cs: append(c.cs, Pattern...)}
}

func (c Punsubscribe) Build() []string {
	return c.cs
}

func (b *Builder) Punsubscribe() (c Punsubscribe) {
	c.cs = append(b.get(), "PUNSUBSCRIBE")
	return
}

type PunsubscribePattern struct {
	cs []string
}

func (c PunsubscribePattern) Pattern(Pattern ...string) PunsubscribePattern {
	return PunsubscribePattern{cs: append(c.cs, Pattern...)}
}

func (c PunsubscribePattern) Build() []string {
	return c.cs
}

type Quit struct {
	cs []string
}

func (c Quit) Build() []string {
	return c.cs
}

func (b *Builder) Quit() (c Quit) {
	c.cs = append(b.get(), "QUIT")
	return
}

type Randomkey struct {
	cs []string
}

func (c Randomkey) Build() []string {
	return c.cs
}

func (b *Builder) Randomkey() (c Randomkey) {
	c.cs = append(b.get(), "RANDOMKEY")
	return
}

type Readonly struct {
	cs []string
}

func (c Readonly) Build() []string {
	return c.cs
}

func (b *Builder) Readonly() (c Readonly) {
	c.cs = append(b.get(), "READONLY")
	return
}

type Readwrite struct {
	cs []string
}

func (c Readwrite) Build() []string {
	return c.cs
}

func (b *Builder) Readwrite() (c Readwrite) {
	c.cs = append(b.get(), "READWRITE")
	return
}

type Rename struct {
	cs []string
}

func (c Rename) Key(Key string) RenameKey {
	return RenameKey{cs: append(c.cs, Key)}
}

func (b *Builder) Rename() (c Rename) {
	c.cs = append(b.get(), "RENAME")
	return
}

type RenameKey struct {
	cs []string
}

func (c RenameKey) Newkey(Newkey string) RenameNewkey {
	return RenameNewkey{cs: append(c.cs, Newkey)}
}

type RenameNewkey struct {
	cs []string
}

func (c RenameNewkey) Build() []string {
	return c.cs
}

type Renamenx struct {
	cs []string
}

func (c Renamenx) Key(Key string) RenamenxKey {
	return RenamenxKey{cs: append(c.cs, Key)}
}

func (b *Builder) Renamenx() (c Renamenx) {
	c.cs = append(b.get(), "RENAMENX")
	return
}

type RenamenxKey struct {
	cs []string
}

func (c RenamenxKey) Newkey(Newkey string) RenamenxNewkey {
	return RenamenxNewkey{cs: append(c.cs, Newkey)}
}

type RenamenxNewkey struct {
	cs []string
}

func (c RenamenxNewkey) Build() []string {
	return c.cs
}

type Replicaof struct {
	cs []string
}

func (c Replicaof) Host(Host string) ReplicaofHost {
	return ReplicaofHost{cs: append(c.cs, Host)}
}

func (b *Builder) Replicaof() (c Replicaof) {
	c.cs = append(b.get(), "REPLICAOF")
	return
}

type ReplicaofHost struct {
	cs []string
}

func (c ReplicaofHost) Port(Port string) ReplicaofPort {
	return ReplicaofPort{cs: append(c.cs, Port)}
}

type ReplicaofPort struct {
	cs []string
}

func (c ReplicaofPort) Build() []string {
	return c.cs
}

type Reset struct {
	cs []string
}

func (c Reset) Build() []string {
	return c.cs
}

func (b *Builder) Reset() (c Reset) {
	c.cs = append(b.get(), "RESET")
	return
}

type Restore struct {
	cs []string
}

func (c Restore) Key(Key string) RestoreKey {
	return RestoreKey{cs: append(c.cs, Key)}
}

func (b *Builder) Restore() (c Restore) {
	c.cs = append(b.get(), "RESTORE")
	return
}

type RestoreAbsttlAbsttl struct {
	cs []string
}

func (c RestoreAbsttlAbsttl) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c RestoreAbsttlAbsttl) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreAbsttlAbsttl) Build() []string {
	return c.cs
}

type RestoreFreq struct {
	cs []string
}

func (c RestoreFreq) Build() []string {
	return c.cs
}

type RestoreIdletime struct {
	cs []string
}

func (c RestoreIdletime) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreIdletime) Build() []string {
	return c.cs
}

type RestoreKey struct {
	cs []string
}

func (c RestoreKey) Ttl(Ttl int64) RestoreTtl {
	return RestoreTtl{cs: append(c.cs, strconv.FormatInt(Ttl, 10))}
}

type RestoreReplaceReplace struct {
	cs []string
}

func (c RestoreReplaceReplace) Absttl() RestoreAbsttlAbsttl {
	return RestoreAbsttlAbsttl{cs: append(c.cs, "ABSTTL")}
}

func (c RestoreReplaceReplace) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c RestoreReplaceReplace) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreReplaceReplace) Build() []string {
	return c.cs
}

type RestoreSerializedValue struct {
	cs []string
}

func (c RestoreSerializedValue) Replace() RestoreReplaceReplace {
	return RestoreReplaceReplace{cs: append(c.cs, "REPLACE")}
}

func (c RestoreSerializedValue) Absttl() RestoreAbsttlAbsttl {
	return RestoreAbsttlAbsttl{cs: append(c.cs, "ABSTTL")}
}

func (c RestoreSerializedValue) Idletime(Seconds int64) RestoreIdletime {
	return RestoreIdletime{cs: append(c.cs, "IDLETIME", strconv.FormatInt(Seconds, 10))}
}

func (c RestoreSerializedValue) Freq(Frequency int64) RestoreFreq {
	return RestoreFreq{cs: append(c.cs, "FREQ", strconv.FormatInt(Frequency, 10))}
}

func (c RestoreSerializedValue) Build() []string {
	return c.cs
}

type RestoreTtl struct {
	cs []string
}

func (c RestoreTtl) SerializedValue(SerializedValue string) RestoreSerializedValue {
	return RestoreSerializedValue{cs: append(c.cs, SerializedValue)}
}

type Role struct {
	cs []string
}

func (c Role) Build() []string {
	return c.cs
}

func (b *Builder) Role() (c Role) {
	c.cs = append(b.get(), "ROLE")
	return
}

type Rpop struct {
	cs []string
}

func (c Rpop) Key(Key string) RpopKey {
	return RpopKey{cs: append(c.cs, Key)}
}

func (b *Builder) Rpop() (c Rpop) {
	c.cs = append(b.get(), "RPOP")
	return
}

type RpopCount struct {
	cs []string
}

func (c RpopCount) Build() []string {
	return c.cs
}

type RpopKey struct {
	cs []string
}

func (c RpopKey) Count(Count int64) RpopCount {
	return RpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c RpopKey) Build() []string {
	return c.cs
}

type Rpoplpush struct {
	cs []string
}

func (c Rpoplpush) Source(Source string) RpoplpushSource {
	return RpoplpushSource{cs: append(c.cs, Source)}
}

func (b *Builder) Rpoplpush() (c Rpoplpush) {
	c.cs = append(b.get(), "RPOPLPUSH")
	return
}

type RpoplpushDestination struct {
	cs []string
}

func (c RpoplpushDestination) Build() []string {
	return c.cs
}

type RpoplpushSource struct {
	cs []string
}

func (c RpoplpushSource) Destination(Destination string) RpoplpushDestination {
	return RpoplpushDestination{cs: append(c.cs, Destination)}
}

type Rpush struct {
	cs []string
}

func (c Rpush) Key(Key string) RpushKey {
	return RpushKey{cs: append(c.cs, Key)}
}

func (b *Builder) Rpush() (c Rpush) {
	c.cs = append(b.get(), "RPUSH")
	return
}

type RpushElement struct {
	cs []string
}

func (c RpushElement) Element(Element ...string) RpushElement {
	return RpushElement{cs: append(c.cs, Element...)}
}

func (c RpushElement) Build() []string {
	return c.cs
}

type RpushKey struct {
	cs []string
}

func (c RpushKey) Element(Element ...string) RpushElement {
	return RpushElement{cs: append(c.cs, Element...)}
}

type Rpushx struct {
	cs []string
}

func (c Rpushx) Key(Key string) RpushxKey {
	return RpushxKey{cs: append(c.cs, Key)}
}

func (b *Builder) Rpushx() (c Rpushx) {
	c.cs = append(b.get(), "RPUSHX")
	return
}

type RpushxElement struct {
	cs []string
}

func (c RpushxElement) Element(Element ...string) RpushxElement {
	return RpushxElement{cs: append(c.cs, Element...)}
}

func (c RpushxElement) Build() []string {
	return c.cs
}

type RpushxKey struct {
	cs []string
}

func (c RpushxKey) Element(Element ...string) RpushxElement {
	return RpushxElement{cs: append(c.cs, Element...)}
}

type Sadd struct {
	cs []string
}

func (c Sadd) Key(Key string) SaddKey {
	return SaddKey{cs: append(c.cs, Key)}
}

func (b *Builder) Sadd() (c Sadd) {
	c.cs = append(b.get(), "SADD")
	return
}

type SaddKey struct {
	cs []string
}

func (c SaddKey) Member(Member ...string) SaddMember {
	return SaddMember{cs: append(c.cs, Member...)}
}

type SaddMember struct {
	cs []string
}

func (c SaddMember) Member(Member ...string) SaddMember {
	return SaddMember{cs: append(c.cs, Member...)}
}

func (c SaddMember) Build() []string {
	return c.cs
}

type Save struct {
	cs []string
}

func (c Save) Build() []string {
	return c.cs
}

func (b *Builder) Save() (c Save) {
	c.cs = append(b.get(), "SAVE")
	return
}

type Scan struct {
	cs []string
}

func (c Scan) Cursor(Cursor int64) ScanCursor {
	return ScanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

func (b *Builder) Scan() (c Scan) {
	c.cs = append(b.get(), "SCAN")
	return
}

type ScanCount struct {
	cs []string
}

func (c ScanCount) Type(Type string) ScanType {
	return ScanType{cs: append(c.cs, "TYPE", Type)}
}

func (c ScanCount) Build() []string {
	return c.cs
}

type ScanCursor struct {
	cs []string
}

func (c ScanCursor) Match(Pattern string) ScanMatch {
	return ScanMatch{cs: append(c.cs, "MATCH", Pattern)}
}

func (c ScanCursor) Count(Count int64) ScanCount {
	return ScanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ScanCursor) Type(Type string) ScanType {
	return ScanType{cs: append(c.cs, "TYPE", Type)}
}

func (c ScanCursor) Build() []string {
	return c.cs
}

type ScanMatch struct {
	cs []string
}

func (c ScanMatch) Count(Count int64) ScanCount {
	return ScanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ScanMatch) Type(Type string) ScanType {
	return ScanType{cs: append(c.cs, "TYPE", Type)}
}

func (c ScanMatch) Build() []string {
	return c.cs
}

type ScanType struct {
	cs []string
}

func (c ScanType) Build() []string {
	return c.cs
}

type Scard struct {
	cs []string
}

func (c Scard) Key(Key string) ScardKey {
	return ScardKey{cs: append(c.cs, Key)}
}

func (b *Builder) Scard() (c Scard) {
	c.cs = append(b.get(), "SCARD")
	return
}

type ScardKey struct {
	cs []string
}

func (c ScardKey) Build() []string {
	return c.cs
}

type ScriptDebug struct {
	cs []string
}

func (c ScriptDebug) Yes() ScriptDebugModeYes {
	return ScriptDebugModeYes{cs: append(c.cs, "YES")}
}

func (c ScriptDebug) Sync() ScriptDebugModeSync {
	return ScriptDebugModeSync{cs: append(c.cs, "SYNC")}
}

func (c ScriptDebug) No() ScriptDebugModeNo {
	return ScriptDebugModeNo{cs: append(c.cs, "NO")}
}

func (b *Builder) ScriptDebug() (c ScriptDebug) {
	c.cs = append(b.get(), "SCRIPT", "DEBUG")
	return
}

type ScriptDebugModeNo struct {
	cs []string
}

func (c ScriptDebugModeNo) Build() []string {
	return c.cs
}

type ScriptDebugModeSync struct {
	cs []string
}

func (c ScriptDebugModeSync) Build() []string {
	return c.cs
}

type ScriptDebugModeYes struct {
	cs []string
}

func (c ScriptDebugModeYes) Build() []string {
	return c.cs
}

type ScriptExists struct {
	cs []string
}

func (c ScriptExists) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cs: append(c.cs, Sha1...)}
}

func (b *Builder) ScriptExists() (c ScriptExists) {
	c.cs = append(b.get(), "SCRIPT", "EXISTS")
	return
}

type ScriptExistsSha1 struct {
	cs []string
}

func (c ScriptExistsSha1) Sha1(Sha1 ...string) ScriptExistsSha1 {
	return ScriptExistsSha1{cs: append(c.cs, Sha1...)}
}

func (c ScriptExistsSha1) Build() []string {
	return c.cs
}

type ScriptFlush struct {
	cs []string
}

func (c ScriptFlush) Async() ScriptFlushAsyncAsync {
	return ScriptFlushAsyncAsync{cs: append(c.cs, "ASYNC")}
}

func (c ScriptFlush) Sync() ScriptFlushAsyncSync {
	return ScriptFlushAsyncSync{cs: append(c.cs, "SYNC")}
}

func (c ScriptFlush) Build() []string {
	return c.cs
}

func (b *Builder) ScriptFlush() (c ScriptFlush) {
	c.cs = append(b.get(), "SCRIPT", "FLUSH")
	return
}

type ScriptFlushAsyncAsync struct {
	cs []string
}

func (c ScriptFlushAsyncAsync) Build() []string {
	return c.cs
}

type ScriptFlushAsyncSync struct {
	cs []string
}

func (c ScriptFlushAsyncSync) Build() []string {
	return c.cs
}

type ScriptKill struct {
	cs []string
}

func (c ScriptKill) Build() []string {
	return c.cs
}

func (b *Builder) ScriptKill() (c ScriptKill) {
	c.cs = append(b.get(), "SCRIPT", "KILL")
	return
}

type ScriptLoad struct {
	cs []string
}

func (c ScriptLoad) Script(Script string) ScriptLoadScript {
	return ScriptLoadScript{cs: append(c.cs, Script)}
}

func (b *Builder) ScriptLoad() (c ScriptLoad) {
	c.cs = append(b.get(), "SCRIPT", "LOAD")
	return
}

type ScriptLoadScript struct {
	cs []string
}

func (c ScriptLoadScript) Build() []string {
	return c.cs
}

type Sdiff struct {
	cs []string
}

func (c Sdiff) Key(Key ...string) SdiffKey {
	return SdiffKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Sdiff() (c Sdiff) {
	c.cs = append(b.get(), "SDIFF")
	return
}

type SdiffKey struct {
	cs []string
}

func (c SdiffKey) Key(Key ...string) SdiffKey {
	return SdiffKey{cs: append(c.cs, Key...)}
}

func (c SdiffKey) Build() []string {
	return c.cs
}

type Sdiffstore struct {
	cs []string
}

func (c Sdiffstore) Destination(Destination string) SdiffstoreDestination {
	return SdiffstoreDestination{cs: append(c.cs, Destination)}
}

func (b *Builder) Sdiffstore() (c Sdiffstore) {
	c.cs = append(b.get(), "SDIFFSTORE")
	return
}

type SdiffstoreDestination struct {
	cs []string
}

func (c SdiffstoreDestination) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cs: append(c.cs, Key...)}
}

type SdiffstoreKey struct {
	cs []string
}

func (c SdiffstoreKey) Key(Key ...string) SdiffstoreKey {
	return SdiffstoreKey{cs: append(c.cs, Key...)}
}

func (c SdiffstoreKey) Build() []string {
	return c.cs
}

type Select struct {
	cs []string
}

func (c Select) Index(Index int64) SelectIndex {
	return SelectIndex{cs: append(c.cs, strconv.FormatInt(Index, 10))}
}

func (b *Builder) Select() (c Select) {
	c.cs = append(b.get(), "SELECT")
	return
}

type SelectIndex struct {
	cs []string
}

func (c SelectIndex) Build() []string {
	return c.cs
}

type Set struct {
	cs []string
}

func (c Set) Key(Key string) SetKey {
	return SetKey{cs: append(c.cs, Key)}
}

func (b *Builder) Set() (c Set) {
	c.cs = append(b.get(), "SET")
	return
}

type SetConditionNx struct {
	cs []string
}

func (c SetConditionNx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetConditionNx) Build() []string {
	return c.cs
}

type SetConditionXx struct {
	cs []string
}

func (c SetConditionXx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetConditionXx) Build() []string {
	return c.cs
}

type SetExpirationEx struct {
	cs []string
}

func (c SetExpirationEx) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX")}
}

func (c SetExpirationEx) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX")}
}

func (c SetExpirationEx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetExpirationEx) Build() []string {
	return c.cs
}

type SetExpirationExat struct {
	cs []string
}

func (c SetExpirationExat) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX")}
}

func (c SetExpirationExat) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX")}
}

func (c SetExpirationExat) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetExpirationExat) Build() []string {
	return c.cs
}

type SetExpirationKeepttl struct {
	cs []string
}

func (c SetExpirationKeepttl) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX")}
}

func (c SetExpirationKeepttl) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX")}
}

func (c SetExpirationKeepttl) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetExpirationKeepttl) Build() []string {
	return c.cs
}

type SetExpirationPx struct {
	cs []string
}

func (c SetExpirationPx) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX")}
}

func (c SetExpirationPx) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX")}
}

func (c SetExpirationPx) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetExpirationPx) Build() []string {
	return c.cs
}

type SetExpirationPxat struct {
	cs []string
}

func (c SetExpirationPxat) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX")}
}

func (c SetExpirationPxat) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX")}
}

func (c SetExpirationPxat) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetExpirationPxat) Build() []string {
	return c.cs
}

type SetGetGet struct {
	cs []string
}

func (c SetGetGet) Build() []string {
	return c.cs
}

type SetKey struct {
	cs []string
}

func (c SetKey) Value(Value string) SetValue {
	return SetValue{cs: append(c.cs, Value)}
}

type SetValue struct {
	cs []string
}

func (c SetValue) Ex(Seconds int64) SetExpirationEx {
	return SetExpirationEx{cs: append(c.cs, "EX", strconv.FormatInt(Seconds, 10))}
}

func (c SetValue) Px(Milliseconds int64) SetExpirationPx {
	return SetExpirationPx{cs: append(c.cs, "PX", strconv.FormatInt(Milliseconds, 10))}
}

func (c SetValue) Exat(Timestamp int64) SetExpirationExat {
	return SetExpirationExat{cs: append(c.cs, "EXAT", strconv.FormatInt(Timestamp, 10))}
}

func (c SetValue) Pxat(Millisecondstimestamp int64) SetExpirationPxat {
	return SetExpirationPxat{cs: append(c.cs, "PXAT", strconv.FormatInt(Millisecondstimestamp, 10))}
}

func (c SetValue) Keepttl() SetExpirationKeepttl {
	return SetExpirationKeepttl{cs: append(c.cs, "KEEPTTL")}
}

func (c SetValue) Nx() SetConditionNx {
	return SetConditionNx{cs: append(c.cs, "NX")}
}

func (c SetValue) Xx() SetConditionXx {
	return SetConditionXx{cs: append(c.cs, "XX")}
}

func (c SetValue) Get() SetGetGet {
	return SetGetGet{cs: append(c.cs, "GET")}
}

func (c SetValue) Build() []string {
	return c.cs
}

type Setbit struct {
	cs []string
}

func (c Setbit) Key(Key string) SetbitKey {
	return SetbitKey{cs: append(c.cs, Key)}
}

func (b *Builder) Setbit() (c Setbit) {
	c.cs = append(b.get(), "SETBIT")
	return
}

type SetbitKey struct {
	cs []string
}

func (c SetbitKey) Offset(Offset int64) SetbitOffset {
	return SetbitOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetbitOffset struct {
	cs []string
}

func (c SetbitOffset) Value(Value int64) SetbitValue {
	return SetbitValue{cs: append(c.cs, strconv.FormatInt(Value, 10))}
}

type SetbitValue struct {
	cs []string
}

func (c SetbitValue) Build() []string {
	return c.cs
}

type Setex struct {
	cs []string
}

func (c Setex) Key(Key string) SetexKey {
	return SetexKey{cs: append(c.cs, Key)}
}

func (b *Builder) Setex() (c Setex) {
	c.cs = append(b.get(), "SETEX")
	return
}

type SetexKey struct {
	cs []string
}

func (c SetexKey) Seconds(Seconds int64) SetexSeconds {
	return SetexSeconds{cs: append(c.cs, strconv.FormatInt(Seconds, 10))}
}

type SetexSeconds struct {
	cs []string
}

func (c SetexSeconds) Value(Value string) SetexValue {
	return SetexValue{cs: append(c.cs, Value)}
}

type SetexValue struct {
	cs []string
}

func (c SetexValue) Build() []string {
	return c.cs
}

type Setnx struct {
	cs []string
}

func (c Setnx) Key(Key string) SetnxKey {
	return SetnxKey{cs: append(c.cs, Key)}
}

func (b *Builder) Setnx() (c Setnx) {
	c.cs = append(b.get(), "SETNX")
	return
}

type SetnxKey struct {
	cs []string
}

func (c SetnxKey) Value(Value string) SetnxValue {
	return SetnxValue{cs: append(c.cs, Value)}
}

type SetnxValue struct {
	cs []string
}

func (c SetnxValue) Build() []string {
	return c.cs
}

type Setrange struct {
	cs []string
}

func (c Setrange) Key(Key string) SetrangeKey {
	return SetrangeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Setrange() (c Setrange) {
	c.cs = append(b.get(), "SETRANGE")
	return
}

type SetrangeKey struct {
	cs []string
}

func (c SetrangeKey) Offset(Offset int64) SetrangeOffset {
	return SetrangeOffset{cs: append(c.cs, strconv.FormatInt(Offset, 10))}
}

type SetrangeOffset struct {
	cs []string
}

func (c SetrangeOffset) Value(Value string) SetrangeValue {
	return SetrangeValue{cs: append(c.cs, Value)}
}

type SetrangeValue struct {
	cs []string
}

func (c SetrangeValue) Build() []string {
	return c.cs
}

type Shutdown struct {
	cs []string
}

func (c Shutdown) Nosave() ShutdownSaveModeNosave {
	return ShutdownSaveModeNosave{cs: append(c.cs, "NOSAVE")}
}

func (c Shutdown) Save() ShutdownSaveModeSave {
	return ShutdownSaveModeSave{cs: append(c.cs, "SAVE")}
}

func (c Shutdown) Build() []string {
	return c.cs
}

func (b *Builder) Shutdown() (c Shutdown) {
	c.cs = append(b.get(), "SHUTDOWN")
	return
}

type ShutdownSaveModeNosave struct {
	cs []string
}

func (c ShutdownSaveModeNosave) Build() []string {
	return c.cs
}

type ShutdownSaveModeSave struct {
	cs []string
}

func (c ShutdownSaveModeSave) Build() []string {
	return c.cs
}

type Sinter struct {
	cs []string
}

func (c Sinter) Key(Key ...string) SinterKey {
	return SinterKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Sinter() (c Sinter) {
	c.cs = append(b.get(), "SINTER")
	return
}

type SinterKey struct {
	cs []string
}

func (c SinterKey) Key(Key ...string) SinterKey {
	return SinterKey{cs: append(c.cs, Key...)}
}

func (c SinterKey) Build() []string {
	return c.cs
}

type Sintercard struct {
	cs []string
}

func (c Sintercard) Key(Key ...string) SintercardKey {
	return SintercardKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Sintercard() (c Sintercard) {
	c.cs = append(b.get(), "SINTERCARD")
	return
}

type SintercardKey struct {
	cs []string
}

func (c SintercardKey) Key(Key ...string) SintercardKey {
	return SintercardKey{cs: append(c.cs, Key...)}
}

func (c SintercardKey) Build() []string {
	return c.cs
}

type Sinterstore struct {
	cs []string
}

func (c Sinterstore) Destination(Destination string) SinterstoreDestination {
	return SinterstoreDestination{cs: append(c.cs, Destination)}
}

func (b *Builder) Sinterstore() (c Sinterstore) {
	c.cs = append(b.get(), "SINTERSTORE")
	return
}

type SinterstoreDestination struct {
	cs []string
}

func (c SinterstoreDestination) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cs: append(c.cs, Key...)}
}

type SinterstoreKey struct {
	cs []string
}

func (c SinterstoreKey) Key(Key ...string) SinterstoreKey {
	return SinterstoreKey{cs: append(c.cs, Key...)}
}

func (c SinterstoreKey) Build() []string {
	return c.cs
}

type Sismember struct {
	cs []string
}

func (c Sismember) Key(Key string) SismemberKey {
	return SismemberKey{cs: append(c.cs, Key)}
}

func (b *Builder) Sismember() (c Sismember) {
	c.cs = append(b.get(), "SISMEMBER")
	return
}

type SismemberKey struct {
	cs []string
}

func (c SismemberKey) Member(Member string) SismemberMember {
	return SismemberMember{cs: append(c.cs, Member)}
}

type SismemberMember struct {
	cs []string
}

func (c SismemberMember) Build() []string {
	return c.cs
}

type Slaveof struct {
	cs []string
}

func (c Slaveof) Host(Host string) SlaveofHost {
	return SlaveofHost{cs: append(c.cs, Host)}
}

func (b *Builder) Slaveof() (c Slaveof) {
	c.cs = append(b.get(), "SLAVEOF")
	return
}

type SlaveofHost struct {
	cs []string
}

func (c SlaveofHost) Port(Port string) SlaveofPort {
	return SlaveofPort{cs: append(c.cs, Port)}
}

type SlaveofPort struct {
	cs []string
}

func (c SlaveofPort) Build() []string {
	return c.cs
}

type Slowlog struct {
	cs []string
}

func (c Slowlog) Subcommand(Subcommand string) SlowlogSubcommand {
	return SlowlogSubcommand{cs: append(c.cs, Subcommand)}
}

func (b *Builder) Slowlog() (c Slowlog) {
	c.cs = append(b.get(), "SLOWLOG")
	return
}

type SlowlogArgument struct {
	cs []string
}

func (c SlowlogArgument) Build() []string {
	return c.cs
}

type SlowlogSubcommand struct {
	cs []string
}

func (c SlowlogSubcommand) Argument(Argument string) SlowlogArgument {
	return SlowlogArgument{cs: append(c.cs, Argument)}
}

func (c SlowlogSubcommand) Build() []string {
	return c.cs
}

type Smembers struct {
	cs []string
}

func (c Smembers) Key(Key string) SmembersKey {
	return SmembersKey{cs: append(c.cs, Key)}
}

func (b *Builder) Smembers() (c Smembers) {
	c.cs = append(b.get(), "SMEMBERS")
	return
}

type SmembersKey struct {
	cs []string
}

func (c SmembersKey) Build() []string {
	return c.cs
}

type Smismember struct {
	cs []string
}

func (c Smismember) Key(Key string) SmismemberKey {
	return SmismemberKey{cs: append(c.cs, Key)}
}

func (b *Builder) Smismember() (c Smismember) {
	c.cs = append(b.get(), "SMISMEMBER")
	return
}

type SmismemberKey struct {
	cs []string
}

func (c SmismemberKey) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cs: append(c.cs, Member...)}
}

type SmismemberMember struct {
	cs []string
}

func (c SmismemberMember) Member(Member ...string) SmismemberMember {
	return SmismemberMember{cs: append(c.cs, Member...)}
}

func (c SmismemberMember) Build() []string {
	return c.cs
}

type Smove struct {
	cs []string
}

func (c Smove) Source(Source string) SmoveSource {
	return SmoveSource{cs: append(c.cs, Source)}
}

func (b *Builder) Smove() (c Smove) {
	c.cs = append(b.get(), "SMOVE")
	return
}

type SmoveDestination struct {
	cs []string
}

func (c SmoveDestination) Member(Member string) SmoveMember {
	return SmoveMember{cs: append(c.cs, Member)}
}

type SmoveMember struct {
	cs []string
}

func (c SmoveMember) Build() []string {
	return c.cs
}

type SmoveSource struct {
	cs []string
}

func (c SmoveSource) Destination(Destination string) SmoveDestination {
	return SmoveDestination{cs: append(c.cs, Destination)}
}

type Sort struct {
	cs []string
}

func (c Sort) Key(Key string) SortKey {
	return SortKey{cs: append(c.cs, Key)}
}

func (b *Builder) Sort() (c Sort) {
	c.cs = append(b.get(), "SORT")
	return
}

type SortBy struct {
	cs []string
}

func (c SortBy) Limit(Offset int64, Count int64) SortLimit {
	return SortLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortBy) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cs: append(c.cs, Pattern...)}
}

func (c SortBy) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortBy) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortBy) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortBy) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination)}
}

func (c SortBy) Build() []string {
	return c.cs
}

type SortGet struct {
	cs []string
}

func (c SortGet) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortGet) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortGet) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortGet) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination)}
}

func (c SortGet) Get(Get ...string) SortGet {
	return SortGet{cs: append(c.cs, Get...)}
}

func (c SortGet) Build() []string {
	return c.cs
}

type SortKey struct {
	cs []string
}

func (c SortKey) By(Pattern string) SortBy {
	return SortBy{cs: append(c.cs, "BY", Pattern)}
}

func (c SortKey) Limit(Offset int64, Count int64) SortLimit {
	return SortLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortKey) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cs: append(c.cs, Pattern...)}
}

func (c SortKey) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortKey) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortKey) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortKey) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination)}
}

func (c SortKey) Build() []string {
	return c.cs
}

type SortLimit struct {
	cs []string
}

func (c SortLimit) Get(Pattern ...string) SortGet {
	c.cs = append(c.cs, "GET")
	return SortGet{cs: append(c.cs, Pattern...)}
}

func (c SortLimit) Asc() SortOrderAsc {
	return SortOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortLimit) Desc() SortOrderDesc {
	return SortOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortLimit) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortLimit) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination)}
}

func (c SortLimit) Build() []string {
	return c.cs
}

type SortOrderAsc struct {
	cs []string
}

func (c SortOrderAsc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortOrderAsc) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination)}
}

func (c SortOrderAsc) Build() []string {
	return c.cs
}

type SortOrderDesc struct {
	cs []string
}

func (c SortOrderDesc) Alpha() SortSortingAlpha {
	return SortSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortOrderDesc) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination)}
}

func (c SortOrderDesc) Build() []string {
	return c.cs
}

type SortRo struct {
	cs []string
}

func (c SortRo) Key(Key string) SortRoKey {
	return SortRoKey{cs: append(c.cs, Key)}
}

func (b *Builder) SortRo() (c SortRo) {
	c.cs = append(b.get(), "SORT_RO")
	return
}

type SortRoBy struct {
	cs []string
}

func (c SortRoBy) Limit(Offset int64, Count int64) SortRoLimit {
	return SortRoLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortRoBy) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cs: append(c.cs, Pattern...)}
}

func (c SortRoBy) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortRoBy) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortRoBy) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortRoBy) Build() []string {
	return c.cs
}

type SortRoGet struct {
	cs []string
}

func (c SortRoGet) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortRoGet) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortRoGet) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortRoGet) Get(Get ...string) SortRoGet {
	return SortRoGet{cs: append(c.cs, Get...)}
}

func (c SortRoGet) Build() []string {
	return c.cs
}

type SortRoKey struct {
	cs []string
}

func (c SortRoKey) By(Pattern string) SortRoBy {
	return SortRoBy{cs: append(c.cs, "BY", Pattern)}
}

func (c SortRoKey) Limit(Offset int64, Count int64) SortRoLimit {
	return SortRoLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c SortRoKey) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cs: append(c.cs, Pattern...)}
}

func (c SortRoKey) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortRoKey) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortRoKey) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortRoKey) Build() []string {
	return c.cs
}

type SortRoLimit struct {
	cs []string
}

func (c SortRoLimit) Get(Pattern ...string) SortRoGet {
	c.cs = append(c.cs, "GET")
	return SortRoGet{cs: append(c.cs, Pattern...)}
}

func (c SortRoLimit) Asc() SortRoOrderAsc {
	return SortRoOrderAsc{cs: append(c.cs, "ASC")}
}

func (c SortRoLimit) Desc() SortRoOrderDesc {
	return SortRoOrderDesc{cs: append(c.cs, "DESC")}
}

func (c SortRoLimit) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortRoLimit) Build() []string {
	return c.cs
}

type SortRoOrderAsc struct {
	cs []string
}

func (c SortRoOrderAsc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortRoOrderAsc) Build() []string {
	return c.cs
}

type SortRoOrderDesc struct {
	cs []string
}

func (c SortRoOrderDesc) Alpha() SortRoSortingAlpha {
	return SortRoSortingAlpha{cs: append(c.cs, "ALPHA")}
}

func (c SortRoOrderDesc) Build() []string {
	return c.cs
}

type SortRoSortingAlpha struct {
	cs []string
}

func (c SortRoSortingAlpha) Build() []string {
	return c.cs
}

type SortSortingAlpha struct {
	cs []string
}

func (c SortSortingAlpha) Store(Destination string) SortStore {
	return SortStore{cs: append(c.cs, "STORE", Destination)}
}

func (c SortSortingAlpha) Build() []string {
	return c.cs
}

type SortStore struct {
	cs []string
}

func (c SortStore) Build() []string {
	return c.cs
}

type Spop struct {
	cs []string
}

func (c Spop) Key(Key string) SpopKey {
	return SpopKey{cs: append(c.cs, Key)}
}

func (b *Builder) Spop() (c Spop) {
	c.cs = append(b.get(), "SPOP")
	return
}

type SpopCount struct {
	cs []string
}

func (c SpopCount) Build() []string {
	return c.cs
}

type SpopKey struct {
	cs []string
}

func (c SpopKey) Count(Count int64) SpopCount {
	return SpopCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SpopKey) Build() []string {
	return c.cs
}

type Srandmember struct {
	cs []string
}

func (c Srandmember) Key(Key string) SrandmemberKey {
	return SrandmemberKey{cs: append(c.cs, Key)}
}

func (b *Builder) Srandmember() (c Srandmember) {
	c.cs = append(b.get(), "SRANDMEMBER")
	return
}

type SrandmemberCount struct {
	cs []string
}

func (c SrandmemberCount) Build() []string {
	return c.cs
}

type SrandmemberKey struct {
	cs []string
}

func (c SrandmemberKey) Count(Count int64) SrandmemberCount {
	return SrandmemberCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c SrandmemberKey) Build() []string {
	return c.cs
}

type Srem struct {
	cs []string
}

func (c Srem) Key(Key string) SremKey {
	return SremKey{cs: append(c.cs, Key)}
}

func (b *Builder) Srem() (c Srem) {
	c.cs = append(b.get(), "SREM")
	return
}

type SremKey struct {
	cs []string
}

func (c SremKey) Member(Member ...string) SremMember {
	return SremMember{cs: append(c.cs, Member...)}
}

type SremMember struct {
	cs []string
}

func (c SremMember) Member(Member ...string) SremMember {
	return SremMember{cs: append(c.cs, Member...)}
}

func (c SremMember) Build() []string {
	return c.cs
}

type Sscan struct {
	cs []string
}

func (c Sscan) Key(Key string) SscanKey {
	return SscanKey{cs: append(c.cs, Key)}
}

func (b *Builder) Sscan() (c Sscan) {
	c.cs = append(b.get(), "SSCAN")
	return
}

type SscanCount struct {
	cs []string
}

func (c SscanCount) Build() []string {
	return c.cs
}

type SscanCursor struct {
	cs []string
}

func (c SscanCursor) Match(Pattern string) SscanMatch {
	return SscanMatch{cs: append(c.cs, "MATCH", Pattern)}
}

func (c SscanCursor) Count(Count int64) SscanCount {
	return SscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SscanCursor) Build() []string {
	return c.cs
}

type SscanKey struct {
	cs []string
}

func (c SscanKey) Cursor(Cursor int64) SscanCursor {
	return SscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type SscanMatch struct {
	cs []string
}

func (c SscanMatch) Count(Count int64) SscanCount {
	return SscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c SscanMatch) Build() []string {
	return c.cs
}

type Stralgo struct {
	cs []string
}

func (c Stralgo) Lcs() StralgoAlgorithmLcs {
	return StralgoAlgorithmLcs{cs: append(c.cs, "LCS")}
}

func (b *Builder) Stralgo() (c Stralgo) {
	c.cs = append(b.get(), "STRALGO")
	return
}

type StralgoAlgoSpecificArgument struct {
	cs []string
}

func (c StralgoAlgoSpecificArgument) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cs: append(c.cs, AlgoSpecificArgument...)}
}

func (c StralgoAlgoSpecificArgument) Build() []string {
	return c.cs
}

type StralgoAlgorithmLcs struct {
	cs []string
}

func (c StralgoAlgorithmLcs) AlgoSpecificArgument(AlgoSpecificArgument ...string) StralgoAlgoSpecificArgument {
	return StralgoAlgoSpecificArgument{cs: append(c.cs, AlgoSpecificArgument...)}
}

type Strlen struct {
	cs []string
}

func (c Strlen) Key(Key string) StrlenKey {
	return StrlenKey{cs: append(c.cs, Key)}
}

func (b *Builder) Strlen() (c Strlen) {
	c.cs = append(b.get(), "STRLEN")
	return
}

type StrlenKey struct {
	cs []string
}

func (c StrlenKey) Build() []string {
	return c.cs
}

type Subscribe struct {
	cs []string
}

func (c Subscribe) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cs: append(c.cs, Channel...)}
}

func (b *Builder) Subscribe() (c Subscribe) {
	c.cs = append(b.get(), "SUBSCRIBE")
	return
}

type SubscribeChannel struct {
	cs []string
}

func (c SubscribeChannel) Channel(Channel ...string) SubscribeChannel {
	return SubscribeChannel{cs: append(c.cs, Channel...)}
}

func (c SubscribeChannel) Build() []string {
	return c.cs
}

type Sunion struct {
	cs []string
}

func (c Sunion) Key(Key ...string) SunionKey {
	return SunionKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Sunion() (c Sunion) {
	c.cs = append(b.get(), "SUNION")
	return
}

type SunionKey struct {
	cs []string
}

func (c SunionKey) Key(Key ...string) SunionKey {
	return SunionKey{cs: append(c.cs, Key...)}
}

func (c SunionKey) Build() []string {
	return c.cs
}

type Sunionstore struct {
	cs []string
}

func (c Sunionstore) Destination(Destination string) SunionstoreDestination {
	return SunionstoreDestination{cs: append(c.cs, Destination)}
}

func (b *Builder) Sunionstore() (c Sunionstore) {
	c.cs = append(b.get(), "SUNIONSTORE")
	return
}

type SunionstoreDestination struct {
	cs []string
}

func (c SunionstoreDestination) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cs: append(c.cs, Key...)}
}

type SunionstoreKey struct {
	cs []string
}

func (c SunionstoreKey) Key(Key ...string) SunionstoreKey {
	return SunionstoreKey{cs: append(c.cs, Key...)}
}

func (c SunionstoreKey) Build() []string {
	return c.cs
}

type Swapdb struct {
	cs []string
}

func (c Swapdb) Index1(Index1 int64) SwapdbIndex1 {
	return SwapdbIndex1{cs: append(c.cs, strconv.FormatInt(Index1, 10))}
}

func (b *Builder) Swapdb() (c Swapdb) {
	c.cs = append(b.get(), "SWAPDB")
	return
}

type SwapdbIndex1 struct {
	cs []string
}

func (c SwapdbIndex1) Index2(Index2 int64) SwapdbIndex2 {
	return SwapdbIndex2{cs: append(c.cs, strconv.FormatInt(Index2, 10))}
}

type SwapdbIndex2 struct {
	cs []string
}

func (c SwapdbIndex2) Build() []string {
	return c.cs
}

type Sync struct {
	cs []string
}

func (c Sync) Build() []string {
	return c.cs
}

func (b *Builder) Sync() (c Sync) {
	c.cs = append(b.get(), "SYNC")
	return
}

type Time struct {
	cs []string
}

func (c Time) Build() []string {
	return c.cs
}

func (b *Builder) Time() (c Time) {
	c.cs = append(b.get(), "TIME")
	return
}

type Touch struct {
	cs []string
}

func (c Touch) Key(Key ...string) TouchKey {
	return TouchKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Touch() (c Touch) {
	c.cs = append(b.get(), "TOUCH")
	return
}

type TouchKey struct {
	cs []string
}

func (c TouchKey) Key(Key ...string) TouchKey {
	return TouchKey{cs: append(c.cs, Key...)}
}

func (c TouchKey) Build() []string {
	return c.cs
}

type Ttl struct {
	cs []string
}

func (c Ttl) Key(Key string) TtlKey {
	return TtlKey{cs: append(c.cs, Key)}
}

func (b *Builder) Ttl() (c Ttl) {
	c.cs = append(b.get(), "TTL")
	return
}

type TtlKey struct {
	cs []string
}

func (c TtlKey) Build() []string {
	return c.cs
}

type Type struct {
	cs []string
}

func (c Type) Key(Key string) TypeKey {
	return TypeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Type() (c Type) {
	c.cs = append(b.get(), "TYPE")
	return
}

type TypeKey struct {
	cs []string
}

func (c TypeKey) Build() []string {
	return c.cs
}

type Unlink struct {
	cs []string
}

func (c Unlink) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Unlink() (c Unlink) {
	c.cs = append(b.get(), "UNLINK")
	return
}

type UnlinkKey struct {
	cs []string
}

func (c UnlinkKey) Key(Key ...string) UnlinkKey {
	return UnlinkKey{cs: append(c.cs, Key...)}
}

func (c UnlinkKey) Build() []string {
	return c.cs
}

type Unsubscribe struct {
	cs []string
}

func (c Unsubscribe) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cs: append(c.cs, Channel...)}
}

func (c Unsubscribe) Build() []string {
	return c.cs
}

func (b *Builder) Unsubscribe() (c Unsubscribe) {
	c.cs = append(b.get(), "UNSUBSCRIBE")
	return
}

type UnsubscribeChannel struct {
	cs []string
}

func (c UnsubscribeChannel) Channel(Channel ...string) UnsubscribeChannel {
	return UnsubscribeChannel{cs: append(c.cs, Channel...)}
}

func (c UnsubscribeChannel) Build() []string {
	return c.cs
}

type Unwatch struct {
	cs []string
}

func (c Unwatch) Build() []string {
	return c.cs
}

func (b *Builder) Unwatch() (c Unwatch) {
	c.cs = append(b.get(), "UNWATCH")
	return
}

type Wait struct {
	cs []string
}

func (c Wait) Numreplicas(Numreplicas int64) WaitNumreplicas {
	return WaitNumreplicas{cs: append(c.cs, strconv.FormatInt(Numreplicas, 10))}
}

func (b *Builder) Wait() (c Wait) {
	c.cs = append(b.get(), "WAIT")
	return
}

type WaitNumreplicas struct {
	cs []string
}

func (c WaitNumreplicas) Timeout(Timeout int64) WaitTimeout {
	return WaitTimeout{cs: append(c.cs, strconv.FormatInt(Timeout, 10))}
}

type WaitTimeout struct {
	cs []string
}

func (c WaitTimeout) Build() []string {
	return c.cs
}

type Watch struct {
	cs []string
}

func (c Watch) Key(Key ...string) WatchKey {
	return WatchKey{cs: append(c.cs, Key...)}
}

func (b *Builder) Watch() (c Watch) {
	c.cs = append(b.get(), "WATCH")
	return
}

type WatchKey struct {
	cs []string
}

func (c WatchKey) Key(Key ...string) WatchKey {
	return WatchKey{cs: append(c.cs, Key...)}
}

func (c WatchKey) Build() []string {
	return c.cs
}

type Xack struct {
	cs []string
}

func (c Xack) Key(Key string) XackKey {
	return XackKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xack() (c Xack) {
	c.cs = append(b.get(), "XACK")
	return
}

type XackGroup struct {
	cs []string
}

func (c XackGroup) Id(Id ...string) XackId {
	return XackId{cs: append(c.cs, Id...)}
}

type XackId struct {
	cs []string
}

func (c XackId) Id(Id ...string) XackId {
	return XackId{cs: append(c.cs, Id...)}
}

func (c XackId) Build() []string {
	return c.cs
}

type XackKey struct {
	cs []string
}

func (c XackKey) Group(Group string) XackGroup {
	return XackGroup{cs: append(c.cs, Group)}
}

type Xadd struct {
	cs []string
}

func (c Xadd) Key(Key string) XaddKey {
	return XaddKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xadd() (c Xadd) {
	c.cs = append(b.get(), "XADD")
	return
}

type XaddFieldValue struct {
	cs []string
}

func (c XaddFieldValue) FieldValue(Field string, Value string) XaddFieldValue {
	return XaddFieldValue{cs: append(c.cs, Field, Value)}
}

func (c XaddFieldValue) Build() []string {
	return c.cs
}

type XaddId struct {
	cs []string
}

func (c XaddId) FieldValue() XaddFieldValue {
	return XaddFieldValue{cs: append(c.cs)}
}

type XaddKey struct {
	cs []string
}

func (c XaddKey) Nomkstream() XaddNomkstream {
	return XaddNomkstream{cs: append(c.cs, "NOMKSTREAM")}
}

func (c XaddKey) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN")}
}

func (c XaddKey) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cs: append(c.cs, "MINID")}
}

func (c XaddKey) Wildcard() XaddWildcard {
	return XaddWildcard{cs: append(c.cs, "*")}
}

func (c XaddKey) Id() XaddId {
	return XaddId{cs: append(c.cs, "ID")}
}

type XaddNomkstream struct {
	cs []string
}

func (c XaddNomkstream) Maxlen() XaddTrimStrategyMaxlen {
	return XaddTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN")}
}

func (c XaddNomkstream) Minid() XaddTrimStrategyMinid {
	return XaddTrimStrategyMinid{cs: append(c.cs, "MINID")}
}

func (c XaddNomkstream) Wildcard() XaddWildcard {
	return XaddWildcard{cs: append(c.cs, "*")}
}

func (c XaddNomkstream) Id() XaddId {
	return XaddId{cs: append(c.cs, "ID")}
}

type XaddTrimLimit struct {
	cs []string
}

func (c XaddTrimLimit) Wildcard() XaddWildcard {
	return XaddWildcard{cs: append(c.cs, "*")}
}

func (c XaddTrimLimit) Id() XaddId {
	return XaddId{cs: append(c.cs, "ID")}
}

type XaddTrimOperatorAlmost struct {
	cs []string
}

func (c XaddTrimOperatorAlmost) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold)}
}

type XaddTrimOperatorExact struct {
	cs []string
}

func (c XaddTrimOperatorExact) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold)}
}

type XaddTrimStrategyMaxlen struct {
	cs []string
}

func (c XaddTrimStrategyMaxlen) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cs: append(c.cs, "=")}
}

func (c XaddTrimStrategyMaxlen) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cs: append(c.cs, "~")}
}

func (c XaddTrimStrategyMaxlen) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold)}
}

type XaddTrimStrategyMinid struct {
	cs []string
}

func (c XaddTrimStrategyMinid) Exact() XaddTrimOperatorExact {
	return XaddTrimOperatorExact{cs: append(c.cs, "=")}
}

func (c XaddTrimStrategyMinid) Almost() XaddTrimOperatorAlmost {
	return XaddTrimOperatorAlmost{cs: append(c.cs, "~")}
}

func (c XaddTrimStrategyMinid) Threshold(Threshold string) XaddTrimThreshold {
	return XaddTrimThreshold{cs: append(c.cs, Threshold)}
}

type XaddTrimThreshold struct {
	cs []string
}

func (c XaddTrimThreshold) Limit(Count int64) XaddTrimLimit {
	return XaddTrimLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XaddTrimThreshold) Wildcard() XaddWildcard {
	return XaddWildcard{cs: append(c.cs, "*")}
}

func (c XaddTrimThreshold) Id() XaddId {
	return XaddId{cs: append(c.cs, "ID")}
}

type XaddWildcard struct {
	cs []string
}

func (c XaddWildcard) FieldValue() XaddFieldValue {
	return XaddFieldValue{cs: append(c.cs)}
}

type Xautoclaim struct {
	cs []string
}

func (c Xautoclaim) Key(Key string) XautoclaimKey {
	return XautoclaimKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xautoclaim() (c Xautoclaim) {
	c.cs = append(b.get(), "XAUTOCLAIM")
	return
}

type XautoclaimConsumer struct {
	cs []string
}

func (c XautoclaimConsumer) MinIdleTime(MinIdleTime string) XautoclaimMinIdleTime {
	return XautoclaimMinIdleTime{cs: append(c.cs, MinIdleTime)}
}

type XautoclaimCount struct {
	cs []string
}

func (c XautoclaimCount) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cs: append(c.cs, "JUSTID")}
}

func (c XautoclaimCount) Build() []string {
	return c.cs
}

type XautoclaimGroup struct {
	cs []string
}

func (c XautoclaimGroup) Consumer(Consumer string) XautoclaimConsumer {
	return XautoclaimConsumer{cs: append(c.cs, Consumer)}
}

type XautoclaimJustidJustid struct {
	cs []string
}

func (c XautoclaimJustidJustid) Build() []string {
	return c.cs
}

type XautoclaimKey struct {
	cs []string
}

func (c XautoclaimKey) Group(Group string) XautoclaimGroup {
	return XautoclaimGroup{cs: append(c.cs, Group)}
}

type XautoclaimMinIdleTime struct {
	cs []string
}

func (c XautoclaimMinIdleTime) Start(Start string) XautoclaimStart {
	return XautoclaimStart{cs: append(c.cs, Start)}
}

type XautoclaimStart struct {
	cs []string
}

func (c XautoclaimStart) Count(Count int64) XautoclaimCount {
	return XautoclaimCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XautoclaimStart) Justid() XautoclaimJustidJustid {
	return XautoclaimJustidJustid{cs: append(c.cs, "JUSTID")}
}

func (c XautoclaimStart) Build() []string {
	return c.cs
}

type Xclaim struct {
	cs []string
}

func (c Xclaim) Key(Key string) XclaimKey {
	return XclaimKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xclaim() (c Xclaim) {
	c.cs = append(b.get(), "XCLAIM")
	return
}

type XclaimConsumer struct {
	cs []string
}

func (c XclaimConsumer) MinIdleTime(MinIdleTime string) XclaimMinIdleTime {
	return XclaimMinIdleTime{cs: append(c.cs, MinIdleTime)}
}

type XclaimForceForce struct {
	cs []string
}

func (c XclaimForceForce) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID")}
}

func (c XclaimForceForce) Build() []string {
	return c.cs
}

type XclaimGroup struct {
	cs []string
}

func (c XclaimGroup) Consumer(Consumer string) XclaimConsumer {
	return XclaimConsumer{cs: append(c.cs, Consumer)}
}

type XclaimId struct {
	cs []string
}

func (c XclaimId) Idle(Ms int64) XclaimIdle {
	return XclaimIdle{cs: append(c.cs, "IDLE", strconv.FormatInt(Ms, 10))}
}

func (c XclaimId) Time(MsUnixTime int64) XclaimTime {
	return XclaimTime{cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10))}
}

func (c XclaimId) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c XclaimId) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE")}
}

func (c XclaimId) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID")}
}

func (c XclaimId) Id(Id ...string) XclaimId {
	return XclaimId{cs: append(c.cs, Id...)}
}

func (c XclaimId) Build() []string {
	return c.cs
}

type XclaimIdle struct {
	cs []string
}

func (c XclaimIdle) Time(MsUnixTime int64) XclaimTime {
	return XclaimTime{cs: append(c.cs, "TIME", strconv.FormatInt(MsUnixTime, 10))}
}

func (c XclaimIdle) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c XclaimIdle) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE")}
}

func (c XclaimIdle) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID")}
}

func (c XclaimIdle) Build() []string {
	return c.cs
}

type XclaimJustidJustid struct {
	cs []string
}

func (c XclaimJustidJustid) Build() []string {
	return c.cs
}

type XclaimKey struct {
	cs []string
}

func (c XclaimKey) Group(Group string) XclaimGroup {
	return XclaimGroup{cs: append(c.cs, Group)}
}

type XclaimMinIdleTime struct {
	cs []string
}

func (c XclaimMinIdleTime) Id(Id ...string) XclaimId {
	return XclaimId{cs: append(c.cs, Id...)}
}

type XclaimRetrycount struct {
	cs []string
}

func (c XclaimRetrycount) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE")}
}

func (c XclaimRetrycount) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID")}
}

func (c XclaimRetrycount) Build() []string {
	return c.cs
}

type XclaimTime struct {
	cs []string
}

func (c XclaimTime) Retrycount(Count int64) XclaimRetrycount {
	return XclaimRetrycount{cs: append(c.cs, "RETRYCOUNT", strconv.FormatInt(Count, 10))}
}

func (c XclaimTime) Force() XclaimForceForce {
	return XclaimForceForce{cs: append(c.cs, "FORCE")}
}

func (c XclaimTime) Justid() XclaimJustidJustid {
	return XclaimJustidJustid{cs: append(c.cs, "JUSTID")}
}

func (c XclaimTime) Build() []string {
	return c.cs
}

type Xdel struct {
	cs []string
}

func (c Xdel) Key(Key string) XdelKey {
	return XdelKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xdel() (c Xdel) {
	c.cs = append(b.get(), "XDEL")
	return
}

type XdelId struct {
	cs []string
}

func (c XdelId) Id(Id ...string) XdelId {
	return XdelId{cs: append(c.cs, Id...)}
}

func (c XdelId) Build() []string {
	return c.cs
}

type XdelKey struct {
	cs []string
}

func (c XdelKey) Id(Id ...string) XdelId {
	return XdelId{cs: append(c.cs, Id...)}
}

type Xgroup struct {
	cs []string
}

func (c Xgroup) Create(Key string, Groupname string) XgroupCreateCreate {
	return XgroupCreateCreate{cs: append(c.cs, "CREATE", Key, Groupname)}
}

func (c Xgroup) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c Xgroup) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c Xgroup) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c Xgroup) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (b *Builder) Xgroup() (c Xgroup) {
	c.cs = append(b.get(), "XGROUP")
	return
}

type XgroupCreateCreate struct {
	cs []string
}

func (c XgroupCreateCreate) Id() XgroupCreateIdId {
	return XgroupCreateIdId{cs: append(c.cs, "ID")}
}

func (c XgroupCreateCreate) Lastid() XgroupCreateIdLastID {
	return XgroupCreateIdLastID{cs: append(c.cs, "$")}
}

type XgroupCreateIdId struct {
	cs []string
}

func (c XgroupCreateIdId) Mkstream() XgroupCreateMkstream {
	return XgroupCreateMkstream{cs: append(c.cs, "MKSTREAM")}
}

func (c XgroupCreateIdId) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c XgroupCreateIdId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupCreateIdId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateIdId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type XgroupCreateIdLastID struct {
	cs []string
}

func (c XgroupCreateIdLastID) Mkstream() XgroupCreateMkstream {
	return XgroupCreateMkstream{cs: append(c.cs, "MKSTREAM")}
}

func (c XgroupCreateIdLastID) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c XgroupCreateIdLastID) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupCreateIdLastID) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateIdLastID) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type XgroupCreateMkstream struct {
	cs []string
}

func (c XgroupCreateMkstream) Setid(Key string, Groupname string) XgroupSetidSetid {
	return XgroupSetidSetid{cs: append(c.cs, "SETID", Key, Groupname)}
}

func (c XgroupCreateMkstream) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupCreateMkstream) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateMkstream) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

type XgroupCreateconsumer struct {
	cs []string
}

func (c XgroupCreateconsumer) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupCreateconsumer) Build() []string {
	return c.cs
}

type XgroupDelconsumer struct {
	cs []string
}

func (c XgroupDelconsumer) Build() []string {
	return c.cs
}

type XgroupDestroy struct {
	cs []string
}

func (c XgroupDestroy) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupDestroy) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupDestroy) Build() []string {
	return c.cs
}

type XgroupSetidIdId struct {
	cs []string
}

func (c XgroupSetidIdId) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupSetidIdId) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdId) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdId) Build() []string {
	return c.cs
}

type XgroupSetidIdLastID struct {
	cs []string
}

func (c XgroupSetidIdLastID) Destroy(Key string, Groupname string) XgroupDestroy {
	return XgroupDestroy{cs: append(c.cs, "DESTROY", Key, Groupname)}
}

func (c XgroupSetidIdLastID) Createconsumer(Key string, Groupname string, Consumername string) XgroupCreateconsumer {
	return XgroupCreateconsumer{cs: append(c.cs, "CREATECONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdLastID) Delconsumer(Key string, Groupname string, Consumername string) XgroupDelconsumer {
	return XgroupDelconsumer{cs: append(c.cs, "DELCONSUMER", Key, Groupname, Consumername)}
}

func (c XgroupSetidIdLastID) Build() []string {
	return c.cs
}

type XgroupSetidSetid struct {
	cs []string
}

func (c XgroupSetidSetid) Id() XgroupSetidIdId {
	return XgroupSetidIdId{cs: append(c.cs, "ID")}
}

func (c XgroupSetidSetid) Lastid() XgroupSetidIdLastID {
	return XgroupSetidIdLastID{cs: append(c.cs, "$")}
}

type Xinfo struct {
	cs []string
}

func (c Xinfo) Consumers(Key string, Groupname string) XinfoConsumers {
	return XinfoConsumers{cs: append(c.cs, "CONSUMERS", Key, Groupname)}
}

func (c Xinfo) Groups(Key string) XinfoGroups {
	return XinfoGroups{cs: append(c.cs, "GROUPS", Key)}
}

func (c Xinfo) Stream(Key string) XinfoStream {
	return XinfoStream{cs: append(c.cs, "STREAM", Key)}
}

func (c Xinfo) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP")}
}

func (c Xinfo) Build() []string {
	return c.cs
}

func (b *Builder) Xinfo() (c Xinfo) {
	c.cs = append(b.get(), "XINFO")
	return
}

type XinfoConsumers struct {
	cs []string
}

func (c XinfoConsumers) Groups(Key string) XinfoGroups {
	return XinfoGroups{cs: append(c.cs, "GROUPS", Key)}
}

func (c XinfoConsumers) Stream(Key string) XinfoStream {
	return XinfoStream{cs: append(c.cs, "STREAM", Key)}
}

func (c XinfoConsumers) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP")}
}

func (c XinfoConsumers) Build() []string {
	return c.cs
}

type XinfoGroups struct {
	cs []string
}

func (c XinfoGroups) Stream(Key string) XinfoStream {
	return XinfoStream{cs: append(c.cs, "STREAM", Key)}
}

func (c XinfoGroups) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP")}
}

func (c XinfoGroups) Build() []string {
	return c.cs
}

type XinfoHelpHelp struct {
	cs []string
}

func (c XinfoHelpHelp) Build() []string {
	return c.cs
}

type XinfoStream struct {
	cs []string
}

func (c XinfoStream) Help() XinfoHelpHelp {
	return XinfoHelpHelp{cs: append(c.cs, "HELP")}
}

func (c XinfoStream) Build() []string {
	return c.cs
}

type Xlen struct {
	cs []string
}

func (c Xlen) Key(Key string) XlenKey {
	return XlenKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xlen() (c Xlen) {
	c.cs = append(b.get(), "XLEN")
	return
}

type XlenKey struct {
	cs []string
}

func (c XlenKey) Build() []string {
	return c.cs
}

type Xpending struct {
	cs []string
}

func (c Xpending) Key(Key string) XpendingKey {
	return XpendingKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xpending() (c Xpending) {
	c.cs = append(b.get(), "XPENDING")
	return
}

type XpendingFiltersConsumer struct {
	cs []string
}

func (c XpendingFiltersConsumer) Build() []string {
	return c.cs
}

type XpendingFiltersCount struct {
	cs []string
}

func (c XpendingFiltersCount) Consumer(Consumer string) XpendingFiltersConsumer {
	return XpendingFiltersConsumer{cs: append(c.cs, Consumer)}
}

func (c XpendingFiltersCount) Build() []string {
	return c.cs
}

type XpendingFiltersEnd struct {
	cs []string
}

func (c XpendingFiltersEnd) Count(Count int64) XpendingFiltersCount {
	return XpendingFiltersCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type XpendingFiltersIdle struct {
	cs []string
}

func (c XpendingFiltersIdle) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cs: append(c.cs, Start)}
}

type XpendingFiltersStart struct {
	cs []string
}

func (c XpendingFiltersStart) End(End string) XpendingFiltersEnd {
	return XpendingFiltersEnd{cs: append(c.cs, End)}
}

type XpendingGroup struct {
	cs []string
}

func (c XpendingGroup) Idle(MinIdleTime int64) XpendingFiltersIdle {
	return XpendingFiltersIdle{cs: append(c.cs, "IDLE", strconv.FormatInt(MinIdleTime, 10))}
}

func (c XpendingGroup) Start(Start string) XpendingFiltersStart {
	return XpendingFiltersStart{cs: append(c.cs, Start)}
}

type XpendingKey struct {
	cs []string
}

func (c XpendingKey) Group(Group string) XpendingGroup {
	return XpendingGroup{cs: append(c.cs, Group)}
}

type Xrange struct {
	cs []string
}

func (c Xrange) Key(Key string) XrangeKey {
	return XrangeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xrange() (c Xrange) {
	c.cs = append(b.get(), "XRANGE")
	return
}

type XrangeCount struct {
	cs []string
}

func (c XrangeCount) Build() []string {
	return c.cs
}

type XrangeEnd struct {
	cs []string
}

func (c XrangeEnd) Count(Count int64) XrangeCount {
	return XrangeCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrangeEnd) Build() []string {
	return c.cs
}

type XrangeKey struct {
	cs []string
}

func (c XrangeKey) Start(Start string) XrangeStart {
	return XrangeStart{cs: append(c.cs, Start)}
}

type XrangeStart struct {
	cs []string
}

func (c XrangeStart) End(End string) XrangeEnd {
	return XrangeEnd{cs: append(c.cs, End)}
}

type Xread struct {
	cs []string
}

func (c Xread) Count(Count int64) XreadCount {
	return XreadCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c Xread) Block(Milliseconds int64) XreadBlock {
	return XreadBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c Xread) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cs: append(c.cs, "STREAMS")}
}

func (b *Builder) Xread() (c Xread) {
	c.cs = append(b.get(), "XREAD")
	return
}

type XreadBlock struct {
	cs []string
}

func (c XreadBlock) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cs: append(c.cs, "STREAMS")}
}

type XreadCount struct {
	cs []string
}

func (c XreadCount) Block(Milliseconds int64) XreadBlock {
	return XreadBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadCount) Streams() XreadStreamsStreams {
	return XreadStreamsStreams{cs: append(c.cs, "STREAMS")}
}

type XreadId struct {
	cs []string
}

func (c XreadId) Id(Id ...string) XreadId {
	return XreadId{cs: append(c.cs, Id...)}
}

func (c XreadId) Build() []string {
	return c.cs
}

type XreadKey struct {
	cs []string
}

func (c XreadKey) Id(Id ...string) XreadId {
	return XreadId{cs: append(c.cs, Id...)}
}

func (c XreadKey) Key(Key ...string) XreadKey {
	return XreadKey{cs: append(c.cs, Key...)}
}

type XreadStreamsStreams struct {
	cs []string
}

func (c XreadStreamsStreams) Key(Key ...string) XreadKey {
	return XreadKey{cs: append(c.cs, Key...)}
}

type Xreadgroup struct {
	cs []string
}

func (c Xreadgroup) Group(Group string, Consumer string) XreadgroupGroup {
	return XreadgroupGroup{cs: append(c.cs, "GROUP", Group, Consumer)}
}

func (b *Builder) Xreadgroup() (c Xreadgroup) {
	c.cs = append(b.get(), "XREADGROUP")
	return
}

type XreadgroupBlock struct {
	cs []string
}

func (c XreadgroupBlock) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cs: append(c.cs, "NOACK")}
}

func (c XreadgroupBlock) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS")}
}

type XreadgroupCount struct {
	cs []string
}

func (c XreadgroupCount) Block(Milliseconds int64) XreadgroupBlock {
	return XreadgroupBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadgroupCount) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cs: append(c.cs, "NOACK")}
}

func (c XreadgroupCount) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS")}
}

type XreadgroupGroup struct {
	cs []string
}

func (c XreadgroupGroup) Count(Count int64) XreadgroupCount {
	return XreadgroupCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XreadgroupGroup) Block(Milliseconds int64) XreadgroupBlock {
	return XreadgroupBlock{cs: append(c.cs, "BLOCK", strconv.FormatInt(Milliseconds, 10))}
}

func (c XreadgroupGroup) Noack() XreadgroupNoackNoack {
	return XreadgroupNoackNoack{cs: append(c.cs, "NOACK")}
}

func (c XreadgroupGroup) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS")}
}

type XreadgroupId struct {
	cs []string
}

func (c XreadgroupId) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cs: append(c.cs, Id...)}
}

func (c XreadgroupId) Build() []string {
	return c.cs
}

type XreadgroupKey struct {
	cs []string
}

func (c XreadgroupKey) Id(Id ...string) XreadgroupId {
	return XreadgroupId{cs: append(c.cs, Id...)}
}

func (c XreadgroupKey) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cs: append(c.cs, Key...)}
}

type XreadgroupNoackNoack struct {
	cs []string
}

func (c XreadgroupNoackNoack) Streams() XreadgroupStreamsStreams {
	return XreadgroupStreamsStreams{cs: append(c.cs, "STREAMS")}
}

type XreadgroupStreamsStreams struct {
	cs []string
}

func (c XreadgroupStreamsStreams) Key(Key ...string) XreadgroupKey {
	return XreadgroupKey{cs: append(c.cs, Key...)}
}

type Xrevrange struct {
	cs []string
}

func (c Xrevrange) Key(Key string) XrevrangeKey {
	return XrevrangeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xrevrange() (c Xrevrange) {
	c.cs = append(b.get(), "XREVRANGE")
	return
}

type XrevrangeCount struct {
	cs []string
}

func (c XrevrangeCount) Build() []string {
	return c.cs
}

type XrevrangeEnd struct {
	cs []string
}

func (c XrevrangeEnd) Start(Start string) XrevrangeStart {
	return XrevrangeStart{cs: append(c.cs, Start)}
}

type XrevrangeKey struct {
	cs []string
}

func (c XrevrangeKey) End(End string) XrevrangeEnd {
	return XrevrangeEnd{cs: append(c.cs, End)}
}

type XrevrangeStart struct {
	cs []string
}

func (c XrevrangeStart) Count(Count int64) XrevrangeCount {
	return XrevrangeCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c XrevrangeStart) Build() []string {
	return c.cs
}

type Xtrim struct {
	cs []string
}

func (c Xtrim) Key(Key string) XtrimKey {
	return XtrimKey{cs: append(c.cs, Key)}
}

func (b *Builder) Xtrim() (c Xtrim) {
	c.cs = append(b.get(), "XTRIM")
	return
}

type XtrimKey struct {
	cs []string
}

func (c XtrimKey) Maxlen() XtrimTrimStrategyMaxlen {
	return XtrimTrimStrategyMaxlen{cs: append(c.cs, "MAXLEN")}
}

func (c XtrimKey) Minid() XtrimTrimStrategyMinid {
	return XtrimTrimStrategyMinid{cs: append(c.cs, "MINID")}
}

type XtrimTrimLimit struct {
	cs []string
}

func (c XtrimTrimLimit) Build() []string {
	return c.cs
}

type XtrimTrimOperatorAlmost struct {
	cs []string
}

func (c XtrimTrimOperatorAlmost) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold)}
}

type XtrimTrimOperatorExact struct {
	cs []string
}

func (c XtrimTrimOperatorExact) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold)}
}

type XtrimTrimStrategyMaxlen struct {
	cs []string
}

func (c XtrimTrimStrategyMaxlen) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cs: append(c.cs, "=")}
}

func (c XtrimTrimStrategyMaxlen) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cs: append(c.cs, "~")}
}

func (c XtrimTrimStrategyMaxlen) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold)}
}

type XtrimTrimStrategyMinid struct {
	cs []string
}

func (c XtrimTrimStrategyMinid) Exact() XtrimTrimOperatorExact {
	return XtrimTrimOperatorExact{cs: append(c.cs, "=")}
}

func (c XtrimTrimStrategyMinid) Almost() XtrimTrimOperatorAlmost {
	return XtrimTrimOperatorAlmost{cs: append(c.cs, "~")}
}

func (c XtrimTrimStrategyMinid) Threshold(Threshold string) XtrimTrimThreshold {
	return XtrimTrimThreshold{cs: append(c.cs, Threshold)}
}

type XtrimTrimThreshold struct {
	cs []string
}

func (c XtrimTrimThreshold) Limit(Count int64) XtrimTrimLimit {
	return XtrimTrimLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Count, 10))}
}

func (c XtrimTrimThreshold) Build() []string {
	return c.cs
}

type Zadd struct {
	cs []string
}

func (c Zadd) Key(Key string) ZaddKey {
	return ZaddKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zadd() (c Zadd) {
	c.cs = append(b.get(), "ZADD")
	return
}

type ZaddChangeCh struct {
	cs []string
}

func (c ZaddChangeCh) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR")}
}

func (c ZaddChangeCh) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs)}
}

type ZaddComparisonGt struct {
	cs []string
}

func (c ZaddComparisonGt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH")}
}

func (c ZaddComparisonGt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR")}
}

func (c ZaddComparisonGt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs)}
}

type ZaddComparisonLt struct {
	cs []string
}

func (c ZaddComparisonLt) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH")}
}

func (c ZaddComparisonLt) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR")}
}

func (c ZaddComparisonLt) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs)}
}

type ZaddConditionNx struct {
	cs []string
}

func (c ZaddConditionNx) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cs: append(c.cs, "GT")}
}

func (c ZaddConditionNx) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cs: append(c.cs, "LT")}
}

func (c ZaddConditionNx) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH")}
}

func (c ZaddConditionNx) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR")}
}

func (c ZaddConditionNx) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs)}
}

type ZaddConditionXx struct {
	cs []string
}

func (c ZaddConditionXx) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cs: append(c.cs, "GT")}
}

func (c ZaddConditionXx) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cs: append(c.cs, "LT")}
}

func (c ZaddConditionXx) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH")}
}

func (c ZaddConditionXx) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR")}
}

func (c ZaddConditionXx) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs)}
}

type ZaddIncrementIncr struct {
	cs []string
}

func (c ZaddIncrementIncr) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs)}
}

type ZaddKey struct {
	cs []string
}

func (c ZaddKey) Nx() ZaddConditionNx {
	return ZaddConditionNx{cs: append(c.cs, "NX")}
}

func (c ZaddKey) Xx() ZaddConditionXx {
	return ZaddConditionXx{cs: append(c.cs, "XX")}
}

func (c ZaddKey) Gt() ZaddComparisonGt {
	return ZaddComparisonGt{cs: append(c.cs, "GT")}
}

func (c ZaddKey) Lt() ZaddComparisonLt {
	return ZaddComparisonLt{cs: append(c.cs, "LT")}
}

func (c ZaddKey) Ch() ZaddChangeCh {
	return ZaddChangeCh{cs: append(c.cs, "CH")}
}

func (c ZaddKey) Incr() ZaddIncrementIncr {
	return ZaddIncrementIncr{cs: append(c.cs, "INCR")}
}

func (c ZaddKey) ScoreMember() ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs)}
}

type ZaddScoreMember struct {
	cs []string
}

func (c ZaddScoreMember) ScoreMember(Score float64, Member string) ZaddScoreMember {
	return ZaddScoreMember{cs: append(c.cs, strconv.FormatFloat(Score, 'f', -1, 64), Member)}
}

func (c ZaddScoreMember) Build() []string {
	return c.cs
}

type Zcard struct {
	cs []string
}

func (c Zcard) Key(Key string) ZcardKey {
	return ZcardKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zcard() (c Zcard) {
	c.cs = append(b.get(), "ZCARD")
	return
}

type ZcardKey struct {
	cs []string
}

func (c ZcardKey) Build() []string {
	return c.cs
}

type Zcount struct {
	cs []string
}

func (c Zcount) Key(Key string) ZcountKey {
	return ZcountKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zcount() (c Zcount) {
	c.cs = append(b.get(), "ZCOUNT")
	return
}

type ZcountKey struct {
	cs []string
}

func (c ZcountKey) Min(Min float64) ZcountMin {
	return ZcountMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZcountMax struct {
	cs []string
}

func (c ZcountMax) Build() []string {
	return c.cs
}

type ZcountMin struct {
	cs []string
}

func (c ZcountMin) Max(Max float64) ZcountMax {
	return ZcountMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type Zdiff struct {
	cs []string
}

func (c Zdiff) Numkeys(Numkeys int64) ZdiffNumkeys {
	return ZdiffNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zdiff() (c Zdiff) {
	c.cs = append(b.get(), "ZDIFF")
	return
}

type ZdiffKey struct {
	cs []string
}

func (c ZdiffKey) Withscores() ZdiffWithscoresWithscores {
	return ZdiffWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZdiffKey) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cs: append(c.cs, Key...)}
}

func (c ZdiffKey) Build() []string {
	return c.cs
}

type ZdiffNumkeys struct {
	cs []string
}

func (c ZdiffNumkeys) Key(Key ...string) ZdiffKey {
	return ZdiffKey{cs: append(c.cs, Key...)}
}

type ZdiffWithscoresWithscores struct {
	cs []string
}

func (c ZdiffWithscoresWithscores) Build() []string {
	return c.cs
}

type Zdiffstore struct {
	cs []string
}

func (c Zdiffstore) Destination(Destination string) ZdiffstoreDestination {
	return ZdiffstoreDestination{cs: append(c.cs, Destination)}
}

func (b *Builder) Zdiffstore() (c Zdiffstore) {
	c.cs = append(b.get(), "ZDIFFSTORE")
	return
}

type ZdiffstoreDestination struct {
	cs []string
}

func (c ZdiffstoreDestination) Numkeys(Numkeys int64) ZdiffstoreNumkeys {
	return ZdiffstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZdiffstoreKey struct {
	cs []string
}

func (c ZdiffstoreKey) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cs: append(c.cs, Key...)}
}

func (c ZdiffstoreKey) Build() []string {
	return c.cs
}

type ZdiffstoreNumkeys struct {
	cs []string
}

func (c ZdiffstoreNumkeys) Key(Key ...string) ZdiffstoreKey {
	return ZdiffstoreKey{cs: append(c.cs, Key...)}
}

type Zincrby struct {
	cs []string
}

func (c Zincrby) Key(Key string) ZincrbyKey {
	return ZincrbyKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zincrby() (c Zincrby) {
	c.cs = append(b.get(), "ZINCRBY")
	return
}

type ZincrbyIncrement struct {
	cs []string
}

func (c ZincrbyIncrement) Member(Member string) ZincrbyMember {
	return ZincrbyMember{cs: append(c.cs, Member)}
}

type ZincrbyKey struct {
	cs []string
}

func (c ZincrbyKey) Increment(Increment int64) ZincrbyIncrement {
	return ZincrbyIncrement{cs: append(c.cs, strconv.FormatInt(Increment, 10))}
}

type ZincrbyMember struct {
	cs []string
}

func (c ZincrbyMember) Build() []string {
	return c.cs
}

type Zinter struct {
	cs []string
}

func (c Zinter) Numkeys(Numkeys int64) ZinterNumkeys {
	return ZinterNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zinter() (c Zinter) {
	c.cs = append(b.get(), "ZINTER")
	return
}

type ZinterAggregateMax struct {
	cs []string
}

func (c ZinterAggregateMax) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMax) Build() []string {
	return c.cs
}

type ZinterAggregateMin struct {
	cs []string
}

func (c ZinterAggregateMin) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateMin) Build() []string {
	return c.cs
}

type ZinterAggregateSum struct {
	cs []string
}

func (c ZinterAggregateSum) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterAggregateSum) Build() []string {
	return c.cs
}

type ZinterKey struct {
	cs []string
}

func (c ZinterKey) Weights(Weight ...int64) ZinterWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterWeights{cs: c.cs}
}

func (c ZinterKey) Sum() ZinterAggregateSum {
	return ZinterAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZinterKey) Min() ZinterAggregateMin {
	return ZinterAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZinterKey) Max() ZinterAggregateMax {
	return ZinterAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZinterKey) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterKey) Key(Key ...string) ZinterKey {
	return ZinterKey{cs: append(c.cs, Key...)}
}

func (c ZinterKey) Build() []string {
	return c.cs
}

type ZinterNumkeys struct {
	cs []string
}

func (c ZinterNumkeys) Key(Key ...string) ZinterKey {
	return ZinterKey{cs: append(c.cs, Key...)}
}

type ZinterWeights struct {
	cs []string
}

func (c ZinterWeights) Sum() ZinterAggregateSum {
	return ZinterAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZinterWeights) Min() ZinterAggregateMin {
	return ZinterAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZinterWeights) Max() ZinterAggregateMax {
	return ZinterAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZinterWeights) Withscores() ZinterWithscoresWithscores {
	return ZinterWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZinterWeights) Weights(Weights ...int64) ZinterWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterWeights{cs: c.cs}
}

func (c ZinterWeights) Build() []string {
	return c.cs
}

type ZinterWithscoresWithscores struct {
	cs []string
}

func (c ZinterWithscoresWithscores) Build() []string {
	return c.cs
}

type Zintercard struct {
	cs []string
}

func (c Zintercard) Numkeys(Numkeys int64) ZintercardNumkeys {
	return ZintercardNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zintercard() (c Zintercard) {
	c.cs = append(b.get(), "ZINTERCARD")
	return
}

type ZintercardKey struct {
	cs []string
}

func (c ZintercardKey) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cs: append(c.cs, Key...)}
}

func (c ZintercardKey) Build() []string {
	return c.cs
}

type ZintercardNumkeys struct {
	cs []string
}

func (c ZintercardNumkeys) Key(Key ...string) ZintercardKey {
	return ZintercardKey{cs: append(c.cs, Key...)}
}

type Zinterstore struct {
	cs []string
}

func (c Zinterstore) Destination(Destination string) ZinterstoreDestination {
	return ZinterstoreDestination{cs: append(c.cs, Destination)}
}

func (b *Builder) Zinterstore() (c Zinterstore) {
	c.cs = append(b.get(), "ZINTERSTORE")
	return
}

type ZinterstoreAggregateMax struct {
	cs []string
}

func (c ZinterstoreAggregateMax) Build() []string {
	return c.cs
}

type ZinterstoreAggregateMin struct {
	cs []string
}

func (c ZinterstoreAggregateMin) Build() []string {
	return c.cs
}

type ZinterstoreAggregateSum struct {
	cs []string
}

func (c ZinterstoreAggregateSum) Build() []string {
	return c.cs
}

type ZinterstoreDestination struct {
	cs []string
}

func (c ZinterstoreDestination) Numkeys(Numkeys int64) ZinterstoreNumkeys {
	return ZinterstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZinterstoreKey struct {
	cs []string
}

func (c ZinterstoreKey) Weights(Weight ...int64) ZinterstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterstoreWeights{cs: c.cs}
}

func (c ZinterstoreKey) Sum() ZinterstoreAggregateSum {
	return ZinterstoreAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZinterstoreKey) Min() ZinterstoreAggregateMin {
	return ZinterstoreAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZinterstoreKey) Max() ZinterstoreAggregateMax {
	return ZinterstoreAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZinterstoreKey) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cs: append(c.cs, Key...)}
}

func (c ZinterstoreKey) Build() []string {
	return c.cs
}

type ZinterstoreNumkeys struct {
	cs []string
}

func (c ZinterstoreNumkeys) Key(Key ...string) ZinterstoreKey {
	return ZinterstoreKey{cs: append(c.cs, Key...)}
}

type ZinterstoreWeights struct {
	cs []string
}

func (c ZinterstoreWeights) Sum() ZinterstoreAggregateSum {
	return ZinterstoreAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZinterstoreWeights) Min() ZinterstoreAggregateMin {
	return ZinterstoreAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZinterstoreWeights) Max() ZinterstoreAggregateMax {
	return ZinterstoreAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZinterstoreWeights) Weights(Weights ...int64) ZinterstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZinterstoreWeights{cs: c.cs}
}

func (c ZinterstoreWeights) Build() []string {
	return c.cs
}

type Zlexcount struct {
	cs []string
}

func (c Zlexcount) Key(Key string) ZlexcountKey {
	return ZlexcountKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zlexcount() (c Zlexcount) {
	c.cs = append(b.get(), "ZLEXCOUNT")
	return
}

type ZlexcountKey struct {
	cs []string
}

func (c ZlexcountKey) Min(Min string) ZlexcountMin {
	return ZlexcountMin{cs: append(c.cs, Min)}
}

type ZlexcountMax struct {
	cs []string
}

func (c ZlexcountMax) Build() []string {
	return c.cs
}

type ZlexcountMin struct {
	cs []string
}

func (c ZlexcountMin) Max(Max string) ZlexcountMax {
	return ZlexcountMax{cs: append(c.cs, Max)}
}

type Zmscore struct {
	cs []string
}

func (c Zmscore) Key(Key string) ZmscoreKey {
	return ZmscoreKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zmscore() (c Zmscore) {
	c.cs = append(b.get(), "ZMSCORE")
	return
}

type ZmscoreKey struct {
	cs []string
}

func (c ZmscoreKey) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cs: append(c.cs, Member...)}
}

type ZmscoreMember struct {
	cs []string
}

func (c ZmscoreMember) Member(Member ...string) ZmscoreMember {
	return ZmscoreMember{cs: append(c.cs, Member...)}
}

func (c ZmscoreMember) Build() []string {
	return c.cs
}

type Zpopmax struct {
	cs []string
}

func (c Zpopmax) Key(Key string) ZpopmaxKey {
	return ZpopmaxKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zpopmax() (c Zpopmax) {
	c.cs = append(b.get(), "ZPOPMAX")
	return
}

type ZpopmaxCount struct {
	cs []string
}

func (c ZpopmaxCount) Build() []string {
	return c.cs
}

type ZpopmaxKey struct {
	cs []string
}

func (c ZpopmaxKey) Count(Count int64) ZpopmaxCount {
	return ZpopmaxCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopmaxKey) Build() []string {
	return c.cs
}

type Zpopmin struct {
	cs []string
}

func (c Zpopmin) Key(Key string) ZpopminKey {
	return ZpopminKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zpopmin() (c Zpopmin) {
	c.cs = append(b.get(), "ZPOPMIN")
	return
}

type ZpopminCount struct {
	cs []string
}

func (c ZpopminCount) Build() []string {
	return c.cs
}

type ZpopminKey struct {
	cs []string
}

func (c ZpopminKey) Count(Count int64) ZpopminCount {
	return ZpopminCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

func (c ZpopminKey) Build() []string {
	return c.cs
}

type Zrandmember struct {
	cs []string
}

func (c Zrandmember) Key(Key string) ZrandmemberKey {
	return ZrandmemberKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrandmember() (c Zrandmember) {
	c.cs = append(b.get(), "ZRANDMEMBER")
	return
}

type ZrandmemberKey struct {
	cs []string
}

func (c ZrandmemberKey) Count(Count int64) ZrandmemberOptionsCount {
	return ZrandmemberOptionsCount{cs: append(c.cs, strconv.FormatInt(Count, 10))}
}

type ZrandmemberOptionsCount struct {
	cs []string
}

func (c ZrandmemberOptionsCount) Withscores() ZrandmemberOptionsWithscoresWithscores {
	return ZrandmemberOptionsWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrandmemberOptionsCount) Build() []string {
	return c.cs
}

type ZrandmemberOptionsWithscoresWithscores struct {
	cs []string
}

func (c ZrandmemberOptionsWithscoresWithscores) Build() []string {
	return c.cs
}

type Zrange struct {
	cs []string
}

func (c Zrange) Key(Key string) ZrangeKey {
	return ZrangeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrange() (c Zrange) {
	c.cs = append(b.get(), "ZRANGE")
	return
}

type ZrangeKey struct {
	cs []string
}

func (c ZrangeKey) Min(Min string) ZrangeMin {
	return ZrangeMin{cs: append(c.cs, Min)}
}

type ZrangeLimit struct {
	cs []string
}

func (c ZrangeLimit) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeLimit) Build() []string {
	return c.cs
}

type ZrangeMax struct {
	cs []string
}

func (c ZrangeMax) Byscore() ZrangeSortbyByscore {
	return ZrangeSortbyByscore{cs: append(c.cs, "BYSCORE")}
}

func (c ZrangeMax) Bylex() ZrangeSortbyBylex {
	return ZrangeSortbyBylex{cs: append(c.cs, "BYLEX")}
}

func (c ZrangeMax) Rev() ZrangeRevRev {
	return ZrangeRevRev{cs: append(c.cs, "REV")}
}

func (c ZrangeMax) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeMax) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeMax) Build() []string {
	return c.cs
}

type ZrangeMin struct {
	cs []string
}

func (c ZrangeMin) Max(Max string) ZrangeMax {
	return ZrangeMax{cs: append(c.cs, Max)}
}

type ZrangeRevRev struct {
	cs []string
}

func (c ZrangeRevRev) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeRevRev) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeRevRev) Build() []string {
	return c.cs
}

type ZrangeSortbyBylex struct {
	cs []string
}

func (c ZrangeSortbyBylex) Rev() ZrangeRevRev {
	return ZrangeRevRev{cs: append(c.cs, "REV")}
}

func (c ZrangeSortbyBylex) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeSortbyBylex) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeSortbyBylex) Build() []string {
	return c.cs
}

type ZrangeSortbyByscore struct {
	cs []string
}

func (c ZrangeSortbyByscore) Rev() ZrangeRevRev {
	return ZrangeRevRev{cs: append(c.cs, "REV")}
}

func (c ZrangeSortbyByscore) Limit(Offset int64, Count int64) ZrangeLimit {
	return ZrangeLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangeSortbyByscore) Withscores() ZrangeWithscoresWithscores {
	return ZrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangeSortbyByscore) Build() []string {
	return c.cs
}

type ZrangeWithscoresWithscores struct {
	cs []string
}

func (c ZrangeWithscoresWithscores) Build() []string {
	return c.cs
}

type Zrangebylex struct {
	cs []string
}

func (c Zrangebylex) Key(Key string) ZrangebylexKey {
	return ZrangebylexKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrangebylex() (c Zrangebylex) {
	c.cs = append(b.get(), "ZRANGEBYLEX")
	return
}

type ZrangebylexKey struct {
	cs []string
}

func (c ZrangebylexKey) Min(Min string) ZrangebylexMin {
	return ZrangebylexMin{cs: append(c.cs, Min)}
}

type ZrangebylexLimit struct {
	cs []string
}

func (c ZrangebylexLimit) Build() []string {
	return c.cs
}

type ZrangebylexMax struct {
	cs []string
}

func (c ZrangebylexMax) Limit(Offset int64, Count int64) ZrangebylexLimit {
	return ZrangebylexLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebylexMax) Build() []string {
	return c.cs
}

type ZrangebylexMin struct {
	cs []string
}

func (c ZrangebylexMin) Max(Max string) ZrangebylexMax {
	return ZrangebylexMax{cs: append(c.cs, Max)}
}

type Zrangebyscore struct {
	cs []string
}

func (c Zrangebyscore) Key(Key string) ZrangebyscoreKey {
	return ZrangebyscoreKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrangebyscore() (c Zrangebyscore) {
	c.cs = append(b.get(), "ZRANGEBYSCORE")
	return
}

type ZrangebyscoreKey struct {
	cs []string
}

func (c ZrangebyscoreKey) Min(Min float64) ZrangebyscoreMin {
	return ZrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZrangebyscoreLimit struct {
	cs []string
}

func (c ZrangebyscoreLimit) Build() []string {
	return c.cs
}

type ZrangebyscoreMax struct {
	cs []string
}

func (c ZrangebyscoreMax) Withscores() ZrangebyscoreWithscoresWithscores {
	return ZrangebyscoreWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrangebyscoreMax) Limit(Offset int64, Count int64) ZrangebyscoreLimit {
	return ZrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebyscoreMax) Build() []string {
	return c.cs
}

type ZrangebyscoreMin struct {
	cs []string
}

func (c ZrangebyscoreMin) Max(Max float64) ZrangebyscoreMax {
	return ZrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type ZrangebyscoreWithscoresWithscores struct {
	cs []string
}

func (c ZrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrangebyscoreLimit {
	return ZrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangebyscoreWithscoresWithscores) Build() []string {
	return c.cs
}

type Zrangestore struct {
	cs []string
}

func (c Zrangestore) Dst(Dst string) ZrangestoreDst {
	return ZrangestoreDst{cs: append(c.cs, Dst)}
}

func (b *Builder) Zrangestore() (c Zrangestore) {
	c.cs = append(b.get(), "ZRANGESTORE")
	return
}

type ZrangestoreDst struct {
	cs []string
}

func (c ZrangestoreDst) Src(Src string) ZrangestoreSrc {
	return ZrangestoreSrc{cs: append(c.cs, Src)}
}

type ZrangestoreLimit struct {
	cs []string
}

func (c ZrangestoreLimit) Build() []string {
	return c.cs
}

type ZrangestoreMax struct {
	cs []string
}

func (c ZrangestoreMax) Byscore() ZrangestoreSortbyByscore {
	return ZrangestoreSortbyByscore{cs: append(c.cs, "BYSCORE")}
}

func (c ZrangestoreMax) Bylex() ZrangestoreSortbyBylex {
	return ZrangestoreSortbyBylex{cs: append(c.cs, "BYLEX")}
}

func (c ZrangestoreMax) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cs: append(c.cs, "REV")}
}

func (c ZrangestoreMax) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreMax) Build() []string {
	return c.cs
}

type ZrangestoreMin struct {
	cs []string
}

func (c ZrangestoreMin) Max(Max string) ZrangestoreMax {
	return ZrangestoreMax{cs: append(c.cs, Max)}
}

type ZrangestoreRevRev struct {
	cs []string
}

func (c ZrangestoreRevRev) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreRevRev) Build() []string {
	return c.cs
}

type ZrangestoreSortbyBylex struct {
	cs []string
}

func (c ZrangestoreSortbyBylex) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cs: append(c.cs, "REV")}
}

func (c ZrangestoreSortbyBylex) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreSortbyBylex) Build() []string {
	return c.cs
}

type ZrangestoreSortbyByscore struct {
	cs []string
}

func (c ZrangestoreSortbyByscore) Rev() ZrangestoreRevRev {
	return ZrangestoreRevRev{cs: append(c.cs, "REV")}
}

func (c ZrangestoreSortbyByscore) Limit(Offset int64, Count int64) ZrangestoreLimit {
	return ZrangestoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrangestoreSortbyByscore) Build() []string {
	return c.cs
}

type ZrangestoreSrc struct {
	cs []string
}

func (c ZrangestoreSrc) Min(Min string) ZrangestoreMin {
	return ZrangestoreMin{cs: append(c.cs, Min)}
}

type Zrank struct {
	cs []string
}

func (c Zrank) Key(Key string) ZrankKey {
	return ZrankKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrank() (c Zrank) {
	c.cs = append(b.get(), "ZRANK")
	return
}

type ZrankKey struct {
	cs []string
}

func (c ZrankKey) Member(Member string) ZrankMember {
	return ZrankMember{cs: append(c.cs, Member)}
}

type ZrankMember struct {
	cs []string
}

func (c ZrankMember) Build() []string {
	return c.cs
}

type Zrem struct {
	cs []string
}

func (c Zrem) Key(Key string) ZremKey {
	return ZremKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrem() (c Zrem) {
	c.cs = append(b.get(), "ZREM")
	return
}

type ZremKey struct {
	cs []string
}

func (c ZremKey) Member(Member ...string) ZremMember {
	return ZremMember{cs: append(c.cs, Member...)}
}

type ZremMember struct {
	cs []string
}

func (c ZremMember) Member(Member ...string) ZremMember {
	return ZremMember{cs: append(c.cs, Member...)}
}

func (c ZremMember) Build() []string {
	return c.cs
}

type Zremrangebylex struct {
	cs []string
}

func (c Zremrangebylex) Key(Key string) ZremrangebylexKey {
	return ZremrangebylexKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebylex() (c Zremrangebylex) {
	c.cs = append(b.get(), "ZREMRANGEBYLEX")
	return
}

type ZremrangebylexKey struct {
	cs []string
}

func (c ZremrangebylexKey) Min(Min string) ZremrangebylexMin {
	return ZremrangebylexMin{cs: append(c.cs, Min)}
}

type ZremrangebylexMax struct {
	cs []string
}

func (c ZremrangebylexMax) Build() []string {
	return c.cs
}

type ZremrangebylexMin struct {
	cs []string
}

func (c ZremrangebylexMin) Max(Max string) ZremrangebylexMax {
	return ZremrangebylexMax{cs: append(c.cs, Max)}
}

type Zremrangebyrank struct {
	cs []string
}

func (c Zremrangebyrank) Key(Key string) ZremrangebyrankKey {
	return ZremrangebyrankKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebyrank() (c Zremrangebyrank) {
	c.cs = append(b.get(), "ZREMRANGEBYRANK")
	return
}

type ZremrangebyrankKey struct {
	cs []string
}

func (c ZremrangebyrankKey) Start(Start int64) ZremrangebyrankStart {
	return ZremrangebyrankStart{cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type ZremrangebyrankStart struct {
	cs []string
}

func (c ZremrangebyrankStart) Stop(Stop int64) ZremrangebyrankStop {
	return ZremrangebyrankStop{cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type ZremrangebyrankStop struct {
	cs []string
}

func (c ZremrangebyrankStop) Build() []string {
	return c.cs
}

type Zremrangebyscore struct {
	cs []string
}

func (c Zremrangebyscore) Key(Key string) ZremrangebyscoreKey {
	return ZremrangebyscoreKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zremrangebyscore() (c Zremrangebyscore) {
	c.cs = append(b.get(), "ZREMRANGEBYSCORE")
	return
}

type ZremrangebyscoreKey struct {
	cs []string
}

func (c ZremrangebyscoreKey) Min(Min float64) ZremrangebyscoreMin {
	return ZremrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZremrangebyscoreMax struct {
	cs []string
}

func (c ZremrangebyscoreMax) Build() []string {
	return c.cs
}

type ZremrangebyscoreMin struct {
	cs []string
}

func (c ZremrangebyscoreMin) Max(Max float64) ZremrangebyscoreMax {
	return ZremrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type Zrevrange struct {
	cs []string
}

func (c Zrevrange) Key(Key string) ZrevrangeKey {
	return ZrevrangeKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrange() (c Zrevrange) {
	c.cs = append(b.get(), "ZREVRANGE")
	return
}

type ZrevrangeKey struct {
	cs []string
}

func (c ZrevrangeKey) Start(Start int64) ZrevrangeStart {
	return ZrevrangeStart{cs: append(c.cs, strconv.FormatInt(Start, 10))}
}

type ZrevrangeStart struct {
	cs []string
}

func (c ZrevrangeStart) Stop(Stop int64) ZrevrangeStop {
	return ZrevrangeStop{cs: append(c.cs, strconv.FormatInt(Stop, 10))}
}

type ZrevrangeStop struct {
	cs []string
}

func (c ZrevrangeStop) Withscores() ZrevrangeWithscoresWithscores {
	return ZrevrangeWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrevrangeStop) Build() []string {
	return c.cs
}

type ZrevrangeWithscoresWithscores struct {
	cs []string
}

func (c ZrevrangeWithscoresWithscores) Build() []string {
	return c.cs
}

type Zrevrangebylex struct {
	cs []string
}

func (c Zrevrangebylex) Key(Key string) ZrevrangebylexKey {
	return ZrevrangebylexKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrangebylex() (c Zrevrangebylex) {
	c.cs = append(b.get(), "ZREVRANGEBYLEX")
	return
}

type ZrevrangebylexKey struct {
	cs []string
}

func (c ZrevrangebylexKey) Max(Max string) ZrevrangebylexMax {
	return ZrevrangebylexMax{cs: append(c.cs, Max)}
}

type ZrevrangebylexLimit struct {
	cs []string
}

func (c ZrevrangebylexLimit) Build() []string {
	return c.cs
}

type ZrevrangebylexMax struct {
	cs []string
}

func (c ZrevrangebylexMax) Min(Min string) ZrevrangebylexMin {
	return ZrevrangebylexMin{cs: append(c.cs, Min)}
}

type ZrevrangebylexMin struct {
	cs []string
}

func (c ZrevrangebylexMin) Limit(Offset int64, Count int64) ZrevrangebylexLimit {
	return ZrevrangebylexLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebylexMin) Build() []string {
	return c.cs
}

type Zrevrangebyscore struct {
	cs []string
}

func (c Zrevrangebyscore) Key(Key string) ZrevrangebyscoreKey {
	return ZrevrangebyscoreKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrangebyscore() (c Zrevrangebyscore) {
	c.cs = append(b.get(), "ZREVRANGEBYSCORE")
	return
}

type ZrevrangebyscoreKey struct {
	cs []string
}

func (c ZrevrangebyscoreKey) Max(Max float64) ZrevrangebyscoreMax {
	return ZrevrangebyscoreMax{cs: append(c.cs, strconv.FormatFloat(Max, 'f', -1, 64))}
}

type ZrevrangebyscoreLimit struct {
	cs []string
}

func (c ZrevrangebyscoreLimit) Build() []string {
	return c.cs
}

type ZrevrangebyscoreMax struct {
	cs []string
}

func (c ZrevrangebyscoreMax) Min(Min float64) ZrevrangebyscoreMin {
	return ZrevrangebyscoreMin{cs: append(c.cs, strconv.FormatFloat(Min, 'f', -1, 64))}
}

type ZrevrangebyscoreMin struct {
	cs []string
}

func (c ZrevrangebyscoreMin) Withscores() ZrevrangebyscoreWithscoresWithscores {
	return ZrevrangebyscoreWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZrevrangebyscoreMin) Limit(Offset int64, Count int64) ZrevrangebyscoreLimit {
	return ZrevrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebyscoreMin) Build() []string {
	return c.cs
}

type ZrevrangebyscoreWithscoresWithscores struct {
	cs []string
}

func (c ZrevrangebyscoreWithscoresWithscores) Limit(Offset int64, Count int64) ZrevrangebyscoreLimit {
	return ZrevrangebyscoreLimit{cs: append(c.cs, "LIMIT", strconv.FormatInt(Offset, 10), strconv.FormatInt(Count, 10))}
}

func (c ZrevrangebyscoreWithscoresWithscores) Build() []string {
	return c.cs
}

type Zrevrank struct {
	cs []string
}

func (c Zrevrank) Key(Key string) ZrevrankKey {
	return ZrevrankKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zrevrank() (c Zrevrank) {
	c.cs = append(b.get(), "ZREVRANK")
	return
}

type ZrevrankKey struct {
	cs []string
}

func (c ZrevrankKey) Member(Member string) ZrevrankMember {
	return ZrevrankMember{cs: append(c.cs, Member)}
}

type ZrevrankMember struct {
	cs []string
}

func (c ZrevrankMember) Build() []string {
	return c.cs
}

type Zscan struct {
	cs []string
}

func (c Zscan) Key(Key string) ZscanKey {
	return ZscanKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zscan() (c Zscan) {
	c.cs = append(b.get(), "ZSCAN")
	return
}

type ZscanCount struct {
	cs []string
}

func (c ZscanCount) Build() []string {
	return c.cs
}

type ZscanCursor struct {
	cs []string
}

func (c ZscanCursor) Match(Pattern string) ZscanMatch {
	return ZscanMatch{cs: append(c.cs, "MATCH", Pattern)}
}

func (c ZscanCursor) Count(Count int64) ZscanCount {
	return ZscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ZscanCursor) Build() []string {
	return c.cs
}

type ZscanKey struct {
	cs []string
}

func (c ZscanKey) Cursor(Cursor int64) ZscanCursor {
	return ZscanCursor{cs: append(c.cs, strconv.FormatInt(Cursor, 10))}
}

type ZscanMatch struct {
	cs []string
}

func (c ZscanMatch) Count(Count int64) ZscanCount {
	return ZscanCount{cs: append(c.cs, "COUNT", strconv.FormatInt(Count, 10))}
}

func (c ZscanMatch) Build() []string {
	return c.cs
}

type Zscore struct {
	cs []string
}

func (c Zscore) Key(Key string) ZscoreKey {
	return ZscoreKey{cs: append(c.cs, Key)}
}

func (b *Builder) Zscore() (c Zscore) {
	c.cs = append(b.get(), "ZSCORE")
	return
}

type ZscoreKey struct {
	cs []string
}

func (c ZscoreKey) Member(Member string) ZscoreMember {
	return ZscoreMember{cs: append(c.cs, Member)}
}

type ZscoreMember struct {
	cs []string
}

func (c ZscoreMember) Build() []string {
	return c.cs
}

type Zunion struct {
	cs []string
}

func (c Zunion) Numkeys(Numkeys int64) ZunionNumkeys {
	return ZunionNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

func (b *Builder) Zunion() (c Zunion) {
	c.cs = append(b.get(), "ZUNION")
	return
}

type ZunionAggregateMax struct {
	cs []string
}

func (c ZunionAggregateMax) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMax) Build() []string {
	return c.cs
}

type ZunionAggregateMin struct {
	cs []string
}

func (c ZunionAggregateMin) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateMin) Build() []string {
	return c.cs
}

type ZunionAggregateSum struct {
	cs []string
}

func (c ZunionAggregateSum) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionAggregateSum) Build() []string {
	return c.cs
}

type ZunionKey struct {
	cs []string
}

func (c ZunionKey) Weights(Weight ...int64) ZunionWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionWeights{cs: c.cs}
}

func (c ZunionKey) Sum() ZunionAggregateSum {
	return ZunionAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZunionKey) Min() ZunionAggregateMin {
	return ZunionAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZunionKey) Max() ZunionAggregateMax {
	return ZunionAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZunionKey) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionKey) Key(Key ...string) ZunionKey {
	return ZunionKey{cs: append(c.cs, Key...)}
}

func (c ZunionKey) Build() []string {
	return c.cs
}

type ZunionNumkeys struct {
	cs []string
}

func (c ZunionNumkeys) Key(Key ...string) ZunionKey {
	return ZunionKey{cs: append(c.cs, Key...)}
}

type ZunionWeights struct {
	cs []string
}

func (c ZunionWeights) Sum() ZunionAggregateSum {
	return ZunionAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZunionWeights) Min() ZunionAggregateMin {
	return ZunionAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZunionWeights) Max() ZunionAggregateMax {
	return ZunionAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZunionWeights) Withscores() ZunionWithscoresWithscores {
	return ZunionWithscoresWithscores{cs: append(c.cs, "WITHSCORES")}
}

func (c ZunionWeights) Weights(Weights ...int64) ZunionWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionWeights{cs: c.cs}
}

func (c ZunionWeights) Build() []string {
	return c.cs
}

type ZunionWithscoresWithscores struct {
	cs []string
}

func (c ZunionWithscoresWithscores) Build() []string {
	return c.cs
}

type Zunionstore struct {
	cs []string
}

func (c Zunionstore) Destination(Destination string) ZunionstoreDestination {
	return ZunionstoreDestination{cs: append(c.cs, Destination)}
}

func (b *Builder) Zunionstore() (c Zunionstore) {
	c.cs = append(b.get(), "ZUNIONSTORE")
	return
}

type ZunionstoreAggregateMax struct {
	cs []string
}

func (c ZunionstoreAggregateMax) Build() []string {
	return c.cs
}

type ZunionstoreAggregateMin struct {
	cs []string
}

func (c ZunionstoreAggregateMin) Build() []string {
	return c.cs
}

type ZunionstoreAggregateSum struct {
	cs []string
}

func (c ZunionstoreAggregateSum) Build() []string {
	return c.cs
}

type ZunionstoreDestination struct {
	cs []string
}

func (c ZunionstoreDestination) Numkeys(Numkeys int64) ZunionstoreNumkeys {
	return ZunionstoreNumkeys{cs: append(c.cs, strconv.FormatInt(Numkeys, 10))}
}

type ZunionstoreKey struct {
	cs []string
}

func (c ZunionstoreKey) Weights(Weight ...int64) ZunionstoreWeights {
	c.cs = append(c.cs, "WEIGHTS")
	for _, n := range Weight {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionstoreWeights{cs: c.cs}
}

func (c ZunionstoreKey) Sum() ZunionstoreAggregateSum {
	return ZunionstoreAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZunionstoreKey) Min() ZunionstoreAggregateMin {
	return ZunionstoreAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZunionstoreKey) Max() ZunionstoreAggregateMax {
	return ZunionstoreAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZunionstoreKey) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cs: append(c.cs, Key...)}
}

func (c ZunionstoreKey) Build() []string {
	return c.cs
}

type ZunionstoreNumkeys struct {
	cs []string
}

func (c ZunionstoreNumkeys) Key(Key ...string) ZunionstoreKey {
	return ZunionstoreKey{cs: append(c.cs, Key...)}
}

type ZunionstoreWeights struct {
	cs []string
}

func (c ZunionstoreWeights) Sum() ZunionstoreAggregateSum {
	return ZunionstoreAggregateSum{cs: append(c.cs, "SUM")}
}

func (c ZunionstoreWeights) Min() ZunionstoreAggregateMin {
	return ZunionstoreAggregateMin{cs: append(c.cs, "MIN")}
}

func (c ZunionstoreWeights) Max() ZunionstoreAggregateMax {
	return ZunionstoreAggregateMax{cs: append(c.cs, "MAX")}
}

func (c ZunionstoreWeights) Weights(Weights ...int64) ZunionstoreWeights {
	for _, n := range Weights {
		c.cs = append(c.cs, strconv.FormatInt(n, 10))
	}
	return ZunionstoreWeights{cs: c.cs}
}

func (c ZunionstoreWeights) Build() []string {
	return c.cs
}
