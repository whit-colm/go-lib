package dat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	Log            *log.Logger
	path           = filepath.FromSlash("./")
	ErrInvalidPath error
)

func NewLog(createDir bool) error {
	currentTime := time.Now().Format(time.RFC3339)
	file, err := os.Create(filepath.Join(path, "/dat/logs@") + currentTime + ".log")
	if err != nil {
		return err
	}
	Log = log.New(file, "", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)
	return nil
}

// Sets the path variable, multi-OS friendly.
func SetPath(p string) {
	path = filepath.FromSlash(p)
}

// Simple check to make sure a... something... exists.
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

var lock sync.Mutex

/* Saves a file to json
* This is a simple helper function to manage data persistance and keep it in
* one place.
*
* Parameters
* - fileName (string) : where the data is to be stored. Separated by /, as it is
*			filtered for the OS later.
* - v (interface{})   : the thing to save saving.
*
* Returns:
* - error : if an error occured in the program.
 */
func Save(fileName string, item interface{}, createDir, createFile bool) error {
	lock.Lock()
	defer lock.Unlock()
	// Checks for directories in fileName, creates them if they aren't real.
	validPath, err := exists(path)
	if err != nil {
		return err
	}
	if !validPath && createDir {
		if err := os.MkdirAll(path, 0775); err != nil {
			return err
		}
	} else if !validPath && !createDir {
		ErrInvalidPath = errors.New("Invalid path " + fileName)
		return ErrInvalidPath
	}
	// Opens file, creates if not real
	file, err := os.OpenFile(filepath.Join(path+"cfg/")+fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.MarshalIndent(item, "", "\t")
	if err != nil {
		return err
	}
	reader := bytes.NewReader(b)
	_, err = io.Copy(file, reader)
	return err
}

/* Loads a file from json
* This is a simple helper function to manage data persistance and keep it in
* one place.
*
* Parameters
* - fileName (string) : where the data is stored. Note that this should
*			be in the form of $MODULENAME/*.json so modules that
*			use the same name config.json don't interfere.
* - v (interface{})   : the thing you're loading. as a pointer
*
* Returns:
* - error : error that occured during execution.
 */
func Load(fileName string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	file, err := os.OpenFile(filepath.Join(path, "cfg/")+fileName, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(v)
	return err
}
