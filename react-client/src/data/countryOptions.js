const countryData = [
  { alpha2: "AD", name: "Andorra" },
  { alpha2: "AE", name: "United Arab Emirates" },
  { alpha2: "AF", name: "Afghanistan" },
  { alpha2: "AG", name: "Antigua and Barbuda" },
  { alpha2: "AI", name: "Anguilla" },
  { alpha2: "AL", name: "Albania" },
  { alpha2: "AM", name: "Armenia" },
  { alpha2: "AO", name: "Angola" },
  { alpha2: "AR", name: "Argentina" },
  { alpha2: "AT", name: "Austria" },
  { alpha2: "AU", name: "Australia" },
  { alpha2: "AZ", name: "Azerbaijan" },
  { alpha2: "BA", name: "Bosnia and Herzegovina" },
  { alpha2: "BB", name: "Barbados" },
  { alpha2: "BD", name: "Bangladesh" },
  { alpha2: "BE", name: "Belgium" },
  { alpha2: "BF", name: "Burkina Faso" },
  { alpha2: "BG", name: "Bulgaria" },
  { alpha2: "BH", name: "Bahrain" },
  { alpha2: "BI", name: "Burundi" },
  { alpha2: "BJ", name: "Benin" },
  { alpha2: "BN", name: "Brunei" },
  { alpha2: "BO", name: "Bolivia" },
  { alpha2: "BR", name: "Brazil" },
  { alpha2: "BS", name: "Bahamas" },
  { alpha2: "BT", name: "Bhutan" },
  { alpha2: "BW", name: "Botswana" },
  { alpha2: "BY", name: "Belarus" },
  { alpha2: "BZ", name: "Belize" },
  { alpha2: "CA", name: "Canada" },
  { alpha2: "CH", name: "Switzerland" },
  { alpha2: "CL", name: "Chile" },
  { alpha2: "CN", name: "China" },
  { alpha2: "CO", name: "Colombia" },
  { alpha2: "CR", name: "Costa Rica" },
  { alpha2: "CU", name: "Cuba" },
  { alpha2: "CV", name: "Cabo Verde" },
  { alpha2: "CY", name: "Cyprus" },
  { alpha2: "CZ", name: "Czechia" },
  { alpha2: "DE", name: "Germany" },
  { alpha2: "DK", name: "Denmark" },
  { alpha2: "DO", name: "Dominican Republic" },
  { alpha2: "DZ", name: "Algeria" },
  { alpha2: "EC", name: "Ecuador" },
  { alpha2: "EE", name: "Estonia" },
  { alpha2: "EG", name: "Egypt" },
  { alpha2: "ES", name: "Spain" },
  { alpha2: "ET", name: "Ethiopia" },
  { alpha2: "FI", name: "Finland" },
  { alpha2: "FR", name: "France" },
  { alpha2: "GB", name: "United Kingdom" },
  { alpha2: "GR", name: "Greece" },
  { alpha2: "GT", name: "Guatemala" },
  { alpha2: "HK", name: "Hong Kong" },
  { alpha2: "HN", name: "Honduras" },
  { alpha2: "HR", name: "Croatia" },
  { alpha2: "HT", name: "Haiti" },
  { alpha2: "HU", name: "Hungary" },
  { alpha2: "ID", name: "Indonesia" },
  { alpha2: "IE", name: "Ireland" },
  { alpha2: "IL", name: "Israel" },
  { alpha2: "IN", name: "India" },
  { alpha2: "IQ", name: "Iraq" },
  { alpha2: "IR", name: "Iran" },
  { alpha2: "IS", name: "Iceland" },
  { alpha2: "IT", name: "Italy" },
  { alpha2: "JP", name: "Japan" },
  { alpha2: "KE", name: "Kenya" },
  { alpha2: "KR", name: "South Korea" },
  { alpha2: "KW", name: "Kuwait" },
  { alpha2: "KZ", name: "Kazakhstan" },
  { alpha2: "LB", name: "Lebanon" },
  { alpha2: "LK", name: "Sri Lanka" },
  { alpha2: "LT", name: "Lithuania" },
  { alpha2: "LU", name: "Luxembourg" },
  { alpha2: "LV", name: "Latvia" },
  { alpha2: "MA", name: "Morocco" },
  { alpha2: "MC", name: "Monaco" },
  { alpha2: "MD", name: "Moldova" },
  { alpha2: "ME", name: "Montenegro" },
  { alpha2: "MG", name: "Madagascar" },
  { alpha2: "MK", name: "North Macedonia" },
  { alpha2: "ML", name: "Mali" },
  { alpha2: "MM", name: "Myanmar" },
  { alpha2: "MN", name: "Mongolia" },
  { alpha2: "MO", name: "Macao" },
  { alpha2: "MT", name: "Malta" },
  { alpha2: "MU", name: "Mauritius" },
  { alpha2: "MX", name: "Mexico" },
  { alpha2: "MY", name: "Malaysia" },
  { alpha2: "MZ", name: "Mozambique" },
  { alpha2: "NA", name: "Namibia" },
  { alpha2: "NG", name: "Nigeria" },
  { alpha2: "NI", name: "Nicaragua" },
  { alpha2: "NL", name: "The Netherlands" },
  { alpha2: "NO", name: "Norway" },
  { alpha2: "NP", name: "Nepal" },
  { alpha2: "NZ", name: "New Zealand" },
  { alpha2: "OM", name: "Oman" },
  { alpha2: "PA", name: "Panama" },
  { alpha2: "PE", name: "Peru" },
  { alpha2: "PH", name: "Philippines" },
  { alpha2: "PK", name: "Pakistan" },
  { alpha2: "PL", name: "Poland" },
  { alpha2: "PT", name: "Portugal" },
  { alpha2: "PY", name: "Paraguay" },
  { alpha2: "QA", name: "Qatar" },
  { alpha2: "RO", name: "Romania" },
  { alpha2: "RS", name: "Serbia" },
  { alpha2: "RU", name: "Russia" },
  { alpha2: "SA", name: "Saudi Arabia" },
  { alpha2: "SE", name: "Sweden" },
  { alpha2: "SG", name: "Singapore" },
  { alpha2: "SI", name: "Slovenia" },
  { alpha2: "SK", name: "Slovakia" },
  { alpha2: "SN", name: "Senegal" },
  { alpha2: "SO", name: "Somalia" },
  { alpha2: "SR", name: "Suriname" },
  { alpha2: "SV", name: "El Salvador" },
  { alpha2: "SY", name: "Syria" },
  { alpha2: "TH", name: "Thailand" },
  { alpha2: "TJ", name: "Tajikistan" },
  { alpha2: "TL", name: "Timor-Leste" },
  { alpha2: "TM", name: "Turkmenistan" },
  { alpha2: "TN", name: "Tunisia" },
  { alpha2: "TR", name: "Türkiye" },
  { alpha2: "TT", name: "Trinidad and Tobago" },
  { alpha2: "TW", name: "Taiwan" },
  { alpha2: "TZ", name: "Tanzania" },
  { alpha2: "UA", name: "Ukraine" },
  { alpha2: "UG", name: "Uganda" },
  { alpha2: "US", name: "United States" },
  { alpha2: "UY", name: "Uruguay" },
  { alpha2: "UZ", name: "Uzbekistan" },
  { alpha2: "VA", name: "Vatican City" },
  { alpha2: "VE", name: "Venezuela" },
  { alpha2: "VN", name: "Vietnam" },
  { alpha2: "ZA", name: "South Africa" },
  { alpha2: "ZM", name: "Zambia" },
  { alpha2: "ZW", name: "Zimbabwe" },
];

export default countryData;