package main

import "strconv"

func DelCommand(c *Client, s *Server) {
	if c.Argc != 2 {
		c.Buff = []byte("invalid get usage")
	}

	var result bool
	DoProcess(func() {
		result = s.Db.DelKey(c.Argv[0].Ptr.(string))
	})

	c.Buff = GetIntResultResponse(result)
}

func ExpireCommand(c *Client, s *Server) {
	if c.Argc != 3 {
		c.Buff = []byte("invalid get usage")
	}
	var result bool
	var err error
	DoProcess(func() {
		key := c.Argv[0].Ptr.(string)
		exist := s.Db.ExistsKey(key)
		if !exist {
			result = false
			return
		}

		time, parseErr := strconv.ParseInt(c.Argv[1].Ptr.(string), 10, 64)
		if parseErr != nil {
			err = parseErr
			return
		}

		result = s.Db.SetKeyExpireTimeBySeconds(key, time)
	})

	if err != nil {
		c.Buff = GetErrorResponse([]byte("value is not an integer or out of range"))
	} else {
		c.Buff = GetIntResultResponse(result)
	}
}
