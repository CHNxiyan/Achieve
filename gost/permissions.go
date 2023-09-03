package gost

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	glob "github.com/ryanuber/go-glob"
)

// Permission is a rule for blacklist and whitelist.
type Permission struct {
	Actions StringSet
	Hosts   StringSet
	Ports   PortSet
}

// PortRange specifies the range of port, such as 1000-2000.
type PortRange struct {
	Min, Max int
}

// ParsePortRange parses the s to a PortRange.
// The s may be a '*' means 0-65535.
func ParsePortRange(s string) (*PortRange, error) {
	if s == "*" {
		return &PortRange{Min: 0, Max: 65535}, nil
	}

	minmax := strings.Split(s, "-")
	switch len(minmax) {
	case 1:
		port, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		if port < 0 || port > 65535 {
			return nil, fmt.Errorf("invalid port: %s", s)
		}
		return &PortRange{Min: port, Max: port}, nil
	case 2:
		min, err := strconv.Atoi(minmax[0])
		if err != nil {
			return nil, err
		}
		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			return nil, err
		}

		realmin := maxint(0, minint(min, max))
		realmax := minint(65535, maxint(min, max))

		return &PortRange{Min: realmin, Max: realmax}, nil
	default:
		return nil, fmt.Errorf("invalid range: %s", s)
	}
}

// Contains checks whether the value is within this range.
func (ir *PortRange) Contains(value int) bool {
	return value >= ir.Min && value <= ir.Max
}

// PortSet is a set of PortRange
type PortSet []PortRange

// ParsePortSet parses the s to a PortSet.
// The s shoud be a comma separated string.
func ParsePortSet(s string) (*PortSet, error) {
	ps := &PortSet{}

	if s == "" {
		return nil, errors.New("must specify at least one port")
	}

	ranges := strings.Split(s, ",")

	for _, r := range ranges {
		portRange, err := ParsePortRange(r)

		if err != nil {
			return nil, err
		}

		*ps = append(*ps, *portRange)
	}

	return ps, nil
}

// Contains checks whether the value is within this port set.
func (ps *PortSet) Contains(value int) bool {
	for _, portRange := range *ps {
		if portRange.Contains(value) {
			return true
		}
	}

	return false
}

// StringSet is a set of string.
type StringSet []string

// ParseStringSet parses the s to a StringSet.
// The s shoud be a comma separated string.
func ParseStringSet(s string) (*StringSet, error) {
	ss := &StringSet{}
	if s == "" {
		return nil, errors.New("cannot be empty")
	}

	*ss = strings.Split(s, ",")

	return ss, nil
}

// Contains checks whether the string subj within this StringSet.
func (ss *StringSet) Contains(subj string) bool {
	for _, s := range *ss {
		if glob.Glob(s, subj) {
			return true
		}
	}

	return false
}

// Permissions is a set of Permission.
type Permissions []Permission

// ParsePermissions parses the s to a Permissions.
func ParsePermissions(s string) (*Permissions, error) {
	ps := &Permissions{}

	if s == "" {
		return &Permissions{}, nil
	}

	perms := strings.Split(s, " ")

	for _, perm := range perms {
		parts := strings.Split(perm, ":")

		switch len(parts) {
		case 3:
			actions, err := ParseStringSet(parts[0])

			if err != nil {
				return nil, fmt.Errorf("action list must look like connect,bind given: %s", parts[0])
			}

			hosts, err := ParseStringSet(parts[1])

			if err != nil {
				return nil, fmt.Errorf("hosts list must look like google.pl,*.google.com given: %s", parts[1])
			}

			ports, err := ParsePortSet(parts[2])

			if err != nil {
				return nil, fmt.Errorf("ports list must look like 80,8000-9000, given: %s", parts[2])
			}

			permission := Permission{Actions: *actions, Hosts: *hosts, Ports: *ports}

			*ps = append(*ps, permission)
		default:
			return nil, fmt.Errorf("permission must have format [actions]:[hosts]:[ports] given: %s", perm)
		}
	}

	return ps, nil
}

// Can tests whether the given action and host:port is allowed by this Permissions.
func (ps *Permissions) Can(action string, host string, port int) bool {
	for _, p := range *ps {
		if p.Actions.Contains(action) && p.Hosts.Contains(host) && p.Ports.Contains(port) {
			return true
		}
	}

	return false
}

func minint(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxint(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Can tests whether the given action and address is allowed by the whitelist and blacklist.
func Can(action string, addr string, whitelist, blacklist *Permissions) bool {
	if !strings.Contains(addr, ":") {
		addr = addr + ":80"
	}
	host, strport, err := net.SplitHostPort(addr)

	if err != nil {
		return false
	}

	port, err := strconv.Atoi(strport)

	if err != nil {
		return false
	}

	return (whitelist == nil || whitelist.Can(action, host, port)) &&
		(blacklist == nil || !blacklist.Can(action, host, port))
}
