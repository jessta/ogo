package ofp10

import (
	"io"
	"net"
	"bytes"
	"encoding/binary"
)

// ofp_action_type 1.0
const (
	OFPAT_OUTPUT = iota
	OFPAT_SET_VLAN_VID
	OFPAT_SET_VLAN_PCP
	OFPAT_STRIP_VLAN
	OFPAT_SET_DL_SRC
	OFPAT_SET_DL_DST
	OFPAT_SET_NW_SRC
	OFPAT_SET_NW_DST
	OFPAT_SET_NW_TOS
	OFPAT_SET_TP_SRC
	OFPAT_SET_TP_DST
	OFPAT_ENQUEUE
	OFPAT_VENDOR = 0xffff
)

// ofp_action_header 1.0
type OfpActionHeader struct {
	Type uint16
	Length uint16
	Pad [4]uint8
}

func (a *OfpActionHeader) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	n, err = buf.Read(b)
	if n == 0 {
		return
	}
	return n, io.EOF
}

func (a *OfpActionHeader) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return 
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return 
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Pad); err != nil {
		return
	}
	n += 4
	return
}

// ofp_action_output 1.0
type OfpActionOutput struct {
	Type uint16
	Length uint16
	Port uint16
	MaxLen uint16
}

func NewActionOutput() *OfpActionOutput {
	act := new(OfpActionOutput)
	act.Type = OFPAT_OUTPUT
	act.Length = 8
	act.Port = OFPP_FLOOD
	return act
}

func (a *OfpActionOutput) Len() (n uint16) {
	return a.Length
}

func (a *OfpActionOutput) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	if n, err = buf.Read(b); n == 0 {
		return
	}
	return n, io.EOF
}

func (a *OfpActionOutput) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Port); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.MaxLen); err != nil {
		return
	}
	n += 2
	return
}

// ofp_action_enqueue 1.0
type OfpActionEnqueue struct {
	Type uint16
	Length uint16
	Port uint16
	Pad [6]uint8
	QueueId uint32
}

func NewActionEnqueue() *OfpActionEnqueue {
	a := new(OfpActionEnqueue)
	a.Type = OFPAT_ENQUEUE
	a.Length = 16
	return a
}

func (a *OfpActionEnqueue) Len() (n uint16) {
	return a.Length
}

func (a *OfpActionEnqueue) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionEnqueue) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Port); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Pad); err != nil {
		return
	}
	n += 6
	if err = binary.Read(buf, binary.BigEndian, &a.QueueId); err != nil {
		return
	}
	n += 4
	return
}

// ofp_action_vlan_vid 1.0
type OfpActionVLANVID struct {
	Type uint16
	Length uint16
	VLANVID uint16
	Pad [2]uint8
}

func NewActionVLANVID() *OfpActionVLANVID {
	a := new(OfpActionVLANVID)
	a.Type = OFPAT_SET_VLAN_VID
	a.Length = 8
	a.VLANVID = 0xffff
	return a
}

func (a *OfpActionVLANVID) Len() (n uint16) {
	return a.Length
}

func (a *OfpActionVLANVID) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionVLANVID) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.VLANVID); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Pad); err != nil {
		return
	}
	n += 2
	return
}

// ofp_action_vlan_pcp 1.0
type OfpActionVLANPCP struct {
	Type uint16
	Length uint16
	VLANPCP uint8
	Pad [3]uint8
}

func NewActionVLANPCP() *OfpActionVLANPCP {
	a := new(OfpActionVLANPCP)
	a.Type = OFPAT_SET_VLAN_PCP
	a.Length = 8
	return a
}

func (a *OfpActionVLANPCP) Len() (n uint16) {
	return 8
}

func (a *OfpActionVLANPCP) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionVLANPCP) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.VLANPCP); err != nil {
		return
	}
	n += 1
	if err = binary.Read(buf, binary.BigEndian, &a.Pad); err != nil {
		return
	}
	n += 3
	return
}

// ofp_action_dl_addr 1.0
type OfpActionDLAddr struct {
	Type uint16
	Length uint16
	DLAddr net.HardwareAddr
	Pad [6]uint8
}

func NewActionDLSrc() *OfpActionDLAddr {
	a := new(OfpActionDLAddr)
	a.Type = OFPAT_SET_DL_SRC
	a.Length = 16
	a.DLAddr = make([]byte, OFP_ETH_ALEN)
	return a
}

func NewActionDLDst() *OfpActionDLAddr {
	a := new(OfpActionDLAddr)
	a.Type = OFPAT_SET_DL_DST
	a.Length = 16
	a.DLAddr = make([]byte, OFP_ETH_ALEN)
	return a
}

func (a *OfpActionDLAddr) Len() (n uint16) {
	return a.Length
}

func (a *OfpActionDLAddr) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a.Type)
	err = binary.Write(buf, binary.BigEndian, a.Length)
	err = binary.Write(buf, binary.BigEndian, a.DLAddr)
	err = binary.Write(buf, binary.BigEndian, a.Pad)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionDLAddr) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	a.DLAddr = make([]byte, OFP_ETH_ALEN)
	if err = binary.Read(buf, binary.BigEndian, &a.DLAddr); err != nil {
		return
	}
	n += OFP_ETH_ALEN
	if err = binary.Read(buf, binary.BigEndian, &a.Pad); err != nil {
		return
	}
	n += 6
	return
}

// ofp_action_nw_addr 1.0
type OfpActionNWAddr struct {
	Type uint16
	Length uint16
	NWAddr net.IP
}

func NewActionNWSrc() *OfpActionNWAddr {
	a := new(OfpActionNWAddr)
	a.Type = OFPAT_SET_NW_SRC
	a.Length = 8
	a.NWAddr = make([]byte, 4)
	return a
}

func NewActionNWDst() *OfpActionNWAddr {
	a := new(OfpActionNWAddr)
	a.Type = OFPAT_SET_NW_DST
	a.Length = 8
	a.NWAddr = make([]byte, 4)
	return a
}

func (a *OfpActionNWAddr) Len() (n uint16) {
	return 8
}

func (a *OfpActionNWAddr) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a.Type)
	err = binary.Write(buf, binary.BigEndian, a.Length)
	err = binary.Write(buf, binary.BigEndian, a.NWAddr)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionNWAddr) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	a.NWAddr = make([]byte, 4)
	if err = binary.Read(buf, binary.BigEndian, &a.NWAddr); err != nil {
		return
	}
	n += 4
	return
}

// ofp_action_nw_tos 1.0
type OfpActionNWTOS struct {
	Type uint16
	Length uint16
	NWTOS uint8
	Pad [3]uint8
}

func NewActionNWTOS() *OfpActionNWTOS {
	a := new(OfpActionNWTOS)
	a.Type = OFPAT_SET_NW_TOS
	a.Length = 8
	return a
}

func (a *OfpActionNWTOS) Len() (n uint16) {
	return a.Length
}

func (a *OfpActionNWTOS) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionNWTOS) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.NWTOS); err != nil {
		return
	}
	n += 1
	if err = binary.Read(buf, binary.BigEndian, &a.Pad); err != nil {
		return
	}
	n += 3
	return
}

// ofp_action_tp_port 1.0
type OfpActionTPPort struct {
	Type uint16
	Length uint16
	TPPort uint16
	Pad [2]uint8
}

func NewActionTPSrc() *OfpActionTPPort {
	a := new(OfpActionTPPort)
	a.Type = OFPAT_SET_TP_SRC
	a.Length = 8
	return a
}

func NewActionTPDst() *OfpActionTPPort {
	a := new(OfpActionTPPort)
	a.Type = OFPAT_SET_TP_DST
	a.Length = 8
	return a
}

func (a *OfpActionTPPort) Len() (n uint16) {
	return a.Length
}

func (a *OfpActionTPPort) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionTPPort) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.TPPort); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Pad); err != nil {
		return
	}
	n += 8
	return
}

// ofp_action_vendor_header 1.0
type OfpActionVendorPort struct {
	Type uint16
	Length uint16
	Vendor uint32
}

func NewActionVendorPort() *OfpActionVendorPort {
	a := new(OfpActionVendorPort)
	a.Type = OFPAT_VENDOR
	a.Length = 8
	return a
}

func (a *OfpActionVendorPort) Len() (n uint16) {
	return a.Length
}

func (a *OfpActionVendorPort) Read(b []byte) (n int, err error) {
	a.Length = a.Len()
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, a)
	n, err = buf.Read(b)
	return
}

func (a *OfpActionVendorPort) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.Vendor); err != nil {
		return
	}
	n += 4
	return
}
