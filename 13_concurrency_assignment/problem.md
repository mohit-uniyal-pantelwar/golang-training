### Problem Statement

You are tasked with implementing a **multithreaded URL status checker** for a list of URLs. Your program should concurrently check the status of each URL and report the results. The program must use **Go routines**, **Wait Groups**, and **Channels** to achieve concurrency.

### Requirements:

1. **Input**:
    - The program should accept a list of URLs.
    - You can hard-code the list of URLs or get them as input arguments.

2. **URL Status Checker**:
    - Implement a function `checkURLStatus(url string) (string, error)` that checks if a URL is reachable (HTTP status 200).
    - The function should return "Success" if the HTTP status code is 200, and an error message if the URL is unreachable or returns an error.

3. **Concurrency**:
    - You need to check the status of each URL **concurrently** using **Go routines**.
    - Use a **WaitGroup** to wait for all Go routines to finish.
    - Use **Channels** to collect the results (success or failure for each URL).

4. **Output**:
    - Once all URLs are checked, the program should output the status of each URL.
    - The output should list each URL, whether it was successful or failed, and any error message if applicable.

5. **Timeout**:
    - Implement a timeout mechanism for each URL check. If a URL takes more than 3 seconds to respond, it should be considered as failed.

6. **Concurrency Limits**:
    - Your program should allow a maximum of 5 concurrent Go routines to check URLs at the same time. Implement this using a buffered channel.

### Additional Considerations:
- Handle any network errors that might occur when making HTTP requests.
- Ensure that the program gracefully handles situations where all URLs fail, and provide a meaningful message to the user.
- Consider edge cases like timeouts, non-200 status codes, and invalid URLs.


---

This assignment should challenge your understanding of Go concurrency mechanisms and how to coordinate tasks in a concurrent program efficiently. Good luck!