package utils

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func GetClientIp(c *fiber.Ctx) string {
	IpAddress := c.Get("X-Real-Ip", "")
	if IpAddress == "" {
		IpAddress = c.Get("X-Forwarded-For")
		splits := strings.Split(IpAddress, ", ")
		return splits[0]
	}
	if IpAddress == "" {
		return "127.0.0.1"
	}

	return IpAddress
}

func GetClientDeviceName(c *fiber.Ctx) string {
	return c.Get("User-Agent")
}

func GetClientInfo(c *fiber.Ctx) map[string]string {
	return map[string]string{
		"IpAddress":  GetClientIp(c),
		"DeviceName": GetClientDeviceName(c),
	}
}
