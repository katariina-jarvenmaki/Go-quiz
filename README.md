# go-quiz
A simple Go quiz game that can be run through the executable or within a Go environment.

```bash
APP: kjc-int
DIRECTORY STRUCTURE: /opt/kjc/int/go-quiz
```

## Installation

1. **Clone the Repositories**:

```bash
git clone git@github.com:KajanCreative/golang-setup.git
git clone git@github.com:KajanCreative/go-quiz.git
```

2. **Copy to the Directory**:
Copy the contents of the go-quiz project into the golang-setup directory at:
```path
/opt/kjc/int/golang-setup/src/app/
```

3. **Navigate to directory***
```path
cd /opt/kjc/int/golang-setup
```

4. **Make the Script Executable**:
```bash
chmod +x ./runexec
```

5. **Add to PATH (Optional but recommended for ease of use)**:
Open your shell configuration file (~/.bashrc, ~/.zshrc, etc.):

```bash
nano ~/.bashrc  # or use your preferred editor
```

Add the following line to include the directory in your PATH:
```bash
export PATH="$PATH:/path/to/runexec"
```

Save the file and reload the configuration:
```bash
source ~/.bashrc  # or source the appropriate file
```

## Usage (golang-setup commands)
Run the script followed by the command you wish to execute:
```bash
./runexec <command>
Replace <command> with the specific one-liner command you want to execute.
```

## Examples (golang-setup commands)

Destroy all docker containers
```bash
./runexec destroy-docker-containers
Output: MSG: All docker containers are destroyed
```

Install or update Docker on Ubuntu -machine
```bash
./runexec install-docker-for-ubuntu
Output: MSG: Docker should now be installed and ready to use
```

Install golang into docker container
```bash
./runexec install-golang
Output: MSG: Install golang into docker container
```

Go inside the container and run the app
```bash
docker run -it --entrypoint /bin/bash golang-app
Output: MSG: Go inside the container
```

Run the app inside the container
```bash
go run main.go
Output: MSG: Run the app inside the container
```

Make an executable inside the container
```bash
go build main.go
Output: MSG: Make an executable of the app
```

Download the executable from the container
```bash
mkdir -p ~/golang-output
docker cp <container-id>:/app/main ~/golang-output/
ls -l ~/golang-output
Output: MSG: Build a node app docker container
```

### Folder structure:
```bash
/opt/kjc/int/go-quiz/
├── go.mod
├── go.sum
├── main                # Executable version of the app
├── main.go             # Main file on app / the quiz
├── package-lock.json   
├── package.json        # Docker package
└── README.md           # Readme
```