package handler

import (
	"encoding/json"
	"github.com/bootcamp-go/live-codings/go-web/functional-tests/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

func Test_SaveProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
        "nombre": "Tester","tipo": "Funcional","cantidad": 10,"precio": 99.99
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
