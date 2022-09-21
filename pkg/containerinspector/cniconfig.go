package containerinspector

type cniconfig struct {
	Kind        string
	ContainerId string
	Config      string
	IfName      string
	NetworkName string
	Result      result
}

type result struct {
	CniVersion string
	Dns        interface{}
	Interfaces []netInterface
	Ips        []ip
	Routes     []route
}

type netInterface struct {
	Mac     string
	Name    string
	Sandbox string
}

type ip struct {
	Address   string
	Gateway   string
	Interface int
}

type route struct {
	Dst string
}
