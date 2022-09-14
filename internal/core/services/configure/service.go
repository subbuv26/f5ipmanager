package configure

import (
	"context"
	ports "f5ipmanager/internal/core/ports/configure"
	spec "f5ipmanager/internal/core/spec/configure"
	"fmt"
	"net"
	"strings"
)

type service struct {
	repo ports.ConfigRepository
}

type Params struct {
	repo ports.ConfigRepository
}

func NewService(params Params) ports.ConfigService {
	return &service{
		repo: params.repo,
	}
}

// AddIPRange adds IPRange
func (s service) AddIPRange(ctx context.Context, request spec.AddIPRangeRequest) (spec.Response, error) {
	label := request.Label
	ipRange := request.Config
	res := spec.Response{Success: false}
	ips, err := parseIPs(ipRange, label)
	if err != nil {
		return res, err
	}
	for _, ip := range ips {
		err := s.repo.AddNewIP(ctx, label, ip)
		if err != nil {
			return res, fmt.Errorf("[CONFIGURE-SVC] Failed to add IP with label:%s, %s, Error: %s",
				ip, label, err)
		}
		//fmt.Printf("Added IP: %s with label: %s ", ip, label)
	}
	//fmt.Printf("Successfully added iprange: %s with label: %s ", ipRange, label)
	res.Success = true
	return res, nil
}

// RemoveIPRange removes IPRange
func (s service) RemoveIPRange(ctx context.Context, request spec.RemoveIPRangeRequest) (spec.Response, error) {
	label := request.Label
	res := spec.Response{Success: false}
	err := s.repo.RemoveIPRange(ctx, label)
	if err != nil {
		return res, fmt.Errorf("[CONFIGURE-SVC] Failed remove IPRange with label:%s, Error: %s", label, err)
	}
	//fmt.Printf("Sucessfully removed IPRange with label: %s", label)
	res.Success = true
	return res, nil
}

// parseIPs parses the ips from the IPRange and returns the list of IPs
func parseIPs(ipRange string, label string) ([]string, error) {
	var ips []string
	ipRangeConfig := strings.Split(ipRange, "-")
	if len(ipRangeConfig) != 2 {
		return nil, fmt.Errorf("[CONFIGURE-SVC] Invalid IP range provided for %s label",
			label)
	}

	startIP := net.ParseIP(ipRangeConfig[0])
	if startIP == nil {
		return nil, fmt.Errorf("[CONFIGURE-SVC] Invalid starting IP %s provided for %s label", ipRangeConfig[0], label)
	}

	endIP := net.ParseIP(ipRangeConfig[1])
	if endIP == nil {
		return nil, fmt.Errorf("[CONFIGURE-SVC] Invalid ending IP %s provided for %s label", ipRangeConfig[1], label)
	}

	for ; startIP.String() != endIP.String(); incIP(startIP) {
		ips = append(ips, startIP.String())
	}
	ips = append(ips, endIP.String())
	return ips, nil
}

// incIP
func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
