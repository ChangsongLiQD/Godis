package main

func SetCommand(c *Client, s *Server) {
	if (c.Argc-1)%2 != 0 {
		// Response error
		c.Buff = GetErrorResponse([]byte("invalid set usage"))
		return
	}

	DoProcess(func() {
		for i := 0; i+1 < c.Argc-1; i = +2 {
			key := c.Argv[i].Ptr.(string)
			obj := &c.Argv[i+1]
			s.Db.SetKey(key, obj)
		}
	})

	c.Buff = GetStringResponse(RespOk)
}

func GetCommand(c *Client, s *Server) {
	if c.Argc > 2 {
		c.Buff = []byte("invalid get usage")
	}

	var data *Object
	DoProcess(func() {
		data = s.Db.GetKey(c.Argv[0].Ptr.(string))
	})

	if data != nil {
		c.Buff = GetBulkBytesResponse([]byte(data.Ptr.(string)))
	} else {
		c.Buff = nil
	}
}
