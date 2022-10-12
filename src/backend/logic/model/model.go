package model

type Numericals struct {
	MethodID          int     `json:"id"`
	MethodTitle       string  `json:"title"`
	MethodDescription string  `json:"description"`
	MethodIcon        *string `json:"icon"`
}

type Task struct {
	TaskName        string `json:"task_name"`
	TaskDescription string `json:"task_description"`
}

type Boundary2ndODEInputs struct {
	GridSize         int     `json:"n"`
	EpsilonParam     float64 `json:"e"`
	DifferenceScheme string  `json:"diff_scheme"`
}

type SolutionData struct {
	XValues []float64 `json:"xVals"`
	YValues []float64 `json:"yVals"`
}

type NumericalSolution struct {
	SolutionData
	Error float64 `json:"err"`
}

//real function
type RF func(float64) float64

type EPSRF func(eps float64) RF

type ODE struct {
	A    RF
	B    RF
	F    RF
	Eta1 float64
	Eta2 float64
	Xi1  float64
	Xi2  float64
	Phi1 RF
	Phi2 RF
}
