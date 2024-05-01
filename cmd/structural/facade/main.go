package main

import facadePkg "github.com/vietpham102301/go-design-patterns-practice/structural_patterns/facade"

func main() {
	facade := facadePkg.NewCloudFacade()
	facade.DeployApplication()
	facade.TeardownApplication()
}
