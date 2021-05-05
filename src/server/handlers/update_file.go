package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"gitlab.computing.dcu.ie/collint9/2021-ca400-collint9-coynemt2/src/database"
	"gitlab.computing.dcu.ie/collint9/2021-ca400-collint9-coynemt2/src/database/repo"
	"gitlab.computing.dcu.ie/collint9/2021-ca400-collint9-coynemt2/src/database/schema"
	"gitlab.computing.dcu.ie/collint9/2021-ca400-collint9-coynemt2/src/server/models"
	"gitlab.computing.dcu.ie/collint9/2021-ca400-collint9-coynemt2/src/server/restapi/operations/files"
	"path/filepath"
)

func ConvertToFileModel(file *schema.File) *models.File {
	return &models.File{
		ID:          uint64(file.ID),
		Name:        file.Name,
		Size:        file.Size,
		Hash:        file.Hash,
		DirectoryID: uint64(file.DirectoryID),
		Available:   file.Available,
	}
}

func getFile(user *schema.User, params files.UpdateFileParams) (*schema.File, error) {
	var file *schema.File
	//get file by file ID
	file, err := repo.GetFileByID(uint(params.FileID))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func updateFileJob(file *schema.File, directory *schema.Directory) (*schema.File, error) {
	file, err := repo.UpdateFileByDirectory(file, directory)
	if err != nil {
		return file, err
	}

	return file, nil
}

func UpdateFile(params files.UpdateFileParams, user *schema.User) middleware.Responder {
	var (
		err                      error
		newFileName              string
		directory                *schema.Directory
		errNotFound              = files.NewUpdateFileDefault(404).WithPayload("file not found")
		errFileDirNotFound       = files.NewUpdateFileDefault(404).WithPayload("file directory not found")
		errRenameFailed          = files.NewUpdateFileDefault(500).WithPayload("failed to rename the file")
		errUpdateDirectoryFailed = files.NewUpdateFileDefault(500).WithPayload("failed to update the file's directory")
	)

	file, err := getFile(user, params)
	if err != nil {
		return errNotFound
	}

	// get directory the file is being moved to by path
	if params.Filepath != nil {
		newFileName = filepath.Base(*params.Filepath)
		homeDir, err := repo.GetHomeDir(user.ID)
		if err != nil {
			return errUpdateDirectoryFailed
		}
		directory, err = repo.GetDirByPath(filepath.Clean(filepath.Join(homeDir.Path, filepath.Dir(*params.Filepath))))
		if err != nil {
			return errNotFound
		}
	}

	// get directory the file is being moved to by ID
	if params.DirectoryID != nil {
		directory, err = repo.GetDirectoryByID(uint(*params.DirectoryID))
		if err != nil {
			return errFileDirNotFound
		}
	}

	// find file
	if file.Directory == nil {
		if err = database.DB.Preload("Directory").Find(file).Error; err != nil {
			return errFileDirNotFound
		}
		if file.Directory == nil {
			return errFileDirNotFound
		}
	}

	newFile, err := updateFileJob(file, directory)
	if err != nil {
		return errFileDirNotFound
	}

	// update filename if needed, both on disk and in DB
	if params.Filepath != nil && newFileName != file.Name {
		newFile, err = repo.UpdateFilename(file, newFileName)
		if err != nil {
			return errRenameFailed
		}
	}

	return files.NewUpdateFileOK().WithPayload(ConvertToFileModel(newFile))
}
