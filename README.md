# Cloud Function Mytodolist
Repositori ini adalah Cloud Function dari aplikasi Mytodolist yang ditulis dengan bahasa Go.

## Import
1. Import Standard Library
    - `fmt` adalah package yang digunakan untuk mencetak teks.
        Contoh penggunaannya:
        ```go
        fmt.Fprintf(w, todo.GCFHandlerDeleteTodo("PASETOPUBLICKEY", "MONGOSTRING", "mytodolist", "todo", r))
        ```
    - `net/http` adalah package yang menyediakan dukungan untuk membangun layanan web (HTTP) dan mengirim permintaan HTTP.
        Contoh penggunaannya:
        ```go
        func MytodolistDeleteTodo(w http.ResponseWriter, r *http.Request) {}
        ```

2. Import External Library
    - `github.com/GoogleCloudPlatform/functions-framework-go/functions` adalah framework go yang memungkinkan Anda membuat fungsi-fungsi dan menjalankannya di lingkungan lokal atau di Google Cloud.
        Contoh penggunaannya:
        ```go
        func init() {
	        functions.HTTP("MytodolistDeleteTodo", MytodolistDeleteTodo)
        }
        ```
    

3. Import Backend Module for this Repository
	- `github.com/mytodolist1/be_p3/modul` adalah modul yang digunakan untuk memanggil handler.
        Contoh penggunaannya:
        ```go
        fmt.Fprintf(w, todo.GCFHandlerDeleteTodo("PASETOPUBLICKEY", "MONGOSTRING", "mytodolist", "todo", r))
        ```

4. Import External Library for Webhook
    - `github.com/gofiber/fiber/v2/middleware/adaptor` adalah middleware di Fiber yang menyediakan antarmuka untuk menyesuaikan atau menyimpan data yang terkait dengan permintaan HTTP dan respons HTTP.
        Contoh penggunaannya:
        ```go
        func init() {
	        functions.HTTP("WebHook", adaptor.FiberHandlerFunc(webhook.PostMessage))
        }
        ```
    - `github.com/whatsauth/webhook` adalah modul webhook yang menggunakan method HTTP POST dengan Header bernama Secret.
        Contoh penggunaannya:
        ```go
        functions.HTTP("WebHook", adaptor.FiberHandlerFunc(webhook.PostMessage))
        ```