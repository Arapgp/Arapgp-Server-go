package sfs

import (
	"errors"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// GetContentByPath is to get content of a file
func GetContentByPath(path, name string) (content string, err error) {
	if !checkFileExistence(path, name) {
		log.WithFields(log.Fields{"path": path, "name": name}).Warningln("File do not exist!")
		return "", errors.New("arapgp.pkg.sfs.GetContentByPath => checkFileExistenc: file do not exist")
	}

	fd, err := os.OpenFile(path+name, os.O_RDONLY, 0444)
	if err != nil {
		errmsg := "arapgp.pkg.sfs.GetContentByPath => File open failed!"
		log.Warningln(errmsg)
		return "", errors.New(errmsg + err.Error())
	}
	defer fd.Close()

	buf, err := ioutil.ReadAll(fd)
	if err != nil {
		errmsg := "arapgp.pkg.sfs.GetContentByPath => File read failed!"
		log.Warningln(errmsg)
		return "", errors.New(errmsg + err.Error())
	}
	return string(buf), nil
}

// WriteContentByPath is to write content to a file
// if that file do not exist, create it
func WriteContentByPath(path, name, content string) (err error) {
	var fd *os.File
	// if-else generate/write file
	if !checkFileExistence(path, name) {
		// create folder
		err = os.MkdirAll(path, os.FileMode(os.O_RDWR))
		if err != nil {
			errmsg := "arapgp.pkg.sfs.GetContentByPath => MkdirAll: create folder recursively failed;"
			log.WithFields(log.Fields{"path": path, "name": name, "err": err.Error()}).Warningln("Open File failed!")
			return errors.New(errmsg + err.Error())
		}

		// create file
		fd, err = os.OpenFile(path+name, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			errmsg := "arapgp.pkg.sfs.GetContentByPath => OpenFile: file created failed;"
			log.WithFields(log.Fields{"path": path, "name": name, "err": err.Error()}).Warningln("Create File failed!")
			return errors.New(errmsg + err.Error())
		}
	} else {
		fd, err = os.OpenFile(path+name, os.O_WRONLY, 0666)
		if err != nil {
			errmsg := "arapgp.pkg.sfs.GetContentByPath => OpenFile: open file failed;"
			log.WithFields(log.Fields{"path": path, "name": name, "err": err.Error()}).Warningln("Open File failed!")
			return errors.New(errmsg + err.Error())
		}
	}
	defer fd.Close()

	cnt, err := fd.WriteString(content)
	if err != nil {
		errmsg := "arapgp.pkg.sfs.GetContentByPath => WriteString: write content failed;"
		log.WithFields(log.Fields{"path+name": path + name, "fd": fd, "cnt": cnt, "err": err.Error()}).Warningln("Write content failed")
		return errors.New(errmsg + err.Error())
	}
	return nil
}

// DeleteFileByPath is to delete file in fs
func DeleteFileByPath(path, name string) (err error) {
	err = os.Remove(path + name)
	if err != nil {
		errmsg := "arapgp.pkg.sfs.GetContentByPath => DeleteFileByPath: delete file failed;"
		log.WithFields(log.Fields{"path+name": path + name, "err": err.Error()}).Warningln("delete file failed")
		return errors.New(errmsg + err.Error())
	}
	return nil
}

// return true  => exist
// return false => do not exist
func checkFileExistence(path, name string) bool {
	_, err := os.Stat(path + name)
	return !os.IsNotExist(err)
}
