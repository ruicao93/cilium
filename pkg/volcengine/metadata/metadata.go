// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package metadata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cilium/cilium/pkg/safeio"
	"github.com/cilium/cilium/pkg/time"
)

const (
	metadataURL = "http://100.96.0.96/latest"
)

func GetInstanceID(ctx context.Context) (string, error) {
	return getMetadata(ctx, "instance_id")
}

func GetInstanceType(ctx context.Context) (string, error) {
	return getMetadata(ctx, "instance_type_id")
}

func GetZoneID(ctx context.Context) (string, error) {
	return getMetadata(ctx, "availability_zone")
}

func GetVPCID(ctx context.Context) (string, error) {
	return getMetadata(ctx, "vpc_id")
}

func GetVPCCIDRBlock(ctx context.Context) (string, error) {
	return getMetadata(ctx, "vpc_cidr_block")
}

func getMetadata(ctx context.Context, path string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := fmt.Sprintf("%s/%s", metadataURL, path)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("metadata service return unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBytes, err := safeio.ReadAllLimit(resp.Body, safeio.MB)
	if err != nil {
		return "", err
	}

	return string(respBytes), nil
}
