package main

func SetCommand(c *Client, s *Server) {
	if (c.Argc-1)%2 != 0 {
		// Response error
		c.Buff = []byte("invalid set usage")
		return
	}

	for i := 0; i+1 < c.Argc-1; i = +2 {
		key := c.Argv[i].Ptr.(string)
		obj := &c.Argv[i+1]
		s.Db.SetKey(key, obj)
	}

	c.Buff = OkResp
}

func GetCommand(c *Client, s *Server) {
	if c.Argc > 2 {
		c.Buff = []byte("invalid get usage")
	}

	data := s.Db.GetKey(c.Argv[0].Ptr.(string))
	if data != nil {
		c.Buff = []byte(data.Ptr.(string))
	} else {
		c.Buff = []byte("wtf")
	}
}
