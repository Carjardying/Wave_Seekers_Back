package Seeders

import (
	"database/sql"
	"log"

	"example/Wave_Seekers_Back/Models"
)

// SeedUsers inserts all the users into the DB
func SeedUsers(db *sql.DB) error {
	users := []Models.User{
		{Email: "wave1@test.com", Password: "wave1"},
		{Email: "wave2@test.com", Password: "wave2"},
		{Email: "wave3@test.com", Password: "wave3"},
		{Email: "wave4@test.com", Password: "wave4"},
		{Email: "wave5@test.com", Password: "wave5"},
	}

	for _, u := range users {
		id, err := Models.AddUser(db, &u)
		if err != nil {
			log.Printf("Insert error user %s : %v\n", u.Email, err)
		} else {
			log.Printf("User %s inserted with ID %d\n", u.Email, id)
		}
	}
	return nil
}

// SeedCountries inserts all the countries into the DB
func SeedCountries(db *sql.DB) error {
	countries := []Models.Country{
		{Name: "Afghanistan"}, {Name: "Albania"}, {Name: "Algeria"}, {Name: "Andorra"}, {Name: "Angola"},
		{Name: "Antigua and Barbuda"}, {Name: "Argentina"}, {Name: "Armenia"}, {Name: "Australia"}, {Name: "Austria"},
		{Name: "Azerbaijan"}, {Name: "Bahamas"}, {Name: "Bahrain"}, {Name: "Bangladesh"}, {Name: "Barbados"},
		{Name: "Belarus"}, {Name: "Belgium"}, {Name: "Belize"}, {Name: "Benin"}, {Name: "Bhutan"},
		{Name: "Bolivia"}, {Name: "Bosnia and Herzegovina"}, {Name: "Botswana"}, {Name: "Brazil"}, {Name: "Brunei"},
		{Name: "Bulgaria"}, {Name: "Burkina Faso"}, {Name: "Burundi"}, {Name: "Cabo Verde"}, {Name: "Cambodia"},
		{Name: "Cameroon"}, {Name: "Canada"}, {Name: "Central African Republic"}, {Name: "Chad"}, {Name: "Chile"},
		{Name: "China"}, {Name: "Colombia"}, {Name: "Comoros"}, {Name: "Congo (Congo-Brazzaville)"}, {Name: "Costa Rica"},
		{Name: "Croatia"}, {Name: "Cuba"}, {Name: "Cyprus"}, {Name: "Czechia"}, {Name: "Democratic Republic of the Congo"},
		{Name: "Denmark"}, {Name: "Djibouti"}, {Name: "Dominica"}, {Name: "Dominican Republic"}, {Name: "Ecuador"},
		{Name: "Egypt"}, {Name: "El Salvador"}, {Name: "Equatorial Guinea"}, {Name: "Eritrea"}, {Name: "Estonia"},
		{Name: "Eswatini"}, {Name: "Ethiopia"}, {Name: "Fiji"}, {Name: "Finland"}, {Name: "France"},
		{Name: "Gabon"}, {Name: "Gambia"}, {Name: "Georgia"}, {Name: "Germany"}, {Name: "Ghana"},
		{Name: "Greece"}, {Name: "Grenada"}, {Name: "Guatemala"}, {Name: "Guinea"}, {Name: "Guinea-Bissau"},
		{Name: "Guyana"}, {Name: "Haiti"}, {Name: "Holy See"}, {Name: "Honduras"}, {Name: "Hungary"},
		{Name: "Iceland"}, {Name: "India"}, {Name: "Indonesia"}, {Name: "Iran"}, {Name: "Iraq"},
		{Name: "Ireland"}, {Name: "Israel"}, {Name: "Italy"}, {Name: "Jamaica"}, {Name: "Japan"},
		{Name: "Jordan"}, {Name: "Kazakhstan"}, {Name: "Kenya"}, {Name: "Kiribati"}, {Name: "Kuwait"},
		{Name: "Kyrgyzstan"}, {Name: "Laos"}, {Name: "Latvia"}, {Name: "Lebanon"}, {Name: "Lesotho"},
		{Name: "Liberia"}, {Name: "Libya"}, {Name: "Liechtenstein"}, {Name: "Lithuania"}, {Name: "Luxembourg"},
		{Name: "Madagascar"}, {Name: "Malawi"}, {Name: "Malaysia"}, {Name: "Maldives"}, {Name: "Mali"},
		{Name: "Malta"}, {Name: "Marshall Islands"}, {Name: "Mauritania"}, {Name: "Mauritius"}, {Name: "Mexico"},
		{Name: "Micronesia"}, {Name: "Moldova"}, {Name: "Monaco"}, {Name: "Mongolia"}, {Name: "Montenegro"},
		{Name: "Morocco"}, {Name: "Mozambique"}, {Name: "Myanmar"}, {Name: "Namibia"}, {Name: "Nauru"},
		{Name: "Nepal"}, {Name: "Netherlands"}, {Name: "New Zealand"}, {Name: "Nicaragua"}, {Name: "Niger"},
		{Name: "Nigeria"}, {Name: "North Korea"}, {Name: "North Macedonia"}, {Name: "Norway"}, {Name: "Oman"},
		{Name: "Pakistan"}, {Name: "Palau"}, {Name: "Palestine"}, {Name: "Panama"}, {Name: "Papua New Guinea"},
		{Name: "Paraguay"}, {Name: "Peru"}, {Name: "Philippines"}, {Name: "Poland"}, {Name: "Portugal"},
		{Name: "Qatar"}, {Name: "Romania"}, {Name: "Russia"}, {Name: "Rwanda"}, {Name: "Saint Kitts and Nevis"},
		{Name: "Saint Lucia"}, {Name: "Saint Vincent and the Grenadines"}, {Name: "Samoa"}, {Name: "San Marino"}, {Name: "Sao Tome and Principe"},
		{Name: "Saudi Arabia"}, {Name: "Senegal"}, {Name: "Serbia"}, {Name: "Seychelles"}, {Name: "Sierra Leone"},
		{Name: "Singapore"}, {Name: "Slovakia"}, {Name: "Slovenia"}, {Name: "Solomon Islands"}, {Name: "Somalia"},
		{Name: "South Africa"}, {Name: "South Korea"}, {Name: "South Sudan"}, {Name: "Spain"}, {Name: "Sri Lanka"},
		{Name: "Sudan"}, {Name: "Suriname"}, {Name: "Sweden"}, {Name: "Switzerland"}, {Name: "Syria"},
		{Name: "Tajikistan"}, {Name: "Tanzania"}, {Name: "Thailand"}, {Name: "Timor-Leste"}, {Name: "Togo"},
		{Name: "Tonga"}, {Name: "Trinidad and Tobago"}, {Name: "Tunisia"}, {Name: "Turkey"}, {Name: "Turkmenistan"},
		{Name: "Tuvalu"}, {Name: "Uganda"}, {Name: "Ukraine"}, {Name: "United Arab Emirates"}, {Name: "United Kingdom"},
		{Name: "United States"}, {Name: "Uruguay"}, {Name: "Uzbekistan"}, {Name: "Vanuatu"}, {Name: "Venezuela"},
		{Name: "Vietnam"}, {Name: "Yemen"}, {Name: "Zambia"}, {Name: "Zimbabwe"},
	}

	for _, c := range countries {
		id, err := Models.AddCountry(db, &c)
		if err != nil {
			log.Printf("Insert error country %s : %v\n", c.Name, err)
		} else {
			log.Printf("Country %s inserted with ID %d\n", c.Name, id)
		}
	}

	return nil
}

// SeedSpots inserts all the spots into the DB
func SeedSpots(db *sql.DB) error {
	spots := []Models.Spot{
		{
			UserID:          1,
			CountryID:       186, // United States
			Destination:     "Oahu North Shore",
			Location:        "Hawaii",
			Lat:             21.466667,
			Long:            -157.983333,
			PeakSeasonStart: "07-22",
			PeakSeasonEnd:   "08-31",
			DifficultyLevel: 4,
			SurfingCulture:  "Oahu's North Shore represents the spiritual birthplace of modern surfing, where ancient Polynesian chiefs rode waves as a sacred practice called he'e nalu over a thousand years ago. The Pipeline break, discovered in the 1960s, became the ultimate proving ground for professional surfers, with its perfect but deadly barrels claiming both legends and lives. Hawaiian surfers like Eddie Aikau and the Da Silva family established a culture of respect, courage, and deep ocean knowledge that influences surfing worldwide. The area hosts the world's most prestigious surfing competitions, including the Vans Triple Crown, where careers are made and broken each winter. Today, the North Shore maintains its position as surfing's most revered arena, where local Hawaiian values of aloha and respect for the ocean remain paramount.",
			ImageURL:        "path1",
		},
		{
			UserID:          2,
			CountryID:       119, // Namibia
			Destination:     "Skeleton Bay",
			Location:        "Kunene",
			Lat:             -22.983333,
			Long:            14.5,
			PeakSeasonStart: "09-01",
			PeakSeasonEnd:   "11-30",
			DifficultyLevel: 5,
			SurfingCulture:  "Skeleton Bay emerged from obscurity in the early 2000s when surf explorers discovered this remote left-hand point break along Namibia's desolate coast. The wave breaks over shallow sandbanks created by the Orange River's sediment deposits, producing rides that can last over a minute across multiple sections. Its isolation in the Namib Desert, accessible only by 4WD vehicles, has preserved a raw, adventurous surfing culture reminiscent of early surf exploration. Professional surfers began pilgrimage-like trips to this spot, often camping in harsh desert conditions for weeks to score perfect sessions. The break has become legendary for producing some of the longest barrel rides on Earth, attracting only the most dedicated wave hunters willing to endure extreme conditions.",
			ImageURL:        "path2",
		},
		{
			UserID:          3,
			CountryID:       9, // Australia
			Destination:     "Superbank",
			Location:        "Gold Coast",
			Lat:             -28.166667,
			Long:            153.55,
			PeakSeasonStart: "11-28",
			PeakSeasonEnd:   "02-01",
			DifficultyLevel: 4,
			SurfingCulture:  "The Superbank was artificially created in the 1990s through a massive sand pumping project that transformed three separate breaks into one continuous 2-kilometer wave. This engineering feat revolutionized professional surfing by creating the world's most consistent high-performance wave, hosting the annual World Surf League Championship Tour event. Australian surfing culture thrives here, with local groms learning alongside world champions in a uniquely egalitarian lineup. The spot embodies Australia's progressive approach to wave riding, where technical innovation and competitive excellence drive the culture forward. Despite being man-made, the Superbank has produced more perfect 10-point rides in professional surfing than any other wave, cementing its status as a modern surfing marvel",
			ImageURL:        "path3",
		},
		{
			UserID:          4,
			CountryID:       123, // New Zealand
			Destination:     "Manu Bay",
			Location:        "Waikato",
			Lat:             -37.65,
			Long:            174.75,
			PeakSeasonStart: "12-01",
			PeakSeasonEnd:   "01-31",
			DifficultyLevel: 2,
			SurfingCulture:  "Manu Bay gained international fame after featuring in the 1966 surf film The Endless Summer, introducing the world to New Zealand's pristine left-hand point break. The wave breaks consistently along a rocky coastline near Raglan, creating long, workable walls perfect for traditional longboard surfing and modern shortboard performance. New Zealand's surfing culture at Manu Bay reflects the country's laid-back, environmentally conscious ethos, with strong emphasis on preserving the natural coastline. Local Maori connections to the ocean add cultural depth to the surfing experience, respecting indigenous relationships with the sea. The spot remains relatively uncrowded compared to other world-class breaks, maintaining its appeal as a peaceful surfing sanctuary in the Southern Hemisphere.",
			ImageURL:        "path4",
		},
		{
			UserID:          5,
			CountryID:       137, // Peru
			Destination:     "Playa Chicama",
			Location:        "La Libertad",
			Lat:             -7.85,
			Long:            -79.433333,
			PeakSeasonStart: "05-01",
			PeakSeasonEnd:   "06-28",
			DifficultyLevel: 3,
			SurfingCulture:  "Chicama boasts the world's longest left-hand wave, with rides potentially lasting over two minutes and covering more than two kilometers when conditions align. Ancient Peruvian civilizations, including the Moche people, used reed boats called caballitos de totora for fishing and wave riding over 3,000 years ago, predating Polynesian surf culture. The modern surf scene developed in the 1960s when international surfers discovered this remote northern Peru gem, accessible only by boat or challenging overland routes. Peruvian surfing culture here blends indigenous maritime traditions with contemporary wave riding, creating a unique atmosphere of respect for ancient ocean wisdom. The break's extreme length and power demand patience and skill, attracting surfers seeking the ultimate point break experience in South America.",
			ImageURL:        "path5",
		},
		{
			UserID:          1,
			CountryID:       164, // Spain (Canary Islands)
			Destination:     "The Bubble",
			Location:        "Fuerteventura",
			Lat:             28.35,
			Long:            -14.033333,
			PeakSeasonStart: "06-01",
			PeakSeasonEnd:   "09-01",
			DifficultyLevel: 3,
			SurfingCulture:  "Fuerteventura's surf culture emerged in the 1970s when European surfers discovered consistent Atlantic swells wrapping around volcanic reefs and sand bottoms. The island's year-round warm climate and reliable waves created a European surf destination that rivals tropical locations, hosting numerous international competitions. Local Canarian culture blends Spanish traditions with African influences, creating a unique island vibe that welcomes international surf travelers. The surf industry now drives much of Fuerteventura's economy, with surf schools, camps, and board shapers establishing a thriving wave-riding community. Trade winds and consistent North Atlantic swells make the island a reliable training ground for European surfers seeking to hone their skills in powerful, consistent waves.",
			ImageURL:        "path6",
		},
		{
			UserID:          2,
			CountryID:       104, // Maldives
			Destination:     "Pasta Point",
			Location:        "North Male Atoll",
			Lat:             3.066667,
			Long:            73.15,
			PeakSeasonStart: "04-01",
			PeakSeasonEnd:   "05-31",
			DifficultyLevel: 3,
			SurfingCulture:  "Pasta Point represents the pinnacle of tropical surfing luxury, where crystal-clear waters break over shallow coral reefs in perfect barrels. The wave was \"discovered\" by surf tourists in the 1990s, though local fishermen had observed these breaks for centuries while navigating between atolls. Maldivian surf culture centers around exclusive surf resorts and boat trips, creating a high-end surf tourism model that protects waves through limited access. The break's perfection comes from deep ocean swells wrapping around coral atolls, creating mechanical waves that seem almost too perfect to be natural. Climate change and coral bleaching now threaten these pristine surf spots, making each session precious and highlighting surfing's relationship with environmental conservation.",
			ImageURL:        "path7",
		},
		{
			UserID:          3,
			CountryID:       161, // South Africa
			Destination:     "Supertubes Beach",
			Location:        "Eastern Cape",
			Lat:             -34.05,
			Long:            24.916667,
			PeakSeasonStart: "08-01",
			PeakSeasonEnd:   "10-09",
			DifficultyLevel: 5,
			SurfingCulture:  "J-Bay established itself as Africa's premier surf destination in the 1960s, when surfers discovered the perfect right-hand point break during apartheid-era South Africa. The wave became internationally famous through surf films and the annual Billabong Pro competition, showcasing South African surfing talent to the world. Local surf culture reflects the country's complex history, with post-apartheid integration creating a more diverse and inclusive surfing community. The town of Jeffreys Bay transformed from a small fishing village into a global surf industry hub, hosting major surfboard manufacturers and surf brands. Great white sharks patrol these waters, adding an element of danger that has shaped the fearless character of South African surfing culture.",
			ImageURL:        "path8",
		},
		{
			UserID:          4,
			CountryID:       186, // United States
			Destination:     "Kitty Hawk",
			Location:        "North Carolina",
			Lat:             36.066667,
			Long:            -75.7,
			PeakSeasonStart: "08-09",
			PeakSeasonEnd:   "10-18",
			DifficultyLevel: 3,
			SurfingCulture:  "The Outer Banks surf scene developed alongside America's East Coast surf culture in the 1960s, with consistent Atlantic hurricane swells creating world-class waves. This barrier island chain holds historical significance as the birthplace of aviation, where the Wright brothers first flew, adding pioneering spirit to the local surf culture. Hurricane season brings powerful swells that can rival any surf destination globally, creating a hardcore community of surfers willing to brave dangerous conditions. The area's relative isolation and harsh winter conditions have fostered a tight-knit surf community that values authenticity over commercialization. Local surfers developed unique techniques for reading rapidly changing sandbars and powerful shore break, creating a distinct East Coast surfing style.",
			ImageURL:        "path9",
		},
		{
			UserID:          5,
			CountryID:       186, // United States
			Destination:     "Rockaway Beach",
			Location:        "Oregon",
			Lat:             45.616667,
			Long:            -123.95,
			PeakSeasonStart: "08-23",
			PeakSeasonEnd:   "10-17",
			DifficultyLevel: 1,
			SurfingCulture:  "Rockaway Beach represents the rugged character of Pacific Northwest surfing, where cold water and powerful waves demand respect and proper equipment. The surf culture here emerged in the 1960s despite frigid water temperatures requiring full wetsuits year-round, attracting dedicated surfers seeking uncrowded waves. Oregon's environmental consciousness influences the local surf community, with strong emphasis on beach cleanups and ocean conservation efforts. The consistent but challenging conditions create skilled surfers who can handle powerful waves in adverse weather conditions. Rockaway's surf culture embodies the Pacific Northwest's independent spirit, where surfers value solitude and connection with raw nature over tropical perfection.",
			ImageURL:        "path10",
		},
	}

	for _, s := range spots {
		id, err := Models.AddSpot(db, &s)
		if err != nil {
			log.Printf("Insert error spot %s : %v\n", s.Destination, err)
		} else {
			log.Printf("Spot %s inserted with ID %d\n", s.Destination, id)
		}
	}
	return nil
}
