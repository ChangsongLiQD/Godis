package main

//hardcode table   struct redisCommand redisCommandTable[]
type Command func(c *Client, s *Server)
