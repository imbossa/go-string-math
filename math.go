package go_string_math

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Add(a, b any) (string, error) {
	aString, err := verify(a)
	if err != nil {
		return "", err
	}
	bString, err := verify(b)
	if err != nil {
		return "", err
	}

	aSplit := strings.Split(aString, ".")
	bSplit := strings.Split(bString, ".")

	if bString[0] == '-' {
		bString = bString[1:]
		return Subtract(aString, bString)
	}

	if aString[0] == '-' && bString[0] != '-' {
		return Subtract(bString, aString[1:])
	}

	if aString[0] == '-' && bString[0] == '-' {
		aString = aString[1:]
		bString = bString[1:]
		sum, err := Add(aString, bString)
		if err != nil {
			return "", err
		}

		return "-" + sum, nil
	}

	sumInt, sumFrac := "", ""
	isProgress := false

	if len(aSplit) == 2 || len(bSplit) == 2 {
		if len(aSplit) != len(bSplit) {
			if len(aSplit) > len(bSplit) {
				sumFrac = aSplit[1]
			} else {
				sumFrac = bSplit[1]
			}
		} else {
			aSplitFracLen := len(aSplit[1])
			bSplitFracLen := len(bSplit[1])

			abs := math.Abs(float64(aSplitFracLen - bSplitFracLen))

			if abs > 0 {
				if aSplitFracLen > bSplitFracLen {
					bSplit[1] = bSplit[1] + strings.Repeat("0", int(abs))
				}
				if aSplitFracLen < bSplitFracLen {
					aSplit[1] = aSplit[1] + strings.Repeat("0", int(abs))
				}
			}

			aSplitFracLen = len(aSplit[1])
			for i := aSplitFracLen; i > 0; i-- {

				aSplitFracCurrent, _ := strconv.Atoi(string(aSplit[1][i-1]))
				bSplitFracCurrent, _ := strconv.Atoi(string(bSplit[1][i-1]))

				currentSum := aSplitFracCurrent + bSplitFracCurrent
				if isProgress {
					currentSum++
					isProgress = false
				}

				if currentSum > 9 {
					currentSum = currentSum - 10
					isProgress = true
				}

				sumFrac = strconv.Itoa(currentSum) + sumFrac
			}
		}
	}

	{
		aSplitIntLen := len(aSplit[0])
		bSplitIntLen := len(bSplit[0])

		abs := math.Abs(float64(aSplitIntLen - bSplitIntLen))
		if abs > 0 {
			if aSplitIntLen > bSplitIntLen {
				bSplit[0] = strings.Repeat("0", int(abs)) + bSplit[0]
			}
			if aSplitIntLen < bSplitIntLen {
				aSplit[0] = strings.Repeat("0", int(abs)) + aSplit[0]
			}
		}

		aSplitIntLen = len(aSplit[0])
		for i := aSplitIntLen; i > 0; i-- {

			aSplitIntCurrent, _ := strconv.Atoi(string(aSplit[0][i-1]))
			bSplitIntCurrent, _ := strconv.Atoi(string(bSplit[0][i-1]))

			currentSum := aSplitIntCurrent + bSplitIntCurrent
			if isProgress {
				currentSum++
				isProgress = false
			}

			if currentSum > 9 {
				currentSum = currentSum - 10
				isProgress = true
			}

			sumInt = strconv.Itoa(currentSum) + sumInt
		}

		if isProgress {
			sumInt = "1" + sumInt
		}
	}

	if sumInt == "" {
		sumInt = "0"
	}
	if sumFrac == "" {
		sumFrac = "0"
	}

	return strings.TrimRight(strings.TrimRight(sumInt+"."+sumFrac, "0"), "."), nil
}

func Subtract(a, b any) (string, error) {
	aString, err := verify(a)
	if err != nil {
		return "", err
	}
	bString, err := verify(b)
	if err != nil {
		return "", err
	}

	if aString == bString {
		if aString[0] == '-' {
			add, err := Add(aString, bString)
			if err != nil {
				return "", err
			}
			return "-" + add, nil
		} else {
			return "0", nil
		}
	}

	if bString[0] == '-' {
		bString = bString[1:]
		return Add(aString, bString)
	}

	if aString[0] == '-' && bString[0] != '-' {
		aString = aString[1:]
		add, err := Add(aString, bString)
		if err != nil {
			return "", err
		}
		return "-" + add, nil
	}

	aSplit := strings.Split(aString, ".")
	bSplit := strings.Split(bString, ".")

	sumInt, sumFrac := "", ""
	isProgress := false

	if len(aSplit) == 2 || len(bSplit) == 2 {
		if len(aSplit) > len(bSplit) {
			sumFrac = aSplit[1]
		} else {
			if len(aSplit) < len(bSplit) {
				aSplit[1] = aSplit[1] + strings.Repeat("0", len(bSplit[1])-len(aSplit[1]))
			}

			for i := len(bSplit[1]); i > 0; i-- {
				aSplitFracCurrent, _ := strconv.Atoi(string(aSplit[1][i-1]))
				bSplitFracCurrent, _ := strconv.Atoi(string(bSplit[1][i-1]))

				currentSum := aSplitFracCurrent - bSplitFracCurrent

				if isProgress {
					currentSum--
					isProgress = false
				}

				if currentSum < 0 {
					currentSum = currentSum + 10
					isProgress = true
				}

				sumFrac = strconv.Itoa(currentSum) + sumFrac
			}
		}
	}

	{
		aSplitIntLen := len(aSplit[0])
		bSplitIntLen := len(bSplit[0])

		if aSplitIntLen < bSplitIntLen {
			aSplit[0] = strings.Repeat("0", len(bSplit[0])-len(aSplit[0])) + aSplit[0]
		}

		for i := len(aSplit[0]); i > 0; i-- {
			aSplitIntCurrent, _ := strconv.Atoi(string(aSplit[0][i-1]))
			bSplitIntCurrent, _ := strconv.Atoi(string(bSplit[0][i-1]))

			currentSum := aSplitIntCurrent - bSplitIntCurrent

			if isProgress {
				currentSum--
				isProgress = false
			}

			if currentSum < 0 {
				currentSum = currentSum + 10
				isProgress = true
			}

			sumInt = strconv.Itoa(currentSum) + sumInt
		}
	}

	if sumInt == "" {
		sumInt = "0"
	}
	if sumFrac == "" {
		sumFrac = "0"
	}

	value := strings.TrimRight(strings.TrimRight(sumInt+"."+sumFrac, "0"), ".")
	if isProgress {
		value = "-" + value
	}

	return value, nil
}

func verify(params any) (string, error) {
	var res string

	switch val := params.(type) {
	case int, int8, int16, int32, int64:
		res = fmt.Sprintf("%d", val)
	case float32, float64:
		res = fmt.Sprintf("%f", val)
	case string:
		regex := regexp.MustCompile("^-?[0-9]+(\\.[0-9]+)?$")
		if regex.MatchString(val) {
			res = val
		} else {
			return "", errors.New("invalid string")
		}
	default:
		return "", errors.New("unsupported type")
	}

	res = strings.TrimRight(strings.TrimRight(res, "0"), ".")
	return res, nil
}
