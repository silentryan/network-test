package local

// Reading some local's network settings

type Localnet struct {
	Hostname  string
	Dnsserver string
	Ips       map[string]string
}

// make sure map is initiaized
func createlocalnet() *Localnet {
	ips := make(map[string]string)
	return &Localnet{
		"",
		"",
		ips,
	}
}

func ReadLocalNetwork() *Localnet {
	var localnet *Localnet = createlocalnet()

	return localnet
}

func GetIntofDefult() string {
	return ""
}
