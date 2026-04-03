package ecies

// Null-scheme
type NullSch struct {
}

// 获取 SUPI
/*func (p *NullSch) GetSupiFromSuci(suci types3gpp.Suci) (*types3gpp.Supi, error) {
	supi := &types3gpp.Supi{}

	supiType := suci.GetSupiType()
	switch supiType {
	case IMSIType:
		imsi, err := suci.GetImsi()
		if err != nil {
			return nil, err
		}
		supi.SetImsi(imsi)
		return supi, nil
	case NAIType:
		//	todo
		return nil, nil
	default:
		return nil, fmt.Errorf("invalid supi type(0-imsi, 1-nai)")
	}
}
*/
