package shared

import (
	"database/sql/driver"
	"fmt"
	"sort"

	commonv1 "github.com/PicoTools/pico/pkg/proto/common/v1"
	"google.golang.org/protobuf/proto"
)

// represents custom type to hold capabilities
type Capability uint32

const (
	CapExit               Capability = 0
	CapSleep              Capability = 1
	CapCp                 Capability = 2 << 0
	CapCd                 Capability = 2 << 1
	CapWhoami             Capability = 2 << 2
	CapJobkill            Capability = 2 << 3
	CapCat                Capability = 2 << 4
	CapExec               Capability = 2 << 5
	CapPwd                Capability = 2 << 6
	CapJobs               Capability = 2 << 7
	CapPs                 Capability = 2 << 8
	CapLs                 Capability = 2 << 9
	CapPause              Capability = 2 << 10
	CapMkdir              Capability = 2 << 11
	CapRm                 Capability = 2 << 12
	CapShell              Capability = 2 << 13
	CapShellcodeInjection Capability = 2 << 14
	CapUpload             Capability = 2 << 15
	CapKill               Capability = 2 << 16
	CapMv                 Capability = 2 << 17
	CapDestroy            Capability = 2 << 18
	CapExecDetach         Capability = 2 << 19
	CapExecAssembly       Capability = 2 << 20
	CapPpid               Capability = 2 << 21
	CapDownload           Capability = 2 << 22
)

// Values returns list of strings represented names of capabilities types
func (Capability) Values() []string {
	return []string{
		CapSleep.String(),
		CapLs.String(),
		CapPwd.String(),
		CapCd.String(),
		CapWhoami.String(),
		CapPs.String(),
		CapCat.String(),
		CapExec.String(),
		CapCp.String(),
		CapJobs.String(),
		CapJobkill.String(),
		CapKill.String(),
		CapMv.String(),
		CapMkdir.String(),
		CapRm.String(),
		CapExecAssembly.String(),
		CapShell.String(),
		CapPpid.String(),
		CapExecDetach.String(),
		CapShellcodeInjection.String(),
		CapDownload.String(),
		CapUpload.String(),
		CapPause.String(),
		CapDestroy.String(),
		CapExit.String(),
	}
}

// Value return database ready value for further processing
func (c Capability) Value() (driver.Value, error) {
	return c.String(), nil
}

// String returns string representation of capabilities
func (c Capability) String() string {
	switch c {
	case CapSleep:
		return "cap_sleep"
	case CapLs:
		return "cap_ls"
	case CapPwd:
		return "cap_pwd"
	case CapCd:
		return "cap_cd"
	case CapWhoami:
		return "cap_whoami"
	case CapPs:
		return "cap_ps"
	case CapCat:
		return "cap_cat"
	case CapExec:
		return "cap_exec"
	case CapCp:
		return "cap_cp"
	case CapJobs:
		return "cap_jobs"
	case CapJobkill:
		return "cap_jobkill"
	case CapKill:
		return "cap_kill"
	case CapMv:
		return "cap_mv"
	case CapMkdir:
		return "cap_mkdir"
	case CapRm:
		return "cap_rm"
	case CapExecAssembly:
		return "cap_exec_assembly"
	case CapShell:
		return "cap_shell"
	case CapPpid:
		return "cap_ppid"
	case CapExecDetach:
		return "cap_exec_detach"
	case CapShellcodeInjection:
		return "cap_shellcode_injection"
	case CapDownload:
		return "cap_download"
	case CapUpload:
		return "cap_upload"
	case CapPause:
		return "cap_pause"
	case CapDestroy:
		return "cap_destroy"
	case CapExit:
		return "cap_exit"
	default:
		return "unknown"
	}
}

// Scan converts value to capability
func (c *Capability) Scan(val any) error {
	var s string

	switch v := val.(type) {
	case nil:
		return nil
	case []uint8:
		s = string(v)
	case string:
		s = v
	}

	switch s {
	case CapSleep.String():
		*c = CapSleep
	case CapLs.String():
		*c = CapLs
	case CapPwd.String():
		*c = CapPwd
	case CapCd.String():
		*c = CapCd
	case CapWhoami.String():
		*c = CapWhoami
	case CapPs.String():
		*c = CapPs
	case CapCat.String():
		*c = CapCat
	case CapExec.String():
		*c = CapExec
	case CapCp.String():
		*c = CapCp
	case CapJobs.String():
		*c = CapJobs
	case CapJobkill.String():
		*c = CapJobkill
	case CapKill.String():
		*c = CapKill
	case CapMv.String():
		*c = CapMv
	case CapMkdir.String():
		*c = CapMkdir
	case CapRm.String():
		*c = CapRm
	case CapExecAssembly.String():
		*c = CapExecAssembly
	case CapShell.String():
		*c = CapShell
	case CapPpid.String():
		*c = CapPpid
	case CapExecDetach.String():
		*c = CapExecDetach
	case CapShellcodeInjection.String():
		*c = CapShellcodeInjection
	case CapDownload.String():
		*c = CapDownload
	case CapUpload.String():
		*c = CapUpload
	case CapPause.String():
		*c = CapPause
	case CapDestroy.String():
		*c = CapDestroy
	case CapExit.String():
		*c = CapExit
	}
	return nil
}

// Marshal converts proto message to byte array
func (c Capability) Marshal(data any) ([]byte, error) {
	switch c {
	case CapSleep:
		v, ok := data.(*commonv1.CapSleep)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapLs:
		v, ok := data.(*commonv1.CapLs)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapPwd:
		v, ok := data.(*commonv1.CapPwd)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapCd:
		v, ok := data.(*commonv1.CapCd)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapWhoami:
		v, ok := data.(*commonv1.CapWhoami)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapPs:
		v, ok := data.(*commonv1.CapPs)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapCat:
		v, ok := data.(*commonv1.CapCat)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapExec:
		v, ok := data.(*commonv1.CapExec)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapCp:
		v, ok := data.(*commonv1.CapCp)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapJobs:
		v, ok := data.(*commonv1.CapJobs)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapJobkill:
		v, ok := data.(*commonv1.CapJobkill)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapKill:
		v, ok := data.(*commonv1.CapKill)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapMv:
		v, ok := data.(*commonv1.CapMv)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapMkdir:
		v, ok := data.(*commonv1.CapMkdir)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapRm:
		v, ok := data.(*commonv1.CapRm)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapExecAssembly:
		v, ok := data.(*commonv1.CapExecAssembly)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapShell:
		v, ok := data.(*commonv1.CapShell)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapPpid:
		v, ok := data.(*commonv1.CapPpid)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapExecDetach:
		v, ok := data.(*commonv1.CapExecDetach)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapShellcodeInjection:
		v, ok := data.(*commonv1.CapShellcodeInjection)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapDownload:
		v, ok := data.(*commonv1.CapDownload)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapUpload:
		v, ok := data.(*commonv1.CapUpload)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapPause:
		v, ok := data.(*commonv1.CapPause)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapDestroy:
		v, ok := data.(*commonv1.CapDestroy)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	case CapExit:
		v, ok := data.(*commonv1.CapExit)
		if !ok {
			return nil, fmt.Errorf("%s: invalid argument to marshal", c.String())
		}
		return proto.Marshal(v)
	default:
		return nil, fmt.Errorf("%s: unknown capability type to marshal", c.String())
	}
}

// Unmarshal converts raw binary to proto message
func (c Capability) Unmarshal(data []byte) (any, error) {
	switch c {
	case CapSleep:
		v := new(commonv1.CapSleep)
		return v, proto.Unmarshal(data, v)
	case CapLs:
		v := new(commonv1.CapLs)
		return v, proto.Unmarshal(data, v)
	case CapPwd:
		v := new(commonv1.CapPwd)
		return v, proto.Unmarshal(data, v)
	case CapCd:
		v := new(commonv1.CapCd)
		return v, proto.Unmarshal(data, v)
	case CapWhoami:
		v := new(commonv1.CapWhoami)
		return v, proto.Unmarshal(data, v)
	case CapPs:
		v := new(commonv1.CapPs)
		return v, proto.Unmarshal(data, v)
	case CapCat:
		v := new(commonv1.CapCat)
		return v, proto.Unmarshal(data, v)
	case CapExec:
		v := new(commonv1.CapExec)
		return v, proto.Unmarshal(data, v)
	case CapCp:
		v := new(commonv1.CapCp)
		return v, proto.Unmarshal(data, v)
	case CapJobs:
		v := new(commonv1.CapJobs)
		return v, proto.Unmarshal(data, v)
	case CapJobkill:
		v := new(commonv1.CapJobkill)
		return v, proto.Unmarshal(data, v)
	case CapKill:
		v := new(commonv1.CapKill)
		return v, proto.Unmarshal(data, v)
	case CapMv:
		v := new(commonv1.CapMv)
		return v, proto.Unmarshal(data, v)
	case CapMkdir:
		v := new(commonv1.CapMkdir)
		return v, proto.Unmarshal(data, v)
	case CapRm:
		v := new(commonv1.CapRm)
		return v, proto.Unmarshal(data, v)
	case CapExecAssembly:
		v := new(commonv1.CapExecAssembly)
		return v, proto.Unmarshal(data, v)
	case CapShell:
		v := new(commonv1.CapShell)
		return v, proto.Unmarshal(data, v)
	case CapPpid:
		v := new(commonv1.CapPpid)
		return v, proto.Unmarshal(data, v)
	case CapExecDetach:
		v := new(commonv1.CapExecDetach)
		return v, proto.Unmarshal(data, v)
	case CapShellcodeInjection:
		v := new(commonv1.CapShellcodeInjection)
		return v, proto.Unmarshal(data, v)
	case CapDownload:
		v := new(commonv1.CapDownload)
		return v, proto.Unmarshal(data, v)
	case CapUpload:
		v := new(commonv1.CapUpload)
		return v, proto.Unmarshal(data, v)
	case CapPause:
		v := new(commonv1.CapPause)
		return v, proto.Unmarshal(data, v)
	case CapDestroy:
		v := new(commonv1.CapDestroy)
		return v, proto.Unmarshal(data, v)
	case CapExit:
		v := new(commonv1.CapExit)
		return v, proto.Unmarshal(data, v)
	default:
		return nil, fmt.Errorf("%s: unknown capability type to unmarshal", c.String())
	}
}

// ValidateMask validates if capability presents in binary mask
func (c Capability) ValidateMask(cap uint32) bool {
	return cap&uint32(c) == uint32(c)
}

// SupportedCaps returns list of capabilities contained in binary mask
func SupportedCaps(mask uint32) []Capability {
	var t []Capability
	if mask&uint32(CapSleep) == uint32(CapSleep) {
		t = append(t, CapSleep)
	}
	if mask&uint32(CapLs) == uint32(CapLs) {
		t = append(t, CapLs)
	}
	if mask&uint32(CapPwd) == uint32(CapPwd) {
		t = append(t, CapPwd)
	}
	if mask&uint32(CapCd) == uint32(CapCd) {
		t = append(t, CapCd)
	}
	if mask&uint32(CapWhoami) == uint32(CapWhoami) {
		t = append(t, CapWhoami)
	}
	if mask&uint32(CapPs) == uint32(CapPs) {
		t = append(t, CapPs)
	}
	if mask&uint32(CapCat) == uint32(CapCat) {
		t = append(t, CapCat)
	}
	if mask&uint32(CapExec) == uint32(CapExec) {
		t = append(t, CapExec)
	}
	if mask&uint32(CapCp) == uint32(CapCp) {
		t = append(t, CapCp)
	}
	if mask&uint32(CapJobs) == uint32(CapJobs) {
		t = append(t, CapJobs)
	}
	if mask&uint32(CapJobkill) == uint32(CapJobkill) {
		t = append(t, CapJobkill)
	}
	if mask&uint32(CapKill) == uint32(CapKill) {
		t = append(t, CapKill)
	}
	if mask&uint32(CapMv) == uint32(CapMv) {
		t = append(t, CapMv)
	}
	if mask&uint32(CapMkdir) == uint32(CapMkdir) {
		t = append(t, CapMkdir)
	}
	if mask&uint32(CapRm) == uint32(CapRm) {
		t = append(t, CapRm)
	}
	if mask&uint32(CapExecAssembly) == uint32(CapExecAssembly) {
		t = append(t, CapExecAssembly)
	}
	if mask&uint32(CapShellcodeInjection) == uint32(CapShellcodeInjection) {
		t = append(t, CapShellcodeInjection)
	}
	if mask&uint32(CapDownload) == uint32(CapDownload) {
		t = append(t, CapDownload)
	}
	if mask&uint32(CapUpload) == uint32(CapUpload) {
		t = append(t, CapUpload)
	}
	if mask&uint32(CapPause) == uint32(CapPause) {
		t = append(t, CapPause)
	}
	if mask&uint32(CapDestroy) == uint32(CapDestroy) {
		t = append(t, CapDestroy)
	}
	if mask&uint32(CapExecDetach) == uint32(CapExecDetach) {
		t = append(t, CapExecDetach)
	}
	if mask&uint32(CapShell) == uint32(CapShell) {
		t = append(t, CapShell)
	}
	if mask&uint32(CapPpid) == uint32(CapPpid) {
		t = append(t, CapPpid)
	}
	if mask&uint32(CapExit) == uint32(CapExit) {
		t = append(t, CapExit)
	}

	// sort capabilities by theirs values
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	return t
}
