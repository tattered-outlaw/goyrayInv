package internal

import "math"

type Pattern interface {
	getCommonState() *PatternCommonState
	localPatternAt(Tuple) Color
}

type PatternCommonState struct {
	transformation *Matrix4x4
}

func newPatternCommonState() *PatternCommonState {
	return &PatternCommonState{transformation: newIdentity4()}
}

type BlankPattern struct {
	commonState *PatternCommonState
	color       Color
}

func newBlankPattern(color Color) *BlankPattern {
	return &BlankPattern{color: color, commonState: newPatternCommonState()}
}

func (p *BlankPattern) getCommonState() *PatternCommonState {
	return p.commonState
}

func (p *BlankPattern) localPatternAt(_ Tuple) Color {
	return p.color
}

type CheckerPattern struct {
	commonState *PatternCommonState
	even, odd   Color
}

func newCheckerPattern(even, odd Color) *CheckerPattern {
	return &CheckerPattern{even: even, odd: odd, commonState: newPatternCommonState()}
}

func (p *CheckerPattern) getCommonState() *PatternCommonState {
	return p.commonState
}

func (p *CheckerPattern) localPatternAt(point Tuple) Color {
	x := math.Round(point[0])
	y := math.Round(point[1])
	z := math.Round(point[2])
	if (int(math.Abs(x))+int(math.Abs(y))+int(math.Abs(z)))%2 == 0 {
		return p.even
	}
	return p.odd
}
