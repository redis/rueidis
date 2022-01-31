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

type goStruct struct {
	Node      *node
	FullName  string
	BuildDef  buildDef
	NextNodes []*node
	Variadic  bool
}

type buildDef struct {
	MethodName string
	Command    []string
	Parameters []parameter
}

type parameter struct {
	Name string
	Type string
}

type command struct {
	Group     string     `json:"group"`
	Since     string     `json:"since"`
	Arguments []argument `json:"arguments"`
}

type argument struct {
	Name     interface{} `json:"name"`
	Type     interface{} `json:"type"`
	Command  string      `json:"command"`
	Enum     []string    `json:"enum"`
	Block    []argument  `json:"block"`
	Multiple bool        `json:"multiple"`
	Optional bool        `json:"optional"`
	Variadic bool        `json:"variadic"`
}

type node struct {
	Parent *node
	Child  *node
	Next   *node
	Cmd    command
	Arg    argument
	Root   bool
}

func (n *node) FindRoot() (root *node) {
	root = n
	for root.Parent != nil {
		root = root.Parent
	}
	return
}

//gocyclo:ignore
func (n *node) GoStructs() (out []goStruct) {
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

	// fix for TS.MRANGE, TS.MREVRANGE, TS.MGET, TS.QUERYINDEX
	if len(n.Arg.Enum) >= 2 && (n.Arg.Enum[0] == "l=v" && n.Arg.Enum[1] == "l!=v") {
		n.Arg.Enum = nil
		n.Arg.Name = "filter"
		n.Arg.Type = "string"
		if !strings.HasSuffix(fn, "Filter") {
			fn += "Filter"
		}
	}

	if len(n.Arg.Block) > 0 {
		panic("GoStructs should not be called on Block node")
	} else if len(n.Arg.Enum) > 0 {
		for _, e := range n.Arg.Enum {
			s := goStruct{
				Node:     n,
				FullName: fn,
				BuildDef: buildDef{
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
					s.BuildDef.Parameters = []parameter{{Name: lcFirst(name(cmds[1])), Type: "integer"}}
				case "*":
					// fix for FT.AGGREGATE
					if cmds[0] == "LOAD" {
						s.BuildDef.Command = append(s.BuildDef.Command, cmds...)
						s.BuildDef.Parameters = nil
					}
				case "label1":
					// fix for TS.MRANGE, TS.MREVRANGE, TS.MGET
					if cmds[0] == "SELECTED_LABELS" {
						s.FullName = strings.TrimRight(s.FullName, "Label1")
						s.BuildDef.MethodName = strings.TrimRight(s.BuildDef.MethodName, "Label1")
						s.BuildDef.Command = append(s.BuildDef.Command, cmds[0])
						s.BuildDef.Parameters = []parameter{{Name: "labels", Type: "[]string"}}
					}
				default:
					panic("unknown enum " + cmds[1])
				}
			}
			out = append(out, s)
		}
		return
	} else {
		s := goStruct{
			Node:     n,
			FullName: fn,
			BuildDef: buildDef{
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
			s.BuildDef.Parameters = []parameter{{Name: lcFirst(name(nm)), Type: n.Arg.Type.(string)}} // not change to go type here, change at render
		case []interface{}:
			for i, nn := range nm {
				s.BuildDef.Parameters = append(s.BuildDef.Parameters, parameter{Name: lcFirst(name(nn.(string))), Type: n.Arg.Type.([]interface{})[i].(string)})
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

func (n *node) Variadic() bool {
	return n.Arg.Multiple || n.Arg.Variadic
}

func (n *node) FullName() (out string) {
	if n.Parent != nil {
		return n.Parent.FullName() + n.Name()
	}
	return n.Name()
}

func (n *node) Name() (out string) {
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

func (n *node) NextNodes() []*node {
	var nodes []*node

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

func (n *node) Walk(fn func(node *node)) {
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
	var commands = map[string]command{}

	for _, p := range []string{
		"./commands.json",
		"./commands_sentinel.json",
		"./commands_json.json",
		"./commands_bloom.json",
		"./commands_search.json",
		"./commands_graph.json",
		"./commands_timeseries.json",
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
	commands["GEORADIUS_RO"] = command{Arguments: filterArgs(commands["GEORADIUS"].Arguments, "STORE")}
	commands["GEORADIUSBYMEMBER_RO"] = command{Arguments: filterArgs(commands["GEORADIUSBYMEMBER"].Arguments, "STORE")}

	var roots []string
	nodes := map[string]*node{}
	for k, cmd := range commands {
		root := &node{Cmd: cmd, Arg: argument{Name: k, Command: k, Type: "command"}, Root: true}
		root.Next = makeChildNodes(root, cmd.Arguments)
		roots = append(roots, k)
		nodes[k] = root
	}
	sort.Strings(roots)

	var structs = map[string]goStruct{}
	for _, name := range roots {
		n := nodes[name]
		n.Walk(func(n *node) {
			for _, s := range n.GoStructs() {
				if v, ok := structs[s.FullName]; ok {
					panic("struct conflict " + v.FullName)
				}
				structs[s.FullName] = s
			}
		})
	}

	gf, err := os.Create("../../internal/cmds/gen.go")
	if err != nil {
		panic(err)
	}
	defer gf.Close()
	tf, err := os.Create("../../internal/cmds/gen_test.go")
	if err != nil {
		panic(err)
	}
	defer tf.Close()

	generate(gf, structs)
	tests(tf, structs)
}

func tests(f io.Writer, structs map[string]goStruct) {
	var names []string
	for name := range structs {
		names = append(names, name)
	}
	sort.Strings(names)

	var pathes [][]goStruct
	for _, name := range names {
		s := structs[name]
		if !s.Node.Root {
			continue
		}
		pathes = makePath(s, nil, pathes)
	}

	fmt.Fprintf(f, "// Code generated DO NOT EDIT\n\npackage cmds\n\n")

	fmt.Fprintf(f, "import \"testing\"\n\n")

	fmt.Fprintf(f, "var s = NewBuilder(InitSlot)\n\n")

	for i, p := range pathes {
		if i%100 == 0 {
			fmt.Fprintf(f, "func TestCommand_%d(t *testing.T) {\n", i/100)
		}
		printPath(f, "s", p, "Build")
		if within(p[0], cacheableCMDs) {
			printPath(f, "s", p, "Cache")
		}
		if i%100 == 99 || i == len(pathes)-1 {
			fmt.Fprintf(f, "}\n\n")
		}
	}

}

func makePath(s goStruct, path []goStruct, pathes [][]goStruct) [][]goStruct {
	path = append(path, s)
	nexts := s.Node.NextNodes()
	if len(path) < 8 {
		for _, n := range nexts {
			if s.Node == n {
				continue
			}
			if n.Parent != nil && n.Parent.Child == n {
				continue
			}
			if n.Child != nil {
				n = n.Child
			}
			for _, ss := range n.GoStructs() {
				clone := make([]goStruct, len(path))
				copy(clone, path)
				pathes = makePath(ss, clone, pathes)
			}
		}
	}
	if allOptional(s.Node, nexts) {
		pathes = append(pathes, path)
	}
	return pathes
}

func testParams(defs []parameter) string {
	var params []string
	for _, param := range defs {
		switch toGoType(param.Type) {
		case "[]string":
			params = append(params, `[]string{"1"}`)
		case "string":
			params = append(params, `"1"`)
		case "int64", "float64":
			params = append(params, `1`)
		}
	}
	return strings.Join(params, ", ")
}

func printPath(f io.Writer, receiver string, path []goStruct, end string) {
	fmt.Fprintf(f, "\t%s.%s()", receiver, path[0].BuildDef.MethodName)
	for _, s := range path[1:] {
		fmt.Fprintf(f, ".%s(", s.BuildDef.MethodName)
		if len(s.BuildDef.Parameters) != 1 && s.Variadic {
			fmt.Fprintf(f, ").%s(", s.BuildDef.MethodName)
		}
		fmt.Fprintf(f, "%s)", testParams(s.BuildDef.Parameters))
		if s.Variadic {
			fmt.Fprintf(f, ".%s(%s)", s.BuildDef.MethodName, testParams(s.BuildDef.Parameters))
		}
	}
	fmt.Fprintf(f, ".%s()\n", end)
}

func generate(f io.Writer, structs map[string]goStruct) {
	var names []string
	for name := range structs {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Fprintf(f, "// Code generated DO NOT EDIT\n\npackage cmds\n\n")
	fmt.Fprintf(f, "import %q\n\n", "strconv")

	for _, name := range names {
		s := structs[name]

		fmt.Fprintf(f, "type %s Completed\n\n", s.FullName)

		if s.Node.Root {
			printRootBuilder(f, s)
		}

		for _, next := range s.NextNodes {
			if next.Child != nil {
				next = next.Child
			}
			for _, ss := range next.GoStructs() {
				printBuilder(f, s, ss)
			}
		}

		if allOptional(s.Node, s.NextNodes) {
			printFinalBuilder(f, s, "Build", "Completed")
			if within(s.Node.FindRoot().GoStructs()[0], cacheableCMDs) {
				printFinalBuilder(f, s, "Cache", "Cacheable")
			}
		}
	}

	checkAllUsed("noRetCMDs", noRetCMDs)
	checkAllUsed("blockingCMDs", blockingCMDs)
	checkAllUsed("cacheableCMDs", cacheableCMDs)
	checkAllUsed("readOnlyCMDs", readOnlyCMDs)
}

func checkAllUsed(name string, tags map[string]bool) {
	for t, v := range tags {
		if !v {
			panic(fmt.Sprintf("unsued tag %s in %s", t, name))
		}
	}
}

func allOptional(s *node, nodes []*node) bool {
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
	case "[]string": // TODO hack for TS.MRANGE, TS.MREVRANGE, TS.MGET
		return "[]string"
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

func printRootBuilder(w io.Writer, root goStruct) {
	fmt.Fprintf(w, "func (b *Builder) %s() (c %s) {\n", root.FullName, root.FullName)

	var appends []string
	for _, cmd := range root.BuildDef.Command {
		appends = append(appends, fmt.Sprintf(`"%s"`, cmd))
	}

	if tag := rootCf(root); tag != "" {
		fmt.Fprintf(w, "\tc = %s{cs: b.get(), ks: b.ks, cf: %s}\n", root.FullName, tag)
	} else {
		fmt.Fprintf(w, "\tc = %s{cs: b.get(), ks: b.ks}\n", root.FullName)
	}
	fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s)\n", strings.Join(appends, ", "))
	fmt.Fprintf(w, "\treturn c\n")
	fmt.Fprintf(w, "}\n\n")
}

func rootCf(root goStruct) (tag string) {
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

func printFinalBuilder(w io.Writer, parent goStruct, method, ss string) {
	fmt.Fprintf(w, "func (c %s) %s() %s {\n", parent.FullName, method, ss)
	fmt.Fprintf(w, "\treturn %s(c)\n", ss)
	fmt.Fprintf(w, "}\n\n")
}

//gocyclo:ignore
func printBuilder(w io.Writer, parent, next goStruct) {
	fmt.Fprintf(w, "func (c %s) %s(", parent.FullName, next.BuildDef.MethodName)
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
	fmt.Fprintf(w, ") %s {\n", next.FullName)

	if len(next.BuildDef.Parameters) == 1 && next.Variadic {
		if next.BuildDef.Parameters[0].Type == "key" {
			fmt.Fprintf(w, "\tif c.ks != NoSlot {\n")
			fmt.Fprintf(w, "\t\tfor _, k := range %s {\n", toGoName(next.BuildDef.Parameters[0].Name))
			fmt.Fprintf(w, "\t\t\tc.ks = check(c.ks, slot(k))\n")
			fmt.Fprintf(w, "\t\t}\n")
			fmt.Fprintf(w, "\t}\n")
		}
	} else {
		if len(next.BuildDef.Parameters) != 1 && next.Variadic && parent.FullName != next.FullName {
			// no parameter
		} else {
			for _, arg := range next.BuildDef.Parameters {
				if arg.Type == "key" {
					fmt.Fprintf(w, "\tif c.ks != NoSlot {\n")
					fmt.Fprintf(w, "\t\tc.ks = check(c.ks, slot(%s))\n", toGoName(arg.Name))
					fmt.Fprintf(w, "\t}\n")
				}
			}
		}
	}

	for _, cmd := range next.BuildDef.Command {
		if cmd == "BLOCK" {
			fmt.Fprintf(w, "\tc.cf = blockTag\n")
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
			if !(len(next.BuildDef.Parameters) != 1 && next.Variadic && parent.FullName == next.FullName) {
				appends = append(appends, fmt.Sprintf(`"%s"`, cmd))
			}
		}
	}

	if len(appends) == 0 && next.Variadic && len(next.BuildDef.Parameters) == 1 && toGoType(next.BuildDef.Parameters[0].Type) == "string" {
		appends = append(appends, toGoName(next.BuildDef.Parameters[0].Name)+"...")
		fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s)\n", strings.Join(appends, ", "))
	} else if len(next.BuildDef.Parameters) != 1 && next.Variadic && parent.FullName != next.FullName {
		// no parameter
		if len(appends) != 0 {
			fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s)\n", strings.Join(appends, ", "))
		}
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
			fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s)\n", strings.Join(appends, ", "))
		} else {
			if len(next.BuildDef.Parameters) == 1 && next.Variadic {
				if len(appends) != 0 {
					fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s)\n", strings.Join(appends, ", "))
				}
				if toGoType(next.BuildDef.Parameters[0].Type) == "string" {
					fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s...)\n", toGoName(next.BuildDef.Parameters[0].Name))
				} else {
					fmt.Fprintf(w, "\tfor _, n := range %s {\n", toGoName(next.BuildDef.Parameters[0].Name))
					switch toGoType(next.BuildDef.Parameters[0].Type) {
					case "float64":
						fmt.Fprintf(w, "\t\tc.cs.s = append(c.cs.s, strconv.FormatFloat(n, 'f', -1, 64))\n")
					case "int64":
						fmt.Fprintf(w, "\t\tc.cs.s = append(c.cs.s, strconv.FormatInt(n, 10))\n")
					default:
						panic("unexpected param type " + next.BuildDef.Parameters[0].Type)
					}
					fmt.Fprintf(w, "\t}\n")
				}
			} else {
				var follows []string
				for _, p := range next.BuildDef.Parameters {
					switch toGoType(p.Type) {
					case "float64":
						appends = append(appends, fmt.Sprintf("strconv.FormatFloat(%s, 'f', -1, 64)", toGoName(p.Name)))
					case "int64":
						appends = append(appends, fmt.Sprintf("strconv.FormatInt(%s, 10)", toGoName(p.Name)))
					case "string":
						appends = append(appends, toGoName(p.Name))
					case "[]string": // TODO hack for TS.MRANGE, TS.MREVRANGE, TS.MGET
						follows = append(follows, toGoName(p.Name)+"...")
					default:
						panic("unexpected param type " + next.BuildDef.Parameters[0].Type)
					}
				}
				fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s)\n", strings.Join(appends, ", "))
				for _, follow := range follows {
					fmt.Fprintf(w, "\tc.cs.s = append(c.cs.s, %s)\n", follow)
				}
			}
		}
	}

	if parent.FullName == next.FullName {
		fmt.Fprintf(w, "\treturn c\n")
	} else {
		fmt.Fprintf(w, "\treturn (%s)(c)\n", next.FullName)
	}
	fmt.Fprintf(w, "}\n\n")
}

func makeChildNodes(parent *node, args []argument) (first *node) {
	if len(args) == 0 {
		return nil
	}
	var nodes []*node
	for _, arg := range args {
		nodes = append(nodes, &node{Parent: parent, Arg: arg})
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

func filterArgs(args []argument, exclude string) (out []argument) {
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
		name += ucFirst(strings.ToLower(n))
	}
	return name
}

func ucFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func lcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func within(cmd goStruct, cmds map[string]bool) bool {
	n := strings.ToLower(cmd.FullName)
	if _, ok := cmds[n]; ok {
		cmds[n] = true
		return true
	}
	return false
}

var noRetCMDs = map[string]bool{
	"subscribe":    false,
	"psubscribe":   false,
	"unsubscribe":  false,
	"punsubscribe": false,
}

var blockingCMDs = map[string]bool{
	"blpop":       false,
	"brpop":       false,
	"brpoplpush":  false,
	"blmove":      false,
	"blmpop":      false,
	"bzpopmin":    false,
	"bzpopmax":    false,
	"bzmpop":      false,
	"clientpause": false,
	"migrate":     false,
	"wait":        false,
}

var cacheableCMDs = map[string]bool{
	"bitcount":            false,
	"bitfieldro":          false,
	"bitpos":              false,
	"expiretime":          false,
	"geodist":             false,
	"geohash":             false,
	"geopos":              false,
	"georadiusro":         false,
	"georadiusbymemberro": false,
	"geosearch":           false,
	"get":                 false,
	"getbit":              false,
	"getrange":            false,
	"hexists":             false,
	"hget":                false,
	"hgetall":             false,
	"hkeys":               false,
	"hlen":                false,
	"hmget":               false,
	"hstrlen":             false,
	"hvals":               false,
	"lindex":              false,
	"llen":                false,
	"lpos":                false,
	"lrange":              false,
	"pexpiretime":         false,
	"pttl":                false,
	"scard":               false,
	"sismember":           false,
	"smembers":            false,
	"smismember":          false,
	"sortro":              false,
	"strlen":              false,
	"ttl":                 false,
	"type":                false,
	"zcard":               false,
	"zcount":              false,
	"zlexcount":           false,
	"zmscore":             false,
	"zrange":              false,
	"zrangebylex":         false,
	"zrangebyscore":       false,
	"zrank":               false,
	"zrevrange":           false,
	"zrevrangebylex":      false,
	"zrevrangebyscore":    false,
	"zrevrank":            false,
	"zscore":              false,
	"jsonget":             false,
	"jsonstrlen":          false,
	"jsonarrindex":        false,
	"jsonarrlen":          false,
	"jsonobjkeys":         false,
	"jsonobjlen":          false,
	"jsontype":            false,
	"jsonresp":            false,
	"bfexists":            false,
	"bfinfo":              false,
	"cfexists":            false,
	"cfcount":             false,
	"cfinfo":              false,
	"cmsquery":            false,
	"cmsinfo":             false,
	"topkquery":           false,
	"topklist":            false,
	"topkinfo":            false,
}

var readOnlyCMDs = map[string]bool{
	"bitcount":            false,
	"bitfieldro":          false,
	"bitpos":              false,
	"dbsize":              false,
	"dump":                false,
	"evalro":              false,
	"evalsharo":           false,
	"exists":              false,
	"expiretime":          false,
	"geodist":             false,
	"geohash":             false,
	"geopos":              false,
	"georadiusro":         false,
	"georadiusbymemberro": false,
	"geosearch":           false,
	"get":                 false,
	"getbit":              false,
	"getrange":            false,
	"hexists":             false,
	"hget":                false,
	"hgetall":             false,
	"hkeys":               false,
	"hlen":                false,
	"hmget":               false,
	"hrandfield":          false,
	"hscan":               false,
	"hstrlen":             false,
	"hvals":               false,
	"keys":                false,
	"lindex":              false,
	"llen":                false,
	"lolwut":              false,
	"lpos":                false,
	"lrange":              false,
	"memorydoctor":        false,
	"memorystats":         false,
	"memoryusage":         false,
	"memorymallocstats":   false,
	"mget":                false,
	"objectencoding":      false,
	"objectfreq":          false,
	"objecthelp":          false,
	"objectidletime":      false,
	"objectrefcount":      false,
	"pexpiretime":         false,
	"pfcount":             false,
	"pttl":                false,
	"pubsubchannels":      false,
	"pubsubnumpat":        false,
	"pubsubnumsub":        false,
	"pubsubhelp":          false,
	"randomkey":           false,
	"scan":                false,
	"scard":               false,
	"sdiff":               false,
	"sinter":              false,
	"sintercard":          false,
	"sismember":           false,
	"slowlogget":          false,
	"slowloglen":          false,
	"slowloghelp":         false,
	"smembers":            false,
	"smismember":          false,
	"sortro":              false,
	"srandmember":         false,
	"sscan":               false,
	"lcs":                 false,
	"strlen":              false,
	"sunion":              false,
	"touch":               false,
	"ttl":                 false,
	"type":                false,
	"xinfoconsumers":      false,
	"xinfogroups":         false,
	"xinfostream":         false,
	"xinfohelp":           false,
	"xlen":                false,
	"xpending":            false,
	"xrange":              false,
	"xread":               false,
	"xrevrange":           false,
	"zcard":               false,
	"zcount":              false,
	"zdiff":               false,
	"zinter":              false,
	"zlexcount":           false,
	"zmscore":             false,
	"zrandmember":         false,
	"zrange":              false,
	"zrangebylex":         false,
	"zrangebyscore":       false,
	"zrank":               false,
	"zrevrange":           false,
	"zrevrangebylex":      false,
	"zrevrangebyscore":    false,
	"zrevrank":            false,
	"zscan":               false,
	"zscore":              false,
	"zunion":              false,
	"zintercard":          false,
	"jsonget":             false,
	"jsonmget":            false,
	"jsonstrlen":          false,
	"jsonarrindex":        false,
	"jsonarrlen":          false,
	"jsonobjkeys":         false,
	"jsonobjlen":          false,
	"jsontype":            false,
	"jsonresp":            false,
	"bfexists":            false,
	"bfmexists":           false,
	"bfscandump":          false,
	"bfinfo":              false,
	"cfexists":            false,
	"cfcount":             false,
	"cfscandump":          false,
	"cfinfo":              false,
	"cmsquery":            false,
	"cmsinfo":             false,
	"topkquery":           false,
	"topklist":            false,
	"topkinfo":            false,
	"tsrange":             false,
	"tsrevrange":          false,
	"tsget":               false,
	"tsinfo":              false,
	"tsqueryindex":        false,
	"graphroquery":        false,
	"graphexplain":        false,
	"graphslowlog":        false,
	"graphconfigget":      false,
	"graphlist":           false,
}
