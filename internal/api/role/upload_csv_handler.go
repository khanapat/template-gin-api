package role

import (
	"encoding/csv"
	"io"
	"net/http"
	"template-gin-api/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type uploadCSV struct{}

func NewUploadCSV() *uploadCSV {
	return &uploadCSV{}
}

func (s *uploadCSV) Handler(c *handler.Ctx) {
	fileHeader, err := c.FormFile("whitelist")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.Read()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res := make([]interface{}, 0)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		row := make(map[string]interface{})
		for i, value := range record {
			row[records[i]] = value
		}

		var data ProgressCsv
		if err := mapstructure.Decode(row, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		res = append(res, data)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// https://www.golinuxcloud.com/go-map-to-struct/
// records, err := reader.Read()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	res := make([]interface{}, 0)
// 	for {
// 		record, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"error": err.Error(),
//			})
//			return
// 		}
// 		row := make(map[string]interface{})
// 		for i, value := range record {
// 			row[records[i]] = value
// 		}
// 		var data ProgressCsv
// 		if err := mapstructure.Decode(row, &data); err != nil {
// 			c.Set("csv.error", err.Error())
// 			c.Next()
// 			return
// 		}
// 		res = append(res, data)
// }
// [
//         {
//             "emp_id": "667125",
//             "quest_name": "Joiner Challenge",
//             "action_date": "2023-11-27 0:00:00"
//         },
//         {
//             "emp_id": "667368",
//             "quest_name": "Joiner Challenge",
//             "action_date": "2023-11-27 0:00:00"
//         }
// ]

// records, err := reader.ReadAll()
// [
//     "emp_id",
//     "quest_name",
//     "action_date"
// ],
// [
//     "667125",
//     "Joiner Challenge",
//     "2023-11-27 0:00:00"
// ]
