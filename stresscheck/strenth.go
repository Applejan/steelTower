package main

import (
	"errors"
	"fmt"
	"math"
)

//GB50017-2017强度验算
func strenth1(sec *section, f *force) (string, error) {
	epstion := math.Sqrt(235 / sec.Fy())
	val := func(sec *section) int {
		is := sec.D / sec.Thick
		switch {
		case is <= 50*math.Pow(epstion, 2):
			return 0
		case is <= 70*math.Pow(epstion, 2):
			return 1
		case is <= 90*math.Pow(epstion, 2):
			return 2
		case is <= 100*math.Pow(epstion, 2):
			return 3
		default:
			return -1
		}
	}(sec)

	gramma := func(v int) float64 {
		switch {
		case v > 2:
			return 1.0
		default:
			return 1.15
		}
	}(val)

	maxval := f.n*1e3/sec.Area() + f.m*1e6/sec.Wn()/gramma
	minval := f.n*1e3/sec.Area() - f.m*1e6/sec.Wn()/gramma
	precent := math.Max(math.Abs(maxval)/sec.Fy(), math.Abs(minval)/sec.Fy())
	if precent > 1 {
		return fmt.Sprintf("%.2f", precent), errors.New("GB50017-2017强度超限\t")
	}
	return fmt.Sprintf("%.2f", precent), nil
}

//DLT5130-2001弯曲强度计算
func strenth2(s *section, f *force) (string, error) {
	precent := f.m * 1e6 * s.C() / s.Ix() / fb(s)
	if precent <= 1 {
		return fmt.Sprintf("%.2f", precent), nil
	}
	return fmt.Sprintf("%.2f", precent), errors.New("DLT5130-2001弯曲强度计算超限\t")
}

//DLT5130-2001剪切强度计算
func strenth3(s *section, f *force) (string, error) {
	precent := f.v * 1e3 * qit(s) / (0.58 * s.Fy())
	if precent <= 1 {
		return fmt.Sprintf("%.2f", precent), nil
	}
	return fmt.Sprintf("%.2f", precent), errors.New("DLT5130-2001剪切强度超限\t")

}

//DLT5130-2001复合受力强度计算
func strenth4(s *section, f *force) (string, error) {
	tmp1 := math.Pow(f.n*1e3/s.Area()+f.m*1e6*s.C()/s.Ix(), 2)
	tmp2 := 3 * math.Pow(f.v*1e3*qit(s), 2)
	precent := (tmp1 + tmp2) / math.Pow(fb(s), 2)
	if precent <= 1 {
		return fmt.Sprintf("%.2f", precent), nil
	}
	return fmt.Sprintf("%.2f", precent), errors.New("DLT5130-2001复合受力强度\t")
}

// Strength implments the check of strength
func Strength(sections map[string]section, f *force) {
	s := sections[f.frameID]
	if pre, err := strenth1(&s, f); err != nil {
		fmt.Print(err, "应力系数：", pre, "\n")
	}
	if pre, err := strenth2(&s, f); err != nil {
		fmt.Print(err, "应力系数：", pre, "\n")
	}
	if pre, err := strenth3(&s, f); err != nil {
		fmt.Print(err, "应力系数：", pre, "\n")
	}
	if pre, err := strenth4(&s, f); err != nil {
		fmt.Print(err, "应力系数：", pre, "\n")
	}
}
