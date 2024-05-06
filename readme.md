# API Stresser

The API Stresser is a simple tool written in Go (Golang) for stress testing APIs by sending a large number of concurrent requests. This tool helps assess the performance and scalability of APIs under heavy load conditions.

## Features

- Sends a configurable number of requests to a specified API endpoint concurrently.
- Measures response times and overall execution time.
- Handles errors gracefully and provides feedback on completion or timeouts.

## Usage

1. **Clone Repository**: Clone this repository to your local machine.

2. **Set API Endpoint**: Open the `main.go` file and set the `apiEnd` variable to the URL of the API endpoint you want to stress test.

3. **Adjust Total Requests**: Modify the `totalRequest` variable to specify the total number of requests you want to send.

4. **Run the Tool**: Execute the `main.go` file to start the stress testing tool.

```bash
go run main.go
```
## Review Results

![Alt Text](test1.mp4)
![Alt Text](test2.mp4)



Upon completion, the tool will print response codes for each request, and it will display the total time taken for all requests to finish. You can review these results to assess the performance and reliability of your API under stress.

## Contributing

Contributions are welcome! Feel free to submit issues, feature requests, or pull requests to improve this tool.

## License

This project is licensed under the [MIT License](LICENSE).
