package library

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"service-all/app/init/global"
	"strconv"
	"sync"
	"time"
)

const bufferSize = 100

var (
	receiveTimeout = time.Duration(0)
	messageID      = 0
	lock           = sync.Mutex{}
)

type Client struct {
	Host           string                      // Host (probably "localhost").
	Port           int                         // Port (OBS default is 4444).
	Password       string                      // Password (OBS default is "").
	Conn           *websocket.Conn             // Underlying connection to OBS.
	receiveTimeout time.Duration               // Maximum blocking time for receiving request responses
	connected      bool                        // True until Disconnect is called.
	handlers       map[string]func(e Event)    // Event handlers.
	respQ          chan map[string]interface{} // Queue of received responses.
}

func GetMessageID() string {
	lock.Lock()
	messageID++
	id := strconv.Itoa(messageID)
	lock.Unlock()
	return id
}

func (c *Client) poll() {
	global.LOG.Info("started polling")

	for c.connected {
		m := make(map[string]interface{})
		if err := c.Conn.ReadJSON(&m); err != nil {
			if !c.connected {
				return
			} else if websocket.IsUnexpectedCloseError(err) {
				c.Disconnect()
			}
			global.LOG.Error("read from WS:", zap.Error(err))
			continue
		}

		if _, ok := m["message-id"]; ok {
			c.handleResponse(m)
		} else {
			c.handleEvent(m)
		}
	}
}

// handleEvent runs an event's handler if it exists.
func (c *Client) handleEvent(m map[string]interface{}) {
	t := m["update-type"].(string)

	switch t {
	case "Heartbeat":

		break
	case "SceneItemTransformChanged":

		break
	case "MediaRestarted":

		break
	case "MediaStarted":

		break
	case "SceneItemVisibilityChanged":

		break
	case "MediaEnded":

		break
	case "SourceDestroyed":

		break
	case "SceneItemSelected":

		break
	case "SceneItemRemoved":

		break
	case "SceneItemDeselected":

		break
	default:
		global.LOG.Info("data", zap.Any("m", m))
		break
	}

	//eventFn, ok := eventMap[t]
	//if !ok {
	//	global.LOG.Error("unknown event type:", zap.Any("data", m["update-type"]))
	//	return
	//}
	//event := eventFn()

	//handler, ok := c.handlers[t]
	//if !ok {
	//	return
	//}
	//
	//if err := mapToStruct(m, event); err != nil {
	//	Logger.Println("event handler failed:", err)
	//	return
	//}
}

// handleResponse sends a response into the queue.
func (c *Client) handleResponse(m map[string]interface{}) {
	//c.respQ <- m
	global.LOG.Info("data", zap.Any("m", m))
}

// Connected returns wheter or not the client is connected.
func (c Client) Connected() bool {
	return c.connected
}

// Disconnect closes the WebSocket connection.
func (c *Client) Disconnect() error {
	c.connected = false
	if err := c.Conn.Close(); err != nil {
		return err
	}
	return nil
}

// connectWS opens the WebSocket connection.
func connectWS(host string, port int) (*websocket.Conn, error) {
	url := fmt.Sprintf("ws://%s:%d", host, port)
	global.LOG.Info("connecting to", zap.String("url", url))
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Connect opens a WebSocket connection and authenticates if necessary.
func (c *Client) Connect() error {
	//c.handlers = make(map[string]func(Event))
	//c.respQ = make(chan map[string]interface{}, bufferSize)

	conn, err := connectWS(c.Host, c.Port)
	if err != nil {
		return err
	}
	c.Conn = conn

	// We can't use SendReceive yet because we haven't started polling.

	//reqGAR := NewGetAuthRequiredRequest()
	//if err = c.Conn.WriteJSON(reqGAR); err != nil {
	//	return err
	//}
	//
	//respGAR := &GetAuthRequiredResponse{}
	//if err = c.Conn.ReadJSON(respGAR); err != nil {
	//	return err
	//}
	//
	//if !respGAR.AuthRequired {
	//	Logger.Println("logged in (no authentication required)")
	//	c.connected = true
	//	go c.poll()
	//	return nil
	//}

	//auth := getAuth(c.Password, respGAR.Salt, respGAR.Challenge)
	//global.LOG.Info("auth:", zap.String("auth", auth))

	//reqA := NewAuthenticateRequest(auth)
	//if err = c.Conn.WriteJSON(reqA); err != nil {
	//	return err
	//}

	//respA := &AuthenticateResponse{}
	//if err = c.Conn.ReadJSON(respA); err != nil {
	//	return err
	//}
	//if respA.Status() != "ok" {
	//	return errors.New(respA.Error())
	//}

	global.LOG.Info("logged in (authentication successful)")
	c.connected = true
	go c.poll()
	return nil
}

// getAuth computes the auth challenge response.
func getAuth(password, salt, challenge string) string {
	sha := sha256.Sum256([]byte(password + salt))
	b64 := base64.StdEncoding.EncodeToString([]byte(sha[:]))

	sha = sha256.Sum256([]byte(b64 + challenge))
	b64 = base64.StdEncoding.EncodeToString([]byte(sha[:]))

	return b64
}
