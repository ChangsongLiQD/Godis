package cmd

import "Godis/app"

//hardcode table   struct redisCommand redisCommandTable[]
type Command func(c *app.Client, s *app.Server)

func populateCommandTable() {

}
