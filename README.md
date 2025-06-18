#  Car Hire Web App – Go + Docker + Jenkins CI/CD

This project is a simple, containerized Go web application built for a fictional car rental company: **Tech Bro Car Hire**. It was developed as part of a real-world DevOps simulation to demonstrate how to integrate code versioning, containerization, and Continuous Integration/Continuous Deployment (CI/CD) pipelines using **Jenkins** on a Virtual Private Server (VPS).

---

##  Features

-  Built with **Go 1.24.4**
-  Clean HTTP server with multiple routes:
  - `/` – Home
  - `/about` – About Tech Bro Car Hire
  - `/contact` – Contact Information
  - `/admin` – Admin dashboard (placeholder)
-  Dockerfile for lightweight containerization
-  Jenkinsfile for full CI/CD automation
-  Clean file structure with `.gitignore` and `.dockerignore`

---

##  Folder Structure
car-hire-app/
├── main.go # Entry point for Go web server
├── go.mod, go.sum # Go module and dependency files
├── Dockerfile # Docker instructions for building the app
├── .dockerignore # Clean image by ignoring unnecessary files
├── .gitignore # Git cleanup rules
├── jenkins/
│ └── Jenkinsfile # Jenkins pipeline for CI/CD automation
├── README.md # Project documentation

---
##  How to Run Locally

To run the app on your machine (without Docker):

```bash
go run main.go
Visit: http://localhost:8080

How to Build & Run with Docker
# Build Docker image
docker build -t car-hire-app .

# Run container (port 8080 is exposed)
docker run -d -p 8080:8080 --name car-hire-app car-hire-app
Visit: http://localhost:8080 or http://<your-vps-ip>:8080

CI/CD with Jenkins

This project includes a Jenkinsfile for automating the following:

Cloning the GitHub repository

Building the Go app

Creating a Docker image

Running the app inside a container on the VPS

This pipeline is executed using a Jenkins server running on the same VPS.


Author
Tech Bro
Cloud & DevOps Engineer in training 



License
This project is open-source and free to use under the MIT License