package api

import (
	"encoding/json"
	"errors"
	"github.com/xjh22222228/github-oauth2/utils"
	"net/http"
	"sort"
	"sync"
	"time"
)

func getRepos(r *http.Request) ([]UserRepo, error) {
	query := r.URL.Query()
	u := query.Get("id")
	userRepo := make([]UserRepo, 0)
	if u == "" {
		return userRepo, errors.New("ID 不能为空")
	}
	res := utils.Fetch(&utils.FetchConf{
		Method: http.MethodGet,
		Url: "https://api.github." +
			"com/users/" + u + "/repos?per_page=100&page=1",
		R: r,
	})
	if err := json.Unmarshal([]byte(res), &userRepo); err != nil {
		return userRepo, errors.New("暂无数据")
	}
	return userRepo, nil
}

var wg sync.WaitGroup

func getCommit(repoName string, r *http.Request) ([]Commit, error) {
	query := r.URL.Query()
	id := query.Get("id")
	commit := make([]Commit, 0)
	res := utils.Fetch(&utils.FetchConf{
		Method: http.MethodGet,
		Url: "https://api.github." +
			"com/repos/" + id + "/" + repoName + "/commits",
		R: r,
	})
	if err := json.Unmarshal([]byte(res), &commit); err != nil {
		return commit, errors.New("暂无数据")
	}
	return commit, nil
}

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	utils.Cors(w)
	var openIssuesCount, starCount = 0, 0
	repos, err := getRepos(r)
	if err != nil {
		utils.Body(w, utils.Stringify(Response{
			Message: err.Error(),
			Data:    nil,
			Status:  http.StatusBadRequest,
		}))
		return
	}

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].StargazersCount > repos[j].StargazersCount
	})

	// 最喜欢编程语言
	likeMap := make(map[string]int)
	likeLanguages := make([]LikeCount, 0)
	for _, item := range repos {
		if item.Language != "" {
			likeMap[item.Language] += 1
		}
	}
	for k, v := range likeMap {
		likeLanguages = append(likeLanguages, LikeCount{
			Name:  k,
			Count: v,
		})
	}
	sort.Slice(likeLanguages, func(i, j int) bool {
		return likeLanguages[i].Count > likeLanguages[j].Count
	})

	// 所有提交信息
	commits := make([]Commit, 0)
	wg.Add(len(repos))
	for _, item := range repos {
		starCount += item.StargazersCount
		openIssuesCount += item.OpenIssuesCount
		go func(item UserRepo) {
			defer wg.Done()
			c, err := getCommit(item.Name, r)
			if err == nil {
				commits = append(commits, c...)
			}
		}(item)
	}

	wg.Wait()

	// 最早开始时间 06:00 - 11:00
	var beforeTime *time.Time
	// 最晚开始时间 00:00 - 05:59
	var afterTime *time.Time
	// 每日提交数量
	dayCommitCountMap := make(map[string]int)
	dayCommits := make([]DateCount, 0)
	// 每小时提交数量
	hourCommitCountMap := make(map[int]int)
	hourCommits := make([]HourCount, 0)

	for _, item := range commits {
		t := item.Commit.Committer.Date
		if t == nil {
			return
		}
		h := t.Hour()
		m := t.Minute()
		k := t.Format(time.RFC3339)

		dayCommitCountMap[k] += 1
		hourCommitCountMap[h] += 1
		if h >= 6 && h <= 11 {
			if beforeTime == nil {
				beforeTime = t
			} else {
				if h <= beforeTime.Hour() && m < beforeTime.Minute() {
					beforeTime = t
				}
			}
		}
		if h >= 0 && (h <= 5 && m <= 59) {
			if afterTime == nil {
				afterTime = t
			} else {
				if h >= afterTime.Hour() && m > afterTime.Minute() {
					afterTime = t
				}
			}
		}
	}
	for k, v := range dayCommitCountMap {
		dayCommits = append(dayCommits, DateCount{
			Date:  k,
			Count: v,
		})
	}
	sort.Slice(dayCommits, func(i, j int) bool {
		return dayCommits[i].Count > dayCommits[j].Count
	})
	for k, v := range hourCommitCountMap {
		hourCommits = append(hourCommits, HourCount{
			Hour:  k,
			Count: v,
		})
	}
	sort.Slice(hourCommits, func(i, j int) bool {
		return hourCommits[i].Count > hourCommits[j].Count
	})
	cLen := len(dayCommits)
	if cLen >= 10 {
		cLen = 10
	}

	utils.Body(w, utils.Stringify(Response{
		Message: "OK",
		Status:  http.StatusOK,
		Data: Map{
			"repos":           repos,
			"likeLanguages":   likeLanguages,
			"beforeTime":      beforeTime,
			"afterTime":       afterTime,
			"dayCommits":      dayCommits[:cLen],
			"hourCommits":     hourCommits,
			"starCount":       starCount,
			"openIssuesCount": openIssuesCount,
			"commitCount":     len(commits),
		},
	}))
}
