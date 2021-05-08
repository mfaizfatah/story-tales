package usecases

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func UploadToFtpProccess(userid int, story *models.Story, file multipart.File, fileHeader *multipart.FileHeader) {
	const SSH_ADDRESS = "178.128.53.127:22"
	const SSH_USERNAME = "sftpdigi"
	const SSH_KEY = ""
	const SSH_PASSWORD = "v8McfYATv2LUZqB9"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
			PublicKeyFile(SSH_KEY),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}
	defer client.Close()

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed create client sftp client. " + err.Error())
	}
	defer sftpClient.Close()

	// /data/sftpdigi/upload/<idUser>/<title>/filename
	dirpath := filepath.Join(strconv.Itoa(userid))
	log.Print(dirpath)
	_, err = sftpClient.Lstat(dirpath)
	if err != nil {
		sftpClient.MkdirAll(dirpath)
	}
	fileLocation := fmt.Sprintf("%v/%v_%v_%v", userid, story.Title, story.Season, fileHeader.Filename)
	// fileLocation := filepath.Join(strconv.Itoa(userid), fileHeader.Filename)
	log.Print(fileLocation)

	fDestination, err := sftpClient.Create(fileLocation)
	if err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}
	defer fDestination.Close()

	if _, err := io.Copy(fDestination, file); err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}

}
