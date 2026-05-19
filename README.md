# Motivate CLI

A simple command-line tool that displays random motivational quotes with beautiful ASCII box framing.

## Installation

### From Source
```bash
git clone https://github.com/demo/motivational-quotes-cli.git
cd motivational-quotes-cli
go build -o motivate .
./motivate
```

### Via Homebrew (after release)
```bash
brew install motivate
```

## Usage

```bash
motivate                    # Show a random quote (default theme)
motivate --theme=rounded   # Use rounded box style
motivate --theme=double    # Use double-line box style
motivate --no-color        # Disable colors (for piping)
motivate --help            # Show help
```

## Available Themes
- `default` - Single-line ASCII
- `bold` - Bold ASCII
- `rounded` - Rounded corners
- `double` - Double-line box
- `minimal` - Minimal style

## Features
- 29 curated motivational quotes
- 5 box themes with colors (TTY only)
- Automatic color disable for non-TTY output
- Random quote selection
- Embed directive for single-binary distribution

## Release
Tags trigger GitHub Actions to:
1. Build for macOS (ARM64, AMD64) and Linux
2. Create Homebrew formula
3. Upload release artifacts