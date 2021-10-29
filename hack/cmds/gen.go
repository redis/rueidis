package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

//go:embed commands.json
var raw []byte
var commands = map[string]struct {
	Group     string     `json:"group"`
	Arguments []Argument `json:"arguments"`
}{}

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
	if a.Command != "" {
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

var nodes = map[string]*CmdNode{}

func main() {
	if err := json.Unmarshal(raw, &commands); err != nil {
		panic(err)
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
			node(cmd, sn, arg, info.Arguments[i+1:], "", nil)
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
	fmt.Printf("package cmds\n\n")

	fmt.Printf("import %q\n\n", "strconv")
	for _, k := range keys {
		node := nodes[k]

		if node.StructName == "append" {
			panic("reserved word")
		}

		fmt.Printf("type %s struct {\n", node.StructName)
		fmt.Printf("\tcs []string\n")
		fmt.Printf("\tcf uint32\n")
		fmt.Printf("}\n\n")

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
			fmt.Printf("func (c %s) %s(", node.StructName, child.Argument.FullName())
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
				case "key", "string", "pattern", "type":
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
				fmt.Printf("%s ...%s", args[0][0], args[0][1])
			} else {
				for i, arg := range args {
					fmt.Printf("%s %s", arg[0], arg[1])
					if i != len(args)-1 {
						fmt.Printf(", ")
					}
				}
			}
			fmt.Printf(") %s {\n", strings.TrimSuffix(child.StructName, "_nocmd"))
			// func body
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
				if len(appends) == 0 && args[0][1] == "string" {
					// one line append
					fmt.Printf("\treturn %s{cf: c.cf, cs: append(c.cs, %s...)}\n", strings.TrimSuffix(child.StructName, "_nocmd"), args[0][0])
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
					if args[0][1] == "string" {
						fmt.Printf("\treturn %s{cf: c.cf, cs: append(c.cs, %s...)}\n", strings.TrimSuffix(child.StructName, "_nocmd"), args[0][0])
					} else {
						fmt.Printf("\tfor _, n := range %s {\n", args[0][0])
						fmt.Printf("\t\tc.cs = append(c.cs, strconv.FormatInt(n, 10))\n")
						fmt.Printf("\t}\n")
						fmt.Printf("\treturn %s{cf: c.cf, cs: c.cs}\n", strings.TrimSuffix(child.StructName, "_nocmd"))
					}
				}
			} else {
				// one line append
				fmt.Printf("\treturn %s{cf: c.cf, cs: append(c.cs, ", strings.TrimSuffix(child.StructName, "_nocmd"))
				for i, ap := range appends {
					fmt.Printf(ap)
					if i != len(appends)-1 {
						fmt.Printf(", ")
					}
				}
				fmt.Printf(")}\n")
			}
			fmt.Printf("}\n\n")
		}

		if allOptional(node.Children) {
			fmt.Printf("func (c %s) Build() Completed {\n", node.StructName)
			fmt.Printf("\treturn Completed(c)\n")
			fmt.Printf("}\n\n")
		}

		if node.Root != nil && supportCaching(node.Root) {
			fmt.Printf("func (c %s) Cache() Cacheable {\n", node.StructName)
			fmt.Printf("\treturn Cacheable(c)\n")
			fmt.Printf("}\n\n")
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
			fmt.Printf("func (b *Builder) %s() (c %s) {\n", node.Argument.FullName(), node.StructName)

			if isBlocking(node) {
				fmt.Printf("\tc.cf = blockTag\n")
			}

			fmt.Printf("\tc.cs = append(b.get(), ")
			for i, ap := range appends {
				fmt.Printf(ap)
				if i != len(appends)-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(")\n")
			fmt.Printf("\treturn\n")
			fmt.Printf("}\n\n")
		}
	}
}

func node(root *CmdNode, prefix string, arg Argument, args []Argument, parent string, parentArgs []Argument) {
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
			case "*":
				cmd.StructName = prefix + UcFirst(name(arg.Name)) + "Wildcard"
				cmd.Argument = Argument{Command: e, Optional: arg.Optional, Multiple: arg.Multiple, Variadic: arg.Variadic}
			case "$":
				cmd.StructName = prefix + UcFirst(name(arg.Name)) + "LastID"
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
		for i, a := range arg.Block {
			node(root, prefix+UcFirst(name(arg.Name)), a, arg.Block[i+1:], prefix, args)
		}
	default:
		sn := UcFirst(prefix) + UcFirst(arg.FullName())
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
	switch arg.Type {
	case "enum":
		for _, e := range arg.Enum {
			switch e {
			case "~":
				ids = append(ids, prefix+UcFirst(name(arg.Name))+"Almost")
			case "=":
				ids = append(ids, prefix+UcFirst(name(arg.Name))+"Exact")
			case "*":
				ids = append(ids, prefix+UcFirst(name(arg.Name))+"Wildcard")
			case "$":
				ids = append(ids, prefix+UcFirst(name(arg.Name))+"LastID")
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
		case "*":
			return "wildcard"
		case "$":
			return "lastid"
		case "\"\"":
			return "empty"
		}

		for _, n := range strings.Split(strings.NewReplacer("-", " ", "_", " ", ":", " ", "/", " ").Replace(s), " ") {
			name += UcFirst(strings.ToLower(n))
		}
		return name
	}
	return ""
}

func allOptional(children []string) bool {
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

func isBlocking(cmd *CmdNode) bool {
	n := strings.ToLower(cmd.StructName)
	for _, v := range blockingCMDs {
		if v == n {
			return true
		}
	}
	return false
}

func supportCaching(cmd *CmdNode) bool {
	n := strings.ToLower(cmd.StructName)
	for _, v := range cacheableCMDs {
		if v == n {
			return true
		}
	}
	return false
}

var blockingCMDs = []string{
	"blpop",
	"brpop",
	"brpoplpush",
	"blmove",
	"blmpop",
	"bzpopmin",
	"bzpopmax",
	"clientpause",
	"migrate",
	"wait",
}

var cacheableCMDs = []string{
	"bitcount",
	"bitfieldro",
	"bitpos",
	"geodist",
	"geohash",
	"geopos",
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
	"pttl",
	"scard",
	"sismember",
	"smembers",
	"smismember",
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
}
