package ws

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"gobase/config"
	"net/http"
	"sync"
)

var (
	wsServer  *WSServer
	mu        sync.Mutex
	MaxWSConn int
)

func GetWSServer() *WSServer {
	mu.Lock()
	if wsServer == nil {
		wsServer = newWSServer()
	}
	mu.Unlock()
	return wsServer
}

type WSServer struct {
	clients    map[string]*websocket.Conn
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

func newWSServer() *WSServer {
	MaxWSConn = config.GetConfig().Websocket.MaxWSConn
	return &WSServer{
		clients:    make(map[string]*websocket.Conn),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (ws *WSServer) Start() {
	for {
		select {
		case conn := <-ws.register:
			// Use address as cliend ID
			clientAddr := conn.RemoteAddr().String()
			if _, connected := ws.clients[clientAddr]; !connected {
				ws.mu.Lock()
				ws.clients[clientAddr] = conn
				ws.mu.Unlock()
				go ws.handleMessage(conn) // start a goroutine to handle messages sent from this client
			}

		case client := <-ws.unregister:
			clientAddr := client.RemoteAddr().String()
			ws.mu.Lock()
			delete(ws.clients, clientAddr)
			ws.mu.Unlock()
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Perform a check here in production code
		// Allow connections from any origin for simplicity
		return true
	},
}

// WebSocket endpoint handler
func (ws *WSServer) HandleConnection(c *gin.Context) {
	// Limit the maximum number of connections
	if len(ws.clients) >= MaxWSConn {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many connections"})
		log.Println("Max ws connections reached.")
		return
	}

	// Upgrade initial GET request to a WebSocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err)
		return
	}
	ws.register <- conn
}

func (ws *WSServer) CloseConnection(clientID string) error {
	conn, hasConn := ws.clients[clientID]
	if !hasConn {
		return nil
	}
	err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Server closed connection."))
	if err != nil {
		log.Errorf("Fail to close ws connection, error: %v", err)
	}
	ws.unregister <- conn
	return nil
}

func (ws *WSServer) handleMessage(conn *websocket.Conn) {
	defer func() {
		ws.unregister <- conn // unregister the client when the message handling function finishes
		err := conn.Close()
		if err != nil {
			log.Println("Close ws connection error:", err)
		}
	}()
	count := 0
	for { // loop forever to read message from this client
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
		count++
	}
}

func (ws *WSServer) sendMessage(clientID string, data []byte) error {
	conn, hasConn := ws.clients[clientID]
	if !hasConn {
		return errors.New("WS connection not found")
	}
	err := conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Errorf("Error when sending message to ws: %v", err)
		return err
	}
	return nil
}
