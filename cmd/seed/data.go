package main

import (
	"fmt"
	"math/rand"
	"time"

	repo "ticketr/internal/repository"

	"github.com/google/uuid"
)

func movies() []repo.CreateMovieParams {
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	return []repo.CreateMovieParams{
		// --- PREVIOUS MONTH (already released) ---
		{Name: "Shadow Strike", Description: "Elite agent uncovers conspiracy.", Casts: []string{"Actor A"}, TrailerUrl: "url1", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, -1, 5), Status: repo.ReleaseStatusRELEASED, Director: "Dir A"},
		{Name: "Midnight Echo", Description: "Dream thriller.", Casts: []string{"Actor B"}, TrailerUrl: "url2", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, -1, 10), Status: repo.ReleaseStatusRELEASED, Director: "Dir B"},
		{Name: "Crimson Code", Description: "Cyber mystery.", Casts: []string{"Actor C"}, TrailerUrl: "url3", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, -1, 15), Status: repo.ReleaseStatusRELEASED, Director: "Dir C"},
		{Name: "Silent Waves", Description: "Musical journey.", Casts: []string{"Actor D"}, TrailerUrl: "url4", Languages: []string{"Tamil"}, ReleaseDate: startOfMonth.AddDate(0, -1, 20), Status: repo.ReleaseStatusRELEASED, Director: "Dir D"},
		{Name: "Iron Will", Description: "Soldier story.", Casts: []string{"Actor E"}, TrailerUrl: "url5", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, -1, 25), Status: repo.ReleaseStatusRELEASED, Director: "Dir E"},
		{Name: "Beyond Stars", Description: "Galaxy trip.", Casts: []string{"Actor Z"}, TrailerUrl: "url26", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, -1, 12), Status: repo.ReleaseStatusRELEASED, Director: "Dir Z"},
		{Name: "Hidden Truth", Description: "Secrets unfold.", Casts: []string{"Actor AA"}, TrailerUrl: "url27", Languages: []string{"Malayalam"}, ReleaseDate: startOfMonth.AddDate(0, -1, 18), Status: repo.ReleaseStatusRELEASED, Director: "Dir AA"},
		{Name: "Desert Storm", Description: "Harsh survival.", Casts: []string{"Actor AB"}, TrailerUrl: "url28", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, -1, 28), Status: repo.ReleaseStatusRELEASED, Director: "Dir AB"},

		// --- CURRENT MONTH ---
		{Name: "City Lights", Description: "Urban love.", Casts: []string{"Actor F"}, TrailerUrl: "url6", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, 0, 2), Status: repo.ReleaseStatusRELEASED, Director: "Dir F"},
		{Name: "Quantum Rift", Description: "Time breaks.", Casts: []string{"Actor G"}, TrailerUrl: "url7", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 0, 5), Status: repo.ReleaseStatusRELEASED, Director: "Dir G"},
		{Name: "Neon Streets", Description: "Cyber chaos.", Casts: []string{"Actor H"}, TrailerUrl: "url8", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 0, 8), Status: repo.ReleaseStatusRELEASED, Director: "Dir H"},
		{Name: "Broken Silence", Description: "Court drama.", Casts: []string{"Actor I"}, TrailerUrl: "url9", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, 0, 10), Status: repo.ReleaseStatusRELEASED, Director: "Dir I"},
		{Name: "Night Chase", Description: "Crime night.", Casts: []string{"Actor J"}, TrailerUrl: "url10", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 0, 12), Status: repo.ReleaseStatusRELEASED, Director: "Dir J"},
		{Name: "Dreamcatcher", Description: "Dream world.", Casts: []string{"Actor K"}, TrailerUrl: "url11", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, 0, 15), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir K"},
		{Name: "Hidden Flame", Description: "Spy rogue.", Casts: []string{"Actor L"}, TrailerUrl: "url12", Languages: []string{"Tamil"}, ReleaseDate: startOfMonth.AddDate(0, 0, 18), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir L"},
		{Name: "Parallel Minds", Description: "Dual timelines.", Casts: []string{"Actor M"}, TrailerUrl: "url13", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 0, 20), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir M"},
		{Name: "Golden Harvest", Description: "Farm fight.", Casts: []string{"Actor N"}, TrailerUrl: "url14", Languages: []string{"Kannada"}, ReleaseDate: startOfMonth.AddDate(0, 0, 22), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir N"},
		{Name: "Ocean Secret", Description: "Underwater city.", Casts: []string{"Actor O"}, TrailerUrl: "url15", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 0, 25), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir O"},
		{Name: "Fire Within", Description: "Boxer rises.", Casts: []string{"Actor AC"}, TrailerUrl: "url29", Languages: []string{"Tamil"}, ReleaseDate: startOfMonth.AddDate(0, 0, 17), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir AC"},
		{Name: "Night Hunter", Description: "Dark chase.", Casts: []string{"Actor AD"}, TrailerUrl: "url30", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 0, 21), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir AD"},
		{Name: "Skyfall Edge", Description: "Air battle.", Casts: []string{"Actor AE"}, TrailerUrl: "url31", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 0, 27), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir AE"},

		// --- NEXT MONTH ---
		{Name: "Final Horizon", Description: "Future journey.", Casts: []string{"Actor P"}, TrailerUrl: "url16", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 2), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir P"},
		{Name: "Galaxy Wars", Description: "Space war.", Casts: []string{"Actor Q"}, TrailerUrl: "url17", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 5), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir Q"},
		{Name: "Dark Horizon", Description: "Mission fails.", Casts: []string{"Actor R"}, TrailerUrl: "url18", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 8), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir R"},
		{Name: "Lost Identity", Description: "Memory loss.", Casts: []string{"Actor S"}, TrailerUrl: "url19", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, 1, 10), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir S"},
		{Name: "Rising Thunder", Description: "Bike gang.", Casts: []string{"Actor T"}, TrailerUrl: "url20", Languages: []string{"Kannada"}, ReleaseDate: startOfMonth.AddDate(0, 1, 12), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir T"},
		{Name: "Burning Roads", Description: "Fast chase.", Casts: []string{"Actor U"}, TrailerUrl: "url21", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, 1, 15), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir U"},
		{Name: "Frozen Path", Description: "Ice survival.", Casts: []string{"Actor V"}, TrailerUrl: "url22", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 18), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir V"},
		{Name: "Street Kings", Description: "Gang war.", Casts: []string{"Actor W"}, TrailerUrl: "url23", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, 1, 20), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir W"},
		{Name: "Echoes Time", Description: "Time travel.", Casts: []string{"Actor X"}, TrailerUrl: "url24", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 22), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir X"},
		{Name: "Final Verdict", Description: "Legal twist.", Casts: []string{"Actor Y"}, TrailerUrl: "url25", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 25), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir Y"},
		{Name: "Deep Shadows", Description: "Mystery unfolds.", Casts: []string{"Actor AF"}, TrailerUrl: "url32", Languages: []string{"Hindi"}, ReleaseDate: startOfMonth.AddDate(0, 1, 14), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir AF"},
		{Name: "Last Signal", Description: "Sci-fi thriller.", Casts: []string{"Actor AG"}, TrailerUrl: "url33", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 17), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir AG"},
		{Name: "Zero Hour", Description: "Race against time.", Casts: []string{"Actor AH"}, TrailerUrl: "url34", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 23), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir AH"},
		{Name: "Phantom Code", Description: "Spy thriller.", Casts: []string{"Actor AI"}, TrailerUrl: "url35", Languages: []string{"English"}, ReleaseDate: startOfMonth.AddDate(0, 1, 27), Status: repo.ReleaseStatusUNRELEASED, Director: "Dir AI"},
	}
}

func cities() []repo.CreateCityParams {
	return []repo.CreateCityParams{
		{City: "Mysuru", State: "Karnataka"},
		{City: "Belagavi", State: "Karnataka"},
		{City: "Shivamogga", State: "Karnataka"},
		{City: "Tumakuru", State: "Karnataka"},
		{City: "Davangere", State: "Karnataka"},
		{City: "Ballari", State: "Karnataka"},
		{City: "Vijayapura", State: "Karnataka"},
		{City: "Kalaburagi", State: "Karnataka"},
		{City: "Chitradurga", State: "Karnataka"},
		{City: "Hassan", State: "Karnataka"},
		{City: "Mandya", State: "Karnataka"},
		{City: "Raichur", State: "Karnataka"},
		{City: "Bidar", State: "Karnataka"},
		{City: "Kolar", State: "Karnataka"},
		{City: "Chikkamagaluru", State: "Karnataka"},
		{City: "Karwar", State: "Karnataka"},
		{City: "Bagalkot", State: "Karnataka"},
		{City: "Gadag", State: "Karnataka"},
		{City: "Yadgir", State: "Karnataka"},
		{City: "Haveri", State: "Karnataka"},
	}
}

func theaters(cityIDs []uuid.UUID) []repo.CreateTheaterParams {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	theaterNames := []string{
		"PVR Cinemas", "INOX", "Cinepolis", "Miraj Cinemas",
		"Wave Cinemas", "Carnival Cinemas", "Asian Cinemas",
		"Movietime", "SRS Cinemas", "Fun Cinemas",
		"SPI Cinemas", "AGS Cinemas", "Q Cinemas",
		"Mukta A2 Cinemas", "Rajhans Cinemas",
		"Prasads Multiplex", "Aries Plex", "EVM Cinemas",
		"Cine Galaxy", "Gold Cinema", "DT Cinemas",
		"City Pride Multiplex", "Roopbani Cinema", "Kamala Cinemas",
		"Urvashi Cinema", "Narthaki Theatre", "Venkateshwara Theatre",
		"Annapurna Theatre", "Gopalan Cinemas", "Cine Square",
	}

	var theaters []repo.CreateTheaterParams

	for i := range 90 {
		cityID := cityIDs[r.Intn(len(cityIDs))]

		pincode := r.Intn(900000) + 100000 // ensures 6-digit

		theater := repo.CreateTheaterParams{
			Name:        theaterNames[r.Intn(len(theaterNames))],
			Description: "A premium movie experience",
			CityID:      cityID,
			Address:     fmt.Sprintf("Address #%d", i+1),
			Pincode:     fmt.Sprintf("%06d", pincode),
		}

		theaters = append(theaters, theater)
	}

	return theaters
}

func screens(theaterIDs []uuid.UUID) []repo.CreateScreenParams {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var screens []repo.CreateScreenParams

	for _, theaterID := range theaterIDs {
		numScreens := 3

		for i := range numScreens {
			screen := repo.CreateScreenParams{
				Name:       fmt.Sprintf("Screen %d", i+1),
				TheaterID:  theaterID,
				TotalSeats: int32(80 + r.Intn(121)), // 80–200 seats
			}

			screens = append(screens, screen)
		}
	}

	return screens
}

func shows(movies []repo.CreateMovieRow, screenIDs []uuid.UUID) []repo.CreateShowParams {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var shows []repo.CreateShowParams

	daysToGenerate := 10

	for _, screenID := range screenIDs {
		for d := range daysToGenerate {
			day := time.Now().AddDate(0, 0, d)

			// Start at 9 AM
			current := time.Date(
				day.Year(), day.Month(), day.Day(),
				9, 0, 0, 0, day.Location(),
			)

			endOfDay := time.Date(
				day.Year(), day.Month(), day.Day(),
				23, 0, 0, 0, day.Location(),
			)

			for current.Before(endOfDay) {

				// Filter movies that are released before this time
				var available []repo.CreateMovieRow
				for _, m := range movies {
					if !m.ReleaseDate.After(current) {
						available = append(available, m)
					}
				}

				if len(available) == 0 {
					// No movies released yet → move forward
					current = current.Add(3 * time.Hour)
					continue
				}

				// Pick random movie
				movie := available[r.Intn(len(available))]

				// Random duration: 2–3 hours
				duration := time.Duration(120+r.Intn(60)) * time.Minute

				start := current
				end := current.Add(duration)

				shows = append(shows, repo.CreateShowParams{
					MovieID:   movie.ID,
					ScreenID:  screenID,
					StartTime: start,
					EndTime:   end,
				})

				// Next show starts exactly when this ends
				current = end
			}
		}
	}

	return shows
}
