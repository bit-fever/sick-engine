//=============================================================================
/*
Copyright © 2024 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package core

import (
	"errors"
	"strconv"
	"strings"
)

//=============================================================================

type Time struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
}

//=============================================================================

func NewTime(hour int, min int) *Time {
	return &Time{
		Hour : hour,
		Min : min,
	}
}

//=============================================================================

func NewTimeFromString(time string) (*Time, error) {
	index := strings.Index(time, ":")
	if index <1 || index>2 || len(time)<4 || len(time)>5 {
		return nil, errors.New("bad time format")
	}

	hour,err1 := strconv.Atoi(time[0:index])
	mins,err2 := strconv.Atoi(time[index:])

	if err1 != nil {
		return nil, errors.New("time hour is not an integer")
	}

	if err2 != nil {
		return nil, errors.New("time minute is not an integer")
	}

	if hour<0 || hour>23 || mins<0 || mins>59 {
		return nil, errors.New("bad hour/minute value")
	}

	return &Time{
		Hour: hour,
		Min : mins,
	}, nil
}

//=============================================================================

func (t *Time) Add(hour,mins int) *Time {
	h := t.Hour + hour
	m := t.Min  + mins

	if m >= 60 {
		m -= 60
		h++
	}

	if h >= 24 {
		h -= 24
	}

	return &Time{
		Hour: h,
		Min : m,
	}
}

//=============================================================================

func (t *Time) Sub(hour,mins int) *Time {
	h := t.Hour - hour
	m := t.Min  - mins

	if m < 0 {
		m += 60
		h--
	}

	if h < 0 {
		h += 24
	}

	return &Time{
		Hour: h,
		Min : m,
	}
}

//=============================================================================
