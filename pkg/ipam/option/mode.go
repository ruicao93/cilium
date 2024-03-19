package option

var (
	ipamCRDSupportModes = map[string]struct{}{
		IPAMENI:          {},
		IPAMAzure:        {},
		IPAMAlibabaCloud: {},
		IPAMVolcengine:   {},
	}
)

func IsSupportedIPAMCRDMode(mode string) bool {
	_, ok := ipamCRDSupportModes[mode]
	return ok
}
