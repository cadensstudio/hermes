![hermes](https://github.com/cadensstudio/hermes/assets/54109914/b040715a-d2b6-4416-ab97-6050ab6a9f1a)

# Hermes - Instantly Download Google Fonts

Hermes is a command-line interface (CLI) application built in Go that accelerates the process of downloading web-optimized Google Font files in the WOFF2 format. Hermes takes an opinionated approach by downloading variable font files, if available. Otherwise, Hermes downloads each individual font weight file separately. Additionally, Hermes generates the necessary CSS code to easily integrate the downloaded fonts into your project.

## Features

- **Efficient Font Downloads**: Optimizes the download process by retrieving only the necessary font files in WOFF2 format.
  
- **Variable Font Support**: Prioritizes downloading a single variable font file (when available) for efficiency.

- **CSS Integration**: Generates CSS code, making it seamless to incorporate the downloaded fonts into your project.

## Getting Started

### Prerequisites

Hermes requires a Google Fonts API Key. Obtain your key [here](https://console.cloud.google.com/apis/credentials).

### Installation

#### Install using Homebrew

```bash
brew tap cadensstudio/tap && brew install hermes
```

#### Download the binary

See [releases](https://github.com/cadensstudio/hermes/releases).

### Usage

Ensure you set your Google Fonts API key by running `export GFONTS_KEY=<YOUR KEY>`.

Run `hermes --help` to view all available hermes commands:

```bash
Usage:
  hermes [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Downloads web-optimized font files for a specified font family
  help        Help about any command
  list        Lists the 10 most trending Google Fonts

Flags:
  -h, --help   help for hermes

Use "hermes [command] --help" for more information about a command.
```

## Contributions

Contributions to Hermes are welcome! Feel free to open issues, submit pull requests, or provide feedback to improve the tool.

### Local Dev Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/cadensstudio/hermes.git
    ```

2. Navigate to the project directory:

    ```bash
    cd hermes
    ```

3. Set your Google Fonts API Key:

    ```bash
    export GFONTS_KEY=<YOUR KEY>
    ```

4. Build Hermes:

    ```bash
    go build
    ```

5. Run Hermes:

    ```bash
    ./hermes get inter
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
