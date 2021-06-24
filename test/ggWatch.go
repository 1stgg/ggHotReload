package ggWatch

type callBackFc func(int, int) int
type GGfsWatch struct {
	path string
}

func (w GGfsWatch) on(wType string, cb callBackFc) {

}
