package server

type Serverinput struct {
	Project     string
	Name        string
	MachineType string
	ImageProj   string
	Imagefamily string
	Zone        string
	Network     string
	Subnet      string
	Region      string
	User        string
	Pem         string
}

type Serveroutput struct {
	Name      string
	PrivateIP string
	PublicIP  string
}

type Svrlistoutput struct {
	Name      string
	PrivateIP string
	PublicIP  string
	Status    string
}

type Serverdelinput struct {
	Project string
	Name    string
	Zone    string
}
