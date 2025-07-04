## 🔐 Password Manager (Go CLI)

A simple and secure CLI password manager written in Go. It allows you to store, manage, generate, and encrypt passwords locally using AES encryption. Built with modular packages and the Cobra CLI framework.

## 🚀 Features

- Add new password entries with username, and password
- Delete password
- List all saved password
- Generate strong random passwords
- Encrypt and decrypt the entire password storage file
- File-based local storage (JSON format)
- Built with modular Go packages
- Cobra-powered CLI interface

## 🧪 Usage
Commands:
- ``add [name] [password]``	Add a new password entry
- ``delete [name]`` delete	Delete an entry by title
- ``list``  list	List all stored titles
- ``generate [name]``	Generate and save a strong random password


## 🔒 Security
- AES encryption is used to secure the storage file.
- Never share your encryption key.
- Store the key securely (e.g., environment variables or password vault).
