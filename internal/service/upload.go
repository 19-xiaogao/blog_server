package service

import (
	"github.com/pkg/errors"
	"mime/multipart"
	"os"
	"xiaolong_blog/global"
	"xiaolong_blog/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(filetype upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {

	fileName := upload.GetFileName(fileHeader.Filename)

	if !upload.CheckContainExt(filetype, fileName) {
		return nil, errors.New("file suffix is not supported.")
	}

	if !upload.CheckMaxSize(filetype, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}
	dst := uploadSavePath + "/" + fileName
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
