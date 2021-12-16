package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"
)

type GoStruct struct {
	Node      *Node
	FullName  string
	BuildDef  BuildDef
	Variadic  bool
	NextNodes []*Node
}

type BuildDef struct {
	MethodName string
	Command    []string
	Parameters []Parameter
}

type Parameter struct {
	Name string
	Type string
}

type Command struct {
	Group     string     `json:"group"`
	Since     string     `json:"since"`
	Arguments []Argument `json:"arguments"`
}

type Argument struct {
	Name     interface{} `json:"name"`
	Type     interface{} `json:"type"`
	Command  string      `json:"command"`
	Enum     []string    `json:"enum"`
	Multiple bool        `json:"multiple"`
	Optional bool        `json:"optional"`
	Variadic bool        `json:"variadic"`
	Block    []Argument  `json:"block"`
}

type Node struct {
	Parent *Node
	Child  *Node
	Next   *Node
	Cmd    Command
	Arg    Argument
	Root   bool
}

func (n *Node) FindRoot() (root *Node) {
	root = n
	for root.Parent != nil {
		root = root.Parent
	}
	return
}

func (n *Node) GoStructs() (out []GoStruct) {
	fn := n.FullName()
	// fix for XGROUP and XADD
	if len(n.Arg.Enum) == 2 && ((n.Arg.Enum[0] == "id" && n.Arg.Enum[1] == "$") || (n.Arg.Enum[0] == "*" && n.Arg.Enum[1] == "ID")) {
		n.Arg.Enum = nil
		n.Arg.Name = "id"
		n.Arg.Type = "string"
		if !strings.HasSuffix(fn, "Id") {
			fn += "Id"
		}
	}

	if len(n.Arg.Block) > 0 {
		panic("GoStructs should not be called on Block Node")
	} else if len(n.Arg.Enum) > 0 {
		for _, e := range n.Arg.Enum {
			s := GoStruct{
				Node:     n,
				FullName: fn,
				BuildDef: BuildDef{
					MethodName: name(e),
					Command:    nil,
					Parameters: nil,
				},
				Variadic:  n.Variadic(),
				NextNodes: n.NextNodes(),
			}
			if len(n.Arg.Command) != 0 {
				if !strings.HasSuffix(s.FullName, name(n.Arg.Command)) {
					s.FullName += name(n.Arg.Command)
				}
				s.BuildDef.MethodName = name(n.Arg.Command) + name(e)
				s.BuildDef.Command = strings.Split(n.Arg.Command, " ")
			}
			if !strings.HasSuffix(s.FullName, name(e)) {
				s.FullName += name(e)
			}
			cmds := strings.Split(e, " ")
			if len(cmds) == 1 {
				s.BuildDef.Command = append(s.BuildDef.Command, cmds...)
				s.BuildDef.Parameters = nil
			} else {
				switch cmds[1] {
				case "seconds", "milliseconds", "timestamp", "milliseconds-timestamp":
					s.BuildDef.Command = append(s.BuildDef.Command, cmds[:1]...)
					s.BuildDef.Parameters = []Parameter{{Name: LcFirst(name(cmds[1])), Type: "integer"}}
				case "*":
					if cmds[0] == "LOAD" {
						s.BuildDef.Command = append(s.BuildDef.Command, cmds...)
						s.BuildDef.Parameters = nil
					}
				default:
					panic("unknown enum")
				}
			}
			out = append(out, s)
		}
		return
	} else {
		s := GoStruct{
			Node:     n,
			FullName: fn,
			BuildDef: BuildDef{
				MethodName: n.Name(),
				Parameters: nil,
			},
			Variadic:  n.Variadic(),
			NextNodes: n.NextNodes(),
		}
		if len(n.Arg.Command) != 0 {
			s.BuildDef.Command = strings.Split(n.Arg.Command, " ")
		}

		switch nm := n.Arg.Name.(type) {
		case string:
			if s.FullName == "FtCreatePrefixPrefix" && nm == "count" && n.Arg.Command == "PREFIX" {
				s.FullName = "FtCreatePrefixCount"
			}
			s.BuildDef.Parameters = []Parameter{{Name: LcFirst(name(nm)), Type: n.Arg.Type.(string)}} // not change to go type here, change at render
		case []interface{}:
			for i, nn := range nm {
				s.BuildDef.Parameters = append(s.BuildDef.Parameters, Parameter{Name: LcFirst(name(nn.(string))), Type: n.Arg.Type.([]interface{})[i].(string)})
			}
		default:
			if n.Arg.Type == nil || (n.Arg.Type != nil && n.Arg.Type.(string) == "command") {
				// ignore
			} else {
				panic("unknown arg name")
			}
		}
		out = append(out, s)
	}
	return
}

func (n *Node) Variadic() bool {
	return n.Arg.Multiple || n.Arg.Variadic
}

func (n *Node) FullName() (out string) {
	if n.Parent != nil {
		return n.Parent.FullName() + n.Name()
	}
	return n.Name()
}

func (n *Node) Name() (out string) {
	var tokens []string
	if n.Arg.Command != "" {
		tokens = append(tokens, name(n.Arg.Command))
	} else {
		switch n := n.Arg.Name.(type) {
		case string:
			tokens = append(tokens, name(n))
		case []interface{}:
			for _, nn := range n {
				tokens = append(tokens, name(nn.(string)))
			}
		}
		if len(tokens) == 0 {
			if n.Child != nil {
				tokens = append(tokens, n.Child.Name())
			}
		}
	}

	duplicate := map[string]bool{}
	for _, t := range tokens {
		if len(t) == 0 || duplicate[t] {
			continue
		}
		duplicate[t] = true
		out += t
	}
	return
}

func (n *Node) NextNodes() []*Node {
	var nodes []*Node

	if n.Child == nil && n.Variadic() {
		nodes = append(nodes, n)
	}

	if n.Child != nil {
		nodes = append(nodes, n.Child)
	}

	parent := n
	for parent != nil {
		next := parent.Next
		for next != nil {
			nodes = append(nodes, next)
			if !next.Arg.Optional {
				return nodes
			}
			next = next.Next
		}

		parent = parent.Parent
		// block variadic
		if parent != nil && parent.Variadic() {
			nodes = append(nodes, parent.Child)
		}
		if parent != nil && parent.Root {
			break // don't climb to root
		}
	}
	return nodes
}

func (n *Node) Walk(fn func(node *Node)) {
	next := n
	for next != nil {
		if next.Child != nil {
			next.Child.Walk(fn)
		} else {
			fn(next)
		}
		next = next.Next
	}
}

func main() {
	var commands = map[string]Command{}

	for _, p := range []string{
		"./commands.json",
		"./commands_json.json",
		"./commands_bloom.json",
		"./commands_search.json",
	} {
		raw, err := os.ReadFile(p)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(raw, &commands); err != nil {
			panic(err)
		}
	}

	// fis missing GEORADIUS_RO and GEORADIUSBYMEMBER_RO
	commands["GEORADIUS_RO"] = Command{Arguments: filterArgs(commands["GEORADIUS"].Arguments, "STORE")}
	commands["GEORADIUSBYMEMBER_RO"] = Command{Arguments: filterArgs(commands["GEORADIUSBYMEMBER"].Arguments, "STORE")}

	var roots []string
	nodes := map[string]*Node{}
	for k, cmd := range commands {
		root := &Node{Cmd: cmd, Arg: Argument{Name: k, Command: k, Type: "command"}, Root: true}
		root.Next = makeChildNodes(root, cmd.Arguments)
		roots = append(roots, k)
		nodes[k] = root
	}
	sort.Strings(roots)

	var structs = map[string]GoStruct{}
	for _, name := range roots {
		node := nodes[name]
		node.Walk(func(n *Node) {
			for _, s := range n.GoStructs() {
				if v, ok := structs[s.FullName]; ok {
					panic("struct conflict " + v.FullName)
				}
				structs[s.FullName] = s
			}
		})
	}

	generate(structs)
}

func generate(structs map[string]GoStruct) {
	var names []string
	for name := range structs {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Fprintf(os.Stdout, "// Code generated DO NOT EDIT\n\npackage cmds\n\n")
	fmt.Fprintf(os.Stdout, "import %q\n\n", "strconv")

	for _, name := range names {
		s := structs[name]

		fmt.Fprintf(os.Stdout, "type %s Completed\n\n", s.FullName)
		fmt.Fprintf(os.Stdout, "type %s SCompleted\n\n", "S"+s.FullName)

		if s.Node.Root {
			printRootBuilder(os.Stdout, s, "")
			printRootBuilder(os.Stdout, s, "S")
		}

		for _, next := range s.NextNodes {
			if next.Child != nil {
				next = next.Child
			}
			for _, ss := range next.GoStructs() {
				printBuilder(os.Stdout, s, ss, "")
				printBuilder(os.Stdout, s, ss, "S")
			}
		}

		if allOptional(s.Node, s.NextNodes) {
			printFinalBuilder(os.Stdout, s, "", "Build", "Completed")
			printFinalBuilder(os.Stdout, s, "S", "Build", "SCompleted")
			if within(s.Node.FindRoot().GoStructs()[0], cacheableCMDs) {
				printFinalBuilder(os.Stdout, s, "", "Cache", "Cacheable")
				printFinalBuilder(os.Stdout, s, "S", "Cache", "SCacheable")
			}
		}
	}
}

func allOptional(s *Node, nodes []*Node) bool {
	for _, n := range nodes {
		if s == n {
			continue
		}
		if n.Parent != nil && n.Parent.Child == n {
			continue
		}
		if !n.Arg.Optional {
			return false
		}
	}
	return true
}

func toGoType(paramType string) string {
	switch paramType {
	case "key", "string", "pattern", "type":
		return "string"
	case "double":
		return "float64"
	case "integer", "posix time":
		return "int64"
	default:
		panic("unknown param type " + paramType)
	}
}

func toGoName(paramName string) string {
	if paramName == "type" {
		return "typ"
	}
	return paramName
}

func printRootBuilder(w io.Writer, root GoStruct, prefix string) {
	fmt.Fprintf(w, "func (b *%sBuilder) %s() %s%s {\n", prefix, root.FullName, prefix, root.FullName)

	var appends []string
	for _, cmd := range root.BuildDef.Command {
		appends = append(appends, fmt.Sprintf(`"%s"`, cmd))
	}

	if tag := rootCf(root); tag != "" {
		fmt.Fprintf(w, "\treturn %s%s{cs: append(b.get(), %s), ks: InitSlot, cf: %s}\n", prefix, root.FullName, strings.Join(appends, ", "), tag)
	} else {
		fmt.Fprintf(w, "\treturn %s%s{cs: append(b.get(), %s), ks: InitSlot}\n", prefix, root.FullName, strings.Join(appends, ", "))
	}

	fmt.Fprintf(w, "}\n\n")
}

func rootCf(root GoStruct) (tag string) {
	if within(root, blockingCMDs) {
		if tag != "" {
			panic("root cf collision")
		}
		tag = "blockTag"
	}

	if within(root, noRetCMDs) {
		if tag != "" {
			panic("root cf collision")
		}
		tag = "noRetTag"
	}

	if within(root, readOnlyCMDs) {
		if tag != "" {
			panic("root cf collision")
		}
		tag = "readonly"
	}

	return tag
}

func printFinalBuilder(w io.Writer, parent GoStruct, prefix, method, ss string) {
	fmt.Fprintf(w, "func (c %s%s) %s() %s {\n", prefix, parent.FullName, method, ss)
	fmt.Fprintf(w, "\treturn %s(c)\n", ss)
	fmt.Fprintf(w, "}\n\n")
}

func printBuilder(w io.Writer, parent, next GoStruct, prefix string) {
	fmt.Fprintf(w, "func (c %s%s) %s(", prefix, parent.FullName, next.BuildDef.MethodName)
	if len(next.BuildDef.Parameters) == 1 && next.Variadic {
		fmt.Fprintf(w, "%s ...%s", toGoName(next.BuildDef.Parameters[0].Name), toGoType(next.BuildDef.Parameters[0].Type))
	} else if next.Variadic && parent.FullName != next.FullName {
		// no parameter
	} else {
		for i, param := range next.BuildDef.Parameters {
			fmt.Fprintf(w, "%s %s", toGoName(param.Name), toGoType(param.Type))
			if i != len(next.BuildDef.Parameters)-1 {
				fmt.Fprintf(w, ", ")
			}
		}
	}
	fmt.Fprintf(w, ") %s%s {\n", prefix, next.FullName)

	if prefix == "S" {
		if len(next.BuildDef.Parameters) == 1 && next.Variadic {
			if next.BuildDef.Parameters[0].Type == "key" {
				fmt.Printf("\tfor _, k := range %s {\n", toGoName(next.BuildDef.Parameters[0].Name))
				fmt.Printf("\t\tc.ks = checkSlot(c.ks, slot(k))\n")
				fmt.Printf("\t}\n")
			}
		} else {
			if len(next.BuildDef.Parameters) != 1 && next.Variadic && parent.FullName != next.FullName {
				// no param
			} else {
				for _, arg := range next.BuildDef.Parameters {
					if arg.Type == "key" {
						fmt.Printf("\tc.ks = checkSlot(c.ks, slot(%s))\n", toGoName(arg.Name))
					}
				}
			}
		}
	}

	for _, cmd := range next.BuildDef.Command {
		if cmd == "BLOCK" {
			fmt.Printf("\tc.cf = blockTag\n")
			break
		}
	}

	var appends []string

	for _, cmd := range next.BuildDef.Command {
		if cmd == "" {
			panic("unexpected empty command on " + next.FullName)
		}
		if cmd == `""` {
			appends = append(appends, `""`)
		} else {
			appends = append(appends, fmt.Sprintf(`"%s"`, cmd))
		}
	}

	if len(appends) == 0 && next.Variadic && len(next.BuildDef.Parameters) == 1 && toGoType(next.BuildDef.Parameters[0].Type) == "string" {
		appends = append(appends, toGoName(next.BuildDef.Parameters[0].Name)+"...")
		fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
	} else if len(next.BuildDef.Parameters) != 1 && next.Variadic && parent.FullName != next.FullName {
		// no parameter
	} else {
		allstring := true
		for _, p := range next.BuildDef.Parameters {
			if toGoType(p.Type) != "string" {
				allstring = false
				break
			}
		}
		if allstring && !next.Variadic {
			for _, p := range next.BuildDef.Parameters {
				appends = append(appends, toGoName(p.Name))
			}
			fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
		} else {
			if len(next.BuildDef.Parameters) == 1 && next.Variadic {
				if len(appends) != 0 {
					fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
				}
				if toGoType(next.BuildDef.Parameters[0].Type) == "string" {
					fmt.Fprintf(w, "\tc.cs = append(c.cs, %s...)\n", toGoName(next.BuildDef.Parameters[0].Name))
				} else {
					fmt.Printf("\tfor _, n := range %s {\n", toGoName(next.BuildDef.Parameters[0].Name))
					switch toGoType(next.BuildDef.Parameters[0].Type) {
					case "float64":
						fmt.Printf("\t\tc.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))\n")
					case "int64":
						fmt.Printf("\t\tc.cs = append(c.cs, strconv.FormatInt(n, 10))\n")
					default:
						panic("unexpected param type " + next.BuildDef.Parameters[0].Type)
					}
					fmt.Printf("\t}\n")
				}
			} else {
				for _, p := range next.BuildDef.Parameters {
					switch toGoType(p.Type) {
					case "float64":
						appends = append(appends, fmt.Sprintf("strconv.FormatFloat(%s, 'f', -1, 64)", toGoName(p.Name)))
					case "int64":
						appends = append(appends, fmt.Sprintf("strconv.FormatInt(%s, 10)", toGoName(p.Name)))
					case "string":
						appends = append(appends, toGoName(p.Name))
					default:
						panic("unexpected param type " + next.BuildDef.Parameters[0].Type)
					}
				}
				fmt.Fprintf(w, "\tc.cs = append(c.cs, %s)\n", strings.Join(appends, ", "))
			}
		}
	}

	if parent.FullName == next.FullName {
		fmt.Fprintf(w, "\treturn c\n")
	} else {
		fmt.Fprintf(w, "\treturn (%s%s)(c)\n", prefix, next.FullName)
	}
	fmt.Fprintf(w, "}\n\n")
}

func makeChildNodes(parent *Node, args []Argument) (first *Node) {
	if len(args) == 0 {
		return nil
	}
	var nodes []*Node
	for _, arg := range args {
		nodes = append(nodes, &Node{Parent: parent, Arg: arg})
	}
	for i, node := range nodes {
		node.Child = makeChildNodes(node, node.Arg.Block)
		if len(node.Arg.Block) != 0 && node.Arg.Command != "" {
			if node.Child.Arg.Optional {
				panic("unexpected block command with optional child")
			}
			if node.Child.Arg.Command == "" {
				node.Child.Arg.Command = node.Arg.Command
			} else {
				node.Child.Arg.Command = node.Arg.Command + " " + node.Child.Arg.Command
			}
		}
		if i != len(nodes)-1 {
			node.Next = nodes[i+1]
		}
	}
	return nodes[0]
}

func filterArgs(args []Argument, exclude string) (out []Argument) {
	for _, a := range args {
		if a.Command != exclude {
			out = append(out, a)
		}
	}
	return out
}

func name(s string) (name string) {
	switch s {
	case "~":
		return "Almost"
	case "=":
		return "Exact"
	case "$":
		return "Last"
	case "\"\"":
		return "Empty"
	}
	for _, n := range strings.Split(strings.NewReplacer("-", " ", "_", " ", ":", " ", "/", " ", ".", " ", "*", "all").Replace(s), " ") {
		name += UcFirst(strings.ToLower(n))
	}
	return name
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func within(cmd GoStruct, cmds []string) bool {
	n := strings.ToLower(cmd.FullName)
	for _, v := range cmds {
		if v == n {
			return true
		}
	}
	return false
}

var noRetCMDs = []string{
	"subscribe",
	"psubscribe",
	"unsubscribe",
	"punsubscribe",
}

var blockingCMDs = []string{
	"blpop",
	"brpop",
	"brpoplpush",
	"blmove",
	"blmpop",
	"bzpopmin",
	"bzpopmax",
	"bzmpop",
	"clientpause",
	"migrate",
	"wait",
}

var cacheableCMDs = []string{
	"bitcount",
	"bitfieldro",
	"bitpos",
	"expiretime",
	"geodist",
	"geohash",
	"geopos",
	"georadiusro",
	"georadiusbymemberro",
	"geosearch",
	"get",
	"getbit",
	"getrange",
	"hexists",
	"hget",
	"hgetall",
	"hkeys",
	"hlen",
	"hmget",
	"hstrlen",
	"hvals",
	"lindex",
	"llen",
	"lpos",
	"lrange",
	"pexpiretime",
	"pttl",
	"scard",
	"sismember",
	"smembers",
	"smismember",
	"sortro",
	"strlen",
	"substr",
	"ttl",
	"type",
	"zcard",
	"zcount",
	"zlexcount",
	"zmscore",
	"zrange",
	"zrangebylex",
	"zrangebyscore",
	"zrank",
	"zrevrange",
	"zrevrangebylex",
	"zrevrangebyscore",
	"zrevrank",
	"zscore",

	"jsonget",
	"jsonstrlen",
	"jsonarrindex",
	"jsonarrlen",
	"jsonobjkeys",
	"jsonobjlen",
	"jsontype",
	"jsonresp",

	"bfexists",
	"bfinfo",
	"cfexists",
	"cfcount",
	"cfinfo",
	"cmsquery",
	"cmsinfo",
	"topkquery",
	"topkcount",
	"topklist",
	"topkinfo",
}

var readOnlyCMDs = []string{
	"bitcount",
	"bitfieldro",
	"bitpos",
	"dbsize",
	"dump",
	"evalro",
	"evalsharo",
	"exists",
	"expiretime",
	"geodist",
	"geohash",
	"geopos",
	"georadiusro",
	"georadiusbymemberro",
	"geosearch",
	"get",
	"getbit",
	"getrange",
	"hexists",
	"hget",
	"hgetall",
	"hkeys",
	"hlen",
	"hmget",
	"hrandfield",
	"hscan",
	"hstrlen",
	"hvals",
	"keys",
	"lindex",
	"llen",
	"lolwut",
	"lpos",
	"lrange",
	"memory",
	"mget",
	"objectencoding",
	"objectfreq",
	"objecthelp",
	"objectidletime",
	"objectrefcount",
	"pexpiretime",
	"pfcount",
	"post",
	"pttl",
	"pubsubchannels",
	"pubsubnumpat",
	"pubsubnumsub",
	"pubsubhelp",
	"randomkey",
	"scan",
	"scard",
	"sdiff",
	"sinter",
	"sintercard",
	"sismember",
	"slowlogget",
	"slowloglen",
	"slowloghelp",
	"smembers",
	"smismember",
	"sortro",
	"srandmember",
	"sscan",
	"lcs",
	"strlen",
	"substr",
	"sunion",
	"touch",
	"ttl",
	"type",
	"xinfoconsumers",
	"xinfogroups",
	"xinfostream",
	"xinfohelp",
	"xlen",
	"xpending",
	"xrange",
	"xread",
	"xrevrange",
	"zcard",
	"zcount",
	"zdiff",
	"zinter",
	"zlexcount",
	"zmscore",
	"zrandmember",
	"zrange",
	"zrangebylex",
	"zrangebyscore",
	"zrank",
	"zrevrange",
	"zrevrangebylex",
	"zrevrangebyscore",
	"zrevrank",
	"zscan",
	"zscore",
	"zunion",
	"zintercard",

	"jsonget",
	"jsonmget",
	"jsonstrlen",
	"jsonarrindex",
	"jsonarrlen",
	"jsonobjkeys",
	"jsonobjlen",
	"jsontype",
	"jsonresp",

	"bfexists",
	"bfmexists",
	"bfscandump",
	"bfinfo",
	"cfexists",
	"cfcount",
	"cfscandump",
	"cfinfo",
	"cmsquery",
	"cmsinfo",
	"topkquery",
	"topkcount",
	"topklist",
	"topkinfo",
}
