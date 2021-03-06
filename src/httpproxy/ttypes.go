// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package httpproxy

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type E_CLIENT_TYPE int64

const (
	E_CLIENT_TYPE_UNKNOW       E_CLIENT_TYPE = 0
	E_CLIENT_TYPE_ANDROID      E_CLIENT_TYPE = 1
	E_CLIENT_TYPE_IOS          E_CLIENT_TYPE = 2
	E_CLIENT_TYPE_WEB          E_CLIENT_TYPE = 3
	E_CLIENT_TYPE_WEB_COSCHOOL E_CLIENT_TYPE = 4
	E_CLIENT_TYPE_WEB_GROUP    E_CLIENT_TYPE = 5
)

func (p E_CLIENT_TYPE) String() string {
	switch p {
	case E_CLIENT_TYPE_UNKNOW:
		return "UNKNOW"
	case E_CLIENT_TYPE_ANDROID:
		return "ANDROID"
	case E_CLIENT_TYPE_IOS:
		return "IOS"
	case E_CLIENT_TYPE_WEB:
		return "WEB"
	case E_CLIENT_TYPE_WEB_COSCHOOL:
		return "WEB_COSCHOOL"
	case E_CLIENT_TYPE_WEB_GROUP:
		return "WEB_GROUP"
	}
	return "<UNSET>"
}

func E_CLIENT_TYPEFromString(s string) (E_CLIENT_TYPE, error) {
	switch s {
	case "UNKNOW":
		return E_CLIENT_TYPE_UNKNOW, nil
	case "ANDROID":
		return E_CLIENT_TYPE_ANDROID, nil
	case "IOS":
		return E_CLIENT_TYPE_IOS, nil
	case "WEB":
		return E_CLIENT_TYPE_WEB, nil
	case "WEB_COSCHOOL":
		return E_CLIENT_TYPE_WEB_COSCHOOL, nil
	case "WEB_GROUP":
		return E_CLIENT_TYPE_WEB_GROUP, nil
	}
	return E_CLIENT_TYPE(0), fmt.Errorf("not a valid E_CLIENT_TYPE string")
}

func E_CLIENT_TYPEPtr(v E_CLIENT_TYPE) *E_CLIENT_TYPE { return &v }

func (p E_CLIENT_TYPE) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *E_CLIENT_TYPE) UnmarshalText(text []byte) error {
	q, err := E_CLIENT_TYPEFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

type E_HTTP_COMMAND int64

const (
	E_HTTP_COMMAND_ECMD_NULL              E_HTTP_COMMAND = 0
	E_HTTP_COMMAND_ECMD_PHOTO_WALL_ADD    E_HTTP_COMMAND = 1000
	E_HTTP_COMMAND_ECMD_PHOTO_WALL_DELETE E_HTTP_COMMAND = 1001
	E_HTTP_COMMAND_ECMD_PHOTO_WALL_QUERY  E_HTTP_COMMAND = 1002
)

func (p E_HTTP_COMMAND) String() string {
	switch p {
	case E_HTTP_COMMAND_ECMD_NULL:
		return "ECMD_NULL"
	case E_HTTP_COMMAND_ECMD_PHOTO_WALL_ADD:
		return "ECMD_PHOTO_WALL_ADD"
	case E_HTTP_COMMAND_ECMD_PHOTO_WALL_DELETE:
		return "ECMD_PHOTO_WALL_DELETE"
	case E_HTTP_COMMAND_ECMD_PHOTO_WALL_QUERY:
		return "ECMD_PHOTO_WALL_QUERY"
	}
	return "<UNSET>"
}

func E_HTTP_COMMANDFromString(s string) (E_HTTP_COMMAND, error) {
	switch s {
	case "ECMD_NULL":
		return E_HTTP_COMMAND_ECMD_NULL, nil
	case "ECMD_PHOTO_WALL_ADD":
		return E_HTTP_COMMAND_ECMD_PHOTO_WALL_ADD, nil
	case "ECMD_PHOTO_WALL_DELETE":
		return E_HTTP_COMMAND_ECMD_PHOTO_WALL_DELETE, nil
	case "ECMD_PHOTO_WALL_QUERY":
		return E_HTTP_COMMAND_ECMD_PHOTO_WALL_QUERY, nil
	}
	return E_HTTP_COMMAND(0), fmt.Errorf("not a valid E_HTTP_COMMAND string")
}

func E_HTTP_COMMANDPtr(v E_HTTP_COMMAND) *E_HTTP_COMMAND { return &v }

func (p E_HTTP_COMMAND) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *E_HTTP_COMMAND) UnmarshalText(text []byte) error {
	q, err := E_HTTP_COMMANDFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - SVer
//  - NClientType
//  - UID
//  - NCMDID
//  - Sequence
//  - Buf
//  - SToken
type ReqHttpPackage struct {
	SVer        string `thrift:"sVer,1" json:"sVer"`
	NClientType int32  `thrift:"nClientType,2" json:"nClientType"`
	UID         string `thrift:"uid,3" json:"uid"`
	NCMDID      int32  `thrift:"nCMDID,4" json:"nCMDID"`
	Sequence    int64  `thrift:"sequence,5" json:"sequence"`
	Buf         []byte `thrift:"buf,6" json:"buf"`
	SToken      string `thrift:"sToken,7" json:"sToken"`
}

func NewReqHttpPackage() *ReqHttpPackage {
	return &ReqHttpPackage{}
}

func (p *ReqHttpPackage) GetSVer() string {
	return p.SVer
}

func (p *ReqHttpPackage) GetNClientType() int32 {
	return p.NClientType
}

func (p *ReqHttpPackage) GetUID() string {
	return p.UID
}

func (p *ReqHttpPackage) GetNCMDID() int32 {
	return p.NCMDID
}

func (p *ReqHttpPackage) GetSequence() int64 {
	return p.Sequence
}

func (p *ReqHttpPackage) GetBuf() []byte {
	return p.Buf
}

func (p *ReqHttpPackage) GetSToken() string {
	return p.SToken
}
func (p *ReqHttpPackage) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ReqHttpPackage) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.SVer = v
	}
	return nil
}

func (p *ReqHttpPackage) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.NClientType = v
	}
	return nil
}

func (p *ReqHttpPackage) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.UID = v
	}
	return nil
}

func (p *ReqHttpPackage) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.NCMDID = v
	}
	return nil
}

func (p *ReqHttpPackage) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Sequence = v
	}
	return nil
}

func (p *ReqHttpPackage) readField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Buf = v
	}
	return nil
}

func (p *ReqHttpPackage) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.SToken = v
	}
	return nil
}

func (p *ReqHttpPackage) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ReqHttpPackage"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ReqHttpPackage) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sVer", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sVer: ", p), err)
	}
	if err := oprot.WriteString(string(p.SVer)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sVer (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sVer: ", p), err)
	}
	return err
}

func (p *ReqHttpPackage) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("nClientType", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:nClientType: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.NClientType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nClientType (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:nClientType: ", p), err)
	}
	return err
}

func (p *ReqHttpPackage) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("uid", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:uid: ", p), err)
	}
	if err := oprot.WriteString(string(p.UID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.uid (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:uid: ", p), err)
	}
	return err
}

func (p *ReqHttpPackage) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("nCMDID", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:nCMDID: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.NCMDID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nCMDID (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:nCMDID: ", p), err)
	}
	return err
}

func (p *ReqHttpPackage) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sequence", thrift.I64, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:sequence: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Sequence)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sequence (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:sequence: ", p), err)
	}
	return err
}

func (p *ReqHttpPackage) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("buf", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:buf: ", p), err)
	}
	if err := oprot.WriteBinary(p.Buf); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.buf (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:buf: ", p), err)
	}
	return err
}

func (p *ReqHttpPackage) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sToken", thrift.STRING, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:sToken: ", p), err)
	}
	if err := oprot.WriteString(string(p.SToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sToken (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:sToken: ", p), err)
	}
	return err
}

func (p *ReqHttpPackage) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReqHttpPackage(%+v)", *p)
}

// Attributes:
//  - NCMDID
//  - Sequence
//  - Result_
//  - SMessage
//  - Buf
type RespHttpPackage struct {
	NCMDID   int32  `thrift:"nCMDID,1" json:"nCMDID"`
	Sequence int64  `thrift:"sequence,2" json:"sequence"`
	Result_  int32  `thrift:"result,3" json:"result"`
	SMessage string `thrift:"sMessage,4" json:"sMessage"`
	Buf      []byte `thrift:"buf,5" json:"buf"`
}

func NewRespHttpPackage() *RespHttpPackage {
	return &RespHttpPackage{}
}

func (p *RespHttpPackage) GetNCMDID() int32 {
	return p.NCMDID
}

func (p *RespHttpPackage) GetSequence() int64 {
	return p.Sequence
}

func (p *RespHttpPackage) GetResult_() int32 {
	return p.Result_
}

func (p *RespHttpPackage) GetSMessage() string {
	return p.SMessage
}

func (p *RespHttpPackage) GetBuf() []byte {
	return p.Buf
}
func (p *RespHttpPackage) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RespHttpPackage) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.NCMDID = v
	}
	return nil
}

func (p *RespHttpPackage) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Sequence = v
	}
	return nil
}

func (p *RespHttpPackage) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Result_ = v
	}
	return nil
}

func (p *RespHttpPackage) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.SMessage = v
	}
	return nil
}

func (p *RespHttpPackage) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Buf = v
	}
	return nil
}

func (p *RespHttpPackage) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("RespHttpPackage"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RespHttpPackage) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("nCMDID", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:nCMDID: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.NCMDID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nCMDID (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:nCMDID: ", p), err)
	}
	return err
}

func (p *RespHttpPackage) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sequence", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sequence: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Sequence)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sequence (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sequence: ", p), err)
	}
	return err
}

func (p *RespHttpPackage) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("result", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:result: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.Result_)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.result (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:result: ", p), err)
	}
	return err
}

func (p *RespHttpPackage) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sMessage", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:sMessage: ", p), err)
	}
	if err := oprot.WriteString(string(p.SMessage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sMessage (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:sMessage: ", p), err)
	}
	return err
}

func (p *RespHttpPackage) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("buf", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:buf: ", p), err)
	}
	if err := oprot.WriteBinary(p.Buf); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.buf (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:buf: ", p), err)
	}
	return err
}

func (p *RespHttpPackage) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RespHttpPackage(%+v)", *p)
}
