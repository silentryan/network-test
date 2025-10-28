package purpose

// The least information of a stardard accessible test
// support layer 3,4,7 accessible test
// support protocols: ICMP, TCP, UDP, HTTP, HTTPS
type Accessible struct {
	// Target's information
	Protocol string
	Dest     string
	Port     string

	// Below information should writed by program
	// Tester's information
	Src      string
	Computer string

	// The informatio of test result
	Reachable bool
	Failure   string
}

// How to judge which interface is using to visit Internet
// Checking route table and find default route
// The default route usually created when get response from DHCP
func (ac *Accessible) GatherSrc() {

}
