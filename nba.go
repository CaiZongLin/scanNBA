package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Schedule struct {
	Context struct {
		User struct {
			CountryCode  string `json:"countryCode"`
			CountryName  string `json:"countryName"`
			Locale       string `json:"locale"`
			TimeZone     string `json:"timeZone"`
			TimeZoneCity string `json:"timeZoneCity"`
		} `json:"user"`
		Device struct {
			Clazz interface{} `json:"clazz"`
		} `json:"device"`
	} `json:"context"`
	Error struct {
		Detail  interface{} `json:"detail"`
		IsError string      `json:"isError"`
		Message interface{} `json:"message"`
	} `json:"error"`
	Payload struct {
		League struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"league"`
		Season struct {
			IsCurrent               string `json:"isCurrent"`
			RosterSeasonType        int    `json:"rosterSeasonType"`
			RosterSeasonYear        string `json:"rosterSeasonYear"`
			RosterSeasonYearDisplay string `json:"rosterSeasonYearDisplay"`
			ScheduleSeasonType      int    `json:"scheduleSeasonType"`
			ScheduleSeasonYear      string `json:"scheduleSeasonYear"`
			ScheduleYearDisplay     string `json:"scheduleYearDisplay"`
			StatsSeasonType         int    `json:"statsSeasonType"`
			StatsSeasonYear         string `json:"statsSeasonYear"`
			StatsSeasonYearDisplay  string `json:"statsSeasonYearDisplay"`
			Year                    string `json:"year"`
			YearDisplay             string `json:"yearDisplay"`
		} `json:"season"`
		Date struct {
			Games []struct {
				Profile struct {
					ArenaLocation string      `json:"arenaLocation"`
					ArenaName     string      `json:"arenaName"`
					AwayTeamID    string      `json:"awayTeamId"`
					DateTimeEt    string      `json:"dateTimeEt"`
					GameID        string      `json:"gameId"`
					HomeTeamID    string      `json:"homeTeamId"`
					Number        string      `json:"number"`
					ScheduleCode  interface{} `json:"scheduleCode"`
					SeasonType    string      `json:"seasonType"`
					Sequence      string      `json:"sequence"`
					UtcMillis     string      `json:"utcMillis"`
				} `json:"profile"`
				Boxscore struct {
					Attendance            string      `json:"attendance"`
					AwayScore             int         `json:"awayScore"`
					GameLength            interface{} `json:"gameLength"`
					HomeScore             int         `json:"homeScore"`
					LeadChanges           interface{} `json:"leadChanges"`
					OfficialsDisplayName1 interface{} `json:"officialsDisplayName1"`
					OfficialsDisplayName2 interface{} `json:"officialsDisplayName2"`
					OfficialsDisplayName3 interface{} `json:"officialsDisplayName3"`
					Period                string      `json:"period"`
					PeriodClock           interface{} `json:"periodClock"`
					Status                string      `json:"status"`
					StatusDesc            interface{} `json:"statusDesc"`
					Ties                  interface{} `json:"ties"`
				} `json:"boxscore"`
				Urls         []interface{} `json:"urls"`
				Broadcasters []interface{} `json:"broadcasters"`
				HomeTeam     struct {
					Profile struct {
						Abbr              string `json:"abbr"`
						City              string `json:"city"`
						CityEn            string `json:"cityEn"`
						Code              string `json:"code"`
						Conference        string `json:"conference"`
						DisplayAbbr       string `json:"displayAbbr"`
						DisplayConference string `json:"displayConference"`
						Division          string `json:"division"`
						ID                string `json:"id"`
						IsAllStarTeam     bool   `json:"isAllStarTeam"`
						IsLeagueTeam      bool   `json:"isLeagueTeam"`
						LeagueID          string `json:"leagueId"`
						Name              string `json:"name"`
						NameEn            string `json:"nameEn"`
					} `json:"profile"`
					Matchup struct {
						ConfRank   string      `json:"confRank"`
						DivRank    string      `json:"divRank"`
						Losses     string      `json:"losses"`
						SeriesText interface{} `json:"seriesText"`
						Wins       string      `json:"wins"`
					} `json:"matchup"`
					Score struct {
						Assists                int     `json:"assists"`
						BiggestLead            int     `json:"biggestLead"`
						Blocks                 int     `json:"blocks"`
						BlocksAgainst          int     `json:"blocksAgainst"`
						DefRebs                int     `json:"defRebs"`
						Disqualifications      int     `json:"disqualifications"`
						Ejections              int     `json:"ejections"`
						FastBreakPoints        int     `json:"fastBreakPoints"`
						Fga                    int     `json:"fga"`
						Fgm                    int     `json:"fgm"`
						Fgpct                  float64 `json:"fgpct"`
						FlagrantFouls          int     `json:"flagrantFouls"`
						Fouls                  int     `json:"fouls"`
						Fta                    int     `json:"fta"`
						Ftm                    int     `json:"ftm"`
						Ftpct                  float64 `json:"ftpct"`
						FullTimeoutsRemaining  int     `json:"fullTimeoutsRemaining"`
						Mins                   int     `json:"mins"`
						OffRebs                int     `json:"offRebs"`
						Ot10Score              int     `json:"ot10Score"`
						Ot1Score               int     `json:"ot1Score"`
						Ot2Score               int     `json:"ot2Score"`
						Ot3Score               int     `json:"ot3Score"`
						Ot4Score               int     `json:"ot4Score"`
						Ot5Score               int     `json:"ot5Score"`
						Ot6Score               int     `json:"ot6Score"`
						Ot7Score               int     `json:"ot7Score"`
						Ot8Score               int     `json:"ot8Score"`
						Ot9Score               int     `json:"ot9Score"`
						PointsInPaint          int     `json:"pointsInPaint"`
						PointsOffTurnovers     int     `json:"pointsOffTurnovers"`
						Q1Score                int     `json:"q1Score"`
						Q2Score                int     `json:"q2Score"`
						Q3Score                int     `json:"q3Score"`
						Q4Score                int     `json:"q4Score"`
						Rebs                   int     `json:"rebs"`
						Score                  int     `json:"score"`
						Seconds                int     `json:"seconds"`
						ShortTimeoutsRemaining int     `json:"shortTimeoutsRemaining"`
						Steals                 int     `json:"steals"`
						TechnicalFouls         int     `json:"technicalFouls"`
						Tpa                    int     `json:"tpa"`
						Tpm                    int     `json:"tpm"`
						Tppct                  float64 `json:"tppct"`
						Turnovers              int     `json:"turnovers"`
					} `json:"score"`
					PointGameLeader   interface{} `json:"pointGameLeader"`
					AssistGameLeader  interface{} `json:"assistGameLeader"`
					ReboundGameLeader interface{} `json:"reboundGameLeader"`
				} `json:"homeTeam"`
				AwayTeam struct {
					Profile struct {
						Abbr              string `json:"abbr"`
						City              string `json:"city"`
						CityEn            string `json:"cityEn"`
						Code              string `json:"code"`
						Conference        string `json:"conference"`
						DisplayAbbr       string `json:"displayAbbr"`
						DisplayConference string `json:"displayConference"`
						Division          string `json:"division"`
						ID                string `json:"id"`
						IsAllStarTeam     bool   `json:"isAllStarTeam"`
						IsLeagueTeam      bool   `json:"isLeagueTeam"`
						LeagueID          string `json:"leagueId"`
						Name              string `json:"name"`
						NameEn            string `json:"nameEn"`
					} `json:"profile"`
					Matchup struct {
						ConfRank   string      `json:"confRank"`
						DivRank    string      `json:"divRank"`
						Losses     string      `json:"losses"`
						SeriesText interface{} `json:"seriesText"`
						Wins       string      `json:"wins"`
					} `json:"matchup"`
					Score struct {
						Assists                int     `json:"assists"`
						BiggestLead            int     `json:"biggestLead"`
						Blocks                 int     `json:"blocks"`
						BlocksAgainst          int     `json:"blocksAgainst"`
						DefRebs                int     `json:"defRebs"`
						Disqualifications      int     `json:"disqualifications"`
						Ejections              int     `json:"ejections"`
						FastBreakPoints        int     `json:"fastBreakPoints"`
						Fga                    int     `json:"fga"`
						Fgm                    int     `json:"fgm"`
						Fgpct                  float64 `json:"fgpct"`
						FlagrantFouls          int     `json:"flagrantFouls"`
						Fouls                  int     `json:"fouls"`
						Fta                    int     `json:"fta"`
						Ftm                    int     `json:"ftm"`
						Ftpct                  float64 `json:"ftpct"`
						FullTimeoutsRemaining  int     `json:"fullTimeoutsRemaining"`
						Mins                   int     `json:"mins"`
						OffRebs                int     `json:"offRebs"`
						Ot10Score              int     `json:"ot10Score"`
						Ot1Score               int     `json:"ot1Score"`
						Ot2Score               int     `json:"ot2Score"`
						Ot3Score               int     `json:"ot3Score"`
						Ot4Score               int     `json:"ot4Score"`
						Ot5Score               int     `json:"ot5Score"`
						Ot6Score               int     `json:"ot6Score"`
						Ot7Score               int     `json:"ot7Score"`
						Ot8Score               int     `json:"ot8Score"`
						Ot9Score               int     `json:"ot9Score"`
						PointsInPaint          int     `json:"pointsInPaint"`
						PointsOffTurnovers     int     `json:"pointsOffTurnovers"`
						Q1Score                int     `json:"q1Score"`
						Q2Score                int     `json:"q2Score"`
						Q3Score                int     `json:"q3Score"`
						Q4Score                int     `json:"q4Score"`
						Rebs                   int     `json:"rebs"`
						Score                  int     `json:"score"`
						Seconds                int     `json:"seconds"`
						ShortTimeoutsRemaining int     `json:"shortTimeoutsRemaining"`
						Steals                 int     `json:"steals"`
						TechnicalFouls         int     `json:"technicalFouls"`
						Tpa                    int     `json:"tpa"`
						Tpm                    int     `json:"tpm"`
						Tppct                  float64 `json:"tppct"`
						Turnovers              int     `json:"turnovers"`
					} `json:"score"`
					PointGameLeader   interface{} `json:"pointGameLeader"`
					AssistGameLeader  interface{} `json:"assistGameLeader"`
					ReboundGameLeader interface{} `json:"reboundGameLeader"`
				} `json:"awayTeam"`
				IfNecessary bool        `json:"ifNecessary"`
				SeriesText  interface{} `json:"seriesText"`
			} `json:"games"`
			DateMillis string `json:"dateMillis"`
			GameCount  string `json:"gameCount"`
		} `json:"date"`
		NextAvailableDateMillis string `json:"nextAvailableDateMillis"`
		UtcMillis               string `json:"utcMillis"`
	} `json:"payload"`
	Timestamp string `json:"timestamp"`
}

//???????????????map
func TeamInit() map[string]string {

	teamMap := make(map[string]string)
	teamMap["Atlanta Hawks"] = "??????"
	teamMap["Boston Celtics"] = "??????"
	teamMap["Brooklyn Nets"] = "??????"
	teamMap["Cleveland Cavaliers"] = "??????"
	teamMap["Charlotte Hornets"] = "??????"
	teamMap["Chicago Bulls"] = "??????"
	teamMap["Dallas Mavericks"] = "??????"
	teamMap["Denver Nuggets"] = "??????"
	teamMap["Detroit Pistons"] = "??????"
	teamMap["Golden State Warriors"] = "??????"
	teamMap["Houston Rockets"] = "??????"
	teamMap["Indiana Pacers"] = "??????"
	teamMap["Los Angeles Lakers"] = "??????"
	teamMap["LA Clippers"] = "??????"
	teamMap["Memphis Grizzlies"] = "??????"
	teamMap["Miami Heat"] = "??????"
	teamMap["Milwaukee Bucks"] = "??????"
	teamMap["Minnesota Timberwolves"] = "??????"
	teamMap["New Orleans Pelicans"] = "??????"
	teamMap["New York Knicks"] = "??????"
	teamMap["Oklahoma City Thunder"] = "??????"
	teamMap["Orlando Magic"] = "??????"
	teamMap["Philadelphia 76ers"] = "76???"
	teamMap["Phoenix Suns"] = "??????"
	teamMap["Portland Trail Blazers"] = "??????"
	teamMap["Sacramento Kings"] = "??????"
	teamMap["San Antonio Spurs"] = "??????"
	teamMap["Toronto Raptors"] = "??????"
	teamMap["Utah Jazz"] = "??????"
	teamMap["Washington Wizards"] = "??????"

	return teamMap
}

//Get the injuriers of the nba team
func getInjury(searchTeam string) (result []string) {

	if searchTeam == "Los Angeles Clippers" {
		searchTeam = "LA Clippers"
	}

	res, err := http.Get("https://www.espn.com/nba/injuries")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".Table__league-injuries").Each(func(i int, s *goquery.Selection) {
		team := s.Find(".injuries__teamName").Text()

		if strings.EqualFold(team, searchTeam) {

			s.Find(".Table__even").Each((func(i int, g *goquery.Selection) {
				name := g.Find(".AnchorLink").Text()
				status := g.Find(".col-stat").Text()
				comment := g.Find(".col-desc").Text()

				if status != "" {
					name = fmt.Sprintf("%-25s", name)
					status = fmt.Sprintf("%-15s", status)
					comment = fmt.Sprintf("%-5s", comment)
					injury := name + status + comment
					result = append(result, injury)
				}
			}))
		}
	})

	return result

}

//??????Injury???comment
func GetInjuryComment(searchTeam string) (result string) {

	if searchTeam == "Los Angeles Clippers" {
		searchTeam = "LA Clippers"
	}

	res, err := http.Get("https://www.espn.com/nba/injuries")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	teamMap := TeamInit()

	doc.Find(".Table__league-injuries").Each(func(i int, s *goquery.Selection) {
		team := s.Find(".injuries__teamName").Text()

		if strings.EqualFold(team, searchTeam) {
			result = result + teamMap[searchTeam] + "--"
			s.Find(".Table__even").Each((func(ii int, g *goquery.Selection) {

				name := g.Find(".AnchorLink").Text()
				status := g.Find(".col-stat").Text()
				comment := g.Find(".col-desc").Text()

				if status != "" {

					commentResult := sortComment(name, comment)
					result = result + commentResult

				}

			}))
		}
	})

	return result
}

//comment ??????
func sortComment(name, comment string) (result string) {
	switch {
	case strings.Contains(comment, "out"):
		result = result + name + "??????/"
	case strings.Contains(comment, "miss"):
		result = result + name + "??????/"
	case strings.Contains(comment, "will not play"):
		result = result + name + "??????/"
	case strings.Contains(comment, "won't"):
		result = result + name + "??????/"
	case strings.Contains(comment, "question"):
		result = result + name + "????????????/"
	case strings.Contains(comment, "doubtful"):
		result = result + name + "????????????/"
	case strings.Contains(comment, "day-to-day"):
		result = result + name + "?????????????????????/"
	case strings.Contains(comment, "probable"):
		result = result + name + "????????????/"
	default:
		result = result + name + "?????????????????????/"
	}

	return result
}

//Get the nba game of the day
func PKTeam() {
	startTime := time.Now()
	//??????UTC-4
	zone := time.FixedZone("", -4*60*60)
	today := time.Now()
	newTime := today.In(zone).Format("2006-01-02")

	//get data api url
	url := "https://in.global.nba.com/stats2/scores/daily.json?gameDate=" + newTime + "&locale=en&tz=%2B8&countryCode=TW#"
	// fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var (
		result                           Schedule
		HomeTeamComment, AwayTeamComment string
	)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if unmarshalErr := json.Unmarshal(body, &result); unmarshalErr != nil {
		panic(unmarshalErr)
	}

	var msg string
	var commentMsg string
	var layout string = "2006-01-02T15:04"
	teamMap := TeamInit()
	if result.Payload.Date.GameCount == "" {
		msg = msg + "?????? " + newTime + " ???????????? \n"
		fmt.Println(msg)
		fmt.Println("Spend Time:", time.Since(startTime))
		return
	}
	msg = msg + "?????? " + newTime + " ??? " + result.Payload.Date.GameCount + " ????????? \n"

	for i, v := range result.Payload.Date.Games {

		t, _ := time.Parse(layout, v.Profile.DateTimeEt)

		if v.AwayTeam.Profile.Name == "Clippers" {
			v.AwayTeam.Profile.City = "Los Angeles"
		}

		if v.HomeTeam.Profile.Name == "Clippers" {
			v.HomeTeam.Profile.City = "Los Angeles"
		}

		AwayTeam := v.AwayTeam.Profile.City + " " + v.AwayTeam.Profile.Name
		HomeTeam := v.HomeTeam.Profile.City + " " + v.HomeTeam.Profile.Name

		AwayTeamInjury := getInjury(AwayTeam)
		HomeTeamInjury := getInjury(HomeTeam)

		msg = msg + fmt.Sprint(i+1) + ". " + AwayTeam + "  " + t.Add(time.Hour*13).Format("15:04") + "  " + HomeTeam + "(???)  " + "\n"

		msg = msg + "\n  ---------------------------------\n"
		msg = msg + "  " + AwayTeam + " injury ?????? \n"

		if len(AwayTeamInjury) == 0 {
			msg = msg + "  ????????????\n\n"
			commentMsg = commentMsg + teamMap[AwayTeam] + "-????????? /"
		} else {
			for _, v := range AwayTeamInjury {
				msg = msg + "  " + v + "\n"
			}
			commentMsg = commentMsg + GetInjuryComment(AwayTeam) + "; "
		}

		msg = msg + "\n  ---------------------------------\n"

		msg = msg + "  " + HomeTeam + " injury ?????? \n"
		if len(HomeTeamInjury) == 0 {
			msg = msg + "  ????????????\n\n"
			commentMsg = commentMsg + teamMap[HomeTeam] + "-?????????"
		} else {
			for _, v := range HomeTeamInjury {
				msg = msg + "  " + v + "\n"
			}
			commentMsg = commentMsg + GetInjuryComment(HomeTeam)
			msg = msg + "\n"
		}

		AwayTeamDish := getDish(AwayTeam)
		HomeTeamDish := getDish(HomeTeam)

		msg = msg + "  " + AwayTeam + " ??????????????????: " + AwayTeamDish[AwayTeam] + "\n"
		msg = msg + "  " + HomeTeam + " ??????????????????: " + HomeTeamDish[HomeTeam] + "\n\n"
		commentMsg = commentMsg + AwayTeamComment + "   " + HomeTeamComment + "\n\n"

	}
	fmt.Println(commentMsg)
	fmt.Println(msg)
	fmt.Println("Spend Time:", time.Since(startTime))
}

//Get the nba game of the day , search by "StartTime"
func PKTeamOnStartTime(st string) {
	if len(st) != 5 {
		fmt.Println("??????????????????" + st + "???????????????eg: '11:00'(????????????) ")
		return
	}

	_, errCheckTime := time.Parse("15:04", st)
	if errCheckTime != nil {
		fmt.Println("??????????????????: ", errCheckTime)
		return
	}

	startTime := time.Now()
	//??????UTC-4
	zone := time.FixedZone("", -4*60*60)
	today := time.Now()
	newTime := today.In(zone).Format("2006-01-02")

	//get data api url
	url := "https://in.global.nba.com/stats2/scores/daily.json?gameDate=" + newTime + "&locale=en&tz=%2B8&countryCode=TW#"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var result Schedule

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if unmarshalErr := json.Unmarshal(body, &result); unmarshalErr != nil {
		panic(unmarshalErr)
	}

	var msg string
	var layout string = "2006-01-02T15:04"
	var count int = 0 //?????????????????????

	for _, v := range result.Payload.Date.Games {

		t, _ := time.Parse(layout, v.Profile.DateTimeEt)

		if v.AwayTeam.Profile.Name == "Clippers" {
			v.AwayTeam.Profile.City = "Los Angeles"
		}

		if v.HomeTeam.Profile.Name == "Clippers" {
			v.HomeTeam.Profile.City = "Los Angeles"
		}

		//???st == ????????????
		if st == t.Add(time.Hour*13).Format("15:04") {
			count++

			AwayTeam := v.AwayTeam.Profile.City + " " + v.AwayTeam.Profile.Name
			HomeTeam := v.HomeTeam.Profile.City + " " + v.HomeTeam.Profile.Name
			AwayTeamInjury := getInjury(AwayTeam)
			HomeTeamInjury := getInjury(HomeTeam)
			msg = msg + fmt.Sprint(count) + ". " + AwayTeam + "  " + t.Add(time.Hour*13).Format("15:04") + "  " + HomeTeam + "(???)  " + "\n"

			msg = msg + "\n  ---------------------------------\n"
			msg = msg + "  " + AwayTeam + " injury ?????? \n"

			if len(AwayTeamInjury) == 0 {
				msg = msg + "  ????????????\n\n"
			} else {
				for _, v := range AwayTeamInjury {
					msg = msg + "  " + v + "\n"
				}

			}

			msg = msg + "\n  ---------------------------------\n"

			msg = msg + "  " + HomeTeam + " injury ?????? \n"
			if len(HomeTeamInjury) == 0 {
				msg = msg + "  ????????????\n\n"
			} else {
				for _, v := range HomeTeamInjury {
					msg = msg + "  " + v + "\n"
				}
				msg = msg + "\n"
			}
			AwayTeamDish := getDish(AwayTeam)
			HomeTeamDish := getDish(HomeTeam)
			msg = msg + "  " + AwayTeam + " ??????????????????: " + AwayTeamDish[AwayTeam] + "\n"
			msg = msg + "  " + HomeTeam + " ??????????????????: " + HomeTeamDish[HomeTeam] + "\n\n"
		}

	}

	if count == 0 {
		msg = msg + "?????? " + st + " ????????????\n"
	} else {
		msg = "?????? " + st + " ??? " + fmt.Sprint(count) + " ????????? \n" + msg
	}

	fmt.Println(msg)
	fmt.Println("Spend Time:", time.Since(startTime))
}

//?????????????????????????????????
func getDish(searchTeam string) map[string]string {
	var (
		season string
	)

	year := time.Now()
	month := time.Now().Format("01")
	monthR, _ := strconv.ParseInt(month, 10, 64)
	if monthR < 7 {
		season = year.AddDate(-1, 0, 0).Format("06") + "-" + year.Format("06")
	} else {
		season = year.Format("06") + "-" + year.AddDate(1, 0, 0).Format("06")
	}
	yearR := year.Format("2006010215")

	// http://nba.titan007.com/cn/LetGoal.aspx?SclassID=1&matchSeason=2022-2023
	url := "https://nba.titan007.com/jsData/letGoal/" + season + "/l1.js?version=" + yearR

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var Team string

	r1, _ := regexp.Compile(";")
	r2 := r1.FindAllStringSubmatchIndex(string(sitemap), -1)

	Team = (string([]byte(sitemap[r2[0][0]+17 : r2[1][0]-1])))
	frontside, _ := regexp.Compile(`\[(.*?)]`)

	frontsideTeam := frontside.FindAllStringSubmatchIndex(string(Team), -1)

	matchMap := make(map[int64]string)
	result := make(map[string]string)

	for i := range frontsideTeam {
		//?????????????????????map????????????????????????
		TeamData := (Team[frontsideTeam[i][0]:frontsideTeam[i][1]])
		match, _ := regexp.Compile(",")
		matchR := match.FindAllStringSubmatchIndex(TeamData, -1)
		TeamNumber, _ := strconv.ParseInt(TeamData[1:matchR[0][0]], 10, 64) //TeamData[1:TeamData[matchR[0][0]-1]]

		if searchTeam == TeamData[matchR[2][1]+1:matchR[3][0]-1] {
			matchMap[TeamNumber] = TeamData[matchR[2][1]+1 : matchR[3][0]-1]
		}
	}

	data := (string([]byte(sitemap[r2[1][0]+20 : r2[2][0]-1])))
	frontsidedata := frontside.FindAllStringSubmatchIndex(string(data), -1)

	//?????????????????????
	for i := range frontsidedata {
		winPercentData := (data[frontsidedata[i][0]:frontsidedata[i][1]])
		match, _ := regexp.Compile(",")
		matchR := match.FindAllStringSubmatchIndex(winPercentData, -1)
		TeamNumber, _ := strconv.ParseInt(winPercentData[matchR[0][0]+1:matchR[1][0]], 10, 64)

		//??????Map??????????????????TeamNumber??????????????????result map?????????
		if _, ok := matchMap[TeamNumber]; ok {

			one := changeWinLose(winPercentData[matchR[11][0]+1 : matchR[12][1]-1])
			two := changeWinLose(winPercentData[matchR[12][0]+1 : matchR[13][1]-1])
			three := changeWinLose(winPercentData[matchR[13][0]+1 : matchR[14][1]-1])
			four := changeWinLose(winPercentData[matchR[14][0]+1 : matchR[15][1]-1])
			five := changeWinLose(winPercentData[matchR[15][0]+1 : matchR[15][1]+1])

			result[matchMap[TeamNumber]] = one + "," + two + "," + three + "," + four + "," + five
		}
	}

	return result

}

func changeWinLose(r string) (result string) {
	if r == "0" {
		result = "???"
	} else {
		result = "???"
	}
	return
}
