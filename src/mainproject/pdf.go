 package main
 
func pdfHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "pdf!")
} 
