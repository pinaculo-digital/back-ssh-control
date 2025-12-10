package util

import (
	randCrypt "crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	app "go_service/core/util/error"
	"go_service/core/util/executor"
	"math/rand/v2"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	_ "database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

func TypeToMap[T any](dataType T) (map[string]interface{}, error) {
	data, err := json.Marshal(dataType)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	return m, err
}

func MapToType[T any](dataMap map[string]interface{}, dataConvert *T) {
	jsonData, _ := json.Marshal(dataMap)
	json.Unmarshal(jsonData, &dataConvert)
}

func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
func NormalizeString(input string) (string, error) {
	return url.QueryUnescape(input)
}

func GenerateRandomBytes() (string, error) {
	key := make([]byte, 64) // 256 bits
	_, err := randCrypt.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}
func GenerateRandomString(max int, dic string) string {
	if max <= 0 || len(dic) == 0 {
		return ""
	}

	length := len(dic)
	var strB = strings.Builder{}

	for i := 0; i < max; i++ {
		randomIndex := rand.IntN(length)
		symbol := dic[randomIndex]
		strB.WriteByte(symbol)
	}
	return strB.String()
}

func StringToDate(input string) time.Time {
	parts := strings.Split(input, ".")
	if len(parts) < 2 {
		return time.Now()
	}

	datePart := strings.TrimSpace(parts[0][12:])

	publishDate, err := time.Parse("02/01/2006", datePart)
	if err != nil {
		return time.Now()
	}
	return publishDate
}

func NormalizeMimeType(mime string) string {
	// Remove leading dot if present
	mime = strings.TrimPrefix(mime, ".")

	// Map of common file extensions to MIME types
	mimeMap := map[string]string{
		"mp4":  "video/mp4",
		"webp": "image/webp",
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"gif":  "image/gif",
		"mp3":  "audio/mpeg",
		"wav":  "audio/wav",
		"pdf":  "application/pdf",
		"txt":  "text/plain",
		"html": "text/html",
		"css":  "text/css",
		"js":   "application/javascript",
	}

	// Convert to lowercase for case-insensitive matching
	mime = strings.ToLower(mime)

	// Look up MIME type in map, return as-is if not found
	if normalized, exists := mimeMap[mime]; exists {
		return normalized
	}
	return mime
}

// Pega uma coleção, verifica se dado item já existe, caso não exista ele da um append
func Adder[T any](collection []T, info T, iteratee func(item T) bool) (dataAppend []T, index int) {

	dataAppend = collection
	_, index, exist := lo.FindIndexOf(collection, iteratee)

	if !exist {
		dataAppend = append(dataAppend, info)
		return dataAppend, len(dataAppend) - 1
	}

	return dataAppend, index

}

func CountSQL(executor executor.Executor, expression string, builder *goqu.SelectDataset) (total int, err error) {
	countQuery := builder.
		ClearSelect().
		ClearOrder().
		ClearOffset().
		ClearLimit()

	countQuery = countQuery.Select(goqu.COUNT(goqu.L(expression))).GroupBy()

	sql, args, err := countQuery.ToSQL()

	if err != nil {
		return 0, fmt.Errorf("erro ao gerar SQL: %w", err)
	}

	rows, err := executor.Query(sql, args...)

	if err != nil {
		return 0, fmt.Errorf("erro ao executar COUNT: %w", err)
	}

	for rows.Next() {
		var tmp int
		rows.Scan(&tmp)
		total += tmp
	}

	return total, nil
}

func ExtractFormFile(ctx *gin.Context) (file multipart.File, fileHeader *multipart.FileHeader, err error) {
	form, err := ctx.MultipartForm()
	fmt.Println(err)
	if err != nil {
		return file, fileHeader, app.InternalServerError(err.Error())
	}

	files := form.File["file"]
	if len(files) == 0 {
		return file, fileHeader, app.BadRequest("No file provided")
	}

	fileHeader = files[0]

	file, err = fileHeader.Open()
	if err != nil {
		return file, fileHeader, app.BadRequest(err.Error())
	}
	return file, fileHeader, err
}
