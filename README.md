# Convert_Post_API_to_Get_API
A Golang Project which convert Post Request to Get Request

# Why to use this Service?

  Actually when you can not change request formats of an API , you need to use a converter, In this project there is a need . As an example there is an POST API that you can not change it to a GET API, thus you need to put minimal changes for the customer.
so you Convert POST request to a new GET request.

# What to do?

## Step 1

Build thew docker file

```
docker build -t PostApi-server:V1 .
```

## Step 2

Run the containers using ```docker compose``` command

```
docker compose up -d
```

## Step 3 

Send a Sample curl or postman POST request

```
curl --location --request POST 'http://127.0.0.1:8080/PostSend' -H 'UserName: username' -H 'Password: password' -H "Content-type: application/json" -d '{"Message":"Hello World...:D","PhoneNumber":"XXXXXXXXXXXXX"}'
```

## Note

This sample os used to send a specific message to a phone number
