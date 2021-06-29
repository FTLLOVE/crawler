package model

type Profile struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthday  string `json:"birthday"`
	Age       string `json:"age"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Income    string `json:"income"`
	Marriage  string `json:"marriage"`
	Education string `json:"education"`
	Hometown  string `json:"hometown"`
	House     string `json:"house"`
}
