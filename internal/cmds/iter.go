//go:build go1.23

package cmds

import "iter"

func (c HmsetFieldValue) FieldValueIter(seq iter.Seq2[string, string]) HmsetFieldValue {
	for field, value := range seq {
		c.cs.s = append(c.cs.s, field, value)
	}
	return c
}

func (c HsetFieldValue) FieldValueIter(seq iter.Seq2[string, string]) HsetFieldValue {
	for field, value := range seq {
		c.cs.s = append(c.cs.s, field, value)
	}
	return c
}

func (c XaddFieldValue) FieldValueIter(seq iter.Seq2[string, string]) XaddFieldValue {
	for field, value := range seq {
		c.cs.s = append(c.cs.s, field, value)
	}
	return c
}
