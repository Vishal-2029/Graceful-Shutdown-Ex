## Graceful Shutdown Example 

- What is Graceful Shutdown?
`
    Graceful shutdown is the process of stopping an application safely by allowing it to complete ongoing tasks before exiting. This ensures:
  ✅ No data loss
  ✅ No corruption of ongoing processes
  ✅ Proper cleanup of resources (e.g., database connections, file handles, goroutines, etc.)

  Without a graceful shutdown, abruptly stopping a program (like using CTRL+C or killing the process) can cause incomplete transactions, corrupted files, or memory leaks.
`

## In this Example I have create a Two Function :

 - 1. Database Graceful Shutdown

  When the `main` Funtion is run then the database is connected. After Press the `ctrl+c` the database the disconnected.

  ## run the code

  ```bash
      go run main.go
  ```

  - 2. Goroutine Graceful Shutdown

    when the `main` function is run then the output is `Worker processing...`. After Press the `ctrl+c` Goroutines is the Stop.

      ## run the code

   ```bash
       go run goroutine.go
   ```

      
