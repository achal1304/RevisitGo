A **JWT (JSON Web Token)** is an open standard used for securely transmitting information between parties as a JSON object. It is commonly used for authentication and information exchange. JWTs are often used to verify the identity of users and ensure that the information they send (such as API requests) is not tampered with.

### How JWT Works:

JWTs typically consist of **three parts**:

1. **Header**: 
   - The header typically consists of two parts: 
     - The type of the token (which is JWT).
     - The signing algorithm used (e.g., HMAC SHA256, RSA).

   Example of a header:
   ```json
   {
     "alg": "HS256",
     "typ": "JWT"
   }
   ```

2. **Payload**: 
   - The payload contains the claims. Claims are statements about an entity (usually the user) and additional metadata. 
   - There are three types of claims:
     - **Registered claims**: Predefined claims that are not mandatory but recommended (e.g., `iss` for issuer, `exp` for expiration time, etc.).
     - **Public claims**: Custom claims agreed upon by the parties involved.
     - **Private claims**: Claims used within a specific context, typically custom to the application.

   Example of a payload (the claim part):
   ```json
   {
     "sub": "1234567890",
     "name": "John Doe",
     "iat": 1516239022
   }
   ```

3. **Signature**:
   - The signature is used to verify that the sender of the JWT is who it says it is, and to ensure that the message wasn't tampered with along the way.
   - The signature is created by taking the encoded header, the encoded payload, a secret key, and the algorithm specified in the header. The resulting signature is then appended to the header and payload.
   
   Example of the signing process:
   - Combine the Base64-encoded header and payload.
   - Apply the signing algorithm (e.g., HMAC SHA256) to the combination of the header, payload, and a secret key.
   - Base64 encode the resulting signature.

### Structure of JWT:

A JWT token is typically a long string with three parts, each separated by a period (`.`):

```
<base64url-encoded header>.<base64url-encoded payload>.<base64url-encoded signature>
```

For example:
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

### Detailed Breakdown of the JWT Flow:

1. **Authentication (Client-side)**:
   - The user logs in (typically with a username and password).
   - The server validates the credentials. If valid, the server creates a JWT, signing it with a secret key or private key, and returns the token to the client.
   
2. **Storing the Token**:
   - The client typically stores the JWT in **localStorage**, **sessionStorage**, or as an **HTTP-only cookie**.
   
3. **Making API Requests**:
   - The client includes the JWT in the Authorization header of every request:
     ```
     Authorization: Bearer <JWT>
     ```
   - The server receives the token, extracts it from the header, and verifies it using the signature and the secret key. The server will decode the payload and check the claims (e.g., `exp` for expiration, `sub` for user identity, etc.).

4. **Verification (Server-side)**:
   - The server verifies the token by:
     - Checking the **signature** to ensure the token has not been tampered with.
     - Checking the **claims** (e.g., ensuring the token is not expired).
   
5. **Token Expiration**:
   - A JWT can include an `exp` claim, which specifies the expiration time of the token. If the token is expired, the server rejects the request, and the user may need to log in again or refresh the token.

6. **Refreshing Tokens (Optional)**:
   - In many cases, the JWT is short-lived for security reasons. To avoid having users log in repeatedly, a **refresh token** mechanism is used. The refresh token can be used to obtain a new JWT without needing the user to log in again.

### Benefits of JWT:
- **Stateless**: The server does not need to store any session information because all the data is contained in the JWT itself.
- **Scalable**: Since the server doesn't need to maintain sessions, it's easier to scale horizontally with multiple server instances.
- **Compact**: JWTs are compact and can be easily passed in HTTP headers, making them ideal for mobile applications, single-page apps (SPAs), and APIs.
- **Self-contained**: JWTs can store user information (e.g., user ID, role) in the payload, reducing the need to query the database on every request.

### Security Considerations:
- **Use HTTPS**: Always transmit JWTs over HTTPS to prevent man-in-the-middle attacks.
- **Secret Key**: The secret key (or private key, in the case of asymmetric encryption) used to sign the JWT should be stored securely.
- **Expiration**: Always set an expiration (`exp`) for your tokens to limit the window of time an attacker can use a stolen token.
- **Revocation**: JWTs cannot be easily revoked since they're self-contained. You might need to implement token revocation manually (e.g., through blacklisting).

### Example of JWT Signing in Go (using `github.com/dgrijalva/jwt-go`):

```go
package main

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

// Secret key for signing the JWT (use a more secure method in production)
var mySigningKey = []byte("mysecret")

// GenerateJWT generates a JWT token
func GenerateJWT() (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"sub":  "1234567890",               // User ID
		"name": "John Doe",                 // User name
		"iat": time.Now().Unix(),           // Issued at time
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiration time (24 hours)
	}

	// Create the token with the specified claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}

func main() {
	// Generate JWT
	token, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	// Output the generated token
	fmt.Println("Generated JWT:", token)
}
```

### Explanation:
- We create a JWT using the **HS256** signing method.
- The `claims` include standard claims like the user ID (`sub`), user name (`name`), issue time (`iat`), and expiration (`exp`).
- The `SignedString` method signs the token with the secret key.

### Decoding and Verifying JWT:

To verify and decode the token, you can use the following code:

```go
func ParseJWT(tokenString string) (*jwt.Token, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is what we expect (HS256 in this case)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return mySigningKey, nil
	})
	return token, err
}
```

This function verifies the JWT's signature and ensures it's valid.

### Conclusion:

JWTs are a powerful and compact method for securely transmitting information between parties. They are widely used for **stateless authentication** and can store user-specific information like user roles and permissions. By using JWTs, you can build scalable and secure authentication systems in your applications.