## What to do


- Tokenize PAN, PIN (raw), expDate
- Tokenize biller id (payment source)
- Generate a fingerprint (optional)

## Data security

- How to secure the data
- How to log data access
- Persist and store the data securely (TODO how?)


## Data encryption

- Where to do data encryption:
    - Client
    - Server


## Interfaces

- Biller ID generate a new card token
- Return:
    - card_tokenized
    - fingerprint
    - biller use the card tokenized subsequently

- Inputs:
    - card_tokenized
    - fingerprint
    - amount
    - biller_id

Processing:

Scenario 1:
    - Offload the payment token
    - complete the request (in server)

Scenraio 2:
    - Offload the payment token
    - Return response to the client
    - Make the client complete the request