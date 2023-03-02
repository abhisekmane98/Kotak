package main
	
	import (
	                "io"
	                "fmt"
	                "log"
	                "net/http"
	                "os"
	                //"errors"
	                "context"
	                "database/sql"
	    "github.com/go-sql-driver/mysql"
	)
	
	func home(w http.ResponseWriter, r *http.Request) {
	                w.Write([]byte("It's Homepage"))
	}
	
	func Ksrp(w http.ResponseWriter, req *http.Request) {
	
	                if req.Method != http.MethodPost {
	                                w.Header().Set("Allow", "POST")
	                                w.Header().Set("Content-Type", "application/json")
	                                w.WriteHeader(http.StatusMethodNotAllowed)
	                                w.Write([]byte("Method Not allowed"))
	                                return
	                }
	
	                body, err := io.ReadAll(req.Body)
	                if err != nil {
	                                w.Write([]byte("Call Failed"))
	                                return
	                                log.Fatal(err)
	                }
	
	                var file *os.File
	                options := os.O_CREATE | os.O_APPEND
	                log.Println("log0")
	                file, err = os.OpenFile("read1.txt", options, os.FileMode(0644))
	                log.Println("log1")
	                if err != nil {
	                                log.Println("log2")
	
	                                w.Write([]byte("Call Failed"))
	                                log.Println("log3")
	
	                                return
	
	                                log.Println("log4")
	                                log.Fatal(err)
	                }
	
	                _, err = file.Write([]byte(string(body)))
	                if err != nil {
	                                w.Write([]byte("Call Failed"))
	                                return
	                                log.Fatal(err)
	                }
	
	                _, err = file.Write([]byte("\n"))
	                if err != nil {
	                                w.Write([]byte("Call Failed"))
	                                return
	                                log.Fatal(err)
	                }
	
	                err = file.Close()
	                if err != nil {
	                                w.Write([]byte("Call Failed"))
	                                return
	                                log.Fatal(err)
	                }
	
	                log.Println("Call Successful")
	
	                w.Write([]byte("Successful"))
	}
	var db *sql.DB
	
	                var server = "10.6.202.36"
	                var port = 1433
	                var user = "web"
	                var password = "cSQmCHVdjlHGWAfl"
	                var database = "itrade"
	
	func main() {
	                // The database is called testDb
	//    db, err := sql.Open("mysql", "web:cSQmCHVdjlHGWAfl@tcp(10.6.202.36:1433)/itrade")
	                
	                connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	                server, user, password, port, database)
	
	                var err error
	     // Create connection pool
	                db, err = sql.Open("sqlserver", connString)
	                if err != nil {
	                                log.Fatal("Error creating connection pool: ", err.Error())
	                }
	                ctx := context.Background()
	                err = db.PingContext(ctx)
	                if err != nil {
	                                log.Fatal(err.Error())
	                }
	                fmt.Printf("Connected!\n")
	
	
	    // defer the close till after the main function has finished
	    // executing
	    defer db.Close()
	                mux := http.NewServeMux()
	                mux.HandleFunc("/", home)
	                mux.HandleFunc("/ksrp", Ksrp)
	
	                log.Println("Starting server on :8080")
	                err = http.ListenAndServe("localhost:8080", mux)
	
	                log.Fatal(err)
	                var w1 http.ResponseWriter
	                w1.Write([]byte("Call Failed"))
	                return
	                log.Fatal(err)
	
	}
	/*
	DEFAULT_SERVER_1 10.6.202.36
	DEFAULT_DB_1 itrade
	DEFAULT_USER_1 web
	DEFAULT_PASS_1 cSQmCHVdjlHGWAfl
	DEFAULT_PORT_1 1433
	*/

