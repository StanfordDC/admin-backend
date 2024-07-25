<a name="readme-top" id="readme-top"></a>
<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About the Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#run-locally">Run locally</a></li>
        <li><a href="#run-using-docker">Run using Docker</a></li>
      </ul>
    </li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project
The project is designed specifically for administrators to efficiently manage waste types and gain insights into mobile app usage.

The backend provides comprehensive functionality for administrators to manage waste types and the logic to track mobile app usage metrics

### Built With
* [![Go][go.com]][go-url]
* [![Firebase][firebase.com]][firebase-url]
* [![Docker][docker.com]][docker-url]


<!-- GETTING STARTED -->
## Getting Started
### Prerequisites
1. Create a Firebase project.
2. Add Firebase to Android application.
3. Download the google-services.json file and rename it to serviceAccountKey.json.
4. Move the serviceAccountKey.json file to the cmd/main folder.

### Installation
Clone the repo
   ```sh
   git clone https://github.com/StanfordDC/admin-backend.git
   ```
### Run locally
1. Install Go
2. Change directory to main folder
   ```sh
   cd admin-backend/cmd/main
   ```
3. Run the application
   ```sh
   go run main.go
   ```

### Run using Docker
1. Change directory to root
   ```sh
   cd admin-backend
   ```
2. Build the docker image
   ```sh
   docker build -t admin-backend ./
   ```
3. Run the docker container
   ```js
   docker run -p 8080:8080 admin-backend
   ```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

[firebase-url]: https://firebase.google.com/
[firebase.com]: https://img.shields.io/badge/firebase-black?style=for-the-badge&logo=firebase&logoColor=color
[go-url]: https://go.dev/
[go.com]: https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[docker-url]: https://www.docker.com/
[docker.com]: https://img.shields.io/badge/docker-black?style=for-the-badge&logo=docker&logoColor=color

