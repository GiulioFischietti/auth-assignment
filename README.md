# Auth Assignment
Short assignment on authentication.

A detailed doc regarding showing design choices will be uploaded in a .pdf file soon.

## Quick Setup

1. Clone this repository 

2. In auth-service rename `.env.example` to `.env` (don't change it, it is already setup)

3. In protected-service rename `.env.example` to `.env` (don't change it, it is already setup)

4. Move in the root folder of the project and Execute the bash script (from wsl linux shell if you're on Windows)  

`./scripts/start.sh`

If for any reason it fails (for example due to windows CRLF in .sh), follow the steps below: you will execute the bash commands manually.

## Manual Setup

1. Go to the root folder /auth-assignment and Create /keys folders for each service:

`mkdir -p auth-service/keys`

`mkdir -p protected-service/keys`


2. Generate a private key

`openssl genpkey -algorithm RSA -out auth-service/keys/private.pem -pkeyopt rsa_keygen_bits:2048`

3. Generate the public key

`openssl rsa -pubout -in auth-service/keys/private.pem -out auth-service/keys/public.pem`

4. Copy the public key in protected-service/keys

`cp auth-service/keys/public.pem protected-service/keys/public.pem`

5. Run the project with docker compose

`docker compose up --build`

## Database seed

Database seed is performed automatically at docker compose startup.
