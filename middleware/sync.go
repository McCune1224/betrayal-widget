package middleware

import (
	"log"
	"sort"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
)

type SyncMiddleware struct {
	models *data.Models
}

func NewSyncMiddleware(db *sqlx.DB) *SyncMiddleware {
	models := data.NewModels(db)
	return &SyncMiddleware{
		models: models,
	}
}

func (s *SyncMiddleware) SyncGameInfo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("game_id") == "" {
			return next(c)
		}
		players, err := s.models.Players.GetAllComplexByGameID(c.Param("game_id"))
		if err != nil {
			return err
		}
		players = util.OrderComplexPlayers(players)
		diff := util.BulkCalculateLuck(players)
		for i := range players {
			if players[i].P.Luck != diff[i].P.Luck {
				err := s.models.Players.UpdateProperty(diff[i].P.ID, "luck", diff[i].P.Luck)
				if err != nil {
					log.Println(err)
					return c.Redirect(302, "/")
				}
			}
		}

		alliances, err := s.models.Alliances.GetAllByGame(c.Param("game_id"))
		if err != nil {
			return err
		}
		sort.Slice(alliances, func(i, j int) bool {
			return alliances[i].Name < alliances[j].Name
		})

		c.Set("players", players)
		c.Set("alliances", alliances)
		return next(c)
	}
}
