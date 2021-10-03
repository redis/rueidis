package cmds

import "strconv"

type aclcat struct {
	cmds []string
}

func (c aclcat) Categoryname(categoryname string) aclcatcategoryname {
	var cmds []string
	cmds = append(c.cmds, categoryname)
	return aclcatcategoryname{cmds: cmds}
}

func (c aclcat) Build() []string {
	return c.cmds
}

func AclCat() (c aclcat) {
	c.cmds = append(c.cmds, "ACL", "CAT")
	return
}

type aclcatcategoryname struct {
	cmds []string
}

func (c aclcatcategoryname) Build() []string {
	return c.cmds
}

type acldeluser struct {
	cmds []string
}

func (c acldeluser) Username(username ...string) acldeluserusername {
	var cmds []string
	cmds = append(cmds, username...)
	return acldeluserusername{cmds: cmds}
}

func AclDeluser() (c acldeluser) {
	c.cmds = append(c.cmds, "ACL", "DELUSER")
	return
}

type acldeluserusername struct {
	cmds []string
}

func (c acldeluserusername) Username(username ...string) acldeluserusername {
	var cmds []string
	cmds = append(cmds, username...)
	return acldeluserusername{cmds: cmds}
}

type aclgenpass struct {
	cmds []string
}

func (c aclgenpass) Bits(bits int64) aclgenpassbits {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(bits, 10))
	return aclgenpassbits{cmds: cmds}
}

func (c aclgenpass) Build() []string {
	return c.cmds
}

func AclGenpass() (c aclgenpass) {
	c.cmds = append(c.cmds, "ACL", "GENPASS")
	return
}

type aclgenpassbits struct {
	cmds []string
}

func (c aclgenpassbits) Build() []string {
	return c.cmds
}

type aclgetuser struct {
	cmds []string
}

func (c aclgetuser) Username(username string) aclgetuserusername {
	var cmds []string
	cmds = append(c.cmds, username)
	return aclgetuserusername{cmds: cmds}
}

func AclGetuser() (c aclgetuser) {
	c.cmds = append(c.cmds, "ACL", "GETUSER")
	return
}

type aclgetuserusername struct {
	cmds []string
}

func (c aclgetuserusername) Build() []string {
	return c.cmds
}

type aclhelp struct {
	cmds []string
}

func (c aclhelp) Build() []string {
	return c.cmds
}

func AclHelp() (c aclhelp) {
	c.cmds = append(c.cmds, "ACL", "HELP")
	return
}

type acllist struct {
	cmds []string
}

func (c acllist) Build() []string {
	return c.cmds
}

func AclList() (c acllist) {
	c.cmds = append(c.cmds, "ACL", "LIST")
	return
}

type aclload struct {
	cmds []string
}

func (c aclload) Build() []string {
	return c.cmds
}

func AclLoad() (c aclload) {
	c.cmds = append(c.cmds, "ACL", "LOAD")
	return
}

type acllog struct {
	cmds []string
}

func (c acllog) CountOrReset(countOrReset string) acllogcountorreset {
	var cmds []string
	cmds = append(c.cmds, countOrReset)
	return acllogcountorreset{cmds: cmds}
}

func (c acllog) Build() []string {
	return c.cmds
}

func AclLog() (c acllog) {
	c.cmds = append(c.cmds, "ACL", "LOG")
	return
}

type acllogcountorreset struct {
	cmds []string
}

func (c acllogcountorreset) Build() []string {
	return c.cmds
}

type aclsave struct {
	cmds []string
}

func (c aclsave) Build() []string {
	return c.cmds
}

func AclSave() (c aclsave) {
	c.cmds = append(c.cmds, "ACL", "SAVE")
	return
}

type aclsetuser struct {
	cmds []string
}

func (c aclsetuser) Username(username string) aclsetuserusername {
	var cmds []string
	cmds = append(c.cmds, username)
	return aclsetuserusername{cmds: cmds}
}

func AclSetuser() (c aclsetuser) {
	c.cmds = append(c.cmds, "ACL", "SETUSER")
	return
}

type aclsetuserrule struct {
	cmds []string
}

func (c aclsetuserrule) Rule(rule ...string) aclsetuserrule {
	var cmds []string
	cmds = append(cmds, rule...)
	return aclsetuserrule{cmds: cmds}
}

func (c aclsetuserrule) Build() []string {
	return c.cmds
}

type aclsetuserusername struct {
	cmds []string
}

func (c aclsetuserusername) Rule(rule ...string) aclsetuserrule {
	var cmds []string
	cmds = append(cmds, rule...)
	return aclsetuserrule{cmds: cmds}
}

func (c aclsetuserusername) Build() []string {
	return c.cmds
}

type aclusers struct {
	cmds []string
}

func (c aclusers) Build() []string {
	return c.cmds
}

func AclUsers() (c aclusers) {
	c.cmds = append(c.cmds, "ACL", "USERS")
	return
}

type aclwhoami struct {
	cmds []string
}

func (c aclwhoami) Build() []string {
	return c.cmds
}

func AclWhoami() (c aclwhoami) {
	c.cmds = append(c.cmds, "ACL", "WHOAMI")
	return
}

type rappend struct {
	cmds []string
}

func (c rappend) Key(key string) appendkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return appendkey{cmds: cmds}
}

func Append() (c rappend) {
	c.cmds = append(c.cmds, "APPEND")
	return
}

type appendkey struct {
	cmds []string
}

func (c appendkey) Value(value string) appendvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return appendvalue{cmds: cmds}
}

type appendvalue struct {
	cmds []string
}

func (c appendvalue) Build() []string {
	return c.cmds
}

type asking struct {
	cmds []string
}

func (c asking) Build() []string {
	return c.cmds
}

func Asking() (c asking) {
	c.cmds = append(c.cmds, "ASKING")
	return
}

type auth struct {
	cmds []string
}

func (c auth) Username(username string) authusername {
	var cmds []string
	cmds = append(c.cmds, username)
	return authusername{cmds: cmds}
}

func (c auth) Password(password string) authpassword {
	var cmds []string
	cmds = append(c.cmds, password)
	return authpassword{cmds: cmds}
}

func Auth() (c auth) {
	c.cmds = append(c.cmds, "AUTH")
	return
}

type authpassword struct {
	cmds []string
}

func (c authpassword) Build() []string {
	return c.cmds
}

type authusername struct {
	cmds []string
}

func (c authusername) Password(password string) authpassword {
	var cmds []string
	cmds = append(c.cmds, password)
	return authpassword{cmds: cmds}
}

type bgrewriteaof struct {
	cmds []string
}

func (c bgrewriteaof) Build() []string {
	return c.cmds
}

func Bgrewriteaof() (c bgrewriteaof) {
	c.cmds = append(c.cmds, "BGREWRITEAOF")
	return
}

type bgsave struct {
	cmds []string
}

func (c bgsave) Schedule() bgsavescheduleschedule {
	var cmds []string
	cmds = append(c.cmds, "SCHEDULE")
	return bgsavescheduleschedule{cmds: cmds}
}

func (c bgsave) Build() []string {
	return c.cmds
}

func Bgsave() (c bgsave) {
	c.cmds = append(c.cmds, "BGSAVE")
	return
}

type bgsavescheduleschedule struct {
	cmds []string
}

func (c bgsavescheduleschedule) Build() []string {
	return c.cmds
}

type bitcount struct {
	cmds []string
}

func (c bitcount) Key(key string) bitcountkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return bitcountkey{cmds: cmds}
}

func Bitcount() (c bitcount) {
	c.cmds = append(c.cmds, "BITCOUNT")
	return
}

type bitcountkey struct {
	cmds []string
}

func (c bitcountkey) StartEnd(start int64, end int64) bitcountstartend {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(start, 10), strconv.FormatInt(end, 10))
	return bitcountstartend{cmds: cmds}
}

func (c bitcountkey) Build() []string {
	return c.cmds
}

type bitcountstartend struct {
	cmds []string
}

func (c bitcountstartend) Build() []string {
	return c.cmds
}

type bitfield struct {
	cmds []string
}

func (c bitfield) Key(key string) bitfieldkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return bitfieldkey{cmds: cmds}
}

func Bitfield() (c bitfield) {
	c.cmds = append(c.cmds, "BITFIELD")
	return
}

type bitfieldfail struct {
	cmds []string
}

func (c bitfieldfail) Build() []string {
	return c.cmds
}

type bitfieldget struct {
	cmds []string
}

func (c bitfieldget) Set(typ string, offset int64, value int64) bitfieldset {
	var cmds []string
	cmds = append(c.cmds, "SET", typ, strconv.FormatInt(offset, 10), strconv.FormatInt(value, 10))
	return bitfieldset{cmds: cmds}
}

func (c bitfieldget) Incrby(typ string, offset int64, increment int64) bitfieldincrby {
	var cmds []string
	cmds = append(c.cmds, "INCRBY", typ, strconv.FormatInt(offset, 10), strconv.FormatInt(increment, 10))
	return bitfieldincrby{cmds: cmds}
}

func (c bitfieldget) Wrap() bitfieldwrap {
	var cmds []string
	cmds = append(c.cmds, "WRAP")
	return bitfieldwrap{cmds: cmds}
}

func (c bitfieldget) Sat() bitfieldsat {
	var cmds []string
	cmds = append(c.cmds, "SAT")
	return bitfieldsat{cmds: cmds}
}

func (c bitfieldget) Fail() bitfieldfail {
	var cmds []string
	cmds = append(c.cmds, "FAIL")
	return bitfieldfail{cmds: cmds}
}

func (c bitfieldget) Build() []string {
	return c.cmds
}

type bitfieldincrby struct {
	cmds []string
}

func (c bitfieldincrby) Wrap() bitfieldwrap {
	var cmds []string
	cmds = append(c.cmds, "WRAP")
	return bitfieldwrap{cmds: cmds}
}

func (c bitfieldincrby) Sat() bitfieldsat {
	var cmds []string
	cmds = append(c.cmds, "SAT")
	return bitfieldsat{cmds: cmds}
}

func (c bitfieldincrby) Fail() bitfieldfail {
	var cmds []string
	cmds = append(c.cmds, "FAIL")
	return bitfieldfail{cmds: cmds}
}

func (c bitfieldincrby) Build() []string {
	return c.cmds
}

type bitfieldkey struct {
	cmds []string
}

func (c bitfieldkey) Get(typ string, offset int64) bitfieldget {
	var cmds []string
	cmds = append(c.cmds, "GET", typ, strconv.FormatInt(offset, 10))
	return bitfieldget{cmds: cmds}
}

func (c bitfieldkey) Set(typ string, offset int64, value int64) bitfieldset {
	var cmds []string
	cmds = append(c.cmds, "SET", typ, strconv.FormatInt(offset, 10), strconv.FormatInt(value, 10))
	return bitfieldset{cmds: cmds}
}

func (c bitfieldkey) Incrby(typ string, offset int64, increment int64) bitfieldincrby {
	var cmds []string
	cmds = append(c.cmds, "INCRBY", typ, strconv.FormatInt(offset, 10), strconv.FormatInt(increment, 10))
	return bitfieldincrby{cmds: cmds}
}

func (c bitfieldkey) Wrap() bitfieldwrap {
	var cmds []string
	cmds = append(c.cmds, "WRAP")
	return bitfieldwrap{cmds: cmds}
}

func (c bitfieldkey) Sat() bitfieldsat {
	var cmds []string
	cmds = append(c.cmds, "SAT")
	return bitfieldsat{cmds: cmds}
}

func (c bitfieldkey) Fail() bitfieldfail {
	var cmds []string
	cmds = append(c.cmds, "FAIL")
	return bitfieldfail{cmds: cmds}
}

func (c bitfieldkey) Build() []string {
	return c.cmds
}

type bitfieldro struct {
	cmds []string
}

func (c bitfieldro) Key(key string) bitfieldrokey {
	var cmds []string
	cmds = append(c.cmds, key)
	return bitfieldrokey{cmds: cmds}
}

func Bitfieldro() (c bitfieldro) {
	c.cmds = append(c.cmds, "BITFIELD_RO")
	return
}

type bitfieldroget struct {
	cmds []string
}

func (c bitfieldroget) Build() []string {
	return c.cmds
}

type bitfieldrokey struct {
	cmds []string
}

func (c bitfieldrokey) Get(typ string, offset int64) bitfieldroget {
	var cmds []string
	cmds = append(c.cmds, "GET", typ, strconv.FormatInt(offset, 10))
	return bitfieldroget{cmds: cmds}
}

type bitfieldsat struct {
	cmds []string
}

func (c bitfieldsat) Build() []string {
	return c.cmds
}

type bitfieldset struct {
	cmds []string
}

func (c bitfieldset) Incrby(typ string, offset int64, increment int64) bitfieldincrby {
	var cmds []string
	cmds = append(c.cmds, "INCRBY", typ, strconv.FormatInt(offset, 10), strconv.FormatInt(increment, 10))
	return bitfieldincrby{cmds: cmds}
}

func (c bitfieldset) Wrap() bitfieldwrap {
	var cmds []string
	cmds = append(c.cmds, "WRAP")
	return bitfieldwrap{cmds: cmds}
}

func (c bitfieldset) Sat() bitfieldsat {
	var cmds []string
	cmds = append(c.cmds, "SAT")
	return bitfieldsat{cmds: cmds}
}

func (c bitfieldset) Fail() bitfieldfail {
	var cmds []string
	cmds = append(c.cmds, "FAIL")
	return bitfieldfail{cmds: cmds}
}

func (c bitfieldset) Build() []string {
	return c.cmds
}

type bitfieldwrap struct {
	cmds []string
}

func (c bitfieldwrap) Build() []string {
	return c.cmds
}

type bitop struct {
	cmds []string
}

func (c bitop) Operation(operation string) bitopoperation {
	var cmds []string
	cmds = append(c.cmds, operation)
	return bitopoperation{cmds: cmds}
}

func Bitop() (c bitop) {
	c.cmds = append(c.cmds, "BITOP")
	return
}

type bitopdestkey struct {
	cmds []string
}

func (c bitopdestkey) Key(key ...string) bitopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return bitopkey{cmds: cmds}
}

type bitopkey struct {
	cmds []string
}

func (c bitopkey) Key(key ...string) bitopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return bitopkey{cmds: cmds}
}

type bitopoperation struct {
	cmds []string
}

func (c bitopoperation) Destkey(destkey string) bitopdestkey {
	var cmds []string
	cmds = append(c.cmds, destkey)
	return bitopdestkey{cmds: cmds}
}

type bitpos struct {
	cmds []string
}

func (c bitpos) Key(key string) bitposkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return bitposkey{cmds: cmds}
}

func Bitpos() (c bitpos) {
	c.cmds = append(c.cmds, "BITPOS")
	return
}

type bitposbit struct {
	cmds []string
}

func (c bitposbit) Start(start int64) bitposindexstart {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(start, 10))
	return bitposindexstart{cmds: cmds}
}

type bitposindexend struct {
	cmds []string
}

func (c bitposindexend) Build() []string {
	return c.cmds
}

type bitposindexstart struct {
	cmds []string
}

func (c bitposindexstart) End(end int64) bitposindexend {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(end, 10))
	return bitposindexend{cmds: cmds}
}

func (c bitposindexstart) Build() []string {
	return c.cmds
}

type bitposkey struct {
	cmds []string
}

func (c bitposkey) Bit(bit int64) bitposbit {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(bit, 10))
	return bitposbit{cmds: cmds}
}

type blmove struct {
	cmds []string
}

func (c blmove) Source(source string) blmovesource {
	var cmds []string
	cmds = append(c.cmds, source)
	return blmovesource{cmds: cmds}
}

func Blmove() (c blmove) {
	c.cmds = append(c.cmds, "BLMOVE")
	return
}

type blmovedestination struct {
	cmds []string
}

func (c blmovedestination) Left() blmovewherefromleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return blmovewherefromleft{cmds: cmds}
}

func (c blmovedestination) Right() blmovewherefromright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return blmovewherefromright{cmds: cmds}
}

type blmovesource struct {
	cmds []string
}

func (c blmovesource) Destination(destination string) blmovedestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return blmovedestination{cmds: cmds}
}

type blmovetimeout struct {
	cmds []string
}

func (c blmovetimeout) Build() []string {
	return c.cmds
}

type blmovewherefromleft struct {
	cmds []string
}

func (c blmovewherefromleft) Left() blmovewheretoleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return blmovewheretoleft{cmds: cmds}
}

func (c blmovewherefromleft) Right() blmovewheretoright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return blmovewheretoright{cmds: cmds}
}

type blmovewherefromright struct {
	cmds []string
}

func (c blmovewherefromright) Left() blmovewheretoleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return blmovewheretoleft{cmds: cmds}
}

func (c blmovewherefromright) Right() blmovewheretoright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return blmovewheretoright{cmds: cmds}
}

type blmovewheretoleft struct {
	cmds []string
}

func (c blmovewheretoleft) Timeout(timeout float64) blmovetimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return blmovetimeout{cmds: cmds}
}

type blmovewheretoright struct {
	cmds []string
}

func (c blmovewheretoright) Timeout(timeout float64) blmovetimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return blmovetimeout{cmds: cmds}
}

type blmpop struct {
	cmds []string
}

func (c blmpop) Timeout(timeout float64) blmpoptimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return blmpoptimeout{cmds: cmds}
}

func Blmpop() (c blmpop) {
	c.cmds = append(c.cmds, "BLMPOP")
	return
}

type blmpopcount struct {
	cmds []string
}

func (c blmpopcount) Build() []string {
	return c.cmds
}

type blmpopkey struct {
	cmds []string
}

func (c blmpopkey) Left() blmpopwhereleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return blmpopwhereleft{cmds: cmds}
}

func (c blmpopkey) Right() blmpopwhereright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return blmpopwhereright{cmds: cmds}
}

func (c blmpopkey) Key(key ...string) blmpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return blmpopkey{cmds: cmds}
}

type blmpopnumkeys struct {
	cmds []string
}

func (c blmpopnumkeys) Key(key ...string) blmpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return blmpopkey{cmds: cmds}
}

func (c blmpopnumkeys) Left() blmpopwhereleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return blmpopwhereleft{cmds: cmds}
}

func (c blmpopnumkeys) Right() blmpopwhereright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return blmpopwhereright{cmds: cmds}
}

type blmpoptimeout struct {
	cmds []string
}

func (c blmpoptimeout) Numkeys(numkeys int64) blmpopnumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return blmpopnumkeys{cmds: cmds}
}

type blmpopwhereleft struct {
	cmds []string
}

func (c blmpopwhereleft) Count(count int64) blmpopcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return blmpopcount{cmds: cmds}
}

func (c blmpopwhereleft) Build() []string {
	return c.cmds
}

type blmpopwhereright struct {
	cmds []string
}

func (c blmpopwhereright) Count(count int64) blmpopcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return blmpopcount{cmds: cmds}
}

func (c blmpopwhereright) Build() []string {
	return c.cmds
}

type blpop struct {
	cmds []string
}

func (c blpop) Key(key ...string) blpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return blpopkey{cmds: cmds}
}

func Blpop() (c blpop) {
	c.cmds = append(c.cmds, "BLPOP")
	return
}

type blpopkey struct {
	cmds []string
}

func (c blpopkey) Timeout(timeout float64) blpoptimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return blpoptimeout{cmds: cmds}
}

func (c blpopkey) Key(key ...string) blpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return blpopkey{cmds: cmds}
}

type blpoptimeout struct {
	cmds []string
}

func (c blpoptimeout) Build() []string {
	return c.cmds
}

type brpop struct {
	cmds []string
}

func (c brpop) Key(key ...string) brpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return brpopkey{cmds: cmds}
}

func Brpop() (c brpop) {
	c.cmds = append(c.cmds, "BRPOP")
	return
}

type brpopkey struct {
	cmds []string
}

func (c brpopkey) Timeout(timeout float64) brpoptimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return brpoptimeout{cmds: cmds}
}

func (c brpopkey) Key(key ...string) brpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return brpopkey{cmds: cmds}
}

type brpoplpush struct {
	cmds []string
}

func (c brpoplpush) Source(source string) brpoplpushsource {
	var cmds []string
	cmds = append(c.cmds, source)
	return brpoplpushsource{cmds: cmds}
}

func Brpoplpush() (c brpoplpush) {
	c.cmds = append(c.cmds, "BRPOPLPUSH")
	return
}

type brpoplpushdestination struct {
	cmds []string
}

func (c brpoplpushdestination) Timeout(timeout float64) brpoplpushtimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return brpoplpushtimeout{cmds: cmds}
}

type brpoplpushsource struct {
	cmds []string
}

func (c brpoplpushsource) Destination(destination string) brpoplpushdestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return brpoplpushdestination{cmds: cmds}
}

type brpoplpushtimeout struct {
	cmds []string
}

func (c brpoplpushtimeout) Build() []string {
	return c.cmds
}

type brpoptimeout struct {
	cmds []string
}

func (c brpoptimeout) Build() []string {
	return c.cmds
}

type bzpopmax struct {
	cmds []string
}

func (c bzpopmax) Key(key ...string) bzpopmaxkey {
	var cmds []string
	cmds = append(cmds, key...)
	return bzpopmaxkey{cmds: cmds}
}

func Bzpopmax() (c bzpopmax) {
	c.cmds = append(c.cmds, "BZPOPMAX")
	return
}

type bzpopmaxkey struct {
	cmds []string
}

func (c bzpopmaxkey) Timeout(timeout float64) bzpopmaxtimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return bzpopmaxtimeout{cmds: cmds}
}

func (c bzpopmaxkey) Key(key ...string) bzpopmaxkey {
	var cmds []string
	cmds = append(cmds, key...)
	return bzpopmaxkey{cmds: cmds}
}

type bzpopmaxtimeout struct {
	cmds []string
}

func (c bzpopmaxtimeout) Build() []string {
	return c.cmds
}

type bzpopmin struct {
	cmds []string
}

func (c bzpopmin) Key(key ...string) bzpopminkey {
	var cmds []string
	cmds = append(cmds, key...)
	return bzpopminkey{cmds: cmds}
}

func Bzpopmin() (c bzpopmin) {
	c.cmds = append(c.cmds, "BZPOPMIN")
	return
}

type bzpopminkey struct {
	cmds []string
}

func (c bzpopminkey) Timeout(timeout float64) bzpopmintimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(timeout, 'f', -1, 64))
	return bzpopmintimeout{cmds: cmds}
}

func (c bzpopminkey) Key(key ...string) bzpopminkey {
	var cmds []string
	cmds = append(cmds, key...)
	return bzpopminkey{cmds: cmds}
}

type bzpopmintimeout struct {
	cmds []string
}

func (c bzpopmintimeout) Build() []string {
	return c.cmds
}

type clientcaching struct {
	cmds []string
}

func (c clientcaching) Yes() clientcachingmodeyes {
	var cmds []string
	cmds = append(c.cmds, "YES")
	return clientcachingmodeyes{cmds: cmds}
}

func (c clientcaching) No() clientcachingmodeno {
	var cmds []string
	cmds = append(c.cmds, "NO")
	return clientcachingmodeno{cmds: cmds}
}

func ClientCaching() (c clientcaching) {
	c.cmds = append(c.cmds, "CLIENT", "CACHING")
	return
}

type clientcachingmodeno struct {
	cmds []string
}

func (c clientcachingmodeno) Build() []string {
	return c.cmds
}

type clientcachingmodeyes struct {
	cmds []string
}

func (c clientcachingmodeyes) Build() []string {
	return c.cmds
}

type clientgetname struct {
	cmds []string
}

func (c clientgetname) Build() []string {
	return c.cmds
}

func ClientGetname() (c clientgetname) {
	c.cmds = append(c.cmds, "CLIENT", "GETNAME")
	return
}

type clientgetredir struct {
	cmds []string
}

func (c clientgetredir) Build() []string {
	return c.cmds
}

func ClientGetredir() (c clientgetredir) {
	c.cmds = append(c.cmds, "CLIENT", "GETREDIR")
	return
}

type clientid struct {
	cmds []string
}

func (c clientid) Build() []string {
	return c.cmds
}

func ClientId() (c clientid) {
	c.cmds = append(c.cmds, "CLIENT", "ID")
	return
}

type clientinfo struct {
	cmds []string
}

func (c clientinfo) Build() []string {
	return c.cmds
}

func ClientInfo() (c clientinfo) {
	c.cmds = append(c.cmds, "CLIENT", "INFO")
	return
}

type clientkill struct {
	cmds []string
}

func (c clientkill) Ipport(ipport string) clientkillipport {
	var cmds []string
	cmds = append(c.cmds, ipport)
	return clientkillipport{cmds: cmds}
}

func (c clientkill) Id(clientid int64) clientkillid {
	var cmds []string
	cmds = append(c.cmds, "ID", strconv.FormatInt(clientid, 10))
	return clientkillid{cmds: cmds}
}

func (c clientkill) Normal() clientkillnormal {
	var cmds []string
	cmds = append(c.cmds, "normal")
	return clientkillnormal{cmds: cmds}
}

func (c clientkill) Master() clientkillmaster {
	var cmds []string
	cmds = append(c.cmds, "master")
	return clientkillmaster{cmds: cmds}
}

func (c clientkill) Slave() clientkillslave {
	var cmds []string
	cmds = append(c.cmds, "slave")
	return clientkillslave{cmds: cmds}
}

func (c clientkill) Pubsub() clientkillpubsub {
	var cmds []string
	cmds = append(c.cmds, "pubsub")
	return clientkillpubsub{cmds: cmds}
}

func (c clientkill) User(username string) clientkilluser {
	var cmds []string
	cmds = append(c.cmds, "USER", username)
	return clientkilluser{cmds: cmds}
}

func (c clientkill) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkill) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkill) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkill) Build() []string {
	return c.cmds
}

func ClientKill() (c clientkill) {
	c.cmds = append(c.cmds, "CLIENT", "KILL")
	return
}

type clientkilladdr struct {
	cmds []string
}

func (c clientkilladdr) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkilladdr) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkilladdr) Build() []string {
	return c.cmds
}

type clientkillid struct {
	cmds []string
}

func (c clientkillid) Normal() clientkillnormal {
	var cmds []string
	cmds = append(c.cmds, "normal")
	return clientkillnormal{cmds: cmds}
}

func (c clientkillid) Master() clientkillmaster {
	var cmds []string
	cmds = append(c.cmds, "master")
	return clientkillmaster{cmds: cmds}
}

func (c clientkillid) Slave() clientkillslave {
	var cmds []string
	cmds = append(c.cmds, "slave")
	return clientkillslave{cmds: cmds}
}

func (c clientkillid) Pubsub() clientkillpubsub {
	var cmds []string
	cmds = append(c.cmds, "pubsub")
	return clientkillpubsub{cmds: cmds}
}

func (c clientkillid) User(username string) clientkilluser {
	var cmds []string
	cmds = append(c.cmds, "USER", username)
	return clientkilluser{cmds: cmds}
}

func (c clientkillid) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkillid) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkillid) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkillid) Build() []string {
	return c.cmds
}

type clientkillipport struct {
	cmds []string
}

func (c clientkillipport) Id(clientid int64) clientkillid {
	var cmds []string
	cmds = append(c.cmds, "ID", strconv.FormatInt(clientid, 10))
	return clientkillid{cmds: cmds}
}

func (c clientkillipport) Normal() clientkillnormal {
	var cmds []string
	cmds = append(c.cmds, "normal")
	return clientkillnormal{cmds: cmds}
}

func (c clientkillipport) Master() clientkillmaster {
	var cmds []string
	cmds = append(c.cmds, "master")
	return clientkillmaster{cmds: cmds}
}

func (c clientkillipport) Slave() clientkillslave {
	var cmds []string
	cmds = append(c.cmds, "slave")
	return clientkillslave{cmds: cmds}
}

func (c clientkillipport) Pubsub() clientkillpubsub {
	var cmds []string
	cmds = append(c.cmds, "pubsub")
	return clientkillpubsub{cmds: cmds}
}

func (c clientkillipport) User(username string) clientkilluser {
	var cmds []string
	cmds = append(c.cmds, "USER", username)
	return clientkilluser{cmds: cmds}
}

func (c clientkillipport) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkillipport) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkillipport) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkillipport) Build() []string {
	return c.cmds
}

type clientkillladdr struct {
	cmds []string
}

func (c clientkillladdr) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkillladdr) Build() []string {
	return c.cmds
}

type clientkillmaster struct {
	cmds []string
}

func (c clientkillmaster) User(username string) clientkilluser {
	var cmds []string
	cmds = append(c.cmds, "USER", username)
	return clientkilluser{cmds: cmds}
}

func (c clientkillmaster) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkillmaster) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkillmaster) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkillmaster) Build() []string {
	return c.cmds
}

type clientkillnormal struct {
	cmds []string
}

func (c clientkillnormal) User(username string) clientkilluser {
	var cmds []string
	cmds = append(c.cmds, "USER", username)
	return clientkilluser{cmds: cmds}
}

func (c clientkillnormal) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkillnormal) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkillnormal) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkillnormal) Build() []string {
	return c.cmds
}

type clientkillpubsub struct {
	cmds []string
}

func (c clientkillpubsub) User(username string) clientkilluser {
	var cmds []string
	cmds = append(c.cmds, "USER", username)
	return clientkilluser{cmds: cmds}
}

func (c clientkillpubsub) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkillpubsub) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkillpubsub) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkillpubsub) Build() []string {
	return c.cmds
}

type clientkillskipme struct {
	cmds []string
}

func (c clientkillskipme) Build() []string {
	return c.cmds
}

type clientkillslave struct {
	cmds []string
}

func (c clientkillslave) User(username string) clientkilluser {
	var cmds []string
	cmds = append(c.cmds, "USER", username)
	return clientkilluser{cmds: cmds}
}

func (c clientkillslave) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkillslave) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkillslave) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkillslave) Build() []string {
	return c.cmds
}

type clientkilluser struct {
	cmds []string
}

func (c clientkilluser) Addr(ipport string) clientkilladdr {
	var cmds []string
	cmds = append(c.cmds, "ADDR", ipport)
	return clientkilladdr{cmds: cmds}
}

func (c clientkilluser) Laddr(ipport string) clientkillladdr {
	var cmds []string
	cmds = append(c.cmds, "LADDR", ipport)
	return clientkillladdr{cmds: cmds}
}

func (c clientkilluser) Skipme(yesno string) clientkillskipme {
	var cmds []string
	cmds = append(c.cmds, "SKIPME", yesno)
	return clientkillskipme{cmds: cmds}
}

func (c clientkilluser) Build() []string {
	return c.cmds
}

type clientlist struct {
	cmds []string
}

func (c clientlist) Normal() clientlistnormal {
	var cmds []string
	cmds = append(c.cmds, "normal")
	return clientlistnormal{cmds: cmds}
}

func (c clientlist) Master() clientlistmaster {
	var cmds []string
	cmds = append(c.cmds, "master")
	return clientlistmaster{cmds: cmds}
}

func (c clientlist) Replica() clientlistreplica {
	var cmds []string
	cmds = append(c.cmds, "replica")
	return clientlistreplica{cmds: cmds}
}

func (c clientlist) Pubsub() clientlistpubsub {
	var cmds []string
	cmds = append(c.cmds, "pubsub")
	return clientlistpubsub{cmds: cmds}
}

func (c clientlist) Id() clientlistidid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return clientlistidid{cmds: cmds}
}

func ClientList() (c clientlist) {
	c.cmds = append(c.cmds, "CLIENT", "LIST")
	return
}

type clientlistidclientid struct {
	cmds []string
}

func (c clientlistidclientid) Clientid(clientid int64) clientlistidclientid {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(clientid, 10))
	return clientlistidclientid{cmds: cmds}
}

type clientlistidid struct {
	cmds []string
}

func (c clientlistidid) Clientid(clientid int64) clientlistidclientid {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(clientid, 10))
	return clientlistidclientid{cmds: cmds}
}

type clientlistmaster struct {
	cmds []string
}

func (c clientlistmaster) Id() clientlistidid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return clientlistidid{cmds: cmds}
}

type clientlistnormal struct {
	cmds []string
}

func (c clientlistnormal) Id() clientlistidid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return clientlistidid{cmds: cmds}
}

type clientlistpubsub struct {
	cmds []string
}

func (c clientlistpubsub) Id() clientlistidid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return clientlistidid{cmds: cmds}
}

type clientlistreplica struct {
	cmds []string
}

func (c clientlistreplica) Id() clientlistidid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return clientlistidid{cmds: cmds}
}

type clientnoevict struct {
	cmds []string
}

func (c clientnoevict) On() clientnoevictenabledon {
	var cmds []string
	cmds = append(c.cmds, "ON")
	return clientnoevictenabledon{cmds: cmds}
}

func (c clientnoevict) Off() clientnoevictenabledoff {
	var cmds []string
	cmds = append(c.cmds, "OFF")
	return clientnoevictenabledoff{cmds: cmds}
}

func ClientNoevict() (c clientnoevict) {
	c.cmds = append(c.cmds, "CLIENT", "NO-EVICT")
	return
}

type clientnoevictenabledoff struct {
	cmds []string
}

func (c clientnoevictenabledoff) Build() []string {
	return c.cmds
}

type clientnoevictenabledon struct {
	cmds []string
}

func (c clientnoevictenabledon) Build() []string {
	return c.cmds
}

type clientpause struct {
	cmds []string
}

func (c clientpause) Timeout(timeout int64) clientpausetimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(timeout, 10))
	return clientpausetimeout{cmds: cmds}
}

func ClientPause() (c clientpause) {
	c.cmds = append(c.cmds, "CLIENT", "PAUSE")
	return
}

type clientpausemodeall struct {
	cmds []string
}

func (c clientpausemodeall) Build() []string {
	return c.cmds
}

type clientpausemodewrite struct {
	cmds []string
}

func (c clientpausemodewrite) Build() []string {
	return c.cmds
}

type clientpausetimeout struct {
	cmds []string
}

func (c clientpausetimeout) Write() clientpausemodewrite {
	var cmds []string
	cmds = append(c.cmds, "WRITE")
	return clientpausemodewrite{cmds: cmds}
}

func (c clientpausetimeout) All() clientpausemodeall {
	var cmds []string
	cmds = append(c.cmds, "ALL")
	return clientpausemodeall{cmds: cmds}
}

func (c clientpausetimeout) Build() []string {
	return c.cmds
}

type clientreply struct {
	cmds []string
}

func (c clientreply) On() clientreplyreplymodeon {
	var cmds []string
	cmds = append(c.cmds, "ON")
	return clientreplyreplymodeon{cmds: cmds}
}

func (c clientreply) Off() clientreplyreplymodeoff {
	var cmds []string
	cmds = append(c.cmds, "OFF")
	return clientreplyreplymodeoff{cmds: cmds}
}

func (c clientreply) Skip() clientreplyreplymodeskip {
	var cmds []string
	cmds = append(c.cmds, "SKIP")
	return clientreplyreplymodeskip{cmds: cmds}
}

func ClientReply() (c clientreply) {
	c.cmds = append(c.cmds, "CLIENT", "REPLY")
	return
}

type clientreplyreplymodeoff struct {
	cmds []string
}

func (c clientreplyreplymodeoff) Build() []string {
	return c.cmds
}

type clientreplyreplymodeon struct {
	cmds []string
}

func (c clientreplyreplymodeon) Build() []string {
	return c.cmds
}

type clientreplyreplymodeskip struct {
	cmds []string
}

func (c clientreplyreplymodeskip) Build() []string {
	return c.cmds
}

type clientsetname struct {
	cmds []string
}

func (c clientsetname) Connectionname(connectionname string) clientsetnameconnectionname {
	var cmds []string
	cmds = append(c.cmds, connectionname)
	return clientsetnameconnectionname{cmds: cmds}
}

func ClientSetname() (c clientsetname) {
	c.cmds = append(c.cmds, "CLIENT", "SETNAME")
	return
}

type clientsetnameconnectionname struct {
	cmds []string
}

func (c clientsetnameconnectionname) Build() []string {
	return c.cmds
}

type clienttracking struct {
	cmds []string
}

func (c clienttracking) On() clienttrackingstatuson {
	var cmds []string
	cmds = append(c.cmds, "ON")
	return clienttrackingstatuson{cmds: cmds}
}

func (c clienttracking) Off() clienttrackingstatusoff {
	var cmds []string
	cmds = append(c.cmds, "OFF")
	return clienttrackingstatusoff{cmds: cmds}
}

func ClientTracking() (c clienttracking) {
	c.cmds = append(c.cmds, "CLIENT", "TRACKING")
	return
}

type clienttrackingbcastbcast struct {
	cmds []string
}

func (c clienttrackingbcastbcast) Optin() clienttrackingoptinoptin {
	var cmds []string
	cmds = append(c.cmds, "OPTIN")
	return clienttrackingoptinoptin{cmds: cmds}
}

func (c clienttrackingbcastbcast) Optout() clienttrackingoptoutoptout {
	var cmds []string
	cmds = append(c.cmds, "OPTOUT")
	return clienttrackingoptoutoptout{cmds: cmds}
}

func (c clienttrackingbcastbcast) Noloop() clienttrackingnoloopnoloop {
	var cmds []string
	cmds = append(c.cmds, "NOLOOP")
	return clienttrackingnoloopnoloop{cmds: cmds}
}

func (c clienttrackingbcastbcast) Build() []string {
	return c.cmds
}

type clienttrackinginfo struct {
	cmds []string
}

func (c clienttrackinginfo) Build() []string {
	return c.cmds
}

func ClientTrackinginfo() (c clienttrackinginfo) {
	c.cmds = append(c.cmds, "CLIENT", "TRACKINGINFO")
	return
}

type clienttrackingnoloopnoloop struct {
	cmds []string
}

func (c clienttrackingnoloopnoloop) Build() []string {
	return c.cmds
}

type clienttrackingoptinoptin struct {
	cmds []string
}

func (c clienttrackingoptinoptin) Optout() clienttrackingoptoutoptout {
	var cmds []string
	cmds = append(c.cmds, "OPTOUT")
	return clienttrackingoptoutoptout{cmds: cmds}
}

func (c clienttrackingoptinoptin) Noloop() clienttrackingnoloopnoloop {
	var cmds []string
	cmds = append(c.cmds, "NOLOOP")
	return clienttrackingnoloopnoloop{cmds: cmds}
}

func (c clienttrackingoptinoptin) Build() []string {
	return c.cmds
}

type clienttrackingoptoutoptout struct {
	cmds []string
}

func (c clienttrackingoptoutoptout) Noloop() clienttrackingnoloopnoloop {
	var cmds []string
	cmds = append(c.cmds, "NOLOOP")
	return clienttrackingnoloopnoloop{cmds: cmds}
}

func (c clienttrackingoptoutoptout) Build() []string {
	return c.cmds
}

type clienttrackingprefix struct {
	cmds []string
}

func (c clienttrackingprefix) Bcast() clienttrackingbcastbcast {
	var cmds []string
	cmds = append(c.cmds, "BCAST")
	return clienttrackingbcastbcast{cmds: cmds}
}

func (c clienttrackingprefix) Optin() clienttrackingoptinoptin {
	var cmds []string
	cmds = append(c.cmds, "OPTIN")
	return clienttrackingoptinoptin{cmds: cmds}
}

func (c clienttrackingprefix) Optout() clienttrackingoptoutoptout {
	var cmds []string
	cmds = append(c.cmds, "OPTOUT")
	return clienttrackingoptoutoptout{cmds: cmds}
}

func (c clienttrackingprefix) Noloop() clienttrackingnoloopnoloop {
	var cmds []string
	cmds = append(c.cmds, "NOLOOP")
	return clienttrackingnoloopnoloop{cmds: cmds}
}

func (c clienttrackingprefix) Prefix(prefix ...string) clienttrackingprefix {
	var cmds []string
	cmds = append(cmds, prefix...)
	return clienttrackingprefix{cmds: cmds}
}

func (c clienttrackingprefix) Build() []string {
	return c.cmds
}

type clienttrackingredirect struct {
	cmds []string
}

func (c clienttrackingredirect) Prefix(prefix ...string) clienttrackingprefix {
	var cmds []string
	cmds = append(c.cmds, "PREFIX")
	cmds = append(cmds, prefix...)
	return clienttrackingprefix{cmds: cmds}
}

func (c clienttrackingredirect) Bcast() clienttrackingbcastbcast {
	var cmds []string
	cmds = append(c.cmds, "BCAST")
	return clienttrackingbcastbcast{cmds: cmds}
}

func (c clienttrackingredirect) Optin() clienttrackingoptinoptin {
	var cmds []string
	cmds = append(c.cmds, "OPTIN")
	return clienttrackingoptinoptin{cmds: cmds}
}

func (c clienttrackingredirect) Optout() clienttrackingoptoutoptout {
	var cmds []string
	cmds = append(c.cmds, "OPTOUT")
	return clienttrackingoptoutoptout{cmds: cmds}
}

func (c clienttrackingredirect) Noloop() clienttrackingnoloopnoloop {
	var cmds []string
	cmds = append(c.cmds, "NOLOOP")
	return clienttrackingnoloopnoloop{cmds: cmds}
}

func (c clienttrackingredirect) Build() []string {
	return c.cmds
}

type clienttrackingstatusoff struct {
	cmds []string
}

func (c clienttrackingstatusoff) Redirect(clientid int64) clienttrackingredirect {
	var cmds []string
	cmds = append(c.cmds, "REDIRECT", strconv.FormatInt(clientid, 10))
	return clienttrackingredirect{cmds: cmds}
}

func (c clienttrackingstatusoff) Prefix(prefix ...string) clienttrackingprefix {
	var cmds []string
	cmds = append(c.cmds, "PREFIX")
	cmds = append(cmds, prefix...)
	return clienttrackingprefix{cmds: cmds}
}

func (c clienttrackingstatusoff) Bcast() clienttrackingbcastbcast {
	var cmds []string
	cmds = append(c.cmds, "BCAST")
	return clienttrackingbcastbcast{cmds: cmds}
}

func (c clienttrackingstatusoff) Optin() clienttrackingoptinoptin {
	var cmds []string
	cmds = append(c.cmds, "OPTIN")
	return clienttrackingoptinoptin{cmds: cmds}
}

func (c clienttrackingstatusoff) Optout() clienttrackingoptoutoptout {
	var cmds []string
	cmds = append(c.cmds, "OPTOUT")
	return clienttrackingoptoutoptout{cmds: cmds}
}

func (c clienttrackingstatusoff) Noloop() clienttrackingnoloopnoloop {
	var cmds []string
	cmds = append(c.cmds, "NOLOOP")
	return clienttrackingnoloopnoloop{cmds: cmds}
}

func (c clienttrackingstatusoff) Build() []string {
	return c.cmds
}

type clienttrackingstatuson struct {
	cmds []string
}

func (c clienttrackingstatuson) Redirect(clientid int64) clienttrackingredirect {
	var cmds []string
	cmds = append(c.cmds, "REDIRECT", strconv.FormatInt(clientid, 10))
	return clienttrackingredirect{cmds: cmds}
}

func (c clienttrackingstatuson) Prefix(prefix ...string) clienttrackingprefix {
	var cmds []string
	cmds = append(c.cmds, "PREFIX")
	cmds = append(cmds, prefix...)
	return clienttrackingprefix{cmds: cmds}
}

func (c clienttrackingstatuson) Bcast() clienttrackingbcastbcast {
	var cmds []string
	cmds = append(c.cmds, "BCAST")
	return clienttrackingbcastbcast{cmds: cmds}
}

func (c clienttrackingstatuson) Optin() clienttrackingoptinoptin {
	var cmds []string
	cmds = append(c.cmds, "OPTIN")
	return clienttrackingoptinoptin{cmds: cmds}
}

func (c clienttrackingstatuson) Optout() clienttrackingoptoutoptout {
	var cmds []string
	cmds = append(c.cmds, "OPTOUT")
	return clienttrackingoptoutoptout{cmds: cmds}
}

func (c clienttrackingstatuson) Noloop() clienttrackingnoloopnoloop {
	var cmds []string
	cmds = append(c.cmds, "NOLOOP")
	return clienttrackingnoloopnoloop{cmds: cmds}
}

func (c clienttrackingstatuson) Build() []string {
	return c.cmds
}

type clientunblock struct {
	cmds []string
}

func (c clientunblock) Clientid(clientid int64) clientunblockclientid {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(clientid, 10))
	return clientunblockclientid{cmds: cmds}
}

func ClientUnblock() (c clientunblock) {
	c.cmds = append(c.cmds, "CLIENT", "UNBLOCK")
	return
}

type clientunblockclientid struct {
	cmds []string
}

func (c clientunblockclientid) Timeout() clientunblockunblocktypetimeout {
	var cmds []string
	cmds = append(c.cmds, "TIMEOUT")
	return clientunblockunblocktypetimeout{cmds: cmds}
}

func (c clientunblockclientid) Error() clientunblockunblocktypeerror {
	var cmds []string
	cmds = append(c.cmds, "ERROR")
	return clientunblockunblocktypeerror{cmds: cmds}
}

func (c clientunblockclientid) Build() []string {
	return c.cmds
}

type clientunblockunblocktypeerror struct {
	cmds []string
}

func (c clientunblockunblocktypeerror) Build() []string {
	return c.cmds
}

type clientunblockunblocktypetimeout struct {
	cmds []string
}

func (c clientunblockunblocktypetimeout) Build() []string {
	return c.cmds
}

type clientunpause struct {
	cmds []string
}

func (c clientunpause) Build() []string {
	return c.cmds
}

func ClientUnpause() (c clientunpause) {
	c.cmds = append(c.cmds, "CLIENT", "UNPAUSE")
	return
}

type clusteraddslots struct {
	cmds []string
}

func (c clusteraddslots) Slot(slot int64) clusteraddslotsslot {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(slot, 10))
	return clusteraddslotsslot{cmds: cmds}
}

func ClusterAddslots() (c clusteraddslots) {
	c.cmds = append(c.cmds, "CLUSTER", "ADDSLOTS")
	return
}

type clusteraddslotsslot struct {
	cmds []string
}

func (c clusteraddslotsslot) Slot(slot int64) clusteraddslotsslot {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(slot, 10))
	return clusteraddslotsslot{cmds: cmds}
}

type clusterbumpepoch struct {
	cmds []string
}

func (c clusterbumpepoch) Build() []string {
	return c.cmds
}

func ClusterBumpepoch() (c clusterbumpepoch) {
	c.cmds = append(c.cmds, "CLUSTER", "BUMPEPOCH")
	return
}

type clustercountfailurereports struct {
	cmds []string
}

func (c clustercountfailurereports) Nodeid(nodeid string) clustercountfailurereportsnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clustercountfailurereportsnodeid{cmds: cmds}
}

func ClusterCountfailurereports() (c clustercountfailurereports) {
	c.cmds = append(c.cmds, "CLUSTER", "COUNT-FAILURE-REPORTS")
	return
}

type clustercountfailurereportsnodeid struct {
	cmds []string
}

func (c clustercountfailurereportsnodeid) Build() []string {
	return c.cmds
}

type clustercountkeysinslot struct {
	cmds []string
}

func (c clustercountkeysinslot) Slot(slot int64) clustercountkeysinslotslot {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(slot, 10))
	return clustercountkeysinslotslot{cmds: cmds}
}

func ClusterCountkeysinslot() (c clustercountkeysinslot) {
	c.cmds = append(c.cmds, "CLUSTER", "COUNTKEYSINSLOT")
	return
}

type clustercountkeysinslotslot struct {
	cmds []string
}

func (c clustercountkeysinslotslot) Build() []string {
	return c.cmds
}

type clusterdelslots struct {
	cmds []string
}

func (c clusterdelslots) Slot(slot int64) clusterdelslotsslot {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(slot, 10))
	return clusterdelslotsslot{cmds: cmds}
}

func ClusterDelslots() (c clusterdelslots) {
	c.cmds = append(c.cmds, "CLUSTER", "DELSLOTS")
	return
}

type clusterdelslotsslot struct {
	cmds []string
}

func (c clusterdelslotsslot) Slot(slot int64) clusterdelslotsslot {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(slot, 10))
	return clusterdelslotsslot{cmds: cmds}
}

type clusterfailover struct {
	cmds []string
}

func (c clusterfailover) Force() clusterfailoveroptionsforce {
	var cmds []string
	cmds = append(c.cmds, "FORCE")
	return clusterfailoveroptionsforce{cmds: cmds}
}

func (c clusterfailover) Takeover() clusterfailoveroptionstakeover {
	var cmds []string
	cmds = append(c.cmds, "TAKEOVER")
	return clusterfailoveroptionstakeover{cmds: cmds}
}

func (c clusterfailover) Build() []string {
	return c.cmds
}

func ClusterFailover() (c clusterfailover) {
	c.cmds = append(c.cmds, "CLUSTER", "FAILOVER")
	return
}

type clusterfailoveroptionsforce struct {
	cmds []string
}

func (c clusterfailoveroptionsforce) Build() []string {
	return c.cmds
}

type clusterfailoveroptionstakeover struct {
	cmds []string
}

func (c clusterfailoveroptionstakeover) Build() []string {
	return c.cmds
}

type clusterflushslots struct {
	cmds []string
}

func (c clusterflushslots) Build() []string {
	return c.cmds
}

func ClusterFlushslots() (c clusterflushslots) {
	c.cmds = append(c.cmds, "CLUSTER", "FLUSHSLOTS")
	return
}

type clusterforget struct {
	cmds []string
}

func (c clusterforget) Nodeid(nodeid string) clusterforgetnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clusterforgetnodeid{cmds: cmds}
}

func ClusterForget() (c clusterforget) {
	c.cmds = append(c.cmds, "CLUSTER", "FORGET")
	return
}

type clusterforgetnodeid struct {
	cmds []string
}

func (c clusterforgetnodeid) Build() []string {
	return c.cmds
}

type clustergetkeysinslot struct {
	cmds []string
}

func (c clustergetkeysinslot) Slot(slot int64) clustergetkeysinslotslot {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(slot, 10))
	return clustergetkeysinslotslot{cmds: cmds}
}

func ClusterGetkeysinslot() (c clustergetkeysinslot) {
	c.cmds = append(c.cmds, "CLUSTER", "GETKEYSINSLOT")
	return
}

type clustergetkeysinslotcount struct {
	cmds []string
}

func (c clustergetkeysinslotcount) Build() []string {
	return c.cmds
}

type clustergetkeysinslotslot struct {
	cmds []string
}

func (c clustergetkeysinslotslot) Count(count int64) clustergetkeysinslotcount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return clustergetkeysinslotcount{cmds: cmds}
}

type clusterinfo struct {
	cmds []string
}

func (c clusterinfo) Build() []string {
	return c.cmds
}

func ClusterInfo() (c clusterinfo) {
	c.cmds = append(c.cmds, "CLUSTER", "INFO")
	return
}

type clusterkeyslot struct {
	cmds []string
}

func (c clusterkeyslot) Key(key string) clusterkeyslotkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return clusterkeyslotkey{cmds: cmds}
}

func ClusterKeyslot() (c clusterkeyslot) {
	c.cmds = append(c.cmds, "CLUSTER", "KEYSLOT")
	return
}

type clusterkeyslotkey struct {
	cmds []string
}

func (c clusterkeyslotkey) Build() []string {
	return c.cmds
}

type clustermeet struct {
	cmds []string
}

func (c clustermeet) Ip(ip string) clustermeetip {
	var cmds []string
	cmds = append(c.cmds, ip)
	return clustermeetip{cmds: cmds}
}

func ClusterMeet() (c clustermeet) {
	c.cmds = append(c.cmds, "CLUSTER", "MEET")
	return
}

type clustermeetip struct {
	cmds []string
}

func (c clustermeetip) Port(port int64) clustermeetport {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(port, 10))
	return clustermeetport{cmds: cmds}
}

type clustermeetport struct {
	cmds []string
}

func (c clustermeetport) Build() []string {
	return c.cmds
}

type clustermyid struct {
	cmds []string
}

func (c clustermyid) Build() []string {
	return c.cmds
}

func ClusterMyid() (c clustermyid) {
	c.cmds = append(c.cmds, "CLUSTER", "MYID")
	return
}

type clusternodes struct {
	cmds []string
}

func (c clusternodes) Build() []string {
	return c.cmds
}

func ClusterNodes() (c clusternodes) {
	c.cmds = append(c.cmds, "CLUSTER", "NODES")
	return
}

type clusterreplicas struct {
	cmds []string
}

func (c clusterreplicas) Nodeid(nodeid string) clusterreplicasnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clusterreplicasnodeid{cmds: cmds}
}

func ClusterReplicas() (c clusterreplicas) {
	c.cmds = append(c.cmds, "CLUSTER", "REPLICAS")
	return
}

type clusterreplicasnodeid struct {
	cmds []string
}

func (c clusterreplicasnodeid) Build() []string {
	return c.cmds
}

type clusterreplicate struct {
	cmds []string
}

func (c clusterreplicate) Nodeid(nodeid string) clusterreplicatenodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clusterreplicatenodeid{cmds: cmds}
}

func ClusterReplicate() (c clusterreplicate) {
	c.cmds = append(c.cmds, "CLUSTER", "REPLICATE")
	return
}

type clusterreplicatenodeid struct {
	cmds []string
}

func (c clusterreplicatenodeid) Build() []string {
	return c.cmds
}

type clusterreset struct {
	cmds []string
}

func (c clusterreset) Hard() clusterresetresettypehard {
	var cmds []string
	cmds = append(c.cmds, "HARD")
	return clusterresetresettypehard{cmds: cmds}
}

func (c clusterreset) Soft() clusterresetresettypesoft {
	var cmds []string
	cmds = append(c.cmds, "SOFT")
	return clusterresetresettypesoft{cmds: cmds}
}

func (c clusterreset) Build() []string {
	return c.cmds
}

func ClusterReset() (c clusterreset) {
	c.cmds = append(c.cmds, "CLUSTER", "RESET")
	return
}

type clusterresetresettypehard struct {
	cmds []string
}

func (c clusterresetresettypehard) Build() []string {
	return c.cmds
}

type clusterresetresettypesoft struct {
	cmds []string
}

func (c clusterresetresettypesoft) Build() []string {
	return c.cmds
}

type clustersaveconfig struct {
	cmds []string
}

func (c clustersaveconfig) Build() []string {
	return c.cmds
}

func ClusterSaveconfig() (c clustersaveconfig) {
	c.cmds = append(c.cmds, "CLUSTER", "SAVECONFIG")
	return
}

type clustersetconfigepoch struct {
	cmds []string
}

func (c clustersetconfigepoch) Configepoch(configepoch int64) clustersetconfigepochconfigepoch {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(configepoch, 10))
	return clustersetconfigepochconfigepoch{cmds: cmds}
}

func ClusterSetconfigepoch() (c clustersetconfigepoch) {
	c.cmds = append(c.cmds, "CLUSTER", "SET-CONFIG-EPOCH")
	return
}

type clustersetconfigepochconfigepoch struct {
	cmds []string
}

func (c clustersetconfigepochconfigepoch) Build() []string {
	return c.cmds
}

type clustersetslot struct {
	cmds []string
}

func (c clustersetslot) Slot(slot int64) clustersetslotslot {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(slot, 10))
	return clustersetslotslot{cmds: cmds}
}

func ClusterSetslot() (c clustersetslot) {
	c.cmds = append(c.cmds, "CLUSTER", "SETSLOT")
	return
}

type clustersetslotnodeid struct {
	cmds []string
}

func (c clustersetslotnodeid) Build() []string {
	return c.cmds
}

type clustersetslotslot struct {
	cmds []string
}

func (c clustersetslotslot) Importing() clustersetslotsubcommandimporting {
	var cmds []string
	cmds = append(c.cmds, "IMPORTING")
	return clustersetslotsubcommandimporting{cmds: cmds}
}

func (c clustersetslotslot) Migrating() clustersetslotsubcommandmigrating {
	var cmds []string
	cmds = append(c.cmds, "MIGRATING")
	return clustersetslotsubcommandmigrating{cmds: cmds}
}

func (c clustersetslotslot) Stable() clustersetslotsubcommandstable {
	var cmds []string
	cmds = append(c.cmds, "STABLE")
	return clustersetslotsubcommandstable{cmds: cmds}
}

func (c clustersetslotslot) Node() clustersetslotsubcommandnode {
	var cmds []string
	cmds = append(c.cmds, "NODE")
	return clustersetslotsubcommandnode{cmds: cmds}
}

type clustersetslotsubcommandimporting struct {
	cmds []string
}

func (c clustersetslotsubcommandimporting) Nodeid(nodeid string) clustersetslotnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clustersetslotnodeid{cmds: cmds}
}

func (c clustersetslotsubcommandimporting) Build() []string {
	return c.cmds
}

type clustersetslotsubcommandmigrating struct {
	cmds []string
}

func (c clustersetslotsubcommandmigrating) Nodeid(nodeid string) clustersetslotnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clustersetslotnodeid{cmds: cmds}
}

func (c clustersetslotsubcommandmigrating) Build() []string {
	return c.cmds
}

type clustersetslotsubcommandnode struct {
	cmds []string
}

func (c clustersetslotsubcommandnode) Nodeid(nodeid string) clustersetslotnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clustersetslotnodeid{cmds: cmds}
}

func (c clustersetslotsubcommandnode) Build() []string {
	return c.cmds
}

type clustersetslotsubcommandstable struct {
	cmds []string
}

func (c clustersetslotsubcommandstable) Nodeid(nodeid string) clustersetslotnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clustersetslotnodeid{cmds: cmds}
}

func (c clustersetslotsubcommandstable) Build() []string {
	return c.cmds
}

type clusterslaves struct {
	cmds []string
}

func (c clusterslaves) Nodeid(nodeid string) clusterslavesnodeid {
	var cmds []string
	cmds = append(c.cmds, nodeid)
	return clusterslavesnodeid{cmds: cmds}
}

func ClusterSlaves() (c clusterslaves) {
	c.cmds = append(c.cmds, "CLUSTER", "SLAVES")
	return
}

type clusterslavesnodeid struct {
	cmds []string
}

func (c clusterslavesnodeid) Build() []string {
	return c.cmds
}

type clusterslots struct {
	cmds []string
}

func (c clusterslots) Build() []string {
	return c.cmds
}

func ClusterSlots() (c clusterslots) {
	c.cmds = append(c.cmds, "CLUSTER", "SLOTS")
	return
}

type command struct {
	cmds []string
}

func (c command) Build() []string {
	return c.cmds
}

func Command() (c command) {
	c.cmds = append(c.cmds, "COMMAND")
	return
}

type commandcount struct {
	cmds []string
}

func (c commandcount) Build() []string {
	return c.cmds
}

func CommandCount() (c commandcount) {
	c.cmds = append(c.cmds, "COMMAND", "COUNT")
	return
}

type commandgetkeys struct {
	cmds []string
}

func (c commandgetkeys) Build() []string {
	return c.cmds
}

func CommandGetkeys() (c commandgetkeys) {
	c.cmds = append(c.cmds, "COMMAND", "GETKEYS")
	return
}

type commandinfo struct {
	cmds []string
}

func (c commandinfo) Commandname(commandname ...string) commandinfocommandname {
	var cmds []string
	cmds = append(cmds, commandname...)
	return commandinfocommandname{cmds: cmds}
}

func CommandInfo() (c commandinfo) {
	c.cmds = append(c.cmds, "COMMAND", "INFO")
	return
}

type commandinfocommandname struct {
	cmds []string
}

func (c commandinfocommandname) Commandname(commandname ...string) commandinfocommandname {
	var cmds []string
	cmds = append(cmds, commandname...)
	return commandinfocommandname{cmds: cmds}
}

type configget struct {
	cmds []string
}

func (c configget) Parameter(parameter string) configgetparameter {
	var cmds []string
	cmds = append(c.cmds, parameter)
	return configgetparameter{cmds: cmds}
}

func ConfigGet() (c configget) {
	c.cmds = append(c.cmds, "CONFIG", "GET")
	return
}

type configgetparameter struct {
	cmds []string
}

func (c configgetparameter) Build() []string {
	return c.cmds
}

type configresetstat struct {
	cmds []string
}

func (c configresetstat) Build() []string {
	return c.cmds
}

func ConfigResetstat() (c configresetstat) {
	c.cmds = append(c.cmds, "CONFIG", "RESETSTAT")
	return
}

type configrewrite struct {
	cmds []string
}

func (c configrewrite) Build() []string {
	return c.cmds
}

func ConfigRewrite() (c configrewrite) {
	c.cmds = append(c.cmds, "CONFIG", "REWRITE")
	return
}

type configset struct {
	cmds []string
}

func (c configset) Parameter(parameter string) configsetparameter {
	var cmds []string
	cmds = append(c.cmds, parameter)
	return configsetparameter{cmds: cmds}
}

func ConfigSet() (c configset) {
	c.cmds = append(c.cmds, "CONFIG", "SET")
	return
}

type configsetparameter struct {
	cmds []string
}

func (c configsetparameter) Value(value string) configsetvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return configsetvalue{cmds: cmds}
}

type configsetvalue struct {
	cmds []string
}

func (c configsetvalue) Build() []string {
	return c.cmds
}

type copy struct {
	cmds []string
}

func (c copy) Source(source string) copysource {
	var cmds []string
	cmds = append(c.cmds, source)
	return copysource{cmds: cmds}
}

func Copy() (c copy) {
	c.cmds = append(c.cmds, "COPY")
	return
}

type copydb struct {
	cmds []string
}

func (c copydb) Replace() copyreplacereplace {
	var cmds []string
	cmds = append(c.cmds, "REPLACE")
	return copyreplacereplace{cmds: cmds}
}

func (c copydb) Build() []string {
	return c.cmds
}

type copydestination struct {
	cmds []string
}

func (c copydestination) Db(destinationdb int64) copydb {
	var cmds []string
	cmds = append(c.cmds, "DB", strconv.FormatInt(destinationdb, 10))
	return copydb{cmds: cmds}
}

func (c copydestination) Replace() copyreplacereplace {
	var cmds []string
	cmds = append(c.cmds, "REPLACE")
	return copyreplacereplace{cmds: cmds}
}

func (c copydestination) Build() []string {
	return c.cmds
}

type copyreplacereplace struct {
	cmds []string
}

func (c copyreplacereplace) Build() []string {
	return c.cmds
}

type copysource struct {
	cmds []string
}

func (c copysource) Destination(destination string) copydestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return copydestination{cmds: cmds}
}

type dbsize struct {
	cmds []string
}

func (c dbsize) Build() []string {
	return c.cmds
}

func Dbsize() (c dbsize) {
	c.cmds = append(c.cmds, "DBSIZE")
	return
}

type debugobject struct {
	cmds []string
}

func (c debugobject) Key(key string) debugobjectkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return debugobjectkey{cmds: cmds}
}

func DebugObject() (c debugobject) {
	c.cmds = append(c.cmds, "DEBUG", "OBJECT")
	return
}

type debugobjectkey struct {
	cmds []string
}

func (c debugobjectkey) Build() []string {
	return c.cmds
}

type debugsegfault struct {
	cmds []string
}

func (c debugsegfault) Build() []string {
	return c.cmds
}

func DebugSegfault() (c debugsegfault) {
	c.cmds = append(c.cmds, "DEBUG", "SEGFAULT")
	return
}

type decr struct {
	cmds []string
}

func (c decr) Key(key string) decrkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return decrkey{cmds: cmds}
}

func Decr() (c decr) {
	c.cmds = append(c.cmds, "DECR")
	return
}

type decrby struct {
	cmds []string
}

func (c decrby) Key(key string) decrbykey {
	var cmds []string
	cmds = append(c.cmds, key)
	return decrbykey{cmds: cmds}
}

func Decrby() (c decrby) {
	c.cmds = append(c.cmds, "DECRBY")
	return
}

type decrbydecrement struct {
	cmds []string
}

func (c decrbydecrement) Build() []string {
	return c.cmds
}

type decrbykey struct {
	cmds []string
}

func (c decrbykey) Decrement(decrement int64) decrbydecrement {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(decrement, 10))
	return decrbydecrement{cmds: cmds}
}

type decrkey struct {
	cmds []string
}

func (c decrkey) Build() []string {
	return c.cmds
}

type del struct {
	cmds []string
}

func (c del) Key(key ...string) delkey {
	var cmds []string
	cmds = append(cmds, key...)
	return delkey{cmds: cmds}
}

func Del() (c del) {
	c.cmds = append(c.cmds, "DEL")
	return
}

type delkey struct {
	cmds []string
}

func (c delkey) Key(key ...string) delkey {
	var cmds []string
	cmds = append(cmds, key...)
	return delkey{cmds: cmds}
}

type discard struct {
	cmds []string
}

func (c discard) Build() []string {
	return c.cmds
}

func Discard() (c discard) {
	c.cmds = append(c.cmds, "DISCARD")
	return
}

type dump struct {
	cmds []string
}

func (c dump) Key(key string) dumpkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return dumpkey{cmds: cmds}
}

func Dump() (c dump) {
	c.cmds = append(c.cmds, "DUMP")
	return
}

type dumpkey struct {
	cmds []string
}

func (c dumpkey) Build() []string {
	return c.cmds
}

type echo struct {
	cmds []string
}

func (c echo) Message(message string) echomessage {
	var cmds []string
	cmds = append(c.cmds, message)
	return echomessage{cmds: cmds}
}

func Echo() (c echo) {
	c.cmds = append(c.cmds, "ECHO")
	return
}

type echomessage struct {
	cmds []string
}

func (c echomessage) Build() []string {
	return c.cmds
}

type eval struct {
	cmds []string
}

func (c eval) Script(script string) evalscript {
	var cmds []string
	cmds = append(c.cmds, script)
	return evalscript{cmds: cmds}
}

func Eval() (c eval) {
	c.cmds = append(c.cmds, "EVAL")
	return
}

type evalarg struct {
	cmds []string
}

func (c evalarg) Arg(arg ...string) evalarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalarg{cmds: cmds}
}

func (c evalarg) Build() []string {
	return c.cmds
}

type evalkey struct {
	cmds []string
}

func (c evalkey) Arg(arg ...string) evalarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalarg{cmds: cmds}
}

func (c evalkey) Key(key ...string) evalkey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalkey{cmds: cmds}
}

func (c evalkey) Build() []string {
	return c.cmds
}

type evalnumkeys struct {
	cmds []string
}

func (c evalnumkeys) Key(key ...string) evalkey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalkey{cmds: cmds}
}

func (c evalnumkeys) Arg(arg ...string) evalarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalarg{cmds: cmds}
}

func (c evalnumkeys) Build() []string {
	return c.cmds
}

type evalro struct {
	cmds []string
}

func (c evalro) Script(script string) evalroscript {
	var cmds []string
	cmds = append(c.cmds, script)
	return evalroscript{cmds: cmds}
}

func Evalro() (c evalro) {
	c.cmds = append(c.cmds, "EVAL_RO")
	return
}

type evalroarg struct {
	cmds []string
}

func (c evalroarg) Arg(arg ...string) evalroarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalroarg{cmds: cmds}
}

type evalrokey struct {
	cmds []string
}

func (c evalrokey) Arg(arg ...string) evalroarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalroarg{cmds: cmds}
}

func (c evalrokey) Key(key ...string) evalrokey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalrokey{cmds: cmds}
}

type evalronumkeys struct {
	cmds []string
}

func (c evalronumkeys) Key(key ...string) evalrokey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalrokey{cmds: cmds}
}

type evalroscript struct {
	cmds []string
}

func (c evalroscript) Numkeys(numkeys int64) evalronumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return evalronumkeys{cmds: cmds}
}

type evalscript struct {
	cmds []string
}

func (c evalscript) Numkeys(numkeys int64) evalnumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return evalnumkeys{cmds: cmds}
}

type evalsha struct {
	cmds []string
}

func (c evalsha) Sha1(sha1 string) evalshasha1 {
	var cmds []string
	cmds = append(c.cmds, sha1)
	return evalshasha1{cmds: cmds}
}

func Evalsha() (c evalsha) {
	c.cmds = append(c.cmds, "EVALSHA")
	return
}

type evalshaarg struct {
	cmds []string
}

func (c evalshaarg) Arg(arg ...string) evalshaarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalshaarg{cmds: cmds}
}

func (c evalshaarg) Build() []string {
	return c.cmds
}

type evalshakey struct {
	cmds []string
}

func (c evalshakey) Arg(arg ...string) evalshaarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalshaarg{cmds: cmds}
}

func (c evalshakey) Key(key ...string) evalshakey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalshakey{cmds: cmds}
}

func (c evalshakey) Build() []string {
	return c.cmds
}

type evalshanumkeys struct {
	cmds []string
}

func (c evalshanumkeys) Key(key ...string) evalshakey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalshakey{cmds: cmds}
}

func (c evalshanumkeys) Arg(arg ...string) evalshaarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalshaarg{cmds: cmds}
}

func (c evalshanumkeys) Build() []string {
	return c.cmds
}

type evalsharo struct {
	cmds []string
}

func (c evalsharo) Sha1(sha1 string) evalsharosha1 {
	var cmds []string
	cmds = append(c.cmds, sha1)
	return evalsharosha1{cmds: cmds}
}

func Evalsharo() (c evalsharo) {
	c.cmds = append(c.cmds, "EVALSHA_RO")
	return
}

type evalsharoarg struct {
	cmds []string
}

func (c evalsharoarg) Arg(arg ...string) evalsharoarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalsharoarg{cmds: cmds}
}

type evalsharokey struct {
	cmds []string
}

func (c evalsharokey) Arg(arg ...string) evalsharoarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return evalsharoarg{cmds: cmds}
}

func (c evalsharokey) Key(key ...string) evalsharokey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalsharokey{cmds: cmds}
}

type evalsharonumkeys struct {
	cmds []string
}

func (c evalsharonumkeys) Key(key ...string) evalsharokey {
	var cmds []string
	cmds = append(cmds, key...)
	return evalsharokey{cmds: cmds}
}

type evalsharosha1 struct {
	cmds []string
}

func (c evalsharosha1) Numkeys(numkeys int64) evalsharonumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return evalsharonumkeys{cmds: cmds}
}

type evalshasha1 struct {
	cmds []string
}

func (c evalshasha1) Numkeys(numkeys int64) evalshanumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return evalshanumkeys{cmds: cmds}
}

type exec struct {
	cmds []string
}

func (c exec) Build() []string {
	return c.cmds
}

func Exec() (c exec) {
	c.cmds = append(c.cmds, "EXEC")
	return
}

type exists struct {
	cmds []string
}

func (c exists) Key(key ...string) existskey {
	var cmds []string
	cmds = append(cmds, key...)
	return existskey{cmds: cmds}
}

func Exists() (c exists) {
	c.cmds = append(c.cmds, "EXISTS")
	return
}

type existskey struct {
	cmds []string
}

func (c existskey) Key(key ...string) existskey {
	var cmds []string
	cmds = append(cmds, key...)
	return existskey{cmds: cmds}
}

type expire struct {
	cmds []string
}

func (c expire) Key(key string) expirekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return expirekey{cmds: cmds}
}

func Expire() (c expire) {
	c.cmds = append(c.cmds, "EXPIRE")
	return
}

type expireat struct {
	cmds []string
}

func (c expireat) Key(key string) expireatkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return expireatkey{cmds: cmds}
}

func Expireat() (c expireat) {
	c.cmds = append(c.cmds, "EXPIREAT")
	return
}

type expireatconditiongt struct {
	cmds []string
}

func (c expireatconditiongt) Build() []string {
	return c.cmds
}

type expireatconditionlt struct {
	cmds []string
}

func (c expireatconditionlt) Build() []string {
	return c.cmds
}

type expireatconditionnx struct {
	cmds []string
}

func (c expireatconditionnx) Build() []string {
	return c.cmds
}

type expireatconditionxx struct {
	cmds []string
}

func (c expireatconditionxx) Build() []string {
	return c.cmds
}

type expireatkey struct {
	cmds []string
}

func (c expireatkey) Timestamp(timestamp int64) expireattimestamp {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(timestamp, 10))
	return expireattimestamp{cmds: cmds}
}

type expireattimestamp struct {
	cmds []string
}

func (c expireattimestamp) Nx() expireatconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return expireatconditionnx{cmds: cmds}
}

func (c expireattimestamp) Xx() expireatconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return expireatconditionxx{cmds: cmds}
}

func (c expireattimestamp) Gt() expireatconditiongt {
	var cmds []string
	cmds = append(c.cmds, "GT")
	return expireatconditiongt{cmds: cmds}
}

func (c expireattimestamp) Lt() expireatconditionlt {
	var cmds []string
	cmds = append(c.cmds, "LT")
	return expireatconditionlt{cmds: cmds}
}

func (c expireattimestamp) Build() []string {
	return c.cmds
}

type expireconditiongt struct {
	cmds []string
}

func (c expireconditiongt) Build() []string {
	return c.cmds
}

type expireconditionlt struct {
	cmds []string
}

func (c expireconditionlt) Build() []string {
	return c.cmds
}

type expireconditionnx struct {
	cmds []string
}

func (c expireconditionnx) Build() []string {
	return c.cmds
}

type expireconditionxx struct {
	cmds []string
}

func (c expireconditionxx) Build() []string {
	return c.cmds
}

type expirekey struct {
	cmds []string
}

func (c expirekey) Seconds(seconds int64) expireseconds {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(seconds, 10))
	return expireseconds{cmds: cmds}
}

type expireseconds struct {
	cmds []string
}

func (c expireseconds) Nx() expireconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return expireconditionnx{cmds: cmds}
}

func (c expireseconds) Xx() expireconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return expireconditionxx{cmds: cmds}
}

func (c expireseconds) Gt() expireconditiongt {
	var cmds []string
	cmds = append(c.cmds, "GT")
	return expireconditiongt{cmds: cmds}
}

func (c expireseconds) Lt() expireconditionlt {
	var cmds []string
	cmds = append(c.cmds, "LT")
	return expireconditionlt{cmds: cmds}
}

func (c expireseconds) Build() []string {
	return c.cmds
}

type expiretime struct {
	cmds []string
}

func (c expiretime) Key(key string) expiretimekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return expiretimekey{cmds: cmds}
}

func Expiretime() (c expiretime) {
	c.cmds = append(c.cmds, "EXPIRETIME")
	return
}

type expiretimekey struct {
	cmds []string
}

func (c expiretimekey) Build() []string {
	return c.cmds
}

type failover struct {
	cmds []string
}

func (c failover) To() failovertargetto {
	var cmds []string
	cmds = append(c.cmds, "TO")
	return failovertargetto{cmds: cmds}
}

func (c failover) Abort() failoverabort {
	var cmds []string
	cmds = append(c.cmds, "ABORT")
	return failoverabort{cmds: cmds}
}

func (c failover) Timeout(milliseconds int64) failovertimeout {
	var cmds []string
	cmds = append(c.cmds, "TIMEOUT", strconv.FormatInt(milliseconds, 10))
	return failovertimeout{cmds: cmds}
}

func Failover() (c failover) {
	c.cmds = append(c.cmds, "FAILOVER")
	return
}

type failoverabort struct {
	cmds []string
}

func (c failoverabort) Timeout(milliseconds int64) failovertimeout {
	var cmds []string
	cmds = append(c.cmds, "TIMEOUT", strconv.FormatInt(milliseconds, 10))
	return failovertimeout{cmds: cmds}
}

func (c failoverabort) Build() []string {
	return c.cmds
}

type failovertargetforce struct {
	cmds []string
}

func (c failovertargetforce) Abort() failoverabort {
	var cmds []string
	cmds = append(c.cmds, "ABORT")
	return failoverabort{cmds: cmds}
}

func (c failovertargetforce) Timeout(milliseconds int64) failovertimeout {
	var cmds []string
	cmds = append(c.cmds, "TIMEOUT", strconv.FormatInt(milliseconds, 10))
	return failovertimeout{cmds: cmds}
}

func (c failovertargetforce) Build() []string {
	return c.cmds
}

type failovertargethost struct {
	cmds []string
}

func (c failovertargethost) Port(port int64) failovertargetport {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(port, 10))
	return failovertargetport{cmds: cmds}
}

type failovertargetport struct {
	cmds []string
}

func (c failovertargetport) Force() failovertargetforce {
	var cmds []string
	cmds = append(c.cmds, "FORCE")
	return failovertargetforce{cmds: cmds}
}

func (c failovertargetport) Abort() failoverabort {
	var cmds []string
	cmds = append(c.cmds, "ABORT")
	return failoverabort{cmds: cmds}
}

func (c failovertargetport) Timeout(milliseconds int64) failovertimeout {
	var cmds []string
	cmds = append(c.cmds, "TIMEOUT", strconv.FormatInt(milliseconds, 10))
	return failovertimeout{cmds: cmds}
}

func (c failovertargetport) Build() []string {
	return c.cmds
}

type failovertargetto struct {
	cmds []string
}

func (c failovertargetto) Host(host string) failovertargethost {
	var cmds []string
	cmds = append(c.cmds, host)
	return failovertargethost{cmds: cmds}
}

type failovertimeout struct {
	cmds []string
}

func (c failovertimeout) Build() []string {
	return c.cmds
}

type flushall struct {
	cmds []string
}

func (c flushall) Async() flushallasyncasync {
	var cmds []string
	cmds = append(c.cmds, "ASYNC")
	return flushallasyncasync{cmds: cmds}
}

func (c flushall) Sync() flushallasyncsync {
	var cmds []string
	cmds = append(c.cmds, "SYNC")
	return flushallasyncsync{cmds: cmds}
}

func (c flushall) Build() []string {
	return c.cmds
}

func Flushall() (c flushall) {
	c.cmds = append(c.cmds, "FLUSHALL")
	return
}

type flushallasyncasync struct {
	cmds []string
}

func (c flushallasyncasync) Build() []string {
	return c.cmds
}

type flushallasyncsync struct {
	cmds []string
}

func (c flushallasyncsync) Build() []string {
	return c.cmds
}

type flushdb struct {
	cmds []string
}

func (c flushdb) Async() flushdbasyncasync {
	var cmds []string
	cmds = append(c.cmds, "ASYNC")
	return flushdbasyncasync{cmds: cmds}
}

func (c flushdb) Sync() flushdbasyncsync {
	var cmds []string
	cmds = append(c.cmds, "SYNC")
	return flushdbasyncsync{cmds: cmds}
}

func (c flushdb) Build() []string {
	return c.cmds
}

func Flushdb() (c flushdb) {
	c.cmds = append(c.cmds, "FLUSHDB")
	return
}

type flushdbasyncasync struct {
	cmds []string
}

func (c flushdbasyncasync) Build() []string {
	return c.cmds
}

type flushdbasyncsync struct {
	cmds []string
}

func (c flushdbasyncsync) Build() []string {
	return c.cmds
}

type geoadd struct {
	cmds []string
}

func (c geoadd) Key(key string) geoaddkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return geoaddkey{cmds: cmds}
}

func Geoadd() (c geoadd) {
	c.cmds = append(c.cmds, "GEOADD")
	return
}

type geoaddchangech struct {
	cmds []string
}

func (c geoaddchangech) LongitudeLatitudeMember(longitude float64, latitude float64, member string) geoaddlongitudelatitudemember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64), member)
	return geoaddlongitudelatitudemember{cmds: cmds}
}

type geoaddconditionnx struct {
	cmds []string
}

func (c geoaddconditionnx) Ch() geoaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return geoaddchangech{cmds: cmds}
}

func (c geoaddconditionnx) LongitudeLatitudeMember(longitude float64, latitude float64, member string) geoaddlongitudelatitudemember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64), member)
	return geoaddlongitudelatitudemember{cmds: cmds}
}

type geoaddconditionxx struct {
	cmds []string
}

func (c geoaddconditionxx) Ch() geoaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return geoaddchangech{cmds: cmds}
}

func (c geoaddconditionxx) LongitudeLatitudeMember(longitude float64, latitude float64, member string) geoaddlongitudelatitudemember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64), member)
	return geoaddlongitudelatitudemember{cmds: cmds}
}

type geoaddkey struct {
	cmds []string
}

func (c geoaddkey) Nx() geoaddconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return geoaddconditionnx{cmds: cmds}
}

func (c geoaddkey) Xx() geoaddconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return geoaddconditionxx{cmds: cmds}
}

func (c geoaddkey) Ch() geoaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return geoaddchangech{cmds: cmds}
}

func (c geoaddkey) LongitudeLatitudeMember(longitude float64, latitude float64, member string) geoaddlongitudelatitudemember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64), member)
	return geoaddlongitudelatitudemember{cmds: cmds}
}

type geoaddlongitudelatitudemember struct {
	cmds []string
}

func (c geoaddlongitudelatitudemember) LongitudeLatitudeMember(longitude float64, latitude float64, member string) geoaddlongitudelatitudemember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64), member)
	return geoaddlongitudelatitudemember{cmds: cmds}
}

type geodist struct {
	cmds []string
}

func (c geodist) Key(key string) geodistkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return geodistkey{cmds: cmds}
}

func Geodist() (c geodist) {
	c.cmds = append(c.cmds, "GEODIST")
	return
}

type geodistkey struct {
	cmds []string
}

func (c geodistkey) Member1(member1 string) geodistmember1 {
	var cmds []string
	cmds = append(c.cmds, member1)
	return geodistmember1{cmds: cmds}
}

type geodistmember1 struct {
	cmds []string
}

func (c geodistmember1) Member2(member2 string) geodistmember2 {
	var cmds []string
	cmds = append(c.cmds, member2)
	return geodistmember2{cmds: cmds}
}

type geodistmember2 struct {
	cmds []string
}

func (c geodistmember2) M() geodistunitm {
	var cmds []string
	cmds = append(c.cmds, "m")
	return geodistunitm{cmds: cmds}
}

func (c geodistmember2) Km() geodistunitkm {
	var cmds []string
	cmds = append(c.cmds, "km")
	return geodistunitkm{cmds: cmds}
}

func (c geodistmember2) Ft() geodistunitft {
	var cmds []string
	cmds = append(c.cmds, "ft")
	return geodistunitft{cmds: cmds}
}

func (c geodistmember2) Mi() geodistunitmi {
	var cmds []string
	cmds = append(c.cmds, "mi")
	return geodistunitmi{cmds: cmds}
}

func (c geodistmember2) Build() []string {
	return c.cmds
}

type geodistunitft struct {
	cmds []string
}

func (c geodistunitft) Build() []string {
	return c.cmds
}

type geodistunitkm struct {
	cmds []string
}

func (c geodistunitkm) Build() []string {
	return c.cmds
}

type geodistunitm struct {
	cmds []string
}

func (c geodistunitm) Build() []string {
	return c.cmds
}

type geodistunitmi struct {
	cmds []string
}

func (c geodistunitmi) Build() []string {
	return c.cmds
}

type geohash struct {
	cmds []string
}

func (c geohash) Key(key string) geohashkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return geohashkey{cmds: cmds}
}

func Geohash() (c geohash) {
	c.cmds = append(c.cmds, "GEOHASH")
	return
}

type geohashkey struct {
	cmds []string
}

func (c geohashkey) Member(member ...string) geohashmember {
	var cmds []string
	cmds = append(cmds, member...)
	return geohashmember{cmds: cmds}
}

type geohashmember struct {
	cmds []string
}

func (c geohashmember) Member(member ...string) geohashmember {
	var cmds []string
	cmds = append(cmds, member...)
	return geohashmember{cmds: cmds}
}

type geopos struct {
	cmds []string
}

func (c geopos) Key(key string) geoposkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return geoposkey{cmds: cmds}
}

func Geopos() (c geopos) {
	c.cmds = append(c.cmds, "GEOPOS")
	return
}

type geoposkey struct {
	cmds []string
}

func (c geoposkey) Member(member ...string) geoposmember {
	var cmds []string
	cmds = append(cmds, member...)
	return geoposmember{cmds: cmds}
}

type geoposmember struct {
	cmds []string
}

func (c geoposmember) Member(member ...string) geoposmember {
	var cmds []string
	cmds = append(cmds, member...)
	return geoposmember{cmds: cmds}
}

type georadius struct {
	cmds []string
}

func (c georadius) Key(key string) georadiuskey {
	var cmds []string
	cmds = append(c.cmds, key)
	return georadiuskey{cmds: cmds}
}

func Georadius() (c georadius) {
	c.cmds = append(c.cmds, "GEORADIUS")
	return
}

type georadiusbymember struct {
	cmds []string
}

func (c georadiusbymember) Key(key string) georadiusbymemberkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return georadiusbymemberkey{cmds: cmds}
}

func Georadiusbymember() (c georadiusbymember) {
	c.cmds = append(c.cmds, "GEORADIUSBYMEMBER")
	return
}

type georadiusbymembercountanyany struct {
	cmds []string
}

func (c georadiusbymembercountanyany) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymembercountanyany) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymembercountanyany) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymembercountanyany) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

func (c georadiusbymembercountanyany) Build() []string {
	return c.cmds
}

type georadiusbymembercountcount struct {
	cmds []string
}

func (c georadiusbymembercountcount) Any() georadiusbymembercountanyany {
	var cmds []string
	cmds = append(c.cmds, "ANY")
	return georadiusbymembercountanyany{cmds: cmds}
}

func (c georadiusbymembercountcount) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymembercountcount) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymembercountcount) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymembercountcount) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

func (c georadiusbymembercountcount) Build() []string {
	return c.cmds
}

type georadiusbymemberkey struct {
	cmds []string
}

func (c georadiusbymemberkey) Member(member string) georadiusbymembermember {
	var cmds []string
	cmds = append(c.cmds, member)
	return georadiusbymembermember{cmds: cmds}
}

type georadiusbymembermember struct {
	cmds []string
}

func (c georadiusbymembermember) Radius(radius float64) georadiusbymemberradius {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(radius, 'f', -1, 64))
	return georadiusbymemberradius{cmds: cmds}
}

type georadiusbymemberorderasc struct {
	cmds []string
}

func (c georadiusbymemberorderasc) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberorderasc) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

func (c georadiusbymemberorderasc) Build() []string {
	return c.cmds
}

type georadiusbymemberorderdesc struct {
	cmds []string
}

func (c georadiusbymemberorderdesc) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberorderdesc) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

func (c georadiusbymemberorderdesc) Build() []string {
	return c.cmds
}

type georadiusbymemberradius struct {
	cmds []string
}

func (c georadiusbymemberradius) M() georadiusbymemberunitm {
	var cmds []string
	cmds = append(c.cmds, "m")
	return georadiusbymemberunitm{cmds: cmds}
}

func (c georadiusbymemberradius) Km() georadiusbymemberunitkm {
	var cmds []string
	cmds = append(c.cmds, "km")
	return georadiusbymemberunitkm{cmds: cmds}
}

func (c georadiusbymemberradius) Ft() georadiusbymemberunitft {
	var cmds []string
	cmds = append(c.cmds, "ft")
	return georadiusbymemberunitft{cmds: cmds}
}

func (c georadiusbymemberradius) Mi() georadiusbymemberunitmi {
	var cmds []string
	cmds = append(c.cmds, "mi")
	return georadiusbymemberunitmi{cmds: cmds}
}

type georadiusbymemberstore struct {
	cmds []string
}

func (c georadiusbymemberstore) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

func (c georadiusbymemberstore) Build() []string {
	return c.cmds
}

type georadiusbymemberstoredist struct {
	cmds []string
}

func (c georadiusbymemberstoredist) Build() []string {
	return c.cmds
}

type georadiusbymemberunitft struct {
	cmds []string
}

func (c georadiusbymemberunitft) Withcoord() georadiusbymemberwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiusbymemberwithcoordwithcoord{cmds: cmds}
}

func (c georadiusbymemberunitft) Withdist() georadiusbymemberwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiusbymemberwithdistwithdist{cmds: cmds}
}

func (c georadiusbymemberunitft) Withhash() georadiusbymemberwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiusbymemberwithhashwithhash{cmds: cmds}
}

func (c georadiusbymemberunitft) Count(count int64) georadiusbymembercountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiusbymembercountcount{cmds: cmds}
}

func (c georadiusbymemberunitft) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymemberunitft) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymemberunitft) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberunitft) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

type georadiusbymemberunitkm struct {
	cmds []string
}

func (c georadiusbymemberunitkm) Withcoord() georadiusbymemberwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiusbymemberwithcoordwithcoord{cmds: cmds}
}

func (c georadiusbymemberunitkm) Withdist() georadiusbymemberwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiusbymemberwithdistwithdist{cmds: cmds}
}

func (c georadiusbymemberunitkm) Withhash() georadiusbymemberwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiusbymemberwithhashwithhash{cmds: cmds}
}

func (c georadiusbymemberunitkm) Count(count int64) georadiusbymembercountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiusbymembercountcount{cmds: cmds}
}

func (c georadiusbymemberunitkm) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymemberunitkm) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymemberunitkm) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberunitkm) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

type georadiusbymemberunitm struct {
	cmds []string
}

func (c georadiusbymemberunitm) Withcoord() georadiusbymemberwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiusbymemberwithcoordwithcoord{cmds: cmds}
}

func (c georadiusbymemberunitm) Withdist() georadiusbymemberwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiusbymemberwithdistwithdist{cmds: cmds}
}

func (c georadiusbymemberunitm) Withhash() georadiusbymemberwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiusbymemberwithhashwithhash{cmds: cmds}
}

func (c georadiusbymemberunitm) Count(count int64) georadiusbymembercountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiusbymembercountcount{cmds: cmds}
}

func (c georadiusbymemberunitm) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymemberunitm) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymemberunitm) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberunitm) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

type georadiusbymemberunitmi struct {
	cmds []string
}

func (c georadiusbymemberunitmi) Withcoord() georadiusbymemberwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiusbymemberwithcoordwithcoord{cmds: cmds}
}

func (c georadiusbymemberunitmi) Withdist() georadiusbymemberwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiusbymemberwithdistwithdist{cmds: cmds}
}

func (c georadiusbymemberunitmi) Withhash() georadiusbymemberwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiusbymemberwithhashwithhash{cmds: cmds}
}

func (c georadiusbymemberunitmi) Count(count int64) georadiusbymembercountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiusbymembercountcount{cmds: cmds}
}

func (c georadiusbymemberunitmi) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymemberunitmi) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymemberunitmi) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberunitmi) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

type georadiusbymemberwithcoordwithcoord struct {
	cmds []string
}

func (c georadiusbymemberwithcoordwithcoord) Withdist() georadiusbymemberwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiusbymemberwithdistwithdist{cmds: cmds}
}

func (c georadiusbymemberwithcoordwithcoord) Withhash() georadiusbymemberwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiusbymemberwithhashwithhash{cmds: cmds}
}

func (c georadiusbymemberwithcoordwithcoord) Count(count int64) georadiusbymembercountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiusbymembercountcount{cmds: cmds}
}

func (c georadiusbymemberwithcoordwithcoord) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymemberwithcoordwithcoord) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymemberwithcoordwithcoord) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberwithcoordwithcoord) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

type georadiusbymemberwithdistwithdist struct {
	cmds []string
}

func (c georadiusbymemberwithdistwithdist) Withhash() georadiusbymemberwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiusbymemberwithhashwithhash{cmds: cmds}
}

func (c georadiusbymemberwithdistwithdist) Count(count int64) georadiusbymembercountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiusbymembercountcount{cmds: cmds}
}

func (c georadiusbymemberwithdistwithdist) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymemberwithdistwithdist) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymemberwithdistwithdist) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberwithdistwithdist) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

type georadiusbymemberwithhashwithhash struct {
	cmds []string
}

func (c georadiusbymemberwithhashwithhash) Count(count int64) georadiusbymembercountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiusbymembercountcount{cmds: cmds}
}

func (c georadiusbymemberwithhashwithhash) Asc() georadiusbymemberorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusbymemberorderasc{cmds: cmds}
}

func (c georadiusbymemberwithhashwithhash) Desc() georadiusbymemberorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusbymemberorderdesc{cmds: cmds}
}

func (c georadiusbymemberwithhashwithhash) Store(key string) georadiusbymemberstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusbymemberstore{cmds: cmds}
}

func (c georadiusbymemberwithhashwithhash) Storedist(key string) georadiusbymemberstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusbymemberstoredist{cmds: cmds}
}

type georadiuscountanyany struct {
	cmds []string
}

func (c georadiuscountanyany) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiuscountanyany) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiuscountanyany) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiuscountanyany) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

func (c georadiuscountanyany) Build() []string {
	return c.cmds
}

type georadiuscountcount struct {
	cmds []string
}

func (c georadiuscountcount) Any() georadiuscountanyany {
	var cmds []string
	cmds = append(c.cmds, "ANY")
	return georadiuscountanyany{cmds: cmds}
}

func (c georadiuscountcount) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiuscountcount) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiuscountcount) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiuscountcount) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

func (c georadiuscountcount) Build() []string {
	return c.cmds
}

type georadiuskey struct {
	cmds []string
}

func (c georadiuskey) Longitude(longitude float64) georadiuslongitude {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(longitude, 'f', -1, 64))
	return georadiuslongitude{cmds: cmds}
}

type georadiuslatitude struct {
	cmds []string
}

func (c georadiuslatitude) Radius(radius float64) georadiusradius {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(radius, 'f', -1, 64))
	return georadiusradius{cmds: cmds}
}

type georadiuslongitude struct {
	cmds []string
}

func (c georadiuslongitude) Latitude(latitude float64) georadiuslatitude {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(latitude, 'f', -1, 64))
	return georadiuslatitude{cmds: cmds}
}

type georadiusorderasc struct {
	cmds []string
}

func (c georadiusorderasc) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiusorderasc) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

func (c georadiusorderasc) Build() []string {
	return c.cmds
}

type georadiusorderdesc struct {
	cmds []string
}

func (c georadiusorderdesc) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiusorderdesc) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

func (c georadiusorderdesc) Build() []string {
	return c.cmds
}

type georadiusradius struct {
	cmds []string
}

func (c georadiusradius) M() georadiusunitm {
	var cmds []string
	cmds = append(c.cmds, "m")
	return georadiusunitm{cmds: cmds}
}

func (c georadiusradius) Km() georadiusunitkm {
	var cmds []string
	cmds = append(c.cmds, "km")
	return georadiusunitkm{cmds: cmds}
}

func (c georadiusradius) Ft() georadiusunitft {
	var cmds []string
	cmds = append(c.cmds, "ft")
	return georadiusunitft{cmds: cmds}
}

func (c georadiusradius) Mi() georadiusunitmi {
	var cmds []string
	cmds = append(c.cmds, "mi")
	return georadiusunitmi{cmds: cmds}
}

type georadiusstore struct {
	cmds []string
}

func (c georadiusstore) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

func (c georadiusstore) Build() []string {
	return c.cmds
}

type georadiusstoredist struct {
	cmds []string
}

func (c georadiusstoredist) Build() []string {
	return c.cmds
}

type georadiusunitft struct {
	cmds []string
}

func (c georadiusunitft) Withcoord() georadiuswithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiuswithcoordwithcoord{cmds: cmds}
}

func (c georadiusunitft) Withdist() georadiuswithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiuswithdistwithdist{cmds: cmds}
}

func (c georadiusunitft) Withhash() georadiuswithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiuswithhashwithhash{cmds: cmds}
}

func (c georadiusunitft) Count(count int64) georadiuscountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiuscountcount{cmds: cmds}
}

func (c georadiusunitft) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiusunitft) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiusunitft) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiusunitft) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

type georadiusunitkm struct {
	cmds []string
}

func (c georadiusunitkm) Withcoord() georadiuswithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiuswithcoordwithcoord{cmds: cmds}
}

func (c georadiusunitkm) Withdist() georadiuswithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiuswithdistwithdist{cmds: cmds}
}

func (c georadiusunitkm) Withhash() georadiuswithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiuswithhashwithhash{cmds: cmds}
}

func (c georadiusunitkm) Count(count int64) georadiuscountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiuscountcount{cmds: cmds}
}

func (c georadiusunitkm) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiusunitkm) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiusunitkm) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiusunitkm) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

type georadiusunitm struct {
	cmds []string
}

func (c georadiusunitm) Withcoord() georadiuswithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiuswithcoordwithcoord{cmds: cmds}
}

func (c georadiusunitm) Withdist() georadiuswithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiuswithdistwithdist{cmds: cmds}
}

func (c georadiusunitm) Withhash() georadiuswithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiuswithhashwithhash{cmds: cmds}
}

func (c georadiusunitm) Count(count int64) georadiuscountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiuscountcount{cmds: cmds}
}

func (c georadiusunitm) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiusunitm) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiusunitm) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiusunitm) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

type georadiusunitmi struct {
	cmds []string
}

func (c georadiusunitmi) Withcoord() georadiuswithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return georadiuswithcoordwithcoord{cmds: cmds}
}

func (c georadiusunitmi) Withdist() georadiuswithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiuswithdistwithdist{cmds: cmds}
}

func (c georadiusunitmi) Withhash() georadiuswithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiuswithhashwithhash{cmds: cmds}
}

func (c georadiusunitmi) Count(count int64) georadiuscountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiuscountcount{cmds: cmds}
}

func (c georadiusunitmi) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiusunitmi) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiusunitmi) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiusunitmi) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

type georadiuswithcoordwithcoord struct {
	cmds []string
}

func (c georadiuswithcoordwithcoord) Withdist() georadiuswithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return georadiuswithdistwithdist{cmds: cmds}
}

func (c georadiuswithcoordwithcoord) Withhash() georadiuswithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiuswithhashwithhash{cmds: cmds}
}

func (c georadiuswithcoordwithcoord) Count(count int64) georadiuscountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiuscountcount{cmds: cmds}
}

func (c georadiuswithcoordwithcoord) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiuswithcoordwithcoord) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiuswithcoordwithcoord) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiuswithcoordwithcoord) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

type georadiuswithdistwithdist struct {
	cmds []string
}

func (c georadiuswithdistwithdist) Withhash() georadiuswithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return georadiuswithhashwithhash{cmds: cmds}
}

func (c georadiuswithdistwithdist) Count(count int64) georadiuscountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiuscountcount{cmds: cmds}
}

func (c georadiuswithdistwithdist) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiuswithdistwithdist) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiuswithdistwithdist) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiuswithdistwithdist) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

type georadiuswithhashwithhash struct {
	cmds []string
}

func (c georadiuswithhashwithhash) Count(count int64) georadiuscountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return georadiuscountcount{cmds: cmds}
}

func (c georadiuswithhashwithhash) Asc() georadiusorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return georadiusorderasc{cmds: cmds}
}

func (c georadiuswithhashwithhash) Desc() georadiusorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return georadiusorderdesc{cmds: cmds}
}

func (c georadiuswithhashwithhash) Store(key string) georadiusstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", key)
	return georadiusstore{cmds: cmds}
}

func (c georadiuswithhashwithhash) Storedist(key string) georadiusstoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST", key)
	return georadiusstoredist{cmds: cmds}
}

type geosearch struct {
	cmds []string
}

func (c geosearch) Key(key string) geosearchkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return geosearchkey{cmds: cmds}
}

func Geosearch() (c geosearch) {
	c.cmds = append(c.cmds, "GEOSEARCH")
	return
}

type geosearchboxbybox struct {
	cmds []string
}

func (c geosearchboxbybox) Height(height float64) geosearchboxheight {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(height, 'f', -1, 64))
	return geosearchboxheight{cmds: cmds}
}

type geosearchboxheight struct {
	cmds []string
}

func (c geosearchboxheight) M() geosearchboxunitm {
	var cmds []string
	cmds = append(c.cmds, "m")
	return geosearchboxunitm{cmds: cmds}
}

func (c geosearchboxheight) Km() geosearchboxunitkm {
	var cmds []string
	cmds = append(c.cmds, "km")
	return geosearchboxunitkm{cmds: cmds}
}

func (c geosearchboxheight) Ft() geosearchboxunitft {
	var cmds []string
	cmds = append(c.cmds, "ft")
	return geosearchboxunitft{cmds: cmds}
}

func (c geosearchboxheight) Mi() geosearchboxunitmi {
	var cmds []string
	cmds = append(c.cmds, "mi")
	return geosearchboxunitmi{cmds: cmds}
}

type geosearchboxunitft struct {
	cmds []string
}

func (c geosearchboxunitft) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchboxunitft) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchboxunitft) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchboxunitft) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchboxunitft) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchboxunitft) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchboxunitkm struct {
	cmds []string
}

func (c geosearchboxunitkm) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchboxunitkm) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchboxunitkm) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchboxunitkm) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchboxunitkm) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchboxunitkm) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchboxunitm struct {
	cmds []string
}

func (c geosearchboxunitm) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchboxunitm) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchboxunitm) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchboxunitm) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchboxunitm) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchboxunitm) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchboxunitmi struct {
	cmds []string
}

func (c geosearchboxunitmi) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchboxunitmi) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchboxunitmi) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchboxunitmi) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchboxunitmi) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchboxunitmi) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchcirclebyradius struct {
	cmds []string
}

func (c geosearchcirclebyradius) M() geosearchcircleunitm {
	var cmds []string
	cmds = append(c.cmds, "m")
	return geosearchcircleunitm{cmds: cmds}
}

func (c geosearchcirclebyradius) Km() geosearchcircleunitkm {
	var cmds []string
	cmds = append(c.cmds, "km")
	return geosearchcircleunitkm{cmds: cmds}
}

func (c geosearchcirclebyradius) Ft() geosearchcircleunitft {
	var cmds []string
	cmds = append(c.cmds, "ft")
	return geosearchcircleunitft{cmds: cmds}
}

func (c geosearchcirclebyradius) Mi() geosearchcircleunitmi {
	var cmds []string
	cmds = append(c.cmds, "mi")
	return geosearchcircleunitmi{cmds: cmds}
}

type geosearchcircleunitft struct {
	cmds []string
}

func (c geosearchcircleunitft) Bybox(width float64) geosearchboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchboxbybox{cmds: cmds}
}

func (c geosearchcircleunitft) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchcircleunitft) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchcircleunitft) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchcircleunitft) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchcircleunitft) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchcircleunitft) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchcircleunitkm struct {
	cmds []string
}

func (c geosearchcircleunitkm) Bybox(width float64) geosearchboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchboxbybox{cmds: cmds}
}

func (c geosearchcircleunitkm) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchcircleunitkm) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchcircleunitkm) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchcircleunitkm) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchcircleunitkm) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchcircleunitkm) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchcircleunitm struct {
	cmds []string
}

func (c geosearchcircleunitm) Bybox(width float64) geosearchboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchboxbybox{cmds: cmds}
}

func (c geosearchcircleunitm) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchcircleunitm) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchcircleunitm) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchcircleunitm) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchcircleunitm) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchcircleunitm) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchcircleunitmi struct {
	cmds []string
}

func (c geosearchcircleunitmi) Bybox(width float64) geosearchboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchboxbybox{cmds: cmds}
}

func (c geosearchcircleunitmi) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchcircleunitmi) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchcircleunitmi) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchcircleunitmi) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchcircleunitmi) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchcircleunitmi) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchcountanyany struct {
	cmds []string
}

func (c geosearchcountanyany) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchcountanyany) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchcountanyany) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

func (c geosearchcountanyany) Build() []string {
	return c.cmds
}

type geosearchcountcount struct {
	cmds []string
}

func (c geosearchcountcount) Any() geosearchcountanyany {
	var cmds []string
	cmds = append(c.cmds, "ANY")
	return geosearchcountanyany{cmds: cmds}
}

func (c geosearchcountcount) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchcountcount) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchcountcount) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

func (c geosearchcountcount) Build() []string {
	return c.cmds
}

type geosearchfromlonlat struct {
	cmds []string
}

func (c geosearchfromlonlat) Byradius(radius float64) geosearchcirclebyradius {
	var cmds []string
	cmds = append(c.cmds, "BYRADIUS", strconv.FormatFloat(radius, 'f', -1, 64))
	return geosearchcirclebyradius{cmds: cmds}
}

func (c geosearchfromlonlat) Bybox(width float64) geosearchboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchboxbybox{cmds: cmds}
}

func (c geosearchfromlonlat) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchfromlonlat) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchfromlonlat) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchfromlonlat) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchfromlonlat) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchfromlonlat) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchfrommember struct {
	cmds []string
}

func (c geosearchfrommember) Fromlonlat(longitude float64, latitude float64) geosearchfromlonlat {
	var cmds []string
	cmds = append(c.cmds, "FROMLONLAT", strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	return geosearchfromlonlat{cmds: cmds}
}

func (c geosearchfrommember) Byradius(radius float64) geosearchcirclebyradius {
	var cmds []string
	cmds = append(c.cmds, "BYRADIUS", strconv.FormatFloat(radius, 'f', -1, 64))
	return geosearchcirclebyradius{cmds: cmds}
}

func (c geosearchfrommember) Bybox(width float64) geosearchboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchboxbybox{cmds: cmds}
}

func (c geosearchfrommember) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchfrommember) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchfrommember) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchfrommember) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchfrommember) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchfrommember) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchkey struct {
	cmds []string
}

func (c geosearchkey) Frommember(member string) geosearchfrommember {
	var cmds []string
	cmds = append(c.cmds, "FROMMEMBER", member)
	return geosearchfrommember{cmds: cmds}
}

func (c geosearchkey) Fromlonlat(longitude float64, latitude float64) geosearchfromlonlat {
	var cmds []string
	cmds = append(c.cmds, "FROMLONLAT", strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	return geosearchfromlonlat{cmds: cmds}
}

func (c geosearchkey) Byradius(radius float64) geosearchcirclebyradius {
	var cmds []string
	cmds = append(c.cmds, "BYRADIUS", strconv.FormatFloat(radius, 'f', -1, 64))
	return geosearchcirclebyradius{cmds: cmds}
}

func (c geosearchkey) Bybox(width float64) geosearchboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchboxbybox{cmds: cmds}
}

func (c geosearchkey) Asc() geosearchorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchorderasc{cmds: cmds}
}

func (c geosearchkey) Desc() geosearchorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchorderdesc{cmds: cmds}
}

func (c geosearchkey) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchkey) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchkey) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchkey) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchorderasc struct {
	cmds []string
}

func (c geosearchorderasc) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchorderasc) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchorderasc) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchorderasc) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchorderdesc struct {
	cmds []string
}

func (c geosearchorderdesc) Count(count int64) geosearchcountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchcountcount{cmds: cmds}
}

func (c geosearchorderdesc) Withcoord() geosearchwithcoordwithcoord {
	var cmds []string
	cmds = append(c.cmds, "WITHCOORD")
	return geosearchwithcoordwithcoord{cmds: cmds}
}

func (c geosearchorderdesc) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchorderdesc) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

type geosearchstore struct {
	cmds []string
}

func (c geosearchstore) Destination(destination string) geosearchstoredestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return geosearchstoredestination{cmds: cmds}
}

func Geosearchstore() (c geosearchstore) {
	c.cmds = append(c.cmds, "GEOSEARCHSTORE")
	return
}

type geosearchstoreboxbybox struct {
	cmds []string
}

func (c geosearchstoreboxbybox) Height(height float64) geosearchstoreboxheight {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(height, 'f', -1, 64))
	return geosearchstoreboxheight{cmds: cmds}
}

type geosearchstoreboxheight struct {
	cmds []string
}

func (c geosearchstoreboxheight) M() geosearchstoreboxunitm {
	var cmds []string
	cmds = append(c.cmds, "m")
	return geosearchstoreboxunitm{cmds: cmds}
}

func (c geosearchstoreboxheight) Km() geosearchstoreboxunitkm {
	var cmds []string
	cmds = append(c.cmds, "km")
	return geosearchstoreboxunitkm{cmds: cmds}
}

func (c geosearchstoreboxheight) Ft() geosearchstoreboxunitft {
	var cmds []string
	cmds = append(c.cmds, "ft")
	return geosearchstoreboxunitft{cmds: cmds}
}

func (c geosearchstoreboxheight) Mi() geosearchstoreboxunitmi {
	var cmds []string
	cmds = append(c.cmds, "mi")
	return geosearchstoreboxunitmi{cmds: cmds}
}

type geosearchstoreboxunitft struct {
	cmds []string
}

func (c geosearchstoreboxunitft) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstoreboxunitft) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstoreboxunitft) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstoreboxunitft) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstoreboxunitkm struct {
	cmds []string
}

func (c geosearchstoreboxunitkm) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstoreboxunitkm) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstoreboxunitkm) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstoreboxunitkm) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstoreboxunitm struct {
	cmds []string
}

func (c geosearchstoreboxunitm) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstoreboxunitm) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstoreboxunitm) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstoreboxunitm) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstoreboxunitmi struct {
	cmds []string
}

func (c geosearchstoreboxunitmi) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstoreboxunitmi) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstoreboxunitmi) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstoreboxunitmi) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstorecirclebyradius struct {
	cmds []string
}

func (c geosearchstorecirclebyradius) M() geosearchstorecircleunitm {
	var cmds []string
	cmds = append(c.cmds, "m")
	return geosearchstorecircleunitm{cmds: cmds}
}

func (c geosearchstorecirclebyradius) Km() geosearchstorecircleunitkm {
	var cmds []string
	cmds = append(c.cmds, "km")
	return geosearchstorecircleunitkm{cmds: cmds}
}

func (c geosearchstorecirclebyradius) Ft() geosearchstorecircleunitft {
	var cmds []string
	cmds = append(c.cmds, "ft")
	return geosearchstorecircleunitft{cmds: cmds}
}

func (c geosearchstorecirclebyradius) Mi() geosearchstorecircleunitmi {
	var cmds []string
	cmds = append(c.cmds, "mi")
	return geosearchstorecircleunitmi{cmds: cmds}
}

type geosearchstorecircleunitft struct {
	cmds []string
}

func (c geosearchstorecircleunitft) Bybox(width float64) geosearchstoreboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchstoreboxbybox{cmds: cmds}
}

func (c geosearchstorecircleunitft) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstorecircleunitft) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstorecircleunitft) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstorecircleunitft) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstorecircleunitkm struct {
	cmds []string
}

func (c geosearchstorecircleunitkm) Bybox(width float64) geosearchstoreboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchstoreboxbybox{cmds: cmds}
}

func (c geosearchstorecircleunitkm) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstorecircleunitkm) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstorecircleunitkm) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstorecircleunitkm) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstorecircleunitm struct {
	cmds []string
}

func (c geosearchstorecircleunitm) Bybox(width float64) geosearchstoreboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchstoreboxbybox{cmds: cmds}
}

func (c geosearchstorecircleunitm) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstorecircleunitm) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstorecircleunitm) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstorecircleunitm) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstorecircleunitmi struct {
	cmds []string
}

func (c geosearchstorecircleunitmi) Bybox(width float64) geosearchstoreboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchstoreboxbybox{cmds: cmds}
}

func (c geosearchstorecircleunitmi) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstorecircleunitmi) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstorecircleunitmi) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstorecircleunitmi) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstorecountanyany struct {
	cmds []string
}

func (c geosearchstorecountanyany) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

func (c geosearchstorecountanyany) Build() []string {
	return c.cmds
}

type geosearchstorecountcount struct {
	cmds []string
}

func (c geosearchstorecountcount) Any() geosearchstorecountanyany {
	var cmds []string
	cmds = append(c.cmds, "ANY")
	return geosearchstorecountanyany{cmds: cmds}
}

func (c geosearchstorecountcount) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

func (c geosearchstorecountcount) Build() []string {
	return c.cmds
}

type geosearchstoredestination struct {
	cmds []string
}

func (c geosearchstoredestination) Source(source string) geosearchstoresource {
	var cmds []string
	cmds = append(c.cmds, source)
	return geosearchstoresource{cmds: cmds}
}

type geosearchstorefromlonlat struct {
	cmds []string
}

func (c geosearchstorefromlonlat) Byradius(radius float64) geosearchstorecirclebyradius {
	var cmds []string
	cmds = append(c.cmds, "BYRADIUS", strconv.FormatFloat(radius, 'f', -1, 64))
	return geosearchstorecirclebyradius{cmds: cmds}
}

func (c geosearchstorefromlonlat) Bybox(width float64) geosearchstoreboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchstoreboxbybox{cmds: cmds}
}

func (c geosearchstorefromlonlat) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstorefromlonlat) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstorefromlonlat) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstorefromlonlat) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstorefrommember struct {
	cmds []string
}

func (c geosearchstorefrommember) Fromlonlat(longitude float64, latitude float64) geosearchstorefromlonlat {
	var cmds []string
	cmds = append(c.cmds, "FROMLONLAT", strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	return geosearchstorefromlonlat{cmds: cmds}
}

func (c geosearchstorefrommember) Byradius(radius float64) geosearchstorecirclebyradius {
	var cmds []string
	cmds = append(c.cmds, "BYRADIUS", strconv.FormatFloat(radius, 'f', -1, 64))
	return geosearchstorecirclebyradius{cmds: cmds}
}

func (c geosearchstorefrommember) Bybox(width float64) geosearchstoreboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchstoreboxbybox{cmds: cmds}
}

func (c geosearchstorefrommember) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstorefrommember) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstorefrommember) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstorefrommember) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstoreorderasc struct {
	cmds []string
}

func (c geosearchstoreorderasc) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstoreorderasc) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstoreorderdesc struct {
	cmds []string
}

func (c geosearchstoreorderdesc) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstoreorderdesc) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstoresource struct {
	cmds []string
}

func (c geosearchstoresource) Frommember(member string) geosearchstorefrommember {
	var cmds []string
	cmds = append(c.cmds, "FROMMEMBER", member)
	return geosearchstorefrommember{cmds: cmds}
}

func (c geosearchstoresource) Fromlonlat(longitude float64, latitude float64) geosearchstorefromlonlat {
	var cmds []string
	cmds = append(c.cmds, "FROMLONLAT", strconv.FormatFloat(longitude, 'f', -1, 64), strconv.FormatFloat(latitude, 'f', -1, 64))
	return geosearchstorefromlonlat{cmds: cmds}
}

func (c geosearchstoresource) Byradius(radius float64) geosearchstorecirclebyradius {
	var cmds []string
	cmds = append(c.cmds, "BYRADIUS", strconv.FormatFloat(radius, 'f', -1, 64))
	return geosearchstorecirclebyradius{cmds: cmds}
}

func (c geosearchstoresource) Bybox(width float64) geosearchstoreboxbybox {
	var cmds []string
	cmds = append(c.cmds, "BYBOX", strconv.FormatFloat(width, 'f', -1, 64))
	return geosearchstoreboxbybox{cmds: cmds}
}

func (c geosearchstoresource) Asc() geosearchstoreorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return geosearchstoreorderasc{cmds: cmds}
}

func (c geosearchstoresource) Desc() geosearchstoreorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return geosearchstoreorderdesc{cmds: cmds}
}

func (c geosearchstoresource) Count(count int64) geosearchstorecountcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return geosearchstorecountcount{cmds: cmds}
}

func (c geosearchstoresource) Storedist() geosearchstorestorediststoredist {
	var cmds []string
	cmds = append(c.cmds, "STOREDIST")
	return geosearchstorestorediststoredist{cmds: cmds}
}

type geosearchstorestorediststoredist struct {
	cmds []string
}

func (c geosearchstorestorediststoredist) Build() []string {
	return c.cmds
}

type geosearchwithcoordwithcoord struct {
	cmds []string
}

func (c geosearchwithcoordwithcoord) Withdist() geosearchwithdistwithdist {
	var cmds []string
	cmds = append(c.cmds, "WITHDIST")
	return geosearchwithdistwithdist{cmds: cmds}
}

func (c geosearchwithcoordwithcoord) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

func (c geosearchwithcoordwithcoord) Build() []string {
	return c.cmds
}

type geosearchwithdistwithdist struct {
	cmds []string
}

func (c geosearchwithdistwithdist) Withhash() geosearchwithhashwithhash {
	var cmds []string
	cmds = append(c.cmds, "WITHHASH")
	return geosearchwithhashwithhash{cmds: cmds}
}

func (c geosearchwithdistwithdist) Build() []string {
	return c.cmds
}

type geosearchwithhashwithhash struct {
	cmds []string
}

func (c geosearchwithhashwithhash) Build() []string {
	return c.cmds
}

type get struct {
	cmds []string
}

func (c get) Key(key string) getkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return getkey{cmds: cmds}
}

func Get() (c get) {
	c.cmds = append(c.cmds, "GET")
	return
}

type getbit struct {
	cmds []string
}

func (c getbit) Key(key string) getbitkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return getbitkey{cmds: cmds}
}

func Getbit() (c getbit) {
	c.cmds = append(c.cmds, "GETBIT")
	return
}

type getbitkey struct {
	cmds []string
}

func (c getbitkey) Offset(offset int64) getbitoffset {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(offset, 10))
	return getbitoffset{cmds: cmds}
}

type getbitoffset struct {
	cmds []string
}

func (c getbitoffset) Build() []string {
	return c.cmds
}

type getdel struct {
	cmds []string
}

func (c getdel) Key(key string) getdelkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return getdelkey{cmds: cmds}
}

func Getdel() (c getdel) {
	c.cmds = append(c.cmds, "GETDEL")
	return
}

type getdelkey struct {
	cmds []string
}

func (c getdelkey) Build() []string {
	return c.cmds
}

type getex struct {
	cmds []string
}

func (c getex) Key(key string) getexkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return getexkey{cmds: cmds}
}

func Getex() (c getex) {
	c.cmds = append(c.cmds, "GETEX")
	return
}

type getexexpirationex struct {
	cmds []string
}

func (c getexexpirationex) Build() []string {
	return c.cmds
}

type getexexpirationexat struct {
	cmds []string
}

func (c getexexpirationexat) Build() []string {
	return c.cmds
}

type getexexpirationpersist struct {
	cmds []string
}

func (c getexexpirationpersist) Build() []string {
	return c.cmds
}

type getexexpirationpx struct {
	cmds []string
}

func (c getexexpirationpx) Build() []string {
	return c.cmds
}

type getexexpirationpxat struct {
	cmds []string
}

func (c getexexpirationpxat) Build() []string {
	return c.cmds
}

type getexkey struct {
	cmds []string
}

func (c getexkey) Ex(seconds int64) getexexpirationex {
	var cmds []string
	cmds = append(c.cmds, "EX", strconv.FormatInt(seconds, 10))
	return getexexpirationex{cmds: cmds}
}

func (c getexkey) Px(milliseconds int64) getexexpirationpx {
	var cmds []string
	cmds = append(c.cmds, "PX", strconv.FormatInt(milliseconds, 10))
	return getexexpirationpx{cmds: cmds}
}

func (c getexkey) Exat(timestamp int64) getexexpirationexat {
	var cmds []string
	cmds = append(c.cmds, "EXAT", strconv.FormatInt(timestamp, 10))
	return getexexpirationexat{cmds: cmds}
}

func (c getexkey) Pxat(millisecondstimestamp int64) getexexpirationpxat {
	var cmds []string
	cmds = append(c.cmds, "PXAT", strconv.FormatInt(millisecondstimestamp, 10))
	return getexexpirationpxat{cmds: cmds}
}

func (c getexkey) Persist() getexexpirationpersist {
	var cmds []string
	cmds = append(c.cmds, "PERSIST")
	return getexexpirationpersist{cmds: cmds}
}

func (c getexkey) Build() []string {
	return c.cmds
}

type getkey struct {
	cmds []string
}

func (c getkey) Build() []string {
	return c.cmds
}

type getrange struct {
	cmds []string
}

func (c getrange) Key(key string) getrangekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return getrangekey{cmds: cmds}
}

func Getrange() (c getrange) {
	c.cmds = append(c.cmds, "GETRANGE")
	return
}

type getrangeend struct {
	cmds []string
}

func (c getrangeend) Build() []string {
	return c.cmds
}

type getrangekey struct {
	cmds []string
}

func (c getrangekey) Start(start int64) getrangestart {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(start, 10))
	return getrangestart{cmds: cmds}
}

type getrangestart struct {
	cmds []string
}

func (c getrangestart) End(end int64) getrangeend {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(end, 10))
	return getrangeend{cmds: cmds}
}

type getset struct {
	cmds []string
}

func (c getset) Key(key string) getsetkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return getsetkey{cmds: cmds}
}

func Getset() (c getset) {
	c.cmds = append(c.cmds, "GETSET")
	return
}

type getsetkey struct {
	cmds []string
}

func (c getsetkey) Value(value string) getsetvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return getsetvalue{cmds: cmds}
}

type getsetvalue struct {
	cmds []string
}

func (c getsetvalue) Build() []string {
	return c.cmds
}

type hdel struct {
	cmds []string
}

func (c hdel) Key(key string) hdelkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hdelkey{cmds: cmds}
}

func Hdel() (c hdel) {
	c.cmds = append(c.cmds, "HDEL")
	return
}

type hdelfield struct {
	cmds []string
}

func (c hdelfield) Field(field ...string) hdelfield {
	var cmds []string
	cmds = append(cmds, field...)
	return hdelfield{cmds: cmds}
}

type hdelkey struct {
	cmds []string
}

func (c hdelkey) Field(field ...string) hdelfield {
	var cmds []string
	cmds = append(cmds, field...)
	return hdelfield{cmds: cmds}
}

type hello struct {
	cmds []string
}

func (c hello) Protover(protover int64) helloargumentsprotover {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(protover, 10))
	return helloargumentsprotover{cmds: cmds}
}

func Hello() (c hello) {
	c.cmds = append(c.cmds, "HELLO")
	return
}

type helloargumentsauth struct {
	cmds []string
}

func (c helloargumentsauth) Setname(clientname string) helloargumentssetname {
	var cmds []string
	cmds = append(c.cmds, "SETNAME", clientname)
	return helloargumentssetname{cmds: cmds}
}

func (c helloargumentsauth) Build() []string {
	return c.cmds
}

type helloargumentsprotover struct {
	cmds []string
}

func (c helloargumentsprotover) Auth(username string, password string) helloargumentsauth {
	var cmds []string
	cmds = append(c.cmds, "AUTH", username, password)
	return helloargumentsauth{cmds: cmds}
}

func (c helloargumentsprotover) Setname(clientname string) helloargumentssetname {
	var cmds []string
	cmds = append(c.cmds, "SETNAME", clientname)
	return helloargumentssetname{cmds: cmds}
}

func (c helloargumentsprotover) Build() []string {
	return c.cmds
}

type helloargumentssetname struct {
	cmds []string
}

func (c helloargumentssetname) Build() []string {
	return c.cmds
}

type hexists struct {
	cmds []string
}

func (c hexists) Key(key string) hexistskey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hexistskey{cmds: cmds}
}

func Hexists() (c hexists) {
	c.cmds = append(c.cmds, "HEXISTS")
	return
}

type hexistsfield struct {
	cmds []string
}

func (c hexistsfield) Build() []string {
	return c.cmds
}

type hexistskey struct {
	cmds []string
}

func (c hexistskey) Field(field string) hexistsfield {
	var cmds []string
	cmds = append(c.cmds, field)
	return hexistsfield{cmds: cmds}
}

type hget struct {
	cmds []string
}

func (c hget) Key(key string) hgetkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hgetkey{cmds: cmds}
}

func Hget() (c hget) {
	c.cmds = append(c.cmds, "HGET")
	return
}

type hgetall struct {
	cmds []string
}

func (c hgetall) Key(key string) hgetallkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hgetallkey{cmds: cmds}
}

func Hgetall() (c hgetall) {
	c.cmds = append(c.cmds, "HGETALL")
	return
}

type hgetallkey struct {
	cmds []string
}

func (c hgetallkey) Build() []string {
	return c.cmds
}

type hgetfield struct {
	cmds []string
}

func (c hgetfield) Build() []string {
	return c.cmds
}

type hgetkey struct {
	cmds []string
}

func (c hgetkey) Field(field string) hgetfield {
	var cmds []string
	cmds = append(c.cmds, field)
	return hgetfield{cmds: cmds}
}

type hincrby struct {
	cmds []string
}

func (c hincrby) Key(key string) hincrbykey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hincrbykey{cmds: cmds}
}

func Hincrby() (c hincrby) {
	c.cmds = append(c.cmds, "HINCRBY")
	return
}

type hincrbyfield struct {
	cmds []string
}

func (c hincrbyfield) Increment(increment int64) hincrbyincrement {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(increment, 10))
	return hincrbyincrement{cmds: cmds}
}

type hincrbyfloat struct {
	cmds []string
}

func (c hincrbyfloat) Key(key string) hincrbyfloatkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hincrbyfloatkey{cmds: cmds}
}

func Hincrbyfloat() (c hincrbyfloat) {
	c.cmds = append(c.cmds, "HINCRBYFLOAT")
	return
}

type hincrbyfloatfield struct {
	cmds []string
}

func (c hincrbyfloatfield) Increment(increment float64) hincrbyfloatincrement {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(increment, 'f', -1, 64))
	return hincrbyfloatincrement{cmds: cmds}
}

type hincrbyfloatincrement struct {
	cmds []string
}

func (c hincrbyfloatincrement) Build() []string {
	return c.cmds
}

type hincrbyfloatkey struct {
	cmds []string
}

func (c hincrbyfloatkey) Field(field string) hincrbyfloatfield {
	var cmds []string
	cmds = append(c.cmds, field)
	return hincrbyfloatfield{cmds: cmds}
}

type hincrbyincrement struct {
	cmds []string
}

func (c hincrbyincrement) Build() []string {
	return c.cmds
}

type hincrbykey struct {
	cmds []string
}

func (c hincrbykey) Field(field string) hincrbyfield {
	var cmds []string
	cmds = append(c.cmds, field)
	return hincrbyfield{cmds: cmds}
}

type hkeys struct {
	cmds []string
}

func (c hkeys) Key(key string) hkeyskey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hkeyskey{cmds: cmds}
}

func Hkeys() (c hkeys) {
	c.cmds = append(c.cmds, "HKEYS")
	return
}

type hkeyskey struct {
	cmds []string
}

func (c hkeyskey) Build() []string {
	return c.cmds
}

type hlen struct {
	cmds []string
}

func (c hlen) Key(key string) hlenkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hlenkey{cmds: cmds}
}

func Hlen() (c hlen) {
	c.cmds = append(c.cmds, "HLEN")
	return
}

type hlenkey struct {
	cmds []string
}

func (c hlenkey) Build() []string {
	return c.cmds
}

type hmget struct {
	cmds []string
}

func (c hmget) Key(key string) hmgetkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hmgetkey{cmds: cmds}
}

func Hmget() (c hmget) {
	c.cmds = append(c.cmds, "HMGET")
	return
}

type hmgetfield struct {
	cmds []string
}

func (c hmgetfield) Field(field ...string) hmgetfield {
	var cmds []string
	cmds = append(cmds, field...)
	return hmgetfield{cmds: cmds}
}

type hmgetkey struct {
	cmds []string
}

func (c hmgetkey) Field(field ...string) hmgetfield {
	var cmds []string
	cmds = append(cmds, field...)
	return hmgetfield{cmds: cmds}
}

type hmset struct {
	cmds []string
}

func (c hmset) Key(key string) hmsetkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hmsetkey{cmds: cmds}
}

func Hmset() (c hmset) {
	c.cmds = append(c.cmds, "HMSET")
	return
}

type hmsetfieldvalue struct {
	cmds []string
}

func (c hmsetfieldvalue) FieldValue(field string, value string) hmsetfieldvalue {
	var cmds []string
	cmds = append(c.cmds, field, value)
	return hmsetfieldvalue{cmds: cmds}
}

type hmsetkey struct {
	cmds []string
}

func (c hmsetkey) FieldValue(field string, value string) hmsetfieldvalue {
	var cmds []string
	cmds = append(c.cmds, field, value)
	return hmsetfieldvalue{cmds: cmds}
}

type hrandfield struct {
	cmds []string
}

func (c hrandfield) Key(key string) hrandfieldkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hrandfieldkey{cmds: cmds}
}

func Hrandfield() (c hrandfield) {
	c.cmds = append(c.cmds, "HRANDFIELD")
	return
}

type hrandfieldkey struct {
	cmds []string
}

func (c hrandfieldkey) Count(count int64) hrandfieldoptionscount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return hrandfieldoptionscount{cmds: cmds}
}

type hrandfieldoptionscount struct {
	cmds []string
}

func (c hrandfieldoptionscount) Withvalues() hrandfieldoptionswithvalueswithvalues {
	var cmds []string
	cmds = append(c.cmds, "WITHVALUES")
	return hrandfieldoptionswithvalueswithvalues{cmds: cmds}
}

func (c hrandfieldoptionscount) Build() []string {
	return c.cmds
}

type hrandfieldoptionswithvalueswithvalues struct {
	cmds []string
}

func (c hrandfieldoptionswithvalueswithvalues) Build() []string {
	return c.cmds
}

type hscan struct {
	cmds []string
}

func (c hscan) Key(key string) hscankey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hscankey{cmds: cmds}
}

func Hscan() (c hscan) {
	c.cmds = append(c.cmds, "HSCAN")
	return
}

type hscancount struct {
	cmds []string
}

func (c hscancount) Build() []string {
	return c.cmds
}

type hscancursor struct {
	cmds []string
}

func (c hscancursor) Match(pattern string) hscanmatch {
	var cmds []string
	cmds = append(c.cmds, "MATCH", pattern)
	return hscanmatch{cmds: cmds}
}

func (c hscancursor) Count(count int64) hscancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return hscancount{cmds: cmds}
}

func (c hscancursor) Build() []string {
	return c.cmds
}

type hscankey struct {
	cmds []string
}

func (c hscankey) Cursor(cursor int64) hscancursor {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(cursor, 10))
	return hscancursor{cmds: cmds}
}

type hscanmatch struct {
	cmds []string
}

func (c hscanmatch) Count(count int64) hscancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return hscancount{cmds: cmds}
}

func (c hscanmatch) Build() []string {
	return c.cmds
}

type hset struct {
	cmds []string
}

func (c hset) Key(key string) hsetkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hsetkey{cmds: cmds}
}

func Hset() (c hset) {
	c.cmds = append(c.cmds, "HSET")
	return
}

type hsetfieldvalue struct {
	cmds []string
}

func (c hsetfieldvalue) FieldValue(field string, value string) hsetfieldvalue {
	var cmds []string
	cmds = append(c.cmds, field, value)
	return hsetfieldvalue{cmds: cmds}
}

type hsetkey struct {
	cmds []string
}

func (c hsetkey) FieldValue(field string, value string) hsetfieldvalue {
	var cmds []string
	cmds = append(c.cmds, field, value)
	return hsetfieldvalue{cmds: cmds}
}

type hsetnx struct {
	cmds []string
}

func (c hsetnx) Key(key string) hsetnxkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hsetnxkey{cmds: cmds}
}

func Hsetnx() (c hsetnx) {
	c.cmds = append(c.cmds, "HSETNX")
	return
}

type hsetnxfield struct {
	cmds []string
}

func (c hsetnxfield) Value(value string) hsetnxvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return hsetnxvalue{cmds: cmds}
}

type hsetnxkey struct {
	cmds []string
}

func (c hsetnxkey) Field(field string) hsetnxfield {
	var cmds []string
	cmds = append(c.cmds, field)
	return hsetnxfield{cmds: cmds}
}

type hsetnxvalue struct {
	cmds []string
}

func (c hsetnxvalue) Build() []string {
	return c.cmds
}

type hstrlen struct {
	cmds []string
}

func (c hstrlen) Key(key string) hstrlenkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hstrlenkey{cmds: cmds}
}

func Hstrlen() (c hstrlen) {
	c.cmds = append(c.cmds, "HSTRLEN")
	return
}

type hstrlenfield struct {
	cmds []string
}

func (c hstrlenfield) Build() []string {
	return c.cmds
}

type hstrlenkey struct {
	cmds []string
}

func (c hstrlenkey) Field(field string) hstrlenfield {
	var cmds []string
	cmds = append(c.cmds, field)
	return hstrlenfield{cmds: cmds}
}

type hvals struct {
	cmds []string
}

func (c hvals) Key(key string) hvalskey {
	var cmds []string
	cmds = append(c.cmds, key)
	return hvalskey{cmds: cmds}
}

func Hvals() (c hvals) {
	c.cmds = append(c.cmds, "HVALS")
	return
}

type hvalskey struct {
	cmds []string
}

func (c hvalskey) Build() []string {
	return c.cmds
}

type incr struct {
	cmds []string
}

func (c incr) Key(key string) incrkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return incrkey{cmds: cmds}
}

func Incr() (c incr) {
	c.cmds = append(c.cmds, "INCR")
	return
}

type incrby struct {
	cmds []string
}

func (c incrby) Key(key string) incrbykey {
	var cmds []string
	cmds = append(c.cmds, key)
	return incrbykey{cmds: cmds}
}

func Incrby() (c incrby) {
	c.cmds = append(c.cmds, "INCRBY")
	return
}

type incrbyfloat struct {
	cmds []string
}

func (c incrbyfloat) Key(key string) incrbyfloatkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return incrbyfloatkey{cmds: cmds}
}

func Incrbyfloat() (c incrbyfloat) {
	c.cmds = append(c.cmds, "INCRBYFLOAT")
	return
}

type incrbyfloatincrement struct {
	cmds []string
}

func (c incrbyfloatincrement) Build() []string {
	return c.cmds
}

type incrbyfloatkey struct {
	cmds []string
}

func (c incrbyfloatkey) Increment(increment float64) incrbyfloatincrement {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(increment, 'f', -1, 64))
	return incrbyfloatincrement{cmds: cmds}
}

type incrbyincrement struct {
	cmds []string
}

func (c incrbyincrement) Build() []string {
	return c.cmds
}

type incrbykey struct {
	cmds []string
}

func (c incrbykey) Increment(increment int64) incrbyincrement {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(increment, 10))
	return incrbyincrement{cmds: cmds}
}

type incrkey struct {
	cmds []string
}

func (c incrkey) Build() []string {
	return c.cmds
}

type info struct {
	cmds []string
}

func (c info) Section(section string) infosection {
	var cmds []string
	cmds = append(c.cmds, section)
	return infosection{cmds: cmds}
}

func (c info) Build() []string {
	return c.cmds
}

func Info() (c info) {
	c.cmds = append(c.cmds, "INFO")
	return
}

type infosection struct {
	cmds []string
}

func (c infosection) Build() []string {
	return c.cmds
}

type keys struct {
	cmds []string
}

func (c keys) Pattern(pattern string) keyspattern {
	var cmds []string
	cmds = append(c.cmds, pattern)
	return keyspattern{cmds: cmds}
}

func Keys() (c keys) {
	c.cmds = append(c.cmds, "KEYS")
	return
}

type keyspattern struct {
	cmds []string
}

func (c keyspattern) Build() []string {
	return c.cmds
}

type lastsave struct {
	cmds []string
}

func (c lastsave) Build() []string {
	return c.cmds
}

func Lastsave() (c lastsave) {
	c.cmds = append(c.cmds, "LASTSAVE")
	return
}

type latencydoctor struct {
	cmds []string
}

func (c latencydoctor) Build() []string {
	return c.cmds
}

func LatencyDoctor() (c latencydoctor) {
	c.cmds = append(c.cmds, "LATENCY", "DOCTOR")
	return
}

type latencygraph struct {
	cmds []string
}

func (c latencygraph) Event(event string) latencygraphevent {
	var cmds []string
	cmds = append(c.cmds, event)
	return latencygraphevent{cmds: cmds}
}

func LatencyGraph() (c latencygraph) {
	c.cmds = append(c.cmds, "LATENCY", "GRAPH")
	return
}

type latencygraphevent struct {
	cmds []string
}

func (c latencygraphevent) Build() []string {
	return c.cmds
}

type latencyhelp struct {
	cmds []string
}

func (c latencyhelp) Build() []string {
	return c.cmds
}

func LatencyHelp() (c latencyhelp) {
	c.cmds = append(c.cmds, "LATENCY", "HELP")
	return
}

type latencyhistory struct {
	cmds []string
}

func (c latencyhistory) Event(event string) latencyhistoryevent {
	var cmds []string
	cmds = append(c.cmds, event)
	return latencyhistoryevent{cmds: cmds}
}

func LatencyHistory() (c latencyhistory) {
	c.cmds = append(c.cmds, "LATENCY", "HISTORY")
	return
}

type latencyhistoryevent struct {
	cmds []string
}

func (c latencyhistoryevent) Build() []string {
	return c.cmds
}

type latencylatest struct {
	cmds []string
}

func (c latencylatest) Build() []string {
	return c.cmds
}

func LatencyLatest() (c latencylatest) {
	c.cmds = append(c.cmds, "LATENCY", "LATEST")
	return
}

type latencyreset struct {
	cmds []string
}

func (c latencyreset) Event(event ...string) latencyresetevent {
	var cmds []string
	cmds = append(cmds, event...)
	return latencyresetevent{cmds: cmds}
}

func (c latencyreset) Build() []string {
	return c.cmds
}

func LatencyReset() (c latencyreset) {
	c.cmds = append(c.cmds, "LATENCY", "RESET")
	return
}

type latencyresetevent struct {
	cmds []string
}

func (c latencyresetevent) Event(event ...string) latencyresetevent {
	var cmds []string
	cmds = append(cmds, event...)
	return latencyresetevent{cmds: cmds}
}

func (c latencyresetevent) Build() []string {
	return c.cmds
}

type lindex struct {
	cmds []string
}

func (c lindex) Key(key string) lindexkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lindexkey{cmds: cmds}
}

func Lindex() (c lindex) {
	c.cmds = append(c.cmds, "LINDEX")
	return
}

type lindexindex struct {
	cmds []string
}

func (c lindexindex) Build() []string {
	return c.cmds
}

type lindexkey struct {
	cmds []string
}

func (c lindexkey) Index(index int64) lindexindex {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(index, 10))
	return lindexindex{cmds: cmds}
}

type linsert struct {
	cmds []string
}

func (c linsert) Key(key string) linsertkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return linsertkey{cmds: cmds}
}

func Linsert() (c linsert) {
	c.cmds = append(c.cmds, "LINSERT")
	return
}

type linsertelement struct {
	cmds []string
}

func (c linsertelement) Build() []string {
	return c.cmds
}

type linsertkey struct {
	cmds []string
}

func (c linsertkey) Before() linsertwherebefore {
	var cmds []string
	cmds = append(c.cmds, "BEFORE")
	return linsertwherebefore{cmds: cmds}
}

func (c linsertkey) After() linsertwhereafter {
	var cmds []string
	cmds = append(c.cmds, "AFTER")
	return linsertwhereafter{cmds: cmds}
}

type linsertpivot struct {
	cmds []string
}

func (c linsertpivot) Element(element string) linsertelement {
	var cmds []string
	cmds = append(c.cmds, element)
	return linsertelement{cmds: cmds}
}

type linsertwhereafter struct {
	cmds []string
}

func (c linsertwhereafter) Pivot(pivot string) linsertpivot {
	var cmds []string
	cmds = append(c.cmds, pivot)
	return linsertpivot{cmds: cmds}
}

type linsertwherebefore struct {
	cmds []string
}

func (c linsertwherebefore) Pivot(pivot string) linsertpivot {
	var cmds []string
	cmds = append(c.cmds, pivot)
	return linsertpivot{cmds: cmds}
}

type llen struct {
	cmds []string
}

func (c llen) Key(key string) llenkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return llenkey{cmds: cmds}
}

func Llen() (c llen) {
	c.cmds = append(c.cmds, "LLEN")
	return
}

type llenkey struct {
	cmds []string
}

func (c llenkey) Build() []string {
	return c.cmds
}

type lmove struct {
	cmds []string
}

func (c lmove) Source(source string) lmovesource {
	var cmds []string
	cmds = append(c.cmds, source)
	return lmovesource{cmds: cmds}
}

func Lmove() (c lmove) {
	c.cmds = append(c.cmds, "LMOVE")
	return
}

type lmovedestination struct {
	cmds []string
}

func (c lmovedestination) Left() lmovewherefromleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return lmovewherefromleft{cmds: cmds}
}

func (c lmovedestination) Right() lmovewherefromright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return lmovewherefromright{cmds: cmds}
}

type lmovesource struct {
	cmds []string
}

func (c lmovesource) Destination(destination string) lmovedestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return lmovedestination{cmds: cmds}
}

type lmovewherefromleft struct {
	cmds []string
}

func (c lmovewherefromleft) Left() lmovewheretoleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return lmovewheretoleft{cmds: cmds}
}

func (c lmovewherefromleft) Right() lmovewheretoright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return lmovewheretoright{cmds: cmds}
}

type lmovewherefromright struct {
	cmds []string
}

func (c lmovewherefromright) Left() lmovewheretoleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return lmovewheretoleft{cmds: cmds}
}

func (c lmovewherefromright) Right() lmovewheretoright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return lmovewheretoright{cmds: cmds}
}

type lmovewheretoleft struct {
	cmds []string
}

func (c lmovewheretoleft) Build() []string {
	return c.cmds
}

type lmovewheretoright struct {
	cmds []string
}

func (c lmovewheretoright) Build() []string {
	return c.cmds
}

type lmpop struct {
	cmds []string
}

func (c lmpop) Numkeys(numkeys int64) lmpopnumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return lmpopnumkeys{cmds: cmds}
}

func Lmpop() (c lmpop) {
	c.cmds = append(c.cmds, "LMPOP")
	return
}

type lmpopcount struct {
	cmds []string
}

func (c lmpopcount) Build() []string {
	return c.cmds
}

type lmpopkey struct {
	cmds []string
}

func (c lmpopkey) Left() lmpopwhereleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return lmpopwhereleft{cmds: cmds}
}

func (c lmpopkey) Right() lmpopwhereright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return lmpopwhereright{cmds: cmds}
}

func (c lmpopkey) Key(key ...string) lmpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return lmpopkey{cmds: cmds}
}

type lmpopnumkeys struct {
	cmds []string
}

func (c lmpopnumkeys) Key(key ...string) lmpopkey {
	var cmds []string
	cmds = append(cmds, key...)
	return lmpopkey{cmds: cmds}
}

func (c lmpopnumkeys) Left() lmpopwhereleft {
	var cmds []string
	cmds = append(c.cmds, "LEFT")
	return lmpopwhereleft{cmds: cmds}
}

func (c lmpopnumkeys) Right() lmpopwhereright {
	var cmds []string
	cmds = append(c.cmds, "RIGHT")
	return lmpopwhereright{cmds: cmds}
}

type lmpopwhereleft struct {
	cmds []string
}

func (c lmpopwhereleft) Count(count int64) lmpopcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return lmpopcount{cmds: cmds}
}

func (c lmpopwhereleft) Build() []string {
	return c.cmds
}

type lmpopwhereright struct {
	cmds []string
}

func (c lmpopwhereright) Count(count int64) lmpopcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return lmpopcount{cmds: cmds}
}

func (c lmpopwhereright) Build() []string {
	return c.cmds
}

type lolwut struct {
	cmds []string
}

func (c lolwut) Version(version int64) lolwutversion {
	var cmds []string
	cmds = append(c.cmds, "VERSION", strconv.FormatInt(version, 10))
	return lolwutversion{cmds: cmds}
}

func (c lolwut) Build() []string {
	return c.cmds
}

func Lolwut() (c lolwut) {
	c.cmds = append(c.cmds, "LOLWUT")
	return
}

type lolwutversion struct {
	cmds []string
}

func (c lolwutversion) Build() []string {
	return c.cmds
}

type lpop struct {
	cmds []string
}

func (c lpop) Key(key string) lpopkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lpopkey{cmds: cmds}
}

func Lpop() (c lpop) {
	c.cmds = append(c.cmds, "LPOP")
	return
}

type lpopcount struct {
	cmds []string
}

func (c lpopcount) Build() []string {
	return c.cmds
}

type lpopkey struct {
	cmds []string
}

func (c lpopkey) Count(count int64) lpopcount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return lpopcount{cmds: cmds}
}

func (c lpopkey) Build() []string {
	return c.cmds
}

type lpos struct {
	cmds []string
}

func (c lpos) Key(key string) lposkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lposkey{cmds: cmds}
}

func Lpos() (c lpos) {
	c.cmds = append(c.cmds, "LPOS")
	return
}

type lposcount struct {
	cmds []string
}

func (c lposcount) Maxlen(len int64) lposmaxlen {
	var cmds []string
	cmds = append(c.cmds, "MAXLEN", strconv.FormatInt(len, 10))
	return lposmaxlen{cmds: cmds}
}

func (c lposcount) Build() []string {
	return c.cmds
}

type lposelement struct {
	cmds []string
}

func (c lposelement) Rank(rank int64) lposrank {
	var cmds []string
	cmds = append(c.cmds, "RANK", strconv.FormatInt(rank, 10))
	return lposrank{cmds: cmds}
}

func (c lposelement) Count(nummatches int64) lposcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(nummatches, 10))
	return lposcount{cmds: cmds}
}

func (c lposelement) Maxlen(len int64) lposmaxlen {
	var cmds []string
	cmds = append(c.cmds, "MAXLEN", strconv.FormatInt(len, 10))
	return lposmaxlen{cmds: cmds}
}

func (c lposelement) Build() []string {
	return c.cmds
}

type lposkey struct {
	cmds []string
}

func (c lposkey) Element(element string) lposelement {
	var cmds []string
	cmds = append(c.cmds, element)
	return lposelement{cmds: cmds}
}

type lposmaxlen struct {
	cmds []string
}

func (c lposmaxlen) Build() []string {
	return c.cmds
}

type lposrank struct {
	cmds []string
}

func (c lposrank) Count(nummatches int64) lposcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(nummatches, 10))
	return lposcount{cmds: cmds}
}

func (c lposrank) Maxlen(len int64) lposmaxlen {
	var cmds []string
	cmds = append(c.cmds, "MAXLEN", strconv.FormatInt(len, 10))
	return lposmaxlen{cmds: cmds}
}

func (c lposrank) Build() []string {
	return c.cmds
}

type lpush struct {
	cmds []string
}

func (c lpush) Key(key string) lpushkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lpushkey{cmds: cmds}
}

func Lpush() (c lpush) {
	c.cmds = append(c.cmds, "LPUSH")
	return
}

type lpushelement struct {
	cmds []string
}

func (c lpushelement) Element(element ...string) lpushelement {
	var cmds []string
	cmds = append(cmds, element...)
	return lpushelement{cmds: cmds}
}

type lpushkey struct {
	cmds []string
}

func (c lpushkey) Element(element ...string) lpushelement {
	var cmds []string
	cmds = append(cmds, element...)
	return lpushelement{cmds: cmds}
}

type lpushx struct {
	cmds []string
}

func (c lpushx) Key(key string) lpushxkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lpushxkey{cmds: cmds}
}

func Lpushx() (c lpushx) {
	c.cmds = append(c.cmds, "LPUSHX")
	return
}

type lpushxelement struct {
	cmds []string
}

func (c lpushxelement) Element(element ...string) lpushxelement {
	var cmds []string
	cmds = append(cmds, element...)
	return lpushxelement{cmds: cmds}
}

type lpushxkey struct {
	cmds []string
}

func (c lpushxkey) Element(element ...string) lpushxelement {
	var cmds []string
	cmds = append(cmds, element...)
	return lpushxelement{cmds: cmds}
}

type lrange struct {
	cmds []string
}

func (c lrange) Key(key string) lrangekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lrangekey{cmds: cmds}
}

func Lrange() (c lrange) {
	c.cmds = append(c.cmds, "LRANGE")
	return
}

type lrangekey struct {
	cmds []string
}

func (c lrangekey) Start(start int64) lrangestart {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(start, 10))
	return lrangestart{cmds: cmds}
}

type lrangestart struct {
	cmds []string
}

func (c lrangestart) Stop(stop int64) lrangestop {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(stop, 10))
	return lrangestop{cmds: cmds}
}

type lrangestop struct {
	cmds []string
}

func (c lrangestop) Build() []string {
	return c.cmds
}

type lrem struct {
	cmds []string
}

func (c lrem) Key(key string) lremkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lremkey{cmds: cmds}
}

func Lrem() (c lrem) {
	c.cmds = append(c.cmds, "LREM")
	return
}

type lremcount struct {
	cmds []string
}

func (c lremcount) Element(element string) lremelement {
	var cmds []string
	cmds = append(c.cmds, element)
	return lremelement{cmds: cmds}
}

type lremelement struct {
	cmds []string
}

func (c lremelement) Build() []string {
	return c.cmds
}

type lremkey struct {
	cmds []string
}

func (c lremkey) Count(count int64) lremcount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return lremcount{cmds: cmds}
}

type lset struct {
	cmds []string
}

func (c lset) Key(key string) lsetkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return lsetkey{cmds: cmds}
}

func Lset() (c lset) {
	c.cmds = append(c.cmds, "LSET")
	return
}

type lsetelement struct {
	cmds []string
}

func (c lsetelement) Build() []string {
	return c.cmds
}

type lsetindex struct {
	cmds []string
}

func (c lsetindex) Element(element string) lsetelement {
	var cmds []string
	cmds = append(c.cmds, element)
	return lsetelement{cmds: cmds}
}

type lsetkey struct {
	cmds []string
}

func (c lsetkey) Index(index int64) lsetindex {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(index, 10))
	return lsetindex{cmds: cmds}
}

type ltrim struct {
	cmds []string
}

func (c ltrim) Key(key string) ltrimkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return ltrimkey{cmds: cmds}
}

func Ltrim() (c ltrim) {
	c.cmds = append(c.cmds, "LTRIM")
	return
}

type ltrimkey struct {
	cmds []string
}

func (c ltrimkey) Start(start int64) ltrimstart {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(start, 10))
	return ltrimstart{cmds: cmds}
}

type ltrimstart struct {
	cmds []string
}

func (c ltrimstart) Stop(stop int64) ltrimstop {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(stop, 10))
	return ltrimstop{cmds: cmds}
}

type ltrimstop struct {
	cmds []string
}

func (c ltrimstop) Build() []string {
	return c.cmds
}

type memorydoctor struct {
	cmds []string
}

func (c memorydoctor) Build() []string {
	return c.cmds
}

func MemoryDoctor() (c memorydoctor) {
	c.cmds = append(c.cmds, "MEMORY", "DOCTOR")
	return
}

type memoryhelp struct {
	cmds []string
}

func (c memoryhelp) Build() []string {
	return c.cmds
}

func MemoryHelp() (c memoryhelp) {
	c.cmds = append(c.cmds, "MEMORY", "HELP")
	return
}

type memorymallocstats struct {
	cmds []string
}

func (c memorymallocstats) Build() []string {
	return c.cmds
}

func MemoryMallocstats() (c memorymallocstats) {
	c.cmds = append(c.cmds, "MEMORY", "MALLOC-STATS")
	return
}

type memorypurge struct {
	cmds []string
}

func (c memorypurge) Build() []string {
	return c.cmds
}

func MemoryPurge() (c memorypurge) {
	c.cmds = append(c.cmds, "MEMORY", "PURGE")
	return
}

type memorystats struct {
	cmds []string
}

func (c memorystats) Build() []string {
	return c.cmds
}

func MemoryStats() (c memorystats) {
	c.cmds = append(c.cmds, "MEMORY", "STATS")
	return
}

type memoryusage struct {
	cmds []string
}

func (c memoryusage) Key(key string) memoryusagekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return memoryusagekey{cmds: cmds}
}

func MemoryUsage() (c memoryusage) {
	c.cmds = append(c.cmds, "MEMORY", "USAGE")
	return
}

type memoryusagekey struct {
	cmds []string
}

func (c memoryusagekey) Samples(count int64) memoryusagesamples {
	var cmds []string
	cmds = append(c.cmds, "SAMPLES", strconv.FormatInt(count, 10))
	return memoryusagesamples{cmds: cmds}
}

func (c memoryusagekey) Build() []string {
	return c.cmds
}

type memoryusagesamples struct {
	cmds []string
}

func (c memoryusagesamples) Build() []string {
	return c.cmds
}

type mget struct {
	cmds []string
}

func (c mget) Key(key ...string) mgetkey {
	var cmds []string
	cmds = append(cmds, key...)
	return mgetkey{cmds: cmds}
}

func Mget() (c mget) {
	c.cmds = append(c.cmds, "MGET")
	return
}

type mgetkey struct {
	cmds []string
}

func (c mgetkey) Key(key ...string) mgetkey {
	var cmds []string
	cmds = append(cmds, key...)
	return mgetkey{cmds: cmds}
}

type migrate struct {
	cmds []string
}

func (c migrate) Host(host string) migratehost {
	var cmds []string
	cmds = append(c.cmds, host)
	return migratehost{cmds: cmds}
}

func Migrate() (c migrate) {
	c.cmds = append(c.cmds, "MIGRATE")
	return
}

type migrateauth struct {
	cmds []string
}

func (c migrateauth) Auth2(usernamePassword string) migrateauth2 {
	var cmds []string
	cmds = append(c.cmds, "AUTH2", usernamePassword)
	return migrateauth2{cmds: cmds}
}

func (c migrateauth) Keys(key ...string) migratekeys {
	var cmds []string
	cmds = append(c.cmds, "KEYS")
	cmds = append(cmds, key...)
	return migratekeys{cmds: cmds}
}

func (c migrateauth) Build() []string {
	return c.cmds
}

type migrateauth2 struct {
	cmds []string
}

func (c migrateauth2) Keys(key ...string) migratekeys {
	var cmds []string
	cmds = append(c.cmds, "KEYS")
	cmds = append(cmds, key...)
	return migratekeys{cmds: cmds}
}

func (c migrateauth2) Build() []string {
	return c.cmds
}

type migratecopycopy struct {
	cmds []string
}

func (c migratecopycopy) Replace() migratereplacereplace {
	var cmds []string
	cmds = append(c.cmds, "REPLACE")
	return migratereplacereplace{cmds: cmds}
}

func (c migratecopycopy) Auth(password string) migrateauth {
	var cmds []string
	cmds = append(c.cmds, "AUTH", password)
	return migrateauth{cmds: cmds}
}

func (c migratecopycopy) Auth2(usernamePassword string) migrateauth2 {
	var cmds []string
	cmds = append(c.cmds, "AUTH2", usernamePassword)
	return migrateauth2{cmds: cmds}
}

func (c migratecopycopy) Keys(key ...string) migratekeys {
	var cmds []string
	cmds = append(c.cmds, "KEYS")
	cmds = append(cmds, key...)
	return migratekeys{cmds: cmds}
}

func (c migratecopycopy) Build() []string {
	return c.cmds
}

type migratedestinationdb struct {
	cmds []string
}

func (c migratedestinationdb) Timeout(timeout int64) migratetimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(timeout, 10))
	return migratetimeout{cmds: cmds}
}

type migratehost struct {
	cmds []string
}

func (c migratehost) Port(port string) migrateport {
	var cmds []string
	cmds = append(c.cmds, port)
	return migrateport{cmds: cmds}
}

type migratekeyempty struct {
	cmds []string
}

func (c migratekeyempty) Destinationdb(destinationdb int64) migratedestinationdb {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(destinationdb, 10))
	return migratedestinationdb{cmds: cmds}
}

type migratekeykey struct {
	cmds []string
}

func (c migratekeykey) Destinationdb(destinationdb int64) migratedestinationdb {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(destinationdb, 10))
	return migratedestinationdb{cmds: cmds}
}

type migratekeys struct {
	cmds []string
}

func (c migratekeys) Key(key ...string) migratekeys {
	var cmds []string
	cmds = append(cmds, key...)
	return migratekeys{cmds: cmds}
}

func (c migratekeys) Build() []string {
	return c.cmds
}

type migrateport struct {
	cmds []string
}

func (c migrateport) Key() migratekeykey {
	var cmds []string
	cmds = append(c.cmds, "key")
	return migratekeykey{cmds: cmds}
}

func (c migrateport) Empty() migratekeyempty {
	var cmds []string
	cmds = append(c.cmds, "\"\"")
	return migratekeyempty{cmds: cmds}
}

type migratereplacereplace struct {
	cmds []string
}

func (c migratereplacereplace) Auth(password string) migrateauth {
	var cmds []string
	cmds = append(c.cmds, "AUTH", password)
	return migrateauth{cmds: cmds}
}

func (c migratereplacereplace) Auth2(usernamePassword string) migrateauth2 {
	var cmds []string
	cmds = append(c.cmds, "AUTH2", usernamePassword)
	return migrateauth2{cmds: cmds}
}

func (c migratereplacereplace) Keys(key ...string) migratekeys {
	var cmds []string
	cmds = append(c.cmds, "KEYS")
	cmds = append(cmds, key...)
	return migratekeys{cmds: cmds}
}

func (c migratereplacereplace) Build() []string {
	return c.cmds
}

type migratetimeout struct {
	cmds []string
}

func (c migratetimeout) Copy() migratecopycopy {
	var cmds []string
	cmds = append(c.cmds, "COPY")
	return migratecopycopy{cmds: cmds}
}

func (c migratetimeout) Replace() migratereplacereplace {
	var cmds []string
	cmds = append(c.cmds, "REPLACE")
	return migratereplacereplace{cmds: cmds}
}

func (c migratetimeout) Auth(password string) migrateauth {
	var cmds []string
	cmds = append(c.cmds, "AUTH", password)
	return migrateauth{cmds: cmds}
}

func (c migratetimeout) Auth2(usernamePassword string) migrateauth2 {
	var cmds []string
	cmds = append(c.cmds, "AUTH2", usernamePassword)
	return migrateauth2{cmds: cmds}
}

func (c migratetimeout) Keys(key ...string) migratekeys {
	var cmds []string
	cmds = append(c.cmds, "KEYS")
	cmds = append(cmds, key...)
	return migratekeys{cmds: cmds}
}

func (c migratetimeout) Build() []string {
	return c.cmds
}

type modulelist struct {
	cmds []string
}

func (c modulelist) Build() []string {
	return c.cmds
}

func ModuleList() (c modulelist) {
	c.cmds = append(c.cmds, "MODULE", "LIST")
	return
}

type moduleload struct {
	cmds []string
}

func (c moduleload) Path(path string) moduleloadpath {
	var cmds []string
	cmds = append(c.cmds, path)
	return moduleloadpath{cmds: cmds}
}

func ModuleLoad() (c moduleload) {
	c.cmds = append(c.cmds, "MODULE", "LOAD")
	return
}

type moduleloadarg struct {
	cmds []string
}

func (c moduleloadarg) Arg(arg ...string) moduleloadarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return moduleloadarg{cmds: cmds}
}

func (c moduleloadarg) Build() []string {
	return c.cmds
}

type moduleloadpath struct {
	cmds []string
}

func (c moduleloadpath) Arg(arg ...string) moduleloadarg {
	var cmds []string
	cmds = append(cmds, arg...)
	return moduleloadarg{cmds: cmds}
}

func (c moduleloadpath) Build() []string {
	return c.cmds
}

type moduleunload struct {
	cmds []string
}

func (c moduleunload) Name(name string) moduleunloadname {
	var cmds []string
	cmds = append(c.cmds, name)
	return moduleunloadname{cmds: cmds}
}

func ModuleUnload() (c moduleunload) {
	c.cmds = append(c.cmds, "MODULE", "UNLOAD")
	return
}

type moduleunloadname struct {
	cmds []string
}

func (c moduleunloadname) Build() []string {
	return c.cmds
}

type monitor struct {
	cmds []string
}

func (c monitor) Build() []string {
	return c.cmds
}

func Monitor() (c monitor) {
	c.cmds = append(c.cmds, "MONITOR")
	return
}

type move struct {
	cmds []string
}

func (c move) Key(key string) movekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return movekey{cmds: cmds}
}

func Move() (c move) {
	c.cmds = append(c.cmds, "MOVE")
	return
}

type movedb struct {
	cmds []string
}

func (c movedb) Build() []string {
	return c.cmds
}

type movekey struct {
	cmds []string
}

func (c movekey) Db(db int64) movedb {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(db, 10))
	return movedb{cmds: cmds}
}

type mset struct {
	cmds []string
}

func (c mset) KeyValue(key string, value string) msetkeyvalue {
	var cmds []string
	cmds = append(c.cmds, key, value)
	return msetkeyvalue{cmds: cmds}
}

func Mset() (c mset) {
	c.cmds = append(c.cmds, "MSET")
	return
}

type msetkeyvalue struct {
	cmds []string
}

func (c msetkeyvalue) KeyValue(key string, value string) msetkeyvalue {
	var cmds []string
	cmds = append(c.cmds, key, value)
	return msetkeyvalue{cmds: cmds}
}

type msetnx struct {
	cmds []string
}

func (c msetnx) KeyValue(key string, value string) msetnxkeyvalue {
	var cmds []string
	cmds = append(c.cmds, key, value)
	return msetnxkeyvalue{cmds: cmds}
}

func Msetnx() (c msetnx) {
	c.cmds = append(c.cmds, "MSETNX")
	return
}

type msetnxkeyvalue struct {
	cmds []string
}

func (c msetnxkeyvalue) KeyValue(key string, value string) msetnxkeyvalue {
	var cmds []string
	cmds = append(c.cmds, key, value)
	return msetnxkeyvalue{cmds: cmds}
}

type multi struct {
	cmds []string
}

func (c multi) Build() []string {
	return c.cmds
}

func Multi() (c multi) {
	c.cmds = append(c.cmds, "MULTI")
	return
}

type object struct {
	cmds []string
}

func (c object) Subcommand(subcommand string) objectsubcommand {
	var cmds []string
	cmds = append(c.cmds, subcommand)
	return objectsubcommand{cmds: cmds}
}

func Object() (c object) {
	c.cmds = append(c.cmds, "OBJECT")
	return
}

type objectarguments struct {
	cmds []string
}

func (c objectarguments) Arguments(arguments ...string) objectarguments {
	var cmds []string
	cmds = append(cmds, arguments...)
	return objectarguments{cmds: cmds}
}

func (c objectarguments) Build() []string {
	return c.cmds
}

type objectsubcommand struct {
	cmds []string
}

func (c objectsubcommand) Arguments(arguments ...string) objectarguments {
	var cmds []string
	cmds = append(cmds, arguments...)
	return objectarguments{cmds: cmds}
}

func (c objectsubcommand) Build() []string {
	return c.cmds
}

type persist struct {
	cmds []string
}

func (c persist) Key(key string) persistkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return persistkey{cmds: cmds}
}

func Persist() (c persist) {
	c.cmds = append(c.cmds, "PERSIST")
	return
}

type persistkey struct {
	cmds []string
}

func (c persistkey) Build() []string {
	return c.cmds
}

type pexpire struct {
	cmds []string
}

func (c pexpire) Key(key string) pexpirekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return pexpirekey{cmds: cmds}
}

func Pexpire() (c pexpire) {
	c.cmds = append(c.cmds, "PEXPIRE")
	return
}

type pexpireat struct {
	cmds []string
}

func (c pexpireat) Key(key string) pexpireatkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return pexpireatkey{cmds: cmds}
}

func Pexpireat() (c pexpireat) {
	c.cmds = append(c.cmds, "PEXPIREAT")
	return
}

type pexpireatconditiongt struct {
	cmds []string
}

func (c pexpireatconditiongt) Build() []string {
	return c.cmds
}

type pexpireatconditionlt struct {
	cmds []string
}

func (c pexpireatconditionlt) Build() []string {
	return c.cmds
}

type pexpireatconditionnx struct {
	cmds []string
}

func (c pexpireatconditionnx) Build() []string {
	return c.cmds
}

type pexpireatconditionxx struct {
	cmds []string
}

func (c pexpireatconditionxx) Build() []string {
	return c.cmds
}

type pexpireatkey struct {
	cmds []string
}

func (c pexpireatkey) Millisecondstimestamp(millisecondstimestamp int64) pexpireatmillisecondstimestamp {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(millisecondstimestamp, 10))
	return pexpireatmillisecondstimestamp{cmds: cmds}
}

type pexpireatmillisecondstimestamp struct {
	cmds []string
}

func (c pexpireatmillisecondstimestamp) Nx() pexpireatconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return pexpireatconditionnx{cmds: cmds}
}

func (c pexpireatmillisecondstimestamp) Xx() pexpireatconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return pexpireatconditionxx{cmds: cmds}
}

func (c pexpireatmillisecondstimestamp) Gt() pexpireatconditiongt {
	var cmds []string
	cmds = append(c.cmds, "GT")
	return pexpireatconditiongt{cmds: cmds}
}

func (c pexpireatmillisecondstimestamp) Lt() pexpireatconditionlt {
	var cmds []string
	cmds = append(c.cmds, "LT")
	return pexpireatconditionlt{cmds: cmds}
}

func (c pexpireatmillisecondstimestamp) Build() []string {
	return c.cmds
}

type pexpireconditiongt struct {
	cmds []string
}

func (c pexpireconditiongt) Build() []string {
	return c.cmds
}

type pexpireconditionlt struct {
	cmds []string
}

func (c pexpireconditionlt) Build() []string {
	return c.cmds
}

type pexpireconditionnx struct {
	cmds []string
}

func (c pexpireconditionnx) Build() []string {
	return c.cmds
}

type pexpireconditionxx struct {
	cmds []string
}

func (c pexpireconditionxx) Build() []string {
	return c.cmds
}

type pexpirekey struct {
	cmds []string
}

func (c pexpirekey) Milliseconds(milliseconds int64) pexpiremilliseconds {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(milliseconds, 10))
	return pexpiremilliseconds{cmds: cmds}
}

type pexpiremilliseconds struct {
	cmds []string
}

func (c pexpiremilliseconds) Nx() pexpireconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return pexpireconditionnx{cmds: cmds}
}

func (c pexpiremilliseconds) Xx() pexpireconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return pexpireconditionxx{cmds: cmds}
}

func (c pexpiremilliseconds) Gt() pexpireconditiongt {
	var cmds []string
	cmds = append(c.cmds, "GT")
	return pexpireconditiongt{cmds: cmds}
}

func (c pexpiremilliseconds) Lt() pexpireconditionlt {
	var cmds []string
	cmds = append(c.cmds, "LT")
	return pexpireconditionlt{cmds: cmds}
}

func (c pexpiremilliseconds) Build() []string {
	return c.cmds
}

type pexpiretime struct {
	cmds []string
}

func (c pexpiretime) Key(key string) pexpiretimekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return pexpiretimekey{cmds: cmds}
}

func Pexpiretime() (c pexpiretime) {
	c.cmds = append(c.cmds, "PEXPIRETIME")
	return
}

type pexpiretimekey struct {
	cmds []string
}

func (c pexpiretimekey) Build() []string {
	return c.cmds
}

type pfadd struct {
	cmds []string
}

func (c pfadd) Key(key string) pfaddkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return pfaddkey{cmds: cmds}
}

func Pfadd() (c pfadd) {
	c.cmds = append(c.cmds, "PFADD")
	return
}

type pfaddelement struct {
	cmds []string
}

func (c pfaddelement) Element(element ...string) pfaddelement {
	var cmds []string
	cmds = append(cmds, element...)
	return pfaddelement{cmds: cmds}
}

func (c pfaddelement) Build() []string {
	return c.cmds
}

type pfaddkey struct {
	cmds []string
}

func (c pfaddkey) Element(element ...string) pfaddelement {
	var cmds []string
	cmds = append(cmds, element...)
	return pfaddelement{cmds: cmds}
}

func (c pfaddkey) Build() []string {
	return c.cmds
}

type pfcount struct {
	cmds []string
}

func (c pfcount) Key(key ...string) pfcountkey {
	var cmds []string
	cmds = append(cmds, key...)
	return pfcountkey{cmds: cmds}
}

func Pfcount() (c pfcount) {
	c.cmds = append(c.cmds, "PFCOUNT")
	return
}

type pfcountkey struct {
	cmds []string
}

func (c pfcountkey) Key(key ...string) pfcountkey {
	var cmds []string
	cmds = append(cmds, key...)
	return pfcountkey{cmds: cmds}
}

type pfmerge struct {
	cmds []string
}

func (c pfmerge) Destkey(destkey string) pfmergedestkey {
	var cmds []string
	cmds = append(c.cmds, destkey)
	return pfmergedestkey{cmds: cmds}
}

func Pfmerge() (c pfmerge) {
	c.cmds = append(c.cmds, "PFMERGE")
	return
}

type pfmergedestkey struct {
	cmds []string
}

func (c pfmergedestkey) Sourcekey(sourcekey ...string) pfmergesourcekey {
	var cmds []string
	cmds = append(cmds, sourcekey...)
	return pfmergesourcekey{cmds: cmds}
}

type pfmergesourcekey struct {
	cmds []string
}

func (c pfmergesourcekey) Sourcekey(sourcekey ...string) pfmergesourcekey {
	var cmds []string
	cmds = append(cmds, sourcekey...)
	return pfmergesourcekey{cmds: cmds}
}

type ping struct {
	cmds []string
}

func (c ping) Message(message string) pingmessage {
	var cmds []string
	cmds = append(c.cmds, message)
	return pingmessage{cmds: cmds}
}

func (c ping) Build() []string {
	return c.cmds
}

func Ping() (c ping) {
	c.cmds = append(c.cmds, "PING")
	return
}

type pingmessage struct {
	cmds []string
}

func (c pingmessage) Build() []string {
	return c.cmds
}

type psetex struct {
	cmds []string
}

func (c psetex) Key(key string) psetexkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return psetexkey{cmds: cmds}
}

func Psetex() (c psetex) {
	c.cmds = append(c.cmds, "PSETEX")
	return
}

type psetexkey struct {
	cmds []string
}

func (c psetexkey) Milliseconds(milliseconds int64) psetexmilliseconds {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(milliseconds, 10))
	return psetexmilliseconds{cmds: cmds}
}

type psetexmilliseconds struct {
	cmds []string
}

func (c psetexmilliseconds) Value(value string) psetexvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return psetexvalue{cmds: cmds}
}

type psetexvalue struct {
	cmds []string
}

func (c psetexvalue) Build() []string {
	return c.cmds
}

type psubscribe struct {
	cmds []string
}

func (c psubscribe) Pattern(pattern ...string) psubscribepattern {
	var cmds []string
	cmds = append(cmds, pattern...)
	return psubscribepattern{cmds: cmds}
}

func Psubscribe() (c psubscribe) {
	c.cmds = append(c.cmds, "PSUBSCRIBE")
	return
}

type psubscribepattern struct {
	cmds []string
}

func (c psubscribepattern) Pattern(pattern ...string) psubscribepattern {
	var cmds []string
	cmds = append(cmds, pattern...)
	return psubscribepattern{cmds: cmds}
}

type psync struct {
	cmds []string
}

func (c psync) Replicationid(replicationid int64) psyncreplicationid {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(replicationid, 10))
	return psyncreplicationid{cmds: cmds}
}

func Psync() (c psync) {
	c.cmds = append(c.cmds, "PSYNC")
	return
}

type psyncoffset struct {
	cmds []string
}

func (c psyncoffset) Build() []string {
	return c.cmds
}

type psyncreplicationid struct {
	cmds []string
}

func (c psyncreplicationid) Offset(offset int64) psyncoffset {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(offset, 10))
	return psyncoffset{cmds: cmds}
}

type pttl struct {
	cmds []string
}

func (c pttl) Key(key string) pttlkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return pttlkey{cmds: cmds}
}

func Pttl() (c pttl) {
	c.cmds = append(c.cmds, "PTTL")
	return
}

type pttlkey struct {
	cmds []string
}

func (c pttlkey) Build() []string {
	return c.cmds
}

type publish struct {
	cmds []string
}

func (c publish) Channel(channel string) publishchannel {
	var cmds []string
	cmds = append(c.cmds, channel)
	return publishchannel{cmds: cmds}
}

func Publish() (c publish) {
	c.cmds = append(c.cmds, "PUBLISH")
	return
}

type publishchannel struct {
	cmds []string
}

func (c publishchannel) Message(message string) publishmessage {
	var cmds []string
	cmds = append(c.cmds, message)
	return publishmessage{cmds: cmds}
}

type publishmessage struct {
	cmds []string
}

func (c publishmessage) Build() []string {
	return c.cmds
}

type pubsub struct {
	cmds []string
}

func (c pubsub) Subcommand(subcommand string) pubsubsubcommand {
	var cmds []string
	cmds = append(c.cmds, subcommand)
	return pubsubsubcommand{cmds: cmds}
}

func Pubsub() (c pubsub) {
	c.cmds = append(c.cmds, "PUBSUB")
	return
}

type pubsubargument struct {
	cmds []string
}

func (c pubsubargument) Argument(argument ...string) pubsubargument {
	var cmds []string
	cmds = append(cmds, argument...)
	return pubsubargument{cmds: cmds}
}

func (c pubsubargument) Build() []string {
	return c.cmds
}

type pubsubsubcommand struct {
	cmds []string
}

func (c pubsubsubcommand) Argument(argument ...string) pubsubargument {
	var cmds []string
	cmds = append(cmds, argument...)
	return pubsubargument{cmds: cmds}
}

func (c pubsubsubcommand) Build() []string {
	return c.cmds
}

type punsubscribe struct {
	cmds []string
}

func (c punsubscribe) Pattern(pattern ...string) punsubscribepattern {
	var cmds []string
	cmds = append(cmds, pattern...)
	return punsubscribepattern{cmds: cmds}
}

func (c punsubscribe) Build() []string {
	return c.cmds
}

func Punsubscribe() (c punsubscribe) {
	c.cmds = append(c.cmds, "PUNSUBSCRIBE")
	return
}

type punsubscribepattern struct {
	cmds []string
}

func (c punsubscribepattern) Pattern(pattern ...string) punsubscribepattern {
	var cmds []string
	cmds = append(cmds, pattern...)
	return punsubscribepattern{cmds: cmds}
}

func (c punsubscribepattern) Build() []string {
	return c.cmds
}

type quit struct {
	cmds []string
}

func (c quit) Build() []string {
	return c.cmds
}

func Quit() (c quit) {
	c.cmds = append(c.cmds, "QUIT")
	return
}

type randomkey struct {
	cmds []string
}

func (c randomkey) Build() []string {
	return c.cmds
}

func Randomkey() (c randomkey) {
	c.cmds = append(c.cmds, "RANDOMKEY")
	return
}

type readonly struct {
	cmds []string
}

func (c readonly) Build() []string {
	return c.cmds
}

func Readonly() (c readonly) {
	c.cmds = append(c.cmds, "READONLY")
	return
}

type readwrite struct {
	cmds []string
}

func (c readwrite) Build() []string {
	return c.cmds
}

func Readwrite() (c readwrite) {
	c.cmds = append(c.cmds, "READWRITE")
	return
}

type rename struct {
	cmds []string
}

func (c rename) Key(key string) renamekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return renamekey{cmds: cmds}
}

func Rename() (c rename) {
	c.cmds = append(c.cmds, "RENAME")
	return
}

type renamekey struct {
	cmds []string
}

func (c renamekey) Newkey(newkey string) renamenewkey {
	var cmds []string
	cmds = append(c.cmds, newkey)
	return renamenewkey{cmds: cmds}
}

type renamenewkey struct {
	cmds []string
}

func (c renamenewkey) Build() []string {
	return c.cmds
}

type renamenx struct {
	cmds []string
}

func (c renamenx) Key(key string) renamenxkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return renamenxkey{cmds: cmds}
}

func Renamenx() (c renamenx) {
	c.cmds = append(c.cmds, "RENAMENX")
	return
}

type renamenxkey struct {
	cmds []string
}

func (c renamenxkey) Newkey(newkey string) renamenxnewkey {
	var cmds []string
	cmds = append(c.cmds, newkey)
	return renamenxnewkey{cmds: cmds}
}

type renamenxnewkey struct {
	cmds []string
}

func (c renamenxnewkey) Build() []string {
	return c.cmds
}

type replicaof struct {
	cmds []string
}

func (c replicaof) Host(host string) replicaofhost {
	var cmds []string
	cmds = append(c.cmds, host)
	return replicaofhost{cmds: cmds}
}

func Replicaof() (c replicaof) {
	c.cmds = append(c.cmds, "REPLICAOF")
	return
}

type replicaofhost struct {
	cmds []string
}

func (c replicaofhost) Port(port string) replicaofport {
	var cmds []string
	cmds = append(c.cmds, port)
	return replicaofport{cmds: cmds}
}

type replicaofport struct {
	cmds []string
}

func (c replicaofport) Build() []string {
	return c.cmds
}

type reset struct {
	cmds []string
}

func (c reset) Build() []string {
	return c.cmds
}

func Reset() (c reset) {
	c.cmds = append(c.cmds, "RESET")
	return
}

type restore struct {
	cmds []string
}

func (c restore) Key(key string) restorekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return restorekey{cmds: cmds}
}

func Restore() (c restore) {
	c.cmds = append(c.cmds, "RESTORE")
	return
}

type restoreabsttlabsttl struct {
	cmds []string
}

func (c restoreabsttlabsttl) Idletime(seconds int64) restoreidletime {
	var cmds []string
	cmds = append(c.cmds, "IDLETIME", strconv.FormatInt(seconds, 10))
	return restoreidletime{cmds: cmds}
}

func (c restoreabsttlabsttl) Freq(frequency int64) restorefreq {
	var cmds []string
	cmds = append(c.cmds, "FREQ", strconv.FormatInt(frequency, 10))
	return restorefreq{cmds: cmds}
}

func (c restoreabsttlabsttl) Build() []string {
	return c.cmds
}

type restorefreq struct {
	cmds []string
}

func (c restorefreq) Build() []string {
	return c.cmds
}

type restoreidletime struct {
	cmds []string
}

func (c restoreidletime) Freq(frequency int64) restorefreq {
	var cmds []string
	cmds = append(c.cmds, "FREQ", strconv.FormatInt(frequency, 10))
	return restorefreq{cmds: cmds}
}

func (c restoreidletime) Build() []string {
	return c.cmds
}

type restorekey struct {
	cmds []string
}

func (c restorekey) Ttl(ttl int64) restorettl {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(ttl, 10))
	return restorettl{cmds: cmds}
}

type restorereplacereplace struct {
	cmds []string
}

func (c restorereplacereplace) Absttl() restoreabsttlabsttl {
	var cmds []string
	cmds = append(c.cmds, "ABSTTL")
	return restoreabsttlabsttl{cmds: cmds}
}

func (c restorereplacereplace) Idletime(seconds int64) restoreidletime {
	var cmds []string
	cmds = append(c.cmds, "IDLETIME", strconv.FormatInt(seconds, 10))
	return restoreidletime{cmds: cmds}
}

func (c restorereplacereplace) Freq(frequency int64) restorefreq {
	var cmds []string
	cmds = append(c.cmds, "FREQ", strconv.FormatInt(frequency, 10))
	return restorefreq{cmds: cmds}
}

func (c restorereplacereplace) Build() []string {
	return c.cmds
}

type restoreserializedvalue struct {
	cmds []string
}

func (c restoreserializedvalue) Replace() restorereplacereplace {
	var cmds []string
	cmds = append(c.cmds, "REPLACE")
	return restorereplacereplace{cmds: cmds}
}

func (c restoreserializedvalue) Absttl() restoreabsttlabsttl {
	var cmds []string
	cmds = append(c.cmds, "ABSTTL")
	return restoreabsttlabsttl{cmds: cmds}
}

func (c restoreserializedvalue) Idletime(seconds int64) restoreidletime {
	var cmds []string
	cmds = append(c.cmds, "IDLETIME", strconv.FormatInt(seconds, 10))
	return restoreidletime{cmds: cmds}
}

func (c restoreserializedvalue) Freq(frequency int64) restorefreq {
	var cmds []string
	cmds = append(c.cmds, "FREQ", strconv.FormatInt(frequency, 10))
	return restorefreq{cmds: cmds}
}

func (c restoreserializedvalue) Build() []string {
	return c.cmds
}

type restorettl struct {
	cmds []string
}

func (c restorettl) Serializedvalue(serializedvalue string) restoreserializedvalue {
	var cmds []string
	cmds = append(c.cmds, serializedvalue)
	return restoreserializedvalue{cmds: cmds}
}

type role struct {
	cmds []string
}

func (c role) Build() []string {
	return c.cmds
}

func Role() (c role) {
	c.cmds = append(c.cmds, "ROLE")
	return
}

type rpop struct {
	cmds []string
}

func (c rpop) Key(key string) rpopkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return rpopkey{cmds: cmds}
}

func Rpop() (c rpop) {
	c.cmds = append(c.cmds, "RPOP")
	return
}

type rpopcount struct {
	cmds []string
}

func (c rpopcount) Build() []string {
	return c.cmds
}

type rpopkey struct {
	cmds []string
}

func (c rpopkey) Count(count int64) rpopcount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return rpopcount{cmds: cmds}
}

func (c rpopkey) Build() []string {
	return c.cmds
}

type rpoplpush struct {
	cmds []string
}

func (c rpoplpush) Source(source string) rpoplpushsource {
	var cmds []string
	cmds = append(c.cmds, source)
	return rpoplpushsource{cmds: cmds}
}

func Rpoplpush() (c rpoplpush) {
	c.cmds = append(c.cmds, "RPOPLPUSH")
	return
}

type rpoplpushdestination struct {
	cmds []string
}

func (c rpoplpushdestination) Build() []string {
	return c.cmds
}

type rpoplpushsource struct {
	cmds []string
}

func (c rpoplpushsource) Destination(destination string) rpoplpushdestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return rpoplpushdestination{cmds: cmds}
}

type rpush struct {
	cmds []string
}

func (c rpush) Key(key string) rpushkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return rpushkey{cmds: cmds}
}

func Rpush() (c rpush) {
	c.cmds = append(c.cmds, "RPUSH")
	return
}

type rpushelement struct {
	cmds []string
}

func (c rpushelement) Element(element ...string) rpushelement {
	var cmds []string
	cmds = append(cmds, element...)
	return rpushelement{cmds: cmds}
}

type rpushkey struct {
	cmds []string
}

func (c rpushkey) Element(element ...string) rpushelement {
	var cmds []string
	cmds = append(cmds, element...)
	return rpushelement{cmds: cmds}
}

type rpushx struct {
	cmds []string
}

func (c rpushx) Key(key string) rpushxkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return rpushxkey{cmds: cmds}
}

func Rpushx() (c rpushx) {
	c.cmds = append(c.cmds, "RPUSHX")
	return
}

type rpushxelement struct {
	cmds []string
}

func (c rpushxelement) Element(element ...string) rpushxelement {
	var cmds []string
	cmds = append(cmds, element...)
	return rpushxelement{cmds: cmds}
}

type rpushxkey struct {
	cmds []string
}

func (c rpushxkey) Element(element ...string) rpushxelement {
	var cmds []string
	cmds = append(cmds, element...)
	return rpushxelement{cmds: cmds}
}

type sadd struct {
	cmds []string
}

func (c sadd) Key(key string) saddkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return saddkey{cmds: cmds}
}

func Sadd() (c sadd) {
	c.cmds = append(c.cmds, "SADD")
	return
}

type saddkey struct {
	cmds []string
}

func (c saddkey) Member(member ...string) saddmember {
	var cmds []string
	cmds = append(cmds, member...)
	return saddmember{cmds: cmds}
}

type saddmember struct {
	cmds []string
}

func (c saddmember) Member(member ...string) saddmember {
	var cmds []string
	cmds = append(cmds, member...)
	return saddmember{cmds: cmds}
}

type save struct {
	cmds []string
}

func (c save) Build() []string {
	return c.cmds
}

func Save() (c save) {
	c.cmds = append(c.cmds, "SAVE")
	return
}

type scan struct {
	cmds []string
}

func (c scan) Cursor(cursor int64) scancursor {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(cursor, 10))
	return scancursor{cmds: cmds}
}

func Scan() (c scan) {
	c.cmds = append(c.cmds, "SCAN")
	return
}

type scancount struct {
	cmds []string
}

func (c scancount) Type(typ string) scantype {
	var cmds []string
	cmds = append(c.cmds, "TYPE", typ)
	return scantype{cmds: cmds}
}

func (c scancount) Build() []string {
	return c.cmds
}

type scancursor struct {
	cmds []string
}

func (c scancursor) Match(pattern string) scanmatch {
	var cmds []string
	cmds = append(c.cmds, "MATCH", pattern)
	return scanmatch{cmds: cmds}
}

func (c scancursor) Count(count int64) scancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return scancount{cmds: cmds}
}

func (c scancursor) Type(typ string) scantype {
	var cmds []string
	cmds = append(c.cmds, "TYPE", typ)
	return scantype{cmds: cmds}
}

func (c scancursor) Build() []string {
	return c.cmds
}

type scanmatch struct {
	cmds []string
}

func (c scanmatch) Count(count int64) scancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return scancount{cmds: cmds}
}

func (c scanmatch) Type(typ string) scantype {
	var cmds []string
	cmds = append(c.cmds, "TYPE", typ)
	return scantype{cmds: cmds}
}

func (c scanmatch) Build() []string {
	return c.cmds
}

type scantype struct {
	cmds []string
}

func (c scantype) Build() []string {
	return c.cmds
}

type scard struct {
	cmds []string
}

func (c scard) Key(key string) scardkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return scardkey{cmds: cmds}
}

func Scard() (c scard) {
	c.cmds = append(c.cmds, "SCARD")
	return
}

type scardkey struct {
	cmds []string
}

func (c scardkey) Build() []string {
	return c.cmds
}

type scriptdebug struct {
	cmds []string
}

func (c scriptdebug) Yes() scriptdebugmodeyes {
	var cmds []string
	cmds = append(c.cmds, "YES")
	return scriptdebugmodeyes{cmds: cmds}
}

func (c scriptdebug) Sync() scriptdebugmodesync {
	var cmds []string
	cmds = append(c.cmds, "SYNC")
	return scriptdebugmodesync{cmds: cmds}
}

func (c scriptdebug) No() scriptdebugmodeno {
	var cmds []string
	cmds = append(c.cmds, "NO")
	return scriptdebugmodeno{cmds: cmds}
}

func ScriptDebug() (c scriptdebug) {
	c.cmds = append(c.cmds, "SCRIPT", "DEBUG")
	return
}

type scriptdebugmodeno struct {
	cmds []string
}

func (c scriptdebugmodeno) Build() []string {
	return c.cmds
}

type scriptdebugmodesync struct {
	cmds []string
}

func (c scriptdebugmodesync) Build() []string {
	return c.cmds
}

type scriptdebugmodeyes struct {
	cmds []string
}

func (c scriptdebugmodeyes) Build() []string {
	return c.cmds
}

type scriptexists struct {
	cmds []string
}

func (c scriptexists) Sha1(sha1 ...string) scriptexistssha1 {
	var cmds []string
	cmds = append(cmds, sha1...)
	return scriptexistssha1{cmds: cmds}
}

func ScriptExists() (c scriptexists) {
	c.cmds = append(c.cmds, "SCRIPT", "EXISTS")
	return
}

type scriptexistssha1 struct {
	cmds []string
}

func (c scriptexistssha1) Sha1(sha1 ...string) scriptexistssha1 {
	var cmds []string
	cmds = append(cmds, sha1...)
	return scriptexistssha1{cmds: cmds}
}

type scriptflush struct {
	cmds []string
}

func (c scriptflush) Async() scriptflushasyncasync {
	var cmds []string
	cmds = append(c.cmds, "ASYNC")
	return scriptflushasyncasync{cmds: cmds}
}

func (c scriptflush) Sync() scriptflushasyncsync {
	var cmds []string
	cmds = append(c.cmds, "SYNC")
	return scriptflushasyncsync{cmds: cmds}
}

func (c scriptflush) Build() []string {
	return c.cmds
}

func ScriptFlush() (c scriptflush) {
	c.cmds = append(c.cmds, "SCRIPT", "FLUSH")
	return
}

type scriptflushasyncasync struct {
	cmds []string
}

func (c scriptflushasyncasync) Build() []string {
	return c.cmds
}

type scriptflushasyncsync struct {
	cmds []string
}

func (c scriptflushasyncsync) Build() []string {
	return c.cmds
}

type scriptkill struct {
	cmds []string
}

func (c scriptkill) Build() []string {
	return c.cmds
}

func ScriptKill() (c scriptkill) {
	c.cmds = append(c.cmds, "SCRIPT", "KILL")
	return
}

type scriptload struct {
	cmds []string
}

func (c scriptload) Script(script string) scriptloadscript {
	var cmds []string
	cmds = append(c.cmds, script)
	return scriptloadscript{cmds: cmds}
}

func ScriptLoad() (c scriptload) {
	c.cmds = append(c.cmds, "SCRIPT", "LOAD")
	return
}

type scriptloadscript struct {
	cmds []string
}

func (c scriptloadscript) Build() []string {
	return c.cmds
}

type sdiff struct {
	cmds []string
}

func (c sdiff) Key(key ...string) sdiffkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sdiffkey{cmds: cmds}
}

func Sdiff() (c sdiff) {
	c.cmds = append(c.cmds, "SDIFF")
	return
}

type sdiffkey struct {
	cmds []string
}

func (c sdiffkey) Key(key ...string) sdiffkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sdiffkey{cmds: cmds}
}

type sdiffstore struct {
	cmds []string
}

func (c sdiffstore) Destination(destination string) sdiffstoredestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return sdiffstoredestination{cmds: cmds}
}

func Sdiffstore() (c sdiffstore) {
	c.cmds = append(c.cmds, "SDIFFSTORE")
	return
}

type sdiffstoredestination struct {
	cmds []string
}

func (c sdiffstoredestination) Key(key ...string) sdiffstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return sdiffstorekey{cmds: cmds}
}

type sdiffstorekey struct {
	cmds []string
}

func (c sdiffstorekey) Key(key ...string) sdiffstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return sdiffstorekey{cmds: cmds}
}

type rselect struct {
	cmds []string
}

func (c rselect) Index(index int64) selectindex {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(index, 10))
	return selectindex{cmds: cmds}
}

func Select() (c rselect) {
	c.cmds = append(c.cmds, "SELECT")
	return
}

type selectindex struct {
	cmds []string
}

func (c selectindex) Build() []string {
	return c.cmds
}

type set struct {
	cmds []string
}

func (c set) Key(key string) setkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return setkey{cmds: cmds}
}

func Set() (c set) {
	c.cmds = append(c.cmds, "SET")
	return
}

type setbit struct {
	cmds []string
}

func (c setbit) Key(key string) setbitkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return setbitkey{cmds: cmds}
}

func Setbit() (c setbit) {
	c.cmds = append(c.cmds, "SETBIT")
	return
}

type setbitkey struct {
	cmds []string
}

func (c setbitkey) Offset(offset int64) setbitoffset {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(offset, 10))
	return setbitoffset{cmds: cmds}
}

type setbitoffset struct {
	cmds []string
}

func (c setbitoffset) Value(value int64) setbitvalue {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(value, 10))
	return setbitvalue{cmds: cmds}
}

type setbitvalue struct {
	cmds []string
}

func (c setbitvalue) Build() []string {
	return c.cmds
}

type setconditionnx struct {
	cmds []string
}

func (c setconditionnx) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setconditionnx) Build() []string {
	return c.cmds
}

type setconditionxx struct {
	cmds []string
}

func (c setconditionxx) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setconditionxx) Build() []string {
	return c.cmds
}

type setex struct {
	cmds []string
}

func (c setex) Key(key string) setexkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return setexkey{cmds: cmds}
}

func Setex() (c setex) {
	c.cmds = append(c.cmds, "SETEX")
	return
}

type setexkey struct {
	cmds []string
}

func (c setexkey) Seconds(seconds int64) setexseconds {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(seconds, 10))
	return setexseconds{cmds: cmds}
}

type setexpirationex struct {
	cmds []string
}

func (c setexpirationex) Nx() setconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return setconditionnx{cmds: cmds}
}

func (c setexpirationex) Xx() setconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return setconditionxx{cmds: cmds}
}

func (c setexpirationex) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setexpirationex) Build() []string {
	return c.cmds
}

type setexpirationexat struct {
	cmds []string
}

func (c setexpirationexat) Nx() setconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return setconditionnx{cmds: cmds}
}

func (c setexpirationexat) Xx() setconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return setconditionxx{cmds: cmds}
}

func (c setexpirationexat) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setexpirationexat) Build() []string {
	return c.cmds
}

type setexpirationkeepttl struct {
	cmds []string
}

func (c setexpirationkeepttl) Nx() setconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return setconditionnx{cmds: cmds}
}

func (c setexpirationkeepttl) Xx() setconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return setconditionxx{cmds: cmds}
}

func (c setexpirationkeepttl) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setexpirationkeepttl) Build() []string {
	return c.cmds
}

type setexpirationpx struct {
	cmds []string
}

func (c setexpirationpx) Nx() setconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return setconditionnx{cmds: cmds}
}

func (c setexpirationpx) Xx() setconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return setconditionxx{cmds: cmds}
}

func (c setexpirationpx) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setexpirationpx) Build() []string {
	return c.cmds
}

type setexpirationpxat struct {
	cmds []string
}

func (c setexpirationpxat) Nx() setconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return setconditionnx{cmds: cmds}
}

func (c setexpirationpxat) Xx() setconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return setconditionxx{cmds: cmds}
}

func (c setexpirationpxat) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setexpirationpxat) Build() []string {
	return c.cmds
}

type setexseconds struct {
	cmds []string
}

func (c setexseconds) Value(value string) setexvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return setexvalue{cmds: cmds}
}

type setexvalue struct {
	cmds []string
}

func (c setexvalue) Build() []string {
	return c.cmds
}

type setgetget struct {
	cmds []string
}

func (c setgetget) Build() []string {
	return c.cmds
}

type setkey struct {
	cmds []string
}

func (c setkey) Value(value string) setvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return setvalue{cmds: cmds}
}

type setnx struct {
	cmds []string
}

func (c setnx) Key(key string) setnxkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return setnxkey{cmds: cmds}
}

func Setnx() (c setnx) {
	c.cmds = append(c.cmds, "SETNX")
	return
}

type setnxkey struct {
	cmds []string
}

func (c setnxkey) Value(value string) setnxvalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return setnxvalue{cmds: cmds}
}

type setnxvalue struct {
	cmds []string
}

func (c setnxvalue) Build() []string {
	return c.cmds
}

type setrange struct {
	cmds []string
}

func (c setrange) Key(key string) setrangekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return setrangekey{cmds: cmds}
}

func Setrange() (c setrange) {
	c.cmds = append(c.cmds, "SETRANGE")
	return
}

type setrangekey struct {
	cmds []string
}

func (c setrangekey) Offset(offset int64) setrangeoffset {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(offset, 10))
	return setrangeoffset{cmds: cmds}
}

type setrangeoffset struct {
	cmds []string
}

func (c setrangeoffset) Value(value string) setrangevalue {
	var cmds []string
	cmds = append(c.cmds, value)
	return setrangevalue{cmds: cmds}
}

type setrangevalue struct {
	cmds []string
}

func (c setrangevalue) Build() []string {
	return c.cmds
}

type setvalue struct {
	cmds []string
}

func (c setvalue) Ex(seconds int64) setexpirationex {
	var cmds []string
	cmds = append(c.cmds, "EX", strconv.FormatInt(seconds, 10))
	return setexpirationex{cmds: cmds}
}

func (c setvalue) Px(milliseconds int64) setexpirationpx {
	var cmds []string
	cmds = append(c.cmds, "PX", strconv.FormatInt(milliseconds, 10))
	return setexpirationpx{cmds: cmds}
}

func (c setvalue) Exat(timestamp int64) setexpirationexat {
	var cmds []string
	cmds = append(c.cmds, "EXAT", strconv.FormatInt(timestamp, 10))
	return setexpirationexat{cmds: cmds}
}

func (c setvalue) Pxat(millisecondstimestamp int64) setexpirationpxat {
	var cmds []string
	cmds = append(c.cmds, "PXAT", strconv.FormatInt(millisecondstimestamp, 10))
	return setexpirationpxat{cmds: cmds}
}

func (c setvalue) Keepttl() setexpirationkeepttl {
	var cmds []string
	cmds = append(c.cmds, "KEEPTTL")
	return setexpirationkeepttl{cmds: cmds}
}

func (c setvalue) Nx() setconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return setconditionnx{cmds: cmds}
}

func (c setvalue) Xx() setconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return setconditionxx{cmds: cmds}
}

func (c setvalue) Get() setgetget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	return setgetget{cmds: cmds}
}

func (c setvalue) Build() []string {
	return c.cmds
}

type shutdown struct {
	cmds []string
}

func (c shutdown) Nosave() shutdownsavemodenosave {
	var cmds []string
	cmds = append(c.cmds, "NOSAVE")
	return shutdownsavemodenosave{cmds: cmds}
}

func (c shutdown) Save() shutdownsavemodesave {
	var cmds []string
	cmds = append(c.cmds, "SAVE")
	return shutdownsavemodesave{cmds: cmds}
}

func (c shutdown) Build() []string {
	return c.cmds
}

func Shutdown() (c shutdown) {
	c.cmds = append(c.cmds, "SHUTDOWN")
	return
}

type shutdownsavemodenosave struct {
	cmds []string
}

func (c shutdownsavemodenosave) Build() []string {
	return c.cmds
}

type shutdownsavemodesave struct {
	cmds []string
}

func (c shutdownsavemodesave) Build() []string {
	return c.cmds
}

type sinter struct {
	cmds []string
}

func (c sinter) Key(key ...string) sinterkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sinterkey{cmds: cmds}
}

func Sinter() (c sinter) {
	c.cmds = append(c.cmds, "SINTER")
	return
}

type sintercard struct {
	cmds []string
}

func (c sintercard) Key(key ...string) sintercardkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sintercardkey{cmds: cmds}
}

func Sintercard() (c sintercard) {
	c.cmds = append(c.cmds, "SINTERCARD")
	return
}

type sintercardkey struct {
	cmds []string
}

func (c sintercardkey) Key(key ...string) sintercardkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sintercardkey{cmds: cmds}
}

type sinterkey struct {
	cmds []string
}

func (c sinterkey) Key(key ...string) sinterkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sinterkey{cmds: cmds}
}

type sinterstore struct {
	cmds []string
}

func (c sinterstore) Destination(destination string) sinterstoredestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return sinterstoredestination{cmds: cmds}
}

func Sinterstore() (c sinterstore) {
	c.cmds = append(c.cmds, "SINTERSTORE")
	return
}

type sinterstoredestination struct {
	cmds []string
}

func (c sinterstoredestination) Key(key ...string) sinterstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return sinterstorekey{cmds: cmds}
}

type sinterstorekey struct {
	cmds []string
}

func (c sinterstorekey) Key(key ...string) sinterstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return sinterstorekey{cmds: cmds}
}

type sismember struct {
	cmds []string
}

func (c sismember) Key(key string) sismemberkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return sismemberkey{cmds: cmds}
}

func Sismember() (c sismember) {
	c.cmds = append(c.cmds, "SISMEMBER")
	return
}

type sismemberkey struct {
	cmds []string
}

func (c sismemberkey) Member(member string) sismembermember {
	var cmds []string
	cmds = append(c.cmds, member)
	return sismembermember{cmds: cmds}
}

type sismembermember struct {
	cmds []string
}

func (c sismembermember) Build() []string {
	return c.cmds
}

type slaveof struct {
	cmds []string
}

func (c slaveof) Host(host string) slaveofhost {
	var cmds []string
	cmds = append(c.cmds, host)
	return slaveofhost{cmds: cmds}
}

func Slaveof() (c slaveof) {
	c.cmds = append(c.cmds, "SLAVEOF")
	return
}

type slaveofhost struct {
	cmds []string
}

func (c slaveofhost) Port(port string) slaveofport {
	var cmds []string
	cmds = append(c.cmds, port)
	return slaveofport{cmds: cmds}
}

type slaveofport struct {
	cmds []string
}

func (c slaveofport) Build() []string {
	return c.cmds
}

type slowlog struct {
	cmds []string
}

func (c slowlog) Subcommand(subcommand string) slowlogsubcommand {
	var cmds []string
	cmds = append(c.cmds, subcommand)
	return slowlogsubcommand{cmds: cmds}
}

func Slowlog() (c slowlog) {
	c.cmds = append(c.cmds, "SLOWLOG")
	return
}

type slowlogargument struct {
	cmds []string
}

func (c slowlogargument) Build() []string {
	return c.cmds
}

type slowlogsubcommand struct {
	cmds []string
}

func (c slowlogsubcommand) Argument(argument string) slowlogargument {
	var cmds []string
	cmds = append(c.cmds, argument)
	return slowlogargument{cmds: cmds}
}

func (c slowlogsubcommand) Build() []string {
	return c.cmds
}

type smembers struct {
	cmds []string
}

func (c smembers) Key(key string) smemberskey {
	var cmds []string
	cmds = append(c.cmds, key)
	return smemberskey{cmds: cmds}
}

func Smembers() (c smembers) {
	c.cmds = append(c.cmds, "SMEMBERS")
	return
}

type smemberskey struct {
	cmds []string
}

func (c smemberskey) Build() []string {
	return c.cmds
}

type smismember struct {
	cmds []string
}

func (c smismember) Key(key string) smismemberkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return smismemberkey{cmds: cmds}
}

func Smismember() (c smismember) {
	c.cmds = append(c.cmds, "SMISMEMBER")
	return
}

type smismemberkey struct {
	cmds []string
}

func (c smismemberkey) Member(member ...string) smismembermember {
	var cmds []string
	cmds = append(cmds, member...)
	return smismembermember{cmds: cmds}
}

type smismembermember struct {
	cmds []string
}

func (c smismembermember) Member(member ...string) smismembermember {
	var cmds []string
	cmds = append(cmds, member...)
	return smismembermember{cmds: cmds}
}

type smove struct {
	cmds []string
}

func (c smove) Source(source string) smovesource {
	var cmds []string
	cmds = append(c.cmds, source)
	return smovesource{cmds: cmds}
}

func Smove() (c smove) {
	c.cmds = append(c.cmds, "SMOVE")
	return
}

type smovedestination struct {
	cmds []string
}

func (c smovedestination) Member(member string) smovemember {
	var cmds []string
	cmds = append(c.cmds, member)
	return smovemember{cmds: cmds}
}

type smovemember struct {
	cmds []string
}

func (c smovemember) Build() []string {
	return c.cmds
}

type smovesource struct {
	cmds []string
}

func (c smovesource) Destination(destination string) smovedestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return smovedestination{cmds: cmds}
}

type sort struct {
	cmds []string
}

func (c sort) Key(key string) sortkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return sortkey{cmds: cmds}
}

func Sort() (c sort) {
	c.cmds = append(c.cmds, "SORT")
	return
}

type sortby struct {
	cmds []string
}

func (c sortby) Limit(offset int64, count int64) sortlimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return sortlimit{cmds: cmds}
}

func (c sortby) Get(pattern ...string) sortget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	cmds = append(cmds, pattern...)
	return sortget{cmds: cmds}
}

func (c sortby) Asc() sortorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortorderasc{cmds: cmds}
}

func (c sortby) Desc() sortorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortorderdesc{cmds: cmds}
}

func (c sortby) Alpha() sortsortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortsortingalpha{cmds: cmds}
}

func (c sortby) Store(destination string) sortstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", destination)
	return sortstore{cmds: cmds}
}

func (c sortby) Build() []string {
	return c.cmds
}

type sortget struct {
	cmds []string
}

func (c sortget) Asc() sortorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortorderasc{cmds: cmds}
}

func (c sortget) Desc() sortorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortorderdesc{cmds: cmds}
}

func (c sortget) Alpha() sortsortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortsortingalpha{cmds: cmds}
}

func (c sortget) Store(destination string) sortstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", destination)
	return sortstore{cmds: cmds}
}

func (c sortget) Pattern(pattern ...string) sortget {
	var cmds []string
	cmds = append(cmds, pattern...)
	return sortget{cmds: cmds}
}

func (c sortget) Build() []string {
	return c.cmds
}

type sortkey struct {
	cmds []string
}

func (c sortkey) By(pattern string) sortby {
	var cmds []string
	cmds = append(c.cmds, "BY", pattern)
	return sortby{cmds: cmds}
}

func (c sortkey) Limit(offset int64, count int64) sortlimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return sortlimit{cmds: cmds}
}

func (c sortkey) Get(pattern ...string) sortget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	cmds = append(cmds, pattern...)
	return sortget{cmds: cmds}
}

func (c sortkey) Asc() sortorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortorderasc{cmds: cmds}
}

func (c sortkey) Desc() sortorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortorderdesc{cmds: cmds}
}

func (c sortkey) Alpha() sortsortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortsortingalpha{cmds: cmds}
}

func (c sortkey) Store(destination string) sortstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", destination)
	return sortstore{cmds: cmds}
}

func (c sortkey) Build() []string {
	return c.cmds
}

type sortlimit struct {
	cmds []string
}

func (c sortlimit) Get(pattern ...string) sortget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	cmds = append(cmds, pattern...)
	return sortget{cmds: cmds}
}

func (c sortlimit) Asc() sortorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortorderasc{cmds: cmds}
}

func (c sortlimit) Desc() sortorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortorderdesc{cmds: cmds}
}

func (c sortlimit) Alpha() sortsortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortsortingalpha{cmds: cmds}
}

func (c sortlimit) Store(destination string) sortstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", destination)
	return sortstore{cmds: cmds}
}

func (c sortlimit) Build() []string {
	return c.cmds
}

type sortorderasc struct {
	cmds []string
}

func (c sortorderasc) Alpha() sortsortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortsortingalpha{cmds: cmds}
}

func (c sortorderasc) Store(destination string) sortstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", destination)
	return sortstore{cmds: cmds}
}

func (c sortorderasc) Build() []string {
	return c.cmds
}

type sortorderdesc struct {
	cmds []string
}

func (c sortorderdesc) Alpha() sortsortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortsortingalpha{cmds: cmds}
}

func (c sortorderdesc) Store(destination string) sortstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", destination)
	return sortstore{cmds: cmds}
}

func (c sortorderdesc) Build() []string {
	return c.cmds
}

type sortro struct {
	cmds []string
}

func (c sortro) Key(key string) sortrokey {
	var cmds []string
	cmds = append(c.cmds, key)
	return sortrokey{cmds: cmds}
}

func Sortro() (c sortro) {
	c.cmds = append(c.cmds, "SORT_RO")
	return
}

type sortroby struct {
	cmds []string
}

func (c sortroby) Limit(offset int64, count int64) sortrolimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return sortrolimit{cmds: cmds}
}

func (c sortroby) Get(pattern ...string) sortroget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	cmds = append(cmds, pattern...)
	return sortroget{cmds: cmds}
}

func (c sortroby) Asc() sortroorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortroorderasc{cmds: cmds}
}

func (c sortroby) Desc() sortroorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortroorderdesc{cmds: cmds}
}

func (c sortroby) Alpha() sortrosortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortrosortingalpha{cmds: cmds}
}

func (c sortroby) Build() []string {
	return c.cmds
}

type sortroget struct {
	cmds []string
}

func (c sortroget) Asc() sortroorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortroorderasc{cmds: cmds}
}

func (c sortroget) Desc() sortroorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortroorderdesc{cmds: cmds}
}

func (c sortroget) Alpha() sortrosortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortrosortingalpha{cmds: cmds}
}

func (c sortroget) Pattern(pattern ...string) sortroget {
	var cmds []string
	cmds = append(cmds, pattern...)
	return sortroget{cmds: cmds}
}

func (c sortroget) Build() []string {
	return c.cmds
}

type sortrokey struct {
	cmds []string
}

func (c sortrokey) By(pattern string) sortroby {
	var cmds []string
	cmds = append(c.cmds, "BY", pattern)
	return sortroby{cmds: cmds}
}

func (c sortrokey) Limit(offset int64, count int64) sortrolimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return sortrolimit{cmds: cmds}
}

func (c sortrokey) Get(pattern ...string) sortroget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	cmds = append(cmds, pattern...)
	return sortroget{cmds: cmds}
}

func (c sortrokey) Asc() sortroorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortroorderasc{cmds: cmds}
}

func (c sortrokey) Desc() sortroorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortroorderdesc{cmds: cmds}
}

func (c sortrokey) Alpha() sortrosortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortrosortingalpha{cmds: cmds}
}

func (c sortrokey) Build() []string {
	return c.cmds
}

type sortrolimit struct {
	cmds []string
}

func (c sortrolimit) Get(pattern ...string) sortroget {
	var cmds []string
	cmds = append(c.cmds, "GET")
	cmds = append(cmds, pattern...)
	return sortroget{cmds: cmds}
}

func (c sortrolimit) Asc() sortroorderasc {
	var cmds []string
	cmds = append(c.cmds, "ASC")
	return sortroorderasc{cmds: cmds}
}

func (c sortrolimit) Desc() sortroorderdesc {
	var cmds []string
	cmds = append(c.cmds, "DESC")
	return sortroorderdesc{cmds: cmds}
}

func (c sortrolimit) Alpha() sortrosortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortrosortingalpha{cmds: cmds}
}

func (c sortrolimit) Build() []string {
	return c.cmds
}

type sortroorderasc struct {
	cmds []string
}

func (c sortroorderasc) Alpha() sortrosortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortrosortingalpha{cmds: cmds}
}

func (c sortroorderasc) Build() []string {
	return c.cmds
}

type sortroorderdesc struct {
	cmds []string
}

func (c sortroorderdesc) Alpha() sortrosortingalpha {
	var cmds []string
	cmds = append(c.cmds, "ALPHA")
	return sortrosortingalpha{cmds: cmds}
}

func (c sortroorderdesc) Build() []string {
	return c.cmds
}

type sortrosortingalpha struct {
	cmds []string
}

func (c sortrosortingalpha) Build() []string {
	return c.cmds
}

type sortsortingalpha struct {
	cmds []string
}

func (c sortsortingalpha) Store(destination string) sortstore {
	var cmds []string
	cmds = append(c.cmds, "STORE", destination)
	return sortstore{cmds: cmds}
}

func (c sortsortingalpha) Build() []string {
	return c.cmds
}

type sortstore struct {
	cmds []string
}

func (c sortstore) Build() []string {
	return c.cmds
}

type spop struct {
	cmds []string
}

func (c spop) Key(key string) spopkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return spopkey{cmds: cmds}
}

func Spop() (c spop) {
	c.cmds = append(c.cmds, "SPOP")
	return
}

type spopcount struct {
	cmds []string
}

func (c spopcount) Build() []string {
	return c.cmds
}

type spopkey struct {
	cmds []string
}

func (c spopkey) Count(count int64) spopcount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return spopcount{cmds: cmds}
}

func (c spopkey) Build() []string {
	return c.cmds
}

type srandmember struct {
	cmds []string
}

func (c srandmember) Key(key string) srandmemberkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return srandmemberkey{cmds: cmds}
}

func Srandmember() (c srandmember) {
	c.cmds = append(c.cmds, "SRANDMEMBER")
	return
}

type srandmembercount struct {
	cmds []string
}

func (c srandmembercount) Build() []string {
	return c.cmds
}

type srandmemberkey struct {
	cmds []string
}

func (c srandmemberkey) Count(count int64) srandmembercount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return srandmembercount{cmds: cmds}
}

func (c srandmemberkey) Build() []string {
	return c.cmds
}

type srem struct {
	cmds []string
}

func (c srem) Key(key string) sremkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return sremkey{cmds: cmds}
}

func Srem() (c srem) {
	c.cmds = append(c.cmds, "SREM")
	return
}

type sremkey struct {
	cmds []string
}

func (c sremkey) Member(member ...string) sremmember {
	var cmds []string
	cmds = append(cmds, member...)
	return sremmember{cmds: cmds}
}

type sremmember struct {
	cmds []string
}

func (c sremmember) Member(member ...string) sremmember {
	var cmds []string
	cmds = append(cmds, member...)
	return sremmember{cmds: cmds}
}

type sscan struct {
	cmds []string
}

func (c sscan) Key(key string) sscankey {
	var cmds []string
	cmds = append(c.cmds, key)
	return sscankey{cmds: cmds}
}

func Sscan() (c sscan) {
	c.cmds = append(c.cmds, "SSCAN")
	return
}

type sscancount struct {
	cmds []string
}

func (c sscancount) Build() []string {
	return c.cmds
}

type sscancursor struct {
	cmds []string
}

func (c sscancursor) Match(pattern string) sscanmatch {
	var cmds []string
	cmds = append(c.cmds, "MATCH", pattern)
	return sscanmatch{cmds: cmds}
}

func (c sscancursor) Count(count int64) sscancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return sscancount{cmds: cmds}
}

func (c sscancursor) Build() []string {
	return c.cmds
}

type sscankey struct {
	cmds []string
}

func (c sscankey) Cursor(cursor int64) sscancursor {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(cursor, 10))
	return sscancursor{cmds: cmds}
}

type sscanmatch struct {
	cmds []string
}

func (c sscanmatch) Count(count int64) sscancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return sscancount{cmds: cmds}
}

func (c sscanmatch) Build() []string {
	return c.cmds
}

type stralgo struct {
	cmds []string
}

func (c stralgo) Lcs() stralgoalgorithmlcs {
	var cmds []string
	cmds = append(c.cmds, "LCS")
	return stralgoalgorithmlcs{cmds: cmds}
}

func Stralgo() (c stralgo) {
	c.cmds = append(c.cmds, "STRALGO")
	return
}

type stralgoalgorithmlcs struct {
	cmds []string
}

func (c stralgoalgorithmlcs) Algospecificargument(algospecificargument ...string) stralgoalgospecificargument {
	var cmds []string
	cmds = append(cmds, algospecificargument...)
	return stralgoalgospecificargument{cmds: cmds}
}

type stralgoalgospecificargument struct {
	cmds []string
}

func (c stralgoalgospecificargument) Algospecificargument(algospecificargument ...string) stralgoalgospecificargument {
	var cmds []string
	cmds = append(cmds, algospecificargument...)
	return stralgoalgospecificargument{cmds: cmds}
}

type strlen struct {
	cmds []string
}

func (c strlen) Key(key string) strlenkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return strlenkey{cmds: cmds}
}

func Strlen() (c strlen) {
	c.cmds = append(c.cmds, "STRLEN")
	return
}

type strlenkey struct {
	cmds []string
}

func (c strlenkey) Build() []string {
	return c.cmds
}

type subscribe struct {
	cmds []string
}

func (c subscribe) Channel(channel ...string) subscribechannel {
	var cmds []string
	cmds = append(cmds, channel...)
	return subscribechannel{cmds: cmds}
}

func Subscribe() (c subscribe) {
	c.cmds = append(c.cmds, "SUBSCRIBE")
	return
}

type subscribechannel struct {
	cmds []string
}

func (c subscribechannel) Channel(channel ...string) subscribechannel {
	var cmds []string
	cmds = append(cmds, channel...)
	return subscribechannel{cmds: cmds}
}

type sunion struct {
	cmds []string
}

func (c sunion) Key(key ...string) sunionkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sunionkey{cmds: cmds}
}

func Sunion() (c sunion) {
	c.cmds = append(c.cmds, "SUNION")
	return
}

type sunionkey struct {
	cmds []string
}

func (c sunionkey) Key(key ...string) sunionkey {
	var cmds []string
	cmds = append(cmds, key...)
	return sunionkey{cmds: cmds}
}

type sunionstore struct {
	cmds []string
}

func (c sunionstore) Destination(destination string) sunionstoredestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return sunionstoredestination{cmds: cmds}
}

func Sunionstore() (c sunionstore) {
	c.cmds = append(c.cmds, "SUNIONSTORE")
	return
}

type sunionstoredestination struct {
	cmds []string
}

func (c sunionstoredestination) Key(key ...string) sunionstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return sunionstorekey{cmds: cmds}
}

type sunionstorekey struct {
	cmds []string
}

func (c sunionstorekey) Key(key ...string) sunionstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return sunionstorekey{cmds: cmds}
}

type swapdb struct {
	cmds []string
}

func (c swapdb) Index1(index1 int64) swapdbindex1 {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(index1, 10))
	return swapdbindex1{cmds: cmds}
}

func Swapdb() (c swapdb) {
	c.cmds = append(c.cmds, "SWAPDB")
	return
}

type swapdbindex1 struct {
	cmds []string
}

func (c swapdbindex1) Index2(index2 int64) swapdbindex2 {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(index2, 10))
	return swapdbindex2{cmds: cmds}
}

type swapdbindex2 struct {
	cmds []string
}

func (c swapdbindex2) Build() []string {
	return c.cmds
}

type sync struct {
	cmds []string
}

func (c sync) Build() []string {
	return c.cmds
}

func Sync() (c sync) {
	c.cmds = append(c.cmds, "SYNC")
	return
}

type time struct {
	cmds []string
}

func (c time) Build() []string {
	return c.cmds
}

func Time() (c time) {
	c.cmds = append(c.cmds, "TIME")
	return
}

type touch struct {
	cmds []string
}

func (c touch) Key(key ...string) touchkey {
	var cmds []string
	cmds = append(cmds, key...)
	return touchkey{cmds: cmds}
}

func Touch() (c touch) {
	c.cmds = append(c.cmds, "TOUCH")
	return
}

type touchkey struct {
	cmds []string
}

func (c touchkey) Key(key ...string) touchkey {
	var cmds []string
	cmds = append(cmds, key...)
	return touchkey{cmds: cmds}
}

type ttl struct {
	cmds []string
}

func (c ttl) Key(key string) ttlkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return ttlkey{cmds: cmds}
}

func Ttl() (c ttl) {
	c.cmds = append(c.cmds, "TTL")
	return
}

type ttlkey struct {
	cmds []string
}

func (c ttlkey) Build() []string {
	return c.cmds
}

type rtype struct {
	cmds []string
}

func (c rtype) Key(key string) typekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return typekey{cmds: cmds}
}

func Type() (c rtype) {
	c.cmds = append(c.cmds, "TYPE")
	return
}

type typekey struct {
	cmds []string
}

func (c typekey) Build() []string {
	return c.cmds
}

type unlink struct {
	cmds []string
}

func (c unlink) Key(key ...string) unlinkkey {
	var cmds []string
	cmds = append(cmds, key...)
	return unlinkkey{cmds: cmds}
}

func Unlink() (c unlink) {
	c.cmds = append(c.cmds, "UNLINK")
	return
}

type unlinkkey struct {
	cmds []string
}

func (c unlinkkey) Key(key ...string) unlinkkey {
	var cmds []string
	cmds = append(cmds, key...)
	return unlinkkey{cmds: cmds}
}

type unsubscribe struct {
	cmds []string
}

func (c unsubscribe) Channel(channel ...string) unsubscribechannel {
	var cmds []string
	cmds = append(cmds, channel...)
	return unsubscribechannel{cmds: cmds}
}

func (c unsubscribe) Build() []string {
	return c.cmds
}

func Unsubscribe() (c unsubscribe) {
	c.cmds = append(c.cmds, "UNSUBSCRIBE")
	return
}

type unsubscribechannel struct {
	cmds []string
}

func (c unsubscribechannel) Channel(channel ...string) unsubscribechannel {
	var cmds []string
	cmds = append(cmds, channel...)
	return unsubscribechannel{cmds: cmds}
}

func (c unsubscribechannel) Build() []string {
	return c.cmds
}

type unwatch struct {
	cmds []string
}

func (c unwatch) Build() []string {
	return c.cmds
}

func Unwatch() (c unwatch) {
	c.cmds = append(c.cmds, "UNWATCH")
	return
}

type wait struct {
	cmds []string
}

func (c wait) Numreplicas(numreplicas int64) waitnumreplicas {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numreplicas, 10))
	return waitnumreplicas{cmds: cmds}
}

func Wait() (c wait) {
	c.cmds = append(c.cmds, "WAIT")
	return
}

type waitnumreplicas struct {
	cmds []string
}

func (c waitnumreplicas) Timeout(timeout int64) waittimeout {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(timeout, 10))
	return waittimeout{cmds: cmds}
}

type waittimeout struct {
	cmds []string
}

func (c waittimeout) Build() []string {
	return c.cmds
}

type watch struct {
	cmds []string
}

func (c watch) Key(key ...string) watchkey {
	var cmds []string
	cmds = append(cmds, key...)
	return watchkey{cmds: cmds}
}

func Watch() (c watch) {
	c.cmds = append(c.cmds, "WATCH")
	return
}

type watchkey struct {
	cmds []string
}

func (c watchkey) Key(key ...string) watchkey {
	var cmds []string
	cmds = append(cmds, key...)
	return watchkey{cmds: cmds}
}

type xack struct {
	cmds []string
}

func (c xack) Key(key string) xackkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xackkey{cmds: cmds}
}

func Xack() (c xack) {
	c.cmds = append(c.cmds, "XACK")
	return
}

type xackgroup struct {
	cmds []string
}

func (c xackgroup) Id(id ...string) xackid {
	var cmds []string
	cmds = append(cmds, id...)
	return xackid{cmds: cmds}
}

type xackid struct {
	cmds []string
}

func (c xackid) Id(id ...string) xackid {
	var cmds []string
	cmds = append(cmds, id...)
	return xackid{cmds: cmds}
}

type xackkey struct {
	cmds []string
}

func (c xackkey) Group(group string) xackgroup {
	var cmds []string
	cmds = append(c.cmds, group)
	return xackgroup{cmds: cmds}
}

type xadd struct {
	cmds []string
}

func (c xadd) Key(key string) xaddkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xaddkey{cmds: cmds}
}

func Xadd() (c xadd) {
	c.cmds = append(c.cmds, "XADD")
	return
}

type xaddfieldvalue struct {
	cmds []string
}

func (c xaddfieldvalue) FieldValue(field string, value string) xaddfieldvalue {
	var cmds []string
	cmds = append(c.cmds, field, value)
	return xaddfieldvalue{cmds: cmds}
}

type xaddid struct {
	cmds []string
}

func (c xaddid) FieldValue(field string, value string) xaddfieldvalue {
	var cmds []string
	cmds = append(c.cmds, field, value)
	return xaddfieldvalue{cmds: cmds}
}

type xaddkey struct {
	cmds []string
}

func (c xaddkey) Nomkstream() xaddnomkstream {
	var cmds []string
	cmds = append(c.cmds, "NOMKSTREAM")
	return xaddnomkstream{cmds: cmds}
}

func (c xaddkey) Maxlen() xaddtrimstrategymaxlen {
	var cmds []string
	cmds = append(c.cmds, "MAXLEN")
	return xaddtrimstrategymaxlen{cmds: cmds}
}

func (c xaddkey) Minid() xaddtrimstrategyminid {
	var cmds []string
	cmds = append(c.cmds, "MINID")
	return xaddtrimstrategyminid{cmds: cmds}
}

func (c xaddkey) Wildcard() xaddwildcard {
	var cmds []string
	cmds = append(c.cmds, "*")
	return xaddwildcard{cmds: cmds}
}

func (c xaddkey) Id() xaddid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return xaddid{cmds: cmds}
}

type xaddnomkstream struct {
	cmds []string
}

func (c xaddnomkstream) Maxlen() xaddtrimstrategymaxlen {
	var cmds []string
	cmds = append(c.cmds, "MAXLEN")
	return xaddtrimstrategymaxlen{cmds: cmds}
}

func (c xaddnomkstream) Minid() xaddtrimstrategyminid {
	var cmds []string
	cmds = append(c.cmds, "MINID")
	return xaddtrimstrategyminid{cmds: cmds}
}

func (c xaddnomkstream) Wildcard() xaddwildcard {
	var cmds []string
	cmds = append(c.cmds, "*")
	return xaddwildcard{cmds: cmds}
}

func (c xaddnomkstream) Id() xaddid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return xaddid{cmds: cmds}
}

type xaddtrimlimit struct {
	cmds []string
}

func (c xaddtrimlimit) Wildcard() xaddwildcard {
	var cmds []string
	cmds = append(c.cmds, "*")
	return xaddwildcard{cmds: cmds}
}

func (c xaddtrimlimit) Id() xaddid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return xaddid{cmds: cmds}
}

type xaddtrimoperatoralmost struct {
	cmds []string
}

func (c xaddtrimoperatoralmost) Threshold(threshold string) xaddtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xaddtrimthreshold{cmds: cmds}
}

type xaddtrimoperatorexact struct {
	cmds []string
}

func (c xaddtrimoperatorexact) Threshold(threshold string) xaddtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xaddtrimthreshold{cmds: cmds}
}

type xaddtrimstrategymaxlen struct {
	cmds []string
}

func (c xaddtrimstrategymaxlen) Exact() xaddtrimoperatorexact {
	var cmds []string
	cmds = append(c.cmds, "=")
	return xaddtrimoperatorexact{cmds: cmds}
}

func (c xaddtrimstrategymaxlen) Almost() xaddtrimoperatoralmost {
	var cmds []string
	cmds = append(c.cmds, "~")
	return xaddtrimoperatoralmost{cmds: cmds}
}

func (c xaddtrimstrategymaxlen) Threshold(threshold string) xaddtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xaddtrimthreshold{cmds: cmds}
}

type xaddtrimstrategyminid struct {
	cmds []string
}

func (c xaddtrimstrategyminid) Exact() xaddtrimoperatorexact {
	var cmds []string
	cmds = append(c.cmds, "=")
	return xaddtrimoperatorexact{cmds: cmds}
}

func (c xaddtrimstrategyminid) Almost() xaddtrimoperatoralmost {
	var cmds []string
	cmds = append(c.cmds, "~")
	return xaddtrimoperatoralmost{cmds: cmds}
}

func (c xaddtrimstrategyminid) Threshold(threshold string) xaddtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xaddtrimthreshold{cmds: cmds}
}

type xaddtrimthreshold struct {
	cmds []string
}

func (c xaddtrimthreshold) Limit(count int64) xaddtrimlimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(count, 10))
	return xaddtrimlimit{cmds: cmds}
}

func (c xaddtrimthreshold) Wildcard() xaddwildcard {
	var cmds []string
	cmds = append(c.cmds, "*")
	return xaddwildcard{cmds: cmds}
}

func (c xaddtrimthreshold) Id() xaddid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return xaddid{cmds: cmds}
}

type xaddwildcard struct {
	cmds []string
}

func (c xaddwildcard) FieldValue(field string, value string) xaddfieldvalue {
	var cmds []string
	cmds = append(c.cmds, field, value)
	return xaddfieldvalue{cmds: cmds}
}

type xautoclaim struct {
	cmds []string
}

func (c xautoclaim) Key(key string) xautoclaimkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xautoclaimkey{cmds: cmds}
}

func Xautoclaim() (c xautoclaim) {
	c.cmds = append(c.cmds, "XAUTOCLAIM")
	return
}

type xautoclaimconsumer struct {
	cmds []string
}

func (c xautoclaimconsumer) Minidletime(minidletime string) xautoclaimminidletime {
	var cmds []string
	cmds = append(c.cmds, minidletime)
	return xautoclaimminidletime{cmds: cmds}
}

type xautoclaimcount struct {
	cmds []string
}

func (c xautoclaimcount) Justid() xautoclaimjustidjustid {
	var cmds []string
	cmds = append(c.cmds, "JUSTID")
	return xautoclaimjustidjustid{cmds: cmds}
}

func (c xautoclaimcount) Build() []string {
	return c.cmds
}

type xautoclaimgroup struct {
	cmds []string
}

func (c xautoclaimgroup) Consumer(consumer string) xautoclaimconsumer {
	var cmds []string
	cmds = append(c.cmds, consumer)
	return xautoclaimconsumer{cmds: cmds}
}

type xautoclaimjustidjustid struct {
	cmds []string
}

func (c xautoclaimjustidjustid) Build() []string {
	return c.cmds
}

type xautoclaimkey struct {
	cmds []string
}

func (c xautoclaimkey) Group(group string) xautoclaimgroup {
	var cmds []string
	cmds = append(c.cmds, group)
	return xautoclaimgroup{cmds: cmds}
}

type xautoclaimminidletime struct {
	cmds []string
}

func (c xautoclaimminidletime) Start(start string) xautoclaimstart {
	var cmds []string
	cmds = append(c.cmds, start)
	return xautoclaimstart{cmds: cmds}
}

type xautoclaimstart struct {
	cmds []string
}

func (c xautoclaimstart) Count(count int64) xautoclaimcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return xautoclaimcount{cmds: cmds}
}

func (c xautoclaimstart) Justid() xautoclaimjustidjustid {
	var cmds []string
	cmds = append(c.cmds, "JUSTID")
	return xautoclaimjustidjustid{cmds: cmds}
}

func (c xautoclaimstart) Build() []string {
	return c.cmds
}

type xclaim struct {
	cmds []string
}

func (c xclaim) Key(key string) xclaimkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xclaimkey{cmds: cmds}
}

func Xclaim() (c xclaim) {
	c.cmds = append(c.cmds, "XCLAIM")
	return
}

type xclaimconsumer struct {
	cmds []string
}

func (c xclaimconsumer) Minidletime(minidletime string) xclaimminidletime {
	var cmds []string
	cmds = append(c.cmds, minidletime)
	return xclaimminidletime{cmds: cmds}
}

type xclaimforceforce struct {
	cmds []string
}

func (c xclaimforceforce) Justid() xclaimjustidjustid {
	var cmds []string
	cmds = append(c.cmds, "JUSTID")
	return xclaimjustidjustid{cmds: cmds}
}

func (c xclaimforceforce) Build() []string {
	return c.cmds
}

type xclaimgroup struct {
	cmds []string
}

func (c xclaimgroup) Consumer(consumer string) xclaimconsumer {
	var cmds []string
	cmds = append(c.cmds, consumer)
	return xclaimconsumer{cmds: cmds}
}

type xclaimid struct {
	cmds []string
}

func (c xclaimid) Idle(ms int64) xclaimidle {
	var cmds []string
	cmds = append(c.cmds, "IDLE", strconv.FormatInt(ms, 10))
	return xclaimidle{cmds: cmds}
}

func (c xclaimid) Time(msunixtime int64) xclaimtime {
	var cmds []string
	cmds = append(c.cmds, "TIME", strconv.FormatInt(msunixtime, 10))
	return xclaimtime{cmds: cmds}
}

func (c xclaimid) Retrycount(count int64) xclaimretrycount {
	var cmds []string
	cmds = append(c.cmds, "RETRYCOUNT", strconv.FormatInt(count, 10))
	return xclaimretrycount{cmds: cmds}
}

func (c xclaimid) Force() xclaimforceforce {
	var cmds []string
	cmds = append(c.cmds, "FORCE")
	return xclaimforceforce{cmds: cmds}
}

func (c xclaimid) Justid() xclaimjustidjustid {
	var cmds []string
	cmds = append(c.cmds, "JUSTID")
	return xclaimjustidjustid{cmds: cmds}
}

func (c xclaimid) Id(id ...string) xclaimid {
	var cmds []string
	cmds = append(cmds, id...)
	return xclaimid{cmds: cmds}
}

type xclaimidle struct {
	cmds []string
}

func (c xclaimidle) Time(msunixtime int64) xclaimtime {
	var cmds []string
	cmds = append(c.cmds, "TIME", strconv.FormatInt(msunixtime, 10))
	return xclaimtime{cmds: cmds}
}

func (c xclaimidle) Retrycount(count int64) xclaimretrycount {
	var cmds []string
	cmds = append(c.cmds, "RETRYCOUNT", strconv.FormatInt(count, 10))
	return xclaimretrycount{cmds: cmds}
}

func (c xclaimidle) Force() xclaimforceforce {
	var cmds []string
	cmds = append(c.cmds, "FORCE")
	return xclaimforceforce{cmds: cmds}
}

func (c xclaimidle) Justid() xclaimjustidjustid {
	var cmds []string
	cmds = append(c.cmds, "JUSTID")
	return xclaimjustidjustid{cmds: cmds}
}

func (c xclaimidle) Build() []string {
	return c.cmds
}

type xclaimjustidjustid struct {
	cmds []string
}

func (c xclaimjustidjustid) Build() []string {
	return c.cmds
}

type xclaimkey struct {
	cmds []string
}

func (c xclaimkey) Group(group string) xclaimgroup {
	var cmds []string
	cmds = append(c.cmds, group)
	return xclaimgroup{cmds: cmds}
}

type xclaimminidletime struct {
	cmds []string
}

func (c xclaimminidletime) Id(id ...string) xclaimid {
	var cmds []string
	cmds = append(cmds, id...)
	return xclaimid{cmds: cmds}
}

type xclaimretrycount struct {
	cmds []string
}

func (c xclaimretrycount) Force() xclaimforceforce {
	var cmds []string
	cmds = append(c.cmds, "FORCE")
	return xclaimforceforce{cmds: cmds}
}

func (c xclaimretrycount) Justid() xclaimjustidjustid {
	var cmds []string
	cmds = append(c.cmds, "JUSTID")
	return xclaimjustidjustid{cmds: cmds}
}

func (c xclaimretrycount) Build() []string {
	return c.cmds
}

type xclaimtime struct {
	cmds []string
}

func (c xclaimtime) Retrycount(count int64) xclaimretrycount {
	var cmds []string
	cmds = append(c.cmds, "RETRYCOUNT", strconv.FormatInt(count, 10))
	return xclaimretrycount{cmds: cmds}
}

func (c xclaimtime) Force() xclaimforceforce {
	var cmds []string
	cmds = append(c.cmds, "FORCE")
	return xclaimforceforce{cmds: cmds}
}

func (c xclaimtime) Justid() xclaimjustidjustid {
	var cmds []string
	cmds = append(c.cmds, "JUSTID")
	return xclaimjustidjustid{cmds: cmds}
}

func (c xclaimtime) Build() []string {
	return c.cmds
}

type xdel struct {
	cmds []string
}

func (c xdel) Key(key string) xdelkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xdelkey{cmds: cmds}
}

func Xdel() (c xdel) {
	c.cmds = append(c.cmds, "XDEL")
	return
}

type xdelid struct {
	cmds []string
}

func (c xdelid) Id(id ...string) xdelid {
	var cmds []string
	cmds = append(cmds, id...)
	return xdelid{cmds: cmds}
}

type xdelkey struct {
	cmds []string
}

func (c xdelkey) Id(id ...string) xdelid {
	var cmds []string
	cmds = append(cmds, id...)
	return xdelid{cmds: cmds}
}

type xgroup struct {
	cmds []string
}

func (c xgroup) Create(key string, groupname string) xgroupcreatecreate {
	var cmds []string
	cmds = append(c.cmds, "CREATE", key, groupname)
	return xgroupcreatecreate{cmds: cmds}
}

func (c xgroup) Setid(key string, groupname string) xgroupsetidsetid {
	var cmds []string
	cmds = append(c.cmds, "SETID", key, groupname)
	return xgroupsetidsetid{cmds: cmds}
}

func (c xgroup) Destroy(key string, groupname string) xgroupdestroy {
	var cmds []string
	cmds = append(c.cmds, "DESTROY", key, groupname)
	return xgroupdestroy{cmds: cmds}
}

func (c xgroup) Createconsumer(key string, groupname string, consumername string) xgroupcreateconsumer {
	var cmds []string
	cmds = append(c.cmds, "CREATECONSUMER", key, groupname, consumername)
	return xgroupcreateconsumer{cmds: cmds}
}

func (c xgroup) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

func Xgroup() (c xgroup) {
	c.cmds = append(c.cmds, "XGROUP")
	return
}

type xgroupcreateconsumer struct {
	cmds []string
}

func (c xgroupcreateconsumer) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

func (c xgroupcreateconsumer) Build() []string {
	return c.cmds
}

type xgroupcreatecreate struct {
	cmds []string
}

func (c xgroupcreatecreate) Id() xgroupcreateidid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return xgroupcreateidid{cmds: cmds}
}

func (c xgroupcreatecreate) Lastid() xgroupcreateidlastid {
	var cmds []string
	cmds = append(c.cmds, "$")
	return xgroupcreateidlastid{cmds: cmds}
}

type xgroupcreateidid struct {
	cmds []string
}

func (c xgroupcreateidid) Mkstream() xgroupcreatemkstream {
	var cmds []string
	cmds = append(c.cmds, "MKSTREAM")
	return xgroupcreatemkstream{cmds: cmds}
}

func (c xgroupcreateidid) Setid(key string, groupname string) xgroupsetidsetid {
	var cmds []string
	cmds = append(c.cmds, "SETID", key, groupname)
	return xgroupsetidsetid{cmds: cmds}
}

func (c xgroupcreateidid) Destroy(key string, groupname string) xgroupdestroy {
	var cmds []string
	cmds = append(c.cmds, "DESTROY", key, groupname)
	return xgroupdestroy{cmds: cmds}
}

func (c xgroupcreateidid) Createconsumer(key string, groupname string, consumername string) xgroupcreateconsumer {
	var cmds []string
	cmds = append(c.cmds, "CREATECONSUMER", key, groupname, consumername)
	return xgroupcreateconsumer{cmds: cmds}
}

func (c xgroupcreateidid) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

type xgroupcreateidlastid struct {
	cmds []string
}

func (c xgroupcreateidlastid) Mkstream() xgroupcreatemkstream {
	var cmds []string
	cmds = append(c.cmds, "MKSTREAM")
	return xgroupcreatemkstream{cmds: cmds}
}

func (c xgroupcreateidlastid) Setid(key string, groupname string) xgroupsetidsetid {
	var cmds []string
	cmds = append(c.cmds, "SETID", key, groupname)
	return xgroupsetidsetid{cmds: cmds}
}

func (c xgroupcreateidlastid) Destroy(key string, groupname string) xgroupdestroy {
	var cmds []string
	cmds = append(c.cmds, "DESTROY", key, groupname)
	return xgroupdestroy{cmds: cmds}
}

func (c xgroupcreateidlastid) Createconsumer(key string, groupname string, consumername string) xgroupcreateconsumer {
	var cmds []string
	cmds = append(c.cmds, "CREATECONSUMER", key, groupname, consumername)
	return xgroupcreateconsumer{cmds: cmds}
}

func (c xgroupcreateidlastid) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

type xgroupcreatemkstream struct {
	cmds []string
}

func (c xgroupcreatemkstream) Setid(key string, groupname string) xgroupsetidsetid {
	var cmds []string
	cmds = append(c.cmds, "SETID", key, groupname)
	return xgroupsetidsetid{cmds: cmds}
}

func (c xgroupcreatemkstream) Destroy(key string, groupname string) xgroupdestroy {
	var cmds []string
	cmds = append(c.cmds, "DESTROY", key, groupname)
	return xgroupdestroy{cmds: cmds}
}

func (c xgroupcreatemkstream) Createconsumer(key string, groupname string, consumername string) xgroupcreateconsumer {
	var cmds []string
	cmds = append(c.cmds, "CREATECONSUMER", key, groupname, consumername)
	return xgroupcreateconsumer{cmds: cmds}
}

func (c xgroupcreatemkstream) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

type xgroupdelconsumer struct {
	cmds []string
}

func (c xgroupdelconsumer) Build() []string {
	return c.cmds
}

type xgroupdestroy struct {
	cmds []string
}

func (c xgroupdestroy) Createconsumer(key string, groupname string, consumername string) xgroupcreateconsumer {
	var cmds []string
	cmds = append(c.cmds, "CREATECONSUMER", key, groupname, consumername)
	return xgroupcreateconsumer{cmds: cmds}
}

func (c xgroupdestroy) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

func (c xgroupdestroy) Build() []string {
	return c.cmds
}

type xgroupsetididid struct {
	cmds []string
}

func (c xgroupsetididid) Destroy(key string, groupname string) xgroupdestroy {
	var cmds []string
	cmds = append(c.cmds, "DESTROY", key, groupname)
	return xgroupdestroy{cmds: cmds}
}

func (c xgroupsetididid) Createconsumer(key string, groupname string, consumername string) xgroupcreateconsumer {
	var cmds []string
	cmds = append(c.cmds, "CREATECONSUMER", key, groupname, consumername)
	return xgroupcreateconsumer{cmds: cmds}
}

func (c xgroupsetididid) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

func (c xgroupsetididid) Build() []string {
	return c.cmds
}

type xgroupsetididlastid struct {
	cmds []string
}

func (c xgroupsetididlastid) Destroy(key string, groupname string) xgroupdestroy {
	var cmds []string
	cmds = append(c.cmds, "DESTROY", key, groupname)
	return xgroupdestroy{cmds: cmds}
}

func (c xgroupsetididlastid) Createconsumer(key string, groupname string, consumername string) xgroupcreateconsumer {
	var cmds []string
	cmds = append(c.cmds, "CREATECONSUMER", key, groupname, consumername)
	return xgroupcreateconsumer{cmds: cmds}
}

func (c xgroupsetididlastid) Delconsumer(key string, groupname string, consumername string) xgroupdelconsumer {
	var cmds []string
	cmds = append(c.cmds, "DELCONSUMER", key, groupname, consumername)
	return xgroupdelconsumer{cmds: cmds}
}

func (c xgroupsetididlastid) Build() []string {
	return c.cmds
}

type xgroupsetidsetid struct {
	cmds []string
}

func (c xgroupsetidsetid) Id() xgroupsetididid {
	var cmds []string
	cmds = append(c.cmds, "ID")
	return xgroupsetididid{cmds: cmds}
}

func (c xgroupsetidsetid) Lastid() xgroupsetididlastid {
	var cmds []string
	cmds = append(c.cmds, "$")
	return xgroupsetididlastid{cmds: cmds}
}

type xinfo struct {
	cmds []string
}

func (c xinfo) Consumers(key string, groupname string) xinfoconsumers {
	var cmds []string
	cmds = append(c.cmds, "CONSUMERS", key, groupname)
	return xinfoconsumers{cmds: cmds}
}

func (c xinfo) Groups(key string) xinfogroups {
	var cmds []string
	cmds = append(c.cmds, "GROUPS", key)
	return xinfogroups{cmds: cmds}
}

func (c xinfo) Stream(key string) xinfostream {
	var cmds []string
	cmds = append(c.cmds, "STREAM", key)
	return xinfostream{cmds: cmds}
}

func (c xinfo) Help() xinfohelphelp {
	var cmds []string
	cmds = append(c.cmds, "HELP")
	return xinfohelphelp{cmds: cmds}
}

func (c xinfo) Build() []string {
	return c.cmds
}

func Xinfo() (c xinfo) {
	c.cmds = append(c.cmds, "XINFO")
	return
}

type xinfoconsumers struct {
	cmds []string
}

func (c xinfoconsumers) Groups(key string) xinfogroups {
	var cmds []string
	cmds = append(c.cmds, "GROUPS", key)
	return xinfogroups{cmds: cmds}
}

func (c xinfoconsumers) Stream(key string) xinfostream {
	var cmds []string
	cmds = append(c.cmds, "STREAM", key)
	return xinfostream{cmds: cmds}
}

func (c xinfoconsumers) Help() xinfohelphelp {
	var cmds []string
	cmds = append(c.cmds, "HELP")
	return xinfohelphelp{cmds: cmds}
}

func (c xinfoconsumers) Build() []string {
	return c.cmds
}

type xinfogroups struct {
	cmds []string
}

func (c xinfogroups) Stream(key string) xinfostream {
	var cmds []string
	cmds = append(c.cmds, "STREAM", key)
	return xinfostream{cmds: cmds}
}

func (c xinfogroups) Help() xinfohelphelp {
	var cmds []string
	cmds = append(c.cmds, "HELP")
	return xinfohelphelp{cmds: cmds}
}

func (c xinfogroups) Build() []string {
	return c.cmds
}

type xinfohelphelp struct {
	cmds []string
}

func (c xinfohelphelp) Build() []string {
	return c.cmds
}

type xinfostream struct {
	cmds []string
}

func (c xinfostream) Help() xinfohelphelp {
	var cmds []string
	cmds = append(c.cmds, "HELP")
	return xinfohelphelp{cmds: cmds}
}

func (c xinfostream) Build() []string {
	return c.cmds
}

type xlen struct {
	cmds []string
}

func (c xlen) Key(key string) xlenkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xlenkey{cmds: cmds}
}

func Xlen() (c xlen) {
	c.cmds = append(c.cmds, "XLEN")
	return
}

type xlenkey struct {
	cmds []string
}

func (c xlenkey) Build() []string {
	return c.cmds
}

type xpending struct {
	cmds []string
}

func (c xpending) Key(key string) xpendingkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xpendingkey{cmds: cmds}
}

func Xpending() (c xpending) {
	c.cmds = append(c.cmds, "XPENDING")
	return
}

type xpendingfiltersconsumer struct {
	cmds []string
}

func (c xpendingfiltersconsumer) Build() []string {
	return c.cmds
}

type xpendingfilterscount struct {
	cmds []string
}

func (c xpendingfilterscount) Consumer(consumer string) xpendingfiltersconsumer {
	var cmds []string
	cmds = append(c.cmds, consumer)
	return xpendingfiltersconsumer{cmds: cmds}
}

func (c xpendingfilterscount) Build() []string {
	return c.cmds
}

type xpendingfiltersend struct {
	cmds []string
}

func (c xpendingfiltersend) Count(count int64) xpendingfilterscount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return xpendingfilterscount{cmds: cmds}
}

type xpendingfiltersidle struct {
	cmds []string
}

func (c xpendingfiltersidle) Start(start string) xpendingfiltersstart {
	var cmds []string
	cmds = append(c.cmds, start)
	return xpendingfiltersstart{cmds: cmds}
}

type xpendingfiltersstart struct {
	cmds []string
}

func (c xpendingfiltersstart) End(end string) xpendingfiltersend {
	var cmds []string
	cmds = append(c.cmds, end)
	return xpendingfiltersend{cmds: cmds}
}

type xpendinggroup struct {
	cmds []string
}

func (c xpendinggroup) Idle(minidletime int64) xpendingfiltersidle {
	var cmds []string
	cmds = append(c.cmds, "IDLE", strconv.FormatInt(minidletime, 10))
	return xpendingfiltersidle{cmds: cmds}
}

func (c xpendinggroup) Start(start string) xpendingfiltersstart {
	var cmds []string
	cmds = append(c.cmds, start)
	return xpendingfiltersstart{cmds: cmds}
}

type xpendingkey struct {
	cmds []string
}

func (c xpendingkey) Group(group string) xpendinggroup {
	var cmds []string
	cmds = append(c.cmds, group)
	return xpendinggroup{cmds: cmds}
}

type xrange struct {
	cmds []string
}

func (c xrange) Key(key string) xrangekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xrangekey{cmds: cmds}
}

func Xrange() (c xrange) {
	c.cmds = append(c.cmds, "XRANGE")
	return
}

type xrangecount struct {
	cmds []string
}

func (c xrangecount) Build() []string {
	return c.cmds
}

type xrangeend struct {
	cmds []string
}

func (c xrangeend) Count(count int64) xrangecount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return xrangecount{cmds: cmds}
}

func (c xrangeend) Build() []string {
	return c.cmds
}

type xrangekey struct {
	cmds []string
}

func (c xrangekey) Start(start string) xrangestart {
	var cmds []string
	cmds = append(c.cmds, start)
	return xrangestart{cmds: cmds}
}

type xrangestart struct {
	cmds []string
}

func (c xrangestart) End(end string) xrangeend {
	var cmds []string
	cmds = append(c.cmds, end)
	return xrangeend{cmds: cmds}
}

type xread struct {
	cmds []string
}

func (c xread) Count(count int64) xreadcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return xreadcount{cmds: cmds}
}

func (c xread) Block(milliseconds int64) xreadblock {
	var cmds []string
	cmds = append(c.cmds, "BLOCK", strconv.FormatInt(milliseconds, 10))
	return xreadblock{cmds: cmds}
}

func (c xread) Streams() xreadstreamsstreams {
	var cmds []string
	cmds = append(c.cmds, "STREAMS")
	return xreadstreamsstreams{cmds: cmds}
}

func Xread() (c xread) {
	c.cmds = append(c.cmds, "XREAD")
	return
}

type xreadblock struct {
	cmds []string
}

func (c xreadblock) Streams() xreadstreamsstreams {
	var cmds []string
	cmds = append(c.cmds, "STREAMS")
	return xreadstreamsstreams{cmds: cmds}
}

type xreadcount struct {
	cmds []string
}

func (c xreadcount) Block(milliseconds int64) xreadblock {
	var cmds []string
	cmds = append(c.cmds, "BLOCK", strconv.FormatInt(milliseconds, 10))
	return xreadblock{cmds: cmds}
}

func (c xreadcount) Streams() xreadstreamsstreams {
	var cmds []string
	cmds = append(c.cmds, "STREAMS")
	return xreadstreamsstreams{cmds: cmds}
}

type xreadgroup struct {
	cmds []string
}

func (c xreadgroup) Group(group string, consumer string) xreadgroupgroup {
	var cmds []string
	cmds = append(c.cmds, "GROUP", group, consumer)
	return xreadgroupgroup{cmds: cmds}
}

func Xreadgroup() (c xreadgroup) {
	c.cmds = append(c.cmds, "XREADGROUP")
	return
}

type xreadgroupblock struct {
	cmds []string
}

func (c xreadgroupblock) Noack() xreadgroupnoacknoack {
	var cmds []string
	cmds = append(c.cmds, "NOACK")
	return xreadgroupnoacknoack{cmds: cmds}
}

func (c xreadgroupblock) Streams() xreadgroupstreamsstreams {
	var cmds []string
	cmds = append(c.cmds, "STREAMS")
	return xreadgroupstreamsstreams{cmds: cmds}
}

type xreadgroupcount struct {
	cmds []string
}

func (c xreadgroupcount) Block(milliseconds int64) xreadgroupblock {
	var cmds []string
	cmds = append(c.cmds, "BLOCK", strconv.FormatInt(milliseconds, 10))
	return xreadgroupblock{cmds: cmds}
}

func (c xreadgroupcount) Noack() xreadgroupnoacknoack {
	var cmds []string
	cmds = append(c.cmds, "NOACK")
	return xreadgroupnoacknoack{cmds: cmds}
}

func (c xreadgroupcount) Streams() xreadgroupstreamsstreams {
	var cmds []string
	cmds = append(c.cmds, "STREAMS")
	return xreadgroupstreamsstreams{cmds: cmds}
}

type xreadgroupgroup struct {
	cmds []string
}

func (c xreadgroupgroup) Count(count int64) xreadgroupcount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return xreadgroupcount{cmds: cmds}
}

func (c xreadgroupgroup) Block(milliseconds int64) xreadgroupblock {
	var cmds []string
	cmds = append(c.cmds, "BLOCK", strconv.FormatInt(milliseconds, 10))
	return xreadgroupblock{cmds: cmds}
}

func (c xreadgroupgroup) Noack() xreadgroupnoacknoack {
	var cmds []string
	cmds = append(c.cmds, "NOACK")
	return xreadgroupnoacknoack{cmds: cmds}
}

func (c xreadgroupgroup) Streams() xreadgroupstreamsstreams {
	var cmds []string
	cmds = append(c.cmds, "STREAMS")
	return xreadgroupstreamsstreams{cmds: cmds}
}

type xreadgroupid struct {
	cmds []string
}

func (c xreadgroupid) Id(id ...string) xreadgroupid {
	var cmds []string
	cmds = append(cmds, id...)
	return xreadgroupid{cmds: cmds}
}

type xreadgroupkey struct {
	cmds []string
}

func (c xreadgroupkey) Id(id ...string) xreadgroupid {
	var cmds []string
	cmds = append(cmds, id...)
	return xreadgroupid{cmds: cmds}
}

func (c xreadgroupkey) Key(key ...string) xreadgroupkey {
	var cmds []string
	cmds = append(cmds, key...)
	return xreadgroupkey{cmds: cmds}
}

type xreadgroupnoacknoack struct {
	cmds []string
}

func (c xreadgroupnoacknoack) Streams() xreadgroupstreamsstreams {
	var cmds []string
	cmds = append(c.cmds, "STREAMS")
	return xreadgroupstreamsstreams{cmds: cmds}
}

type xreadgroupstreamsstreams struct {
	cmds []string
}

func (c xreadgroupstreamsstreams) Key(key ...string) xreadgroupkey {
	var cmds []string
	cmds = append(cmds, key...)
	return xreadgroupkey{cmds: cmds}
}

type xreadid struct {
	cmds []string
}

func (c xreadid) Id(id ...string) xreadid {
	var cmds []string
	cmds = append(cmds, id...)
	return xreadid{cmds: cmds}
}

type xreadkey struct {
	cmds []string
}

func (c xreadkey) Id(id ...string) xreadid {
	var cmds []string
	cmds = append(cmds, id...)
	return xreadid{cmds: cmds}
}

func (c xreadkey) Key(key ...string) xreadkey {
	var cmds []string
	cmds = append(cmds, key...)
	return xreadkey{cmds: cmds}
}

type xreadstreamsstreams struct {
	cmds []string
}

func (c xreadstreamsstreams) Key(key ...string) xreadkey {
	var cmds []string
	cmds = append(cmds, key...)
	return xreadkey{cmds: cmds}
}

type xrevrange struct {
	cmds []string
}

func (c xrevrange) Key(key string) xrevrangekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xrevrangekey{cmds: cmds}
}

func Xrevrange() (c xrevrange) {
	c.cmds = append(c.cmds, "XREVRANGE")
	return
}

type xrevrangecount struct {
	cmds []string
}

func (c xrevrangecount) Build() []string {
	return c.cmds
}

type xrevrangeend struct {
	cmds []string
}

func (c xrevrangeend) Start(start string) xrevrangestart {
	var cmds []string
	cmds = append(c.cmds, start)
	return xrevrangestart{cmds: cmds}
}

type xrevrangekey struct {
	cmds []string
}

func (c xrevrangekey) End(end string) xrevrangeend {
	var cmds []string
	cmds = append(c.cmds, end)
	return xrevrangeend{cmds: cmds}
}

type xrevrangestart struct {
	cmds []string
}

func (c xrevrangestart) Count(count int64) xrevrangecount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return xrevrangecount{cmds: cmds}
}

func (c xrevrangestart) Build() []string {
	return c.cmds
}

type xtrim struct {
	cmds []string
}

func (c xtrim) Key(key string) xtrimkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return xtrimkey{cmds: cmds}
}

func Xtrim() (c xtrim) {
	c.cmds = append(c.cmds, "XTRIM")
	return
}

type xtrimkey struct {
	cmds []string
}

func (c xtrimkey) Maxlen() xtrimtrimstrategymaxlen {
	var cmds []string
	cmds = append(c.cmds, "MAXLEN")
	return xtrimtrimstrategymaxlen{cmds: cmds}
}

func (c xtrimkey) Minid() xtrimtrimstrategyminid {
	var cmds []string
	cmds = append(c.cmds, "MINID")
	return xtrimtrimstrategyminid{cmds: cmds}
}

type xtrimtrimlimit struct {
	cmds []string
}

func (c xtrimtrimlimit) Build() []string {
	return c.cmds
}

type xtrimtrimoperatoralmost struct {
	cmds []string
}

func (c xtrimtrimoperatoralmost) Threshold(threshold string) xtrimtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xtrimtrimthreshold{cmds: cmds}
}

type xtrimtrimoperatorexact struct {
	cmds []string
}

func (c xtrimtrimoperatorexact) Threshold(threshold string) xtrimtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xtrimtrimthreshold{cmds: cmds}
}

type xtrimtrimstrategymaxlen struct {
	cmds []string
}

func (c xtrimtrimstrategymaxlen) Exact() xtrimtrimoperatorexact {
	var cmds []string
	cmds = append(c.cmds, "=")
	return xtrimtrimoperatorexact{cmds: cmds}
}

func (c xtrimtrimstrategymaxlen) Almost() xtrimtrimoperatoralmost {
	var cmds []string
	cmds = append(c.cmds, "~")
	return xtrimtrimoperatoralmost{cmds: cmds}
}

func (c xtrimtrimstrategymaxlen) Threshold(threshold string) xtrimtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xtrimtrimthreshold{cmds: cmds}
}

type xtrimtrimstrategyminid struct {
	cmds []string
}

func (c xtrimtrimstrategyminid) Exact() xtrimtrimoperatorexact {
	var cmds []string
	cmds = append(c.cmds, "=")
	return xtrimtrimoperatorexact{cmds: cmds}
}

func (c xtrimtrimstrategyminid) Almost() xtrimtrimoperatoralmost {
	var cmds []string
	cmds = append(c.cmds, "~")
	return xtrimtrimoperatoralmost{cmds: cmds}
}

func (c xtrimtrimstrategyminid) Threshold(threshold string) xtrimtrimthreshold {
	var cmds []string
	cmds = append(c.cmds, threshold)
	return xtrimtrimthreshold{cmds: cmds}
}

type xtrimtrimthreshold struct {
	cmds []string
}

func (c xtrimtrimthreshold) Limit(count int64) xtrimtrimlimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(count, 10))
	return xtrimtrimlimit{cmds: cmds}
}

func (c xtrimtrimthreshold) Build() []string {
	return c.cmds
}

type zadd struct {
	cmds []string
}

func (c zadd) Key(key string) zaddkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zaddkey{cmds: cmds}
}

func Zadd() (c zadd) {
	c.cmds = append(c.cmds, "ZADD")
	return
}

type zaddchangech struct {
	cmds []string
}

func (c zaddchangech) Incr() zaddincrementincr {
	var cmds []string
	cmds = append(c.cmds, "INCR")
	return zaddincrementincr{cmds: cmds}
}

func (c zaddchangech) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zaddcomparisongt struct {
	cmds []string
}

func (c zaddcomparisongt) Ch() zaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return zaddchangech{cmds: cmds}
}

func (c zaddcomparisongt) Incr() zaddincrementincr {
	var cmds []string
	cmds = append(c.cmds, "INCR")
	return zaddincrementincr{cmds: cmds}
}

func (c zaddcomparisongt) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zaddcomparisonlt struct {
	cmds []string
}

func (c zaddcomparisonlt) Ch() zaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return zaddchangech{cmds: cmds}
}

func (c zaddcomparisonlt) Incr() zaddincrementincr {
	var cmds []string
	cmds = append(c.cmds, "INCR")
	return zaddincrementincr{cmds: cmds}
}

func (c zaddcomparisonlt) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zaddconditionnx struct {
	cmds []string
}

func (c zaddconditionnx) Gt() zaddcomparisongt {
	var cmds []string
	cmds = append(c.cmds, "GT")
	return zaddcomparisongt{cmds: cmds}
}

func (c zaddconditionnx) Lt() zaddcomparisonlt {
	var cmds []string
	cmds = append(c.cmds, "LT")
	return zaddcomparisonlt{cmds: cmds}
}

func (c zaddconditionnx) Ch() zaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return zaddchangech{cmds: cmds}
}

func (c zaddconditionnx) Incr() zaddincrementincr {
	var cmds []string
	cmds = append(c.cmds, "INCR")
	return zaddincrementincr{cmds: cmds}
}

func (c zaddconditionnx) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zaddconditionxx struct {
	cmds []string
}

func (c zaddconditionxx) Gt() zaddcomparisongt {
	var cmds []string
	cmds = append(c.cmds, "GT")
	return zaddcomparisongt{cmds: cmds}
}

func (c zaddconditionxx) Lt() zaddcomparisonlt {
	var cmds []string
	cmds = append(c.cmds, "LT")
	return zaddcomparisonlt{cmds: cmds}
}

func (c zaddconditionxx) Ch() zaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return zaddchangech{cmds: cmds}
}

func (c zaddconditionxx) Incr() zaddincrementincr {
	var cmds []string
	cmds = append(c.cmds, "INCR")
	return zaddincrementincr{cmds: cmds}
}

func (c zaddconditionxx) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zaddincrementincr struct {
	cmds []string
}

func (c zaddincrementincr) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zaddkey struct {
	cmds []string
}

func (c zaddkey) Nx() zaddconditionnx {
	var cmds []string
	cmds = append(c.cmds, "NX")
	return zaddconditionnx{cmds: cmds}
}

func (c zaddkey) Xx() zaddconditionxx {
	var cmds []string
	cmds = append(c.cmds, "XX")
	return zaddconditionxx{cmds: cmds}
}

func (c zaddkey) Gt() zaddcomparisongt {
	var cmds []string
	cmds = append(c.cmds, "GT")
	return zaddcomparisongt{cmds: cmds}
}

func (c zaddkey) Lt() zaddcomparisonlt {
	var cmds []string
	cmds = append(c.cmds, "LT")
	return zaddcomparisonlt{cmds: cmds}
}

func (c zaddkey) Ch() zaddchangech {
	var cmds []string
	cmds = append(c.cmds, "CH")
	return zaddchangech{cmds: cmds}
}

func (c zaddkey) Incr() zaddincrementincr {
	var cmds []string
	cmds = append(c.cmds, "INCR")
	return zaddincrementincr{cmds: cmds}
}

func (c zaddkey) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zaddscoremember struct {
	cmds []string
}

func (c zaddscoremember) ScoreMember(score float64, member string) zaddscoremember {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(score, 'f', -1, 64), member)
	return zaddscoremember{cmds: cmds}
}

type zcard struct {
	cmds []string
}

func (c zcard) Key(key string) zcardkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zcardkey{cmds: cmds}
}

func Zcard() (c zcard) {
	c.cmds = append(c.cmds, "ZCARD")
	return
}

type zcardkey struct {
	cmds []string
}

func (c zcardkey) Build() []string {
	return c.cmds
}

type zcount struct {
	cmds []string
}

func (c zcount) Key(key string) zcountkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zcountkey{cmds: cmds}
}

func Zcount() (c zcount) {
	c.cmds = append(c.cmds, "ZCOUNT")
	return
}

type zcountkey struct {
	cmds []string
}

func (c zcountkey) Min(min float64) zcountmin {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(min, 'f', -1, 64))
	return zcountmin{cmds: cmds}
}

type zcountmax struct {
	cmds []string
}

func (c zcountmax) Build() []string {
	return c.cmds
}

type zcountmin struct {
	cmds []string
}

func (c zcountmin) Max(max float64) zcountmax {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(max, 'f', -1, 64))
	return zcountmax{cmds: cmds}
}

type zdiff struct {
	cmds []string
}

func (c zdiff) Numkeys(numkeys int64) zdiffnumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return zdiffnumkeys{cmds: cmds}
}

func Zdiff() (c zdiff) {
	c.cmds = append(c.cmds, "ZDIFF")
	return
}

type zdiffkey struct {
	cmds []string
}

func (c zdiffkey) Withscores() zdiffwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zdiffwithscoreswithscores{cmds: cmds}
}

func (c zdiffkey) Key(key ...string) zdiffkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zdiffkey{cmds: cmds}
}

type zdiffnumkeys struct {
	cmds []string
}

func (c zdiffnumkeys) Key(key ...string) zdiffkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zdiffkey{cmds: cmds}
}

type zdiffstore struct {
	cmds []string
}

func (c zdiffstore) Destination(destination string) zdiffstoredestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return zdiffstoredestination{cmds: cmds}
}

func Zdiffstore() (c zdiffstore) {
	c.cmds = append(c.cmds, "ZDIFFSTORE")
	return
}

type zdiffstoredestination struct {
	cmds []string
}

func (c zdiffstoredestination) Numkeys(numkeys int64) zdiffstorenumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return zdiffstorenumkeys{cmds: cmds}
}

type zdiffstorekey struct {
	cmds []string
}

func (c zdiffstorekey) Key(key ...string) zdiffstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return zdiffstorekey{cmds: cmds}
}

type zdiffstorenumkeys struct {
	cmds []string
}

func (c zdiffstorenumkeys) Key(key ...string) zdiffstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return zdiffstorekey{cmds: cmds}
}

type zdiffwithscoreswithscores struct {
	cmds []string
}

func (c zdiffwithscoreswithscores) Build() []string {
	return c.cmds
}

type zincrby struct {
	cmds []string
}

func (c zincrby) Key(key string) zincrbykey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zincrbykey{cmds: cmds}
}

func Zincrby() (c zincrby) {
	c.cmds = append(c.cmds, "ZINCRBY")
	return
}

type zincrbyincrement struct {
	cmds []string
}

func (c zincrbyincrement) Member(member string) zincrbymember {
	var cmds []string
	cmds = append(c.cmds, member)
	return zincrbymember{cmds: cmds}
}

type zincrbykey struct {
	cmds []string
}

func (c zincrbykey) Increment(increment int64) zincrbyincrement {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(increment, 10))
	return zincrbyincrement{cmds: cmds}
}

type zincrbymember struct {
	cmds []string
}

func (c zincrbymember) Build() []string {
	return c.cmds
}

type zinter struct {
	cmds []string
}

func (c zinter) Numkeys(numkeys int64) zinternumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return zinternumkeys{cmds: cmds}
}

func Zinter() (c zinter) {
	c.cmds = append(c.cmds, "ZINTER")
	return
}

type zinteraggregatemax struct {
	cmds []string
}

func (c zinteraggregatemax) Withscores() zinterwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zinterwithscoreswithscores{cmds: cmds}
}

func (c zinteraggregatemax) Build() []string {
	return c.cmds
}

type zinteraggregatemin struct {
	cmds []string
}

func (c zinteraggregatemin) Withscores() zinterwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zinterwithscoreswithscores{cmds: cmds}
}

func (c zinteraggregatemin) Build() []string {
	return c.cmds
}

type zinteraggregatesum struct {
	cmds []string
}

func (c zinteraggregatesum) Withscores() zinterwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zinterwithscoreswithscores{cmds: cmds}
}

func (c zinteraggregatesum) Build() []string {
	return c.cmds
}

type zintercard struct {
	cmds []string
}

func (c zintercard) Numkeys(numkeys int64) zintercardnumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return zintercardnumkeys{cmds: cmds}
}

func Zintercard() (c zintercard) {
	c.cmds = append(c.cmds, "ZINTERCARD")
	return
}

type zintercardkey struct {
	cmds []string
}

func (c zintercardkey) Key(key ...string) zintercardkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zintercardkey{cmds: cmds}
}

type zintercardnumkeys struct {
	cmds []string
}

func (c zintercardnumkeys) Key(key ...string) zintercardkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zintercardkey{cmds: cmds}
}

type zinterkey struct {
	cmds []string
}

func (c zinterkey) Weights(weight int64) zinterweights {
	var cmds []string
	cmds = append(c.cmds, "WEIGHTS", strconv.FormatInt(weight, 10))
	return zinterweights{cmds: cmds}
}

func (c zinterkey) Sum() zinteraggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zinteraggregatesum{cmds: cmds}
}

func (c zinterkey) Min() zinteraggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zinteraggregatemin{cmds: cmds}
}

func (c zinterkey) Max() zinteraggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zinteraggregatemax{cmds: cmds}
}

func (c zinterkey) Withscores() zinterwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zinterwithscoreswithscores{cmds: cmds}
}

func (c zinterkey) Key(key ...string) zinterkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zinterkey{cmds: cmds}
}

type zinternumkeys struct {
	cmds []string
}

func (c zinternumkeys) Key(key ...string) zinterkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zinterkey{cmds: cmds}
}

type zinterstore struct {
	cmds []string
}

func (c zinterstore) Destination(destination string) zinterstoredestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return zinterstoredestination{cmds: cmds}
}

func Zinterstore() (c zinterstore) {
	c.cmds = append(c.cmds, "ZINTERSTORE")
	return
}

type zinterstoreaggregatemax struct {
	cmds []string
}

func (c zinterstoreaggregatemax) Build() []string {
	return c.cmds
}

type zinterstoreaggregatemin struct {
	cmds []string
}

func (c zinterstoreaggregatemin) Build() []string {
	return c.cmds
}

type zinterstoreaggregatesum struct {
	cmds []string
}

func (c zinterstoreaggregatesum) Build() []string {
	return c.cmds
}

type zinterstoredestination struct {
	cmds []string
}

func (c zinterstoredestination) Numkeys(numkeys int64) zinterstorenumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return zinterstorenumkeys{cmds: cmds}
}

type zinterstorekey struct {
	cmds []string
}

func (c zinterstorekey) Weights(weight int64) zinterstoreweights {
	var cmds []string
	cmds = append(c.cmds, "WEIGHTS", strconv.FormatInt(weight, 10))
	return zinterstoreweights{cmds: cmds}
}

func (c zinterstorekey) Sum() zinterstoreaggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zinterstoreaggregatesum{cmds: cmds}
}

func (c zinterstorekey) Min() zinterstoreaggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zinterstoreaggregatemin{cmds: cmds}
}

func (c zinterstorekey) Max() zinterstoreaggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zinterstoreaggregatemax{cmds: cmds}
}

func (c zinterstorekey) Key(key ...string) zinterstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return zinterstorekey{cmds: cmds}
}

type zinterstorenumkeys struct {
	cmds []string
}

func (c zinterstorenumkeys) Key(key ...string) zinterstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return zinterstorekey{cmds: cmds}
}

type zinterstoreweights struct {
	cmds []string
}

func (c zinterstoreweights) Sum() zinterstoreaggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zinterstoreaggregatesum{cmds: cmds}
}

func (c zinterstoreweights) Min() zinterstoreaggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zinterstoreaggregatemin{cmds: cmds}
}

func (c zinterstoreweights) Max() zinterstoreaggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zinterstoreaggregatemax{cmds: cmds}
}

func (c zinterstoreweights) Weight(weight int64) zinterstoreweights {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(weight, 10))
	return zinterstoreweights{cmds: cmds}
}

func (c zinterstoreweights) Build() []string {
	return c.cmds
}

type zinterweights struct {
	cmds []string
}

func (c zinterweights) Sum() zinteraggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zinteraggregatesum{cmds: cmds}
}

func (c zinterweights) Min() zinteraggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zinteraggregatemin{cmds: cmds}
}

func (c zinterweights) Max() zinteraggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zinteraggregatemax{cmds: cmds}
}

func (c zinterweights) Withscores() zinterwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zinterwithscoreswithscores{cmds: cmds}
}

func (c zinterweights) Weight(weight int64) zinterweights {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(weight, 10))
	return zinterweights{cmds: cmds}
}

func (c zinterweights) Build() []string {
	return c.cmds
}

type zinterwithscoreswithscores struct {
	cmds []string
}

func (c zinterwithscoreswithscores) Build() []string {
	return c.cmds
}

type zlexcount struct {
	cmds []string
}

func (c zlexcount) Key(key string) zlexcountkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zlexcountkey{cmds: cmds}
}

func Zlexcount() (c zlexcount) {
	c.cmds = append(c.cmds, "ZLEXCOUNT")
	return
}

type zlexcountkey struct {
	cmds []string
}

func (c zlexcountkey) Min(min string) zlexcountmin {
	var cmds []string
	cmds = append(c.cmds, min)
	return zlexcountmin{cmds: cmds}
}

type zlexcountmax struct {
	cmds []string
}

func (c zlexcountmax) Build() []string {
	return c.cmds
}

type zlexcountmin struct {
	cmds []string
}

func (c zlexcountmin) Max(max string) zlexcountmax {
	var cmds []string
	cmds = append(c.cmds, max)
	return zlexcountmax{cmds: cmds}
}

type zmscore struct {
	cmds []string
}

func (c zmscore) Key(key string) zmscorekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zmscorekey{cmds: cmds}
}

func Zmscore() (c zmscore) {
	c.cmds = append(c.cmds, "ZMSCORE")
	return
}

type zmscorekey struct {
	cmds []string
}

func (c zmscorekey) Member(member ...string) zmscoremember {
	var cmds []string
	cmds = append(cmds, member...)
	return zmscoremember{cmds: cmds}
}

type zmscoremember struct {
	cmds []string
}

func (c zmscoremember) Member(member ...string) zmscoremember {
	var cmds []string
	cmds = append(cmds, member...)
	return zmscoremember{cmds: cmds}
}

type zpopmax struct {
	cmds []string
}

func (c zpopmax) Key(key string) zpopmaxkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zpopmaxkey{cmds: cmds}
}

func Zpopmax() (c zpopmax) {
	c.cmds = append(c.cmds, "ZPOPMAX")
	return
}

type zpopmaxcount struct {
	cmds []string
}

func (c zpopmaxcount) Build() []string {
	return c.cmds
}

type zpopmaxkey struct {
	cmds []string
}

func (c zpopmaxkey) Count(count int64) zpopmaxcount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return zpopmaxcount{cmds: cmds}
}

func (c zpopmaxkey) Build() []string {
	return c.cmds
}

type zpopmin struct {
	cmds []string
}

func (c zpopmin) Key(key string) zpopminkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zpopminkey{cmds: cmds}
}

func Zpopmin() (c zpopmin) {
	c.cmds = append(c.cmds, "ZPOPMIN")
	return
}

type zpopmincount struct {
	cmds []string
}

func (c zpopmincount) Build() []string {
	return c.cmds
}

type zpopminkey struct {
	cmds []string
}

func (c zpopminkey) Count(count int64) zpopmincount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return zpopmincount{cmds: cmds}
}

func (c zpopminkey) Build() []string {
	return c.cmds
}

type zrandmember struct {
	cmds []string
}

func (c zrandmember) Key(key string) zrandmemberkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrandmemberkey{cmds: cmds}
}

func Zrandmember() (c zrandmember) {
	c.cmds = append(c.cmds, "ZRANDMEMBER")
	return
}

type zrandmemberkey struct {
	cmds []string
}

func (c zrandmemberkey) Count(count int64) zrandmemberoptionscount {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(count, 10))
	return zrandmemberoptionscount{cmds: cmds}
}

type zrandmemberoptionscount struct {
	cmds []string
}

func (c zrandmemberoptionscount) Withscores() zrandmemberoptionswithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrandmemberoptionswithscoreswithscores{cmds: cmds}
}

func (c zrandmemberoptionscount) Build() []string {
	return c.cmds
}

type zrandmemberoptionswithscoreswithscores struct {
	cmds []string
}

func (c zrandmemberoptionswithscoreswithscores) Build() []string {
	return c.cmds
}

type zrange struct {
	cmds []string
}

func (c zrange) Key(key string) zrangekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrangekey{cmds: cmds}
}

func Zrange() (c zrange) {
	c.cmds = append(c.cmds, "ZRANGE")
	return
}

type zrangebylex struct {
	cmds []string
}

func (c zrangebylex) Key(key string) zrangebylexkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrangebylexkey{cmds: cmds}
}

func Zrangebylex() (c zrangebylex) {
	c.cmds = append(c.cmds, "ZRANGEBYLEX")
	return
}

type zrangebylexkey struct {
	cmds []string
}

func (c zrangebylexkey) Min(min string) zrangebylexmin {
	var cmds []string
	cmds = append(c.cmds, min)
	return zrangebylexmin{cmds: cmds}
}

type zrangebylexlimit struct {
	cmds []string
}

func (c zrangebylexlimit) Build() []string {
	return c.cmds
}

type zrangebylexmax struct {
	cmds []string
}

func (c zrangebylexmax) Limit(offset int64, count int64) zrangebylexlimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangebylexlimit{cmds: cmds}
}

func (c zrangebylexmax) Build() []string {
	return c.cmds
}

type zrangebylexmin struct {
	cmds []string
}

func (c zrangebylexmin) Max(max string) zrangebylexmax {
	var cmds []string
	cmds = append(c.cmds, max)
	return zrangebylexmax{cmds: cmds}
}

type zrangebyscore struct {
	cmds []string
}

func (c zrangebyscore) Key(key string) zrangebyscorekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrangebyscorekey{cmds: cmds}
}

func Zrangebyscore() (c zrangebyscore) {
	c.cmds = append(c.cmds, "ZRANGEBYSCORE")
	return
}

type zrangebyscorekey struct {
	cmds []string
}

func (c zrangebyscorekey) Min(min float64) zrangebyscoremin {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(min, 'f', -1, 64))
	return zrangebyscoremin{cmds: cmds}
}

type zrangebyscorelimit struct {
	cmds []string
}

func (c zrangebyscorelimit) Build() []string {
	return c.cmds
}

type zrangebyscoremax struct {
	cmds []string
}

func (c zrangebyscoremax) Withscores() zrangebyscorewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrangebyscorewithscoreswithscores{cmds: cmds}
}

func (c zrangebyscoremax) Limit(offset int64, count int64) zrangebyscorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangebyscorelimit{cmds: cmds}
}

func (c zrangebyscoremax) Build() []string {
	return c.cmds
}

type zrangebyscoremin struct {
	cmds []string
}

func (c zrangebyscoremin) Max(max float64) zrangebyscoremax {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(max, 'f', -1, 64))
	return zrangebyscoremax{cmds: cmds}
}

type zrangebyscorewithscoreswithscores struct {
	cmds []string
}

func (c zrangebyscorewithscoreswithscores) Limit(offset int64, count int64) zrangebyscorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangebyscorelimit{cmds: cmds}
}

func (c zrangebyscorewithscoreswithscores) Build() []string {
	return c.cmds
}

type zrangekey struct {
	cmds []string
}

func (c zrangekey) Min(min string) zrangemin {
	var cmds []string
	cmds = append(c.cmds, min)
	return zrangemin{cmds: cmds}
}

type zrangelimit struct {
	cmds []string
}

func (c zrangelimit) Withscores() zrangewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrangewithscoreswithscores{cmds: cmds}
}

func (c zrangelimit) Build() []string {
	return c.cmds
}

type zrangemax struct {
	cmds []string
}

func (c zrangemax) Byscore() zrangesortbybyscore {
	var cmds []string
	cmds = append(c.cmds, "BYSCORE")
	return zrangesortbybyscore{cmds: cmds}
}

func (c zrangemax) Bylex() zrangesortbybylex {
	var cmds []string
	cmds = append(c.cmds, "BYLEX")
	return zrangesortbybylex{cmds: cmds}
}

func (c zrangemax) Rev() zrangerevrev {
	var cmds []string
	cmds = append(c.cmds, "REV")
	return zrangerevrev{cmds: cmds}
}

func (c zrangemax) Limit(offset int64, count int64) zrangelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangelimit{cmds: cmds}
}

func (c zrangemax) Withscores() zrangewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrangewithscoreswithscores{cmds: cmds}
}

func (c zrangemax) Build() []string {
	return c.cmds
}

type zrangemin struct {
	cmds []string
}

func (c zrangemin) Max(max string) zrangemax {
	var cmds []string
	cmds = append(c.cmds, max)
	return zrangemax{cmds: cmds}
}

type zrangerevrev struct {
	cmds []string
}

func (c zrangerevrev) Limit(offset int64, count int64) zrangelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangelimit{cmds: cmds}
}

func (c zrangerevrev) Withscores() zrangewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrangewithscoreswithscores{cmds: cmds}
}

func (c zrangerevrev) Build() []string {
	return c.cmds
}

type zrangesortbybylex struct {
	cmds []string
}

func (c zrangesortbybylex) Rev() zrangerevrev {
	var cmds []string
	cmds = append(c.cmds, "REV")
	return zrangerevrev{cmds: cmds}
}

func (c zrangesortbybylex) Limit(offset int64, count int64) zrangelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangelimit{cmds: cmds}
}

func (c zrangesortbybylex) Withscores() zrangewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrangewithscoreswithscores{cmds: cmds}
}

func (c zrangesortbybylex) Build() []string {
	return c.cmds
}

type zrangesortbybyscore struct {
	cmds []string
}

func (c zrangesortbybyscore) Rev() zrangerevrev {
	var cmds []string
	cmds = append(c.cmds, "REV")
	return zrangerevrev{cmds: cmds}
}

func (c zrangesortbybyscore) Limit(offset int64, count int64) zrangelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangelimit{cmds: cmds}
}

func (c zrangesortbybyscore) Withscores() zrangewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrangewithscoreswithscores{cmds: cmds}
}

func (c zrangesortbybyscore) Build() []string {
	return c.cmds
}

type zrangestore struct {
	cmds []string
}

func (c zrangestore) Dst(dst string) zrangestoredst {
	var cmds []string
	cmds = append(c.cmds, dst)
	return zrangestoredst{cmds: cmds}
}

func Zrangestore() (c zrangestore) {
	c.cmds = append(c.cmds, "ZRANGESTORE")
	return
}

type zrangestoredst struct {
	cmds []string
}

func (c zrangestoredst) Src(src string) zrangestoresrc {
	var cmds []string
	cmds = append(c.cmds, src)
	return zrangestoresrc{cmds: cmds}
}

type zrangestorelimit struct {
	cmds []string
}

func (c zrangestorelimit) Build() []string {
	return c.cmds
}

type zrangestoremax struct {
	cmds []string
}

func (c zrangestoremax) Byscore() zrangestoresortbybyscore {
	var cmds []string
	cmds = append(c.cmds, "BYSCORE")
	return zrangestoresortbybyscore{cmds: cmds}
}

func (c zrangestoremax) Bylex() zrangestoresortbybylex {
	var cmds []string
	cmds = append(c.cmds, "BYLEX")
	return zrangestoresortbybylex{cmds: cmds}
}

func (c zrangestoremax) Rev() zrangestorerevrev {
	var cmds []string
	cmds = append(c.cmds, "REV")
	return zrangestorerevrev{cmds: cmds}
}

func (c zrangestoremax) Limit(offset int64, count int64) zrangestorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangestorelimit{cmds: cmds}
}

func (c zrangestoremax) Build() []string {
	return c.cmds
}

type zrangestoremin struct {
	cmds []string
}

func (c zrangestoremin) Max(max string) zrangestoremax {
	var cmds []string
	cmds = append(c.cmds, max)
	return zrangestoremax{cmds: cmds}
}

type zrangestorerevrev struct {
	cmds []string
}

func (c zrangestorerevrev) Limit(offset int64, count int64) zrangestorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangestorelimit{cmds: cmds}
}

func (c zrangestorerevrev) Build() []string {
	return c.cmds
}

type zrangestoresortbybylex struct {
	cmds []string
}

func (c zrangestoresortbybylex) Rev() zrangestorerevrev {
	var cmds []string
	cmds = append(c.cmds, "REV")
	return zrangestorerevrev{cmds: cmds}
}

func (c zrangestoresortbybylex) Limit(offset int64, count int64) zrangestorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangestorelimit{cmds: cmds}
}

func (c zrangestoresortbybylex) Build() []string {
	return c.cmds
}

type zrangestoresortbybyscore struct {
	cmds []string
}

func (c zrangestoresortbybyscore) Rev() zrangestorerevrev {
	var cmds []string
	cmds = append(c.cmds, "REV")
	return zrangestorerevrev{cmds: cmds}
}

func (c zrangestoresortbybyscore) Limit(offset int64, count int64) zrangestorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrangestorelimit{cmds: cmds}
}

func (c zrangestoresortbybyscore) Build() []string {
	return c.cmds
}

type zrangestoresrc struct {
	cmds []string
}

func (c zrangestoresrc) Min(min string) zrangestoremin {
	var cmds []string
	cmds = append(c.cmds, min)
	return zrangestoremin{cmds: cmds}
}

type zrangewithscoreswithscores struct {
	cmds []string
}

func (c zrangewithscoreswithscores) Build() []string {
	return c.cmds
}

type zrank struct {
	cmds []string
}

func (c zrank) Key(key string) zrankkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrankkey{cmds: cmds}
}

func Zrank() (c zrank) {
	c.cmds = append(c.cmds, "ZRANK")
	return
}

type zrankkey struct {
	cmds []string
}

func (c zrankkey) Member(member string) zrankmember {
	var cmds []string
	cmds = append(c.cmds, member)
	return zrankmember{cmds: cmds}
}

type zrankmember struct {
	cmds []string
}

func (c zrankmember) Build() []string {
	return c.cmds
}

type zrem struct {
	cmds []string
}

func (c zrem) Key(key string) zremkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zremkey{cmds: cmds}
}

func Zrem() (c zrem) {
	c.cmds = append(c.cmds, "ZREM")
	return
}

type zremkey struct {
	cmds []string
}

func (c zremkey) Member(member ...string) zremmember {
	var cmds []string
	cmds = append(cmds, member...)
	return zremmember{cmds: cmds}
}

type zremmember struct {
	cmds []string
}

func (c zremmember) Member(member ...string) zremmember {
	var cmds []string
	cmds = append(cmds, member...)
	return zremmember{cmds: cmds}
}

type zremrangebylex struct {
	cmds []string
}

func (c zremrangebylex) Key(key string) zremrangebylexkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zremrangebylexkey{cmds: cmds}
}

func Zremrangebylex() (c zremrangebylex) {
	c.cmds = append(c.cmds, "ZREMRANGEBYLEX")
	return
}

type zremrangebylexkey struct {
	cmds []string
}

func (c zremrangebylexkey) Min(min string) zremrangebylexmin {
	var cmds []string
	cmds = append(c.cmds, min)
	return zremrangebylexmin{cmds: cmds}
}

type zremrangebylexmax struct {
	cmds []string
}

func (c zremrangebylexmax) Build() []string {
	return c.cmds
}

type zremrangebylexmin struct {
	cmds []string
}

func (c zremrangebylexmin) Max(max string) zremrangebylexmax {
	var cmds []string
	cmds = append(c.cmds, max)
	return zremrangebylexmax{cmds: cmds}
}

type zremrangebyrank struct {
	cmds []string
}

func (c zremrangebyrank) Key(key string) zremrangebyrankkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zremrangebyrankkey{cmds: cmds}
}

func Zremrangebyrank() (c zremrangebyrank) {
	c.cmds = append(c.cmds, "ZREMRANGEBYRANK")
	return
}

type zremrangebyrankkey struct {
	cmds []string
}

func (c zremrangebyrankkey) Start(start int64) zremrangebyrankstart {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(start, 10))
	return zremrangebyrankstart{cmds: cmds}
}

type zremrangebyrankstart struct {
	cmds []string
}

func (c zremrangebyrankstart) Stop(stop int64) zremrangebyrankstop {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(stop, 10))
	return zremrangebyrankstop{cmds: cmds}
}

type zremrangebyrankstop struct {
	cmds []string
}

func (c zremrangebyrankstop) Build() []string {
	return c.cmds
}

type zremrangebyscore struct {
	cmds []string
}

func (c zremrangebyscore) Key(key string) zremrangebyscorekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zremrangebyscorekey{cmds: cmds}
}

func Zremrangebyscore() (c zremrangebyscore) {
	c.cmds = append(c.cmds, "ZREMRANGEBYSCORE")
	return
}

type zremrangebyscorekey struct {
	cmds []string
}

func (c zremrangebyscorekey) Min(min float64) zremrangebyscoremin {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(min, 'f', -1, 64))
	return zremrangebyscoremin{cmds: cmds}
}

type zremrangebyscoremax struct {
	cmds []string
}

func (c zremrangebyscoremax) Build() []string {
	return c.cmds
}

type zremrangebyscoremin struct {
	cmds []string
}

func (c zremrangebyscoremin) Max(max float64) zremrangebyscoremax {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(max, 'f', -1, 64))
	return zremrangebyscoremax{cmds: cmds}
}

type zrevrange struct {
	cmds []string
}

func (c zrevrange) Key(key string) zrevrangekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrevrangekey{cmds: cmds}
}

func Zrevrange() (c zrevrange) {
	c.cmds = append(c.cmds, "ZREVRANGE")
	return
}

type zrevrangebylex struct {
	cmds []string
}

func (c zrevrangebylex) Key(key string) zrevrangebylexkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrevrangebylexkey{cmds: cmds}
}

func Zrevrangebylex() (c zrevrangebylex) {
	c.cmds = append(c.cmds, "ZREVRANGEBYLEX")
	return
}

type zrevrangebylexkey struct {
	cmds []string
}

func (c zrevrangebylexkey) Max(max string) zrevrangebylexmax {
	var cmds []string
	cmds = append(c.cmds, max)
	return zrevrangebylexmax{cmds: cmds}
}

type zrevrangebylexlimit struct {
	cmds []string
}

func (c zrevrangebylexlimit) Build() []string {
	return c.cmds
}

type zrevrangebylexmax struct {
	cmds []string
}

func (c zrevrangebylexmax) Min(min string) zrevrangebylexmin {
	var cmds []string
	cmds = append(c.cmds, min)
	return zrevrangebylexmin{cmds: cmds}
}

type zrevrangebylexmin struct {
	cmds []string
}

func (c zrevrangebylexmin) Limit(offset int64, count int64) zrevrangebylexlimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrevrangebylexlimit{cmds: cmds}
}

func (c zrevrangebylexmin) Build() []string {
	return c.cmds
}

type zrevrangebyscore struct {
	cmds []string
}

func (c zrevrangebyscore) Key(key string) zrevrangebyscorekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrevrangebyscorekey{cmds: cmds}
}

func Zrevrangebyscore() (c zrevrangebyscore) {
	c.cmds = append(c.cmds, "ZREVRANGEBYSCORE")
	return
}

type zrevrangebyscorekey struct {
	cmds []string
}

func (c zrevrangebyscorekey) Max(max float64) zrevrangebyscoremax {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(max, 'f', -1, 64))
	return zrevrangebyscoremax{cmds: cmds}
}

type zrevrangebyscorelimit struct {
	cmds []string
}

func (c zrevrangebyscorelimit) Build() []string {
	return c.cmds
}

type zrevrangebyscoremax struct {
	cmds []string
}

func (c zrevrangebyscoremax) Min(min float64) zrevrangebyscoremin {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatFloat(min, 'f', -1, 64))
	return zrevrangebyscoremin{cmds: cmds}
}

type zrevrangebyscoremin struct {
	cmds []string
}

func (c zrevrangebyscoremin) Withscores() zrevrangebyscorewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrevrangebyscorewithscoreswithscores{cmds: cmds}
}

func (c zrevrangebyscoremin) Limit(offset int64, count int64) zrevrangebyscorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrevrangebyscorelimit{cmds: cmds}
}

func (c zrevrangebyscoremin) Build() []string {
	return c.cmds
}

type zrevrangebyscorewithscoreswithscores struct {
	cmds []string
}

func (c zrevrangebyscorewithscoreswithscores) Limit(offset int64, count int64) zrevrangebyscorelimit {
	var cmds []string
	cmds = append(c.cmds, "LIMIT", strconv.FormatInt(offset, 10), strconv.FormatInt(count, 10))
	return zrevrangebyscorelimit{cmds: cmds}
}

func (c zrevrangebyscorewithscoreswithscores) Build() []string {
	return c.cmds
}

type zrevrangekey struct {
	cmds []string
}

func (c zrevrangekey) Start(start int64) zrevrangestart {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(start, 10))
	return zrevrangestart{cmds: cmds}
}

type zrevrangestart struct {
	cmds []string
}

func (c zrevrangestart) Stop(stop int64) zrevrangestop {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(stop, 10))
	return zrevrangestop{cmds: cmds}
}

type zrevrangestop struct {
	cmds []string
}

func (c zrevrangestop) Withscores() zrevrangewithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zrevrangewithscoreswithscores{cmds: cmds}
}

func (c zrevrangestop) Build() []string {
	return c.cmds
}

type zrevrangewithscoreswithscores struct {
	cmds []string
}

func (c zrevrangewithscoreswithscores) Build() []string {
	return c.cmds
}

type zrevrank struct {
	cmds []string
}

func (c zrevrank) Key(key string) zrevrankkey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zrevrankkey{cmds: cmds}
}

func Zrevrank() (c zrevrank) {
	c.cmds = append(c.cmds, "ZREVRANK")
	return
}

type zrevrankkey struct {
	cmds []string
}

func (c zrevrankkey) Member(member string) zrevrankmember {
	var cmds []string
	cmds = append(c.cmds, member)
	return zrevrankmember{cmds: cmds}
}

type zrevrankmember struct {
	cmds []string
}

func (c zrevrankmember) Build() []string {
	return c.cmds
}

type zscan struct {
	cmds []string
}

func (c zscan) Key(key string) zscankey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zscankey{cmds: cmds}
}

func Zscan() (c zscan) {
	c.cmds = append(c.cmds, "ZSCAN")
	return
}

type zscancount struct {
	cmds []string
}

func (c zscancount) Build() []string {
	return c.cmds
}

type zscancursor struct {
	cmds []string
}

func (c zscancursor) Match(pattern string) zscanmatch {
	var cmds []string
	cmds = append(c.cmds, "MATCH", pattern)
	return zscanmatch{cmds: cmds}
}

func (c zscancursor) Count(count int64) zscancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return zscancount{cmds: cmds}
}

func (c zscancursor) Build() []string {
	return c.cmds
}

type zscankey struct {
	cmds []string
}

func (c zscankey) Cursor(cursor int64) zscancursor {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(cursor, 10))
	return zscancursor{cmds: cmds}
}

type zscanmatch struct {
	cmds []string
}

func (c zscanmatch) Count(count int64) zscancount {
	var cmds []string
	cmds = append(c.cmds, "COUNT", strconv.FormatInt(count, 10))
	return zscancount{cmds: cmds}
}

func (c zscanmatch) Build() []string {
	return c.cmds
}

type zscore struct {
	cmds []string
}

func (c zscore) Key(key string) zscorekey {
	var cmds []string
	cmds = append(c.cmds, key)
	return zscorekey{cmds: cmds}
}

func Zscore() (c zscore) {
	c.cmds = append(c.cmds, "ZSCORE")
	return
}

type zscorekey struct {
	cmds []string
}

func (c zscorekey) Member(member string) zscoremember {
	var cmds []string
	cmds = append(c.cmds, member)
	return zscoremember{cmds: cmds}
}

type zscoremember struct {
	cmds []string
}

func (c zscoremember) Build() []string {
	return c.cmds
}

type zunion struct {
	cmds []string
}

func (c zunion) Numkeys(numkeys int64) zunionnumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return zunionnumkeys{cmds: cmds}
}

func Zunion() (c zunion) {
	c.cmds = append(c.cmds, "ZUNION")
	return
}

type zunionaggregatemax struct {
	cmds []string
}

func (c zunionaggregatemax) Withscores() zunionwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zunionwithscoreswithscores{cmds: cmds}
}

func (c zunionaggregatemax) Build() []string {
	return c.cmds
}

type zunionaggregatemin struct {
	cmds []string
}

func (c zunionaggregatemin) Withscores() zunionwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zunionwithscoreswithscores{cmds: cmds}
}

func (c zunionaggregatemin) Build() []string {
	return c.cmds
}

type zunionaggregatesum struct {
	cmds []string
}

func (c zunionaggregatesum) Withscores() zunionwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zunionwithscoreswithscores{cmds: cmds}
}

func (c zunionaggregatesum) Build() []string {
	return c.cmds
}

type zunionkey struct {
	cmds []string
}

func (c zunionkey) Weights(weight int64) zunionweights {
	var cmds []string
	cmds = append(c.cmds, "WEIGHTS", strconv.FormatInt(weight, 10))
	return zunionweights{cmds: cmds}
}

func (c zunionkey) Sum() zunionaggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zunionaggregatesum{cmds: cmds}
}

func (c zunionkey) Min() zunionaggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zunionaggregatemin{cmds: cmds}
}

func (c zunionkey) Max() zunionaggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zunionaggregatemax{cmds: cmds}
}

func (c zunionkey) Withscores() zunionwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zunionwithscoreswithscores{cmds: cmds}
}

func (c zunionkey) Key(key ...string) zunionkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zunionkey{cmds: cmds}
}

type zunionnumkeys struct {
	cmds []string
}

func (c zunionnumkeys) Key(key ...string) zunionkey {
	var cmds []string
	cmds = append(cmds, key...)
	return zunionkey{cmds: cmds}
}

type zunionstore struct {
	cmds []string
}

func (c zunionstore) Destination(destination string) zunionstoredestination {
	var cmds []string
	cmds = append(c.cmds, destination)
	return zunionstoredestination{cmds: cmds}
}

func Zunionstore() (c zunionstore) {
	c.cmds = append(c.cmds, "ZUNIONSTORE")
	return
}

type zunionstoreaggregatemax struct {
	cmds []string
}

func (c zunionstoreaggregatemax) Build() []string {
	return c.cmds
}

type zunionstoreaggregatemin struct {
	cmds []string
}

func (c zunionstoreaggregatemin) Build() []string {
	return c.cmds
}

type zunionstoreaggregatesum struct {
	cmds []string
}

func (c zunionstoreaggregatesum) Build() []string {
	return c.cmds
}

type zunionstoredestination struct {
	cmds []string
}

func (c zunionstoredestination) Numkeys(numkeys int64) zunionstorenumkeys {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(numkeys, 10))
	return zunionstorenumkeys{cmds: cmds}
}

type zunionstorekey struct {
	cmds []string
}

func (c zunionstorekey) Weights(weight int64) zunionstoreweights {
	var cmds []string
	cmds = append(c.cmds, "WEIGHTS", strconv.FormatInt(weight, 10))
	return zunionstoreweights{cmds: cmds}
}

func (c zunionstorekey) Sum() zunionstoreaggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zunionstoreaggregatesum{cmds: cmds}
}

func (c zunionstorekey) Min() zunionstoreaggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zunionstoreaggregatemin{cmds: cmds}
}

func (c zunionstorekey) Max() zunionstoreaggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zunionstoreaggregatemax{cmds: cmds}
}

func (c zunionstorekey) Key(key ...string) zunionstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return zunionstorekey{cmds: cmds}
}

type zunionstorenumkeys struct {
	cmds []string
}

func (c zunionstorenumkeys) Key(key ...string) zunionstorekey {
	var cmds []string
	cmds = append(cmds, key...)
	return zunionstorekey{cmds: cmds}
}

type zunionstoreweights struct {
	cmds []string
}

func (c zunionstoreweights) Sum() zunionstoreaggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zunionstoreaggregatesum{cmds: cmds}
}

func (c zunionstoreweights) Min() zunionstoreaggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zunionstoreaggregatemin{cmds: cmds}
}

func (c zunionstoreweights) Max() zunionstoreaggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zunionstoreaggregatemax{cmds: cmds}
}

func (c zunionstoreweights) Weight(weight int64) zunionstoreweights {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(weight, 10))
	return zunionstoreweights{cmds: cmds}
}

func (c zunionstoreweights) Build() []string {
	return c.cmds
}

type zunionweights struct {
	cmds []string
}

func (c zunionweights) Sum() zunionaggregatesum {
	var cmds []string
	cmds = append(c.cmds, "SUM")
	return zunionaggregatesum{cmds: cmds}
}

func (c zunionweights) Min() zunionaggregatemin {
	var cmds []string
	cmds = append(c.cmds, "MIN")
	return zunionaggregatemin{cmds: cmds}
}

func (c zunionweights) Max() zunionaggregatemax {
	var cmds []string
	cmds = append(c.cmds, "MAX")
	return zunionaggregatemax{cmds: cmds}
}

func (c zunionweights) Withscores() zunionwithscoreswithscores {
	var cmds []string
	cmds = append(c.cmds, "WITHSCORES")
	return zunionwithscoreswithscores{cmds: cmds}
}

func (c zunionweights) Weight(weight int64) zunionweights {
	var cmds []string
	cmds = append(c.cmds, strconv.FormatInt(weight, 10))
	return zunionweights{cmds: cmds}
}

func (c zunionweights) Build() []string {
	return c.cmds
}

type zunionwithscoreswithscores struct {
	cmds []string
}

func (c zunionwithscoreswithscores) Build() []string {
	return c.cmds
}
