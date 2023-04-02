# :book: Build Bazel with GoLang

This project it's to understand how works the Bazel with GoLang Projects

## :hammer: Setup and Tools used

- Bazel 6.0.0
- Gazelle
- GoLang 1.20
- Docker

This list does not complete yet, will increase with the time.

## :eyeglasses: Goals

- Understading the system build with Bazel and GoLang

- Implementation using Docker to configuration in the enviroment

- Add tests to use with Bazel

## :rocket: Run

- If you want to use the bazel to build and run the project locally on your machine
follow these instructions:
```console
user@machine:~$ bazel run gazelle
```
This command will execute the basic configuration to work on the project. **If you make any changes to the code, it is mandatory to execute this command again.**
- Update the dependencies:
```console
user@machine:~$ bazel run gazelle-update-repos
```
- Finally, start the project. Note that the start of the project it's based on the name of binary in **BUILD** file.
```console
user@machine:~$ bazel run bazel-project-example
```
- Open another terminal and execute the command to call the endpoint: */hello*
```console
user@machine:~$ curl http://localhost:3333/hello
```
You will see in the terminal: **Hello World with Chi and Bazel**