// This program parsed my spotty weight data and then outputs it in json for
// the fitbit api. It also converts the date to yyyy-mm-dd format. Instead of
// writing to a file, I print on stdout and write stdout to a file.
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// https://dev.fitbit.com/build/reference/web-api/body/#log-weight
type weightData struct {
	Date   string  `json:"date"`
	Weight float64 `json:"weight"`
}

func main() {
	// value convert float to int to avoid float math
	const multiplier = 100.0
	weightMap := make(map[string]int)
	weightSlice := strings.Split(weight, "\n")
	keys := []string{}
	for _, wl := range weightSlice {
		strs := strings.Split(wl, " ")
		d, wstr := strs[0], strs[1]
		wfl, _ := strconv.ParseFloat(wstr, 64)
		wfl = wfl * multiplier
		w := int(wfl)
		cw, exists := weightMap[d]
		if exists {
			//fmt.Println()
			//fmt.Printf("%v %v\n", w, cw)
			w = (cw + w) / 2
			//fmt.Println(w)
		} else {
			keys = append(keys, d)
		}
		weightMap[d] = w
	}

	weightDataSlice := []weightData{}
	for _, d := range keys {
		w := float64(weightMap[d]) / multiplier
		date, _ := time.Parse("01/02/2006", d)
		dStr := date.Format("2006-01-02")

		fmt.Printf("%v %.1f\n", dStr, w)
		wd := weightData{Date: dStr, Weight: math.Round(w*10) / 10}
		j, _ := json.Marshal(wd)
		weightDataSlice = append(weightDataSlice, wd)
		_ = j
		//fmt.Println(string(j))
	}

}

// Sadly my weight data from joes goals did sometimes had off by one errors in
// the dates. Probably related to UTC issues. Unfortunately the export only has
// the dates, no times. As such, I decided to just average any duplicate dates,
// and import those to fitbit. Sure, I won't have as many data points, though
// they will still be close to correct.
const weight = ` 12/12/2013 166.6
12/12/2013 168.4
12/12/2013 170.2
12/12/2013 167.6
12/13/2013 167.6
12/15/2013 166.0
12/15/2013 165.0
12/16/2013 165.2
12/17/2013 165.0
12/18/2013 167.8
12/19/2013 166.0
12/21/2013 164.4
12/22/2013 165.4
12/22/2013 163
12/23/2013 162.8
12/25/2013 163.6
12/25/2013 163.0
12/26/2013 164.0
12/28/2013 164.4
12/30/2013 163
12/30/2013 162.4
12/30/2013 165.0
01/01/2014 164.0
01/01/2014 162.4
01/02/2014 165.8
01/03/2014 165.4
01/07/2014 166.0
01/07/2014 165.8
01/09/2014 166.0
01/10/2014 166.0
01/11/2014 164.8
01/16/2014 164.8
01/16/2014 165.8
01/17/2014 161.8
01/17/2014 162.6
01/23/2014 161.2
01/26/2014 160.0
01/27/2014 162.0
01/28/2014 163.8
01/30/2014 163.0
02/02/2014 160
02/02/2014 162.2
02/04/2014 161
02/04/2014 162.2
02/05/2014 161.2
02/06/2014 160.8
02/11/2014 163.0
02/11/2014 164.4
02/11/2014 162.2
02/12/2014 160.8
02/13/2014 162.0
02/19/2014 161.2
02/19/2014 158.8
02/21/2014 158.8
02/22/2014 161.0
02/24/2014 158.8
02/25/2014 160.0
02/28/2014 161.0
02/28/2014 161.4
03/01/2014 163.0
03/02/2014 160.8
03/06/2014 160
03/06/2014 159.2
03/07/2014 159.4
03/10/2014 160.4
03/10/2014 159.6
03/11/2014 160.8
03/12/2014 160.8
03/13/2014 160.2
03/15/2014 158.0
03/17/2014 158.0
03/17/2014 162
03/18/2014 159
03/19/2014 158.2
03/20/2014 159.8
03/21/2014 158
03/22/2014 158.9
03/24/2014 159
03/24/2014 160.4
03/25/2014 160.4
03/26/2014 159.2
03/27/2014 158.0
03/28/2014 157.8
03/31/2014 160.2
03/31/2014 160
04/01/2014 163
04/02/2014 161.3
04/03/2014 161.4
04/04/2014 161.3
04/08/2014 161.4
04/08/2014 158.6
04/09/2014 160.2
04/09/2014 160.4
04/10/2014 158.6
04/12/2014 158.2
04/12/2014 160.4
04/14/2014 157.8
04/15/2014 159.2
04/16/2014 160.6
04/18/2014 160.6
04/18/2014 159.4
04/20/2014 157.8
04/21/2014 160
04/21/2014 158.8
04/23/2014 158.0
04/25/2014 158.8
04/25/2014 160.0
04/29/2014 157.8
04/29/2014 160
04/30/2014 158.8
05/02/2014 161.2
05/05/2014 161.2
05/06/2014 158.2
05/08/2014 157.8
05/08/2014 159.2
05/10/2014 158.6
05/11/2014 160.2
05/12/2014 160.2
05/13/2014 157.4
05/14/2014 157.2
05/20/2014 156
05/22/2014 158.6
05/22/2014 158.5
05/23/2014 159.6
05/27/2014 156.6
05/29/2014 157.0
05/29/2014 157.8
06/08/2014 159.6
06/08/2014 158.6
06/09/2014 159.7
06/10/2014 158.7
06/11/2014 158.2
06/12/2014 158.2
06/13/2014 160.2
06/14/2014 159.8
06/15/2014 157.2
06/17/2014 155.6
06/17/2014 157.8
06/19/2014 158.2
06/20/2014 159
06/21/2014 159.2
06/22/2014 157.6
06/23/2014 160.0
06/24/2014 159.0
06/25/2014 159.0
06/26/2014 157.0
06/27/2014 157.0
06/30/2014 158.8
07/01/2014 160.4
07/02/2014 159.6
07/03/2014 161
07/07/2014 163
07/08/2014 163.4
07/09/2014 159.0
07/10/2014 159.4
07/11/2014 157
07/12/2014 157
07/14/2014 157.2
07/15/2014 157.2
07/16/2014 157.0
07/17/2014 158.8
07/22/2014 161
07/22/2014 160.2
07/24/2014 160
07/29/2014 160
07/31/2014 158.8
07/31/2014 157.2
08/02/2014 157.4
08/02/2014 158.2
08/03/2014 158.2
08/04/2014 158.2
08/06/2014 160.2
08/06/2014 161.2
08/08/2014 159
08/13/2014 159.8
08/13/2014 160
08/13/2014 157
08/14/2014 157.4
08/17/2014 158.6
08/19/2014 155.2
08/19/2014 157.4
08/23/2014 155.8
09/17/2014 159.6
09/18/2014 158.2
09/19/2014 159.2
09/22/2014 158.0
09/23/2014 160
09/24/2014 160.0
09/25/2014 158.6
09/26/2014 158.6
09/27/2014 158.0
09/28/2014 158.2
09/29/2014 161.0
09/30/2014 162.2
10/02/2014 160.2
10/02/2014 160
10/07/2014 161
10/07/2014 161.2
10/10/2014 161.2
10/13/2014 160.8
10/15/2014 160.8
10/15/2014 160
10/16/2014 161.8
10/17/2014 160.2
10/19/2014 161.8
10/20/2014 161.6
10/23/2014 160.6
10/23/2014 161.2
10/27/2014 159.4
10/30/2014 159.8
10/30/2014 160.8
11/03/2014 165
11/04/2014 163.2
11/05/2014 160
11/07/2014 160
11/07/2014 162
11/10/2014 161.4
11/13/2014 160.8
11/14/2014 163.4
11/18/2014 159.2
11/18/2014 161
11/20/2014 161
11/20/2014 164
11/21/2014 164.4
11/24/2014 163.6
11/25/2014 163.8
11/25/2014 166
11/26/2014 163.8
11/27/2014 161.6
11/28/2014 162.4
12/01/2014 161.8
12/02/2014 159.8
12/06/2014 162.8
12/09/2014 164.0
12/20/2014 163.8
12/22/2014 162
12/25/2014 163
12/25/2014 162
12/25/2014 160
12/28/2014 160
12/28/2014 161
01/07/2015 166.0
01/08/2015 166.0
01/13/2015 163.6
01/14/2015 164.6
01/16/2015 160.4
01/26/2015 162.4
02/08/2015 159.2
02/09/2015 158.6
02/11/2015 157.6
02/11/2015 156.6
02/13/2015 156.8
02/16/2015 161.4
02/17/2015 157
02/20/2015 157.6
02/20/2015 157
02/22/2015 155
02/24/2015 157.2
02/26/2015 155.8
02/27/2015 155.0
02/27/2015 155.8
02/28/2015 157.2
03/02/2015 155.6
03/02/2015 154.6
03/03/2015 152.8
03/04/2015 154.6
03/05/2015 155.2
03/07/2015 155.2
03/07/2015 154.2
03/09/2015 152.8
03/10/2015 153.4
03/15/2015 153.4
03/17/2015 159
03/23/2015 155
03/23/2015 152.2
03/28/2015 153.2
03/30/2015 154.6
03/30/2015 151.4
04/03/2015 151.0
04/05/2015 150.6
04/06/2015 148.8
04/07/2015 151.0
04/10/2015 148.2
04/11/2015 148.0
04/15/2015 150.6
04/15/2015 150
04/17/2015 150.2
04/20/2015 150
04/25/2015 148.8
04/27/2015 153.6
04/29/2015 154.6
05/02/2015 151.4
05/04/2015 151.4
05/05/2015 151.8
05/07/2015 151.0
05/08/2015 150
05/11/2015 153.2
05/14/2015 154.2
05/17/2015 155
05/18/2015 151.4
05/22/2015 150.6
05/25/2015 151.4
05/26/2015 151.6
05/27/2015 153.4
05/31/2015 155.0
05/31/2015 153.3
06/01/2015 153.4
06/04/2015 155.6
06/09/2015 155
06/22/2015 157
06/23/2015 156.8
09/09/2015 152.5
09/14/2015 155.2
09/15/2015 155.2
09/16/2015 155.8
09/18/2015 154.4
09/19/2015 156.6
09/21/2015 156.2
09/26/2015 153.4
09/28/2015 154.2
09/29/2015 153.2
10/01/2015 154.2
10/04/2015 154.6
10/07/2015 151.8
10/07/2015 153.0
10/08/2015 154.2
10/09/2015 154.6
10/11/2015 152.8
10/12/2015 152.8
10/16/2015 152.8
10/16/2015 151.8`
