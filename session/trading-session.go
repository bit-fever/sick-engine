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

package session

import (
	"github.com/bit-fever/sick-engine/core"
	"time"
)

//=============================================================================

type TradingSession struct {
	Days []*SessionDay `json:"days"`
}

//=============================================================================

type SessionDay struct {
	Day    int        `json:"day"`
	Start  *core.Time `json:"start"`
	End    *core.Time `json:"end"`
	Pauses []*Pause   `json:"pauses"`
}

//=============================================================================

type Pause struct {
	From *core.Time `json:"from"`
	To   *core.Time `json:"to"`
}

//=============================================================================
//===
//=== TradingSession
//===
//=============================================================================

func (td *TradingSession) IsStartOfSession(t time.Time, timeframe int) bool {
	h,m,_ := t.Clock()
	time  := core.NewTime(h,m).Sub(0, timeframe)
	dow   := int(t.Weekday())

	for _, sd := range td.Days {
		if sd.isStartOfSession(dow, time) {
			return true
		}
	}

	return false
}

//=============================================================================
//===
//=== SessionDay
//===
//=============================================================================

func (sd *SessionDay) isStartOfSession(dow int, t *core.Time) bool {
	if dow != sd.Day {
		return false
	}

	return sd.Start == t
}

//=============================================================================
