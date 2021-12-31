package client

type AppClient struct {
	Title string
}

func (c *AppClient) String() string {
	return c.Title
}