package main

type Profile struct {
	Username          string
	Firstname         string
	Lastname          string
	Age               int
	EncryptedPassword string
}

// type Address struct {
// 	Detail  string
// 	Country string
// 	Zip     string
// }

// type Profile struct {
// 	Username          string
// 	Firstname         string
// 	Lastname          string
// 	Age               int
// 	EncryptedPassword string
// 	FavoriteMovies    [5]string
// 	LoggedIPHistories []string
// 	MetaData          map[string]string
// 	Address           Address
// }

func main() {
	// var profile Profile
	// fmt.Println(profile)
	// fmt.Printf("%+v\n", profile)
	// profile.Username = "por"
	// profile.Firstname = "Weerasak"
	// profile.Lastname = "Chongnguluam"
	// profile.Age = 35
	// profile.EncryptedPassword = "encyptedpass1234"
	// fmt.Printf("%+v\n", profile)
	// fmt.Println("Username:", profile.Username)
	// fmt.Println("Firstname:", profile.Firstname)
	// fmt.Println("Lastname:", profile.Lastname)
	// fmt.Println("Age:", profile.Age)
	// fmt.Println("EncryptedPassword:", profile.EncryptedPassword)
	// profile := Profile{
	// 	Username:  "por",
	// 	Firstname: "Weerasak",
	// 	Lastname:  "Chongnguluam",
	// 	// Age:               35,
	// 	// EncryptedPassword: "encyptedpass1234",
	// }
	// profile := Profile{"por", "Weerasak", "Chongnguluam", 35, "encyptedpass1234"}
	// fmt.Println("Username:", profile.Username)
	// fmt.Println("Firstname:", profile.Firstname)
	// fmt.Println("Lastname:", profile.Lastname)
	// fmt.Println("Age:", profile.Age)
	// fmt.Println("EncryptedPassword:", profile.EncryptedPassword)

	// profile := Profile{
	// 	Username:          "por",
	// 	Firstname:         "Weerasak",
	// 	Lastname:          "Chongnguluam",
	// 	Age:               35,
	// 	EncryptedPassword: "encyptedpass1234",
	// 	FavoriteMovies: [5]string{
	// 		"Dark",
	// 		"Game of Thrones",
	// 		"A beautiful mind",
	// 		"Avengers End game",
	// 		"Iron Man",
	// 	},
	// 	LoggedIPHistories: []string{
	// 		"192.168.1.1",
	// 	},
	// 	MetaData: map[string]string{
	// 		"last-order": "Pizza",
	// 	},
	// 	Address: Address{
	// 		Detail:  "12/34",
	// 		Country: "Thailand",
	// 		Zip:     "111111",
	// 	},
	// }
	// fmt.Println(profile)
	// fmt.Println(profile.FavoriteMovies[0])
	// profile.FavoriteMovies[0] = "Spider Man"
	// fmt.Println(profile.FavoriteMovies[0])

	// fmt.Println(profile.LoggedIPHistories[0])
	// profile.LoggedIPHistories = append(profile.LoggedIPHistories, "127.0.0.1")
	// fmt.Println(profile.LoggedIPHistories)

	// fmt.Println(profile.MetaData["last-order"])
	// profile.MetaData["last-order-price"] = "256"
	// fmt.Println(profile.MetaData)

	// fmt.Println(profile.Address.Detail)
	// profile.Address.Zip = "12345"
	// fmt.Println(profile.Address)

	// profileAccessCount := make(map[Profile]int)
	// p1 := Profile{
	// 	Username: "Por",
	// }
	// profileAccessCount[p1] = 1
	// p2 := Profile{
	// 	Username: "May",
	// }
	// profileAccessCount[p2] = 2

	// for k, v := range profileAccessCount {
	// 	fmt.Println(k.Username, v)
	// }

	// cacheProfile := make(map[int]Profile)
	// p1 := Profile{
	// 	Username: "Por",
	// }
	// cacheProfile[1] = p1
	// p2 := Profile{
	// 	Username: "May",
	// }
	// cacheProfile[2] = p2

	// cacheProfile := map[int]Profile{
	// 	1: {
	// 		Username: "Por",
	// 	},
	// 	2: {
	// 		Username: "May",
	// 	},
	// }

	// fmt.Println(cacheProfile[1].Username)
	// fmt.Println(cacheProfile[2].Username)

	// cacheFiveFavorites := make(map[int][5]int)
	// cacheFiveFavorites[1] = [5]int{1, 2, 3, 4, 5}
	// cacheFiveFavorites[2] = [5]int{6, 7, 8, 9, 10}
	// fmt.Println(cacheFiveFavorites[1])
	// fmt.Println(cacheFiveFavorites[2])

	// cacheFiveFavorites := make(map[int][]int)
	// cacheFiveFavorites[1] = []int{1, 2, 3, 4, 5}
	// cacheFiveFavorites[2] = []int{6, 7, 8, 9, 10}
	// fmt.Println(cacheFiveFavorites[1][1])
	// fmt.Println(cacheFiveFavorites[2][1])

	// twoProfiles := [2]Profile{
	// 	{Username: "Por"},
	// 	{Username: "May"},
	// }
	// fmt.Println(twoProfiles[0].Username)
	// fmt.Println(twoProfiles[1].Username)

	// twoMaps := [2]map[string]int{
	// 	{
	// 		"hello": 1,
	// 	},
	// 	{
	// 		"world": 2,
	// 	},
	// }
	// fmt.Println(twoMaps[0]["hello"])
	// fmt.Println(twoMaps[1]["world"])

	// twoSlices := [2][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// fmt.Println(twoSlices[0][0])
	// fmt.Println(twoSlices[1][1])

	// twoArrays := [2][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// fmt.Println(twoArrays[0][0])
	// fmt.Println(twoArrays[1][1])

	// profiles := []Profile{
	// 	{Username: "Por"},
	// 	{Username: "May"},
	// }
	// fmt.Println(profiles[0].Username)
	// fmt.Println(profiles[1].Username)
	// profiles = append(profiles, Profile{Username: "Game"})
	// fmt.Println(profiles[2].Username)

	// maps := []map[string]int{
	// 	{
	// 		"hello": 1,
	// 	},
	// 	{
	// 		"world": 2,
	// 	},
	// }
	// fmt.Println(maps[0]["hello"])
	// fmt.Println(maps[1]["world"])
	// maps = append(maps, map[string]int{"go": 3})
	// fmt.Println(maps[2]["go"])

	// slices := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// fmt.Println(slices[0][0])
	// fmt.Println(slices[1][1])
	// slices = append(slices, []int{5, 6, 7})
	// fmt.Println(slices[2][2])

	// arrays := [][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// }

	// fmt.Println(arrays[0][0])
	// fmt.Println(arrays[1][1])
	// arrays = append(arrays, [2]int{5, 6})
	// fmt.Println(arrays[2][0])
}
