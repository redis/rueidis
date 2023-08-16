// Code generated DO NOT EDIT

package cmds

import "strconv"

type Tfcall Completed

func (b Builder) Tfcall() (c Tfcall) {
	c = Tfcall{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "TFCALL")
	return c
}

func (c Tfcall) LibraryFunction(libraryFunction string) TfcallLibraryFunction {
	c.cs.s = append(c.cs.s, libraryFunction)
	return (TfcallLibraryFunction)(c)
}

type TfcallArg Completed

func (c TfcallArg) Arg(arg ...string) TfcallArg {
	c.cs.s = append(c.cs.s, arg...)
	return c
}

func (c TfcallArg) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfcallKey Completed

func (c TfcallKey) Key(key ...string) TfcallKey {
	if c.ks&NoSlot == NoSlot {
		for _, k := range key {
			c.ks = NoSlot | slot(k)
			break
		}
	} else {
		for _, k := range key {
			c.ks = check(c.ks, slot(k))
		}
	}
	c.cs.s = append(c.cs.s, key...)
	return c
}

func (c TfcallKey) Arg(arg ...string) TfcallArg {
	c.cs.s = append(c.cs.s, arg...)
	return (TfcallArg)(c)
}

func (c TfcallKey) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfcallLibraryFunction Completed

func (c TfcallLibraryFunction) Numkeys(numkeys int64) TfcallNumkeys {
	c.cs.s = append(c.cs.s, strconv.FormatInt(numkeys, 10))
	return (TfcallNumkeys)(c)
}

type TfcallNumkeys Completed

func (c TfcallNumkeys) Key(key ...string) TfcallKey {
	if c.ks&NoSlot == NoSlot {
		for _, k := range key {
			c.ks = NoSlot | slot(k)
			break
		}
	} else {
		for _, k := range key {
			c.ks = check(c.ks, slot(k))
		}
	}
	c.cs.s = append(c.cs.s, key...)
	return (TfcallKey)(c)
}

func (c TfcallNumkeys) Arg(arg ...string) TfcallArg {
	c.cs.s = append(c.cs.s, arg...)
	return (TfcallArg)(c)
}

func (c TfcallNumkeys) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type Tfcallasync Completed

func (b Builder) Tfcallasync() (c Tfcallasync) {
	c = Tfcallasync{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "TFCALLASYNC")
	return c
}

func (c Tfcallasync) LibraryFunction(libraryFunction string) TfcallasyncLibraryFunction {
	c.cs.s = append(c.cs.s, libraryFunction)
	return (TfcallasyncLibraryFunction)(c)
}

type TfcallasyncArg Completed

func (c TfcallasyncArg) Arg(arg ...string) TfcallasyncArg {
	c.cs.s = append(c.cs.s, arg...)
	return c
}

func (c TfcallasyncArg) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfcallasyncKey Completed

func (c TfcallasyncKey) Key(key ...string) TfcallasyncKey {
	if c.ks&NoSlot == NoSlot {
		for _, k := range key {
			c.ks = NoSlot | slot(k)
			break
		}
	} else {
		for _, k := range key {
			c.ks = check(c.ks, slot(k))
		}
	}
	c.cs.s = append(c.cs.s, key...)
	return c
}

func (c TfcallasyncKey) Arg(arg ...string) TfcallasyncArg {
	c.cs.s = append(c.cs.s, arg...)
	return (TfcallasyncArg)(c)
}

func (c TfcallasyncKey) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfcallasyncLibraryFunction Completed

func (c TfcallasyncLibraryFunction) Numkeys(numkeys int64) TfcallasyncNumkeys {
	c.cs.s = append(c.cs.s, strconv.FormatInt(numkeys, 10))
	return (TfcallasyncNumkeys)(c)
}

type TfcallasyncNumkeys Completed

func (c TfcallasyncNumkeys) Key(key ...string) TfcallasyncKey {
	if c.ks&NoSlot == NoSlot {
		for _, k := range key {
			c.ks = NoSlot | slot(k)
			break
		}
	} else {
		for _, k := range key {
			c.ks = check(c.ks, slot(k))
		}
	}
	c.cs.s = append(c.cs.s, key...)
	return (TfcallasyncKey)(c)
}

func (c TfcallasyncNumkeys) Arg(arg ...string) TfcallasyncArg {
	c.cs.s = append(c.cs.s, arg...)
	return (TfcallasyncArg)(c)
}

func (c TfcallasyncNumkeys) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionDelete Completed

func (b Builder) TfunctionDelete() (c TfunctionDelete) {
	c = TfunctionDelete{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "TFUNCTION", "DELETE")
	return c
}

func (c TfunctionDelete) LibraryName(libraryName string) TfunctionDeleteLibraryName {
	c.cs.s = append(c.cs.s, libraryName)
	return (TfunctionDeleteLibraryName)(c)
}

type TfunctionDeleteLibraryName Completed

func (c TfunctionDeleteLibraryName) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionList Completed

func (b Builder) TfunctionList() (c TfunctionList) {
	c = TfunctionList{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "TFUNCTION", "LIST")
	return c
}

func (c TfunctionList) LibraryName(libraryName string) TfunctionListLibraryName {
	c.cs.s = append(c.cs.s, libraryName)
	return (TfunctionListLibraryName)(c)
}

func (c TfunctionList) Withcode() TfunctionListWithcode {
	c.cs.s = append(c.cs.s, "WITHCODE")
	return (TfunctionListWithcode)(c)
}

func (c TfunctionList) Verbose() TfunctionListVerbose {
	c.cs.s = append(c.cs.s, "VERBOSE")
	return (TfunctionListVerbose)(c)
}

func (c TfunctionList) V() TfunctionListV {
	c.cs.s = append(c.cs.s, "V")
	return (TfunctionListV)(c)
}

func (c TfunctionList) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionListLibraryName Completed

func (c TfunctionListLibraryName) Withcode() TfunctionListWithcode {
	c.cs.s = append(c.cs.s, "WITHCODE")
	return (TfunctionListWithcode)(c)
}

func (c TfunctionListLibraryName) Verbose() TfunctionListVerbose {
	c.cs.s = append(c.cs.s, "VERBOSE")
	return (TfunctionListVerbose)(c)
}

func (c TfunctionListLibraryName) V() TfunctionListV {
	c.cs.s = append(c.cs.s, "V")
	return (TfunctionListV)(c)
}

func (c TfunctionListLibraryName) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionListV Completed

func (c TfunctionListV) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionListVerbose Completed

func (c TfunctionListVerbose) V() TfunctionListV {
	c.cs.s = append(c.cs.s, "V")
	return (TfunctionListV)(c)
}

func (c TfunctionListVerbose) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionListWithcode Completed

func (c TfunctionListWithcode) Verbose() TfunctionListVerbose {
	c.cs.s = append(c.cs.s, "VERBOSE")
	return (TfunctionListVerbose)(c)
}

func (c TfunctionListWithcode) V() TfunctionListV {
	c.cs.s = append(c.cs.s, "V")
	return (TfunctionListV)(c)
}

func (c TfunctionListWithcode) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionLoad Completed

func (b Builder) TfunctionLoad() (c TfunctionLoad) {
	c = TfunctionLoad{cs: get(), ks: b.ks}
	c.cs.s = append(c.cs.s, "TFUNCTION", "LOAD")
	return c
}

func (c TfunctionLoad) Replace() TfunctionLoadReplace {
	c.cs.s = append(c.cs.s, "REPLACE")
	return (TfunctionLoadReplace)(c)
}

func (c TfunctionLoad) Config(config string) TfunctionLoadConfig {
	c.cs.s = append(c.cs.s, config)
	return (TfunctionLoadConfig)(c)
}

func (c TfunctionLoad) LibraryCode(libraryCode string) TfunctionLoadLibraryCode {
	c.cs.s = append(c.cs.s, libraryCode)
	return (TfunctionLoadLibraryCode)(c)
}

type TfunctionLoadConfig Completed

func (c TfunctionLoadConfig) LibraryCode(libraryCode string) TfunctionLoadLibraryCode {
	c.cs.s = append(c.cs.s, libraryCode)
	return (TfunctionLoadLibraryCode)(c)
}

type TfunctionLoadLibraryCode Completed

func (c TfunctionLoadLibraryCode) Build() Completed {
	c.cs.Build()
	return Completed(c)
}

type TfunctionLoadReplace Completed

func (c TfunctionLoadReplace) Config(config string) TfunctionLoadConfig {
	c.cs.s = append(c.cs.s, config)
	return (TfunctionLoadConfig)(c)
}

func (c TfunctionLoadReplace) LibraryCode(libraryCode string) TfunctionLoadLibraryCode {
	c.cs.s = append(c.cs.s, libraryCode)
	return (TfunctionLoadLibraryCode)(c)
}
