package dns

import "github.com/Kuadrant/multicluster-gateway-controller/pkg/_internal/slice"

type countryCodes struct {
	numericCode int
	countryName string
	alpha2Code  string
	alpha3Code  string
}

var iso3166Codes = []*countryCodes{
	{4, "Afghanistan", "AF", "AFG"},
	{8, "Albania", "AL", "ALB"},
	{12, "Algeria", "DZ", "DZA"},
	{16, "American Samoa", "AS", "ASM"},
	{20, "Andorra", "AD", "AND"},
	{24, "Angola", "AO", "AGO"},
	{660, "Anguilla", "AI", "AIA"},
	{10, "Antarctica", "AQ", "ATA"},
	{28, "Antigua and Barbuda", "AG", "ATG"},
	{32, "Argentina", "AR", "ARG"},
	{51, "Armenia", "AM", "ARM"},
	{533, "Aruba", "AW", "ABW"},
	{36, "Australia", "AU", "AUS"},
	{40, "Austria", "AT", "AUT"},
	{31, "Azerbaijan", "AZ", "AZE"},
	{44, "Bahamas", "BS", "BHS"},
	{48, "Bahrain", "BH", "BHR"},
	{50, "Bangladesh", "BD", "BGD"},
	{52, "Barbados", "BB", "BRB"},
	{112, "Belarus", "BY", "BLR"},
	{56, "Belgium", "BE", "BEL"},
	{84, "Belize", "BZ", "BLZ"},
	{204, "Benin", "BJ", "BEN"},
	{60, "Bermuda", "BM", "BMU"},
	{248, "Åland Islands", "AX", "ALA"},
	{64, "Bhutan", "BT", "BTN"},
	{68, "Bolivia (Plurinational State of)", "BO", "BOL"},
	{535, "Bonaire, Sint Eustatius and Saba", "BQ", "BES"},
	{70, "Bosnia and Herzegovina", "BA", "BIH"},
	{72, "Botswana", "BW", "BWA"},
	{74, "Bouvet Island", "BV", "BVT"},
	{76, "Brazil", "BR", "BRA"},
	{86, "British Indian Ocean Territory", "IO", "IOT"},
	{96, "Brunei Darussalam", "BN", "BRN"},
	{100, "Bulgaria", "BG", "BGR"},
	{854, "Burkina Faso", "BF", "BFA"},
	{108, "Burundi", "BI", "BDI"},
	{132, "Cabo Verde", "CV", "CPV"},
	{116, "Cambodia", "KH", "KHM"},
	{120, "Cameroon", "CM", "CMR"},
	{124, "Canada", "CA", "CAN"},
	{136, "Cayman Islands", "KY", "CYM"},
	{140, "Central African Republic", "CF", "CAF"},
	{148, "Chad", "TD", "TCD"},
	{152, "Chile", "CL", "CHL"},
	{156, "China", "CN", "CHN"},
	{162, "Christmas Island", "CX", "CXR"},
	{166, "Cocos (Keeling) Islands", "CC", "CCK"},
	{170, "Colombia", "CO", "COL"},
	{174, "Comoros", "KM", "COM"},
	{180, "Congo (the Democratic Republic of the)", "CD", "COD"},
	{178, "Congo", "CG", "COG"},
	{184, "Cook Islands", "CK", "COK"},
	{188, "Costa Rica", "CR", "CRI"},
	{191, "Croatia", "HR", "HRV"},
	{192, "Cuba", "CU", "CUB"},
	{531, "Curaçao", "CW", "CUW"},
	{196, "Cyprus", "CY", "CYP"},
	{203, "Czechia", "CZ", "CZE"},
	{384, "Côte d'Ivoire", "CI", "CIV"},
	{208, "Denmark", "DK", "DNK"},
	{262, "Djibouti", "DJ", "DJI"},
	{212, "Dominica", "DM", "DMA"},
	{214, "Dominican Republic", "DO", "DOM"},
	{218, "Ecuador", "EC", "ECU"},
	{818, "Egypt", "EG", "EGY"},
	{222, "El Salvador", "SV", "SLV"},
	{226, "Equatorial Guinea", "GQ", "GNQ"},
	{232, "Eritrea", "ER", "ERI"},
	{233, "Estonia", "EE", "EST"},
	{748, "Eswatini", "SZ", "SWZ"},
	{231, "Ethiopia", "ET", "ETH"},
	{238, "Falkland Islands  [Malvinas]", "FK", "FLK"},
	{234, "Faroe Islands", "FO", "FRO"},
	{242, "Fiji", "FJ", "FJI"},
	{246, "Finland", "FI", "FIN"},
	{250, "France", "FR", "FRA"},
	{254, "French Guiana", "GF", "GUF"},
	{258, "French Polynesia", "PF", "PYF"},
	{260, "French Southern Territories", "TF", "ATF"},
	{266, "Gabon", "GA", "GAB"},
	{270, "Gambia", "GM", "GMB"},
	{268, "Georgia", "GE", "GEO"},
	{276, "Germany", "DE", "DEU"},
	{288, "Ghana", "GH", "GHA"},
	{292, "Gibraltar", "GI", "GIB"},
	{300, "Greece", "GR", "GRC"},
	{304, "Greenland", "GL", "GRL"},
	{308, "Grenada", "GD", "GRD"},
	{312, "Guadeloupe", "GP", "GLP"},
	{316, "Guam", "GU", "GUM"},
	{320, "Guatemala", "GT", "GTM"},
	{831, "Guernsey", "GG", "GGY"},
	{324, "Guinea", "GN", "GIN"},
	{624, "Guinea-Bissau", "GW", "GNB"},
	{328, "Guyana", "GY", "GUY"},
	{332, "Haiti", "HT", "HTI"},
	{334, "Heard Island and McDonald Islands", "HM", "HMD"},
	{336, "Holy See", "VA", "VAT"},
	{340, "Honduras", "HN", "HND"},
	{344, "Hong Kong", "HK", "HKG"},
	{348, "Hungary", "HU", "HUN"},
	{352, "Iceland", "IS", "ISL"},
	{356, "India", "IN", "IND"},
	{360, "Indonesia", "ID", "IDN"},
	{364, "Iran (Islamic Republic of)", "IR", "IRN"},
	{368, "Iraq", "IQ", "IRQ"},
	{372, "Ireland", "IE", "IRL"},
	{833, "Isle of Man", "IM", "IMN"},
	{376, "Israel", "IL", "ISR"},
	{380, "Italy", "IT", "ITA"},
	{388, "Jamaica", "JM", "JAM"},
	{392, "Japan", "JP", "JPN"},
	{832, "Jersey", "JE", "JEY"},
	{400, "Jordan", "JO", "JOR"},
	{398, "Kazakhstan", "KZ", "KAZ"},
	{404, "Kenya", "KE", "KEN"},
	{296, "Kiribati", "KI", "KIR"},
	{408, "Korea (the Democratic People's Republic of)", "KP", "PRK"},
	{410, "Korea (the Republic of)", "KR", "KOR"},
	{414, "Kuwait", "KW", "KWT"},
	{417, "Kyrgyzstan", "KG", "KGZ"},
	{418, "Lao People's Democratic Republic", "LA", "LAO"},
	{428, "Latvia", "LV", "LVA"},
	{422, "Lebanon", "LB", "LBN"},
	{426, "Lesotho", "LS", "LSO"},
	{430, "Liberia", "LR", "LBR"},
	{434, "Libya", "LY", "LBY"},
	{438, "Liechtenstein", "LI", "LIE"},
	{440, "Lithuania", "LT", "LTU"},
	{442, "Luxembourg", "LU", "LUX"},
	{446, "Macao", "MO", "MAC"},
	{450, "Madagascar", "MG", "MDG"},
	{454, "Malawi", "MW", "MWI"},
	{458, "Malaysia", "MY", "MYS"},
	{462, "Maldives", "MV", "MDV"},
	{466, "Mali", "ML", "MLI"},
	{470, "Malta", "MT", "MLT"},
	{584, "Marshall Islands", "MH", "MHL"},
	{474, "Martinique", "MQ", "MTQ"},
	{478, "Mauritania", "MR", "MRT"},
	{480, "Mauritius", "MU", "MUS"},
	{175, "Mayotte", "YT", "MYT"},
	{484, "Mexico", "MX", "MEX"},
	{583, "Micronesia (Federated States of)", "FM", "FSM"},
	{498, "Moldova (the Republic of)", "MD", "MDA"},
	{492, "Monaco", "MC", "MCO"},
	{496, "Mongolia", "MN", "MNG"},
	{499, "Montenegro", "ME", "MNE"},
	{500, "Montserrat", "MS", "MSR"},
	{504, "Morocco", "MA", "MAR"},
	{508, "Mozambique", "MZ", "MOZ"},
	{104, "Myanmar", "MM", "MMR"},
	{516, "Namibia", "NA", "NAM"},
	{520, "Nauru", "NR", "NRU"},
	{524, "Nepal", "NP", "NPL"},
	{528, "Netherlands", "NL", "NLD"},
	{540, "New Caledonia", "NC", "NCL"},
	{554, "New Zealand", "NZ", "NZL"},
	{558, "Nicaragua", "NI", "NIC"},
	{562, "Niger", "NE", "NER"},
	{566, "Nigeria", "NG", "NGA"},
	{570, "Niue", "NU", "NIU"},
	{574, "Norfolk Island", "NF", "NFK"},
	{807, "North Macedonia", "MK", "MKD"},
	{580, "Northern Mariana Islands", "MP", "MNP"},
	{578, "Norway", "NO", "NOR"},
	{512, "Oman", "OM", "OMN"},
	{586, "Pakistan", "PK", "PAK"},
	{585, "Palau", "PW", "PLW"},
	{275, "Palestine, State of", "PS", "PSE"},
	{591, "Panama", "PA", "PAN"},
	{598, "Papua New Guinea", "PG", "PNG"},
	{600, "Paraguay", "PY", "PRY"},
	{604, "Peru", "PE", "PER"},
	{608, "Philippines", "PH", "PHL"},
	{612, "Pitcairn", "PN", "PCN"},
	{616, "Poland", "PL", "POL"},
	{620, "Portugal", "PT", "PRT"},
	{630, "Puerto Rico", "PR", "PRI"},
	{634, "Qatar", "QA", "QAT"},
	{642, "Romania", "RO", "ROU"},
	{643, "Russian Federation", "RU", "RUS"},
	{646, "Rwanda", "RW", "RWA"},
	{638, "Réunion", "RE", "REU"},
	{652, "Saint Barthélemy", "BL", "BLM"},
	{654, "Saint Helena, Ascension and Tristan da Cunha", "SH", "SHN"},
	{659, "Saint Kitts and Nevis", "KN", "KNA"},
	{662, "Saint Lucia", "LC", "LCA"},
	{663, "Saint Martin (French part)", "MF", "MAF"},
	{666, "Saint Pierre and Miquelon", "PM", "SPM"},
	{670, "Saint Vincent and the Grenadines", "VC", "VCT"},
	{882, "Samoa", "WS", "WSM"},
	{674, "San Marino", "SM", "SMR"},
	{678, "Sao Tome and Principe", "ST", "STP"},
	{682, "Saudi Arabia", "SA", "SAU"},
	{686, "Senegal", "SN", "SEN"},
	{688, "Serbia", "RS", "SRB"},
	{690, "Seychelles", "SC", "SYC"},
	{694, "Sierra Leone", "SL", "SLE"},
	{702, "Singapore", "SG", "SGP"},
	{534, "Sint Maarten (Dutch part)", "SX", "SXM"},
	{703, "Slovakia", "SK", "SVK"},
	{705, "Slovenia", "SI", "SVN"},
	{90, "Solomon Islands", "SB", "SLB"},
	{706, "Somalia", "SO", "SOM"},
	{710, "South Africa", "ZA", "ZAF"},
	{239, "South Georgia and the South Sandwich Islands", "GS", "SGS"},
	{728, "South Sudan", "SS", "SSD"},
	{724, "Spain", "ES", "ESP"},
	{144, "Sri Lanka", "LK", "LKA"},
	{729, "Sudan", "SD", "SDN"},
	{740, "Suriname", "SR", "SUR"},
	{744, "Svalbard and Jan Mayen", "SJ", "SJM"},
	{752, "Sweden", "SE", "SWE"},
	{756, "Switzerland", "CH", "CHE"},
	{760, "Syrian Arab Republic", "SY", "SYR"},
	{158, "Taiwan (Province of China)", "TW", "TWN"},
	{762, "Tajikistan", "TJ", "TJK"},
	{834, "Tanzania, the United Republic of", "TZ", "TZA"},
	{764, "Thailand", "TH", "THA"},
	{626, "Timor-Leste", "TL", "TLS"},
	{768, "Togo", "TG", "TGO"},
	{772, "Tokelau", "TK", "TKL"},
	{776, "Tonga", "TO", "TON"},
	{780, "Trinidad and Tobago", "TT", "TTO"},
	{788, "Tunisia", "TN", "TUN"},
	{792, "Turkey", "TR", "TUR"},
	{795, "Turkmenistan", "TM", "TKM"},
	{796, "Turks and Caicos Islands", "TC", "TCA"},
	{798, "Tuvalu", "TV", "TUV"},
	{800, "Uganda", "UG", "UGA"},
	{804, "Ukraine", "UA", "UKR"},
	{784, "United Arab Emirates", "AE", "ARE"},
	{826, "United Kingdom of Great Britain and Northern Ireland", "GB", "GBR"},
	{581, "United States Minor Outlying Islands", "UM", "UMI"},
	{840, "United States of America", "US", "USA"},
	{858, "Uruguay", "UY", "URY"},
	{860, "Uzbekistan", "UZ", "UZB"},
	{548, "Vanuatu", "VU", "VUT"},
	{862, "Venezuela (Bolivarian Republic of)", "VE", "VEN"},
	{704, "Viet Nam", "VN", "VNM"},
	{92, "Virgin Islands (British)", "VG", "VGB"},
	{850, "Virgin Islands (U.S.)", "VI", "VIR"},
	{876, "Wallis and Futuna", "WF", "WLF"},
	{732, "Western Sahara*", "EH", "ESH"},
	{887, "Yemen", "YE", "YEM"},
	{894, "Zambia", "ZM", "ZMB"},
	{716, "Zimbabwe", "ZW", "ZWE"},
}

func GetISO3166Alpha2Codes() []string {
	var codes []string
	for _, v := range iso3166Codes {
		codes = append(codes, v.alpha2Code)
	}
	return codes
}

// IsISO3166Alpha2Code returns true if it's a valid ISO_3166 Alpha 2 country code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
func IsISO3166Alpha2Code(code string) bool {
	return slice.ContainsString(GetISO3166Alpha2Codes(), code)
}
