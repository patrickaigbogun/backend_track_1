# Backend Track

This is a simple backend service built in Go that provides an API endpoint which returns user information in JSON format. The service includes a basic handler for returning the current time in UTC, along with other user details.

## Features

- Returns user details in JSON format.
- Current timestamp in UTC.
- Exposes a simple HTTP API endpoint.

## API Endpoint

- **GET `https://goservice-c85-7000.prg1.zerops.app/api/clasify-number?number=4`**: Returns a JSON object 


### Example Response:

```json
{"digit_sum":4,"fun_fact":"4 is the smallest composite number that is equal to the sum of its prime factors.","is_perfect":"false","is_prime":"false","number":4,"properties":["armstrong","odd"]}
```

## Requirements

- Go (version 1.18+)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/patrickaigbogun/backend_track_1.git
   ```

2. Navigate to the project directory:

   ```bash
   cd backend_track_1
   ```
   
3. Run the application:

   ```bash
   air
   ```

   The server will start on port `7000`.

## Usage

Once the server is running, you can access the endpoint by visiting:

```
http://localhost:7000/api/classify-number?number=<any number of your choice>
```

This will return the JSON response with a fun fact about the number you entered.

## File Structure

.
├── go.mod
├── main.go
├── README.md
├── tmp
│   ├── build-errors.log
│   └── main
└── zerops.yml
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Additional Notes:
- **Dependencies:** The Go module (`go.mod`) file manages the dependencies for the project.
- **Error Handling:** Basic error handling is in place to ensure that the JSON is properly marshaled before being sent as the HTTP response.
- **Hiring:** Experienced with Golang and need a job? [go here](https://hng.tech/hire/golang-developers)
