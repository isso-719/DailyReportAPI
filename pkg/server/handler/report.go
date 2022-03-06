package handler

import (
	"DailyReportAPI/pkg/constant"
	"DailyReportAPI/pkg/http/response"
	"DailyReportAPI/pkg/server/domain/model"
	"DailyReportAPI/pkg/server/infra"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Report のモデルを定義
type reportGetResponse struct {
	ID      int       `json:"id"`
	Date    time.Time `json:"date"`
	Body    string    `json:"body"`
	Weather string    `json:"weather"`
}

// Report を取得する
func HandleGetReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reports, err := infra.SelectAllReports()
		if err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		response.Success(w, reports)
	}
}

// Report を登録する
func HandlePostReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// パラメータ body を取得する
		body := r.FormValue("body")

		weather, err := getWeather()
		if err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		// レポートを保存する
		report := model.Report{
			Body:    body,
			Weather: weather,
		}

		if err := infra.InsertReport(report); err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		// レスポンスを返す
		response.Success(w, nil)
	}
}

func HandleGetReportWithLastID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// パラメータ lastID を取得する
		lastID, err := strconv.Atoi(r.FormValue("lastID"))
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		// レポートを取得する
		report, err := infra.SelectReportWithLastID(lastID)
		if err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		// レスポンスを返す
		response.Success(w, report)
	}
}

func HandleUpdateReportWithLastID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// パラメータ lastID を取得する
		lastID, err := strconv.Atoi(r.FormValue("lastID"))
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		body := r.FormValue("body")

		// レポートを取得する
		report, err := infra.SelectReportWithLastID(lastID)
		if err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		report.Body = body

		// レポートを更新する
		if err := infra.UpdateReportWithLastID(lastID, report); err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		// レスポンスを返す
		response.Success(w, nil)
	}
}

func getWeather() (string, error) {
	// API を叩く
	res, err := http.Get(constant.WeatherEndpoint + constant.WeatherMainLocation)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// レスポンスを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// レスポンスをパースする
	var data model.WeatherAPIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	// 天気を返す
	return data.Forecasts[0].Telop, nil
}
