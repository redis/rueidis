package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

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

func (a *Argument) FullName() (nn string) {
	var tokens []string
	if a.Command != "" && a.Command != "LOAD *" {
		return UcFirst(name(a.Command))
	} else {
		switch n := a.Name.(type) {
		case string:
			return UcFirst(name(n))
		case []interface{}:
			for _, nn := range n {
				tokens = append(tokens, strings.Split(nn.(string), " ")...)
			}
		}
	}
	for _, t := range tokens {
		nn += UcFirst(name(t))
	}
	if nn == "" {
		panic("no name")
	}
	return
}

type CmdNode struct {
	Argument
	StructName string
	Children   []string
	Root       *CmdNode
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func main() {
	fmt.Printf("// Code generated DO NOT EDIT\n\npackage cmds\n\n")
	fmt.Printf("import %q\n\n", "strconv")

	generate("")
	generate("S")
}

func generate(prefix string) {
	var nodes = map[string]*CmdNode{}
	var commands = map[string]struct {
		Group     string     `json:"group"`
		Arguments []Argument `json:"arguments"`
	}{}

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
	commands["GEORADIUS_RO"] = struct {
		Group     string     `json:"group"`
		Arguments []Argument `json:"arguments"`
	}{
		Arguments: filterArgs(commands["GEORADIUS"].Arguments, "STORE"),
	}
	commands["GEORADIUSBYMEMBER_RO"] = struct {
		Group     string     `json:"group"`
		Arguments []Argument `json:"arguments"`
	}{
		Arguments: filterArgs(commands["GEORADIUSBYMEMBER"].Arguments, "STORE"),
	}

	for k, info := range commands {
		sn := name(k)
		cmd := &CmdNode{Root: nil, StructName: sn, Argument: Argument{Command: k}}
		if _, ok := nodes[cmd.StructName]; ok {
			panic("StructName conflict " + cmd.StructName)
		}
		nodes[cmd.StructName] = cmd
		for _, arg := range info.Arguments {
			cmd.Children = append(cmd.Children, argIDs(sn, arg)...)
			if !arg.Optional {
				break
			}
		}
		for i, arg := range info.Arguments {
			node(nodes, cmd, sn, arg, info.Arguments[i+1:], "", nil)
		}
	}
	var keys []string
	for k := range nodes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		node := nodes[k]
		for _, c := range node.Children {
			if _, ok := nodes[c]; !ok {
				panic("missing node " + c)
			}
		}
	}

	for _, k := range keys {
		node := nodes[k]

		if node.StructName == "append" {
			panic("reserved word")
		}

		fmt.Printf("type %s %sCompleted\n\n", prefix+node.StructName, prefix)

		if node.Multiple || node.Variadic {
			// put node itself (no command) as child
			nocmd := &CmdNode{Root: node, StructName: node.StructName + "_nocmd", Argument: node.Argument}
			if nocmd.Argument.Command != "" {
				nocmd.Argument.Name = strings.ToLower(nocmd.Argument.Command)
				nocmd.Argument.Command = ""
			}
			nodes[nocmd.StructName] = nocmd
			node.Children = append(node.Children, nocmd.StructName)
		}

		for _, c := range node.Children {
			child := nodes[c]
			fmt.Printf("func (c %s) %s(", prefix+node.StructName, child.Argument.FullName())
			// func args
			var args [][2]string
			switch nn := child.Argument.Name.(type) {
			case string:
				if child.Argument.Type != nil {
					args = append(args, [2]string{nn, child.Argument.Type.(string)})
				}
			case []interface{}:
				for i, n := range nn {
					args = append(args, [2]string{n.(string), child.Argument.Type.([]interface{})[i].(string)})
				}
			case nil:
			default:
				panic("unknown name type")
			}
			for i, arg := range args {
				arg[0] = name(arg[0])
				if arg[0] == "type" {
					arg[0] = "typ"
				}
				switch arg[1] {
				case "key":
				case "string", "pattern", "type":
					arg[1] = "string"
				case "integer", "posix time":
					arg[1] = "int64"
				case "double":
					arg[1] = "float64"
				default:
					panic("unknown type " + arg[1])
				}
				args[i] = arg
			}
			if len(args) > 1 && (child.Multiple || child.Variadic) && !strings.HasSuffix(child.StructName, "_nocmd") {
				args = nil
			}
			if len(args) == 1 && (child.Multiple || child.Variadic) {
				if args[0][1] == "key" {
					fmt.Printf("%s ...string", args[0][0])
				} else {
					fmt.Printf("%s ...%s", args[0][0], args[0][1])
				}
			} else {
				for i, arg := range args {
					if arg[1] == "key" {
						fmt.Printf("%s string", arg[0])
					} else {
						fmt.Printf("%s %s", arg[0], arg[1])
					}
					if i != len(args)-1 {
						fmt.Printf(", ")
					}
				}
			}
			fmt.Printf(") %s {\n", prefix+strings.TrimSuffix(child.StructName, "_nocmd"))
			// func body

			if prefix == "S" {
				if len(args) == 1 && (child.Multiple || child.Variadic) {
					if args[0][1] == "key" {
						fmt.Printf("\tfor _, k := range %s {\n", args[0][0])
						fmt.Printf("\t\tc.ks = checkSlot(c.ks, slot(k))\n")
						fmt.Printf("\t}\n")
					}
				} else {
					for _, arg := range args {
						if arg[1] == "key" {
							fmt.Printf("\tc.ks = checkSlot(c.ks, slot(%s))\n", arg[0])
						}
					}
				}
			}

			if child.Command == "BLOCK" {
				fmt.Printf("\tc.cf = blockTag\n")
			}

			var appends []string
			if len(child.Command) > 0 {
				for _, c := range strings.Split(child.Command, " ") {
					if c == "empty" {
						appends = append(appends, "\"\"")
					} else {
						appends = append(appends, fmt.Sprintf("%q", c))
					}
				}
			}
			if len(args) == 1 && (child.Multiple || child.Variadic) {
				// nothing
			} else {
				for i, arg := range args {
					if args[i][1] == "int64" {
						appends = append(appends, fmt.Sprintf("strconv.FormatInt(%s, 10)", arg[0]))
					} else if args[i][1] == "float64" {
						appends = append(appends, fmt.Sprintf("strconv.FormatFloat(%s, 'f', -1, 64)", arg[0]))
					} else {
						appends = append(appends, arg[0])
					}
				}
			}
			if len(args) == 1 && (child.Multiple || child.Variadic) {
				if len(appends) == 0 && (args[0][1] == "string" || args[0][1] == "key") {
					// one line append
					fmt.Printf("\treturn %s{cs: append(c.cs, %s...), cf: c.cf, ks: c.ks}\n", prefix+strings.TrimSuffix(child.StructName, "_nocmd"), args[0][0])
				} else {
					if len(appends) != 0 {
						fmt.Printf("\tc.cs = append(c.cs, ")
						for i, ap := range appends {
							fmt.Printf(ap)
							if i != len(appends)-1 {
								fmt.Printf(", ")
							}
						}
						fmt.Printf(")\n")
					}
					if args[0][1] == "string" || args[0][1] == "key" {
						fmt.Printf("\treturn %s{cs: append(c.cs, %s...), cf: c.cf, ks: c.ks}\n", prefix+strings.TrimSuffix(child.StructName, "_nocmd"), args[0][0])
					} else {
						fmt.Printf("\tfor _, n := range %s {\n", args[0][0])
						if args[0][1] == "float64" {
							fmt.Printf("\t\tc.cs = append(c.cs, strconv.FormatFloat(n, 'f', -1, 64))\n")
						} else {
							fmt.Printf("\t\tc.cs = append(c.cs, strconv.FormatInt(n, 10))\n")
						}
						fmt.Printf("\t}\n")
						fmt.Printf("\treturn %s{cs: c.cs, cf: c.cf, ks: c.ks}\n", prefix+strings.TrimSuffix(child.StructName, "_nocmd"))
					}
				}
			} else {
				// one line append
				if len(appends) > 0 {
					fmt.Printf("\treturn %s{cs: append(c.cs, ", prefix+strings.TrimSuffix(child.StructName, "_nocmd"))
					for i, ap := range appends {
						fmt.Printf(ap)
						if i != len(appends)-1 {
							fmt.Printf(", ")
						}
					}
					fmt.Printf("), cf: c.cf, ks: c.ks}\n")
				} else {
					fmt.Printf("\treturn %s{cs: c.cs, cf: c.cf, ks: c.ks}\n", prefix+strings.TrimSuffix(child.StructName, "_nocmd"))
				}
			}
			fmt.Printf("}\n\n")
		}

		if allOptional(nodes, node.Children) {
			fmt.Printf("func (c %s) Build() %sCompleted {\n", prefix+node.StructName, prefix)
			fmt.Printf("\treturn %sCompleted(c)\n", prefix)
			fmt.Printf("}\n\n")

			if node.Root != nil && within(node.Root, cacheableCMDs) {
				fmt.Printf("func (c %s) Cache() %sCacheable {\n", prefix+node.StructName, prefix)
				fmt.Printf("\treturn %sCacheable(c)\n", prefix)
				fmt.Printf("}\n\n")
			}
		}

		if node.Root == nil {
			var appends []string
			for _, c := range strings.Split(node.Command, " ") {
				if c == "empty" {
					appends = append(appends, "\"\"")
				} else {
					appends = append(appends, fmt.Sprintf("%q", c))
				}
			}
			fmt.Printf("func (b *%sBuilder) %s() (c %s) {\n", prefix, node.Argument.FullName(), prefix+node.StructName)

			fmt.Printf("\tc.cs = append(b.get(), ")
			for i, ap := range appends {
				fmt.Printf(ap)
				if i != len(appends)-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(")\n")

			if within(node, blockingCMDs) {
				fmt.Printf("\tc.cf = blockTag\n")
			}

			if within(node, noRetCMDs) {
				fmt.Printf("\tc.cf = noRetTag\n")
			}

			if within(node, readOnlyCMDs) {
				fmt.Printf("\tc.cf = readonly\n")
			}

			if prefix == "S" {
				fmt.Printf("\tc.ks = InitSlot\n")
			}

			fmt.Printf("\treturn\n")
			fmt.Printf("}\n\n")
		}
	}
}

func node(nodes map[string]*CmdNode, root *CmdNode, prefix string, arg Argument, args []Argument, parent string, parentArgs []Argument) {
	if len(arg.Enum) > 0 && arg.Type == nil {
		arg.Type = "enum"
	}
	if len(arg.Block) > 0 && arg.Type == nil {
		arg.Type = "block"
	}
	if len(arg.Enum) > 0 && arg.Type != "enum" {
		panic("wrong input")
	}
	if len(arg.Block) > 0 && arg.Type != "block" {
		panic("wrong input")
	}
	// fix for XGROUP
	if len(arg.Enum) == 2 && (arg.Enum[0] == "ID" || arg.Enum[0] == "id") && arg.Enum[1] == "$" {
		arg.Name = "id"
		arg.Type = "string"
		arg.Enum = nil
	}
	// fix for XADD
	if len(arg.Enum) == 2 && arg.Enum[0] == "*" && arg.Enum[1] == "ID" {
		arg.Name = "id"
		arg.Type = "string"
		arg.Enum = nil
	}
	// fix for FT.AGGREGATE
	if len(arg.Enum) == 1 && arg.Enum[0] == "LOAD *" {
		arg.Type = nil
		arg.Command = "LOAD *"
		arg.Enum = nil
	}
	switch arg.Type {
	case "enum":
		for _, e := range arg.Enum {
			cmd := &CmdNode{Root: root}
			switch e {
			case "~":
				cmd.StructName = prefix + UcFirst(name(arg.Name)) + "Almost"
				cmd.Argument = Argument{Command: e, Optional: arg.Optional, Multiple: arg.Multiple, Variadic: arg.Variadic}
			case "=":
				cmd.StructName = prefix + UcFirst(name(arg.Name)) + "Exact"
				cmd.Argument = Argument{Command: e, Optional: arg.Optional, Multiple: arg.Multiple, Variadic: arg.Variadic}
			case "\"\"":
				cmd.StructName = prefix + UcFirst(name(arg.Name)) + "Empty"
				cmd.Argument = Argument{Command: e, Optional: arg.Optional, Multiple: arg.Multiple, Variadic: arg.Variadic}
			default:
				if len(e) == 1 && e != "m" {
					panic("unknown e: " + e)
				}
				cmd.StructName = prefix + UcFirst(name(arg.Name)) + UcFirst(strings.ToLower(strings.Split(e, " ")[0]))
				cmd.Argument = Argument{Command: strings.Split(e, " ")[0], Optional: arg.Optional, Multiple: arg.Multiple, Variadic: arg.Variadic}
				if tokens := strings.Split(e, " "); len(tokens) > 1 {
					switch tokens[1] {
					case "seconds":
						cmd.Argument.Name = "seconds"
						cmd.Argument.Type = "integer"
					case "milliseconds":
						cmd.Argument.Name = "milliseconds"
						cmd.Argument.Type = "integer"
					case "timestamp":
						cmd.Argument.Name = "timestamp"
						cmd.Argument.Type = "integer"
					case "milliseconds-timestamp":
						cmd.Argument.Name = "millisecondsTimestamp"
						cmd.Argument.Type = "integer"
					default:
						panic("unknown enum " + tokens[1])
					}
				}
			}
			if _, ok := nodes[cmd.StructName]; ok {
				panic("StructName conflict " + cmd.StructName)
			}
			nodes[cmd.StructName] = cmd
			breaked := false
			for _, a := range args {
				cmd.Children = append(cmd.Children, argIDs(prefix, a)...)
				if !a.Optional {
					breaked = true
					break
				}
			}
			if !breaked {
				for _, a := range parentArgs {
					cmd.Children = append(cmd.Children, argIDs(parent, a)...)
					if !a.Optional {
						break
					}
				}
			}
		}
	case "block":
		if arg.Optional {
			arg.Block[0].Optional = true
		}
		for i, a := range arg.Block {
			node(nodes, root, prefix+UcFirst(name(arg.Name)), a, arg.Block[i+1:], prefix, args)
		}
	default:
		sn := UcFirst(prefix) + UcFirst(arg.FullName())
		// fix for FtCreatePrefix
		if sn == "FtCreatePrefix" && arg.Name == "count" {
			sn = "FtCreatePrefixCount"
		}
		cmd := &CmdNode{Root: root, StructName: sn, Argument: arg}
		if _, ok := nodes[cmd.StructName]; ok {
			panic("StructName conflict " + cmd.StructName)
		}
		nodes[cmd.StructName] = cmd
		breaked := false
		for _, a := range args {
			cmd.Children = append(cmd.Children, argIDs(prefix, a)...)
			if !a.Optional {
				breaked = true
				break
			}
		}
		if !breaked {
			for _, a := range parentArgs {
				cmd.Children = append(cmd.Children, argIDs(parent, a)...)
				if !a.Optional {
					break
				}
			}
		}
	}
}

func argIDs(prefix string, arg Argument) (ids []string) {
	if len(arg.Enum) > 0 && arg.Type == nil {
		arg.Type = "enum"
	}
	if len(arg.Block) > 0 && arg.Type == nil {
		arg.Type = "block"
	}
	if len(arg.Enum) > 0 && arg.Type != "enum" {
		panic("wrong input")
	}
	if len(arg.Block) > 0 && arg.Type != "block" {
		panic("wrong input")
	}
	// fix for XGROUP
	if len(arg.Enum) == 2 && (arg.Enum[0] == "ID" || arg.Enum[0] == "id") && arg.Enum[1] == "$" {
		arg.Name = "id"
		arg.Type = "string"
		arg.Enum = nil
	}
	// fix for XADD
	if len(arg.Enum) == 2 && arg.Enum[0] == "*" && arg.Enum[1] == "ID" {
		arg.Name = "id"
		arg.Type = "string"
		arg.Enum = nil
	}
	// fix for FT.AGGREGATE
	if len(arg.Enum) == 1 && arg.Enum[0] == "LOAD *" {
		arg.Type = "command"
		arg.Command = "LOAD *"
		arg.Enum = nil
	}
	switch arg.Type {
	case "enum":
		for _, e := range arg.Enum {
			switch e {
			case "~":
				ids = append(ids, prefix+UcFirst(name(arg.Name))+"Almost")
			case "=":
				ids = append(ids, prefix+UcFirst(name(arg.Name))+"Exact")
			case "\"\"":
				ids = append(ids, prefix+UcFirst(name(arg.Name))+"Empty")
			default:
				if len(e) == 1 && e != "m" {
					panic("unknown e: " + e)
				}
				ids = append(ids, prefix+UcFirst(name(arg.Name))+UcFirst(strings.ToLower(strings.Split(e, " ")[0])))
			}
		}
	case "block":
		for _, a := range arg.Block {
			ids = append(ids, argIDs(prefix+UcFirst(name(arg.Name)), a)...)
			if !a.Optional {
				break
			}
		}
	default:
		return []string{UcFirst(prefix) + UcFirst(arg.FullName())}
	}
	return ids
}

func name(n interface{}) (name string) {
	if s, ok := n.(string); ok {
		switch s {
		case "~":
			return "almost"
		case "=":
			return "exact"
		case "\"\"":
			return "empty"
		}

		for _, n := range strings.Split(strings.NewReplacer("-", " ", "_", " ", ":", " ", "/", " ", ".", " ").Replace(s), " ") {
			name += UcFirst(strings.ToLower(n))
		}
		return name
	}
	return ""
}

func allOptional(nodes map[string]*CmdNode, children []string) bool {
	for _, c := range children {
		if strings.HasSuffix(c, "_nocmd") {
			continue
		}
		if ch, ok := nodes[c]; !ok {
			panic("missing child " + c)
		} else if !ch.Optional {
			return false
		}
	}
	return true
}

func within(cmd *CmdNode, cmds []string) bool {
	n := strings.ToLower(cmd.StructName)
	for _, v := range cmds {
		if v == n {
			return true
		}
	}
	return false
}

func filterArgs(args []Argument, exclude string) (out []Argument) {
	for _, a := range args {
		bs, _ := json.Marshal(a)
		cp := Argument{}
		json.Unmarshal(bs, &cp)
		if cp.Command != exclude {
			out = append(out, cp)
		}
	}
	return out
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
	"stralgo",
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
