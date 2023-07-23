package types

type PackageService struct {
	Baadbaan  string `json:"baadbaan"`
	Technical string `json:"technical"`
}

type Packages struct {
	Version        string `json:"version"`
	PackageService `json:"package_version"`
}
