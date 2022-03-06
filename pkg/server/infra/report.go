package infra

import (
	"DailyReportAPI/pkg/db"
	"DailyReportAPI/pkg/server/domain/model"
	"fmt"
)

func init() {
	db.Conn.AutoMigrate(&model.Report{})
}

func SelectAllReports() (*model.Reports, error) {
	var reports model.Reports
	err := db.Conn.Order("id desc").Find(&reports.Reports).Error
	return &reports, err
}

func SelectReportWithLastID(id int) (*model.Report, error) {
	var report model.Report
	// 最後から数えて id 番目のレコードを取得する
	err := db.Conn.Order("id desc").Limit(1).Offset(id).Find(&report).Error

	if report.CreatedAt.IsZero() {
		err := fmt.Errorf("no record")
		return nil, err
	}

	return &report, err
}

func InsertReport(report model.Report) error {
	return db.Conn.Create(&report).Error
}

func UpdateReportWithLastID(lastID int, report *model.Report) error {
	var reportOld model.Report
	if err := db.Conn.Order("id desc").Limit(1).Offset(lastID).Find(&reportOld).Error; err != nil {
		return err
	}
	reportOld.Body = report.Body
	return db.Conn.Save(&reportOld).Error
}
