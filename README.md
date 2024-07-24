<!-- ABOUT THE PROJECT -->
## About The Project
This project is a backend system designed to support an administrative interface for managing waste types. It provides comprehensive functionality for administrators to oversee and update various waste types. It also tracks and reports key usage metrics, including the total number of application usages and the accuracy of classifications. 

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
1. Change directory to main folder
   ```sh
   cd admin-backend/cmd/main
   ```
2. Run the application
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


[firebase-url]: https://firebase.google.com/
[firebase.com]: https://img.shields.io/badge/firebase-black?style=for-the-badge&logo=firebase&logoColor=color
[go-url]: https://go.dev/
[go.com]: https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[docker-url]: https://www.docker.com/
[docker.com]: https://img.shields.io/badge/docker-black?style=for-the-badge&logo=docker&logoColor=color

