package main

func DelCommand(c *Client, s *Server) {
	if c.Argc != 2 {
		c.Buff = []byte("invalid get usage")
	}

	var result bool
	DoProcess(func() {
		result = s.Db.DelKey(c.Argv[0].Ptr.(string))
	})

	if result {
		c.Buff = RespIntSuccess
	} else {
		c.Buff = RespIntFail
	}
}
