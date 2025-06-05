# 📷 OpenMVGO - Go Wrapper for OpenMVG and OpenMVS

![OpenMVGO](https://img.shields.io/badge/OpenMVGO-v1.0.0-blue.svg) ![GitHub Releases](https://img.shields.io/github/release/mamofbi/openmvgo.svg) ![Go](https://img.shields.io/badge/Go-1.16+-blue.svg)

Welcome to **OpenMVGO**, a powerful Go wrapper for the OpenMVG and OpenMVS libraries. This project simplifies the use of photogrammetry tools, making it easier for developers to integrate these libraries into their applications. With OpenMVGO, you can streamline your workflows in 3D reconstruction and related fields.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Introduction

Photogrammetry is the science of making measurements from photographs. OpenMVG (Multiple View Geometry) and OpenMVS (Multi-View Stereo) are two libraries that excel in this domain. OpenMVGO serves as a Go wrapper for these libraries, allowing developers to leverage their capabilities in a Go environment.

This project is ideal for those who want to create applications that require 3D modeling, reconstruction, and other photogrammetry-related tasks. By using OpenMVGO, you can access the full power of OpenMVG and OpenMVS while enjoying the simplicity and efficiency of Go.

## Features

- **Easy Integration**: Seamlessly integrate OpenMVG and OpenMVS into your Go applications.
- **Command-Line Interface**: Use a simple CLI for quick operations.
- **Docker Support**: Run your applications in isolated environments using Docker.
- **Comprehensive Documentation**: Access clear documentation to help you get started quickly.
- **Active Development**: The project is actively maintained and updated.

## Installation

To get started with OpenMVGO, you need to download and execute the latest release. You can find the releases [here](https://github.com/mamofbi/openmvgo/releases). Follow these steps:

1. Visit the releases page.
2. Download the appropriate binary for your system.
3. Execute the binary.

Make sure you have Go installed on your machine. If you haven't installed Go yet, follow the instructions on the [official Go website](https://golang.org/doc/install).

## Usage

Using OpenMVGO is straightforward. Here’s a simple example to help you get started:

1. **Import the Library**:

   ```go
   import "github.com/mamofbi/openmvgo"
   ```

2. **Initialize the Library**:

   ```go
   mvgo := openmvgo.New()
   ```

3. **Run Photogrammetry Tasks**:

   You can run various tasks such as feature extraction, matching, and 3D reconstruction. Here’s an example of running a feature extraction task:

   ```go
   mvgo.ExtractFeatures("path/to/image.jpg")
   ```

4. **View Results**:

   After running your tasks, you can view the results directly in your application or export them to files.

For more detailed examples and API references, please check the [Documentation](#documentation).

## Documentation

Comprehensive documentation is available to help you navigate through the features and functionalities of OpenMVGO. You can find it in the `docs` folder of this repository or visit our [Wiki](https://github.com/mamofbi/openmvgo/wiki).

## Contributing

We welcome contributions from the community! If you would like to contribute to OpenMVGO, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch to your fork.
5. Create a pull request to the main repository.

Please ensure that your code follows the project's coding standards and that you include tests for new features.

## License

OpenMVGO is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

For questions, suggestions, or feedback, please reach out to the maintainers:

- [Your Name](mailto:your.email@example.com)

You can also check the releases section for updates and new features. Visit the releases page [here](https://github.com/mamofbi/openmvgo/releases) to stay informed.

---

We hope you find OpenMVGO useful for your photogrammetry projects. Happy coding!