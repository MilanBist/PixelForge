package database

import (
	"context"
	"fmt"
)

// get the image's storageKey based on the imageId and userId
func(p *PostgresData) GetImageStorageKey(imageId, userId int) (string, error){
	var storageKey string
	query := `SELECT "storageKey" FROM "images" WHERE "userId" = $1 AND "id" = $2`
	err := p.Db.QueryRow(
		context.Background(),
		query,
		userId,
		imageId,
	).Scan(&storageKey)

	if err != nil{
		fmt.Println("Package Database. Error: ", err)
		return "", err
	}
	return storageKey, nil
}