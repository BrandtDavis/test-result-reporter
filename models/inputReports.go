package models

type MochawesomeReport struct {
	start string `json:"start"`
	end   string `json:"end"`
	tests []Test `json:tests`
}

type Test struct {
	desc     string   `json:"desc"`
	testData TestData `json:"testData"`
}

type TestData struct {
	testName  string `json:"testName"`
	numPasses int64  `json:"numPasses"`
	numFails  int64  `json:"numFails"`
}
