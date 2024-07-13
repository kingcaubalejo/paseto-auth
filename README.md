# Qualitech PASETO Auth Example

This project demonstrates a simple Go web server with three main routes: login, protected, and unprotected. It utilizes PASETO (Platform-Agnostic Security Tokens) v2 for secure tokenization and authentication.

## Routes

1. **/login**:
   - Handles user login requests.
   - Generates and returns a PASETO token upon successful authentication.

2. **/protected**:
   - A protected endpoint that requires a valid PASETO token for access.
   - Uses middleware to authenticate requests based on the token.

3. **/unprotected**:
   - An open endpoint that can be accessed without any authentication.

## PASETO v2 Integration

- **Token Creation**:
  - Generates PASETO tokens using a symmetric key.
  - Includes custom payloads for additional data within the token.

- **Token Verification**:
  - Validates incoming PASETO tokens to ensure they are legitimate.
  - Checks for token expiration and custom claims.

## Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/yourusername/qualitech-paseto-auth.git
   cd qualitech-paseto-auth

2. **Install Dependencies**:
    ```sh
    go mod init qualitech.paseto-auth
    go get github.com/o1egl/paseto

## Usage

1. **Start Server**:
    ```sh
    go run main.go


## License
This project is licensed under the MIT License.