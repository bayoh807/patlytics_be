package services

import (
	"backend/data"
	"backend/resource"
	"backend/utils"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize/english"
	"regexp"
	"sort"
	"time"
)

type reportServ struct {
}

var ReportServ reportServ

func (s *reportServ) Analyze(val *resource.ReportReq) (interface{}, error) {

	if patents := data.Resource.GetData("patents", *val.Patent); patents == nil {
		return map[string]interface{}{}, nil
	} else if companies := data.Resource.GetData("companies", *val.Company); companies == nil {
		return map[string]interface{}{}, nil
	} else {
		prompt := s.makeMessage(*val.Company, *val.Patent, companies.(map[string]interface{}), patents.(map[string]interface{}))
		for count := 0; count < 2; count++ {
			res, err := utils.Claude.SendMessage(prompt)
			if err != nil || res == "" {
				return "Sorry, something is wrong. Please try it later.", nil
			} else if toJson, err := s.toJson(res); err == nil {
				dateObj := time.Date(2023, time.September, 15, 12, 0, 0, 0, time.UTC)
				date := dateObj.Format("2006-01-02")
				return resource.ReportRes{
					Company: *val.Company,
					Patent:  *val.Patent,
					Date:    date,
					Analyze: toJson,
				}, nil
			} else {
				count++
				return "Sorry, something is wrong. Please try it later.", nil
			}
		}
		return "Sorry, something is wrong. Please try it later.", nil
	}
}

func (s *reportServ) toJson(jsonString interface{}) (interface{}, error) {
	var anyType interface{}
	str, ok := jsonString.(string)
	if !ok {
		return nil, fmt.Errorf("%s", "This is not a JSON.")
	}

	json.Unmarshal([]byte(str), &anyType)
	return anyType, nil

}

func (s *reportServ) makeMessage(name, ID string, company, patent map[string]interface{}) string {

	mapToString := func(obj map[string]interface{}) string {
		toJSON, _ := json.Marshal(obj)
		return string(toJSON)
	}
	prompt := fmt.Sprintf("Help me analyze all products of %s and evaluate their relevance and potential infringement regarding patent: %s. "+
		"company : %s"+
		"patent : %s\n"+
		"Please return only the following JSON string based on your analysis and judgment. Replace {advice} with your recommendation, and do not include any additional text or explanation:\n"+
		"{\n  \"top_infringing_products\": [\n    {\n      \"product_name\": \"\",\n      \"infringement_likelihood\": \"High\",\n      \"relevant_claims\": [],\n      \"explanation\": \"\",\n      \"specific_features\": []\n    }\n  ],\n  \"overall_risk_assessment\": \"{advice}\"\n}",
		name, ID, mapToString(company), mapToString(patent))
	return prompt
}

func (s *reportServ) SearchKeyword(req resource.SearchReq) interface{} {

	key := english.PluralWord(2, req.Type, "")

	items := data.Resource.GetData(key)
	res := []string{}
	pattern := fmt.Sprintf("(?i)%s", req.Keyword)
	regex, _ := regexp.Compile(pattern)

	for index, _ := range items.(map[string]map[string]interface{}) {
		if regex.MatchString(index) {
			res = append(res, index)

		}
	}
	sort.Strings(res)
	return res
}
