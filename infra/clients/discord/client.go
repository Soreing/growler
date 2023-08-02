package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	baseUrl string
	token   string
}

func NewClient(
	optn *Options,
) *Client {
	return &Client{
		baseUrl: optn.BaseUrl,
		token:   optn.Token,
	}
}

func (c *Client) CreateDMhannel(
	ctx context.Context,
	userId string,
) (DMChannel, error) {
	msg := CreateDM{
		RecipientId: userId,
	}

	dat, err := json.Marshal(msg)
	if err != nil {
		return DMChannel{}, err
	}

	req, err := http.NewRequest(
		"POST",
		c.baseUrl+"/users/@me/channels",
		bytes.NewBuffer(dat),
	)
	if err != nil {
		return DMChannel{}, err
	}

	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("Content-Type", "application/json")

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return DMChannel{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return DMChannel{}, err
	} else if resp.StatusCode != 200 {
		return DMChannel{}, fmt.Errorf(string(body))
	}

	var dm DMChannel
	err = json.Unmarshal(body, &dm)
	if err != nil {
		return DMChannel{}, err
	}

	return dm, nil
}

func (c *Client) CreateMessage(
	ctx context.Context,
	channelId string,
	message []byte,
) error {
	req, err := http.NewRequest(
		"POST",
		c.baseUrl+"/channels/"+channelId+"/messages",
		bytes.NewBuffer(message),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bot "+c.token)
	req.Header.Set("Content-Type", "application/json")

	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err == nil && resp.StatusCode != 200 {
		body, ioerr := ioutil.ReadAll(resp.Body)
		if ioerr != nil {
			err = ioerr
		} else {
			err = fmt.Errorf(string(body))
		}
	}
	if err != nil {
		return err
	}

	return nil
}
