package main

import "fmt"

func main() {
	x := make([]string, 50, 51)
	x[0] = "Alabama"
	x[1] = "Alaska"
	x[2] = "Arizona"
	x[3] = "Arkansas"
	x[4] = "California"
	x[5] = "Colorado"
	x[6] = "Connecticut"
	x[7] = "Delaware"
	x[8] = "Florida"
	x[9] = "Georgia"
	x[10] = "Hawaii"
	x[11] = "Idaho"
	x[12] = "Illinois"
	x[13] = "Indiana"
	x[14] = "Iowa"
	x[15] = "Kansas"
	x[16] = "Kentucky"
	x[17] = "Louisiana"
	x[18] = "Maine"
	x[19] = "Maryland"
	x[20] = "Massachusetts"
	x[21] = "Michigan"
	x[22] = "Minnesota"
	x[23] = "Mississippi"
	x[24] = "Missouri"
	x[25] = "Montana"
	x[26] = "Nebraska"
	x[27] = "Nevada"
	x[28] = "New Hampshire"
	x[29] = "New Jersey"
	x[30] = "New Mexico"
	x[31] = "New York"
	x[32] = "North Carolina"
	x[33] = "North Dakota"
	x[34] = "Ohio"
	x[35] = "Oklahoma"
	x[36] = "Oregon"
	x[37] = "Pennsylvania"
	x[38] = "Rhode Island"
	x[39] = "South Carolina"
	x[40] = "South Dakota"
	x[41] = "Tennessee"
	x[42] = "Texas"
	x[43] = "Utah"
	x[44] = "Vermont"
	x[45] = "Virginia"
	x[46] = "Washington"
	x[47] = "West Virginia"
	x[48] = "Wisconsin"
	x[49] = "Wyoming"
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)
}
