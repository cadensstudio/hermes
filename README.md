# Hermes - Google Fonts Downloader

Hermes is a command-line interface (CLI) application built in Go that simplifies the process of downloading web-optimized Google Font files in the WOFF2 format. Hermes takes an opinionated approach, aiming to download the single variable font file if available; otherwise, it downloads each individual font weight file separately. Additionally, Hermes generates the necessary CSS code to easily integrate the downloaded fonts into your project.

## Features

- **Efficient Font Downloads**: Hermes optimizes the download process by retrieving only the necessary font files in WOFF2 format.
  
- **Variable Font Support**: When available, Hermes prioritizes the download of a single variable font file for efficiency.

- **CSS Integration**: The tool generates CSS code, making it seamless to incorporate the downloaded fonts into your project.

## Getting Started

### Prerequisites

Hermes requires a Google Fonts API Key to run commands. Obtain your key [here](https://console.cloud.google.com/apis/credentials).

### Installation

#### Install using Homebrew

```bash
brew tap cadensstudio/tap && brew install hermes
```

#### Download the binary files

See releases for binary files.

### Usage

Hermes currently provides the following commands

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

3. Setup your Google Fonts API Key in a `.env` file:

    ```bash
    cp .env.example .env
    ```

4. Build Hermes:

    ```bash
    go build
    ```

5. Run Hermes:

    ```bash
    ./hermes get roboto
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
