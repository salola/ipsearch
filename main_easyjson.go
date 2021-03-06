// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson89aae3efDecodeGithubComIpsearch1(in *jlexer.Lexer, out *kafkaMsg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "srcip":
			out.SrcIP = string(in.String())
		case "dstip":
			out.DstIP = string(in.String())
		case "logsource_time":
			out.Time = string(in.String())
		case "action":
			out.Action = string(in.String())
		case "badip":
			out.BadIP = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89aae3efEncodeGithubComIpsearch1(out *jwriter.Writer, in kafkaMsg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"srcip\":"
		out.RawString(prefix[1:])
		out.String(string(in.SrcIP))
	}
	{
		const prefix string = ",\"dstip\":"
		out.RawString(prefix)
		out.String(string(in.DstIP))
	}
	{
		const prefix string = ",\"logsource_time\":"
		out.RawString(prefix)
		out.String(string(in.Time))
	}
	{
		const prefix string = ",\"action\":"
		out.RawString(prefix)
		out.String(string(in.Action))
	}
	{
		const prefix string = ",\"badip\":"
		out.RawString(prefix)
		out.String(string(in.BadIP))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v kafkaMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89aae3efEncodeGithubComIpsearch1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v kafkaMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89aae3efEncodeGithubComIpsearch1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *kafkaMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89aae3efDecodeGithubComIpsearch1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *kafkaMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89aae3efDecodeGithubComIpsearch1(l, v)
}
