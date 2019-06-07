package Env

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type YearConfig struct {
	Year    int      `yaml:"year"`
	Terms   []Term   `yaml:"terms"`
	Periods []Period `yaml:"periods"`
}

type Term struct {
	TermName  string    `yaml:"termName"`
	StartDate time.Time `yaml:"startDate"`
	EndDate   time.Time `yaml:"endDate"`
}

type Period struct {
	PeriodName string `yaml:"periodName"`
	StartTime  string `yaml:"startTime"`
	EndTime    string `yaml:"endTime"`
}

var ypath = "year-config.yml"

var yearConfig YearConfig

func GetYearConfig() *YearConfig {
	return &yearConfig
}

func parseYearConfig() error {
	fmt.Println("parsing yaml school config file")
	file, err := ioutil.ReadFile(ypath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(file, &yearConfig); err != nil {
		return err
	}
	return nil
}
