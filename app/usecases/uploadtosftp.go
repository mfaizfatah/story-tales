package usecases

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
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

func (r *uc) UploadToFtpProccess(ctx context.Context, userid int, path string, file multipart.File, fileHeader *multipart.FileHeader) (context.Context, string, error) {
	var (
		dir           string
		baseUriImages = os.Getenv("IMAGES_URI")
	)

	sshConfig := &ssh.ClientConfig{
		User:            os.Getenv("SFTP_USERNAME"),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("SFTP_PASSWORD")),
		},
	}

	client, err := ssh.Dial("tcp", os.Getenv("SFTP_ADDRESS"), sshConfig)
	if err != nil {
		return ctx, dir, err
	}
	defer client.Close()

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return ctx, dir, err
	}
	defer sftpClient.Close()

	// /data/sftpdigi/upload/<idUser>/<title>/filename
	dirpath := fmt.Sprintf("%v/%v", userid, path)
	_, err = sftpClient.Lstat(dirpath)
	if err != nil {
		sftpClient.MkdirAll(dirpath)
	}
	fileLocation := fmt.Sprintf("%v/%v/%v", userid, path, fileHeader.Filename)
	ctx = logger.Logf(ctx, "file location() => %v", fileLocation)

	fDestination, err := sftpClient.Create(fileLocation)
	if err != nil {
		return ctx, dir, err
	}
	defer fDestination.Close()

	if _, err := io.Copy(fDestination, file); err != nil {
		return ctx, dir, err
	}

	res := fmt.Sprintf("%v/%v", baseUriImages, fileLocation)
	return ctx, res, nil
}
