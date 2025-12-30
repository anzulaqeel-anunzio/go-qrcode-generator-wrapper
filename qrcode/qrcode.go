// Developed for Anunzio International by Anzul Aqeel. Contact +971545822608 or +971585515742. Linkedin Profile: linkedin.com/in/anzulaqeel

/*
 * Developed for Anunzio International by Anzul Aqeel
 * Contact +971545822608 or +971585515742
 */

package qrcode

import (
	go_qrcode "github.com/skip2/go-qrcode"
)

// RecoveryLevel represents the error recovery level
type RecoveryLevel int

const (
	// Low recovery level (7%)
	Low RecoveryLevel = iota
	// Medium recovery level (15%)
	Medium
	// High recovery level (25%)
	High
	// Highest recovery level (30%)
	Highest
)

// Generate creates a QR code PNG image
func Generate(content string, size int, filename string) error {
	return go_qrcode.WriteFile(content, go_qrcode.Medium, size, filename)
}

// GenerateWithRecovery creates a QR code with a specific recovery level
func GenerateWithRecovery(content string, size int, filename string, level RecoveryLevel) error {
	var qLevel go_qrcode.RecoveryLevel
	switch level {
	case Low:
		qLevel = go_qrcode.Low
	case Medium:
		qLevel = go_qrcode.Medium
	case High:
		qLevel = go_qrcode.High
	case Highest:
		qLevel = go_qrcode.Highest
	default:
		qLevel = go_qrcode.Medium
	}
	
	return go_qrcode.WriteFile(content, qLevel, size, filename)
}

// Developed for Anunzio International by Anzul Aqeel. Contact +971545822608 or +971585515742. Linkedin Profile: linkedin.com/in/anzulaqeel
