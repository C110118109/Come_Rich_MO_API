package build_file

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

//專屬FOR excel pdf io.Reader版本的s3 上傳
func NewS3Client(ctx context.Context) *s3.Client {
	cfg, err := config.LoadDefaultConfig(ctx, config.
		WithCredentialsProvider(credentials.
			NewStaticCredentialsProvider("AKIASN3ZFKGX6PNNAXBN", "iBrN+9EzNTeWMO1sBq7TK5vPO8DmT9/89QZ36d85", "")),
			config.WithRegion("ap-northeast-1"))

	if err != nil {
		panic(err)
	}
	return s3.NewFromConfig(cfg) // Create an Amazon S3 service client
}

func CreateGetObjectInput(bucket string, key string) *s3.GetObjectInput {
    return &s3.GetObjectInput{
        Bucket: &bucket,
        Key:    &key,
    }
}

func GetObjectContent(ctx context.Context, client *s3.Client, input *s3.GetObjectInput) io.Reader {
    output, err := client.GetObject(ctx, input)
    if err != nil {
        panic(err)
    }

    // b, err := ioutil.ReadAll(output.Body) // Body is Reader
    // if err != nil {
    //     panic(err)
	// }
	//ioutil.WriteFile(filename, b, 0644)
    
	return output.Body
}


func CreateInput(bucket string, key string, data io.Reader) *s3.PutObjectInput {
	return &s3.PutObjectInput{
        Bucket: &bucket,
        Key:    &key,
        Body:   data,
    }
}
