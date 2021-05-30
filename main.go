package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"os"
	"strings"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	var port int

	flag.IntVar(&port, "p", 8000, "Server port number")
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	
	var username string
	var password string

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file := path.Join(cwd, "config.toml")
	
	if _, err := os.Stat(file); err == nil {
		// config file exists so lets read it
		config, err := os.Open("config.toml")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(config)

		scanner.Split(bufio.ScanWords)
		count := 0
		
		for scanner.Scan() {
			if (count > 6) {
				log.Fatal("Incorrectly formatted config.toml file")
			}
			switch count {
				case 2:
					username = strings.TrimPrefix(strings.TrimSuffix(scanner.Text(), "\""), "\"")
				case 5:
					password = strings.TrimPrefix(strings.TrimSuffix(scanner.Text(), "\""), "\"")
			}
			count++
		}

		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				username: password,
			},
		}))
	}

	app.Get("/", func(c *fiber.Ctx) error {
		data, err := Asset("assets/index.html")
		if err != nil {
			fmt.Println(err)
		}

		c.Type("html", "utf-8")
		return c.SendString(string(data))
	})

	app.Get("/list-services", func(c *fiber.Ctx) error {
		out, err := exec.Command("systemctl", "list-units", "--state=running", "--no-pager").Output()

		if err != nil {
			fmt.Printf("%s", err)
		}

		return c.SendString(string(out[:]))
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		var messageType int = 1
		var message []byte
		var err error
		var cmd *exec.Cmd

		if c.Params("id") == "all" {
			cmd = exec.Command("journalctl", "-b", "-f", "-n", "10", "-o", "json")
		} else {
			cmd = exec.Command("journalctl", "-b", "-u", fmt.Sprintf("%s.service", c.Params("id")), "-f", "-n", "10", "-o", "json")
		}

		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			message = []byte(scanner.Text())
			if err = c.WriteMessage(messageType, message); err != nil {
				log.Println(err)
			}
		}

		cmd.Wait()
	}, websocket.Config{
		WriteBufferSize: 8192,
	}))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}