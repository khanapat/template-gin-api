package account

import (
	"net/http"
	"template-gin-api/internal/handler"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type registerAccountCopyFrom struct {
	repo registerAccountCopyFromRepo
}

func NewRegisterAccountCopyFrom(repo registerAccountCopyFromRepo) *registerAccountCopyFrom {
	return &registerAccountCopyFrom{
		repo: repo,
	}
}

func (s *registerAccountCopyFrom) Handler(c *handler.Ctx) {
	accounts := []CreateAccount{
		{
			Id:        uuid.NewString(),
			FirstName: "Leanne",
			LastName:  "Graham",
			Email:     "Sincere@april.biz",
			Balance:   10.25,
			RoleId:    1,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Ervin",
			LastName:  "Howell",
			Email:     "Shanna@melissa.tv",
			Balance:   91.75,
			RoleId:    1,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Clementine",
			LastName:  "Bauch",
			Email:     "Nathan@yesenia.net",
			Balance:   3,
			RoleId:    1,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Patricia",
			LastName:  "Lebsack",
			Email:     "Julianne.OConner@kory.org",
			Balance:   10.25,
			RoleId:    1,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Chelsey",
			LastName:  "Dietrich",
			Email:     "Lucio_Hettinger@annie.ca",
			Balance:   10.25,
			RoleId:    1,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Dennis",
			LastName:  "Schulist",
			Email:     "Karley_Dach@jasper.info",
			Balance:   1074,
			RoleId:    2,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Kurtis",
			LastName:  "Weissnat",
			Email:     "Telly.Hoeger@billy.biz",
			Balance:   347,
			RoleId:    2,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Nicholas",
			LastName:  "Runolfsdottir",
			Email:     "Sherwood@rosamond.me",
			Balance:   32,
			RoleId:    2,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Glenna",
			LastName:  "Reichert",
			Email:     "Chaim_McDermott@dana.io",
			Balance:   32,
			RoleId:    2,
		},
		{
			Id:        uuid.NewString(),
			FirstName: "Clementina",
			LastName:  "DuBuque",
			Email:     "Rey.Padberg@karina.biz",
			Balance:   6,
			RoleId:    2,
		},
	}

	start := time.Now()
	if err := s.repo.CreateAccountWithCopyFrom(c.Context, accounts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, time.Since(start).Seconds())
}
