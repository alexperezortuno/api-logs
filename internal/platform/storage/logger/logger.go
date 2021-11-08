package logger

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

func UnmarshalInputRequest(data []byte) (InputRequest, error) {
	var r InputRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InputRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InputRequest struct {
	Message    *string       `json:"message,omitempty"`
	Additional []interface{} `json:"additional,omitempty"`
	Level      *int32        `json:"level,omitempty"`
	//Level      *string       `json:"level,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	FileName  *string `json:"fileName,omitempty"`
	//LineNumber *string       `json:"lineNumber,omitempty"`
	LineNumber *int32 `json:"lineNumber,omitempty"`
}

func LogRequest(p InputRequest) {
	logPath := os.Getenv("APP_LOG_PATH")
	var f *os.File

	if logPath == "" {
		logPath = "/tmp"
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err := os.MkdirAll(logPath, 0700)

		if err != nil {
			return
		}
	}

	if _, err := os.Stat(fmt.Sprintf("%s/apilog.log", logPath)); err == nil {
		log.Println("Ok!, log file exist")
	} else if errors.Is(err, os.ErrNotExist) {
		log.Printf("log file not exist: %v", err)
		f, err := os.Create(fmt.Sprintf("%s/apilog.log", logPath))

		if err != nil {
			log.Fatalf("error creating file: %v", err)
		}

		defer f.Close()
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/apilog.log", logPath), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)
	logData, err := p.Marshal()

	if err != nil {
		log.Fatalf("error unmarshal data: %v", err)
	}

	log.Println(string(logData))
}
