package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "myapp/handlers"
    _ "github.com/lib/pq"
)

func main() {
    // SSL接続用に修正: sslmode=require
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"))

    var err error
    for i := 0; i < 30; i++ {
        // データベース接続を試行
        handlers.DB, err = sql.Open("postgres", dsn)
        if err == nil {
            // データベース接続確認
            err = handlers.DB.Ping()
            if err == nil {
                break
            }
        }
        log.Printf("データベース接続を試行中... (%d/30)", i+1)
        time.Sleep(2 * time.Second)
    }
    if err != nil {
        log.Fatalf("データベース接続に失敗しました: %v", err)
    }
    defer handlers.DB.Close()

    log.Println("サーバーがポート8080で起動しています...")
    http.HandleFunc("/", handlers.ShowRecords)
    http.HandleFunc("/add", handlers.AddRecord)
    http.HandleFunc("/download", handlers.DownloadCSV) // 新しいルートを追加
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
