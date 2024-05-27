# GRBG (Garbage)

<p align="center">
    <img src="logo.webp" alt="GRBG Logo" width="500" height="500">
</p>

GRBG is a Go-based fuzzer designed for CTFs (Capture The Flag competitions). It is a lightweight tool with no external dependencies, making it easy to use and deploy.

## Active Development

Currently, GRBG has one exploiter implemented, which is the format string exploiter. This exploiter allows you to perform format string attacks during CTFs. We are actively working on adding more exploiters to enhance the capabilities of GRBG.

## Features

- **No Dependencies**: GRBG does not rely on any external libraries or packages, ensuring a seamless setup process.
- **Exploiter Interface**: The fuzzer can be easily expanded and customized using the exploiter interface, allowing you to tailor it to your specific needs.

## Installation

To install GRBG, follow these steps:

1. Clone the GRBG repository:

   ```shell
   git clone https://github.com/your-username/GRBG.git
   ```

2. Build the project:

   ```shell
   cd GRBG
   go build
   ```

3. Run GRBG:

   ```shell
   ./GRBG -success <success_flag> -fail <fail_flag> -bin <relative_path_to_executable>
   ```

## Usage

To use GRBG, simply execute the binary and provide the necessary input flags. For example:

- success flag: `./GRBG -success <success_flag>`
- fail flag: `./GRBG -fail <fail_flag>`
- binary flag: `./GRBG -bin <relative_path_to_executable>`

## Contributions

Contributions to GRBG are welcome! If you would like to contribute, please follow these steps:

1. Fork the GRBG repository on GitHub.
2. Clone your forked repository to your local machine:

   ```shell
   git clone https://github.com/your-username/GRBG.git
   ```

3. Create a new branch for your contribution:

   ```shell
   git checkout -b my-contribution
   ```

4. Make your changes and commit them:

   ```shell
   git commit -m "Add my contribution"
   ```

5. Push your changes to your forked repository:

   ```shell
   git push origin my-contribution
   ```

6. Open a pull request on the original GRBG repository to submit your contribution.
