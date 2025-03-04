package constants

type Country struct {
	Name       string
	Code       string
	Population int64
	Area       int64
}

var Countries = []Country{
	{Name: "Afghanistan", Code: "AFG", Population: 29121286, Area: 647500},
	{Name: "Albania", Code: "ALB", Population: 2986952, Area: 28748},
	{Name: "Algeria", Code: "DZA", Population: 34586184, Area: 2381740},
	{Name: "American Samoa", Code: "ASM", Population: 57881, Area: 199},
	{Name: "Andorra", Code: "AND", Population: 84000, Area: 468},
	{Name: "Angola", Code: "AGO", Population: 13068161, Area: 1246700},
	{Name: "Anguilla", Code: "AIA", Population: 13254, Area: 102},
	{Name: "Antarctica", Code: "ATA", Population: 0, Area: 14000000},
	{Name: "Antigua and Barbuda", Code: "ATG", Population: 86754, Area: 443},
	{Name: "Argentina", Code: "ARG", Population: 41343201, Area: 2766890},
	{Name: "Armenia", Code: "ARM", Population: 2968000, Area: 29800},
	{Name: "Aruba", Code: "ABW", Population: 71566, Area: 193},
	{Name: "Australia", Code: "AUS", Population: 21515754, Area: 7686850},
	{Name: "Austria", Code: "AUT", Population: 8205000, Area: 83858},
	{Name: "Azerbaijan", Code: "AZE", Population: 8303512, Area: 86600},
	{Name: "Bahamas", Code: "BHS", Population: 301790, Area: 13940},
	{Name: "Bahrain", Code: "BHR", Population: 738004, Area: 665},
	{Name: "Bangladesh", Code: "BGD", Population: 156118464, Area: 144000},
	{Name: "Barbados", Code: "BRB", Population: 285653, Area: 431},
	{Name: "Belarus", Code: "BLR", Population: 9685000, Area: 207600},
	{Name: "Belgium", Code: "BEL", Population: 10403000, Area: 30510},
	{Name: "Belize", Code: "BLZ", Population: 314522, Area: 22966},
	{Name: "Benin", Code: "BEN", Population: 9056010, Area: 112620},
	{Name: "Bermuda", Code: "BMU", Population: 65365, Area: 53},
	{Name: "Bhutan", Code: "BTN", Population: 699847, Area: 47000},
	{Name: "Bolivia", Code: "BOL", Population: 9947418, Area: 1098580},
	{Name: "Bosnia and Herzegovina", Code: "BIH", Population: 4590000, Area: 51129},
	{Name: "Botswana", Code: "BWA", Population: 2029307, Area: 600370},
	{Name: "Brazil", Code: "BRA", Population: 201103330, Area: 8511965},
	{Name: "British Indian Ocean Territory", Code: "IOT", Population: 4000, Area: 60},
	{Name: "British Virgin Islands", Code: "VGB", Population: 21730, Area: 153},
	{Name: "Brunei", Code: "BRN", Population: 395027, Area: 5770},
	{Name: "Bulgaria", Code: "BGR", Population: 7148785, Area: 110910},
	{Name: "Burkina Faso", Code: "BFA", Population: 16241811, Area: 274200},
	{Name: "Burundi", Code: "BDI", Population: 9863117, Area: 27830},
	{Name: "Cambodia", Code: "KHM", Population: 14453680, Area: 181040},
	{Name: "Cameroon", Code: "CMR", Population: 19294149, Area: 475440},
	{Name: "Canada", Code: "CAN", Population: 33679000, Area: 9984670},
	{Name: "Cape Verde", Code: "CPV", Population: 508659, Area: 4033},
	{Name: "Cayman Islands", Code: "CYM", Population: 44270, Area: 262},
	{Name: "Central African Republic", Code: "CAF", Population: 4844927, Area: 622984},
	{Name: "Chad", Code: "TCD", Population: 10543464, Area: 1284000},
	{Name: "Chile", Code: "CHL", Population: 16746491, Area: 756950},
	{Name: "China", Code: "CHN", Population: 1330044000, Area: 9596960},
	{Name: "Christmas Island", Code: "CXR", Population: 1500, Area: 135},
	{Name: "Cocos Islands", Code: "CCK", Population: 628, Area: 14},
	{Name: "Colombia", Code: "COL", Population: 47790000, Area: 1138910},
	{Name: "Comoros", Code: "COM", Population: 773407, Area: 2170},
	{Name: "Cook Islands", Code: "COK", Population: 21388, Area: 240},
	{Name: "Costa Rica", Code: "CRI", Population: 4516220, Area: 51100},
	{Name: "Croatia", Code: "HRV", Population: 4491000, Area: 56542},
	{Name: "Cuba", Code: "CUB", Population: 11423000, Area: 110860},
	{Name: "Curacao", Code: "CUW", Population: 141766, Area: 444},
	{Name: "Cyprus", Code: "CYP", Population: 1102677, Area: 9250},
	{Name: "Czech Republic", Code: "CZE", Population: 10476000, Area: 78866},
	{Name: "Democratic Republic of the Congo", Code: "COD", Population: 70916439, Area: 2345410},
	{Name: "Denmark", Code: "DNK", Population: 5484000, Area: 43094},
	{Name: "Djibouti", Code: "DJI", Population: 740528, Area: 23000},
	{Name: "Dominica", Code: "DMA", Population: 72813, Area: 754},
	{Name: "Dominican Republic", Code: "DOM", Population: 9823821, Area: 48730},
	{Name: "East Timor", Code: "TLS", Population: 1154625, Area: 15007},
	{Name: "Ecuador", Code: "ECU", Population: 14790608, Area: 283560},
	{Name: "Egypt", Code: "EGY", Population: 80471869, Area: 1001450},
	{Name: "El Salvador", Code: "SLV", Population: 6052064, Area: 21040},
	{Name: "Equatorial Guinea", Code: "GNQ", Population: 1014999, Area: 28051},
	{Name: "Eritrea", Code: "ERI", Population: 5792984, Area: 121320},
	{Name: "Estonia", Code: "EST", Population: 1291170, Area: 45226},
	{Name: "Ethiopia", Code: "ETH", Population: 88013491, Area: 1127127},
	{Name: "Falkland Islands", Code: "FLK", Population: 2638, Area: 12173},
	{Name: "Faroe Islands", Code: "FRO", Population: 48228, Area: 1399},
	{Name: "Fiji", Code: "FJI", Population: 875983, Area: 18270},
	{Name: "Finland", Code: "FIN", Population: 5244000, Area: 337030},
	{Name: "France", Code: "FRA", Population: 64768389, Area: 547030},
	{Name: "French Polynesia", Code: "PYF", Population: 270485, Area: 4167},
	{Name: "Gabon", Code: "GAB", Population: 1545255, Area: 267667},
	{Name: "Gambia", Code: "GMB", Population: 1593256, Area: 11300},
	{Name: "Georgia", Code: "GEO", Population: 4630000, Area: 69700},
	{Name: "Germany", Code: "DEU", Population: 81802257, Area: 357021},
	{Name: "Ghana", Code: "GHA", Population: 24339838, Area: 239460},
	{Name: "Gibraltar", Code: "GIB", Population: 27884, Area: 7},
	{Name: "Greece", Code: "GRC", Population: 11000000, Area: 131940},
	{Name: "Greenland", Code: "GRL", Population: 56375, Area: 2166086},
	{Name: "Grenada", Code: "GRD", Population: 107818, Area: 344},
	{Name: "Guam", Code: "GUM", Population: 159358, Area: 549},
	{Name: "Guatemala", Code: "GTM", Population: 13550440, Area: 108890},
	{Name: "Guernsey", Code: "GGY", Population: 65228, Area: 78},
	{Name: "Guinea", Code: "GIN", Population: 10324025, Area: 245857},
	{Name: "Guinea-Bissau", Code: "GNB", Population: 1565126, Area: 36120},
	{Name: "Guyana", Code: "GUY", Population: 748486, Area: 214970},
	{Name: "Haiti", Code: "HTI", Population: 9648924, Area: 27750},
	{Name: "Honduras", Code: "HND", Population: 7989415, Area: 112090},
	{Name: "Hong Kong", Code: "HKG", Population: 6898686, Area: 1092},
	{Name: "Hungary", Code: "HUN", Population: 9982000, Area: 93030},
	{Name: "Iceland", Code: "ISL", Population: 308910, Area: 103000},
	{Name: "India", Code: "IND", Population: 1173108018, Area: 3287590},
	{Name: "Indonesia", Code: "IDN", Population: 242968342, Area: 1919440},
	{Name: "Iran", Code: "IRN", Population: 76923300, Area: 1648000},
	{Name: "Iraq", Code: "IRQ", Population: 29671605, Area: 437072},
	{Name: "Ireland", Code: "IRL", Population: 4622917, Area: 70280},
	{Name: "Isle of Man", Code: "IMN", Population: 75049, Area: 572},
	{Name: "Israel", Code: "ISR", Population: 7353985, Area: 20770},
	{Name: "Italy", Code: "ITA", Population: 60340328, Area: 301230},
	{Name: "Ivory Coast", Code: "CIV", Population: 21058798, Area: 322460},
	{Name: "Jamaica", Code: "JAM", Population: 2847232, Area: 10991},
	{Name: "Japan", Code: "JPN", Population: 127288000, Area: 377835},
	{Name: "Jersey", Code: "JEY", Population: 90812, Area: 116},
	{Name: "Jordan", Code: "JOR", Population: 6407085, Area: 92300},
	{Name: "Kazakhstan", Code: "KAZ", Population: 15340000, Area: 2717300},
	{Name: "Kenya", Code: "KEN", Population: 40046566, Area: 582650},
	{Name: "Kiribati", Code: "KIR", Population: 92533, Area: 811},
	{Name: "Kosovo", Code: "XKX", Population: 1800000, Area: 10887},
	{Name: "Kuwait", Code: "KWT", Population: 2789132, Area: 17820},
	{Name: "Kyrgyzstan", Code: "KGZ", Population: 5508626, Area: 198500},
	{Name: "Laos", Code: "LAO", Population: 6368162, Area: 236800},
	{Name: "Latvia", Code: "LVA", Population: 2217969, Area: 64589},
	{Name: "Lebanon", Code: "LBN", Population: 4125247, Area: 10400},
	{Name: "Lesotho", Code: "LSO", Population: 1919552, Area: 30355},
	{Name: "Liberia", Code: "LBR", Population: 3685076, Area: 111370},
	{Name: "Libya", Code: "LBY", Population: 6461454, Area: 1759540},
	{Name: "Liechtenstein", Code: "LIE", Population: 35000, Area: 160},
	{Name: "Lithuania", Code: "LTU", Population: 2944459, Area: 65200},
	{Name: "Luxembourg", Code: "LUX", Population: 497538, Area: 2586},
	{Name: "Macau", Code: "MAC", Population: 449198, Area: 254},
	{Name: "Macedonia", Code: "MKD", Population: 2062294, Area: 25333},
	{Name: "Madagascar", Code: "MDG", Population: 21281844, Area: 587040},
	{Name: "Malawi", Code: "MWI", Population: 15447500, Area: 118480},
	{Name: "Malaysia", Code: "MYS", Population: 28274729, Area: 329750},
	{Name: "Maldives", Code: "MDV", Population: 395650, Area: 300},
	{Name: "Mali", Code: "MLI", Population: 13796354, Area: 1240000},
	{Name: "Malta", Code: "MLT", Population: 403000, Area: 316},
	{Name: "Marshall Islands", Code: "MHL", Population: 65859, Area: 181},
	{Name: "Mauritania", Code: "MRT", Population: 3205060, Area: 1030700},
	{Name: "Mauritius", Code: "MUS", Population: 1294104, Area: 2040},
	{Name: "Mexico", Code: "MEX", Population: 112468855, Area: 1972550},
	{Name: "Micronesia", Code: "FSM", Population: 107708, Area: 702},
	{Name: "Moldova", Code: "MDA", Population: 4324000, Area: 33843},
	{Name: "Monaco", Code: "MCO", Population: 32965, Area: 2},
	{Name: "Mongolia", Code: "MNG", Population: 3086918, Area: 1565000},
	{Name: "Montenegro", Code: "MNE", Population: 666730, Area: 14026},
	{Name: "Morocco", Code: "MAR", Population: 31627428, Area: 446550},
	{Name: "Mozambique", Code: "MOZ", Population: 22061451, Area: 801590},
	{Name: "Myanmar", Code: "MMR", Population: 53414374, Area: 678500},
	{Name: "Namibia", Code: "NAM", Population: 2128471, Area: 825418},
	{Name: "Nauru", Code: "NRU", Population: 10065, Area: 21},
	{Name: "Nepal", Code: "NPL", Population: 28951852, Area: 140800},
	{Name: "Netherlands", Code: "NLD", Population: 16645000, Area: 41526},
	{Name: "New Caledonia", Code: "NCL", Population: 216494, Area: 19060},
	{Name: "New Zealand", Code: "NZL", Population: 4252277, Area: 268680},
	{Name: "Nicaragua", Code: "NIC", Population: 5995928, Area: 129494},
	{Name: "Niger", Code: "NER", Population: 15878271, Area: 1267000},
	{Name: "Nigeria", Code: "NGA", Population: 154000000, Area: 923768},
	{Name: "North Korea", Code: "PRK", Population: 22912177, Area: 120540},
	{Name: "Norway", Code: "NOR", Population: 5009150, Area: 324220},
	{Name: "Oman", Code: "OMN", Population: 2967717, Area: 212460},
	{Name: "Pakistan", Code: "PAK", Population: 184404791, Area: 803940},
	{Name: "Palau", Code: "PLW", Population: 19907, Area: 458},
	{Name: "Palestine", Code: "PSE", Population: 3800000, Area: 5970},
	{Name: "Panama", Code: "PAN", Population: 3410676, Area: 78200},
	{Name: "Papua New Guinea", Code: "PNG", Population: 6064515, Area: 462840},
	{Name: "Paraguay", Code: "PRY", Population: 6375830, Area: 406750},
	{Name: "Peru", Code: "PER", Population: 29907003, Area: 1285220},
	{Name: "Philippines", Code: "PHL", Population: 99900177, Area: 300000},
	{Name: "Poland", Code: "POL", Population: 38500000, Area: 312685},
	{Name: "Portugal", Code: "PRT", Population: 10676000, Area: 92391},
	{Name: "Puerto Rico", Code: "PRI", Population: 3916632, Area: 9104},
	{Name: "Qatar", Code: "QAT", Population: 840926, Area: 11437},
	{Name: "Republic of the Congo", Code: "COG", Population: 3039126, Area: 342000},
	{Name: "Romania", Code: "ROU", Population: 21959278, Area: 237500},
	{Name: "Russia", Code: "RUS", Population: 140702000, Area: 17100000},
	{Name: "Rwanda", Code: "RWA", Population: 11055976, Area: 26338},
	{Name: "Saint Barthelemy", Code: "BLM", Population: 8450, Area: 21},
	{Name: "Saint Helena", Code: "SHN", Population: 7460, Area: 410},
	{Name: "Saint Kitts and Nevis", Code: "KNA", Population: 51134, Area: 261},
	{Name: "Saint Lucia", Code: "LCA", Population: 160922, Area: 616},
	{Name: "Saint Martin", Code: "MAF", Population: 35925, Area: 53},
	{Name: "Saint Pierre and Miquelon", Code: "SPM", Population: 7012, Area: 242},
	{Name: "Saint Vincent and the Grenadines", Code: "VCT", Population: 104217, Area: 389},
	{Name: "Samoa", Code: "WSM", Population: 192001, Area: 2944},
	{Name: "San Marino", Code: "SMR", Population: 31477, Area: 61},
	{Name: "Sao Tome and Principe", Code: "STP", Population: 175808, Area: 1001},
	{Name: "Saudi Arabia", Code: "SAU", Population: 25731776, Area: 1960582},
	{Name: "Senegal", Code: "SEN", Population: 12323252, Area: 196190},
	{Name: "Serbia", Code: "SRB", Population: 7344847, Area: 88361},
	{Name: "Seychelles", Code: "SYC", Population: 88340, Area: 455},
	{Name: "Sierra Leone", Code: "SLE", Population: 5245695, Area: 71740},
	{Name: "Singapore", Code: "SGP", Population: 4701069, Area: 693},
	{Name: "Sint Maarten", Code: "SXM", Population: 37429, Area: 34},
	{Name: "Slovakia", Code: "SVK", Population: 5455000, Area: 48845},
	{Name: "Slovenia", Code: "SVN", Population: 2007000, Area: 20273},
	{Name: "Solomon Islands", Code: "SLB", Population: 559198, Area: 28450},
	{Name: "Somalia", Code: "SOM", Population: 10112453, Area: 637657},
	{Name: "South Africa", Code: "ZAF", Population: 49000000, Area: 1219912},
	{Name: "South Korea", Code: "KOR", Population: 48422644, Area: 98480},
	{Name: "South Sudan", Code: "SSD", Population: 8260490, Area: 644329},
	{Name: "Spain", Code: "ESP", Population: 46505963, Area: 504782},
	{Name: "Sri Lanka", Code: "LKA", Population: 21513990, Area: 65610},
	{Name: "Sudan", Code: "SDN", Population: 35000000, Area: 1861484},
	{Name: "Suriname", Code: "SUR", Population: 492829, Area: 163270},
	{Name: "Svalbard and Jan Mayen", Code: "SJM", Population: 2550, Area: 62049},
	{Name: "Swaziland", Code: "SWZ", Population: 1354051, Area: 17363},
	{Name: "Sweden", Code: "SWE", Population: 9555893, Area: 449964},
	{Name: "Switzerland", Code: "CHE", Population: 7581000, Area: 41290},
	{Name: "Syria", Code: "SYR", Population: 22198110, Area: 185180},
	{Name: "Taiwan", Code: "TWN", Population: 22894384, Area: 35980},
	{Name: "Tajikistan", Code: "TJK", Population: 7487489, Area: 143100},
	{Name: "Tanzania", Code: "TZA", Population: 41892895, Area: 945087},
	{Name: "Thailand", Code: "THA", Population: 67089500, Area: 514000},
	{Name: "Togo", Code: "TGO", Population: 6587239, Area: 56785},
	{Name: "Tokelau", Code: "TKL", Population: 1466, Area: 10},
	{Name: "Tonga", Code: "TON", Population: 122580, Area: 748},
	{Name: "Trinidad and Tobago", Code: "TTO", Population: 1228691, Area: 5128},
	{Name: "Tunisia", Code: "TUN", Population: 10589025, Area: 163610},
	{Name: "Turkey", Code: "TUR", Population: 77804122, Area: 780580},
	{Name: "Turkmenistan", Code: "TKM", Population: 4940916, Area: 488100},
	{Name: "Turks and Caicos Islands", Code: "TCA", Population: 20556, Area: 430},
	{Name: "Tuvalu", Code: "TUV", Population: 10472, Area: 26},
	{Name: "U.S. Virgin Islands", Code: "VIR", Population: 108708, Area: 352},
	{Name: "Uganda", Code: "UGA", Population: 33398682, Area: 236040},
	{Name: "Ukraine", Code: "UKR", Population: 45415596, Area: 603700},
	{Name: "United Arab Emirates", Code: "ARE", Population: 4975593, Area: 82880},
	{Name: "United Kingdom", Code: "GBR", Population: 62348447, Area: 244820},
	{Name: "United States", Code: "USA", Population: 310232863, Area: 9629091},
	{Name: "Uruguay", Code: "URY", Population: 3477000, Area: 176220},
	{Name: "Uzbekistan", Code: "UZB", Population: 27865738, Area: 447400},
	{Name: "Vanuatu", Code: "VUT", Population: 221552, Area: 12200},
	{Name: "Vatican", Code: "VAT", Population: 921, Area: 0},
	{Name: "Venezuela", Code: "VEN", Population: 27223228, Area: 912050},
	{Name: "Vietnam", Code: "VNM", Population: 89571130, Area: 329560},
	{Name: "Wallis and Futuna", Code: "WLF", Population: 16025, Area: 274},
	{Name: "Western Sahara", Code: "ESH", Population: 273008, Area: 266000},
	{Name: "Yemen", Code: "YEM", Population: 23495361, Area: 527970},
	{Name: "Zambia", Code: "ZMB", Population: 13460305, Area: 752614},
	{Name: "Zimbabwe", Code: "ZWE", Population: 11651858, Area: 390580},
}
