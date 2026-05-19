package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//go:embed quotes.json
var quotesFS embed.FS

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var themes = map[string]struct {
	top     string
	bottom  string
	content string
}{
	"default": {"┌┐", "└┘", "│"},
	"bold":    {"┏┓", "┗┛", "┃"},
	"rounded": {"╭╮", "╰╯", "│"},
	"double":  {"╔╗", "╚╝", "║"},
	"minimal": {"+=", "=+", "|"},
}

var quotes []Quote

func init() {
	data, err := quotesFS.ReadFile("quotes.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading quotes:", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(data, &quotes); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing quotes:", err)
		os.Exit(1)
	}

	if len(quotes) == 0 {
		fmt.Fprintln(os.Stderr, "No quotes found")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
}

func color(code string, text string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", code, text)
}

func printBox(text, author, themeName string, noColor bool) {
	theme, ok := themes[themeName]
	if !ok {
		theme = themes["default"]
	}

	topLeft := theme.top[:len(theme.top)/2]
	topRight := theme.top[len(theme.top)/2:]
	botLeft := theme.bottom[:len(theme.bottom)/2]
	botRight := theme.bottom[len(theme.bottom)/2:]

	maxLen := 60
	lines := wrapText(text, maxLen)

	topLine := topLeft + strings.Repeat("─", maxLen+2) + topRight
	botLine := botLeft + strings.Repeat("─", maxLen+2) + botRight

	if noColor {
		fmt.Println(topLine)
		for _, line := range lines {
			fmt.Printf("%s %-*s %s\n", theme.content, maxLen, line, theme.content)
		}
		fmt.Println(botLine)
		fmt.Printf("  — %s\n", author)
	} else {
		topC := color("32;1", topLine)
		botC := color("32;1", botLine)
		vBarC := color("32;1", theme.content)
		fmt.Println(topC)
		for _, line := range lines {
			fmt.Printf("%s %-*s %s\n", vBarC, maxLen, line, vBarC)
		}
		fmt.Println(botC)
		colored := color("36", fmt.Sprintf("  — %s", author))
		fmt.Println(colored)
	}
}

func wrapText(text string, width int) []string {
	if len(text) <= width {
		return []string{text}
	}

	var lines []string
	words := strings.Fields(text)
	current := ""

	for _, word := range words {
		if len(current)+len(word)+1 <= width {
			if current != "" {
				current += " "
			}
			current += word
		} else {
			if current != "" {
				lines = append(lines, current)
			}
			current = word
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func printHelp() {
	fmt.Println("motivational-quotes - Display inspiring quotes")
	fmt.Println()
	fmt.Println("Usage: motivational-quotes [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --theme=<name>   Box style: default, bold, rounded, double, minimal")
	fmt.Println("  --no-color      Disable colored output")
	fmt.Println("  --help          Show this help message")
	fmt.Println()
	fmt.Println("Available themes:")
	for name := range themes {
		fmt.Printf("  %s\n", name)
	}
}

func main() {
	var themeName string
	var noColor bool
	var showHelp bool

	flag.StringVar(&themeName, "theme", "default", "Box style theme")
	flag.BoolVar(&noColor, "no-color", false, "Disable colors")
	flag.BoolVar(&showHelp, "help", false, "Show help")

	flag.Parse()

	if showHelp {
		printHelp()
		return
	}

	// Auto-disable color if not a TTY
	if !isTerminal() {
		noColor = true
	}

	if _, ok := themes[themeName]; !ok {
		fmt.Fprintf(os.Stderr, "Unknown theme: %s\n", themeName)
		fmt.Fprintln(os.Stderr, "Available themes: default, bold, rounded, double, minimal")
		os.Exit(1)
	}

	idx := rand.Intn(len(quotes))
	quote := quotes[idx]
	printBox(quote.Text, quote.Author, themeName, noColor)
}

func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}