package main

import (
        "fmt"
        "net/http"
        "time"
)

func main () {
    fmt.Println("DeepSeek Backup Server starting...")


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprintf(w, `
            <!DOCTYPE html>
            <html>
            <head><title>Backup Server</title></head>
            <body>
                    <h1>DeepSeek Backup Server</h1>
			<p>Status: <span style="color: green;">Running</span></p>
			<p>Time: %s</p>
			<p>Made by: jemus666</p>
		</body>
		</html>
        `, time.Now().Format("2006-01-02 15:04:05"))
    })

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "healthy", "timestamp": "%s"}`, 
			time.Now().Format(time.RFC3339))
	})
	
	fmt.Println("Server listening on :8080")
	fmt.Println("Open http://localhost:8080 in browser")
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
            fmt.Printf("Server error: %v\n", err)
    }        
}
