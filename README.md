![Gopher](./dist/static/img/twitter_header_photo_2.png)

Quickly create "pads" and share with others

[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/r/justkato/freepad)
[![Ko-Fi](https://img.shields.io/badge/Ko--fi-F16061?style=for-the-badge&logo=ko-fi&logoColor=white)](https://ko-fi.com/justkato)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

![MariaDB](https://img.shields.io/badge/MariaDB-003545?style=for-the-badge&logo=mariadb&logoColor=white)

# **FreePad**
**FreePad** is a simple `Go` project to help you juggle temporary notes that you might wanna pass from one device to another, or from a person to another with memorable and easy to communicate online "Pads".

The project is absolutely free to use, you can extend the code and even contribute, I am more than happy to be corrected in my horrible beginner code.

The current maintainer and creator is `Kato Twofold`

![Gopher](./dist/static/img/banner_prerequisites.png)

Before getting started there are a couple things you should configure before proceeding, such as the database storage type and a couple limits, now if you really want to you can skip these but it's better to know what you're running as to not wake up with a not-so-nice surprise.

![Gopher](./dist/static/img/banner_environment.png)

The `.env` file contains all of the available options and you should use it to change those said variables, these are really important to customizing and self hosting this experience for yourself.

If you need any help with any setting you can always open an issue over on github and get help from me.

If you are barely getting started with hosting your own services, or even Sys admin stuff in general or writing code my suggestion is to just copy `.env` and leave it as is until you get it running with the defaults running fine, afterwards you can play with it a little and who knows, maybe even get to learn something!

![Gopher](./dist/static/img/banner_building.png)


## From Source
Building from source isn't exactly recommended as it's a hasle 
```bash

# Clone the repo
git clone https://github.com/JustKato/FreePad FreePad

# Get in it!
cd FreePad

# Install golang
sudo apt install golang # Obviously use your distro's package manager

# Run the build Script
./build.sh

# Check out the ./dist folder
cd ./dist

# Make sure you change settings here
cp ../.env.example ./.env

# Run the program
./freepad

```

## Running the Binary
```bash
# Download the latest version of FreePad
freepad.1.0.3.zip

# Extract to wherever
unzip freepad.1.0.3.zip

# Get into the directory
cd FreePad

# ( Optionaly but recommended ) Edit the .env file
vim .env

# Run the program
./freepad

```

## Starting with Docker-Compose [ WIP ]
```bash
# Copy the example docker-compose file to anywhere
wget https://raw.githubusercontent.com/JustKato/FreePad/master/docker-compose.example.yaml

# Yoink the example .env file while we're at it
wget https://raw.githubusercontent.com/JustKato/FreePad/master/.env.example

# Rename the files
mv docker-compose.example.yaml docker-compose.yaml
mv .env.example .env

# ! Please take a look at the files and edit them before running
docker-compose up -d;
```