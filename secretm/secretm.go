package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/petaluser/awsgo"
	"github.com/petaluser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJSON, error) {
	var datosSecret models.SecretRDSJSON
	fmt.Println(" > Pido Secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println("Error al obtener el secreto" + err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Secreto obtenido " + datosSecret.Username)
	return datosSecret, nil
}
