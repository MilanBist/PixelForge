package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/image-generator/internal/models"
)

func(p *PostgresData) AddUploadedFiles(uploadedMetaData models.UploadedFilesMetaData)(int, error){
	
	var id int64

	err := p.Db.QueryRow(
		context.Background(),
		`INSERT INTO "uploadedFiles" ("userId", "fileName", "storageKey", "fileType", "mimeType", "fileSize") 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id"`,
		uploadedMetaData.UserId,
		uploadedMetaData.Filename,
		uploadedMetaData.StorageKey,
		uploadedMetaData.FileType,
		uploadedMetaData.Mimetype,
		uploadedMetaData.FileSize,
	).Scan(&id)

	if err != nil{
		fmt.Println("Package Database. Error: ", err)
		return -1, err
	}
		return int(id), nil
}

func(p *PostgresData) AddUploadedImageFiles(uploadedMetaData models.UploadedFilesMetaData)(int, error){
	
	var id int64

	err := p.Db.QueryRow(
		context.Background(),
		`INSERT INTO "uploadedFiles" ("userId", "fileName", "storageKey", "fileType", "mimeType", "fileSize", "height", "width") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING "id"`,
		uploadedMetaData.UserId,
		uploadedMetaData.Filename,
		uploadedMetaData.StorageKey,
		uploadedMetaData.FileType,
		uploadedMetaData.Mimetype,
		uploadedMetaData.FileSize,
		uploadedMetaData.Height,
		uploadedMetaData.Width,
	).Scan(&id)

	if err != nil{
		fmt.Println("Package Database. Error: ", err)
		return -1, err
	}
		return int(id), nil
}

	

func(p *PostgresData) AddGeneratedFiles(generatedFilesMetaData models.GeneratedImageMetaData)(int64, error){
	var id int64
	err := p.Db.QueryRow(
		context.Background(),
		`INSERT INTO "images" ("userId", "sourceFileId", "storageKey", "fileName", "mimeType", "width", "height", "fileSize") 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING "id"`,
		generatedFilesMetaData.Userid,
		generatedFilesMetaData.SourceFieldId,
		generatedFilesMetaData.StorageKey,
		generatedFilesMetaData.Filename,
		generatedFilesMetaData.Mimetype,
		generatedFilesMetaData.Width,
		generatedFilesMetaData.Height,
		generatedFilesMetaData.FileSize,
	).Scan(&id)
	if err != nil{
		fmt.Println("Package Database. Error: ", err)
		return id, err
	}
	return id, nil
}
	
func (p *PostgresData) AddToHistoryOfUser(historyData models.History) (error){
	var id int
	query := `INSERT INTO "history"("userId", "sourceFileId", "outputImageId", "operationType", "parameters", "status")
			VALUES($1, $2, $3, $4, $5, $6) RETURNING "id"
	`
	err := p.Db.QueryRow(
		context.Background(),
		query, 
		historyData.UserID, 
		historyData.SourceFileID, 
		historyData.OutputImageID,
		historyData.OperationType,
		historyData.Parameters,
		historyData.Status,
	).Scan(&id)

	fmt.Println("The error is: ", err)
	if err != nil{
		fmt.Println(err)
		return errors.New("Error in adding to database.")
	}
	return nil
}

func(p *PostgresData) GetAllUploadedFiles(userId int)([]models.BaseFileData, error){
	query := `SELECT "id", "fileName","fileType", "createdAt"
			FROM "uploadedFiles" WHERE "userId" = $1
	`

	ctx := context.Background()

	rows, err := p.Db.Query(ctx, query, userId)
	

	if err != nil{
		fmt.Println("Error in getting the uploadedFiles: ", err)
		return []models.BaseFileData{}, errors.New("error in getting data")
	}

	defer rows.Close()

	var uploadedFiles []models.BaseFileData

	for rows.Next(){
		var uploadedModel models.BaseFileData
		err := rows.Scan(&uploadedModel.Id, 
		&uploadedModel.Filename,
		&uploadedModel.FileType,
		&uploadedModel.CreatedAt)

		if err != nil{
			fmt.Println(err)
			fmt.Println("Error in gettting metadata from database.")
			return []models.BaseFileData{}, errors.New("Error in getting the value of rows.")
		}


		uploadedFiles = append(uploadedFiles, uploadedModel)
	}


	return uploadedFiles, nil

}


func (p *PostgresData) GetHistoryDataOfUser(userId int) ([]models.HistoricalData, error){
	query := `
		SELECT
		"uploadedFiles"."id",
		"uploadedFiles"."fileName",
		"uploadedFiles"."createdAt",
		"images"."id",
		"images"."fileName",
		"images"."mimeType",
		"images"."height",
		"images"."width",
		"images"."fileSize"
		FROM "uploadedFiles"
		JOIN "images"
			ON "images"."sourceFileId" = "uploadedFiles"."id"
		WHERE "uploadedFiles"."userId" = $1
		ORDER BY
			"uploadedFiles"."createdAt" DESC,
			"images"."id" ASC;
	`

	ctx := context.Background()
	var data []models.HistoricalData
	var counter int
	rows, err := p.Db.Query(ctx, query, userId)
	if err != nil{
		fmt.Println("Error is: ",err)
		return nil, errors.New("Error in generating response.")
	}
	defer rows.Close()

	for rows.Next(){
		var d models.HistoricalData
		err := rows.Scan(&d.UploadedFileId,
		&d.UploadedFileName,
		&d.CreatedAt,
		&d.ImageId,
		&d.ImageName,
		&d.MimeType,
		&d.Width,
		&d.Height,
		&d.FileSize,
	)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}

	fmt.Println(d)
	fmt.Println("counter: ",counter)
	counter += 1
	data = append(data, d)
	}

	fmt.Println(data)
	return data, nil
}